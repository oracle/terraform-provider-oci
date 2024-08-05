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

// ListBuildRunSnapshotsRequest wrapper for the ListBuildRunSnapshots operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListBuildRunSnapshots.go.html to see an example of how to use ListBuildRunSnapshotsRequest.
type ListBuildRunSnapshotsRequest struct {

	// unique PullRequest identifier
	PullRequestId *string `mandatory:"true" contributesTo:"path" name:"pullRequestId"`

	// Unique build pipeline identifier.
	PipelineId *string `mandatory:"false" contributesTo:"query" name:"pipelineId"`

	// Unique build run identifier.
	BuildRunId *string `mandatory:"false" contributesTo:"query" name:"buildRunId"`

	// Commit ID in a repository.
	CommitId *string `mandatory:"false" contributesTo:"query" name:"commitId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListBuildRunSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for name is ascending. If no value is specified time created is default.
	SortBy ListBuildRunSnapshotsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBuildRunSnapshotsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBuildRunSnapshotsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBuildRunSnapshotsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBuildRunSnapshotsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBuildRunSnapshotsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBuildRunSnapshotsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBuildRunSnapshotsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBuildRunSnapshotsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBuildRunSnapshotsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBuildRunSnapshotsResponse wrapper for the ListBuildRunSnapshots operation
type ListBuildRunSnapshotsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BuildRunSnapshotCollection instances
	BuildRunSnapshotCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBuildRunSnapshotsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBuildRunSnapshotsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBuildRunSnapshotsSortOrderEnum Enum with underlying type: string
type ListBuildRunSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListBuildRunSnapshotsSortOrderEnum
const (
	ListBuildRunSnapshotsSortOrderAsc  ListBuildRunSnapshotsSortOrderEnum = "ASC"
	ListBuildRunSnapshotsSortOrderDesc ListBuildRunSnapshotsSortOrderEnum = "DESC"
)

var mappingListBuildRunSnapshotsSortOrderEnum = map[string]ListBuildRunSnapshotsSortOrderEnum{
	"ASC":  ListBuildRunSnapshotsSortOrderAsc,
	"DESC": ListBuildRunSnapshotsSortOrderDesc,
}

var mappingListBuildRunSnapshotsSortOrderEnumLowerCase = map[string]ListBuildRunSnapshotsSortOrderEnum{
	"asc":  ListBuildRunSnapshotsSortOrderAsc,
	"desc": ListBuildRunSnapshotsSortOrderDesc,
}

// GetListBuildRunSnapshotsSortOrderEnumValues Enumerates the set of values for ListBuildRunSnapshotsSortOrderEnum
func GetListBuildRunSnapshotsSortOrderEnumValues() []ListBuildRunSnapshotsSortOrderEnum {
	values := make([]ListBuildRunSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListBuildRunSnapshotsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBuildRunSnapshotsSortOrderEnumStringValues Enumerates the set of values in String for ListBuildRunSnapshotsSortOrderEnum
func GetListBuildRunSnapshotsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBuildRunSnapshotsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBuildRunSnapshotsSortOrderEnum(val string) (ListBuildRunSnapshotsSortOrderEnum, bool) {
	enum, ok := mappingListBuildRunSnapshotsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBuildRunSnapshotsSortByEnum Enum with underlying type: string
type ListBuildRunSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListBuildRunSnapshotsSortByEnum
const (
	ListBuildRunSnapshotsSortByTimecreated ListBuildRunSnapshotsSortByEnum = "timeCreated"
	ListBuildRunSnapshotsSortByName        ListBuildRunSnapshotsSortByEnum = "name"
)

var mappingListBuildRunSnapshotsSortByEnum = map[string]ListBuildRunSnapshotsSortByEnum{
	"timeCreated": ListBuildRunSnapshotsSortByTimecreated,
	"name":        ListBuildRunSnapshotsSortByName,
}

var mappingListBuildRunSnapshotsSortByEnumLowerCase = map[string]ListBuildRunSnapshotsSortByEnum{
	"timecreated": ListBuildRunSnapshotsSortByTimecreated,
	"name":        ListBuildRunSnapshotsSortByName,
}

// GetListBuildRunSnapshotsSortByEnumValues Enumerates the set of values for ListBuildRunSnapshotsSortByEnum
func GetListBuildRunSnapshotsSortByEnumValues() []ListBuildRunSnapshotsSortByEnum {
	values := make([]ListBuildRunSnapshotsSortByEnum, 0)
	for _, v := range mappingListBuildRunSnapshotsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBuildRunSnapshotsSortByEnumStringValues Enumerates the set of values in String for ListBuildRunSnapshotsSortByEnum
func GetListBuildRunSnapshotsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListBuildRunSnapshotsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBuildRunSnapshotsSortByEnum(val string) (ListBuildRunSnapshotsSortByEnum, bool) {
	enum, ok := mappingListBuildRunSnapshotsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
