package server

import (
	"ClickHitsCount/infrastructure/logs"

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

func Start() {
	// Set gin release mode.
	gin.SetMode(gin.ReleaseMode)
	logFile, err := setServerLog()
	if err == nil {
		defer logFile.Close()
	}

	// Create gin engine.
	engine := gin.Default()
	engineLocal := gin.Default()

	// Registering router.
	router(engine)
	routerLocal(engineLocal)

	// Register pprof.
	pprof.Register(engineLocal)

	// Set port.
	port := viper.GetString("port.open")
	if port == "" {
		port = "0.0.0.0:8123"
	}
	portLoacl := viper.GetString("port.local")
	if portLoacl == "" {
		portLoacl = "127.0.0.1:8223"
	}

	// Set http server parameter.
	ser := &http.Server{
		Addr:           port,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      setServerTLS(),
	}

	serLocal := &http.Server{
		Addr:           portLoacl,
		Handler:        engineLocal,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Set http2.
	http2.ConfigureServer(ser, &http2.Server{})
	http2.ConfigureServer(serLocal, &http2.Server{})

	// Start listen.
	listen(ser, serLocal)

	// Wait shutdown.
	waitShutdown(ser, serLocal)

	// Print Exited info.
	logs.Info.Println("Server exiting")
}

func listen(ser, serLocal *http.Server) {
	go func(ser *http.Server) {
		err := ser.ListenAndServeTLS("ssl/server.crt", "ssl/server.key")
		if err != nil && err != http.ErrServerClosed {
			logs.Error.Fatalf("Listen: %s\n", err)
		}
	}(ser)

	go func(serLocal *http.Server) {
		err := serLocal.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logs.Error.Fatal("Local Listen: %s", err)
		}
	}(serLocal)
}

func waitShutdown(ser, serLocal *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Info.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctxLocal, cancelLocal := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelLocal()

	if err := ser.Shutdown(ctx); err != nil {
		logs.Error.Fatal("Server Shutdown: ", err)
	}

	if err := serLocal.Shutdown(ctxLocal); err != nil {
		logs.Error.Fatal("Server Shutdown: ", err)
	}
}
