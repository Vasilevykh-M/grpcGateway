package server

import "time"

type ArticleRequest struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Rating    int64     `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
}

type JoinArticlePost struct {
	NameArticle string    `json:"article_name"`
	Rating      int64     `json:"article_rating"`
	CreatedAt   time.Time `json:"article_created_at"`
	IDPost      int64     `json:"post_id"`
	NamePost    string    `json:"post_name"`
	Sales       int64     `json:"post_sales"`
}

type PostRequest struct {
	ID       int64  `json:"id"`
	IdAuthor int64  `json:"id_author"`
	Name     string `json:"name"`
	Sales    int64  `json:"sales"`
}
