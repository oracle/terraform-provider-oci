// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeAwrSourcesSummariesRequest wrapper for the SummarizeAwrSourcesSummaries operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrSourcesSummaries.go.html to see an example of how to use SummarizeAwrSourcesSummariesRequest.
type SummarizeAwrSourcesSummariesRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Name for an Awr source database
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

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

	// The order in which Awr sources summary records are listed
	SortBy SummarizeAwrSourcesSummariesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAwrSourcesSummariesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAwrSourcesSummariesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAwrSourcesSummariesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAwrSourcesSummariesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAwrSourcesSummariesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeAwrSourcesSummariesResponse wrapper for the SummarizeAwrSourcesSummaries operation
type SummarizeAwrSourcesSummariesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeAwrSourcesSummariesCollection instances
	SummarizeAwrSourcesSummariesCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAwrSourcesSummariesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAwrSourcesSummariesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAwrSourcesSummariesSortByEnum Enum with underlying type: string
type SummarizeAwrSourcesSummariesSortByEnum string

// Set of constants representing the allowable values for SummarizeAwrSourcesSummariesSortByEnum
const (
	SummarizeAwrSourcesSummariesSortBySnapshotsuploaded SummarizeAwrSourcesSummariesSortByEnum = "snapshotsUploaded"
	SummarizeAwrSourcesSummariesSortByName              SummarizeAwrSourcesSummariesSortByEnum = "name"
)

var mappingSummarizeAwrSourcesSummariesSortBy = map[string]SummarizeAwrSourcesSummariesSortByEnum{
	"snapshotsUploaded": SummarizeAwrSourcesSummariesSortBySnapshotsuploaded,
	"name":              SummarizeAwrSourcesSummariesSortByName,
}

// GetSummarizeAwrSourcesSummariesSortByEnumValues Enumerates the set of values for SummarizeAwrSourcesSummariesSortByEnum
func GetSummarizeAwrSourcesSummariesSortByEnumValues() []SummarizeAwrSourcesSummariesSortByEnum {
	values := make([]SummarizeAwrSourcesSummariesSortByEnum, 0)
	for _, v := range mappingSummarizeAwrSourcesSummariesSortBy {
		values = append(values, v)
	}
	return values
}

// SummarizeAwrSourcesSummariesSortOrderEnum Enum with underlying type: string
type SummarizeAwrSourcesSummariesSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAwrSourcesSummariesSortOrderEnum
const (
	SummarizeAwrSourcesSummariesSortOrderAsc  SummarizeAwrSourcesSummariesSortOrderEnum = "ASC"
	SummarizeAwrSourcesSummariesSortOrderDesc SummarizeAwrSourcesSummariesSortOrderEnum = "DESC"
)

var mappingSummarizeAwrSourcesSummariesSortOrder = map[string]SummarizeAwrSourcesSummariesSortOrderEnum{
	"ASC":  SummarizeAwrSourcesSummariesSortOrderAsc,
	"DESC": SummarizeAwrSourcesSummariesSortOrderDesc,
}

// GetSummarizeAwrSourcesSummariesSortOrderEnumValues Enumerates the set of values for SummarizeAwrSourcesSummariesSortOrderEnum
func GetSummarizeAwrSourcesSummariesSortOrderEnumValues() []SummarizeAwrSourcesSummariesSortOrderEnum {
	values := make([]SummarizeAwrSourcesSummariesSortOrderEnum, 0)
	for _, v := range mappingSummarizeAwrSourcesSummariesSortOrder {
		values = append(values, v)
	}
	return values
}
