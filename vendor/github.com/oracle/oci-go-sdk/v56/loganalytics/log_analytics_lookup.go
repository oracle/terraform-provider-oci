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

// LogAnalyticsLookup LogAnalyticsLookup
type LogAnalyticsLookup struct {

	// The active edit version.
	ActiveEditVersion *int64 `mandatory:"false" json:"activeEditVersion"`

	// The canonical link.
	CanonicalLink *string `mandatory:"false" json:"canonicalLink"`

	// The lookup description.
	Description *string `mandatory:"false" json:"description"`

	// The edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The lookup fields.
	Fields []LookupField `mandatory:"false" json:"fields"`

	// The lookup reference as an integer.
	LookupReference *int64 `mandatory:"false" json:"lookupReference"`

	// The lookup reference as a string.
	LookupReferenceString *string `mandatory:"false" json:"lookupReferenceString"`

	// The lookup type.  Valid values are LOOKUP or DICTIONARY.
	Type LogAnalyticsLookupTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The lookup name.
	Name *string `mandatory:"false" json:"name"`

	// A flag indicating if the lookup is custom (user-defined) or
	// built in.
	IsBuiltIn *int64 `mandatory:"false" json:"isBuiltIn"`

	// A flag indicating if the lookup is hidden or not.  A hidden lookup will
	// not be returned in list operations by default.
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// The lookup display name.
	LookupDisplayName *string `mandatory:"false" json:"lookupDisplayName"`

	ReferringSources *AutoLookups `mandatory:"false" json:"referringSources"`

	StatusSummary *StatusSummary `mandatory:"false" json:"statusSummary"`

	// The last updated date.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An array of categories assigned to this lookup.
	// The isSystem flag denotes if each category assignment is user-created or Oracle-defined.
	Categories []LogAnalyticsCategory `mandatory:"false" json:"categories"`
}

func (m LogAnalyticsLookup) String() string {
	return common.PointerString(m)
}

// LogAnalyticsLookupTypeEnum Enum with underlying type: string
type LogAnalyticsLookupTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsLookupTypeEnum
const (
	LogAnalyticsLookupTypeLookup     LogAnalyticsLookupTypeEnum = "Lookup"
	LogAnalyticsLookupTypeDictionary LogAnalyticsLookupTypeEnum = "Dictionary"
)

var mappingLogAnalyticsLookupType = map[string]LogAnalyticsLookupTypeEnum{
	"Lookup":     LogAnalyticsLookupTypeLookup,
	"Dictionary": LogAnalyticsLookupTypeDictionary,
}

// GetLogAnalyticsLookupTypeEnumValues Enumerates the set of values for LogAnalyticsLookupTypeEnum
func GetLogAnalyticsLookupTypeEnumValues() []LogAnalyticsLookupTypeEnum {
	values := make([]LogAnalyticsLookupTypeEnum, 0)
	for _, v := range mappingLogAnalyticsLookupType {
		values = append(values, v)
	}
	return values
}
