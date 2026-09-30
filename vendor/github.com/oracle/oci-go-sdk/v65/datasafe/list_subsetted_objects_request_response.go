// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubsettedObjectsRequest wrapper for the ListSubsettedObjects operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettedObjects.go.html to see an example of how to use ListSubsettedObjectsRequest.
type ListSubsettedObjectsRequest struct {

	// The OCID of the subsetting report.
	SubsettingReportId *string `mandatory:"true" contributesTo:"path" name:"subsettingReportId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettedObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for all the fields is ascending.
	SortBy ListSubsettedObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettedObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettedObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettedObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettedObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettedObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettedObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettedObjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettedObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettedObjectsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettedObjectsResponse wrapper for the ListSubsettedObjects operation
type ListSubsettedObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettedObjectCollection instances
	SubsettedObjectCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettedObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettedObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettedObjectsSortOrderEnum Enum with underlying type: string
type ListSubsettedObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettedObjectsSortOrderEnum
const (
	ListSubsettedObjectsSortOrderAsc  ListSubsettedObjectsSortOrderEnum = "ASC"
	ListSubsettedObjectsSortOrderDesc ListSubsettedObjectsSortOrderEnum = "DESC"
)

var mappingListSubsettedObjectsSortOrderEnum = map[string]ListSubsettedObjectsSortOrderEnum{
	"ASC":  ListSubsettedObjectsSortOrderAsc,
	"DESC": ListSubsettedObjectsSortOrderDesc,
}

var mappingListSubsettedObjectsSortOrderEnumLowerCase = map[string]ListSubsettedObjectsSortOrderEnum{
	"asc":  ListSubsettedObjectsSortOrderAsc,
	"desc": ListSubsettedObjectsSortOrderDesc,
}

// GetListSubsettedObjectsSortOrderEnumValues Enumerates the set of values for ListSubsettedObjectsSortOrderEnum
func GetListSubsettedObjectsSortOrderEnumValues() []ListSubsettedObjectsSortOrderEnum {
	values := make([]ListSubsettedObjectsSortOrderEnum, 0)
	for _, v := range mappingListSubsettedObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettedObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettedObjectsSortOrderEnum
func GetListSubsettedObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettedObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettedObjectsSortOrderEnum(val string) (ListSubsettedObjectsSortOrderEnum, bool) {
	enum, ok := mappingListSubsettedObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettedObjectsSortByEnum Enum with underlying type: string
type ListSubsettedObjectsSortByEnum string

// Set of constants representing the allowable values for ListSubsettedObjectsSortByEnum
const (
	ListSubsettedObjectsSortBySchemaname ListSubsettedObjectsSortByEnum = "schemaName"
	ListSubsettedObjectsSortByObjectname ListSubsettedObjectsSortByEnum = "objectName"
)

var mappingListSubsettedObjectsSortByEnum = map[string]ListSubsettedObjectsSortByEnum{
	"schemaName": ListSubsettedObjectsSortBySchemaname,
	"objectName": ListSubsettedObjectsSortByObjectname,
}

var mappingListSubsettedObjectsSortByEnumLowerCase = map[string]ListSubsettedObjectsSortByEnum{
	"schemaname": ListSubsettedObjectsSortBySchemaname,
	"objectname": ListSubsettedObjectsSortByObjectname,
}

// GetListSubsettedObjectsSortByEnumValues Enumerates the set of values for ListSubsettedObjectsSortByEnum
func GetListSubsettedObjectsSortByEnumValues() []ListSubsettedObjectsSortByEnum {
	values := make([]ListSubsettedObjectsSortByEnum, 0)
	for _, v := range mappingListSubsettedObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettedObjectsSortByEnumStringValues Enumerates the set of values in String for ListSubsettedObjectsSortByEnum
func GetListSubsettedObjectsSortByEnumStringValues() []string {
	return []string{
		"schemaName",
		"objectName",
	}
}

// GetMappingListSubsettedObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettedObjectsSortByEnum(val string) (ListSubsettedObjectsSortByEnum, bool) {
	enum, ok := mappingListSubsettedObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
