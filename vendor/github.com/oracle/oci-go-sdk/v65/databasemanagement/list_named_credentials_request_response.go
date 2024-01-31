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

// ListNamedCredentialsRequest wrapper for the ListNamedCredentials operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListNamedCredentials.go.html to see an example of how to use ListNamedCredentialsRequest.
type ListNamedCredentialsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The resource associated to the named credential.
	AssociatedResource *string `mandatory:"false" contributesTo:"query" name:"associatedResource"`

	// The type of database that is associated to the named credential.
	Type ListNamedCredentialsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The scope of named credential.
	Scope ListNamedCredentialsScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The name of the named credential.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListNamedCredentialsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListNamedCredentialsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNamedCredentialsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNamedCredentialsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNamedCredentialsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNamedCredentialsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNamedCredentialsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNamedCredentialsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListNamedCredentialsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNamedCredentialsScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetListNamedCredentialsScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNamedCredentialsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNamedCredentialsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNamedCredentialsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNamedCredentialsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNamedCredentialsResponse wrapper for the ListNamedCredentials operation
type ListNamedCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NamedCredentialCollection instances
	NamedCredentialCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNamedCredentialsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNamedCredentialsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNamedCredentialsTypeEnum Enum with underlying type: string
type ListNamedCredentialsTypeEnum string

// Set of constants representing the allowable values for ListNamedCredentialsTypeEnum
const (
	ListNamedCredentialsTypeOracleDb ListNamedCredentialsTypeEnum = "ORACLE_DB"
)

var mappingListNamedCredentialsTypeEnum = map[string]ListNamedCredentialsTypeEnum{
	"ORACLE_DB": ListNamedCredentialsTypeOracleDb,
}

var mappingListNamedCredentialsTypeEnumLowerCase = map[string]ListNamedCredentialsTypeEnum{
	"oracle_db": ListNamedCredentialsTypeOracleDb,
}

// GetListNamedCredentialsTypeEnumValues Enumerates the set of values for ListNamedCredentialsTypeEnum
func GetListNamedCredentialsTypeEnumValues() []ListNamedCredentialsTypeEnum {
	values := make([]ListNamedCredentialsTypeEnum, 0)
	for _, v := range mappingListNamedCredentialsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListNamedCredentialsTypeEnumStringValues Enumerates the set of values in String for ListNamedCredentialsTypeEnum
func GetListNamedCredentialsTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DB",
	}
}

// GetMappingListNamedCredentialsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNamedCredentialsTypeEnum(val string) (ListNamedCredentialsTypeEnum, bool) {
	enum, ok := mappingListNamedCredentialsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNamedCredentialsScopeEnum Enum with underlying type: string
type ListNamedCredentialsScopeEnum string

// Set of constants representing the allowable values for ListNamedCredentialsScopeEnum
const (
	ListNamedCredentialsScopeResource ListNamedCredentialsScopeEnum = "RESOURCE"
	ListNamedCredentialsScopeGlobal   ListNamedCredentialsScopeEnum = "GLOBAL"
)

var mappingListNamedCredentialsScopeEnum = map[string]ListNamedCredentialsScopeEnum{
	"RESOURCE": ListNamedCredentialsScopeResource,
	"GLOBAL":   ListNamedCredentialsScopeGlobal,
}

var mappingListNamedCredentialsScopeEnumLowerCase = map[string]ListNamedCredentialsScopeEnum{
	"resource": ListNamedCredentialsScopeResource,
	"global":   ListNamedCredentialsScopeGlobal,
}

// GetListNamedCredentialsScopeEnumValues Enumerates the set of values for ListNamedCredentialsScopeEnum
func GetListNamedCredentialsScopeEnumValues() []ListNamedCredentialsScopeEnum {
	values := make([]ListNamedCredentialsScopeEnum, 0)
	for _, v := range mappingListNamedCredentialsScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetListNamedCredentialsScopeEnumStringValues Enumerates the set of values in String for ListNamedCredentialsScopeEnum
func GetListNamedCredentialsScopeEnumStringValues() []string {
	return []string{
		"RESOURCE",
		"GLOBAL",
	}
}

// GetMappingListNamedCredentialsScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNamedCredentialsScopeEnum(val string) (ListNamedCredentialsScopeEnum, bool) {
	enum, ok := mappingListNamedCredentialsScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNamedCredentialsSortByEnum Enum with underlying type: string
type ListNamedCredentialsSortByEnum string

// Set of constants representing the allowable values for ListNamedCredentialsSortByEnum
const (
	ListNamedCredentialsSortByTimecreated ListNamedCredentialsSortByEnum = "TIMECREATED"
	ListNamedCredentialsSortByName        ListNamedCredentialsSortByEnum = "NAME"
)

var mappingListNamedCredentialsSortByEnum = map[string]ListNamedCredentialsSortByEnum{
	"TIMECREATED": ListNamedCredentialsSortByTimecreated,
	"NAME":        ListNamedCredentialsSortByName,
}

var mappingListNamedCredentialsSortByEnumLowerCase = map[string]ListNamedCredentialsSortByEnum{
	"timecreated": ListNamedCredentialsSortByTimecreated,
	"name":        ListNamedCredentialsSortByName,
}

// GetListNamedCredentialsSortByEnumValues Enumerates the set of values for ListNamedCredentialsSortByEnum
func GetListNamedCredentialsSortByEnumValues() []ListNamedCredentialsSortByEnum {
	values := make([]ListNamedCredentialsSortByEnum, 0)
	for _, v := range mappingListNamedCredentialsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNamedCredentialsSortByEnumStringValues Enumerates the set of values in String for ListNamedCredentialsSortByEnum
func GetListNamedCredentialsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListNamedCredentialsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNamedCredentialsSortByEnum(val string) (ListNamedCredentialsSortByEnum, bool) {
	enum, ok := mappingListNamedCredentialsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNamedCredentialsSortOrderEnum Enum with underlying type: string
type ListNamedCredentialsSortOrderEnum string

// Set of constants representing the allowable values for ListNamedCredentialsSortOrderEnum
const (
	ListNamedCredentialsSortOrderAsc  ListNamedCredentialsSortOrderEnum = "ASC"
	ListNamedCredentialsSortOrderDesc ListNamedCredentialsSortOrderEnum = "DESC"
)

var mappingListNamedCredentialsSortOrderEnum = map[string]ListNamedCredentialsSortOrderEnum{
	"ASC":  ListNamedCredentialsSortOrderAsc,
	"DESC": ListNamedCredentialsSortOrderDesc,
}

var mappingListNamedCredentialsSortOrderEnumLowerCase = map[string]ListNamedCredentialsSortOrderEnum{
	"asc":  ListNamedCredentialsSortOrderAsc,
	"desc": ListNamedCredentialsSortOrderDesc,
}

// GetListNamedCredentialsSortOrderEnumValues Enumerates the set of values for ListNamedCredentialsSortOrderEnum
func GetListNamedCredentialsSortOrderEnumValues() []ListNamedCredentialsSortOrderEnum {
	values := make([]ListNamedCredentialsSortOrderEnum, 0)
	for _, v := range mappingListNamedCredentialsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNamedCredentialsSortOrderEnumStringValues Enumerates the set of values in String for ListNamedCredentialsSortOrderEnum
func GetListNamedCredentialsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNamedCredentialsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNamedCredentialsSortOrderEnum(val string) (ListNamedCredentialsSortOrderEnum, bool) {
	enum, ok := mappingListNamedCredentialsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
