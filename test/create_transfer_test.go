package test

import (
	"context"
	"fmt"
	"simple-bank/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func createFakeTransfer(ctx context.Context, iTransferRepo repository.ITransferRepo, args repository.CreateTransferDto) (repository.CreateTransferResultDto, error) {
	result, err := iTransferRepo.CreateTransfer(ctx, args)
	return result, err

}
func TestCreateTransfer(t *testing.T) {
	ctx := context.Background()

	account1, _ := createFakeAccountWithFakeUser(ctx)
	account2, _ := createFakeAccountWithFakeUser(ctx)

	t.Log(">> before: ", account1.Balance, account2.Balance)

	//run a concurrent transfer transcations
	n := 5
	amount := float64(10)

	errChan := make(chan error)
	resultChan := make(chan repository.CreateTransferResultDto)

	for i := 0; i < n; i++ {
		go func() {
			ctx2 := context.Background()

			transferRepo := repository.NewTransferRepo(testDb)
			result, err := createFakeTransfer(ctx2, transferRepo, repository.CreateTransferDto{FromAccountID: account1.ID.String(), ToAccountID: account2.ID.String(), Amount: amount})
			errChan <- err
			resultChan <- result
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errChan
		result := <-resultChan

		require.NoError(t, err)
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreateTime)

		transferFromAccount := testDb.Transfer.QueryFromAccount(transfer).OnlyIDX(ctx)
		transferToAccount := testDb.Transfer.QueryToAccount(transfer).OnlyIDX(ctx)
		require.Equal(t, transferFromAccount, account1.ID)
		require.Equal(t, transferToAccount, account2.ID)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreateTime)

		_, err = testDb.Transfer.Get(ctx, transfer.ID)
		require.NoError(t, err)

		fromEntryAccount := testDb.Entry.QueryAccount(fromEntry).OnlyX(ctx)
		require.Equal(t, fromEntryAccount.ID, account1.ID)
		require.Equal(t, -amount, fromEntry.Amount)

		_, err = testDb.Entry.Get(ctx, fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreateTime)

		toEntryAccount := testDb.Entry.QueryAccount(toEntry).OnlyX(ctx)
		require.Equal(t, toEntryAccount.ID, account2.ID)
		require.Equal(t, amount, toEntry.Amount)

		_, err = testDb.Entry.Get(ctx, toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		// check accounts
		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check accounts balance
		fmt.Println(">> tx: ", fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, int(diff1)%int(amount) == 0.0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true

	}

	//check the final updated balance
	updateAccount1, err := testDb.Account.Get(ctx, account1.ID)
	require.NoError(t, err)
	updateAccount2, err := testDb.Account.Get(ctx, account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after: ", updateAccount1.Balance, updateAccount2.Balance)

	require.Equal(t, account1.Balance-float64(n)*amount, updateAccount1.Balance)
	require.Equal(t, account2.Balance+float64(n)*amount, updateAccount2.Balance)
}

func TestCreateTransferDeadlock(t *testing.T) {
	ctx := context.Background()

	account1, _ := createFakeAccountWithFakeUser(ctx)
	account2, _ := createFakeAccountWithFakeUser(ctx)

	t.Log(">> before: ", account1.Balance, account2.Balance)

	//run a concurrent transfer transcations
	n := 6
	amount := float64(10)

	errChan := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			ctx2 := context.Background()
			transferRepo := repository.NewTransferRepo(testDb)
			_, err := createFakeTransfer(ctx2, transferRepo, repository.CreateTransferDto{FromAccountID: fromAccountID.String(), ToAccountID: toAccountID.String(), Amount: amount})

			errChan <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errChan

		require.NoError(t, err)
	}

	//check the final updated balance
	updateAccount1, err := testDb.Account.Get(ctx, account1.ID)
	require.NoError(t, err)
	updateAccount2, err := testDb.Account.Get(ctx, account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after: ", updateAccount1.Balance, updateAccount2.Balance)

	require.Equal(t, account1.Balance, updateAccount1.Balance)
	require.Equal(t, account2.Balance, updateAccount2.Balance)
}
