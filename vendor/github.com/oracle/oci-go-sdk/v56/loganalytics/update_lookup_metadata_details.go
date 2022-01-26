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

// UpdateLookupMetadataDetails UpdateLookupMetadataDetails
type UpdateLookupMetadataDetails struct {

	// The default match value.
	DefaultMatchValue *string `mandatory:"false" json:"defaultMatchValue"`

	// The lookup description.
	Description *string `mandatory:"false" json:"description"`

	// The lookup fields.
	Fields []LogAnalyticsLookupFields `mandatory:"false" json:"fields"`

	// The maximum number of matches.
	MaxMatches *int64 `mandatory:"false" json:"maxMatches"`

	// An array of categories to assign to the lookup. Specifying the name attribute for each category would suffice.
	// Oracle-defined category assignments cannot be removed.
	Categories []LogAnalyticsCategory `mandatory:"false" json:"categories"`
}

func (m UpdateLookupMetadataDetails) String() string {
	return common.PointerString(m)
}
