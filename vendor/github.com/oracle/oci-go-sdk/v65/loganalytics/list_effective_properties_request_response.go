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

// ListEffectivePropertiesRequest wrapper for the ListEffectiveProperties operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListEffectiveProperties.go.html to see an example of how to use ListEffectivePropertiesRequest.
type ListEffectivePropertiesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The agent ocid.
	AgentId *string `mandatory:"false" contributesTo:"query" name:"agentId"`

	// The source name.
	SourceName *string `mandatory:"false" contributesTo:"query" name:"sourceName"`

	// The include pattern flag.
	IsIncludePatterns *bool `mandatory:"false" contributesTo:"query" name:"isIncludePatterns"`

	// The entity ocid.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// The pattern id.
	PatternId *int `mandatory:"false" contributesTo:"query" name:"patternId"`

	// The property name used for filtering.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListEffectivePropertiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned properties
	SortBy ListEffectivePropertiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEffectivePropertiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEffectivePropertiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEffectivePropertiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEffectivePropertiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEffectivePropertiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEffectivePropertiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEffectivePropertiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEffectivePropertiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEffectivePropertiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEffectivePropertiesResponse wrapper for the ListEffectiveProperties operation
type ListEffectivePropertiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EffectivePropertyCollection instances
	EffectivePropertyCollection `presentIn:"body"`

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

func (response ListEffectivePropertiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEffectivePropertiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEffectivePropertiesSortOrderEnum Enum with underlying type: string
type ListEffectivePropertiesSortOrderEnum string

// Set of constants representing the allowable values for ListEffectivePropertiesSortOrderEnum
const (
	ListEffectivePropertiesSortOrderAsc  ListEffectivePropertiesSortOrderEnum = "ASC"
	ListEffectivePropertiesSortOrderDesc ListEffectivePropertiesSortOrderEnum = "DESC"
)

var mappingListEffectivePropertiesSortOrderEnum = map[string]ListEffectivePropertiesSortOrderEnum{
	"ASC":  ListEffectivePropertiesSortOrderAsc,
	"DESC": ListEffectivePropertiesSortOrderDesc,
}

var mappingListEffectivePropertiesSortOrderEnumLowerCase = map[string]ListEffectivePropertiesSortOrderEnum{
	"asc":  ListEffectivePropertiesSortOrderAsc,
	"desc": ListEffectivePropertiesSortOrderDesc,
}

// GetListEffectivePropertiesSortOrderEnumValues Enumerates the set of values for ListEffectivePropertiesSortOrderEnum
func GetListEffectivePropertiesSortOrderEnumValues() []ListEffectivePropertiesSortOrderEnum {
	values := make([]ListEffectivePropertiesSortOrderEnum, 0)
	for _, v := range mappingListEffectivePropertiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEffectivePropertiesSortOrderEnumStringValues Enumerates the set of values in String for ListEffectivePropertiesSortOrderEnum
func GetListEffectivePropertiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEffectivePropertiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEffectivePropertiesSortOrderEnum(val string) (ListEffectivePropertiesSortOrderEnum, bool) {
	enum, ok := mappingListEffectivePropertiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEffectivePropertiesSortByEnum Enum with underlying type: string
type ListEffectivePropertiesSortByEnum string

// Set of constants representing the allowable values for ListEffectivePropertiesSortByEnum
const (
	ListEffectivePropertiesSortByName        ListEffectivePropertiesSortByEnum = "name"
	ListEffectivePropertiesSortByDisplayname ListEffectivePropertiesSortByEnum = "displayName"
)

var mappingListEffectivePropertiesSortByEnum = map[string]ListEffectivePropertiesSortByEnum{
	"name":        ListEffectivePropertiesSortByName,
	"displayName": ListEffectivePropertiesSortByDisplayname,
}

var mappingListEffectivePropertiesSortByEnumLowerCase = map[string]ListEffectivePropertiesSortByEnum{
	"name":        ListEffectivePropertiesSortByName,
	"displayname": ListEffectivePropertiesSortByDisplayname,
}

// GetListEffectivePropertiesSortByEnumValues Enumerates the set of values for ListEffectivePropertiesSortByEnum
func GetListEffectivePropertiesSortByEnumValues() []ListEffectivePropertiesSortByEnum {
	values := make([]ListEffectivePropertiesSortByEnum, 0)
	for _, v := range mappingListEffectivePropertiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEffectivePropertiesSortByEnumStringValues Enumerates the set of values in String for ListEffectivePropertiesSortByEnum
func GetListEffectivePropertiesSortByEnumStringValues() []string {
	return []string{
		"name",
		"displayName",
	}
}

// GetMappingListEffectivePropertiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEffectivePropertiesSortByEnum(val string) (ListEffectivePropertiesSortByEnum, bool) {
	enum, ok := mappingListEffectivePropertiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
