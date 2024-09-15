package controllers

import (
	"context"
	"net/http"
	"net/url"
	"splitter/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type GroupController struct {
	DB *bun.DB
}

func (gc *GroupController) GetGroup(c *gin.Context) {
	ctx := context.Background()
	tx, err := gc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	ur := repositories.UserRepository{TX: &tx}
	gr := repositories.GroupRepository{TX: &tx}
	group, err := gr.FindByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	users, err := ur.FindByGroupID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "group": group})
}

func (gc *GroupController) PostGroup(c *gin.Context) {
	var getData struct {
		Name  string   `json:"name"`
		Users []string `json:"users"`
	}

	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if getData.Name == "" || len(getData.Users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group or users"})
		return
	}

	ctx := context.Background()
	tx, err := gc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupRepository := repositories.GroupRepository{TX: &tx}
	group, err := groupRepository.Create(ctx, getData.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	userRepository := repositories.UserRepository{TX: &tx}
	users, err := userRepository.Create(ctx, getData.Users, group.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"group": group, "users": users})
}

func (gc *GroupController) PatchGroup(c *gin.Context) {
	id := c.Param("id")
	rawName := c.DefaultQuery("name", "Name")

	name, err := url.QueryUnescape(rawName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name"})
		return
	}

	ctx := context.Background()
	tx, err := gc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groupRepository := repositories.GroupRepository{TX: &tx}
	group, err := groupRepository.UpdateName(ctx, id, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": group.Name, "updated_at": group.UpdatedAt})
}
