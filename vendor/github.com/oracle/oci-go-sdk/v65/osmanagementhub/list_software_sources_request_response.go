// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSoftwareSourcesRequest wrapper for the ListSoftwareSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListSoftwareSources.go.html to see an example of how to use ListSoftwareSourcesRequest.
type ListSoftwareSourcesRequest struct {

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID for the software source.
	SoftwareSourceId *string `mandatory:"false" contributesTo:"query" name:"softwareSourceId"`

	// The type of the software source.
	SoftwareSourceType []SoftwareSourceTypeEnum `contributesTo:"query" name:"softwareSourceType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only profiles that match the given vendorName.
	VendorName ListSoftwareSourcesVendorNameEnum `mandatory:"false" contributesTo:"query" name:"vendorName" omitEmpty:"true"`

	// A filter to return only instances whose OS family type matches the given OS family.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only instances whose architecture type matches the given architecture.
	ArchType []ArchTypeEnum `contributesTo:"query" name:"archType" omitEmpty:"true" collectionFormat:"multi"`

	// The availabilities of the software source for a tenant.
	Availability []AvailabilityEnum `contributesTo:"query" name:"availability" omitEmpty:"true" collectionFormat:"multi"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A multi filter to return resources that do not contains the given display names.
	DisplayNameNotEqualTo []string `contributesTo:"query" name:"displayNameNotEqualTo" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSoftwareSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListSoftwareSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleStates.
	LifecycleState []SoftwareSourceLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwareSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwareSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.SoftwareSourceType {
		if _, ok := GetMappingSoftwareSourceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareSourceType: %s. Supported values are: %s.", val, strings.Join(GetSoftwareSourceTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListSoftwareSourcesVendorNameEnum(string(request.VendorName)); !ok && request.VendorName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", request.VendorName, strings.Join(GetListSoftwareSourcesVendorNameEnumStringValues(), ",")))
	}
	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ArchType {
		if _, ok := GetMappingArchTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", val, strings.Join(GetArchTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.Availability {
		if _, ok := GetMappingAvailabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Availability: %s. Supported values are: %s.", val, strings.Join(GetAvailabilityEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListSoftwareSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwareSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwareSourcesSortByEnumStringValues(), ",")))
	}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingSoftwareSourceLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetSoftwareSourceLifecycleStateEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwareSourcesResponse wrapper for the ListSoftwareSources operation
type ListSoftwareSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwareSourceCollection instances
	SoftwareSourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwareSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareSourcesVendorNameEnum Enum with underlying type: string
type ListSoftwareSourcesVendorNameEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesVendorNameEnum
const (
	ListSoftwareSourcesVendorNameOracle ListSoftwareSourcesVendorNameEnum = "ORACLE"
)

var mappingListSoftwareSourcesVendorNameEnum = map[string]ListSoftwareSourcesVendorNameEnum{
	"ORACLE": ListSoftwareSourcesVendorNameOracle,
}

var mappingListSoftwareSourcesVendorNameEnumLowerCase = map[string]ListSoftwareSourcesVendorNameEnum{
	"oracle": ListSoftwareSourcesVendorNameOracle,
}

// GetListSoftwareSourcesVendorNameEnumValues Enumerates the set of values for ListSoftwareSourcesVendorNameEnum
func GetListSoftwareSourcesVendorNameEnumValues() []ListSoftwareSourcesVendorNameEnum {
	values := make([]ListSoftwareSourcesVendorNameEnum, 0)
	for _, v := range mappingListSoftwareSourcesVendorNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourcesVendorNameEnumStringValues Enumerates the set of values in String for ListSoftwareSourcesVendorNameEnum
func GetListSoftwareSourcesVendorNameEnumStringValues() []string {
	return []string{
		"ORACLE",
	}
}

// GetMappingListSoftwareSourcesVendorNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourcesVendorNameEnum(val string) (ListSoftwareSourcesVendorNameEnum, bool) {
	enum, ok := mappingListSoftwareSourcesVendorNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareSourcesSortOrderEnum Enum with underlying type: string
type ListSoftwareSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesSortOrderEnum
const (
	ListSoftwareSourcesSortOrderAsc  ListSoftwareSourcesSortOrderEnum = "ASC"
	ListSoftwareSourcesSortOrderDesc ListSoftwareSourcesSortOrderEnum = "DESC"
)

var mappingListSoftwareSourcesSortOrderEnum = map[string]ListSoftwareSourcesSortOrderEnum{
	"ASC":  ListSoftwareSourcesSortOrderAsc,
	"DESC": ListSoftwareSourcesSortOrderDesc,
}

var mappingListSoftwareSourcesSortOrderEnumLowerCase = map[string]ListSoftwareSourcesSortOrderEnum{
	"asc":  ListSoftwareSourcesSortOrderAsc,
	"desc": ListSoftwareSourcesSortOrderDesc,
}

// GetListSoftwareSourcesSortOrderEnumValues Enumerates the set of values for ListSoftwareSourcesSortOrderEnum
func GetListSoftwareSourcesSortOrderEnumValues() []ListSoftwareSourcesSortOrderEnum {
	values := make([]ListSoftwareSourcesSortOrderEnum, 0)
	for _, v := range mappingListSoftwareSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwareSourcesSortOrderEnum
func GetListSoftwareSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwareSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourcesSortOrderEnum(val string) (ListSoftwareSourcesSortOrderEnum, bool) {
	enum, ok := mappingListSoftwareSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareSourcesSortByEnum Enum with underlying type: string
type ListSoftwareSourcesSortByEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesSortByEnum
const (
	ListSoftwareSourcesSortByTimecreated ListSoftwareSourcesSortByEnum = "timeCreated"
	ListSoftwareSourcesSortByDisplayname ListSoftwareSourcesSortByEnum = "displayName"
)

var mappingListSoftwareSourcesSortByEnum = map[string]ListSoftwareSourcesSortByEnum{
	"timeCreated": ListSoftwareSourcesSortByTimecreated,
	"displayName": ListSoftwareSourcesSortByDisplayname,
}

var mappingListSoftwareSourcesSortByEnumLowerCase = map[string]ListSoftwareSourcesSortByEnum{
	"timecreated": ListSoftwareSourcesSortByTimecreated,
	"displayname": ListSoftwareSourcesSortByDisplayname,
}

// GetListSoftwareSourcesSortByEnumValues Enumerates the set of values for ListSoftwareSourcesSortByEnum
func GetListSoftwareSourcesSortByEnumValues() []ListSoftwareSourcesSortByEnum {
	values := make([]ListSoftwareSourcesSortByEnum, 0)
	for _, v := range mappingListSoftwareSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourcesSortByEnumStringValues Enumerates the set of values in String for ListSoftwareSourcesSortByEnum
func GetListSoftwareSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSoftwareSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourcesSortByEnum(val string) (ListSoftwareSourcesSortByEnum, bool) {
	enum, ok := mappingListSoftwareSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
