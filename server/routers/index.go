package routers

import (
	"database/sql"
	routers "main/routers/routes"

	"path/filepath"

	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	// router.LoadHTMLGlob("site/templates/*")
	// router.SetHTMLTemplate(template.Must(template.ParseFiles(
	// 	"site/templates/layout.tmpl",
	// 	"site/templates/home.tmpl",
	// 	"site/templates/create-post.tmpl",
	// )))

	router.HTMLRender = loadTemplates("./site/templates")

	// Load scripts
	router.Static("/site/scripts/", "./site/scripts/")

	// 404 Route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	routers.IndexRoutes(router)
	routers.PostRoutes(router)
	routers.DatabaseRoutes(router, db)
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/site/templates/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, layout)
		r.AddFromFiles(filepath.Base(layout), files...)
	}
	return r
}
