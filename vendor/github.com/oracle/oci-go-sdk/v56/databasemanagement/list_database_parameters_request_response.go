// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDatabaseParametersRequest wrapper for the ListDatabaseParameters operation
//
// See also
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

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

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

var mappingListDatabaseParametersSource = map[string]ListDatabaseParametersSourceEnum{
	"CURRENT": ListDatabaseParametersSourceCurrent,
	"SPFILE":  ListDatabaseParametersSourceSpfile,
}

// GetListDatabaseParametersSourceEnumValues Enumerates the set of values for ListDatabaseParametersSourceEnum
func GetListDatabaseParametersSourceEnumValues() []ListDatabaseParametersSourceEnum {
	values := make([]ListDatabaseParametersSourceEnum, 0)
	for _, v := range mappingListDatabaseParametersSource {
		values = append(values, v)
	}
	return values
}

// ListDatabaseParametersSortByEnum Enum with underlying type: string
type ListDatabaseParametersSortByEnum string

// Set of constants representing the allowable values for ListDatabaseParametersSortByEnum
const (
	ListDatabaseParametersSortByName ListDatabaseParametersSortByEnum = "NAME"
)

var mappingListDatabaseParametersSortBy = map[string]ListDatabaseParametersSortByEnum{
	"NAME": ListDatabaseParametersSortByName,
}

// GetListDatabaseParametersSortByEnumValues Enumerates the set of values for ListDatabaseParametersSortByEnum
func GetListDatabaseParametersSortByEnumValues() []ListDatabaseParametersSortByEnum {
	values := make([]ListDatabaseParametersSortByEnum, 0)
	for _, v := range mappingListDatabaseParametersSortBy {
		values = append(values, v)
	}
	return values
}

// ListDatabaseParametersSortOrderEnum Enum with underlying type: string
type ListDatabaseParametersSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseParametersSortOrderEnum
const (
	ListDatabaseParametersSortOrderAsc  ListDatabaseParametersSortOrderEnum = "ASC"
	ListDatabaseParametersSortOrderDesc ListDatabaseParametersSortOrderEnum = "DESC"
)

var mappingListDatabaseParametersSortOrder = map[string]ListDatabaseParametersSortOrderEnum{
	"ASC":  ListDatabaseParametersSortOrderAsc,
	"DESC": ListDatabaseParametersSortOrderDesc,
}

// GetListDatabaseParametersSortOrderEnumValues Enumerates the set of values for ListDatabaseParametersSortOrderEnum
func GetListDatabaseParametersSortOrderEnumValues() []ListDatabaseParametersSortOrderEnum {
	values := make([]ListDatabaseParametersSortOrderEnum, 0)
	for _, v := range mappingListDatabaseParametersSortOrder {
		values = append(values, v)
	}
	return values
}
