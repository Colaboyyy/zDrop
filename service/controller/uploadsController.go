package controller

import (
	"gee"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func getUploadsDir() (uploads string) {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	uploads = filepath.Join(dir, "uploads")
	return
}

func UploadsController(c *gee.Context) {
	if path := c.Param("path"); path != "" {
		target := filepath.Join(getUploadsDir(), path)
		c.SetHeader("Content-Description", "File Transfer")
		c.SetHeader("Content-Transfer-Encoding", "binary")
		c.SetHeader("Content-Disposition", "attachment; filename="+path)
		c.SetHeader("Content-Type", "application/octet-stream")
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}
}
