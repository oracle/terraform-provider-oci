// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// RequestSummarizedMonitoringTemplatesResourcesDetails Filtering criteria data to be specified in the request. Either monitoringTemplateId or compartmentId must be passed even when no other filter property is passed.
type RequestSummarizedMonitoringTemplatesResourcesDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Monitoring Template resource
	MonitoringTemplateId *string `mandatory:"false" json:"monitoringTemplateId"`

	// Multiple resource type to which Monitoring Template applies
	ResourceTypes []string `mandatory:"false" json:"resourceTypes"`

	// Compartment Identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Filter to return monitoring templates based on input enable status i.e. Applied/Not Applied/Partial Applied
	AssociationStatus MonitoringTemplateLifeCycleDetailsEnum `mandatory:"false" json:"associationStatus,omitempty"`

	// The OCID's (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Monitored Resource
	ResourceIds []string `mandatory:"false" json:"resourceIds"`

	// The field to group by
	GroupBy RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum `mandatory:"false" json:"groupBy,omitempty"`

	// Result will ne sorted by this parameter value
	SortBy RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`
}

func (m RequestSummarizedMonitoringTemplatesResourcesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestSummarizedMonitoringTemplatesResourcesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMonitoringTemplateLifeCycleDetailsEnum(string(m.AssociationStatus)); !ok && m.AssociationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssociationStatus: %s. Supported values are: %s.", m.AssociationStatus, strings.Join(GetMonitoringTemplateLifeCycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum(string(m.GroupBy)); !ok && m.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", m.GroupBy, strings.Join(GetRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum Enum with underlying type: string
type RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum string

// Set of constants representing the allowable values for RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum
const (
	RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByMonitoringtemplateid RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum = "monitoringTemplateId"
)

var mappingRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum = map[string]RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum{
	"monitoringTemplateId": RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByMonitoringtemplateid,
}

var mappingRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnumLowerCase = map[string]RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum{
	"monitoringtemplateid": RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByMonitoringtemplateid,
}

// GetRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnumValues Enumerates the set of values for RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum
func GetRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnumValues() []RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum {
	values := make([]RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum, 0)
	for _, v := range mappingRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnumStringValues Enumerates the set of values in String for RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum
func GetRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnumStringValues() []string {
	return []string{
		"monitoringTemplateId",
	}
}

// GetMappingRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum(val string) (RequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnum, bool) {
	enum, ok := mappingRequestSummarizedMonitoringTemplatesResourcesDetailsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum Enum with underlying type: string
type RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum string

// Set of constants representing the allowable values for RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum
const (
	RequestSummarizedMonitoringTemplatesResourcesDetailsSortByCount RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum = "count"
)

var mappingRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum = map[string]RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum{
	"count": RequestSummarizedMonitoringTemplatesResourcesDetailsSortByCount,
}

var mappingRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnumLowerCase = map[string]RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum{
	"count": RequestSummarizedMonitoringTemplatesResourcesDetailsSortByCount,
}

// GetRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnumValues Enumerates the set of values for RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum
func GetRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnumValues() []RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum {
	values := make([]RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum, 0)
	for _, v := range mappingRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnumStringValues Enumerates the set of values in String for RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum
func GetRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnumStringValues() []string {
	return []string{
		"count",
	}
}

// GetMappingRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum(val string) (RequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnum, bool) {
	enum, ok := mappingRequestSummarizedMonitoringTemplatesResourcesDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
