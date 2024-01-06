package handler

import (
	"fmt"
	"net/http"

	"github.com/gabriel-O-C/godin/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "show opening", opening)
}
