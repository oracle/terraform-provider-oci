// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUsersRequest wrapper for the ListUsers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListUsers.go.html to see an example of how to use ListUsersRequest.
type ListUsersRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The id of the identity provider.
	IdentityProviderId *string `mandatory:"false" contributesTo:"query" name:"identityProviderId"`

	// The id of a user in the identity provider.
	ExternalIdentifier *string `mandatory:"false" contributesTo:"query" name:"externalIdentifier"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState UserLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUsersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUsersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUserLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetUserLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUsersResponse wrapper for the ListUsers operation
type ListUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []User instances
	Items []User `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUsersSortByEnum Enum with underlying type: string
type ListUsersSortByEnum string

// Set of constants representing the allowable values for ListUsersSortByEnum
const (
	ListUsersSortByTimecreated ListUsersSortByEnum = "TIMECREATED"
	ListUsersSortByName        ListUsersSortByEnum = "NAME"
)

var mappingListUsersSortByEnum = map[string]ListUsersSortByEnum{
	"TIMECREATED": ListUsersSortByTimecreated,
	"NAME":        ListUsersSortByName,
}

var mappingListUsersSortByEnumLowerCase = map[string]ListUsersSortByEnum{
	"timecreated": ListUsersSortByTimecreated,
	"name":        ListUsersSortByName,
}

// GetListUsersSortByEnumValues Enumerates the set of values for ListUsersSortByEnum
func GetListUsersSortByEnumValues() []ListUsersSortByEnum {
	values := make([]ListUsersSortByEnum, 0)
	for _, v := range mappingListUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsersSortByEnumStringValues Enumerates the set of values in String for ListUsersSortByEnum
func GetListUsersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsersSortByEnum(val string) (ListUsersSortByEnum, bool) {
	enum, ok := mappingListUsersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUsersSortOrderEnum Enum with underlying type: string
type ListUsersSortOrderEnum string

// Set of constants representing the allowable values for ListUsersSortOrderEnum
const (
	ListUsersSortOrderAsc  ListUsersSortOrderEnum = "ASC"
	ListUsersSortOrderDesc ListUsersSortOrderEnum = "DESC"
)

var mappingListUsersSortOrderEnum = map[string]ListUsersSortOrderEnum{
	"ASC":  ListUsersSortOrderAsc,
	"DESC": ListUsersSortOrderDesc,
}

var mappingListUsersSortOrderEnumLowerCase = map[string]ListUsersSortOrderEnum{
	"asc":  ListUsersSortOrderAsc,
	"desc": ListUsersSortOrderDesc,
}

// GetListUsersSortOrderEnumValues Enumerates the set of values for ListUsersSortOrderEnum
func GetListUsersSortOrderEnumValues() []ListUsersSortOrderEnum {
	values := make([]ListUsersSortOrderEnum, 0)
	for _, v := range mappingListUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsersSortOrderEnumStringValues Enumerates the set of values in String for ListUsersSortOrderEnum
func GetListUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsersSortOrderEnum(val string) (ListUsersSortOrderEnum, bool) {
	enum, ok := mappingListUsersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
