package db

import (
	"context"
	"testing"
	"time"

	"github.com/404th/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	ent := createRandomEntry(t, account)

	entry, err := testQueries.GetEntry(context.Background(), ent.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, ent.AccountID, entry.AccountID)
	require.Equal(t, ent.Amount, entry.Amount)
	require.Equal(t, ent.ID, entry.ID)
	require.WithinDuration(t, ent.CreatedAt, entry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	args := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)

	for _, ent := range entries {
		require.NotEmpty(t, ent)
		require.Equal(t, args.AccountID, ent.AccountID)
	}
}

func createRandomEntry(t *testing.T, account Account) Entry {
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	res, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, res.AccountID, args.AccountID)
	require.Equal(t, res.Amount, args.Amount)

	require.NotZero(t, res.CreatedAt)
	require.NotZero(t, res.ID)

	return res
}
