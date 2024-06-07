package easycsv

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadColumns(t *testing.T) {
	input := `a,b,c
1,2,3
4,5,6
`
	tests := []struct {
		name    string
		columns []string
		input   string
		want    [][]string
	}{
		{"good-ac", []string{"a", "c"}, input, [][]string{{"1", "3"}, {"4", "6"}}},
		{"good-ba", []string{"b", "a"}, input, [][]string{{"2", "1"}, {"5", "4"}}},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result, err := ReadColumns(strings.NewReader(test.input), test.columns)
			if err != nil {
				t.Errorf("got %v; want nil", err)
			}
			if !reflect.DeepEqual(result, test.want) {
				t.Errorf("got %v; want %v", result, test.want)
			}
		})
	}
	errorTests := []struct {
		name    string
		columns []string
		input   string
	}{
		{"bad-column", []string{"a", "b", "x"}, input},
		{"bad-empty", []string{"a", "b", "c"}, ""},
		{"bad-missing", []string{"a", "b", "c"}, "a,b,c\n1,2\n4,5,6"},
		{"bad-extra", []string{"a", "b", "c"}, "a,b,c\n1,2,3,7\n4,5,6"},
	}
	for _, test := range errorTests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			_, err := ReadColumns(strings.NewReader(test.input), test.columns)
			if err == nil {
				t.Fail()
			}
		})
	}
}
