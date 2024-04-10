// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListZprTagsRequest wrapper for the ListZprTags operation
type ListZprTagsRequest struct {

	// The OCID of the zpr tag namespace.
	ZprTagNamespaceId *string `mandatory:"true" contributesTo:"path" name:"zprTagNamespaceId"`

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListZprTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListZprTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState ZprTagLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZprTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZprTagsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZprTagsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZprTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListZprTagsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListZprTagsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListZprTagsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprTagsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListZprTagsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingZprTagLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetZprTagLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListZprTagsResponse wrapper for the ListZprTags operation
type ListZprTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ZprTagCollection instances
	ZprTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of Zpr tags. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListZprTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZprTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZprTagsSortOrderEnum Enum with underlying type: string
type ListZprTagsSortOrderEnum string

// Set of constants representing the allowable values for ListZprTagsSortOrderEnum
const (
	ListZprTagsSortOrderAsc  ListZprTagsSortOrderEnum = "ASC"
	ListZprTagsSortOrderDesc ListZprTagsSortOrderEnum = "DESC"
)

var mappingListZprTagsSortOrderEnum = map[string]ListZprTagsSortOrderEnum{
	"ASC":  ListZprTagsSortOrderAsc,
	"DESC": ListZprTagsSortOrderDesc,
}

var mappingListZprTagsSortOrderEnumLowerCase = map[string]ListZprTagsSortOrderEnum{
	"asc":  ListZprTagsSortOrderAsc,
	"desc": ListZprTagsSortOrderDesc,
}

// GetListZprTagsSortOrderEnumValues Enumerates the set of values for ListZprTagsSortOrderEnum
func GetListZprTagsSortOrderEnumValues() []ListZprTagsSortOrderEnum {
	values := make([]ListZprTagsSortOrderEnum, 0)
	for _, v := range mappingListZprTagsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprTagsSortOrderEnumStringValues Enumerates the set of values in String for ListZprTagsSortOrderEnum
func GetListZprTagsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListZprTagsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprTagsSortOrderEnum(val string) (ListZprTagsSortOrderEnum, bool) {
	enum, ok := mappingListZprTagsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprTagsSortByEnum Enum with underlying type: string
type ListZprTagsSortByEnum string

// Set of constants representing the allowable values for ListZprTagsSortByEnum
const (
	ListZprTagsSortByTimecreated ListZprTagsSortByEnum = "TIMECREATED"
	ListZprTagsSortByName        ListZprTagsSortByEnum = "NAME"
)

var mappingListZprTagsSortByEnum = map[string]ListZprTagsSortByEnum{
	"TIMECREATED": ListZprTagsSortByTimecreated,
	"NAME":        ListZprTagsSortByName,
}

var mappingListZprTagsSortByEnumLowerCase = map[string]ListZprTagsSortByEnum{
	"timecreated": ListZprTagsSortByTimecreated,
	"name":        ListZprTagsSortByName,
}

// GetListZprTagsSortByEnumValues Enumerates the set of values for ListZprTagsSortByEnum
func GetListZprTagsSortByEnumValues() []ListZprTagsSortByEnum {
	values := make([]ListZprTagsSortByEnum, 0)
	for _, v := range mappingListZprTagsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprTagsSortByEnumStringValues Enumerates the set of values in String for ListZprTagsSortByEnum
func GetListZprTagsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListZprTagsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprTagsSortByEnum(val string) (ListZprTagsSortByEnum, bool) {
	enum, ok := mappingListZprTagsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
