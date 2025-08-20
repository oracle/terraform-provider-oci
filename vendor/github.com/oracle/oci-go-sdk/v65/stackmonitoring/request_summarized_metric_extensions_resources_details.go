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

// RequestSummarizedMetricExtensionsResourcesDetails Filtering criteria data to be specified in the request. Either metricExtensionId or compartmentId must be passed even when no other filter property is passed.
type RequestSummarizedMetricExtensionsResourcesDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Metric Extension resource
	MetricExtensionId *string `mandatory:"false" json:"metricExtensionId"`

	// Resource type to which Metric Extension applies
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// Compartment Identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Filter to return metric extensions based on input enable status i.e. Enabled/Disabled
	AssociationStatus RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum `mandatory:"false" json:"associationStatus,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Monitored Resource
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The field to group by
	GroupBy RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum `mandatory:"false" json:"groupBy,omitempty"`

	// Result will ne sorted by this parameter value
	SortBy RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// Sort orders
	SortOrder RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`
}

func (m RequestSummarizedMetricExtensionsResourcesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestSummarizedMetricExtensionsResourcesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum(string(m.AssociationStatus)); !ok && m.AssociationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssociationStatus: %s. Supported values are: %s.", m.AssociationStatus, strings.Join(GetRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum(string(m.GroupBy)); !ok && m.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", m.GroupBy, strings.Join(GetRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestSummarizedMetricExtensionsResourcesDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetRequestSummarizedMetricExtensionsResourcesDetailsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum Enum with underlying type: string
type RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum string

// Set of constants representing the allowable values for RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum
const (
	RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnabled  RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum = "ENABLED"
	RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusDisabled RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum = "DISABLED"
)

var mappingRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum = map[string]RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum{
	"ENABLED":  RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnabled,
	"DISABLED": RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusDisabled,
}

var mappingRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnumLowerCase = map[string]RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum{
	"enabled":  RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnabled,
	"disabled": RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusDisabled,
}

// GetRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnumValues Enumerates the set of values for RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum
func GetRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnumValues() []RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum {
	values := make([]RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum, 0)
	for _, v := range mappingRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnumStringValues Enumerates the set of values in String for RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum
func GetRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum(val string) (RequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnum, bool) {
	enum, ok := mappingRequestSummarizedMetricExtensionsResourcesDetailsAssociationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum Enum with underlying type: string
type RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum string

// Set of constants representing the allowable values for RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum
const (
	RequestSummarizedMetricExtensionsResourcesDetailsGroupByMetricExtensionId RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum = "METRIC_EXTENSION_ID"
)

var mappingRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum = map[string]RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum{
	"METRIC_EXTENSION_ID": RequestSummarizedMetricExtensionsResourcesDetailsGroupByMetricExtensionId,
}

var mappingRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnumLowerCase = map[string]RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum{
	"metric_extension_id": RequestSummarizedMetricExtensionsResourcesDetailsGroupByMetricExtensionId,
}

// GetRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnumValues Enumerates the set of values for RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum
func GetRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnumValues() []RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum {
	values := make([]RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum, 0)
	for _, v := range mappingRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnumStringValues Enumerates the set of values in String for RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum
func GetRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnumStringValues() []string {
	return []string{
		"METRIC_EXTENSION_ID",
	}
}

// GetMappingRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum(val string) (RequestSummarizedMetricExtensionsResourcesDetailsGroupByEnum, bool) {
	enum, ok := mappingRequestSummarizedMetricExtensionsResourcesDetailsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum Enum with underlying type: string
type RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum string

// Set of constants representing the allowable values for RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum
const (
	RequestSummarizedMetricExtensionsResourcesDetailsSortByCount RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum = "COUNT"
)

var mappingRequestSummarizedMetricExtensionsResourcesDetailsSortByEnum = map[string]RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum{
	"COUNT": RequestSummarizedMetricExtensionsResourcesDetailsSortByCount,
}

var mappingRequestSummarizedMetricExtensionsResourcesDetailsSortByEnumLowerCase = map[string]RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum{
	"count": RequestSummarizedMetricExtensionsResourcesDetailsSortByCount,
}

// GetRequestSummarizedMetricExtensionsResourcesDetailsSortByEnumValues Enumerates the set of values for RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum
func GetRequestSummarizedMetricExtensionsResourcesDetailsSortByEnumValues() []RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum {
	values := make([]RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum, 0)
	for _, v := range mappingRequestSummarizedMetricExtensionsResourcesDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedMetricExtensionsResourcesDetailsSortByEnumStringValues Enumerates the set of values in String for RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum
func GetRequestSummarizedMetricExtensionsResourcesDetailsSortByEnumStringValues() []string {
	return []string{
		"COUNT",
	}
}

// GetMappingRequestSummarizedMetricExtensionsResourcesDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedMetricExtensionsResourcesDetailsSortByEnum(val string) (RequestSummarizedMetricExtensionsResourcesDetailsSortByEnum, bool) {
	enum, ok := mappingRequestSummarizedMetricExtensionsResourcesDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum Enum with underlying type: string
type RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum string

// Set of constants representing the allowable values for RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum
const (
	RequestSummarizedMetricExtensionsResourcesDetailsSortOrderAsc  RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum = "ASC"
	RequestSummarizedMetricExtensionsResourcesDetailsSortOrderDesc RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum = "DESC"
)

var mappingRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum = map[string]RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum{
	"ASC":  RequestSummarizedMetricExtensionsResourcesDetailsSortOrderAsc,
	"DESC": RequestSummarizedMetricExtensionsResourcesDetailsSortOrderDesc,
}

var mappingRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnumLowerCase = map[string]RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum{
	"asc":  RequestSummarizedMetricExtensionsResourcesDetailsSortOrderAsc,
	"desc": RequestSummarizedMetricExtensionsResourcesDetailsSortOrderDesc,
}

// GetRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnumValues Enumerates the set of values for RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum
func GetRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnumValues() []RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum {
	values := make([]RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum, 0)
	for _, v := range mappingRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnumStringValues Enumerates the set of values in String for RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum
func GetRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum(val string) (RequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnum, bool) {
	enum, ok := mappingRequestSummarizedMetricExtensionsResourcesDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
