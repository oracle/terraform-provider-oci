// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
