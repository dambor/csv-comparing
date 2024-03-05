package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file1 := "file1.csv"
	file2 := "file2.csv"

	addedFile, err := os.Create("added.csv")
	if err != nil {
		panic(err)
	}
	defer addedFile.Close()

	deletedFile, err := os.Create("deleted.csv")
	if err != nil {
		panic(err)
	}
	defer deletedFile.Close()

	changedFile, err := os.Create("changed.csv")
	if err != nil {
		panic(err)
	}
	defer changedFile.Close()

	file1Records := readCSV(file1)
	file2Records := readCSV(file2)

	// Search for added, deleted, and changed records
	for _, record1 := range file1Records {
		found := false
		for _, record2 := range file2Records {
			if record1[0] == record2[0] { // Assuming first column is ID
				found = true
				if !equal(record1, record2) {
					writeCSV(changedFile, record1)
				}
				break
			}
		}
		if !found {
			writeCSV(deletedFile, record1)
		}
	}

	for _, record2 := range file2Records {
		found := false
		for _, record1 := range file1Records {
			if record2[0] == record1[0] { // Assuming first column is ID
				found = true
				break
			}
		}
		if !found {
			writeCSV(addedFile, record2)
		}
	}

	fmt.Println("Processing complete.")
}

func readCSV(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func writeCSV(file *os.File, record []string) {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(record)
	if err != nil {
		panic(err)
	}
}

func equal(record1, record2 []string) bool {
	if len(record1) != len(record2) {
		return false
	}
	for i := range record1 {
		if record1[i] != record2[i] {
			return false
		}
	}
	return true
}
