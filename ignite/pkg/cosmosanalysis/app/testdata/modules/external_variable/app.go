package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	abci "github.com/tendermint/tendermint/abci/types"

	"app/modules"
)

var ModuleBasics = module.NewBasicManager(modules.Basic...)

type Foo struct{}

func (Foo) Name() string {
	return "foo"
}

func (Foo) BeginBlocker(sdk.Context, abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return abci.ResponseBeginBlock{}
}

func (Foo) EndBlocker(sdk.Context, abci.RequestEndBlock) abci.ResponseEndBlock {
	return abci.ResponseEndBlock{}
}

func (Foo) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	_ = apiSvr.ClientCtx
}

func (Foo) GetKey(storeKey string) *storetypes.KVStoreKey { return nil }

func (Foo) TxConfig() client.TxConfig { return nil }
