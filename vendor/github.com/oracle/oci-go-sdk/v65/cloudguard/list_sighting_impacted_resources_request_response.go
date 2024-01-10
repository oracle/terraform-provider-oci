// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSightingImpactedResourcesRequest wrapper for the ListSightingImpactedResources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSightingImpactedResources.go.html to see an example of how to use ListSightingImpactedResourcesRequest.
type ListSightingImpactedResourcesRequest struct {

	// OCID of the sighting.
	SightingId *string `mandatory:"true" contributesTo:"path" name:"sightingId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSightingImpactedResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListSightingImpactedResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSightingImpactedResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSightingImpactedResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSightingImpactedResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSightingImpactedResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSightingImpactedResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSightingImpactedResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSightingImpactedResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSightingImpactedResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSightingImpactedResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSightingImpactedResourcesResponse wrapper for the ListSightingImpactedResources operation
type ListSightingImpactedResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SightingImpactedResourceCollection instances
	SightingImpactedResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSightingImpactedResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSightingImpactedResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSightingImpactedResourcesSortOrderEnum Enum with underlying type: string
type ListSightingImpactedResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSightingImpactedResourcesSortOrderEnum
const (
	ListSightingImpactedResourcesSortOrderAsc  ListSightingImpactedResourcesSortOrderEnum = "ASC"
	ListSightingImpactedResourcesSortOrderDesc ListSightingImpactedResourcesSortOrderEnum = "DESC"
)

var mappingListSightingImpactedResourcesSortOrderEnum = map[string]ListSightingImpactedResourcesSortOrderEnum{
	"ASC":  ListSightingImpactedResourcesSortOrderAsc,
	"DESC": ListSightingImpactedResourcesSortOrderDesc,
}

var mappingListSightingImpactedResourcesSortOrderEnumLowerCase = map[string]ListSightingImpactedResourcesSortOrderEnum{
	"asc":  ListSightingImpactedResourcesSortOrderAsc,
	"desc": ListSightingImpactedResourcesSortOrderDesc,
}

// GetListSightingImpactedResourcesSortOrderEnumValues Enumerates the set of values for ListSightingImpactedResourcesSortOrderEnum
func GetListSightingImpactedResourcesSortOrderEnumValues() []ListSightingImpactedResourcesSortOrderEnum {
	values := make([]ListSightingImpactedResourcesSortOrderEnum, 0)
	for _, v := range mappingListSightingImpactedResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSightingImpactedResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListSightingImpactedResourcesSortOrderEnum
func GetListSightingImpactedResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSightingImpactedResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSightingImpactedResourcesSortOrderEnum(val string) (ListSightingImpactedResourcesSortOrderEnum, bool) {
	enum, ok := mappingListSightingImpactedResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSightingImpactedResourcesSortByEnum Enum with underlying type: string
type ListSightingImpactedResourcesSortByEnum string

// Set of constants representing the allowable values for ListSightingImpactedResourcesSortByEnum
const (
	ListSightingImpactedResourcesSortByTimecreated ListSightingImpactedResourcesSortByEnum = "timeCreated"
)

var mappingListSightingImpactedResourcesSortByEnum = map[string]ListSightingImpactedResourcesSortByEnum{
	"timeCreated": ListSightingImpactedResourcesSortByTimecreated,
}

var mappingListSightingImpactedResourcesSortByEnumLowerCase = map[string]ListSightingImpactedResourcesSortByEnum{
	"timecreated": ListSightingImpactedResourcesSortByTimecreated,
}

// GetListSightingImpactedResourcesSortByEnumValues Enumerates the set of values for ListSightingImpactedResourcesSortByEnum
func GetListSightingImpactedResourcesSortByEnumValues() []ListSightingImpactedResourcesSortByEnum {
	values := make([]ListSightingImpactedResourcesSortByEnum, 0)
	for _, v := range mappingListSightingImpactedResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSightingImpactedResourcesSortByEnumStringValues Enumerates the set of values in String for ListSightingImpactedResourcesSortByEnum
func GetListSightingImpactedResourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListSightingImpactedResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSightingImpactedResourcesSortByEnum(val string) (ListSightingImpactedResourcesSortByEnum, bool) {
	enum, ok := mappingListSightingImpactedResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
