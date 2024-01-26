package main

import (
	"context"
	"flag"

	"github.com/blueambertech-demos/xrp-blockchain-viewer/api"
	"github.com/blueambertech-demos/xrp-blockchain-viewer/pkg/bytestore"
	"github.com/blueambertech-demos/xrp-blockchain-viewer/pkg/ledger"
	"github.com/gin-gonic/gin"
)

func main() {
	var serverAddress string
	flag.StringVar(&serverAddress, "sa", "https://s.altnet.rippletest.net:51234", "")
	flag.Parse()

	exitCtx, canc := context.WithCancel(context.Background())
	defer canc()
	ledger.SetMemoryStore(exitCtx, bytestore.NewStore())
	r := gin.Default()
	api.XRPNetAddress = serverAddress
	api.RegisterHandlers(r)

	r.Run(":7770")
}
