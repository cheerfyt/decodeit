package pkg

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeEthRawTx(rawTx string) (*types.Transaction, error) {
	tx := new(types.Transaction)
	err := tx.UnmarshalBinary(common.FromHex(rawTx))
	if err != nil {
		return nil, err
	}
	return tx, err
}

func GetFromAddress(tx *types.Transaction) string {
	msg, _ := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), big.NewInt(0))
	return strings.ToLower(msg.From().Hex())
}
