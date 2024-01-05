// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTaskRunLineagesRequest wrapper for the ListTaskRunLineages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskRunLineages.go.html to see an example of how to use ListTaskRunLineagesRequest.
type ListTaskRunLineagesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTaskRunLineagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTaskRunLineagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// This filter parameter can be used to filter by model specific queryable fields of the object <br><br><B>Examples:-</B><br> <ul> <li><B>?filter=status eq Failed</B> returns all objects that have a status field with value Failed</li> </ul>
	Filter []string `contributesTo:"query" name:"filter" collectionFormat:"multi"`

	// This parameter allows users to get objects which were updated after a certain time. The format of timeUpdatedGreaterThan is "YYYY-MM-dd'T'HH:mm:ss.SSS'Z'"
	TimeUpdatedGreaterThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedGreaterThan"`

	// This parameter allows users to get objects which were updated after and at a certain time. The format of timeUpdatedGreaterThanOrEqualTo is "YYYY-MM-dd'T'HH:mm:ss.SSS'Z'"
	TimeUpdatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedGreaterThanOrEqualTo"`

	// This parameter allows users to get objects which were updated before a certain time. The format of timeUpatedLessThan is "YYYY-MM-dd'T'HH:mm:ss.SSS'Z'"
	TimeUpatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpatedLessThan"`

	// This parameter allows users to get objects which were updated before and at a certain time. The format of timeUpatedLessThanOrEqualTo is "YYYY-MM-dd'T'HH:mm:ss.SSS'Z'"
	TimeUpatedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpatedLessThanOrEqualTo"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTaskRunLineagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTaskRunLineagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTaskRunLineagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTaskRunLineagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTaskRunLineagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTaskRunLineagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTaskRunLineagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTaskRunLineagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTaskRunLineagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTaskRunLineagesResponse wrapper for the ListTaskRunLineages operation
type ListTaskRunLineagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TaskRunLineageSummaryCollection instances
	TaskRunLineageSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `TaskRunLineage`s. If this header appears in the response, then this
	// is a partial list of TaskRunLineage. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of TaskRunLineages.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTaskRunLineagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTaskRunLineagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTaskRunLineagesSortOrderEnum Enum with underlying type: string
type ListTaskRunLineagesSortOrderEnum string

// Set of constants representing the allowable values for ListTaskRunLineagesSortOrderEnum
const (
	ListTaskRunLineagesSortOrderAsc  ListTaskRunLineagesSortOrderEnum = "ASC"
	ListTaskRunLineagesSortOrderDesc ListTaskRunLineagesSortOrderEnum = "DESC"
)

var mappingListTaskRunLineagesSortOrderEnum = map[string]ListTaskRunLineagesSortOrderEnum{
	"ASC":  ListTaskRunLineagesSortOrderAsc,
	"DESC": ListTaskRunLineagesSortOrderDesc,
}

var mappingListTaskRunLineagesSortOrderEnumLowerCase = map[string]ListTaskRunLineagesSortOrderEnum{
	"asc":  ListTaskRunLineagesSortOrderAsc,
	"desc": ListTaskRunLineagesSortOrderDesc,
}

// GetListTaskRunLineagesSortOrderEnumValues Enumerates the set of values for ListTaskRunLineagesSortOrderEnum
func GetListTaskRunLineagesSortOrderEnumValues() []ListTaskRunLineagesSortOrderEnum {
	values := make([]ListTaskRunLineagesSortOrderEnum, 0)
	for _, v := range mappingListTaskRunLineagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskRunLineagesSortOrderEnumStringValues Enumerates the set of values in String for ListTaskRunLineagesSortOrderEnum
func GetListTaskRunLineagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTaskRunLineagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskRunLineagesSortOrderEnum(val string) (ListTaskRunLineagesSortOrderEnum, bool) {
	enum, ok := mappingListTaskRunLineagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTaskRunLineagesSortByEnum Enum with underlying type: string
type ListTaskRunLineagesSortByEnum string

// Set of constants representing the allowable values for ListTaskRunLineagesSortByEnum
const (
	ListTaskRunLineagesSortByTimeCreated ListTaskRunLineagesSortByEnum = "TIME_CREATED"
	ListTaskRunLineagesSortByDisplayName ListTaskRunLineagesSortByEnum = "DISPLAY_NAME"
	ListTaskRunLineagesSortByTimeUpdated ListTaskRunLineagesSortByEnum = "TIME_UPDATED"
)

var mappingListTaskRunLineagesSortByEnum = map[string]ListTaskRunLineagesSortByEnum{
	"TIME_CREATED": ListTaskRunLineagesSortByTimeCreated,
	"DISPLAY_NAME": ListTaskRunLineagesSortByDisplayName,
	"TIME_UPDATED": ListTaskRunLineagesSortByTimeUpdated,
}

var mappingListTaskRunLineagesSortByEnumLowerCase = map[string]ListTaskRunLineagesSortByEnum{
	"time_created": ListTaskRunLineagesSortByTimeCreated,
	"display_name": ListTaskRunLineagesSortByDisplayName,
	"time_updated": ListTaskRunLineagesSortByTimeUpdated,
}

// GetListTaskRunLineagesSortByEnumValues Enumerates the set of values for ListTaskRunLineagesSortByEnum
func GetListTaskRunLineagesSortByEnumValues() []ListTaskRunLineagesSortByEnum {
	values := make([]ListTaskRunLineagesSortByEnum, 0)
	for _, v := range mappingListTaskRunLineagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskRunLineagesSortByEnumStringValues Enumerates the set of values in String for ListTaskRunLineagesSortByEnum
func GetListTaskRunLineagesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListTaskRunLineagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskRunLineagesSortByEnum(val string) (ListTaskRunLineagesSortByEnum, bool) {
	enum, ok := mappingListTaskRunLineagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
