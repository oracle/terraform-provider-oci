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

// ListFindingAnalyticsRequest wrapper for the ListFindingAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindingAnalytics.go.html to see an example of how to use ListFindingAnalyticsRequest.
type ListFindingAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListFindingAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only the findings that are marked as top findings.
	IsTopFinding *bool `mandatory:"false" contributesTo:"query" name:"isTopFinding"`

	// Attribute by which the finding analytics data should be grouped.
	GroupBy ListFindingAnalyticsGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// An optional filter to return only the top finding that match the specified status.
	TopFindingStatus FindingAnalyticsDimensionsTopFindingStatusEnum `mandatory:"false" contributesTo:"query" name:"topFindingStatus" omitEmpty:"true"`

	// A filter to return only findings of a particular risk level.
	Severity ListFindingAnalyticsSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// The unique key that identifies the finding. It is a string and unique within a security assessment.
	FindingKey *string `mandatory:"false" contributesTo:"query" name:"findingKey"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFindingAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFindingAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFindingAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFindingAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFindingAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFindingAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListFindingAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFindingAnalyticsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetListFindingAnalyticsGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFindingAnalyticsDimensionsTopFindingStatusEnum(string(request.TopFindingStatus)); !ok && request.TopFindingStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TopFindingStatus: %s. Supported values are: %s.", request.TopFindingStatus, strings.Join(GetFindingAnalyticsDimensionsTopFindingStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFindingAnalyticsSeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetListFindingAnalyticsSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFindingAnalyticsResponse wrapper for the ListFindingAnalytics operation
type ListFindingAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FindingAnalyticsCollection instances
	FindingAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListFindingAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFindingAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFindingAnalyticsAccessLevelEnum Enum with underlying type: string
type ListFindingAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListFindingAnalyticsAccessLevelEnum
const (
	ListFindingAnalyticsAccessLevelRestricted ListFindingAnalyticsAccessLevelEnum = "RESTRICTED"
	ListFindingAnalyticsAccessLevelAccessible ListFindingAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListFindingAnalyticsAccessLevelEnum = map[string]ListFindingAnalyticsAccessLevelEnum{
	"RESTRICTED": ListFindingAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListFindingAnalyticsAccessLevelAccessible,
}

var mappingListFindingAnalyticsAccessLevelEnumLowerCase = map[string]ListFindingAnalyticsAccessLevelEnum{
	"restricted": ListFindingAnalyticsAccessLevelRestricted,
	"accessible": ListFindingAnalyticsAccessLevelAccessible,
}

// GetListFindingAnalyticsAccessLevelEnumValues Enumerates the set of values for ListFindingAnalyticsAccessLevelEnum
func GetListFindingAnalyticsAccessLevelEnumValues() []ListFindingAnalyticsAccessLevelEnum {
	values := make([]ListFindingAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListFindingAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListFindingAnalyticsAccessLevelEnum
func GetListFindingAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListFindingAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingAnalyticsAccessLevelEnum(val string) (ListFindingAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListFindingAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFindingAnalyticsGroupByEnum Enum with underlying type: string
type ListFindingAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListFindingAnalyticsGroupByEnum
const (
	ListFindingAnalyticsGroupByFindingkeyandtopfindingstatus ListFindingAnalyticsGroupByEnum = "findingKeyAndTopFindingStatus"
	ListFindingAnalyticsGroupByFindingkeyandseverity         ListFindingAnalyticsGroupByEnum = "findingKeyAndSeverity"
)

var mappingListFindingAnalyticsGroupByEnum = map[string]ListFindingAnalyticsGroupByEnum{
	"findingKeyAndTopFindingStatus": ListFindingAnalyticsGroupByFindingkeyandtopfindingstatus,
	"findingKeyAndSeverity":         ListFindingAnalyticsGroupByFindingkeyandseverity,
}

var mappingListFindingAnalyticsGroupByEnumLowerCase = map[string]ListFindingAnalyticsGroupByEnum{
	"findingkeyandtopfindingstatus": ListFindingAnalyticsGroupByFindingkeyandtopfindingstatus,
	"findingkeyandseverity":         ListFindingAnalyticsGroupByFindingkeyandseverity,
}

// GetListFindingAnalyticsGroupByEnumValues Enumerates the set of values for ListFindingAnalyticsGroupByEnum
func GetListFindingAnalyticsGroupByEnumValues() []ListFindingAnalyticsGroupByEnum {
	values := make([]ListFindingAnalyticsGroupByEnum, 0)
	for _, v := range mappingListFindingAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListFindingAnalyticsGroupByEnum
func GetListFindingAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"findingKeyAndTopFindingStatus",
		"findingKeyAndSeverity",
	}
}

// GetMappingListFindingAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingAnalyticsGroupByEnum(val string) (ListFindingAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListFindingAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFindingAnalyticsSeverityEnum Enum with underlying type: string
type ListFindingAnalyticsSeverityEnum string

// Set of constants representing the allowable values for ListFindingAnalyticsSeverityEnum
const (
	ListFindingAnalyticsSeverityHigh     ListFindingAnalyticsSeverityEnum = "HIGH"
	ListFindingAnalyticsSeverityMedium   ListFindingAnalyticsSeverityEnum = "MEDIUM"
	ListFindingAnalyticsSeverityLow      ListFindingAnalyticsSeverityEnum = "LOW"
	ListFindingAnalyticsSeverityEvaluate ListFindingAnalyticsSeverityEnum = "EVALUATE"
	ListFindingAnalyticsSeverityAdvisory ListFindingAnalyticsSeverityEnum = "ADVISORY"
	ListFindingAnalyticsSeverityPass     ListFindingAnalyticsSeverityEnum = "PASS"
	ListFindingAnalyticsSeverityDeferred ListFindingAnalyticsSeverityEnum = "DEFERRED"
)

var mappingListFindingAnalyticsSeverityEnum = map[string]ListFindingAnalyticsSeverityEnum{
	"HIGH":     ListFindingAnalyticsSeverityHigh,
	"MEDIUM":   ListFindingAnalyticsSeverityMedium,
	"LOW":      ListFindingAnalyticsSeverityLow,
	"EVALUATE": ListFindingAnalyticsSeverityEvaluate,
	"ADVISORY": ListFindingAnalyticsSeverityAdvisory,
	"PASS":     ListFindingAnalyticsSeverityPass,
	"DEFERRED": ListFindingAnalyticsSeverityDeferred,
}

var mappingListFindingAnalyticsSeverityEnumLowerCase = map[string]ListFindingAnalyticsSeverityEnum{
	"high":     ListFindingAnalyticsSeverityHigh,
	"medium":   ListFindingAnalyticsSeverityMedium,
	"low":      ListFindingAnalyticsSeverityLow,
	"evaluate": ListFindingAnalyticsSeverityEvaluate,
	"advisory": ListFindingAnalyticsSeverityAdvisory,
	"pass":     ListFindingAnalyticsSeverityPass,
	"deferred": ListFindingAnalyticsSeverityDeferred,
}

// GetListFindingAnalyticsSeverityEnumValues Enumerates the set of values for ListFindingAnalyticsSeverityEnum
func GetListFindingAnalyticsSeverityEnumValues() []ListFindingAnalyticsSeverityEnum {
	values := make([]ListFindingAnalyticsSeverityEnum, 0)
	for _, v := range mappingListFindingAnalyticsSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingAnalyticsSeverityEnumStringValues Enumerates the set of values in String for ListFindingAnalyticsSeverityEnum
func GetListFindingAnalyticsSeverityEnumStringValues() []string {
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

// GetMappingListFindingAnalyticsSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingAnalyticsSeverityEnum(val string) (ListFindingAnalyticsSeverityEnum, bool) {
	enum, ok := mappingListFindingAnalyticsSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
