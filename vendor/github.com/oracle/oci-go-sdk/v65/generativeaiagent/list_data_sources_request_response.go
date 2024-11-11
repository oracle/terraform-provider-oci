// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataSourcesRequest wrapper for the ListDataSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagent/ListDataSources.go.html to see an example of how to use ListDataSourcesRequest.
type ListDataSourcesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledge base.
	KnowledgeBaseId *string `mandatory:"false" contributesTo:"query" name:"knowledgeBaseId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState DataSourceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDataSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListDataSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataSourceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDataSourceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataSourcesResponse wrapper for the ListDataSources operation
type ListDataSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataSourceCollection instances
	DataSourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSourcesSortOrderEnum Enum with underlying type: string
type ListDataSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListDataSourcesSortOrderEnum
const (
	ListDataSourcesSortOrderAsc  ListDataSourcesSortOrderEnum = "ASC"
	ListDataSourcesSortOrderDesc ListDataSourcesSortOrderEnum = "DESC"
)

var mappingListDataSourcesSortOrderEnum = map[string]ListDataSourcesSortOrderEnum{
	"ASC":  ListDataSourcesSortOrderAsc,
	"DESC": ListDataSourcesSortOrderDesc,
}

var mappingListDataSourcesSortOrderEnumLowerCase = map[string]ListDataSourcesSortOrderEnum{
	"asc":  ListDataSourcesSortOrderAsc,
	"desc": ListDataSourcesSortOrderDesc,
}

// GetListDataSourcesSortOrderEnumValues Enumerates the set of values for ListDataSourcesSortOrderEnum
func GetListDataSourcesSortOrderEnumValues() []ListDataSourcesSortOrderEnum {
	values := make([]ListDataSourcesSortOrderEnum, 0)
	for _, v := range mappingListDataSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListDataSourcesSortOrderEnum
func GetListDataSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesSortOrderEnum(val string) (ListDataSourcesSortOrderEnum, bool) {
	enum, ok := mappingListDataSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSourcesSortByEnum Enum with underlying type: string
type ListDataSourcesSortByEnum string

// Set of constants representing the allowable values for ListDataSourcesSortByEnum
const (
	ListDataSourcesSortByTimecreated ListDataSourcesSortByEnum = "timeCreated"
	ListDataSourcesSortByDisplayname ListDataSourcesSortByEnum = "displayName"
)

var mappingListDataSourcesSortByEnum = map[string]ListDataSourcesSortByEnum{
	"timeCreated": ListDataSourcesSortByTimecreated,
	"displayName": ListDataSourcesSortByDisplayname,
}

var mappingListDataSourcesSortByEnumLowerCase = map[string]ListDataSourcesSortByEnum{
	"timecreated": ListDataSourcesSortByTimecreated,
	"displayname": ListDataSourcesSortByDisplayname,
}

// GetListDataSourcesSortByEnumValues Enumerates the set of values for ListDataSourcesSortByEnum
func GetListDataSourcesSortByEnumValues() []ListDataSourcesSortByEnum {
	values := make([]ListDataSourcesSortByEnum, 0)
	for _, v := range mappingListDataSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourcesSortByEnumStringValues Enumerates the set of values in String for ListDataSourcesSortByEnum
func GetListDataSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDataSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourcesSortByEnum(val string) (ListDataSourcesSortByEnum, bool) {
	enum, ok := mappingListDataSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
