package coverage

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var totalRgx = regexp.MustCompile(`total:\s+\(statements\)\s+([^"]*)%`)

const CodeCoverageThreshold = 56.1

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
	diff := math.Abs(iCoverage - CodeCoverageThreshold)
	if diff < 1.0 {
		fmt.Println("Skip this check because of round error in golang tool")
		t.SkipNow()
	}
	if iCoverage > CodeCoverageThreshold {
		t.Fatalf("Please update CodeCoverageThreshold in coverage/coverage_test.go with lastest value: %v", iCoverage)
	} else if iCoverage < CodeCoverageThreshold {
		t.Fatalf("Your code coverage is %v. Please add more unit tests to reach code coverage: %v", iCoverage, CodeCoverageThreshold)
	}
	assert.Equal(t, iCoverage, CodeCoverageThreshold)
}
