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

// ListServiceProviderActionsRequest wrapper for the ListServiceProviderActions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListServiceProviderActions.go.html to see an example of how to use ListServiceProviderActionsRequest.
type ListServiceProviderActionsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the given resource type.
	ResourceType ListServiceProviderActionsResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"resourceType" omitEmpty:"true"`

	// A filter to return only resources that match the given Service Provider service type.
	ServiceProviderServiceType []ServiceProviderServiceTypeEnum `contributesTo:"query" name:"serviceProviderServiceType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources whose lifecycleState matches the given Service Provider Action lifecycleState.
	LifecycleState ServiceProviderActionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListServiceProviderActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending. If no value is specified, default is timeCreated.
	SortBy ListServiceProviderActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServiceProviderActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServiceProviderActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListServiceProviderActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServiceProviderActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListServiceProviderActionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListServiceProviderActionsResourceTypeEnum(string(request.ResourceType)); !ok && request.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", request.ResourceType, strings.Join(GetListServiceProviderActionsResourceTypeEnumStringValues(), ",")))
	}
	for _, val := range request.ServiceProviderServiceType {
		if _, ok := GetMappingServiceProviderServiceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceProviderServiceType: %s. Supported values are: %s.", val, strings.Join(GetServiceProviderServiceTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingServiceProviderActionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetServiceProviderActionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceProviderActionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListServiceProviderActionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceProviderActionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListServiceProviderActionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListServiceProviderActionsResponse wrapper for the ListServiceProviderActions operation
type ListServiceProviderActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServiceProviderActionSummaryCollection instances
	ServiceProviderActionSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListServiceProviderActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServiceProviderActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServiceProviderActionsResourceTypeEnum Enum with underlying type: string
type ListServiceProviderActionsResourceTypeEnum string

// Set of constants representing the allowable values for ListServiceProviderActionsResourceTypeEnum
const (
	ListServiceProviderActionsResourceTypeVmcluster      ListServiceProviderActionsResourceTypeEnum = "VMCLUSTER"
	ListServiceProviderActionsResourceTypeCloudvmcluster ListServiceProviderActionsResourceTypeEnum = "CLOUDVMCLUSTER"
)

var mappingListServiceProviderActionsResourceTypeEnum = map[string]ListServiceProviderActionsResourceTypeEnum{
	"VMCLUSTER":      ListServiceProviderActionsResourceTypeVmcluster,
	"CLOUDVMCLUSTER": ListServiceProviderActionsResourceTypeCloudvmcluster,
}

var mappingListServiceProviderActionsResourceTypeEnumLowerCase = map[string]ListServiceProviderActionsResourceTypeEnum{
	"vmcluster":      ListServiceProviderActionsResourceTypeVmcluster,
	"cloudvmcluster": ListServiceProviderActionsResourceTypeCloudvmcluster,
}

// GetListServiceProviderActionsResourceTypeEnumValues Enumerates the set of values for ListServiceProviderActionsResourceTypeEnum
func GetListServiceProviderActionsResourceTypeEnumValues() []ListServiceProviderActionsResourceTypeEnum {
	values := make([]ListServiceProviderActionsResourceTypeEnum, 0)
	for _, v := range mappingListServiceProviderActionsResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceProviderActionsResourceTypeEnumStringValues Enumerates the set of values in String for ListServiceProviderActionsResourceTypeEnum
func GetListServiceProviderActionsResourceTypeEnumStringValues() []string {
	return []string{
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
	}
}

// GetMappingListServiceProviderActionsResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceProviderActionsResourceTypeEnum(val string) (ListServiceProviderActionsResourceTypeEnum, bool) {
	enum, ok := mappingListServiceProviderActionsResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceProviderActionsSortOrderEnum Enum with underlying type: string
type ListServiceProviderActionsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceProviderActionsSortOrderEnum
const (
	ListServiceProviderActionsSortOrderAsc  ListServiceProviderActionsSortOrderEnum = "ASC"
	ListServiceProviderActionsSortOrderDesc ListServiceProviderActionsSortOrderEnum = "DESC"
)

var mappingListServiceProviderActionsSortOrderEnum = map[string]ListServiceProviderActionsSortOrderEnum{
	"ASC":  ListServiceProviderActionsSortOrderAsc,
	"DESC": ListServiceProviderActionsSortOrderDesc,
}

var mappingListServiceProviderActionsSortOrderEnumLowerCase = map[string]ListServiceProviderActionsSortOrderEnum{
	"asc":  ListServiceProviderActionsSortOrderAsc,
	"desc": ListServiceProviderActionsSortOrderDesc,
}

// GetListServiceProviderActionsSortOrderEnumValues Enumerates the set of values for ListServiceProviderActionsSortOrderEnum
func GetListServiceProviderActionsSortOrderEnumValues() []ListServiceProviderActionsSortOrderEnum {
	values := make([]ListServiceProviderActionsSortOrderEnum, 0)
	for _, v := range mappingListServiceProviderActionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceProviderActionsSortOrderEnumStringValues Enumerates the set of values in String for ListServiceProviderActionsSortOrderEnum
func GetListServiceProviderActionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListServiceProviderActionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceProviderActionsSortOrderEnum(val string) (ListServiceProviderActionsSortOrderEnum, bool) {
	enum, ok := mappingListServiceProviderActionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceProviderActionsSortByEnum Enum with underlying type: string
type ListServiceProviderActionsSortByEnum string

// Set of constants representing the allowable values for ListServiceProviderActionsSortByEnum
const (
	ListServiceProviderActionsSortByTimecreated ListServiceProviderActionsSortByEnum = "timeCreated"
	ListServiceProviderActionsSortByName        ListServiceProviderActionsSortByEnum = "name"
)

var mappingListServiceProviderActionsSortByEnum = map[string]ListServiceProviderActionsSortByEnum{
	"timeCreated": ListServiceProviderActionsSortByTimecreated,
	"name":        ListServiceProviderActionsSortByName,
}

var mappingListServiceProviderActionsSortByEnumLowerCase = map[string]ListServiceProviderActionsSortByEnum{
	"timecreated": ListServiceProviderActionsSortByTimecreated,
	"name":        ListServiceProviderActionsSortByName,
}

// GetListServiceProviderActionsSortByEnumValues Enumerates the set of values for ListServiceProviderActionsSortByEnum
func GetListServiceProviderActionsSortByEnumValues() []ListServiceProviderActionsSortByEnum {
	values := make([]ListServiceProviderActionsSortByEnum, 0)
	for _, v := range mappingListServiceProviderActionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceProviderActionsSortByEnumStringValues Enumerates the set of values in String for ListServiceProviderActionsSortByEnum
func GetListServiceProviderActionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListServiceProviderActionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceProviderActionsSortByEnum(val string) (ListServiceProviderActionsSortByEnum, bool) {
	enum, ok := mappingListServiceProviderActionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
