package db

import (
	"context"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func NewRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	NewRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// Create new account
	account1 := NewRandomAccount(t)
	// Get that account with id
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	// Check equal
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	// require.Equal(t, account1.Owner, account2.Owner)
	// require.Equal(t, account1.Balance, account2.Balance)
	// require.Equal(t, account1.Currency, account2.Currency)

	// require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	// require.Equal(t, account1.ID, account2.ID)
	// Faster way
	if diff := cmp.Diff(account1, account2); diff != "" {
		t.Errorf("accounts mismatch (-account1 +account2):\n%s", diff)
	}
}

func TestGetListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 1; i <= 10; i++ {
		lastAccount = NewRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Offset: 0,
		Limit:  5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func TestUpdateAccount(t *testing.T) {
	account1 := NewRandomAccount(t)

	arg := UpdateAccountParams{
		Balance: util.RandomMoney(),
		ID:      account1.ID,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)

	require.Equal(t, arg.Balance, account2.Balance)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := NewRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	// require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, account2)

}
