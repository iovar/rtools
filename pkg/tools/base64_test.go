package tools

import "testing"

func TestBase64Encode(t *testing.T) {
	input := "TestString"
	want := "VGVzdFN0cmluZw=="

	res := Base64Encode(input)

	if res != want {
		t.Errorf("Base64Encode output %s did not match expectation: %s", res, want)
	}
}

func TestBase64Decode(t *testing.T) {
	input := "VGVzdFN0cmluZw=="
	want := "TestString"

	res := Base64Decode(input)

	if res != want {
		t.Errorf("Base64Decode output %s did not match expectation: %s", res, want)
	}
}
