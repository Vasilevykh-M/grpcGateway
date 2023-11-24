package main

import (
	"awesomeProject/config"
	"awesomeProject/internal/db"
	"awesomeProject/internal/grpc_serv"
	"awesomeProject/internal/repository/postgres"
	server_2 "awesomeProject/internal/server"
	"awesomeProject/pkg/articles"
	"awesomeProject/pkg/logger"
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	_ "github.com/opentracing/opentracing-go"
	config_2 "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	flag.Parse()

	var adrGrpc string
	var adrGateway string

	flag.StringVar(&adrGrpc, "adrGrpc", ":50051", "Add for messages server")
	flag.StringVar(&adrGateway, "adrGateway", ":9001", "Add for messages server")

	go runRest(adrGateway)

	if err := run(ctx, adrGrpc); err != nil {
		log.Fatal(err)
	}
}

func runRest(addr string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := articles.RegisterArticlesHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 9001")
	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}
}

func run(ctx context.Context, addr string) error {

	cfg := config_2.Configuration{
		Sampler: &config_2.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config_2.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := cfg.New(
		"articles-service",
	)
	if err != nil {
		fmt.Printf("cannot create tracer: %v\n", err)
		log.Fatal(err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	conStrDB, err := config.ReadConfDBConn("config/dbConf.json")
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewDB(ctx, conStrDB)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()

	articleRepo := postgres.NewArticles(database)

	zapLogger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	logger.SetGlobal(
		zapLogger.With(zap.String("component", "gateway")),
	)

	server := grpc.NewServer()

	articles.RegisterArticlesServer(server, &grpc_serv.Impl{Server: &server_2.Server{Repo: articleRepo}})

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logger.Infof(ctx, "service gateway listening on %q", addr)
	log.Printf("service messages listening on %q", addr)
	return server.Serve(lis)
}
