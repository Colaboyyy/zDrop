package main

import (
	"github.com/Colaboyyy/zGee/gee"
	"github.com/zserge/lorca"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		r := gee.New()
		r.Use(gee.Logger())
		r.LoadHTMLGlob("./frontend/*")
		r.Static("/assets", "./static")
		r.GET("/web", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "index.html", nil)
		})
		r.RUN(":9999")
	}()

	// Create UI with basic HTML passed via data URI
	var ui lorca.UI
	ui, _ = lorca.New("http://127.0.0.1:9999/web", "", 800, 600, "--disable-sync", "--disable-translate")
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ui.Done():
	case <-chSignal:
	}
	ui.Close()
}
