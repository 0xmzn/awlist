package awesomestore

import (
	"os"
	"reflect"
	"testing"

	"github.com/0xmzn/awelist/internal/model"
	"gopkg.in/yaml.v3"
)

func TestSortData_Success(t *testing.T) {
	store, err := NewStore("testdata/sample_unsorted.yaml")
	if err != nil {
		t.Fatalf("Failed to initialize unsorted store: %v", err)
	}

	expectedSortedDataBytes, err := os.ReadFile("testdata/sample_sorted.yaml")
	if err != nil {
		t.Fatalf("Failed to read sample_sorted.yaml: %v", err)
	}

	var expectedSortedData model.AwesomeData
	err = yaml.Unmarshal(expectedSortedDataBytes, &expectedSortedData)
	if err != nil {
		t.Fatalf("Failed to unmarshal sample_sorted.yaml: %v", err)
	}

	if !reflect.DeepEqual(store.Data(), expectedSortedData) {
		t.Errorf("Data after sorting is not as expected.\nExpected: %+v\nGot: %+v", expectedSortedData, store.Data())
	}
}

func TestNewStore_Success(t *testing.T) {
	store, err := NewStore("testdata/sample_unsorted.yaml")
	if err != nil {
		t.Fatalf("NewStore failed with valid path: %v", err)
	}
	if store == nil {
		t.Fatal("NewStore returned nil store with valid path")
	}
	if len(store.Data()) == 0 {
		t.Error("NewStore did not load data or data is empty")
	}
}

func TestNewStore_FileNotFound(t *testing.T) {
	_, err := NewStore("testdata/non_existent_file.yaml")
	if err == nil {
		t.Fatal("NewStore did not return an error for a non-existent file")
	}
}

func TestData_ReturnsData(t *testing.T) {
	store, err := NewStore("testdata/sample_sorted.yaml")
	if err != nil {
		t.Fatalf("NewStore failed with valid path: %v", err)
	}

	expectedStore, err := NewStore("testdata/sample_sorted.yaml")
	if err != nil {
		t.Fatalf("Failed to load expected data for comparison: %v", err)
	}

	err = expectedStore.Load()
	if err != nil {
		t.Fatalf("Failed to load data into expectedStore: %v", err)
	}

	retrievedData := store.Data()
	if !reflect.DeepEqual(retrievedData, expectedStore.Data()) {
		t.Errorf("Data() method did not return the expected data.\nExpected: %+v\nGot: %+v", expectedStore.Data(), retrievedData)
	}

	if len(retrievedData) == 0 {
		t.Error("Data() method returned empty data for a non-empty store")
	}
}
