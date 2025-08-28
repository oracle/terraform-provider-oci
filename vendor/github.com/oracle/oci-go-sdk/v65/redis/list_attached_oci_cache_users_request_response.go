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

// ListAttachedOciCacheUsersRequest wrapper for the ListAttachedOciCacheUsers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListAttachedOciCacheUsers.go.html to see an example of how to use ListAttachedOciCacheUsersRequest.
type ListAttachedOciCacheUsersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.
	RedisClusterId *string `mandatory:"true" contributesTo:"path" name:"redisClusterId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAttachedOciCacheUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAttachedOciCacheUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAttachedOciCacheUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAttachedOciCacheUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAttachedOciCacheUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAttachedOciCacheUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAttachedOciCacheUsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAttachedOciCacheUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAttachedOciCacheUsersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttachedOciCacheUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAttachedOciCacheUsersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAttachedOciCacheUsersResponse wrapper for the ListAttachedOciCacheUsers operation
type ListAttachedOciCacheUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AttachedOciCacheUser instances
	Items []AttachedOciCacheUser `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAttachedOciCacheUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAttachedOciCacheUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAttachedOciCacheUsersSortOrderEnum Enum with underlying type: string
type ListAttachedOciCacheUsersSortOrderEnum string

// Set of constants representing the allowable values for ListAttachedOciCacheUsersSortOrderEnum
const (
	ListAttachedOciCacheUsersSortOrderAsc  ListAttachedOciCacheUsersSortOrderEnum = "ASC"
	ListAttachedOciCacheUsersSortOrderDesc ListAttachedOciCacheUsersSortOrderEnum = "DESC"
)

var mappingListAttachedOciCacheUsersSortOrderEnum = map[string]ListAttachedOciCacheUsersSortOrderEnum{
	"ASC":  ListAttachedOciCacheUsersSortOrderAsc,
	"DESC": ListAttachedOciCacheUsersSortOrderDesc,
}

var mappingListAttachedOciCacheUsersSortOrderEnumLowerCase = map[string]ListAttachedOciCacheUsersSortOrderEnum{
	"asc":  ListAttachedOciCacheUsersSortOrderAsc,
	"desc": ListAttachedOciCacheUsersSortOrderDesc,
}

// GetListAttachedOciCacheUsersSortOrderEnumValues Enumerates the set of values for ListAttachedOciCacheUsersSortOrderEnum
func GetListAttachedOciCacheUsersSortOrderEnumValues() []ListAttachedOciCacheUsersSortOrderEnum {
	values := make([]ListAttachedOciCacheUsersSortOrderEnum, 0)
	for _, v := range mappingListAttachedOciCacheUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttachedOciCacheUsersSortOrderEnumStringValues Enumerates the set of values in String for ListAttachedOciCacheUsersSortOrderEnum
func GetListAttachedOciCacheUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAttachedOciCacheUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttachedOciCacheUsersSortOrderEnum(val string) (ListAttachedOciCacheUsersSortOrderEnum, bool) {
	enum, ok := mappingListAttachedOciCacheUsersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttachedOciCacheUsersSortByEnum Enum with underlying type: string
type ListAttachedOciCacheUsersSortByEnum string

// Set of constants representing the allowable values for ListAttachedOciCacheUsersSortByEnum
const (
	ListAttachedOciCacheUsersSortByTimecreated ListAttachedOciCacheUsersSortByEnum = "timeCreated"
	ListAttachedOciCacheUsersSortByDisplayname ListAttachedOciCacheUsersSortByEnum = "displayName"
)

var mappingListAttachedOciCacheUsersSortByEnum = map[string]ListAttachedOciCacheUsersSortByEnum{
	"timeCreated": ListAttachedOciCacheUsersSortByTimecreated,
	"displayName": ListAttachedOciCacheUsersSortByDisplayname,
}

var mappingListAttachedOciCacheUsersSortByEnumLowerCase = map[string]ListAttachedOciCacheUsersSortByEnum{
	"timecreated": ListAttachedOciCacheUsersSortByTimecreated,
	"displayname": ListAttachedOciCacheUsersSortByDisplayname,
}

// GetListAttachedOciCacheUsersSortByEnumValues Enumerates the set of values for ListAttachedOciCacheUsersSortByEnum
func GetListAttachedOciCacheUsersSortByEnumValues() []ListAttachedOciCacheUsersSortByEnum {
	values := make([]ListAttachedOciCacheUsersSortByEnum, 0)
	for _, v := range mappingListAttachedOciCacheUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttachedOciCacheUsersSortByEnumStringValues Enumerates the set of values in String for ListAttachedOciCacheUsersSortByEnum
func GetListAttachedOciCacheUsersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAttachedOciCacheUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttachedOciCacheUsersSortByEnum(val string) (ListAttachedOciCacheUsersSortByEnum, bool) {
	enum, ok := mappingListAttachedOciCacheUsersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
