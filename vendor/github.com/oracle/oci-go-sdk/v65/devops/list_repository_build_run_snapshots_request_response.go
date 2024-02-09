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

// ListRepositoryBuildRunSnapshotsRequest wrapper for the ListRepositoryBuildRunSnapshots operation
type ListRepositoryBuildRunSnapshotsRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

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
	SortOrder ListRepositoryBuildRunSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for name is ascending. If no value is specified time created is default.
	SortBy ListRepositoryBuildRunSnapshotsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRepositoryBuildRunSnapshotsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRepositoryBuildRunSnapshotsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRepositoryBuildRunSnapshotsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRepositoryBuildRunSnapshotsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRepositoryBuildRunSnapshotsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRepositoryBuildRunSnapshotsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRepositoryBuildRunSnapshotsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoryBuildRunSnapshotsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRepositoryBuildRunSnapshotsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRepositoryBuildRunSnapshotsResponse wrapper for the ListRepositoryBuildRunSnapshots operation
type ListRepositoryBuildRunSnapshotsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BuildRunSnapshotCollection instances
	BuildRunSnapshotCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRepositoryBuildRunSnapshotsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRepositoryBuildRunSnapshotsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRepositoryBuildRunSnapshotsSortOrderEnum Enum with underlying type: string
type ListRepositoryBuildRunSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListRepositoryBuildRunSnapshotsSortOrderEnum
const (
	ListRepositoryBuildRunSnapshotsSortOrderAsc  ListRepositoryBuildRunSnapshotsSortOrderEnum = "ASC"
	ListRepositoryBuildRunSnapshotsSortOrderDesc ListRepositoryBuildRunSnapshotsSortOrderEnum = "DESC"
)

var mappingListRepositoryBuildRunSnapshotsSortOrderEnum = map[string]ListRepositoryBuildRunSnapshotsSortOrderEnum{
	"ASC":  ListRepositoryBuildRunSnapshotsSortOrderAsc,
	"DESC": ListRepositoryBuildRunSnapshotsSortOrderDesc,
}

var mappingListRepositoryBuildRunSnapshotsSortOrderEnumLowerCase = map[string]ListRepositoryBuildRunSnapshotsSortOrderEnum{
	"asc":  ListRepositoryBuildRunSnapshotsSortOrderAsc,
	"desc": ListRepositoryBuildRunSnapshotsSortOrderDesc,
}

// GetListRepositoryBuildRunSnapshotsSortOrderEnumValues Enumerates the set of values for ListRepositoryBuildRunSnapshotsSortOrderEnum
func GetListRepositoryBuildRunSnapshotsSortOrderEnumValues() []ListRepositoryBuildRunSnapshotsSortOrderEnum {
	values := make([]ListRepositoryBuildRunSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListRepositoryBuildRunSnapshotsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryBuildRunSnapshotsSortOrderEnumStringValues Enumerates the set of values in String for ListRepositoryBuildRunSnapshotsSortOrderEnum
func GetListRepositoryBuildRunSnapshotsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRepositoryBuildRunSnapshotsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryBuildRunSnapshotsSortOrderEnum(val string) (ListRepositoryBuildRunSnapshotsSortOrderEnum, bool) {
	enum, ok := mappingListRepositoryBuildRunSnapshotsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRepositoryBuildRunSnapshotsSortByEnum Enum with underlying type: string
type ListRepositoryBuildRunSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListRepositoryBuildRunSnapshotsSortByEnum
const (
	ListRepositoryBuildRunSnapshotsSortByTimecreated ListRepositoryBuildRunSnapshotsSortByEnum = "timeCreated"
	ListRepositoryBuildRunSnapshotsSortByName        ListRepositoryBuildRunSnapshotsSortByEnum = "name"
)

var mappingListRepositoryBuildRunSnapshotsSortByEnum = map[string]ListRepositoryBuildRunSnapshotsSortByEnum{
	"timeCreated": ListRepositoryBuildRunSnapshotsSortByTimecreated,
	"name":        ListRepositoryBuildRunSnapshotsSortByName,
}

var mappingListRepositoryBuildRunSnapshotsSortByEnumLowerCase = map[string]ListRepositoryBuildRunSnapshotsSortByEnum{
	"timecreated": ListRepositoryBuildRunSnapshotsSortByTimecreated,
	"name":        ListRepositoryBuildRunSnapshotsSortByName,
}

// GetListRepositoryBuildRunSnapshotsSortByEnumValues Enumerates the set of values for ListRepositoryBuildRunSnapshotsSortByEnum
func GetListRepositoryBuildRunSnapshotsSortByEnumValues() []ListRepositoryBuildRunSnapshotsSortByEnum {
	values := make([]ListRepositoryBuildRunSnapshotsSortByEnum, 0)
	for _, v := range mappingListRepositoryBuildRunSnapshotsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoryBuildRunSnapshotsSortByEnumStringValues Enumerates the set of values in String for ListRepositoryBuildRunSnapshotsSortByEnum
func GetListRepositoryBuildRunSnapshotsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListRepositoryBuildRunSnapshotsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoryBuildRunSnapshotsSortByEnum(val string) (ListRepositoryBuildRunSnapshotsSortByEnum, bool) {
	enum, ok := mappingListRepositoryBuildRunSnapshotsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
