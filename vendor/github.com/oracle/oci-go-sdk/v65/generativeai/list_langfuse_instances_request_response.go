// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLangfuseInstancesRequest wrapper for the ListLangfuseInstances operation
type ListLangfuseInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycle state matches the given value.
	LifecycleState LangfuseInstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Langfuse instance.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListLangfuseInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListLangfuseInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLangfuseInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLangfuseInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLangfuseInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLangfuseInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLangfuseInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLangfuseInstanceLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetLangfuseInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLangfuseInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLangfuseInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLangfuseInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLangfuseInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLangfuseInstancesResponse wrapper for the ListLangfuseInstances operation
type ListLangfuseInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LangfuseInstanceCollection instances
	LangfuseInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLangfuseInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLangfuseInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLangfuseInstancesSortOrderEnum Enum with underlying type: string
type ListLangfuseInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListLangfuseInstancesSortOrderEnum
const (
	ListLangfuseInstancesSortOrderAsc  ListLangfuseInstancesSortOrderEnum = "ASC"
	ListLangfuseInstancesSortOrderDesc ListLangfuseInstancesSortOrderEnum = "DESC"
)

var mappingListLangfuseInstancesSortOrderEnum = map[string]ListLangfuseInstancesSortOrderEnum{
	"ASC":  ListLangfuseInstancesSortOrderAsc,
	"DESC": ListLangfuseInstancesSortOrderDesc,
}

var mappingListLangfuseInstancesSortOrderEnumLowerCase = map[string]ListLangfuseInstancesSortOrderEnum{
	"asc":  ListLangfuseInstancesSortOrderAsc,
	"desc": ListLangfuseInstancesSortOrderDesc,
}

// GetListLangfuseInstancesSortOrderEnumValues Enumerates the set of values for ListLangfuseInstancesSortOrderEnum
func GetListLangfuseInstancesSortOrderEnumValues() []ListLangfuseInstancesSortOrderEnum {
	values := make([]ListLangfuseInstancesSortOrderEnum, 0)
	for _, v := range mappingListLangfuseInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLangfuseInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListLangfuseInstancesSortOrderEnum
func GetListLangfuseInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLangfuseInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLangfuseInstancesSortOrderEnum(val string) (ListLangfuseInstancesSortOrderEnum, bool) {
	enum, ok := mappingListLangfuseInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLangfuseInstancesSortByEnum Enum with underlying type: string
type ListLangfuseInstancesSortByEnum string

// Set of constants representing the allowable values for ListLangfuseInstancesSortByEnum
const (
	ListLangfuseInstancesSortByDisplayname ListLangfuseInstancesSortByEnum = "displayName"
	ListLangfuseInstancesSortByTimecreated ListLangfuseInstancesSortByEnum = "timeCreated"
)

var mappingListLangfuseInstancesSortByEnum = map[string]ListLangfuseInstancesSortByEnum{
	"displayName": ListLangfuseInstancesSortByDisplayname,
	"timeCreated": ListLangfuseInstancesSortByTimecreated,
}

var mappingListLangfuseInstancesSortByEnumLowerCase = map[string]ListLangfuseInstancesSortByEnum{
	"displayname": ListLangfuseInstancesSortByDisplayname,
	"timecreated": ListLangfuseInstancesSortByTimecreated,
}

// GetListLangfuseInstancesSortByEnumValues Enumerates the set of values for ListLangfuseInstancesSortByEnum
func GetListLangfuseInstancesSortByEnumValues() []ListLangfuseInstancesSortByEnum {
	values := make([]ListLangfuseInstancesSortByEnum, 0)
	for _, v := range mappingListLangfuseInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLangfuseInstancesSortByEnumStringValues Enumerates the set of values in String for ListLangfuseInstancesSortByEnum
func GetListLangfuseInstancesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListLangfuseInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLangfuseInstancesSortByEnum(val string) (ListLangfuseInstancesSortByEnum, bool) {
	enum, ok := mappingListLangfuseInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
