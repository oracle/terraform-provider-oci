// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListObjectStorageLinksRequest wrapper for the ListObjectStorageLinks operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListObjectStorageLinks.go.html to see an example of how to use ListObjectStorageLinksRequest.
type ListObjectStorageLinksRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ObjectStorageLinkLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Object Storage link.
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
	SortOrder ListObjectStorageLinksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListObjectStorageLinksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.
	LustreFileSystemId *string `mandatory:"false" contributesTo:"query" name:"lustreFileSystemId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListObjectStorageLinksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListObjectStorageLinksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListObjectStorageLinksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListObjectStorageLinksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListObjectStorageLinksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingObjectStorageLinkLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetObjectStorageLinkLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListObjectStorageLinksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListObjectStorageLinksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListObjectStorageLinksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListObjectStorageLinksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListObjectStorageLinksResponse wrapper for the ListObjectStorageLinks operation
type ListObjectStorageLinksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ObjectStorageLinkCollection instances
	ObjectStorageLinkCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListObjectStorageLinksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListObjectStorageLinksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListObjectStorageLinksSortOrderEnum Enum with underlying type: string
type ListObjectStorageLinksSortOrderEnum string

// Set of constants representing the allowable values for ListObjectStorageLinksSortOrderEnum
const (
	ListObjectStorageLinksSortOrderAsc  ListObjectStorageLinksSortOrderEnum = "ASC"
	ListObjectStorageLinksSortOrderDesc ListObjectStorageLinksSortOrderEnum = "DESC"
)

var mappingListObjectStorageLinksSortOrderEnum = map[string]ListObjectStorageLinksSortOrderEnum{
	"ASC":  ListObjectStorageLinksSortOrderAsc,
	"DESC": ListObjectStorageLinksSortOrderDesc,
}

var mappingListObjectStorageLinksSortOrderEnumLowerCase = map[string]ListObjectStorageLinksSortOrderEnum{
	"asc":  ListObjectStorageLinksSortOrderAsc,
	"desc": ListObjectStorageLinksSortOrderDesc,
}

// GetListObjectStorageLinksSortOrderEnumValues Enumerates the set of values for ListObjectStorageLinksSortOrderEnum
func GetListObjectStorageLinksSortOrderEnumValues() []ListObjectStorageLinksSortOrderEnum {
	values := make([]ListObjectStorageLinksSortOrderEnum, 0)
	for _, v := range mappingListObjectStorageLinksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListObjectStorageLinksSortOrderEnumStringValues Enumerates the set of values in String for ListObjectStorageLinksSortOrderEnum
func GetListObjectStorageLinksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListObjectStorageLinksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListObjectStorageLinksSortOrderEnum(val string) (ListObjectStorageLinksSortOrderEnum, bool) {
	enum, ok := mappingListObjectStorageLinksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListObjectStorageLinksSortByEnum Enum with underlying type: string
type ListObjectStorageLinksSortByEnum string

// Set of constants representing the allowable values for ListObjectStorageLinksSortByEnum
const (
	ListObjectStorageLinksSortByTimecreated ListObjectStorageLinksSortByEnum = "timeCreated"
	ListObjectStorageLinksSortByDisplayname ListObjectStorageLinksSortByEnum = "displayName"
)

var mappingListObjectStorageLinksSortByEnum = map[string]ListObjectStorageLinksSortByEnum{
	"timeCreated": ListObjectStorageLinksSortByTimecreated,
	"displayName": ListObjectStorageLinksSortByDisplayname,
}

var mappingListObjectStorageLinksSortByEnumLowerCase = map[string]ListObjectStorageLinksSortByEnum{
	"timecreated": ListObjectStorageLinksSortByTimecreated,
	"displayname": ListObjectStorageLinksSortByDisplayname,
}

// GetListObjectStorageLinksSortByEnumValues Enumerates the set of values for ListObjectStorageLinksSortByEnum
func GetListObjectStorageLinksSortByEnumValues() []ListObjectStorageLinksSortByEnum {
	values := make([]ListObjectStorageLinksSortByEnum, 0)
	for _, v := range mappingListObjectStorageLinksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListObjectStorageLinksSortByEnumStringValues Enumerates the set of values in String for ListObjectStorageLinksSortByEnum
func GetListObjectStorageLinksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListObjectStorageLinksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListObjectStorageLinksSortByEnum(val string) (ListObjectStorageLinksSortByEnum, bool) {
	enum, ok := mappingListObjectStorageLinksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
