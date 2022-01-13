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
	klayClient "github.com/klaytn/klaytn/client"
)

var klaytndial *klayClient.Client = nil

var IMAGE_PATH string = "../CollectNFT/KlaytnImages/"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("GOMAXPROCS : ", runtime.GOMAXPROCS(0))

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

	// block number 13717846 (Nov-30-2021 11:59:50 PM +UTC)
	// block number 13527859 (Nov-01-2021 12:00:07 AM +UTC)
	// block number 13527858 (Oct-31-2021 11:59:20 PM +UTC)
	// block number 13330090 (Oct-01-2021 12:00:00 AM +UTC)
	// block number 13330089 (Sep-30-2021 11:59:56 PM +UTC)

	var fromBlockNumber int64 = 13340710 //13717846
	var toBlockNumber int64 = 13340710

	if *fromNum != 0 {
		fromBlockNumber = *fromNum
		toBlockNumber = *toNum
	}

	i := fromBlockNumber

	for i <= toBlockNumber {

		blockNum := big.NewInt(i)

		block, err := klaytndial.BlockByNumber(context.Background(), blockNum)
		if err != nil {
			log.Fatal("BlockByNumber : ", err)
		}

		blocktime := block.Time().Int64()
		blocktimestring := time.Unix(blocktime, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("blockdata time : %s\n", blocktimestring)
		i = i + 1
	}
}
