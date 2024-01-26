package bytestore

import (
	"bytes"
	"context"
	"sync"
	"testing"
	"time"
)

func TestGetData(t *testing.T) {
	testKey, testData, expiry := createTestData()
	store := NewStore()
	store.SetData(context.Background(), testKey, testData, expiry)

	d, ok := store.GetData(testKey)
	if !ok {
		t.Error("data not found")
		return
	}
	if !bytes.Equal(d, testData) {
		t.Errorf("wrong data retrieved. Expected %v got %v", testData, d)
	}
}

func TestGetMissing(t *testing.T) {
	testKey, _, _ := createTestData()
	store := NewStore()

	_, ok := store.GetData(testKey)
	if ok {
		t.Error("data found, expected failure")
		return
	}
}

func TestSetData(t *testing.T) {
	testKey, testData, expiry := createTestData()
	store := NewStore()
	store.SetData(context.Background(), testKey, testData, expiry)
}

func TestSetDataConcurrent(t *testing.T) {
	testKey, testData, expiry := createTestData()
	n := 50
	wg := sync.WaitGroup{}
	wg.Add(n)
	store := NewStore()
	for i := 0; i < n; i++ {
		go func(k string, d []byte, w *sync.WaitGroup) {
			defer w.Done()
			store.SetData(context.Background(), testKey, testData, expiry)
		}(testKey, testData, &wg)
	}
	wg.Wait()
}

func TestSetDataExpiry(t *testing.T) {
	testKey, testData, expiry := createTestData()
	store := NewStore()
	store.SetData(context.Background(), testKey, testData, expiry)

	_, ok := store.GetData(testKey)
	if !ok {
		t.Error("data not found")
		return
	}

	time.Sleep(time.Second * 3)

	_, ok = store.GetData(testKey)
	if ok {
		t.Error("data still found after expiry")
		return
	}
}

func TestSetDataWithCancellation(t *testing.T) {
	testKey, testData, expiry := createTestData()
	store := NewStore()

	cancCtx, canc := context.WithCancel(context.Background())
	store.SetData(cancCtx, testKey, testData, expiry)

	_, ok := store.GetData(testKey)
	if !ok {
		t.Error("data not found")
		return
	}

	canc()

	time.Sleep(time.Second * 3)

	_, ok = store.GetData(testKey)
	if !ok {
		t.Error("data expected to be found after monitor routine was cancelled")
		return
	}
}

func TestDeleteData(t *testing.T) {
	testKey, testData, expiry := createTestData()
	store := NewStore()
	store.SetData(context.Background(), testKey, testData, expiry)

	_, ok := store.GetData(testKey)
	if !ok {
		t.Error("data not found")
		return
	}

	store.DeleteData(testKey)

	_, ok = store.GetData(testKey)
	if ok {
		t.Error("data found (should've been deleted)")
		return
	}
}

func createTestData() (string, []byte, int64) {
	testKey := "testkey"
	testData := []byte("here is some data")
	exp := time.Now().Add(time.Second * 1)
	return testKey, testData, exp.Unix()
}
