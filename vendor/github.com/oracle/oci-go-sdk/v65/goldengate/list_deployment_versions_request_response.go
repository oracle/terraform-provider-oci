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

// ListDeploymentVersionsRequest wrapper for the ListDeploymentVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentVersions.go.html to see an example of how to use ListDeploymentVersionsRequest.
type ListDeploymentVersionsRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment in which to list resources.
	DeploymentId *string `mandatory:"false" contributesTo:"query" name:"deploymentId"`

	// The type of deployment, the value determines the exact 'type' of the service executed in the deployment. Default value is DATABASE_ORACLE.
	DeploymentType ListDeploymentVersionsDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"deploymentType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListDeploymentVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDeploymentVersionsDeploymentTypeEnum(string(request.DeploymentType)); !ok && request.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", request.DeploymentType, strings.Join(GetListDeploymentVersionsDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentVersionsResponse wrapper for the ListDeploymentVersions operation
type ListDeploymentVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentVersionCollection instances
	DeploymentVersionCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentVersionsDeploymentTypeEnum Enum with underlying type: string
type ListDeploymentVersionsDeploymentTypeEnum string

// Set of constants representing the allowable values for ListDeploymentVersionsDeploymentTypeEnum
const (
	ListDeploymentVersionsDeploymentTypeOgg                        ListDeploymentVersionsDeploymentTypeEnum = "OGG"
	ListDeploymentVersionsDeploymentTypeDatabaseOracle             ListDeploymentVersionsDeploymentTypeEnum = "DATABASE_ORACLE"
	ListDeploymentVersionsDeploymentTypeBigdata                    ListDeploymentVersionsDeploymentTypeEnum = "BIGDATA"
	ListDeploymentVersionsDeploymentTypeDatabaseMicrosoftSqlserver ListDeploymentVersionsDeploymentTypeEnum = "DATABASE_MICROSOFT_SQLSERVER"
	ListDeploymentVersionsDeploymentTypeDatabaseMysql              ListDeploymentVersionsDeploymentTypeEnum = "DATABASE_MYSQL"
	ListDeploymentVersionsDeploymentTypeDatabasePostgresql         ListDeploymentVersionsDeploymentTypeEnum = "DATABASE_POSTGRESQL"
	ListDeploymentVersionsDeploymentTypeDatabaseDb2zos             ListDeploymentVersionsDeploymentTypeEnum = "DATABASE_DB2ZOS"
	ListDeploymentVersionsDeploymentTypeDataTransforms             ListDeploymentVersionsDeploymentTypeEnum = "DATA_TRANSFORMS"
)

var mappingListDeploymentVersionsDeploymentTypeEnum = map[string]ListDeploymentVersionsDeploymentTypeEnum{
	"OGG":                          ListDeploymentVersionsDeploymentTypeOgg,
	"DATABASE_ORACLE":              ListDeploymentVersionsDeploymentTypeDatabaseOracle,
	"BIGDATA":                      ListDeploymentVersionsDeploymentTypeBigdata,
	"DATABASE_MICROSOFT_SQLSERVER": ListDeploymentVersionsDeploymentTypeDatabaseMicrosoftSqlserver,
	"DATABASE_MYSQL":               ListDeploymentVersionsDeploymentTypeDatabaseMysql,
	"DATABASE_POSTGRESQL":          ListDeploymentVersionsDeploymentTypeDatabasePostgresql,
	"DATABASE_DB2ZOS":              ListDeploymentVersionsDeploymentTypeDatabaseDb2zos,
	"DATA_TRANSFORMS":              ListDeploymentVersionsDeploymentTypeDataTransforms,
}

var mappingListDeploymentVersionsDeploymentTypeEnumLowerCase = map[string]ListDeploymentVersionsDeploymentTypeEnum{
	"ogg":                          ListDeploymentVersionsDeploymentTypeOgg,
	"database_oracle":              ListDeploymentVersionsDeploymentTypeDatabaseOracle,
	"bigdata":                      ListDeploymentVersionsDeploymentTypeBigdata,
	"database_microsoft_sqlserver": ListDeploymentVersionsDeploymentTypeDatabaseMicrosoftSqlserver,
	"database_mysql":               ListDeploymentVersionsDeploymentTypeDatabaseMysql,
	"database_postgresql":          ListDeploymentVersionsDeploymentTypeDatabasePostgresql,
	"database_db2zos":              ListDeploymentVersionsDeploymentTypeDatabaseDb2zos,
	"data_transforms":              ListDeploymentVersionsDeploymentTypeDataTransforms,
}

// GetListDeploymentVersionsDeploymentTypeEnumValues Enumerates the set of values for ListDeploymentVersionsDeploymentTypeEnum
func GetListDeploymentVersionsDeploymentTypeEnumValues() []ListDeploymentVersionsDeploymentTypeEnum {
	values := make([]ListDeploymentVersionsDeploymentTypeEnum, 0)
	for _, v := range mappingListDeploymentVersionsDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentVersionsDeploymentTypeEnumStringValues Enumerates the set of values in String for ListDeploymentVersionsDeploymentTypeEnum
func GetListDeploymentVersionsDeploymentTypeEnumStringValues() []string {
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

// GetMappingListDeploymentVersionsDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentVersionsDeploymentTypeEnum(val string) (ListDeploymentVersionsDeploymentTypeEnum, bool) {
	enum, ok := mappingListDeploymentVersionsDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentVersionsSortOrderEnum Enum with underlying type: string
type ListDeploymentVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentVersionsSortOrderEnum
const (
	ListDeploymentVersionsSortOrderAsc  ListDeploymentVersionsSortOrderEnum = "ASC"
	ListDeploymentVersionsSortOrderDesc ListDeploymentVersionsSortOrderEnum = "DESC"
)

var mappingListDeploymentVersionsSortOrderEnum = map[string]ListDeploymentVersionsSortOrderEnum{
	"ASC":  ListDeploymentVersionsSortOrderAsc,
	"DESC": ListDeploymentVersionsSortOrderDesc,
}

var mappingListDeploymentVersionsSortOrderEnumLowerCase = map[string]ListDeploymentVersionsSortOrderEnum{
	"asc":  ListDeploymentVersionsSortOrderAsc,
	"desc": ListDeploymentVersionsSortOrderDesc,
}

// GetListDeploymentVersionsSortOrderEnumValues Enumerates the set of values for ListDeploymentVersionsSortOrderEnum
func GetListDeploymentVersionsSortOrderEnumValues() []ListDeploymentVersionsSortOrderEnum {
	values := make([]ListDeploymentVersionsSortOrderEnum, 0)
	for _, v := range mappingListDeploymentVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentVersionsSortOrderEnum
func GetListDeploymentVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentVersionsSortOrderEnum(val string) (ListDeploymentVersionsSortOrderEnum, bool) {
	enum, ok := mappingListDeploymentVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentVersionsSortByEnum Enum with underlying type: string
type ListDeploymentVersionsSortByEnum string

// Set of constants representing the allowable values for ListDeploymentVersionsSortByEnum
const (
	ListDeploymentVersionsSortByTimecreated ListDeploymentVersionsSortByEnum = "timeCreated"
	ListDeploymentVersionsSortByDisplayname ListDeploymentVersionsSortByEnum = "displayName"
)

var mappingListDeploymentVersionsSortByEnum = map[string]ListDeploymentVersionsSortByEnum{
	"timeCreated": ListDeploymentVersionsSortByTimecreated,
	"displayName": ListDeploymentVersionsSortByDisplayname,
}

var mappingListDeploymentVersionsSortByEnumLowerCase = map[string]ListDeploymentVersionsSortByEnum{
	"timecreated": ListDeploymentVersionsSortByTimecreated,
	"displayname": ListDeploymentVersionsSortByDisplayname,
}

// GetListDeploymentVersionsSortByEnumValues Enumerates the set of values for ListDeploymentVersionsSortByEnum
func GetListDeploymentVersionsSortByEnumValues() []ListDeploymentVersionsSortByEnum {
	values := make([]ListDeploymentVersionsSortByEnum, 0)
	for _, v := range mappingListDeploymentVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentVersionsSortByEnumStringValues Enumerates the set of values in String for ListDeploymentVersionsSortByEnum
func GetListDeploymentVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeploymentVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentVersionsSortByEnum(val string) (ListDeploymentVersionsSortByEnum, bool) {
	enum, ok := mappingListDeploymentVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
