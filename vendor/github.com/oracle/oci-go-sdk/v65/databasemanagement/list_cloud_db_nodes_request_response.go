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

// ListCloudDbNodesRequest wrapper for the ListCloudDbNodes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudDbNodes.go.html to see an example of how to use ListCloudDbNodesRequest.
type ListCloudDbNodesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	CloudDbSystemId *string `mandatory:"false" contributesTo:"query" name:"cloudDbSystemId"`

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
	SortBy ListCloudDbNodesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudDbNodesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudDbNodesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudDbNodesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudDbNodesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudDbNodesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudDbNodesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudDbNodesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudDbNodesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudDbNodesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudDbNodesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudDbNodesResponse wrapper for the ListCloudDbNodes operation
type ListCloudDbNodesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudDbNodeCollection instances
	CloudDbNodeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudDbNodesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudDbNodesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudDbNodesSortByEnum Enum with underlying type: string
type ListCloudDbNodesSortByEnum string

// Set of constants representing the allowable values for ListCloudDbNodesSortByEnum
const (
	ListCloudDbNodesSortByTimecreated ListCloudDbNodesSortByEnum = "TIMECREATED"
	ListCloudDbNodesSortByDisplayname ListCloudDbNodesSortByEnum = "DISPLAYNAME"
)

var mappingListCloudDbNodesSortByEnum = map[string]ListCloudDbNodesSortByEnum{
	"TIMECREATED": ListCloudDbNodesSortByTimecreated,
	"DISPLAYNAME": ListCloudDbNodesSortByDisplayname,
}

var mappingListCloudDbNodesSortByEnumLowerCase = map[string]ListCloudDbNodesSortByEnum{
	"timecreated": ListCloudDbNodesSortByTimecreated,
	"displayname": ListCloudDbNodesSortByDisplayname,
}

// GetListCloudDbNodesSortByEnumValues Enumerates the set of values for ListCloudDbNodesSortByEnum
func GetListCloudDbNodesSortByEnumValues() []ListCloudDbNodesSortByEnum {
	values := make([]ListCloudDbNodesSortByEnum, 0)
	for _, v := range mappingListCloudDbNodesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbNodesSortByEnumStringValues Enumerates the set of values in String for ListCloudDbNodesSortByEnum
func GetListCloudDbNodesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudDbNodesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbNodesSortByEnum(val string) (ListCloudDbNodesSortByEnum, bool) {
	enum, ok := mappingListCloudDbNodesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudDbNodesSortOrderEnum Enum with underlying type: string
type ListCloudDbNodesSortOrderEnum string

// Set of constants representing the allowable values for ListCloudDbNodesSortOrderEnum
const (
	ListCloudDbNodesSortOrderAsc  ListCloudDbNodesSortOrderEnum = "ASC"
	ListCloudDbNodesSortOrderDesc ListCloudDbNodesSortOrderEnum = "DESC"
)

var mappingListCloudDbNodesSortOrderEnum = map[string]ListCloudDbNodesSortOrderEnum{
	"ASC":  ListCloudDbNodesSortOrderAsc,
	"DESC": ListCloudDbNodesSortOrderDesc,
}

var mappingListCloudDbNodesSortOrderEnumLowerCase = map[string]ListCloudDbNodesSortOrderEnum{
	"asc":  ListCloudDbNodesSortOrderAsc,
	"desc": ListCloudDbNodesSortOrderDesc,
}

// GetListCloudDbNodesSortOrderEnumValues Enumerates the set of values for ListCloudDbNodesSortOrderEnum
func GetListCloudDbNodesSortOrderEnumValues() []ListCloudDbNodesSortOrderEnum {
	values := make([]ListCloudDbNodesSortOrderEnum, 0)
	for _, v := range mappingListCloudDbNodesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbNodesSortOrderEnumStringValues Enumerates the set of values in String for ListCloudDbNodesSortOrderEnum
func GetListCloudDbNodesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudDbNodesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbNodesSortOrderEnum(val string) (ListCloudDbNodesSortOrderEnum, bool) {
	enum, ok := mappingListCloudDbNodesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
