// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOperationsRequest wrapper for the ListOperations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataconnectivity/ListOperations.go.html to see an example of how to use ListOperationsRequest.
type ListOperationsRequest struct {

	// The registry OCID.
	RegistryId *string `mandatory:"true" contributesTo:"path" name:"registryId"`

	// The connection key.
	ConnectionKey *string `mandatory:"true" contributesTo:"path" name:"connectionKey"`

	// The schema resource name used for retrieving schemas.
	SchemaResourceName *string `mandatory:"true" contributesTo:"path" name:"schemaResourceName"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order are by relevance score in descending order).
	SortBy ListOperationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListOperationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Endpoint ID used for getDataAssetFullDetails.
	EndpointId *string `mandatory:"false" contributesTo:"query" name:"endpointId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOperationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOperationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOperationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOperationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOperationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOperationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOperationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOperationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOperationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOperationsResponse wrapper for the ListOperations operation
type ListOperationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OperationSummaryCollection instances
	OperationSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListOperationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOperationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOperationsSortByEnum Enum with underlying type: string
type ListOperationsSortByEnum string

// Set of constants representing the allowable values for ListOperationsSortByEnum
const (
	ListOperationsSortById          ListOperationsSortByEnum = "id"
	ListOperationsSortByTimecreated ListOperationsSortByEnum = "timeCreated"
	ListOperationsSortByDisplayname ListOperationsSortByEnum = "displayName"
)

var mappingListOperationsSortByEnum = map[string]ListOperationsSortByEnum{
	"id":          ListOperationsSortById,
	"timeCreated": ListOperationsSortByTimecreated,
	"displayName": ListOperationsSortByDisplayname,
}

var mappingListOperationsSortByEnumLowerCase = map[string]ListOperationsSortByEnum{
	"id":          ListOperationsSortById,
	"timecreated": ListOperationsSortByTimecreated,
	"displayname": ListOperationsSortByDisplayname,
}

// GetListOperationsSortByEnumValues Enumerates the set of values for ListOperationsSortByEnum
func GetListOperationsSortByEnumValues() []ListOperationsSortByEnum {
	values := make([]ListOperationsSortByEnum, 0)
	for _, v := range mappingListOperationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperationsSortByEnumStringValues Enumerates the set of values in String for ListOperationsSortByEnum
func GetListOperationsSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOperationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperationsSortByEnum(val string) (ListOperationsSortByEnum, bool) {
	enum, ok := mappingListOperationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOperationsSortOrderEnum Enum with underlying type: string
type ListOperationsSortOrderEnum string

// Set of constants representing the allowable values for ListOperationsSortOrderEnum
const (
	ListOperationsSortOrderAsc  ListOperationsSortOrderEnum = "ASC"
	ListOperationsSortOrderDesc ListOperationsSortOrderEnum = "DESC"
)

var mappingListOperationsSortOrderEnum = map[string]ListOperationsSortOrderEnum{
	"ASC":  ListOperationsSortOrderAsc,
	"DESC": ListOperationsSortOrderDesc,
}

var mappingListOperationsSortOrderEnumLowerCase = map[string]ListOperationsSortOrderEnum{
	"asc":  ListOperationsSortOrderAsc,
	"desc": ListOperationsSortOrderDesc,
}

// GetListOperationsSortOrderEnumValues Enumerates the set of values for ListOperationsSortOrderEnum
func GetListOperationsSortOrderEnumValues() []ListOperationsSortOrderEnum {
	values := make([]ListOperationsSortOrderEnum, 0)
	for _, v := range mappingListOperationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperationsSortOrderEnumStringValues Enumerates the set of values in String for ListOperationsSortOrderEnum
func GetListOperationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOperationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperationsSortOrderEnum(val string) (ListOperationsSortOrderEnum, bool) {
	enum, ok := mappingListOperationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
