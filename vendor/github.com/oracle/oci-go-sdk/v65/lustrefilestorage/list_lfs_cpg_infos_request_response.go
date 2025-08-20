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

// ListLfsCpgInfosRequest wrapper for the ListLfsCpgInfos operation
type ListLfsCpgInfosRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the LFS service CPG.
	LfsCpgId *string `mandatory:"false" contributesTo:"query" name:"lfsCpgId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLfsCpgInfosSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLfsCpgInfosRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLfsCpgInfosRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLfsCpgInfosRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLfsCpgInfosRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLfsCpgInfosRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLfsCpgInfosSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLfsCpgInfosSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLfsCpgInfosResponse wrapper for the ListLfsCpgInfos operation
type ListLfsCpgInfosResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LfsCpgInfoCollection instances
	LfsCpgInfoCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLfsCpgInfosResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLfsCpgInfosResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLfsCpgInfosSortOrderEnum Enum with underlying type: string
type ListLfsCpgInfosSortOrderEnum string

// Set of constants representing the allowable values for ListLfsCpgInfosSortOrderEnum
const (
	ListLfsCpgInfosSortOrderAsc  ListLfsCpgInfosSortOrderEnum = "ASC"
	ListLfsCpgInfosSortOrderDesc ListLfsCpgInfosSortOrderEnum = "DESC"
)

var mappingListLfsCpgInfosSortOrderEnum = map[string]ListLfsCpgInfosSortOrderEnum{
	"ASC":  ListLfsCpgInfosSortOrderAsc,
	"DESC": ListLfsCpgInfosSortOrderDesc,
}

var mappingListLfsCpgInfosSortOrderEnumLowerCase = map[string]ListLfsCpgInfosSortOrderEnum{
	"asc":  ListLfsCpgInfosSortOrderAsc,
	"desc": ListLfsCpgInfosSortOrderDesc,
}

// GetListLfsCpgInfosSortOrderEnumValues Enumerates the set of values for ListLfsCpgInfosSortOrderEnum
func GetListLfsCpgInfosSortOrderEnumValues() []ListLfsCpgInfosSortOrderEnum {
	values := make([]ListLfsCpgInfosSortOrderEnum, 0)
	for _, v := range mappingListLfsCpgInfosSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLfsCpgInfosSortOrderEnumStringValues Enumerates the set of values in String for ListLfsCpgInfosSortOrderEnum
func GetListLfsCpgInfosSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLfsCpgInfosSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLfsCpgInfosSortOrderEnum(val string) (ListLfsCpgInfosSortOrderEnum, bool) {
	enum, ok := mappingListLfsCpgInfosSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
