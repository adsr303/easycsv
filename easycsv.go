// Package easycsv provides some utilities that simplify reading CSV data
// using the [encoding/csv] package.
package easycsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
)

// ReadColumns reads specified CSV columns from reader and returns a slice of records.
// Returned rows contain values in the order provided in columns.
// The header row is not included in the returned result.
// If any of the column names cannot be found in the CSV header row,
// ReadColumns returns an error.
func ReadColumns(reader io.Reader, columns []string) ([][]string, error) {
	var result [][]string
	more := func(values []string) error {
		result = append(result, values)
		return nil
	}
	if err := ReadColumnsFunc(reader, columns, more); err != nil {
		return nil, err
	}
	return result, nil
}

// ReadColumnsFunc reads specified CSV columns from reader and calls f for every record.
// Function f is called for every row with values in the order provided in columns.
// ReadColumnsFunc doesn't call f for the header row.
// If any of the column names cannot be found in the CSV header row,
// ReadColumnsFunc returns an error.
func ReadColumnsFunc(reader io.Reader, columns []string, f func([]string) error) error {
	r := csv.NewReader(reader)
	r.ReuseRecord = true
	record, err := r.Read()
	if err != nil {
		return fmt.Errorf("reading header: %w", err)
	}
	indices := make([]int, len(columns))
	for i, c := range columns {
		index := slices.Index(record, c)
		if index == -1 {
			return fmt.Errorf("column %s not found", c)
		}
		indices[i] = index
	}
	for {
		record, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading data: %w", err)
		}
		r := make([]string, len(indices))
		for i, ii := range indices {
			r[i] = record[ii]
		}
		if err = f(r); err != nil {
			return err
		}
	}
	return nil
}
