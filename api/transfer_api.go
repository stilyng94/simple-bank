package api

import (
	"errors"
	"fmt"
	"net/http"
	"simple-bank/ent"
	"simple-bank/repository"

	"github.com/gin-gonic/gin"
)

func (server *Server) createTransfer(ctx *gin.Context) {

	var kwargs repository.CreateTransferDto
	if err := ctx.ShouldBindJSON(&kwargs); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, kwargs.FromAccountID, kwargs.Currency) {
		return
	}

	response, err := server.iTransferRepo.CreateTransfer(ctx, kwargs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (server *Server) validAccount(ctx *gin.Context, accountID string, currency string) bool {
	account, err := server.iAccountRepo.GetAccount(ctx, repository.GetAccountDto{ID: accountID})
	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}
