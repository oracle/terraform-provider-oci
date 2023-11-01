// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDiscoveryJobResultsRequest wrapper for the ListDiscoveryJobResults operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryJobResults.go.html to see an example of how to use ListDiscoveryJobResultsRequest.
type ListDiscoveryJobResultsRequest struct {

	// The OCID of the discovery job.
	DiscoveryJobId *string `mandatory:"true" contributesTo:"path" name:"discoveryJobId"`

	// A filter to return only the resources that match the specified discovery type.
	DiscoveryType DiscoveryJobDiscoveryTypeEnum `mandatory:"false" contributesTo:"query" name:"discoveryType" omitEmpty:"true"`

	// A filter to return only the resources that match the specified planned action.
	PlannedAction DiscoveryJobResultPlannedActionEnum `mandatory:"false" contributesTo:"query" name:"plannedAction" omitEmpty:"true"`

	// A filter to return the discovery result resources based on the value of their isResultApplied attribute.
	IsResultApplied *bool `mandatory:"false" contributesTo:"query" name:"isResultApplied"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDiscoveryJobResultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeFinished is descending.
	// The default order for discoveryType, schemaName, objectName, columnName and plannedAction is ascending.
	SortBy ListDiscoveryJobResultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiscoveryJobResultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiscoveryJobResultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiscoveryJobResultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiscoveryJobResultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiscoveryJobResultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryJobDiscoveryTypeEnum(string(request.DiscoveryType)); !ok && request.DiscoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryType: %s. Supported values are: %s.", request.DiscoveryType, strings.Join(GetDiscoveryJobDiscoveryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDiscoveryJobResultPlannedActionEnum(string(request.PlannedAction)); !ok && request.PlannedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlannedAction: %s. Supported values are: %s.", request.PlannedAction, strings.Join(GetDiscoveryJobResultPlannedActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryJobResultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDiscoveryJobResultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryJobResultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDiscoveryJobResultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiscoveryJobResultsResponse wrapper for the ListDiscoveryJobResults operation
type ListDiscoveryJobResultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiscoveryJobResultCollection instances
	DiscoveryJobResultCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDiscoveryJobResultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiscoveryJobResultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiscoveryJobResultsSortOrderEnum Enum with underlying type: string
type ListDiscoveryJobResultsSortOrderEnum string

// Set of constants representing the allowable values for ListDiscoveryJobResultsSortOrderEnum
const (
	ListDiscoveryJobResultsSortOrderAsc  ListDiscoveryJobResultsSortOrderEnum = "ASC"
	ListDiscoveryJobResultsSortOrderDesc ListDiscoveryJobResultsSortOrderEnum = "DESC"
)

var mappingListDiscoveryJobResultsSortOrderEnum = map[string]ListDiscoveryJobResultsSortOrderEnum{
	"ASC":  ListDiscoveryJobResultsSortOrderAsc,
	"DESC": ListDiscoveryJobResultsSortOrderDesc,
}

var mappingListDiscoveryJobResultsSortOrderEnumLowerCase = map[string]ListDiscoveryJobResultsSortOrderEnum{
	"asc":  ListDiscoveryJobResultsSortOrderAsc,
	"desc": ListDiscoveryJobResultsSortOrderDesc,
}

// GetListDiscoveryJobResultsSortOrderEnumValues Enumerates the set of values for ListDiscoveryJobResultsSortOrderEnum
func GetListDiscoveryJobResultsSortOrderEnumValues() []ListDiscoveryJobResultsSortOrderEnum {
	values := make([]ListDiscoveryJobResultsSortOrderEnum, 0)
	for _, v := range mappingListDiscoveryJobResultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobResultsSortOrderEnumStringValues Enumerates the set of values in String for ListDiscoveryJobResultsSortOrderEnum
func GetListDiscoveryJobResultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDiscoveryJobResultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobResultsSortOrderEnum(val string) (ListDiscoveryJobResultsSortOrderEnum, bool) {
	enum, ok := mappingListDiscoveryJobResultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryJobResultsSortByEnum Enum with underlying type: string
type ListDiscoveryJobResultsSortByEnum string

// Set of constants representing the allowable values for ListDiscoveryJobResultsSortByEnum
const (
	ListDiscoveryJobResultsSortByDiscoverytype ListDiscoveryJobResultsSortByEnum = "discoveryType"
	ListDiscoveryJobResultsSortByTimefinished  ListDiscoveryJobResultsSortByEnum = "timeFinished"
	ListDiscoveryJobResultsSortBySchemaname    ListDiscoveryJobResultsSortByEnum = "schemaName"
	ListDiscoveryJobResultsSortByObjectname    ListDiscoveryJobResultsSortByEnum = "objectName"
	ListDiscoveryJobResultsSortByColumnname    ListDiscoveryJobResultsSortByEnum = "columnName"
	ListDiscoveryJobResultsSortByPlannedaction ListDiscoveryJobResultsSortByEnum = "plannedAction"
)

var mappingListDiscoveryJobResultsSortByEnum = map[string]ListDiscoveryJobResultsSortByEnum{
	"discoveryType": ListDiscoveryJobResultsSortByDiscoverytype,
	"timeFinished":  ListDiscoveryJobResultsSortByTimefinished,
	"schemaName":    ListDiscoveryJobResultsSortBySchemaname,
	"objectName":    ListDiscoveryJobResultsSortByObjectname,
	"columnName":    ListDiscoveryJobResultsSortByColumnname,
	"plannedAction": ListDiscoveryJobResultsSortByPlannedaction,
}

var mappingListDiscoveryJobResultsSortByEnumLowerCase = map[string]ListDiscoveryJobResultsSortByEnum{
	"discoverytype": ListDiscoveryJobResultsSortByDiscoverytype,
	"timefinished":  ListDiscoveryJobResultsSortByTimefinished,
	"schemaname":    ListDiscoveryJobResultsSortBySchemaname,
	"objectname":    ListDiscoveryJobResultsSortByObjectname,
	"columnname":    ListDiscoveryJobResultsSortByColumnname,
	"plannedaction": ListDiscoveryJobResultsSortByPlannedaction,
}

// GetListDiscoveryJobResultsSortByEnumValues Enumerates the set of values for ListDiscoveryJobResultsSortByEnum
func GetListDiscoveryJobResultsSortByEnumValues() []ListDiscoveryJobResultsSortByEnum {
	values := make([]ListDiscoveryJobResultsSortByEnum, 0)
	for _, v := range mappingListDiscoveryJobResultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryJobResultsSortByEnumStringValues Enumerates the set of values in String for ListDiscoveryJobResultsSortByEnum
func GetListDiscoveryJobResultsSortByEnumStringValues() []string {
	return []string{
		"discoveryType",
		"timeFinished",
		"schemaName",
		"objectName",
		"columnName",
		"plannedAction",
	}
}

// GetMappingListDiscoveryJobResultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryJobResultsSortByEnum(val string) (ListDiscoveryJobResultsSortByEnum, bool) {
	enum, ok := mappingListDiscoveryJobResultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
