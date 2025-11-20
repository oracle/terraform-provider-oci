// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailIpPoolsRequest wrapper for the ListEmailIpPools operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/email/ListEmailIpPools.go.html to see an example of how to use ListEmailIpPoolsRequest.
type ListEmailIpPoolsRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState EmailIpPoolLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailIpPoolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the Email IpPools.
	// Default: `TIMECREATED`
	// * **TIMECREATED:** Sorts by timeCreated.
	SortBy ListEmailIpPoolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailIpPoolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailIpPoolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailIpPoolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailIpPoolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailIpPoolsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailIpPoolLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailIpPoolLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailIpPoolsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailIpPoolsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailIpPoolsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailIpPoolsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailIpPoolsResponse wrapper for the ListEmailIpPools operation
type ListEmailIpPoolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailIpPoolCollection instances
	EmailIpPoolCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEmailIpPoolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailIpPoolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailIpPoolsSortOrderEnum Enum with underlying type: string
type ListEmailIpPoolsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailIpPoolsSortOrderEnum
const (
	ListEmailIpPoolsSortOrderAsc  ListEmailIpPoolsSortOrderEnum = "ASC"
	ListEmailIpPoolsSortOrderDesc ListEmailIpPoolsSortOrderEnum = "DESC"
)

var mappingListEmailIpPoolsSortOrderEnum = map[string]ListEmailIpPoolsSortOrderEnum{
	"ASC":  ListEmailIpPoolsSortOrderAsc,
	"DESC": ListEmailIpPoolsSortOrderDesc,
}

var mappingListEmailIpPoolsSortOrderEnumLowerCase = map[string]ListEmailIpPoolsSortOrderEnum{
	"asc":  ListEmailIpPoolsSortOrderAsc,
	"desc": ListEmailIpPoolsSortOrderDesc,
}

// GetListEmailIpPoolsSortOrderEnumValues Enumerates the set of values for ListEmailIpPoolsSortOrderEnum
func GetListEmailIpPoolsSortOrderEnumValues() []ListEmailIpPoolsSortOrderEnum {
	values := make([]ListEmailIpPoolsSortOrderEnum, 0)
	for _, v := range mappingListEmailIpPoolsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailIpPoolsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailIpPoolsSortOrderEnum
func GetListEmailIpPoolsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailIpPoolsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailIpPoolsSortOrderEnum(val string) (ListEmailIpPoolsSortOrderEnum, bool) {
	enum, ok := mappingListEmailIpPoolsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailIpPoolsSortByEnum Enum with underlying type: string
type ListEmailIpPoolsSortByEnum string

// Set of constants representing the allowable values for ListEmailIpPoolsSortByEnum
const (
	ListEmailIpPoolsSortByTimecreated ListEmailIpPoolsSortByEnum = "TIMECREATED"
	ListEmailIpPoolsSortByName        ListEmailIpPoolsSortByEnum = "NAME"
)

var mappingListEmailIpPoolsSortByEnum = map[string]ListEmailIpPoolsSortByEnum{
	"TIMECREATED": ListEmailIpPoolsSortByTimecreated,
	"NAME":        ListEmailIpPoolsSortByName,
}

var mappingListEmailIpPoolsSortByEnumLowerCase = map[string]ListEmailIpPoolsSortByEnum{
	"timecreated": ListEmailIpPoolsSortByTimecreated,
	"name":        ListEmailIpPoolsSortByName,
}

// GetListEmailIpPoolsSortByEnumValues Enumerates the set of values for ListEmailIpPoolsSortByEnum
func GetListEmailIpPoolsSortByEnumValues() []ListEmailIpPoolsSortByEnum {
	values := make([]ListEmailIpPoolsSortByEnum, 0)
	for _, v := range mappingListEmailIpPoolsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailIpPoolsSortByEnumStringValues Enumerates the set of values in String for ListEmailIpPoolsSortByEnum
func GetListEmailIpPoolsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListEmailIpPoolsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailIpPoolsSortByEnum(val string) (ListEmailIpPoolsSortByEnum, bool) {
	enum, ok := mappingListEmailIpPoolsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
