package sqlc

import (
	"context"
	"testing"

	"github.com/Q2key/go-bank/utils"
	"github.com/stretchr/testify/require"
)

func TestListAccounts(t *testing.T) {
	_, err := testQueries.ListAccounts(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
}

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    utils.MakeRandomName(),
		Balance:  utils.MakeRandomInt64(),
		Currency: utils.MakeRandonCurrency(),
	}

	res, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.Equal(t, args.Owner, res.Owner)
	require.Equal(t, args.Balance, res.Balance)
	require.Equal(t, args.Currency, res.Currency)

	return res
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
