// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDelegationControlsRequest wrapper for the ListDelegationControls operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegationControls.go.html to see an example of how to use ListDelegationControlsRequest.
type ListDelegationControlsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only Delegation Control resources whose lifecycleState matches the given Delegation Control lifecycle state.
	LifecycleState DelegationControlLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Delegation Control resources that match the given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given resource type.
	ResourceType ListDelegationControlsResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"resourceType" omitEmpty:"true"`

	// A filter to return Delegation Control resources that match the given resource ID.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDelegationControlsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListDelegationControlsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDelegationControlsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDelegationControlsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDelegationControlsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDelegationControlsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDelegationControlsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDelegationControlLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDelegationControlLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegationControlsResourceTypeEnum(string(request.ResourceType)); !ok && request.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", request.ResourceType, strings.Join(GetListDelegationControlsResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegationControlsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDelegationControlsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegationControlsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDelegationControlsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDelegationControlsResponse wrapper for the ListDelegationControls operation
type ListDelegationControlsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DelegationControlSummaryCollection instances
	DelegationControlSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDelegationControlsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDelegationControlsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDelegationControlsResourceTypeEnum Enum with underlying type: string
type ListDelegationControlsResourceTypeEnum string

// Set of constants representing the allowable values for ListDelegationControlsResourceTypeEnum
const (
	ListDelegationControlsResourceTypeVmcluster      ListDelegationControlsResourceTypeEnum = "VMCLUSTER"
	ListDelegationControlsResourceTypeCloudvmcluster ListDelegationControlsResourceTypeEnum = "CLOUDVMCLUSTER"
)

var mappingListDelegationControlsResourceTypeEnum = map[string]ListDelegationControlsResourceTypeEnum{
	"VMCLUSTER":      ListDelegationControlsResourceTypeVmcluster,
	"CLOUDVMCLUSTER": ListDelegationControlsResourceTypeCloudvmcluster,
}

var mappingListDelegationControlsResourceTypeEnumLowerCase = map[string]ListDelegationControlsResourceTypeEnum{
	"vmcluster":      ListDelegationControlsResourceTypeVmcluster,
	"cloudvmcluster": ListDelegationControlsResourceTypeCloudvmcluster,
}

// GetListDelegationControlsResourceTypeEnumValues Enumerates the set of values for ListDelegationControlsResourceTypeEnum
func GetListDelegationControlsResourceTypeEnumValues() []ListDelegationControlsResourceTypeEnum {
	values := make([]ListDelegationControlsResourceTypeEnum, 0)
	for _, v := range mappingListDelegationControlsResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegationControlsResourceTypeEnumStringValues Enumerates the set of values in String for ListDelegationControlsResourceTypeEnum
func GetListDelegationControlsResourceTypeEnumStringValues() []string {
	return []string{
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
	}
}

// GetMappingListDelegationControlsResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegationControlsResourceTypeEnum(val string) (ListDelegationControlsResourceTypeEnum, bool) {
	enum, ok := mappingListDelegationControlsResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDelegationControlsSortOrderEnum Enum with underlying type: string
type ListDelegationControlsSortOrderEnum string

// Set of constants representing the allowable values for ListDelegationControlsSortOrderEnum
const (
	ListDelegationControlsSortOrderAsc  ListDelegationControlsSortOrderEnum = "ASC"
	ListDelegationControlsSortOrderDesc ListDelegationControlsSortOrderEnum = "DESC"
)

var mappingListDelegationControlsSortOrderEnum = map[string]ListDelegationControlsSortOrderEnum{
	"ASC":  ListDelegationControlsSortOrderAsc,
	"DESC": ListDelegationControlsSortOrderDesc,
}

var mappingListDelegationControlsSortOrderEnumLowerCase = map[string]ListDelegationControlsSortOrderEnum{
	"asc":  ListDelegationControlsSortOrderAsc,
	"desc": ListDelegationControlsSortOrderDesc,
}

// GetListDelegationControlsSortOrderEnumValues Enumerates the set of values for ListDelegationControlsSortOrderEnum
func GetListDelegationControlsSortOrderEnumValues() []ListDelegationControlsSortOrderEnum {
	values := make([]ListDelegationControlsSortOrderEnum, 0)
	for _, v := range mappingListDelegationControlsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegationControlsSortOrderEnumStringValues Enumerates the set of values in String for ListDelegationControlsSortOrderEnum
func GetListDelegationControlsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDelegationControlsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegationControlsSortOrderEnum(val string) (ListDelegationControlsSortOrderEnum, bool) {
	enum, ok := mappingListDelegationControlsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDelegationControlsSortByEnum Enum with underlying type: string
type ListDelegationControlsSortByEnum string

// Set of constants representing the allowable values for ListDelegationControlsSortByEnum
const (
	ListDelegationControlsSortByTimecreated ListDelegationControlsSortByEnum = "timeCreated"
	ListDelegationControlsSortByDisplayname ListDelegationControlsSortByEnum = "displayName"
)

var mappingListDelegationControlsSortByEnum = map[string]ListDelegationControlsSortByEnum{
	"timeCreated": ListDelegationControlsSortByTimecreated,
	"displayName": ListDelegationControlsSortByDisplayname,
}

var mappingListDelegationControlsSortByEnumLowerCase = map[string]ListDelegationControlsSortByEnum{
	"timecreated": ListDelegationControlsSortByTimecreated,
	"displayname": ListDelegationControlsSortByDisplayname,
}

// GetListDelegationControlsSortByEnumValues Enumerates the set of values for ListDelegationControlsSortByEnum
func GetListDelegationControlsSortByEnumValues() []ListDelegationControlsSortByEnum {
	values := make([]ListDelegationControlsSortByEnum, 0)
	for _, v := range mappingListDelegationControlsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegationControlsSortByEnumStringValues Enumerates the set of values in String for ListDelegationControlsSortByEnum
func GetListDelegationControlsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDelegationControlsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegationControlsSortByEnum(val string) (ListDelegationControlsSortByEnum, bool) {
	enum, ok := mappingListDelegationControlsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
