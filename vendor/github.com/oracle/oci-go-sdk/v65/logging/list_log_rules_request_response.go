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

// ListLogRulesRequest wrapper for the ListLogRules operation
type ListLogRulesRequest struct {

	// Compartment OCID to list resources in. See compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Severity of the log rule.
	Severity LogRuleSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// Frequency in minutes.
	Frequency *string `mandatory:"false" contributesTo:"query" name:"frequency"`

	// Resource name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle state of the log rule.
	LifecycleState LogRuleLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only) for log rules. Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListLogRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogRuleSeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetLogRuleSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogRuleLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetLogRuleLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogRulesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogRulesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogRulesResponse wrapper for the ListLogRules operation
type ListLogRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogRuleSummaryCollection instances
	LogRuleSummaryCollection `presentIn:"body"`

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

func (response ListLogRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogRulesSortByEnum Enum with underlying type: string
type ListLogRulesSortByEnum string

// Set of constants representing the allowable values for ListLogRulesSortByEnum
const (
	ListLogRulesSortByTimecreated    ListLogRulesSortByEnum = "timeCreated"
	ListLogRulesSortByDisplayname    ListLogRulesSortByEnum = "displayName"
	ListLogRulesSortBySeverity       ListLogRulesSortByEnum = "severity"
	ListLogRulesSortByFrequency      ListLogRulesSortByEnum = "frequency"
	ListLogRulesSortByLifecyclestate ListLogRulesSortByEnum = "lifecycleState"
)

var mappingListLogRulesSortByEnum = map[string]ListLogRulesSortByEnum{
	"timeCreated":    ListLogRulesSortByTimecreated,
	"displayName":    ListLogRulesSortByDisplayname,
	"severity":       ListLogRulesSortBySeverity,
	"frequency":      ListLogRulesSortByFrequency,
	"lifecycleState": ListLogRulesSortByLifecyclestate,
}

var mappingListLogRulesSortByEnumLowerCase = map[string]ListLogRulesSortByEnum{
	"timecreated":    ListLogRulesSortByTimecreated,
	"displayname":    ListLogRulesSortByDisplayname,
	"severity":       ListLogRulesSortBySeverity,
	"frequency":      ListLogRulesSortByFrequency,
	"lifecyclestate": ListLogRulesSortByLifecyclestate,
}

// GetListLogRulesSortByEnumValues Enumerates the set of values for ListLogRulesSortByEnum
func GetListLogRulesSortByEnumValues() []ListLogRulesSortByEnum {
	values := make([]ListLogRulesSortByEnum, 0)
	for _, v := range mappingListLogRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogRulesSortByEnumStringValues Enumerates the set of values in String for ListLogRulesSortByEnum
func GetListLogRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"severity",
		"frequency",
		"lifecycleState",
	}
}

// GetMappingListLogRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogRulesSortByEnum(val string) (ListLogRulesSortByEnum, bool) {
	enum, ok := mappingListLogRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogRulesSortOrderEnum Enum with underlying type: string
type ListLogRulesSortOrderEnum string

// Set of constants representing the allowable values for ListLogRulesSortOrderEnum
const (
	ListLogRulesSortOrderAsc  ListLogRulesSortOrderEnum = "ASC"
	ListLogRulesSortOrderDesc ListLogRulesSortOrderEnum = "DESC"
)

var mappingListLogRulesSortOrderEnum = map[string]ListLogRulesSortOrderEnum{
	"ASC":  ListLogRulesSortOrderAsc,
	"DESC": ListLogRulesSortOrderDesc,
}

var mappingListLogRulesSortOrderEnumLowerCase = map[string]ListLogRulesSortOrderEnum{
	"asc":  ListLogRulesSortOrderAsc,
	"desc": ListLogRulesSortOrderDesc,
}

// GetListLogRulesSortOrderEnumValues Enumerates the set of values for ListLogRulesSortOrderEnum
func GetListLogRulesSortOrderEnumValues() []ListLogRulesSortOrderEnum {
	values := make([]ListLogRulesSortOrderEnum, 0)
	for _, v := range mappingListLogRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogRulesSortOrderEnumStringValues Enumerates the set of values in String for ListLogRulesSortOrderEnum
func GetListLogRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogRulesSortOrderEnum(val string) (ListLogRulesSortOrderEnum, bool) {
	enum, ok := mappingListLogRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
