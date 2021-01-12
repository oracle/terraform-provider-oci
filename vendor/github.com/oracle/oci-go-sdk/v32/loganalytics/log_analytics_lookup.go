// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// LogAnalyticsLookup LogAnalyticsLookup
type LogAnalyticsLookup struct {

	// active edit version
	ActiveEditVersion *int64 `mandatory:"false" json:"activeEditVersion"`

	// canonical link
	CanonicalLink *string `mandatory:"false" json:"canonicalLink"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// fields
	Fields []LookupField `mandatory:"false" json:"fields"`

	// lookupReference
	LookupReference *int64 `mandatory:"false" json:"lookupReference"`

	// iname
	Name *string `mandatory:"false" json:"name"`

	// is built in
	IsBuiltIn *int64 `mandatory:"false" json:"isBuiltIn"`

	// is hidden
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// name
	LookupDisplayName *string `mandatory:"false" json:"lookupDisplayName"`

	ReferringSources *AutoLookups `mandatory:"false" json:"referringSources"`

	StatusSummary *StatusSummary `mandatory:"false" json:"statusSummary"`

	// last updated date
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m LogAnalyticsLookup) String() string {
	return common.PointerString(m)
}
