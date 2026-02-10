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

// ListAssessmentsRequest wrapper for the ListAssessments operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAssessments.go.html to see an example of how to use ListAssessmentsRequest.
type ListAssessmentsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	SortBy ListAssessmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAssessmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The lifecycle state of the Assessment.
	LifecycleState ListAssessmentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The lifecycle detailed status of the Migration.
	LifecycleDetails ListAssessmentsLifecycleDetailsEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetails" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssessmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssessmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssessmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssessmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssessmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssessmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssessmentsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssessmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssessmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssessmentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAssessmentsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssessmentsLifecycleDetailsEnum(string(request.LifecycleDetails)); !ok && request.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", request.LifecycleDetails, strings.Join(GetListAssessmentsLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssessmentsResponse wrapper for the ListAssessments operation
type ListAssessmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssessmentCollection instances
	AssessmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssessmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssessmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssessmentsSortByEnum Enum with underlying type: string
type ListAssessmentsSortByEnum string

// Set of constants representing the allowable values for ListAssessmentsSortByEnum
const (
	ListAssessmentsSortByTimecreated ListAssessmentsSortByEnum = "timeCreated"
	ListAssessmentsSortByDisplayname ListAssessmentsSortByEnum = "displayName"
)

var mappingListAssessmentsSortByEnum = map[string]ListAssessmentsSortByEnum{
	"timeCreated": ListAssessmentsSortByTimecreated,
	"displayName": ListAssessmentsSortByDisplayname,
}

var mappingListAssessmentsSortByEnumLowerCase = map[string]ListAssessmentsSortByEnum{
	"timecreated": ListAssessmentsSortByTimecreated,
	"displayname": ListAssessmentsSortByDisplayname,
}

// GetListAssessmentsSortByEnumValues Enumerates the set of values for ListAssessmentsSortByEnum
func GetListAssessmentsSortByEnumValues() []ListAssessmentsSortByEnum {
	values := make([]ListAssessmentsSortByEnum, 0)
	for _, v := range mappingListAssessmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessmentsSortByEnumStringValues Enumerates the set of values in String for ListAssessmentsSortByEnum
func GetListAssessmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAssessmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessmentsSortByEnum(val string) (ListAssessmentsSortByEnum, bool) {
	enum, ok := mappingListAssessmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssessmentsSortOrderEnum Enum with underlying type: string
type ListAssessmentsSortOrderEnum string

// Set of constants representing the allowable values for ListAssessmentsSortOrderEnum
const (
	ListAssessmentsSortOrderAsc  ListAssessmentsSortOrderEnum = "ASC"
	ListAssessmentsSortOrderDesc ListAssessmentsSortOrderEnum = "DESC"
)

var mappingListAssessmentsSortOrderEnum = map[string]ListAssessmentsSortOrderEnum{
	"ASC":  ListAssessmentsSortOrderAsc,
	"DESC": ListAssessmentsSortOrderDesc,
}

var mappingListAssessmentsSortOrderEnumLowerCase = map[string]ListAssessmentsSortOrderEnum{
	"asc":  ListAssessmentsSortOrderAsc,
	"desc": ListAssessmentsSortOrderDesc,
}

// GetListAssessmentsSortOrderEnumValues Enumerates the set of values for ListAssessmentsSortOrderEnum
func GetListAssessmentsSortOrderEnumValues() []ListAssessmentsSortOrderEnum {
	values := make([]ListAssessmentsSortOrderEnum, 0)
	for _, v := range mappingListAssessmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessmentsSortOrderEnumStringValues Enumerates the set of values in String for ListAssessmentsSortOrderEnum
func GetListAssessmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssessmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessmentsSortOrderEnum(val string) (ListAssessmentsSortOrderEnum, bool) {
	enum, ok := mappingListAssessmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssessmentsLifecycleStateEnum Enum with underlying type: string
type ListAssessmentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAssessmentsLifecycleStateEnum
const (
	ListAssessmentsLifecycleStateCreating       ListAssessmentsLifecycleStateEnum = "CREATING"
	ListAssessmentsLifecycleStateUpdating       ListAssessmentsLifecycleStateEnum = "UPDATING"
	ListAssessmentsLifecycleStateActive         ListAssessmentsLifecycleStateEnum = "ACTIVE"
	ListAssessmentsLifecycleStateSucceeded      ListAssessmentsLifecycleStateEnum = "SUCCEEDED"
	ListAssessmentsLifecycleStateInProgress     ListAssessmentsLifecycleStateEnum = "IN_PROGRESS"
	ListAssessmentsLifecycleStateNeedsAttention ListAssessmentsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListAssessmentsLifecycleStateDeleting       ListAssessmentsLifecycleStateEnum = "DELETING"
	ListAssessmentsLifecycleStateDeleted        ListAssessmentsLifecycleStateEnum = "DELETED"
	ListAssessmentsLifecycleStateFailed         ListAssessmentsLifecycleStateEnum = "FAILED"
)

var mappingListAssessmentsLifecycleStateEnum = map[string]ListAssessmentsLifecycleStateEnum{
	"CREATING":        ListAssessmentsLifecycleStateCreating,
	"UPDATING":        ListAssessmentsLifecycleStateUpdating,
	"ACTIVE":          ListAssessmentsLifecycleStateActive,
	"SUCCEEDED":       ListAssessmentsLifecycleStateSucceeded,
	"IN_PROGRESS":     ListAssessmentsLifecycleStateInProgress,
	"NEEDS_ATTENTION": ListAssessmentsLifecycleStateNeedsAttention,
	"DELETING":        ListAssessmentsLifecycleStateDeleting,
	"DELETED":         ListAssessmentsLifecycleStateDeleted,
	"FAILED":          ListAssessmentsLifecycleStateFailed,
}

var mappingListAssessmentsLifecycleStateEnumLowerCase = map[string]ListAssessmentsLifecycleStateEnum{
	"creating":        ListAssessmentsLifecycleStateCreating,
	"updating":        ListAssessmentsLifecycleStateUpdating,
	"active":          ListAssessmentsLifecycleStateActive,
	"succeeded":       ListAssessmentsLifecycleStateSucceeded,
	"in_progress":     ListAssessmentsLifecycleStateInProgress,
	"needs_attention": ListAssessmentsLifecycleStateNeedsAttention,
	"deleting":        ListAssessmentsLifecycleStateDeleting,
	"deleted":         ListAssessmentsLifecycleStateDeleted,
	"failed":          ListAssessmentsLifecycleStateFailed,
}

// GetListAssessmentsLifecycleStateEnumValues Enumerates the set of values for ListAssessmentsLifecycleStateEnum
func GetListAssessmentsLifecycleStateEnumValues() []ListAssessmentsLifecycleStateEnum {
	values := make([]ListAssessmentsLifecycleStateEnum, 0)
	for _, v := range mappingListAssessmentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessmentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAssessmentsLifecycleStateEnum
func GetListAssessmentsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"SUCCEEDED",
		"IN_PROGRESS",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListAssessmentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessmentsLifecycleStateEnum(val string) (ListAssessmentsLifecycleStateEnum, bool) {
	enum, ok := mappingListAssessmentsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssessmentsLifecycleDetailsEnum Enum with underlying type: string
type ListAssessmentsLifecycleDetailsEnum string

// Set of constants representing the allowable values for ListAssessmentsLifecycleDetailsEnum
const (
	ListAssessmentsLifecycleDetailsReady      ListAssessmentsLifecycleDetailsEnum = "READY"
	ListAssessmentsLifecycleDetailsAborting   ListAssessmentsLifecycleDetailsEnum = "ABORTING"
	ListAssessmentsLifecycleDetailsValidating ListAssessmentsLifecycleDetailsEnum = "VALIDATING"
	ListAssessmentsLifecycleDetailsValidated  ListAssessmentsLifecycleDetailsEnum = "VALIDATED"
	ListAssessmentsLifecycleDetailsAssessed   ListAssessmentsLifecycleDetailsEnum = "ASSESSED"
	ListAssessmentsLifecycleDetailsWaiting    ListAssessmentsLifecycleDetailsEnum = "WAITING"
	ListAssessmentsLifecycleDetailsMigrating  ListAssessmentsLifecycleDetailsEnum = "MIGRATING"
	ListAssessmentsLifecycleDetailsDone       ListAssessmentsLifecycleDetailsEnum = "DONE"
)

var mappingListAssessmentsLifecycleDetailsEnum = map[string]ListAssessmentsLifecycleDetailsEnum{
	"READY":      ListAssessmentsLifecycleDetailsReady,
	"ABORTING":   ListAssessmentsLifecycleDetailsAborting,
	"VALIDATING": ListAssessmentsLifecycleDetailsValidating,
	"VALIDATED":  ListAssessmentsLifecycleDetailsValidated,
	"ASSESSED":   ListAssessmentsLifecycleDetailsAssessed,
	"WAITING":    ListAssessmentsLifecycleDetailsWaiting,
	"MIGRATING":  ListAssessmentsLifecycleDetailsMigrating,
	"DONE":       ListAssessmentsLifecycleDetailsDone,
}

var mappingListAssessmentsLifecycleDetailsEnumLowerCase = map[string]ListAssessmentsLifecycleDetailsEnum{
	"ready":      ListAssessmentsLifecycleDetailsReady,
	"aborting":   ListAssessmentsLifecycleDetailsAborting,
	"validating": ListAssessmentsLifecycleDetailsValidating,
	"validated":  ListAssessmentsLifecycleDetailsValidated,
	"assessed":   ListAssessmentsLifecycleDetailsAssessed,
	"waiting":    ListAssessmentsLifecycleDetailsWaiting,
	"migrating":  ListAssessmentsLifecycleDetailsMigrating,
	"done":       ListAssessmentsLifecycleDetailsDone,
}

// GetListAssessmentsLifecycleDetailsEnumValues Enumerates the set of values for ListAssessmentsLifecycleDetailsEnum
func GetListAssessmentsLifecycleDetailsEnumValues() []ListAssessmentsLifecycleDetailsEnum {
	values := make([]ListAssessmentsLifecycleDetailsEnum, 0)
	for _, v := range mappingListAssessmentsLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssessmentsLifecycleDetailsEnumStringValues Enumerates the set of values in String for ListAssessmentsLifecycleDetailsEnum
func GetListAssessmentsLifecycleDetailsEnumStringValues() []string {
	return []string{
		"READY",
		"ABORTING",
		"VALIDATING",
		"VALIDATED",
		"ASSESSED",
		"WAITING",
		"MIGRATING",
		"DONE",
	}
}

// GetMappingListAssessmentsLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssessmentsLifecycleDetailsEnum(val string) (ListAssessmentsLifecycleDetailsEnum, bool) {
	enum, ok := mappingListAssessmentsLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
