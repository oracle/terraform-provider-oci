// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apiaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApiMetadataRequest wrapper for the ListApiMetadata operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ListApiMetadata.go.html to see an example of how to use ListApiMetadataRequest.
type ListApiMetadataRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ApiMetadataLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only lists of resources that match the entire given service type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListApiMetadataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListApiMetadataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApiMetadataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApiMetadataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApiMetadataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApiMetadataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApiMetadataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiMetadataLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetApiMetadataLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApiMetadataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApiMetadataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApiMetadataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApiMetadataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApiMetadataResponse wrapper for the ListApiMetadata operation
type ListApiMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApiMetadataCollection instances
	ApiMetadataCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApiMetadataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApiMetadataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApiMetadataSortOrderEnum Enum with underlying type: string
type ListApiMetadataSortOrderEnum string

// Set of constants representing the allowable values for ListApiMetadataSortOrderEnum
const (
	ListApiMetadataSortOrderAsc  ListApiMetadataSortOrderEnum = "ASC"
	ListApiMetadataSortOrderDesc ListApiMetadataSortOrderEnum = "DESC"
)

var mappingListApiMetadataSortOrderEnum = map[string]ListApiMetadataSortOrderEnum{
	"ASC":  ListApiMetadataSortOrderAsc,
	"DESC": ListApiMetadataSortOrderDesc,
}

var mappingListApiMetadataSortOrderEnumLowerCase = map[string]ListApiMetadataSortOrderEnum{
	"asc":  ListApiMetadataSortOrderAsc,
	"desc": ListApiMetadataSortOrderDesc,
}

// GetListApiMetadataSortOrderEnumValues Enumerates the set of values for ListApiMetadataSortOrderEnum
func GetListApiMetadataSortOrderEnumValues() []ListApiMetadataSortOrderEnum {
	values := make([]ListApiMetadataSortOrderEnum, 0)
	for _, v := range mappingListApiMetadataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApiMetadataSortOrderEnumStringValues Enumerates the set of values in String for ListApiMetadataSortOrderEnum
func GetListApiMetadataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApiMetadataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApiMetadataSortOrderEnum(val string) (ListApiMetadataSortOrderEnum, bool) {
	enum, ok := mappingListApiMetadataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApiMetadataSortByEnum Enum with underlying type: string
type ListApiMetadataSortByEnum string

// Set of constants representing the allowable values for ListApiMetadataSortByEnum
const (
	ListApiMetadataSortByTimecreated ListApiMetadataSortByEnum = "timeCreated"
	ListApiMetadataSortByDisplayname ListApiMetadataSortByEnum = "displayName"
)

var mappingListApiMetadataSortByEnum = map[string]ListApiMetadataSortByEnum{
	"timeCreated": ListApiMetadataSortByTimecreated,
	"displayName": ListApiMetadataSortByDisplayname,
}

var mappingListApiMetadataSortByEnumLowerCase = map[string]ListApiMetadataSortByEnum{
	"timecreated": ListApiMetadataSortByTimecreated,
	"displayname": ListApiMetadataSortByDisplayname,
}

// GetListApiMetadataSortByEnumValues Enumerates the set of values for ListApiMetadataSortByEnum
func GetListApiMetadataSortByEnumValues() []ListApiMetadataSortByEnum {
	values := make([]ListApiMetadataSortByEnum, 0)
	for _, v := range mappingListApiMetadataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApiMetadataSortByEnumStringValues Enumerates the set of values in String for ListApiMetadataSortByEnum
func GetListApiMetadataSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListApiMetadataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApiMetadataSortByEnum(val string) (ListApiMetadataSortByEnum, bool) {
	enum, ok := mappingListApiMetadataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
