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

// ListDatabaseParametersRequest wrapper for the ListDatabaseParameters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListDatabaseParameters.go.html to see an example of how to use ListDatabaseParametersRequest.
type ListDatabaseParametersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The source used to list database parameters. `CURRENT` is used to get the
	// database parameters that are currently in effect for the database
	// instance. `SPFILE` is used to list parameters from the server parameter
	// file. Default is `CURRENT`.
	Source ListDatabaseParametersSourceEnum `mandatory:"false" contributesTo:"query" name:"source" omitEmpty:"true"`

	// A filter to return all parameters that have the text given in their names.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// When true, results include a list of valid values for parameters (if applicable).
	IsAllowedValuesIncluded *bool `mandatory:"false" contributesTo:"query" name:"isAllowedValuesIncluded"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `NAME` is ascending and it is case-sensitive.
	SortBy ListDatabaseParametersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListDatabaseParametersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseParametersSourceEnum(string(request.Source)); !ok && request.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", request.Source, strings.Join(GetListDatabaseParametersSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseParametersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseParametersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseParametersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseParametersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseParametersResponse wrapper for the ListDatabaseParameters operation
type ListDatabaseParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DatabaseParametersCollection instance
	DatabaseParametersCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDatabaseParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseParametersSourceEnum Enum with underlying type: string
type ListDatabaseParametersSourceEnum string

// Set of constants representing the allowable values for ListDatabaseParametersSourceEnum
const (
	ListDatabaseParametersSourceCurrent ListDatabaseParametersSourceEnum = "CURRENT"
	ListDatabaseParametersSourceSpfile  ListDatabaseParametersSourceEnum = "SPFILE"
)

var mappingListDatabaseParametersSourceEnum = map[string]ListDatabaseParametersSourceEnum{
	"CURRENT": ListDatabaseParametersSourceCurrent,
	"SPFILE":  ListDatabaseParametersSourceSpfile,
}

var mappingListDatabaseParametersSourceEnumLowerCase = map[string]ListDatabaseParametersSourceEnum{
	"current": ListDatabaseParametersSourceCurrent,
	"spfile":  ListDatabaseParametersSourceSpfile,
}

// GetListDatabaseParametersSourceEnumValues Enumerates the set of values for ListDatabaseParametersSourceEnum
func GetListDatabaseParametersSourceEnumValues() []ListDatabaseParametersSourceEnum {
	values := make([]ListDatabaseParametersSourceEnum, 0)
	for _, v := range mappingListDatabaseParametersSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseParametersSourceEnumStringValues Enumerates the set of values in String for ListDatabaseParametersSourceEnum
func GetListDatabaseParametersSourceEnumStringValues() []string {
	return []string{
		"CURRENT",
		"SPFILE",
	}
}

// GetMappingListDatabaseParametersSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseParametersSourceEnum(val string) (ListDatabaseParametersSourceEnum, bool) {
	enum, ok := mappingListDatabaseParametersSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseParametersSortByEnum Enum with underlying type: string
type ListDatabaseParametersSortByEnum string

// Set of constants representing the allowable values for ListDatabaseParametersSortByEnum
const (
	ListDatabaseParametersSortByName ListDatabaseParametersSortByEnum = "NAME"
)

var mappingListDatabaseParametersSortByEnum = map[string]ListDatabaseParametersSortByEnum{
	"NAME": ListDatabaseParametersSortByName,
}

var mappingListDatabaseParametersSortByEnumLowerCase = map[string]ListDatabaseParametersSortByEnum{
	"name": ListDatabaseParametersSortByName,
}

// GetListDatabaseParametersSortByEnumValues Enumerates the set of values for ListDatabaseParametersSortByEnum
func GetListDatabaseParametersSortByEnumValues() []ListDatabaseParametersSortByEnum {
	values := make([]ListDatabaseParametersSortByEnum, 0)
	for _, v := range mappingListDatabaseParametersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseParametersSortByEnumStringValues Enumerates the set of values in String for ListDatabaseParametersSortByEnum
func GetListDatabaseParametersSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListDatabaseParametersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseParametersSortByEnum(val string) (ListDatabaseParametersSortByEnum, bool) {
	enum, ok := mappingListDatabaseParametersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseParametersSortOrderEnum Enum with underlying type: string
type ListDatabaseParametersSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseParametersSortOrderEnum
const (
	ListDatabaseParametersSortOrderAsc  ListDatabaseParametersSortOrderEnum = "ASC"
	ListDatabaseParametersSortOrderDesc ListDatabaseParametersSortOrderEnum = "DESC"
)

var mappingListDatabaseParametersSortOrderEnum = map[string]ListDatabaseParametersSortOrderEnum{
	"ASC":  ListDatabaseParametersSortOrderAsc,
	"DESC": ListDatabaseParametersSortOrderDesc,
}

var mappingListDatabaseParametersSortOrderEnumLowerCase = map[string]ListDatabaseParametersSortOrderEnum{
	"asc":  ListDatabaseParametersSortOrderAsc,
	"desc": ListDatabaseParametersSortOrderDesc,
}

// GetListDatabaseParametersSortOrderEnumValues Enumerates the set of values for ListDatabaseParametersSortOrderEnum
func GetListDatabaseParametersSortOrderEnumValues() []ListDatabaseParametersSortOrderEnum {
	values := make([]ListDatabaseParametersSortOrderEnum, 0)
	for _, v := range mappingListDatabaseParametersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseParametersSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseParametersSortOrderEnum
func GetListDatabaseParametersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseParametersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseParametersSortOrderEnum(val string) (ListDatabaseParametersSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseParametersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
