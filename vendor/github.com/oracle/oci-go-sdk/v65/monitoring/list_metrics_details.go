// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ListMetricsDetails The request details for retrieving metric definitions. Specify optional properties to filter the returned results.
// Use an asterisk (&#42;) as a wildcard character, placed anywhere in the string.
// For example, to search for all metrics with names that begin with "disk", specify "name" as "disk&#42;".
// If no properties are specified, then all metric definitions within the request scope are returned.
type ListMetricsDetails struct {

	// The metric name to use when searching for metric definitions.
	// Example: `CpuUtilization`
	Name *string `mandatory:"false" json:"name"`

	// The source service or application to use when searching for metric definitions.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"false" json:"namespace"`

	// Resource group that you want to match. A null value returns only metric data that has no resource groups. The specified resource group must exist in the definition of the posted metric. Only one resource group can be applied per metric.
	// A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	// Example: `frontend-fleet`
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// Qualifiers that you want to use when searching for metric definitions.
	// Available dimensions vary by metric namespace. Each dimension takes the form of a key-value pair.
	// Example: `{"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"}`
	DimensionFilters map[string]string `mandatory:"false" json:"dimensionFilters"`

	// Group metrics by these fields in the response. For example, to list all metric namespaces available
	//           in a compartment, groupBy the "namespace" field. Supported fields: namespace, name, resourceGroup.
	// If `groupBy` is used, then `dimensionFilters` is ignored.
	// Example - group by namespace:
	// `[ "namespace" ]`
	GroupBy []string `mandatory:"false" json:"groupBy"`

	// The field to use when sorting returned metric definitions. Only one sorting level is provided.
	// Example: `NAMESPACE`
	SortBy ListMetricsDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// The sort order to use when sorting returned metric definitions. Ascending (ASC) or
	// descending (DESC).
	// Example: `ASC`
	SortOrder ListMetricsDetailsSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`
}

func (m ListMetricsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListMetricsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingListMetricsDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetListMetricsDetailsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMetricsDetailsSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetListMetricsDetailsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMetricsDetailsSortByEnum Enum with underlying type: string
type ListMetricsDetailsSortByEnum string

// Set of constants representing the allowable values for ListMetricsDetailsSortByEnum
const (
	ListMetricsDetailsSortByNamespace     ListMetricsDetailsSortByEnum = "NAMESPACE"
	ListMetricsDetailsSortByName          ListMetricsDetailsSortByEnum = "NAME"
	ListMetricsDetailsSortByResourcegroup ListMetricsDetailsSortByEnum = "RESOURCEGROUP"
)

var mappingListMetricsDetailsSortByEnum = map[string]ListMetricsDetailsSortByEnum{
	"NAMESPACE":     ListMetricsDetailsSortByNamespace,
	"NAME":          ListMetricsDetailsSortByName,
	"RESOURCEGROUP": ListMetricsDetailsSortByResourcegroup,
}

var mappingListMetricsDetailsSortByEnumLowerCase = map[string]ListMetricsDetailsSortByEnum{
	"namespace":     ListMetricsDetailsSortByNamespace,
	"name":          ListMetricsDetailsSortByName,
	"resourcegroup": ListMetricsDetailsSortByResourcegroup,
}

// GetListMetricsDetailsSortByEnumValues Enumerates the set of values for ListMetricsDetailsSortByEnum
func GetListMetricsDetailsSortByEnumValues() []ListMetricsDetailsSortByEnum {
	values := make([]ListMetricsDetailsSortByEnum, 0)
	for _, v := range mappingListMetricsDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetricsDetailsSortByEnumStringValues Enumerates the set of values in String for ListMetricsDetailsSortByEnum
func GetListMetricsDetailsSortByEnumStringValues() []string {
	return []string{
		"NAMESPACE",
		"NAME",
		"RESOURCEGROUP",
	}
}

// GetMappingListMetricsDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetricsDetailsSortByEnum(val string) (ListMetricsDetailsSortByEnum, bool) {
	enum, ok := mappingListMetricsDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMetricsDetailsSortOrderEnum Enum with underlying type: string
type ListMetricsDetailsSortOrderEnum string

// Set of constants representing the allowable values for ListMetricsDetailsSortOrderEnum
const (
	ListMetricsDetailsSortOrderAsc  ListMetricsDetailsSortOrderEnum = "ASC"
	ListMetricsDetailsSortOrderDesc ListMetricsDetailsSortOrderEnum = "DESC"
)

var mappingListMetricsDetailsSortOrderEnum = map[string]ListMetricsDetailsSortOrderEnum{
	"ASC":  ListMetricsDetailsSortOrderAsc,
	"DESC": ListMetricsDetailsSortOrderDesc,
}

var mappingListMetricsDetailsSortOrderEnumLowerCase = map[string]ListMetricsDetailsSortOrderEnum{
	"asc":  ListMetricsDetailsSortOrderAsc,
	"desc": ListMetricsDetailsSortOrderDesc,
}

// GetListMetricsDetailsSortOrderEnumValues Enumerates the set of values for ListMetricsDetailsSortOrderEnum
func GetListMetricsDetailsSortOrderEnumValues() []ListMetricsDetailsSortOrderEnum {
	values := make([]ListMetricsDetailsSortOrderEnum, 0)
	for _, v := range mappingListMetricsDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetricsDetailsSortOrderEnumStringValues Enumerates the set of values in String for ListMetricsDetailsSortOrderEnum
func GetListMetricsDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMetricsDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetricsDetailsSortOrderEnum(val string) (ListMetricsDetailsSortOrderEnum, bool) {
	enum, ok := mappingListMetricsDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
