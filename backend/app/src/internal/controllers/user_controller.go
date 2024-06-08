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
	// Path
	// group_id string(uuid)

	// Process
	// 1. Find users in 'users' where 'group_id' is equal to request path parameter.

	// Validation
	groupID := c.Param("group_id")
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group ID is not UUID format"})
		return
	}

	// Initialization
	ctx := context.Background()
	tx, err := uc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
		return
	}

	// 1. Find users in 'users' where 'group_id' is equal to request path parameter.
	ur := repositories.UserRepository{TX: &tx}
	users, err := ur.FindByGroupID(ctx, groupUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error"})
		tx.Rollback()
		return
	}

	if len(users) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users of group not found"})
		tx.Rollback()
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) PostUser(c *gin.Context) {
	// Path
	// group_id string(uuid)

	// Query
	// name string

	// Process
	// 1. Create user in existing group

	// Validation
	groupID := c.Param("group_id")
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group ID is not UUID format"})
		return
	}

	queryName := c.DefaultQuery("name", "Name")
	name, err := url.QueryUnescape(queryName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name"})
		return
	}

	// Initialization
	ctx := context.Background()
	tx, err := uc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Create user in existing group
	userRepository := repositories.UserRepository{TX: &tx}
	_, err = userRepository.Create(ctx, []string{name}, groupUUID)
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
	c.Status(http.StatusOK)
}

func (uc *UserController) PatchUser(c *gin.Context) {
	// Path
	// user_id string(uuid)

	// Query
	// name string

	// Process
	// 1. Update user name

	// Validation
	userID := c.Param("user_id")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group ID is not UUID format"})
		return
	}

	queryName := c.DefaultQuery("name", "Name")
	name, err := url.QueryUnescape(queryName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name"})
		return
	}

	// Initialization
	ctx := context.Background()
	tx, err := uc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Update user name
	userRepository := repositories.UserRepository{TX: &tx}
	_, err = userRepository.UpdateName(ctx, name, userUUID)
	if err != nil {
		responseStatus := http.StatusInternalServerError
		errorMessage := "Server Error"
		if err.Error() == "Not Found" {
			responseStatus = http.StatusNotFound
			errorMessage = err.Error()
		}
		c.JSON(responseStatus, gin.H{"error": errorMessage})
		tx.Rollback()
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response
	c.Status(http.StatusOK)
}
