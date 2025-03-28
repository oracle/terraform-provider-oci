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

// ListExternalListenersRequest wrapper for the ListExternalListeners operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalListeners.go.html to see an example of how to use ListExternalListenersRequest.
type ListExternalListenersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
	ExternalDbSystemId *string `mandatory:"false" contributesTo:"query" name:"externalDbSystemId"`

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
	SortBy ListExternalListenersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalListenersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalListenersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalListenersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalListenersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalListenersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalListenersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalListenersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalListenersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalListenersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalListenersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalListenersResponse wrapper for the ListExternalListeners operation
type ListExternalListenersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalListenerCollection instances
	ExternalListenerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalListenersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalListenersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalListenersSortByEnum Enum with underlying type: string
type ListExternalListenersSortByEnum string

// Set of constants representing the allowable values for ListExternalListenersSortByEnum
const (
	ListExternalListenersSortByTimecreated ListExternalListenersSortByEnum = "TIMECREATED"
	ListExternalListenersSortByDisplayname ListExternalListenersSortByEnum = "DISPLAYNAME"
)

var mappingListExternalListenersSortByEnum = map[string]ListExternalListenersSortByEnum{
	"TIMECREATED": ListExternalListenersSortByTimecreated,
	"DISPLAYNAME": ListExternalListenersSortByDisplayname,
}

var mappingListExternalListenersSortByEnumLowerCase = map[string]ListExternalListenersSortByEnum{
	"timecreated": ListExternalListenersSortByTimecreated,
	"displayname": ListExternalListenersSortByDisplayname,
}

// GetListExternalListenersSortByEnumValues Enumerates the set of values for ListExternalListenersSortByEnum
func GetListExternalListenersSortByEnumValues() []ListExternalListenersSortByEnum {
	values := make([]ListExternalListenersSortByEnum, 0)
	for _, v := range mappingListExternalListenersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalListenersSortByEnumStringValues Enumerates the set of values in String for ListExternalListenersSortByEnum
func GetListExternalListenersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExternalListenersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalListenersSortByEnum(val string) (ListExternalListenersSortByEnum, bool) {
	enum, ok := mappingListExternalListenersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalListenersSortOrderEnum Enum with underlying type: string
type ListExternalListenersSortOrderEnum string

// Set of constants representing the allowable values for ListExternalListenersSortOrderEnum
const (
	ListExternalListenersSortOrderAsc  ListExternalListenersSortOrderEnum = "ASC"
	ListExternalListenersSortOrderDesc ListExternalListenersSortOrderEnum = "DESC"
)

var mappingListExternalListenersSortOrderEnum = map[string]ListExternalListenersSortOrderEnum{
	"ASC":  ListExternalListenersSortOrderAsc,
	"DESC": ListExternalListenersSortOrderDesc,
}

var mappingListExternalListenersSortOrderEnumLowerCase = map[string]ListExternalListenersSortOrderEnum{
	"asc":  ListExternalListenersSortOrderAsc,
	"desc": ListExternalListenersSortOrderDesc,
}

// GetListExternalListenersSortOrderEnumValues Enumerates the set of values for ListExternalListenersSortOrderEnum
func GetListExternalListenersSortOrderEnumValues() []ListExternalListenersSortOrderEnum {
	values := make([]ListExternalListenersSortOrderEnum, 0)
	for _, v := range mappingListExternalListenersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalListenersSortOrderEnumStringValues Enumerates the set of values in String for ListExternalListenersSortOrderEnum
func GetListExternalListenersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalListenersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalListenersSortOrderEnum(val string) (ListExternalListenersSortOrderEnum, bool) {
	enum, ok := mappingListExternalListenersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
