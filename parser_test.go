package parser

import (
	"reflect"
	"testing"
)

func expectNotError(t *testing.T, err error) {
	if err != nil {
		t.Errorf(`Fail, expecting not error but got err: %v`, err)
	}
}
func TestDecodeGreyScale(t *testing.T) {
	output, err := DecodeGreyScale([]byte{1, 1, 0, 0, 0, 0, 1}, 7)
	expectNotError(t, err)
	if output != "a" {
		t.Errorf(`Fail, expecting 'a', output %s`, output)
	}
	output, err = DecodeGreyScale([]byte{31, 31, 20, 20, 20, 21, 0}, 7)
	expectNotError(t, err)
	if output != "b" {
		t.Errorf(`Fail, expecting 'b', output %s`, output)
	}
	output, err = DecodeGreyScale([]byte{1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1}, 7)
	expectNotError(t, err)
	if output != "aa" {
		t.Errorf(`Fail, expecting 'aa', output %s`, output)
	}
	_, err = DecodeGreyScale([]byte{1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1}, 7)
	if err == nil {
		t.Error(`Fail, expecting error`)
	}
}

func TestBytesStringToBytesSlice(t *testing.T) {
	expected := []byte{1, 1, 0, 1, 0, 1, 0}
	output, err := bytesStringToBytesSlice("1 1 0 1 0 1 0", " ")
	expectNotError(t, err)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf(`Fail, expecting %v but got %v`, expected, output)
	}
	output, err = bytesStringToBytesSlice("1,1,0,1,0,1,0", ",")
	expectNotError(t, err)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf(`Fail, expecting %v but got %v`, expected, output)
	}
	output, err = bytesStringToBytesSlice("1,1,0,1,0,1,a", ",")
	if err == nil {
		t.Error(`Fail, expecting error`)
	}
}

//1101010 1110101 1101101 1100001 1110100 1101000 1100101 1101011 0100000 1100011 1101111 1101101 1100101 1100010 1100001 1100011 1101011
func TestParse7BitStringGreyScale(t *testing.T) {
	output, err := Parse7BitStringGreyScale("1 1 0 1 0 1 0 1 1 1 0 1 0 1 1 1 0 1 1 0 1 1 1 0 0 0 0 1 1 1 1 0 1 0 0 1 1 0 1 0 0 0 1 1 0 0 1 0 1 1 1 0 1 0 1 1 0 1 0 0 0 0 0 1 1 0 0 0 1 1 1 1 0 1 1 1 1 1 1 0 1 1 0 1 1 1 0 0 1 0 1 1 1 0 0 0 1 0 1 1 0 0 0 0 1 1 1 0 0 0 1 1 1 1 0 1 0 1 1")
	expectNotError(t, err)
	if output != "jumathek comeback" {
		t.Errorf(`Fail, expecting 'jumathek comeback', output %s`, output)
	}
	output, err = Parse7BitStringGreyScale("1 a 0 1 0 1 0 1 1 1 0 1 0 1 1 1 0 1 1 0 1 1 1 0 0 0 0 1 1 1 1 0 1 0 0 1 1 0 1 0 0 0 1 1 0 0 1 0 1 1 1 0 1 0 1 1 0 1 0 0 0 0 0 1 1 0 0 0 1 1 1 1 0 1 1 1 1 1 1 0 1 1 0 1 1 1 0 0 1 0 1 1 1 0 0 0 1 0 1 1 0 0 0 0 1 1 1 0 0 0 1 1 1 1 0 1 0 1 1")
	if err == nil {
		t.Error(`Fail, expecting error`)
	}
	output, err = Parse7BitStringGreyScale("1 1 0 1 0 1 0 1 1 1 0 1 0 1 1 1 0 1 1 0 1 1 1 0 0 0 0 1 1 1 1 0 1 0 0 1 1 0 1 0 0 0 1 1 0 0 1 0 1 1 1 0 1 0 1 1 0 1 0 0 0 0 0 1 1 0 0 0 1 1 1 1 0 1 1 1 1 1 1 0 1 1 0 1 1 1 0 0 1 0 1 1 1 0 0 0 1 0 1 1 0 0 0 0 1 1 1 0 0 0 1 1 1 1 0 1 0 1")
	if err == nil {
		t.Error(`Fail, expecting error`)
	}
}
