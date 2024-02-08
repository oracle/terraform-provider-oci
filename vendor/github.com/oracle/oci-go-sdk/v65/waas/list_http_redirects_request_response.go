// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListHttpRedirectsRequest wrapper for the ListHttpRedirects operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListHttpRedirects.go.html to see an example of how to use ListHttpRedirectsRequest.
type ListHttpRedirectsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The value of the sorting direction of resources in a paginated 'List' call. If unspecified, defaults to `DESC`.
	SortOrder ListHttpRedirectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the results of the List query.
	SortBy ListHttpRedirectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter redirects using a list of redirect OCIDs.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Filter redirects using a display name.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// Filter redirects using a list of lifecycle states.
	LifecycleState []LifecycleStatesEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter that matches redirects created on or after the specified date and time.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// A filter that matches redirects created before the specified date-time. Default to 1 day before now.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHttpRedirectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHttpRedirectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHttpRedirectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHttpRedirectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHttpRedirectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListHttpRedirectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHttpRedirectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHttpRedirectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHttpRedirectsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingLifecycleStatesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHttpRedirectsResponse wrapper for the ListHttpRedirects operation
type ListHttpRedirectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []HttpRedirectSummary instances
	Items []HttpRedirectSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListHttpRedirectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHttpRedirectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHttpRedirectsSortOrderEnum Enum with underlying type: string
type ListHttpRedirectsSortOrderEnum string

// Set of constants representing the allowable values for ListHttpRedirectsSortOrderEnum
const (
	ListHttpRedirectsSortOrderAsc  ListHttpRedirectsSortOrderEnum = "ASC"
	ListHttpRedirectsSortOrderDesc ListHttpRedirectsSortOrderEnum = "DESC"
)

var mappingListHttpRedirectsSortOrderEnum = map[string]ListHttpRedirectsSortOrderEnum{
	"ASC":  ListHttpRedirectsSortOrderAsc,
	"DESC": ListHttpRedirectsSortOrderDesc,
}

var mappingListHttpRedirectsSortOrderEnumLowerCase = map[string]ListHttpRedirectsSortOrderEnum{
	"asc":  ListHttpRedirectsSortOrderAsc,
	"desc": ListHttpRedirectsSortOrderDesc,
}

// GetListHttpRedirectsSortOrderEnumValues Enumerates the set of values for ListHttpRedirectsSortOrderEnum
func GetListHttpRedirectsSortOrderEnumValues() []ListHttpRedirectsSortOrderEnum {
	values := make([]ListHttpRedirectsSortOrderEnum, 0)
	for _, v := range mappingListHttpRedirectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHttpRedirectsSortOrderEnumStringValues Enumerates the set of values in String for ListHttpRedirectsSortOrderEnum
func GetListHttpRedirectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHttpRedirectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHttpRedirectsSortOrderEnum(val string) (ListHttpRedirectsSortOrderEnum, bool) {
	enum, ok := mappingListHttpRedirectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHttpRedirectsSortByEnum Enum with underlying type: string
type ListHttpRedirectsSortByEnum string

// Set of constants representing the allowable values for ListHttpRedirectsSortByEnum
const (
	ListHttpRedirectsSortById          ListHttpRedirectsSortByEnum = "id"
	ListHttpRedirectsSortByDomain      ListHttpRedirectsSortByEnum = "domain"
	ListHttpRedirectsSortByTarget      ListHttpRedirectsSortByEnum = "target"
	ListHttpRedirectsSortByDisplayname ListHttpRedirectsSortByEnum = "displayName"
)

var mappingListHttpRedirectsSortByEnum = map[string]ListHttpRedirectsSortByEnum{
	"id":          ListHttpRedirectsSortById,
	"domain":      ListHttpRedirectsSortByDomain,
	"target":      ListHttpRedirectsSortByTarget,
	"displayName": ListHttpRedirectsSortByDisplayname,
}

var mappingListHttpRedirectsSortByEnumLowerCase = map[string]ListHttpRedirectsSortByEnum{
	"id":          ListHttpRedirectsSortById,
	"domain":      ListHttpRedirectsSortByDomain,
	"target":      ListHttpRedirectsSortByTarget,
	"displayname": ListHttpRedirectsSortByDisplayname,
}

// GetListHttpRedirectsSortByEnumValues Enumerates the set of values for ListHttpRedirectsSortByEnum
func GetListHttpRedirectsSortByEnumValues() []ListHttpRedirectsSortByEnum {
	values := make([]ListHttpRedirectsSortByEnum, 0)
	for _, v := range mappingListHttpRedirectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHttpRedirectsSortByEnumStringValues Enumerates the set of values in String for ListHttpRedirectsSortByEnum
func GetListHttpRedirectsSortByEnumStringValues() []string {
	return []string{
		"id",
		"domain",
		"target",
		"displayName",
	}
}

// GetMappingListHttpRedirectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHttpRedirectsSortByEnum(val string) (ListHttpRedirectsSortByEnum, bool) {
	enum, ok := mappingListHttpRedirectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
