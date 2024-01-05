// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousContainerDatabaseVersionsRequest wrapper for the ListAutonomousContainerDatabaseVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousContainerDatabaseVersions.go.html to see an example of how to use ListAutonomousContainerDatabaseVersionsRequest.
type ListAutonomousContainerDatabaseVersionsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The service component to use, either ADBD or EXACC.
	ServiceComponent ListAutonomousContainerDatabaseVersionsServiceComponentEnum `mandatory:"true" contributesTo:"query" name:"serviceComponent" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousContainerDatabaseVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousContainerDatabaseVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousContainerDatabaseVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousContainerDatabaseVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousContainerDatabaseVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousContainerDatabaseVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutonomousContainerDatabaseVersionsServiceComponentEnum(string(request.ServiceComponent)); !ok && request.ServiceComponent != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceComponent: %s. Supported values are: %s.", request.ServiceComponent, strings.Join(GetListAutonomousContainerDatabaseVersionsServiceComponentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousContainerDatabaseVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutonomousContainerDatabaseVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousContainerDatabaseVersionsResponse wrapper for the ListAutonomousContainerDatabaseVersions operation
type ListAutonomousContainerDatabaseVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousContainerDatabaseVersionSummary instances
	Items []AutonomousContainerDatabaseVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousContainerDatabaseVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousContainerDatabaseVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousContainerDatabaseVersionsServiceComponentEnum Enum with underlying type: string
type ListAutonomousContainerDatabaseVersionsServiceComponentEnum string

// Set of constants representing the allowable values for ListAutonomousContainerDatabaseVersionsServiceComponentEnum
const (
	ListAutonomousContainerDatabaseVersionsServiceComponentAdbd  ListAutonomousContainerDatabaseVersionsServiceComponentEnum = "ADBD"
	ListAutonomousContainerDatabaseVersionsServiceComponentExacc ListAutonomousContainerDatabaseVersionsServiceComponentEnum = "EXACC"
)

var mappingListAutonomousContainerDatabaseVersionsServiceComponentEnum = map[string]ListAutonomousContainerDatabaseVersionsServiceComponentEnum{
	"ADBD":  ListAutonomousContainerDatabaseVersionsServiceComponentAdbd,
	"EXACC": ListAutonomousContainerDatabaseVersionsServiceComponentExacc,
}

var mappingListAutonomousContainerDatabaseVersionsServiceComponentEnumLowerCase = map[string]ListAutonomousContainerDatabaseVersionsServiceComponentEnum{
	"adbd":  ListAutonomousContainerDatabaseVersionsServiceComponentAdbd,
	"exacc": ListAutonomousContainerDatabaseVersionsServiceComponentExacc,
}

// GetListAutonomousContainerDatabaseVersionsServiceComponentEnumValues Enumerates the set of values for ListAutonomousContainerDatabaseVersionsServiceComponentEnum
func GetListAutonomousContainerDatabaseVersionsServiceComponentEnumValues() []ListAutonomousContainerDatabaseVersionsServiceComponentEnum {
	values := make([]ListAutonomousContainerDatabaseVersionsServiceComponentEnum, 0)
	for _, v := range mappingListAutonomousContainerDatabaseVersionsServiceComponentEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousContainerDatabaseVersionsServiceComponentEnumStringValues Enumerates the set of values in String for ListAutonomousContainerDatabaseVersionsServiceComponentEnum
func GetListAutonomousContainerDatabaseVersionsServiceComponentEnumStringValues() []string {
	return []string{
		"ADBD",
		"EXACC",
	}
}

// GetMappingListAutonomousContainerDatabaseVersionsServiceComponentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousContainerDatabaseVersionsServiceComponentEnum(val string) (ListAutonomousContainerDatabaseVersionsServiceComponentEnum, bool) {
	enum, ok := mappingListAutonomousContainerDatabaseVersionsServiceComponentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutonomousContainerDatabaseVersionsSortOrderEnum Enum with underlying type: string
type ListAutonomousContainerDatabaseVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousContainerDatabaseVersionsSortOrderEnum
const (
	ListAutonomousContainerDatabaseVersionsSortOrderAsc  ListAutonomousContainerDatabaseVersionsSortOrderEnum = "ASC"
	ListAutonomousContainerDatabaseVersionsSortOrderDesc ListAutonomousContainerDatabaseVersionsSortOrderEnum = "DESC"
)

var mappingListAutonomousContainerDatabaseVersionsSortOrderEnum = map[string]ListAutonomousContainerDatabaseVersionsSortOrderEnum{
	"ASC":  ListAutonomousContainerDatabaseVersionsSortOrderAsc,
	"DESC": ListAutonomousContainerDatabaseVersionsSortOrderDesc,
}

var mappingListAutonomousContainerDatabaseVersionsSortOrderEnumLowerCase = map[string]ListAutonomousContainerDatabaseVersionsSortOrderEnum{
	"asc":  ListAutonomousContainerDatabaseVersionsSortOrderAsc,
	"desc": ListAutonomousContainerDatabaseVersionsSortOrderDesc,
}

// GetListAutonomousContainerDatabaseVersionsSortOrderEnumValues Enumerates the set of values for ListAutonomousContainerDatabaseVersionsSortOrderEnum
func GetListAutonomousContainerDatabaseVersionsSortOrderEnumValues() []ListAutonomousContainerDatabaseVersionsSortOrderEnum {
	values := make([]ListAutonomousContainerDatabaseVersionsSortOrderEnum, 0)
	for _, v := range mappingListAutonomousContainerDatabaseVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousContainerDatabaseVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListAutonomousContainerDatabaseVersionsSortOrderEnum
func GetListAutonomousContainerDatabaseVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutonomousContainerDatabaseVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousContainerDatabaseVersionsSortOrderEnum(val string) (ListAutonomousContainerDatabaseVersionsSortOrderEnum, bool) {
	enum, ok := mappingListAutonomousContainerDatabaseVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
