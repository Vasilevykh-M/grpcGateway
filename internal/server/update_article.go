package server

import (
	"awesomeProject/internal/repository"
	"context"
	"errors"
	"net/http"
)

func (s *Server) UpdateArticle(ctx context.Context, unm *ArticleRequest) (int, *ArticleRequest) {
	articleRepo := &repository.Article{
		ID:     unm.ID,
		Name:   unm.Name,
		Rating: unm.Rating,
	}

	id, err := s.Repo.Update(ctx, articleRepo)

	if errors.Is(err, repository.ErrObjectNotFound) {
		return http.StatusNotFound, nil
	}

	if err != nil {
		return http.StatusInternalServerError, nil
	}
	unm.ID = id
	return http.StatusOK, unm
}
