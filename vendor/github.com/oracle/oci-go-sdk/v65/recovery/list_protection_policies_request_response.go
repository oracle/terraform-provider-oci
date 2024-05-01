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

// ListProtectionPoliciesRequest wrapper for the ListProtectionPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListProtectionPolicies.go.html to see an example of how to use ListProtectionPoliciesRequest.
type ListProtectionPoliciesRequest struct {

	// The compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListProtectionPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire 'displayname' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The protection policy OCID.
	ProtectionPolicyId *string `mandatory:"false" contributesTo:"query" name:"protectionPolicyId"`

	// A filter to return only the policies that match the owner as 'Customer' or 'Oracle'.
	Owner ListProtectionPoliciesOwnerEnum `mandatory:"false" contributesTo:"query" name:"owner" omitEmpty:"true"`

	// The maximum number of items to return. Specify a value greater than 4.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	// Allowed values are:
	//   - ASC
	//   - DESC
	SortOrder ListProtectionPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (sortOrder). Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If you do not specify a value, then TIMECREATED is used as the default sort order.
	// Allowed values are:
	//   - TIMECREATED
	//   - DISPLAYNAME
	SortBy ListProtectionPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProtectionPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProtectionPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProtectionPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProtectionPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProtectionPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProtectionPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListProtectionPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectionPoliciesOwnerEnum(string(request.Owner)); !ok && request.Owner != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Owner: %s. Supported values are: %s.", request.Owner, strings.Join(GetListProtectionPoliciesOwnerEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectionPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProtectionPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectionPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProtectionPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProtectionPoliciesResponse wrapper for the ListProtectionPolicies operation
type ListProtectionPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProtectionPolicyCollection instances
	ProtectionPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProtectionPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProtectionPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProtectionPoliciesLifecycleStateEnum Enum with underlying type: string
type ListProtectionPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListProtectionPoliciesLifecycleStateEnum
const (
	ListProtectionPoliciesLifecycleStateCreating        ListProtectionPoliciesLifecycleStateEnum = "CREATING"
	ListProtectionPoliciesLifecycleStateUpdating        ListProtectionPoliciesLifecycleStateEnum = "UPDATING"
	ListProtectionPoliciesLifecycleStateActive          ListProtectionPoliciesLifecycleStateEnum = "ACTIVE"
	ListProtectionPoliciesLifecycleStateDeleteScheduled ListProtectionPoliciesLifecycleStateEnum = "DELETE_SCHEDULED"
	ListProtectionPoliciesLifecycleStateDeleting        ListProtectionPoliciesLifecycleStateEnum = "DELETING"
	ListProtectionPoliciesLifecycleStateDeleted         ListProtectionPoliciesLifecycleStateEnum = "DELETED"
	ListProtectionPoliciesLifecycleStateFailed          ListProtectionPoliciesLifecycleStateEnum = "FAILED"
)

var mappingListProtectionPoliciesLifecycleStateEnum = map[string]ListProtectionPoliciesLifecycleStateEnum{
	"CREATING":         ListProtectionPoliciesLifecycleStateCreating,
	"UPDATING":         ListProtectionPoliciesLifecycleStateUpdating,
	"ACTIVE":           ListProtectionPoliciesLifecycleStateActive,
	"DELETE_SCHEDULED": ListProtectionPoliciesLifecycleStateDeleteScheduled,
	"DELETING":         ListProtectionPoliciesLifecycleStateDeleting,
	"DELETED":          ListProtectionPoliciesLifecycleStateDeleted,
	"FAILED":           ListProtectionPoliciesLifecycleStateFailed,
}

var mappingListProtectionPoliciesLifecycleStateEnumLowerCase = map[string]ListProtectionPoliciesLifecycleStateEnum{
	"creating":         ListProtectionPoliciesLifecycleStateCreating,
	"updating":         ListProtectionPoliciesLifecycleStateUpdating,
	"active":           ListProtectionPoliciesLifecycleStateActive,
	"delete_scheduled": ListProtectionPoliciesLifecycleStateDeleteScheduled,
	"deleting":         ListProtectionPoliciesLifecycleStateDeleting,
	"deleted":          ListProtectionPoliciesLifecycleStateDeleted,
	"failed":           ListProtectionPoliciesLifecycleStateFailed,
}

// GetListProtectionPoliciesLifecycleStateEnumValues Enumerates the set of values for ListProtectionPoliciesLifecycleStateEnum
func GetListProtectionPoliciesLifecycleStateEnumValues() []ListProtectionPoliciesLifecycleStateEnum {
	values := make([]ListProtectionPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListProtectionPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListProtectionPoliciesLifecycleStateEnum
func GetListProtectionPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETE_SCHEDULED",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListProtectionPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionPoliciesLifecycleStateEnum(val string) (ListProtectionPoliciesLifecycleStateEnum, bool) {
	enum, ok := mappingListProtectionPoliciesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProtectionPoliciesOwnerEnum Enum with underlying type: string
type ListProtectionPoliciesOwnerEnum string

// Set of constants representing the allowable values for ListProtectionPoliciesOwnerEnum
const (
	ListProtectionPoliciesOwnerOracle   ListProtectionPoliciesOwnerEnum = "oracle"
	ListProtectionPoliciesOwnerCustomer ListProtectionPoliciesOwnerEnum = "customer"
)

var mappingListProtectionPoliciesOwnerEnum = map[string]ListProtectionPoliciesOwnerEnum{
	"oracle":   ListProtectionPoliciesOwnerOracle,
	"customer": ListProtectionPoliciesOwnerCustomer,
}

var mappingListProtectionPoliciesOwnerEnumLowerCase = map[string]ListProtectionPoliciesOwnerEnum{
	"oracle":   ListProtectionPoliciesOwnerOracle,
	"customer": ListProtectionPoliciesOwnerCustomer,
}

// GetListProtectionPoliciesOwnerEnumValues Enumerates the set of values for ListProtectionPoliciesOwnerEnum
func GetListProtectionPoliciesOwnerEnumValues() []ListProtectionPoliciesOwnerEnum {
	values := make([]ListProtectionPoliciesOwnerEnum, 0)
	for _, v := range mappingListProtectionPoliciesOwnerEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionPoliciesOwnerEnumStringValues Enumerates the set of values in String for ListProtectionPoliciesOwnerEnum
func GetListProtectionPoliciesOwnerEnumStringValues() []string {
	return []string{
		"oracle",
		"customer",
	}
}

// GetMappingListProtectionPoliciesOwnerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionPoliciesOwnerEnum(val string) (ListProtectionPoliciesOwnerEnum, bool) {
	enum, ok := mappingListProtectionPoliciesOwnerEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProtectionPoliciesSortOrderEnum Enum with underlying type: string
type ListProtectionPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListProtectionPoliciesSortOrderEnum
const (
	ListProtectionPoliciesSortOrderAsc  ListProtectionPoliciesSortOrderEnum = "ASC"
	ListProtectionPoliciesSortOrderDesc ListProtectionPoliciesSortOrderEnum = "DESC"
)

var mappingListProtectionPoliciesSortOrderEnum = map[string]ListProtectionPoliciesSortOrderEnum{
	"ASC":  ListProtectionPoliciesSortOrderAsc,
	"DESC": ListProtectionPoliciesSortOrderDesc,
}

var mappingListProtectionPoliciesSortOrderEnumLowerCase = map[string]ListProtectionPoliciesSortOrderEnum{
	"asc":  ListProtectionPoliciesSortOrderAsc,
	"desc": ListProtectionPoliciesSortOrderDesc,
}

// GetListProtectionPoliciesSortOrderEnumValues Enumerates the set of values for ListProtectionPoliciesSortOrderEnum
func GetListProtectionPoliciesSortOrderEnumValues() []ListProtectionPoliciesSortOrderEnum {
	values := make([]ListProtectionPoliciesSortOrderEnum, 0)
	for _, v := range mappingListProtectionPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListProtectionPoliciesSortOrderEnum
func GetListProtectionPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProtectionPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionPoliciesSortOrderEnum(val string) (ListProtectionPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListProtectionPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProtectionPoliciesSortByEnum Enum with underlying type: string
type ListProtectionPoliciesSortByEnum string

// Set of constants representing the allowable values for ListProtectionPoliciesSortByEnum
const (
	ListProtectionPoliciesSortByTimecreated ListProtectionPoliciesSortByEnum = "timeCreated"
	ListProtectionPoliciesSortByDisplayname ListProtectionPoliciesSortByEnum = "displayName"
)

var mappingListProtectionPoliciesSortByEnum = map[string]ListProtectionPoliciesSortByEnum{
	"timeCreated": ListProtectionPoliciesSortByTimecreated,
	"displayName": ListProtectionPoliciesSortByDisplayname,
}

var mappingListProtectionPoliciesSortByEnumLowerCase = map[string]ListProtectionPoliciesSortByEnum{
	"timecreated": ListProtectionPoliciesSortByTimecreated,
	"displayname": ListProtectionPoliciesSortByDisplayname,
}

// GetListProtectionPoliciesSortByEnumValues Enumerates the set of values for ListProtectionPoliciesSortByEnum
func GetListProtectionPoliciesSortByEnumValues() []ListProtectionPoliciesSortByEnum {
	values := make([]ListProtectionPoliciesSortByEnum, 0)
	for _, v := range mappingListProtectionPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionPoliciesSortByEnumStringValues Enumerates the set of values in String for ListProtectionPoliciesSortByEnum
func GetListProtectionPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListProtectionPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionPoliciesSortByEnum(val string) (ListProtectionPoliciesSortByEnum, bool) {
	enum, ok := mappingListProtectionPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
