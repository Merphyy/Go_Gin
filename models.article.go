package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(articleID int) (*article, error){
	for _, a := range articleList{
		if a.ID == articleID {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}