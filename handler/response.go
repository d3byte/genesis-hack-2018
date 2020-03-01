package handler

import (
	"github.com/gin-gonic/gin"
	"google-service/models"
)

// respondWithJSON write json response format
func RespondWithJSON(ctx *gin.Context, status int, payload interface{}) {
	su := models.HTTPSuccess{
		Code:    status,
		Data:    payload,
	}

	ctx.JSON(status, su)
}

func RespondWithError(ctx *gin.Context, status int, err error) {
	er := models.HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	ctx.JSON(status, er)
}
