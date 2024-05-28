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

	u := r.Group("/users")
	{
		userController := controllers.UserController{DB: db}
		u.GET("/:group_id", userController.GetUser)
		u.POST("/:group_id", userController.PostUser)
		u.PATCH("/:user_id", userController.PatchUser)
	}

	p := r.Group("/payments")
	{
		paymentController := controllers.PaymentController{DB: db}
		p.GET("/:group_id", paymentController.GetPayment)
	}
}
