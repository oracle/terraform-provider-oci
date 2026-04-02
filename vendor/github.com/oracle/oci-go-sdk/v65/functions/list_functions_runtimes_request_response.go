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

// ListFunctionsRuntimesRequest wrapper for the ListFunctionsRuntimes operation
type ListFunctionsRuntimesRequest struct {

	// unique FunctionsRuntime identifier
	FunctionsRuntimeId *string `mandatory:"false" contributesTo:"query" name:"functionsRuntimeId"`

	// A filter to return only resources that match the entire FunctionsRuntime name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that contain the supplied filter text in the FunctionsRuntime name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// A filter to return only resources that start with the supplied filter text in the FunctionsRuntime name given.
	NameStartsWith *string `mandatory:"false" contributesTo:"query" name:"nameStartsWith"`

	// A filter to return only resources that match the entire os name given.
	Os *string `mandatory:"false" contributesTo:"query" name:"os"`

	// A filter to return only resources that match the entire language name given.
	Language *string `mandatory:"false" contributesTo:"query" name:"language"`

	// A filter to return only resources where their lifecycleState matches the given lifecycleState.
	LifecycleState FunctionsRuntimeLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token for a list query returned by a previous operation
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListFunctionsRuntimesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListFunctionsRuntimesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFunctionsRuntimesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFunctionsRuntimesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFunctionsRuntimesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListFunctionsRuntimesRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFunctionsRuntimesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFunctionsRuntimesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFunctionsRuntimeLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFunctionsRuntimeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFunctionsRuntimesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFunctionsRuntimesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFunctionsRuntimesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFunctionsRuntimesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFunctionsRuntimesResponse wrapper for the ListFunctionsRuntimes operation
type ListFunctionsRuntimesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FunctionsRuntimeCollection instances
	FunctionsRuntimeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFunctionsRuntimesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFunctionsRuntimesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFunctionsRuntimesSortOrderEnum Enum with underlying type: string
type ListFunctionsRuntimesSortOrderEnum string

// Set of constants representing the allowable values for ListFunctionsRuntimesSortOrderEnum
const (
	ListFunctionsRuntimesSortOrderAsc  ListFunctionsRuntimesSortOrderEnum = "ASC"
	ListFunctionsRuntimesSortOrderDesc ListFunctionsRuntimesSortOrderEnum = "DESC"
)

var mappingListFunctionsRuntimesSortOrderEnum = map[string]ListFunctionsRuntimesSortOrderEnum{
	"ASC":  ListFunctionsRuntimesSortOrderAsc,
	"DESC": ListFunctionsRuntimesSortOrderDesc,
}

var mappingListFunctionsRuntimesSortOrderEnumLowerCase = map[string]ListFunctionsRuntimesSortOrderEnum{
	"asc":  ListFunctionsRuntimesSortOrderAsc,
	"desc": ListFunctionsRuntimesSortOrderDesc,
}

// GetListFunctionsRuntimesSortOrderEnumValues Enumerates the set of values for ListFunctionsRuntimesSortOrderEnum
func GetListFunctionsRuntimesSortOrderEnumValues() []ListFunctionsRuntimesSortOrderEnum {
	values := make([]ListFunctionsRuntimesSortOrderEnum, 0)
	for _, v := range mappingListFunctionsRuntimesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFunctionsRuntimesSortOrderEnumStringValues Enumerates the set of values in String for ListFunctionsRuntimesSortOrderEnum
func GetListFunctionsRuntimesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFunctionsRuntimesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFunctionsRuntimesSortOrderEnum(val string) (ListFunctionsRuntimesSortOrderEnum, bool) {
	enum, ok := mappingListFunctionsRuntimesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFunctionsRuntimesSortByEnum Enum with underlying type: string
type ListFunctionsRuntimesSortByEnum string

// Set of constants representing the allowable values for ListFunctionsRuntimesSortByEnum
const (
	ListFunctionsRuntimesSortByTimecreated        ListFunctionsRuntimesSortByEnum = "timeCreated"
	ListFunctionsRuntimesSortByName               ListFunctionsRuntimesSortByEnum = "name"
	ListFunctionsRuntimesSortByTimedeprecated     ListFunctionsRuntimesSortByEnum = "timeDeprecated"
	ListFunctionsRuntimesSortByTimedecommissioned ListFunctionsRuntimesSortByEnum = "timeDecommissioned"
)

var mappingListFunctionsRuntimesSortByEnum = map[string]ListFunctionsRuntimesSortByEnum{
	"timeCreated":        ListFunctionsRuntimesSortByTimecreated,
	"name":               ListFunctionsRuntimesSortByName,
	"timeDeprecated":     ListFunctionsRuntimesSortByTimedeprecated,
	"timeDecommissioned": ListFunctionsRuntimesSortByTimedecommissioned,
}

var mappingListFunctionsRuntimesSortByEnumLowerCase = map[string]ListFunctionsRuntimesSortByEnum{
	"timecreated":        ListFunctionsRuntimesSortByTimecreated,
	"name":               ListFunctionsRuntimesSortByName,
	"timedeprecated":     ListFunctionsRuntimesSortByTimedeprecated,
	"timedecommissioned": ListFunctionsRuntimesSortByTimedecommissioned,
}

// GetListFunctionsRuntimesSortByEnumValues Enumerates the set of values for ListFunctionsRuntimesSortByEnum
func GetListFunctionsRuntimesSortByEnumValues() []ListFunctionsRuntimesSortByEnum {
	values := make([]ListFunctionsRuntimesSortByEnum, 0)
	for _, v := range mappingListFunctionsRuntimesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFunctionsRuntimesSortByEnumStringValues Enumerates the set of values in String for ListFunctionsRuntimesSortByEnum
func GetListFunctionsRuntimesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"timeDeprecated",
		"timeDecommissioned",
	}
}

// GetMappingListFunctionsRuntimesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFunctionsRuntimesSortByEnum(val string) (ListFunctionsRuntimesSortByEnum, bool) {
	enum, ok := mappingListFunctionsRuntimesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
