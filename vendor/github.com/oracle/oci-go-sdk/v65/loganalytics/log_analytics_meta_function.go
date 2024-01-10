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

// LogAnalyticsMetaFunction LogAnalyticsMetaFunction
type LogAnalyticsMetaFunction struct {

	// An array of meta function arguments.
	MetaFunctionArgument []LogAnalyticsMetaFunctionArgument `mandatory:"false" json:"metaFunctionArgument"`

	// The component.
	Component *string `mandatory:"false" json:"component"`

	// The description.
	Description *string `mandatory:"false" json:"description"`

	// The edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The meta function unique identifier.
	MetaFunctionId *int64 `mandatory:"false" json:"metaFunctionId"`

	// The java class name.
	JavaClassName *string `mandatory:"false" json:"javaClassName"`

	// The meta function name.
	Name *string `mandatory:"false" json:"name"`
}

func (m LogAnalyticsMetaFunction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsMetaFunction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
