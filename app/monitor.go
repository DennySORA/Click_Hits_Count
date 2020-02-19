package app

import (
	"ClickHitsCount/infrastructure/logs"

	"github.com/gin-gonic/gin"
)

func GetChapterHits(c *gin.Context) {
	var uid int
	ip, name, chapter := c.ClientIP(), c.Query("name"), c.Query("chapter")
	if name != "" && chapter != "" {
		var err error
		// Check novel is exist and get novel uid.
		uid, err = GetChapterUID(name, chapter)
		if err != nil {
			logs.Warning.Printf("&v error : %v", ip, err.Error())
			c.String(404, "Bad Request.")
			return
		}
	} else {
		uid = -1
	}
	data, err := GetChapterHitsCounts(uid)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.JSON(200, data)
	return
}

func GetAllChapterHits(c *gin.Context) {
	ip := c.ClientIP()
	data, err := GetAllChapterHitsCounts()
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.JSON(200, data)
	return
}

func GetAllData(c *gin.Context) {
	ip := c.ClientIP()
	data, err := GetAllHitsData()
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.JSON(200, data)
	return
}
