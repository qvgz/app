// tool合集
package main

import (
	"app/webtool/dfb"
	"app/webtool/ktds"
	"app/webtool/lib"
	"bytes"
	"flag"
	"os"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	port := flag.String("port", "8081", "监听端口")
	flag.Parse()

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formatFloat": lib.FormatFloat,
		"formatTime": func(t time.Time) string {
			return t.Format("2006/01/02")
		},
	})
	r.LoadHTMLGlob("templates/*.html")
	// r := r.Group("/tool")
	r.GET("/", func(ctx *gin.Context) {
		file, _ := os.ReadFile("README.md")
		md := goldmark.New(
			goldmark.WithExtensions(extension.GFM),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
			),
		)
		var buf bytes.Buffer
		if err := md.Convert(file, &buf); err != nil {
			panic(err)
		}

		ctx.Data(200, "text/html; charset=utf-8", buf.Bytes())
	})
	r.GET("/dfb", dfb.Get)
	r.POST("/dfb", dfb.Post)

	r.GET("/ktds", ktds.Get)
	r.POST("/ktds", ktds.Post)

	r.Run(":" + *port)
}
