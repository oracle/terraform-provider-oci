// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSendersRequest wrapper for the ListSenders operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/email/ListSenders.go.html to see an example of how to use ListSendersRequest.
type ListSendersRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The current state of a sender.
	LifecycleState SenderLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to only return resources that match the given domain exactly.
	Domain *string `mandatory:"false" contributesTo:"query" name:"domain"`

	// The email address of the approved sender.
	EmailAddress *string `mandatory:"false" contributesTo:"query" name:"emailAddress"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. The `TIMECREATED` value returns the list in in
	// descending order by default. The `EMAILADDRESS` value returns the list in
	// ascending order by default. Use the `SortOrderQueryParam` to change the
	// direction of the returned list of items.
	SortBy ListSendersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListSendersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSendersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSendersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSendersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSendersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSendersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSenderLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSenderLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSendersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSendersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSendersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSendersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSendersResponse wrapper for the ListSenders operation
type ListSendersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SenderSummary instances
	Items []SenderSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// The total number of items returned from the request.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListSendersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSendersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSendersSortByEnum Enum with underlying type: string
type ListSendersSortByEnum string

// Set of constants representing the allowable values for ListSendersSortByEnum
const (
	ListSendersSortByTimecreated  ListSendersSortByEnum = "TIMECREATED"
	ListSendersSortByEmailaddress ListSendersSortByEnum = "EMAILADDRESS"
)

var mappingListSendersSortByEnum = map[string]ListSendersSortByEnum{
	"TIMECREATED":  ListSendersSortByTimecreated,
	"EMAILADDRESS": ListSendersSortByEmailaddress,
}

var mappingListSendersSortByEnumLowerCase = map[string]ListSendersSortByEnum{
	"timecreated":  ListSendersSortByTimecreated,
	"emailaddress": ListSendersSortByEmailaddress,
}

// GetListSendersSortByEnumValues Enumerates the set of values for ListSendersSortByEnum
func GetListSendersSortByEnumValues() []ListSendersSortByEnum {
	values := make([]ListSendersSortByEnum, 0)
	for _, v := range mappingListSendersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSendersSortByEnumStringValues Enumerates the set of values in String for ListSendersSortByEnum
func GetListSendersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"EMAILADDRESS",
	}
}

// GetMappingListSendersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSendersSortByEnum(val string) (ListSendersSortByEnum, bool) {
	enum, ok := mappingListSendersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSendersSortOrderEnum Enum with underlying type: string
type ListSendersSortOrderEnum string

// Set of constants representing the allowable values for ListSendersSortOrderEnum
const (
	ListSendersSortOrderAsc  ListSendersSortOrderEnum = "ASC"
	ListSendersSortOrderDesc ListSendersSortOrderEnum = "DESC"
)

var mappingListSendersSortOrderEnum = map[string]ListSendersSortOrderEnum{
	"ASC":  ListSendersSortOrderAsc,
	"DESC": ListSendersSortOrderDesc,
}

var mappingListSendersSortOrderEnumLowerCase = map[string]ListSendersSortOrderEnum{
	"asc":  ListSendersSortOrderAsc,
	"desc": ListSendersSortOrderDesc,
}

// GetListSendersSortOrderEnumValues Enumerates the set of values for ListSendersSortOrderEnum
func GetListSendersSortOrderEnumValues() []ListSendersSortOrderEnum {
	values := make([]ListSendersSortOrderEnum, 0)
	for _, v := range mappingListSendersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSendersSortOrderEnumStringValues Enumerates the set of values in String for ListSendersSortOrderEnum
func GetListSendersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSendersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSendersSortOrderEnum(val string) (ListSendersSortOrderEnum, bool) {
	enum, ok := mappingListSendersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
