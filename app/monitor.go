package app

import (
	"ClickHitsCount/infrastructure/logs"

	"github.com/gin-gonic/gin"
)

type Novel struct {
	chapter_id   interface{}
	novel_name   interface{}
	chapter_name interface{}
	ipaddress    interface{}
	episode      interface{}
	chapter      interface{}
}

func GetChapterIP(c *gin.Context) {
	ip := c.ClientIP()
	novel := Novel{
		chapter_id:   CheckString(c.Query("chapter_id")),
		novel_name:   CheckString(c.Query("novel_name")),
		chapter_name: CheckString(c.Query("chapter_name")),
		ipaddress:    CheckString(c.Query("ip")),
		episode:      CheckString(c.Query("ep")),
		chapter:      CheckString(c.Query("chapter")),
	}
	data, err := GetChapterIPFromDatabase(novel)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.JSON(200, data)
	return
}

func GetChapterHits(c *gin.Context) {
	ip := c.ClientIP()
	novel := Novel{
		chapter_id:   CheckString(c.Query("chapter_id")),
		novel_name:   CheckString(c.Query("novel_name")),
		chapter_name: CheckString(c.Query("chapter_name")),
		ipaddress:    CheckString(c.Query("ip")),
		episode:      CheckString(c.Query("ep")),
		chapter:      CheckString(c.Query("chapter")),
	}
	data, err := GetChapterHitsFromDatabase(novel)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.JSON(200, data)
	return
}

func GetNovelHits(c *gin.Context) {
	ip := c.ClientIP()
	novelName := CheckString(c.Query("novel_name"))
	data, err := GetNovelHitsFromDatabase(novelName)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.JSON(200, data)
	return
}

func GetEpisodeHits(c *gin.Context) {
	ip := c.ClientIP()
	novelName := CheckString(c.Query("novel_name"))
	data, err := GetEpisodeHitsFromDatabase(novelName)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.JSON(200, data)
	return
}

func GetAllHitsIP(c *gin.Context) {
	ip := c.ClientIP()
	data, err := GetAllHitsIPFromDatabase()
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.JSON(200, data)
	return
}
