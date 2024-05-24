package routers

import (
	"splitter/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func SetupRouter(r *gin.Engine, db *bun.DB) {

	g := r.Group("/groups")
	{
		groupController := controllers.GroupController{DB: db}
		g.GET("/:id", groupController.GetGroup)
		g.POST("", groupController.PostGroup)
		g.PATCH("/:id", groupController.PatchGroup)
	}
}
