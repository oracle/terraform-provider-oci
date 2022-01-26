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

// LogAnalyticsMetaSourceType LogAnalyticsMetaSourceType
type LogAnalyticsMetaSourceType struct {

	// The built in parser name.
	BuiltInParserName *string `mandatory:"false" json:"builtInParserName"`

	// The source type description.
	Description *string `mandatory:"false" json:"description"`

	// The source type display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The entity display name.
	EntityDisplayName *string `mandatory:"false" json:"entityDisplayName"`

	// The entity internal name.
	EntityName *string `mandatory:"false" json:"entityName"`

	// The source type name.
	Name *string `mandatory:"false" json:"name"`

	// The maximum exclude pattern.
	MaximumExcludePattern *int64 `mandatory:"false" json:"maximumExcludePattern"`

	// The maximum include pattern.
	MaximumIncludePattern *int64 `mandatory:"false" json:"maximumIncludePattern"`
}

func (m LogAnalyticsMetaSourceType) String() string {
	return common.PointerString(m)
}
