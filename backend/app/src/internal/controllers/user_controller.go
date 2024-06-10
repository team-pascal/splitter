package controllers

import (
	"context"
	"net/http"
	"net/url"
	"splitter/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserController struct {
	DB *bun.DB
}

func (uc *UserController) GetUser(c *gin.Context) {
	ctx := context.Background()
	tx, err := uc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	group_id := c.Param("group_id")

	ur := repositories.UserRepository{TX: &tx}

	users, err := ur.FindByGroupID(ctx, group_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) PostUser(c *gin.Context) {
	rawGroupID := c.Param("group_id")
	groupID, err := uuid.Parse(rawGroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rawName := c.DefaultQuery("name", "Name")

	name, err := url.QueryUnescape(rawName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name"})
		return
	}

	ctx := context.Background()
	tx, err := uc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRepository := repositories.UserRepository{TX: &tx}
	user, err := userRepository.Create(ctx, []string{name}, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user[0])
}

func (uc *UserController) PatchUser(c *gin.Context) {
	userID := c.Param("user_id")

	rawName := c.DefaultQuery("name", "Name")

	name, err := url.QueryUnescape(rawName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name"})
		return
	}

	ctx := context.Background()
	tx, err := uc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRepository := repositories.UserRepository{TX: &tx}
	user, err := userRepository.UpdateName(ctx, name, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": user.Name, "updated_at": user.UpdatedAt})
}
