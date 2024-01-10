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

// ListExternalPublicationsRequest wrapper for the ListExternalPublications operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListExternalPublications.go.html to see an example of how to use ListExternalPublicationsRequest.
type ListExternalPublicationsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The task key.
	TaskKey *string `mandatory:"true" contributesTo:"path" name:"taskKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListExternalPublicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListExternalPublicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalPublicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalPublicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalPublicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalPublicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalPublicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalPublicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalPublicationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalPublicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalPublicationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalPublicationsResponse wrapper for the ListExternalPublications operation
type ListExternalPublicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalPublicationSummaryCollection instances
	ExternalPublicationSummaryCollection `presentIn:"body"`

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

func (response ListExternalPublicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalPublicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalPublicationsSortOrderEnum Enum with underlying type: string
type ListExternalPublicationsSortOrderEnum string

// Set of constants representing the allowable values for ListExternalPublicationsSortOrderEnum
const (
	ListExternalPublicationsSortOrderAsc  ListExternalPublicationsSortOrderEnum = "ASC"
	ListExternalPublicationsSortOrderDesc ListExternalPublicationsSortOrderEnum = "DESC"
)

var mappingListExternalPublicationsSortOrderEnum = map[string]ListExternalPublicationsSortOrderEnum{
	"ASC":  ListExternalPublicationsSortOrderAsc,
	"DESC": ListExternalPublicationsSortOrderDesc,
}

var mappingListExternalPublicationsSortOrderEnumLowerCase = map[string]ListExternalPublicationsSortOrderEnum{
	"asc":  ListExternalPublicationsSortOrderAsc,
	"desc": ListExternalPublicationsSortOrderDesc,
}

// GetListExternalPublicationsSortOrderEnumValues Enumerates the set of values for ListExternalPublicationsSortOrderEnum
func GetListExternalPublicationsSortOrderEnumValues() []ListExternalPublicationsSortOrderEnum {
	values := make([]ListExternalPublicationsSortOrderEnum, 0)
	for _, v := range mappingListExternalPublicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalPublicationsSortOrderEnumStringValues Enumerates the set of values in String for ListExternalPublicationsSortOrderEnum
func GetListExternalPublicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalPublicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalPublicationsSortOrderEnum(val string) (ListExternalPublicationsSortOrderEnum, bool) {
	enum, ok := mappingListExternalPublicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalPublicationsSortByEnum Enum with underlying type: string
type ListExternalPublicationsSortByEnum string

// Set of constants representing the allowable values for ListExternalPublicationsSortByEnum
const (
	ListExternalPublicationsSortByTimeCreated ListExternalPublicationsSortByEnum = "TIME_CREATED"
	ListExternalPublicationsSortByDisplayName ListExternalPublicationsSortByEnum = "DISPLAY_NAME"
	ListExternalPublicationsSortByTimeUpdated ListExternalPublicationsSortByEnum = "TIME_UPDATED"
)

var mappingListExternalPublicationsSortByEnum = map[string]ListExternalPublicationsSortByEnum{
	"TIME_CREATED": ListExternalPublicationsSortByTimeCreated,
	"DISPLAY_NAME": ListExternalPublicationsSortByDisplayName,
	"TIME_UPDATED": ListExternalPublicationsSortByTimeUpdated,
}

var mappingListExternalPublicationsSortByEnumLowerCase = map[string]ListExternalPublicationsSortByEnum{
	"time_created": ListExternalPublicationsSortByTimeCreated,
	"display_name": ListExternalPublicationsSortByDisplayName,
	"time_updated": ListExternalPublicationsSortByTimeUpdated,
}

// GetListExternalPublicationsSortByEnumValues Enumerates the set of values for ListExternalPublicationsSortByEnum
func GetListExternalPublicationsSortByEnumValues() []ListExternalPublicationsSortByEnum {
	values := make([]ListExternalPublicationsSortByEnum, 0)
	for _, v := range mappingListExternalPublicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalPublicationsSortByEnumStringValues Enumerates the set of values in String for ListExternalPublicationsSortByEnum
func GetListExternalPublicationsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListExternalPublicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalPublicationsSortByEnum(val string) (ListExternalPublicationsSortByEnum, bool) {
	enum, ok := mappingListExternalPublicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
