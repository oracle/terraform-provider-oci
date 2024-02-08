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

// ListDisApplicationTaskRunLineagesRequest wrapper for the ListDisApplicationTaskRunLineages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDisApplicationTaskRunLineages.go.html to see an example of how to use ListDisApplicationTaskRunLineagesRequest.
type ListDisApplicationTaskRunLineagesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The OCID of the DIS Application.
	DisApplicationId *string `mandatory:"true" contributesTo:"path" name:"disApplicationId"`

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
	SortOrder ListDisApplicationTaskRunLineagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListDisApplicationTaskRunLineagesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

func (request ListDisApplicationTaskRunLineagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDisApplicationTaskRunLineagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDisApplicationTaskRunLineagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDisApplicationTaskRunLineagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDisApplicationTaskRunLineagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDisApplicationTaskRunLineagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDisApplicationTaskRunLineagesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDisApplicationTaskRunLineagesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDisApplicationTaskRunLineagesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDisApplicationTaskRunLineagesResponse wrapper for the ListDisApplicationTaskRunLineages operation
type ListDisApplicationTaskRunLineagesResponse struct {

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

func (response ListDisApplicationTaskRunLineagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDisApplicationTaskRunLineagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDisApplicationTaskRunLineagesSortOrderEnum Enum with underlying type: string
type ListDisApplicationTaskRunLineagesSortOrderEnum string

// Set of constants representing the allowable values for ListDisApplicationTaskRunLineagesSortOrderEnum
const (
	ListDisApplicationTaskRunLineagesSortOrderAsc  ListDisApplicationTaskRunLineagesSortOrderEnum = "ASC"
	ListDisApplicationTaskRunLineagesSortOrderDesc ListDisApplicationTaskRunLineagesSortOrderEnum = "DESC"
)

var mappingListDisApplicationTaskRunLineagesSortOrderEnum = map[string]ListDisApplicationTaskRunLineagesSortOrderEnum{
	"ASC":  ListDisApplicationTaskRunLineagesSortOrderAsc,
	"DESC": ListDisApplicationTaskRunLineagesSortOrderDesc,
}

var mappingListDisApplicationTaskRunLineagesSortOrderEnumLowerCase = map[string]ListDisApplicationTaskRunLineagesSortOrderEnum{
	"asc":  ListDisApplicationTaskRunLineagesSortOrderAsc,
	"desc": ListDisApplicationTaskRunLineagesSortOrderDesc,
}

// GetListDisApplicationTaskRunLineagesSortOrderEnumValues Enumerates the set of values for ListDisApplicationTaskRunLineagesSortOrderEnum
func GetListDisApplicationTaskRunLineagesSortOrderEnumValues() []ListDisApplicationTaskRunLineagesSortOrderEnum {
	values := make([]ListDisApplicationTaskRunLineagesSortOrderEnum, 0)
	for _, v := range mappingListDisApplicationTaskRunLineagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDisApplicationTaskRunLineagesSortOrderEnumStringValues Enumerates the set of values in String for ListDisApplicationTaskRunLineagesSortOrderEnum
func GetListDisApplicationTaskRunLineagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDisApplicationTaskRunLineagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDisApplicationTaskRunLineagesSortOrderEnum(val string) (ListDisApplicationTaskRunLineagesSortOrderEnum, bool) {
	enum, ok := mappingListDisApplicationTaskRunLineagesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDisApplicationTaskRunLineagesSortByEnum Enum with underlying type: string
type ListDisApplicationTaskRunLineagesSortByEnum string

// Set of constants representing the allowable values for ListDisApplicationTaskRunLineagesSortByEnum
const (
	ListDisApplicationTaskRunLineagesSortByTimeCreated ListDisApplicationTaskRunLineagesSortByEnum = "TIME_CREATED"
	ListDisApplicationTaskRunLineagesSortByDisplayName ListDisApplicationTaskRunLineagesSortByEnum = "DISPLAY_NAME"
	ListDisApplicationTaskRunLineagesSortByTimeUpdated ListDisApplicationTaskRunLineagesSortByEnum = "TIME_UPDATED"
)

var mappingListDisApplicationTaskRunLineagesSortByEnum = map[string]ListDisApplicationTaskRunLineagesSortByEnum{
	"TIME_CREATED": ListDisApplicationTaskRunLineagesSortByTimeCreated,
	"DISPLAY_NAME": ListDisApplicationTaskRunLineagesSortByDisplayName,
	"TIME_UPDATED": ListDisApplicationTaskRunLineagesSortByTimeUpdated,
}

var mappingListDisApplicationTaskRunLineagesSortByEnumLowerCase = map[string]ListDisApplicationTaskRunLineagesSortByEnum{
	"time_created": ListDisApplicationTaskRunLineagesSortByTimeCreated,
	"display_name": ListDisApplicationTaskRunLineagesSortByDisplayName,
	"time_updated": ListDisApplicationTaskRunLineagesSortByTimeUpdated,
}

// GetListDisApplicationTaskRunLineagesSortByEnumValues Enumerates the set of values for ListDisApplicationTaskRunLineagesSortByEnum
func GetListDisApplicationTaskRunLineagesSortByEnumValues() []ListDisApplicationTaskRunLineagesSortByEnum {
	values := make([]ListDisApplicationTaskRunLineagesSortByEnum, 0)
	for _, v := range mappingListDisApplicationTaskRunLineagesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDisApplicationTaskRunLineagesSortByEnumStringValues Enumerates the set of values in String for ListDisApplicationTaskRunLineagesSortByEnum
func GetListDisApplicationTaskRunLineagesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListDisApplicationTaskRunLineagesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDisApplicationTaskRunLineagesSortByEnum(val string) (ListDisApplicationTaskRunLineagesSortByEnum, bool) {
	enum, ok := mappingListDisApplicationTaskRunLineagesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
