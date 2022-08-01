// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListContinuousQueryRequest wrapper for the ListContinuousQuery operation
type ListContinuousQueryRequest struct {

	// Compartment OCID to list resources in. See compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Severity of the continuous query.
	Severity ContinuousQuerySeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// Frequency in minutes.
	Frequency *string `mandatory:"false" contributesTo:"query" name:"frequency"`

	// Resource name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle state of the Continuous Query
	LifecycleState ContinuousQueryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only) for continuous queries. Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListContinuousQuerySortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListContinuousQuerySortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListContinuousQueryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListContinuousQueryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListContinuousQueryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListContinuousQueryRequest) RetryPolicy() common.OCIRetry {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListContinuousQueryRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContinuousQuerySeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetContinuousQuerySeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingContinuousQueryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetContinuousQueryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContinuousQuerySortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListContinuousQuerySortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListContinuousQuerySortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListContinuousQuerySortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListContinuousQueryResponse wrapper for the ListContinuousQuery operation
type ListContinuousQueryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ContinuousQuerySummaryCollection instances
	ContinuousQuerySummaryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListContinuousQueryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListContinuousQueryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListContinuousQuerySortByEnum Enum with underlying type: string
type ListContinuousQuerySortByEnum string

// Set of constants representing the allowable values for ListContinuousQuerySortByEnum
const (
	ListContinuousQuerySortByTimecreated    ListContinuousQuerySortByEnum = "timeCreated"
	ListContinuousQuerySortByDisplayname    ListContinuousQuerySortByEnum = "displayName"
	ListContinuousQuerySortBySeverity       ListContinuousQuerySortByEnum = "severity"
	ListContinuousQuerySortByFrequency      ListContinuousQuerySortByEnum = "frequency"
	ListContinuousQuerySortByLifecyclestate ListContinuousQuerySortByEnum = "lifecycleState"
)

var mappingListContinuousQuerySortByEnum = map[string]ListContinuousQuerySortByEnum{
	"timeCreated":    ListContinuousQuerySortByTimecreated,
	"displayName":    ListContinuousQuerySortByDisplayname,
	"severity":       ListContinuousQuerySortBySeverity,
	"frequency":      ListContinuousQuerySortByFrequency,
	"lifecycleState": ListContinuousQuerySortByLifecyclestate,
}

var mappingListContinuousQuerySortByEnumLowerCase = map[string]ListContinuousQuerySortByEnum{
	"timecreated":    ListContinuousQuerySortByTimecreated,
	"displayname":    ListContinuousQuerySortByDisplayname,
	"severity":       ListContinuousQuerySortBySeverity,
	"frequency":      ListContinuousQuerySortByFrequency,
	"lifecyclestate": ListContinuousQuerySortByLifecyclestate,
}

// GetListContinuousQuerySortByEnumValues Enumerates the set of values for ListContinuousQuerySortByEnum
func GetListContinuousQuerySortByEnumValues() []ListContinuousQuerySortByEnum {
	values := make([]ListContinuousQuerySortByEnum, 0)
	for _, v := range mappingListContinuousQuerySortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListContinuousQuerySortByEnumStringValues Enumerates the set of values in String for ListContinuousQuerySortByEnum
func GetListContinuousQuerySortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"severity",
		"frequency",
		"lifecycleState",
	}
}

// GetMappingListContinuousQuerySortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContinuousQuerySortByEnum(val string) (ListContinuousQuerySortByEnum, bool) {
	enum, ok := mappingListContinuousQuerySortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListContinuousQuerySortOrderEnum Enum with underlying type: string
type ListContinuousQuerySortOrderEnum string

// Set of constants representing the allowable values for ListContinuousQuerySortOrderEnum
const (
	ListContinuousQuerySortOrderAsc  ListContinuousQuerySortOrderEnum = "ASC"
	ListContinuousQuerySortOrderDesc ListContinuousQuerySortOrderEnum = "DESC"
)

var mappingListContinuousQuerySortOrderEnum = map[string]ListContinuousQuerySortOrderEnum{
	"ASC":  ListContinuousQuerySortOrderAsc,
	"DESC": ListContinuousQuerySortOrderDesc,
}

var mappingListContinuousQuerySortOrderEnumLowerCase = map[string]ListContinuousQuerySortOrderEnum{
	"asc":  ListContinuousQuerySortOrderAsc,
	"desc": ListContinuousQuerySortOrderDesc,
}

// GetListContinuousQuerySortOrderEnumValues Enumerates the set of values for ListContinuousQuerySortOrderEnum
func GetListContinuousQuerySortOrderEnumValues() []ListContinuousQuerySortOrderEnum {
	values := make([]ListContinuousQuerySortOrderEnum, 0)
	for _, v := range mappingListContinuousQuerySortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListContinuousQuerySortOrderEnumStringValues Enumerates the set of values in String for ListContinuousQuerySortOrderEnum
func GetListContinuousQuerySortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListContinuousQuerySortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListContinuousQuerySortOrderEnum(val string) (ListContinuousQuerySortOrderEnum, bool) {
	enum, ok := mappingListContinuousQuerySortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
