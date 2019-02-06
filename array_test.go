package main

import (
	"reflect"
	"testing"
)

func TestMaxInt(t *testing.T) {
	input := []int{1, 0, 100, 5, 10, -2, -100, 100}
	expected := 100

	actual := MaxInt(input)

	if actual != expected {
		t.Fatalf("TestMaxInt Failed! (expected: %d actual: %d)", expected, actual)
	}
}

func TestMinInt(t *testing.T) {
	input := []int{1, 0, 100, 5, 10, -2, -100, 100}
	expected := -100

	actual := MinInt(input)

	if actual != expected {
		t.Fatalf("TestMinInt Failed! (expected: %d actual: %d)", expected, actual)
	}
}

func TestReverseInts(t *testing.T) {
	input := []int{1, 0, 100, 5, 10, -2, -100, 100}
	expected := []int{100, -100, -2, 10, 5, 100, 0, 1}

	actual := ReverseInts(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestReverseInts#1 Failed! (expected: %d actual: %d)", expected, actual)
	}

	input = []int{}
	expected = []int{}

	actual = ReverseInts(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestReverseInts#2 Failed! (expected: %d actual: %d)", expected, actual)
	}
}
