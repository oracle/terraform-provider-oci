// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNodeShapesRequest wrapper for the ListNodeShapes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managedkafka/ListNodeShapes.go.html to see an example of how to use ListNodeShapesRequest.
type ListNodeShapesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The name to filter on.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListNodeShapesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListNodeShapesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNodeShapesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNodeShapesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNodeShapesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNodeShapesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNodeShapesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNodeShapesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNodeShapesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNodeShapesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNodeShapesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNodeShapesResponse wrapper for the ListNodeShapes operation
type ListNodeShapesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NodeShapeCollection instances
	NodeShapeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNodeShapesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNodeShapesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNodeShapesSortOrderEnum Enum with underlying type: string
type ListNodeShapesSortOrderEnum string

// Set of constants representing the allowable values for ListNodeShapesSortOrderEnum
const (
	ListNodeShapesSortOrderAsc  ListNodeShapesSortOrderEnum = "ASC"
	ListNodeShapesSortOrderDesc ListNodeShapesSortOrderEnum = "DESC"
)

var mappingListNodeShapesSortOrderEnum = map[string]ListNodeShapesSortOrderEnum{
	"ASC":  ListNodeShapesSortOrderAsc,
	"DESC": ListNodeShapesSortOrderDesc,
}

var mappingListNodeShapesSortOrderEnumLowerCase = map[string]ListNodeShapesSortOrderEnum{
	"asc":  ListNodeShapesSortOrderAsc,
	"desc": ListNodeShapesSortOrderDesc,
}

// GetListNodeShapesSortOrderEnumValues Enumerates the set of values for ListNodeShapesSortOrderEnum
func GetListNodeShapesSortOrderEnumValues() []ListNodeShapesSortOrderEnum {
	values := make([]ListNodeShapesSortOrderEnum, 0)
	for _, v := range mappingListNodeShapesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNodeShapesSortOrderEnumStringValues Enumerates the set of values in String for ListNodeShapesSortOrderEnum
func GetListNodeShapesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNodeShapesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNodeShapesSortOrderEnum(val string) (ListNodeShapesSortOrderEnum, bool) {
	enum, ok := mappingListNodeShapesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNodeShapesSortByEnum Enum with underlying type: string
type ListNodeShapesSortByEnum string

// Set of constants representing the allowable values for ListNodeShapesSortByEnum
const (
	ListNodeShapesSortByTimecreated ListNodeShapesSortByEnum = "timeCreated"
	ListNodeShapesSortByDisplayname ListNodeShapesSortByEnum = "displayName"
)

var mappingListNodeShapesSortByEnum = map[string]ListNodeShapesSortByEnum{
	"timeCreated": ListNodeShapesSortByTimecreated,
	"displayName": ListNodeShapesSortByDisplayname,
}

var mappingListNodeShapesSortByEnumLowerCase = map[string]ListNodeShapesSortByEnum{
	"timecreated": ListNodeShapesSortByTimecreated,
	"displayname": ListNodeShapesSortByDisplayname,
}

// GetListNodeShapesSortByEnumValues Enumerates the set of values for ListNodeShapesSortByEnum
func GetListNodeShapesSortByEnumValues() []ListNodeShapesSortByEnum {
	values := make([]ListNodeShapesSortByEnum, 0)
	for _, v := range mappingListNodeShapesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNodeShapesSortByEnumStringValues Enumerates the set of values in String for ListNodeShapesSortByEnum
func GetListNodeShapesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNodeShapesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNodeShapesSortByEnum(val string) (ListNodeShapesSortByEnum, bool) {
	enum, ok := mappingListNodeShapesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
