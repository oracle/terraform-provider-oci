// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListShareSetsRequest wrapper for the ListShareSets operation
type ListShareSetsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A comment. It does not have to be unique, and it is changeable.
	// Example: `My share`
	Comment *string `mandatory:"false" contributesTo:"query" name:"comment"`

	// A customer-provided DNS name. This name plays a critical role in
	// establishing the server's place in the customer SMB security hierarchy.
	// For example, if an SMB server has a DNS name of
	// register5.store34.california.usa.marks-hats.com, then this particular
	// server is part of the store34.california.usa.marks-hats.com security
	// domain which in turn is part of the california.usa.marks-hats.com domain, which
	// in turn is part of the usa.marks-hats.com domain,
	// which in turn is part of the marks-hats.com security domain.
	// Must be unique across all FQDNs in the subnet and comply
	// with RFC 952 (https://tools.ietf.org/html/rfc952)
	// and RFC 1123 (https://tools.ietf.org/html/rfc1123).
	CustomFqdn *string `mandatory:"false" contributesTo:"query" name:"customFqdn"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the shareSets resource type.
	LifecycleState ShareSetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter results by OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for
	// the resouce type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field to sort by. You can provide either value, but not both.
	// By default, when you sort by time created, results are shown
	// in descending order. When you sort by comment, results are
	// shown in ascending order.
	SortBy ListShareSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListShareSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListShareSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListShareSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListShareSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListShareSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListShareSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShareSetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetShareSetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListShareSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListShareSetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListShareSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListShareSetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListShareSetsResponse wrapper for the ListShareSets operation
type ListShareSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ShareSetSummary instances
	Items []ShareSetSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListShareSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListShareSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListShareSetsSortByEnum Enum with underlying type: string
type ListShareSetsSortByEnum string

// Set of constants representing the allowable values for ListShareSetsSortByEnum
const (
	ListShareSetsSortByTimecreated ListShareSetsSortByEnum = "TIMECREATED"
	ListShareSetsSortByComment     ListShareSetsSortByEnum = "COMMENT"
)

var mappingListShareSetsSortByEnum = map[string]ListShareSetsSortByEnum{
	"TIMECREATED": ListShareSetsSortByTimecreated,
	"COMMENT":     ListShareSetsSortByComment,
}

var mappingListShareSetsSortByEnumLowerCase = map[string]ListShareSetsSortByEnum{
	"timecreated": ListShareSetsSortByTimecreated,
	"comment":     ListShareSetsSortByComment,
}

// GetListShareSetsSortByEnumValues Enumerates the set of values for ListShareSetsSortByEnum
func GetListShareSetsSortByEnumValues() []ListShareSetsSortByEnum {
	values := make([]ListShareSetsSortByEnum, 0)
	for _, v := range mappingListShareSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListShareSetsSortByEnumStringValues Enumerates the set of values in String for ListShareSetsSortByEnum
func GetListShareSetsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"COMMENT",
	}
}

// GetMappingListShareSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListShareSetsSortByEnum(val string) (ListShareSetsSortByEnum, bool) {
	enum, ok := mappingListShareSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListShareSetsSortOrderEnum Enum with underlying type: string
type ListShareSetsSortOrderEnum string

// Set of constants representing the allowable values for ListShareSetsSortOrderEnum
const (
	ListShareSetsSortOrderAsc  ListShareSetsSortOrderEnum = "ASC"
	ListShareSetsSortOrderDesc ListShareSetsSortOrderEnum = "DESC"
)

var mappingListShareSetsSortOrderEnum = map[string]ListShareSetsSortOrderEnum{
	"ASC":  ListShareSetsSortOrderAsc,
	"DESC": ListShareSetsSortOrderDesc,
}

var mappingListShareSetsSortOrderEnumLowerCase = map[string]ListShareSetsSortOrderEnum{
	"asc":  ListShareSetsSortOrderAsc,
	"desc": ListShareSetsSortOrderDesc,
}

// GetListShareSetsSortOrderEnumValues Enumerates the set of values for ListShareSetsSortOrderEnum
func GetListShareSetsSortOrderEnumValues() []ListShareSetsSortOrderEnum {
	values := make([]ListShareSetsSortOrderEnum, 0)
	for _, v := range mappingListShareSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListShareSetsSortOrderEnumStringValues Enumerates the set of values in String for ListShareSetsSortOrderEnum
func GetListShareSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListShareSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListShareSetsSortOrderEnum(val string) (ListShareSetsSortOrderEnum, bool) {
	enum, ok := mappingListShareSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
