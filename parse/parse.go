package parse

import (
	"fmt"

	logger "CollectNFTDataKlaytn/logger"
	"github.com/klaytn/klaytn/blockchain/types"

	erc721 "CollectNFTDataKlaytn/token/ERC721"
)

func ParseReceipt(rept *types.Receipt) {

	fmt.Println("ParseReceipt")

	for _, m := range rept.Logs {

		contractAddr := m.Address.Hex()

		logger.InfoLog("--ContractAddr ContractAddr[%s] , m.Topics[0][%s] \n", contractAddr, m.Topics[0].Hex())

		instance, err := erc721.NewErc721(m.Address, client)
		if err != nil {
			logger.InfoLog("-------getDataERC721 NewErc721 contractAddressHex[%s] , error[%s] ", m.Address.Hex(), err.Error())
			return
		}

	}
}
