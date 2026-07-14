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

// ListVmInstancesRequest wrapper for the ListVmInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListVmInstances.go.html to see an example of how to use ListVmInstancesRequest.
type ListVmInstancesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	// For list operations, you may provide the tenant [OCID] in this field. When a tenant OCID is provided,
	// it will be validated against the caller's tenant and then treated as tenant scope (compartmentId filtering is not applied).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return resources that match the entire display name given. The match is case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that match the specified lifecycle state.
	LifecycleState []VmInstanceLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// The maximum number of items to return in a page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which you want to start retrieving results. This token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order that you want to use, which is either `ASC` or `DESC`.
	SortOrder ListVmInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which you want to sort. You can provide only one type of sort order. The default order for `timeCreated` is descending. The default order for `displayName` is ascending. If no value is specified, then `timeCreated` is the default.
	// When listing software images within the same `version`, using `sortBy=buildIdentifier` is recommended. `buildIdentifier` is a monotonically increasing, time-ordered string marker (yyyy-mm-dd-hh:mm:ss) stored with the image.
	SortBy ListVmInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
	InfrastructureId *string `mandatory:"false" contributesTo:"query" name:"infrastructureId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure Server Id.
	BaseServerId *string `mandatory:"false" contributesTo:"query" name:"baseServerId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVmInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVmInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVmInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVmInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVmInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingVmInstanceLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetVmInstanceLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListVmInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVmInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVmInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVmInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVmInstancesResponse wrapper for the ListVmInstances operation
type ListVmInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VmInstanceCollection instances
	VmInstanceCollection `presentIn:"body"`

	// Unique identifier assigned by Oracle for the request. If you need to contact
	// Oracle about a particular request, then please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then it can mean that a partial list was returned. To obtain the next batch of items, include this value as the `page` parameter for your next GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVmInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVmInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVmInstancesSortOrderEnum Enum with underlying type: string
type ListVmInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListVmInstancesSortOrderEnum
const (
	ListVmInstancesSortOrderAsc  ListVmInstancesSortOrderEnum = "ASC"
	ListVmInstancesSortOrderDesc ListVmInstancesSortOrderEnum = "DESC"
)

var mappingListVmInstancesSortOrderEnum = map[string]ListVmInstancesSortOrderEnum{
	"ASC":  ListVmInstancesSortOrderAsc,
	"DESC": ListVmInstancesSortOrderDesc,
}

var mappingListVmInstancesSortOrderEnumLowerCase = map[string]ListVmInstancesSortOrderEnum{
	"asc":  ListVmInstancesSortOrderAsc,
	"desc": ListVmInstancesSortOrderDesc,
}

// GetListVmInstancesSortOrderEnumValues Enumerates the set of values for ListVmInstancesSortOrderEnum
func GetListVmInstancesSortOrderEnumValues() []ListVmInstancesSortOrderEnum {
	values := make([]ListVmInstancesSortOrderEnum, 0)
	for _, v := range mappingListVmInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListVmInstancesSortOrderEnum
func GetListVmInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVmInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmInstancesSortOrderEnum(val string) (ListVmInstancesSortOrderEnum, bool) {
	enum, ok := mappingListVmInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVmInstancesSortByEnum Enum with underlying type: string
type ListVmInstancesSortByEnum string

// Set of constants representing the allowable values for ListVmInstancesSortByEnum
const (
	ListVmInstancesSortByTimecreated     ListVmInstancesSortByEnum = "timeCreated"
	ListVmInstancesSortByDisplayname     ListVmInstancesSortByEnum = "displayName"
	ListVmInstancesSortByBuildidentifier ListVmInstancesSortByEnum = "buildIdentifier"
)

var mappingListVmInstancesSortByEnum = map[string]ListVmInstancesSortByEnum{
	"timeCreated":     ListVmInstancesSortByTimecreated,
	"displayName":     ListVmInstancesSortByDisplayname,
	"buildIdentifier": ListVmInstancesSortByBuildidentifier,
}

var mappingListVmInstancesSortByEnumLowerCase = map[string]ListVmInstancesSortByEnum{
	"timecreated":     ListVmInstancesSortByTimecreated,
	"displayname":     ListVmInstancesSortByDisplayname,
	"buildidentifier": ListVmInstancesSortByBuildidentifier,
}

// GetListVmInstancesSortByEnumValues Enumerates the set of values for ListVmInstancesSortByEnum
func GetListVmInstancesSortByEnumValues() []ListVmInstancesSortByEnum {
	values := make([]ListVmInstancesSortByEnum, 0)
	for _, v := range mappingListVmInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVmInstancesSortByEnumStringValues Enumerates the set of values in String for ListVmInstancesSortByEnum
func GetListVmInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"buildIdentifier",
	}
}

// GetMappingListVmInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVmInstancesSortByEnum(val string) (ListVmInstancesSortByEnum, bool) {
	enum, ok := mappingListVmInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
