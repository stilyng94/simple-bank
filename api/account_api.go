package api

import (
	"net/http"
	"simple-bank/repository"

	"github.com/gin-gonic/gin"
)

func (server *Server) createAccount(ctx *gin.Context) {

	var kwargs repository.CreateAccountDto
	if err := ctx.ShouldBindJSON(&kwargs); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.iAccountRepo.CreateAccount(ctx, kwargs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, account)

}

func (server *Server) getAccount(ctx *gin.Context) {
	var kwargs repository.GetAccountDto

	if err := ctx.ShouldBindUri(&kwargs); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.iAccountRepo.GetAccount(ctx, kwargs)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)

}

func (server *Server) getAccounts(ctx *gin.Context) {
	var kwargs repository.GetAccountsDto

	if err := ctx.ShouldBindQuery(&kwargs); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.iAccountRepo.GetAccounts(ctx, kwargs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)

}
