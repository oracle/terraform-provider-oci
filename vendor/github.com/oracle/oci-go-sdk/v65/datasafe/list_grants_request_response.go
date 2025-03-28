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

// ListGrantsRequest wrapper for the ListGrants operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListGrants.go.html to see an example of how to use ListGrantsRequest.
type ListGrantsRequest struct {

	// The OCID of the user assessment.
	UserAssessmentId *string `mandatory:"true" contributesTo:"path" name:"userAssessmentId"`

	// The unique user key. This is a system-generated identifier. ListUsers gets the user key for a user.
	UserKey *string `mandatory:"true" contributesTo:"path" name:"userKey"`

	// A filter to return only items that match the specified user grant key.
	GrantKey *string `mandatory:"false" contributesTo:"query" name:"grantKey"`

	// A filter to return only items that match the specified user grant name.
	GrantName *string `mandatory:"false" contributesTo:"query" name:"grantName"`

	// A filter to return only items that match the specified privilege grant type.
	PrivilegeType *string `mandatory:"false" contributesTo:"query" name:"privilegeType"`

	// A filter to return only items that match the specified user privilege category.
	PrivilegeCategory *string `mandatory:"false" contributesTo:"query" name:"privilegeCategory"`

	// A filter to return only items that match the specified user grant depth level.
	DepthLevel *int `mandatory:"false" contributesTo:"query" name:"depthLevel"`

	// A filter to return only items that are at a level greater than or equal to the specified user grant depth level.
	DepthLevelGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"depthLevelGreaterThanOrEqualTo"`

	// A filter to return only items that are at a level less than the specified user grant depth level.
	DepthLevelLessThan *int `mandatory:"false" contributesTo:"query" name:"depthLevelLessThan"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListGrantsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for grantName is ascending.
	SortBy ListGrantsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGrantsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGrantsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGrantsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGrantsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGrantsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGrantsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGrantsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGrantsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGrantsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGrantsResponse wrapper for the ListGrants operation
type ListGrantsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []GrantSummary instances
	Items []GrantSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListGrantsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGrantsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGrantsSortOrderEnum Enum with underlying type: string
type ListGrantsSortOrderEnum string

// Set of constants representing the allowable values for ListGrantsSortOrderEnum
const (
	ListGrantsSortOrderAsc  ListGrantsSortOrderEnum = "ASC"
	ListGrantsSortOrderDesc ListGrantsSortOrderEnum = "DESC"
)

var mappingListGrantsSortOrderEnum = map[string]ListGrantsSortOrderEnum{
	"ASC":  ListGrantsSortOrderAsc,
	"DESC": ListGrantsSortOrderDesc,
}

var mappingListGrantsSortOrderEnumLowerCase = map[string]ListGrantsSortOrderEnum{
	"asc":  ListGrantsSortOrderAsc,
	"desc": ListGrantsSortOrderDesc,
}

// GetListGrantsSortOrderEnumValues Enumerates the set of values for ListGrantsSortOrderEnum
func GetListGrantsSortOrderEnumValues() []ListGrantsSortOrderEnum {
	values := make([]ListGrantsSortOrderEnum, 0)
	for _, v := range mappingListGrantsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGrantsSortOrderEnumStringValues Enumerates the set of values in String for ListGrantsSortOrderEnum
func GetListGrantsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGrantsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGrantsSortOrderEnum(val string) (ListGrantsSortOrderEnum, bool) {
	enum, ok := mappingListGrantsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGrantsSortByEnum Enum with underlying type: string
type ListGrantsSortByEnum string

// Set of constants representing the allowable values for ListGrantsSortByEnum
const (
	ListGrantsSortByGrantname         ListGrantsSortByEnum = "grantName"
	ListGrantsSortByGranttype         ListGrantsSortByEnum = "grantType"
	ListGrantsSortByPrivilegecategory ListGrantsSortByEnum = "privilegeCategory"
	ListGrantsSortByDepthlevel        ListGrantsSortByEnum = "depthLevel"
	ListGrantsSortByKey               ListGrantsSortByEnum = "key"
)

var mappingListGrantsSortByEnum = map[string]ListGrantsSortByEnum{
	"grantName":         ListGrantsSortByGrantname,
	"grantType":         ListGrantsSortByGranttype,
	"privilegeCategory": ListGrantsSortByPrivilegecategory,
	"depthLevel":        ListGrantsSortByDepthlevel,
	"key":               ListGrantsSortByKey,
}

var mappingListGrantsSortByEnumLowerCase = map[string]ListGrantsSortByEnum{
	"grantname":         ListGrantsSortByGrantname,
	"granttype":         ListGrantsSortByGranttype,
	"privilegecategory": ListGrantsSortByPrivilegecategory,
	"depthlevel":        ListGrantsSortByDepthlevel,
	"key":               ListGrantsSortByKey,
}

// GetListGrantsSortByEnumValues Enumerates the set of values for ListGrantsSortByEnum
func GetListGrantsSortByEnumValues() []ListGrantsSortByEnum {
	values := make([]ListGrantsSortByEnum, 0)
	for _, v := range mappingListGrantsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGrantsSortByEnumStringValues Enumerates the set of values in String for ListGrantsSortByEnum
func GetListGrantsSortByEnumStringValues() []string {
	return []string{
		"grantName",
		"grantType",
		"privilegeCategory",
		"depthLevel",
		"key",
	}
}

// GetMappingListGrantsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGrantsSortByEnum(val string) (ListGrantsSortByEnum, bool) {
	enum, ok := mappingListGrantsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
