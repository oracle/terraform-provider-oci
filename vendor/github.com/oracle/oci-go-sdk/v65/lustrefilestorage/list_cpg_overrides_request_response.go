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

// ListCpgOverridesRequest wrapper for the ListCpgOverrides operation
type ListCpgOverridesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer CPG.
	CustomerCpgId *string `mandatory:"false" contributesTo:"query" name:"customerCpgId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer tenancy.
	CustomerTenancyId *string `mandatory:"false" contributesTo:"query" name:"customerTenancyId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCpgOverridesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCpgOverridesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCpgOverridesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCpgOverridesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCpgOverridesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCpgOverridesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCpgOverridesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCpgOverridesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCpgOverridesResponse wrapper for the ListCpgOverrides operation
type ListCpgOverridesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CpgOverrideCollection instances
	CpgOverrideCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCpgOverridesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCpgOverridesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCpgOverridesSortOrderEnum Enum with underlying type: string
type ListCpgOverridesSortOrderEnum string

// Set of constants representing the allowable values for ListCpgOverridesSortOrderEnum
const (
	ListCpgOverridesSortOrderAsc  ListCpgOverridesSortOrderEnum = "ASC"
	ListCpgOverridesSortOrderDesc ListCpgOverridesSortOrderEnum = "DESC"
)

var mappingListCpgOverridesSortOrderEnum = map[string]ListCpgOverridesSortOrderEnum{
	"ASC":  ListCpgOverridesSortOrderAsc,
	"DESC": ListCpgOverridesSortOrderDesc,
}

var mappingListCpgOverridesSortOrderEnumLowerCase = map[string]ListCpgOverridesSortOrderEnum{
	"asc":  ListCpgOverridesSortOrderAsc,
	"desc": ListCpgOverridesSortOrderDesc,
}

// GetListCpgOverridesSortOrderEnumValues Enumerates the set of values for ListCpgOverridesSortOrderEnum
func GetListCpgOverridesSortOrderEnumValues() []ListCpgOverridesSortOrderEnum {
	values := make([]ListCpgOverridesSortOrderEnum, 0)
	for _, v := range mappingListCpgOverridesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCpgOverridesSortOrderEnumStringValues Enumerates the set of values in String for ListCpgOverridesSortOrderEnum
func GetListCpgOverridesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCpgOverridesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCpgOverridesSortOrderEnum(val string) (ListCpgOverridesSortOrderEnum, bool) {
	enum, ok := mappingListCpgOverridesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
