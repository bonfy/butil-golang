package butil

import "testing"

func TestStringToUint_1(t *testing.T) {
	// 8 can convert to uint
	uid, err := StringToUint("8")
	if err != nil {
		t.Fatalf("err should be nil")
	}
	if uid != 8 {
		t.Fatalf("Converting Error")
	}
}

func TestStringToUint_2(t *testing.T) {
	// -1 can not convert to uint
	_, err := StringToUint("-1")
	// t.Log(err.Error())
	if err == nil {
		t.Fatalf("err should not be nil: -1")
	}
}
