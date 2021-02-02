// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v35/common"
)

// LogAnalyticsMetaFunctionArgument LogAnalyticsMetaFunctionArgument
type LogAnalyticsMetaFunctionArgument struct {

	// override output fields
	IsOverrideOutputFields *bool `mandatory:"false" json:"isOverrideOutputFields"`

	// argument display name
	ArgumentDisplayName *string `mandatory:"false" json:"argumentDisplayName"`

	// argument example
	ArgumentExample *string `mandatory:"false" json:"argumentExample"`

	// argument service
	ArgumentService *string `mandatory:"false" json:"argumentService"`

	// argument data type
	ArgumentDataType *string `mandatory:"false" json:"argumentDataType"`

	// argument description
	ArgumentDescription *string `mandatory:"false" json:"argumentDescription"`

	// argument name
	ArgumentName *string `mandatory:"false" json:"argumentName"`

	// argument order
	ArgumentOrder *int64 `mandatory:"false" json:"argumentOrder"`

	// argument type
	ArgumentType *int64 `mandatory:"false" json:"argumentType"`

	// meta function id
	ArgumentId *int64 `mandatory:"false" json:"argumentId"`

	// column
	ArgumentLookupColumn *string `mandatory:"false" json:"argumentLookupColumn"`

	// column position
	ArgumentLookupColumnPosition *int64 `mandatory:"false" json:"argumentLookupColumnPosition"`

	// value
	ArgumentValue *string `mandatory:"false" json:"argumentValue"`
}

func (m LogAnalyticsMetaFunctionArgument) String() string {
	return common.PointerString(m)
}
