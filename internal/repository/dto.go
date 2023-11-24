package repository

import "time"

type JoinArticlePost struct {
	NameArticle string    `db:"name_article"`
	Rating      int64     `db:"rating"`
	CreatedAt   time.Time `db:"created_at"`
	IDPost      *int64    `db:"id"`
	NamePost    *string   `db:"name_post"`
	Sales       *int64    `db:"sales"`
}

type Post struct {
	ID       int64  `db:"id"`
	IdAuthor int64  `db:"id_author"`
	Name     string `db:"name"`
	Sales    int64  `db:"sales"`
}

type Article struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name_article"`
	Rating    int64     `db:"rating"`
	CreatedAt time.Time `db:"created_at"`
}
