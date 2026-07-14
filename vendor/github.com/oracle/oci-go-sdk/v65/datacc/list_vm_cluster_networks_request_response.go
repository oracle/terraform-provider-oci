// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVmClusterNetworksRequest wrapper for the ListVmClusterNetworks operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListVmClusterNetworks.go.html to see an example of how to use ListVmClusterNetworksRequest.
type ListVmClusterNetworksRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	// For list operations, you may provide the tenant [OCID] in this field. When a tenant OCID is provided,
	// it will be validated against the caller's tenant and then treated as tenant scope (compartmentId filtering is not applied).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the entire display name given. The match is case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that match the specified lifecycle state.
	LifecycleState []VmClusterNetworkLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return VM cluster network resources that matches the specified value.
	IsScanEnabled *bool `mandatory:"false" contributesTo:"query" name:"isScanEnabled"`

	// The maximum number of items to return in a page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which you want to start retrieving results. This token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order that you want to use, which is either `ASC` or `DESC`.
	SortOrder ListVmClusterNetworksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which you want to sort. You can provide only one type of sort order. The default order for `timeCreated` is descending. The default order for `displayName` is ascending. If no value is specified, then `timeCreated` is the default.
	// When listing software images within the same `version`, using `sortBy=buildIdentifier` is recommended. `buildIdentifier` is a monotonically increasing, time-ordered string marker (yyyy-mm-dd-hh:mm:ss) stored with the image.
	SortBy ListVmClusterNetworksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
	InfrastructureId *string `mandatory:"false" contributesTo:"query" name:"infrastructureId"`

	// Count of virtual machines in this VM cluster.
	NodeCount *int `mandatory:"false" contributesTo:"query" name:"nodeCount"`

	// VM network consumer type.
	VmNetworkConsumerType ListVmClusterNetworksVmNetworkConsumerTypeEnum `mandatory:"false" contributesTo:"query" name:"vmNetworkConsumerType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVmClusterNetworksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVmClusterNetworksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVmClusterNetworksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVmClusterNetworksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVmClusterNetworksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingVmClusterNetworkLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetVmClusterNetworkLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListVmClusterNetworksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVmClusterNetworksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVmClusterNetworksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVmClusterNetworksSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVmClusterNetworksVmNetworkConsumerTypeEnum(string(request.VmNetworkConsumerType)); !ok && request.VmNetworkConsumerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VmNetworkConsumerType: %s. Supported values are: %s.", request.VmNetworkConsumerType, strings.Join(GetListVmClusterNetworksVmNetworkConsumerTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVmClusterNetworksResponse wrapper for the ListVmClusterNetworks operation
type ListVmClusterNetworksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VmClusterNetworkCollection instances
	VmClusterNetworkCollection `presentIn:"body"`

	// Unique identifier assigned by Oracle for the request. If you need to contact
	// Oracle about a particular request, then please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then it can mean that a partial list was returned. To obtain the next batch of items, include this value as the `page` parameter for your next GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVmClusterNetworksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVmClusterNetworksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVmClusterNetworksSortOrderEnum Enum with underlying type: string
type ListVmClusterNetworksSortOrderEnum string

// Set of constants representing the allowable values for ListVmClusterNetworksSortOrderEnum
const (
	ListVmClusterNetworksSortOrderAsc  ListVmClusterNetworksSortOrderEnum = "ASC"
	ListVmClusterNetworksSortOrderDesc ListVmClusterNetworksSortOrderEnum = "DESC"
)

var mappingListVmClusterNetworksSortOrderEnum = map[string]ListVmClusterNetworksSortOrderEnum{
	"ASC":  ListVmClusterNetworksSortOrderAsc,
	"DESC": ListVmClusterNetworksSortOrderDesc,
}

var mappingListVmClusterNetworksSortOrderEnumLowerCase = map[string]ListVmClusterNetworksSortOrderEnum{
	"asc":  ListVmClusterNetworksSortOrderAsc,
	"desc": ListVmClusterNetworksSortOrderDesc,
}

// GetListVmClusterNetworksSortOrderEnumValues Enumerates the set of values for ListVmClusterNetworksSortOrderEnum
func GetListVmClusterNetworksSortOrderEnumValues() []ListVmClusterNetworksSortOrderEnum {
	values := make([]ListVmClusterNetworksSortOrderEnum, 0)
	for _, v := range mappingListVmClusterNetworksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmClusterNetworksSortOrderEnumStringValues Enumerates the set of values in String for ListVmClusterNetworksSortOrderEnum
func GetListVmClusterNetworksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVmClusterNetworksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmClusterNetworksSortOrderEnum(val string) (ListVmClusterNetworksSortOrderEnum, bool) {
	enum, ok := mappingListVmClusterNetworksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVmClusterNetworksSortByEnum Enum with underlying type: string
type ListVmClusterNetworksSortByEnum string

// Set of constants representing the allowable values for ListVmClusterNetworksSortByEnum
const (
	ListVmClusterNetworksSortByTimecreated     ListVmClusterNetworksSortByEnum = "timeCreated"
	ListVmClusterNetworksSortByDisplayname     ListVmClusterNetworksSortByEnum = "displayName"
	ListVmClusterNetworksSortByBuildidentifier ListVmClusterNetworksSortByEnum = "buildIdentifier"
)

var mappingListVmClusterNetworksSortByEnum = map[string]ListVmClusterNetworksSortByEnum{
	"timeCreated":     ListVmClusterNetworksSortByTimecreated,
	"displayName":     ListVmClusterNetworksSortByDisplayname,
	"buildIdentifier": ListVmClusterNetworksSortByBuildidentifier,
}

var mappingListVmClusterNetworksSortByEnumLowerCase = map[string]ListVmClusterNetworksSortByEnum{
	"timecreated":     ListVmClusterNetworksSortByTimecreated,
	"displayname":     ListVmClusterNetworksSortByDisplayname,
	"buildidentifier": ListVmClusterNetworksSortByBuildidentifier,
}

// GetListVmClusterNetworksSortByEnumValues Enumerates the set of values for ListVmClusterNetworksSortByEnum
func GetListVmClusterNetworksSortByEnumValues() []ListVmClusterNetworksSortByEnum {
	values := make([]ListVmClusterNetworksSortByEnum, 0)
	for _, v := range mappingListVmClusterNetworksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmClusterNetworksSortByEnumStringValues Enumerates the set of values in String for ListVmClusterNetworksSortByEnum
func GetListVmClusterNetworksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"buildIdentifier",
	}
}

// GetMappingListVmClusterNetworksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmClusterNetworksSortByEnum(val string) (ListVmClusterNetworksSortByEnum, bool) {
	enum, ok := mappingListVmClusterNetworksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVmClusterNetworksVmNetworkConsumerTypeEnum Enum with underlying type: string
type ListVmClusterNetworksVmNetworkConsumerTypeEnum string

// Set of constants representing the allowable values for ListVmClusterNetworksVmNetworkConsumerTypeEnum
const (
	ListVmClusterNetworksVmNetworkConsumerTypeInstance ListVmClusterNetworksVmNetworkConsumerTypeEnum = "INSTANCE"
	ListVmClusterNetworksVmNetworkConsumerTypeCluster  ListVmClusterNetworksVmNetworkConsumerTypeEnum = "CLUSTER"
)

var mappingListVmClusterNetworksVmNetworkConsumerTypeEnum = map[string]ListVmClusterNetworksVmNetworkConsumerTypeEnum{
	"INSTANCE": ListVmClusterNetworksVmNetworkConsumerTypeInstance,
	"CLUSTER":  ListVmClusterNetworksVmNetworkConsumerTypeCluster,
}

var mappingListVmClusterNetworksVmNetworkConsumerTypeEnumLowerCase = map[string]ListVmClusterNetworksVmNetworkConsumerTypeEnum{
	"instance": ListVmClusterNetworksVmNetworkConsumerTypeInstance,
	"cluster":  ListVmClusterNetworksVmNetworkConsumerTypeCluster,
}

// GetListVmClusterNetworksVmNetworkConsumerTypeEnumValues Enumerates the set of values for ListVmClusterNetworksVmNetworkConsumerTypeEnum
func GetListVmClusterNetworksVmNetworkConsumerTypeEnumValues() []ListVmClusterNetworksVmNetworkConsumerTypeEnum {
	values := make([]ListVmClusterNetworksVmNetworkConsumerTypeEnum, 0)
	for _, v := range mappingListVmClusterNetworksVmNetworkConsumerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmClusterNetworksVmNetworkConsumerTypeEnumStringValues Enumerates the set of values in String for ListVmClusterNetworksVmNetworkConsumerTypeEnum
func GetListVmClusterNetworksVmNetworkConsumerTypeEnumStringValues() []string {
	return []string{
		"INSTANCE",
		"CLUSTER",
	}
}

// GetMappingListVmClusterNetworksVmNetworkConsumerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmClusterNetworksVmNetworkConsumerTypeEnum(val string) (ListVmClusterNetworksVmNetworkConsumerTypeEnum, bool) {
	enum, ok := mappingListVmClusterNetworksVmNetworkConsumerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
