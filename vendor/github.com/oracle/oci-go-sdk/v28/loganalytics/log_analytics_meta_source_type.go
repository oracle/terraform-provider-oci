// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v28/common"
)

// LogAnalyticsMetaSourceType LogAnalyticsMetaSourceType
type LogAnalyticsMetaSourceType struct {

	// built in parser name
	BuiltInParserName *string `mandatory:"false" json:"builtInParserName"`

	// type description
	Description *string `mandatory:"false" json:"description"`

	// display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// entity display name
	EntityDisplayName *string `mandatory:"false" json:"entityDisplayName"`

	// entity name
	EntityName *string `mandatory:"false" json:"entityName"`

	// source type name
	Name *string `mandatory:"false" json:"name"`

	// maximum exclude pattern
	MaximumExcludePattern *int64 `mandatory:"false" json:"maximumExcludePattern"`

	// maximum include pattern
	MaximumIncludePattern *int64 `mandatory:"false" json:"maximumIncludePattern"`
}

func (m LogAnalyticsMetaSourceType) String() string {
	return common.PointerString(m)
}
