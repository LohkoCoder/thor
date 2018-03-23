package api

import (
	"github.com/gorilla/mux"
	"github.com/vechain/thor/chain"
	"github.com/vechain/thor/logdb"
	"github.com/vechain/thor/state"
	"github.com/vechain/thor/txpool"
)

//NewHTTPHandler return api router
func NewHTTPHandler(chain *chain.Chain, stateCreator *state.Creator, txPool *txpool.TxPool, ldb *logdb.LogDB) *mux.Router {
	router := mux.NewRouter()
	NewAccountHTTPRouter(router, NewAccountInterface(chain, stateCreator))
	NewTransactionHTTPRouter(router, NewTransactionInterface(chain, txPool))
	NewBlockHTTPRouter(router, NewBlockInterface(chain))
	NewLogHTTPRouter(router, NewLogInterface(ldb))
	NewContractHTTPRouter(router, NewContractInterface(chain, stateCreator))
	return router
}
