// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMetastoresRequest wrapper for the ListMetastores operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListMetastores.go.html to see an example of how to use ListMetastoresRequest.
type ListMetastoresRequest struct {

	// The OCID of the compartment where you want to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListMetastoresLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMetastoresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListMetastoresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMetastoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMetastoresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMetastoresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMetastoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMetastoresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMetastoresLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMetastoresLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMetastoresSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMetastoresSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMetastoresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMetastoresSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMetastoresResponse wrapper for the ListMetastores operation
type ListMetastoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []MetastoreSummary instances
	Items []MetastoreSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMetastoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMetastoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMetastoresLifecycleStateEnum Enum with underlying type: string
type ListMetastoresLifecycleStateEnum string

// Set of constants representing the allowable values for ListMetastoresLifecycleStateEnum
const (
	ListMetastoresLifecycleStateCreating ListMetastoresLifecycleStateEnum = "CREATING"
	ListMetastoresLifecycleStateActive   ListMetastoresLifecycleStateEnum = "ACTIVE"
	ListMetastoresLifecycleStateInactive ListMetastoresLifecycleStateEnum = "INACTIVE"
	ListMetastoresLifecycleStateUpdating ListMetastoresLifecycleStateEnum = "UPDATING"
	ListMetastoresLifecycleStateDeleting ListMetastoresLifecycleStateEnum = "DELETING"
	ListMetastoresLifecycleStateDeleted  ListMetastoresLifecycleStateEnum = "DELETED"
	ListMetastoresLifecycleStateFailed   ListMetastoresLifecycleStateEnum = "FAILED"
	ListMetastoresLifecycleStateMoving   ListMetastoresLifecycleStateEnum = "MOVING"
)

var mappingListMetastoresLifecycleStateEnum = map[string]ListMetastoresLifecycleStateEnum{
	"CREATING": ListMetastoresLifecycleStateCreating,
	"ACTIVE":   ListMetastoresLifecycleStateActive,
	"INACTIVE": ListMetastoresLifecycleStateInactive,
	"UPDATING": ListMetastoresLifecycleStateUpdating,
	"DELETING": ListMetastoresLifecycleStateDeleting,
	"DELETED":  ListMetastoresLifecycleStateDeleted,
	"FAILED":   ListMetastoresLifecycleStateFailed,
	"MOVING":   ListMetastoresLifecycleStateMoving,
}

var mappingListMetastoresLifecycleStateEnumLowerCase = map[string]ListMetastoresLifecycleStateEnum{
	"creating": ListMetastoresLifecycleStateCreating,
	"active":   ListMetastoresLifecycleStateActive,
	"inactive": ListMetastoresLifecycleStateInactive,
	"updating": ListMetastoresLifecycleStateUpdating,
	"deleting": ListMetastoresLifecycleStateDeleting,
	"deleted":  ListMetastoresLifecycleStateDeleted,
	"failed":   ListMetastoresLifecycleStateFailed,
	"moving":   ListMetastoresLifecycleStateMoving,
}

// GetListMetastoresLifecycleStateEnumValues Enumerates the set of values for ListMetastoresLifecycleStateEnum
func GetListMetastoresLifecycleStateEnumValues() []ListMetastoresLifecycleStateEnum {
	values := make([]ListMetastoresLifecycleStateEnum, 0)
	for _, v := range mappingListMetastoresLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetastoresLifecycleStateEnumStringValues Enumerates the set of values in String for ListMetastoresLifecycleStateEnum
func GetListMetastoresLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"MOVING",
	}
}

// GetMappingListMetastoresLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetastoresLifecycleStateEnum(val string) (ListMetastoresLifecycleStateEnum, bool) {
	enum, ok := mappingListMetastoresLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMetastoresSortOrderEnum Enum with underlying type: string
type ListMetastoresSortOrderEnum string

// Set of constants representing the allowable values for ListMetastoresSortOrderEnum
const (
	ListMetastoresSortOrderAsc  ListMetastoresSortOrderEnum = "ASC"
	ListMetastoresSortOrderDesc ListMetastoresSortOrderEnum = "DESC"
)

var mappingListMetastoresSortOrderEnum = map[string]ListMetastoresSortOrderEnum{
	"ASC":  ListMetastoresSortOrderAsc,
	"DESC": ListMetastoresSortOrderDesc,
}

var mappingListMetastoresSortOrderEnumLowerCase = map[string]ListMetastoresSortOrderEnum{
	"asc":  ListMetastoresSortOrderAsc,
	"desc": ListMetastoresSortOrderDesc,
}

// GetListMetastoresSortOrderEnumValues Enumerates the set of values for ListMetastoresSortOrderEnum
func GetListMetastoresSortOrderEnumValues() []ListMetastoresSortOrderEnum {
	values := make([]ListMetastoresSortOrderEnum, 0)
	for _, v := range mappingListMetastoresSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetastoresSortOrderEnumStringValues Enumerates the set of values in String for ListMetastoresSortOrderEnum
func GetListMetastoresSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMetastoresSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetastoresSortOrderEnum(val string) (ListMetastoresSortOrderEnum, bool) {
	enum, ok := mappingListMetastoresSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMetastoresSortByEnum Enum with underlying type: string
type ListMetastoresSortByEnum string

// Set of constants representing the allowable values for ListMetastoresSortByEnum
const (
	ListMetastoresSortByTimecreated ListMetastoresSortByEnum = "TIMECREATED"
	ListMetastoresSortByDisplayname ListMetastoresSortByEnum = "DISPLAYNAME"
)

var mappingListMetastoresSortByEnum = map[string]ListMetastoresSortByEnum{
	"TIMECREATED": ListMetastoresSortByTimecreated,
	"DISPLAYNAME": ListMetastoresSortByDisplayname,
}

var mappingListMetastoresSortByEnumLowerCase = map[string]ListMetastoresSortByEnum{
	"timecreated": ListMetastoresSortByTimecreated,
	"displayname": ListMetastoresSortByDisplayname,
}

// GetListMetastoresSortByEnumValues Enumerates the set of values for ListMetastoresSortByEnum
func GetListMetastoresSortByEnumValues() []ListMetastoresSortByEnum {
	values := make([]ListMetastoresSortByEnum, 0)
	for _, v := range mappingListMetastoresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetastoresSortByEnumStringValues Enumerates the set of values in String for ListMetastoresSortByEnum
func GetListMetastoresSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListMetastoresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetastoresSortByEnum(val string) (ListMetastoresSortByEnum, bool) {
	enum, ok := mappingListMetastoresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
