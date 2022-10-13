package controller

import (
	"fmt"
	"gee"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func QrcodeController(c *gee.Context) {
	if content := c.Query("content"); content != "" {
		fmt.Println("=== content:", content)
		png, err := qrcode.Encode(content, qrcode.Medium, 256)
		if err != nil {
			log.Fatal("qrcode encode error!", err)
		}
		c.SetHeader("Content-Type", "image/png")
		c.Data(http.StatusOK, png)
	} else {
		c.Status(http.StatusBadRequest)
	}
}
