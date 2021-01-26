// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v34/common"
)

// LogAnalyticsSourceEntityType LogAnalyticsSourceEntityType
type LogAnalyticsSourceEntityType struct {

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// Entity type
	EntityType *string `mandatory:"false" json:"entityType"`

	// type category
	EntityTypeCategory *string `mandatory:"false" json:"entityTypeCategory"`

	// Entity type display name
	EntityTypeDisplayName *string `mandatory:"false" json:"entityTypeDisplayName"`
}

func (m LogAnalyticsSourceEntityType) String() string {
	return common.PointerString(m)
}
