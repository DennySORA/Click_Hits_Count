package app

import (
	"ClickHitsCount/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func CreateNovel(c *gin.Context) {
	name := c.Query("name")
	_, err := database.HitsDatbase.Stmt["CreateNovel"].Exec(name)
	if err != nil {
		c.String(404, err.Error())
		return
	}
	c.String(200, "OK")
}

func CreateChapter(c *gin.Context) {
	id := c.Query("id")
	ep := c.Query("ep")
	chapter := c.Query("chapter")
	name := c.Query("name")
	uid, err := GetNovelUID(id)
	if err != nil {
		c.String(404, err.Error())
		return
	}
	_, err = database.HitsDatbase.Stmt["CreateNovelChapter"].Exec(uid, ep, chapter, name)
	if err != nil {
		c.String(404, err.Error())
		return
	}
	c.String(200, "OK")
}
