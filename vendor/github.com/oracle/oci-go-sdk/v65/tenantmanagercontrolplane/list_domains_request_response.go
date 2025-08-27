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

// ListDomainsRequest wrapper for the ListDomains operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListDomains.go.html to see an example of how to use ListDomainsRequest.
type ListDomainsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The domain OCID.
	DomainId *string `mandatory:"false" contributesTo:"query" name:"domainId"`

	// The lifecycle state of the resource.
	LifecycleState ListDomainsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The status of the domain.
	Status DomainStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

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
	SortBy ListDomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListDomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDomainsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDomainsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDomainsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDomainStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetDomainStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDomainsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDomainsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDomainsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDomainsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDomainsResponse wrapper for the ListDomains operation
type ListDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DomainCollection instances
	DomainCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDomainsLifecycleStateEnum Enum with underlying type: string
type ListDomainsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDomainsLifecycleStateEnum
const (
	ListDomainsLifecycleStateCreating   ListDomainsLifecycleStateEnum = "CREATING"
	ListDomainsLifecycleStateActive     ListDomainsLifecycleStateEnum = "ACTIVE"
	ListDomainsLifecycleStateInactive   ListDomainsLifecycleStateEnum = "INACTIVE"
	ListDomainsLifecycleStateUpdating   ListDomainsLifecycleStateEnum = "UPDATING"
	ListDomainsLifecycleStateFailed     ListDomainsLifecycleStateEnum = "FAILED"
	ListDomainsLifecycleStateTerminated ListDomainsLifecycleStateEnum = "TERMINATED"
)

var mappingListDomainsLifecycleStateEnum = map[string]ListDomainsLifecycleStateEnum{
	"CREATING":   ListDomainsLifecycleStateCreating,
	"ACTIVE":     ListDomainsLifecycleStateActive,
	"INACTIVE":   ListDomainsLifecycleStateInactive,
	"UPDATING":   ListDomainsLifecycleStateUpdating,
	"FAILED":     ListDomainsLifecycleStateFailed,
	"TERMINATED": ListDomainsLifecycleStateTerminated,
}

var mappingListDomainsLifecycleStateEnumLowerCase = map[string]ListDomainsLifecycleStateEnum{
	"creating":   ListDomainsLifecycleStateCreating,
	"active":     ListDomainsLifecycleStateActive,
	"inactive":   ListDomainsLifecycleStateInactive,
	"updating":   ListDomainsLifecycleStateUpdating,
	"failed":     ListDomainsLifecycleStateFailed,
	"terminated": ListDomainsLifecycleStateTerminated,
}

// GetListDomainsLifecycleStateEnumValues Enumerates the set of values for ListDomainsLifecycleStateEnum
func GetListDomainsLifecycleStateEnumValues() []ListDomainsLifecycleStateEnum {
	values := make([]ListDomainsLifecycleStateEnum, 0)
	for _, v := range mappingListDomainsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDomainsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDomainsLifecycleStateEnum
func GetListDomainsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"FAILED",
		"TERMINATED",
	}
}

// GetMappingListDomainsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDomainsLifecycleStateEnum(val string) (ListDomainsLifecycleStateEnum, bool) {
	enum, ok := mappingListDomainsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDomainsSortByEnum Enum with underlying type: string
type ListDomainsSortByEnum string

// Set of constants representing the allowable values for ListDomainsSortByEnum
const (
	ListDomainsSortByTimecreated ListDomainsSortByEnum = "timeCreated"
	ListDomainsSortByDisplayname ListDomainsSortByEnum = "displayName"
)

var mappingListDomainsSortByEnum = map[string]ListDomainsSortByEnum{
	"timeCreated": ListDomainsSortByTimecreated,
	"displayName": ListDomainsSortByDisplayname,
}

var mappingListDomainsSortByEnumLowerCase = map[string]ListDomainsSortByEnum{
	"timecreated": ListDomainsSortByTimecreated,
	"displayname": ListDomainsSortByDisplayname,
}

// GetListDomainsSortByEnumValues Enumerates the set of values for ListDomainsSortByEnum
func GetListDomainsSortByEnumValues() []ListDomainsSortByEnum {
	values := make([]ListDomainsSortByEnum, 0)
	for _, v := range mappingListDomainsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDomainsSortByEnumStringValues Enumerates the set of values in String for ListDomainsSortByEnum
func GetListDomainsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDomainsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDomainsSortByEnum(val string) (ListDomainsSortByEnum, bool) {
	enum, ok := mappingListDomainsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDomainsSortOrderEnum Enum with underlying type: string
type ListDomainsSortOrderEnum string

// Set of constants representing the allowable values for ListDomainsSortOrderEnum
const (
	ListDomainsSortOrderAsc  ListDomainsSortOrderEnum = "ASC"
	ListDomainsSortOrderDesc ListDomainsSortOrderEnum = "DESC"
)

var mappingListDomainsSortOrderEnum = map[string]ListDomainsSortOrderEnum{
	"ASC":  ListDomainsSortOrderAsc,
	"DESC": ListDomainsSortOrderDesc,
}

var mappingListDomainsSortOrderEnumLowerCase = map[string]ListDomainsSortOrderEnum{
	"asc":  ListDomainsSortOrderAsc,
	"desc": ListDomainsSortOrderDesc,
}

// GetListDomainsSortOrderEnumValues Enumerates the set of values for ListDomainsSortOrderEnum
func GetListDomainsSortOrderEnumValues() []ListDomainsSortOrderEnum {
	values := make([]ListDomainsSortOrderEnum, 0)
	for _, v := range mappingListDomainsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDomainsSortOrderEnumStringValues Enumerates the set of values in String for ListDomainsSortOrderEnum
func GetListDomainsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDomainsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDomainsSortOrderEnum(val string) (ListDomainsSortOrderEnum, bool) {
	enum, ok := mappingListDomainsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
