package types

import (
	"github.com/klaytn/klaytn/common"
)

type TokenInfo struct {
	TransactionHash  common.Hash
	BlockTime        string
	Contractaddress  common.Address
	ContractName     string
	ContractSymbol   string
	TokenID          string
	WrapTokenAddress common.Address
	WrapTokenName    string
	WrapTokenSymbol  string
	KlayValue        string
}
