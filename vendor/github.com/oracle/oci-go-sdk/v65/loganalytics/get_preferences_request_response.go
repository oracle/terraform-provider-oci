// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetPreferencesRequest wrapper for the GetPreferences operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetPreferences.go.html to see an example of how to use GetPreferencesRequest.
type GetPreferencesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder GetPreferencesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned preferences.
	SortBy GetPreferencesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPreferencesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPreferencesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPreferencesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPreferencesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetPreferencesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetPreferencesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetPreferencesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetPreferencesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetPreferencesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetPreferencesResponse wrapper for the GetPreferences operation
type GetPreferencesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsPreferenceCollection instances
	LogAnalyticsPreferenceCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPreferencesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPreferencesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetPreferencesSortOrderEnum Enum with underlying type: string
type GetPreferencesSortOrderEnum string

// Set of constants representing the allowable values for GetPreferencesSortOrderEnum
const (
	GetPreferencesSortOrderAsc  GetPreferencesSortOrderEnum = "ASC"
	GetPreferencesSortOrderDesc GetPreferencesSortOrderEnum = "DESC"
)

var mappingGetPreferencesSortOrderEnum = map[string]GetPreferencesSortOrderEnum{
	"ASC":  GetPreferencesSortOrderAsc,
	"DESC": GetPreferencesSortOrderDesc,
}

var mappingGetPreferencesSortOrderEnumLowerCase = map[string]GetPreferencesSortOrderEnum{
	"asc":  GetPreferencesSortOrderAsc,
	"desc": GetPreferencesSortOrderDesc,
}

// GetGetPreferencesSortOrderEnumValues Enumerates the set of values for GetPreferencesSortOrderEnum
func GetGetPreferencesSortOrderEnumValues() []GetPreferencesSortOrderEnum {
	values := make([]GetPreferencesSortOrderEnum, 0)
	for _, v := range mappingGetPreferencesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPreferencesSortOrderEnumStringValues Enumerates the set of values in String for GetPreferencesSortOrderEnum
func GetGetPreferencesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetPreferencesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPreferencesSortOrderEnum(val string) (GetPreferencesSortOrderEnum, bool) {
	enum, ok := mappingGetPreferencesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetPreferencesSortByEnum Enum with underlying type: string
type GetPreferencesSortByEnum string

// Set of constants representing the allowable values for GetPreferencesSortByEnum
const (
	GetPreferencesSortByName GetPreferencesSortByEnum = "name"
)

var mappingGetPreferencesSortByEnum = map[string]GetPreferencesSortByEnum{
	"name": GetPreferencesSortByName,
}

var mappingGetPreferencesSortByEnumLowerCase = map[string]GetPreferencesSortByEnum{
	"name": GetPreferencesSortByName,
}

// GetGetPreferencesSortByEnumValues Enumerates the set of values for GetPreferencesSortByEnum
func GetGetPreferencesSortByEnumValues() []GetPreferencesSortByEnum {
	values := make([]GetPreferencesSortByEnum, 0)
	for _, v := range mappingGetPreferencesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPreferencesSortByEnumStringValues Enumerates the set of values in String for GetPreferencesSortByEnum
func GetGetPreferencesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingGetPreferencesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPreferencesSortByEnum(val string) (GetPreferencesSortByEnum, bool) {
	enum, ok := mappingGetPreferencesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
