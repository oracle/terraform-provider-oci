// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package autoscaling

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutoScalingPoliciesRequest wrapper for the ListAutoScalingPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/autoscaling/ListAutoScalingPolicies.go.html to see an example of how to use ListAutoScalingPoliciesRequest.
type ListAutoScalingPoliciesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the autoscaling configuration.
	AutoScalingConfigurationId *string `mandatory:"true" contributesTo:"path" name:"autoScalingConfigurationId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return in a paginated "List" call. For important details
	// about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call. For important
	// details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	SortBy ListAutoScalingPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListAutoScalingPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutoScalingPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutoScalingPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutoScalingPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutoScalingPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutoScalingPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutoScalingPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutoScalingPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutoScalingPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutoScalingPoliciesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutoScalingPoliciesResponse wrapper for the ListAutoScalingPolicies operation
type ListAutoScalingPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutoScalingPolicySummary instances
	Items []AutoScalingPolicySummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAutoScalingPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutoScalingPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutoScalingPoliciesSortByEnum Enum with underlying type: string
type ListAutoScalingPoliciesSortByEnum string

// Set of constants representing the allowable values for ListAutoScalingPoliciesSortByEnum
const (
	ListAutoScalingPoliciesSortByTimecreated ListAutoScalingPoliciesSortByEnum = "TIMECREATED"
	ListAutoScalingPoliciesSortByDisplayname ListAutoScalingPoliciesSortByEnum = "DISPLAYNAME"
)

var mappingListAutoScalingPoliciesSortByEnum = map[string]ListAutoScalingPoliciesSortByEnum{
	"TIMECREATED": ListAutoScalingPoliciesSortByTimecreated,
	"DISPLAYNAME": ListAutoScalingPoliciesSortByDisplayname,
}

var mappingListAutoScalingPoliciesSortByEnumLowerCase = map[string]ListAutoScalingPoliciesSortByEnum{
	"timecreated": ListAutoScalingPoliciesSortByTimecreated,
	"displayname": ListAutoScalingPoliciesSortByDisplayname,
}

// GetListAutoScalingPoliciesSortByEnumValues Enumerates the set of values for ListAutoScalingPoliciesSortByEnum
func GetListAutoScalingPoliciesSortByEnumValues() []ListAutoScalingPoliciesSortByEnum {
	values := make([]ListAutoScalingPoliciesSortByEnum, 0)
	for _, v := range mappingListAutoScalingPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutoScalingPoliciesSortByEnumStringValues Enumerates the set of values in String for ListAutoScalingPoliciesSortByEnum
func GetListAutoScalingPoliciesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAutoScalingPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutoScalingPoliciesSortByEnum(val string) (ListAutoScalingPoliciesSortByEnum, bool) {
	enum, ok := mappingListAutoScalingPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutoScalingPoliciesSortOrderEnum Enum with underlying type: string
type ListAutoScalingPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListAutoScalingPoliciesSortOrderEnum
const (
	ListAutoScalingPoliciesSortOrderAsc  ListAutoScalingPoliciesSortOrderEnum = "ASC"
	ListAutoScalingPoliciesSortOrderDesc ListAutoScalingPoliciesSortOrderEnum = "DESC"
)

var mappingListAutoScalingPoliciesSortOrderEnum = map[string]ListAutoScalingPoliciesSortOrderEnum{
	"ASC":  ListAutoScalingPoliciesSortOrderAsc,
	"DESC": ListAutoScalingPoliciesSortOrderDesc,
}

var mappingListAutoScalingPoliciesSortOrderEnumLowerCase = map[string]ListAutoScalingPoliciesSortOrderEnum{
	"asc":  ListAutoScalingPoliciesSortOrderAsc,
	"desc": ListAutoScalingPoliciesSortOrderDesc,
}

// GetListAutoScalingPoliciesSortOrderEnumValues Enumerates the set of values for ListAutoScalingPoliciesSortOrderEnum
func GetListAutoScalingPoliciesSortOrderEnumValues() []ListAutoScalingPoliciesSortOrderEnum {
	values := make([]ListAutoScalingPoliciesSortOrderEnum, 0)
	for _, v := range mappingListAutoScalingPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutoScalingPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListAutoScalingPoliciesSortOrderEnum
func GetListAutoScalingPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutoScalingPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutoScalingPoliciesSortOrderEnum(val string) (ListAutoScalingPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListAutoScalingPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
