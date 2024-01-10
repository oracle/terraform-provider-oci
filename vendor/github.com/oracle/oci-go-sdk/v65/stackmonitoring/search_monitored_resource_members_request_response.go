// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SearchMonitoredResourceMembersRequest wrapper for the SearchMonitoredResourceMembers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/SearchMonitoredResourceMembers.go.html to see an example of how to use SearchMonitoredResourceMembersRequest.
type SearchMonitoredResourceMembersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of monitored resource.
	MonitoredResourceId *string `mandatory:"true" contributesTo:"path" name:"monitoredResourceId"`

	// Search criteria for listing member monitored resources.
	SearchMonitoredResourceMembersDetails `contributesTo:"body"`

	// If this query parameter is specified, the result is sorted by this query parameter value.
	SortBy SearchMonitoredResourceMembersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SearchMonitoredResourceMembersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SearchMonitoredResourceMembersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SearchMonitoredResourceMembersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SearchMonitoredResourceMembersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SearchMonitoredResourceMembersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SearchMonitoredResourceMembersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSearchMonitoredResourceMembersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSearchMonitoredResourceMembersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSearchMonitoredResourceMembersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSearchMonitoredResourceMembersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SearchMonitoredResourceMembersResponse wrapper for the SearchMonitoredResourceMembers operation
type SearchMonitoredResourceMembersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoredResourceMembersCollection instances
	MonitoredResourceMembersCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response SearchMonitoredResourceMembersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SearchMonitoredResourceMembersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SearchMonitoredResourceMembersSortByEnum Enum with underlying type: string
type SearchMonitoredResourceMembersSortByEnum string

// Set of constants representing the allowable values for SearchMonitoredResourceMembersSortByEnum
const (
	SearchMonitoredResourceMembersSortByResourcename       SearchMonitoredResourceMembersSortByEnum = "resourceName"
	SearchMonitoredResourceMembersSortByResourcetype       SearchMonitoredResourceMembersSortByEnum = "resourceType"
	SearchMonitoredResourceMembersSortBySourceresourcetype SearchMonitoredResourceMembersSortByEnum = "sourceResourceType"
)

var mappingSearchMonitoredResourceMembersSortByEnum = map[string]SearchMonitoredResourceMembersSortByEnum{
	"resourceName":       SearchMonitoredResourceMembersSortByResourcename,
	"resourceType":       SearchMonitoredResourceMembersSortByResourcetype,
	"sourceResourceType": SearchMonitoredResourceMembersSortBySourceresourcetype,
}

var mappingSearchMonitoredResourceMembersSortByEnumLowerCase = map[string]SearchMonitoredResourceMembersSortByEnum{
	"resourcename":       SearchMonitoredResourceMembersSortByResourcename,
	"resourcetype":       SearchMonitoredResourceMembersSortByResourcetype,
	"sourceresourcetype": SearchMonitoredResourceMembersSortBySourceresourcetype,
}

// GetSearchMonitoredResourceMembersSortByEnumValues Enumerates the set of values for SearchMonitoredResourceMembersSortByEnum
func GetSearchMonitoredResourceMembersSortByEnumValues() []SearchMonitoredResourceMembersSortByEnum {
	values := make([]SearchMonitoredResourceMembersSortByEnum, 0)
	for _, v := range mappingSearchMonitoredResourceMembersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchMonitoredResourceMembersSortByEnumStringValues Enumerates the set of values in String for SearchMonitoredResourceMembersSortByEnum
func GetSearchMonitoredResourceMembersSortByEnumStringValues() []string {
	return []string{
		"resourceName",
		"resourceType",
		"sourceResourceType",
	}
}

// GetMappingSearchMonitoredResourceMembersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchMonitoredResourceMembersSortByEnum(val string) (SearchMonitoredResourceMembersSortByEnum, bool) {
	enum, ok := mappingSearchMonitoredResourceMembersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SearchMonitoredResourceMembersSortOrderEnum Enum with underlying type: string
type SearchMonitoredResourceMembersSortOrderEnum string

// Set of constants representing the allowable values for SearchMonitoredResourceMembersSortOrderEnum
const (
	SearchMonitoredResourceMembersSortOrderAsc  SearchMonitoredResourceMembersSortOrderEnum = "ASC"
	SearchMonitoredResourceMembersSortOrderDesc SearchMonitoredResourceMembersSortOrderEnum = "DESC"
)

var mappingSearchMonitoredResourceMembersSortOrderEnum = map[string]SearchMonitoredResourceMembersSortOrderEnum{
	"ASC":  SearchMonitoredResourceMembersSortOrderAsc,
	"DESC": SearchMonitoredResourceMembersSortOrderDesc,
}

var mappingSearchMonitoredResourceMembersSortOrderEnumLowerCase = map[string]SearchMonitoredResourceMembersSortOrderEnum{
	"asc":  SearchMonitoredResourceMembersSortOrderAsc,
	"desc": SearchMonitoredResourceMembersSortOrderDesc,
}

// GetSearchMonitoredResourceMembersSortOrderEnumValues Enumerates the set of values for SearchMonitoredResourceMembersSortOrderEnum
func GetSearchMonitoredResourceMembersSortOrderEnumValues() []SearchMonitoredResourceMembersSortOrderEnum {
	values := make([]SearchMonitoredResourceMembersSortOrderEnum, 0)
	for _, v := range mappingSearchMonitoredResourceMembersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchMonitoredResourceMembersSortOrderEnumStringValues Enumerates the set of values in String for SearchMonitoredResourceMembersSortOrderEnum
func GetSearchMonitoredResourceMembersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSearchMonitoredResourceMembersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchMonitoredResourceMembersSortOrderEnum(val string) (SearchMonitoredResourceMembersSortOrderEnum, bool) {
	enum, ok := mappingSearchMonitoredResourceMembersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
