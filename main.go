package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/*.html")
	pageStore := persistence.NewInMemoryStore(time.Hour * 24)


	r.GET("/", cache.CachePage(pageStore, time.Minute,func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}))

	r.Static("/public", "./public")

	_ = r.Run( )
}
