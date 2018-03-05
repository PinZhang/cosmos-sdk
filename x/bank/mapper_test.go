package bank

import (
	"fmt"
	"os"
	"testing"

	crypto "github.com/tendermint/go-crypto"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/examples/basecoin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	//"github.com/stretchr/testify/assert"
)

func defaultLogger() log.Logger {
	return log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "sdk/app")
}

func newBaseApp(name string) *baseapp.BaseApp {
	logger := defaultLogger()
	db := dbm.NewMemDB()
	return baseapp.NewBaseApp(name, logger, db)
}

func TestNewCoinKeeper(t *testing.T) {

}

/*
func TestSubtractCoins(t *testing.T) {
	bapp := newBaseApp()

	//accMapper := ?
	coinKeeper := NewCoinKeeper(accMapper)

	ctx := bapp.BaseApp.NewContext(true, abci.Header{})

	coins, err := coinKeeper.SubtractCoins(ctx, addr, amt)
	assert.NotNil(t, err)
	assert.Equal(t, coins, something)

}

func TestAddCoins(t *testing.T) {

	coins, err := coinKeeper.AddCoins(ctx, addr, amt)
	assert.NotNil(t, err)
	assert.Equal(t, coins, something)
}*/
