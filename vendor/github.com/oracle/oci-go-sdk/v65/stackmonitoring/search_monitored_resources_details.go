// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// SearchMonitoredResourcesDetails The property search criteria for listing monitored resources.
type SearchMonitoredResourcesDetails struct {

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A filter to return resources that match exact resource name.
	Name *string `mandatory:"false" json:"name"`

	// A filter to return resources that match resource name pattern given. The match is not case sensitive.
	NameContains *string `mandatory:"false" json:"nameContains"`

	// A filter to return resources that match resource type.
	Type *string `mandatory:"false" json:"type"`

	// A filter to return resources with host name match.
	HostName *string `mandatory:"false" json:"hostName"`

	// External resource is any OCI resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	// which is not a Stack Monitoring service resource.
	// Currently supports only following resource types - Container database, non-container database,
	// pluggable database and OCI compute instance.
	ExternalId *string `mandatory:"false" json:"externalId"`

	// A filter to return resources with host name pattern.
	HostNameContains *string `mandatory:"false" json:"hostNameContains"`

	// A filter to return resources with matching management agent id.
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// A filter to return resources with matching lifecycle state.
	LifecycleState ResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// License edition of the monitored resource.
	License LicenseTypeEnum `mandatory:"false" json:"license,omitempty"`

	// Search for resources that were created within a specific date range,
	// using this parameter to specify the earliest creation date for the
	// returned list (inclusive). Specifying this parameter without the
	// corresponding `timeCreatedLessThan` parameter will retrieve resources created from the
	// given `timeCreatedGreaterThanOrEqualTo` to the current time, in "YYYY-MM-ddThh:mmZ" format with a
	// Z offset, as defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" json:"timeCreatedGreaterThanOrEqualTo"`

	// Search for resources that were created within a specific date range,
	// using this parameter to specify the latest creation date for the returned
	// list (exclusive). Specifying this parameter without the corresponding
	// `timeCreatedGreaterThanOrEqualTo` parameter will retrieve all resources created before the
	// specified end date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" json:"timeCreatedLessThan"`

	// Search for resources that were updated within a specific date range,
	// using this parameter to specify the earliest update date for the
	// returned list (inclusive). Specifying this parameter without the
	// corresponding `timeUpdatedLessThan` parameter will retrieve resources updated from the
	// given `timeUpdatedGreaterThanOrEqualTo` to the current time, in "YYYY-MM-ddThh:mmZ" format with a
	// Z offset, as defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeUpdatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" json:"timeUpdatedGreaterThanOrEqualTo"`

	// Search for resources that were updated within a specific date range,
	// using this parameter to specify the latest creation date for the returned
	// list (exclusive). Specifying this parameter without the corresponding
	// `timeUpdatedGreaterThanOrEqualTo` parameter will retrieve all resources updated before the
	// specified end date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeUpdatedLessThan *common.SDKTime `mandatory:"false" json:"timeUpdatedLessThan"`

	// Time zone in the form of tz database canonical zone ID. Specifies the preference with
	// a value that uses the IANA Time Zone Database format (x-obmcs-time-zone).
	// For example - America/Los_Angeles
	ResourceTimeZone *string `mandatory:"false" json:"resourceTimeZone"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for resources is ascending.
	SortBy SearchMonitoredResourcesDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// Criteria based on resource property.
	PropertyEquals map[string]string `mandatory:"false" json:"propertyEquals"`
}

func (m SearchMonitoredResourcesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SearchMonitoredResourcesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseTypeEnum(string(m.License)); !ok && m.License != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for License: %s. Supported values are: %s.", m.License, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSearchMonitoredResourcesDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetSearchMonitoredResourcesDetailsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SearchMonitoredResourcesDetailsSortByEnum Enum with underlying type: string
type SearchMonitoredResourcesDetailsSortByEnum string

// Set of constants representing the allowable values for SearchMonitoredResourcesDetailsSortByEnum
const (
	SearchMonitoredResourcesDetailsSortByTimeCreated  SearchMonitoredResourcesDetailsSortByEnum = "TIME_CREATED"
	SearchMonitoredResourcesDetailsSortByResourceName SearchMonitoredResourcesDetailsSortByEnum = "RESOURCE_NAME"
)

var mappingSearchMonitoredResourcesDetailsSortByEnum = map[string]SearchMonitoredResourcesDetailsSortByEnum{
	"TIME_CREATED":  SearchMonitoredResourcesDetailsSortByTimeCreated,
	"RESOURCE_NAME": SearchMonitoredResourcesDetailsSortByResourceName,
}

var mappingSearchMonitoredResourcesDetailsSortByEnumLowerCase = map[string]SearchMonitoredResourcesDetailsSortByEnum{
	"time_created":  SearchMonitoredResourcesDetailsSortByTimeCreated,
	"resource_name": SearchMonitoredResourcesDetailsSortByResourceName,
}

// GetSearchMonitoredResourcesDetailsSortByEnumValues Enumerates the set of values for SearchMonitoredResourcesDetailsSortByEnum
func GetSearchMonitoredResourcesDetailsSortByEnumValues() []SearchMonitoredResourcesDetailsSortByEnum {
	values := make([]SearchMonitoredResourcesDetailsSortByEnum, 0)
	for _, v := range mappingSearchMonitoredResourcesDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchMonitoredResourcesDetailsSortByEnumStringValues Enumerates the set of values in String for SearchMonitoredResourcesDetailsSortByEnum
func GetSearchMonitoredResourcesDetailsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"RESOURCE_NAME",
	}
}

// GetMappingSearchMonitoredResourcesDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchMonitoredResourcesDetailsSortByEnum(val string) (SearchMonitoredResourcesDetailsSortByEnum, bool) {
	enum, ok := mappingSearchMonitoredResourcesDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
