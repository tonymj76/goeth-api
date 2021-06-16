package hanler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	Models "github.com/tonymj76/goeth-api/models"
	Modules "github.com/tonymj76/goeth-api/modules"
)

type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get paramenter from url request
	vars := mux.Vars(r)
	module := vars["module"]

	// Get the query parameters from url request
	address := r.URL.Query().Get("address")
	hash := r.URL.Query().Get("hash")

	// Set our response header
	w.Header().Set("Content-Type", "application/json")

	// Handle each request using the module parameter

	switch module {
	case "latest-block":
		_block := Modules.GetLatestBlock(*client.Client)
		json.NewEncoder(w).Encode(_block)
	case "get-tx":
		if hash == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		txHash := common.HexToHash(hash)
		_tx := Modules.GetTxByHash(*client.Client, txHash)
		if _tx != nil {
			json.NewEncoder(w).Encode(_tx)
			return
		}
		json.NewEncoder(w).Encode(&Models.Error{
			Code:    404,
			Message: "Tx Not Found!",
		})
	case "send-eth":
		decoder := json.NewDecoder(r.Body)
		var t Models.TransferEthRequest
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_hash, err := Modules.TransferEth(*client.Client, t.PrivKey, t.To, t.Amount)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}
		json.NewEncoder(w).Encode(&Models.HashResponse{
			Hash: _hash,
		})
	case "get-balance":
		if address == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		balance, err := Modules.GetAddressBalance(*client.Client, address)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}
		json.NewEncoder(w).Encode(&Models.BalanceResponse{
			Address: address,
			Balance: balance,
			Symbol:  "Ether",
			Units:   "Wei",
		})
	}
}