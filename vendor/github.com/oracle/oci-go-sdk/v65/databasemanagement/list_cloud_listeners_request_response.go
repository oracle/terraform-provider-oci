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

// ListCloudListenersRequest wrapper for the ListCloudListeners operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudListeners.go.html to see an example of how to use ListCloudListenersRequest.
type ListCloudListenersRequest struct {

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
	SortBy ListCloudListenersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudListenersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudListenersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudListenersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudListenersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudListenersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudListenersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudListenersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudListenersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudListenersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudListenersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudListenersResponse wrapper for the ListCloudListeners operation
type ListCloudListenersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudListenerCollection instances
	CloudListenerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudListenersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudListenersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudListenersSortByEnum Enum with underlying type: string
type ListCloudListenersSortByEnum string

// Set of constants representing the allowable values for ListCloudListenersSortByEnum
const (
	ListCloudListenersSortByTimecreated ListCloudListenersSortByEnum = "TIMECREATED"
	ListCloudListenersSortByDisplayname ListCloudListenersSortByEnum = "DISPLAYNAME"
)

var mappingListCloudListenersSortByEnum = map[string]ListCloudListenersSortByEnum{
	"TIMECREATED": ListCloudListenersSortByTimecreated,
	"DISPLAYNAME": ListCloudListenersSortByDisplayname,
}

var mappingListCloudListenersSortByEnumLowerCase = map[string]ListCloudListenersSortByEnum{
	"timecreated": ListCloudListenersSortByTimecreated,
	"displayname": ListCloudListenersSortByDisplayname,
}

// GetListCloudListenersSortByEnumValues Enumerates the set of values for ListCloudListenersSortByEnum
func GetListCloudListenersSortByEnumValues() []ListCloudListenersSortByEnum {
	values := make([]ListCloudListenersSortByEnum, 0)
	for _, v := range mappingListCloudListenersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudListenersSortByEnumStringValues Enumerates the set of values in String for ListCloudListenersSortByEnum
func GetListCloudListenersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudListenersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudListenersSortByEnum(val string) (ListCloudListenersSortByEnum, bool) {
	enum, ok := mappingListCloudListenersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudListenersSortOrderEnum Enum with underlying type: string
type ListCloudListenersSortOrderEnum string

// Set of constants representing the allowable values for ListCloudListenersSortOrderEnum
const (
	ListCloudListenersSortOrderAsc  ListCloudListenersSortOrderEnum = "ASC"
	ListCloudListenersSortOrderDesc ListCloudListenersSortOrderEnum = "DESC"
)

var mappingListCloudListenersSortOrderEnum = map[string]ListCloudListenersSortOrderEnum{
	"ASC":  ListCloudListenersSortOrderAsc,
	"DESC": ListCloudListenersSortOrderDesc,
}

var mappingListCloudListenersSortOrderEnumLowerCase = map[string]ListCloudListenersSortOrderEnum{
	"asc":  ListCloudListenersSortOrderAsc,
	"desc": ListCloudListenersSortOrderDesc,
}

// GetListCloudListenersSortOrderEnumValues Enumerates the set of values for ListCloudListenersSortOrderEnum
func GetListCloudListenersSortOrderEnumValues() []ListCloudListenersSortOrderEnum {
	values := make([]ListCloudListenersSortOrderEnum, 0)
	for _, v := range mappingListCloudListenersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudListenersSortOrderEnumStringValues Enumerates the set of values in String for ListCloudListenersSortOrderEnum
func GetListCloudListenersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudListenersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudListenersSortOrderEnum(val string) (ListCloudListenersSortOrderEnum, bool) {
	enum, ok := mappingListCloudListenersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
