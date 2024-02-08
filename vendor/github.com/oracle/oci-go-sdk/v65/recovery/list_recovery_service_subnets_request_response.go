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

// ListRecoveryServiceSubnetsRequest wrapper for the ListRecoveryServiceSubnets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListRecoveryServiceSubnets.go.html to see an example of how to use ListRecoveryServiceSubnetsRequest.
type ListRecoveryServiceSubnetsRequest struct {

	// The compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the specified lifecycle state.
	// Allowed values are:
	//   - CREATING
	//   - UPDATING
	//   - ACTIVE
	//   - DELETING
	//   - DELETED
	//   - FAILED
	LifecycleState ListRecoveryServiceSubnetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire 'displayname' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The recovery service subnet OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID of the virtual cloud network (VCN) associated with the recovery service subnet.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	// Allowed values are:
	//   - ASC
	//   - DESC
	SortOrder ListRecoveryServiceSubnetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (sortOrder). Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If you do not specify a value, then TIMECREATED is used as the default sort order.
	// Allowed values are:
	//   - TIMECREATED
	//   - DISPLAYNAME
	SortBy ListRecoveryServiceSubnetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecoveryServiceSubnetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecoveryServiceSubnetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecoveryServiceSubnetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecoveryServiceSubnetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecoveryServiceSubnetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecoveryServiceSubnetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRecoveryServiceSubnetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecoveryServiceSubnetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRecoveryServiceSubnetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecoveryServiceSubnetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRecoveryServiceSubnetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecoveryServiceSubnetsResponse wrapper for the ListRecoveryServiceSubnets operation
type ListRecoveryServiceSubnetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecoveryServiceSubnetCollection instances
	RecoveryServiceSubnetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRecoveryServiceSubnetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecoveryServiceSubnetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecoveryServiceSubnetsLifecycleStateEnum Enum with underlying type: string
type ListRecoveryServiceSubnetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListRecoveryServiceSubnetsLifecycleStateEnum
const (
	ListRecoveryServiceSubnetsLifecycleStateCreating ListRecoveryServiceSubnetsLifecycleStateEnum = "CREATING"
	ListRecoveryServiceSubnetsLifecycleStateUpdating ListRecoveryServiceSubnetsLifecycleStateEnum = "UPDATING"
	ListRecoveryServiceSubnetsLifecycleStateActive   ListRecoveryServiceSubnetsLifecycleStateEnum = "ACTIVE"
	ListRecoveryServiceSubnetsLifecycleStateDeleting ListRecoveryServiceSubnetsLifecycleStateEnum = "DELETING"
	ListRecoveryServiceSubnetsLifecycleStateDeleted  ListRecoveryServiceSubnetsLifecycleStateEnum = "DELETED"
	ListRecoveryServiceSubnetsLifecycleStateFailed   ListRecoveryServiceSubnetsLifecycleStateEnum = "FAILED"
)

var mappingListRecoveryServiceSubnetsLifecycleStateEnum = map[string]ListRecoveryServiceSubnetsLifecycleStateEnum{
	"CREATING": ListRecoveryServiceSubnetsLifecycleStateCreating,
	"UPDATING": ListRecoveryServiceSubnetsLifecycleStateUpdating,
	"ACTIVE":   ListRecoveryServiceSubnetsLifecycleStateActive,
	"DELETING": ListRecoveryServiceSubnetsLifecycleStateDeleting,
	"DELETED":  ListRecoveryServiceSubnetsLifecycleStateDeleted,
	"FAILED":   ListRecoveryServiceSubnetsLifecycleStateFailed,
}

var mappingListRecoveryServiceSubnetsLifecycleStateEnumLowerCase = map[string]ListRecoveryServiceSubnetsLifecycleStateEnum{
	"creating": ListRecoveryServiceSubnetsLifecycleStateCreating,
	"updating": ListRecoveryServiceSubnetsLifecycleStateUpdating,
	"active":   ListRecoveryServiceSubnetsLifecycleStateActive,
	"deleting": ListRecoveryServiceSubnetsLifecycleStateDeleting,
	"deleted":  ListRecoveryServiceSubnetsLifecycleStateDeleted,
	"failed":   ListRecoveryServiceSubnetsLifecycleStateFailed,
}

// GetListRecoveryServiceSubnetsLifecycleStateEnumValues Enumerates the set of values for ListRecoveryServiceSubnetsLifecycleStateEnum
func GetListRecoveryServiceSubnetsLifecycleStateEnumValues() []ListRecoveryServiceSubnetsLifecycleStateEnum {
	values := make([]ListRecoveryServiceSubnetsLifecycleStateEnum, 0)
	for _, v := range mappingListRecoveryServiceSubnetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecoveryServiceSubnetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListRecoveryServiceSubnetsLifecycleStateEnum
func GetListRecoveryServiceSubnetsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListRecoveryServiceSubnetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecoveryServiceSubnetsLifecycleStateEnum(val string) (ListRecoveryServiceSubnetsLifecycleStateEnum, bool) {
	enum, ok := mappingListRecoveryServiceSubnetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecoveryServiceSubnetsSortOrderEnum Enum with underlying type: string
type ListRecoveryServiceSubnetsSortOrderEnum string

// Set of constants representing the allowable values for ListRecoveryServiceSubnetsSortOrderEnum
const (
	ListRecoveryServiceSubnetsSortOrderAsc  ListRecoveryServiceSubnetsSortOrderEnum = "ASC"
	ListRecoveryServiceSubnetsSortOrderDesc ListRecoveryServiceSubnetsSortOrderEnum = "DESC"
)

var mappingListRecoveryServiceSubnetsSortOrderEnum = map[string]ListRecoveryServiceSubnetsSortOrderEnum{
	"ASC":  ListRecoveryServiceSubnetsSortOrderAsc,
	"DESC": ListRecoveryServiceSubnetsSortOrderDesc,
}

var mappingListRecoveryServiceSubnetsSortOrderEnumLowerCase = map[string]ListRecoveryServiceSubnetsSortOrderEnum{
	"asc":  ListRecoveryServiceSubnetsSortOrderAsc,
	"desc": ListRecoveryServiceSubnetsSortOrderDesc,
}

// GetListRecoveryServiceSubnetsSortOrderEnumValues Enumerates the set of values for ListRecoveryServiceSubnetsSortOrderEnum
func GetListRecoveryServiceSubnetsSortOrderEnumValues() []ListRecoveryServiceSubnetsSortOrderEnum {
	values := make([]ListRecoveryServiceSubnetsSortOrderEnum, 0)
	for _, v := range mappingListRecoveryServiceSubnetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecoveryServiceSubnetsSortOrderEnumStringValues Enumerates the set of values in String for ListRecoveryServiceSubnetsSortOrderEnum
func GetListRecoveryServiceSubnetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRecoveryServiceSubnetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecoveryServiceSubnetsSortOrderEnum(val string) (ListRecoveryServiceSubnetsSortOrderEnum, bool) {
	enum, ok := mappingListRecoveryServiceSubnetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecoveryServiceSubnetsSortByEnum Enum with underlying type: string
type ListRecoveryServiceSubnetsSortByEnum string

// Set of constants representing the allowable values for ListRecoveryServiceSubnetsSortByEnum
const (
	ListRecoveryServiceSubnetsSortByTimecreated ListRecoveryServiceSubnetsSortByEnum = "timeCreated"
	ListRecoveryServiceSubnetsSortByDisplayname ListRecoveryServiceSubnetsSortByEnum = "displayName"
)

var mappingListRecoveryServiceSubnetsSortByEnum = map[string]ListRecoveryServiceSubnetsSortByEnum{
	"timeCreated": ListRecoveryServiceSubnetsSortByTimecreated,
	"displayName": ListRecoveryServiceSubnetsSortByDisplayname,
}

var mappingListRecoveryServiceSubnetsSortByEnumLowerCase = map[string]ListRecoveryServiceSubnetsSortByEnum{
	"timecreated": ListRecoveryServiceSubnetsSortByTimecreated,
	"displayname": ListRecoveryServiceSubnetsSortByDisplayname,
}

// GetListRecoveryServiceSubnetsSortByEnumValues Enumerates the set of values for ListRecoveryServiceSubnetsSortByEnum
func GetListRecoveryServiceSubnetsSortByEnumValues() []ListRecoveryServiceSubnetsSortByEnum {
	values := make([]ListRecoveryServiceSubnetsSortByEnum, 0)
	for _, v := range mappingListRecoveryServiceSubnetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecoveryServiceSubnetsSortByEnumStringValues Enumerates the set of values in String for ListRecoveryServiceSubnetsSortByEnum
func GetListRecoveryServiceSubnetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRecoveryServiceSubnetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecoveryServiceSubnetsSortByEnum(val string) (ListRecoveryServiceSubnetsSortByEnum, bool) {
	enum, ok := mappingListRecoveryServiceSubnetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
