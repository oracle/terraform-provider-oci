// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMigrationParametersRequest wrapper for the ListMigrationParameters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListMigrationParameters.go.html to see an example of how to use ListMigrationParametersRequest.
type ListMigrationParametersRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match a certain Migration Type.
	MigrationType ListMigrationParametersMigrationTypeEnum `mandatory:"false" contributesTo:"query" name:"migrationType" omitEmpty:"true"`

	// A filter to return only resources that match a certain Database Combination.
	DatabaseCombination ListMigrationParametersDatabaseCombinationEnum `mandatory:"false" contributesTo:"query" name:"databaseCombination" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListMigrationParametersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMigrationParametersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMigrationParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMigrationParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMigrationParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMigrationParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMigrationParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMigrationParametersMigrationTypeEnum(string(request.MigrationType)); !ok && request.MigrationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MigrationType: %s. Supported values are: %s.", request.MigrationType, strings.Join(GetListMigrationParametersMigrationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationParametersDatabaseCombinationEnum(string(request.DatabaseCombination)); !ok && request.DatabaseCombination != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseCombination: %s. Supported values are: %s.", request.DatabaseCombination, strings.Join(GetListMigrationParametersDatabaseCombinationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationParametersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMigrationParametersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationParametersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMigrationParametersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMigrationParametersResponse wrapper for the ListMigrationParameters operation
type ListMigrationParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MigrationParameterSummaryCollection instances
	MigrationParameterSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMigrationParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMigrationParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMigrationParametersMigrationTypeEnum Enum with underlying type: string
type ListMigrationParametersMigrationTypeEnum string

// Set of constants representing the allowable values for ListMigrationParametersMigrationTypeEnum
const (
	ListMigrationParametersMigrationTypeOnline  ListMigrationParametersMigrationTypeEnum = "ONLINE"
	ListMigrationParametersMigrationTypeOffline ListMigrationParametersMigrationTypeEnum = "OFFLINE"
)

var mappingListMigrationParametersMigrationTypeEnum = map[string]ListMigrationParametersMigrationTypeEnum{
	"ONLINE":  ListMigrationParametersMigrationTypeOnline,
	"OFFLINE": ListMigrationParametersMigrationTypeOffline,
}

var mappingListMigrationParametersMigrationTypeEnumLowerCase = map[string]ListMigrationParametersMigrationTypeEnum{
	"online":  ListMigrationParametersMigrationTypeOnline,
	"offline": ListMigrationParametersMigrationTypeOffline,
}

// GetListMigrationParametersMigrationTypeEnumValues Enumerates the set of values for ListMigrationParametersMigrationTypeEnum
func GetListMigrationParametersMigrationTypeEnumValues() []ListMigrationParametersMigrationTypeEnum {
	values := make([]ListMigrationParametersMigrationTypeEnum, 0)
	for _, v := range mappingListMigrationParametersMigrationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationParametersMigrationTypeEnumStringValues Enumerates the set of values in String for ListMigrationParametersMigrationTypeEnum
func GetListMigrationParametersMigrationTypeEnumStringValues() []string {
	return []string{
		"ONLINE",
		"OFFLINE",
	}
}

// GetMappingListMigrationParametersMigrationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationParametersMigrationTypeEnum(val string) (ListMigrationParametersMigrationTypeEnum, bool) {
	enum, ok := mappingListMigrationParametersMigrationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationParametersDatabaseCombinationEnum Enum with underlying type: string
type ListMigrationParametersDatabaseCombinationEnum string

// Set of constants representing the allowable values for ListMigrationParametersDatabaseCombinationEnum
const (
	ListMigrationParametersDatabaseCombinationMysql  ListMigrationParametersDatabaseCombinationEnum = "MYSQL"
	ListMigrationParametersDatabaseCombinationOracle ListMigrationParametersDatabaseCombinationEnum = "ORACLE"
)

var mappingListMigrationParametersDatabaseCombinationEnum = map[string]ListMigrationParametersDatabaseCombinationEnum{
	"MYSQL":  ListMigrationParametersDatabaseCombinationMysql,
	"ORACLE": ListMigrationParametersDatabaseCombinationOracle,
}

var mappingListMigrationParametersDatabaseCombinationEnumLowerCase = map[string]ListMigrationParametersDatabaseCombinationEnum{
	"mysql":  ListMigrationParametersDatabaseCombinationMysql,
	"oracle": ListMigrationParametersDatabaseCombinationOracle,
}

// GetListMigrationParametersDatabaseCombinationEnumValues Enumerates the set of values for ListMigrationParametersDatabaseCombinationEnum
func GetListMigrationParametersDatabaseCombinationEnumValues() []ListMigrationParametersDatabaseCombinationEnum {
	values := make([]ListMigrationParametersDatabaseCombinationEnum, 0)
	for _, v := range mappingListMigrationParametersDatabaseCombinationEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationParametersDatabaseCombinationEnumStringValues Enumerates the set of values in String for ListMigrationParametersDatabaseCombinationEnum
func GetListMigrationParametersDatabaseCombinationEnumStringValues() []string {
	return []string{
		"MYSQL",
		"ORACLE",
	}
}

// GetMappingListMigrationParametersDatabaseCombinationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationParametersDatabaseCombinationEnum(val string) (ListMigrationParametersDatabaseCombinationEnum, bool) {
	enum, ok := mappingListMigrationParametersDatabaseCombinationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationParametersSortByEnum Enum with underlying type: string
type ListMigrationParametersSortByEnum string

// Set of constants representing the allowable values for ListMigrationParametersSortByEnum
const (
	ListMigrationParametersSortByTimecreated ListMigrationParametersSortByEnum = "timeCreated"
	ListMigrationParametersSortByDisplayname ListMigrationParametersSortByEnum = "displayName"
)

var mappingListMigrationParametersSortByEnum = map[string]ListMigrationParametersSortByEnum{
	"timeCreated": ListMigrationParametersSortByTimecreated,
	"displayName": ListMigrationParametersSortByDisplayname,
}

var mappingListMigrationParametersSortByEnumLowerCase = map[string]ListMigrationParametersSortByEnum{
	"timecreated": ListMigrationParametersSortByTimecreated,
	"displayname": ListMigrationParametersSortByDisplayname,
}

// GetListMigrationParametersSortByEnumValues Enumerates the set of values for ListMigrationParametersSortByEnum
func GetListMigrationParametersSortByEnumValues() []ListMigrationParametersSortByEnum {
	values := make([]ListMigrationParametersSortByEnum, 0)
	for _, v := range mappingListMigrationParametersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationParametersSortByEnumStringValues Enumerates the set of values in String for ListMigrationParametersSortByEnum
func GetListMigrationParametersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMigrationParametersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationParametersSortByEnum(val string) (ListMigrationParametersSortByEnum, bool) {
	enum, ok := mappingListMigrationParametersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationParametersSortOrderEnum Enum with underlying type: string
type ListMigrationParametersSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationParametersSortOrderEnum
const (
	ListMigrationParametersSortOrderAsc  ListMigrationParametersSortOrderEnum = "ASC"
	ListMigrationParametersSortOrderDesc ListMigrationParametersSortOrderEnum = "DESC"
)

var mappingListMigrationParametersSortOrderEnum = map[string]ListMigrationParametersSortOrderEnum{
	"ASC":  ListMigrationParametersSortOrderAsc,
	"DESC": ListMigrationParametersSortOrderDesc,
}

var mappingListMigrationParametersSortOrderEnumLowerCase = map[string]ListMigrationParametersSortOrderEnum{
	"asc":  ListMigrationParametersSortOrderAsc,
	"desc": ListMigrationParametersSortOrderDesc,
}

// GetListMigrationParametersSortOrderEnumValues Enumerates the set of values for ListMigrationParametersSortOrderEnum
func GetListMigrationParametersSortOrderEnumValues() []ListMigrationParametersSortOrderEnum {
	values := make([]ListMigrationParametersSortOrderEnum, 0)
	for _, v := range mappingListMigrationParametersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationParametersSortOrderEnumStringValues Enumerates the set of values in String for ListMigrationParametersSortOrderEnum
func GetListMigrationParametersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMigrationParametersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationParametersSortOrderEnum(val string) (ListMigrationParametersSortOrderEnum, bool) {
	enum, ok := mappingListMigrationParametersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
