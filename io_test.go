package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TempFile(str string) (*os.File, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(f, str)
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func TestReadInt(t *testing.T) {
	f, err := TempFile("3\n")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	actual := ReadInt(f)

	expected := 3
	assert.Equal(t, expected, actual)
}

func TestReadInts(t *testing.T) {
	f, err := TempFile("3 5 7\n")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	actual := ReadInts(f, 3)

	expected := []int{3, 5, 7}
	assert.Equal(t, expected, actual)
}

func TestReadLine(t *testing.T) {
	expected := "おっす！おら悟空！"
	f, err := TempFile(expected + "\n")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	actual := ReadLine(f)

	assert.Equal(t, expected, actual)
}

func TestReadWords(t *testing.T) {
	f, err := TempFile("hop step jump\n")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	actual := ReadWords(f, 3)

	expected := []string{"hop", "step", "jump"}
	assert.Equal(t, expected, actual)
}
