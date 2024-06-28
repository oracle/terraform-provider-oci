// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedDatabasesRequest wrapper for the ListManagedDatabases operation
//
// # See also
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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExternalExadataInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"externalExadataInfrastructureId"`

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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedDatabasesManagementOptionEnum(string(request.ManagementOption)); !ok && request.ManagementOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementOption: %s. Supported values are: %s.", request.ManagementOption, strings.Join(GetListManagedDatabasesManagementOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedDatabasesDeploymentTypeEnum(string(request.DeploymentType)); !ok && request.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", request.DeploymentType, strings.Join(GetListManagedDatabasesDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedDatabasesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListManagedDatabasesManagementOptionEnum = map[string]ListManagedDatabasesManagementOptionEnum{
	"BASIC":    ListManagedDatabasesManagementOptionBasic,
	"ADVANCED": ListManagedDatabasesManagementOptionAdvanced,
}

var mappingListManagedDatabasesManagementOptionEnumLowerCase = map[string]ListManagedDatabasesManagementOptionEnum{
	"basic":    ListManagedDatabasesManagementOptionBasic,
	"advanced": ListManagedDatabasesManagementOptionAdvanced,
}

// GetListManagedDatabasesManagementOptionEnumValues Enumerates the set of values for ListManagedDatabasesManagementOptionEnum
func GetListManagedDatabasesManagementOptionEnumValues() []ListManagedDatabasesManagementOptionEnum {
	values := make([]ListManagedDatabasesManagementOptionEnum, 0)
	for _, v := range mappingListManagedDatabasesManagementOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedDatabasesManagementOptionEnumStringValues Enumerates the set of values in String for ListManagedDatabasesManagementOptionEnum
func GetListManagedDatabasesManagementOptionEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
	}
}

// GetMappingListManagedDatabasesManagementOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedDatabasesManagementOptionEnum(val string) (ListManagedDatabasesManagementOptionEnum, bool) {
	enum, ok := mappingListManagedDatabasesManagementOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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
	ListManagedDatabasesDeploymentTypeExadataXs  ListManagedDatabasesDeploymentTypeEnum = "EXADATA_XS"
)

var mappingListManagedDatabasesDeploymentTypeEnum = map[string]ListManagedDatabasesDeploymentTypeEnum{
	"ONPREMISE":  ListManagedDatabasesDeploymentTypeOnpremise,
	"BM":         ListManagedDatabasesDeploymentTypeBm,
	"VM":         ListManagedDatabasesDeploymentTypeVm,
	"EXADATA":    ListManagedDatabasesDeploymentTypeExadata,
	"EXADATA_CC": ListManagedDatabasesDeploymentTypeExadataCc,
	"AUTONOMOUS": ListManagedDatabasesDeploymentTypeAutonomous,
	"EXADATA_XS": ListManagedDatabasesDeploymentTypeExadataXs,
}

var mappingListManagedDatabasesDeploymentTypeEnumLowerCase = map[string]ListManagedDatabasesDeploymentTypeEnum{
	"onpremise":  ListManagedDatabasesDeploymentTypeOnpremise,
	"bm":         ListManagedDatabasesDeploymentTypeBm,
	"vm":         ListManagedDatabasesDeploymentTypeVm,
	"exadata":    ListManagedDatabasesDeploymentTypeExadata,
	"exadata_cc": ListManagedDatabasesDeploymentTypeExadataCc,
	"autonomous": ListManagedDatabasesDeploymentTypeAutonomous,
	"exadata_xs": ListManagedDatabasesDeploymentTypeExadataXs,
}

// GetListManagedDatabasesDeploymentTypeEnumValues Enumerates the set of values for ListManagedDatabasesDeploymentTypeEnum
func GetListManagedDatabasesDeploymentTypeEnumValues() []ListManagedDatabasesDeploymentTypeEnum {
	values := make([]ListManagedDatabasesDeploymentTypeEnum, 0)
	for _, v := range mappingListManagedDatabasesDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedDatabasesDeploymentTypeEnumStringValues Enumerates the set of values in String for ListManagedDatabasesDeploymentTypeEnum
func GetListManagedDatabasesDeploymentTypeEnumStringValues() []string {
	return []string{
		"ONPREMISE",
		"BM",
		"VM",
		"EXADATA",
		"EXADATA_CC",
		"AUTONOMOUS",
		"EXADATA_XS",
	}
}

// GetMappingListManagedDatabasesDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedDatabasesDeploymentTypeEnum(val string) (ListManagedDatabasesDeploymentTypeEnum, bool) {
	enum, ok := mappingListManagedDatabasesDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedDatabasesSortByEnum Enum with underlying type: string
type ListManagedDatabasesSortByEnum string

// Set of constants representing the allowable values for ListManagedDatabasesSortByEnum
const (
	ListManagedDatabasesSortByTimecreated ListManagedDatabasesSortByEnum = "TIMECREATED"
	ListManagedDatabasesSortByName        ListManagedDatabasesSortByEnum = "NAME"
)

var mappingListManagedDatabasesSortByEnum = map[string]ListManagedDatabasesSortByEnum{
	"TIMECREATED": ListManagedDatabasesSortByTimecreated,
	"NAME":        ListManagedDatabasesSortByName,
}

var mappingListManagedDatabasesSortByEnumLowerCase = map[string]ListManagedDatabasesSortByEnum{
	"timecreated": ListManagedDatabasesSortByTimecreated,
	"name":        ListManagedDatabasesSortByName,
}

// GetListManagedDatabasesSortByEnumValues Enumerates the set of values for ListManagedDatabasesSortByEnum
func GetListManagedDatabasesSortByEnumValues() []ListManagedDatabasesSortByEnum {
	values := make([]ListManagedDatabasesSortByEnum, 0)
	for _, v := range mappingListManagedDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedDatabasesSortByEnumStringValues Enumerates the set of values in String for ListManagedDatabasesSortByEnum
func GetListManagedDatabasesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListManagedDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedDatabasesSortByEnum(val string) (ListManagedDatabasesSortByEnum, bool) {
	enum, ok := mappingListManagedDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedDatabasesSortOrderEnum Enum with underlying type: string
type ListManagedDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedDatabasesSortOrderEnum
const (
	ListManagedDatabasesSortOrderAsc  ListManagedDatabasesSortOrderEnum = "ASC"
	ListManagedDatabasesSortOrderDesc ListManagedDatabasesSortOrderEnum = "DESC"
)

var mappingListManagedDatabasesSortOrderEnum = map[string]ListManagedDatabasesSortOrderEnum{
	"ASC":  ListManagedDatabasesSortOrderAsc,
	"DESC": ListManagedDatabasesSortOrderDesc,
}

var mappingListManagedDatabasesSortOrderEnumLowerCase = map[string]ListManagedDatabasesSortOrderEnum{
	"asc":  ListManagedDatabasesSortOrderAsc,
	"desc": ListManagedDatabasesSortOrderDesc,
}

// GetListManagedDatabasesSortOrderEnumValues Enumerates the set of values for ListManagedDatabasesSortOrderEnum
func GetListManagedDatabasesSortOrderEnumValues() []ListManagedDatabasesSortOrderEnum {
	values := make([]ListManagedDatabasesSortOrderEnum, 0)
	for _, v := range mappingListManagedDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedDatabasesSortOrderEnum
func GetListManagedDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedDatabasesSortOrderEnum(val string) (ListManagedDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListManagedDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
