// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmconfig

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataFilesRequest wrapper for the ListDataFiles operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/ListDataFiles.go.html to see an example of how to use ListDataFilesRequest.
type ListDataFilesRequest struct {

	// The APM Domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// The type of the data file.
	ApmType *string `mandatory:"false" contributesTo:"query" name:"apmType"`

	// A filter to return resources that match the specified name. Supports regular expressions to filter data files.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Return data files with time 'timeLastModified' before the specified time, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-19T22:47:12.613Z`
	TimeLastModifiedBefore *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastModifiedBefore"`

	// Return data files with the 'timeLastModified' after the specified time, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2020-02-19T22:47:12.613Z`
	TimeLastModifiedAfter *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastModifiedAfter"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The maximum number of results per page, or items to return in a paginated "List" call. For information on
	// how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The displayName sort order
	// is case-sensitive.
	SortOrder ListDataFilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one "sortBy" value. The default order for displayName, timeCreated
	// and timeUpdated is ascending. The displayName sort by is case-sensitive.
	SortBy ListDataFilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A string containing a JSON-encoded object with metadata related to the uploaded file or resource.
	// Example:
	//   {"fileName":"report.pdf","uploader":"jane.doe","category":"financial"}
	Metadata *string `mandatory:"false" contributesTo:"header" name:"metadata"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataFilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataFilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataFilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataFilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataFilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataFilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataFilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataFilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataFilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataFilesResponse wrapper for the ListDataFiles operation
type ListDataFilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataFileSummaryCollection instances
	DataFileSummaryCollection `presentIn:"body"`

	// The client request ID.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataFilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataFilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataFilesSortOrderEnum Enum with underlying type: string
type ListDataFilesSortOrderEnum string

// Set of constants representing the allowable values for ListDataFilesSortOrderEnum
const (
	ListDataFilesSortOrderAsc  ListDataFilesSortOrderEnum = "ASC"
	ListDataFilesSortOrderDesc ListDataFilesSortOrderEnum = "DESC"
)

var mappingListDataFilesSortOrderEnum = map[string]ListDataFilesSortOrderEnum{
	"ASC":  ListDataFilesSortOrderAsc,
	"DESC": ListDataFilesSortOrderDesc,
}

var mappingListDataFilesSortOrderEnumLowerCase = map[string]ListDataFilesSortOrderEnum{
	"asc":  ListDataFilesSortOrderAsc,
	"desc": ListDataFilesSortOrderDesc,
}

// GetListDataFilesSortOrderEnumValues Enumerates the set of values for ListDataFilesSortOrderEnum
func GetListDataFilesSortOrderEnumValues() []ListDataFilesSortOrderEnum {
	values := make([]ListDataFilesSortOrderEnum, 0)
	for _, v := range mappingListDataFilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataFilesSortOrderEnumStringValues Enumerates the set of values in String for ListDataFilesSortOrderEnum
func GetListDataFilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataFilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataFilesSortOrderEnum(val string) (ListDataFilesSortOrderEnum, bool) {
	enum, ok := mappingListDataFilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataFilesSortByEnum Enum with underlying type: string
type ListDataFilesSortByEnum string

// Set of constants representing the allowable values for ListDataFilesSortByEnum
const (
	ListDataFilesSortByDisplayname ListDataFilesSortByEnum = "displayName"
	ListDataFilesSortByTimecreated ListDataFilesSortByEnum = "timeCreated"
	ListDataFilesSortByTimeupdated ListDataFilesSortByEnum = "timeUpdated"
)

var mappingListDataFilesSortByEnum = map[string]ListDataFilesSortByEnum{
	"displayName": ListDataFilesSortByDisplayname,
	"timeCreated": ListDataFilesSortByTimecreated,
	"timeUpdated": ListDataFilesSortByTimeupdated,
}

var mappingListDataFilesSortByEnumLowerCase = map[string]ListDataFilesSortByEnum{
	"displayname": ListDataFilesSortByDisplayname,
	"timecreated": ListDataFilesSortByTimecreated,
	"timeupdated": ListDataFilesSortByTimeupdated,
}

// GetListDataFilesSortByEnumValues Enumerates the set of values for ListDataFilesSortByEnum
func GetListDataFilesSortByEnumValues() []ListDataFilesSortByEnum {
	values := make([]ListDataFilesSortByEnum, 0)
	for _, v := range mappingListDataFilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataFilesSortByEnumStringValues Enumerates the set of values in String for ListDataFilesSortByEnum
func GetListDataFilesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListDataFilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataFilesSortByEnum(val string) (ListDataFilesSortByEnum, bool) {
	enum, ok := mappingListDataFilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
