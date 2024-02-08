// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmcontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApmDomainsRequest wrapper for the ListApmDomains operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmcontrolplane/ListApmDomains.go.html to see an example of how to use ListApmDomainsRequest.
type ListApmDomainsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given life-cycle state.
	LifecycleState ListApmDomainsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListApmDomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, timeCreated is default.
	SortBy ListApmDomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApmDomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApmDomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApmDomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApmDomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApmDomainsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListApmDomainsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListApmDomainsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApmDomainsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApmDomainsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApmDomainsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApmDomainsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApmDomainsResponse wrapper for the ListApmDomains operation
type ListApmDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ApmDomainSummary instances
	Items []ApmDomainSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApmDomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApmDomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApmDomainsLifecycleStateEnum Enum with underlying type: string
type ListApmDomainsLifecycleStateEnum string

// Set of constants representing the allowable values for ListApmDomainsLifecycleStateEnum
const (
	ListApmDomainsLifecycleStateCreating ListApmDomainsLifecycleStateEnum = "CREATING"
	ListApmDomainsLifecycleStateUpdating ListApmDomainsLifecycleStateEnum = "UPDATING"
	ListApmDomainsLifecycleStateActive   ListApmDomainsLifecycleStateEnum = "ACTIVE"
	ListApmDomainsLifecycleStateDeleting ListApmDomainsLifecycleStateEnum = "DELETING"
	ListApmDomainsLifecycleStateDeleted  ListApmDomainsLifecycleStateEnum = "DELETED"
	ListApmDomainsLifecycleStateFailed   ListApmDomainsLifecycleStateEnum = "FAILED"
)

var mappingListApmDomainsLifecycleStateEnum = map[string]ListApmDomainsLifecycleStateEnum{
	"CREATING": ListApmDomainsLifecycleStateCreating,
	"UPDATING": ListApmDomainsLifecycleStateUpdating,
	"ACTIVE":   ListApmDomainsLifecycleStateActive,
	"DELETING": ListApmDomainsLifecycleStateDeleting,
	"DELETED":  ListApmDomainsLifecycleStateDeleted,
	"FAILED":   ListApmDomainsLifecycleStateFailed,
}

var mappingListApmDomainsLifecycleStateEnumLowerCase = map[string]ListApmDomainsLifecycleStateEnum{
	"creating": ListApmDomainsLifecycleStateCreating,
	"updating": ListApmDomainsLifecycleStateUpdating,
	"active":   ListApmDomainsLifecycleStateActive,
	"deleting": ListApmDomainsLifecycleStateDeleting,
	"deleted":  ListApmDomainsLifecycleStateDeleted,
	"failed":   ListApmDomainsLifecycleStateFailed,
}

// GetListApmDomainsLifecycleStateEnumValues Enumerates the set of values for ListApmDomainsLifecycleStateEnum
func GetListApmDomainsLifecycleStateEnumValues() []ListApmDomainsLifecycleStateEnum {
	values := make([]ListApmDomainsLifecycleStateEnum, 0)
	for _, v := range mappingListApmDomainsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListApmDomainsLifecycleStateEnumStringValues Enumerates the set of values in String for ListApmDomainsLifecycleStateEnum
func GetListApmDomainsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListApmDomainsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApmDomainsLifecycleStateEnum(val string) (ListApmDomainsLifecycleStateEnum, bool) {
	enum, ok := mappingListApmDomainsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApmDomainsSortOrderEnum Enum with underlying type: string
type ListApmDomainsSortOrderEnum string

// Set of constants representing the allowable values for ListApmDomainsSortOrderEnum
const (
	ListApmDomainsSortOrderAsc  ListApmDomainsSortOrderEnum = "ASC"
	ListApmDomainsSortOrderDesc ListApmDomainsSortOrderEnum = "DESC"
)

var mappingListApmDomainsSortOrderEnum = map[string]ListApmDomainsSortOrderEnum{
	"ASC":  ListApmDomainsSortOrderAsc,
	"DESC": ListApmDomainsSortOrderDesc,
}

var mappingListApmDomainsSortOrderEnumLowerCase = map[string]ListApmDomainsSortOrderEnum{
	"asc":  ListApmDomainsSortOrderAsc,
	"desc": ListApmDomainsSortOrderDesc,
}

// GetListApmDomainsSortOrderEnumValues Enumerates the set of values for ListApmDomainsSortOrderEnum
func GetListApmDomainsSortOrderEnumValues() []ListApmDomainsSortOrderEnum {
	values := make([]ListApmDomainsSortOrderEnum, 0)
	for _, v := range mappingListApmDomainsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApmDomainsSortOrderEnumStringValues Enumerates the set of values in String for ListApmDomainsSortOrderEnum
func GetListApmDomainsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApmDomainsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApmDomainsSortOrderEnum(val string) (ListApmDomainsSortOrderEnum, bool) {
	enum, ok := mappingListApmDomainsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApmDomainsSortByEnum Enum with underlying type: string
type ListApmDomainsSortByEnum string

// Set of constants representing the allowable values for ListApmDomainsSortByEnum
const (
	ListApmDomainsSortByTimecreated ListApmDomainsSortByEnum = "timeCreated"
	ListApmDomainsSortByDisplayname ListApmDomainsSortByEnum = "displayName"
)

var mappingListApmDomainsSortByEnum = map[string]ListApmDomainsSortByEnum{
	"timeCreated": ListApmDomainsSortByTimecreated,
	"displayName": ListApmDomainsSortByDisplayname,
}

var mappingListApmDomainsSortByEnumLowerCase = map[string]ListApmDomainsSortByEnum{
	"timecreated": ListApmDomainsSortByTimecreated,
	"displayname": ListApmDomainsSortByDisplayname,
}

// GetListApmDomainsSortByEnumValues Enumerates the set of values for ListApmDomainsSortByEnum
func GetListApmDomainsSortByEnumValues() []ListApmDomainsSortByEnum {
	values := make([]ListApmDomainsSortByEnum, 0)
	for _, v := range mappingListApmDomainsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApmDomainsSortByEnumStringValues Enumerates the set of values in String for ListApmDomainsSortByEnum
func GetListApmDomainsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListApmDomainsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApmDomainsSortByEnum(val string) (ListApmDomainsSortByEnum, bool) {
	enum, ok := mappingListApmDomainsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
