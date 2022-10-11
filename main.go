package main

import (
	"embed"
	"gee"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var Fs embed.FS

func main() {
	r := gee.Default()
	staticFiles, _ := fs.Sub(Fs, "dist")
	r.StaticFs("/static", http.FS(staticFiles))

	r.RUN(":27149")
}
