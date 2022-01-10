package main

import (
	//"context"

	"encoding/json"
	"fmt"

	//"errors"
	"io/ioutil"
	//"math/big"

	//"strconv"

	"bytes"

	//"time"
	//"strconv"
	//"encoding/json"

	//"io"
	//"log"
	//"os"

	//"math"
	//"encoding/hex"

	//"strings"

	klaytntypes "github.com/klaytn/klaytn/blockchain/types"

	logger "KalytnProj/logger"

	"net/http"
)

type ErrorType uint8

const (
	Ok                  ErrorType = 200 // GET, POST Only
	Created                       = 201 // POST Only
	Accepted                      = 202 // PUT, DELETE Only
	BadRequest                    = 400 // GET Only
	Unauthorized                  = 401
	NotFound                      = 404
	MethodNotAllowed              = 405
	InternalServerError           = 500
	ServiceUnavailable            = 503
)

type JsonRequestStruct struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type JsonResponsetStruct struct {
	JsonRPC string             `json:"jsonrpc"`
	ID      int                `json:"id"`
	Result  klaytntypes.Header `json:"result"`
}

func main() {

	logger.LoggerInit()

	dataStruct := JsonRequestStruct{}

	// JsonRPC: "2.0",
	// 	Method:  "klay_getBlockByNumber",
	// 	Params:  []interface{"0x1b4" , true},
	// 	ID:      1,
	dataStruct.JsonRPC = "2.0"
	dataStruct.Method = "klay_getBlockByNumber"

	//dataStruct.Method = "klay_blockNumber"

	dataStruct.Params = make([]interface{}, 0)
	dataStruct.Params = append(dataStruct.Params, "0x4c0c248", true)

	dataStruct.ID = 1

	url := "https://node-api.klaytnapi.com/v1/klaytn"
	pbytes, _ := json.Marshal(dataStruct)
	buff := bytes.NewBuffer(pbytes)

	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		logger.ErrorLog("Http Post : fail NewRequest err : %s", err.Error())

	}
	//KASKRBBYWZ90A4R2PALA3IQ2
	//Content-Type 헤더 추가
	req.Header.Add("Authorization", "Basic S0FTS1JCQllXWjkwQTRSMlBBTEEzSVEyOkktVThybTUxQnY1d1l2WWtXXzlqdG84UDRLSXpuZTNBVGJKVnNseDE=")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-chain-id", "8217")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.ErrorLog("Http Post : fail DefaultClient DO() err : %s", err.Error())
		return

	}

	defer res.Body.Close()

	statuscode := ErrorType(res.StatusCode)

	// Response 체크.
	if statuscode.IsSuccess() == false {
		logger.ErrorLog("Http Post : statuscode [%d]", res.StatusCode)
		return
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.ErrorLog("Http Post : Fail ReadAll err :  [%s]", err.Error())

	}

	responseData := JsonResponsetStruct{}

	err = json.Unmarshal(respBody, &responseData)
	if err != nil {
		fmt.Println("err responseData :", err)
	}

	fmt.Println("respBody responseData.Result.Number : ", responseData.Result)

}

func HttpPost(url string, object interface{}) ([]byte, bool) {
	pbytes, _ := json.Marshal(object)
	buff := bytes.NewBuffer(pbytes)

	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		//logger.ServerLogError("Http Post : fail NewRequest")
		return nil, false
	}

	//Content-Type 헤더 추가
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		//logger.ServerLogError("Http Post : fail DefaultClient DO()")
		return nil, false
	}

	statuscode := ErrorType(res.StatusCode)

	defer res.Body.Close()

	// Response 체크.
	if statuscode.IsSuccess() == false {
		//logger.ServerLogError("Http Post : statuscode [%d]", res.StatusCode)
		return nil, false

	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//logger.ServerLogError("Http Post : Fail ReadAll [%v]", err.Error())
		return nil, false
	}
	return respBody, true
}

func (errType ErrorType) IsSuccess() bool {
	return (errType == Created || errType == Accepted || errType == Ok)
	//return int(errType) < 400
}
