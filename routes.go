package main

func initialRoutes() {
	router.GET("/", showIndexPage)	
	router.GET("/article/view/:article_id", getArticle)
}
