// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConnectionsRequest wrapper for the ListConnections operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListConnections.go.html to see an example of how to use ListConnectionsRequest.
type ListConnectionsRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The array of technology types.
	TechnologyType []TechnologyTypeEnum `contributesTo:"query" name:"technologyType" omitEmpty:"true" collectionFormat:"multi"`

	// The array of connection types.
	ConnectionType []ConnectionTypeEnum `contributesTo:"query" name:"connectionType" omitEmpty:"true" collectionFormat:"multi"`

	// The OCID of the deployment which for the connection must be assigned.
	AssignedDeploymentId *string `mandatory:"false" contributesTo:"query" name:"assignedDeploymentId"`

	// Filters for compatible connections which can be, but currently not assigned to the deployment specified by its id.
	AssignableDeploymentId *string `mandatory:"false" contributesTo:"query" name:"assignableDeploymentId"`

	// Filters for connections which can be assigned to the latest version of the specified deployment type.
	AssignableDeploymentType ListConnectionsAssignableDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"assignableDeploymentType" omitEmpty:"true"`

	// A filter to return only connections having the 'lifecycleState' given.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConnectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListConnectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConnectionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConnectionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.TechnologyType {
		if _, ok := GetMappingTechnologyTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", val, strings.Join(GetTechnologyTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ConnectionType {
		if _, ok := GetMappingConnectionTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionType: %s. Supported values are: %s.", val, strings.Join(GetConnectionTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListConnectionsAssignableDeploymentTypeEnum(string(request.AssignableDeploymentType)); !ok && request.AssignableDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssignableDeploymentType: %s. Supported values are: %s.", request.AssignableDeploymentType, strings.Join(GetListConnectionsAssignableDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConnectionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConnectionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConnectionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConnectionsResponse wrapper for the ListConnections operation
type ListConnectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectionCollection instances
	ConnectionCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConnectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectionsAssignableDeploymentTypeEnum Enum with underlying type: string
type ListConnectionsAssignableDeploymentTypeEnum string

// Set of constants representing the allowable values for ListConnectionsAssignableDeploymentTypeEnum
const (
	ListConnectionsAssignableDeploymentTypeOgg                        ListConnectionsAssignableDeploymentTypeEnum = "OGG"
	ListConnectionsAssignableDeploymentTypeDatabaseOracle             ListConnectionsAssignableDeploymentTypeEnum = "DATABASE_ORACLE"
	ListConnectionsAssignableDeploymentTypeBigdata                    ListConnectionsAssignableDeploymentTypeEnum = "BIGDATA"
	ListConnectionsAssignableDeploymentTypeDatabaseMicrosoftSqlserver ListConnectionsAssignableDeploymentTypeEnum = "DATABASE_MICROSOFT_SQLSERVER"
	ListConnectionsAssignableDeploymentTypeDatabaseMysql              ListConnectionsAssignableDeploymentTypeEnum = "DATABASE_MYSQL"
	ListConnectionsAssignableDeploymentTypeDatabasePostgresql         ListConnectionsAssignableDeploymentTypeEnum = "DATABASE_POSTGRESQL"
	ListConnectionsAssignableDeploymentTypeDatabaseDb2zos             ListConnectionsAssignableDeploymentTypeEnum = "DATABASE_DB2ZOS"
	ListConnectionsAssignableDeploymentTypeDataTransforms             ListConnectionsAssignableDeploymentTypeEnum = "DATA_TRANSFORMS"
)

var mappingListConnectionsAssignableDeploymentTypeEnum = map[string]ListConnectionsAssignableDeploymentTypeEnum{
	"OGG":                          ListConnectionsAssignableDeploymentTypeOgg,
	"DATABASE_ORACLE":              ListConnectionsAssignableDeploymentTypeDatabaseOracle,
	"BIGDATA":                      ListConnectionsAssignableDeploymentTypeBigdata,
	"DATABASE_MICROSOFT_SQLSERVER": ListConnectionsAssignableDeploymentTypeDatabaseMicrosoftSqlserver,
	"DATABASE_MYSQL":               ListConnectionsAssignableDeploymentTypeDatabaseMysql,
	"DATABASE_POSTGRESQL":          ListConnectionsAssignableDeploymentTypeDatabasePostgresql,
	"DATABASE_DB2ZOS":              ListConnectionsAssignableDeploymentTypeDatabaseDb2zos,
	"DATA_TRANSFORMS":              ListConnectionsAssignableDeploymentTypeDataTransforms,
}

var mappingListConnectionsAssignableDeploymentTypeEnumLowerCase = map[string]ListConnectionsAssignableDeploymentTypeEnum{
	"ogg":                          ListConnectionsAssignableDeploymentTypeOgg,
	"database_oracle":              ListConnectionsAssignableDeploymentTypeDatabaseOracle,
	"bigdata":                      ListConnectionsAssignableDeploymentTypeBigdata,
	"database_microsoft_sqlserver": ListConnectionsAssignableDeploymentTypeDatabaseMicrosoftSqlserver,
	"database_mysql":               ListConnectionsAssignableDeploymentTypeDatabaseMysql,
	"database_postgresql":          ListConnectionsAssignableDeploymentTypeDatabasePostgresql,
	"database_db2zos":              ListConnectionsAssignableDeploymentTypeDatabaseDb2zos,
	"data_transforms":              ListConnectionsAssignableDeploymentTypeDataTransforms,
}

// GetListConnectionsAssignableDeploymentTypeEnumValues Enumerates the set of values for ListConnectionsAssignableDeploymentTypeEnum
func GetListConnectionsAssignableDeploymentTypeEnumValues() []ListConnectionsAssignableDeploymentTypeEnum {
	values := make([]ListConnectionsAssignableDeploymentTypeEnum, 0)
	for _, v := range mappingListConnectionsAssignableDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsAssignableDeploymentTypeEnumStringValues Enumerates the set of values in String for ListConnectionsAssignableDeploymentTypeEnum
func GetListConnectionsAssignableDeploymentTypeEnumStringValues() []string {
	return []string{
		"OGG",
		"DATABASE_ORACLE",
		"BIGDATA",
		"DATABASE_MICROSOFT_SQLSERVER",
		"DATABASE_MYSQL",
		"DATABASE_POSTGRESQL",
		"DATABASE_DB2ZOS",
		"DATA_TRANSFORMS",
	}
}

// GetMappingListConnectionsAssignableDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsAssignableDeploymentTypeEnum(val string) (ListConnectionsAssignableDeploymentTypeEnum, bool) {
	enum, ok := mappingListConnectionsAssignableDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsSortOrderEnum Enum with underlying type: string
type ListConnectionsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectionsSortOrderEnum
const (
	ListConnectionsSortOrderAsc  ListConnectionsSortOrderEnum = "ASC"
	ListConnectionsSortOrderDesc ListConnectionsSortOrderEnum = "DESC"
)

var mappingListConnectionsSortOrderEnum = map[string]ListConnectionsSortOrderEnum{
	"ASC":  ListConnectionsSortOrderAsc,
	"DESC": ListConnectionsSortOrderDesc,
}

var mappingListConnectionsSortOrderEnumLowerCase = map[string]ListConnectionsSortOrderEnum{
	"asc":  ListConnectionsSortOrderAsc,
	"desc": ListConnectionsSortOrderDesc,
}

// GetListConnectionsSortOrderEnumValues Enumerates the set of values for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumValues() []ListConnectionsSortOrderEnum {
	values := make([]ListConnectionsSortOrderEnum, 0)
	for _, v := range mappingListConnectionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsSortOrderEnumStringValues Enumerates the set of values in String for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConnectionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsSortOrderEnum(val string) (ListConnectionsSortOrderEnum, bool) {
	enum, ok := mappingListConnectionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsSortByEnum Enum with underlying type: string
type ListConnectionsSortByEnum string

// Set of constants representing the allowable values for ListConnectionsSortByEnum
const (
	ListConnectionsSortByTimecreated ListConnectionsSortByEnum = "timeCreated"
	ListConnectionsSortByDisplayname ListConnectionsSortByEnum = "displayName"
)

var mappingListConnectionsSortByEnum = map[string]ListConnectionsSortByEnum{
	"timeCreated": ListConnectionsSortByTimecreated,
	"displayName": ListConnectionsSortByDisplayname,
}

var mappingListConnectionsSortByEnumLowerCase = map[string]ListConnectionsSortByEnum{
	"timecreated": ListConnectionsSortByTimecreated,
	"displayname": ListConnectionsSortByDisplayname,
}

// GetListConnectionsSortByEnumValues Enumerates the set of values for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumValues() []ListConnectionsSortByEnum {
	values := make([]ListConnectionsSortByEnum, 0)
	for _, v := range mappingListConnectionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsSortByEnumStringValues Enumerates the set of values in String for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListConnectionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsSortByEnum(val string) (ListConnectionsSortByEnum, bool) {
	enum, ok := mappingListConnectionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
