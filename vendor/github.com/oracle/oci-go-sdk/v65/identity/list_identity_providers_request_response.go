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

// ListIdentityProvidersRequest wrapper for the ListIdentityProviders operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIdentityProviders.go.html to see an example of how to use ListIdentityProvidersRequest.
type ListIdentityProvidersRequest struct {

	// The protocol used for federation.
	Protocol ListIdentityProvidersProtocolEnum `mandatory:"true" contributesTo:"query" name:"protocol" omitEmpty:"true"`

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListIdentityProvidersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListIdentityProvidersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState IdentityProviderLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIdentityProvidersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIdentityProvidersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIdentityProvidersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIdentityProvidersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIdentityProvidersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIdentityProvidersProtocolEnum(string(request.Protocol)); !ok && request.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", request.Protocol, strings.Join(GetListIdentityProvidersProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIdentityProvidersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIdentityProvidersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIdentityProvidersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIdentityProvidersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIdentityProviderLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetIdentityProviderLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIdentityProvidersResponse wrapper for the ListIdentityProviders operation
type ListIdentityProvidersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IdentityProvider instances
	Items []IdentityProvider `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIdentityProvidersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIdentityProvidersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIdentityProvidersProtocolEnum Enum with underlying type: string
type ListIdentityProvidersProtocolEnum string

// Set of constants representing the allowable values for ListIdentityProvidersProtocolEnum
const (
	ListIdentityProvidersProtocolSaml2 ListIdentityProvidersProtocolEnum = "SAML2"
)

var mappingListIdentityProvidersProtocolEnum = map[string]ListIdentityProvidersProtocolEnum{
	"SAML2": ListIdentityProvidersProtocolSaml2,
}

var mappingListIdentityProvidersProtocolEnumLowerCase = map[string]ListIdentityProvidersProtocolEnum{
	"saml2": ListIdentityProvidersProtocolSaml2,
}

// GetListIdentityProvidersProtocolEnumValues Enumerates the set of values for ListIdentityProvidersProtocolEnum
func GetListIdentityProvidersProtocolEnumValues() []ListIdentityProvidersProtocolEnum {
	values := make([]ListIdentityProvidersProtocolEnum, 0)
	for _, v := range mappingListIdentityProvidersProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetListIdentityProvidersProtocolEnumStringValues Enumerates the set of values in String for ListIdentityProvidersProtocolEnum
func GetListIdentityProvidersProtocolEnumStringValues() []string {
	return []string{
		"SAML2",
	}
}

// GetMappingListIdentityProvidersProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIdentityProvidersProtocolEnum(val string) (ListIdentityProvidersProtocolEnum, bool) {
	enum, ok := mappingListIdentityProvidersProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIdentityProvidersSortByEnum Enum with underlying type: string
type ListIdentityProvidersSortByEnum string

// Set of constants representing the allowable values for ListIdentityProvidersSortByEnum
const (
	ListIdentityProvidersSortByTimecreated ListIdentityProvidersSortByEnum = "TIMECREATED"
	ListIdentityProvidersSortByName        ListIdentityProvidersSortByEnum = "NAME"
)

var mappingListIdentityProvidersSortByEnum = map[string]ListIdentityProvidersSortByEnum{
	"TIMECREATED": ListIdentityProvidersSortByTimecreated,
	"NAME":        ListIdentityProvidersSortByName,
}

var mappingListIdentityProvidersSortByEnumLowerCase = map[string]ListIdentityProvidersSortByEnum{
	"timecreated": ListIdentityProvidersSortByTimecreated,
	"name":        ListIdentityProvidersSortByName,
}

// GetListIdentityProvidersSortByEnumValues Enumerates the set of values for ListIdentityProvidersSortByEnum
func GetListIdentityProvidersSortByEnumValues() []ListIdentityProvidersSortByEnum {
	values := make([]ListIdentityProvidersSortByEnum, 0)
	for _, v := range mappingListIdentityProvidersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIdentityProvidersSortByEnumStringValues Enumerates the set of values in String for ListIdentityProvidersSortByEnum
func GetListIdentityProvidersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListIdentityProvidersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIdentityProvidersSortByEnum(val string) (ListIdentityProvidersSortByEnum, bool) {
	enum, ok := mappingListIdentityProvidersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIdentityProvidersSortOrderEnum Enum with underlying type: string
type ListIdentityProvidersSortOrderEnum string

// Set of constants representing the allowable values for ListIdentityProvidersSortOrderEnum
const (
	ListIdentityProvidersSortOrderAsc  ListIdentityProvidersSortOrderEnum = "ASC"
	ListIdentityProvidersSortOrderDesc ListIdentityProvidersSortOrderEnum = "DESC"
)

var mappingListIdentityProvidersSortOrderEnum = map[string]ListIdentityProvidersSortOrderEnum{
	"ASC":  ListIdentityProvidersSortOrderAsc,
	"DESC": ListIdentityProvidersSortOrderDesc,
}

var mappingListIdentityProvidersSortOrderEnumLowerCase = map[string]ListIdentityProvidersSortOrderEnum{
	"asc":  ListIdentityProvidersSortOrderAsc,
	"desc": ListIdentityProvidersSortOrderDesc,
}

// GetListIdentityProvidersSortOrderEnumValues Enumerates the set of values for ListIdentityProvidersSortOrderEnum
func GetListIdentityProvidersSortOrderEnumValues() []ListIdentityProvidersSortOrderEnum {
	values := make([]ListIdentityProvidersSortOrderEnum, 0)
	for _, v := range mappingListIdentityProvidersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIdentityProvidersSortOrderEnumStringValues Enumerates the set of values in String for ListIdentityProvidersSortOrderEnum
func GetListIdentityProvidersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIdentityProvidersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIdentityProvidersSortOrderEnum(val string) (ListIdentityProvidersSortOrderEnum, bool) {
	enum, ok := mappingListIdentityProvidersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
