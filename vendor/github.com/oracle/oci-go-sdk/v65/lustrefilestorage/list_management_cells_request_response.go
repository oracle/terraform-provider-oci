// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagementCellsRequest wrapper for the ListManagementCells operation
type ListManagementCellsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ManagementCell.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListManagementCellsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementCellsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementCellsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementCellsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementCellsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagementCellsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagementCellsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagementCellsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagementCellsResponse wrapper for the ListManagementCells operation
type ListManagementCellsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementCellCollection instances
	ManagementCellCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementCellsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementCellsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementCellsSortOrderEnum Enum with underlying type: string
type ListManagementCellsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementCellsSortOrderEnum
const (
	ListManagementCellsSortOrderAsc  ListManagementCellsSortOrderEnum = "ASC"
	ListManagementCellsSortOrderDesc ListManagementCellsSortOrderEnum = "DESC"
)

var mappingListManagementCellsSortOrderEnum = map[string]ListManagementCellsSortOrderEnum{
	"ASC":  ListManagementCellsSortOrderAsc,
	"DESC": ListManagementCellsSortOrderDesc,
}

var mappingListManagementCellsSortOrderEnumLowerCase = map[string]ListManagementCellsSortOrderEnum{
	"asc":  ListManagementCellsSortOrderAsc,
	"desc": ListManagementCellsSortOrderDesc,
}

// GetListManagementCellsSortOrderEnumValues Enumerates the set of values for ListManagementCellsSortOrderEnum
func GetListManagementCellsSortOrderEnumValues() []ListManagementCellsSortOrderEnum {
	values := make([]ListManagementCellsSortOrderEnum, 0)
	for _, v := range mappingListManagementCellsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagementCellsSortOrderEnumStringValues Enumerates the set of values in String for ListManagementCellsSortOrderEnum
func GetListManagementCellsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagementCellsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagementCellsSortOrderEnum(val string) (ListManagementCellsSortOrderEnum, bool) {
	enum, ok := mappingListManagementCellsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
