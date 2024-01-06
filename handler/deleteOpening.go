package handler

import (
	"fmt"
	"net/http"

	"github.com/gabriel-O-C/godin/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query parameter").Error())
		return
	}

	opening := schemas.Opening{}

	err := db.First(&opening, id).Error
	if err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error on deleting opening")
		return
	}

	sendSuccess(ctx, "delete opening", opening)
}
