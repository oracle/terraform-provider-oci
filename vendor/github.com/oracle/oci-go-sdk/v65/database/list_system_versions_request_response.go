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

// ListSystemVersionsRequest wrapper for the ListSystemVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListSystemVersions.go.html to see an example of how to use ListSystemVersionsRequest.
type ListSystemVersionsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Specifies shape query parameter.
	Shape *string `mandatory:"true" contributesTo:"query" name:"shape"`

	// Specifies gi version query parameter.
	GiVersion *string `mandatory:"true" contributesTo:"query" name:"giVersion"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSystemVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSystemVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSystemVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSystemVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSystemVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSystemVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSystemVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSystemVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSystemVersionsResponse wrapper for the ListSystemVersions operation
type ListSystemVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SystemVersionCollection instances
	SystemVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSystemVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSystemVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSystemVersionsSortOrderEnum Enum with underlying type: string
type ListSystemVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListSystemVersionsSortOrderEnum
const (
	ListSystemVersionsSortOrderAsc  ListSystemVersionsSortOrderEnum = "ASC"
	ListSystemVersionsSortOrderDesc ListSystemVersionsSortOrderEnum = "DESC"
)

var mappingListSystemVersionsSortOrderEnum = map[string]ListSystemVersionsSortOrderEnum{
	"ASC":  ListSystemVersionsSortOrderAsc,
	"DESC": ListSystemVersionsSortOrderDesc,
}

var mappingListSystemVersionsSortOrderEnumLowerCase = map[string]ListSystemVersionsSortOrderEnum{
	"asc":  ListSystemVersionsSortOrderAsc,
	"desc": ListSystemVersionsSortOrderDesc,
}

// GetListSystemVersionsSortOrderEnumValues Enumerates the set of values for ListSystemVersionsSortOrderEnum
func GetListSystemVersionsSortOrderEnumValues() []ListSystemVersionsSortOrderEnum {
	values := make([]ListSystemVersionsSortOrderEnum, 0)
	for _, v := range mappingListSystemVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSystemVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListSystemVersionsSortOrderEnum
func GetListSystemVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSystemVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSystemVersionsSortOrderEnum(val string) (ListSystemVersionsSortOrderEnum, bool) {
	enum, ok := mappingListSystemVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
