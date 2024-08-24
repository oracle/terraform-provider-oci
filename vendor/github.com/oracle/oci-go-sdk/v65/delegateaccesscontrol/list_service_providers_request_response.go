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

// ListServiceProvidersRequest wrapper for the ListServiceProviders operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListServiceProviders.go.html to see an example of how to use ListServiceProvidersRequest.
type ListServiceProvidersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only Service Provider resources whose lifecycleState matches the given Service Provider lifecycle state.
	LifecycleState ServiceProviderLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Service Provider resources that match the given name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only Service Provider resources whose supported resource type matches the given resource type.
	SupportedResourceType ListServiceProvidersSupportedResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"supportedResourceType" omitEmpty:"true"`

	// A filter to return only Service Provider resources whose provider type matches the given provider type.
	ServiceProviderType ServiceProviderServiceProviderTypeEnum `mandatory:"false" contributesTo:"query" name:"serviceProviderType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListServiceProvidersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending. If no value is specified, timeCreated is default.
	SortBy ListServiceProvidersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServiceProvidersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServiceProvidersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListServiceProvidersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServiceProvidersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListServiceProvidersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceProviderLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetServiceProviderLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceProvidersSupportedResourceTypeEnum(string(request.SupportedResourceType)); !ok && request.SupportedResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedResourceType: %s. Supported values are: %s.", request.SupportedResourceType, strings.Join(GetListServiceProvidersSupportedResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingServiceProviderServiceProviderTypeEnum(string(request.ServiceProviderType)); !ok && request.ServiceProviderType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceProviderType: %s. Supported values are: %s.", request.ServiceProviderType, strings.Join(GetServiceProviderServiceProviderTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceProvidersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListServiceProvidersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceProvidersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListServiceProvidersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListServiceProvidersResponse wrapper for the ListServiceProviders operation
type ListServiceProvidersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServiceProviderSummaryCollection instances
	ServiceProviderSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListServiceProvidersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServiceProvidersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServiceProvidersSupportedResourceTypeEnum Enum with underlying type: string
type ListServiceProvidersSupportedResourceTypeEnum string

// Set of constants representing the allowable values for ListServiceProvidersSupportedResourceTypeEnum
const (
	ListServiceProvidersSupportedResourceTypeVmcluster      ListServiceProvidersSupportedResourceTypeEnum = "VMCLUSTER"
	ListServiceProvidersSupportedResourceTypeCloudvmcluster ListServiceProvidersSupportedResourceTypeEnum = "CLOUDVMCLUSTER"
)

var mappingListServiceProvidersSupportedResourceTypeEnum = map[string]ListServiceProvidersSupportedResourceTypeEnum{
	"VMCLUSTER":      ListServiceProvidersSupportedResourceTypeVmcluster,
	"CLOUDVMCLUSTER": ListServiceProvidersSupportedResourceTypeCloudvmcluster,
}

var mappingListServiceProvidersSupportedResourceTypeEnumLowerCase = map[string]ListServiceProvidersSupportedResourceTypeEnum{
	"vmcluster":      ListServiceProvidersSupportedResourceTypeVmcluster,
	"cloudvmcluster": ListServiceProvidersSupportedResourceTypeCloudvmcluster,
}

// GetListServiceProvidersSupportedResourceTypeEnumValues Enumerates the set of values for ListServiceProvidersSupportedResourceTypeEnum
func GetListServiceProvidersSupportedResourceTypeEnumValues() []ListServiceProvidersSupportedResourceTypeEnum {
	values := make([]ListServiceProvidersSupportedResourceTypeEnum, 0)
	for _, v := range mappingListServiceProvidersSupportedResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceProvidersSupportedResourceTypeEnumStringValues Enumerates the set of values in String for ListServiceProvidersSupportedResourceTypeEnum
func GetListServiceProvidersSupportedResourceTypeEnumStringValues() []string {
	return []string{
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
	}
}

// GetMappingListServiceProvidersSupportedResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceProvidersSupportedResourceTypeEnum(val string) (ListServiceProvidersSupportedResourceTypeEnum, bool) {
	enum, ok := mappingListServiceProvidersSupportedResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceProvidersSortOrderEnum Enum with underlying type: string
type ListServiceProvidersSortOrderEnum string

// Set of constants representing the allowable values for ListServiceProvidersSortOrderEnum
const (
	ListServiceProvidersSortOrderAsc  ListServiceProvidersSortOrderEnum = "ASC"
	ListServiceProvidersSortOrderDesc ListServiceProvidersSortOrderEnum = "DESC"
)

var mappingListServiceProvidersSortOrderEnum = map[string]ListServiceProvidersSortOrderEnum{
	"ASC":  ListServiceProvidersSortOrderAsc,
	"DESC": ListServiceProvidersSortOrderDesc,
}

var mappingListServiceProvidersSortOrderEnumLowerCase = map[string]ListServiceProvidersSortOrderEnum{
	"asc":  ListServiceProvidersSortOrderAsc,
	"desc": ListServiceProvidersSortOrderDesc,
}

// GetListServiceProvidersSortOrderEnumValues Enumerates the set of values for ListServiceProvidersSortOrderEnum
func GetListServiceProvidersSortOrderEnumValues() []ListServiceProvidersSortOrderEnum {
	values := make([]ListServiceProvidersSortOrderEnum, 0)
	for _, v := range mappingListServiceProvidersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceProvidersSortOrderEnumStringValues Enumerates the set of values in String for ListServiceProvidersSortOrderEnum
func GetListServiceProvidersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListServiceProvidersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceProvidersSortOrderEnum(val string) (ListServiceProvidersSortOrderEnum, bool) {
	enum, ok := mappingListServiceProvidersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceProvidersSortByEnum Enum with underlying type: string
type ListServiceProvidersSortByEnum string

// Set of constants representing the allowable values for ListServiceProvidersSortByEnum
const (
	ListServiceProvidersSortByTimecreated ListServiceProvidersSortByEnum = "timeCreated"
	ListServiceProvidersSortByDisplayname ListServiceProvidersSortByEnum = "displayName"
)

var mappingListServiceProvidersSortByEnum = map[string]ListServiceProvidersSortByEnum{
	"timeCreated": ListServiceProvidersSortByTimecreated,
	"displayName": ListServiceProvidersSortByDisplayname,
}

var mappingListServiceProvidersSortByEnumLowerCase = map[string]ListServiceProvidersSortByEnum{
	"timecreated": ListServiceProvidersSortByTimecreated,
	"displayname": ListServiceProvidersSortByDisplayname,
}

// GetListServiceProvidersSortByEnumValues Enumerates the set of values for ListServiceProvidersSortByEnum
func GetListServiceProvidersSortByEnumValues() []ListServiceProvidersSortByEnum {
	values := make([]ListServiceProvidersSortByEnum, 0)
	for _, v := range mappingListServiceProvidersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceProvidersSortByEnumStringValues Enumerates the set of values in String for ListServiceProvidersSortByEnum
func GetListServiceProvidersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListServiceProvidersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceProvidersSortByEnum(val string) (ListServiceProvidersSortByEnum, bool) {
	enum, ok := mappingListServiceProvidersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
