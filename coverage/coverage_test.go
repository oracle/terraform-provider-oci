package coverage

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var totalRgx = regexp.MustCompile(`total:\s+\(statements\)\s+([^"]*)%`)

const CodeCoverageThreshold = 40.0

func TestCoverage(t *testing.T) {
	if os.Getenv("CHECK_COVERAGE") != "true" {
		t.Skip("Skipping coverage check")
	}

	coverageOutput, err := ioutil.ReadFile("coverage.txt")

	if err != nil {
		assert.Fail(t, "unable to parse coverage test")
		return
	}

	// parse coverage output for total
	coverageAmt := totalRgx.FindStringSubmatch(string(coverageOutput))
	if len(coverageAmt) < 2 {
		assert.Fail(t, "no coverage total found")
		return
	}

	iCoverage, err := strconv.ParseFloat(coverageAmt[1], 64)
	if err != nil {
		assert.Fail(t, "unable to parse coverage amount")
		return
	}

	fmt.Println("coverage: ", iCoverage)
	assert.GreaterOrEqual(t, iCoverage, CodeCoverageThreshold)
}
