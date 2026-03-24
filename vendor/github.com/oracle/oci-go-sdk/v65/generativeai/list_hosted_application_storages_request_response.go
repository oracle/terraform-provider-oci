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

// ListHostedApplicationStoragesRequest wrapper for the ListHostedApplicationStorages operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListHostedApplicationStorages.go.html to see an example of how to use ListHostedApplicationStoragesRequest.
type ListHostedApplicationStoragesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the hosted applications that their lifecycle state matches the given lifecycle state.
	LifecycleState HostedApplicationStorageLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The type of the hosted application storage.
	HostedApplicationStorageType HostedApplicationStorageStorageTypeEnum `mandatory:"false" contributesTo:"query" name:"hostedApplicationStorageType" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the hosted application storage.
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
	SortOrder ListHostedApplicationStoragesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListHostedApplicationStoragesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHostedApplicationStoragesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHostedApplicationStoragesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHostedApplicationStoragesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHostedApplicationStoragesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHostedApplicationStoragesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHostedApplicationStorageLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetHostedApplicationStorageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHostedApplicationStorageStorageTypeEnum(string(request.HostedApplicationStorageType)); !ok && request.HostedApplicationStorageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HostedApplicationStorageType: %s. Supported values are: %s.", request.HostedApplicationStorageType, strings.Join(GetHostedApplicationStorageStorageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHostedApplicationStoragesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHostedApplicationStoragesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHostedApplicationStoragesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHostedApplicationStoragesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHostedApplicationStoragesResponse wrapper for the ListHostedApplicationStorages operation
type ListHostedApplicationStoragesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of HostedApplicationStorageCollection instances
	HostedApplicationStorageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHostedApplicationStoragesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHostedApplicationStoragesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHostedApplicationStoragesSortOrderEnum Enum with underlying type: string
type ListHostedApplicationStoragesSortOrderEnum string

// Set of constants representing the allowable values for ListHostedApplicationStoragesSortOrderEnum
const (
	ListHostedApplicationStoragesSortOrderAsc  ListHostedApplicationStoragesSortOrderEnum = "ASC"
	ListHostedApplicationStoragesSortOrderDesc ListHostedApplicationStoragesSortOrderEnum = "DESC"
)

var mappingListHostedApplicationStoragesSortOrderEnum = map[string]ListHostedApplicationStoragesSortOrderEnum{
	"ASC":  ListHostedApplicationStoragesSortOrderAsc,
	"DESC": ListHostedApplicationStoragesSortOrderDesc,
}

var mappingListHostedApplicationStoragesSortOrderEnumLowerCase = map[string]ListHostedApplicationStoragesSortOrderEnum{
	"asc":  ListHostedApplicationStoragesSortOrderAsc,
	"desc": ListHostedApplicationStoragesSortOrderDesc,
}

// GetListHostedApplicationStoragesSortOrderEnumValues Enumerates the set of values for ListHostedApplicationStoragesSortOrderEnum
func GetListHostedApplicationStoragesSortOrderEnumValues() []ListHostedApplicationStoragesSortOrderEnum {
	values := make([]ListHostedApplicationStoragesSortOrderEnum, 0)
	for _, v := range mappingListHostedApplicationStoragesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostedApplicationStoragesSortOrderEnumStringValues Enumerates the set of values in String for ListHostedApplicationStoragesSortOrderEnum
func GetListHostedApplicationStoragesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHostedApplicationStoragesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostedApplicationStoragesSortOrderEnum(val string) (ListHostedApplicationStoragesSortOrderEnum, bool) {
	enum, ok := mappingListHostedApplicationStoragesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHostedApplicationStoragesSortByEnum Enum with underlying type: string
type ListHostedApplicationStoragesSortByEnum string

// Set of constants representing the allowable values for ListHostedApplicationStoragesSortByEnum
const (
	ListHostedApplicationStoragesSortByTimecreated    ListHostedApplicationStoragesSortByEnum = "timeCreated"
	ListHostedApplicationStoragesSortByDisplayname    ListHostedApplicationStoragesSortByEnum = "displayName"
	ListHostedApplicationStoragesSortByLifecyclestate ListHostedApplicationStoragesSortByEnum = "lifecycleState"
)

var mappingListHostedApplicationStoragesSortByEnum = map[string]ListHostedApplicationStoragesSortByEnum{
	"timeCreated":    ListHostedApplicationStoragesSortByTimecreated,
	"displayName":    ListHostedApplicationStoragesSortByDisplayname,
	"lifecycleState": ListHostedApplicationStoragesSortByLifecyclestate,
}

var mappingListHostedApplicationStoragesSortByEnumLowerCase = map[string]ListHostedApplicationStoragesSortByEnum{
	"timecreated":    ListHostedApplicationStoragesSortByTimecreated,
	"displayname":    ListHostedApplicationStoragesSortByDisplayname,
	"lifecyclestate": ListHostedApplicationStoragesSortByLifecyclestate,
}

// GetListHostedApplicationStoragesSortByEnumValues Enumerates the set of values for ListHostedApplicationStoragesSortByEnum
func GetListHostedApplicationStoragesSortByEnumValues() []ListHostedApplicationStoragesSortByEnum {
	values := make([]ListHostedApplicationStoragesSortByEnum, 0)
	for _, v := range mappingListHostedApplicationStoragesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHostedApplicationStoragesSortByEnumStringValues Enumerates the set of values in String for ListHostedApplicationStoragesSortByEnum
func GetListHostedApplicationStoragesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
	}
}

// GetMappingListHostedApplicationStoragesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHostedApplicationStoragesSortByEnum(val string) (ListHostedApplicationStoragesSortByEnum, bool) {
	enum, ok := mappingListHostedApplicationStoragesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
