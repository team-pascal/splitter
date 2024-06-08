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

type SplitController struct {
	DB *bun.DB
}

type relatedSplitUser struct {
	UserID string `json:"user_id"`
	Amount int    `json:"amount"`
}

type resSplitLessee struct {
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt bun.NullTime
	Calc      int
	Amount    int
	To        []relatedSplitUser
}

func (rL resSplitLessee) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID    string             `json:"user_id"`
		CreatedAt time.Time          `json:"created_at"`
		UpdatedAt time.Time          `json:"updated_at"`
		DeletedAt bun.NullTime       `json:"deleted_at"`
		Amount    int                `json:"amount"`
		To        []relatedSplitUser `json:"to"`
	}{
		UserID:    rL.UserID,
		CreatedAt: rL.CreatedAt,
		UpdatedAt: rL.UpdatedAt,
		DeletedAt: rL.DeletedAt,
		Amount:    rL.Amount,
		To:        rL.To,
	})
}

type resSplitLessor struct {
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt bun.NullTime
	Calc      int
	Amount    int
	From      []relatedSplitUser
}

type tmpLessor struct {
	UserID string `json:"user_id"`
	Amount uint   `json:"amount"`
}

type getData struct {
	Title   string      `json:"title"`
	GroupID string      `json:"group_id"`
	Lessors []tmpLessor `json:"lessors"`
	Lessees []string    `json:"lessees"`
}

func (rL resSplitLessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID    string             `json:"user_id"`
		CreatedAt time.Time          `json:"created_at"`
		UpdatedAt time.Time          `json:"updated_at"`
		DeletedAt bun.NullTime       `json:"deleted_at"`
		Amount    int                `json:"amount"`
		From      []relatedSplitUser `json:"from"`
	}{
		UserID:    rL.UserID,
		CreatedAt: rL.CreatedAt,
		UpdatedAt: rL.UpdatedAt,
		DeletedAt: rL.DeletedAt,
		Amount:    rL.Amount,
		From:      rL.From,
	})
}

func (sc *SplitController) GetSplit(c *gin.Context) {
	ctx := context.Background()
	tx, err := sc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	split_id := c.Param("split_id")

	splitRepository := repositories.SplitRepository{TX: &tx}
	split, err := splitRepository.FindByID(ctx, split_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	splitLessorRepository := repositories.SplitLessorRepository{TX: &tx}
	rawLessors, err := splitLessorRepository.FindBySplitID(ctx, split_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	splitLesseeRepository := repositories.SplitLesseeRepository{TX: &tx}
	rawLessees, err := splitLesseeRepository.FindBySplitID(ctx, split_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	lessees := make([]resSplitLessee, 0)
	lessors := make([]resSplitLessor, 0)
	amount := int(split.Amount)
	count := (len(rawLessees) + len(rawLessors))
	ave := amount / count
	remainder := amount % count

	for i, lessee := range rawLessees {
		pls := 0
		if i < remainder {
			pls = 1
		}
		lessees = append(lessees, resSplitLessee{
			UserID:    lessee.UserID.String(),
			CreatedAt: lessee.CreatedAt,
			UpdatedAt: lessee.UpdatedAt,
			DeletedAt: lessee.DeletedAt,
			Calc:      ave + pls,
			Amount:    0,
		})
	}

	for i, lessor := range rawLessors {
		pls := 0
		if i+len(rawLessees) < remainder {
			pls = 1
		}
		userAmount := int(lessor.Amount) - ave - pls
		if userAmount < 0 {
			lessees = append(lessees, resSplitLessee{
				UserID:    lessor.UserID.String(),
				CreatedAt: lessor.CreatedAt,
				UpdatedAt: lessor.UpdatedAt,
				DeletedAt: lessor.DeletedAt,
				Calc:      0 - userAmount,
				Amount:    int(lessor.Amount),
			})
			continue
		}
		lessors = append(lessors, resSplitLessor{
			UserID:    lessor.UserID.String(),
			CreatedAt: lessor.CreatedAt,
			UpdatedAt: lessor.UpdatedAt,
			DeletedAt: lessor.DeletedAt,
			Calc:      userAmount,
			Amount:    int(lessor.Amount),
		})
	}

	for i, j := 0, len(lessors)-1; i < j; i, j = i+1, j-1 {
		lessors[i], lessors[j] = lessors[j], lessors[i]
	}

	for i, j := 0, 0; i < len(lessees) && j < len(lessors); {
		if lessors[j].Calc == 0 {
			break
		}
		if lessees[i].Calc < lessors[j].Calc {
			debt := lessees[i].Calc
			lessees[i].To = append(lessees[i].To, relatedSplitUser{UserID: lessors[j].UserID, Amount: debt})
			lessors[j].From = append(lessors[j].From, relatedSplitUser{UserID: lessees[i].UserID, Amount: debt})
			lessees[i].Calc -= debt
			lessors[j].Calc -= debt
			i++
			continue
		}
		if lessees[i].Calc > lessors[j].Calc {
			debt := lessors[j].Calc
			lessees[i].To = append(lessees[i].To, relatedSplitUser{UserID: lessors[j].UserID, Amount: debt})
			lessors[j].From = append(lessors[j].From, relatedSplitUser{UserID: lessees[i].UserID, Amount: debt})
			lessees[i].Calc -= debt
			lessors[j].Calc -= debt
			j++
			continue
		}
		if lessees[i].Calc == lessors[j].Calc {
			debt := lessees[i].Calc
			lessees[i].To = append(lessees[i].To, relatedSplitUser{UserID: lessors[j].UserID, Amount: debt})
			lessors[j].From = append(lessors[j].From, relatedSplitUser{UserID: lessees[i].UserID, Amount: debt})
			lessees[i].Calc -= debt
			lessors[j].Calc -= debt
			j++
			i++
		}
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"split": gin.H{
			"id":         split.ID,
			"title":      split.Title,
			"amount":     split.Amount,
			"average":    ave,
			"remainder":  remainder,
			"done":       split.Done,
			"group_id":   split.GroupID,
			"created_at": split.CreatedAt,
			"updated_at": split.UpdatedAt,
			"deleted_at": split.DeletedAt,
		},
		"lessees": lessees,
		"lessors": lessors,
	}

	c.JSON(http.StatusOK, response)
}

func (sc *SplitController) PostSplit(c *gin.Context) {

	var getData getData

	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var amount uint = 0

	for _, lessor := range getData.Lessors {
		amount += lessor.Amount
	}

	ctx := context.Background()
	tx, err := sc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	splitRepository := repositories.SplitRepository{TX: &tx}
	split, err := splitRepository.Create(ctx, getData.Title, getData.GroupID, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	lessors := make([]models.SplitLessor, len(getData.Lessors))

	for i, getDataLessor := range getData.Lessors {
		userID, uuidError := uuid.Parse(getDataLessor.UserID)
		if uuidError != nil {
			err = uuidError
			break
		}
		lessors[i] = models.SplitLessor{
			UserID:  userID,
			Amount:  getDataLessor.Amount,
			SplitID: split.ID,
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	splitLessorRepository := repositories.SplitLessorRepository{TX: &tx}
	lessors, err = splitLessorRepository.Create(ctx, lessors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	splitLesseeRepository := repositories.SplitLesseeRepository{TX: &tx}
	lessees, err := splitLesseeRepository.Create(ctx, getData.Lessees, split.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	c.JSON(http.StatusOK, gin.H{"split": split, "lessors": lessors, "lessees": lessees})
}

func (sc *SplitController) PutSplit(c *gin.Context) {
	// path
	// required split_id

	// json
	// option title string
	// required lessees []object
	// required lessors []string

	// 1. Update split table
	// 2. Update lessors data
	// 3. Update lessees data

	// Initialization
	splitID := c.Param("split_id")

	var getData getData

	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	tx, err := sc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Update split table

	splitRepository := repositories.SplitRepository{TX: &tx}

	var total uint = 0

	for _, lessor := range getData.Lessors {
		total += lessor.Amount
	}

	split, err := splitRepository.Update(ctx, getData.Title, total, splitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 2. Create temporary table 'temp_lessors'

	splitLessorRepository := repositories.SplitLessorRepository{TX: &tx}

	tmpl := make([]models.SplitLessor, len(getData.Lessors))
	for i, lessor := range getData.Lessors {
		userUUID, err := uuid.Parse(lessor.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}
		splitUUID, err := uuid.Parse(splitID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}
		tmpl[i] = models.SplitLessor{UserID: userUUID, SplitID: splitUUID, Amount: lessor.Amount}
	}

	err = splitLessorRepository.Update(ctx, tmpl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 3. Update lessees data

	splitLesseeRepository := repositories.SplitLesseeRepository{TX: &tx}

	err = splitLesseeRepository.Update(ctx, getData.Lessees, splitID)
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
	resSplit := gin.H{
		"update_at": split.UpdatedAt,
		"amount":    split.Amount,
	}
	if split.Title != "" {
		resSplit["title"] = split.Title
	}

	c.Status(http.StatusOK)
}

func (sc *SplitController) PatchDoneSplit(c *gin.Context) {
	// path
	// split_id

	// 1. 'done' in splits table be changed true.

	// Initialization
	splitID := c.Param("split_id")

	ctx := context.Background()
	tx, err := sc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. 'done' in splits table be changed true.

	splitRepository := repositories.SplitRepository{TX: &tx}

	if err := splitRepository.UpdateDone(ctx, true, splitID); err != nil {
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

func (sc *SplitController) PatchDoingSplit(c *gin.Context) {
	// path
	// split_id

	// 1. 'done' in splits table be changed false.

	// Initialization
	splitID := c.Param("split_id")

	ctx := context.Background()
	tx, err := sc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. 'done' in splits table be changed false.

	splitRepository := repositories.SplitRepository{TX: &tx}

	if err := splitRepository.UpdateDone(ctx, false, splitID); err != nil {
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

func (sc *SplitController) DeleteSplit(c *gin.Context) {
	// path
	// split_id

	// 1. Delete the data in split_lessors where split_id in the table is equal to split_id of request path
	// 2. Delete the data in split_lessees where split_id in the table is equal to split_id of request path
	// 3. Delete the data in splits where split_id in the table is equal to split_id of request path

	// Initialization

	splitID := c.Param("split_id")

	ctx := context.Background()
	tx, err := sc.DB.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1. Delete the data in split_lessors where split_id in the table is equal to split_id of request path

	splitLesseeRepository := repositories.SplitLesseeRepository{TX: &tx}
	if err := splitLesseeRepository.Delete(ctx, splitID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 2. Delete the data in split_lessees where split_id in the table is equal to split_id of request path

	splitLessorRepository := repositories.SplitLessorRepository{TX: &tx}
	if err := splitLessorRepository.Delete(ctx, splitID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	// 3. Delete the data in splits where split_id in the table is equal to split_id of request path

	splitRepository := repositories.SplitRepository{TX: &tx}
	if err := splitRepository.Delete(ctx, splitID); err != nil {
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
