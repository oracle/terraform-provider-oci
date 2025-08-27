// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestSummarizedMetricExtensionsMetricsDetails Filtering criteria data to be specified in the request. Either metricExtensionId or compartmentId must be passed even when no other filter property is passed.
type RequestSummarizedMetricExtensionsMetricsDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Metric Extension resource
	MetricExtensionId *string `mandatory:"false" json:"metricExtensionId"`

	// Resource type to which Metric Extension applies
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// Compartment Identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Filter for metric extension resources which contain the given metric name
	ContainsMetricWithName *string `mandatory:"false" json:"containsMetricWithName"`

	// Result will ne sorted by this parameter value
	SortBy RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// Sort orders
	SortOrder RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`
}

func (m RequestSummarizedMetricExtensionsMetricsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestSummarizedMetricExtensionsMetricsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRequestSummarizedMetricExtensionsMetricsDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetRequestSummarizedMetricExtensionsMetricsDetailsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum Enum with underlying type: string
type RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum string

// Set of constants representing the allowable values for RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum
const (
	RequestSummarizedMetricExtensionsMetricsDetailsSortByCount RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum = "COUNT"
)

var mappingRequestSummarizedMetricExtensionsMetricsDetailsSortByEnum = map[string]RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum{
	"COUNT": RequestSummarizedMetricExtensionsMetricsDetailsSortByCount,
}

var mappingRequestSummarizedMetricExtensionsMetricsDetailsSortByEnumLowerCase = map[string]RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum{
	"count": RequestSummarizedMetricExtensionsMetricsDetailsSortByCount,
}

// GetRequestSummarizedMetricExtensionsMetricsDetailsSortByEnumValues Enumerates the set of values for RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum
func GetRequestSummarizedMetricExtensionsMetricsDetailsSortByEnumValues() []RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum {
	values := make([]RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum, 0)
	for _, v := range mappingRequestSummarizedMetricExtensionsMetricsDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedMetricExtensionsMetricsDetailsSortByEnumStringValues Enumerates the set of values in String for RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum
func GetRequestSummarizedMetricExtensionsMetricsDetailsSortByEnumStringValues() []string {
	return []string{
		"COUNT",
	}
}

// GetMappingRequestSummarizedMetricExtensionsMetricsDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedMetricExtensionsMetricsDetailsSortByEnum(val string) (RequestSummarizedMetricExtensionsMetricsDetailsSortByEnum, bool) {
	enum, ok := mappingRequestSummarizedMetricExtensionsMetricsDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum Enum with underlying type: string
type RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum string

// Set of constants representing the allowable values for RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum
const (
	RequestSummarizedMetricExtensionsMetricsDetailsSortOrderAsc  RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum = "ASC"
	RequestSummarizedMetricExtensionsMetricsDetailsSortOrderDesc RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum = "DESC"
)

var mappingRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum = map[string]RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum{
	"ASC":  RequestSummarizedMetricExtensionsMetricsDetailsSortOrderAsc,
	"DESC": RequestSummarizedMetricExtensionsMetricsDetailsSortOrderDesc,
}

var mappingRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnumLowerCase = map[string]RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum{
	"asc":  RequestSummarizedMetricExtensionsMetricsDetailsSortOrderAsc,
	"desc": RequestSummarizedMetricExtensionsMetricsDetailsSortOrderDesc,
}

// GetRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnumValues Enumerates the set of values for RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum
func GetRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnumValues() []RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum {
	values := make([]RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum, 0)
	for _, v := range mappingRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnumStringValues Enumerates the set of values in String for RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum
func GetRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum(val string) (RequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnum, bool) {
	enum, ok := mappingRequestSummarizedMetricExtensionsMetricsDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
