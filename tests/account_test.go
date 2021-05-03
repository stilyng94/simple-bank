package tests

import (
	"context"
	"simple-bank/ent"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createFakeAccount(ctx context.Context) (*ent.Account, error) {
	account, err := testDb.Account.Create().SetOwner(util.RandomOwner()).SetBalance(util.RandomAmount()).SetCurrency(util.RandomCurrency()).Save(ctx)
	return account, err

}

func TestCreateAccount(t *testing.T) {
	ctx := context.Background()
	account, err := createFakeAccount(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, account)

}

func TestGetAccounts(t *testing.T) {
	ctx := context.Background()
	accounts, err := testDb.Account.Query().All(ctx)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(accounts), 5)

}
