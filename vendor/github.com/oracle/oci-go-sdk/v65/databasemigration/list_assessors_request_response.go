// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAssessorsRequest wrapper for the ListAssessors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessors.go.html to see an example of how to use ListAssessorsRequest.
type ListAssessorsRequest struct {

	// The OCID of the Assessment
	AssessmentId *string `mandatory:"true" contributesTo:"path" name:"assessmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAssessorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAssessorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The lifecycle state of the Assessor.
	LifecycleState ListAssessorsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssessorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssessorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssessorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssessorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssessorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssessorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssessorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssessorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssessorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssessorsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAssessorsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssessorsResponse wrapper for the ListAssessors operation
type ListAssessorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssessorSummaryCollection instances
	AssessorSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssessorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssessorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssessorsSortByEnum Enum with underlying type: string
type ListAssessorsSortByEnum string

// Set of constants representing the allowable values for ListAssessorsSortByEnum
const (
	ListAssessorsSortByTimecreated ListAssessorsSortByEnum = "timeCreated"
	ListAssessorsSortByDisplayname ListAssessorsSortByEnum = "displayName"
)

var mappingListAssessorsSortByEnum = map[string]ListAssessorsSortByEnum{
	"timeCreated": ListAssessorsSortByTimecreated,
	"displayName": ListAssessorsSortByDisplayname,
}

var mappingListAssessorsSortByEnumLowerCase = map[string]ListAssessorsSortByEnum{
	"timecreated": ListAssessorsSortByTimecreated,
	"displayname": ListAssessorsSortByDisplayname,
}

// GetListAssessorsSortByEnumValues Enumerates the set of values for ListAssessorsSortByEnum
func GetListAssessorsSortByEnumValues() []ListAssessorsSortByEnum {
	values := make([]ListAssessorsSortByEnum, 0)
	for _, v := range mappingListAssessorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessorsSortByEnumStringValues Enumerates the set of values in String for ListAssessorsSortByEnum
func GetListAssessorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAssessorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessorsSortByEnum(val string) (ListAssessorsSortByEnum, bool) {
	enum, ok := mappingListAssessorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssessorsSortOrderEnum Enum with underlying type: string
type ListAssessorsSortOrderEnum string

// Set of constants representing the allowable values for ListAssessorsSortOrderEnum
const (
	ListAssessorsSortOrderAsc  ListAssessorsSortOrderEnum = "ASC"
	ListAssessorsSortOrderDesc ListAssessorsSortOrderEnum = "DESC"
)

var mappingListAssessorsSortOrderEnum = map[string]ListAssessorsSortOrderEnum{
	"ASC":  ListAssessorsSortOrderAsc,
	"DESC": ListAssessorsSortOrderDesc,
}

var mappingListAssessorsSortOrderEnumLowerCase = map[string]ListAssessorsSortOrderEnum{
	"asc":  ListAssessorsSortOrderAsc,
	"desc": ListAssessorsSortOrderDesc,
}

// GetListAssessorsSortOrderEnumValues Enumerates the set of values for ListAssessorsSortOrderEnum
func GetListAssessorsSortOrderEnumValues() []ListAssessorsSortOrderEnum {
	values := make([]ListAssessorsSortOrderEnum, 0)
	for _, v := range mappingListAssessorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessorsSortOrderEnumStringValues Enumerates the set of values in String for ListAssessorsSortOrderEnum
func GetListAssessorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssessorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessorsSortOrderEnum(val string) (ListAssessorsSortOrderEnum, bool) {
	enum, ok := mappingListAssessorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssessorsLifecycleStateEnum Enum with underlying type: string
type ListAssessorsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAssessorsLifecycleStateEnum
const (
	ListAssessorsLifecycleStateAccepted       ListAssessorsLifecycleStateEnum = "ACCEPTED"
	ListAssessorsLifecycleStateInProgress     ListAssessorsLifecycleStateEnum = "IN_PROGRESS"
	ListAssessorsLifecycleStateSucceeded      ListAssessorsLifecycleStateEnum = "SUCCEEDED"
	ListAssessorsLifecycleStateNeedsAttention ListAssessorsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListAssessorsLifecycleStateFailed         ListAssessorsLifecycleStateEnum = "FAILED"
)

var mappingListAssessorsLifecycleStateEnum = map[string]ListAssessorsLifecycleStateEnum{
	"ACCEPTED":        ListAssessorsLifecycleStateAccepted,
	"IN_PROGRESS":     ListAssessorsLifecycleStateInProgress,
	"SUCCEEDED":       ListAssessorsLifecycleStateSucceeded,
	"NEEDS_ATTENTION": ListAssessorsLifecycleStateNeedsAttention,
	"FAILED":          ListAssessorsLifecycleStateFailed,
}

var mappingListAssessorsLifecycleStateEnumLowerCase = map[string]ListAssessorsLifecycleStateEnum{
	"accepted":        ListAssessorsLifecycleStateAccepted,
	"in_progress":     ListAssessorsLifecycleStateInProgress,
	"succeeded":       ListAssessorsLifecycleStateSucceeded,
	"needs_attention": ListAssessorsLifecycleStateNeedsAttention,
	"failed":          ListAssessorsLifecycleStateFailed,
}

// GetListAssessorsLifecycleStateEnumValues Enumerates the set of values for ListAssessorsLifecycleStateEnum
func GetListAssessorsLifecycleStateEnumValues() []ListAssessorsLifecycleStateEnum {
	values := make([]ListAssessorsLifecycleStateEnum, 0)
	for _, v := range mappingListAssessorsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessorsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAssessorsLifecycleStateEnum
func GetListAssessorsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingListAssessorsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessorsLifecycleStateEnum(val string) (ListAssessorsLifecycleStateEnum, bool) {
	enum, ok := mappingListAssessorsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
