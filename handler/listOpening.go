package handler

import (
	"net/http"

	"github.com/gabriel-O-C/godin/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	err := db.Find(&openings).Error
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "List openings", openings)
}
