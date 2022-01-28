package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"runtime"
	"strconv"
	"time"

	"CollectNFTDataKlaytn/config"
	"CollectNFTDataKlaytn/kas"
	. "CollectNFTDataKlaytn/types"

	//"CollectNFTDataKlaytn/parse"

	kip17 "CollectNFTDataKlaytn/contract/KIP17"
	kip7 "CollectNFTDataKlaytn/contract/KIP7"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	klayClient "github.com/klaytn/klaytn/client"

	"github.com/klaytn/klaytn/common"

	//"github.com/klaytn/klaytn/accounts/abi/bind"

	logger "CollectNFTDataKlaytn/logger"
)

var klaytndial *klayClient.Client = nil

var IMAGE_PATH string = "../CollectNFT/ImagesKlaytn/"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("GOMAXPROCS : ", runtime.GOMAXPROCS(0))

	logger.LoggerInit()

	fromNum := flag.Int64("fromblock", 0, "FromBlockNumber")
	toNum := flag.Int64("toblock", 0, "ToBlockNumber")

	flag.Parse()

	configData, err := config.LoadConfigration("config.json")
	if err != nil {
		log.Fatal("LoadConfigration :", err)
	}

	ksconfig := kas.Config{}

	ksconfig.AccessKeyID = configData.AccessKeyID
	ksconfig.SecretAccessKey = configData.SecretAccessKey
	ksconfig.Endpoint = configData.Endpoint
	ksconfig.Network = configData.Network

	k, err := kas.Dial(ksconfig)
	if err != nil {
		log.Fatal("Dial : ", err)
	}

	klaytndial = k

	chainId, err := klaytndial.ChainID(context.Background())
	if err != nil {
		log.Fatal("ChainID error : ", err)
	}
	fmt.Println("ChainID : ", chainId.Int64())

	// klaytn block list
	// klaytn lastest block num 80292389

	// block number 73967410 (Nov 01, 2021 00:00:00 / UTC+9)
	// block number 73967409 (Oct 31, 2021 23:59:59 / UTC+9)
	// block number 71291139 (Oct 01, 2021 00:00:00 / UTC+9)
	// block number 71291138 (Sep 30, 2021 23:59:59 / UTC+9)

	var fromBlockNumber int64 = 71291139 // 80520318 // 71291139
	var toBlockNumber int64 = 73967409   // 80520318   // 71299139

	if *fromNum != 0 {
		fromBlockNumber = *fromNum
		toBlockNumber = *toNum
	}

	logger.InfoLog("-----Start fromBlockNumber :  %d , toBlockNumber : %d", fromBlockNumber, toBlockNumber)

	for fromBlockNumber < toBlockNumber {

		divideToBlockNumber := fromBlockNumber + 1000
		CollectTrxProcess(fromBlockNumber, divideToBlockNumber)
		fromBlockNumber = divideToBlockNumber + 1

	}

}

func CollectTrxProcess(fromBlockNumber, toBlockNumber int64) {

	var minKlayValue int = 90000000000 // 100 klay 100000000000000000000 인데 0 10개 뺀다 10000000000 , 100000000000 1000klay
	//2149000000000000000000
	//100000000000000000000

	address := "0xfd844c2fca5e595004b17615f891620d1cb9bbb2"

	wklayContractAddress := common.HexToAddress(address)

	transferHash := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") // transfer topic[0]

	logger.InfoLog("-----Start CollectTrxProcess filterQuery fromBlockNumber[%d] , toBlockNumber[%d]", fromBlockNumber, toBlockNumber)

	query := klaytn.FilterQuery{
		FromBlock: big.NewInt(fromBlockNumber),
		ToBlock:   big.NewInt(toBlockNumber),
		Addresses: []common.Address{
			wklayContractAddress,
		},
		Topics: [][]common.Hash{
			{transferHash},
		},
	}

	logs, err := klaytndial.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("FilterLogs Count : ", len(logs))

	for _, m := range logs { // address wklay log

		//0xfd844c2fca5e595004b17615f891620d1cb9bbb2 wklay address
		if m.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" { //wklay contract의 transfer
			//	logger.InfoLog("------Event Log Tx[%s] , ContractAddr[%s]\n", m.TxHash.Hex(), m.Address.Hex())

			// klay := m.Topics[3].Big()

			// fmt.Println("klay.Int64()", klay.Int64())
			// logger.InfoLog("----Log Topicvalue[%d]\n", klay.Int64())

			wrapTokenAddress := m.Address // wklay 혹은 다른 token

			instance, err := kip7.NewKip7(wrapTokenAddress, klaytndial)
			if err != nil {
				logger.InfoLog("----Error NewKip7 TxHash[%s] Error[%s]\n", m.TxHash.Hex(), err.Error())
				continue
			}

			name, err := instance.Name(&bind.CallOpts{})
			if err != nil {
				logger.InfoLog("----Error Name TxHash[%s] Error[%s]\n", m.TxHash.Hex(), err.Error())

			}

			wrapTokenName := name

			symbol, err := instance.Symbol(&bind.CallOpts{})
			if err != nil {
				logger.InfoLog("----Error Symbol TxHash[%s] Error[%s]\n", m.TxHash.Hex(), err.Error())

			}

			wrapTokenSymbol := symbol

			kip7Transfer, err := instance.ParseTransfer(m)
			if err != nil {
				logger.InfoLog("----Error NewKip7  Error[%s]\n", err.Error())
				continue
			}

			wKlayString := kip7Transfer.Value.String()

			if len(wKlayString) < 12 {
				continue
			}
			wklayint := ChangeWklayValue(wKlayString)

			//   1000000000000000000 이게 1 klay
			//2149000000000000000000

			//var tokenid int64 = 0
			if wklayint >= minKlayValue {
				// 특정 klay 이상 value 만 체크

				blocknum := m.BlockNumber
				blocknumNew := big.NewInt(int64(blocknum))
				txhash := m.TxHash

				wklayLast := fmt.Sprintf("%f", float64(wklayint)/100000000)

				block, err := klaytndial.BlockByNumber(context.Background(), blocknumNew)
				if err != nil {
					logger.InfoLog("----Error BlockByNumber num[%d] Error[%s]\n", blocknumNew.Int64(), err.Error())
				}

				blocktime := block.Time().Int64()
				blocktimestring := time.Unix(blocktime, 0).Format("2006-01-02 15:04:05")

				// 해당 트랜잭션의 영수증
				rept, err := klaytndial.TransactionReceipt(context.Background(), txhash)
				if err != nil {
					logger.InfoLog("!!!!!!!!!!!!!!!!!!!!!!!!!!TransactionReceiptt Error vLog.TxHash[%s] , err[%s]\n", txhash, err.Error())
					continue
				}

				//var cName string = ""
				//var cSymbol string = ""
				for _, b := range rept.Logs {

					if b.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" { //nft contract 의 transfer
						//3번째 로그의 transfer(첫 transfer event) 가 KIP17 CONTRACT 의 event log 다
						kip17instance, err := kip17.NewKip17(b.Address, klaytndial)
						if err != nil {
							logger.InfoLog("----Error NewKip17 Error[%s]\n", err.Error())
							continue
						}

						name, err := kip17instance.Name(&bind.CallOpts{})
						if err != nil {
							logger.InfoLog("----Error kip17instance Name Error[%s]\n", err.Error())
						}

						symbol, err := kip17instance.Symbol(&bind.CallOpts{})
						if err != nil {
							logger.InfoLog("----Error kip17instance Symbol Error[%s]\n", err.Error())
						}

						kip17transfer, err := kip17instance.ParseTransfer(*b)
						if err != nil {
							///logger.InfoLog("----Error kip17instance ParseTransfer Maybe Not Kip17 Error[%s]\n", err.Error())
							continue
						}

						cName := name
						cSymbol := symbol
						tokenID := kip17transfer.TokenId.Int64()
						cAddress := b.Address

						tokeninfo := &TokenInfo{}

						tokeninfo.BlockTime = blocktimestring
						tokeninfo.TransactionHash = txhash
						tokeninfo.ContractName = cName
						tokeninfo.Contractaddress = cAddress
						tokeninfo.ContractSymbol = cSymbol

						tokenIDStr := fmt.Sprintf("%d", tokenID)
						tokeninfo.TokenID = tokenIDStr

						tokeninfo.WrapTokenAddress = wrapTokenAddress
						tokeninfo.WrapTokenName = wrapTokenName
						tokeninfo.WrapTokenSymbol = wrapTokenSymbol
						tokeninfo.KlayValue = wklayLast

						PrintTokenData(tokeninfo)

						//logger.InfoLog("--Blocknum[%d] , BlockTime[%s] txHash[%s] , wklayLast[%s] , ContractAddress[%s] , ContractName[%s] , ContractSymbol[%s], TokenID[%d] \n", blocknum, blocktimestring, txhash.Hex(), wklayLast, cAddress.Hex(), cName, cSymbol, tokenID)

					}
					//logger.InfoLog("----Log num[%d] , b.Topics[0][%s] \n", a, b.Topics[0].Hex())

				}

			}
		}

	}
}

func ChangeWklayValue(ValueString string) int {

	wKlayString := ValueString

	//logger.InfoLog("----wKlayString[%s]\n", wKlayString)

	wKlayrune := []rune(wKlayString)

	//logger.InfoLog("----wKlayrune[%s]\n", wKlayrune)
	rune10length := len(wKlayrune) - 10 // 전체 길이에서 10을 뺀다
	//logger.InfoLog("----rune10length[%d]\n", rune10length)

	wklayMinimal := string(wKlayrune[:rune10length])

	//logger.InfoLog("----wklayMinimal[%s]\n", wklayMinimal)

	wklayint, err := strconv.Atoi(wklayMinimal)
	if err != nil {
		return -1
	}

	return wklayint

}

func PrintTokenData(logdata *TokenInfo) {

	transaction := logdata.TransactionHash.Hex()
	blockTime := logdata.BlockTime[:10]
	contractAddress := logdata.Contractaddress.Hex()
	contractName := logdata.ContractName
	contractSymbol := logdata.ContractSymbol
	tokenID := logdata.TokenID

	wrapTokenAddress := logdata.WrapTokenAddress.Hex()
	wrapTokenName := logdata.WrapTokenName
	wrapTokenSymbol := logdata.WrapTokenSymbol
	klayValue := logdata.KlayValue //float64

	var b bytes.Buffer

	b.WriteString(blockTime)
	b.WriteString(",")
	b.WriteString(transaction)
	b.WriteString(",")
	b.WriteString(contractAddress)
	b.WriteString(",")
	b.WriteString(contractName)
	b.WriteString(",")
	b.WriteString(contractSymbol)
	b.WriteString(",")
	b.WriteString(tokenID)
	b.WriteString(",")

	b.WriteString(wrapTokenAddress)
	b.WriteString(",")
	b.WriteString(wrapTokenName)
	b.WriteString(",")
	b.WriteString(wrapTokenSymbol)
	b.WriteString(",")
	b.WriteString(klayValue)

	logger.TokenLog(b.String())

}

// WKLAY CONTRACT ADDRESS
//https://scope.klaytn.com/account/0xfd844c2fca5e595004b17615f891620d1cb9bbb2?tabId=txList

//KLAYTN 거래 예제
//https://scope.klaytn.com/tx/0x8386fa5a13e4ec15347e46bce5f7ed446c3ecf934205eada121252dab15c901c?tabId=eventLog
//https://scope.klaytn.com/tx/0x6cea865be1d1c0b0ae584ffcbbae11475f6d936652391225e80f811effeb774f?tabId=eventLog
//https://scope.klaytn.com/tx/0x1a59f3ee5de6a5581f858994c73715627f1e06b24959ae8798acfb37ed58d5ca?tabId=eventLog

//KLAY 치킨
//https://opensea.io/assets/klaytn/0x6b8f71aa8d5817d94056103886a1f07d12e78ce5/6142
//https://scope.klaytn.com/tx/0xe59be99f7540f24c0ef932a7f52ca718287023e3060a1c15b52bb20b68f802a7?tabId=eventLog

//i := fromBlockNumber

// for i <= toBlockNumber {

// 	logger.InfoLog("----- Block Num :  %d , Time : %s", i, time.Now())

// 	blockNum := big.NewInt(i)

// 	block, err := klaytndial.BlockByNumber(context.Background(), blockNum)
// 	if err != nil {
// 		log.Fatal("BlockByNumber : ", err)
// 	}

// 	blocktime := block.Time().Int64()
// 	blocktimestring := time.Unix(blocktime, 0).Format("2006-01-02 15:04:05")
// 	fmt.Printf("blockdata time : %s\n", blocktimestring)

// for _, txs := range block.Transactions() {

// 	klayint64 := txs.Value().Int64()

// 	// if klayint64 < minKlayValue {
// 	// 	continue
// 	// }

// 	txhash := txs.Hash()

// 	if txhash.Hex() != "0x5cf031ef3e2422b936323fd71772452adbbad540358081b2b636dce6f5e118f0" {
// 		continue
// 	}

// 	klaystring := fmt.Sprintf("%f", float64(klayint64)/1000000000000000000)

// 	logger.InfoLog("------ start parse Transaction TxHash[%s] Klay[%s]\n", txhash.Hex(), klaystring)

// 	// 해당 트랜잭션의 영수증
// 	rept, err := klaytndial.TransactionReceipt(context.Background(), txhash)
// 	if err != nil {
// 		logger.InfoLog("--ransactionReceipt Error vLog.TxHash[%s] , err[%s]\n", txhash, err.Error())
// 		continue
// 	}

// 	if len(rept.Logs) == 0 { //event log 가없으면 일반 거래일것이다
// 		continue
// 	}

// 	for _, m := range rept.Logs {

// 		//0번쨰는 opensea Contract 인것 같다
// 		//1번쨰도 opensea Contract 인것 같다
// 		//2번째는 해당 NFT contract  TRANSFER KIP17
// 		//3번째는 WKLAY 거래 총 비용 TRANFER KIP7
// 		//4번쨰 부터 수수료 개념 TRANFSER
// 		contractAddr := m.Address.Hex()
// 		logger.InfoLog("--ContractAddr ContractAddr[%s] , m.Topics[0][%s] \n", contractAddr, m.Topics[0].Hex())

// 		instance, err := kip17.NewKip17(m.Address, klaytndial)
// 		if err != nil {
// 			logger.InfoLog("------- NewKip17 contractAddressHex[%s] , error[%s] ", m.Address.Hex(), err.Error())
// 			return
// 		}

// 		if m.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {

// 			Name, err := instance.Name(&bind.CallOpts{})
// 			if err != nil {
// 				logger.InfoLog("GetDataERC721 instance.Name error[%s] ", err.Error())

// 			}

// 			Symbol, err := instance.Symbol(&bind.CallOpts{})
// 			if err != nil {
// 				logger.InfoLog("GetDataERC721 instance.Symbol error[%s] ", err.Error())

// 			}

// 			logger.InfoLog("------- NewKip17 contractAddressHex[%s] , Name[%s] , Symbol[%s]", m.Address.Hex(), Name, Symbol)

// 			// 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef  Event Transfer
// 			kip17Transfer, err := instance.ParseTransfer(*m)
// 			if err != nil {
// 				logger.InfoLog("GetDataERC721 instance.ParseTransfer  error[%s] ", err.Error())
// 				//return
// 			}

// 			//TokenID = fmt.Sprintf("%s", kip17Transfer.TokenId)
// 			if err == nil {
// 				logger.InfoLog("Get KIP17 Transfer  From[%s] , To[%s]  , TokenID[%d]", kip17Transfer.From.Hex(), kip17Transfer.To.Hex(), kip17Transfer.TokenId.Int64())
// 			}

// 			kip7instance, err := kip7.NewKip7(m.Address, klaytndial)
// 			if err != nil {
// 				logger.InfoLog("------- New kip7 contractAddressHex[%s] , error[%s] ", m.Address.Hex(), err.Error())
// 				//return
// 			}

// 			kip7Transfer, err := kip7instance.ParseTransfer(*m)
// 			if err != nil {
// 				logger.InfoLog("GetDataERC721 instance.ParseTransfer  error[%s] ", err.Error())
// 				//return
// 			}

// 			if err == nil {
// 				logger.InfoLog("Get KIP7 Transfer  From[%s] , To[%s]  , Value[%s]", kip7Transfer.From.Hex(), kip7Transfer.To.Hex(), kip7Transfer.Value.String())
// 			}

// 		}

// 	}

// }

//	i = i + 1
//}
