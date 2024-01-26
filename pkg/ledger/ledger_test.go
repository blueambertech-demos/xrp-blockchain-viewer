package ledger

import (
	"context"
	"os"
	"testing"

	"github.com/blueambertech-demos/xrp-blockchain-viewer/pkg/mock"
)

var xrpTestNet = "https://s.altnet.rippletest.net:51234"

func TestMain(m *testing.M) {
	SetMemoryStore(context.Background(), &mock.Store{})
	os.Exit(m.Run())
}

func TestInfoValidated(t *testing.T) {
	response, err := Info(xrpTestNet, Validated)
	if err != nil {
		t.Error(err)
		return
	}
	if len(response) == 0 {
		t.Error("ledger response empty")
	}
}

func TestInfoClosed(t *testing.T) {
	response, err := Info(xrpTestNet, Closed)
	if err != nil {
		t.Error(err)
		return
	}
	if len(response) == 0 {
		t.Error("ledger response empty")
	}
}

func TestInfoCurrent(t *testing.T) {
	response, err := Info(xrpTestNet, Current)
	if err != nil {
		t.Error(err)
		return
	}
	if len(response) == 0 {
		t.Error("ledger response empty")
	}
}

func TestInfoBadAddress(t *testing.T) {
	_, err := Info("http://bad", Validated)
	if err == nil {
		t.Error("expected error, received nil")
	}
}
