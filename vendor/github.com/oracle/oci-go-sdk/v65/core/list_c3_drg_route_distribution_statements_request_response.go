// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListC3DrgRouteDistributionStatementsRequest wrapper for the ListC3DrgRouteDistributionStatements operation
type ListC3DrgRouteDistributionStatementsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route distribution.
	DrgRouteDistributionId *string `mandatory:"true" contributesTo:"path" name:"drgRouteDistributionId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.
	SortBy ListC3DrgRouteDistributionStatementsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListC3DrgRouteDistributionStatementsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListC3DrgRouteDistributionStatementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListC3DrgRouteDistributionStatementsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListC3DrgRouteDistributionStatementsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListC3DrgRouteDistributionStatementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListC3DrgRouteDistributionStatementsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListC3DrgRouteDistributionStatementsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListC3DrgRouteDistributionStatementsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListC3DrgRouteDistributionStatementsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListC3DrgRouteDistributionStatementsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListC3DrgRouteDistributionStatementsResponse wrapper for the ListC3DrgRouteDistributionStatements operation
type ListC3DrgRouteDistributionStatementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DrgRouteDistributionStatement instances
	Items []DrgRouteDistributionStatement `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListC3DrgRouteDistributionStatementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListC3DrgRouteDistributionStatementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListC3DrgRouteDistributionStatementsSortByEnum Enum with underlying type: string
type ListC3DrgRouteDistributionStatementsSortByEnum string

// Set of constants representing the allowable values for ListC3DrgRouteDistributionStatementsSortByEnum
const (
	ListC3DrgRouteDistributionStatementsSortByTimecreated ListC3DrgRouteDistributionStatementsSortByEnum = "TIMECREATED"
)

var mappingListC3DrgRouteDistributionStatementsSortByEnum = map[string]ListC3DrgRouteDistributionStatementsSortByEnum{
	"TIMECREATED": ListC3DrgRouteDistributionStatementsSortByTimecreated,
}

var mappingListC3DrgRouteDistributionStatementsSortByEnumLowerCase = map[string]ListC3DrgRouteDistributionStatementsSortByEnum{
	"timecreated": ListC3DrgRouteDistributionStatementsSortByTimecreated,
}

// GetListC3DrgRouteDistributionStatementsSortByEnumValues Enumerates the set of values for ListC3DrgRouteDistributionStatementsSortByEnum
func GetListC3DrgRouteDistributionStatementsSortByEnumValues() []ListC3DrgRouteDistributionStatementsSortByEnum {
	values := make([]ListC3DrgRouteDistributionStatementsSortByEnum, 0)
	for _, v := range mappingListC3DrgRouteDistributionStatementsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListC3DrgRouteDistributionStatementsSortByEnumStringValues Enumerates the set of values in String for ListC3DrgRouteDistributionStatementsSortByEnum
func GetListC3DrgRouteDistributionStatementsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
	}
}

// GetMappingListC3DrgRouteDistributionStatementsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListC3DrgRouteDistributionStatementsSortByEnum(val string) (ListC3DrgRouteDistributionStatementsSortByEnum, bool) {
	enum, ok := mappingListC3DrgRouteDistributionStatementsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListC3DrgRouteDistributionStatementsSortOrderEnum Enum with underlying type: string
type ListC3DrgRouteDistributionStatementsSortOrderEnum string

// Set of constants representing the allowable values for ListC3DrgRouteDistributionStatementsSortOrderEnum
const (
	ListC3DrgRouteDistributionStatementsSortOrderAsc  ListC3DrgRouteDistributionStatementsSortOrderEnum = "ASC"
	ListC3DrgRouteDistributionStatementsSortOrderDesc ListC3DrgRouteDistributionStatementsSortOrderEnum = "DESC"
)

var mappingListC3DrgRouteDistributionStatementsSortOrderEnum = map[string]ListC3DrgRouteDistributionStatementsSortOrderEnum{
	"ASC":  ListC3DrgRouteDistributionStatementsSortOrderAsc,
	"DESC": ListC3DrgRouteDistributionStatementsSortOrderDesc,
}

var mappingListC3DrgRouteDistributionStatementsSortOrderEnumLowerCase = map[string]ListC3DrgRouteDistributionStatementsSortOrderEnum{
	"asc":  ListC3DrgRouteDistributionStatementsSortOrderAsc,
	"desc": ListC3DrgRouteDistributionStatementsSortOrderDesc,
}

// GetListC3DrgRouteDistributionStatementsSortOrderEnumValues Enumerates the set of values for ListC3DrgRouteDistributionStatementsSortOrderEnum
func GetListC3DrgRouteDistributionStatementsSortOrderEnumValues() []ListC3DrgRouteDistributionStatementsSortOrderEnum {
	values := make([]ListC3DrgRouteDistributionStatementsSortOrderEnum, 0)
	for _, v := range mappingListC3DrgRouteDistributionStatementsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListC3DrgRouteDistributionStatementsSortOrderEnumStringValues Enumerates the set of values in String for ListC3DrgRouteDistributionStatementsSortOrderEnum
func GetListC3DrgRouteDistributionStatementsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListC3DrgRouteDistributionStatementsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListC3DrgRouteDistributionStatementsSortOrderEnum(val string) (ListC3DrgRouteDistributionStatementsSortOrderEnum, bool) {
	enum, ok := mappingListC3DrgRouteDistributionStatementsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
