package api

import (
	"net/http"
	"simple-bank/repository"

	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(ctx *gin.Context) {
	var kwargs repository.CreateUserDto
	if err := ctx.ShouldBindJSON(&kwargs); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.iUserRepo.CreateUser(ctx, kwargs)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, user)

}

func (server *Server) getUser(ctx *gin.Context) {
	var kwargs repository.GetUserDto
	if err := ctx.ShouldBindUri(&kwargs); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.iUserRepo.GetUser(ctx, kwargs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)

}
