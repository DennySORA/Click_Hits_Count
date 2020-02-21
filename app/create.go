package app

import (
	"ClickHitsCount/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func CreateNovel(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.String(404, "Please input name.")
	}
	_, err := database.HitsDatbase.Stmt["CreateNovelFromDatabase"].Exec(name)
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
	if name == "" || id == "" || chapter == "" || ep == "" {
		c.String(404, "Please input name.")
	}
	uid, err := GetNovelUIDFromDatabase(id)
	if err != nil {
		c.String(404, err.Error())
		return
	}
	_, err = database.HitsDatbase.Stmt["CreateChapterFromDatabase"].Exec(uid, ep, chapter, name)
	if err != nil {
		c.String(404, err.Error())
		return
	}
	c.String(200, "OK")
}
