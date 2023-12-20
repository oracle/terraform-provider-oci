// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ListIpInventoryDetails Input Parameters for retrieving Ip Inventory Data within specified compartments of a region.
type ListIpInventoryDetails struct {

	// The List of Regions selected
	RegionList []string `mandatory:"true" json:"regionList"`

	// The list of OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartments.
	CompartmentList []string `mandatory:"true" json:"compartmentList"`

	// The selected filters
	OverrideFilters *bool `mandatory:"false" json:"overrideFilters"`

	// The CIDR(s) Utilization of a VCN
	Utilization *float32 `mandatory:"false" json:"utilization"`

	// The List of Overlapping VCN's.
	OverlappingVcnsOnly *bool `mandatory:"false" json:"overlappingVcnsOnly"`

	// List of Address types of the IP Consumed by Customer
	AddressTypeList []AddressTypeEnum `mandatory:"false" json:"addressTypeList"`

	// List of Resource types of the VCN
	ResourceTypeList []ListIpInventoryDetailsResourceTypeListEnum `mandatory:"false" json:"resourceTypeList,omitempty"`

	// filter the results for the given searchKey
	SearchKeyword *string `mandatory:"false" json:"searchKeyword"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListIpInventoryDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListIpInventoryDetailsSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	PaginationOffset *int `mandatory:"false" json:"paginationOffset"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	PaginationLimit *int `mandatory:"false" json:"paginationLimit"`
}

func (m ListIpInventoryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListIpInventoryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.ResourceTypeList {
		if _, ok := GetMappingListIpInventoryDetailsResourceTypeListEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceTypeList: %s. Supported values are: %s.", val, strings.Join(GetListIpInventoryDetailsResourceTypeListEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListIpInventoryDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetListIpInventoryDetailsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIpInventoryDetailsSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetListIpInventoryDetailsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIpInventoryDetailsResourceTypeListEnum Enum with underlying type: string
type ListIpInventoryDetailsResourceTypeListEnum string

// Set of constants representing the allowable values for ListIpInventoryDetailsResourceTypeListEnum
const (
	ListIpInventoryDetailsResourceTypeListResource ListIpInventoryDetailsResourceTypeListEnum = "Resource"
)

var mappingListIpInventoryDetailsResourceTypeListEnum = map[string]ListIpInventoryDetailsResourceTypeListEnum{
	"Resource": ListIpInventoryDetailsResourceTypeListResource,
}

var mappingListIpInventoryDetailsResourceTypeListEnumLowerCase = map[string]ListIpInventoryDetailsResourceTypeListEnum{
	"resource": ListIpInventoryDetailsResourceTypeListResource,
}

// GetListIpInventoryDetailsResourceTypeListEnumValues Enumerates the set of values for ListIpInventoryDetailsResourceTypeListEnum
func GetListIpInventoryDetailsResourceTypeListEnumValues() []ListIpInventoryDetailsResourceTypeListEnum {
	values := make([]ListIpInventoryDetailsResourceTypeListEnum, 0)
	for _, v := range mappingListIpInventoryDetailsResourceTypeListEnum {
		values = append(values, v)
	}
	return values
}

// GetListIpInventoryDetailsResourceTypeListEnumStringValues Enumerates the set of values in String for ListIpInventoryDetailsResourceTypeListEnum
func GetListIpInventoryDetailsResourceTypeListEnumStringValues() []string {
	return []string{
		"Resource",
	}
}

// GetMappingListIpInventoryDetailsResourceTypeListEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIpInventoryDetailsResourceTypeListEnum(val string) (ListIpInventoryDetailsResourceTypeListEnum, bool) {
	enum, ok := mappingListIpInventoryDetailsResourceTypeListEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIpInventoryDetailsSortByEnum Enum with underlying type: string
type ListIpInventoryDetailsSortByEnum string

// Set of constants representing the allowable values for ListIpInventoryDetailsSortByEnum
const (
	ListIpInventoryDetailsSortByDisplayname ListIpInventoryDetailsSortByEnum = "DISPLAYNAME"
	ListIpInventoryDetailsSortByUtilization ListIpInventoryDetailsSortByEnum = "UTILIZATION"
	ListIpInventoryDetailsSortByDnsHostname ListIpInventoryDetailsSortByEnum = "DNS_HOSTNAME"
	ListIpInventoryDetailsSortByRegion      ListIpInventoryDetailsSortByEnum = "REGION"
)

var mappingListIpInventoryDetailsSortByEnum = map[string]ListIpInventoryDetailsSortByEnum{
	"DISPLAYNAME":  ListIpInventoryDetailsSortByDisplayname,
	"UTILIZATION":  ListIpInventoryDetailsSortByUtilization,
	"DNS_HOSTNAME": ListIpInventoryDetailsSortByDnsHostname,
	"REGION":       ListIpInventoryDetailsSortByRegion,
}

var mappingListIpInventoryDetailsSortByEnumLowerCase = map[string]ListIpInventoryDetailsSortByEnum{
	"displayname":  ListIpInventoryDetailsSortByDisplayname,
	"utilization":  ListIpInventoryDetailsSortByUtilization,
	"dns_hostname": ListIpInventoryDetailsSortByDnsHostname,
	"region":       ListIpInventoryDetailsSortByRegion,
}

// GetListIpInventoryDetailsSortByEnumValues Enumerates the set of values for ListIpInventoryDetailsSortByEnum
func GetListIpInventoryDetailsSortByEnumValues() []ListIpInventoryDetailsSortByEnum {
	values := make([]ListIpInventoryDetailsSortByEnum, 0)
	for _, v := range mappingListIpInventoryDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIpInventoryDetailsSortByEnumStringValues Enumerates the set of values in String for ListIpInventoryDetailsSortByEnum
func GetListIpInventoryDetailsSortByEnumStringValues() []string {
	return []string{
		"DISPLAYNAME",
		"UTILIZATION",
		"DNS_HOSTNAME",
		"REGION",
	}
}

// GetMappingListIpInventoryDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIpInventoryDetailsSortByEnum(val string) (ListIpInventoryDetailsSortByEnum, bool) {
	enum, ok := mappingListIpInventoryDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIpInventoryDetailsSortOrderEnum Enum with underlying type: string
type ListIpInventoryDetailsSortOrderEnum string

// Set of constants representing the allowable values for ListIpInventoryDetailsSortOrderEnum
const (
	ListIpInventoryDetailsSortOrderAsc  ListIpInventoryDetailsSortOrderEnum = "ASC"
	ListIpInventoryDetailsSortOrderDesc ListIpInventoryDetailsSortOrderEnum = "DESC"
)

var mappingListIpInventoryDetailsSortOrderEnum = map[string]ListIpInventoryDetailsSortOrderEnum{
	"ASC":  ListIpInventoryDetailsSortOrderAsc,
	"DESC": ListIpInventoryDetailsSortOrderDesc,
}

var mappingListIpInventoryDetailsSortOrderEnumLowerCase = map[string]ListIpInventoryDetailsSortOrderEnum{
	"asc":  ListIpInventoryDetailsSortOrderAsc,
	"desc": ListIpInventoryDetailsSortOrderDesc,
}

// GetListIpInventoryDetailsSortOrderEnumValues Enumerates the set of values for ListIpInventoryDetailsSortOrderEnum
func GetListIpInventoryDetailsSortOrderEnumValues() []ListIpInventoryDetailsSortOrderEnum {
	values := make([]ListIpInventoryDetailsSortOrderEnum, 0)
	for _, v := range mappingListIpInventoryDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIpInventoryDetailsSortOrderEnumStringValues Enumerates the set of values in String for ListIpInventoryDetailsSortOrderEnum
func GetListIpInventoryDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIpInventoryDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIpInventoryDetailsSortOrderEnum(val string) (ListIpInventoryDetailsSortOrderEnum, bool) {
	enum, ok := mappingListIpInventoryDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
