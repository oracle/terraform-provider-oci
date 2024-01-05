// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the download engine of the Java Management Service.
//

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestSummarizedJavaDownloadCountsDetails Attributes to summarize the Java download counts in a tenancy.
type RequestSummarizedJavaDownloadCountsDetails struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) here should be the tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Group as property specifying the aggregation type for download counts.
	GroupAs JavaDownloadCountAggregationTypeEnum `mandatory:"true" json:"groupAs"`

	// Unique Java family version identifier.
	FamilyVersion *string `mandatory:"false" json:"familyVersion"`

	// Unique Java release version identifier.
	ReleaseVersion *string `mandatory:"false" json:"releaseVersion"`

	// The start time from when download data has to be included (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The end time until when the download data has to be included (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The property to be used for sorting the aggregated report.
	SortBy AggregationSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// The sort order for the aggregated report.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" json:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous call.
	Page *string `mandatory:"false" json:"page"`
}

func (m RequestSummarizedJavaDownloadCountsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestSummarizedJavaDownloadCountsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJavaDownloadCountAggregationTypeEnum(string(m.GroupAs)); !ok && m.GroupAs != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupAs: %s. Supported values are: %s.", m.GroupAs, strings.Join(GetJavaDownloadCountAggregationTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAggregationSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetAggregationSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
