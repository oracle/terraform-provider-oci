// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// LogAnalyticsMetaFunction LogAnalyticsMetaFunction
type LogAnalyticsMetaFunction struct {

	// meta function argument object
	MetaFunctionArgument []LogAnalyticsMetaFunctionArgument `mandatory:"false" json:"metaFunctionArgument"`

	// component
	Component *string `mandatory:"false" json:"component"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// meta function Id
	MetaFunctionId *int64 `mandatory:"false" json:"metaFunctionId"`

	// java class name
	JavaClassName *string `mandatory:"false" json:"javaClassName"`

	// meta function name
	Name *string `mandatory:"false" json:"name"`
}

func (m LogAnalyticsMetaFunction) String() string {
	return common.PointerString(m)
}
