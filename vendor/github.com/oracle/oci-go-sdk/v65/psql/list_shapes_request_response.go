// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListShapesRequest wrapper for the ListShapes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psql/ListShapes.go.html to see an example of how to use ListShapesRequest.
type ListShapesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return the feature by the shape name.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListShapesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListShapesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListShapesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListShapesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListShapesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListShapesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListShapesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListShapesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListShapesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListShapesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListShapesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListShapesResponse wrapper for the ListShapes operation
type ListShapesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ShapeCollection instances
	ShapeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListShapesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListShapesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListShapesSortOrderEnum Enum with underlying type: string
type ListShapesSortOrderEnum string

// Set of constants representing the allowable values for ListShapesSortOrderEnum
const (
	ListShapesSortOrderAsc  ListShapesSortOrderEnum = "ASC"
	ListShapesSortOrderDesc ListShapesSortOrderEnum = "DESC"
)

var mappingListShapesSortOrderEnum = map[string]ListShapesSortOrderEnum{
	"ASC":  ListShapesSortOrderAsc,
	"DESC": ListShapesSortOrderDesc,
}

var mappingListShapesSortOrderEnumLowerCase = map[string]ListShapesSortOrderEnum{
	"asc":  ListShapesSortOrderAsc,
	"desc": ListShapesSortOrderDesc,
}

// GetListShapesSortOrderEnumValues Enumerates the set of values for ListShapesSortOrderEnum
func GetListShapesSortOrderEnumValues() []ListShapesSortOrderEnum {
	values := make([]ListShapesSortOrderEnum, 0)
	for _, v := range mappingListShapesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListShapesSortOrderEnumStringValues Enumerates the set of values in String for ListShapesSortOrderEnum
func GetListShapesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListShapesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListShapesSortOrderEnum(val string) (ListShapesSortOrderEnum, bool) {
	enum, ok := mappingListShapesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListShapesSortByEnum Enum with underlying type: string
type ListShapesSortByEnum string

// Set of constants representing the allowable values for ListShapesSortByEnum
const (
	ListShapesSortByTimecreated ListShapesSortByEnum = "timeCreated"
	ListShapesSortByDisplayname ListShapesSortByEnum = "displayName"
)

var mappingListShapesSortByEnum = map[string]ListShapesSortByEnum{
	"timeCreated": ListShapesSortByTimecreated,
	"displayName": ListShapesSortByDisplayname,
}

var mappingListShapesSortByEnumLowerCase = map[string]ListShapesSortByEnum{
	"timecreated": ListShapesSortByTimecreated,
	"displayname": ListShapesSortByDisplayname,
}

// GetListShapesSortByEnumValues Enumerates the set of values for ListShapesSortByEnum
func GetListShapesSortByEnumValues() []ListShapesSortByEnum {
	values := make([]ListShapesSortByEnum, 0)
	for _, v := range mappingListShapesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListShapesSortByEnumStringValues Enumerates the set of values in String for ListShapesSortByEnum
func GetListShapesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListShapesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListShapesSortByEnum(val string) (ListShapesSortByEnum, bool) {
	enum, ok := mappingListShapesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
