package homework

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestReverse(t *testing.T) {
	var reverseInput = []int64{1, 2, 5, 15}
	var reverseExpected = []int64{15, 5, 2, 1}

	var reverseResult = reverse(reverseInput)

	var cmpResult = reflect.DeepEqual(reverseExpected, reverseResult)
	assert.Assert(t, cmpResult, "Unexpected result\tExpected: %v\n\tGot: %v",
		reverseExpected, reverseResult)

}
