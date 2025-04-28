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

// ListTargetAlertPolicyAssociationDeviationsRequest wrapper for the ListTargetAlertPolicyAssociationDeviations operation
type ListTargetAlertPolicyAssociationDeviationsRequest struct {

	// The OCID of the target-alert policy association.
	TargetAlertPolicyAssociationId *string `mandatory:"true" contributesTo:"path" name:"targetAlertPolicyAssociationId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListTargetAlertPolicyAssociationDeviationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTargetAlertPolicyAssociationDeviationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetAlertPolicyAssociationDeviationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetAlertPolicyAssociationDeviationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetAlertPolicyAssociationDeviationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetAlertPolicyAssociationDeviationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetAlertPolicyAssociationDeviationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetAlertPolicyAssociationDeviationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetAlertPolicyAssociationDeviationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetAlertPolicyAssociationDeviationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetAlertPolicyAssociationDeviationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetAlertPolicyAssociationDeviationsResponse wrapper for the ListTargetAlertPolicyAssociationDeviations operation
type ListTargetAlertPolicyAssociationDeviationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetAlertPolicyAssociationDeviationCollection instances
	TargetAlertPolicyAssociationDeviationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTargetAlertPolicyAssociationDeviationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetAlertPolicyAssociationDeviationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetAlertPolicyAssociationDeviationsSortByEnum Enum with underlying type: string
type ListTargetAlertPolicyAssociationDeviationsSortByEnum string

// Set of constants representing the allowable values for ListTargetAlertPolicyAssociationDeviationsSortByEnum
const (
	ListTargetAlertPolicyAssociationDeviationsSortByTimecreated ListTargetAlertPolicyAssociationDeviationsSortByEnum = "TIMECREATED"
	ListTargetAlertPolicyAssociationDeviationsSortByDisplayname ListTargetAlertPolicyAssociationDeviationsSortByEnum = "DISPLAYNAME"
)

var mappingListTargetAlertPolicyAssociationDeviationsSortByEnum = map[string]ListTargetAlertPolicyAssociationDeviationsSortByEnum{
	"TIMECREATED": ListTargetAlertPolicyAssociationDeviationsSortByTimecreated,
	"DISPLAYNAME": ListTargetAlertPolicyAssociationDeviationsSortByDisplayname,
}

var mappingListTargetAlertPolicyAssociationDeviationsSortByEnumLowerCase = map[string]ListTargetAlertPolicyAssociationDeviationsSortByEnum{
	"timecreated": ListTargetAlertPolicyAssociationDeviationsSortByTimecreated,
	"displayname": ListTargetAlertPolicyAssociationDeviationsSortByDisplayname,
}

// GetListTargetAlertPolicyAssociationDeviationsSortByEnumValues Enumerates the set of values for ListTargetAlertPolicyAssociationDeviationsSortByEnum
func GetListTargetAlertPolicyAssociationDeviationsSortByEnumValues() []ListTargetAlertPolicyAssociationDeviationsSortByEnum {
	values := make([]ListTargetAlertPolicyAssociationDeviationsSortByEnum, 0)
	for _, v := range mappingListTargetAlertPolicyAssociationDeviationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAlertPolicyAssociationDeviationsSortByEnumStringValues Enumerates the set of values in String for ListTargetAlertPolicyAssociationDeviationsSortByEnum
func GetListTargetAlertPolicyAssociationDeviationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListTargetAlertPolicyAssociationDeviationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAlertPolicyAssociationDeviationsSortByEnum(val string) (ListTargetAlertPolicyAssociationDeviationsSortByEnum, bool) {
	enum, ok := mappingListTargetAlertPolicyAssociationDeviationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetAlertPolicyAssociationDeviationsSortOrderEnum Enum with underlying type: string
type ListTargetAlertPolicyAssociationDeviationsSortOrderEnum string

// Set of constants representing the allowable values for ListTargetAlertPolicyAssociationDeviationsSortOrderEnum
const (
	ListTargetAlertPolicyAssociationDeviationsSortOrderAsc  ListTargetAlertPolicyAssociationDeviationsSortOrderEnum = "ASC"
	ListTargetAlertPolicyAssociationDeviationsSortOrderDesc ListTargetAlertPolicyAssociationDeviationsSortOrderEnum = "DESC"
)

var mappingListTargetAlertPolicyAssociationDeviationsSortOrderEnum = map[string]ListTargetAlertPolicyAssociationDeviationsSortOrderEnum{
	"ASC":  ListTargetAlertPolicyAssociationDeviationsSortOrderAsc,
	"DESC": ListTargetAlertPolicyAssociationDeviationsSortOrderDesc,
}

var mappingListTargetAlertPolicyAssociationDeviationsSortOrderEnumLowerCase = map[string]ListTargetAlertPolicyAssociationDeviationsSortOrderEnum{
	"asc":  ListTargetAlertPolicyAssociationDeviationsSortOrderAsc,
	"desc": ListTargetAlertPolicyAssociationDeviationsSortOrderDesc,
}

// GetListTargetAlertPolicyAssociationDeviationsSortOrderEnumValues Enumerates the set of values for ListTargetAlertPolicyAssociationDeviationsSortOrderEnum
func GetListTargetAlertPolicyAssociationDeviationsSortOrderEnumValues() []ListTargetAlertPolicyAssociationDeviationsSortOrderEnum {
	values := make([]ListTargetAlertPolicyAssociationDeviationsSortOrderEnum, 0)
	for _, v := range mappingListTargetAlertPolicyAssociationDeviationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetAlertPolicyAssociationDeviationsSortOrderEnumStringValues Enumerates the set of values in String for ListTargetAlertPolicyAssociationDeviationsSortOrderEnum
func GetListTargetAlertPolicyAssociationDeviationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetAlertPolicyAssociationDeviationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetAlertPolicyAssociationDeviationsSortOrderEnum(val string) (ListTargetAlertPolicyAssociationDeviationsSortOrderEnum, bool) {
	enum, ok := mappingListTargetAlertPolicyAssociationDeviationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
