package baseapp

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"
)

// InitChain implements the ABCI interface. It runs the initialization logic
// directly on the CommitMultiStore.
func (app *BaseApp) InitChain(req abci.RequestInitChain) (res abci.ResponseInitChain) {
	fmt.Printf("BaseApp.InitChain %+v\n", req)
	return app.parent.InitChain(RequestInitChainWrapper{
		RequestInitChain: req,
		ChainID:          1,
		ChainType:        ChainType_Cosmos,
	})
}

// Info implements the ABCI interface.
func (app *BaseApp) Info(req abci.RequestInfo) abci.ResponseInfo {
	return app.parent.Info(RequestInfoWrapper{
		RequestInfo: req,
		ChainID:     1,
		ChainType:   ChainType_Cosmos,
	})
}

// SetOption implements the ABCI interface.
func (app *BaseApp) SetOption(req abci.RequestSetOption) (res abci.ResponseSetOption) {
	return app.parent.SetOption(RequestSetOptionWrapper{
		RequestSetOption: req,
		ChainID:          1,
		ChainType:        ChainType_Cosmos,
	})
}

// BeginBlock implements the ABCI application interface.
func (app *BaseApp) BeginBlock(req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	return app.parent.BeginBlock(RequestBeginBlockWrapper{
		RequestBeginBlock: req,
		ChainID:           1,
		ChainType:         ChainType_Cosmos,
	})
}

// EndBlock implements the ABCI interface.
func (app *BaseApp) EndBlock(req abci.RequestEndBlock) (res abci.ResponseEndBlock) {
	return app.parent.EndBlock(RequestEndBlockWrapper{
		RequestEndBlock: req,
		ChainID:         1,
		ChainType:       ChainType_Cosmos,
	})
}

// checkTx implements the ABCI interface and executes a tx in CheckTx mode. In
// CheckTx mode, messages are not executed. This means messages are only validated
// and only the AnteHandler is executed. State is persisted to the BaseApp's
// internal CheckTx state if the AnteHandler passes. Otherwise, the ResponseCheckTx
// will contain releveant error information. Regardless of tx execution outcome,
// the ResponseCheckTx will contain relevant gas execution context.
func (app *BaseApp) CheckTx(req abci.RequestCheckTx) abci.ResponseCheckTx {
	return app.parent.CheckTx(RequestCheckTxWrapper{
		RequestCheckTx: req,
		ChainID:        1,
		ChainType:      ChainType_Cosmos,
	})
}

// DeliverTx implements the ABCI interface and executes a tx in DeliverTx mode.
// State only gets persisted if all messages are valid and get executed successfully.
// Otherwise, the ResponseDeliverTx will contain releveant error information.
// Regardless of tx execution outcome, the ResponseDeliverTx will contain relevant
// gas execution context.
func (app *BaseApp) DeliverTx(req abci.RequestDeliverTx) abci.ResponseDeliverTx {
	return app.parent.DeliverTx(RequestDeliverTxWrapper{
		RequestDeliverTx: req,
		ChainID:          1,
		ChainType:        ChainType_Cosmos,
	})
}

// Commit implements the ABCI interface. It will commit all state that exists in
// the deliver state's multi-store and includes the resulting commit ID in the
// returned abci.ResponseCommit. Commit will set the check state based on the
// latest header and reset the deliver state. Also, if a non-zero halt height is
// defined in config, Commit will execute a deferred function call to check
// against that height and gracefully halt if it matches the latest committed
// height.
func (app *BaseApp) Commit() (res abci.ResponseCommit) {
	return app.parent.Commit(RequestCommitWrapper{
		ChainID:   1,
		ChainType: ChainType_Cosmos,
	})
}

// Query implements the ABCI interface. It delegates to CommitMultiStore if it
// implements Queryable.
func (app *BaseApp) Query(req abci.RequestQuery) (res abci.ResponseQuery) {
	return app.parent.Query(RequestQueryWrapper{
		RequestQuery: req,
		ChainID:      1,
		ChainType:    ChainType_Cosmos,
	})
}

// ListSnapshots implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) ListSnapshots(req abci.RequestListSnapshots) abci.ResponseListSnapshots {
	return app.parent.ListSnapshots(RequestListSnapshotsWrapper{
		RequestListSnapshots: req,
		ChainID:              1,
		ChainType:            ChainType_Cosmos,
	})
}

// LoadSnapshotChunk implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) LoadSnapshotChunk(req abci.RequestLoadSnapshotChunk) abci.ResponseLoadSnapshotChunk {
	return app.parent.LoadSnapshotChunk(RequestLoadSnapshotChunkWrapper{
		RequestLoadSnapshotChunk: req,
		ChainID:                  1,
		ChainType:                ChainType_Cosmos,
	})
}

// OfferSnapshot implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) OfferSnapshot(req abci.RequestOfferSnapshot) abci.ResponseOfferSnapshot {
	return app.parent.OfferSnapshot(RequestOfferSnapshotWrapper{
		RequestOfferSnapshot: req,
		ChainID:              1,
		ChainType:            ChainType_Cosmos,
	})
}

// ApplySnapshotChunk implements the ABCI interface. It delegates to app.snapshotManager if set.
func (app *BaseApp) ApplySnapshotChunk(req abci.RequestApplySnapshotChunk) abci.ResponseApplySnapshotChunk {
	return app.parent.ApplySnapshotChunk(RequestApplySnapshotChunkWrapper{
		RequestApplySnapshotChunk: req,
		ChainID:                   1,
		ChainType:                 ChainType_Cosmos,
	})

}
