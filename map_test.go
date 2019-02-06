package main

import (
	"reflect"
	"testing"
)

func TestSortMapByValue(t *testing.T) {
	input := map[string]int{
		"Yui Aragaki":      30,
		"Masami Nagasawa":  31,
		"Kasumi Arimura":   25,
		"Mitsuki Takahata": 27,
		"Satomi Ishihara":  32,
		"Ryoko Hirosue":    38,
		"Miho Kanno":       41,
		"Kyoko Hukada":     36,
		"Haruka Ayase":     33,
		"Kanna Hashimoto":  20,
	}
	expected := PairList{
		Pair{"Miho Kanno", 41},
		Pair{"Ryoko Hirosue", 38},
		Pair{"Kyoko Hukada", 36},
		Pair{"Haruka Ayase", 33},
		Pair{"Satomi Ishihara", 32},
		Pair{"Masami Nagasawa", 31},
		Pair{"Yui Aragaki", 30},
		Pair{"Mitsuki Takahata", 27},
		Pair{"Kasumi Arimura", 25},
		Pair{"Kanna Hashimoto", 20},
	}

	actual := sortMapByValue(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("TestSortMapByValue#1 Failed! \nexpected: %v, \nactual: %v\n)", expected, actual)
	}
}
