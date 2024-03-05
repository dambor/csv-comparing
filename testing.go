package main

import (
	"reflect"
	"testing"
)

func TestCompareCSVFiles(t *testing.T) {
	expectedAdded := [][]string{{"3", "noam", "noam@ownbackup.com"}}
	expectedDeleted := [][]string{{"1", "yoni", "yoni@ownbackup.com"}}
	expectedChanged := [][]string{{"2", "roy", "roy@ownbackup.com"}}

	added, deleted, changed, err := CompareCSVFiles("file1.csv", "file2.csv")
	if err != nil {
		t.Errorf("Error comparing CSV files: %v", err)
	}

	if !reflect.DeepEqual(added, expectedAdded) {
		t.Errorf("Added records mismatch, expected: %v, got: %v", expectedAdded, added)
	}

	if !reflect.DeepEqual(deleted, expectedDeleted) {
		t.Errorf("Deleted records mismatch, expected: %v, got: %v", expectedDeleted, deleted)
	}

	if !reflect.DeepEqual(changed, expectedChanged) {
		t.Errorf("Changed records mismatch, expected: %v, got: %v", expectedChanged, changed)
	}
}
