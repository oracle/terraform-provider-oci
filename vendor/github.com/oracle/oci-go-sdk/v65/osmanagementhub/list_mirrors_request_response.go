// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMirrorsRequest wrapper for the ListMirrors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListMirrors.go.html to see an example of how to use ListMirrorsRequest.
type ListMirrorsRequest struct {

	// The OCID of the management station.
	ManagementStationId *string `mandatory:"true" contributesTo:"path" name:"managementStationId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMirrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListMirrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// List of Mirror state to filter by
	MirrorStates []MirrorStateEnum `contributesTo:"query" name:"mirrorStates" omitEmpty:"true" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMirrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMirrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMirrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMirrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMirrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMirrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMirrorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMirrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMirrorsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.MirrorStates {
		if _, ok := GetMappingMirrorStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MirrorStates: %s. Supported values are: %s.", val, strings.Join(GetMirrorStateEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMirrorsResponse wrapper for the ListMirrors operation
type ListMirrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MirrorsCollection instances
	MirrorsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMirrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMirrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMirrorsSortOrderEnum Enum with underlying type: string
type ListMirrorsSortOrderEnum string

// Set of constants representing the allowable values for ListMirrorsSortOrderEnum
const (
	ListMirrorsSortOrderAsc  ListMirrorsSortOrderEnum = "ASC"
	ListMirrorsSortOrderDesc ListMirrorsSortOrderEnum = "DESC"
)

var mappingListMirrorsSortOrderEnum = map[string]ListMirrorsSortOrderEnum{
	"ASC":  ListMirrorsSortOrderAsc,
	"DESC": ListMirrorsSortOrderDesc,
}

var mappingListMirrorsSortOrderEnumLowerCase = map[string]ListMirrorsSortOrderEnum{
	"asc":  ListMirrorsSortOrderAsc,
	"desc": ListMirrorsSortOrderDesc,
}

// GetListMirrorsSortOrderEnumValues Enumerates the set of values for ListMirrorsSortOrderEnum
func GetListMirrorsSortOrderEnumValues() []ListMirrorsSortOrderEnum {
	values := make([]ListMirrorsSortOrderEnum, 0)
	for _, v := range mappingListMirrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMirrorsSortOrderEnumStringValues Enumerates the set of values in String for ListMirrorsSortOrderEnum
func GetListMirrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMirrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMirrorsSortOrderEnum(val string) (ListMirrorsSortOrderEnum, bool) {
	enum, ok := mappingListMirrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMirrorsSortByEnum Enum with underlying type: string
type ListMirrorsSortByEnum string

// Set of constants representing the allowable values for ListMirrorsSortByEnum
const (
	ListMirrorsSortByTimecreated ListMirrorsSortByEnum = "timeCreated"
	ListMirrorsSortByDisplayname ListMirrorsSortByEnum = "displayName"
)

var mappingListMirrorsSortByEnum = map[string]ListMirrorsSortByEnum{
	"timeCreated": ListMirrorsSortByTimecreated,
	"displayName": ListMirrorsSortByDisplayname,
}

var mappingListMirrorsSortByEnumLowerCase = map[string]ListMirrorsSortByEnum{
	"timecreated": ListMirrorsSortByTimecreated,
	"displayname": ListMirrorsSortByDisplayname,
}

// GetListMirrorsSortByEnumValues Enumerates the set of values for ListMirrorsSortByEnum
func GetListMirrorsSortByEnumValues() []ListMirrorsSortByEnum {
	values := make([]ListMirrorsSortByEnum, 0)
	for _, v := range mappingListMirrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMirrorsSortByEnumStringValues Enumerates the set of values in String for ListMirrorsSortByEnum
func GetListMirrorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMirrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMirrorsSortByEnum(val string) (ListMirrorsSortByEnum, bool) {
	enum, ok := mappingListMirrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
