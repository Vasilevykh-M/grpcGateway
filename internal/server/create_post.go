package server

import (
	"awesomeProject/internal/repository"
	"context"
	"errors"
	"net/http"
)

func (s *Server) CreatePost(ctx context.Context, id int64, unm *PostRequest) (int, *PostRequest) {
	postRepo := &repository.Post{
		Name:     unm.Name,
		IdAuthor: id,
		Sales:    unm.Sales,
	}

	id, err := s.Repo.AddPost(ctx, postRepo)
	if errors.Is(err, repository.ErrObjectNotFound) {
		return http.StatusNotFound, nil
	}

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	unm.ID = id
	return http.StatusOK, unm
}
