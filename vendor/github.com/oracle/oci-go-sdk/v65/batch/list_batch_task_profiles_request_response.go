// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package batch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBatchTaskProfilesRequest wrapper for the ListBatchTaskProfiles operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/batch/ListBatchTaskProfiles.go.html to see an example of how to use ListBatchTaskProfilesRequest.
type ListBatchTaskProfilesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState BatchTaskProfileLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task profile.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListBatchTaskProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListBatchTaskProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBatchTaskProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBatchTaskProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBatchTaskProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBatchTaskProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBatchTaskProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchTaskProfileLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBatchTaskProfileLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchTaskProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBatchTaskProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBatchTaskProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBatchTaskProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBatchTaskProfilesResponse wrapper for the ListBatchTaskProfiles operation
type ListBatchTaskProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BatchTaskProfileCollection instances
	BatchTaskProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBatchTaskProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBatchTaskProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBatchTaskProfilesSortOrderEnum Enum with underlying type: string
type ListBatchTaskProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListBatchTaskProfilesSortOrderEnum
const (
	ListBatchTaskProfilesSortOrderAsc  ListBatchTaskProfilesSortOrderEnum = "ASC"
	ListBatchTaskProfilesSortOrderDesc ListBatchTaskProfilesSortOrderEnum = "DESC"
)

var mappingListBatchTaskProfilesSortOrderEnum = map[string]ListBatchTaskProfilesSortOrderEnum{
	"ASC":  ListBatchTaskProfilesSortOrderAsc,
	"DESC": ListBatchTaskProfilesSortOrderDesc,
}

var mappingListBatchTaskProfilesSortOrderEnumLowerCase = map[string]ListBatchTaskProfilesSortOrderEnum{
	"asc":  ListBatchTaskProfilesSortOrderAsc,
	"desc": ListBatchTaskProfilesSortOrderDesc,
}

// GetListBatchTaskProfilesSortOrderEnumValues Enumerates the set of values for ListBatchTaskProfilesSortOrderEnum
func GetListBatchTaskProfilesSortOrderEnumValues() []ListBatchTaskProfilesSortOrderEnum {
	values := make([]ListBatchTaskProfilesSortOrderEnum, 0)
	for _, v := range mappingListBatchTaskProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchTaskProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListBatchTaskProfilesSortOrderEnum
func GetListBatchTaskProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBatchTaskProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchTaskProfilesSortOrderEnum(val string) (ListBatchTaskProfilesSortOrderEnum, bool) {
	enum, ok := mappingListBatchTaskProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBatchTaskProfilesSortByEnum Enum with underlying type: string
type ListBatchTaskProfilesSortByEnum string

// Set of constants representing the allowable values for ListBatchTaskProfilesSortByEnum
const (
	ListBatchTaskProfilesSortByTimecreated ListBatchTaskProfilesSortByEnum = "timeCreated"
	ListBatchTaskProfilesSortByDisplayname ListBatchTaskProfilesSortByEnum = "displayName"
)

var mappingListBatchTaskProfilesSortByEnum = map[string]ListBatchTaskProfilesSortByEnum{
	"timeCreated": ListBatchTaskProfilesSortByTimecreated,
	"displayName": ListBatchTaskProfilesSortByDisplayname,
}

var mappingListBatchTaskProfilesSortByEnumLowerCase = map[string]ListBatchTaskProfilesSortByEnum{
	"timecreated": ListBatchTaskProfilesSortByTimecreated,
	"displayname": ListBatchTaskProfilesSortByDisplayname,
}

// GetListBatchTaskProfilesSortByEnumValues Enumerates the set of values for ListBatchTaskProfilesSortByEnum
func GetListBatchTaskProfilesSortByEnumValues() []ListBatchTaskProfilesSortByEnum {
	values := make([]ListBatchTaskProfilesSortByEnum, 0)
	for _, v := range mappingListBatchTaskProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBatchTaskProfilesSortByEnumStringValues Enumerates the set of values in String for ListBatchTaskProfilesSortByEnum
func GetListBatchTaskProfilesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBatchTaskProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBatchTaskProfilesSortByEnum(val string) (ListBatchTaskProfilesSortByEnum, bool) {
	enum, ok := mappingListBatchTaskProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
