// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAssociatedOciCacheClustersRequest wrapper for the ListAssociatedOciCacheClusters operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListAssociatedOciCacheClusters.go.html to see an example of how to use ListAssociatedOciCacheClustersRequest.
type ListAssociatedOciCacheClustersRequest struct {

	// Unique OCI Cache Config Set identifier.
	OciCacheConfigSetId *string `mandatory:"true" contributesTo:"path" name:"ociCacheConfigSetId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAssociatedOciCacheClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAssociatedOciCacheClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociatedOciCacheClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociatedOciCacheClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociatedOciCacheClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociatedOciCacheClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssociatedOciCacheClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssociatedOciCacheClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssociatedOciCacheClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssociatedOciCacheClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssociatedOciCacheClustersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssociatedOciCacheClustersResponse wrapper for the ListAssociatedOciCacheClusters operation
type ListAssociatedOciCacheClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssociatedOciCacheClusterCollection instances
	AssociatedOciCacheClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssociatedOciCacheClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociatedOciCacheClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociatedOciCacheClustersSortOrderEnum Enum with underlying type: string
type ListAssociatedOciCacheClustersSortOrderEnum string

// Set of constants representing the allowable values for ListAssociatedOciCacheClustersSortOrderEnum
const (
	ListAssociatedOciCacheClustersSortOrderAsc  ListAssociatedOciCacheClustersSortOrderEnum = "ASC"
	ListAssociatedOciCacheClustersSortOrderDesc ListAssociatedOciCacheClustersSortOrderEnum = "DESC"
)

var mappingListAssociatedOciCacheClustersSortOrderEnum = map[string]ListAssociatedOciCacheClustersSortOrderEnum{
	"ASC":  ListAssociatedOciCacheClustersSortOrderAsc,
	"DESC": ListAssociatedOciCacheClustersSortOrderDesc,
}

var mappingListAssociatedOciCacheClustersSortOrderEnumLowerCase = map[string]ListAssociatedOciCacheClustersSortOrderEnum{
	"asc":  ListAssociatedOciCacheClustersSortOrderAsc,
	"desc": ListAssociatedOciCacheClustersSortOrderDesc,
}

// GetListAssociatedOciCacheClustersSortOrderEnumValues Enumerates the set of values for ListAssociatedOciCacheClustersSortOrderEnum
func GetListAssociatedOciCacheClustersSortOrderEnumValues() []ListAssociatedOciCacheClustersSortOrderEnum {
	values := make([]ListAssociatedOciCacheClustersSortOrderEnum, 0)
	for _, v := range mappingListAssociatedOciCacheClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociatedOciCacheClustersSortOrderEnumStringValues Enumerates the set of values in String for ListAssociatedOciCacheClustersSortOrderEnum
func GetListAssociatedOciCacheClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssociatedOciCacheClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociatedOciCacheClustersSortOrderEnum(val string) (ListAssociatedOciCacheClustersSortOrderEnum, bool) {
	enum, ok := mappingListAssociatedOciCacheClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssociatedOciCacheClustersSortByEnum Enum with underlying type: string
type ListAssociatedOciCacheClustersSortByEnum string

// Set of constants representing the allowable values for ListAssociatedOciCacheClustersSortByEnum
const (
	ListAssociatedOciCacheClustersSortByTimecreated ListAssociatedOciCacheClustersSortByEnum = "timeCreated"
	ListAssociatedOciCacheClustersSortByDisplayname ListAssociatedOciCacheClustersSortByEnum = "displayName"
)

var mappingListAssociatedOciCacheClustersSortByEnum = map[string]ListAssociatedOciCacheClustersSortByEnum{
	"timeCreated": ListAssociatedOciCacheClustersSortByTimecreated,
	"displayName": ListAssociatedOciCacheClustersSortByDisplayname,
}

var mappingListAssociatedOciCacheClustersSortByEnumLowerCase = map[string]ListAssociatedOciCacheClustersSortByEnum{
	"timecreated": ListAssociatedOciCacheClustersSortByTimecreated,
	"displayname": ListAssociatedOciCacheClustersSortByDisplayname,
}

// GetListAssociatedOciCacheClustersSortByEnumValues Enumerates the set of values for ListAssociatedOciCacheClustersSortByEnum
func GetListAssociatedOciCacheClustersSortByEnumValues() []ListAssociatedOciCacheClustersSortByEnum {
	values := make([]ListAssociatedOciCacheClustersSortByEnum, 0)
	for _, v := range mappingListAssociatedOciCacheClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociatedOciCacheClustersSortByEnumStringValues Enumerates the set of values in String for ListAssociatedOciCacheClustersSortByEnum
func GetListAssociatedOciCacheClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAssociatedOciCacheClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociatedOciCacheClustersSortByEnum(val string) (ListAssociatedOciCacheClustersSortByEnum, bool) {
	enum, ok := mappingListAssociatedOciCacheClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
