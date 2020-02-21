package server

import (
	"ClickHitsCount/app"

	"github.com/gin-gonic/gin"
)

func router(engine *gin.Engine) {
	engine.GET("/hits", app.Hits)
}

func routerLocal(engineLocal *gin.Engine) {
	// Create part.
	engineLocal.GET("/create/novel", app.CreateNovel)
	engineLocal.GET("/create/chapter", app.CreateChapter)
	// Get part.
	engineLocal.GET("/get/all/hits/ip", app.GetAllHitsIP)
	engineLocal.GET("/get/chapter/hits", app.GetChapterHits)
	engineLocal.GET("/get/chapter/ip", app.GetChapterIP)
	engineLocal.GET("/get/novel/hits", app.GetNovelHits)
	engineLocal.GET("/get/ep/hits",app.GetEpisodeHits)
}
