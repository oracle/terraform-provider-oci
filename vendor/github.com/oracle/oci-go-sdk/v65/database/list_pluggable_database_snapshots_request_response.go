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

// ListPluggableDatabaseSnapshotsRequest wrapper for the ListPluggableDatabaseSnapshots operation
type ListPluggableDatabaseSnapshotsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListPluggableDatabaseSnapshotsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given. The match is not case sensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPluggableDatabaseSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only Exadata Pluggable Database Snapshots that match the given lifecycle state exactly.
	LifecycleState PluggableDatabaseSnapshotLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only Exadata Pluggable Database Snapshots that match the given exadb VM cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ExadbVmClusterId *string `mandatory:"false" contributesTo:"query" name:"exadbVmClusterId"`

	// A filter to return only Exadata Pluggable Database Snapshots that match the given database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	PluggableDatabaseId *string `mandatory:"false" contributesTo:"query" name:"pluggableDatabaseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPluggableDatabaseSnapshotsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPluggableDatabaseSnapshotsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPluggableDatabaseSnapshotsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPluggableDatabaseSnapshotsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPluggableDatabaseSnapshotsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPluggableDatabaseSnapshotsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPluggableDatabaseSnapshotsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPluggableDatabaseSnapshotsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPluggableDatabaseSnapshotsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPluggableDatabaseSnapshotLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPluggableDatabaseSnapshotLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPluggableDatabaseSnapshotsResponse wrapper for the ListPluggableDatabaseSnapshots operation
type ListPluggableDatabaseSnapshotsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PluggableDatabaseSnapshotSummary instances
	Items []PluggableDatabaseSnapshotSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPluggableDatabaseSnapshotsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPluggableDatabaseSnapshotsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPluggableDatabaseSnapshotsSortByEnum Enum with underlying type: string
type ListPluggableDatabaseSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListPluggableDatabaseSnapshotsSortByEnum
const (
	ListPluggableDatabaseSnapshotsSortByTimecreated ListPluggableDatabaseSnapshotsSortByEnum = "TIMECREATED"
	ListPluggableDatabaseSnapshotsSortByName        ListPluggableDatabaseSnapshotsSortByEnum = "NAME"
)

var mappingListPluggableDatabaseSnapshotsSortByEnum = map[string]ListPluggableDatabaseSnapshotsSortByEnum{
	"TIMECREATED": ListPluggableDatabaseSnapshotsSortByTimecreated,
	"NAME":        ListPluggableDatabaseSnapshotsSortByName,
}

var mappingListPluggableDatabaseSnapshotsSortByEnumLowerCase = map[string]ListPluggableDatabaseSnapshotsSortByEnum{
	"timecreated": ListPluggableDatabaseSnapshotsSortByTimecreated,
	"name":        ListPluggableDatabaseSnapshotsSortByName,
}

// GetListPluggableDatabaseSnapshotsSortByEnumValues Enumerates the set of values for ListPluggableDatabaseSnapshotsSortByEnum
func GetListPluggableDatabaseSnapshotsSortByEnumValues() []ListPluggableDatabaseSnapshotsSortByEnum {
	values := make([]ListPluggableDatabaseSnapshotsSortByEnum, 0)
	for _, v := range mappingListPluggableDatabaseSnapshotsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPluggableDatabaseSnapshotsSortByEnumStringValues Enumerates the set of values in String for ListPluggableDatabaseSnapshotsSortByEnum
func GetListPluggableDatabaseSnapshotsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListPluggableDatabaseSnapshotsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPluggableDatabaseSnapshotsSortByEnum(val string) (ListPluggableDatabaseSnapshotsSortByEnum, bool) {
	enum, ok := mappingListPluggableDatabaseSnapshotsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPluggableDatabaseSnapshotsSortOrderEnum Enum with underlying type: string
type ListPluggableDatabaseSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListPluggableDatabaseSnapshotsSortOrderEnum
const (
	ListPluggableDatabaseSnapshotsSortOrderAsc  ListPluggableDatabaseSnapshotsSortOrderEnum = "ASC"
	ListPluggableDatabaseSnapshotsSortOrderDesc ListPluggableDatabaseSnapshotsSortOrderEnum = "DESC"
)

var mappingListPluggableDatabaseSnapshotsSortOrderEnum = map[string]ListPluggableDatabaseSnapshotsSortOrderEnum{
	"ASC":  ListPluggableDatabaseSnapshotsSortOrderAsc,
	"DESC": ListPluggableDatabaseSnapshotsSortOrderDesc,
}

var mappingListPluggableDatabaseSnapshotsSortOrderEnumLowerCase = map[string]ListPluggableDatabaseSnapshotsSortOrderEnum{
	"asc":  ListPluggableDatabaseSnapshotsSortOrderAsc,
	"desc": ListPluggableDatabaseSnapshotsSortOrderDesc,
}

// GetListPluggableDatabaseSnapshotsSortOrderEnumValues Enumerates the set of values for ListPluggableDatabaseSnapshotsSortOrderEnum
func GetListPluggableDatabaseSnapshotsSortOrderEnumValues() []ListPluggableDatabaseSnapshotsSortOrderEnum {
	values := make([]ListPluggableDatabaseSnapshotsSortOrderEnum, 0)
	for _, v := range mappingListPluggableDatabaseSnapshotsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPluggableDatabaseSnapshotsSortOrderEnumStringValues Enumerates the set of values in String for ListPluggableDatabaseSnapshotsSortOrderEnum
func GetListPluggableDatabaseSnapshotsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPluggableDatabaseSnapshotsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPluggableDatabaseSnapshotsSortOrderEnum(val string) (ListPluggableDatabaseSnapshotsSortOrderEnum, bool) {
	enum, ok := mappingListPluggableDatabaseSnapshotsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
