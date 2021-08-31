package offchain

import (
	"os"
	"testing"

	"github.com/leverwwz/go-substrate-rpc-client/client"
	"github.com/leverwwz/go-substrate-rpc-client/config"
)

var offchain *Offchain

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	offchain = NewOffchain(cl)
	os.Exit(m.Run())
}
