package httputil

import "testing"

type testObj struct {
	value1 string
	value2 string
}

const baseTestAddress = "http://httpbin.org"

func TestPostJson(t *testing.T) {
	obj := testObj{
		value1: "hi",
		value2: "bonjour",
	}
	r, err := PostJson(baseTestAddress+"/anything", obj)
	if err != nil {
		t.Error(err)
	} else if len(r) == 0 {
		t.Error("no response")
	}
}

func TestPostJsonBadUrl(t *testing.T) {
	obj := testObj{
		value1: "hi",
		value2: "bonjour",
	}
	_, err := PostJson("hjbdsuhjbs", obj)
	if err == nil {
		t.Error("expected error, received nil")
	}
}

func TestPostJsonNonOKResponse(t *testing.T) {
	obj := testObj{
		value1: "hi",
		value2: "bonjour",
	}
	_, err := PostJson(baseTestAddress+"/status/400", obj)
	if err == nil {
		t.Error("expected error, received nil")
	}
}
