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

// ListManagedInstanceAvailableWindowsUpdatesRequest wrapper for the ListManagedInstanceAvailableWindowsUpdates operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceAvailableWindowsUpdates.go.html to see an example of how to use ListManagedInstanceAvailableWindowsUpdatesRequest.
type ListManagedInstanceAvailableWindowsUpdatesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A filter to return only packages that match the given update classification type.
	ClassificationType []ClassificationTypesEnum `contributesTo:"query" name:"classificationType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter based on the unique identifier for the Windows update. Note that this is not an OCID, but is a unique identifier assigned by Microsoft.
	// Example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed'
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Indicates if the update can be installed by the OS Management Hub service.
	IsInstallable WindowsUpdateInstallableEnum `mandatory:"false" contributesTo:"query" name:"isInstallable" omitEmpty:"true"`

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
	SortOrder ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeInstalled is descending. Default order for name or displayName is ascending.
	SortBy ListManagedInstanceAvailableWindowsUpdatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceAvailableWindowsUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceAvailableWindowsUpdatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceAvailableWindowsUpdatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceAvailableWindowsUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceAvailableWindowsUpdatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ClassificationType {
		if _, ok := GetMappingClassificationTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationType: %s. Supported values are: %s.", val, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingWindowsUpdateInstallableEnum(string(request.IsInstallable)); !ok && request.IsInstallable != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsInstallable: %s. Supported values are: %s.", request.IsInstallable, strings.Join(GetWindowsUpdateInstallableEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceAvailableWindowsUpdatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceAvailableWindowsUpdatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceAvailableWindowsUpdatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceAvailableWindowsUpdatesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceAvailableWindowsUpdatesResponse wrapper for the ListManagedInstanceAvailableWindowsUpdates operation
type ListManagedInstanceAvailableWindowsUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AvailableWindowsUpdateCollection instances
	AvailableWindowsUpdateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceAvailableWindowsUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceAvailableWindowsUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum
const (
	ListManagedInstanceAvailableWindowsUpdatesSortOrderAsc  ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum = "ASC"
	ListManagedInstanceAvailableWindowsUpdatesSortOrderDesc ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceAvailableWindowsUpdatesSortOrderEnum = map[string]ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum{
	"ASC":  ListManagedInstanceAvailableWindowsUpdatesSortOrderAsc,
	"DESC": ListManagedInstanceAvailableWindowsUpdatesSortOrderDesc,
}

var mappingListManagedInstanceAvailableWindowsUpdatesSortOrderEnumLowerCase = map[string]ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum{
	"asc":  ListManagedInstanceAvailableWindowsUpdatesSortOrderAsc,
	"desc": ListManagedInstanceAvailableWindowsUpdatesSortOrderDesc,
}

// GetListManagedInstanceAvailableWindowsUpdatesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum
func GetListManagedInstanceAvailableWindowsUpdatesSortOrderEnumValues() []ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum {
	values := make([]ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceAvailableWindowsUpdatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceAvailableWindowsUpdatesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum
func GetListManagedInstanceAvailableWindowsUpdatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceAvailableWindowsUpdatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceAvailableWindowsUpdatesSortOrderEnum(val string) (ListManagedInstanceAvailableWindowsUpdatesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceAvailableWindowsUpdatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceAvailableWindowsUpdatesSortByEnum Enum with underlying type: string
type ListManagedInstanceAvailableWindowsUpdatesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceAvailableWindowsUpdatesSortByEnum
const (
	ListManagedInstanceAvailableWindowsUpdatesSortByTimecreated ListManagedInstanceAvailableWindowsUpdatesSortByEnum = "timeCreated"
	ListManagedInstanceAvailableWindowsUpdatesSortByName        ListManagedInstanceAvailableWindowsUpdatesSortByEnum = "name"
	ListManagedInstanceAvailableWindowsUpdatesSortByDisplayname ListManagedInstanceAvailableWindowsUpdatesSortByEnum = "displayName"
)

var mappingListManagedInstanceAvailableWindowsUpdatesSortByEnum = map[string]ListManagedInstanceAvailableWindowsUpdatesSortByEnum{
	"timeCreated": ListManagedInstanceAvailableWindowsUpdatesSortByTimecreated,
	"name":        ListManagedInstanceAvailableWindowsUpdatesSortByName,
	"displayName": ListManagedInstanceAvailableWindowsUpdatesSortByDisplayname,
}

var mappingListManagedInstanceAvailableWindowsUpdatesSortByEnumLowerCase = map[string]ListManagedInstanceAvailableWindowsUpdatesSortByEnum{
	"timecreated": ListManagedInstanceAvailableWindowsUpdatesSortByTimecreated,
	"name":        ListManagedInstanceAvailableWindowsUpdatesSortByName,
	"displayname": ListManagedInstanceAvailableWindowsUpdatesSortByDisplayname,
}

// GetListManagedInstanceAvailableWindowsUpdatesSortByEnumValues Enumerates the set of values for ListManagedInstanceAvailableWindowsUpdatesSortByEnum
func GetListManagedInstanceAvailableWindowsUpdatesSortByEnumValues() []ListManagedInstanceAvailableWindowsUpdatesSortByEnum {
	values := make([]ListManagedInstanceAvailableWindowsUpdatesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceAvailableWindowsUpdatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceAvailableWindowsUpdatesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceAvailableWindowsUpdatesSortByEnum
func GetListManagedInstanceAvailableWindowsUpdatesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"displayName",
	}
}

// GetMappingListManagedInstanceAvailableWindowsUpdatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceAvailableWindowsUpdatesSortByEnum(val string) (ListManagedInstanceAvailableWindowsUpdatesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceAvailableWindowsUpdatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
