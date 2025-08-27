// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceServerInstalledPatchesRequest wrapper for the ListManagedInstanceServerInstalledPatches operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListManagedInstanceServerInstalledPatches.go.html to see an example of how to use ListManagedInstanceServerInstalledPatchesRequest.
type ListManagedInstanceServerInstalledPatchesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The unique identifier of a server.
	// **Note:** Not an OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ServerId *string `mandatory:"true" contributesTo:"path" name:"serverId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceServerInstalledPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified, _displayName_ is default.
	SortBy ListManagedInstanceServerInstalledPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceServerInstalledPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceServerInstalledPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceServerInstalledPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceServerInstalledPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceServerInstalledPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceServerInstalledPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceServerInstalledPatchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceServerInstalledPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceServerInstalledPatchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceServerInstalledPatchesResponse wrapper for the ListManagedInstanceServerInstalledPatches operation
type ListManagedInstanceServerInstalledPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstalledPatchCollection instances
	InstalledPatchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceServerInstalledPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceServerInstalledPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceServerInstalledPatchesSortOrderEnum Enum with underlying type: string
type ListManagedInstanceServerInstalledPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceServerInstalledPatchesSortOrderEnum
const (
	ListManagedInstanceServerInstalledPatchesSortOrderAsc  ListManagedInstanceServerInstalledPatchesSortOrderEnum = "ASC"
	ListManagedInstanceServerInstalledPatchesSortOrderDesc ListManagedInstanceServerInstalledPatchesSortOrderEnum = "DESC"
)

var mappingListManagedInstanceServerInstalledPatchesSortOrderEnum = map[string]ListManagedInstanceServerInstalledPatchesSortOrderEnum{
	"ASC":  ListManagedInstanceServerInstalledPatchesSortOrderAsc,
	"DESC": ListManagedInstanceServerInstalledPatchesSortOrderDesc,
}

var mappingListManagedInstanceServerInstalledPatchesSortOrderEnumLowerCase = map[string]ListManagedInstanceServerInstalledPatchesSortOrderEnum{
	"asc":  ListManagedInstanceServerInstalledPatchesSortOrderAsc,
	"desc": ListManagedInstanceServerInstalledPatchesSortOrderDesc,
}

// GetListManagedInstanceServerInstalledPatchesSortOrderEnumValues Enumerates the set of values for ListManagedInstanceServerInstalledPatchesSortOrderEnum
func GetListManagedInstanceServerInstalledPatchesSortOrderEnumValues() []ListManagedInstanceServerInstalledPatchesSortOrderEnum {
	values := make([]ListManagedInstanceServerInstalledPatchesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceServerInstalledPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceServerInstalledPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceServerInstalledPatchesSortOrderEnum
func GetListManagedInstanceServerInstalledPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceServerInstalledPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceServerInstalledPatchesSortOrderEnum(val string) (ListManagedInstanceServerInstalledPatchesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceServerInstalledPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceServerInstalledPatchesSortByEnum Enum with underlying type: string
type ListManagedInstanceServerInstalledPatchesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceServerInstalledPatchesSortByEnum
const (
	ListManagedInstanceServerInstalledPatchesSortByDisplayname ListManagedInstanceServerInstalledPatchesSortByEnum = "displayName"
)

var mappingListManagedInstanceServerInstalledPatchesSortByEnum = map[string]ListManagedInstanceServerInstalledPatchesSortByEnum{
	"displayName": ListManagedInstanceServerInstalledPatchesSortByDisplayname,
}

var mappingListManagedInstanceServerInstalledPatchesSortByEnumLowerCase = map[string]ListManagedInstanceServerInstalledPatchesSortByEnum{
	"displayname": ListManagedInstanceServerInstalledPatchesSortByDisplayname,
}

// GetListManagedInstanceServerInstalledPatchesSortByEnumValues Enumerates the set of values for ListManagedInstanceServerInstalledPatchesSortByEnum
func GetListManagedInstanceServerInstalledPatchesSortByEnumValues() []ListManagedInstanceServerInstalledPatchesSortByEnum {
	values := make([]ListManagedInstanceServerInstalledPatchesSortByEnum, 0)
	for _, v := range mappingListManagedInstanceServerInstalledPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceServerInstalledPatchesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceServerInstalledPatchesSortByEnum
func GetListManagedInstanceServerInstalledPatchesSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListManagedInstanceServerInstalledPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceServerInstalledPatchesSortByEnum(val string) (ListManagedInstanceServerInstalledPatchesSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceServerInstalledPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
