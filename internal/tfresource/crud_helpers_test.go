// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

type TestResource struct {
	GetError          error
	GetAttempts       int
	ActualGetAttempts int
}

func (t *TestResource) Get() error {
	t.ActualGetAttempts++
	t.GetAttempts--
	if t.GetAttempts <= 0 {
		return t.GetError
	}
	return nil
}

// issue-routing-tag: terraform/default
func TestUnitWaitForResourceCondition_basic(t *testing.T) {
	if httpreplay.ModeRecordReplay() {
		t.Skip("Skip TestWaitForResourceCondition_basic test in HttpReplay mode.")
	}
	getAttempts := 1
	testResource := &TestResource{GetError: nil, GetAttempts: getAttempts}
	finalStateFunc := func() bool {
		return testResource.GetAttempts == 0
	}

	// Test normal case
	err := WaitForResourceCondition(testResource, finalStateFunc, 0)
	if err != nil {
		t.Errorf("Got unexpected error '%q' from single attempt", err)
		return
	}

	// Test normal case with multiple attempts
	testResource = &TestResource{GetError: nil, GetAttempts: 3}
	err = WaitForResourceCondition(testResource, finalStateFunc, time.Minute)
	if err != nil {
		t.Errorf("Got unexpected error '%q' from multiple attempts", err)
		return
	}

	// Test case where Get returns error after 1 attempt
	testResource = &TestResource{GetError: fmt.Errorf("GetError"), GetAttempts: 1}
	err = WaitForResourceCondition(testResource, finalStateFunc, 0)
	if err == nil || !strings.HasPrefix(err.Error(), "GetError") {
		t.Errorf("Got unexpected error '%q' after single attempt, expected a GetError", err)
		return
	}

	// Test case where Get returns error after multiple attempts
	testResource = &TestResource{GetError: fmt.Errorf("GetError"), GetAttempts: 3}
	err = WaitForResourceCondition(testResource, finalStateFunc, time.Minute)
	if err == nil || !strings.HasPrefix(err.Error(), "GetError") {
		t.Errorf("Got unexpected error '%q' after multiple attempts, expected a GetError", err)
		return
	}

	// Test timing out with zero timeout duration
	testResource = &TestResource{GetError: nil, GetAttempts: 10}
	err = WaitForResourceCondition(testResource, finalStateFunc, 0)
	if err == nil || !strings.HasPrefix(err.Error(), "Timed out") {
		t.Errorf("Got unexpected error '%q' after a single attempt, expected a timeout error", err)
		return
	}

	// Test timing out with non-zero timeout duration, also validate that we got expected number of Get attempts due to exponential backoff
	testResource = &TestResource{GetError: nil, GetAttempts: 10}
	err = WaitForResourceCondition(testResource, finalStateFunc, 20*time.Second)
	if err == nil || !strings.HasPrefix(err.Error(), "Timed out") {
		t.Errorf("Got unexpected error '%q' after a single attempt, expected a timeout error", err)
		return
	}

	// Expected Get attempts at: 0, 2, 6, 14, 20 seconds
	if testResource.ActualGetAttempts != 5 {
		t.Errorf("Expected 5 Get attempts, got %d instead", testResource.ActualGetAttempts)
		return
	}
}
