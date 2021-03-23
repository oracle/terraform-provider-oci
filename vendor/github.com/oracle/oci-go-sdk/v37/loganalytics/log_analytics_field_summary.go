// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v37/common"
)

// LogAnalyticsFieldSummary summary of fields
type LogAnalyticsFieldSummary struct {

	// The name this field is given in the common event expression standard from mitre.org.
	// This is used for reference when exporting content conforming to CEE standard
	CeeAlias *string `mandatory:"false" json:"ceeAlias"`

	// data type
	DataType *string `mandatory:"false" json:"dataType"`

	// default regular expression
	RegularExpression *string `mandatory:"false" json:"regularExpression"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// facet priority
	FacetPriority *int64 `mandatory:"false" json:"facetPriority"`

	// internal name
	Name *string `mandatory:"false" json:"name"`

	// is facet eligible flag
	IsFacetEligible *bool `mandatory:"false" json:"isFacetEligible"`

	// is high cardinality flag
	IsHighCardinality *bool `mandatory:"false" json:"isHighCardinality"`

	// is larget data flag
	IsLargeData *bool `mandatory:"false" json:"isLargeData"`

	// is multi-valued flag
	IsMultiValued *bool `mandatory:"false" json:"isMultiValued"`

	// is primary flag
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// is summarizable flag
	IsSummarizable *bool `mandatory:"false" json:"isSummarizable"`

	// mappable
	MappedValue *string `mandatory:"false" json:"mappedValue"`

	// metric key eligible
	IsMetricKeyEligible *bool `mandatory:"false" json:"isMetricKeyEligible"`

	// metric value eligible
	IsMetricValueEligible *bool `mandatory:"false" json:"isMetricValueEligible"`

	// range facet eligible
	RangeFacetEligible *int64 `mandatory:"false" json:"rangeFacetEligible"`

	// table eligible
	IsTableEligible *bool `mandatory:"false" json:"isTableEligible"`

	// unit type
	UnitType *string `mandatory:"false" json:"unitType"`
}

func (m LogAnalyticsFieldSummary) String() string {
	return common.PointerString(m)
}
