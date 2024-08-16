// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRedisClusterNodesRequest wrapper for the ListRedisClusterNodes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListRedisClusterNodes.go.html to see an example of how to use ListRedisClusterNodesRequest.
type ListRedisClusterNodesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.
	RedisClusterId *string `mandatory:"true" contributesTo:"path" name:"redisClusterId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRedisClusterNodesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListRedisClusterNodesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRedisClusterNodesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRedisClusterNodesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRedisClusterNodesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRedisClusterNodesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRedisClusterNodesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRedisClusterNodesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRedisClusterNodesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRedisClusterNodesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRedisClusterNodesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRedisClusterNodesResponse wrapper for the ListRedisClusterNodes operation
type ListRedisClusterNodesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RedisNodeCollection instances
	RedisNodeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRedisClusterNodesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRedisClusterNodesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRedisClusterNodesSortOrderEnum Enum with underlying type: string
type ListRedisClusterNodesSortOrderEnum string

// Set of constants representing the allowable values for ListRedisClusterNodesSortOrderEnum
const (
	ListRedisClusterNodesSortOrderAsc  ListRedisClusterNodesSortOrderEnum = "ASC"
	ListRedisClusterNodesSortOrderDesc ListRedisClusterNodesSortOrderEnum = "DESC"
)

var mappingListRedisClusterNodesSortOrderEnum = map[string]ListRedisClusterNodesSortOrderEnum{
	"ASC":  ListRedisClusterNodesSortOrderAsc,
	"DESC": ListRedisClusterNodesSortOrderDesc,
}

var mappingListRedisClusterNodesSortOrderEnumLowerCase = map[string]ListRedisClusterNodesSortOrderEnum{
	"asc":  ListRedisClusterNodesSortOrderAsc,
	"desc": ListRedisClusterNodesSortOrderDesc,
}

// GetListRedisClusterNodesSortOrderEnumValues Enumerates the set of values for ListRedisClusterNodesSortOrderEnum
func GetListRedisClusterNodesSortOrderEnumValues() []ListRedisClusterNodesSortOrderEnum {
	values := make([]ListRedisClusterNodesSortOrderEnum, 0)
	for _, v := range mappingListRedisClusterNodesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRedisClusterNodesSortOrderEnumStringValues Enumerates the set of values in String for ListRedisClusterNodesSortOrderEnum
func GetListRedisClusterNodesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRedisClusterNodesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRedisClusterNodesSortOrderEnum(val string) (ListRedisClusterNodesSortOrderEnum, bool) {
	enum, ok := mappingListRedisClusterNodesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRedisClusterNodesSortByEnum Enum with underlying type: string
type ListRedisClusterNodesSortByEnum string

// Set of constants representing the allowable values for ListRedisClusterNodesSortByEnum
const (
	ListRedisClusterNodesSortByTimecreated ListRedisClusterNodesSortByEnum = "timeCreated"
	ListRedisClusterNodesSortByDisplayname ListRedisClusterNodesSortByEnum = "displayName"
)

var mappingListRedisClusterNodesSortByEnum = map[string]ListRedisClusterNodesSortByEnum{
	"timeCreated": ListRedisClusterNodesSortByTimecreated,
	"displayName": ListRedisClusterNodesSortByDisplayname,
}

var mappingListRedisClusterNodesSortByEnumLowerCase = map[string]ListRedisClusterNodesSortByEnum{
	"timecreated": ListRedisClusterNodesSortByTimecreated,
	"displayname": ListRedisClusterNodesSortByDisplayname,
}

// GetListRedisClusterNodesSortByEnumValues Enumerates the set of values for ListRedisClusterNodesSortByEnum
func GetListRedisClusterNodesSortByEnumValues() []ListRedisClusterNodesSortByEnum {
	values := make([]ListRedisClusterNodesSortByEnum, 0)
	for _, v := range mappingListRedisClusterNodesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRedisClusterNodesSortByEnumStringValues Enumerates the set of values in String for ListRedisClusterNodesSortByEnum
func GetListRedisClusterNodesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRedisClusterNodesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRedisClusterNodesSortByEnum(val string) (ListRedisClusterNodesSortByEnum, bool) {
	enum, ok := mappingListRedisClusterNodesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
