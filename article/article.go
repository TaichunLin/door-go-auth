package article

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []article {
	return articleList
}

func ShowIndexPage(c *gin.Context) {
	articles := GetAllArticles()

	c.HTML(
		200,
		"index.html",
		gin.H{
			"is_logged_in": c.MustGet("is_logged_in").(bool),
			"Title":        "Home Page",
			"payload":      articles,
		},
	)

}

func GetArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}

func GetArticle(c *gin.Context) {

	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {

		if article, err := GetArticleByID(articleID); err == nil {
			c.HTML(
				200,
				"article.html",
				gin.H{
					"is_logged_in": c.MustGet("is_logged_in").(bool),
					"Title":        article.Title,
					"payload":      article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//submission

func CreateNewArticle(title, content string) (*article, error) {
	a := article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, a)

	return &a, nil
}

func ShowArticleCreationPage(c *gin.Context) {
	c.HTML(200, "create-article.html", gin.H{
		"is_logged_in": c.MustGet("is_logged_in").(bool),
		"Title":        "Create New Article",
	})
}

func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := CreateNewArticle(title, content); err == nil {
		c.HTML(200, "submission-successful.html", gin.H{
			"is_logged_in": c.MustGet("is_logged_in").(bool),
			"Title":        "Submission Successful",
			"payload":      a,
		})

	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
