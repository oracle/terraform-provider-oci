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

// ListImportableAgentEntitiesRequest wrapper for the ListImportableAgentEntities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListImportableAgentEntities.go.html to see an example of how to use ListImportableAgentEntitiesRequest.
type ListImportableAgentEntitiesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	SortOrder ListImportableAgentEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Hosted entity list sort options.
	SortBy ListImportableAgentEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListImportableAgentEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListImportableAgentEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListImportableAgentEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListImportableAgentEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListImportableAgentEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListImportableAgentEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListImportableAgentEntitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListImportableAgentEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListImportableAgentEntitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListImportableAgentEntitiesResponse wrapper for the ListImportableAgentEntities operation
type ListImportableAgentEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ImportableAgentEntitySummaryCollection instances
	ImportableAgentEntitySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListImportableAgentEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListImportableAgentEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListImportableAgentEntitiesSortOrderEnum Enum with underlying type: string
type ListImportableAgentEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListImportableAgentEntitiesSortOrderEnum
const (
	ListImportableAgentEntitiesSortOrderAsc  ListImportableAgentEntitiesSortOrderEnum = "ASC"
	ListImportableAgentEntitiesSortOrderDesc ListImportableAgentEntitiesSortOrderEnum = "DESC"
)

var mappingListImportableAgentEntitiesSortOrderEnum = map[string]ListImportableAgentEntitiesSortOrderEnum{
	"ASC":  ListImportableAgentEntitiesSortOrderAsc,
	"DESC": ListImportableAgentEntitiesSortOrderDesc,
}

var mappingListImportableAgentEntitiesSortOrderEnumLowerCase = map[string]ListImportableAgentEntitiesSortOrderEnum{
	"asc":  ListImportableAgentEntitiesSortOrderAsc,
	"desc": ListImportableAgentEntitiesSortOrderDesc,
}

// GetListImportableAgentEntitiesSortOrderEnumValues Enumerates the set of values for ListImportableAgentEntitiesSortOrderEnum
func GetListImportableAgentEntitiesSortOrderEnumValues() []ListImportableAgentEntitiesSortOrderEnum {
	values := make([]ListImportableAgentEntitiesSortOrderEnum, 0)
	for _, v := range mappingListImportableAgentEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportableAgentEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListImportableAgentEntitiesSortOrderEnum
func GetListImportableAgentEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListImportableAgentEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportableAgentEntitiesSortOrderEnum(val string) (ListImportableAgentEntitiesSortOrderEnum, bool) {
	enum, ok := mappingListImportableAgentEntitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListImportableAgentEntitiesSortByEnum Enum with underlying type: string
type ListImportableAgentEntitiesSortByEnum string

// Set of constants representing the allowable values for ListImportableAgentEntitiesSortByEnum
const (
	ListImportableAgentEntitiesSortByEntityname ListImportableAgentEntitiesSortByEnum = "entityName"
	ListImportableAgentEntitiesSortByEntitytype ListImportableAgentEntitiesSortByEnum = "entityType"
)

var mappingListImportableAgentEntitiesSortByEnum = map[string]ListImportableAgentEntitiesSortByEnum{
	"entityName": ListImportableAgentEntitiesSortByEntityname,
	"entityType": ListImportableAgentEntitiesSortByEntitytype,
}

var mappingListImportableAgentEntitiesSortByEnumLowerCase = map[string]ListImportableAgentEntitiesSortByEnum{
	"entityname": ListImportableAgentEntitiesSortByEntityname,
	"entitytype": ListImportableAgentEntitiesSortByEntitytype,
}

// GetListImportableAgentEntitiesSortByEnumValues Enumerates the set of values for ListImportableAgentEntitiesSortByEnum
func GetListImportableAgentEntitiesSortByEnumValues() []ListImportableAgentEntitiesSortByEnum {
	values := make([]ListImportableAgentEntitiesSortByEnum, 0)
	for _, v := range mappingListImportableAgentEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListImportableAgentEntitiesSortByEnumStringValues Enumerates the set of values in String for ListImportableAgentEntitiesSortByEnum
func GetListImportableAgentEntitiesSortByEnumStringValues() []string {
	return []string{
		"entityName",
		"entityType",
	}
}

// GetMappingListImportableAgentEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListImportableAgentEntitiesSortByEnum(val string) (ListImportableAgentEntitiesSortByEnum, bool) {
	enum, ok := mappingListImportableAgentEntitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
