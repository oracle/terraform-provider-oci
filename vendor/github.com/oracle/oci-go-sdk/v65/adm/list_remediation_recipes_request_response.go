// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRemediationRecipesRequest wrapper for the ListRemediationRecipes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/adm/ListRemediationRecipes.go.html to see an example of how to use ListRemediationRecipesRequest.
type ListRemediationRecipesRequest struct {

	// A filter to return only resources that match the specified identifier.
	// Required only if the compartmentId query parameter is not specified.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The field used to sort Remediation Recipes. Only one sort order is allowed.
	// Default order for _displayName_ is **ascending alphabetical order**.
	// Default order for _lifecyleState_ is the following sequence: **CREATING, ACTIVE, UPDATING, INACTIVE, FAILED, DELETING, and DELETED**.
	// Default order for _timeCreated_ is **descending**.
	// Default order for _timeUpdated_ is **descending**.
	// Default order for _type_ is the following sequence: **ADM**.
	SortBy ListRemediationRecipesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only Remediation Recipes that match the specified lifecycleState.
	LifecycleState RemediationRecipeLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRemediationRecipesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that belong to the specified compartment identifier.
	// Required only if the id query param is not specified.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRemediationRecipesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRemediationRecipesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRemediationRecipesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRemediationRecipesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRemediationRecipesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRemediationRecipesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRemediationRecipesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRemediationRecipeLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRemediationRecipeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRemediationRecipesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRemediationRecipesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRemediationRecipesResponse wrapper for the ListRemediationRecipes operation
type ListRemediationRecipesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RemediationRecipeCollection instances
	RemediationRecipeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRemediationRecipesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRemediationRecipesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRemediationRecipesSortByEnum Enum with underlying type: string
type ListRemediationRecipesSortByEnum string

// Set of constants representing the allowable values for ListRemediationRecipesSortByEnum
const (
	ListRemediationRecipesSortByDisplayName    ListRemediationRecipesSortByEnum = "DISPLAY_NAME"
	ListRemediationRecipesSortByLifecycleState ListRemediationRecipesSortByEnum = "LIFECYCLE_STATE"
	ListRemediationRecipesSortByTimeCreated    ListRemediationRecipesSortByEnum = "TIME_CREATED"
	ListRemediationRecipesSortByTimeUpdated    ListRemediationRecipesSortByEnum = "TIME_UPDATED"
	ListRemediationRecipesSortByType           ListRemediationRecipesSortByEnum = "TYPE"
)

var mappingListRemediationRecipesSortByEnum = map[string]ListRemediationRecipesSortByEnum{
	"DISPLAY_NAME":    ListRemediationRecipesSortByDisplayName,
	"LIFECYCLE_STATE": ListRemediationRecipesSortByLifecycleState,
	"TIME_CREATED":    ListRemediationRecipesSortByTimeCreated,
	"TIME_UPDATED":    ListRemediationRecipesSortByTimeUpdated,
	"TYPE":            ListRemediationRecipesSortByType,
}

var mappingListRemediationRecipesSortByEnumLowerCase = map[string]ListRemediationRecipesSortByEnum{
	"display_name":    ListRemediationRecipesSortByDisplayName,
	"lifecycle_state": ListRemediationRecipesSortByLifecycleState,
	"time_created":    ListRemediationRecipesSortByTimeCreated,
	"time_updated":    ListRemediationRecipesSortByTimeUpdated,
	"type":            ListRemediationRecipesSortByType,
}

// GetListRemediationRecipesSortByEnumValues Enumerates the set of values for ListRemediationRecipesSortByEnum
func GetListRemediationRecipesSortByEnumValues() []ListRemediationRecipesSortByEnum {
	values := make([]ListRemediationRecipesSortByEnum, 0)
	for _, v := range mappingListRemediationRecipesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRemediationRecipesSortByEnumStringValues Enumerates the set of values in String for ListRemediationRecipesSortByEnum
func GetListRemediationRecipesSortByEnumStringValues() []string {
	return []string{
		"DISPLAY_NAME",
		"LIFECYCLE_STATE",
		"TIME_CREATED",
		"TIME_UPDATED",
		"TYPE",
	}
}

// GetMappingListRemediationRecipesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRemediationRecipesSortByEnum(val string) (ListRemediationRecipesSortByEnum, bool) {
	enum, ok := mappingListRemediationRecipesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRemediationRecipesSortOrderEnum Enum with underlying type: string
type ListRemediationRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListRemediationRecipesSortOrderEnum
const (
	ListRemediationRecipesSortOrderAsc  ListRemediationRecipesSortOrderEnum = "ASC"
	ListRemediationRecipesSortOrderDesc ListRemediationRecipesSortOrderEnum = "DESC"
)

var mappingListRemediationRecipesSortOrderEnum = map[string]ListRemediationRecipesSortOrderEnum{
	"ASC":  ListRemediationRecipesSortOrderAsc,
	"DESC": ListRemediationRecipesSortOrderDesc,
}

var mappingListRemediationRecipesSortOrderEnumLowerCase = map[string]ListRemediationRecipesSortOrderEnum{
	"asc":  ListRemediationRecipesSortOrderAsc,
	"desc": ListRemediationRecipesSortOrderDesc,
}

// GetListRemediationRecipesSortOrderEnumValues Enumerates the set of values for ListRemediationRecipesSortOrderEnum
func GetListRemediationRecipesSortOrderEnumValues() []ListRemediationRecipesSortOrderEnum {
	values := make([]ListRemediationRecipesSortOrderEnum, 0)
	for _, v := range mappingListRemediationRecipesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRemediationRecipesSortOrderEnumStringValues Enumerates the set of values in String for ListRemediationRecipesSortOrderEnum
func GetListRemediationRecipesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRemediationRecipesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRemediationRecipesSortOrderEnum(val string) (ListRemediationRecipesSortOrderEnum, bool) {
	enum, ok := mappingListRemediationRecipesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
