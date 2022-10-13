package service

import (
	"gee"
	"github.com/Colaboyyy/zDrop/service/controller"
	"github.com/Colaboyyy/zDrop/service/ws"
)

func RUN(r *gee.Engine) {
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

	//r.RUN(":27149")
}
