// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListHostedEntitiesRequest wrapper for the ListHostedEntities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostedEntities.go.html to see an example of how to use ListHostedEntitiesRequest.
type ListHostedEntitiesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Required OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" contributesTo:"query" name:"id"`

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timeInterval is specified, then timeIntervalStart and timeIntervalEnd will be ignored.
	// Examples  P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months), . Maximum value allowed is 25 months prior to current time (P25M).
	AnalysisTimeInterval *string `mandatory:"false" contributesTo:"query" name:"analysisTimeInterval"`

	// Analysis start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// The minimum allowed value is 2 years prior to the current day.
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	// If analysisTimeInterval is specified, this parameter is ignored.
	TimeIntervalStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalStart"`

	// Analysis end time in UTC in ISO 8601 format(exclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd are used together.
	// If timeIntervalEnd is not specified, current time is used as timeIntervalEnd.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalEnd"`

	// Filter by one or more platform types.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS].
	PlatformType []ListHostedEntitiesPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
	ExadataInsightId *string `mandatory:"false" contributesTo:"query" name:"exadataInsightId"`

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
	SortOrder ListHostedEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Hosted entity list sort options.
	SortBy ListHostedEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHostedEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHostedEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHostedEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHostedEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListHostedEntitiesResponse wrapper for the ListHostedEntities operation
type ListHostedEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of HostedEntityCollection instances
	HostedEntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHostedEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHostedEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHostedEntitiesPlatformTypeEnum Enum with underlying type: string
type ListHostedEntitiesPlatformTypeEnum string

// Set of constants representing the allowable values for ListHostedEntitiesPlatformTypeEnum
const (
	ListHostedEntitiesPlatformTypeLinux   ListHostedEntitiesPlatformTypeEnum = "LINUX"
	ListHostedEntitiesPlatformTypeSolaris ListHostedEntitiesPlatformTypeEnum = "SOLARIS"
	ListHostedEntitiesPlatformTypeSunos   ListHostedEntitiesPlatformTypeEnum = "SUNOS"
)

var mappingListHostedEntitiesPlatformType = map[string]ListHostedEntitiesPlatformTypeEnum{
	"LINUX":   ListHostedEntitiesPlatformTypeLinux,
	"SOLARIS": ListHostedEntitiesPlatformTypeSolaris,
	"SUNOS":   ListHostedEntitiesPlatformTypeSunos,
}

// GetListHostedEntitiesPlatformTypeEnumValues Enumerates the set of values for ListHostedEntitiesPlatformTypeEnum
func GetListHostedEntitiesPlatformTypeEnumValues() []ListHostedEntitiesPlatformTypeEnum {
	values := make([]ListHostedEntitiesPlatformTypeEnum, 0)
	for _, v := range mappingListHostedEntitiesPlatformType {
		values = append(values, v)
	}
	return values
}

// ListHostedEntitiesSortOrderEnum Enum with underlying type: string
type ListHostedEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListHostedEntitiesSortOrderEnum
const (
	ListHostedEntitiesSortOrderAsc  ListHostedEntitiesSortOrderEnum = "ASC"
	ListHostedEntitiesSortOrderDesc ListHostedEntitiesSortOrderEnum = "DESC"
)

var mappingListHostedEntitiesSortOrder = map[string]ListHostedEntitiesSortOrderEnum{
	"ASC":  ListHostedEntitiesSortOrderAsc,
	"DESC": ListHostedEntitiesSortOrderDesc,
}

// GetListHostedEntitiesSortOrderEnumValues Enumerates the set of values for ListHostedEntitiesSortOrderEnum
func GetListHostedEntitiesSortOrderEnumValues() []ListHostedEntitiesSortOrderEnum {
	values := make([]ListHostedEntitiesSortOrderEnum, 0)
	for _, v := range mappingListHostedEntitiesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListHostedEntitiesSortByEnum Enum with underlying type: string
type ListHostedEntitiesSortByEnum string

// Set of constants representing the allowable values for ListHostedEntitiesSortByEnum
const (
	ListHostedEntitiesSortByEntityname ListHostedEntitiesSortByEnum = "entityName"
	ListHostedEntitiesSortByEntitytype ListHostedEntitiesSortByEnum = "entityType"
)

var mappingListHostedEntitiesSortBy = map[string]ListHostedEntitiesSortByEnum{
	"entityName": ListHostedEntitiesSortByEntityname,
	"entityType": ListHostedEntitiesSortByEntitytype,
}

// GetListHostedEntitiesSortByEnumValues Enumerates the set of values for ListHostedEntitiesSortByEnum
func GetListHostedEntitiesSortByEnumValues() []ListHostedEntitiesSortByEnum {
	values := make([]ListHostedEntitiesSortByEnum, 0)
	for _, v := range mappingListHostedEntitiesSortBy {
		values = append(values, v)
	}
	return values
}
