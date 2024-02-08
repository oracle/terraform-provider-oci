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

// ListTechniquesRequest wrapper for the ListTechniques operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTechniques.go.html to see an example of how to use ListTechniquesRequest.
type ListTechniquesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the list of tactics given.
	Tactics []string `contributesTo:"query" name:"tactics" collectionFormat:"multi"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListTechniquesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTechniquesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListTechniquesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTechniquesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTechniquesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTechniquesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTechniquesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTechniquesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTechniquesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTechniquesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTechniquesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTechniquesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTechniquesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTechniquesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTechniquesResponse wrapper for the ListTechniques operation
type ListTechniquesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TechniqueCollection instances
	TechniqueCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTechniquesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTechniquesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTechniquesLifecycleStateEnum Enum with underlying type: string
type ListTechniquesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTechniquesLifecycleStateEnum
const (
	ListTechniquesLifecycleStateCreating ListTechniquesLifecycleStateEnum = "CREATING"
	ListTechniquesLifecycleStateUpdating ListTechniquesLifecycleStateEnum = "UPDATING"
	ListTechniquesLifecycleStateActive   ListTechniquesLifecycleStateEnum = "ACTIVE"
	ListTechniquesLifecycleStateInactive ListTechniquesLifecycleStateEnum = "INACTIVE"
	ListTechniquesLifecycleStateDeleting ListTechniquesLifecycleStateEnum = "DELETING"
	ListTechniquesLifecycleStateDeleted  ListTechniquesLifecycleStateEnum = "DELETED"
	ListTechniquesLifecycleStateFailed   ListTechniquesLifecycleStateEnum = "FAILED"
)

var mappingListTechniquesLifecycleStateEnum = map[string]ListTechniquesLifecycleStateEnum{
	"CREATING": ListTechniquesLifecycleStateCreating,
	"UPDATING": ListTechniquesLifecycleStateUpdating,
	"ACTIVE":   ListTechniquesLifecycleStateActive,
	"INACTIVE": ListTechniquesLifecycleStateInactive,
	"DELETING": ListTechniquesLifecycleStateDeleting,
	"DELETED":  ListTechniquesLifecycleStateDeleted,
	"FAILED":   ListTechniquesLifecycleStateFailed,
}

var mappingListTechniquesLifecycleStateEnumLowerCase = map[string]ListTechniquesLifecycleStateEnum{
	"creating": ListTechniquesLifecycleStateCreating,
	"updating": ListTechniquesLifecycleStateUpdating,
	"active":   ListTechniquesLifecycleStateActive,
	"inactive": ListTechniquesLifecycleStateInactive,
	"deleting": ListTechniquesLifecycleStateDeleting,
	"deleted":  ListTechniquesLifecycleStateDeleted,
	"failed":   ListTechniquesLifecycleStateFailed,
}

// GetListTechniquesLifecycleStateEnumValues Enumerates the set of values for ListTechniquesLifecycleStateEnum
func GetListTechniquesLifecycleStateEnumValues() []ListTechniquesLifecycleStateEnum {
	values := make([]ListTechniquesLifecycleStateEnum, 0)
	for _, v := range mappingListTechniquesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTechniquesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTechniquesLifecycleStateEnum
func GetListTechniquesLifecycleStateEnumStringValues() []string {
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

// GetMappingListTechniquesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTechniquesLifecycleStateEnum(val string) (ListTechniquesLifecycleStateEnum, bool) {
	enum, ok := mappingListTechniquesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTechniquesSortOrderEnum Enum with underlying type: string
type ListTechniquesSortOrderEnum string

// Set of constants representing the allowable values for ListTechniquesSortOrderEnum
const (
	ListTechniquesSortOrderAsc  ListTechniquesSortOrderEnum = "ASC"
	ListTechniquesSortOrderDesc ListTechniquesSortOrderEnum = "DESC"
)

var mappingListTechniquesSortOrderEnum = map[string]ListTechniquesSortOrderEnum{
	"ASC":  ListTechniquesSortOrderAsc,
	"DESC": ListTechniquesSortOrderDesc,
}

var mappingListTechniquesSortOrderEnumLowerCase = map[string]ListTechniquesSortOrderEnum{
	"asc":  ListTechniquesSortOrderAsc,
	"desc": ListTechniquesSortOrderDesc,
}

// GetListTechniquesSortOrderEnumValues Enumerates the set of values for ListTechniquesSortOrderEnum
func GetListTechniquesSortOrderEnumValues() []ListTechniquesSortOrderEnum {
	values := make([]ListTechniquesSortOrderEnum, 0)
	for _, v := range mappingListTechniquesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTechniquesSortOrderEnumStringValues Enumerates the set of values in String for ListTechniquesSortOrderEnum
func GetListTechniquesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTechniquesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTechniquesSortOrderEnum(val string) (ListTechniquesSortOrderEnum, bool) {
	enum, ok := mappingListTechniquesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTechniquesSortByEnum Enum with underlying type: string
type ListTechniquesSortByEnum string

// Set of constants representing the allowable values for ListTechniquesSortByEnum
const (
	ListTechniquesSortByDisplayname ListTechniquesSortByEnum = "displayName"
)

var mappingListTechniquesSortByEnum = map[string]ListTechniquesSortByEnum{
	"displayName": ListTechniquesSortByDisplayname,
}

var mappingListTechniquesSortByEnumLowerCase = map[string]ListTechniquesSortByEnum{
	"displayname": ListTechniquesSortByDisplayname,
}

// GetListTechniquesSortByEnumValues Enumerates the set of values for ListTechniquesSortByEnum
func GetListTechniquesSortByEnumValues() []ListTechniquesSortByEnum {
	values := make([]ListTechniquesSortByEnum, 0)
	for _, v := range mappingListTechniquesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTechniquesSortByEnumStringValues Enumerates the set of values in String for ListTechniquesSortByEnum
func GetListTechniquesSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListTechniquesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTechniquesSortByEnum(val string) (ListTechniquesSortByEnum, bool) {
	enum, ok := mappingListTechniquesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
