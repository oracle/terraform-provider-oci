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

// LogAnalyticsResourceCategory A resource and its category.
type LogAnalyticsResourceCategory struct {

	// The unique identifier of the resource, usually a name or ocid.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The resource type.
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// The category name to which this resource belongs.
	CategoryName *string `mandatory:"false" json:"categoryName"`

	// The system flag. A value of false denotes a user-created category assignment.
	// A value of true denotes an Oracle-defined category assignment.
	IsSystem *bool `mandatory:"false" json:"isSystem"`
}

func (m LogAnalyticsResourceCategory) String() string {
	return common.PointerString(m)
}
