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

// ListAddressListsRequest wrapper for the ListAddressLists operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListAddressLists.go.html to see an example of how to use ListAddressListsRequest.
type ListAddressListsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `10`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The value by which address lists are sorted in a paginated 'List' call. If unspecified, defaults to `timeCreated`.
	SortBy ListAddressListsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The value of the sorting direction of resources in a paginated 'List' call. If unspecified, defaults to `DESC`.
	SortOrder ListAddressListsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter address lists using a list of address lists OCIDs.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Filter address lists using a list of names.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// Filter address lists using a list of lifecycle states.
	LifecycleState []LifecycleStatesEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter that matches address lists created on or after the specified date-time.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// A filter that matches address lists created before the specified date-time.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAddressListsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAddressListsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAddressListsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAddressListsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAddressListsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAddressListsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAddressListsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAddressListsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAddressListsSortOrderEnumStringValues(), ",")))
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

// ListAddressListsResponse wrapper for the ListAddressLists operation
type ListAddressListsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AddressListSummary instances
	Items []AddressListSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAddressListsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAddressListsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAddressListsSortByEnum Enum with underlying type: string
type ListAddressListsSortByEnum string

// Set of constants representing the allowable values for ListAddressListsSortByEnum
const (
	ListAddressListsSortById          ListAddressListsSortByEnum = "id"
	ListAddressListsSortByName        ListAddressListsSortByEnum = "name"
	ListAddressListsSortByTimecreated ListAddressListsSortByEnum = "timeCreated"
)

var mappingListAddressListsSortByEnum = map[string]ListAddressListsSortByEnum{
	"id":          ListAddressListsSortById,
	"name":        ListAddressListsSortByName,
	"timeCreated": ListAddressListsSortByTimecreated,
}

// GetListAddressListsSortByEnumValues Enumerates the set of values for ListAddressListsSortByEnum
func GetListAddressListsSortByEnumValues() []ListAddressListsSortByEnum {
	values := make([]ListAddressListsSortByEnum, 0)
	for _, v := range mappingListAddressListsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddressListsSortByEnumStringValues Enumerates the set of values in String for ListAddressListsSortByEnum
func GetListAddressListsSortByEnumStringValues() []string {
	return []string{
		"id",
		"name",
		"timeCreated",
	}
}

// GetMappingListAddressListsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddressListsSortByEnum(val string) (ListAddressListsSortByEnum, bool) {
	mappingListAddressListsSortByEnumIgnoreCase := make(map[string]ListAddressListsSortByEnum)
	for k, v := range mappingListAddressListsSortByEnum {
		mappingListAddressListsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAddressListsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAddressListsSortOrderEnum Enum with underlying type: string
type ListAddressListsSortOrderEnum string

// Set of constants representing the allowable values for ListAddressListsSortOrderEnum
const (
	ListAddressListsSortOrderAsc  ListAddressListsSortOrderEnum = "ASC"
	ListAddressListsSortOrderDesc ListAddressListsSortOrderEnum = "DESC"
)

var mappingListAddressListsSortOrderEnum = map[string]ListAddressListsSortOrderEnum{
	"ASC":  ListAddressListsSortOrderAsc,
	"DESC": ListAddressListsSortOrderDesc,
}

// GetListAddressListsSortOrderEnumValues Enumerates the set of values for ListAddressListsSortOrderEnum
func GetListAddressListsSortOrderEnumValues() []ListAddressListsSortOrderEnum {
	values := make([]ListAddressListsSortOrderEnum, 0)
	for _, v := range mappingListAddressListsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddressListsSortOrderEnumStringValues Enumerates the set of values in String for ListAddressListsSortOrderEnum
func GetListAddressListsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAddressListsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddressListsSortOrderEnum(val string) (ListAddressListsSortOrderEnum, bool) {
	mappingListAddressListsSortOrderEnumIgnoreCase := make(map[string]ListAddressListsSortOrderEnum)
	for k, v := range mappingListAddressListsSortOrderEnum {
		mappingListAddressListsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAddressListsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
