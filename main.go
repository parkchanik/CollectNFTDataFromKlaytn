package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	"runtime"
	"time"

	"CollectNFTDataKlaytn/config"
	"CollectNFTDataKlaytn/kas"
	"CollectNFTDataKlaytn/parse"
	klayClient "github.com/klaytn/klaytn/client"

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

	var fromBlockNumber int64 = 80520318 //13717846
	var toBlockNumber int64 = 80520318

	if *fromNum != 0 {
		fromBlockNumber = *fromNum
		toBlockNumber = *toNum
	}

	i := fromBlockNumber

	for i <= toBlockNumber {

		logger.InfoLog("----- Block Num :  %d , Time : %s", i, time.Now())

		blockNum := big.NewInt(i)

		block, err := klaytndial.BlockByNumber(context.Background(), blockNum)
		if err != nil {
			log.Fatal("BlockByNumber : ", err)
		}

		blocktime := block.Time().Int64()
		blocktimestring := time.Unix(blocktime, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("blockdata time : %s\n", blocktimestring)

		for _, txs := range block.Transactions() {

			//etherint64 := txs.Value().Int64()

			// if etherint64 < minETHValue {
			// 	continue
			// }

			txhash := txs.Hash()

			if txhash.Hex() != "0x5cf031ef3e2422b936323fd71772452adbbad540358081b2b636dce6f5e118f0" {
				continue
			}

			logger.InfoLog("------ start parse Transaction TxHash[%s]\n", txhash.Hex())

			// 해당 트랜잭션의 영수증
			rept, err := klaytndial.TransactionReceipt(context.Background(), txhash)
			if err != nil {
				logger.InfoLog("--ransactionReceipt Error vLog.TxHash[%s] , err[%s]\n", txhash, err.Error())
				continue
			}

			if len(rept.Logs) == 0 { //event log 가없으면 일반 거래일것이다
				continue
			}

			parse.ParseReceipt(rept)

		}

		i = i + 1
	}
}
