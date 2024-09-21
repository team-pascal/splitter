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

type GroupController struct {
	DB *bun.DB
}

func (gc *GroupController) GetGroup(c *gin.Context) {
	// Path
	// group_id string(uuid)

	// Process
	// 1. Find group in 'groups' where 'group_id' is equal to request path parameter.
	// 2. Find users in 'users' where 'group_id' is equal to request path parameter.

	// Validation
	groupID := c.Param("id")
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your Group ID is not UUID format"})
	}

	// Initialization
	ctx := context.Background()
	tx, err := gc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Find group in 'groups' where 'group_id' is equal to request path parameter.
	gr := repositories.GroupRepository{TX: &tx}
	group, err := gr.FindByID(context.Background(), groupUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 2. Find users in 'users' where 'group_id' is equal to request path parameter.
	ur := repositories.UserRepository{TX: &tx}
	users, err := ur.FindByGroupID(context.Background(), groupUUID)
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
	c.JSON(http.StatusOK, gin.H{"users": users, "group": group})
}

func (gc *GroupController) PostGroup(c *gin.Context) {
	// Body JSON
	// name string
	// users []string

	// Process
	// 1. Register group
	// 2. Register users in the group

	// Validation
	var getData struct {
		Name  string   `json:"name"`
		Users []string `json:"users"`
	}

	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if getData.Name == "" || len(getData.Name) > 300 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group name"})
		return
	}

	if len(getData.Users) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user"})
		return
	}

	for _, user := range getData.Users {
		if user == "" || len(user) > 300 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user name"})
			return
		}
	}

	// Initialization
	ctx := context.Background()
	tx, err := gc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Register group
	groupRepository := repositories.GroupRepository{TX: &tx}
	group, err := groupRepository.Create(ctx, getData.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 2. Register users in the group
	userRepository := repositories.UserRepository{TX: &tx}
	users, err := userRepository.Create(ctx, getData.Users, group.ID)
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
	c.JSON(http.StatusOK, gin.H{"group": group, "users": users})
}

func (gc *GroupController) PatchGroup(c *gin.Context) {
	// Process
	// 1. Update group name

	// Validation
	id := c.Param("id")
	groupUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group ID is not UUID format"})
		return
	}

	rawName := c.DefaultQuery("name", "Name")
	name, err := url.QueryUnescape(rawName)
	if err != nil || len(name) <= 0 || len(name) > 300 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name"})
		return
	}

	// Initialization
	ctx := context.Background()
	tx, err := gc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Update group name
	groupRepository := repositories.GroupRepository{TX: &tx}
	group, err := groupRepository.UpdateName(ctx, groupUUID, name)
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
	c.JSON(http.StatusOK, gin.H{"name": group.Name, "updated_at": group.UpdatedAt})
}
