// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsMetaFunctionArgument LogAnalyticsMetaFunctionArgument
type LogAnalyticsMetaFunctionArgument struct {

	// The override output fields.
	IsOverrideOutputFields *bool `mandatory:"false" json:"isOverrideOutputFields"`

	// The argument display name.
	ArgumentDisplayName *string `mandatory:"false" json:"argumentDisplayName"`

	// The argument example.
	ArgumentExample *string `mandatory:"false" json:"argumentExample"`

	// The argument service.
	ArgumentService *string `mandatory:"false" json:"argumentService"`

	// The argument data type.
	ArgumentDataType *string `mandatory:"false" json:"argumentDataType"`

	// The argument description.
	ArgumentDescription *string `mandatory:"false" json:"argumentDescription"`

	// The argument name.
	ArgumentName *string `mandatory:"false" json:"argumentName"`

	// The argument order.
	ArgumentOrder *int64 `mandatory:"false" json:"argumentOrder"`

	// The argument type.
	ArgumentType *int64 `mandatory:"false" json:"argumentType"`

	// The argument unique identifier.
	ArgumentId *int64 `mandatory:"false" json:"argumentId"`

	// The lookup column.
	ArgumentLookupColumn *string `mandatory:"false" json:"argumentLookupColumn"`

	// The lookup column position.
	ArgumentLookupColumnPosition *int64 `mandatory:"false" json:"argumentLookupColumnPosition"`

	// The argument value.
	ArgumentValue *string `mandatory:"false" json:"argumentValue"`

	// The argument unique identifier as a string.
	ArgumentReference *string `mandatory:"false" json:"argumentReference"`
}

func (m LogAnalyticsMetaFunctionArgument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsMetaFunctionArgument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
