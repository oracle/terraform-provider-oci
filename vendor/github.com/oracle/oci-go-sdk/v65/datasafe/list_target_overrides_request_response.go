// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTargetOverridesRequest wrapper for the ListTargetOverrides operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetOverrides.go.html to see an example of how to use ListTargetOverridesRequest.
type ListTargetOverridesRequest struct {

	// The OCID of the audit.
	AuditProfileId *string `mandatory:"true" contributesTo:"path" name:"auditProfileId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListTargetOverridesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTargetOverridesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetOverridesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetOverridesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetOverridesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetOverridesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetOverridesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetOverridesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetOverridesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetOverridesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetOverridesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetOverridesResponse wrapper for the ListTargetOverrides operation
type ListTargetOverridesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetOverrideCollection instances
	TargetOverrideCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTargetOverridesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetOverridesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetOverridesSortByEnum Enum with underlying type: string
type ListTargetOverridesSortByEnum string

// Set of constants representing the allowable values for ListTargetOverridesSortByEnum
const (
	ListTargetOverridesSortByTimecreated ListTargetOverridesSortByEnum = "TIMECREATED"
	ListTargetOverridesSortByDisplayname ListTargetOverridesSortByEnum = "DISPLAYNAME"
)

var mappingListTargetOverridesSortByEnum = map[string]ListTargetOverridesSortByEnum{
	"TIMECREATED": ListTargetOverridesSortByTimecreated,
	"DISPLAYNAME": ListTargetOverridesSortByDisplayname,
}

var mappingListTargetOverridesSortByEnumLowerCase = map[string]ListTargetOverridesSortByEnum{
	"timecreated": ListTargetOverridesSortByTimecreated,
	"displayname": ListTargetOverridesSortByDisplayname,
}

// GetListTargetOverridesSortByEnumValues Enumerates the set of values for ListTargetOverridesSortByEnum
func GetListTargetOverridesSortByEnumValues() []ListTargetOverridesSortByEnum {
	values := make([]ListTargetOverridesSortByEnum, 0)
	for _, v := range mappingListTargetOverridesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetOverridesSortByEnumStringValues Enumerates the set of values in String for ListTargetOverridesSortByEnum
func GetListTargetOverridesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListTargetOverridesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetOverridesSortByEnum(val string) (ListTargetOverridesSortByEnum, bool) {
	enum, ok := mappingListTargetOverridesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetOverridesSortOrderEnum Enum with underlying type: string
type ListTargetOverridesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetOverridesSortOrderEnum
const (
	ListTargetOverridesSortOrderAsc  ListTargetOverridesSortOrderEnum = "ASC"
	ListTargetOverridesSortOrderDesc ListTargetOverridesSortOrderEnum = "DESC"
)

var mappingListTargetOverridesSortOrderEnum = map[string]ListTargetOverridesSortOrderEnum{
	"ASC":  ListTargetOverridesSortOrderAsc,
	"DESC": ListTargetOverridesSortOrderDesc,
}

var mappingListTargetOverridesSortOrderEnumLowerCase = map[string]ListTargetOverridesSortOrderEnum{
	"asc":  ListTargetOverridesSortOrderAsc,
	"desc": ListTargetOverridesSortOrderDesc,
}

// GetListTargetOverridesSortOrderEnumValues Enumerates the set of values for ListTargetOverridesSortOrderEnum
func GetListTargetOverridesSortOrderEnumValues() []ListTargetOverridesSortOrderEnum {
	values := make([]ListTargetOverridesSortOrderEnum, 0)
	for _, v := range mappingListTargetOverridesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetOverridesSortOrderEnumStringValues Enumerates the set of values in String for ListTargetOverridesSortOrderEnum
func GetListTargetOverridesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetOverridesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetOverridesSortOrderEnum(val string) (ListTargetOverridesSortOrderEnum, bool) {
	enum, ok := mappingListTargetOverridesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
