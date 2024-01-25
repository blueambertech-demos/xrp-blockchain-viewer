package bytestore

import (
	"bytes"
	"sync"
	"testing"
)

func TestGetData(t *testing.T) {
	testKey, testData := createTestData()
	store := NewStore()
	store.SetData(testKey, testData)

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
	testKey, _ := createTestData()
	store := NewStore()

	_, ok := store.GetData(testKey)
	if ok {
		t.Error("data found, expected failure")
		return
	}
}

func TestSetData(t *testing.T) {
	testKey, testData := createTestData()
	store := NewStore()
	store.SetData(testKey, testData)
}

func TestSetDataConcurrent(t *testing.T) {
	testKey, testData := createTestData()
	n := 50
	wg := sync.WaitGroup{}
	wg.Add(n)
	store := NewStore()
	for i := 0; i < n; i++ {
		go func(k string, d []byte, w *sync.WaitGroup) {
			defer w.Done()
			store.SetData(testKey, testData)
		}(testKey, testData, &wg)
	}
	wg.Wait()
}

func createTestData() (string, []byte) {
	testKey := "testkey"
	testData := []byte("here is some data")
	return testKey, testData
}
