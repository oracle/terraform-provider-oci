// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v34/common"
)

// RegexMatchResult RegexMatchResult
type RegexMatchResult struct {

	// matchedLogEntryEndIndex
	MatchedLogEntryEndIndex *int `mandatory:"false" json:"matchedLogEntryEndIndex"`

	// regexScore
	RegexScore *int `mandatory:"false" json:"regexScore"`

	// regexStepsInfo
	RegexStepsInfo []StepInfo `mandatory:"false" json:"regexStepsInfo"`

	// stepCount
	StepCount *int `mandatory:"false" json:"stepCount"`

	// subRegexesMatchInfo
	SubRegexesMatchInfo map[string]MatchInfo `mandatory:"false" json:"subRegexesMatchInfo"`
}

func (m RegexMatchResult) String() string {
	return common.PointerString(m)
}
