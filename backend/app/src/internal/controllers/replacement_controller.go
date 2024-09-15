package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"splitter/internal/models"
	"splitter/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReplacementController struct {
	DB *bun.DB
}

type tempReplacementUser struct {
	UserID string `json:"user_id"`
	Amount uint   `json:"amount"`
}

type resReplacementLessor struct {
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt bun.NullTime
	Calc      int
	Amount    int
	From      []tempReplacementUser
}

func (rrl resReplacementLessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID    string                `json:"user_id"`
		CreatedAt time.Time             `json:"created_at"`
		UpdatedAt time.Time             `json:"updated_at"`
		DeletedAt bun.NullTime          `json:"deleted_at"`
		Amount    int                   `json:"amount"`
		From      []tempReplacementUser `json:"from"`
	}{
		UserID:    rrl.UserID,
		CreatedAt: rrl.CreatedAt,
		UpdatedAt: rrl.UpdatedAt,
		DeletedAt: rrl.DeletedAt,
		Amount:    rrl.Amount,
		From:      rrl.From,
	})
}

func fromReplacementLessor(rl models.ReplacementLessor) resReplacementLessor {
	return resReplacementLessor{
		UserID:    rl.UserID.String(),
		CreatedAt: rl.CreatedAt,
		UpdatedAt: rl.UpdatedAt,
		DeletedAt: rl.DeletedAt,
		Calc:      int(rl.Amount),
		Amount:    int(rl.Amount),
	}
}

type resReplacementLessee struct {
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt bun.NullTime
	Calc      int
	Amount    int
	To        []tempReplacementUser
}

func (rrl resReplacementLessee) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID    string                `json:"user_id"`
		CreatedAt time.Time             `json:"created_at"`
		UpdatedAt time.Time             `json:"updated_at"`
		DeletedAt bun.NullTime          `json:"deleted_at"`
		Amount    int                   `json:"amount"`
		To        []tempReplacementUser `json:"to"`
	}{
		UserID:    rrl.UserID,
		CreatedAt: rrl.CreatedAt,
		UpdatedAt: rrl.UpdatedAt,
		DeletedAt: rrl.DeletedAt,
		Amount:    rrl.Amount,
		To:        rrl.To,
	})
}

func fromReplacementLessee(rl models.ReplacementLessee) resReplacementLessee {
	return resReplacementLessee{
		UserID:    rl.UserID.String(),
		CreatedAt: rl.CreatedAt,
		UpdatedAt: rl.UpdatedAt,
		DeletedAt: rl.DeletedAt,
		Calc:      int(rl.Amount),
		Amount:    int(rl.Amount),
	}
}

type reqBody struct {
	Title   string                `json:"title"`
	GroupID string                `json:"group_id"`
	Lessors []tempReplacementUser `json:"lessors"`
	Lessees []tempReplacementUser `json:"lessees"`
}

func (rc *ReplacementController) GetReplacement(c *gin.Context) {
	// Path
	// replacement_id string(uuid) required

	// Process
	// 1. Find data in 'replacements', 'replacement_lessors' and 'replacement_lessees' where replacement_id in the talbes is equal to replacement_id of path
	// 2. Compute who will be returned the money from lessees.

	// Initialization

	replacementID := c.Param("replacement_id")

	replacementUUID, err := uuid.Parse(replacementID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tx, err := rc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Find data in 'replacements', 'replacement_lessors' and 'replacement_lessees' where replacement_id in the talbes is equal to replacement_id of path

	replacementRepository := repositories.ReplacementRepository{TX: &tx}
	rawReplacement, err := replacementRepository.FindByID(ctx, replacementUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	replacementLessorRepository := repositories.ReplacementLessorRepository{TX: &tx}
	rawReplacementLessors, err := replacementLessorRepository.FindByReplacementID(ctx, replacementUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	replacementLesseeRepository := repositories.ReplacementLesseeRepository{TX: &tx}
	rawReplacementLessees, err := replacementLesseeRepository.FindByReplacementID(ctx, replacementUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 2. Compute who will be returned the money from lessees.

	lessors := make([]resReplacementLessor, len(rawReplacementLessors))
	for i, lessor := range rawReplacementLessors {
		lessors[i] = fromReplacementLessor(lessor)
	}

	lessees := make([]resReplacementLessee, len(rawReplacementLessees))
	for i, lessee := range rawReplacementLessees {
		lessees[i] = fromReplacementLessee(lessee)
	}

	for i, j := 0, 0; i < len(lessees) && j < len(lessors); {
		if lessors[j].Calc == 0 {
			break
		}
		if lessees[i].Calc < lessors[j].Calc {
			debt := lessees[i].Calc
			lessees[i].To = append(lessees[i].To, tempReplacementUser{UserID: lessors[j].UserID, Amount: uint(debt)})
			lessors[j].From = append(lessors[j].From, tempReplacementUser{UserID: lessees[i].UserID, Amount: uint(debt)})
			lessees[i].Calc -= debt
			lessors[j].Calc -= debt
			i++
			continue
		}
		if lessees[i].Calc > lessors[j].Calc {
			debt := lessors[j].Calc
			lessees[i].To = append(lessees[i].To, tempReplacementUser{UserID: lessors[j].UserID, Amount: uint(debt)})
			lessors[j].From = append(lessors[j].From, tempReplacementUser{UserID: lessees[i].UserID, Amount: uint(debt)})
			lessees[i].Calc -= debt
			lessors[j].Calc -= debt
			j++
			continue
		}
		if lessees[i].Calc == lessors[j].Calc {
			debt := lessees[i].Calc
			lessees[i].To = append(lessees[i].To, tempReplacementUser{UserID: lessors[j].UserID, Amount: uint(debt)})
			lessors[j].From = append(lessors[j].From, tempReplacementUser{UserID: lessees[i].UserID, Amount: uint(debt)})
			lessees[i].Calc -= debt
			lessors[j].Calc -= debt
			j++
			i++
		}
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"replacement": rawReplacement,
		"lessors":     lessors,
		"lessees":     lessees,
	})
}

func (rc *ReplacementController) PostReplacement(c *gin.Context) {
	// Request body json
	// title string required
	// group_id string(uuid) required
	// lessors []object(tempReplacementUser) required
	// lessees []object(tempReplacementUser) required

	// Process
	// 1. Create new data of replacements
	// 2. Create new data of replacement_lessors using replacement_id
	// 3. Create new data of replacement_lessees using replacement_id

	// Initialization

	var rb reqBody
	if err := c.ShouldBindJSON(&rb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var lessorsTotal uint = 0
	for _, lessor := range rb.Lessors {
		lessorsTotal += lessor.Amount
	}

	var lesseesTotal uint = 0
	for _, lessee := range rb.Lessees {
		lesseesTotal += lessee.Amount
	}

	if lessorsTotal != lesseesTotal {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lessors total amount is not equal to Lesses total amount."})
		return
	}

	groupUUID, err := uuid.Parse(rb.GroupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "group_id is not uuid format."})
		return
	}

	ctx := context.Background()
	tx, err := rc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Create new data of replacements

	replacementRepository := repositories.ReplacementRepository{TX: &tx}
	replacement, err := replacementRepository.Create(ctx, rb.Title, lessorsTotal, groupUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 2. Create new data of replacement_lessors using replacement_id

	lessors := make([]models.ReplacementLessor, len(rb.Lessors))

	for i, lessor := range rb.Lessors {
		lessorUUID, err := uuid.Parse(lessor.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}
		lessors[i] = models.ReplacementLessor{
			UserID:        lessorUUID,
			Amount:        lessor.Amount,
			ReplacementID: replacement.ID,
		}
	}

	replacementLessorRepository := repositories.ReplacementLessorRepository{TX: &tx}
	lessors, err = replacementLessorRepository.Create(ctx, lessors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 3. Create new data of replacement_lessees using replacement_id

	lessees := make([]models.ReplacementLessee, len(rb.Lessees))

	for i, lessee := range rb.Lessees {
		lesseeUUID, err := uuid.Parse(lessee.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}
		lessees[i] = models.ReplacementLessee{
			UserID:        lesseeUUID,
			Amount:        lessee.Amount,
			ReplacementID: replacement.ID,
		}
	}

	replacementLesseeRepository := repositories.ReplacementLesseeRepository{TX: &tx}
	lessees, err = replacementLesseeRepository.Create(ctx, lessees)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{"replacement": replacement, "lessors": lessors, "lessees": lessees})
}

func (rc *ReplacementController) PutReplacement(c *gin.Context) {
	// Path
	// replacement_id uuid required

	// Request body json
	// title string option
	// lessors []object(tempReplacementUser) required
	// lessees []object(tempReplacementUser) required

	// 1. Update 'replacements' table
	// 2. Update 'replacement_lessors' data
	// 3. Update 'replacement_lessees' data

	// Initialization
	replacementID := c.Param("replacement_id")
	replacementUUID, err := uuid.Parse(replacementID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var rb reqBody
	if err := c.ShouldBindJSON(&rb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var lessorsTotal uint = 0
	for _, lessor := range rb.Lessors {
		lessorsTotal += lessor.Amount
	}

	var lesseesTotal uint = 0
	for _, lessee := range rb.Lessees {
		lesseesTotal += lessee.Amount
	}

	if lessorsTotal != lesseesTotal {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lessors total amount is not equal to Lesses total amount."})
		return
	}

	newLessors := make([]models.ReplacementLessor, len(rb.Lessors))
	for i, lessor := range rb.Lessors {
		userUUID, err := uuid.Parse(lessor.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newLessors[i] = models.ReplacementLessor{
			UserID:        userUUID,
			ReplacementID: replacementUUID,
			Amount:        lessor.Amount,
		}
	}

	newLessees := make([]models.ReplacementLessee, len(rb.Lessees))
	for i, lessee := range rb.Lessees {
		userUUID, err := uuid.Parse(lessee.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newLessees[i] = models.ReplacementLessee{
			UserID:        userUUID,
			ReplacementID: replacementUUID,
			Amount:        lessee.Amount,
		}
	}

	ctx := context.Background()
	tx, err := rc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Update 'replacements' table

	replacementRepository := repositories.ReplacementRepository{TX: &tx}
	_, err = replacementRepository.Update(ctx, rb.Title, lessorsTotal, replacementUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + " replacements"})
		tx.Rollback()
		return
	}

	// 2. Update 'replacement_lessors' data
	replacementLessorRepository := repositories.ReplacementLessorRepository{TX: &tx}
	err = replacementLessorRepository.Update(ctx, newLessors, replacementUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + " replacement_lessors"})
		tx.Rollback()
		return
	}

	// 3. Update 'replacement_lessees' data
	replacementLesseeRepository := repositories.ReplacementLesseeRepository{TX: &tx}
	err = replacementLesseeRepository.Update(ctx, newLessees, replacementUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + " replacement_lessees"})
		tx.Rollback()
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// Response
	c.Status(200)
}

func (rc *ReplacementController) PatchDoneReplacement(c *gin.Context) {
	// Path
	// replacement_id uuid

	// Process
	// 1. 'done' in 'replacements' be changed true

	// Initialization
	replacementID := c.Param("replacement_id")
	replacementUUID, err := uuid.Parse(replacementID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tx, err := rc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. 'done' in 'replacements' be changed true
	replacementRepository := repositories.ReplacementRepository{TX: &tx}
	if err := replacementRepository.UpdateDone(ctx, true, replacementUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	c.Status(http.StatusOK)
}

func (rc *ReplacementController) PatchDoingReplacement(c *gin.Context) {
	// Path
	// replacement_id

	// 1. 'done' in 'replacements' be changed false.

	// Initialization
	replacementID := c.Param("replacement_id")
	replacementUUID, err := uuid.Parse(replacementID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tx, err := rc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. 'done' in 'replacements' be changed false.
	replacementRepository := repositories.ReplacementRepository{TX: &tx}

	if err := replacementRepository.UpdateDone(ctx, false, replacementUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	c.Status(http.StatusOK)
}

func (rc *ReplacementController) DeleteReplacement(c *gin.Context) {
	// Path
	// replacement_id

	// 1. Delete the data in replacement_lessees where replacement_id in the table is equal to replacement_id of request path
	// 2. Delete the data in replacement_lessors where replacement_id in the table is equal to replacement_id of request path
	// 3. Delete the data in replacements where replacement_id in the table is equal to replacement_id of request path

	// Initialization
	replacementID := c.Param("replacement_id")
	replacementUUID, err := uuid.Parse(replacementID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tx, err := rc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 1. Delete the data in replacement_lessees where replacement_id in the table is equal to replacement_id of request path
	replacementLesseeRepository := repositories.ReplacementLesseeRepository{TX: &tx}
	if err := replacementLesseeRepository.Delete(ctx, replacementUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 2. Delete the data in replacement_lessors where replacement_id in the table is equal to replacement_id of request path
	replacementLessorRepository := repositories.ReplacementLessorRepository{TX: &tx}
	if err := replacementLessorRepository.Delete(ctx, replacementUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. Delete the data in replacements where replacement_id in the table is equal to replacement_id of request path
	replacementRepository := repositories.ReplacementRepository{TX: &tx}
	if err := replacementRepository.Delete(ctx, replacementUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	c.Status(http.StatusOK)
}
