// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceInstalledWindowsUpdatesRequest wrapper for the ListManagedInstanceInstalledWindowsUpdates operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceInstalledWindowsUpdates.go.html to see an example of how to use ListManagedInstanceInstalledWindowsUpdatesRequest.
type ListManagedInstanceInstalledWindowsUpdatesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A filter based on the unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.
	// Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed'
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeInstalled is descending. Default order for name or displayName is ascending.
	SortBy ListManagedInstanceInstalledWindowsUpdatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceInstalledWindowsUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceInstalledWindowsUpdatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceInstalledWindowsUpdatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceInstalledWindowsUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceInstalledWindowsUpdatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceInstalledWindowsUpdatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceInstalledWindowsUpdatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceInstalledWindowsUpdatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceInstalledWindowsUpdatesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceInstalledWindowsUpdatesResponse wrapper for the ListManagedInstanceInstalledWindowsUpdates operation
type ListManagedInstanceInstalledWindowsUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstalledWindowsUpdateCollection instances
	InstalledWindowsUpdateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceInstalledWindowsUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceInstalledWindowsUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum
const (
	ListManagedInstanceInstalledWindowsUpdatesSortOrderAsc  ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum = "ASC"
	ListManagedInstanceInstalledWindowsUpdatesSortOrderDesc ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceInstalledWindowsUpdatesSortOrderEnum = map[string]ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum{
	"ASC":  ListManagedInstanceInstalledWindowsUpdatesSortOrderAsc,
	"DESC": ListManagedInstanceInstalledWindowsUpdatesSortOrderDesc,
}

var mappingListManagedInstanceInstalledWindowsUpdatesSortOrderEnumLowerCase = map[string]ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum{
	"asc":  ListManagedInstanceInstalledWindowsUpdatesSortOrderAsc,
	"desc": ListManagedInstanceInstalledWindowsUpdatesSortOrderDesc,
}

// GetListManagedInstanceInstalledWindowsUpdatesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum
func GetListManagedInstanceInstalledWindowsUpdatesSortOrderEnumValues() []ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum {
	values := make([]ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceInstalledWindowsUpdatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceInstalledWindowsUpdatesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum
func GetListManagedInstanceInstalledWindowsUpdatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceInstalledWindowsUpdatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceInstalledWindowsUpdatesSortOrderEnum(val string) (ListManagedInstanceInstalledWindowsUpdatesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceInstalledWindowsUpdatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceInstalledWindowsUpdatesSortByEnum Enum with underlying type: string
type ListManagedInstanceInstalledWindowsUpdatesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceInstalledWindowsUpdatesSortByEnum
const (
	ListManagedInstanceInstalledWindowsUpdatesSortByTimecreated ListManagedInstanceInstalledWindowsUpdatesSortByEnum = "timeCreated"
	ListManagedInstanceInstalledWindowsUpdatesSortByName        ListManagedInstanceInstalledWindowsUpdatesSortByEnum = "name"
	ListManagedInstanceInstalledWindowsUpdatesSortByDisplayname ListManagedInstanceInstalledWindowsUpdatesSortByEnum = "displayName"
)

var mappingListManagedInstanceInstalledWindowsUpdatesSortByEnum = map[string]ListManagedInstanceInstalledWindowsUpdatesSortByEnum{
	"timeCreated": ListManagedInstanceInstalledWindowsUpdatesSortByTimecreated,
	"name":        ListManagedInstanceInstalledWindowsUpdatesSortByName,
	"displayName": ListManagedInstanceInstalledWindowsUpdatesSortByDisplayname,
}

var mappingListManagedInstanceInstalledWindowsUpdatesSortByEnumLowerCase = map[string]ListManagedInstanceInstalledWindowsUpdatesSortByEnum{
	"timecreated": ListManagedInstanceInstalledWindowsUpdatesSortByTimecreated,
	"name":        ListManagedInstanceInstalledWindowsUpdatesSortByName,
	"displayname": ListManagedInstanceInstalledWindowsUpdatesSortByDisplayname,
}

// GetListManagedInstanceInstalledWindowsUpdatesSortByEnumValues Enumerates the set of values for ListManagedInstanceInstalledWindowsUpdatesSortByEnum
func GetListManagedInstanceInstalledWindowsUpdatesSortByEnumValues() []ListManagedInstanceInstalledWindowsUpdatesSortByEnum {
	values := make([]ListManagedInstanceInstalledWindowsUpdatesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceInstalledWindowsUpdatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceInstalledWindowsUpdatesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceInstalledWindowsUpdatesSortByEnum
func GetListManagedInstanceInstalledWindowsUpdatesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"displayName",
	}
}

// GetMappingListManagedInstanceInstalledWindowsUpdatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceInstalledWindowsUpdatesSortByEnum(val string) (ListManagedInstanceInstalledWindowsUpdatesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceInstalledWindowsUpdatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
