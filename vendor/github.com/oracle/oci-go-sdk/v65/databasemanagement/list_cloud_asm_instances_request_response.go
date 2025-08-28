// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCloudAsmInstancesRequest wrapper for the ListCloudAsmInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudAsmInstances.go.html to see an example of how to use ListCloudAsmInstancesRequest.
type ListCloudAsmInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM.
	CloudAsmId *string `mandatory:"false" contributesTo:"query" name:"cloudAsmId"`

	// A filter to only return the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for `TIMECREATED` is descending and the default sort order for `DISPLAYNAME` is ascending.
	// The `DISPLAYNAME` sort order is case-sensitive.
	SortBy ListCloudAsmInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudAsmInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudAsmInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudAsmInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudAsmInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudAsmInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudAsmInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudAsmInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudAsmInstancesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudAsmInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudAsmInstancesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudAsmInstancesResponse wrapper for the ListCloudAsmInstances operation
type ListCloudAsmInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudAsmInstanceCollection instances
	CloudAsmInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudAsmInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudAsmInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudAsmInstancesSortByEnum Enum with underlying type: string
type ListCloudAsmInstancesSortByEnum string

// Set of constants representing the allowable values for ListCloudAsmInstancesSortByEnum
const (
	ListCloudAsmInstancesSortByTimecreated ListCloudAsmInstancesSortByEnum = "TIMECREATED"
	ListCloudAsmInstancesSortByDisplayname ListCloudAsmInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListCloudAsmInstancesSortByEnum = map[string]ListCloudAsmInstancesSortByEnum{
	"TIMECREATED": ListCloudAsmInstancesSortByTimecreated,
	"DISPLAYNAME": ListCloudAsmInstancesSortByDisplayname,
}

var mappingListCloudAsmInstancesSortByEnumLowerCase = map[string]ListCloudAsmInstancesSortByEnum{
	"timecreated": ListCloudAsmInstancesSortByTimecreated,
	"displayname": ListCloudAsmInstancesSortByDisplayname,
}

// GetListCloudAsmInstancesSortByEnumValues Enumerates the set of values for ListCloudAsmInstancesSortByEnum
func GetListCloudAsmInstancesSortByEnumValues() []ListCloudAsmInstancesSortByEnum {
	values := make([]ListCloudAsmInstancesSortByEnum, 0)
	for _, v := range mappingListCloudAsmInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAsmInstancesSortByEnumStringValues Enumerates the set of values in String for ListCloudAsmInstancesSortByEnum
func GetListCloudAsmInstancesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudAsmInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAsmInstancesSortByEnum(val string) (ListCloudAsmInstancesSortByEnum, bool) {
	enum, ok := mappingListCloudAsmInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudAsmInstancesSortOrderEnum Enum with underlying type: string
type ListCloudAsmInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListCloudAsmInstancesSortOrderEnum
const (
	ListCloudAsmInstancesSortOrderAsc  ListCloudAsmInstancesSortOrderEnum = "ASC"
	ListCloudAsmInstancesSortOrderDesc ListCloudAsmInstancesSortOrderEnum = "DESC"
)

var mappingListCloudAsmInstancesSortOrderEnum = map[string]ListCloudAsmInstancesSortOrderEnum{
	"ASC":  ListCloudAsmInstancesSortOrderAsc,
	"DESC": ListCloudAsmInstancesSortOrderDesc,
}

var mappingListCloudAsmInstancesSortOrderEnumLowerCase = map[string]ListCloudAsmInstancesSortOrderEnum{
	"asc":  ListCloudAsmInstancesSortOrderAsc,
	"desc": ListCloudAsmInstancesSortOrderDesc,
}

// GetListCloudAsmInstancesSortOrderEnumValues Enumerates the set of values for ListCloudAsmInstancesSortOrderEnum
func GetListCloudAsmInstancesSortOrderEnumValues() []ListCloudAsmInstancesSortOrderEnum {
	values := make([]ListCloudAsmInstancesSortOrderEnum, 0)
	for _, v := range mappingListCloudAsmInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAsmInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListCloudAsmInstancesSortOrderEnum
func GetListCloudAsmInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudAsmInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAsmInstancesSortOrderEnum(val string) (ListCloudAsmInstancesSortOrderEnum, bool) {
	enum, ok := mappingListCloudAsmInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
