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

// ListDetectorsRequest wrapper for the ListDetectors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDetectors.go.html to see an example of how to use ListDetectorsRequest.
type ListDetectorsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListDetectorsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListDetectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDetectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDetectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDetectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDetectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDetectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDetectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDetectorsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDetectorsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDetectorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDetectorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDetectorsResponse wrapper for the ListDetectors operation
type ListDetectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DetectorCollection instances
	DetectorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDetectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDetectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDetectorsLifecycleStateEnum Enum with underlying type: string
type ListDetectorsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDetectorsLifecycleStateEnum
const (
	ListDetectorsLifecycleStateCreating ListDetectorsLifecycleStateEnum = "CREATING"
	ListDetectorsLifecycleStateUpdating ListDetectorsLifecycleStateEnum = "UPDATING"
	ListDetectorsLifecycleStateActive   ListDetectorsLifecycleStateEnum = "ACTIVE"
	ListDetectorsLifecycleStateInactive ListDetectorsLifecycleStateEnum = "INACTIVE"
	ListDetectorsLifecycleStateDeleting ListDetectorsLifecycleStateEnum = "DELETING"
	ListDetectorsLifecycleStateDeleted  ListDetectorsLifecycleStateEnum = "DELETED"
	ListDetectorsLifecycleStateFailed   ListDetectorsLifecycleStateEnum = "FAILED"
)

var mappingListDetectorsLifecycleStateEnum = map[string]ListDetectorsLifecycleStateEnum{
	"CREATING": ListDetectorsLifecycleStateCreating,
	"UPDATING": ListDetectorsLifecycleStateUpdating,
	"ACTIVE":   ListDetectorsLifecycleStateActive,
	"INACTIVE": ListDetectorsLifecycleStateInactive,
	"DELETING": ListDetectorsLifecycleStateDeleting,
	"DELETED":  ListDetectorsLifecycleStateDeleted,
	"FAILED":   ListDetectorsLifecycleStateFailed,
}

var mappingListDetectorsLifecycleStateEnumLowerCase = map[string]ListDetectorsLifecycleStateEnum{
	"creating": ListDetectorsLifecycleStateCreating,
	"updating": ListDetectorsLifecycleStateUpdating,
	"active":   ListDetectorsLifecycleStateActive,
	"inactive": ListDetectorsLifecycleStateInactive,
	"deleting": ListDetectorsLifecycleStateDeleting,
	"deleted":  ListDetectorsLifecycleStateDeleted,
	"failed":   ListDetectorsLifecycleStateFailed,
}

// GetListDetectorsLifecycleStateEnumValues Enumerates the set of values for ListDetectorsLifecycleStateEnum
func GetListDetectorsLifecycleStateEnumValues() []ListDetectorsLifecycleStateEnum {
	values := make([]ListDetectorsLifecycleStateEnum, 0)
	for _, v := range mappingListDetectorsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDetectorsLifecycleStateEnum
func GetListDetectorsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDetectorsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorsLifecycleStateEnum(val string) (ListDetectorsLifecycleStateEnum, bool) {
	enum, ok := mappingListDetectorsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectorsSortOrderEnum Enum with underlying type: string
type ListDetectorsSortOrderEnum string

// Set of constants representing the allowable values for ListDetectorsSortOrderEnum
const (
	ListDetectorsSortOrderAsc  ListDetectorsSortOrderEnum = "ASC"
	ListDetectorsSortOrderDesc ListDetectorsSortOrderEnum = "DESC"
)

var mappingListDetectorsSortOrderEnum = map[string]ListDetectorsSortOrderEnum{
	"ASC":  ListDetectorsSortOrderAsc,
	"DESC": ListDetectorsSortOrderDesc,
}

var mappingListDetectorsSortOrderEnumLowerCase = map[string]ListDetectorsSortOrderEnum{
	"asc":  ListDetectorsSortOrderAsc,
	"desc": ListDetectorsSortOrderDesc,
}

// GetListDetectorsSortOrderEnumValues Enumerates the set of values for ListDetectorsSortOrderEnum
func GetListDetectorsSortOrderEnumValues() []ListDetectorsSortOrderEnum {
	values := make([]ListDetectorsSortOrderEnum, 0)
	for _, v := range mappingListDetectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorsSortOrderEnumStringValues Enumerates the set of values in String for ListDetectorsSortOrderEnum
func GetListDetectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDetectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorsSortOrderEnum(val string) (ListDetectorsSortOrderEnum, bool) {
	enum, ok := mappingListDetectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectorsSortByEnum Enum with underlying type: string
type ListDetectorsSortByEnum string

// Set of constants representing the allowable values for ListDetectorsSortByEnum
const (
	ListDetectorsSortByTimecreated ListDetectorsSortByEnum = "timeCreated"
	ListDetectorsSortByDisplayname ListDetectorsSortByEnum = "displayName"
)

var mappingListDetectorsSortByEnum = map[string]ListDetectorsSortByEnum{
	"timeCreated": ListDetectorsSortByTimecreated,
	"displayName": ListDetectorsSortByDisplayname,
}

var mappingListDetectorsSortByEnumLowerCase = map[string]ListDetectorsSortByEnum{
	"timecreated": ListDetectorsSortByTimecreated,
	"displayname": ListDetectorsSortByDisplayname,
}

// GetListDetectorsSortByEnumValues Enumerates the set of values for ListDetectorsSortByEnum
func GetListDetectorsSortByEnumValues() []ListDetectorsSortByEnum {
	values := make([]ListDetectorsSortByEnum, 0)
	for _, v := range mappingListDetectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorsSortByEnumStringValues Enumerates the set of values in String for ListDetectorsSortByEnum
func GetListDetectorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDetectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorsSortByEnum(val string) (ListDetectorsSortByEnum, bool) {
	enum, ok := mappingListDetectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
