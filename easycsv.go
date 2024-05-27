package easycsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
)

func ReadColumns(reader io.Reader, columns []string) ([][]string, error) {
	r := csv.NewReader(reader)
	record, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("reading header: %w", err)
	}
	indices := make([]int, len(columns))
	for i, c := range columns {
		index := slices.Index(record, c)
		if index == -1 {
			return nil, fmt.Errorf("column %s not found", c)
		}
		indices[i] = index
	}
	result := make([][]string, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading data: %w", err)
		}
		r := make([]string, len(indices))
		for i, ii := range indices {
			r[i] = record[ii]
		}
		result = append(result, r)
	}
	return result, nil
}
