// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDatabaseInsightsRequest wrapper for the ListDatabaseInsights operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseInsights.go.html to see an example of how to use ListDatabaseInsightsRequest.
type ListDatabaseInsightsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique Enterprise Manager bridge identifier
	EnterpriseManagerBridgeId *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerBridgeId"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Lifecycle states
	LifecycleState []LifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []ListDatabaseInsightsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Specifies the fields to return in a database summary response. By default all fields are returned if omitted.
	Fields []ListDatabaseInsightsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

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
	SortOrder ListDatabaseInsightsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Database insight list sort options. If `fields` parameter is selected, the `sortBy` parameter must be one of the fields specified.
	SortBy ListDatabaseInsightsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
	ExadataInsightId *string `mandatory:"false" contributesTo:"query" name:"exadataInsightId"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseInsightsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseInsightsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseInsightsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseInsightsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDatabaseInsightsResponse wrapper for the ListDatabaseInsights operation
type ListDatabaseInsightsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseInsightsCollection instances
	DatabaseInsightsCollection `presentIn:"body"`

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

func (response ListDatabaseInsightsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseInsightsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseInsightsDatabaseTypeEnum Enum with underlying type: string
type ListDatabaseInsightsDatabaseTypeEnum string

// Set of constants representing the allowable values for ListDatabaseInsightsDatabaseTypeEnum
const (
	ListDatabaseInsightsDatabaseTypeAdwS           ListDatabaseInsightsDatabaseTypeEnum = "ADW-S"
	ListDatabaseInsightsDatabaseTypeAtpS           ListDatabaseInsightsDatabaseTypeEnum = "ATP-S"
	ListDatabaseInsightsDatabaseTypeAdwD           ListDatabaseInsightsDatabaseTypeEnum = "ADW-D"
	ListDatabaseInsightsDatabaseTypeAtpD           ListDatabaseInsightsDatabaseTypeEnum = "ATP-D"
	ListDatabaseInsightsDatabaseTypeExternalPdb    ListDatabaseInsightsDatabaseTypeEnum = "EXTERNAL-PDB"
	ListDatabaseInsightsDatabaseTypeExternalNoncdb ListDatabaseInsightsDatabaseTypeEnum = "EXTERNAL-NONCDB"
)

var mappingListDatabaseInsightsDatabaseType = map[string]ListDatabaseInsightsDatabaseTypeEnum{
	"ADW-S":           ListDatabaseInsightsDatabaseTypeAdwS,
	"ATP-S":           ListDatabaseInsightsDatabaseTypeAtpS,
	"ADW-D":           ListDatabaseInsightsDatabaseTypeAdwD,
	"ATP-D":           ListDatabaseInsightsDatabaseTypeAtpD,
	"EXTERNAL-PDB":    ListDatabaseInsightsDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB": ListDatabaseInsightsDatabaseTypeExternalNoncdb,
}

// GetListDatabaseInsightsDatabaseTypeEnumValues Enumerates the set of values for ListDatabaseInsightsDatabaseTypeEnum
func GetListDatabaseInsightsDatabaseTypeEnumValues() []ListDatabaseInsightsDatabaseTypeEnum {
	values := make([]ListDatabaseInsightsDatabaseTypeEnum, 0)
	for _, v := range mappingListDatabaseInsightsDatabaseType {
		values = append(values, v)
	}
	return values
}

// ListDatabaseInsightsFieldsEnum Enum with underlying type: string
type ListDatabaseInsightsFieldsEnum string

// Set of constants representing the allowable values for ListDatabaseInsightsFieldsEnum
const (
	ListDatabaseInsightsFieldsCompartmentid       ListDatabaseInsightsFieldsEnum = "compartmentId"
	ListDatabaseInsightsFieldsDatabasename        ListDatabaseInsightsFieldsEnum = "databaseName"
	ListDatabaseInsightsFieldsDatabasedisplayname ListDatabaseInsightsFieldsEnum = "databaseDisplayName"
	ListDatabaseInsightsFieldsDatabasetype        ListDatabaseInsightsFieldsEnum = "databaseType"
	ListDatabaseInsightsFieldsDatabaseversion     ListDatabaseInsightsFieldsEnum = "databaseVersion"
	ListDatabaseInsightsFieldsDatabasehostnames   ListDatabaseInsightsFieldsEnum = "databaseHostNames"
	ListDatabaseInsightsFieldsFreeformtags        ListDatabaseInsightsFieldsEnum = "freeformTags"
	ListDatabaseInsightsFieldsDefinedtags         ListDatabaseInsightsFieldsEnum = "definedTags"
)

var mappingListDatabaseInsightsFields = map[string]ListDatabaseInsightsFieldsEnum{
	"compartmentId":       ListDatabaseInsightsFieldsCompartmentid,
	"databaseName":        ListDatabaseInsightsFieldsDatabasename,
	"databaseDisplayName": ListDatabaseInsightsFieldsDatabasedisplayname,
	"databaseType":        ListDatabaseInsightsFieldsDatabasetype,
	"databaseVersion":     ListDatabaseInsightsFieldsDatabaseversion,
	"databaseHostNames":   ListDatabaseInsightsFieldsDatabasehostnames,
	"freeformTags":        ListDatabaseInsightsFieldsFreeformtags,
	"definedTags":         ListDatabaseInsightsFieldsDefinedtags,
}

// GetListDatabaseInsightsFieldsEnumValues Enumerates the set of values for ListDatabaseInsightsFieldsEnum
func GetListDatabaseInsightsFieldsEnumValues() []ListDatabaseInsightsFieldsEnum {
	values := make([]ListDatabaseInsightsFieldsEnum, 0)
	for _, v := range mappingListDatabaseInsightsFields {
		values = append(values, v)
	}
	return values
}

// ListDatabaseInsightsSortOrderEnum Enum with underlying type: string
type ListDatabaseInsightsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseInsightsSortOrderEnum
const (
	ListDatabaseInsightsSortOrderAsc  ListDatabaseInsightsSortOrderEnum = "ASC"
	ListDatabaseInsightsSortOrderDesc ListDatabaseInsightsSortOrderEnum = "DESC"
)

var mappingListDatabaseInsightsSortOrder = map[string]ListDatabaseInsightsSortOrderEnum{
	"ASC":  ListDatabaseInsightsSortOrderAsc,
	"DESC": ListDatabaseInsightsSortOrderDesc,
}

// GetListDatabaseInsightsSortOrderEnumValues Enumerates the set of values for ListDatabaseInsightsSortOrderEnum
func GetListDatabaseInsightsSortOrderEnumValues() []ListDatabaseInsightsSortOrderEnum {
	values := make([]ListDatabaseInsightsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseInsightsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDatabaseInsightsSortByEnum Enum with underlying type: string
type ListDatabaseInsightsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseInsightsSortByEnum
const (
	ListDatabaseInsightsSortByDatabasename        ListDatabaseInsightsSortByEnum = "databaseName"
	ListDatabaseInsightsSortByDatabasedisplayname ListDatabaseInsightsSortByEnum = "databaseDisplayName"
	ListDatabaseInsightsSortByDatabasetype        ListDatabaseInsightsSortByEnum = "databaseType"
)

var mappingListDatabaseInsightsSortBy = map[string]ListDatabaseInsightsSortByEnum{
	"databaseName":        ListDatabaseInsightsSortByDatabasename,
	"databaseDisplayName": ListDatabaseInsightsSortByDatabasedisplayname,
	"databaseType":        ListDatabaseInsightsSortByDatabasetype,
}

// GetListDatabaseInsightsSortByEnumValues Enumerates the set of values for ListDatabaseInsightsSortByEnum
func GetListDatabaseInsightsSortByEnumValues() []ListDatabaseInsightsSortByEnum {
	values := make([]ListDatabaseInsightsSortByEnum, 0)
	for _, v := range mappingListDatabaseInsightsSortBy {
		values = append(values, v)
	}
	return values
}
