// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDynamicSetsRequest wrapper for the ListDynamicSets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListDynamicSets.go.html to see an example of how to use ListDynamicSetsRequest.
type ListDynamicSetsRequest struct {

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dynamic set. This filter returns resources associated with this dynamic set.
	DynamicSetId *string `mandatory:"false" contributesTo:"query" name:"dynamicSetId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDynamicSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDynamicSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDynamicSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDynamicSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDynamicSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDynamicSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDynamicSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDynamicSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDynamicSetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDynamicSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDynamicSetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDynamicSetsResponse wrapper for the ListDynamicSets operation
type ListDynamicSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DynamicSetCollection instances
	DynamicSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items in the result. Used for pagination of a list of items.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListDynamicSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDynamicSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDynamicSetsSortOrderEnum Enum with underlying type: string
type ListDynamicSetsSortOrderEnum string

// Set of constants representing the allowable values for ListDynamicSetsSortOrderEnum
const (
	ListDynamicSetsSortOrderAsc  ListDynamicSetsSortOrderEnum = "ASC"
	ListDynamicSetsSortOrderDesc ListDynamicSetsSortOrderEnum = "DESC"
)

var mappingListDynamicSetsSortOrderEnum = map[string]ListDynamicSetsSortOrderEnum{
	"ASC":  ListDynamicSetsSortOrderAsc,
	"DESC": ListDynamicSetsSortOrderDesc,
}

var mappingListDynamicSetsSortOrderEnumLowerCase = map[string]ListDynamicSetsSortOrderEnum{
	"asc":  ListDynamicSetsSortOrderAsc,
	"desc": ListDynamicSetsSortOrderDesc,
}

// GetListDynamicSetsSortOrderEnumValues Enumerates the set of values for ListDynamicSetsSortOrderEnum
func GetListDynamicSetsSortOrderEnumValues() []ListDynamicSetsSortOrderEnum {
	values := make([]ListDynamicSetsSortOrderEnum, 0)
	for _, v := range mappingListDynamicSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDynamicSetsSortOrderEnumStringValues Enumerates the set of values in String for ListDynamicSetsSortOrderEnum
func GetListDynamicSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDynamicSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDynamicSetsSortOrderEnum(val string) (ListDynamicSetsSortOrderEnum, bool) {
	enum, ok := mappingListDynamicSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDynamicSetsSortByEnum Enum with underlying type: string
type ListDynamicSetsSortByEnum string

// Set of constants representing the allowable values for ListDynamicSetsSortByEnum
const (
	ListDynamicSetsSortByTimecreated ListDynamicSetsSortByEnum = "timeCreated"
	ListDynamicSetsSortByDisplayname ListDynamicSetsSortByEnum = "displayName"
)

var mappingListDynamicSetsSortByEnum = map[string]ListDynamicSetsSortByEnum{
	"timeCreated": ListDynamicSetsSortByTimecreated,
	"displayName": ListDynamicSetsSortByDisplayname,
}

var mappingListDynamicSetsSortByEnumLowerCase = map[string]ListDynamicSetsSortByEnum{
	"timecreated": ListDynamicSetsSortByTimecreated,
	"displayname": ListDynamicSetsSortByDisplayname,
}

// GetListDynamicSetsSortByEnumValues Enumerates the set of values for ListDynamicSetsSortByEnum
func GetListDynamicSetsSortByEnumValues() []ListDynamicSetsSortByEnum {
	values := make([]ListDynamicSetsSortByEnum, 0)
	for _, v := range mappingListDynamicSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDynamicSetsSortByEnumStringValues Enumerates the set of values in String for ListDynamicSetsSortByEnum
func GetListDynamicSetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDynamicSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDynamicSetsSortByEnum(val string) (ListDynamicSetsSortByEnum, bool) {
	enum, ok := mappingListDynamicSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
