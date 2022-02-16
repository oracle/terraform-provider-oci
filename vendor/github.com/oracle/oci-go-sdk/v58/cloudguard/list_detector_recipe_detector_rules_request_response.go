// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDetectorRecipeDetectorRulesRequest wrapper for the ListDetectorRecipeDetectorRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDetectorRecipeDetectorRules.go.html to see an example of how to use ListDetectorRecipeDetectorRulesRequest.
type ListDetectorRecipeDetectorRulesRequest struct {

	// DetectorRecipe OCID
	DetectorRecipeId *string `mandatory:"true" contributesTo:"path" name:"detectorRecipeId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListDetectorRecipeDetectorRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDetectorRecipeDetectorRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListDetectorRecipeDetectorRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDetectorRecipeDetectorRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDetectorRecipeDetectorRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDetectorRecipeDetectorRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDetectorRecipeDetectorRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDetectorRecipeDetectorRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDetectorRecipeDetectorRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDetectorRecipeDetectorRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectorRecipeDetectorRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDetectorRecipeDetectorRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectorRecipeDetectorRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDetectorRecipeDetectorRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDetectorRecipeDetectorRulesResponse wrapper for the ListDetectorRecipeDetectorRules operation
type ListDetectorRecipeDetectorRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DetectorRecipeDetectorRuleCollection instances
	DetectorRecipeDetectorRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDetectorRecipeDetectorRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDetectorRecipeDetectorRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDetectorRecipeDetectorRulesLifecycleStateEnum Enum with underlying type: string
type ListDetectorRecipeDetectorRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDetectorRecipeDetectorRulesLifecycleStateEnum
const (
	ListDetectorRecipeDetectorRulesLifecycleStateCreating ListDetectorRecipeDetectorRulesLifecycleStateEnum = "CREATING"
	ListDetectorRecipeDetectorRulesLifecycleStateUpdating ListDetectorRecipeDetectorRulesLifecycleStateEnum = "UPDATING"
	ListDetectorRecipeDetectorRulesLifecycleStateActive   ListDetectorRecipeDetectorRulesLifecycleStateEnum = "ACTIVE"
	ListDetectorRecipeDetectorRulesLifecycleStateInactive ListDetectorRecipeDetectorRulesLifecycleStateEnum = "INACTIVE"
	ListDetectorRecipeDetectorRulesLifecycleStateDeleting ListDetectorRecipeDetectorRulesLifecycleStateEnum = "DELETING"
	ListDetectorRecipeDetectorRulesLifecycleStateDeleted  ListDetectorRecipeDetectorRulesLifecycleStateEnum = "DELETED"
	ListDetectorRecipeDetectorRulesLifecycleStateFailed   ListDetectorRecipeDetectorRulesLifecycleStateEnum = "FAILED"
)

var mappingListDetectorRecipeDetectorRulesLifecycleStateEnum = map[string]ListDetectorRecipeDetectorRulesLifecycleStateEnum{
	"CREATING": ListDetectorRecipeDetectorRulesLifecycleStateCreating,
	"UPDATING": ListDetectorRecipeDetectorRulesLifecycleStateUpdating,
	"ACTIVE":   ListDetectorRecipeDetectorRulesLifecycleStateActive,
	"INACTIVE": ListDetectorRecipeDetectorRulesLifecycleStateInactive,
	"DELETING": ListDetectorRecipeDetectorRulesLifecycleStateDeleting,
	"DELETED":  ListDetectorRecipeDetectorRulesLifecycleStateDeleted,
	"FAILED":   ListDetectorRecipeDetectorRulesLifecycleStateFailed,
}

// GetListDetectorRecipeDetectorRulesLifecycleStateEnumValues Enumerates the set of values for ListDetectorRecipeDetectorRulesLifecycleStateEnum
func GetListDetectorRecipeDetectorRulesLifecycleStateEnumValues() []ListDetectorRecipeDetectorRulesLifecycleStateEnum {
	values := make([]ListDetectorRecipeDetectorRulesLifecycleStateEnum, 0)
	for _, v := range mappingListDetectorRecipeDetectorRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipeDetectorRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDetectorRecipeDetectorRulesLifecycleStateEnum
func GetListDetectorRecipeDetectorRulesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDetectorRecipeDetectorRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipeDetectorRulesLifecycleStateEnum(val string) (ListDetectorRecipeDetectorRulesLifecycleStateEnum, bool) {
	mappingListDetectorRecipeDetectorRulesLifecycleStateEnumIgnoreCase := make(map[string]ListDetectorRecipeDetectorRulesLifecycleStateEnum)
	for k, v := range mappingListDetectorRecipeDetectorRulesLifecycleStateEnum {
		mappingListDetectorRecipeDetectorRulesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDetectorRecipeDetectorRulesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectorRecipeDetectorRulesSortOrderEnum Enum with underlying type: string
type ListDetectorRecipeDetectorRulesSortOrderEnum string

// Set of constants representing the allowable values for ListDetectorRecipeDetectorRulesSortOrderEnum
const (
	ListDetectorRecipeDetectorRulesSortOrderAsc  ListDetectorRecipeDetectorRulesSortOrderEnum = "ASC"
	ListDetectorRecipeDetectorRulesSortOrderDesc ListDetectorRecipeDetectorRulesSortOrderEnum = "DESC"
)

var mappingListDetectorRecipeDetectorRulesSortOrderEnum = map[string]ListDetectorRecipeDetectorRulesSortOrderEnum{
	"ASC":  ListDetectorRecipeDetectorRulesSortOrderAsc,
	"DESC": ListDetectorRecipeDetectorRulesSortOrderDesc,
}

// GetListDetectorRecipeDetectorRulesSortOrderEnumValues Enumerates the set of values for ListDetectorRecipeDetectorRulesSortOrderEnum
func GetListDetectorRecipeDetectorRulesSortOrderEnumValues() []ListDetectorRecipeDetectorRulesSortOrderEnum {
	values := make([]ListDetectorRecipeDetectorRulesSortOrderEnum, 0)
	for _, v := range mappingListDetectorRecipeDetectorRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipeDetectorRulesSortOrderEnumStringValues Enumerates the set of values in String for ListDetectorRecipeDetectorRulesSortOrderEnum
func GetListDetectorRecipeDetectorRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDetectorRecipeDetectorRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipeDetectorRulesSortOrderEnum(val string) (ListDetectorRecipeDetectorRulesSortOrderEnum, bool) {
	mappingListDetectorRecipeDetectorRulesSortOrderEnumIgnoreCase := make(map[string]ListDetectorRecipeDetectorRulesSortOrderEnum)
	for k, v := range mappingListDetectorRecipeDetectorRulesSortOrderEnum {
		mappingListDetectorRecipeDetectorRulesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDetectorRecipeDetectorRulesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectorRecipeDetectorRulesSortByEnum Enum with underlying type: string
type ListDetectorRecipeDetectorRulesSortByEnum string

// Set of constants representing the allowable values for ListDetectorRecipeDetectorRulesSortByEnum
const (
	ListDetectorRecipeDetectorRulesSortByDisplayname ListDetectorRecipeDetectorRulesSortByEnum = "displayName"
	ListDetectorRecipeDetectorRulesSortByRisklevel   ListDetectorRecipeDetectorRulesSortByEnum = "riskLevel"
)

var mappingListDetectorRecipeDetectorRulesSortByEnum = map[string]ListDetectorRecipeDetectorRulesSortByEnum{
	"displayName": ListDetectorRecipeDetectorRulesSortByDisplayname,
	"riskLevel":   ListDetectorRecipeDetectorRulesSortByRisklevel,
}

// GetListDetectorRecipeDetectorRulesSortByEnumValues Enumerates the set of values for ListDetectorRecipeDetectorRulesSortByEnum
func GetListDetectorRecipeDetectorRulesSortByEnumValues() []ListDetectorRecipeDetectorRulesSortByEnum {
	values := make([]ListDetectorRecipeDetectorRulesSortByEnum, 0)
	for _, v := range mappingListDetectorRecipeDetectorRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipeDetectorRulesSortByEnumStringValues Enumerates the set of values in String for ListDetectorRecipeDetectorRulesSortByEnum
func GetListDetectorRecipeDetectorRulesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"riskLevel",
	}
}

// GetMappingListDetectorRecipeDetectorRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipeDetectorRulesSortByEnum(val string) (ListDetectorRecipeDetectorRulesSortByEnum, bool) {
	mappingListDetectorRecipeDetectorRulesSortByEnumIgnoreCase := make(map[string]ListDetectorRecipeDetectorRulesSortByEnum)
	for k, v := range mappingListDetectorRecipeDetectorRulesSortByEnum {
		mappingListDetectorRecipeDetectorRulesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDetectorRecipeDetectorRulesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
