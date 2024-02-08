// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPbfListingVersionsRequest wrapper for the ListPbfListingVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListPbfListingVersions.go.html to see an example of how to use ListPbfListingVersionsRequest.
type ListPbfListingVersionsRequest struct {

	// unique PbfListing identifier
	PbfListingId *string `mandatory:"true" contributesTo:"query" name:"pbfListingId"`

	// unique PbfListingVersion identifier
	PbfListingVersionId *string `mandatory:"false" contributesTo:"query" name:"pbfListingVersionId"`

	// Matches a PbfListingVersion based on a provided semantic version name for a PbfListingVersion.
	// Each PbfListingVersion name is unique with respect to its associated PbfListing.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Matches the current version (the most recently added version with an Active
	// lifecycleState) associated with a PbfListing.
	IsCurrentVersion *bool `mandatory:"false" contributesTo:"query" name:"isCurrentVersion"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState PbfListingVersionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token for a list query returned by a previous operation
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListPbfListingVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListPbfListingVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPbfListingVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPbfListingVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPbfListingVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPbfListingVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPbfListingVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPbfListingVersionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPbfListingVersionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPbfListingVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPbfListingVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPbfListingVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPbfListingVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPbfListingVersionsResponse wrapper for the ListPbfListingVersions operation
type ListPbfListingVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PbfListingVersionsCollection instances
	PbfListingVersionsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPbfListingVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPbfListingVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPbfListingVersionsSortOrderEnum Enum with underlying type: string
type ListPbfListingVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListPbfListingVersionsSortOrderEnum
const (
	ListPbfListingVersionsSortOrderAsc  ListPbfListingVersionsSortOrderEnum = "ASC"
	ListPbfListingVersionsSortOrderDesc ListPbfListingVersionsSortOrderEnum = "DESC"
)

var mappingListPbfListingVersionsSortOrderEnum = map[string]ListPbfListingVersionsSortOrderEnum{
	"ASC":  ListPbfListingVersionsSortOrderAsc,
	"DESC": ListPbfListingVersionsSortOrderDesc,
}

var mappingListPbfListingVersionsSortOrderEnumLowerCase = map[string]ListPbfListingVersionsSortOrderEnum{
	"asc":  ListPbfListingVersionsSortOrderAsc,
	"desc": ListPbfListingVersionsSortOrderDesc,
}

// GetListPbfListingVersionsSortOrderEnumValues Enumerates the set of values for ListPbfListingVersionsSortOrderEnum
func GetListPbfListingVersionsSortOrderEnumValues() []ListPbfListingVersionsSortOrderEnum {
	values := make([]ListPbfListingVersionsSortOrderEnum, 0)
	for _, v := range mappingListPbfListingVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPbfListingVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListPbfListingVersionsSortOrderEnum
func GetListPbfListingVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPbfListingVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPbfListingVersionsSortOrderEnum(val string) (ListPbfListingVersionsSortOrderEnum, bool) {
	enum, ok := mappingListPbfListingVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPbfListingVersionsSortByEnum Enum with underlying type: string
type ListPbfListingVersionsSortByEnum string

// Set of constants representing the allowable values for ListPbfListingVersionsSortByEnum
const (
	ListPbfListingVersionsSortByTimecreated ListPbfListingVersionsSortByEnum = "timeCreated"
	ListPbfListingVersionsSortByName        ListPbfListingVersionsSortByEnum = "name"
)

var mappingListPbfListingVersionsSortByEnum = map[string]ListPbfListingVersionsSortByEnum{
	"timeCreated": ListPbfListingVersionsSortByTimecreated,
	"name":        ListPbfListingVersionsSortByName,
}

var mappingListPbfListingVersionsSortByEnumLowerCase = map[string]ListPbfListingVersionsSortByEnum{
	"timecreated": ListPbfListingVersionsSortByTimecreated,
	"name":        ListPbfListingVersionsSortByName,
}

// GetListPbfListingVersionsSortByEnumValues Enumerates the set of values for ListPbfListingVersionsSortByEnum
func GetListPbfListingVersionsSortByEnumValues() []ListPbfListingVersionsSortByEnum {
	values := make([]ListPbfListingVersionsSortByEnum, 0)
	for _, v := range mappingListPbfListingVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPbfListingVersionsSortByEnumStringValues Enumerates the set of values in String for ListPbfListingVersionsSortByEnum
func GetListPbfListingVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListPbfListingVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPbfListingVersionsSortByEnum(val string) (ListPbfListingVersionsSortByEnum, bool) {
	enum, ok := mappingListPbfListingVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
