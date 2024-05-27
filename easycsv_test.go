package easycsv

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadColumns(t *testing.T) {
	input := `a,b,c
1,2,3
4,5,6`
	want := [][]string{{"1", "3"}, {"4", "6"}}
	result, err := ReadColumns(strings.NewReader(input), []string{"a", "c"})
	if err != nil {
		t.Errorf("got %v; want nil", err)
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("got %v; want %v", result, want)
	}
}
