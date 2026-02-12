// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPatchGroupPendingPatchesRequest wrapper for the ListPatchGroupPendingPatches operation
type ListPatchGroupPendingPatchesRequest struct {

	// Unique Fleet identifier.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// Unique FleetTarget identifier.
	TargetId *string `mandatory:"true" contributesTo:"path" name:"targetId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListPatchGroupPendingPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListPatchGroupPendingPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatchGroupPendingPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatchGroupPendingPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatchGroupPendingPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatchGroupPendingPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPatchGroupPendingPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPatchGroupPendingPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPatchGroupPendingPatchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPatchGroupPendingPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPatchGroupPendingPatchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPatchGroupPendingPatchesResponse wrapper for the ListPatchGroupPendingPatches operation
type ListPatchGroupPendingPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatchGroupPendingPatchesCollection instances
	PatchGroupPendingPatchesCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPatchGroupPendingPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatchGroupPendingPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatchGroupPendingPatchesSortOrderEnum Enum with underlying type: string
type ListPatchGroupPendingPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListPatchGroupPendingPatchesSortOrderEnum
const (
	ListPatchGroupPendingPatchesSortOrderAsc  ListPatchGroupPendingPatchesSortOrderEnum = "ASC"
	ListPatchGroupPendingPatchesSortOrderDesc ListPatchGroupPendingPatchesSortOrderEnum = "DESC"
)

var mappingListPatchGroupPendingPatchesSortOrderEnum = map[string]ListPatchGroupPendingPatchesSortOrderEnum{
	"ASC":  ListPatchGroupPendingPatchesSortOrderAsc,
	"DESC": ListPatchGroupPendingPatchesSortOrderDesc,
}

var mappingListPatchGroupPendingPatchesSortOrderEnumLowerCase = map[string]ListPatchGroupPendingPatchesSortOrderEnum{
	"asc":  ListPatchGroupPendingPatchesSortOrderAsc,
	"desc": ListPatchGroupPendingPatchesSortOrderDesc,
}

// GetListPatchGroupPendingPatchesSortOrderEnumValues Enumerates the set of values for ListPatchGroupPendingPatchesSortOrderEnum
func GetListPatchGroupPendingPatchesSortOrderEnumValues() []ListPatchGroupPendingPatchesSortOrderEnum {
	values := make([]ListPatchGroupPendingPatchesSortOrderEnum, 0)
	for _, v := range mappingListPatchGroupPendingPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchGroupPendingPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListPatchGroupPendingPatchesSortOrderEnum
func GetListPatchGroupPendingPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPatchGroupPendingPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchGroupPendingPatchesSortOrderEnum(val string) (ListPatchGroupPendingPatchesSortOrderEnum, bool) {
	enum, ok := mappingListPatchGroupPendingPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPatchGroupPendingPatchesSortByEnum Enum with underlying type: string
type ListPatchGroupPendingPatchesSortByEnum string

// Set of constants representing the allowable values for ListPatchGroupPendingPatchesSortByEnum
const (
	ListPatchGroupPendingPatchesSortByTimecreated ListPatchGroupPendingPatchesSortByEnum = "timeCreated"
	ListPatchGroupPendingPatchesSortByDisplayname ListPatchGroupPendingPatchesSortByEnum = "displayName"
)

var mappingListPatchGroupPendingPatchesSortByEnum = map[string]ListPatchGroupPendingPatchesSortByEnum{
	"timeCreated": ListPatchGroupPendingPatchesSortByTimecreated,
	"displayName": ListPatchGroupPendingPatchesSortByDisplayname,
}

var mappingListPatchGroupPendingPatchesSortByEnumLowerCase = map[string]ListPatchGroupPendingPatchesSortByEnum{
	"timecreated": ListPatchGroupPendingPatchesSortByTimecreated,
	"displayname": ListPatchGroupPendingPatchesSortByDisplayname,
}

// GetListPatchGroupPendingPatchesSortByEnumValues Enumerates the set of values for ListPatchGroupPendingPatchesSortByEnum
func GetListPatchGroupPendingPatchesSortByEnumValues() []ListPatchGroupPendingPatchesSortByEnum {
	values := make([]ListPatchGroupPendingPatchesSortByEnum, 0)
	for _, v := range mappingListPatchGroupPendingPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPatchGroupPendingPatchesSortByEnumStringValues Enumerates the set of values in String for ListPatchGroupPendingPatchesSortByEnum
func GetListPatchGroupPendingPatchesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPatchGroupPendingPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPatchGroupPendingPatchesSortByEnum(val string) (ListPatchGroupPendingPatchesSortByEnum, bool) {
	enum, ok := mappingListPatchGroupPendingPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
