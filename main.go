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

type Info struct {
	BlockNumber      uint64
	WrapTokenName    string
	WrapTokenSymbol  string
	WrapTokenAddress common.Address
	Klay             int
}

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

	latestBlockNum, err := klaytndial.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Lastest Block Num : %s : ", latestBlockNum.String())
	}

	// klaytn block list
	// klaytn lastest block num 80292389

	// block number 73967410 (Nov 01, 2021 00:00:00 / UTC+9)
	// block number 73967409 (Oct 31, 2021 23:59:59 / UTC+9)
	// block number 71291139 (Oct 01, 2021 00:00:00 / UTC+9)
	// block number 71291138 (Sep 30, 2021 23:59:59 / UTC+9)

	// block number 66021022 (Aug 01, 2021 00:00:00 / UTC+9)

	logger.InfoLog("-----Check Lastest Block Num :  %d ", latestBlockNum.Int64())

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

	var minKlayValue int = 10000000 // 500000000 5 klay , 1000000000 10 klay 100 klay 100000000000000000000 인데 0 10개 뺀다 10000000000 , 100000000000 1000klay
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

	TxHashInfo := make(map[common.Hash]*Info, 0)

	//fmt.Println("FilterLogs Count : ", len(logs))

	for _, m := range logs { // address wklay log

		mtxhash := m.TxHash

		logger.InfoLog("-----FilterLogs :TxHash[%s] , m.Topics[0].Hex()[%s] , m.Address[%s]", mtxhash.Hex(), m.Topics[0].Hex(), m.Address.Hex())
		//0xfd844c2fca5e595004b17615f891620d1cb9bbb2 wklay address
		//if m.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" { //wklay contract의 transfer
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

		wrapTokenName, err := instance.Name(&bind.CallOpts{})
		if err != nil {
			logger.InfoLog("----Error Name TxHash[%s] Error[%s]\n", m.TxHash.Hex(), err.Error())

		}

		wrapTokenSymbol, err := instance.Symbol(&bind.CallOpts{})
		if err != nil {
			logger.InfoLog("----Error Symbol TxHash[%s] Error[%s]\n", m.TxHash.Hex(), err.Error())

		}

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

		//logger.InfoLog("-----FilterLogs :wrapTokenName[%s] , symbol[%s] , wklayint[%d]", wrapTokenName, wrapTokenSymbol, wklayint)

		//   1000000000000000000 이게 1 klay
		//2149000000000000000000

		//var tokenid int64 = 0
		if wklayint >= minKlayValue {
			// 특정 klay 이상 value 만 체크
			txhashinfo, ex := TxHashInfo[mtxhash]
			if !ex {
				info := &Info{}
				info.BlockNumber = m.BlockNumber
				info.Klay = wklayint
				info.WrapTokenName = wrapTokenName
				info.WrapTokenSymbol = wrapTokenSymbol
				info.WrapTokenAddress = wrapTokenAddress

				TxHashInfo[mtxhash] = info
			} else {

				if txhashinfo.Klay < wklayint { // 등록 되어있는 klay 가 작으면 새로 수정
					txhashinfo.Klay = wklayint
				}
			}

		}

	} // for

	for txHash, info := range TxHashInfo {

		fmt.Printf("TxHashInfo : k[%s] blockNumber[%d] Klay[%d]\n", txHash.Hex(), info.BlockNumber, info.Klay)

		blocknum := info.BlockNumber
		blocknumNew := big.NewInt(int64(blocknum))
		txhash := txHash

		wklayint := info.Klay
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

		for _, log := range rept.Logs {

			//logger.InfoLog("-----rept.Logs :TxHash[%s] , m.Topics[0].Hex()[%s] , m.Address[%s]", txHash.Hex(), log.Topics[0].Hex(), log.Address.Hex())

			if log.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" { //nft contract 의 transfer
				//3번째 로그의 transfer(첫 transfer event) 가 KIP17 CONTRACT 의 event log 다
				kip17instance, err := kip17.NewKip17(log.Address, klaytndial)
				if err != nil {
					logger.InfoLog("----Error NewKip17 Error[%s]\n", err.Error())
					continue
				}

				kip17transfer, err := kip17instance.ParseTransfer(*log) // 이 로그 parse 가 에러 나면 kip17 이 아니다
				if err != nil {
					///logger.InfoLog("----Error kip17instance ParseTransfer Maybe Not Kip17 Error[%s]\n", err.Error())
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

				cName := name
				cSymbol := symbol
				tokenID := kip17transfer.TokenId.Int64()
				cAddress := log.Address

				tokeninfo := &TokenInfo{}

				tokeninfo.BlockTime = blocktimestring
				tokeninfo.TransactionHash = txhash
				tokeninfo.ContractName = cName
				tokeninfo.Contractaddress = cAddress
				tokeninfo.ContractSymbol = cSymbol

				tokenIDStr := fmt.Sprintf("%d", tokenID)
				tokeninfo.TokenID = tokenIDStr

				tokeninfo.WrapTokenAddress = info.WrapTokenAddress
				tokeninfo.WrapTokenName = info.WrapTokenName
				tokeninfo.WrapTokenSymbol = info.WrapTokenSymbol
				tokeninfo.KlayValue = wklayLast

				PrintTokenData(tokeninfo)

			}

		}

	}

}

func ChangeWklayValue(ValueString string) int {

	wKlayString := ValueString

	wKlayrune := []rune(wKlayString)

	rune10length := len(wKlayrune) - 10 // 전체 길이에서 10을 뺀다

	wklayMinimal := string(wKlayrune[:rune10length])

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
