package api

import (
	"github.com/blueambertech-demos/xrp-blockchain-viewer/pkg/ledger"
	"github.com/gin-gonic/gin"
)

var XRPNetAddress string

func RegisterHandlers(e *gin.Engine) {
	e.GET("/ledger/validated/info", handleLedgerValidatedInfo)
	e.GET("/ledger/closed/info", handleLedgerClosedInfo)
	e.GET("/ledger/current/info", handleLedgerCurrentInfo)
}

func handleLedgerValidatedInfo(ctx *gin.Context) {
	ledgerInf, err := ledger.Info(XRPNetAddress, ledger.Validated)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(500, ctx.Errors.JSON())
		return
	}
	ctx.String(200, string(ledgerInf))
}

func handleLedgerClosedInfo(ctx *gin.Context) {
	ledgerInf, err := ledger.Info(XRPNetAddress, ledger.Closed)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(500, ctx.Errors.JSON())
		return
	}
	ctx.String(200, string(ledgerInf))
}

func handleLedgerCurrentInfo(ctx *gin.Context) {
	ledgerInf, err := ledger.Info(XRPNetAddress, ledger.Current)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(500, ctx.Errors.JSON())
		return
	}
	ctx.String(200, string(ledgerInf))
}
