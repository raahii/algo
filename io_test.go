package algo

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

func TestReadIntF(t *testing.T) {
	f, err := TempFile("3\n")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	actual := ReadIntF(f)

	expected := 3
	assert.Equal(t, expected, actual)
}

func TestReadIntsF(t *testing.T) {
	f, err := TempFile("3 5 7\n")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	actual := ReadIntsF(f, 3)

	expected := []int{3, 5, 7}
	assert.Equal(t, expected, actual)
}

func TestReadLineF(t *testing.T) {
	str := "010101001001001000001010101010011111111111101100101000010101001000000000\n"
	f, err := TempFile(str)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expected := "010101001001001000001010101010011111111111101100101000010101001000000000"
	actual := ReadLineF(f)
	assert.Equal(t, expected, actual)
}

func TestReadLinesF(t *testing.T) {
	str := "オッス！おら悟空！\nいっちょやってみっか！\n"
	f, err := TempFile(str)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expected := []string{
		"オッス！おら悟空！",
		"いっちょやってみっか！",
	}
	actual := ReadLinesF(f, 2)
	assert.Equal(t, expected, actual)
}
