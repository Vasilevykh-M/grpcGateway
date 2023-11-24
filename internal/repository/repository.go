//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository

package repository

import (
	"context"
)

type IArticleRepo interface {
	AddArticles(ctx context.Context, article *Article) (int64, error)
	AddPost(ctx context.Context, post *Post) (int64, error)
	GetByID(ctx context.Context, id int64) ([]*JoinArticlePost, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, article *Article) (int64, error)
}
