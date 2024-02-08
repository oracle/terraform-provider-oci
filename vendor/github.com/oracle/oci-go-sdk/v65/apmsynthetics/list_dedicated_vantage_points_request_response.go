// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDedicatedVantagePointsRequest wrapper for the ListDedicatedVantagePoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmsynthetics/ListDedicatedVantagePoints.go.html to see an example of how to use ListDedicatedVantagePointsRequest.
type ListDedicatedVantagePointsRequest struct {

	// The APM domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The maximum number of results per page, or items to return in a paginated
	// "List" call. For information on how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). Default sort order is ascending.
	SortOrder ListDedicatedVantagePointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order of displayName is ascending.
	// Default order of timeCreated and timeUpdated is descending.
	// The displayName sort by is case-sensitive.
	SortBy ListDedicatedVantagePointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only the dedicated vantage points that match a given status.
	Status ListDedicatedVantagePointsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDedicatedVantagePointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDedicatedVantagePointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDedicatedVantagePointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDedicatedVantagePointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDedicatedVantagePointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDedicatedVantagePointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDedicatedVantagePointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDedicatedVantagePointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDedicatedVantagePointsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDedicatedVantagePointsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListDedicatedVantagePointsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDedicatedVantagePointsResponse wrapper for the ListDedicatedVantagePoints operation
type ListDedicatedVantagePointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DedicatedVantagePointCollection instances
	DedicatedVantagePointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDedicatedVantagePointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDedicatedVantagePointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDedicatedVantagePointsSortOrderEnum Enum with underlying type: string
type ListDedicatedVantagePointsSortOrderEnum string

// Set of constants representing the allowable values for ListDedicatedVantagePointsSortOrderEnum
const (
	ListDedicatedVantagePointsSortOrderAsc  ListDedicatedVantagePointsSortOrderEnum = "ASC"
	ListDedicatedVantagePointsSortOrderDesc ListDedicatedVantagePointsSortOrderEnum = "DESC"
)

var mappingListDedicatedVantagePointsSortOrderEnum = map[string]ListDedicatedVantagePointsSortOrderEnum{
	"ASC":  ListDedicatedVantagePointsSortOrderAsc,
	"DESC": ListDedicatedVantagePointsSortOrderDesc,
}

var mappingListDedicatedVantagePointsSortOrderEnumLowerCase = map[string]ListDedicatedVantagePointsSortOrderEnum{
	"asc":  ListDedicatedVantagePointsSortOrderAsc,
	"desc": ListDedicatedVantagePointsSortOrderDesc,
}

// GetListDedicatedVantagePointsSortOrderEnumValues Enumerates the set of values for ListDedicatedVantagePointsSortOrderEnum
func GetListDedicatedVantagePointsSortOrderEnumValues() []ListDedicatedVantagePointsSortOrderEnum {
	values := make([]ListDedicatedVantagePointsSortOrderEnum, 0)
	for _, v := range mappingListDedicatedVantagePointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDedicatedVantagePointsSortOrderEnumStringValues Enumerates the set of values in String for ListDedicatedVantagePointsSortOrderEnum
func GetListDedicatedVantagePointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDedicatedVantagePointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDedicatedVantagePointsSortOrderEnum(val string) (ListDedicatedVantagePointsSortOrderEnum, bool) {
	enum, ok := mappingListDedicatedVantagePointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDedicatedVantagePointsSortByEnum Enum with underlying type: string
type ListDedicatedVantagePointsSortByEnum string

// Set of constants representing the allowable values for ListDedicatedVantagePointsSortByEnum
const (
	ListDedicatedVantagePointsSortByDisplayname ListDedicatedVantagePointsSortByEnum = "displayName"
	ListDedicatedVantagePointsSortByName        ListDedicatedVantagePointsSortByEnum = "name"
	ListDedicatedVantagePointsSortByTimecreated ListDedicatedVantagePointsSortByEnum = "timeCreated"
	ListDedicatedVantagePointsSortByTimeupdated ListDedicatedVantagePointsSortByEnum = "timeUpdated"
	ListDedicatedVantagePointsSortByStatus      ListDedicatedVantagePointsSortByEnum = "status"
)

var mappingListDedicatedVantagePointsSortByEnum = map[string]ListDedicatedVantagePointsSortByEnum{
	"displayName": ListDedicatedVantagePointsSortByDisplayname,
	"name":        ListDedicatedVantagePointsSortByName,
	"timeCreated": ListDedicatedVantagePointsSortByTimecreated,
	"timeUpdated": ListDedicatedVantagePointsSortByTimeupdated,
	"status":      ListDedicatedVantagePointsSortByStatus,
}

var mappingListDedicatedVantagePointsSortByEnumLowerCase = map[string]ListDedicatedVantagePointsSortByEnum{
	"displayname": ListDedicatedVantagePointsSortByDisplayname,
	"name":        ListDedicatedVantagePointsSortByName,
	"timecreated": ListDedicatedVantagePointsSortByTimecreated,
	"timeupdated": ListDedicatedVantagePointsSortByTimeupdated,
	"status":      ListDedicatedVantagePointsSortByStatus,
}

// GetListDedicatedVantagePointsSortByEnumValues Enumerates the set of values for ListDedicatedVantagePointsSortByEnum
func GetListDedicatedVantagePointsSortByEnumValues() []ListDedicatedVantagePointsSortByEnum {
	values := make([]ListDedicatedVantagePointsSortByEnum, 0)
	for _, v := range mappingListDedicatedVantagePointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDedicatedVantagePointsSortByEnumStringValues Enumerates the set of values in String for ListDedicatedVantagePointsSortByEnum
func GetListDedicatedVantagePointsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"name",
		"timeCreated",
		"timeUpdated",
		"status",
	}
}

// GetMappingListDedicatedVantagePointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDedicatedVantagePointsSortByEnum(val string) (ListDedicatedVantagePointsSortByEnum, bool) {
	enum, ok := mappingListDedicatedVantagePointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDedicatedVantagePointsStatusEnum Enum with underlying type: string
type ListDedicatedVantagePointsStatusEnum string

// Set of constants representing the allowable values for ListDedicatedVantagePointsStatusEnum
const (
	ListDedicatedVantagePointsStatusEnabled  ListDedicatedVantagePointsStatusEnum = "ENABLED"
	ListDedicatedVantagePointsStatusDisabled ListDedicatedVantagePointsStatusEnum = "DISABLED"
)

var mappingListDedicatedVantagePointsStatusEnum = map[string]ListDedicatedVantagePointsStatusEnum{
	"ENABLED":  ListDedicatedVantagePointsStatusEnabled,
	"DISABLED": ListDedicatedVantagePointsStatusDisabled,
}

var mappingListDedicatedVantagePointsStatusEnumLowerCase = map[string]ListDedicatedVantagePointsStatusEnum{
	"enabled":  ListDedicatedVantagePointsStatusEnabled,
	"disabled": ListDedicatedVantagePointsStatusDisabled,
}

// GetListDedicatedVantagePointsStatusEnumValues Enumerates the set of values for ListDedicatedVantagePointsStatusEnum
func GetListDedicatedVantagePointsStatusEnumValues() []ListDedicatedVantagePointsStatusEnum {
	values := make([]ListDedicatedVantagePointsStatusEnum, 0)
	for _, v := range mappingListDedicatedVantagePointsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListDedicatedVantagePointsStatusEnumStringValues Enumerates the set of values in String for ListDedicatedVantagePointsStatusEnum
func GetListDedicatedVantagePointsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingListDedicatedVantagePointsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDedicatedVantagePointsStatusEnum(val string) (ListDedicatedVantagePointsStatusEnum, bool) {
	enum, ok := mappingListDedicatedVantagePointsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
