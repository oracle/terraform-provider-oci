// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package artifacts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListRepositoriesRequest wrapper for the ListRepositories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListRepositories.go.html to see an example of how to use ListRepositoriesRequest.
type ListRepositoriesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return the resources for the specified OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that match the isImmutable value.
	IsImmutable *bool `mandatory:"false" contributesTo:"query" name:"isImmutable"`

	// A filter to return only resources that match the given lifecycle state name exactly.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListRepositoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListRepositoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRepositoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRepositoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRepositoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRepositoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRepositoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRepositoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRepositoriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRepositoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRepositoriesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRepositoriesResponse wrapper for the ListRepositories operation
type ListRepositoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryCollection instances
	RepositoryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListRepositoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRepositoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRepositoriesSortByEnum Enum with underlying type: string
type ListRepositoriesSortByEnum string

// Set of constants representing the allowable values for ListRepositoriesSortByEnum
const (
	ListRepositoriesSortByTimecreated ListRepositoriesSortByEnum = "TIMECREATED"
	ListRepositoriesSortByDisplayname ListRepositoriesSortByEnum = "DISPLAYNAME"
)

var mappingListRepositoriesSortByEnum = map[string]ListRepositoriesSortByEnum{
	"TIMECREATED": ListRepositoriesSortByTimecreated,
	"DISPLAYNAME": ListRepositoriesSortByDisplayname,
}

// GetListRepositoriesSortByEnumValues Enumerates the set of values for ListRepositoriesSortByEnum
func GetListRepositoriesSortByEnumValues() []ListRepositoriesSortByEnum {
	values := make([]ListRepositoriesSortByEnum, 0)
	for _, v := range mappingListRepositoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoriesSortByEnumStringValues Enumerates the set of values in String for ListRepositoriesSortByEnum
func GetListRepositoriesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListRepositoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoriesSortByEnum(val string) (ListRepositoriesSortByEnum, bool) {
	mappingListRepositoriesSortByEnumIgnoreCase := make(map[string]ListRepositoriesSortByEnum)
	for k, v := range mappingListRepositoriesSortByEnum {
		mappingListRepositoriesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRepositoriesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListRepositoriesSortOrderEnum Enum with underlying type: string
type ListRepositoriesSortOrderEnum string

// Set of constants representing the allowable values for ListRepositoriesSortOrderEnum
const (
	ListRepositoriesSortOrderAsc  ListRepositoriesSortOrderEnum = "ASC"
	ListRepositoriesSortOrderDesc ListRepositoriesSortOrderEnum = "DESC"
)

var mappingListRepositoriesSortOrderEnum = map[string]ListRepositoriesSortOrderEnum{
	"ASC":  ListRepositoriesSortOrderAsc,
	"DESC": ListRepositoriesSortOrderDesc,
}

// GetListRepositoriesSortOrderEnumValues Enumerates the set of values for ListRepositoriesSortOrderEnum
func GetListRepositoriesSortOrderEnumValues() []ListRepositoriesSortOrderEnum {
	values := make([]ListRepositoriesSortOrderEnum, 0)
	for _, v := range mappingListRepositoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRepositoriesSortOrderEnumStringValues Enumerates the set of values in String for ListRepositoriesSortOrderEnum
func GetListRepositoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRepositoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRepositoriesSortOrderEnum(val string) (ListRepositoriesSortOrderEnum, bool) {
	mappingListRepositoriesSortOrderEnumIgnoreCase := make(map[string]ListRepositoriesSortOrderEnum)
	for k, v := range mappingListRepositoriesSortOrderEnum {
		mappingListRepositoriesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListRepositoriesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
