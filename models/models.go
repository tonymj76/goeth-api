package models

type (
	// Block data structure
	Block struct {
		BlockNumber      int64         `json:"blockNumber"`
		Timestamp        uint64        `json:"timestamp"`
		Difficulty       uint64        `json:"diffiulty"`
		Hash             string        `json:"hash"`
		TransactionCount int           `json:"transactionCount"`
		Transactions     []Transaction `json:"transactions"`
	}

	// Transaction data struction
	Transaction struct {
		Hash     string `json:"hash"`
		Value    string `json:"value"`
		Gas      uint64 `json:"gas"`
		GasPrice uint64 `json:"gasPrice"`
		Nonce    uint64 `json:"nonce"`
		To       string `json:"to"`
		Pending  bool   `json:"pending"`
	}

	// TransferEthRequest data structure
	TransferEthRequest struct {
		PrivKey string `json:"privKey"`
		To      string `json:"to"`
		Amount  int64  `json:"amount"`
	}

	// HashResponse data structure
	HashResponse struct {
		Hash string `json:"hash"`
	}

	// BalanceResponse data structure
	BalanceResponse struct {
		Address string `json:"address"`
		Balance string `json:"balance"`
		Symbol  string `json:"symbol"`
		Units   string `json:"units"`
	}

	// Error data structure
	Error struct {
		Code    uint64 `json:"code"`
		Message string `json:"message"`
	}
)
