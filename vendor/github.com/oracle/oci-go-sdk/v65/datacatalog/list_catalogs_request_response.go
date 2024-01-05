// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCatalogsRequest wrapper for the ListCatalogs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListCatalogs.go.html to see an example of how to use ListCatalogsRequest.
type ListCatalogsRequest struct {

	// The OCID of the compartment where you want to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListCatalogsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCatalogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListCatalogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
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
	if _, ok := GetMappingListCatalogsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListCatalogsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCatalogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCatalogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCatalogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
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

	// A list of []CatalogSummary instances
	Items []CatalogSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCatalogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCatalogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCatalogsLifecycleStateEnum Enum with underlying type: string
type ListCatalogsLifecycleStateEnum string

// Set of constants representing the allowable values for ListCatalogsLifecycleStateEnum
const (
	ListCatalogsLifecycleStateCreating ListCatalogsLifecycleStateEnum = "CREATING"
	ListCatalogsLifecycleStateActive   ListCatalogsLifecycleStateEnum = "ACTIVE"
	ListCatalogsLifecycleStateInactive ListCatalogsLifecycleStateEnum = "INACTIVE"
	ListCatalogsLifecycleStateUpdating ListCatalogsLifecycleStateEnum = "UPDATING"
	ListCatalogsLifecycleStateDeleting ListCatalogsLifecycleStateEnum = "DELETING"
	ListCatalogsLifecycleStateDeleted  ListCatalogsLifecycleStateEnum = "DELETED"
	ListCatalogsLifecycleStateFailed   ListCatalogsLifecycleStateEnum = "FAILED"
	ListCatalogsLifecycleStateMoving   ListCatalogsLifecycleStateEnum = "MOVING"
)

var mappingListCatalogsLifecycleStateEnum = map[string]ListCatalogsLifecycleStateEnum{
	"CREATING": ListCatalogsLifecycleStateCreating,
	"ACTIVE":   ListCatalogsLifecycleStateActive,
	"INACTIVE": ListCatalogsLifecycleStateInactive,
	"UPDATING": ListCatalogsLifecycleStateUpdating,
	"DELETING": ListCatalogsLifecycleStateDeleting,
	"DELETED":  ListCatalogsLifecycleStateDeleted,
	"FAILED":   ListCatalogsLifecycleStateFailed,
	"MOVING":   ListCatalogsLifecycleStateMoving,
}

var mappingListCatalogsLifecycleStateEnumLowerCase = map[string]ListCatalogsLifecycleStateEnum{
	"creating": ListCatalogsLifecycleStateCreating,
	"active":   ListCatalogsLifecycleStateActive,
	"inactive": ListCatalogsLifecycleStateInactive,
	"updating": ListCatalogsLifecycleStateUpdating,
	"deleting": ListCatalogsLifecycleStateDeleting,
	"deleted":  ListCatalogsLifecycleStateDeleted,
	"failed":   ListCatalogsLifecycleStateFailed,
	"moving":   ListCatalogsLifecycleStateMoving,
}

// GetListCatalogsLifecycleStateEnumValues Enumerates the set of values for ListCatalogsLifecycleStateEnum
func GetListCatalogsLifecycleStateEnumValues() []ListCatalogsLifecycleStateEnum {
	values := make([]ListCatalogsLifecycleStateEnum, 0)
	for _, v := range mappingListCatalogsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListCatalogsLifecycleStateEnumStringValues Enumerates the set of values in String for ListCatalogsLifecycleStateEnum
func GetListCatalogsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"MOVING",
	}
}

// GetMappingListCatalogsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogsLifecycleStateEnum(val string) (ListCatalogsLifecycleStateEnum, bool) {
	enum, ok := mappingListCatalogsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListCatalogsSortOrderEnumLowerCase = map[string]ListCatalogsSortOrderEnum{
	"asc":  ListCatalogsSortOrderAsc,
	"desc": ListCatalogsSortOrderDesc,
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

// GetMappingListCatalogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogsSortOrderEnum(val string) (ListCatalogsSortOrderEnum, bool) {
	enum, ok := mappingListCatalogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCatalogsSortByEnum Enum with underlying type: string
type ListCatalogsSortByEnum string

// Set of constants representing the allowable values for ListCatalogsSortByEnum
const (
	ListCatalogsSortByTimecreated ListCatalogsSortByEnum = "TIMECREATED"
	ListCatalogsSortByDisplayname ListCatalogsSortByEnum = "DISPLAYNAME"
)

var mappingListCatalogsSortByEnum = map[string]ListCatalogsSortByEnum{
	"TIMECREATED": ListCatalogsSortByTimecreated,
	"DISPLAYNAME": ListCatalogsSortByDisplayname,
}

var mappingListCatalogsSortByEnumLowerCase = map[string]ListCatalogsSortByEnum{
	"timecreated": ListCatalogsSortByTimecreated,
	"displayname": ListCatalogsSortByDisplayname,
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
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCatalogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCatalogsSortByEnum(val string) (ListCatalogsSortByEnum, bool) {
	enum, ok := mappingListCatalogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
