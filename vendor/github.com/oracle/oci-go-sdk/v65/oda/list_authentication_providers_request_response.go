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

// ListAuthenticationProvidersRequest wrapper for the ListAuthenticationProviders operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListAuthenticationProviders.go.html to see an example of how to use ListAuthenticationProvidersRequest.
type ListAuthenticationProvidersRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Authentication Provider identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// List only Authentication Providers for this Identity Provider.
	IdentityProvider ListAuthenticationProvidersIdentityProviderEnum `mandatory:"false" contributesTo:"query" name:"identityProvider" omitEmpty:"true"`

	// List only the information for Authentication Providers with this name. Authentication Provider names are unique and may not change.
	// Example: `MyProvider`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// List only the resources that are in this lifecycle state.
	LifecycleState ListAuthenticationProvidersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListAuthenticationProvidersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `timeCreated`.
	// The default sort order for `timeCreated` and `timeUpdated` is descending.
	// For all other sort fields the default sort order is ascending.
	SortBy ListAuthenticationProvidersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuthenticationProvidersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuthenticationProvidersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuthenticationProvidersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuthenticationProvidersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuthenticationProvidersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuthenticationProvidersIdentityProviderEnum(string(request.IdentityProvider)); !ok && request.IdentityProvider != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdentityProvider: %s. Supported values are: %s.", request.IdentityProvider, strings.Join(GetListAuthenticationProvidersIdentityProviderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuthenticationProvidersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAuthenticationProvidersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuthenticationProvidersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAuthenticationProvidersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuthenticationProvidersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAuthenticationProvidersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuthenticationProvidersResponse wrapper for the ListAuthenticationProviders operation
type ListAuthenticationProvidersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuthenticationProviderCollection instances
	AuthenticationProviderCollection `presentIn:"body"`

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

func (response ListAuthenticationProvidersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuthenticationProvidersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuthenticationProvidersIdentityProviderEnum Enum with underlying type: string
type ListAuthenticationProvidersIdentityProviderEnum string

// Set of constants representing the allowable values for ListAuthenticationProvidersIdentityProviderEnum
const (
	ListAuthenticationProvidersIdentityProviderGeneric   ListAuthenticationProvidersIdentityProviderEnum = "GENERIC"
	ListAuthenticationProvidersIdentityProviderOam       ListAuthenticationProvidersIdentityProviderEnum = "OAM"
	ListAuthenticationProvidersIdentityProviderGoogle    ListAuthenticationProvidersIdentityProviderEnum = "GOOGLE"
	ListAuthenticationProvidersIdentityProviderMicrosoft ListAuthenticationProvidersIdentityProviderEnum = "MICROSOFT"
)

var mappingListAuthenticationProvidersIdentityProviderEnum = map[string]ListAuthenticationProvidersIdentityProviderEnum{
	"GENERIC":   ListAuthenticationProvidersIdentityProviderGeneric,
	"OAM":       ListAuthenticationProvidersIdentityProviderOam,
	"GOOGLE":    ListAuthenticationProvidersIdentityProviderGoogle,
	"MICROSOFT": ListAuthenticationProvidersIdentityProviderMicrosoft,
}

var mappingListAuthenticationProvidersIdentityProviderEnumLowerCase = map[string]ListAuthenticationProvidersIdentityProviderEnum{
	"generic":   ListAuthenticationProvidersIdentityProviderGeneric,
	"oam":       ListAuthenticationProvidersIdentityProviderOam,
	"google":    ListAuthenticationProvidersIdentityProviderGoogle,
	"microsoft": ListAuthenticationProvidersIdentityProviderMicrosoft,
}

// GetListAuthenticationProvidersIdentityProviderEnumValues Enumerates the set of values for ListAuthenticationProvidersIdentityProviderEnum
func GetListAuthenticationProvidersIdentityProviderEnumValues() []ListAuthenticationProvidersIdentityProviderEnum {
	values := make([]ListAuthenticationProvidersIdentityProviderEnum, 0)
	for _, v := range mappingListAuthenticationProvidersIdentityProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuthenticationProvidersIdentityProviderEnumStringValues Enumerates the set of values in String for ListAuthenticationProvidersIdentityProviderEnum
func GetListAuthenticationProvidersIdentityProviderEnumStringValues() []string {
	return []string{
		"GENERIC",
		"OAM",
		"GOOGLE",
		"MICROSOFT",
	}
}

// GetMappingListAuthenticationProvidersIdentityProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuthenticationProvidersIdentityProviderEnum(val string) (ListAuthenticationProvidersIdentityProviderEnum, bool) {
	enum, ok := mappingListAuthenticationProvidersIdentityProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuthenticationProvidersLifecycleStateEnum Enum with underlying type: string
type ListAuthenticationProvidersLifecycleStateEnum string

// Set of constants representing the allowable values for ListAuthenticationProvidersLifecycleStateEnum
const (
	ListAuthenticationProvidersLifecycleStateCreating ListAuthenticationProvidersLifecycleStateEnum = "CREATING"
	ListAuthenticationProvidersLifecycleStateUpdating ListAuthenticationProvidersLifecycleStateEnum = "UPDATING"
	ListAuthenticationProvidersLifecycleStateActive   ListAuthenticationProvidersLifecycleStateEnum = "ACTIVE"
	ListAuthenticationProvidersLifecycleStateInactive ListAuthenticationProvidersLifecycleStateEnum = "INACTIVE"
	ListAuthenticationProvidersLifecycleStateDeleting ListAuthenticationProvidersLifecycleStateEnum = "DELETING"
	ListAuthenticationProvidersLifecycleStateDeleted  ListAuthenticationProvidersLifecycleStateEnum = "DELETED"
	ListAuthenticationProvidersLifecycleStateFailed   ListAuthenticationProvidersLifecycleStateEnum = "FAILED"
)

var mappingListAuthenticationProvidersLifecycleStateEnum = map[string]ListAuthenticationProvidersLifecycleStateEnum{
	"CREATING": ListAuthenticationProvidersLifecycleStateCreating,
	"UPDATING": ListAuthenticationProvidersLifecycleStateUpdating,
	"ACTIVE":   ListAuthenticationProvidersLifecycleStateActive,
	"INACTIVE": ListAuthenticationProvidersLifecycleStateInactive,
	"DELETING": ListAuthenticationProvidersLifecycleStateDeleting,
	"DELETED":  ListAuthenticationProvidersLifecycleStateDeleted,
	"FAILED":   ListAuthenticationProvidersLifecycleStateFailed,
}

var mappingListAuthenticationProvidersLifecycleStateEnumLowerCase = map[string]ListAuthenticationProvidersLifecycleStateEnum{
	"creating": ListAuthenticationProvidersLifecycleStateCreating,
	"updating": ListAuthenticationProvidersLifecycleStateUpdating,
	"active":   ListAuthenticationProvidersLifecycleStateActive,
	"inactive": ListAuthenticationProvidersLifecycleStateInactive,
	"deleting": ListAuthenticationProvidersLifecycleStateDeleting,
	"deleted":  ListAuthenticationProvidersLifecycleStateDeleted,
	"failed":   ListAuthenticationProvidersLifecycleStateFailed,
}

// GetListAuthenticationProvidersLifecycleStateEnumValues Enumerates the set of values for ListAuthenticationProvidersLifecycleStateEnum
func GetListAuthenticationProvidersLifecycleStateEnumValues() []ListAuthenticationProvidersLifecycleStateEnum {
	values := make([]ListAuthenticationProvidersLifecycleStateEnum, 0)
	for _, v := range mappingListAuthenticationProvidersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuthenticationProvidersLifecycleStateEnumStringValues Enumerates the set of values in String for ListAuthenticationProvidersLifecycleStateEnum
func GetListAuthenticationProvidersLifecycleStateEnumStringValues() []string {
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

// GetMappingListAuthenticationProvidersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuthenticationProvidersLifecycleStateEnum(val string) (ListAuthenticationProvidersLifecycleStateEnum, bool) {
	enum, ok := mappingListAuthenticationProvidersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuthenticationProvidersSortOrderEnum Enum with underlying type: string
type ListAuthenticationProvidersSortOrderEnum string

// Set of constants representing the allowable values for ListAuthenticationProvidersSortOrderEnum
const (
	ListAuthenticationProvidersSortOrderAsc  ListAuthenticationProvidersSortOrderEnum = "ASC"
	ListAuthenticationProvidersSortOrderDesc ListAuthenticationProvidersSortOrderEnum = "DESC"
)

var mappingListAuthenticationProvidersSortOrderEnum = map[string]ListAuthenticationProvidersSortOrderEnum{
	"ASC":  ListAuthenticationProvidersSortOrderAsc,
	"DESC": ListAuthenticationProvidersSortOrderDesc,
}

var mappingListAuthenticationProvidersSortOrderEnumLowerCase = map[string]ListAuthenticationProvidersSortOrderEnum{
	"asc":  ListAuthenticationProvidersSortOrderAsc,
	"desc": ListAuthenticationProvidersSortOrderDesc,
}

// GetListAuthenticationProvidersSortOrderEnumValues Enumerates the set of values for ListAuthenticationProvidersSortOrderEnum
func GetListAuthenticationProvidersSortOrderEnumValues() []ListAuthenticationProvidersSortOrderEnum {
	values := make([]ListAuthenticationProvidersSortOrderEnum, 0)
	for _, v := range mappingListAuthenticationProvidersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuthenticationProvidersSortOrderEnumStringValues Enumerates the set of values in String for ListAuthenticationProvidersSortOrderEnum
func GetListAuthenticationProvidersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAuthenticationProvidersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuthenticationProvidersSortOrderEnum(val string) (ListAuthenticationProvidersSortOrderEnum, bool) {
	enum, ok := mappingListAuthenticationProvidersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuthenticationProvidersSortByEnum Enum with underlying type: string
type ListAuthenticationProvidersSortByEnum string

// Set of constants representing the allowable values for ListAuthenticationProvidersSortByEnum
const (
	ListAuthenticationProvidersSortByTimecreated      ListAuthenticationProvidersSortByEnum = "timeCreated"
	ListAuthenticationProvidersSortByTimeupdated      ListAuthenticationProvidersSortByEnum = "timeUpdated"
	ListAuthenticationProvidersSortByName             ListAuthenticationProvidersSortByEnum = "name"
	ListAuthenticationProvidersSortByIdentityprovider ListAuthenticationProvidersSortByEnum = "identityProvider"
)

var mappingListAuthenticationProvidersSortByEnum = map[string]ListAuthenticationProvidersSortByEnum{
	"timeCreated":      ListAuthenticationProvidersSortByTimecreated,
	"timeUpdated":      ListAuthenticationProvidersSortByTimeupdated,
	"name":             ListAuthenticationProvidersSortByName,
	"identityProvider": ListAuthenticationProvidersSortByIdentityprovider,
}

var mappingListAuthenticationProvidersSortByEnumLowerCase = map[string]ListAuthenticationProvidersSortByEnum{
	"timecreated":      ListAuthenticationProvidersSortByTimecreated,
	"timeupdated":      ListAuthenticationProvidersSortByTimeupdated,
	"name":             ListAuthenticationProvidersSortByName,
	"identityprovider": ListAuthenticationProvidersSortByIdentityprovider,
}

// GetListAuthenticationProvidersSortByEnumValues Enumerates the set of values for ListAuthenticationProvidersSortByEnum
func GetListAuthenticationProvidersSortByEnumValues() []ListAuthenticationProvidersSortByEnum {
	values := make([]ListAuthenticationProvidersSortByEnum, 0)
	for _, v := range mappingListAuthenticationProvidersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuthenticationProvidersSortByEnumStringValues Enumerates the set of values in String for ListAuthenticationProvidersSortByEnum
func GetListAuthenticationProvidersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
		"identityProvider",
	}
}

// GetMappingListAuthenticationProvidersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuthenticationProvidersSortByEnum(val string) (ListAuthenticationProvidersSortByEnum, bool) {
	enum, ok := mappingListAuthenticationProvidersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
