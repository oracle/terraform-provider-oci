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

// ListEmailOutboundIpsRequest wrapper for the ListEmailOutboundIps operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/email/ListEmailOutboundIps.go.html to see an example of how to use ListEmailOutboundIpsRequest.
type ListEmailOutboundIpsRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The outbound IP address assigned to the tenancy.
	OutboundIp *string `mandatory:"false" contributesTo:"query" name:"outboundIp"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailOutboundIpsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState EmailOutboundIpSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter returned list by specified assignment state.
	AssignmentState EmailOutboundIpSummaryAssignmentStateEnum `mandatory:"false" contributesTo:"query" name:"assignmentState" omitEmpty:"true"`

	// The field to sort by. The `OUTBOUNDIP` value returns the list in
	// ascending order or Outbound Ip address by default. Use the `SortOrderQueryParam` to change the
	// direction of the returned list of items.
	SortBy ListEmailOutboundIpsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailOutboundIpsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailOutboundIpsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailOutboundIpsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailOutboundIpsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailOutboundIpsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailOutboundIpsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailOutboundIpsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailOutboundIpSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailOutboundIpSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailOutboundIpSummaryAssignmentStateEnum(string(request.AssignmentState)); !ok && request.AssignmentState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssignmentState: %s. Supported values are: %s.", request.AssignmentState, strings.Join(GetEmailOutboundIpSummaryAssignmentStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailOutboundIpsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailOutboundIpsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailOutboundIpsResponse wrapper for the ListEmailOutboundIps operation
type ListEmailOutboundIpsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailOutboundIpCollection instances
	EmailOutboundIpCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEmailOutboundIpsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailOutboundIpsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailOutboundIpsSortOrderEnum Enum with underlying type: string
type ListEmailOutboundIpsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailOutboundIpsSortOrderEnum
const (
	ListEmailOutboundIpsSortOrderAsc  ListEmailOutboundIpsSortOrderEnum = "ASC"
	ListEmailOutboundIpsSortOrderDesc ListEmailOutboundIpsSortOrderEnum = "DESC"
)

var mappingListEmailOutboundIpsSortOrderEnum = map[string]ListEmailOutboundIpsSortOrderEnum{
	"ASC":  ListEmailOutboundIpsSortOrderAsc,
	"DESC": ListEmailOutboundIpsSortOrderDesc,
}

var mappingListEmailOutboundIpsSortOrderEnumLowerCase = map[string]ListEmailOutboundIpsSortOrderEnum{
	"asc":  ListEmailOutboundIpsSortOrderAsc,
	"desc": ListEmailOutboundIpsSortOrderDesc,
}

// GetListEmailOutboundIpsSortOrderEnumValues Enumerates the set of values for ListEmailOutboundIpsSortOrderEnum
func GetListEmailOutboundIpsSortOrderEnumValues() []ListEmailOutboundIpsSortOrderEnum {
	values := make([]ListEmailOutboundIpsSortOrderEnum, 0)
	for _, v := range mappingListEmailOutboundIpsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailOutboundIpsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailOutboundIpsSortOrderEnum
func GetListEmailOutboundIpsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailOutboundIpsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailOutboundIpsSortOrderEnum(val string) (ListEmailOutboundIpsSortOrderEnum, bool) {
	enum, ok := mappingListEmailOutboundIpsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailOutboundIpsSortByEnum Enum with underlying type: string
type ListEmailOutboundIpsSortByEnum string

// Set of constants representing the allowable values for ListEmailOutboundIpsSortByEnum
const (
	ListEmailOutboundIpsSortByOutboundip ListEmailOutboundIpsSortByEnum = "OUTBOUNDIP"
)

var mappingListEmailOutboundIpsSortByEnum = map[string]ListEmailOutboundIpsSortByEnum{
	"OUTBOUNDIP": ListEmailOutboundIpsSortByOutboundip,
}

var mappingListEmailOutboundIpsSortByEnumLowerCase = map[string]ListEmailOutboundIpsSortByEnum{
	"outboundip": ListEmailOutboundIpsSortByOutboundip,
}

// GetListEmailOutboundIpsSortByEnumValues Enumerates the set of values for ListEmailOutboundIpsSortByEnum
func GetListEmailOutboundIpsSortByEnumValues() []ListEmailOutboundIpsSortByEnum {
	values := make([]ListEmailOutboundIpsSortByEnum, 0)
	for _, v := range mappingListEmailOutboundIpsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailOutboundIpsSortByEnumStringValues Enumerates the set of values in String for ListEmailOutboundIpsSortByEnum
func GetListEmailOutboundIpsSortByEnumStringValues() []string {
	return []string{
		"OUTBOUNDIP",
	}
}

// GetMappingListEmailOutboundIpsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailOutboundIpsSortByEnum(val string) (ListEmailOutboundIpsSortByEnum, bool) {
	enum, ok := mappingListEmailOutboundIpsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
