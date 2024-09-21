package routers

import (
	"splitter/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func SetupRouter(router *gin.Engine, db *bun.DB) {

	groups := router.Group("/groups")
	{
		groupController := controllers.GroupController{DB: db}
		groups.GET("/:id", groupController.GetGroup)
		groups.POST("", groupController.PostGroup)
		groups.PATCH("/:id", groupController.PatchGroup)
	}

	users := router.Group("/users")
	{
		userController := controllers.UserController{DB: db}
		users.GET("/:group_id", userController.GetUser)
		users.POST("/:group_id", userController.PostUser)
		users.PATCH("/:user_id", userController.PatchUser)
	}

	payments := router.Group("/payments")
	{
		paymentController := controllers.PaymentController{DB: db}
		payments.GET("/:group_id", paymentController.GetPayment)
	}

	splits := router.Group("/splits")
	{
		splitController := controllers.SplitController{DB: db}
		splits.GET("/:split_id", splitController.GetSplit)
		splits.POST("", splitController.PostSplit)
		splits.PUT("/:split_id", splitController.PutSplit)
		splits.PATCH("/done/:split_id", splitController.PatchDoneSplit)
		splits.PATCH("/doing/:split_id", splitController.PatchDoingSplit)
		splits.DELETE("/:split_id", splitController.DeleteSplit)
	}

	replacements := router.Group("/replacements")
	{
		replacementController := controllers.ReplacementController{DB: db}
		replacements.GET("/:replacement_id", replacementController.GetReplacement)
		replacements.POST("", replacementController.PostReplacement)
		replacements.PUT("/:replacement_id", replacementController.PutReplacement)
		replacements.PATCH("/done/:replacement_id", replacementController.PatchDoneReplacement)
		replacements.PATCH("/doing/:replacement_id", replacementController.PatchDoingReplacement)
		replacements.DELETE("/:replacement_id", replacementController.DeleteReplacement)
	}
}
