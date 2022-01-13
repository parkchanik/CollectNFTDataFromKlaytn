package kas

////////////////////////////////////////////////////////////////////////////////////
////참고 https://github.com/airbloc/airbloc-go/tree/6823a42e4d827ee8f8cea406eb95a2147ec27a6a/pkg/kas
/////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	klayClient "github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/networks/rpc"
	"github.com/klaytn/klaytn/params"
	"net/http"
)

type Config struct {
	Endpoint        string `default:"https://node-api.klaytnapi.com/v1/klaytn"`
	Network         string `default:"cypress"`
	AccessKeyID     string
	SecretAccessKey string
}

var networkNameToChainIDs = map[string]uint64{
	"cypress": params.CypressNetworkId,
	"baobab":  params.BaobabNetworkId,
}

// Dial connects to Klaytn API Service and returns a JSON-RPC client.
func Dial(cfg Config) (*klayClient.Client, error) {
	if _, ok := networkNameToChainIDs[cfg.Network]; !ok {
		return nil, fmt.Errorf("validate KAS config: unknown network: \"%s\"", cfg.Network)
	}
	if cfg.AccessKeyID == "" {
		return nil, fmt.Errorf("validate KAS config: access key ID is empty")
	}
	if cfg.SecretAccessKey == "" {
		return nil, fmt.Errorf("validate KAS config: secret access key is empty")
	}
	cli := new(http.Client)
	cli.Transport = kasAuthTransport{cfg}
	rpcCli, err := rpc.DialHTTPWithClient(cfg.Endpoint, cli)
	if err != nil {
		return nil, err
	}
	return klayClient.NewClient(rpcCli), nil
}

// kasAuthTransport wraps http.RoundTripper and adds authentication header to HTTP request
// for communicating with Klaytn API Service.
type kasAuthTransport struct {
	cfg Config
}

// RoundTrip adds authentication header to JSONRPC request.
// For details, please refer to KAS documentation (https://console.klaytnapi.com/ko/service/node)
func (k kasAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(k.cfg.AccessKeyID, k.cfg.SecretAccessKey)
	req.Header.Set("x-chain-id", fmt.Sprintf("%d", networkNameToChainIDs[k.cfg.Network]))

	return http.DefaultTransport.RoundTrip(req)
}
