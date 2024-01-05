// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAwrDatabaseSnapshotsRequest wrapper for the ListAwrDatabaseSnapshots operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrDatabaseSnapshots.go.html to see an example of how to use ListAwrDatabaseSnapshotsRequest.
type ListAwrDatabaseSnapshotsRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabases
	AwrSourceDatabaseIdentifier *string `mandatory:"true" contributesTo:"query" name:"awrSourceDatabaseIdentifier"`

	// The optional single value query parameter to filter by database instance number.
	InstanceNumber *string `mandatory:"false" contributesTo:"query" name:"instanceNumber"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnapshotIdentifierGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnapshotIdentifierGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot Identifier.
	EndSnapshotIdentifierLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnapshotIdentifierLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the AWR snapshot summary data.
	SortBy ListAwrDatabaseSnapshotsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAwrDatabaseSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAwrDatabaseSnapshotsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAwrDatabaseSnapshotsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAwrDatabaseSnapshotsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAwrDatabaseSnapshotsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAwrDatabaseSnapshotsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAwrDatabaseSnapshotsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAwrDatabaseSnapshotsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAwrDatabaseSnapshotsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAwrDatabaseSnapshotsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAwrDatabaseSnapshotsResponse wrapper for the ListAwrDatabaseSnapshots operation
type ListAwrDatabaseSnapshotsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDatabaseSnapshotCollection instances
	AwrDatabaseSnapshotCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAwrDatabaseSnapshotsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAwrDatabaseSnapshotsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAwrDatabaseSnapshotsSortByEnum Enum with underlying type: string
type ListAwrDatabaseSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListAwrDatabaseSnapshotsSortByEnum
const (
	ListAwrDatabaseSnapshotsSortByTimeBegin  ListAwrDatabaseSnapshotsSortByEnum = "TIME_BEGIN"
	ListAwrDatabaseSnapshotsSortBySnapshotId ListAwrDatabaseSnapshotsSortByEnum = "SNAPSHOT_ID"
)

var mappingListAwrDatabaseSnapshotsSortByEnum = map[string]ListAwrDatabaseSnapshotsSortByEnum{
	"TIME_BEGIN":  ListAwrDatabaseSnapshotsSortByTimeBegin,
	"SNAPSHOT_ID": ListAwrDatabaseSnapshotsSortBySnapshotId,
}

var mappingListAwrDatabaseSnapshotsSortByEnumLowerCase = map[string]ListAwrDatabaseSnapshotsSortByEnum{
	"time_begin":  ListAwrDatabaseSnapshotsSortByTimeBegin,
	"snapshot_id": ListAwrDatabaseSnapshotsSortBySnapshotId,
}

// GetListAwrDatabaseSnapshotsSortByEnumValues Enumerates the set of values for ListAwrDatabaseSnapshotsSortByEnum
func GetListAwrDatabaseSnapshotsSortByEnumValues() []ListAwrDatabaseSnapshotsSortByEnum {
	values := make([]ListAwrDatabaseSnapshotsSortByEnum, 0)
	for _, v := range mappingListAwrDatabaseSnapshotsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrDatabaseSnapshotsSortByEnumStringValues Enumerates the set of values in String for ListAwrDatabaseSnapshotsSortByEnum
func GetListAwrDatabaseSnapshotsSortByEnumStringValues() []string {
	return []string{
		"TIME_BEGIN",
		"SNAPSHOT_ID",
	}
}

// GetMappingListAwrDatabaseSnapshotsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrDatabaseSnapshotsSortByEnum(val string) (ListAwrDatabaseSnapshotsSortByEnum, bool) {
	enum, ok := mappingListAwrDatabaseSnapshotsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAwrDatabaseSnapshotsSortOrderEnum Enum with underlying type: string
type ListAwrDatabaseSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListAwrDatabaseSnapshotsSortOrderEnum
const (
	ListAwrDatabaseSnapshotsSortOrderAsc  ListAwrDatabaseSnapshotsSortOrderEnum = "ASC"
	ListAwrDatabaseSnapshotsSortOrderDesc ListAwrDatabaseSnapshotsSortOrderEnum = "DESC"
)

var mappingListAwrDatabaseSnapshotsSortOrderEnum = map[string]ListAwrDatabaseSnapshotsSortOrderEnum{
	"ASC":  ListAwrDatabaseSnapshotsSortOrderAsc,
	"DESC": ListAwrDatabaseSnapshotsSortOrderDesc,
}

var mappingListAwrDatabaseSnapshotsSortOrderEnumLowerCase = map[string]ListAwrDatabaseSnapshotsSortOrderEnum{
	"asc":  ListAwrDatabaseSnapshotsSortOrderAsc,
	"desc": ListAwrDatabaseSnapshotsSortOrderDesc,
}

// GetListAwrDatabaseSnapshotsSortOrderEnumValues Enumerates the set of values for ListAwrDatabaseSnapshotsSortOrderEnum
func GetListAwrDatabaseSnapshotsSortOrderEnumValues() []ListAwrDatabaseSnapshotsSortOrderEnum {
	values := make([]ListAwrDatabaseSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListAwrDatabaseSnapshotsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrDatabaseSnapshotsSortOrderEnumStringValues Enumerates the set of values in String for ListAwrDatabaseSnapshotsSortOrderEnum
func GetListAwrDatabaseSnapshotsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAwrDatabaseSnapshotsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrDatabaseSnapshotsSortOrderEnum(val string) (ListAwrDatabaseSnapshotsSortOrderEnum, bool) {
	enum, ok := mappingListAwrDatabaseSnapshotsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
