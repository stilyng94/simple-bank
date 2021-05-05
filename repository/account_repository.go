package repository

import (
	"context"
	"simple-bank/ent"

	"github.com/google/uuid"
)

type CreateAccountDto struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency"  binding:"required,currency"`
}

type GetAccountDto struct {
	ID string `uri:"id" binding:"required"`
}

type GetAccountsDto struct {
	PageID    int32  `form:"page_id" binding:"required,min=1"`
	PageSize  int32  `form:"page_size" binding:"required,min=5,max=10"`
	Direction string `form:"direction" binding:"required,oneof=asc desc"`
}

type IAccountRepo interface {
	CreateAccount(ctx context.Context, args CreateAccountDto) (account *ent.Account, err error)
	GetAccount(ctx context.Context, args GetAccountDto) (account *ent.Account, err error)
	GetAccounts(ctx context.Context, args GetAccountsDto) (accounts []*ent.Account, err error)
}

type AccountRepo struct {
	dbClient *ent.Client
}

func NewAccountRepo(dbClient *ent.Client) IAccountRepo {
	return &AccountRepo{dbClient: dbClient}
}

func (accountRepo *AccountRepo) CreateAccount(ctx context.Context, args CreateAccountDto) (account *ent.Account, err error) {
	account, err = accountRepo.dbClient.Account.Create().SetOwner((args.Owner)).SetCurrency(args.Currency).Save(ctx)
	return
}

func (accountRepo *AccountRepo) GetAccounts(ctx context.Context, args GetAccountsDto) (accounts []*ent.Account, err error) {
	orderBy := ent.Desc()
	if args.Direction == "asc" {
		orderBy = ent.Asc()
	}
	accounts, err = accountRepo.dbClient.Account.Query().Limit(int(args.PageSize)).Offset(int(args.PageSize) * int(args.PageID-1)).Order(orderBy).All(ctx)
	return
}

func (accountRepo *AccountRepo) GetAccount(ctx context.Context, args GetAccountDto) (account *ent.Account, err error) {
	account, err = accountRepo.dbClient.Account.Get(ctx, uuid.MustParse(args.ID))
	return
}
