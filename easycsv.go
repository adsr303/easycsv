package easycsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
)

func ReadColumns(reader io.Reader, columns []string) ([][]string, error) {
	result := make([][]string, 0)
	more := func(values []string) error {
		result = append(result, values)
		return nil
	}
	err := ReadColumnsFunc(reader, columns, more)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ReadColumnsFunc(reader io.Reader, columns []string, f func([]string) error) error {
	r := csv.NewReader(reader)
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
		record, err := r.Read()
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
		err = f(r)
		if err != nil {
			return err
		}
	}
	return nil
}
