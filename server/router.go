package server

import (
	"ClickHitsCount/app"

	"github.com/gin-gonic/gin"
)

func router(engine *gin.Engine) {
	engine.GET("/hits", app.Hits)
}

func routerLocal(engineLocal *gin.Engine) {
	engineLocal.GET("/create/novel", app.CreateNovel)
	engineLocal.GET("/create/chapter", app.CreateChapter)
}
