// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListManagedDatabasesRequest wrapper for the ListManagedDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedDatabases.go.html to see an example of how to use ListManagedDatabasesRequest.
type ListManagedDatabasesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The identifier of the resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return Managed Databases with the specified management option.
	ManagementOption ListManagedDatabasesManagementOptionEnum `mandatory:"false" contributesTo:"query" name:"managementOption" omitEmpty:"true"`

	// A filter to return Managed Databases of the specified deployment type.
	DeploymentType ListManagedDatabasesDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"deploymentType" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListManagedDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListManagedDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagedDatabasesResponse wrapper for the ListManagedDatabases operation
type ListManagedDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedDatabaseCollection instances
	ManagedDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedDatabasesManagementOptionEnum Enum with underlying type: string
type ListManagedDatabasesManagementOptionEnum string

// Set of constants representing the allowable values for ListManagedDatabasesManagementOptionEnum
const (
	ListManagedDatabasesManagementOptionBasic    ListManagedDatabasesManagementOptionEnum = "BASIC"
	ListManagedDatabasesManagementOptionAdvanced ListManagedDatabasesManagementOptionEnum = "ADVANCED"
)

var mappingListManagedDatabasesManagementOption = map[string]ListManagedDatabasesManagementOptionEnum{
	"BASIC":    ListManagedDatabasesManagementOptionBasic,
	"ADVANCED": ListManagedDatabasesManagementOptionAdvanced,
}

// GetListManagedDatabasesManagementOptionEnumValues Enumerates the set of values for ListManagedDatabasesManagementOptionEnum
func GetListManagedDatabasesManagementOptionEnumValues() []ListManagedDatabasesManagementOptionEnum {
	values := make([]ListManagedDatabasesManagementOptionEnum, 0)
	for _, v := range mappingListManagedDatabasesManagementOption {
		values = append(values, v)
	}
	return values
}

// ListManagedDatabasesDeploymentTypeEnum Enum with underlying type: string
type ListManagedDatabasesDeploymentTypeEnum string

// Set of constants representing the allowable values for ListManagedDatabasesDeploymentTypeEnum
const (
	ListManagedDatabasesDeploymentTypeOnpremise  ListManagedDatabasesDeploymentTypeEnum = "ONPREMISE"
	ListManagedDatabasesDeploymentTypeBm         ListManagedDatabasesDeploymentTypeEnum = "BM"
	ListManagedDatabasesDeploymentTypeVm         ListManagedDatabasesDeploymentTypeEnum = "VM"
	ListManagedDatabasesDeploymentTypeExadata    ListManagedDatabasesDeploymentTypeEnum = "EXADATA"
	ListManagedDatabasesDeploymentTypeExadataCc  ListManagedDatabasesDeploymentTypeEnum = "EXADATA_CC"
	ListManagedDatabasesDeploymentTypeAutonomous ListManagedDatabasesDeploymentTypeEnum = "AUTONOMOUS"
)

var mappingListManagedDatabasesDeploymentType = map[string]ListManagedDatabasesDeploymentTypeEnum{
	"ONPREMISE":  ListManagedDatabasesDeploymentTypeOnpremise,
	"BM":         ListManagedDatabasesDeploymentTypeBm,
	"VM":         ListManagedDatabasesDeploymentTypeVm,
	"EXADATA":    ListManagedDatabasesDeploymentTypeExadata,
	"EXADATA_CC": ListManagedDatabasesDeploymentTypeExadataCc,
	"AUTONOMOUS": ListManagedDatabasesDeploymentTypeAutonomous,
}

// GetListManagedDatabasesDeploymentTypeEnumValues Enumerates the set of values for ListManagedDatabasesDeploymentTypeEnum
func GetListManagedDatabasesDeploymentTypeEnumValues() []ListManagedDatabasesDeploymentTypeEnum {
	values := make([]ListManagedDatabasesDeploymentTypeEnum, 0)
	for _, v := range mappingListManagedDatabasesDeploymentType {
		values = append(values, v)
	}
	return values
}

// ListManagedDatabasesSortByEnum Enum with underlying type: string
type ListManagedDatabasesSortByEnum string

// Set of constants representing the allowable values for ListManagedDatabasesSortByEnum
const (
	ListManagedDatabasesSortByTimecreated ListManagedDatabasesSortByEnum = "TIMECREATED"
	ListManagedDatabasesSortByName        ListManagedDatabasesSortByEnum = "NAME"
)

var mappingListManagedDatabasesSortBy = map[string]ListManagedDatabasesSortByEnum{
	"TIMECREATED": ListManagedDatabasesSortByTimecreated,
	"NAME":        ListManagedDatabasesSortByName,
}

// GetListManagedDatabasesSortByEnumValues Enumerates the set of values for ListManagedDatabasesSortByEnum
func GetListManagedDatabasesSortByEnumValues() []ListManagedDatabasesSortByEnum {
	values := make([]ListManagedDatabasesSortByEnum, 0)
	for _, v := range mappingListManagedDatabasesSortBy {
		values = append(values, v)
	}
	return values
}

// ListManagedDatabasesSortOrderEnum Enum with underlying type: string
type ListManagedDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedDatabasesSortOrderEnum
const (
	ListManagedDatabasesSortOrderAsc  ListManagedDatabasesSortOrderEnum = "ASC"
	ListManagedDatabasesSortOrderDesc ListManagedDatabasesSortOrderEnum = "DESC"
)

var mappingListManagedDatabasesSortOrder = map[string]ListManagedDatabasesSortOrderEnum{
	"ASC":  ListManagedDatabasesSortOrderAsc,
	"DESC": ListManagedDatabasesSortOrderDesc,
}

// GetListManagedDatabasesSortOrderEnumValues Enumerates the set of values for ListManagedDatabasesSortOrderEnum
func GetListManagedDatabasesSortOrderEnumValues() []ListManagedDatabasesSortOrderEnum {
	values := make([]ListManagedDatabasesSortOrderEnum, 0)
	for _, v := range mappingListManagedDatabasesSortOrder {
		values = append(values, v)
	}
	return values
}
