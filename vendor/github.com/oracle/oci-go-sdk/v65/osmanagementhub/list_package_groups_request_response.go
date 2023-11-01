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

// ListPackageGroupsRequest wrapper for the ListPackageGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListPackageGroups.go.html to see an example of how to use ListPackageGroupsRequest.
type ListPackageGroupsRequest struct {

	// The software source OCID.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// The OCID of the compartment that contains the resources to list.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The name of the entity to be queried.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return resources that may partially match the name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// A filter to return only package groups of the specified type.
	GroupType []PackageGroupGroupTypeEnum `contributesTo:"query" name:"groupType" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPackageGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListPackageGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPackageGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPackageGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPackageGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPackageGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPackageGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.GroupType {
		if _, ok := GetMappingPackageGroupGroupTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", val, strings.Join(GetPackageGroupGroupTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListPackageGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPackageGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPackageGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPackageGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPackageGroupsResponse wrapper for the ListPackageGroups operation
type ListPackageGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PackageGroupCollection instances
	PackageGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPackageGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPackageGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPackageGroupsSortOrderEnum Enum with underlying type: string
type ListPackageGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListPackageGroupsSortOrderEnum
const (
	ListPackageGroupsSortOrderAsc  ListPackageGroupsSortOrderEnum = "ASC"
	ListPackageGroupsSortOrderDesc ListPackageGroupsSortOrderEnum = "DESC"
)

var mappingListPackageGroupsSortOrderEnum = map[string]ListPackageGroupsSortOrderEnum{
	"ASC":  ListPackageGroupsSortOrderAsc,
	"DESC": ListPackageGroupsSortOrderDesc,
}

var mappingListPackageGroupsSortOrderEnumLowerCase = map[string]ListPackageGroupsSortOrderEnum{
	"asc":  ListPackageGroupsSortOrderAsc,
	"desc": ListPackageGroupsSortOrderDesc,
}

// GetListPackageGroupsSortOrderEnumValues Enumerates the set of values for ListPackageGroupsSortOrderEnum
func GetListPackageGroupsSortOrderEnumValues() []ListPackageGroupsSortOrderEnum {
	values := make([]ListPackageGroupsSortOrderEnum, 0)
	for _, v := range mappingListPackageGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPackageGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListPackageGroupsSortOrderEnum
func GetListPackageGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPackageGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPackageGroupsSortOrderEnum(val string) (ListPackageGroupsSortOrderEnum, bool) {
	enum, ok := mappingListPackageGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPackageGroupsSortByEnum Enum with underlying type: string
type ListPackageGroupsSortByEnum string

// Set of constants representing the allowable values for ListPackageGroupsSortByEnum
const (
	ListPackageGroupsSortByTimecreated ListPackageGroupsSortByEnum = "timeCreated"
	ListPackageGroupsSortByDisplayname ListPackageGroupsSortByEnum = "displayName"
)

var mappingListPackageGroupsSortByEnum = map[string]ListPackageGroupsSortByEnum{
	"timeCreated": ListPackageGroupsSortByTimecreated,
	"displayName": ListPackageGroupsSortByDisplayname,
}

var mappingListPackageGroupsSortByEnumLowerCase = map[string]ListPackageGroupsSortByEnum{
	"timecreated": ListPackageGroupsSortByTimecreated,
	"displayname": ListPackageGroupsSortByDisplayname,
}

// GetListPackageGroupsSortByEnumValues Enumerates the set of values for ListPackageGroupsSortByEnum
func GetListPackageGroupsSortByEnumValues() []ListPackageGroupsSortByEnum {
	values := make([]ListPackageGroupsSortByEnum, 0)
	for _, v := range mappingListPackageGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPackageGroupsSortByEnumStringValues Enumerates the set of values in String for ListPackageGroupsSortByEnum
func GetListPackageGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPackageGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPackageGroupsSortByEnum(val string) (ListPackageGroupsSortByEnum, bool) {
	enum, ok := mappingListPackageGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
