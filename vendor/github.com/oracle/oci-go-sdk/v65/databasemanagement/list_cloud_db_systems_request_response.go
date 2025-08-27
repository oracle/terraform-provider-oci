// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCloudDbSystemsRequest wrapper for the ListCloudDbSystems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudDbSystems.go.html to see an example of how to use ListCloudDbSystemsRequest.
type ListCloudDbSystemsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dbaas parent infrastructure of the cloud DB system.
	DbaasParentInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"dbaasParentInfrastructureId"`

	// A filter to return cloud DB systems of the specified deployment type.
	DeploymentType ListCloudDbSystemsDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"deploymentType" omitEmpty:"true"`

	// A filter to only return the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The lifecycle state of a resource.
	LifecycleState ListCloudDbSystemsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for `TIMECREATED` is descending and the default sort order for `DISPLAYNAME` is ascending.
	// The `DISPLAYNAME` sort order is case-sensitive.
	SortBy ListCloudDbSystemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudDbSystemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudDbSystemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudDbSystemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudDbSystemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudDbSystemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudDbSystemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudDbSystemsDeploymentTypeEnum(string(request.DeploymentType)); !ok && request.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", request.DeploymentType, strings.Join(GetListCloudDbSystemsDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudDbSystemsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListCloudDbSystemsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudDbSystemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudDbSystemsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudDbSystemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudDbSystemsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudDbSystemsResponse wrapper for the ListCloudDbSystems operation
type ListCloudDbSystemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudDbSystemCollection instances
	CloudDbSystemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudDbSystemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudDbSystemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudDbSystemsDeploymentTypeEnum Enum with underlying type: string
type ListCloudDbSystemsDeploymentTypeEnum string

// Set of constants representing the allowable values for ListCloudDbSystemsDeploymentTypeEnum
const (
	ListCloudDbSystemsDeploymentTypeVm        ListCloudDbSystemsDeploymentTypeEnum = "VM"
	ListCloudDbSystemsDeploymentTypeExadata   ListCloudDbSystemsDeploymentTypeEnum = "EXADATA"
	ListCloudDbSystemsDeploymentTypeExadataCc ListCloudDbSystemsDeploymentTypeEnum = "EXADATA_CC"
	ListCloudDbSystemsDeploymentTypeExadataXs ListCloudDbSystemsDeploymentTypeEnum = "EXADATA_XS"
)

var mappingListCloudDbSystemsDeploymentTypeEnum = map[string]ListCloudDbSystemsDeploymentTypeEnum{
	"VM":         ListCloudDbSystemsDeploymentTypeVm,
	"EXADATA":    ListCloudDbSystemsDeploymentTypeExadata,
	"EXADATA_CC": ListCloudDbSystemsDeploymentTypeExadataCc,
	"EXADATA_XS": ListCloudDbSystemsDeploymentTypeExadataXs,
}

var mappingListCloudDbSystemsDeploymentTypeEnumLowerCase = map[string]ListCloudDbSystemsDeploymentTypeEnum{
	"vm":         ListCloudDbSystemsDeploymentTypeVm,
	"exadata":    ListCloudDbSystemsDeploymentTypeExadata,
	"exadata_cc": ListCloudDbSystemsDeploymentTypeExadataCc,
	"exadata_xs": ListCloudDbSystemsDeploymentTypeExadataXs,
}

// GetListCloudDbSystemsDeploymentTypeEnumValues Enumerates the set of values for ListCloudDbSystemsDeploymentTypeEnum
func GetListCloudDbSystemsDeploymentTypeEnumValues() []ListCloudDbSystemsDeploymentTypeEnum {
	values := make([]ListCloudDbSystemsDeploymentTypeEnum, 0)
	for _, v := range mappingListCloudDbSystemsDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbSystemsDeploymentTypeEnumStringValues Enumerates the set of values in String for ListCloudDbSystemsDeploymentTypeEnum
func GetListCloudDbSystemsDeploymentTypeEnumStringValues() []string {
	return []string{
		"VM",
		"EXADATA",
		"EXADATA_CC",
		"EXADATA_XS",
	}
}

// GetMappingListCloudDbSystemsDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbSystemsDeploymentTypeEnum(val string) (ListCloudDbSystemsDeploymentTypeEnum, bool) {
	enum, ok := mappingListCloudDbSystemsDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudDbSystemsLifecycleStateEnum Enum with underlying type: string
type ListCloudDbSystemsLifecycleStateEnum string

// Set of constants representing the allowable values for ListCloudDbSystemsLifecycleStateEnum
const (
	ListCloudDbSystemsLifecycleStateCreating ListCloudDbSystemsLifecycleStateEnum = "CREATING"
	ListCloudDbSystemsLifecycleStateActive   ListCloudDbSystemsLifecycleStateEnum = "ACTIVE"
	ListCloudDbSystemsLifecycleStateUpdating ListCloudDbSystemsLifecycleStateEnum = "UPDATING"
	ListCloudDbSystemsLifecycleStateDeleting ListCloudDbSystemsLifecycleStateEnum = "DELETING"
	ListCloudDbSystemsLifecycleStateDeleted  ListCloudDbSystemsLifecycleStateEnum = "DELETED"
	ListCloudDbSystemsLifecycleStateInactive ListCloudDbSystemsLifecycleStateEnum = "INACTIVE"
)

var mappingListCloudDbSystemsLifecycleStateEnum = map[string]ListCloudDbSystemsLifecycleStateEnum{
	"CREATING": ListCloudDbSystemsLifecycleStateCreating,
	"ACTIVE":   ListCloudDbSystemsLifecycleStateActive,
	"UPDATING": ListCloudDbSystemsLifecycleStateUpdating,
	"DELETING": ListCloudDbSystemsLifecycleStateDeleting,
	"DELETED":  ListCloudDbSystemsLifecycleStateDeleted,
	"INACTIVE": ListCloudDbSystemsLifecycleStateInactive,
}

var mappingListCloudDbSystemsLifecycleStateEnumLowerCase = map[string]ListCloudDbSystemsLifecycleStateEnum{
	"creating": ListCloudDbSystemsLifecycleStateCreating,
	"active":   ListCloudDbSystemsLifecycleStateActive,
	"updating": ListCloudDbSystemsLifecycleStateUpdating,
	"deleting": ListCloudDbSystemsLifecycleStateDeleting,
	"deleted":  ListCloudDbSystemsLifecycleStateDeleted,
	"inactive": ListCloudDbSystemsLifecycleStateInactive,
}

// GetListCloudDbSystemsLifecycleStateEnumValues Enumerates the set of values for ListCloudDbSystemsLifecycleStateEnum
func GetListCloudDbSystemsLifecycleStateEnumValues() []ListCloudDbSystemsLifecycleStateEnum {
	values := make([]ListCloudDbSystemsLifecycleStateEnum, 0)
	for _, v := range mappingListCloudDbSystemsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbSystemsLifecycleStateEnumStringValues Enumerates the set of values in String for ListCloudDbSystemsLifecycleStateEnum
func GetListCloudDbSystemsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"INACTIVE",
	}
}

// GetMappingListCloudDbSystemsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbSystemsLifecycleStateEnum(val string) (ListCloudDbSystemsLifecycleStateEnum, bool) {
	enum, ok := mappingListCloudDbSystemsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudDbSystemsSortByEnum Enum with underlying type: string
type ListCloudDbSystemsSortByEnum string

// Set of constants representing the allowable values for ListCloudDbSystemsSortByEnum
const (
	ListCloudDbSystemsSortByTimecreated ListCloudDbSystemsSortByEnum = "TIMECREATED"
	ListCloudDbSystemsSortByDisplayname ListCloudDbSystemsSortByEnum = "DISPLAYNAME"
)

var mappingListCloudDbSystemsSortByEnum = map[string]ListCloudDbSystemsSortByEnum{
	"TIMECREATED": ListCloudDbSystemsSortByTimecreated,
	"DISPLAYNAME": ListCloudDbSystemsSortByDisplayname,
}

var mappingListCloudDbSystemsSortByEnumLowerCase = map[string]ListCloudDbSystemsSortByEnum{
	"timecreated": ListCloudDbSystemsSortByTimecreated,
	"displayname": ListCloudDbSystemsSortByDisplayname,
}

// GetListCloudDbSystemsSortByEnumValues Enumerates the set of values for ListCloudDbSystemsSortByEnum
func GetListCloudDbSystemsSortByEnumValues() []ListCloudDbSystemsSortByEnum {
	values := make([]ListCloudDbSystemsSortByEnum, 0)
	for _, v := range mappingListCloudDbSystemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbSystemsSortByEnumStringValues Enumerates the set of values in String for ListCloudDbSystemsSortByEnum
func GetListCloudDbSystemsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudDbSystemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbSystemsSortByEnum(val string) (ListCloudDbSystemsSortByEnum, bool) {
	enum, ok := mappingListCloudDbSystemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudDbSystemsSortOrderEnum Enum with underlying type: string
type ListCloudDbSystemsSortOrderEnum string

// Set of constants representing the allowable values for ListCloudDbSystemsSortOrderEnum
const (
	ListCloudDbSystemsSortOrderAsc  ListCloudDbSystemsSortOrderEnum = "ASC"
	ListCloudDbSystemsSortOrderDesc ListCloudDbSystemsSortOrderEnum = "DESC"
)

var mappingListCloudDbSystemsSortOrderEnum = map[string]ListCloudDbSystemsSortOrderEnum{
	"ASC":  ListCloudDbSystemsSortOrderAsc,
	"DESC": ListCloudDbSystemsSortOrderDesc,
}

var mappingListCloudDbSystemsSortOrderEnumLowerCase = map[string]ListCloudDbSystemsSortOrderEnum{
	"asc":  ListCloudDbSystemsSortOrderAsc,
	"desc": ListCloudDbSystemsSortOrderDesc,
}

// GetListCloudDbSystemsSortOrderEnumValues Enumerates the set of values for ListCloudDbSystemsSortOrderEnum
func GetListCloudDbSystemsSortOrderEnumValues() []ListCloudDbSystemsSortOrderEnum {
	values := make([]ListCloudDbSystemsSortOrderEnum, 0)
	for _, v := range mappingListCloudDbSystemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudDbSystemsSortOrderEnumStringValues Enumerates the set of values in String for ListCloudDbSystemsSortOrderEnum
func GetListCloudDbSystemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudDbSystemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudDbSystemsSortOrderEnum(val string) (ListCloudDbSystemsSortOrderEnum, bool) {
	enum, ok := mappingListCloudDbSystemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
