// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCryptoAssessmentsRequest wrapper for the ListCryptoAssessments operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessments.go.html to see an example of how to use ListCryptoAssessmentsRequest.
type ListCryptoAssessmentsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListCryptoAssessmentsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only crypto assessments that match the specified type.
	Type CryptoAssessmentTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter to return only resources associated with the specified crypto assessment OCID.
	AssessmentId *string `mandatory:"false" contributesTo:"query" name:"assessmentId"`

	// A filter to return only crypto assessments associated with the specified target OCID. When provided, targetType must also be specified.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only resources associated with any of the specified target OCIDs.
	TargetIds []string `contributesTo:"query" name:"targetIds" collectionFormat:"multi"`

	// A filter to return the target database group that matches the specified OCID.
	TargetDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupId"`

	// A filter to return crypto assessments belonging to the specified target type. `ListCryptoAssessments` returns assessment rows; use `targetDatabaseGroupId` to list the underlying target database assessments for a group.
	TargetType CryptoAssessmentTargetTypeEnum `mandatory:"false" contributesTo:"query" name:"targetType" omitEmpty:"true"`

	// A filter to return only crypto assessments that match any of the specified posture categories.
	PostureCategory []CryptoPostureCategoryEnum `contributesTo:"query" name:"postureCategory" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only crypto assessments whose scheduled execution state matches the specified value.
	IsAssessmentScheduled *bool `mandatory:"false" contributesTo:"query" name:"isAssessmentScheduled"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState CryptoAssessmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field used to sort crypto assessments. You can specify only one sort order (sortOrder).
	SortBy ListCryptoAssessmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCryptoAssessmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCryptoAssessmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAssessmentsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCryptoAssessmentsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTargetTypeEnum(string(request.TargetType)); !ok && request.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", request.TargetType, strings.Join(GetCryptoAssessmentTargetTypeEnumStringValues(), ",")))
	}
	for _, val := range request.PostureCategory {
		if _, ok := GetMappingCryptoPostureCategoryEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PostureCategory: %s. Supported values are: %s.", val, strings.Join(GetCryptoPostureCategoryEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingCryptoAssessmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCryptoAssessmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAssessmentsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAssessmentsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentsResponse wrapper for the ListCryptoAssessments operation
type ListCryptoAssessmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentCollection instances
	CryptoAssessmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCryptoAssessmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentsAccessLevelEnum Enum with underlying type: string
type ListCryptoAssessmentsAccessLevelEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentsAccessLevelEnum
const (
	ListCryptoAssessmentsAccessLevelRestricted ListCryptoAssessmentsAccessLevelEnum = "RESTRICTED"
	ListCryptoAssessmentsAccessLevelAccessible ListCryptoAssessmentsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCryptoAssessmentsAccessLevelEnum = map[string]ListCryptoAssessmentsAccessLevelEnum{
	"RESTRICTED": ListCryptoAssessmentsAccessLevelRestricted,
	"ACCESSIBLE": ListCryptoAssessmentsAccessLevelAccessible,
}

var mappingListCryptoAssessmentsAccessLevelEnumLowerCase = map[string]ListCryptoAssessmentsAccessLevelEnum{
	"restricted": ListCryptoAssessmentsAccessLevelRestricted,
	"accessible": ListCryptoAssessmentsAccessLevelAccessible,
}

// GetListCryptoAssessmentsAccessLevelEnumValues Enumerates the set of values for ListCryptoAssessmentsAccessLevelEnum
func GetListCryptoAssessmentsAccessLevelEnumValues() []ListCryptoAssessmentsAccessLevelEnum {
	values := make([]ListCryptoAssessmentsAccessLevelEnum, 0)
	for _, v := range mappingListCryptoAssessmentsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentsAccessLevelEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentsAccessLevelEnum
func GetListCryptoAssessmentsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCryptoAssessmentsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentsAccessLevelEnum(val string) (ListCryptoAssessmentsAccessLevelEnum, bool) {
	enum, ok := mappingListCryptoAssessmentsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentsSortByEnum Enum with underlying type: string
type ListCryptoAssessmentsSortByEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentsSortByEnum
const (
	ListCryptoAssessmentsSortByIssuecount  ListCryptoAssessmentsSortByEnum = "issueCount"
	ListCryptoAssessmentsSortByTimecreated ListCryptoAssessmentsSortByEnum = "timeCreated"
	ListCryptoAssessmentsSortByTimeupdated ListCryptoAssessmentsSortByEnum = "timeUpdated"
)

var mappingListCryptoAssessmentsSortByEnum = map[string]ListCryptoAssessmentsSortByEnum{
	"issueCount":  ListCryptoAssessmentsSortByIssuecount,
	"timeCreated": ListCryptoAssessmentsSortByTimecreated,
	"timeUpdated": ListCryptoAssessmentsSortByTimeupdated,
}

var mappingListCryptoAssessmentsSortByEnumLowerCase = map[string]ListCryptoAssessmentsSortByEnum{
	"issuecount":  ListCryptoAssessmentsSortByIssuecount,
	"timecreated": ListCryptoAssessmentsSortByTimecreated,
	"timeupdated": ListCryptoAssessmentsSortByTimeupdated,
}

// GetListCryptoAssessmentsSortByEnumValues Enumerates the set of values for ListCryptoAssessmentsSortByEnum
func GetListCryptoAssessmentsSortByEnumValues() []ListCryptoAssessmentsSortByEnum {
	values := make([]ListCryptoAssessmentsSortByEnum, 0)
	for _, v := range mappingListCryptoAssessmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentsSortByEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentsSortByEnum
func GetListCryptoAssessmentsSortByEnumStringValues() []string {
	return []string{
		"issueCount",
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListCryptoAssessmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentsSortByEnum(val string) (ListCryptoAssessmentsSortByEnum, bool) {
	enum, ok := mappingListCryptoAssessmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentsSortOrderEnum Enum with underlying type: string
type ListCryptoAssessmentsSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentsSortOrderEnum
const (
	ListCryptoAssessmentsSortOrderAsc  ListCryptoAssessmentsSortOrderEnum = "ASC"
	ListCryptoAssessmentsSortOrderDesc ListCryptoAssessmentsSortOrderEnum = "DESC"
)

var mappingListCryptoAssessmentsSortOrderEnum = map[string]ListCryptoAssessmentsSortOrderEnum{
	"ASC":  ListCryptoAssessmentsSortOrderAsc,
	"DESC": ListCryptoAssessmentsSortOrderDesc,
}

var mappingListCryptoAssessmentsSortOrderEnumLowerCase = map[string]ListCryptoAssessmentsSortOrderEnum{
	"asc":  ListCryptoAssessmentsSortOrderAsc,
	"desc": ListCryptoAssessmentsSortOrderDesc,
}

// GetListCryptoAssessmentsSortOrderEnumValues Enumerates the set of values for ListCryptoAssessmentsSortOrderEnum
func GetListCryptoAssessmentsSortOrderEnumValues() []ListCryptoAssessmentsSortOrderEnum {
	values := make([]ListCryptoAssessmentsSortOrderEnum, 0)
	for _, v := range mappingListCryptoAssessmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentsSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentsSortOrderEnum
func GetListCryptoAssessmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAssessmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentsSortOrderEnum(val string) (ListCryptoAssessmentsSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAssessmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
