package baseapp

import (
	abci "github.com/tendermint/tendermint/abci/types"
)

type RequestInfoWrapper struct {
	abci.RequestInfo

	ChainID   int
	ChainType ChainType
}

type RequestCheckTxWrapper struct {
	abci.RequestCheckTx

	ChainID   int
	ChainType ChainType
}
type (
	ChainType int
)

const (
	ChainType_Cosmos ChainType = 1
)
