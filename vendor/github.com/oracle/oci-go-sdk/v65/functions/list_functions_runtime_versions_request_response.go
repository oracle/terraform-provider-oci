// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFunctionsRuntimeVersionsRequest wrapper for the ListFunctionsRuntimeVersions operation
type ListFunctionsRuntimeVersionsRequest struct {

	// unique FunctionsRuntime identifier
	FunctionsRuntimeId *string `mandatory:"false" contributesTo:"query" name:"functionsRuntimeId"`

	// A filter to return only resources that match the entire FunctionsRuntime name given.
	FunctionsRuntimeName *string `mandatory:"false" contributesTo:"query" name:"functionsRuntimeName"`

	// unique FunctionsRuntimeVersion identifier
	FunctionsRuntimeVersionId *string `mandatory:"false" contributesTo:"query" name:"functionsRuntimeVersionId"`

	// A filter to return only resources that match the entire FunctionsRuntimeVersion name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the entire osVersion name given.
	OsVersion *string `mandatory:"false" contributesTo:"query" name:"osVersion"`

	// A filter to return only resources that match the entire languageVersion name given.
	LanguageVersion *string `mandatory:"false" contributesTo:"query" name:"languageVersion"`

	// Matches the current version associated with a FunctionsRuntime.
	IsCurrentVersion *bool `mandatory:"false" contributesTo:"query" name:"isCurrentVersion"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState FunctionsRuntimeVersionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListFunctionsRuntimeVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token for a list query returned by a previous operation
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListFunctionsRuntimeVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFunctionsRuntimeVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFunctionsRuntimeVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFunctionsRuntimeVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListFunctionsRuntimeVersionsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFunctionsRuntimeVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFunctionsRuntimeVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFunctionsRuntimeVersionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFunctionsRuntimeVersionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFunctionsRuntimeVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFunctionsRuntimeVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFunctionsRuntimeVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFunctionsRuntimeVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFunctionsRuntimeVersionsResponse wrapper for the ListFunctionsRuntimeVersions operation
type ListFunctionsRuntimeVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FunctionsRuntimeVersionCollection instances
	FunctionsRuntimeVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFunctionsRuntimeVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFunctionsRuntimeVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFunctionsRuntimeVersionsSortByEnum Enum with underlying type: string
type ListFunctionsRuntimeVersionsSortByEnum string

// Set of constants representing the allowable values for ListFunctionsRuntimeVersionsSortByEnum
const (
	ListFunctionsRuntimeVersionsSortByTimecreated ListFunctionsRuntimeVersionsSortByEnum = "timeCreated"
	ListFunctionsRuntimeVersionsSortByName        ListFunctionsRuntimeVersionsSortByEnum = "name"
)

var mappingListFunctionsRuntimeVersionsSortByEnum = map[string]ListFunctionsRuntimeVersionsSortByEnum{
	"timeCreated": ListFunctionsRuntimeVersionsSortByTimecreated,
	"name":        ListFunctionsRuntimeVersionsSortByName,
}

var mappingListFunctionsRuntimeVersionsSortByEnumLowerCase = map[string]ListFunctionsRuntimeVersionsSortByEnum{
	"timecreated": ListFunctionsRuntimeVersionsSortByTimecreated,
	"name":        ListFunctionsRuntimeVersionsSortByName,
}

// GetListFunctionsRuntimeVersionsSortByEnumValues Enumerates the set of values for ListFunctionsRuntimeVersionsSortByEnum
func GetListFunctionsRuntimeVersionsSortByEnumValues() []ListFunctionsRuntimeVersionsSortByEnum {
	values := make([]ListFunctionsRuntimeVersionsSortByEnum, 0)
	for _, v := range mappingListFunctionsRuntimeVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFunctionsRuntimeVersionsSortByEnumStringValues Enumerates the set of values in String for ListFunctionsRuntimeVersionsSortByEnum
func GetListFunctionsRuntimeVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListFunctionsRuntimeVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFunctionsRuntimeVersionsSortByEnum(val string) (ListFunctionsRuntimeVersionsSortByEnum, bool) {
	enum, ok := mappingListFunctionsRuntimeVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFunctionsRuntimeVersionsSortOrderEnum Enum with underlying type: string
type ListFunctionsRuntimeVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListFunctionsRuntimeVersionsSortOrderEnum
const (
	ListFunctionsRuntimeVersionsSortOrderAsc  ListFunctionsRuntimeVersionsSortOrderEnum = "ASC"
	ListFunctionsRuntimeVersionsSortOrderDesc ListFunctionsRuntimeVersionsSortOrderEnum = "DESC"
)

var mappingListFunctionsRuntimeVersionsSortOrderEnum = map[string]ListFunctionsRuntimeVersionsSortOrderEnum{
	"ASC":  ListFunctionsRuntimeVersionsSortOrderAsc,
	"DESC": ListFunctionsRuntimeVersionsSortOrderDesc,
}

var mappingListFunctionsRuntimeVersionsSortOrderEnumLowerCase = map[string]ListFunctionsRuntimeVersionsSortOrderEnum{
	"asc":  ListFunctionsRuntimeVersionsSortOrderAsc,
	"desc": ListFunctionsRuntimeVersionsSortOrderDesc,
}

// GetListFunctionsRuntimeVersionsSortOrderEnumValues Enumerates the set of values for ListFunctionsRuntimeVersionsSortOrderEnum
func GetListFunctionsRuntimeVersionsSortOrderEnumValues() []ListFunctionsRuntimeVersionsSortOrderEnum {
	values := make([]ListFunctionsRuntimeVersionsSortOrderEnum, 0)
	for _, v := range mappingListFunctionsRuntimeVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFunctionsRuntimeVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListFunctionsRuntimeVersionsSortOrderEnum
func GetListFunctionsRuntimeVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFunctionsRuntimeVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFunctionsRuntimeVersionsSortOrderEnum(val string) (ListFunctionsRuntimeVersionsSortOrderEnum, bool) {
	enum, ok := mappingListFunctionsRuntimeVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
