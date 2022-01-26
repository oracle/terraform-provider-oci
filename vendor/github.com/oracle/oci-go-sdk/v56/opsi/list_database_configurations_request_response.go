// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDatabaseConfigurationsRequest wrapper for the ListDatabaseConfigurations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseConfigurations.go.html to see an example of how to use ListDatabaseConfigurationsRequest.
type ListDatabaseConfigurationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique Enterprise Manager bridge identifier
	EnterpriseManagerBridgeId *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerBridgeId"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []ListDatabaseConfigurationsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

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
	SortOrder ListDatabaseConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Database configuration list sort options. If `fields` parameter is selected, the `sortBy` parameter must be one of the fields specified.
	SortBy ListDatabaseConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDatabaseConfigurationsResponse wrapper for the ListDatabaseConfigurations operation
type ListDatabaseConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseConfigurationCollection instances
	DatabaseConfigurationCollection `presentIn:"body"`

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

func (response ListDatabaseConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseConfigurationsDatabaseTypeEnum Enum with underlying type: string
type ListDatabaseConfigurationsDatabaseTypeEnum string

// Set of constants representing the allowable values for ListDatabaseConfigurationsDatabaseTypeEnum
const (
	ListDatabaseConfigurationsDatabaseTypeAdwS           ListDatabaseConfigurationsDatabaseTypeEnum = "ADW-S"
	ListDatabaseConfigurationsDatabaseTypeAtpS           ListDatabaseConfigurationsDatabaseTypeEnum = "ATP-S"
	ListDatabaseConfigurationsDatabaseTypeAdwD           ListDatabaseConfigurationsDatabaseTypeEnum = "ADW-D"
	ListDatabaseConfigurationsDatabaseTypeAtpD           ListDatabaseConfigurationsDatabaseTypeEnum = "ATP-D"
	ListDatabaseConfigurationsDatabaseTypeExternalPdb    ListDatabaseConfigurationsDatabaseTypeEnum = "EXTERNAL-PDB"
	ListDatabaseConfigurationsDatabaseTypeExternalNoncdb ListDatabaseConfigurationsDatabaseTypeEnum = "EXTERNAL-NONCDB"
)

var mappingListDatabaseConfigurationsDatabaseType = map[string]ListDatabaseConfigurationsDatabaseTypeEnum{
	"ADW-S":           ListDatabaseConfigurationsDatabaseTypeAdwS,
	"ATP-S":           ListDatabaseConfigurationsDatabaseTypeAtpS,
	"ADW-D":           ListDatabaseConfigurationsDatabaseTypeAdwD,
	"ATP-D":           ListDatabaseConfigurationsDatabaseTypeAtpD,
	"EXTERNAL-PDB":    ListDatabaseConfigurationsDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB": ListDatabaseConfigurationsDatabaseTypeExternalNoncdb,
}

// GetListDatabaseConfigurationsDatabaseTypeEnumValues Enumerates the set of values for ListDatabaseConfigurationsDatabaseTypeEnum
func GetListDatabaseConfigurationsDatabaseTypeEnumValues() []ListDatabaseConfigurationsDatabaseTypeEnum {
	values := make([]ListDatabaseConfigurationsDatabaseTypeEnum, 0)
	for _, v := range mappingListDatabaseConfigurationsDatabaseType {
		values = append(values, v)
	}
	return values
}

// ListDatabaseConfigurationsSortOrderEnum Enum with underlying type: string
type ListDatabaseConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseConfigurationsSortOrderEnum
const (
	ListDatabaseConfigurationsSortOrderAsc  ListDatabaseConfigurationsSortOrderEnum = "ASC"
	ListDatabaseConfigurationsSortOrderDesc ListDatabaseConfigurationsSortOrderEnum = "DESC"
)

var mappingListDatabaseConfigurationsSortOrder = map[string]ListDatabaseConfigurationsSortOrderEnum{
	"ASC":  ListDatabaseConfigurationsSortOrderAsc,
	"DESC": ListDatabaseConfigurationsSortOrderDesc,
}

// GetListDatabaseConfigurationsSortOrderEnumValues Enumerates the set of values for ListDatabaseConfigurationsSortOrderEnum
func GetListDatabaseConfigurationsSortOrderEnumValues() []ListDatabaseConfigurationsSortOrderEnum {
	values := make([]ListDatabaseConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseConfigurationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDatabaseConfigurationsSortByEnum Enum with underlying type: string
type ListDatabaseConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseConfigurationsSortByEnum
const (
	ListDatabaseConfigurationsSortByDatabasename        ListDatabaseConfigurationsSortByEnum = "databaseName"
	ListDatabaseConfigurationsSortByDatabasedisplayname ListDatabaseConfigurationsSortByEnum = "databaseDisplayName"
	ListDatabaseConfigurationsSortByDatabasetype        ListDatabaseConfigurationsSortByEnum = "databaseType"
)

var mappingListDatabaseConfigurationsSortBy = map[string]ListDatabaseConfigurationsSortByEnum{
	"databaseName":        ListDatabaseConfigurationsSortByDatabasename,
	"databaseDisplayName": ListDatabaseConfigurationsSortByDatabasedisplayname,
	"databaseType":        ListDatabaseConfigurationsSortByDatabasetype,
}

// GetListDatabaseConfigurationsSortByEnumValues Enumerates the set of values for ListDatabaseConfigurationsSortByEnum
func GetListDatabaseConfigurationsSortByEnumValues() []ListDatabaseConfigurationsSortByEnum {
	values := make([]ListDatabaseConfigurationsSortByEnum, 0)
	for _, v := range mappingListDatabaseConfigurationsSortBy {
		values = append(values, v)
	}
	return values
}
