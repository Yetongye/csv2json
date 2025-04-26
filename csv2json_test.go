package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

//
// This test file verifies the functionality of:
// - parseRow(): converting a CSV row into a House struct
// - writeJSONLines(): writing a slice of House structs as JSON Lines (.jl) file
//

// TestParseRow_Valid checks if a valid CSV row is parsed correctly into a House struct
func TestParseRow_Valid(t *testing.T) {
	// Example row from the CSV file
	row := []string{"452600", "8.3252", "41", "880", "129", "322", "126"}

	// Expected result after parsing
	expected := House{
		Value:    452600,
		Income:   8.3252,
		Age:      41,
		Rooms:    880,
		Bedrooms: 129,
		Pop:      322,
		HH:       126,
	}

	// Call the function under test
	house, err := parseRow(row)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Assert that the parsed result matches the expected struct
	if house != expected {
		t.Errorf("Parsed house does not match expected.\nGot: %+v\nExpected: %+v", house, expected)
	}
}

// TestParseRow_Invalid checks how the parser handles invalid input data (non-numeric values)
func TestParseRow_Invalid(t *testing.T) {
	// Invalid input: non-numeric strings
	row := []string{"abc", "xyz", "41", "880", "129", "322", "126"}

	// Expect an error to be returned
	_, err := parseRow(row)
	if err == nil {
		t.Error("Expected error for invalid input row, got none")
	}
}

// TestWriteJSONLines tests writing a list of House structs to a valid .jl (JSON Lines) file
func TestWriteJSONLines(t *testing.T) {
	// Create test data
	houses := []House{
		{Value: 452600, Income: 8.3252, Age: 41, Rooms: 880, Bedrooms: 129, Pop: 322, HH: 126},
		{Value: 358500, Income: 8.3014, Age: 21, Rooms: 7099, Bedrooms: 1106, Pop: 2401, HH: 1138},
	}

	// Name of the temporary output file
	testFile := "test_output.jl"
	// Ensure the file gets deleted after the test
	defer os.Remove(testFile)

	// Call the function to write JSON Lines
	err := writeJSONLines(testFile, houses)
	if err != nil {
		t.Fatalf("Error writing JSON Lines: %v", err)
	}

	// Read the contents of the output file
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Error reading test output file: %v", err)
	}

	// Split the file content by line
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	// There should be exactly 2 lines (2 JSON objects)
	if len(lines) != 2 {
		t.Errorf("Expected 2 lines in output, got %d", len(lines))
	}

	// Check that the first line is valid JSON
	var h House
	if err := json.Unmarshal([]byte(lines[0]), &h); err != nil {
		t.Errorf("First line is not valid JSON: %v", err)
	}
}
