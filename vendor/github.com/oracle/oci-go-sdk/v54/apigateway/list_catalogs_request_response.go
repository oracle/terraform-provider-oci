// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"net/http"
	"strings"
)

// ListCatalogsRequest wrapper for the ListCatalogs operation
type ListCatalogsRequest struct {

	// The ocid of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state.
	// Example: `CREATING`
	LifecycleState CatalogLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'. The default order depends on the sortBy value.
	SortOrder ListCatalogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for `timeCreated` is descending. Default order for
	// `displayName` is ascending. The `displayName` sort order is case
	// sensitive.
	SortBy ListCatalogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request id for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCatalogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCatalogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCatalogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCatalogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCatalogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingCatalogLifecycleStateEnum[string(request.LifecycleState)]; !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCatalogLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := mappingListCatalogsSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCatalogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListCatalogsSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCatalogsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCatalogsResponse wrapper for the ListCatalogs operation
type ListCatalogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CatalogCollection instances
	CatalogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// id.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain. For important details about how
	// pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response,
	// additional pages of results were seen previously. For important details
	// about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCatalogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCatalogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCatalogsSortOrderEnum Enum with underlying type: string
type ListCatalogsSortOrderEnum string

// Set of constants representing the allowable values for ListCatalogsSortOrderEnum
const (
	ListCatalogsSortOrderAsc  ListCatalogsSortOrderEnum = "ASC"
	ListCatalogsSortOrderDesc ListCatalogsSortOrderEnum = "DESC"
)

var mappingListCatalogsSortOrderEnum = map[string]ListCatalogsSortOrderEnum{
	"ASC":  ListCatalogsSortOrderAsc,
	"DESC": ListCatalogsSortOrderDesc,
}

// GetListCatalogsSortOrderEnumValues Enumerates the set of values for ListCatalogsSortOrderEnum
func GetListCatalogsSortOrderEnumValues() []ListCatalogsSortOrderEnum {
	values := make([]ListCatalogsSortOrderEnum, 0)
	for _, v := range mappingListCatalogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogsSortOrderEnumStringValues Enumerates the set of values in String for ListCatalogsSortOrderEnum
func GetListCatalogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListCatalogsSortByEnum Enum with underlying type: string
type ListCatalogsSortByEnum string

// Set of constants representing the allowable values for ListCatalogsSortByEnum
const (
	ListCatalogsSortByTimecreated ListCatalogsSortByEnum = "timeCreated"
	ListCatalogsSortByDisplayname ListCatalogsSortByEnum = "displayName"
)

var mappingListCatalogsSortByEnum = map[string]ListCatalogsSortByEnum{
	"timeCreated": ListCatalogsSortByTimecreated,
	"displayName": ListCatalogsSortByDisplayname,
}

// GetListCatalogsSortByEnumValues Enumerates the set of values for ListCatalogsSortByEnum
func GetListCatalogsSortByEnumValues() []ListCatalogsSortByEnum {
	values := make([]ListCatalogsSortByEnum, 0)
	for _, v := range mappingListCatalogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogsSortByEnumStringValues Enumerates the set of values in String for ListCatalogsSortByEnum
func GetListCatalogsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}
