package postgres

import (
	"awesomeProject/internal/db"
	"awesomeProject/internal/repository"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

type ArticleRepo struct {
	db db.DBops
}

func NewArticles(database db.DBops) *ArticleRepo {
	return &ArticleRepo{db: database}
}

func (r *ArticleRepo) AddArticles(ctx context.Context, article *repository.Article) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO articles(name_article, rating, created_at) VALUES($1,$2,$3) RETURNING id;`, article.Name, article.Rating, article.CreatedAt).Scan(&id)
	return id, err
}

func (r *ArticleRepo) AddPost(ctx context.Context, post *repository.Post) (int64, error) {

	tx, err := r.db.Begin(ctx)

	if err != nil {
		return 0, err
	}

	defer tx.Rollback(ctx)

	var a repository.Article
	err = r.db.Get(ctx, &a, "SELECT id,name_article,rating,created_at FROM articles WHERE id=$1", post.IdAuthor)

	if errors.Is(err, pgx.ErrNoRows) {
		return 0, repository.ErrObjectNotFound
	}

	if err != nil {
		return 0, repository.ErrObjectNotFound
	}

	var id int64
	err = r.db.ExecQueryRow(ctx, `INSERT INTO post(id_author, name_post, sales) VALUES($1, $2, $3) RETURNING id;`, post.IdAuthor, post.Name, post.Sales).Scan(&id)

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (r *ArticleRepo) GetByID(ctx context.Context, id int64) ([]*repository.JoinArticlePost, error) {
	var a []*repository.JoinArticlePost
	err := r.db.Select(ctx, &a, `SELECT articles.name_article, articles.rating, articles.created_at, post.id, post.name_post, post.sales FROM articles LEFT JOIN post ON articles.id = post.id_author WHERE articles.id=$1;`, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrObjectNotFound
	}
	if err != nil {
		return nil, repository.ErrNetwork
	}

	if a == nil {
		return nil, repository.ErrObjectNotFound
	}
	return a, nil
}

func (r *ArticleRepo) Delete(ctx context.Context, id int64) error {
	err := r.db.Exec(ctx, `DELETE FROM articles WHERE id = $1;`, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return repository.ErrObjectNotFound
	}
	if err != nil {
		return repository.ErrNetwork
	}
	return nil
}

func (r *ArticleRepo) Update(ctx context.Context, article *repository.Article) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `UPDATE articles SET (name_article, rating)=($1,$2) WHERE id = $3 RETURNING id;`, article.Name, article.Rating, article.ID).Scan(&id)

	if errors.Is(err, pgx.ErrNoRows) {
		return 0, repository.ErrObjectNotFound
	}

	if err != nil {
		return 0, repository.ErrNetwork
	}
	return id, err
}
