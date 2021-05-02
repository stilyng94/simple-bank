package repository

import (
	"context"
	"simple-bank/ent"

	"github.com/google/uuid"
)

type CreateTransferArgs struct {
	FromAccountID string `json:"from_account_id"`
	ToAccountID   string `json:"to_account_id"`
	Amount        int32  `json:"amount"`
}

type CreateTransferResult struct {
	Transfer    *ent.Transfer `json:"transfer"`
	FromAccount *ent.Account  `json:"to_account_id"`
	ToAccount   *ent.Account  `json:"amount"`
	FromEntry   *ent.Entry    `json:"from_entry"`
	ToEntry     *ent.Entry    `json:"to_entry"`
}

type ITransferRepo interface {
	CreateTransfer(ctx context.Context, args CreateTransferArgs) (createTransferResult CreateTransferResult, err error)
}

type TransferRepo struct {
	dbClient *ent.Client
}

func NewTransferRepository(dbClient *ent.Client) *TransferRepo {
	return &TransferRepo{dbClient: dbClient}
}

func (transferRepo *TransferRepo) CreateTransfer(ctx context.Context, args CreateTransferArgs) (createTransferResult CreateTransferResult, err error) {

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

	if args.FromAccountID < args.ToAccountID {
		createTransferResult.FromAccount, createTransferResult.ToAccount, err = addMoney(ctx, txClient, args.FromAccountID, -args.Amount, args.ToAccountID, args.Amount)
		if err != nil {
			return
		}
	} else {
		createTransferResult.ToAccount, createTransferResult.FromAccount, err = addMoney(ctx, txClient, args.ToAccountID, args.Amount, args.FromAccountID, -args.Amount)
		if err != nil {
			return
		}
	}

	err = txClient.Commit()

	return

}

func addMoney(ctx context.Context, txClient *ent.Tx, account1ID string, amount1 int32, account2ID string, amount2 int32) (account1 *ent.Account, account2 *ent.Account, err error) {
	account1, err = txClient.Account.UpdateOneID(uuid.MustParse(account1ID)).AddBalance(amount1).Save(ctx)
	if err != nil {
		return
	}
	account2, err = txClient.Account.UpdateOneID(uuid.MustParse(account2ID)).AddBalance(amount2).Save(ctx)
	if err != nil {
		return
	}
	return
}
