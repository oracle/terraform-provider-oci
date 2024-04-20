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

// ListResourceProfileImpactedResourcesRequest wrapper for the ListResourceProfileImpactedResources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceProfileImpactedResources.go.html to see an example of how to use ListResourceProfileImpactedResourcesRequest.
type ListResourceProfileImpactedResourcesRequest struct {

	// OCID of the resource profile.
	ResourceProfileId *string `mandatory:"true" contributesTo:"path" name:"resourceProfileId"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListResourceProfileImpactedResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListResourceProfileImpactedResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceProfileImpactedResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceProfileImpactedResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceProfileImpactedResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceProfileImpactedResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceProfileImpactedResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceProfileImpactedResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceProfileImpactedResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceProfileImpactedResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceProfileImpactedResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceProfileImpactedResourcesResponse wrapper for the ListResourceProfileImpactedResources operation
type ListResourceProfileImpactedResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceProfileImpactedResourceCollection instances
	ResourceProfileImpactedResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceProfileImpactedResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceProfileImpactedResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceProfileImpactedResourcesSortOrderEnum Enum with underlying type: string
type ListResourceProfileImpactedResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListResourceProfileImpactedResourcesSortOrderEnum
const (
	ListResourceProfileImpactedResourcesSortOrderAsc  ListResourceProfileImpactedResourcesSortOrderEnum = "ASC"
	ListResourceProfileImpactedResourcesSortOrderDesc ListResourceProfileImpactedResourcesSortOrderEnum = "DESC"
)

var mappingListResourceProfileImpactedResourcesSortOrderEnum = map[string]ListResourceProfileImpactedResourcesSortOrderEnum{
	"ASC":  ListResourceProfileImpactedResourcesSortOrderAsc,
	"DESC": ListResourceProfileImpactedResourcesSortOrderDesc,
}

var mappingListResourceProfileImpactedResourcesSortOrderEnumLowerCase = map[string]ListResourceProfileImpactedResourcesSortOrderEnum{
	"asc":  ListResourceProfileImpactedResourcesSortOrderAsc,
	"desc": ListResourceProfileImpactedResourcesSortOrderDesc,
}

// GetListResourceProfileImpactedResourcesSortOrderEnumValues Enumerates the set of values for ListResourceProfileImpactedResourcesSortOrderEnum
func GetListResourceProfileImpactedResourcesSortOrderEnumValues() []ListResourceProfileImpactedResourcesSortOrderEnum {
	values := make([]ListResourceProfileImpactedResourcesSortOrderEnum, 0)
	for _, v := range mappingListResourceProfileImpactedResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceProfileImpactedResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListResourceProfileImpactedResourcesSortOrderEnum
func GetListResourceProfileImpactedResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceProfileImpactedResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceProfileImpactedResourcesSortOrderEnum(val string) (ListResourceProfileImpactedResourcesSortOrderEnum, bool) {
	enum, ok := mappingListResourceProfileImpactedResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceProfileImpactedResourcesSortByEnum Enum with underlying type: string
type ListResourceProfileImpactedResourcesSortByEnum string

// Set of constants representing the allowable values for ListResourceProfileImpactedResourcesSortByEnum
const (
	ListResourceProfileImpactedResourcesSortByTimecreated ListResourceProfileImpactedResourcesSortByEnum = "timeCreated"
)

var mappingListResourceProfileImpactedResourcesSortByEnum = map[string]ListResourceProfileImpactedResourcesSortByEnum{
	"timeCreated": ListResourceProfileImpactedResourcesSortByTimecreated,
}

var mappingListResourceProfileImpactedResourcesSortByEnumLowerCase = map[string]ListResourceProfileImpactedResourcesSortByEnum{
	"timecreated": ListResourceProfileImpactedResourcesSortByTimecreated,
}

// GetListResourceProfileImpactedResourcesSortByEnumValues Enumerates the set of values for ListResourceProfileImpactedResourcesSortByEnum
func GetListResourceProfileImpactedResourcesSortByEnumValues() []ListResourceProfileImpactedResourcesSortByEnum {
	values := make([]ListResourceProfileImpactedResourcesSortByEnum, 0)
	for _, v := range mappingListResourceProfileImpactedResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceProfileImpactedResourcesSortByEnumStringValues Enumerates the set of values in String for ListResourceProfileImpactedResourcesSortByEnum
func GetListResourceProfileImpactedResourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListResourceProfileImpactedResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceProfileImpactedResourcesSortByEnum(val string) (ListResourceProfileImpactedResourcesSortByEnum, bool) {
	enum, ok := mappingListResourceProfileImpactedResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
