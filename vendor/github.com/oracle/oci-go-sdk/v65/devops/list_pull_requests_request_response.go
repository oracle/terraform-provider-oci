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

// ListPullRequestsRequest wrapper for the ListPullRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListPullRequests.go.html to see an example of how to use ListPullRequestsRequest.
type ListPullRequestsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only pull requests that match the given lifecycle state.
	LifecycleState PullRequestLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only pull requests that match the given lifecycle state.
	LifecycleDetails PullRequestLifecycleDetailsEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetails" omitEmpty:"true"`

	// The OCID of the repository in which to list resources.
	RepositoryId *string `mandatory:"false" contributesTo:"query" name:"repositoryId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// PullRequest Created By
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// PullRequest Target Branch
	DestinationBranch *string `mandatory:"false" contributesTo:"query" name:"destinationBranch"`

	// PullRequest Source Branch.
	SourceBranch *string `mandatory:"false" contributesTo:"query" name:"sourceBranch"`

	// PullRequest Reviewer Id
	ReviewerPrincipalId *string `mandatory:"false" contributesTo:"query" name:"reviewerPrincipalId"`

	// PullRequest Source Repository Id
	SourceRepositoryId *string `mandatory:"false" contributesTo:"query" name:"sourceRepositoryId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListPullRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListPullRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPullRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPullRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPullRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPullRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPullRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPullRequestLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPullRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPullRequestLifecycleDetailsEnum(string(request.LifecycleDetails)); !ok && request.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", request.LifecycleDetails, strings.Join(GetPullRequestLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPullRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPullRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPullRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPullRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPullRequestsResponse wrapper for the ListPullRequests operation
type ListPullRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PullRequestCollection instances
	PullRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPullRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPullRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPullRequestsSortOrderEnum Enum with underlying type: string
type ListPullRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListPullRequestsSortOrderEnum
const (
	ListPullRequestsSortOrderAsc  ListPullRequestsSortOrderEnum = "ASC"
	ListPullRequestsSortOrderDesc ListPullRequestsSortOrderEnum = "DESC"
)

var mappingListPullRequestsSortOrderEnum = map[string]ListPullRequestsSortOrderEnum{
	"ASC":  ListPullRequestsSortOrderAsc,
	"DESC": ListPullRequestsSortOrderDesc,
}

var mappingListPullRequestsSortOrderEnumLowerCase = map[string]ListPullRequestsSortOrderEnum{
	"asc":  ListPullRequestsSortOrderAsc,
	"desc": ListPullRequestsSortOrderDesc,
}

// GetListPullRequestsSortOrderEnumValues Enumerates the set of values for ListPullRequestsSortOrderEnum
func GetListPullRequestsSortOrderEnumValues() []ListPullRequestsSortOrderEnum {
	values := make([]ListPullRequestsSortOrderEnum, 0)
	for _, v := range mappingListPullRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPullRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListPullRequestsSortOrderEnum
func GetListPullRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPullRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPullRequestsSortOrderEnum(val string) (ListPullRequestsSortOrderEnum, bool) {
	enum, ok := mappingListPullRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPullRequestsSortByEnum Enum with underlying type: string
type ListPullRequestsSortByEnum string

// Set of constants representing the allowable values for ListPullRequestsSortByEnum
const (
	ListPullRequestsSortByTimecreated ListPullRequestsSortByEnum = "timeCreated"
	ListPullRequestsSortByDisplayname ListPullRequestsSortByEnum = "displayName"
)

var mappingListPullRequestsSortByEnum = map[string]ListPullRequestsSortByEnum{
	"timeCreated": ListPullRequestsSortByTimecreated,
	"displayName": ListPullRequestsSortByDisplayname,
}

var mappingListPullRequestsSortByEnumLowerCase = map[string]ListPullRequestsSortByEnum{
	"timecreated": ListPullRequestsSortByTimecreated,
	"displayname": ListPullRequestsSortByDisplayname,
}

// GetListPullRequestsSortByEnumValues Enumerates the set of values for ListPullRequestsSortByEnum
func GetListPullRequestsSortByEnumValues() []ListPullRequestsSortByEnum {
	values := make([]ListPullRequestsSortByEnum, 0)
	for _, v := range mappingListPullRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPullRequestsSortByEnumStringValues Enumerates the set of values in String for ListPullRequestsSortByEnum
func GetListPullRequestsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPullRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPullRequestsSortByEnum(val string) (ListPullRequestsSortByEnum, bool) {
	enum, ok := mappingListPullRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
