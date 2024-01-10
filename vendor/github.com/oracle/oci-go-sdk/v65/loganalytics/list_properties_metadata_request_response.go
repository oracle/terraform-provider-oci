// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPropertiesMetadataRequest wrapper for the ListPropertiesMetadata operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListPropertiesMetadata.go.html to see an example of how to use ListPropertiesMetadataRequest.
type ListPropertiesMetadataRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The property name used for filtering.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The property display text used for filtering. Only properties matching the specified display
	// name or description will be returned.
	DisplayText *string `mandatory:"false" contributesTo:"query" name:"displayText"`

	// The level for which applicable properties are to be listed.
	Level *string `mandatory:"false" contributesTo:"query" name:"level"`

	// The constraints that apply to the properties at a certain level.
	Constraints *string `mandatory:"false" contributesTo:"query" name:"constraints"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPropertiesMetadataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned properties
	SortBy ListPropertiesMetadataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPropertiesMetadataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPropertiesMetadataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPropertiesMetadataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPropertiesMetadataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPropertiesMetadataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPropertiesMetadataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPropertiesMetadataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPropertiesMetadataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPropertiesMetadataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPropertiesMetadataResponse wrapper for the ListPropertiesMetadata operation
type ListPropertiesMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PropertyMetadataSummaryCollection instances
	PropertyMetadataSummaryCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPropertiesMetadataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPropertiesMetadataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPropertiesMetadataSortOrderEnum Enum with underlying type: string
type ListPropertiesMetadataSortOrderEnum string

// Set of constants representing the allowable values for ListPropertiesMetadataSortOrderEnum
const (
	ListPropertiesMetadataSortOrderAsc  ListPropertiesMetadataSortOrderEnum = "ASC"
	ListPropertiesMetadataSortOrderDesc ListPropertiesMetadataSortOrderEnum = "DESC"
)

var mappingListPropertiesMetadataSortOrderEnum = map[string]ListPropertiesMetadataSortOrderEnum{
	"ASC":  ListPropertiesMetadataSortOrderAsc,
	"DESC": ListPropertiesMetadataSortOrderDesc,
}

var mappingListPropertiesMetadataSortOrderEnumLowerCase = map[string]ListPropertiesMetadataSortOrderEnum{
	"asc":  ListPropertiesMetadataSortOrderAsc,
	"desc": ListPropertiesMetadataSortOrderDesc,
}

// GetListPropertiesMetadataSortOrderEnumValues Enumerates the set of values for ListPropertiesMetadataSortOrderEnum
func GetListPropertiesMetadataSortOrderEnumValues() []ListPropertiesMetadataSortOrderEnum {
	values := make([]ListPropertiesMetadataSortOrderEnum, 0)
	for _, v := range mappingListPropertiesMetadataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPropertiesMetadataSortOrderEnumStringValues Enumerates the set of values in String for ListPropertiesMetadataSortOrderEnum
func GetListPropertiesMetadataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPropertiesMetadataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPropertiesMetadataSortOrderEnum(val string) (ListPropertiesMetadataSortOrderEnum, bool) {
	enum, ok := mappingListPropertiesMetadataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPropertiesMetadataSortByEnum Enum with underlying type: string
type ListPropertiesMetadataSortByEnum string

// Set of constants representing the allowable values for ListPropertiesMetadataSortByEnum
const (
	ListPropertiesMetadataSortByName        ListPropertiesMetadataSortByEnum = "name"
	ListPropertiesMetadataSortByDisplayname ListPropertiesMetadataSortByEnum = "displayName"
)

var mappingListPropertiesMetadataSortByEnum = map[string]ListPropertiesMetadataSortByEnum{
	"name":        ListPropertiesMetadataSortByName,
	"displayName": ListPropertiesMetadataSortByDisplayname,
}

var mappingListPropertiesMetadataSortByEnumLowerCase = map[string]ListPropertiesMetadataSortByEnum{
	"name":        ListPropertiesMetadataSortByName,
	"displayname": ListPropertiesMetadataSortByDisplayname,
}

// GetListPropertiesMetadataSortByEnumValues Enumerates the set of values for ListPropertiesMetadataSortByEnum
func GetListPropertiesMetadataSortByEnumValues() []ListPropertiesMetadataSortByEnum {
	values := make([]ListPropertiesMetadataSortByEnum, 0)
	for _, v := range mappingListPropertiesMetadataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPropertiesMetadataSortByEnumStringValues Enumerates the set of values in String for ListPropertiesMetadataSortByEnum
func GetListPropertiesMetadataSortByEnumStringValues() []string {
	return []string{
		"name",
		"displayName",
	}
}

// GetMappingListPropertiesMetadataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPropertiesMetadataSortByEnum(val string) (ListPropertiesMetadataSortByEnum, bool) {
	enum, ok := mappingListPropertiesMetadataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
