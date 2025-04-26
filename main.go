package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// House represents a single row of housing data
type House struct {
	Value    int     `json:"value"`
	Income   float64 `json:"income"`
	Age      float64 `json:"age"`
	Rooms    int     `json:"rooms"`
	Bedrooms int     `json:"bedrooms"`
	Pop      int     `json:"pop"`
	HH       int     `json:"hh"`
}

func main() {
	var inputPath, outputPath string

	// check if arguments are passed
	if len(os.Args) == 3 {
		inputPath = os.Args[1]
		outputPath = os.Args[2]
	} else {
		fmt.Println("No input arguments detected.")
		fmt.Println("Please enter the input CSV file path:")
		fmt.Scanln(&inputPath)
		fmt.Println("Please enter the output JL file path:")
		fmt.Scanln(&outputPath)
	}

	houses, err := readCSV(inputPath)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	err = writeJSONLines(outputPath, houses)
	if err != nil {
		fmt.Println("Error writing JSON Lines:", err)
		os.Exit(1)
	}

	fmt.Println("Conversion complete. Output saved to", outputPath)
}

// readCSV reads housing data from a CSV file and returns a slice of House structs
func readCSV(path string) ([]House, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var houses []House
	for i, row := range records {
		if i == 0 {
			continue // Skip header
		}
		house, err := parseRow(row)
		if err != nil {
			fmt.Println("Skipping invalid row:", err)
			continue
		}
		houses = append(houses, house)
	}
	return houses, nil
}

// use ParseFloat + cast to int to support scientific notation like "1e+05"
func mustParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("Invalid float: %s", s))
	}
	return f
}

// parseRow converts a slice of strings into a House struct
func parseRow(row []string) (House, error) {
	return House{
		Value:    int(mustParseFloat(row[0])),
		Income:   mustParseFloat(row[1]),
		Age:      mustParseFloat(row[2]),
		Rooms:    int(mustParseFloat(row[3])),
		Bedrooms: int(mustParseFloat(row[4])),
		Pop:      int(mustParseFloat(row[5])),
		HH:       int(mustParseFloat(row[6])),
	}, nil

}

// writeJSONLines writes a slice of House structs to a .jl file (JSON Lines format)
func writeJSONLines(path string, houses []House) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, h := range houses {
		line, err := json.Marshal(h)
		if err != nil {
			return err
		}
		_, err = file.WriteString(string(line) + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
