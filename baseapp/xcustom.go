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

// Info return application info
func (a *Application) Info(requestInfo RequestInfoWrapper) abciTypes.ResponseInfo {
	fmt.Println("=======================Info")
	fmt.Println("=======================Info")
	fmt.Println("=======================Info")
	fmt.Println("=======================Info")
	fmt.Println("=======================Info")
	fmt.Println("=======================Info")
	fmt.Println("=======================Info")
	fmt.Println("=======================Info")
	fmt.Println("=======================Info")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.Info(requestInfo.RequestInfo)
	default:
	}

	return abciTypes.ResponseInfo{
		Data:             "",
		Version:          "",
		AppVersion:       0,
		LastBlockHeight:  0,
		LastBlockAppHash: []byte{},
	}
}

// CheckTx validate a tx for the mempool
func (a *Application) CheckTx(requestInfo RequestCheckTxWrapper) abciTypes.ResponseCheckTx {
	fmt.Println("=======================CheckTx")
	fmt.Println("=======================CheckTx")
	fmt.Println("=======================CheckTx")
	fmt.Println("=======================CheckTx")
	fmt.Println("=======================CheckTx")
	fmt.Println("=======================CheckTx")
	fmt.Println("=======================CheckTx")
	fmt.Println("=======================CheckTx")
	fmt.Println("=======================CheckTx")
	switch requestInfo.ChainType {
	case ChainType_Cosmos:
		return a.BaseApp.CheckTx(requestInfo.RequestCheckTx)
	default:
	}

	return abciTypes.ResponseCheckTx{}
}
