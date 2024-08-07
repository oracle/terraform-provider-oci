// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPullRequestActivitiesRequest wrapper for the ListPullRequestActivities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListPullRequestActivities.go.html to see an example of how to use ListPullRequestActivitiesRequest.
type ListPullRequestActivitiesRequest struct {

	// unique PullRequest identifier
	PullRequestId *string `mandatory:"true" contributesTo:"path" name:"pullRequestId"`

	// An optional filter to list activities based on activity type. If no value is specified, all activity types will returned.
	ActivityType ListPullRequestActivitiesActivityTypeEnum `mandatory:"false" contributesTo:"query" name:"activityType" omitEmpty:"true"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListPullRequestActivitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPullRequestActivitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPullRequestActivitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPullRequestActivitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPullRequestActivitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPullRequestActivitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPullRequestActivitiesActivityTypeEnum(string(request.ActivityType)); !ok && request.ActivityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActivityType: %s. Supported values are: %s.", request.ActivityType, strings.Join(GetListPullRequestActivitiesActivityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPullRequestActivitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPullRequestActivitiesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPullRequestActivitiesResponse wrapper for the ListPullRequestActivities operation
type ListPullRequestActivitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PullRequestActivityCollection instances
	PullRequestActivityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPullRequestActivitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPullRequestActivitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPullRequestActivitiesActivityTypeEnum Enum with underlying type: string
type ListPullRequestActivitiesActivityTypeEnum string

// Set of constants representing the allowable values for ListPullRequestActivitiesActivityTypeEnum
const (
	ListPullRequestActivitiesActivityTypeLifecycle ListPullRequestActivitiesActivityTypeEnum = "LIFECYCLE"
	ListPullRequestActivitiesActivityTypeApproval  ListPullRequestActivitiesActivityTypeEnum = "APPROVAL"
	ListPullRequestActivitiesActivityTypeCommit    ListPullRequestActivitiesActivityTypeEnum = "COMMIT"
	ListPullRequestActivitiesActivityTypeReviewer  ListPullRequestActivitiesActivityTypeEnum = "REVIEWER"
	ListPullRequestActivitiesActivityTypeComment   ListPullRequestActivitiesActivityTypeEnum = "COMMENT"
)

var mappingListPullRequestActivitiesActivityTypeEnum = map[string]ListPullRequestActivitiesActivityTypeEnum{
	"LIFECYCLE": ListPullRequestActivitiesActivityTypeLifecycle,
	"APPROVAL":  ListPullRequestActivitiesActivityTypeApproval,
	"COMMIT":    ListPullRequestActivitiesActivityTypeCommit,
	"REVIEWER":  ListPullRequestActivitiesActivityTypeReviewer,
	"COMMENT":   ListPullRequestActivitiesActivityTypeComment,
}

var mappingListPullRequestActivitiesActivityTypeEnumLowerCase = map[string]ListPullRequestActivitiesActivityTypeEnum{
	"lifecycle": ListPullRequestActivitiesActivityTypeLifecycle,
	"approval":  ListPullRequestActivitiesActivityTypeApproval,
	"commit":    ListPullRequestActivitiesActivityTypeCommit,
	"reviewer":  ListPullRequestActivitiesActivityTypeReviewer,
	"comment":   ListPullRequestActivitiesActivityTypeComment,
}

// GetListPullRequestActivitiesActivityTypeEnumValues Enumerates the set of values for ListPullRequestActivitiesActivityTypeEnum
func GetListPullRequestActivitiesActivityTypeEnumValues() []ListPullRequestActivitiesActivityTypeEnum {
	values := make([]ListPullRequestActivitiesActivityTypeEnum, 0)
	for _, v := range mappingListPullRequestActivitiesActivityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListPullRequestActivitiesActivityTypeEnumStringValues Enumerates the set of values in String for ListPullRequestActivitiesActivityTypeEnum
func GetListPullRequestActivitiesActivityTypeEnumStringValues() []string {
	return []string{
		"LIFECYCLE",
		"APPROVAL",
		"COMMIT",
		"REVIEWER",
		"COMMENT",
	}
}

// GetMappingListPullRequestActivitiesActivityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPullRequestActivitiesActivityTypeEnum(val string) (ListPullRequestActivitiesActivityTypeEnum, bool) {
	enum, ok := mappingListPullRequestActivitiesActivityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPullRequestActivitiesSortOrderEnum Enum with underlying type: string
type ListPullRequestActivitiesSortOrderEnum string

// Set of constants representing the allowable values for ListPullRequestActivitiesSortOrderEnum
const (
	ListPullRequestActivitiesSortOrderAsc  ListPullRequestActivitiesSortOrderEnum = "ASC"
	ListPullRequestActivitiesSortOrderDesc ListPullRequestActivitiesSortOrderEnum = "DESC"
)

var mappingListPullRequestActivitiesSortOrderEnum = map[string]ListPullRequestActivitiesSortOrderEnum{
	"ASC":  ListPullRequestActivitiesSortOrderAsc,
	"DESC": ListPullRequestActivitiesSortOrderDesc,
}

var mappingListPullRequestActivitiesSortOrderEnumLowerCase = map[string]ListPullRequestActivitiesSortOrderEnum{
	"asc":  ListPullRequestActivitiesSortOrderAsc,
	"desc": ListPullRequestActivitiesSortOrderDesc,
}

// GetListPullRequestActivitiesSortOrderEnumValues Enumerates the set of values for ListPullRequestActivitiesSortOrderEnum
func GetListPullRequestActivitiesSortOrderEnumValues() []ListPullRequestActivitiesSortOrderEnum {
	values := make([]ListPullRequestActivitiesSortOrderEnum, 0)
	for _, v := range mappingListPullRequestActivitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPullRequestActivitiesSortOrderEnumStringValues Enumerates the set of values in String for ListPullRequestActivitiesSortOrderEnum
func GetListPullRequestActivitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPullRequestActivitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPullRequestActivitiesSortOrderEnum(val string) (ListPullRequestActivitiesSortOrderEnum, bool) {
	enum, ok := mappingListPullRequestActivitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
