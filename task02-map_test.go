package homework

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func smvEval(t *testing.T, smvInput map[int]string, smvExpected []string) {
	var smvResult = sortMapValues(smvInput)

	var cmpResult = reflect.DeepEqual(smvExpected, smvResult)
	assert.Assert(t, cmpResult, "Unexpected result\tExpected: %v\n\tGot: %v",
		smvExpected, smvResult)
}

func TestSortMapValues(t *testing.T) {

	var smvInput = map[int]string{2: "a", 0: "b", 1: "c"}
	var smvExpected = []string{"b", "c", "a"}

	smvEval(t, smvInput, smvExpected)

	smvInput = map[int]string{10: "aa", 0: "bb", 500: "cc"}
	smvExpected = []string{"bb", "aa", "cc"}

	smvEval(t, smvInput, smvExpected)
}
