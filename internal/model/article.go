package model

import "github.com/herrhu97/blog-service/pkg/app"

// article.go
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
