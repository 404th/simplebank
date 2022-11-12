package db

import (
	"context"
	"testing"
	"time"

	"github.com/404th/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, acc1, acc2 Account) Transfer {
	args := CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        util.RandomMoney(),
	}

	res, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, args.FromAccountID, res.FromAccountID)
	require.Equal(t, args.ToAccountID, res.ToAccountID)
	require.Equal(t, args.Amount, res.Amount)

	return res
}

func TestCreateTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	createRandomTransfer(t, acc1, acc2)
}

func TestGetTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	transfer := createRandomTransfer(t, acc1, acc2)

	actual_transfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actual_transfer)

	require.Equal(t, acc1.ID, actual_transfer.FromAccountID)
	require.Equal(t, acc2.ID, actual_transfer.ToAccountID)
	require.Equal(t, transfer.Amount, actual_transfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, actual_transfer.CreatedAt, time.Second)

	item, err := testQueries.GetTransfer(context.Background(), 12356)
	require.Error(t, err)
	require.Empty(t, item)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
	}

	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	actual_transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, actual_transfers, 5)

	for _, tr := range actual_transfers {
		require.NotEmpty(t, tr)
	}
}
