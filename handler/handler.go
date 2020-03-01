package handler

import (
	"github.com/gin-gonic/gin"
	"google-service/models"
	"google-service/repository"
	"google-service/repository/orm"
	"net/http"
)

func GoogleHandler() *Google {
	return &Google{
		repo: orm.GoogleRepo(),
	}
}

type Google struct {
	repo repository.GoogleRepo
}

// Copy Google Table godoc
// @Summary Copy Google Table
// @Description Copy Google Table
// @Tags Spreadsheet
// @Param data body models.Spreadsheet true "Fields are required"
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPSuccess
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /spreadsheet/copy [post]
func (g *Google) CopyTable(ctx *gin.Context) {
	var spreadsheet *models.Spreadsheet

	err := ctx.BindJSON(&spreadsheet)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := g.repo.CopyTable(ctx, spreadsheet)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, res)
}

// Create Google Table godoc
// @Summary Create Google Table
// @Description Create Google Table
// @Tags Spreadsheet
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPSuccess
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /spreadsheet [post]
func (g *Google) CreateTable(ctx *gin.Context) {
	res, err := g.repo.CreateTable(ctx)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, res)
}

// Append Data godoc
// @Summary Append Data
// @Description Append Data
// @Tags Spreadsheet
// @Param data body models.Spreadsheet true "Fields are required"
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPSuccess
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /spreadsheet/append [post]
func (g *Google) AppendData(ctx *gin.Context) {
	var spreadsheet *models.Spreadsheet

	err := ctx.BindJSON(&spreadsheet)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	err = g.repo.AppendData(ctx, spreadsheet)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, map[string]interface{}{"message": "Success append"})
}

// Clear Data godoc
// @Summary Clear Data
// @Description Clear Data
// @Tags Spreadsheet
// @Param data body models.SpreadsheetClear true "Fields are required"
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPSuccess
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /spreadsheet/clear [post]
func (g *Google) ClearData(ctx *gin.Context) {
	var spreadsheet *models.SpreadsheetClear

	err := ctx.BindJSON(&spreadsheet)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	err = g.repo.ClearData(ctx, spreadsheet)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, map[string]interface{}{"message": "Success updated"})
}
