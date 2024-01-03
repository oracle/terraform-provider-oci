// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbnodeSnapshotsRequest wrapper for the ListDbnodeSnapshots operation
type ListDbnodeSnapshotsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListDbnodeSnapshotsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given. The match is not case sensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDbnodeSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only Exadata Database Snapshots that match the given lifecycle state exactly.
	LifecycleState DbnodeSnapshotLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only Exadata Database Node Snapshots that match the given exaDB VM cluster.
	ExadbVmClusterId *string `mandatory:"false" contributesTo:"query" name:"exadbVmClusterId"`

	// A filter to return only Exadata Database Snapshots that match the given database node.
	SourceDbnodeId *string `mandatory:"false" contributesTo:"query" name:"sourceDbnodeId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbnodeSnapshotsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbnodeSnapshotsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbnodeSnapshotsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbnodeSnapshotsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbnodeSnapshotsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbnodeSnapshotsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbnodeSnapshotsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbnodeSnapshotsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbnodeSnapshotsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbnodeSnapshotLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbnodeSnapshotLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbnodeSnapshotsResponse wrapper for the ListDbnodeSnapshots operation
type ListDbnodeSnapshotsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbnodeSnapshotSummary instances
	Items []DbnodeSnapshotSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbnodeSnapshotsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbnodeSnapshotsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbnodeSnapshotsSortByEnum Enum with underlying type: string
type ListDbnodeSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListDbnodeSnapshotsSortByEnum
const (
	ListDbnodeSnapshotsSortByTimecreated ListDbnodeSnapshotsSortByEnum = "TIMECREATED"
	ListDbnodeSnapshotsSortByName        ListDbnodeSnapshotsSortByEnum = "NAME"
)

var mappingListDbnodeSnapshotsSortByEnum = map[string]ListDbnodeSnapshotsSortByEnum{
	"TIMECREATED": ListDbnodeSnapshotsSortByTimecreated,
	"NAME":        ListDbnodeSnapshotsSortByName,
}

var mappingListDbnodeSnapshotsSortByEnumLowerCase = map[string]ListDbnodeSnapshotsSortByEnum{
	"timecreated": ListDbnodeSnapshotsSortByTimecreated,
	"name":        ListDbnodeSnapshotsSortByName,
}

// GetListDbnodeSnapshotsSortByEnumValues Enumerates the set of values for ListDbnodeSnapshotsSortByEnum
func GetListDbnodeSnapshotsSortByEnumValues() []ListDbnodeSnapshotsSortByEnum {
	values := make([]ListDbnodeSnapshotsSortByEnum, 0)
	for _, v := range mappingListDbnodeSnapshotsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbnodeSnapshotsSortByEnumStringValues Enumerates the set of values in String for ListDbnodeSnapshotsSortByEnum
func GetListDbnodeSnapshotsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListDbnodeSnapshotsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbnodeSnapshotsSortByEnum(val string) (ListDbnodeSnapshotsSortByEnum, bool) {
	enum, ok := mappingListDbnodeSnapshotsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbnodeSnapshotsSortOrderEnum Enum with underlying type: string
type ListDbnodeSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListDbnodeSnapshotsSortOrderEnum
const (
	ListDbnodeSnapshotsSortOrderAsc  ListDbnodeSnapshotsSortOrderEnum = "ASC"
	ListDbnodeSnapshotsSortOrderDesc ListDbnodeSnapshotsSortOrderEnum = "DESC"
)

var mappingListDbnodeSnapshotsSortOrderEnum = map[string]ListDbnodeSnapshotsSortOrderEnum{
	"ASC":  ListDbnodeSnapshotsSortOrderAsc,
	"DESC": ListDbnodeSnapshotsSortOrderDesc,
}

var mappingListDbnodeSnapshotsSortOrderEnumLowerCase = map[string]ListDbnodeSnapshotsSortOrderEnum{
	"asc":  ListDbnodeSnapshotsSortOrderAsc,
	"desc": ListDbnodeSnapshotsSortOrderDesc,
}

// GetListDbnodeSnapshotsSortOrderEnumValues Enumerates the set of values for ListDbnodeSnapshotsSortOrderEnum
func GetListDbnodeSnapshotsSortOrderEnumValues() []ListDbnodeSnapshotsSortOrderEnum {
	values := make([]ListDbnodeSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListDbnodeSnapshotsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbnodeSnapshotsSortOrderEnumStringValues Enumerates the set of values in String for ListDbnodeSnapshotsSortOrderEnum
func GetListDbnodeSnapshotsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbnodeSnapshotsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbnodeSnapshotsSortOrderEnum(val string) (ListDbnodeSnapshotsSortOrderEnum, bool) {
	enum, ok := mappingListDbnodeSnapshotsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
