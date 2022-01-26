// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListImportableAgentEntitiesRequest wrapper for the ListImportableAgentEntities operation
//
// See also
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

var mappingListImportableAgentEntitiesSortOrder = map[string]ListImportableAgentEntitiesSortOrderEnum{
	"ASC":  ListImportableAgentEntitiesSortOrderAsc,
	"DESC": ListImportableAgentEntitiesSortOrderDesc,
}

// GetListImportableAgentEntitiesSortOrderEnumValues Enumerates the set of values for ListImportableAgentEntitiesSortOrderEnum
func GetListImportableAgentEntitiesSortOrderEnumValues() []ListImportableAgentEntitiesSortOrderEnum {
	values := make([]ListImportableAgentEntitiesSortOrderEnum, 0)
	for _, v := range mappingListImportableAgentEntitiesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListImportableAgentEntitiesSortByEnum Enum with underlying type: string
type ListImportableAgentEntitiesSortByEnum string

// Set of constants representing the allowable values for ListImportableAgentEntitiesSortByEnum
const (
	ListImportableAgentEntitiesSortByEntityname ListImportableAgentEntitiesSortByEnum = "entityName"
	ListImportableAgentEntitiesSortByEntitytype ListImportableAgentEntitiesSortByEnum = "entityType"
)

var mappingListImportableAgentEntitiesSortBy = map[string]ListImportableAgentEntitiesSortByEnum{
	"entityName": ListImportableAgentEntitiesSortByEntityname,
	"entityType": ListImportableAgentEntitiesSortByEntitytype,
}

// GetListImportableAgentEntitiesSortByEnumValues Enumerates the set of values for ListImportableAgentEntitiesSortByEnum
func GetListImportableAgentEntitiesSortByEnumValues() []ListImportableAgentEntitiesSortByEnum {
	values := make([]ListImportableAgentEntitiesSortByEnum, 0)
	for _, v := range mappingListImportableAgentEntitiesSortBy {
		values = append(values, v)
	}
	return values
}
