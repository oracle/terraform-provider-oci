// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNewsReportsRequest wrapper for the ListNewsReports operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListNewsReports.go.html to see an example of how to use ListNewsReportsRequest.
type ListNewsReportsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique Ops Insights news report identifier
	NewsReportId *string `mandatory:"false" contributesTo:"query" name:"newsReportId"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Lifecycle states
	LifecycleState []LifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListNewsReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// News report list sort options. If `fields` parameter is selected, the `sortBy` parameter must be one of the fields specified.
	SortBy ListNewsReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNewsReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNewsReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNewsReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNewsReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNewsReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Status {
		if _, ok := GetMappingResourceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetResourceStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.LifecycleState {
		if _, ok := GetMappingLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListNewsReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNewsReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNewsReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNewsReportsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNewsReportsResponse wrapper for the ListNewsReports operation
type ListNewsReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NewsReportCollection instances
	NewsReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNewsReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNewsReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNewsReportsSortOrderEnum Enum with underlying type: string
type ListNewsReportsSortOrderEnum string

// Set of constants representing the allowable values for ListNewsReportsSortOrderEnum
const (
	ListNewsReportsSortOrderAsc  ListNewsReportsSortOrderEnum = "ASC"
	ListNewsReportsSortOrderDesc ListNewsReportsSortOrderEnum = "DESC"
)

var mappingListNewsReportsSortOrderEnum = map[string]ListNewsReportsSortOrderEnum{
	"ASC":  ListNewsReportsSortOrderAsc,
	"DESC": ListNewsReportsSortOrderDesc,
}

var mappingListNewsReportsSortOrderEnumLowerCase = map[string]ListNewsReportsSortOrderEnum{
	"asc":  ListNewsReportsSortOrderAsc,
	"desc": ListNewsReportsSortOrderDesc,
}

// GetListNewsReportsSortOrderEnumValues Enumerates the set of values for ListNewsReportsSortOrderEnum
func GetListNewsReportsSortOrderEnumValues() []ListNewsReportsSortOrderEnum {
	values := make([]ListNewsReportsSortOrderEnum, 0)
	for _, v := range mappingListNewsReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNewsReportsSortOrderEnumStringValues Enumerates the set of values in String for ListNewsReportsSortOrderEnum
func GetListNewsReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNewsReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNewsReportsSortOrderEnum(val string) (ListNewsReportsSortOrderEnum, bool) {
	enum, ok := mappingListNewsReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNewsReportsSortByEnum Enum with underlying type: string
type ListNewsReportsSortByEnum string

// Set of constants representing the allowable values for ListNewsReportsSortByEnum
const (
	ListNewsReportsSortByName          ListNewsReportsSortByEnum = "name"
	ListNewsReportsSortByNewsfrequency ListNewsReportsSortByEnum = "newsFrequency"
)

var mappingListNewsReportsSortByEnum = map[string]ListNewsReportsSortByEnum{
	"name":          ListNewsReportsSortByName,
	"newsFrequency": ListNewsReportsSortByNewsfrequency,
}

var mappingListNewsReportsSortByEnumLowerCase = map[string]ListNewsReportsSortByEnum{
	"name":          ListNewsReportsSortByName,
	"newsfrequency": ListNewsReportsSortByNewsfrequency,
}

// GetListNewsReportsSortByEnumValues Enumerates the set of values for ListNewsReportsSortByEnum
func GetListNewsReportsSortByEnumValues() []ListNewsReportsSortByEnum {
	values := make([]ListNewsReportsSortByEnum, 0)
	for _, v := range mappingListNewsReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNewsReportsSortByEnumStringValues Enumerates the set of values in String for ListNewsReportsSortByEnum
func GetListNewsReportsSortByEnumStringValues() []string {
	return []string{
		"name",
		"newsFrequency",
	}
}

// GetMappingListNewsReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNewsReportsSortByEnum(val string) (ListNewsReportsSortByEnum, bool) {
	enum, ok := mappingListNewsReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
