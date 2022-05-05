package homework

import (
	"testing"

	"gotest.tools/assert"
)

func TestAverage(t *testing.T) {
	var averageArgs = [15]float32{1, 2, 3, 4, 5, 6, 7, 3.5, 1, 2, 3, 4, 5, 6, 7}
	var averageExpected float32 = 3.5
	var averageResult = average(averageArgs)

	assert.Equal(t, averageExpected, averageResult)
	if averageExpected != averageResult {
		const msg = "Unexpected result" +
			"\tExpected: %e\n\tGot: %e"
		t.Errorf(msg, averageExpected, averageResult)
	}

}
