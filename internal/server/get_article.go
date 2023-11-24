package server

import (
	"awesomeProject/internal/repository"
	"context"
	"errors"
	"net/http"
)

func (s *Server) Get(ctx context.Context, id int64) (int, []*repository.JoinArticlePost) {

	joinArticlePost, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return http.StatusNotFound, nil
		}
		return http.StatusInternalServerError, nil
	}
	return http.StatusOK, joinArticlePost
}
