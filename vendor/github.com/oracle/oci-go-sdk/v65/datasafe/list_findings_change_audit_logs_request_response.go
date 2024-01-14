// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFindingsChangeAuditLogsRequest wrapper for the ListFindingsChangeAuditLogs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindingsChangeAuditLogs.go.html to see an example of how to use ListFindingsChangeAuditLogsRequest.
type ListFindingsChangeAuditLogsRequest struct {

	// The OCID of the security assessment.
	SecurityAssessmentId *string `mandatory:"true" contributesTo:"path" name:"securityAssessmentId"`

	// A filter to return only findings of a particular risk level.
	Severity ListFindingsChangeAuditLogsSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// The unique key that identifies the finding. It is a string and unique within a security assessment.
	FindingKey *string `mandatory:"false" contributesTo:"query" name:"findingKey"`

	// The unique title that identifies the finding. It is a string and unique within a security assessment.
	FindingTitle *string `mandatory:"false" contributesTo:"query" name:"findingTitle"`

	// A filter to check findings whose risks were deferred by the user.
	IsRiskDeferred *bool `mandatory:"false" contributesTo:"query" name:"isRiskDeferred"`

	// A filter to check which user modified the risk level of the finding.
	ModifiedBy *string `mandatory:"false" contributesTo:"query" name:"modifiedBy"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListFindingsChangeAuditLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order(sortOrder). The default order for timeUpdated is descending.
	SortBy ListFindingsChangeAuditLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifying `TimeValidUntilGreaterThanOrEqualToQueryParam` parameter
	// will retrieve all items for which the risk level modification by user will
	// no longer be valid greater than the date and time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T00:00:00.000Z
	TimeValidUntilGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeValidUntilGreaterThanOrEqualTo"`

	// Specifying `TimeValidUntilLessThanQueryParam` parameter
	// will retrieve all items for which the risk level modification by user will
	// be valid until less than the date and time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T00:00:00.000Z
	TimeValidUntilLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeValidUntilLessThan"`

	// Search for resources that were updated after a specific date.
	// Specifying this parameter corresponding `timeUpdatedGreaterThanOrEqualTo`
	// parameter will retrieve all resources updated after the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	TimeUpdatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedGreaterThanOrEqualTo"`

	// Search for resources that were updated before a specific date.
	// Specifying this parameter corresponding `timeUpdatedLessThan`
	// parameter will retrieve all resources updated before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	TimeUpdatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedLessThan"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFindingsChangeAuditLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFindingsChangeAuditLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFindingsChangeAuditLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFindingsChangeAuditLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFindingsChangeAuditLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFindingsChangeAuditLogsSeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetListFindingsChangeAuditLogsSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFindingsChangeAuditLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFindingsChangeAuditLogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFindingsChangeAuditLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFindingsChangeAuditLogsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFindingsChangeAuditLogsResponse wrapper for the ListFindingsChangeAuditLogs operation
type ListFindingsChangeAuditLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FindingsChangeAuditLogCollection instances
	FindingsChangeAuditLogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListFindingsChangeAuditLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFindingsChangeAuditLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFindingsChangeAuditLogsSeverityEnum Enum with underlying type: string
type ListFindingsChangeAuditLogsSeverityEnum string

// Set of constants representing the allowable values for ListFindingsChangeAuditLogsSeverityEnum
const (
	ListFindingsChangeAuditLogsSeverityHigh     ListFindingsChangeAuditLogsSeverityEnum = "HIGH"
	ListFindingsChangeAuditLogsSeverityMedium   ListFindingsChangeAuditLogsSeverityEnum = "MEDIUM"
	ListFindingsChangeAuditLogsSeverityLow      ListFindingsChangeAuditLogsSeverityEnum = "LOW"
	ListFindingsChangeAuditLogsSeverityEvaluate ListFindingsChangeAuditLogsSeverityEnum = "EVALUATE"
	ListFindingsChangeAuditLogsSeverityAdvisory ListFindingsChangeAuditLogsSeverityEnum = "ADVISORY"
	ListFindingsChangeAuditLogsSeverityPass     ListFindingsChangeAuditLogsSeverityEnum = "PASS"
	ListFindingsChangeAuditLogsSeverityDeferred ListFindingsChangeAuditLogsSeverityEnum = "DEFERRED"
)

var mappingListFindingsChangeAuditLogsSeverityEnum = map[string]ListFindingsChangeAuditLogsSeverityEnum{
	"HIGH":     ListFindingsChangeAuditLogsSeverityHigh,
	"MEDIUM":   ListFindingsChangeAuditLogsSeverityMedium,
	"LOW":      ListFindingsChangeAuditLogsSeverityLow,
	"EVALUATE": ListFindingsChangeAuditLogsSeverityEvaluate,
	"ADVISORY": ListFindingsChangeAuditLogsSeverityAdvisory,
	"PASS":     ListFindingsChangeAuditLogsSeverityPass,
	"DEFERRED": ListFindingsChangeAuditLogsSeverityDeferred,
}

var mappingListFindingsChangeAuditLogsSeverityEnumLowerCase = map[string]ListFindingsChangeAuditLogsSeverityEnum{
	"high":     ListFindingsChangeAuditLogsSeverityHigh,
	"medium":   ListFindingsChangeAuditLogsSeverityMedium,
	"low":      ListFindingsChangeAuditLogsSeverityLow,
	"evaluate": ListFindingsChangeAuditLogsSeverityEvaluate,
	"advisory": ListFindingsChangeAuditLogsSeverityAdvisory,
	"pass":     ListFindingsChangeAuditLogsSeverityPass,
	"deferred": ListFindingsChangeAuditLogsSeverityDeferred,
}

// GetListFindingsChangeAuditLogsSeverityEnumValues Enumerates the set of values for ListFindingsChangeAuditLogsSeverityEnum
func GetListFindingsChangeAuditLogsSeverityEnumValues() []ListFindingsChangeAuditLogsSeverityEnum {
	values := make([]ListFindingsChangeAuditLogsSeverityEnum, 0)
	for _, v := range mappingListFindingsChangeAuditLogsSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingsChangeAuditLogsSeverityEnumStringValues Enumerates the set of values in String for ListFindingsChangeAuditLogsSeverityEnum
func GetListFindingsChangeAuditLogsSeverityEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
		"ADVISORY",
		"PASS",
		"DEFERRED",
	}
}

// GetMappingListFindingsChangeAuditLogsSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingsChangeAuditLogsSeverityEnum(val string) (ListFindingsChangeAuditLogsSeverityEnum, bool) {
	enum, ok := mappingListFindingsChangeAuditLogsSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFindingsChangeAuditLogsSortOrderEnum Enum with underlying type: string
type ListFindingsChangeAuditLogsSortOrderEnum string

// Set of constants representing the allowable values for ListFindingsChangeAuditLogsSortOrderEnum
const (
	ListFindingsChangeAuditLogsSortOrderAsc  ListFindingsChangeAuditLogsSortOrderEnum = "ASC"
	ListFindingsChangeAuditLogsSortOrderDesc ListFindingsChangeAuditLogsSortOrderEnum = "DESC"
)

var mappingListFindingsChangeAuditLogsSortOrderEnum = map[string]ListFindingsChangeAuditLogsSortOrderEnum{
	"ASC":  ListFindingsChangeAuditLogsSortOrderAsc,
	"DESC": ListFindingsChangeAuditLogsSortOrderDesc,
}

var mappingListFindingsChangeAuditLogsSortOrderEnumLowerCase = map[string]ListFindingsChangeAuditLogsSortOrderEnum{
	"asc":  ListFindingsChangeAuditLogsSortOrderAsc,
	"desc": ListFindingsChangeAuditLogsSortOrderDesc,
}

// GetListFindingsChangeAuditLogsSortOrderEnumValues Enumerates the set of values for ListFindingsChangeAuditLogsSortOrderEnum
func GetListFindingsChangeAuditLogsSortOrderEnumValues() []ListFindingsChangeAuditLogsSortOrderEnum {
	values := make([]ListFindingsChangeAuditLogsSortOrderEnum, 0)
	for _, v := range mappingListFindingsChangeAuditLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingsChangeAuditLogsSortOrderEnumStringValues Enumerates the set of values in String for ListFindingsChangeAuditLogsSortOrderEnum
func GetListFindingsChangeAuditLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFindingsChangeAuditLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingsChangeAuditLogsSortOrderEnum(val string) (ListFindingsChangeAuditLogsSortOrderEnum, bool) {
	enum, ok := mappingListFindingsChangeAuditLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFindingsChangeAuditLogsSortByEnum Enum with underlying type: string
type ListFindingsChangeAuditLogsSortByEnum string

// Set of constants representing the allowable values for ListFindingsChangeAuditLogsSortByEnum
const (
	ListFindingsChangeAuditLogsSortByTimeupdated    ListFindingsChangeAuditLogsSortByEnum = "timeUpdated"
	ListFindingsChangeAuditLogsSortByModifiedby     ListFindingsChangeAuditLogsSortByEnum = "modifiedBy"
	ListFindingsChangeAuditLogsSortByIsriskdeferred ListFindingsChangeAuditLogsSortByEnum = "isRiskDeferred"
	ListFindingsChangeAuditLogsSortByTimevaliduntil ListFindingsChangeAuditLogsSortByEnum = "timeValidUntil"
)

var mappingListFindingsChangeAuditLogsSortByEnum = map[string]ListFindingsChangeAuditLogsSortByEnum{
	"timeUpdated":    ListFindingsChangeAuditLogsSortByTimeupdated,
	"modifiedBy":     ListFindingsChangeAuditLogsSortByModifiedby,
	"isRiskDeferred": ListFindingsChangeAuditLogsSortByIsriskdeferred,
	"timeValidUntil": ListFindingsChangeAuditLogsSortByTimevaliduntil,
}

var mappingListFindingsChangeAuditLogsSortByEnumLowerCase = map[string]ListFindingsChangeAuditLogsSortByEnum{
	"timeupdated":    ListFindingsChangeAuditLogsSortByTimeupdated,
	"modifiedby":     ListFindingsChangeAuditLogsSortByModifiedby,
	"isriskdeferred": ListFindingsChangeAuditLogsSortByIsriskdeferred,
	"timevaliduntil": ListFindingsChangeAuditLogsSortByTimevaliduntil,
}

// GetListFindingsChangeAuditLogsSortByEnumValues Enumerates the set of values for ListFindingsChangeAuditLogsSortByEnum
func GetListFindingsChangeAuditLogsSortByEnumValues() []ListFindingsChangeAuditLogsSortByEnum {
	values := make([]ListFindingsChangeAuditLogsSortByEnum, 0)
	for _, v := range mappingListFindingsChangeAuditLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingsChangeAuditLogsSortByEnumStringValues Enumerates the set of values in String for ListFindingsChangeAuditLogsSortByEnum
func GetListFindingsChangeAuditLogsSortByEnumStringValues() []string {
	return []string{
		"timeUpdated",
		"modifiedBy",
		"isRiskDeferred",
		"timeValidUntil",
	}
}

// GetMappingListFindingsChangeAuditLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingsChangeAuditLogsSortByEnum(val string) (ListFindingsChangeAuditLogsSortByEnum, bool) {
	enum, ok := mappingListFindingsChangeAuditLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
