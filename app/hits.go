package app

import (
	"ClickHitsCount/infrastructure/logs"

	"github.com/gin-gonic/gin"
)

func Hits(c *gin.Context) {
	ip, name, chapter := c.ClientIP(), c.Query("name"), c.Query("chapter")
	// Check parameter is exist.
	if name == "" || chapter == "" {
		logs.Warning.Printf("&v error : %v", ip, "Not input name or chapter.")
		c.String(404, "Bad Request.")
		return
	}
	// Check novel is exist and get novel uid.
	uid, err := GetChapterUID(name, chapter)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	// Check access ip.
	err = AddAccessIP(uid, ip)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	// Get novel access count.
	count, err := GetNovelHits(uid)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	// Get count image.
	returnImageData, err := GetImage(count)
	if err != nil {
		logs.Warning.Printf("&v error : %v", ip, err.Error())
		c.String(404, "Bad Request.")
		return
	}
	c.Header("content-type", "image/svg+xml;charset=utf-8")
	c.Data(200, "OK", returnImageData)
	return
}
