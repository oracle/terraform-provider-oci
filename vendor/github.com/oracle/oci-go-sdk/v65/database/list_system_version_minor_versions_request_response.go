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

// ListSystemVersionMinorVersionsRequest wrapper for the ListSystemVersionMinorVersions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListSystemVersionMinorVersions.go.html to see an example of how to use ListSystemVersionMinorVersionsRequest.
type ListSystemVersionMinorVersionsRequest struct {

	// The System major version.
	MajorVersion *string `mandatory:"true" contributesTo:"path" name:"majorVersion"`

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Specifies gi version query parameter.
	GiVersion *string `mandatory:"true" contributesTo:"query" name:"giVersion"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSystemVersionMinorVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If provided, filters the results for the given shape.
	Shape *string `mandatory:"false" contributesTo:"query" name:"shape"`

	// If provided, filters the results for the specified resource Id.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// If provided, return highest versions from each major version family.
	IsLatest *bool `mandatory:"false" contributesTo:"query" name:"isLatest"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSystemVersionMinorVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSystemVersionMinorVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSystemVersionMinorVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSystemVersionMinorVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSystemVersionMinorVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSystemVersionMinorVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSystemVersionMinorVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSystemVersionMinorVersionsResponse wrapper for the ListSystemVersionMinorVersions operation
type ListSystemVersionMinorVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SystemVersionMinorVersionCollection instances
	SystemVersionMinorVersionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSystemVersionMinorVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSystemVersionMinorVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSystemVersionMinorVersionsSortOrderEnum Enum with underlying type: string
type ListSystemVersionMinorVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListSystemVersionMinorVersionsSortOrderEnum
const (
	ListSystemVersionMinorVersionsSortOrderAsc  ListSystemVersionMinorVersionsSortOrderEnum = "ASC"
	ListSystemVersionMinorVersionsSortOrderDesc ListSystemVersionMinorVersionsSortOrderEnum = "DESC"
)

var mappingListSystemVersionMinorVersionsSortOrderEnum = map[string]ListSystemVersionMinorVersionsSortOrderEnum{
	"ASC":  ListSystemVersionMinorVersionsSortOrderAsc,
	"DESC": ListSystemVersionMinorVersionsSortOrderDesc,
}

var mappingListSystemVersionMinorVersionsSortOrderEnumLowerCase = map[string]ListSystemVersionMinorVersionsSortOrderEnum{
	"asc":  ListSystemVersionMinorVersionsSortOrderAsc,
	"desc": ListSystemVersionMinorVersionsSortOrderDesc,
}

// GetListSystemVersionMinorVersionsSortOrderEnumValues Enumerates the set of values for ListSystemVersionMinorVersionsSortOrderEnum
func GetListSystemVersionMinorVersionsSortOrderEnumValues() []ListSystemVersionMinorVersionsSortOrderEnum {
	values := make([]ListSystemVersionMinorVersionsSortOrderEnum, 0)
	for _, v := range mappingListSystemVersionMinorVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSystemVersionMinorVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListSystemVersionMinorVersionsSortOrderEnum
func GetListSystemVersionMinorVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSystemVersionMinorVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSystemVersionMinorVersionsSortOrderEnum(val string) (ListSystemVersionMinorVersionsSortOrderEnum, bool) {
	enum, ok := mappingListSystemVersionMinorVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
