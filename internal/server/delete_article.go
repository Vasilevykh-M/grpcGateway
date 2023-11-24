package server

import (
	"awesomeProject/internal/repository"
	"context"
	"errors"
	"net/http"
)

func (s *Server) DeleteArticle(ctx context.Context, id int64) int {

	err := s.Repo.Delete(ctx, id)

	if errors.Is(err, repository.ErrObjectNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, repository.ErrNetwork) {
		return http.StatusInternalServerError
	}

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
