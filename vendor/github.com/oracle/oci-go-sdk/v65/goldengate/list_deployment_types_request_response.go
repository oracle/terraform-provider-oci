// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDeploymentTypesRequest wrapper for the ListDeploymentTypes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentTypes.go.html to see an example of how to use ListDeploymentTypesRequest.
type ListDeploymentTypesRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The type of deployment, the value determines the exact 'type' of the service executed in the deployment. Default value is DATABASE_ORACLE.
	DeploymentType ListDeploymentTypesDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"deploymentType" omitEmpty:"true"`

	// Allows to query by a specific GoldenGate version.
	OggVersion *string `mandatory:"false" contributesTo:"query" name:"oggVersion"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListDeploymentTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDeploymentTypesDeploymentTypeEnum(string(request.DeploymentType)); !ok && request.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", request.DeploymentType, strings.Join(GetListDeploymentTypesDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentTypesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentTypesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentTypesResponse wrapper for the ListDeploymentTypes operation
type ListDeploymentTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentTypeCollection instances
	DeploymentTypeCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentTypesDeploymentTypeEnum Enum with underlying type: string
type ListDeploymentTypesDeploymentTypeEnum string

// Set of constants representing the allowable values for ListDeploymentTypesDeploymentTypeEnum
const (
	ListDeploymentTypesDeploymentTypeOgg                        ListDeploymentTypesDeploymentTypeEnum = "OGG"
	ListDeploymentTypesDeploymentTypeDatabaseOracle             ListDeploymentTypesDeploymentTypeEnum = "DATABASE_ORACLE"
	ListDeploymentTypesDeploymentTypeBigdata                    ListDeploymentTypesDeploymentTypeEnum = "BIGDATA"
	ListDeploymentTypesDeploymentTypeDatabaseMicrosoftSqlserver ListDeploymentTypesDeploymentTypeEnum = "DATABASE_MICROSOFT_SQLSERVER"
	ListDeploymentTypesDeploymentTypeDatabaseMysql              ListDeploymentTypesDeploymentTypeEnum = "DATABASE_MYSQL"
	ListDeploymentTypesDeploymentTypeDatabasePostgresql         ListDeploymentTypesDeploymentTypeEnum = "DATABASE_POSTGRESQL"
	ListDeploymentTypesDeploymentTypeDatabaseDb2zos             ListDeploymentTypesDeploymentTypeEnum = "DATABASE_DB2ZOS"
	ListDeploymentTypesDeploymentTypeGgsa                       ListDeploymentTypesDeploymentTypeEnum = "GGSA"
	ListDeploymentTypesDeploymentTypeDataTransforms             ListDeploymentTypesDeploymentTypeEnum = "DATA_TRANSFORMS"
)

var mappingListDeploymentTypesDeploymentTypeEnum = map[string]ListDeploymentTypesDeploymentTypeEnum{
	"OGG":                          ListDeploymentTypesDeploymentTypeOgg,
	"DATABASE_ORACLE":              ListDeploymentTypesDeploymentTypeDatabaseOracle,
	"BIGDATA":                      ListDeploymentTypesDeploymentTypeBigdata,
	"DATABASE_MICROSOFT_SQLSERVER": ListDeploymentTypesDeploymentTypeDatabaseMicrosoftSqlserver,
	"DATABASE_MYSQL":               ListDeploymentTypesDeploymentTypeDatabaseMysql,
	"DATABASE_POSTGRESQL":          ListDeploymentTypesDeploymentTypeDatabasePostgresql,
	"DATABASE_DB2ZOS":              ListDeploymentTypesDeploymentTypeDatabaseDb2zos,
	"GGSA":                         ListDeploymentTypesDeploymentTypeGgsa,
	"DATA_TRANSFORMS":              ListDeploymentTypesDeploymentTypeDataTransforms,
}

var mappingListDeploymentTypesDeploymentTypeEnumLowerCase = map[string]ListDeploymentTypesDeploymentTypeEnum{
	"ogg":                          ListDeploymentTypesDeploymentTypeOgg,
	"database_oracle":              ListDeploymentTypesDeploymentTypeDatabaseOracle,
	"bigdata":                      ListDeploymentTypesDeploymentTypeBigdata,
	"database_microsoft_sqlserver": ListDeploymentTypesDeploymentTypeDatabaseMicrosoftSqlserver,
	"database_mysql":               ListDeploymentTypesDeploymentTypeDatabaseMysql,
	"database_postgresql":          ListDeploymentTypesDeploymentTypeDatabasePostgresql,
	"database_db2zos":              ListDeploymentTypesDeploymentTypeDatabaseDb2zos,
	"ggsa":                         ListDeploymentTypesDeploymentTypeGgsa,
	"data_transforms":              ListDeploymentTypesDeploymentTypeDataTransforms,
}

// GetListDeploymentTypesDeploymentTypeEnumValues Enumerates the set of values for ListDeploymentTypesDeploymentTypeEnum
func GetListDeploymentTypesDeploymentTypeEnumValues() []ListDeploymentTypesDeploymentTypeEnum {
	values := make([]ListDeploymentTypesDeploymentTypeEnum, 0)
	for _, v := range mappingListDeploymentTypesDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentTypesDeploymentTypeEnumStringValues Enumerates the set of values in String for ListDeploymentTypesDeploymentTypeEnum
func GetListDeploymentTypesDeploymentTypeEnumStringValues() []string {
	return []string{
		"OGG",
		"DATABASE_ORACLE",
		"BIGDATA",
		"DATABASE_MICROSOFT_SQLSERVER",
		"DATABASE_MYSQL",
		"DATABASE_POSTGRESQL",
		"DATABASE_DB2ZOS",
		"GGSA",
		"DATA_TRANSFORMS",
	}
}

// GetMappingListDeploymentTypesDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentTypesDeploymentTypeEnum(val string) (ListDeploymentTypesDeploymentTypeEnum, bool) {
	enum, ok := mappingListDeploymentTypesDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentTypesSortOrderEnum Enum with underlying type: string
type ListDeploymentTypesSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentTypesSortOrderEnum
const (
	ListDeploymentTypesSortOrderAsc  ListDeploymentTypesSortOrderEnum = "ASC"
	ListDeploymentTypesSortOrderDesc ListDeploymentTypesSortOrderEnum = "DESC"
)

var mappingListDeploymentTypesSortOrderEnum = map[string]ListDeploymentTypesSortOrderEnum{
	"ASC":  ListDeploymentTypesSortOrderAsc,
	"DESC": ListDeploymentTypesSortOrderDesc,
}

var mappingListDeploymentTypesSortOrderEnumLowerCase = map[string]ListDeploymentTypesSortOrderEnum{
	"asc":  ListDeploymentTypesSortOrderAsc,
	"desc": ListDeploymentTypesSortOrderDesc,
}

// GetListDeploymentTypesSortOrderEnumValues Enumerates the set of values for ListDeploymentTypesSortOrderEnum
func GetListDeploymentTypesSortOrderEnumValues() []ListDeploymentTypesSortOrderEnum {
	values := make([]ListDeploymentTypesSortOrderEnum, 0)
	for _, v := range mappingListDeploymentTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentTypesSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentTypesSortOrderEnum
func GetListDeploymentTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentTypesSortOrderEnum(val string) (ListDeploymentTypesSortOrderEnum, bool) {
	enum, ok := mappingListDeploymentTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentTypesSortByEnum Enum with underlying type: string
type ListDeploymentTypesSortByEnum string

// Set of constants representing the allowable values for ListDeploymentTypesSortByEnum
const (
	ListDeploymentTypesSortByTimecreated ListDeploymentTypesSortByEnum = "timeCreated"
	ListDeploymentTypesSortByDisplayname ListDeploymentTypesSortByEnum = "displayName"
)

var mappingListDeploymentTypesSortByEnum = map[string]ListDeploymentTypesSortByEnum{
	"timeCreated": ListDeploymentTypesSortByTimecreated,
	"displayName": ListDeploymentTypesSortByDisplayname,
}

var mappingListDeploymentTypesSortByEnumLowerCase = map[string]ListDeploymentTypesSortByEnum{
	"timecreated": ListDeploymentTypesSortByTimecreated,
	"displayname": ListDeploymentTypesSortByDisplayname,
}

// GetListDeploymentTypesSortByEnumValues Enumerates the set of values for ListDeploymentTypesSortByEnum
func GetListDeploymentTypesSortByEnumValues() []ListDeploymentTypesSortByEnum {
	values := make([]ListDeploymentTypesSortByEnum, 0)
	for _, v := range mappingListDeploymentTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentTypesSortByEnumStringValues Enumerates the set of values in String for ListDeploymentTypesSortByEnum
func GetListDeploymentTypesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeploymentTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentTypesSortByEnum(val string) (ListDeploymentTypesSortByEnum, bool) {
	enum, ok := mappingListDeploymentTypesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
