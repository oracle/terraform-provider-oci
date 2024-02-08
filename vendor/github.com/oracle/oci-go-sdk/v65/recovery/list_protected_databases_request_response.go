// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProtectedDatabasesRequest wrapper for the ListProtectedDatabases operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListProtectedDatabases.go.html to see an example of how to use ListProtectedDatabasesRequest.
type ListProtectedDatabasesRequest struct {

	// The compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the specified lifecycle state.
	LifecycleState ListProtectedDatabasesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire 'displayname' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The protected database OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The protection policy OCID.
	ProtectionPolicyId *string `mandatory:"false" contributesTo:"query" name:"protectionPolicyId"`

	// The recovery service subnet OCID.
	RecoveryServiceSubnetId *string `mandatory:"false" contributesTo:"query" name:"recoveryServiceSubnetId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	// Allowed values are:
	//   - ASC
	//   - DESC
	SortOrder ListProtectedDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (sortOrder). Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If you do not specify a value, then TIMECREATED is used as the default sort order.
	// Allowed values are:
	//   - TIMECREATED
	//   - DISPLAYNAME
	SortBy ListProtectedDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProtectedDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProtectedDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProtectedDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProtectedDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProtectedDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProtectedDatabasesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListProtectedDatabasesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectedDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProtectedDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectedDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProtectedDatabasesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProtectedDatabasesResponse wrapper for the ListProtectedDatabases operation
type ListProtectedDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProtectedDatabaseCollection instances
	ProtectedDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProtectedDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProtectedDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProtectedDatabasesLifecycleStateEnum Enum with underlying type: string
type ListProtectedDatabasesLifecycleStateEnum string

// Set of constants representing the allowable values for ListProtectedDatabasesLifecycleStateEnum
const (
	ListProtectedDatabasesLifecycleStateCreating ListProtectedDatabasesLifecycleStateEnum = "CREATING"
	ListProtectedDatabasesLifecycleStateUpdating ListProtectedDatabasesLifecycleStateEnum = "UPDATING"
	ListProtectedDatabasesLifecycleStateActive   ListProtectedDatabasesLifecycleStateEnum = "ACTIVE"
	ListProtectedDatabasesLifecycleStateDeleting ListProtectedDatabasesLifecycleStateEnum = "DELETING"
	ListProtectedDatabasesLifecycleStateDeleted  ListProtectedDatabasesLifecycleStateEnum = "DELETED"
	ListProtectedDatabasesLifecycleStateFailed   ListProtectedDatabasesLifecycleStateEnum = "FAILED"
)

var mappingListProtectedDatabasesLifecycleStateEnum = map[string]ListProtectedDatabasesLifecycleStateEnum{
	"CREATING": ListProtectedDatabasesLifecycleStateCreating,
	"UPDATING": ListProtectedDatabasesLifecycleStateUpdating,
	"ACTIVE":   ListProtectedDatabasesLifecycleStateActive,
	"DELETING": ListProtectedDatabasesLifecycleStateDeleting,
	"DELETED":  ListProtectedDatabasesLifecycleStateDeleted,
	"FAILED":   ListProtectedDatabasesLifecycleStateFailed,
}

var mappingListProtectedDatabasesLifecycleStateEnumLowerCase = map[string]ListProtectedDatabasesLifecycleStateEnum{
	"creating": ListProtectedDatabasesLifecycleStateCreating,
	"updating": ListProtectedDatabasesLifecycleStateUpdating,
	"active":   ListProtectedDatabasesLifecycleStateActive,
	"deleting": ListProtectedDatabasesLifecycleStateDeleting,
	"deleted":  ListProtectedDatabasesLifecycleStateDeleted,
	"failed":   ListProtectedDatabasesLifecycleStateFailed,
}

// GetListProtectedDatabasesLifecycleStateEnumValues Enumerates the set of values for ListProtectedDatabasesLifecycleStateEnum
func GetListProtectedDatabasesLifecycleStateEnumValues() []ListProtectedDatabasesLifecycleStateEnum {
	values := make([]ListProtectedDatabasesLifecycleStateEnum, 0)
	for _, v := range mappingListProtectedDatabasesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectedDatabasesLifecycleStateEnumStringValues Enumerates the set of values in String for ListProtectedDatabasesLifecycleStateEnum
func GetListProtectedDatabasesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListProtectedDatabasesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectedDatabasesLifecycleStateEnum(val string) (ListProtectedDatabasesLifecycleStateEnum, bool) {
	enum, ok := mappingListProtectedDatabasesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProtectedDatabasesSortOrderEnum Enum with underlying type: string
type ListProtectedDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListProtectedDatabasesSortOrderEnum
const (
	ListProtectedDatabasesSortOrderAsc  ListProtectedDatabasesSortOrderEnum = "ASC"
	ListProtectedDatabasesSortOrderDesc ListProtectedDatabasesSortOrderEnum = "DESC"
)

var mappingListProtectedDatabasesSortOrderEnum = map[string]ListProtectedDatabasesSortOrderEnum{
	"ASC":  ListProtectedDatabasesSortOrderAsc,
	"DESC": ListProtectedDatabasesSortOrderDesc,
}

var mappingListProtectedDatabasesSortOrderEnumLowerCase = map[string]ListProtectedDatabasesSortOrderEnum{
	"asc":  ListProtectedDatabasesSortOrderAsc,
	"desc": ListProtectedDatabasesSortOrderDesc,
}

// GetListProtectedDatabasesSortOrderEnumValues Enumerates the set of values for ListProtectedDatabasesSortOrderEnum
func GetListProtectedDatabasesSortOrderEnumValues() []ListProtectedDatabasesSortOrderEnum {
	values := make([]ListProtectedDatabasesSortOrderEnum, 0)
	for _, v := range mappingListProtectedDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectedDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListProtectedDatabasesSortOrderEnum
func GetListProtectedDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProtectedDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectedDatabasesSortOrderEnum(val string) (ListProtectedDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListProtectedDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProtectedDatabasesSortByEnum Enum with underlying type: string
type ListProtectedDatabasesSortByEnum string

// Set of constants representing the allowable values for ListProtectedDatabasesSortByEnum
const (
	ListProtectedDatabasesSortByTimecreated ListProtectedDatabasesSortByEnum = "timeCreated"
	ListProtectedDatabasesSortByDisplayname ListProtectedDatabasesSortByEnum = "displayName"
)

var mappingListProtectedDatabasesSortByEnum = map[string]ListProtectedDatabasesSortByEnum{
	"timeCreated": ListProtectedDatabasesSortByTimecreated,
	"displayName": ListProtectedDatabasesSortByDisplayname,
}

var mappingListProtectedDatabasesSortByEnumLowerCase = map[string]ListProtectedDatabasesSortByEnum{
	"timecreated": ListProtectedDatabasesSortByTimecreated,
	"displayname": ListProtectedDatabasesSortByDisplayname,
}

// GetListProtectedDatabasesSortByEnumValues Enumerates the set of values for ListProtectedDatabasesSortByEnum
func GetListProtectedDatabasesSortByEnumValues() []ListProtectedDatabasesSortByEnum {
	values := make([]ListProtectedDatabasesSortByEnum, 0)
	for _, v := range mappingListProtectedDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectedDatabasesSortByEnumStringValues Enumerates the set of values in String for ListProtectedDatabasesSortByEnum
func GetListProtectedDatabasesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListProtectedDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectedDatabasesSortByEnum(val string) (ListProtectedDatabasesSortByEnum, bool) {
	enum, ok := mappingListProtectedDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
