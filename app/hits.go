package app

import (
	"ClickHitsCount/infrastructure/database"
	"ClickHitsCount/infrastructure/logs"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

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

func GetNovelUID(id string) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetNovel"].QueryRow(id))
}

func GetChapterUID(name string, chapter string) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetChapter"].QueryRow(name, chapter))
}

func GetNovelHits(uid int) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetHits"].QueryRow(uid))
}

func GetOneValue(row *sql.Row) (int, error) {
	var data int
	if err := row.Scan(&data); err != nil {
		return 0, err
	}
	return data, nil
}

func AddAccessIP(uid int, ip string) error {
	_, err := database.HitsDatbase.Stmt["CreateCounts"].Exec(ip, uid)
	if err != nil {
		return err
	}
	return nil
}

func GetImage(count int) ([]byte, error) {
	url := fmt.Sprintf("https://img.shields.io/badge/hits-%d-green", count)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
