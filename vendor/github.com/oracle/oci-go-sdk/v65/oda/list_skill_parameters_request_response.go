// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSkillParametersRequest wrapper for the ListSkillParameters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListSkillParameters.go.html to see an example of how to use ListSkillParametersRequest.
type ListSkillParametersRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Skill identifier.
	SkillId *string `mandatory:"true" contributesTo:"path" name:"skillId"`

	// List only Parameters with this name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// List only the resources that are in this lifecycle state.
	LifecycleState ListSkillParametersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListSkillParametersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `name`.
	// The default sort order is ascending.
	SortBy ListSkillParametersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSkillParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSkillParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSkillParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSkillParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSkillParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSkillParametersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSkillParametersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSkillParametersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSkillParametersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSkillParametersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSkillParametersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSkillParametersResponse wrapper for the ListSkillParameters operation
type ListSkillParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SkillParameterCollection instances
	SkillParameterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of results that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListSkillParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSkillParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSkillParametersLifecycleStateEnum Enum with underlying type: string
type ListSkillParametersLifecycleStateEnum string

// Set of constants representing the allowable values for ListSkillParametersLifecycleStateEnum
const (
	ListSkillParametersLifecycleStateCreating ListSkillParametersLifecycleStateEnum = "CREATING"
	ListSkillParametersLifecycleStateUpdating ListSkillParametersLifecycleStateEnum = "UPDATING"
	ListSkillParametersLifecycleStateActive   ListSkillParametersLifecycleStateEnum = "ACTIVE"
	ListSkillParametersLifecycleStateInactive ListSkillParametersLifecycleStateEnum = "INACTIVE"
	ListSkillParametersLifecycleStateDeleting ListSkillParametersLifecycleStateEnum = "DELETING"
	ListSkillParametersLifecycleStateDeleted  ListSkillParametersLifecycleStateEnum = "DELETED"
	ListSkillParametersLifecycleStateFailed   ListSkillParametersLifecycleStateEnum = "FAILED"
)

var mappingListSkillParametersLifecycleStateEnum = map[string]ListSkillParametersLifecycleStateEnum{
	"CREATING": ListSkillParametersLifecycleStateCreating,
	"UPDATING": ListSkillParametersLifecycleStateUpdating,
	"ACTIVE":   ListSkillParametersLifecycleStateActive,
	"INACTIVE": ListSkillParametersLifecycleStateInactive,
	"DELETING": ListSkillParametersLifecycleStateDeleting,
	"DELETED":  ListSkillParametersLifecycleStateDeleted,
	"FAILED":   ListSkillParametersLifecycleStateFailed,
}

var mappingListSkillParametersLifecycleStateEnumLowerCase = map[string]ListSkillParametersLifecycleStateEnum{
	"creating": ListSkillParametersLifecycleStateCreating,
	"updating": ListSkillParametersLifecycleStateUpdating,
	"active":   ListSkillParametersLifecycleStateActive,
	"inactive": ListSkillParametersLifecycleStateInactive,
	"deleting": ListSkillParametersLifecycleStateDeleting,
	"deleted":  ListSkillParametersLifecycleStateDeleted,
	"failed":   ListSkillParametersLifecycleStateFailed,
}

// GetListSkillParametersLifecycleStateEnumValues Enumerates the set of values for ListSkillParametersLifecycleStateEnum
func GetListSkillParametersLifecycleStateEnumValues() []ListSkillParametersLifecycleStateEnum {
	values := make([]ListSkillParametersLifecycleStateEnum, 0)
	for _, v := range mappingListSkillParametersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSkillParametersLifecycleStateEnumStringValues Enumerates the set of values in String for ListSkillParametersLifecycleStateEnum
func GetListSkillParametersLifecycleStateEnumStringValues() []string {
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

// GetMappingListSkillParametersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSkillParametersLifecycleStateEnum(val string) (ListSkillParametersLifecycleStateEnum, bool) {
	enum, ok := mappingListSkillParametersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSkillParametersSortOrderEnum Enum with underlying type: string
type ListSkillParametersSortOrderEnum string

// Set of constants representing the allowable values for ListSkillParametersSortOrderEnum
const (
	ListSkillParametersSortOrderAsc  ListSkillParametersSortOrderEnum = "ASC"
	ListSkillParametersSortOrderDesc ListSkillParametersSortOrderEnum = "DESC"
)

var mappingListSkillParametersSortOrderEnum = map[string]ListSkillParametersSortOrderEnum{
	"ASC":  ListSkillParametersSortOrderAsc,
	"DESC": ListSkillParametersSortOrderDesc,
}

var mappingListSkillParametersSortOrderEnumLowerCase = map[string]ListSkillParametersSortOrderEnum{
	"asc":  ListSkillParametersSortOrderAsc,
	"desc": ListSkillParametersSortOrderDesc,
}

// GetListSkillParametersSortOrderEnumValues Enumerates the set of values for ListSkillParametersSortOrderEnum
func GetListSkillParametersSortOrderEnumValues() []ListSkillParametersSortOrderEnum {
	values := make([]ListSkillParametersSortOrderEnum, 0)
	for _, v := range mappingListSkillParametersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSkillParametersSortOrderEnumStringValues Enumerates the set of values in String for ListSkillParametersSortOrderEnum
func GetListSkillParametersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSkillParametersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSkillParametersSortOrderEnum(val string) (ListSkillParametersSortOrderEnum, bool) {
	enum, ok := mappingListSkillParametersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSkillParametersSortByEnum Enum with underlying type: string
type ListSkillParametersSortByEnum string

// Set of constants representing the allowable values for ListSkillParametersSortByEnum
const (
	ListSkillParametersSortByName        ListSkillParametersSortByEnum = "name"
	ListSkillParametersSortByDisplayname ListSkillParametersSortByEnum = "displayName"
	ListSkillParametersSortByType        ListSkillParametersSortByEnum = "type"
)

var mappingListSkillParametersSortByEnum = map[string]ListSkillParametersSortByEnum{
	"name":        ListSkillParametersSortByName,
	"displayName": ListSkillParametersSortByDisplayname,
	"type":        ListSkillParametersSortByType,
}

var mappingListSkillParametersSortByEnumLowerCase = map[string]ListSkillParametersSortByEnum{
	"name":        ListSkillParametersSortByName,
	"displayname": ListSkillParametersSortByDisplayname,
	"type":        ListSkillParametersSortByType,
}

// GetListSkillParametersSortByEnumValues Enumerates the set of values for ListSkillParametersSortByEnum
func GetListSkillParametersSortByEnumValues() []ListSkillParametersSortByEnum {
	values := make([]ListSkillParametersSortByEnum, 0)
	for _, v := range mappingListSkillParametersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSkillParametersSortByEnumStringValues Enumerates the set of values in String for ListSkillParametersSortByEnum
func GetListSkillParametersSortByEnumStringValues() []string {
	return []string{
		"name",
		"displayName",
		"type",
	}
}

// GetMappingListSkillParametersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSkillParametersSortByEnum(val string) (ListSkillParametersSortByEnum, bool) {
	enum, ok := mappingListSkillParametersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
