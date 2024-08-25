package rest

import (
	"net/http"

	"github.com/co1seam/Taskify/service/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(ctx *gin.Context)  {
	var input models.User
	
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		`id`: id,
	})
}

func (h *Handler) SignIn(ctx *gin.Context) {
	
}