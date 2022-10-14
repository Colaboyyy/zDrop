package main

import (
	"embed"
	"gee"
	"github.com/Colaboyyy/zDrop/service/controller"
	"github.com/Colaboyyy/zDrop/service/ws"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/dict/*
var Fs embed.FS

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	staticFiles, _ := fs.Sub(Fs, "frontend/dict")
	//r.StaticFs("/static", http.FS(staticFiles))
	r.LoadHTMLGlob("/Users/tommyyzhang/GolandProjects/zDrop/frontend/dict/index.html")
	r.Static("/assets", "./frontend/dict/assets")
	r.GET("/static", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	//service.RUN(r)
	v1 := r.Group("/api/v1")
	{
		v1.POST("/texts", controller.TextController)
		v1.POST("/files", controller.FileController)
		v1.GET("/addresses", controller.AddressesController)
		v1.GET("/qrcodes", controller.QrcodeController)
	}
	r.GET("/uploads/:path", controller.UploadsController)

	// ws
	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", func(c *gee.Context) {
		ws.WsController(c, hub)
	})
	r.NoRoute(func(c *gee.Context) {
		path := c.Req.URL.Path
		if strings.HasPrefix(path, "/static/") {
			log.Println("--HasPrefix static")
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal("staticFiles.Open failed, ", err)
			}
			defer reader.Close()
			_, err = reader.Stat()
			if err != nil {
				log.Fatal("reader.Stat failed, ", err)
			}
			c.DataFromReader(http.StatusOK, "text/html;charset=utf-8", reader, nil)
		} else {
			log.Println("-- Don't HasPrefix static")
			c.Status(http.StatusNotFound)
		}
	})

	r.RUN(":27149")
}
