package server

import (
	"awesomeProject/internal/repository"
	"context"
	"net/http"
	"time"
)

func (s *Server) CreateArticle(ctx context.Context, unm *ArticleRequest) (int, *ArticleRequest) {
	articleRepo := &repository.Article{
		Name:      unm.Name,
		Rating:    unm.Rating,
		CreatedAt: time.Now(),
	}

	id, err := s.Repo.AddArticles(ctx, articleRepo)
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	unm.ID = id

	return http.StatusOK, unm
}
