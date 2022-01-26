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

// LogAnalyticsCategory A category into which resources can be placed.
type LogAnalyticsCategory struct {

	// The unique name that identifies the category.
	Name *string `mandatory:"false" json:"name"`

	// The category description.
	Description *string `mandatory:"false" json:"description"`

	// The category display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The category type. Values include "PRODUCT", "TIER", "VENDOR" and "GENERIC".
	Type *string `mandatory:"false" json:"type"`

	// The system flag. A value of false denotes a user-created
	// category. A value of true denotes an Oracle-defined category.
	IsSystem *bool `mandatory:"false" json:"isSystem"`
}

func (m LogAnalyticsCategory) String() string {
	return common.PointerString(m)
}
