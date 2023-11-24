package grpc_serv

import (
	"awesomeProject/internal/server"
	"awesomeProject/pkg/articles"
	"awesomeProject/pkg/logger"
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

type Impl struct {
	articles.UnimplementedArticlesServer
	Server *server.Server
}

func (i *Impl) GetArticleByID(ctx context.Context, Id *articles.Id) (*articles.JoinArticlePost, error) {

	l := logger.FromContext(ctx)

	ctx = logger.ToContext(ctx, l.With(zap.String("method", "GetArticleByID")))

	span, ctx := opentracing.StartSpanFromContext(ctx, "app: get_by_id")
	defer span.Finish()

	codeReq, artReq := i.Server.Get(ctx, Id.Id)

	if codeReq != http.StatusOK {
		err := fmt.Errorf("errots is %d", http.StatusNotFound)
		logger.Errorf(ctx, "id", Id.Id, "err", err)
		return nil, err
	}

	if artReq == nil {
		err := fmt.Errorf("errots is %d", http.StatusNotFound)
		logger.Errorf(ctx, "id", Id.Id, "err", err)
		return nil, err
	}

	if len(artReq) == 0 {
		err := fmt.Errorf("errots is %d", http.StatusNotFound)
		logger.Errorf(ctx, "id", Id.Id, "err", err)
		return nil, err
	}

	article := articles.JoinArticlePost{}

	posts := make([]*articles.Post, 0)

	for _, art := range artReq {
		if art.IDPost == nil {
			break
		}
		posts = append(posts, &articles.Post{
			Id:       *art.IDPost,
			IdAuthor: Id.Id,
			Name:     *art.NamePost,
			Sales:    *art.Sales,
		})
	}

	article.Article = &articles.Article{
		Id:     Id.Id,
		Name:   artReq[0].NameArticle,
		Rating: artReq[0].Rating,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: artReq[0].CreatedAt.Unix(),
			Nanos:   int32(artReq[0].CreatedAt.Nanosecond()),
		},
	}
	article.Post = posts

	logger.Infof(ctx, "GetArticleByID success", "id", Id.Id)

	return &article, nil
}

func (i *Impl) CreateArticle(ctx context.Context, Article *articles.Article) (*articles.Article, error) {

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "CreateArticle")))

	span, ctx := opentracing.StartSpanFromContext(ctx, "app: create_article")
	defer span.Finish()

	dbArticle := server.ArticleRequest{
		ID:        Article.Id,
		Name:      Article.Name,
		Rating:    Article.Rating,
		CreatedAt: Article.CreatedAt.AsTime(),
	}

	codeReq, artReq := i.Server.CreateArticle(ctx, &dbArticle)

	if codeReq != http.StatusOK {
		err := fmt.Errorf("errots is %d", http.StatusNotFound)
		logger.Errorf(ctx, "name", Article.Name, "rating", Article.Rating, "err", err)
		return nil, err
	}

	result := articles.Article{
		Id:     artReq.ID,
		Name:   artReq.Name,
		Rating: artReq.Rating,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: artReq.CreatedAt.Unix(),
			Nanos:   int32(artReq.CreatedAt.Nanosecond()),
		},
	}

	logger.Infof(ctx, "CreateArticle success", "name", Article.Name, "rating", Article.Rating)

	return &result, nil
}

func (i *Impl) DeleteArticle(ctx context.Context, Id *articles.Id) (*articles.Id, error) {
	codeReq := i.Server.DeleteArticle(ctx, Id.Id)

	l := logger.FromContext(ctx)

	ctx = logger.ToContext(ctx, l.With(zap.String("method", "DeleteArticle")))

	span, ctx := opentracing.StartSpanFromContext(ctx, "app: delete_article")
	defer span.Finish()

	if codeReq != http.StatusOK {
		err := fmt.Errorf("errots is %d", http.StatusNotFound)
		logger.Errorf(ctx, "id", Id.Id, "err", err)
		return nil, err
	}

	logger.Infof(ctx, "DeleteArticle success", "id", Id.Id)

	return Id, nil
}

func (i *Impl) CreatePost(ctx context.Context, Post *articles.Post) (*articles.Post, error) {

	l := logger.FromContext(ctx)

	ctx = logger.ToContext(ctx, l.With(zap.String("method", "CreatePost")))

	span, ctx := opentracing.StartSpanFromContext(ctx, "app: create_post")
	defer span.Finish()

	dbPost := server.PostRequest{
		ID:       Post.Id,
		IdAuthor: Post.IdAuthor,
		Name:     Post.Name,
		Sales:    Post.Sales,
	}

	codeReq, artReq := i.Server.CreatePost(ctx, dbPost.IdAuthor, &dbPost)

	if codeReq == http.StatusBadRequest {
		err := fmt.Errorf("errots is %d", http.StatusNotFound)
		logger.Errorf(ctx, "name", Post.Name, "sales", Post.Sales, "err", err)
		return nil, err
	}

	Post.Id = artReq.ID

	logger.Infof(ctx, "CreatePost success", "name", Post.Name, "sales", Post.Sales)

	return Post, nil
}

func (i *Impl) UpdateArticle(ctx context.Context, Article *articles.Article) (*articles.Article, error) {

	l := logger.FromContext(ctx)

	ctx = logger.ToContext(ctx, l.With(zap.String("method", "UpdateArticle")))

	span, ctx := opentracing.StartSpanFromContext(ctx, "app: update_post")
	defer span.Finish()

	dbArticle := server.ArticleRequest{
		ID:        Article.Id,
		Name:      Article.Name,
		Rating:    Article.Rating,
		CreatedAt: Article.CreatedAt.AsTime(),
	}

	codeReq, artReq := i.Server.UpdateArticle(ctx, &dbArticle)

	if codeReq != http.StatusOK {
		err := fmt.Errorf("errots is %d", http.StatusNotFound)
		logger.Errorf(ctx, "name", Article.Name, "rating", Article.Rating, "err", err)
		return nil, err
	}

	result := articles.Article{
		Id:     artReq.ID,
		Name:   artReq.Name,
		Rating: artReq.Rating,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: artReq.CreatedAt.Unix(),
			Nanos:   int32(artReq.CreatedAt.Nanosecond()),
		},
	}

	logger.Infof(ctx, "UpdateArticle success", "name", Article.Name, "rating", Article.Rating)

	return &result, nil
}
