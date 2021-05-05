package repository

import (
	"context"
	"simple-bank/ent"

	"github.com/google/uuid"
)

type CreateTransferDto struct {
	FromAccountID string  `json:"from_account_id" binding:"required"`
	ToAccountID   string  `json:"to_account_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0.0"`
	Currency      string  `json:"currency"  binding:"required,currency"`
}

type CreateTransferResultDto struct {
	Transfer    *ent.Transfer `json:"transfer"`
	FromAccount *ent.Account  `json:"from_account"`
	ToAccount   *ent.Account  `json:"to_account"`
	FromEntry   *ent.Entry    `json:"from_entry"`
	ToEntry     *ent.Entry    `json:"to_entry"`
}

type ITransferRepo interface {
	CreateTransfer(ctx context.Context, args CreateTransferDto) (createTransferResult CreateTransferResultDto, err error)
}

type TransferRepo struct {
	dbClient *ent.Client
}

func NewTransferRepo(dbClient *ent.Client) ITransferRepo {
	return &TransferRepo{dbClient: dbClient}
}

func (transferRepo *TransferRepo) CreateTransfer(ctx context.Context, args CreateTransferDto) (createTransferResult CreateTransferResultDto, err error) {

	txClient, err := transferRepo.dbClient.Tx(ctx)
	if err != nil {
		return
	}

	tc := txClient.Transfer.Create()
	tc.SetFromAccountID(uuid.MustParse(args.FromAccountID))
	tc.SetToAccountID(uuid.MustParse(args.ToAccountID))
	tc.SetAmount(args.Amount)
	createTransferResult.Transfer, err = tc.Save(ctx)
	if err != nil {
		return
	}

	ecF := txClient.Client().Entry.Create()
	ecF.SetAccountID(uuid.MustParse(args.FromAccountID))
	ecF.SetAmount(-args.Amount)
	createTransferResult.FromEntry, err = ecF.Save(ctx)
	if err != nil {
		return
	}

	ecT := txClient.Client().Entry.Create()
	ecT.SetAccountID(uuid.MustParse(args.ToAccountID))
	ecT.SetAmount(args.Amount)
	createTransferResult.ToEntry, err = ecT.Save(ctx)
	if err != nil {
		return
	}

	//TODO: Lock for update balance

	if uuid.MustParse(args.FromAccountID).String() < uuid.MustParse(args.ToAccountID).String() {
		createTransferResult.FromAccount, createTransferResult.ToAccount, err = addMoney(ctx, txClient, uuid.MustParse(args.FromAccountID), -args.Amount, uuid.MustParse(args.ToAccountID), args.Amount)
		if err != nil {
			return
		}
	} else {
		createTransferResult.ToAccount, createTransferResult.FromAccount, err = addMoney(ctx, txClient, uuid.MustParse(args.ToAccountID), args.Amount, uuid.MustParse(args.FromAccountID), -args.Amount)
		if err != nil {
			return
		}
	}

	err = txClient.Commit()

	return

}

func addMoney(ctx context.Context, txClient *ent.Tx, account1ID uuid.UUID, amount1 float64, account2ID uuid.UUID, amount2 float64) (account1 *ent.Account, account2 *ent.Account, err error) {
	account1, err = txClient.Account.UpdateOneID(account1ID).AddBalance(amount1).Save(ctx)
	if err != nil {
		return
	}
	account2, err = txClient.Account.UpdateOneID(account2ID).AddBalance(amount2).Save(ctx)
	if err != nil {
		return
	}
	return
}
