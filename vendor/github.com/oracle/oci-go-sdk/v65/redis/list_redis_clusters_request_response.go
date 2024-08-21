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

// ListRedisClustersRequest wrapper for the ListRedisClusters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListRedisClusters.go.html to see an example of how to use ListRedisClustersRequest.
type ListRedisClustersRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState RedisClusterLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRedisClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListRedisClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRedisClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRedisClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRedisClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRedisClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRedisClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRedisClusterLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRedisClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRedisClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRedisClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRedisClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRedisClustersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRedisClustersResponse wrapper for the ListRedisClusters operation
type ListRedisClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RedisClusterCollection instances
	RedisClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRedisClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRedisClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRedisClustersSortOrderEnum Enum with underlying type: string
type ListRedisClustersSortOrderEnum string

// Set of constants representing the allowable values for ListRedisClustersSortOrderEnum
const (
	ListRedisClustersSortOrderAsc  ListRedisClustersSortOrderEnum = "ASC"
	ListRedisClustersSortOrderDesc ListRedisClustersSortOrderEnum = "DESC"
)

var mappingListRedisClustersSortOrderEnum = map[string]ListRedisClustersSortOrderEnum{
	"ASC":  ListRedisClustersSortOrderAsc,
	"DESC": ListRedisClustersSortOrderDesc,
}

var mappingListRedisClustersSortOrderEnumLowerCase = map[string]ListRedisClustersSortOrderEnum{
	"asc":  ListRedisClustersSortOrderAsc,
	"desc": ListRedisClustersSortOrderDesc,
}

// GetListRedisClustersSortOrderEnumValues Enumerates the set of values for ListRedisClustersSortOrderEnum
func GetListRedisClustersSortOrderEnumValues() []ListRedisClustersSortOrderEnum {
	values := make([]ListRedisClustersSortOrderEnum, 0)
	for _, v := range mappingListRedisClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRedisClustersSortOrderEnumStringValues Enumerates the set of values in String for ListRedisClustersSortOrderEnum
func GetListRedisClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRedisClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRedisClustersSortOrderEnum(val string) (ListRedisClustersSortOrderEnum, bool) {
	enum, ok := mappingListRedisClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRedisClustersSortByEnum Enum with underlying type: string
type ListRedisClustersSortByEnum string

// Set of constants representing the allowable values for ListRedisClustersSortByEnum
const (
	ListRedisClustersSortByTimecreated ListRedisClustersSortByEnum = "timeCreated"
	ListRedisClustersSortByDisplayname ListRedisClustersSortByEnum = "displayName"
)

var mappingListRedisClustersSortByEnum = map[string]ListRedisClustersSortByEnum{
	"timeCreated": ListRedisClustersSortByTimecreated,
	"displayName": ListRedisClustersSortByDisplayname,
}

var mappingListRedisClustersSortByEnumLowerCase = map[string]ListRedisClustersSortByEnum{
	"timecreated": ListRedisClustersSortByTimecreated,
	"displayname": ListRedisClustersSortByDisplayname,
}

// GetListRedisClustersSortByEnumValues Enumerates the set of values for ListRedisClustersSortByEnum
func GetListRedisClustersSortByEnumValues() []ListRedisClustersSortByEnum {
	values := make([]ListRedisClustersSortByEnum, 0)
	for _, v := range mappingListRedisClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRedisClustersSortByEnumStringValues Enumerates the set of values in String for ListRedisClustersSortByEnum
func GetListRedisClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRedisClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRedisClustersSortByEnum(val string) (ListRedisClustersSortByEnum, bool) {
	enum, ok := mappingListRedisClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
