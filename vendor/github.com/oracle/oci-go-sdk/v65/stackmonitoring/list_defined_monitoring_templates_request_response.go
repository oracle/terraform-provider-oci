// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDefinedMonitoringTemplatesRequest wrapper for the ListDefinedMonitoringTemplates operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListDefinedMonitoringTemplates.go.html to see an example of how to use ListDefinedMonitoringTemplatesRequest.
type ListDefinedMonitoringTemplatesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy(root) for which
	// defined monitored templates should be listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for 'namespace' is ascending.
	SortBy ListDefinedMonitoringTemplatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return monitoring template based on name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDefinedMonitoringTemplatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Multiple resource types filter.
	ResourceTypes []string `contributesTo:"query" name:"resourceTypes" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDefinedMonitoringTemplatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDefinedMonitoringTemplatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDefinedMonitoringTemplatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDefinedMonitoringTemplatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDefinedMonitoringTemplatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDefinedMonitoringTemplatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDefinedMonitoringTemplatesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDefinedMonitoringTemplatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDefinedMonitoringTemplatesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDefinedMonitoringTemplatesResponse wrapper for the ListDefinedMonitoringTemplates operation
type ListDefinedMonitoringTemplatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DefinedMonitoringTemplateCollection instances
	DefinedMonitoringTemplateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDefinedMonitoringTemplatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDefinedMonitoringTemplatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDefinedMonitoringTemplatesSortByEnum Enum with underlying type: string
type ListDefinedMonitoringTemplatesSortByEnum string

// Set of constants representing the allowable values for ListDefinedMonitoringTemplatesSortByEnum
const (
	ListDefinedMonitoringTemplatesSortByNamespace ListDefinedMonitoringTemplatesSortByEnum = "namespace"
)

var mappingListDefinedMonitoringTemplatesSortByEnum = map[string]ListDefinedMonitoringTemplatesSortByEnum{
	"namespace": ListDefinedMonitoringTemplatesSortByNamespace,
}

var mappingListDefinedMonitoringTemplatesSortByEnumLowerCase = map[string]ListDefinedMonitoringTemplatesSortByEnum{
	"namespace": ListDefinedMonitoringTemplatesSortByNamespace,
}

// GetListDefinedMonitoringTemplatesSortByEnumValues Enumerates the set of values for ListDefinedMonitoringTemplatesSortByEnum
func GetListDefinedMonitoringTemplatesSortByEnumValues() []ListDefinedMonitoringTemplatesSortByEnum {
	values := make([]ListDefinedMonitoringTemplatesSortByEnum, 0)
	for _, v := range mappingListDefinedMonitoringTemplatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDefinedMonitoringTemplatesSortByEnumStringValues Enumerates the set of values in String for ListDefinedMonitoringTemplatesSortByEnum
func GetListDefinedMonitoringTemplatesSortByEnumStringValues() []string {
	return []string{
		"namespace",
	}
}

// GetMappingListDefinedMonitoringTemplatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDefinedMonitoringTemplatesSortByEnum(val string) (ListDefinedMonitoringTemplatesSortByEnum, bool) {
	enum, ok := mappingListDefinedMonitoringTemplatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDefinedMonitoringTemplatesSortOrderEnum Enum with underlying type: string
type ListDefinedMonitoringTemplatesSortOrderEnum string

// Set of constants representing the allowable values for ListDefinedMonitoringTemplatesSortOrderEnum
const (
	ListDefinedMonitoringTemplatesSortOrderAsc  ListDefinedMonitoringTemplatesSortOrderEnum = "ASC"
	ListDefinedMonitoringTemplatesSortOrderDesc ListDefinedMonitoringTemplatesSortOrderEnum = "DESC"
)

var mappingListDefinedMonitoringTemplatesSortOrderEnum = map[string]ListDefinedMonitoringTemplatesSortOrderEnum{
	"ASC":  ListDefinedMonitoringTemplatesSortOrderAsc,
	"DESC": ListDefinedMonitoringTemplatesSortOrderDesc,
}

var mappingListDefinedMonitoringTemplatesSortOrderEnumLowerCase = map[string]ListDefinedMonitoringTemplatesSortOrderEnum{
	"asc":  ListDefinedMonitoringTemplatesSortOrderAsc,
	"desc": ListDefinedMonitoringTemplatesSortOrderDesc,
}

// GetListDefinedMonitoringTemplatesSortOrderEnumValues Enumerates the set of values for ListDefinedMonitoringTemplatesSortOrderEnum
func GetListDefinedMonitoringTemplatesSortOrderEnumValues() []ListDefinedMonitoringTemplatesSortOrderEnum {
	values := make([]ListDefinedMonitoringTemplatesSortOrderEnum, 0)
	for _, v := range mappingListDefinedMonitoringTemplatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDefinedMonitoringTemplatesSortOrderEnumStringValues Enumerates the set of values in String for ListDefinedMonitoringTemplatesSortOrderEnum
func GetListDefinedMonitoringTemplatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDefinedMonitoringTemplatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDefinedMonitoringTemplatesSortOrderEnum(val string) (ListDefinedMonitoringTemplatesSortOrderEnum, bool) {
	enum, ok := mappingListDefinedMonitoringTemplatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
