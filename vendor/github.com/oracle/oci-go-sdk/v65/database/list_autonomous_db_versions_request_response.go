// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousDbVersionsRequest wrapper for the ListAutonomousDbVersions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDbVersions.go.html to see an example of how to use ListAutonomousDbVersionsRequest.
type ListAutonomousDbVersionsRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only autonomous database resources that match the specified workload type.
	DbWorkload AutonomousDatabaseSummaryDbWorkloadEnum `mandatory:"false" contributesTo:"query" name:"dbWorkload" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousDbVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDbVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDbVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousDbVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDbVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousDbVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseSummaryDbWorkloadEnum(string(request.DbWorkload)); !ok && request.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", request.DbWorkload, strings.Join(GetAutonomousDatabaseSummaryDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousDbVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutonomousDbVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousDbVersionsResponse wrapper for the ListAutonomousDbVersions operation
type ListAutonomousDbVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousDbVersionSummary instances
	Items []AutonomousDbVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDbVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDbVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousDbVersionsSortOrderEnum Enum with underlying type: string
type ListAutonomousDbVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousDbVersionsSortOrderEnum
const (
	ListAutonomousDbVersionsSortOrderAsc  ListAutonomousDbVersionsSortOrderEnum = "ASC"
	ListAutonomousDbVersionsSortOrderDesc ListAutonomousDbVersionsSortOrderEnum = "DESC"
)

var mappingListAutonomousDbVersionsSortOrderEnum = map[string]ListAutonomousDbVersionsSortOrderEnum{
	"ASC":  ListAutonomousDbVersionsSortOrderAsc,
	"DESC": ListAutonomousDbVersionsSortOrderDesc,
}

var mappingListAutonomousDbVersionsSortOrderEnumLowerCase = map[string]ListAutonomousDbVersionsSortOrderEnum{
	"asc":  ListAutonomousDbVersionsSortOrderAsc,
	"desc": ListAutonomousDbVersionsSortOrderDesc,
}

// GetListAutonomousDbVersionsSortOrderEnumValues Enumerates the set of values for ListAutonomousDbVersionsSortOrderEnum
func GetListAutonomousDbVersionsSortOrderEnumValues() []ListAutonomousDbVersionsSortOrderEnum {
	values := make([]ListAutonomousDbVersionsSortOrderEnum, 0)
	for _, v := range mappingListAutonomousDbVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDbVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListAutonomousDbVersionsSortOrderEnum
func GetListAutonomousDbVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutonomousDbVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDbVersionsSortOrderEnum(val string) (ListAutonomousDbVersionsSortOrderEnum, bool) {
	enum, ok := mappingListAutonomousDbVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
