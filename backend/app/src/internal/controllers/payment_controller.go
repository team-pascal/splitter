package controllers

import (
	"context"
	"net/http"
	"splitter/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type PaymentController struct {
	DB *bun.DB
}

func (pc *PaymentController) GetPayment(c *gin.Context) {
	ctx := context.Background()
	tx, err := pc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	group_id := c.Param("group_id")

	pr := repositories.PaymentRepository{TX: &tx}

	payments, err := pr.FindByGroupID(ctx, group_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}
