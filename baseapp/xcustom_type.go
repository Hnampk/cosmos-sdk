package baseapp

import (
	abci "github.com/tendermint/tendermint/abci/types"
)

type (
	ChainType int
)

const (
	ChainType_Cosmos ChainType = 1
)

type RequestInitChainWrapper struct {
	abci.RequestInitChain

	ChainID   int
	ChainType ChainType
}

type RequestInfoWrapper struct {
	abci.RequestInfo

	ChainID   int
	ChainType ChainType
}
type RequestSetOptionWrapper struct {
	abci.RequestSetOption

	ChainID   int
	ChainType ChainType
}

type RequestBeginBlockWrapper struct {
	abci.RequestBeginBlock

	ChainID   int
	ChainType ChainType
}

type RequestEndBlockWrapper struct {
	abci.RequestEndBlock

	ChainID   int
	ChainType ChainType
}

type RequestCheckTxWrapper struct {
	abci.RequestCheckTx

	ChainID   int
	ChainType ChainType
}

type RequestDeliverTxWrapper struct {
	abci.RequestDeliverTx

	ChainID   int
	ChainType ChainType
}

type RequestCommitWrapper struct {
	abci.RequestCommit

	ChainID   int
	ChainType ChainType
}

type RequestQueryWrapper struct {
	abci.RequestQuery

	ChainID   int
	ChainType ChainType
}

type RequestListSnapshotsWrapper struct {
	abci.RequestListSnapshots

	ChainID   int
	ChainType ChainType
}

type RequestLoadSnapshotChunkWrapper struct {
	abci.RequestLoadSnapshotChunk

	ChainID   int
	ChainType ChainType
}

type RequestOfferSnapshotWrapper struct {
	abci.RequestOfferSnapshot

	ChainID   int
	ChainType ChainType
}
type RequestApplySnapshotChunkWrapper struct {
	abci.RequestApplySnapshotChunk

	ChainID   int
	ChainType ChainType
}
