package controllers

import (
	"context"
	"net/http"
	"splitter/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaymentController struct {
	DB *bun.DB
}

func (pc *PaymentController) GetPayment(c *gin.Context) {
	// Validation
	groupID := c.Param("group_id")
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group ID is not UUID format"})
		return
	}

	// Initialization
	ctx := context.Background()
	tx, err := pc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Get payments (splits and replacemnets) where 'group_id' of query param is equal to 'group_id' in database.
	pr := repositories.PaymentRepository{TX: &tx}
	payments, err := pr.FindByGroupID(ctx, groupUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response
	c.JSON(http.StatusOK, payments)
}
