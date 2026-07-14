// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTargetAlertPolicyUnassociatedMembersRequest wrapper for the ListTargetAlertPolicyUnassociatedMembers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetAlertPolicyUnassociatedMembers.go.html to see an example of how to use ListTargetAlertPolicyUnassociatedMembersRequest.
type ListTargetAlertPolicyUnassociatedMembersRequest struct {

	// The OCID of the target-alert policy association.
	TargetAlertPolicyAssociationId *string `mandatory:"true" contributesTo:"path" name:"targetAlertPolicyAssociationId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort parameter may be provided.
	SortBy ListTargetAlertPolicyUnassociatedMembersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTargetAlertPolicyUnassociatedMembersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetAlertPolicyUnassociatedMembersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetAlertPolicyUnassociatedMembersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetAlertPolicyUnassociatedMembersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetAlertPolicyUnassociatedMembersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetAlertPolicyUnassociatedMembersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetAlertPolicyUnassociatedMembersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetAlertPolicyUnassociatedMembersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetAlertPolicyUnassociatedMembersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetAlertPolicyUnassociatedMembersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetAlertPolicyUnassociatedMembersResponse wrapper for the ListTargetAlertPolicyUnassociatedMembers operation
type ListTargetAlertPolicyUnassociatedMembersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetAlertPolicyUnassociatedCollection instances
	TargetAlertPolicyUnassociatedCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTargetAlertPolicyUnassociatedMembersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetAlertPolicyUnassociatedMembersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetAlertPolicyUnassociatedMembersSortByEnum Enum with underlying type: string
type ListTargetAlertPolicyUnassociatedMembersSortByEnum string

// Set of constants representing the allowable values for ListTargetAlertPolicyUnassociatedMembersSortByEnum
const (
	ListTargetAlertPolicyUnassociatedMembersSortByTargetid         ListTargetAlertPolicyUnassociatedMembersSortByEnum = "targetId"
	ListTargetAlertPolicyUnassociatedMembersSortByNotappliedreason ListTargetAlertPolicyUnassociatedMembersSortByEnum = "notAppliedReason"
)

var mappingListTargetAlertPolicyUnassociatedMembersSortByEnum = map[string]ListTargetAlertPolicyUnassociatedMembersSortByEnum{
	"targetId":         ListTargetAlertPolicyUnassociatedMembersSortByTargetid,
	"notAppliedReason": ListTargetAlertPolicyUnassociatedMembersSortByNotappliedreason,
}

var mappingListTargetAlertPolicyUnassociatedMembersSortByEnumLowerCase = map[string]ListTargetAlertPolicyUnassociatedMembersSortByEnum{
	"targetid":         ListTargetAlertPolicyUnassociatedMembersSortByTargetid,
	"notappliedreason": ListTargetAlertPolicyUnassociatedMembersSortByNotappliedreason,
}

// GetListTargetAlertPolicyUnassociatedMembersSortByEnumValues Enumerates the set of values for ListTargetAlertPolicyUnassociatedMembersSortByEnum
func GetListTargetAlertPolicyUnassociatedMembersSortByEnumValues() []ListTargetAlertPolicyUnassociatedMembersSortByEnum {
	values := make([]ListTargetAlertPolicyUnassociatedMembersSortByEnum, 0)
	for _, v := range mappingListTargetAlertPolicyUnassociatedMembersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAlertPolicyUnassociatedMembersSortByEnumStringValues Enumerates the set of values in String for ListTargetAlertPolicyUnassociatedMembersSortByEnum
func GetListTargetAlertPolicyUnassociatedMembersSortByEnumStringValues() []string {
	return []string{
		"targetId",
		"notAppliedReason",
	}
}

// GetMappingListTargetAlertPolicyUnassociatedMembersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAlertPolicyUnassociatedMembersSortByEnum(val string) (ListTargetAlertPolicyUnassociatedMembersSortByEnum, bool) {
	enum, ok := mappingListTargetAlertPolicyUnassociatedMembersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetAlertPolicyUnassociatedMembersSortOrderEnum Enum with underlying type: string
type ListTargetAlertPolicyUnassociatedMembersSortOrderEnum string

// Set of constants representing the allowable values for ListTargetAlertPolicyUnassociatedMembersSortOrderEnum
const (
	ListTargetAlertPolicyUnassociatedMembersSortOrderAsc  ListTargetAlertPolicyUnassociatedMembersSortOrderEnum = "ASC"
	ListTargetAlertPolicyUnassociatedMembersSortOrderDesc ListTargetAlertPolicyUnassociatedMembersSortOrderEnum = "DESC"
)

var mappingListTargetAlertPolicyUnassociatedMembersSortOrderEnum = map[string]ListTargetAlertPolicyUnassociatedMembersSortOrderEnum{
	"ASC":  ListTargetAlertPolicyUnassociatedMembersSortOrderAsc,
	"DESC": ListTargetAlertPolicyUnassociatedMembersSortOrderDesc,
}

var mappingListTargetAlertPolicyUnassociatedMembersSortOrderEnumLowerCase = map[string]ListTargetAlertPolicyUnassociatedMembersSortOrderEnum{
	"asc":  ListTargetAlertPolicyUnassociatedMembersSortOrderAsc,
	"desc": ListTargetAlertPolicyUnassociatedMembersSortOrderDesc,
}

// GetListTargetAlertPolicyUnassociatedMembersSortOrderEnumValues Enumerates the set of values for ListTargetAlertPolicyUnassociatedMembersSortOrderEnum
func GetListTargetAlertPolicyUnassociatedMembersSortOrderEnumValues() []ListTargetAlertPolicyUnassociatedMembersSortOrderEnum {
	values := make([]ListTargetAlertPolicyUnassociatedMembersSortOrderEnum, 0)
	for _, v := range mappingListTargetAlertPolicyUnassociatedMembersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAlertPolicyUnassociatedMembersSortOrderEnumStringValues Enumerates the set of values in String for ListTargetAlertPolicyUnassociatedMembersSortOrderEnum
func GetListTargetAlertPolicyUnassociatedMembersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetAlertPolicyUnassociatedMembersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAlertPolicyUnassociatedMembersSortOrderEnum(val string) (ListTargetAlertPolicyUnassociatedMembersSortOrderEnum, bool) {
	enum, ok := mappingListTargetAlertPolicyUnassociatedMembersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
