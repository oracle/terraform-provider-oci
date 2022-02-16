// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListWaasPoliciesRequest wrapper for the ListWaasPolicies operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListWaasPolicies.go.html to see an example of how to use ListWaasPoliciesRequest.
type ListWaasPoliciesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The value by which policies are sorted in a paginated 'List' call.  If unspecified, defaults to `timeCreated`.
	SortBy ListWaasPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The value of the sorting direction of resources in a paginated 'List' call. If unspecified, defaults to `DESC`.
	SortOrder ListWaasPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter policies using a list of policy OCIDs.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Filter policies using a list of display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// Filter policies using a list of lifecycle states.
	LifecycleState []LifecycleStatesEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter that matches policies created on or after the specified date and time.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// A filter that matches policies created before the specified date-time.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWaasPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWaasPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWaasPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWaasPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWaasPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWaasPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWaasPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWaasPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWaasPoliciesSortOrderEnumStringValues(), ",")))
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

// ListWaasPoliciesResponse wrapper for the ListWaasPolicies operation
type ListWaasPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WaasPolicySummary instances
	Items []WaasPolicySummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListWaasPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWaasPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWaasPoliciesSortByEnum Enum with underlying type: string
type ListWaasPoliciesSortByEnum string

// Set of constants representing the allowable values for ListWaasPoliciesSortByEnum
const (
	ListWaasPoliciesSortById          ListWaasPoliciesSortByEnum = "id"
	ListWaasPoliciesSortByDisplayname ListWaasPoliciesSortByEnum = "displayName"
	ListWaasPoliciesSortByTimecreated ListWaasPoliciesSortByEnum = "timeCreated"
)

var mappingListWaasPoliciesSortByEnum = map[string]ListWaasPoliciesSortByEnum{
	"id":          ListWaasPoliciesSortById,
	"displayName": ListWaasPoliciesSortByDisplayname,
	"timeCreated": ListWaasPoliciesSortByTimecreated,
}

// GetListWaasPoliciesSortByEnumValues Enumerates the set of values for ListWaasPoliciesSortByEnum
func GetListWaasPoliciesSortByEnumValues() []ListWaasPoliciesSortByEnum {
	values := make([]ListWaasPoliciesSortByEnum, 0)
	for _, v := range mappingListWaasPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWaasPoliciesSortByEnumStringValues Enumerates the set of values in String for ListWaasPoliciesSortByEnum
func GetListWaasPoliciesSortByEnumStringValues() []string {
	return []string{
		"id",
		"displayName",
		"timeCreated",
	}
}

// GetMappingListWaasPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWaasPoliciesSortByEnum(val string) (ListWaasPoliciesSortByEnum, bool) {
	mappingListWaasPoliciesSortByEnumIgnoreCase := make(map[string]ListWaasPoliciesSortByEnum)
	for k, v := range mappingListWaasPoliciesSortByEnum {
		mappingListWaasPoliciesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWaasPoliciesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListWaasPoliciesSortOrderEnum Enum with underlying type: string
type ListWaasPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListWaasPoliciesSortOrderEnum
const (
	ListWaasPoliciesSortOrderAsc  ListWaasPoliciesSortOrderEnum = "ASC"
	ListWaasPoliciesSortOrderDesc ListWaasPoliciesSortOrderEnum = "DESC"
)

var mappingListWaasPoliciesSortOrderEnum = map[string]ListWaasPoliciesSortOrderEnum{
	"ASC":  ListWaasPoliciesSortOrderAsc,
	"DESC": ListWaasPoliciesSortOrderDesc,
}

// GetListWaasPoliciesSortOrderEnumValues Enumerates the set of values for ListWaasPoliciesSortOrderEnum
func GetListWaasPoliciesSortOrderEnumValues() []ListWaasPoliciesSortOrderEnum {
	values := make([]ListWaasPoliciesSortOrderEnum, 0)
	for _, v := range mappingListWaasPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWaasPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListWaasPoliciesSortOrderEnum
func GetListWaasPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWaasPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWaasPoliciesSortOrderEnum(val string) (ListWaasPoliciesSortOrderEnum, bool) {
	mappingListWaasPoliciesSortOrderEnumIgnoreCase := make(map[string]ListWaasPoliciesSortOrderEnum)
	for k, v := range mappingListWaasPoliciesSortOrderEnum {
		mappingListWaasPoliciesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWaasPoliciesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
