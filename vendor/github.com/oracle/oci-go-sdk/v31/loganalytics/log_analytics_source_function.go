// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// LogAnalyticsSourceFunction LogAnalyticsSourceFunction
type LogAnalyticsSourceFunction struct {

	// argument
	Arguments []LogAnalyticsMetaFunctionArgument `mandatory:"false" json:"arguments"`

	// enabled flag
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	Function *LogAnalyticsMetaFunction `mandatory:"false" json:"function"`

	// source function Id
	FunctionId *int64 `mandatory:"false" json:"functionId"`

	// source function order
	Order *int64 `mandatory:"false" json:"order"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// column
	LookupColumn *string `mandatory:"false" json:"lookupColumn"`

	// column position
	LookupColumnPosition *int64 `mandatory:"false" json:"lookupColumnPosition"`

	// lookup display name
	LookupDisplayName *string `mandatory:"false" json:"lookupDisplayName"`

	// lookup mode
	LookupMode *int64 `mandatory:"false" json:"lookupMode"`

	// lookup table
	LookupTable *string `mandatory:"false" json:"lookupTable"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`
}

func (m LogAnalyticsSourceFunction) String() string {
	return common.PointerString(m)
}
