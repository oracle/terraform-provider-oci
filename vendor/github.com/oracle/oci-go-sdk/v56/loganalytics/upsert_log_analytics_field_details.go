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

// UpsertLogAnalyticsFieldDetails Upsert LogAnalytics Field Details
type UpsertLogAnalyticsFieldDetails struct {

	// The data type.
	DataType *string `mandatory:"false" json:"dataType"`

	// A flag indicating whether or not the field is multi-valued.
	IsMultiValued *bool `mandatory:"false" json:"isMultiValued"`

	// The field description.
	Description *string `mandatory:"false" json:"description"`

	// The field display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The field internal name.
	Name *string `mandatory:"false" json:"name"`
}

func (m UpsertLogAnalyticsFieldDetails) String() string {
	return common.PointerString(m)
}
