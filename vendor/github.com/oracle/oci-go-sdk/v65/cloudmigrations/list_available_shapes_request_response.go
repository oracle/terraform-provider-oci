// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableShapesRequest wrapper for the ListAvailableShapes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudmigrations/ListAvailableShapes.go.html to see an example of how to use ListAvailableShapesRequest.
type ListAvailableShapesRequest struct {

	// Unique migration plan identifier
	MigrationPlanId *string `mandatory:"true" contributesTo:"path" name:"migrationPlanId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the Dvh in which to list resources.
	DvhHostId *string `mandatory:"false" contributesTo:"query" name:"dvhHostId"`

	// The availability domain in which to list resources.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The reserved capacity ID for which to list resources.
	ReservedCapacityId *string `mandatory:"false" contributesTo:"query" name:"reservedCapacityId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of the previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAvailableShapesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for 'timeCreated' is descending. The default order for 'displayName' is ascending.
	SortBy ListAvailableShapesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableShapesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableShapesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableShapesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableShapesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableShapesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableShapesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableShapesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableShapesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableShapesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableShapesResponse wrapper for the ListAvailableShapes operation
type ListAvailableShapesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AvailableShapesCollection instances
	AvailableShapesCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableShapesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableShapesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableShapesSortOrderEnum Enum with underlying type: string
type ListAvailableShapesSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableShapesSortOrderEnum
const (
	ListAvailableShapesSortOrderAsc  ListAvailableShapesSortOrderEnum = "ASC"
	ListAvailableShapesSortOrderDesc ListAvailableShapesSortOrderEnum = "DESC"
)

var mappingListAvailableShapesSortOrderEnum = map[string]ListAvailableShapesSortOrderEnum{
	"ASC":  ListAvailableShapesSortOrderAsc,
	"DESC": ListAvailableShapesSortOrderDesc,
}

var mappingListAvailableShapesSortOrderEnumLowerCase = map[string]ListAvailableShapesSortOrderEnum{
	"asc":  ListAvailableShapesSortOrderAsc,
	"desc": ListAvailableShapesSortOrderDesc,
}

// GetListAvailableShapesSortOrderEnumValues Enumerates the set of values for ListAvailableShapesSortOrderEnum
func GetListAvailableShapesSortOrderEnumValues() []ListAvailableShapesSortOrderEnum {
	values := make([]ListAvailableShapesSortOrderEnum, 0)
	for _, v := range mappingListAvailableShapesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableShapesSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableShapesSortOrderEnum
func GetListAvailableShapesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableShapesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableShapesSortOrderEnum(val string) (ListAvailableShapesSortOrderEnum, bool) {
	enum, ok := mappingListAvailableShapesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableShapesSortByEnum Enum with underlying type: string
type ListAvailableShapesSortByEnum string

// Set of constants representing the allowable values for ListAvailableShapesSortByEnum
const (
	ListAvailableShapesSortByTimecreated ListAvailableShapesSortByEnum = "timeCreated"
	ListAvailableShapesSortByDisplayname ListAvailableShapesSortByEnum = "displayName"
)

var mappingListAvailableShapesSortByEnum = map[string]ListAvailableShapesSortByEnum{
	"timeCreated": ListAvailableShapesSortByTimecreated,
	"displayName": ListAvailableShapesSortByDisplayname,
}

var mappingListAvailableShapesSortByEnumLowerCase = map[string]ListAvailableShapesSortByEnum{
	"timecreated": ListAvailableShapesSortByTimecreated,
	"displayname": ListAvailableShapesSortByDisplayname,
}

// GetListAvailableShapesSortByEnumValues Enumerates the set of values for ListAvailableShapesSortByEnum
func GetListAvailableShapesSortByEnumValues() []ListAvailableShapesSortByEnum {
	values := make([]ListAvailableShapesSortByEnum, 0)
	for _, v := range mappingListAvailableShapesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableShapesSortByEnumStringValues Enumerates the set of values in String for ListAvailableShapesSortByEnum
func GetListAvailableShapesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAvailableShapesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableShapesSortByEnum(val string) (ListAvailableShapesSortByEnum, bool) {
	enum, ok := mappingListAvailableShapesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
