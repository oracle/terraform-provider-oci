// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListAwrSnapshotsRequest wrapper for the ListAwrSnapshots operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrSnapshots.go.html to see an example of how to use ListAwrSnapshotsRequest.
type ListAwrSnapshotsRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// AWR source database identifier.
	AwrSourceDatabaseIdentifier *string `mandatory:"true" contributesTo:"query" name:"awrSourceDatabaseIdentifier"`

	// The optional greater than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAwrSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The option to sort the AWR snapshot summary data. Default sort is by timeBegin.
	SortBy ListAwrSnapshotsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAwrSnapshotsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAwrSnapshotsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAwrSnapshotsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAwrSnapshotsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAwrSnapshotsResponse wrapper for the ListAwrSnapshots operation
type ListAwrSnapshotsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrSnapshotCollection instances
	AwrSnapshotCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAwrSnapshotsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAwrSnapshotsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAwrSnapshotsSortOrderEnum Enum with underlying type: string
type ListAwrSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListAwrSnapshotsSortOrderEnum
const (
	ListAwrSnapshotsSortOrderAsc  ListAwrSnapshotsSortOrderEnum = "ASC"
	ListAwrSnapshotsSortOrderDesc ListAwrSnapshotsSortOrderEnum = "DESC"
)

var mappingListAwrSnapshotsSortOrder = map[string]ListAwrSnapshotsSortOrderEnum{
	"ASC":  ListAwrSnapshotsSortOrderAsc,
	"DESC": ListAwrSnapshotsSortOrderDesc,
}

// GetListAwrSnapshotsSortOrderEnumValues Enumerates the set of values for ListAwrSnapshotsSortOrderEnum
func GetListAwrSnapshotsSortOrderEnumValues() []ListAwrSnapshotsSortOrderEnum {
	values := make([]ListAwrSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListAwrSnapshotsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAwrSnapshotsSortByEnum Enum with underlying type: string
type ListAwrSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListAwrSnapshotsSortByEnum
const (
	ListAwrSnapshotsSortByTimebegin  ListAwrSnapshotsSortByEnum = "timeBegin"
	ListAwrSnapshotsSortBySnapshotid ListAwrSnapshotsSortByEnum = "snapshotId"
)

var mappingListAwrSnapshotsSortBy = map[string]ListAwrSnapshotsSortByEnum{
	"timeBegin":  ListAwrSnapshotsSortByTimebegin,
	"snapshotId": ListAwrSnapshotsSortBySnapshotid,
}

// GetListAwrSnapshotsSortByEnumValues Enumerates the set of values for ListAwrSnapshotsSortByEnum
func GetListAwrSnapshotsSortByEnumValues() []ListAwrSnapshotsSortByEnum {
	values := make([]ListAwrSnapshotsSortByEnum, 0)
	for _, v := range mappingListAwrSnapshotsSortBy {
		values = append(values, v)
	}
	return values
}
