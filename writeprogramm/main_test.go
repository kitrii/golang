package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOkData = `1
2
3
4
5`

var testOkResult = `1
2
3
4
5
`
var testDataNeg = `1
2
1`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOkData))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("Test failed")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("Test failed")
	}
}

func TestOfNegative(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testDataNeg))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("Negative test failed")
	}

}
