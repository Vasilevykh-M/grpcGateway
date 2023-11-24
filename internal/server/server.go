package server

import (
	"awesomeProject/internal/repository"
	"context"
)

type IServer interface {
	CreateArticle(ctx context.Context, unm *ArticleRequest) (int, *ArticleRequest)
	CreatePost(ctx context.Context, id int64, unm *PostRequest) (int, *PostRequest)
	DeleteArticle(ctx context.Context, id int64) int
	Get(ctx context.Context, id int64) (int, []*repository.JoinArticlePost)
	UpdateArticle(ctx context.Context, unm *ArticleRequest) (int, *ArticleRequest)
}

type Server struct {
	Repo repository.IArticleRepo
}

func NewServer(articleRepo repository.IArticleRepo) *Server {
	return &Server{Repo: articleRepo}
}
