// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMultiCloudResourceDiscoveriesRequest wrapper for the ListMultiCloudResourceDiscoveries operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListMultiCloudResourceDiscoveries.go.html to see an example of how to use ListMultiCloudResourceDiscoveriesRequest.
type ListMultiCloudResourceDiscoveriesRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Display Name of the Multi Cloud Discovery Resource.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multi Cloud Discovery Resource.
	MultiCloudResourceDiscoveryId *string `mandatory:"false" contributesTo:"query" name:"multiCloudResourceDiscoveryId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState MultiCloudResourceDiscoveryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Oracle DB Azure Blob Mount Resources.
	OracleDbAzureConnectorId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureConnectorId"`

	// The type of Multi Cloud Resource.
	ResourceType MultiCloudResourceDiscoveryResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"resourceType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMultiCloudResourceDiscoveriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListMultiCloudResourceDiscoveriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMultiCloudResourceDiscoveriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMultiCloudResourceDiscoveriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMultiCloudResourceDiscoveriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMultiCloudResourceDiscoveriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMultiCloudResourceDiscoveriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMultiCloudResourceDiscoveryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMultiCloudResourceDiscoveryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMultiCloudResourceDiscoveryResourceTypeEnum(string(request.ResourceType)); !ok && request.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", request.ResourceType, strings.Join(GetMultiCloudResourceDiscoveryResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMultiCloudResourceDiscoveriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMultiCloudResourceDiscoveriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMultiCloudResourceDiscoveriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMultiCloudResourceDiscoveriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMultiCloudResourceDiscoveriesResponse wrapper for the ListMultiCloudResourceDiscoveries operation
type ListMultiCloudResourceDiscoveriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MultiCloudResourceDiscoverySummaryCollection instances
	MultiCloudResourceDiscoverySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMultiCloudResourceDiscoveriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMultiCloudResourceDiscoveriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMultiCloudResourceDiscoveriesSortOrderEnum Enum with underlying type: string
type ListMultiCloudResourceDiscoveriesSortOrderEnum string

// Set of constants representing the allowable values for ListMultiCloudResourceDiscoveriesSortOrderEnum
const (
	ListMultiCloudResourceDiscoveriesSortOrderAsc  ListMultiCloudResourceDiscoveriesSortOrderEnum = "ASC"
	ListMultiCloudResourceDiscoveriesSortOrderDesc ListMultiCloudResourceDiscoveriesSortOrderEnum = "DESC"
)

var mappingListMultiCloudResourceDiscoveriesSortOrderEnum = map[string]ListMultiCloudResourceDiscoveriesSortOrderEnum{
	"ASC":  ListMultiCloudResourceDiscoveriesSortOrderAsc,
	"DESC": ListMultiCloudResourceDiscoveriesSortOrderDesc,
}

var mappingListMultiCloudResourceDiscoveriesSortOrderEnumLowerCase = map[string]ListMultiCloudResourceDiscoveriesSortOrderEnum{
	"asc":  ListMultiCloudResourceDiscoveriesSortOrderAsc,
	"desc": ListMultiCloudResourceDiscoveriesSortOrderDesc,
}

// GetListMultiCloudResourceDiscoveriesSortOrderEnumValues Enumerates the set of values for ListMultiCloudResourceDiscoveriesSortOrderEnum
func GetListMultiCloudResourceDiscoveriesSortOrderEnumValues() []ListMultiCloudResourceDiscoveriesSortOrderEnum {
	values := make([]ListMultiCloudResourceDiscoveriesSortOrderEnum, 0)
	for _, v := range mappingListMultiCloudResourceDiscoveriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMultiCloudResourceDiscoveriesSortOrderEnumStringValues Enumerates the set of values in String for ListMultiCloudResourceDiscoveriesSortOrderEnum
func GetListMultiCloudResourceDiscoveriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMultiCloudResourceDiscoveriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMultiCloudResourceDiscoveriesSortOrderEnum(val string) (ListMultiCloudResourceDiscoveriesSortOrderEnum, bool) {
	enum, ok := mappingListMultiCloudResourceDiscoveriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMultiCloudResourceDiscoveriesSortByEnum Enum with underlying type: string
type ListMultiCloudResourceDiscoveriesSortByEnum string

// Set of constants representing the allowable values for ListMultiCloudResourceDiscoveriesSortByEnum
const (
	ListMultiCloudResourceDiscoveriesSortByTimecreated ListMultiCloudResourceDiscoveriesSortByEnum = "timeCreated"
	ListMultiCloudResourceDiscoveriesSortByDisplayname ListMultiCloudResourceDiscoveriesSortByEnum = "displayName"
)

var mappingListMultiCloudResourceDiscoveriesSortByEnum = map[string]ListMultiCloudResourceDiscoveriesSortByEnum{
	"timeCreated": ListMultiCloudResourceDiscoveriesSortByTimecreated,
	"displayName": ListMultiCloudResourceDiscoveriesSortByDisplayname,
}

var mappingListMultiCloudResourceDiscoveriesSortByEnumLowerCase = map[string]ListMultiCloudResourceDiscoveriesSortByEnum{
	"timecreated": ListMultiCloudResourceDiscoveriesSortByTimecreated,
	"displayname": ListMultiCloudResourceDiscoveriesSortByDisplayname,
}

// GetListMultiCloudResourceDiscoveriesSortByEnumValues Enumerates the set of values for ListMultiCloudResourceDiscoveriesSortByEnum
func GetListMultiCloudResourceDiscoveriesSortByEnumValues() []ListMultiCloudResourceDiscoveriesSortByEnum {
	values := make([]ListMultiCloudResourceDiscoveriesSortByEnum, 0)
	for _, v := range mappingListMultiCloudResourceDiscoveriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMultiCloudResourceDiscoveriesSortByEnumStringValues Enumerates the set of values in String for ListMultiCloudResourceDiscoveriesSortByEnum
func GetListMultiCloudResourceDiscoveriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMultiCloudResourceDiscoveriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMultiCloudResourceDiscoveriesSortByEnum(val string) (ListMultiCloudResourceDiscoveriesSortByEnum, bool) {
	enum, ok := mappingListMultiCloudResourceDiscoveriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
