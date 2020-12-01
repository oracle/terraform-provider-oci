// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// LogAnalyticsParserFilter LogAnalyticsParserFilter
type LogAnalyticsParserFilter struct {

	// id
	Id *string `mandatory:"false" json:"id"`

	Parser *LogAnalyticsParser `mandatory:"false" json:"parser"`

	// agent version
	AgentVersion *string `mandatory:"false" json:"agentVersion"`

	// is in use flag
	IsInUse *int64 `mandatory:"false" json:"isInUse"`

	// operating system
	OperatingSystem *string `mandatory:"false" json:"operatingSystem"`

	// parser Id
	ParserId *int64 `mandatory:"false" json:"parserId"`

	// version
	Version *string `mandatory:"false" json:"version"`
}

func (m LogAnalyticsParserFilter) String() string {
	return common.PointerString(m)
}
