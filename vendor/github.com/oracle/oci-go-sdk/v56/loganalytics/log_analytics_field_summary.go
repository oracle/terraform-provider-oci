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

// LogAnalyticsFieldSummary A summary of a field.
type LogAnalyticsFieldSummary struct {

	// The name this field is given in the common event expression standard from mitre.org.
	// This is used for reference when exporting content conforming to CEE standard
	CeeAlias *string `mandatory:"false" json:"ceeAlias"`

	// The field data type.
	DataType *string `mandatory:"false" json:"dataType"`

	// The field default regular expression.
	RegularExpression *string `mandatory:"false" json:"regularExpression"`

	// The field description.
	Description *string `mandatory:"false" json:"description"`

	// The field display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The field edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The facet priority.
	FacetPriority *int64 `mandatory:"false" json:"facetPriority"`

	// The field internal name.
	Name *string `mandatory:"false" json:"name"`

	// A flag inidcating whether or not the facet is elibigle for use.
	IsFacetEligible *bool `mandatory:"false" json:"isFacetEligible"`

	// A flag inidcating whether or not the cardinality of the field is high.
	IsHighCardinality *bool `mandatory:"false" json:"isHighCardinality"`

	// A flag inidcating whether or not the field is a large data field.
	IsLargeData *bool `mandatory:"false" json:"isLargeData"`

	// A flag indicating whether or not the field is multi-valued.
	IsMultiValued *bool `mandatory:"false" json:"isMultiValued"`

	// A flag inidcating whether or not this is a primary field.
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// A flag inidcating whether or not the field can be summarized.
	IsSummarizable *bool `mandatory:"false" json:"isSummarizable"`

	// The mapped value.
	MappedValue *string `mandatory:"false" json:"mappedValue"`

	// A flag inidcating whether or not the field is metric key eligible.
	IsMetricKeyEligible *bool `mandatory:"false" json:"isMetricKeyEligible"`

	// A flag inidcating whether or not the field is metric value eligible.
	IsMetricValueEligible *bool `mandatory:"false" json:"isMetricValueEligible"`

	// A flag inidcating whether or not the field is range facet eligible.
	RangeFacetEligible *int64 `mandatory:"false" json:"rangeFacetEligible"`

	// A flag inidcating whether or not the field is table eligible.
	IsTableEligible *bool `mandatory:"false" json:"isTableEligible"`

	// The field unit type.
	UnitType *string `mandatory:"false" json:"unitType"`
}

func (m LogAnalyticsFieldSummary) String() string {
	return common.PointerString(m)
}
