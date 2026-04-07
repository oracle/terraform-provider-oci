// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceSnapsRequest wrapper for the ListManagedInstanceSnaps operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstanceSnaps.go.html to see an example of how to use ListManagedInstanceSnapsRequest.
type ListManagedInstanceSnapsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return resources that may partially match the name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceSnapsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for name is ascending.
	SortBy ListManagedInstanceSnapsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceSnapsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceSnapsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceSnapsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceSnapsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceSnapsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceSnapsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceSnapsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceSnapsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceSnapsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceSnapsResponse wrapper for the ListManagedInstanceSnaps operation
type ListManagedInstanceSnapsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SnapCollection instances
	SnapCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the asynchronous work. You can use this to query its status.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceSnapsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceSnapsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceSnapsSortOrderEnum Enum with underlying type: string
type ListManagedInstanceSnapsSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceSnapsSortOrderEnum
const (
	ListManagedInstanceSnapsSortOrderAsc  ListManagedInstanceSnapsSortOrderEnum = "ASC"
	ListManagedInstanceSnapsSortOrderDesc ListManagedInstanceSnapsSortOrderEnum = "DESC"
)

var mappingListManagedInstanceSnapsSortOrderEnum = map[string]ListManagedInstanceSnapsSortOrderEnum{
	"ASC":  ListManagedInstanceSnapsSortOrderAsc,
	"DESC": ListManagedInstanceSnapsSortOrderDesc,
}

var mappingListManagedInstanceSnapsSortOrderEnumLowerCase = map[string]ListManagedInstanceSnapsSortOrderEnum{
	"asc":  ListManagedInstanceSnapsSortOrderAsc,
	"desc": ListManagedInstanceSnapsSortOrderDesc,
}

// GetListManagedInstanceSnapsSortOrderEnumValues Enumerates the set of values for ListManagedInstanceSnapsSortOrderEnum
func GetListManagedInstanceSnapsSortOrderEnumValues() []ListManagedInstanceSnapsSortOrderEnum {
	values := make([]ListManagedInstanceSnapsSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceSnapsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceSnapsSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceSnapsSortOrderEnum
func GetListManagedInstanceSnapsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceSnapsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceSnapsSortOrderEnum(val string) (ListManagedInstanceSnapsSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceSnapsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceSnapsSortByEnum Enum with underlying type: string
type ListManagedInstanceSnapsSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceSnapsSortByEnum
const (
	ListManagedInstanceSnapsSortByName ListManagedInstanceSnapsSortByEnum = "name"
)

var mappingListManagedInstanceSnapsSortByEnum = map[string]ListManagedInstanceSnapsSortByEnum{
	"name": ListManagedInstanceSnapsSortByName,
}

var mappingListManagedInstanceSnapsSortByEnumLowerCase = map[string]ListManagedInstanceSnapsSortByEnum{
	"name": ListManagedInstanceSnapsSortByName,
}

// GetListManagedInstanceSnapsSortByEnumValues Enumerates the set of values for ListManagedInstanceSnapsSortByEnum
func GetListManagedInstanceSnapsSortByEnumValues() []ListManagedInstanceSnapsSortByEnum {
	values := make([]ListManagedInstanceSnapsSortByEnum, 0)
	for _, v := range mappingListManagedInstanceSnapsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceSnapsSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceSnapsSortByEnum
func GetListManagedInstanceSnapsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListManagedInstanceSnapsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceSnapsSortByEnum(val string) (ListManagedInstanceSnapsSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceSnapsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
