// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDomainGovernancesRequest wrapper for the ListDomainGovernances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListDomainGovernances.go.html to see an example of how to use ListDomainGovernancesRequest.
type ListDomainGovernancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The domain OCID.
	DomainId *string `mandatory:"false" contributesTo:"query" name:"domainId"`

	// The domain governance OCID.
	DomainGovernanceId *string `mandatory:"false" contributesTo:"query" name:"domainGovernanceId"`

	// The lifecycle state of the resource.
	LifecycleState ListDomainGovernancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that exactly match the name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order can be provided.
	// * The default order for timeCreated is descending.
	// * The default order for displayName is ascending.
	// * If no value is specified, timeCreated is the default.
	SortBy ListDomainGovernancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListDomainGovernancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDomainGovernancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDomainGovernancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDomainGovernancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDomainGovernancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDomainGovernancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDomainGovernancesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDomainGovernancesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDomainGovernancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDomainGovernancesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDomainGovernancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDomainGovernancesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDomainGovernancesResponse wrapper for the ListDomainGovernances operation
type ListDomainGovernancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DomainGovernanceCollection instances
	DomainGovernanceCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDomainGovernancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDomainGovernancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDomainGovernancesLifecycleStateEnum Enum with underlying type: string
type ListDomainGovernancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDomainGovernancesLifecycleStateEnum
const (
	ListDomainGovernancesLifecycleStateCreating   ListDomainGovernancesLifecycleStateEnum = "CREATING"
	ListDomainGovernancesLifecycleStateActive     ListDomainGovernancesLifecycleStateEnum = "ACTIVE"
	ListDomainGovernancesLifecycleStateInactive   ListDomainGovernancesLifecycleStateEnum = "INACTIVE"
	ListDomainGovernancesLifecycleStateUpdating   ListDomainGovernancesLifecycleStateEnum = "UPDATING"
	ListDomainGovernancesLifecycleStateFailed     ListDomainGovernancesLifecycleStateEnum = "FAILED"
	ListDomainGovernancesLifecycleStateTerminated ListDomainGovernancesLifecycleStateEnum = "TERMINATED"
)

var mappingListDomainGovernancesLifecycleStateEnum = map[string]ListDomainGovernancesLifecycleStateEnum{
	"CREATING":   ListDomainGovernancesLifecycleStateCreating,
	"ACTIVE":     ListDomainGovernancesLifecycleStateActive,
	"INACTIVE":   ListDomainGovernancesLifecycleStateInactive,
	"UPDATING":   ListDomainGovernancesLifecycleStateUpdating,
	"FAILED":     ListDomainGovernancesLifecycleStateFailed,
	"TERMINATED": ListDomainGovernancesLifecycleStateTerminated,
}

var mappingListDomainGovernancesLifecycleStateEnumLowerCase = map[string]ListDomainGovernancesLifecycleStateEnum{
	"creating":   ListDomainGovernancesLifecycleStateCreating,
	"active":     ListDomainGovernancesLifecycleStateActive,
	"inactive":   ListDomainGovernancesLifecycleStateInactive,
	"updating":   ListDomainGovernancesLifecycleStateUpdating,
	"failed":     ListDomainGovernancesLifecycleStateFailed,
	"terminated": ListDomainGovernancesLifecycleStateTerminated,
}

// GetListDomainGovernancesLifecycleStateEnumValues Enumerates the set of values for ListDomainGovernancesLifecycleStateEnum
func GetListDomainGovernancesLifecycleStateEnumValues() []ListDomainGovernancesLifecycleStateEnum {
	values := make([]ListDomainGovernancesLifecycleStateEnum, 0)
	for _, v := range mappingListDomainGovernancesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDomainGovernancesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDomainGovernancesLifecycleStateEnum
func GetListDomainGovernancesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"FAILED",
		"TERMINATED",
	}
}

// GetMappingListDomainGovernancesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDomainGovernancesLifecycleStateEnum(val string) (ListDomainGovernancesLifecycleStateEnum, bool) {
	enum, ok := mappingListDomainGovernancesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDomainGovernancesSortByEnum Enum with underlying type: string
type ListDomainGovernancesSortByEnum string

// Set of constants representing the allowable values for ListDomainGovernancesSortByEnum
const (
	ListDomainGovernancesSortByTimecreated ListDomainGovernancesSortByEnum = "timeCreated"
	ListDomainGovernancesSortByDisplayname ListDomainGovernancesSortByEnum = "displayName"
)

var mappingListDomainGovernancesSortByEnum = map[string]ListDomainGovernancesSortByEnum{
	"timeCreated": ListDomainGovernancesSortByTimecreated,
	"displayName": ListDomainGovernancesSortByDisplayname,
}

var mappingListDomainGovernancesSortByEnumLowerCase = map[string]ListDomainGovernancesSortByEnum{
	"timecreated": ListDomainGovernancesSortByTimecreated,
	"displayname": ListDomainGovernancesSortByDisplayname,
}

// GetListDomainGovernancesSortByEnumValues Enumerates the set of values for ListDomainGovernancesSortByEnum
func GetListDomainGovernancesSortByEnumValues() []ListDomainGovernancesSortByEnum {
	values := make([]ListDomainGovernancesSortByEnum, 0)
	for _, v := range mappingListDomainGovernancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDomainGovernancesSortByEnumStringValues Enumerates the set of values in String for ListDomainGovernancesSortByEnum
func GetListDomainGovernancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDomainGovernancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDomainGovernancesSortByEnum(val string) (ListDomainGovernancesSortByEnum, bool) {
	enum, ok := mappingListDomainGovernancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDomainGovernancesSortOrderEnum Enum with underlying type: string
type ListDomainGovernancesSortOrderEnum string

// Set of constants representing the allowable values for ListDomainGovernancesSortOrderEnum
const (
	ListDomainGovernancesSortOrderAsc  ListDomainGovernancesSortOrderEnum = "ASC"
	ListDomainGovernancesSortOrderDesc ListDomainGovernancesSortOrderEnum = "DESC"
)

var mappingListDomainGovernancesSortOrderEnum = map[string]ListDomainGovernancesSortOrderEnum{
	"ASC":  ListDomainGovernancesSortOrderAsc,
	"DESC": ListDomainGovernancesSortOrderDesc,
}

var mappingListDomainGovernancesSortOrderEnumLowerCase = map[string]ListDomainGovernancesSortOrderEnum{
	"asc":  ListDomainGovernancesSortOrderAsc,
	"desc": ListDomainGovernancesSortOrderDesc,
}

// GetListDomainGovernancesSortOrderEnumValues Enumerates the set of values for ListDomainGovernancesSortOrderEnum
func GetListDomainGovernancesSortOrderEnumValues() []ListDomainGovernancesSortOrderEnum {
	values := make([]ListDomainGovernancesSortOrderEnum, 0)
	for _, v := range mappingListDomainGovernancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDomainGovernancesSortOrderEnumStringValues Enumerates the set of values in String for ListDomainGovernancesSortOrderEnum
func GetListDomainGovernancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDomainGovernancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDomainGovernancesSortOrderEnum(val string) (ListDomainGovernancesSortOrderEnum, bool) {
	enum, ok := mappingListDomainGovernancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
