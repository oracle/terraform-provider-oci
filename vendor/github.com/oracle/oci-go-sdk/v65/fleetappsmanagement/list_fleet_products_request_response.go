// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFleetProductsRequest wrapper for the ListFleetProducts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListFleetProducts.go.html to see an example of how to use ListFleetProductsRequest.
type ListFleetProductsRequest struct {

	// Unique Fleet identifier.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Resource Identifier
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Resource Display Name.
	ResourceDisplayName *string `mandatory:"false" contributesTo:"query" name:"resourceDisplayName"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFleetProductsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName and resourceDisplayName is ascending.
	SortBy ListFleetProductsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetProductsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetProductsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetProductsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetProductsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFleetProductsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFleetProductsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFleetProductsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetProductsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFleetProductsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFleetProductsResponse wrapper for the ListFleetProducts operation
type ListFleetProductsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetProductCollection instances
	FleetProductCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetProductsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetProductsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetProductsSortOrderEnum Enum with underlying type: string
type ListFleetProductsSortOrderEnum string

// Set of constants representing the allowable values for ListFleetProductsSortOrderEnum
const (
	ListFleetProductsSortOrderAsc  ListFleetProductsSortOrderEnum = "ASC"
	ListFleetProductsSortOrderDesc ListFleetProductsSortOrderEnum = "DESC"
)

var mappingListFleetProductsSortOrderEnum = map[string]ListFleetProductsSortOrderEnum{
	"ASC":  ListFleetProductsSortOrderAsc,
	"DESC": ListFleetProductsSortOrderDesc,
}

var mappingListFleetProductsSortOrderEnumLowerCase = map[string]ListFleetProductsSortOrderEnum{
	"asc":  ListFleetProductsSortOrderAsc,
	"desc": ListFleetProductsSortOrderDesc,
}

// GetListFleetProductsSortOrderEnumValues Enumerates the set of values for ListFleetProductsSortOrderEnum
func GetListFleetProductsSortOrderEnumValues() []ListFleetProductsSortOrderEnum {
	values := make([]ListFleetProductsSortOrderEnum, 0)
	for _, v := range mappingListFleetProductsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetProductsSortOrderEnumStringValues Enumerates the set of values in String for ListFleetProductsSortOrderEnum
func GetListFleetProductsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFleetProductsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetProductsSortOrderEnum(val string) (ListFleetProductsSortOrderEnum, bool) {
	enum, ok := mappingListFleetProductsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetProductsSortByEnum Enum with underlying type: string
type ListFleetProductsSortByEnum string

// Set of constants representing the allowable values for ListFleetProductsSortByEnum
const (
	ListFleetProductsSortByDisplayname         ListFleetProductsSortByEnum = "displayName"
	ListFleetProductsSortByResourcedisplayname ListFleetProductsSortByEnum = "resourceDisplayName"
)

var mappingListFleetProductsSortByEnum = map[string]ListFleetProductsSortByEnum{
	"displayName":         ListFleetProductsSortByDisplayname,
	"resourceDisplayName": ListFleetProductsSortByResourcedisplayname,
}

var mappingListFleetProductsSortByEnumLowerCase = map[string]ListFleetProductsSortByEnum{
	"displayname":         ListFleetProductsSortByDisplayname,
	"resourcedisplayname": ListFleetProductsSortByResourcedisplayname,
}

// GetListFleetProductsSortByEnumValues Enumerates the set of values for ListFleetProductsSortByEnum
func GetListFleetProductsSortByEnumValues() []ListFleetProductsSortByEnum {
	values := make([]ListFleetProductsSortByEnum, 0)
	for _, v := range mappingListFleetProductsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetProductsSortByEnumStringValues Enumerates the set of values in String for ListFleetProductsSortByEnum
func GetListFleetProductsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"resourceDisplayName",
	}
}

// GetMappingListFleetProductsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetProductsSortByEnum(val string) (ListFleetProductsSortByEnum, bool) {
	enum, ok := mappingListFleetProductsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
