// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInventoryResourcesRequest wrapper for the ListInventoryResources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListInventoryResources.go.html to see an example of how to use ListInventoryResourcesRequest.
type ListInventoryResourcesRequest struct {

	// A filter to return only resources whose base Compartment ID(TenancyId) matches the given base Compartment ID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose resource Compartment ID matches the given resource Compartment ID.
	ResourceCompartmentId *string `mandatory:"true" contributesTo:"query" name:"resourceCompartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Resource Region
	ResourceRegion *string `mandatory:"false" contributesTo:"query" name:"resourceRegion"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}={value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	// Example: Identification.Development=Yes
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of inventory properties filters to apply.
	// The key for each inventory property and value for each resource type is "{resourceType}.{inventoryProperty}={value}".
	// Example: Instance.displayName=TEST_INSTANCE
	InventoryProperties []string `contributesTo:"query" name:"inventoryProperties" collectionFormat:"multi"`

	// Fetch resources matching ANY or ALL criteria passed as params in "tags" and "inventoryProperties".
	// Example: matchingCriteria=ANY
	MatchingCriteria *string `mandatory:"false" contributesTo:"query" name:"matchingCriteria"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInventoryResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListInventoryResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInventoryResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInventoryResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInventoryResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInventoryResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInventoryResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInventoryResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInventoryResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInventoryResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInventoryResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInventoryResourcesResponse wrapper for the ListInventoryResources operation
type ListInventoryResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InventoryResourceCollection instances
	InventoryResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInventoryResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInventoryResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInventoryResourcesSortOrderEnum Enum with underlying type: string
type ListInventoryResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListInventoryResourcesSortOrderEnum
const (
	ListInventoryResourcesSortOrderAsc  ListInventoryResourcesSortOrderEnum = "ASC"
	ListInventoryResourcesSortOrderDesc ListInventoryResourcesSortOrderEnum = "DESC"
)

var mappingListInventoryResourcesSortOrderEnum = map[string]ListInventoryResourcesSortOrderEnum{
	"ASC":  ListInventoryResourcesSortOrderAsc,
	"DESC": ListInventoryResourcesSortOrderDesc,
}

var mappingListInventoryResourcesSortOrderEnumLowerCase = map[string]ListInventoryResourcesSortOrderEnum{
	"asc":  ListInventoryResourcesSortOrderAsc,
	"desc": ListInventoryResourcesSortOrderDesc,
}

// GetListInventoryResourcesSortOrderEnumValues Enumerates the set of values for ListInventoryResourcesSortOrderEnum
func GetListInventoryResourcesSortOrderEnumValues() []ListInventoryResourcesSortOrderEnum {
	values := make([]ListInventoryResourcesSortOrderEnum, 0)
	for _, v := range mappingListInventoryResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInventoryResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListInventoryResourcesSortOrderEnum
func GetListInventoryResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInventoryResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInventoryResourcesSortOrderEnum(val string) (ListInventoryResourcesSortOrderEnum, bool) {
	enum, ok := mappingListInventoryResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInventoryResourcesSortByEnum Enum with underlying type: string
type ListInventoryResourcesSortByEnum string

// Set of constants representing the allowable values for ListInventoryResourcesSortByEnum
const (
	ListInventoryResourcesSortByTimecreated ListInventoryResourcesSortByEnum = "timeCreated"
	ListInventoryResourcesSortByDisplayname ListInventoryResourcesSortByEnum = "displayName"
)

var mappingListInventoryResourcesSortByEnum = map[string]ListInventoryResourcesSortByEnum{
	"timeCreated": ListInventoryResourcesSortByTimecreated,
	"displayName": ListInventoryResourcesSortByDisplayname,
}

var mappingListInventoryResourcesSortByEnumLowerCase = map[string]ListInventoryResourcesSortByEnum{
	"timecreated": ListInventoryResourcesSortByTimecreated,
	"displayname": ListInventoryResourcesSortByDisplayname,
}

// GetListInventoryResourcesSortByEnumValues Enumerates the set of values for ListInventoryResourcesSortByEnum
func GetListInventoryResourcesSortByEnumValues() []ListInventoryResourcesSortByEnum {
	values := make([]ListInventoryResourcesSortByEnum, 0)
	for _, v := range mappingListInventoryResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInventoryResourcesSortByEnumStringValues Enumerates the set of values in String for ListInventoryResourcesSortByEnum
func GetListInventoryResourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListInventoryResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInventoryResourcesSortByEnum(val string) (ListInventoryResourcesSortByEnum, bool) {
	enum, ok := mappingListInventoryResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
