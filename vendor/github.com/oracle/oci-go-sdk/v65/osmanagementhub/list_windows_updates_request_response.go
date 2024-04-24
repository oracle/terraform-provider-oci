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

// ListWindowsUpdatesRequest wrapper for the ListWindowsUpdates operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListWindowsUpdates.go.html to see an example of how to use ListWindowsUpdatesRequest.
type ListWindowsUpdatesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only packages that match the given update classification type.
	ClassificationType []ClassificationTypesEnum `contributesTo:"query" name:"classificationType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter based on the unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.
	// Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed'
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListWindowsUpdatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeInstalled is descending. Default order for name or displayName is ascending.
	SortBy ListWindowsUpdatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWindowsUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWindowsUpdatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWindowsUpdatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWindowsUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWindowsUpdatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ClassificationType {
		if _, ok := GetMappingClassificationTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationType: %s. Supported values are: %s.", val, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListWindowsUpdatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWindowsUpdatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWindowsUpdatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWindowsUpdatesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWindowsUpdatesResponse wrapper for the ListWindowsUpdates operation
type ListWindowsUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WindowsUpdateCollection instances
	WindowsUpdateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWindowsUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWindowsUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWindowsUpdatesSortOrderEnum Enum with underlying type: string
type ListWindowsUpdatesSortOrderEnum string

// Set of constants representing the allowable values for ListWindowsUpdatesSortOrderEnum
const (
	ListWindowsUpdatesSortOrderAsc  ListWindowsUpdatesSortOrderEnum = "ASC"
	ListWindowsUpdatesSortOrderDesc ListWindowsUpdatesSortOrderEnum = "DESC"
)

var mappingListWindowsUpdatesSortOrderEnum = map[string]ListWindowsUpdatesSortOrderEnum{
	"ASC":  ListWindowsUpdatesSortOrderAsc,
	"DESC": ListWindowsUpdatesSortOrderDesc,
}

var mappingListWindowsUpdatesSortOrderEnumLowerCase = map[string]ListWindowsUpdatesSortOrderEnum{
	"asc":  ListWindowsUpdatesSortOrderAsc,
	"desc": ListWindowsUpdatesSortOrderDesc,
}

// GetListWindowsUpdatesSortOrderEnumValues Enumerates the set of values for ListWindowsUpdatesSortOrderEnum
func GetListWindowsUpdatesSortOrderEnumValues() []ListWindowsUpdatesSortOrderEnum {
	values := make([]ListWindowsUpdatesSortOrderEnum, 0)
	for _, v := range mappingListWindowsUpdatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWindowsUpdatesSortOrderEnumStringValues Enumerates the set of values in String for ListWindowsUpdatesSortOrderEnum
func GetListWindowsUpdatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWindowsUpdatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWindowsUpdatesSortOrderEnum(val string) (ListWindowsUpdatesSortOrderEnum, bool) {
	enum, ok := mappingListWindowsUpdatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWindowsUpdatesSortByEnum Enum with underlying type: string
type ListWindowsUpdatesSortByEnum string

// Set of constants representing the allowable values for ListWindowsUpdatesSortByEnum
const (
	ListWindowsUpdatesSortByTimecreated ListWindowsUpdatesSortByEnum = "timeCreated"
	ListWindowsUpdatesSortByName        ListWindowsUpdatesSortByEnum = "name"
	ListWindowsUpdatesSortByDisplayname ListWindowsUpdatesSortByEnum = "displayName"
)

var mappingListWindowsUpdatesSortByEnum = map[string]ListWindowsUpdatesSortByEnum{
	"timeCreated": ListWindowsUpdatesSortByTimecreated,
	"name":        ListWindowsUpdatesSortByName,
	"displayName": ListWindowsUpdatesSortByDisplayname,
}

var mappingListWindowsUpdatesSortByEnumLowerCase = map[string]ListWindowsUpdatesSortByEnum{
	"timecreated": ListWindowsUpdatesSortByTimecreated,
	"name":        ListWindowsUpdatesSortByName,
	"displayname": ListWindowsUpdatesSortByDisplayname,
}

// GetListWindowsUpdatesSortByEnumValues Enumerates the set of values for ListWindowsUpdatesSortByEnum
func GetListWindowsUpdatesSortByEnumValues() []ListWindowsUpdatesSortByEnum {
	values := make([]ListWindowsUpdatesSortByEnum, 0)
	for _, v := range mappingListWindowsUpdatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWindowsUpdatesSortByEnumStringValues Enumerates the set of values in String for ListWindowsUpdatesSortByEnum
func GetListWindowsUpdatesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"displayName",
	}
}

// GetMappingListWindowsUpdatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWindowsUpdatesSortByEnum(val string) (ListWindowsUpdatesSortByEnum, bool) {
	enum, ok := mappingListWindowsUpdatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
