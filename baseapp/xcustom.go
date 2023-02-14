package baseapp

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Application struct {
	abci.Application
	*BaseApp
}

// ----- Info/Query Connection

func (a *Application) InitChain(requestInfo RequestInitChainWrapper) (res abciTypes.ResponseInitChain) {
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.initChain(requestInfo.RequestInitChain)
	default:
	}

	return abciTypes.ResponseInitChain{}
}

// Info return application info
func (a *Application) Info(requestInfo RequestInfoWrapper) abciTypes.ResponseInfo {
	fmt.Println("====nampkh Info")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.info(requestInfo.RequestInfo)
	default:
	}

	return abciTypes.ResponseInfo{}
}

func (a *Application) SetOption(requestInfo RequestSetOptionWrapper) (res abciTypes.ResponseSetOption) {
	fmt.Println("====nampkh SetOption")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.setOption(requestInfo.RequestSetOption)
	default:
	}

	return abciTypes.ResponseSetOption{}
}

func (a *Application) BeginBlock(requestInfo RequestBeginBlockWrapper) (res abciTypes.ResponseBeginBlock) {
	fmt.Println("====nampkh BeginBlock")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.beginBlock(requestInfo.RequestBeginBlock)
	default:
	}

	return abciTypes.ResponseBeginBlock{}
}

func (a *Application) EndBlock(requestInfo RequestEndBlockWrapper) (res abciTypes.ResponseEndBlock) {
	fmt.Println("====nampkh EndBlock")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.endBlock(requestInfo.RequestEndBlock)
	default:
	}

	return abciTypes.ResponseEndBlock{}
}

// CheckTx validate a tx for the mempool
func (a *Application) CheckTx(requestInfo RequestCheckTxWrapper) abciTypes.ResponseCheckTx {
	fmt.Printf("====nampkh CheckTx: %+v\n", string(requestInfo.RequestCheckTx.GetTx()))
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.checkTx(requestInfo.RequestCheckTx)
	default:
	}

	return abciTypes.ResponseCheckTx{}
}

func (a *Application) DeliverTx(requestInfo RequestDeliverTxWrapper) abciTypes.ResponseDeliverTx {
	fmt.Printf("====nampkh DeliverTx: %+v\n", requestInfo.RequestDeliverTx)
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.deliverTx(requestInfo.RequestDeliverTx)
	default:
	}

	return abciTypes.ResponseDeliverTx{}
}

func (a *Application) Commit(requestInfo RequestCommitWrapper) abciTypes.ResponseCommit {
	fmt.Printf("====nampkh Commit%+v\n", requestInfo.RequestCommit)
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.commit()
	default:
	}

	return abciTypes.ResponseCommit{}
}

func (a *Application) Query(requestInfo RequestQueryWrapper) abciTypes.ResponseQuery {
	fmt.Printf("====nampkh Query: %+v\n", string(requestInfo.RequestQuery.GetData()))
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.query(requestInfo.RequestQuery)
	default:
	}

	return abciTypes.ResponseQuery{}
}

func (a *Application) ListSnapshots(requestInfo RequestListSnapshotsWrapper) abciTypes.ResponseListSnapshots {
	fmt.Println("====nampkh ListSnapshots")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.listSnapshots(requestInfo.RequestListSnapshots)
	default:
	}

	return abciTypes.ResponseListSnapshots{}
}

func (a *Application) LoadSnapshotChunk(requestInfo RequestLoadSnapshotChunkWrapper) abciTypes.ResponseLoadSnapshotChunk {
	fmt.Println("====nampkh LoadSnapshotChunk")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.loadSnapshotChunk(requestInfo.RequestLoadSnapshotChunk)
	default:
	}

	return abciTypes.ResponseLoadSnapshotChunk{}
}

func (a *Application) OfferSnapshot(requestInfo RequestOfferSnapshotWrapper) abciTypes.ResponseOfferSnapshot {
	fmt.Println("====nampkh OfferSnapshot")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.offerSnapshot(requestInfo.RequestOfferSnapshot)
	default:
	}

	return abciTypes.ResponseOfferSnapshot{}
}

func (a *Application) ApplySnapshotChunk(requestInfo RequestApplySnapshotChunkWrapper) abciTypes.ResponseApplySnapshotChunk {
	fmt.Println("====nampkh ApplySnapshotChunk")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.applySnapshotChunk(requestInfo.RequestApplySnapshotChunk)
	default:
	}

	return abciTypes.ResponseApplySnapshotChunk{}
}
