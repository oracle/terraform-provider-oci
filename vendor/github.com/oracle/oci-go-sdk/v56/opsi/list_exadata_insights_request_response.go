// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListExadataInsightsRequest wrapper for the ListExadataInsights operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListExadataInsights.go.html to see an example of how to use ListExadataInsightsRequest.
type ListExadataInsightsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique Enterprise Manager bridge identifier
	EnterpriseManagerBridgeId *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerBridgeId"`

	// Optional list of Exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Lifecycle states
	LifecycleState []LifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// Filter by one or more Exadata types.
	// Possible value are DBMACHINE, EXACS, and EXACC.
	ExadataType []string `contributesTo:"query" name:"exadataType" collectionFormat:"multi"`

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
	SortOrder ListExadataInsightsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Exadata insight list sort options. If `fields` parameter is selected, the `sortBy` parameter must be one of the fields specified. Default order for timeCreated is descending. Default order for exadataName is ascending. If no value is specified timeCreated is default.
	SortBy ListExadataInsightsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExadataInsightsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExadataInsightsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExadataInsightsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExadataInsightsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListExadataInsightsResponse wrapper for the ListExadataInsights operation
type ListExadataInsightsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExadataInsightSummaryCollection instances
	ExadataInsightSummaryCollection `presentIn:"body"`

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

func (response ListExadataInsightsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExadataInsightsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExadataInsightsSortOrderEnum Enum with underlying type: string
type ListExadataInsightsSortOrderEnum string

// Set of constants representing the allowable values for ListExadataInsightsSortOrderEnum
const (
	ListExadataInsightsSortOrderAsc  ListExadataInsightsSortOrderEnum = "ASC"
	ListExadataInsightsSortOrderDesc ListExadataInsightsSortOrderEnum = "DESC"
)

var mappingListExadataInsightsSortOrder = map[string]ListExadataInsightsSortOrderEnum{
	"ASC":  ListExadataInsightsSortOrderAsc,
	"DESC": ListExadataInsightsSortOrderDesc,
}

// GetListExadataInsightsSortOrderEnumValues Enumerates the set of values for ListExadataInsightsSortOrderEnum
func GetListExadataInsightsSortOrderEnumValues() []ListExadataInsightsSortOrderEnum {
	values := make([]ListExadataInsightsSortOrderEnum, 0)
	for _, v := range mappingListExadataInsightsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListExadataInsightsSortByEnum Enum with underlying type: string
type ListExadataInsightsSortByEnum string

// Set of constants representing the allowable values for ListExadataInsightsSortByEnum
const (
	ListExadataInsightsSortByTimecreated ListExadataInsightsSortByEnum = "timeCreated"
	ListExadataInsightsSortByExadataname ListExadataInsightsSortByEnum = "exadataName"
)

var mappingListExadataInsightsSortBy = map[string]ListExadataInsightsSortByEnum{
	"timeCreated": ListExadataInsightsSortByTimecreated,
	"exadataName": ListExadataInsightsSortByExadataname,
}

// GetListExadataInsightsSortByEnumValues Enumerates the set of values for ListExadataInsightsSortByEnum
func GetListExadataInsightsSortByEnumValues() []ListExadataInsightsSortByEnum {
	values := make([]ListExadataInsightsSortByEnum, 0)
	for _, v := range mappingListExadataInsightsSortBy {
		values = append(values, v)
	}
	return values
}
