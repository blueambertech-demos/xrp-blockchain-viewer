package api

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/blueambertech-demos/xrp-blockchain-viewer/pkg/ledger"
	"github.com/blueambertech-demos/xrp-blockchain-viewer/pkg/mock"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	XRPNetAddress = "https://s.altnet.rippletest.net:51234"
	ledger.SetMemoryStore(context.Background(), &mock.Store{})
	os.Exit(m.Run())
}

func TestRegisterHandlers(t *testing.T) {
	_, e, _ := ginTestSetup()
	RegisterHandlers(e)
}

func TestLedgerValidatedInfo(t *testing.T) {
	ctx, _, w := ginTestSetup()
	mockJson(ctx, "POST")
	handleLedgerValidatedInfo(ctx)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect response code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if len(respBody) == 0 {
		t.Error("no body returned")
	}
}

func TestLedgerClosedInfo(t *testing.T) {
	ctx, _, w := ginTestSetup()
	mockJson(ctx, "POST")
	handleLedgerClosedInfo(ctx)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect response code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if len(respBody) == 0 {
		t.Error("no body returned")
	}
}

func TestLedgerCurrentInfo(t *testing.T) {
	ctx, _, w := ginTestSetup()
	mockJson(ctx, "POST")
	handleLedgerCurrentInfo(ctx)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Incorrect response code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if len(respBody) == 0 {
		t.Error("no body returned")
	}
}

func ginTestSetup() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, e := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx, e, w
}

func mockJson(c *gin.Context, m string) {
	c.Request.Method = m
	c.Request.Header.Set("Content-Type", "application/json")
}
