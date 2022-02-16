// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListAgentImagesRequest wrapper for the ListAgentImages operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAgentImages.go.html to see an example of how to use ListAgentImagesRequest.
type ListAgentImagesRequest struct {

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAgentImagesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAgentImagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAgentImagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAgentImagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAgentImagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAgentImagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAgentImagesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAgentImagesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAgentImagesResponse wrapper for the ListAgentImages operation
type ListAgentImagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AgentImageCollection instances
	AgentImageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAgentImagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAgentImagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAgentImagesSortOrderEnum Enum with underlying type: string
type ListAgentImagesSortOrderEnum string

// Set of constants representing the allowable values for ListAgentImagesSortOrderEnum
const (
	ListAgentImagesSortOrderAsc  ListAgentImagesSortOrderEnum = "ASC"
	ListAgentImagesSortOrderDesc ListAgentImagesSortOrderEnum = "DESC"
)

var mappingListAgentImagesSortOrderEnum = map[string]ListAgentImagesSortOrderEnum{
	"ASC":  ListAgentImagesSortOrderAsc,
	"DESC": ListAgentImagesSortOrderDesc,
}

// GetListAgentImagesSortOrderEnumValues Enumerates the set of values for ListAgentImagesSortOrderEnum
func GetListAgentImagesSortOrderEnumValues() []ListAgentImagesSortOrderEnum {
	values := make([]ListAgentImagesSortOrderEnum, 0)
	for _, v := range mappingListAgentImagesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAgentImagesSortOrderEnumStringValues Enumerates the set of values in String for ListAgentImagesSortOrderEnum
func GetListAgentImagesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAgentImagesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAgentImagesSortOrderEnum(val string) (ListAgentImagesSortOrderEnum, bool) {
	mappingListAgentImagesSortOrderEnumIgnoreCase := make(map[string]ListAgentImagesSortOrderEnum)
	for k, v := range mappingListAgentImagesSortOrderEnum {
		mappingListAgentImagesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAgentImagesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
