package test

import (
	"context"
	"fmt"
	"simple-bank/ent"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createFakeUser(ctx context.Context) *ent.User {
	user := testDb.User.Create().
		SetID(util.RandomUsername()).SetPassword(util.RandomPassword()).SetFullName(fmt.Sprintf("%s %s", util.RandomUsername(), util.RandomUsername())).SetEmail(util.RandomEmail()).
		SaveX(ctx)
	return user

}

func createFakeAccountWithFakeUser(ctx context.Context) (*ent.Account, error) {
	user := createFakeUser(ctx)
	account, err := testDb.Account.Create().
		SetOwner(user.ID).
		SetBalance(util.RandomAmount()).
		SetCurrency(util.RandomCurrency()).
		Save(ctx)
	return account, err
}

func TestCreateAccount(t *testing.T) {
	ctx := context.Background()
	account, err := createFakeAccountWithFakeUser(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, account)

}

func TestGetAccounts(t *testing.T) {
	ctx := context.Background()
	accounts, err := testDb.Account.Query().All(ctx)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(accounts), 1)

}
