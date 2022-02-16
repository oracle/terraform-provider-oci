// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAwrSnapshotsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAwrSnapshotsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAwrSnapshotsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAwrSnapshotsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAwrSnapshotsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListAwrSnapshotsSortOrderEnum = map[string]ListAwrSnapshotsSortOrderEnum{
	"ASC":  ListAwrSnapshotsSortOrderAsc,
	"DESC": ListAwrSnapshotsSortOrderDesc,
}

// GetListAwrSnapshotsSortOrderEnumValues Enumerates the set of values for ListAwrSnapshotsSortOrderEnum
func GetListAwrSnapshotsSortOrderEnumValues() []ListAwrSnapshotsSortOrderEnum {
	values := make([]ListAwrSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListAwrSnapshotsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrSnapshotsSortOrderEnumStringValues Enumerates the set of values in String for ListAwrSnapshotsSortOrderEnum
func GetListAwrSnapshotsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAwrSnapshotsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrSnapshotsSortOrderEnum(val string) (ListAwrSnapshotsSortOrderEnum, bool) {
	mappingListAwrSnapshotsSortOrderEnumIgnoreCase := make(map[string]ListAwrSnapshotsSortOrderEnum)
	for k, v := range mappingListAwrSnapshotsSortOrderEnum {
		mappingListAwrSnapshotsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAwrSnapshotsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAwrSnapshotsSortByEnum Enum with underlying type: string
type ListAwrSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListAwrSnapshotsSortByEnum
const (
	ListAwrSnapshotsSortByTimebegin  ListAwrSnapshotsSortByEnum = "timeBegin"
	ListAwrSnapshotsSortBySnapshotid ListAwrSnapshotsSortByEnum = "snapshotId"
)

var mappingListAwrSnapshotsSortByEnum = map[string]ListAwrSnapshotsSortByEnum{
	"timeBegin":  ListAwrSnapshotsSortByTimebegin,
	"snapshotId": ListAwrSnapshotsSortBySnapshotid,
}

// GetListAwrSnapshotsSortByEnumValues Enumerates the set of values for ListAwrSnapshotsSortByEnum
func GetListAwrSnapshotsSortByEnumValues() []ListAwrSnapshotsSortByEnum {
	values := make([]ListAwrSnapshotsSortByEnum, 0)
	for _, v := range mappingListAwrSnapshotsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrSnapshotsSortByEnumStringValues Enumerates the set of values in String for ListAwrSnapshotsSortByEnum
func GetListAwrSnapshotsSortByEnumStringValues() []string {
	return []string{
		"timeBegin",
		"snapshotId",
	}
}

// GetMappingListAwrSnapshotsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrSnapshotsSortByEnum(val string) (ListAwrSnapshotsSortByEnum, bool) {
	mappingListAwrSnapshotsSortByEnumIgnoreCase := make(map[string]ListAwrSnapshotsSortByEnum)
	for k, v := range mappingListAwrSnapshotsSortByEnum {
		mappingListAwrSnapshotsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAwrSnapshotsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
