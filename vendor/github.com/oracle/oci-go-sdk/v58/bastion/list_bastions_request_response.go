// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bastion

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListBastionsRequest wrapper for the ListBastions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bastion/ListBastions.go.html to see an example of how to use ListBastionsRequest.
type ListBastionsRequest struct {

	// The unique identifier (OCID) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	BastionLifecycleState ListBastionsBastionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"bastionLifecycleState" omitEmpty:"true"`

	// The unique identifier (OCID) of the bastion in which to list resources.
	BastionId *string `mandatory:"false" contributesTo:"query" name:"bastionId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBastionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending. If no value is specified timeCreated is default.
	SortBy ListBastionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBastionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBastionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBastionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBastionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBastionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBastionsBastionLifecycleStateEnum(string(request.BastionLifecycleState)); !ok && request.BastionLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BastionLifecycleState: %s. Supported values are: %s.", request.BastionLifecycleState, strings.Join(GetListBastionsBastionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBastionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBastionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBastionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBastionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBastionsResponse wrapper for the ListBastions operation
type ListBastionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BastionSummary instances
	Items []BastionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBastionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBastionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBastionsBastionLifecycleStateEnum Enum with underlying type: string
type ListBastionsBastionLifecycleStateEnum string

// Set of constants representing the allowable values for ListBastionsBastionLifecycleStateEnum
const (
	ListBastionsBastionLifecycleStateCreating ListBastionsBastionLifecycleStateEnum = "CREATING"
	ListBastionsBastionLifecycleStateUpdating ListBastionsBastionLifecycleStateEnum = "UPDATING"
	ListBastionsBastionLifecycleStateActive   ListBastionsBastionLifecycleStateEnum = "ACTIVE"
	ListBastionsBastionLifecycleStateDeleting ListBastionsBastionLifecycleStateEnum = "DELETING"
	ListBastionsBastionLifecycleStateDeleted  ListBastionsBastionLifecycleStateEnum = "DELETED"
	ListBastionsBastionLifecycleStateFailed   ListBastionsBastionLifecycleStateEnum = "FAILED"
)

var mappingListBastionsBastionLifecycleStateEnum = map[string]ListBastionsBastionLifecycleStateEnum{
	"CREATING": ListBastionsBastionLifecycleStateCreating,
	"UPDATING": ListBastionsBastionLifecycleStateUpdating,
	"ACTIVE":   ListBastionsBastionLifecycleStateActive,
	"DELETING": ListBastionsBastionLifecycleStateDeleting,
	"DELETED":  ListBastionsBastionLifecycleStateDeleted,
	"FAILED":   ListBastionsBastionLifecycleStateFailed,
}

// GetListBastionsBastionLifecycleStateEnumValues Enumerates the set of values for ListBastionsBastionLifecycleStateEnum
func GetListBastionsBastionLifecycleStateEnumValues() []ListBastionsBastionLifecycleStateEnum {
	values := make([]ListBastionsBastionLifecycleStateEnum, 0)
	for _, v := range mappingListBastionsBastionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListBastionsBastionLifecycleStateEnumStringValues Enumerates the set of values in String for ListBastionsBastionLifecycleStateEnum
func GetListBastionsBastionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListBastionsBastionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBastionsBastionLifecycleStateEnum(val string) (ListBastionsBastionLifecycleStateEnum, bool) {
	mappingListBastionsBastionLifecycleStateEnumIgnoreCase := make(map[string]ListBastionsBastionLifecycleStateEnum)
	for k, v := range mappingListBastionsBastionLifecycleStateEnum {
		mappingListBastionsBastionLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBastionsBastionLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListBastionsSortOrderEnum Enum with underlying type: string
type ListBastionsSortOrderEnum string

// Set of constants representing the allowable values for ListBastionsSortOrderEnum
const (
	ListBastionsSortOrderAsc  ListBastionsSortOrderEnum = "ASC"
	ListBastionsSortOrderDesc ListBastionsSortOrderEnum = "DESC"
)

var mappingListBastionsSortOrderEnum = map[string]ListBastionsSortOrderEnum{
	"ASC":  ListBastionsSortOrderAsc,
	"DESC": ListBastionsSortOrderDesc,
}

// GetListBastionsSortOrderEnumValues Enumerates the set of values for ListBastionsSortOrderEnum
func GetListBastionsSortOrderEnumValues() []ListBastionsSortOrderEnum {
	values := make([]ListBastionsSortOrderEnum, 0)
	for _, v := range mappingListBastionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBastionsSortOrderEnumStringValues Enumerates the set of values in String for ListBastionsSortOrderEnum
func GetListBastionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBastionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBastionsSortOrderEnum(val string) (ListBastionsSortOrderEnum, bool) {
	mappingListBastionsSortOrderEnumIgnoreCase := make(map[string]ListBastionsSortOrderEnum)
	for k, v := range mappingListBastionsSortOrderEnum {
		mappingListBastionsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBastionsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListBastionsSortByEnum Enum with underlying type: string
type ListBastionsSortByEnum string

// Set of constants representing the allowable values for ListBastionsSortByEnum
const (
	ListBastionsSortByTimecreated ListBastionsSortByEnum = "timeCreated"
	ListBastionsSortByName        ListBastionsSortByEnum = "name"
)

var mappingListBastionsSortByEnum = map[string]ListBastionsSortByEnum{
	"timeCreated": ListBastionsSortByTimecreated,
	"name":        ListBastionsSortByName,
}

// GetListBastionsSortByEnumValues Enumerates the set of values for ListBastionsSortByEnum
func GetListBastionsSortByEnumValues() []ListBastionsSortByEnum {
	values := make([]ListBastionsSortByEnum, 0)
	for _, v := range mappingListBastionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBastionsSortByEnumStringValues Enumerates the set of values in String for ListBastionsSortByEnum
func GetListBastionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListBastionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBastionsSortByEnum(val string) (ListBastionsSortByEnum, bool) {
	mappingListBastionsSortByEnumIgnoreCase := make(map[string]ListBastionsSortByEnum)
	for k, v := range mappingListBastionsSortByEnum {
		mappingListBastionsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListBastionsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
