// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListExadataConfigurationsRequest wrapper for the ListExadataConfigurations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListExadataConfigurations.go.html to see an example of how to use ListExadataConfigurationsRequest.
type ListExadataConfigurationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

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
	SortOrder ListExadataConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Exadata configuration list sort options. If `fields` parameter is selected, the `sortBy` parameter must be one of the fields specified.
	SortBy ListExadataConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExadataConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExadataConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExadataConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExadataConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListExadataConfigurationsResponse wrapper for the ListExadataConfigurations operation
type ListExadataConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExadataConfigurationCollection instances
	ExadataConfigurationCollection `presentIn:"body"`

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

func (response ListExadataConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExadataConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExadataConfigurationsSortOrderEnum Enum with underlying type: string
type ListExadataConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListExadataConfigurationsSortOrderEnum
const (
	ListExadataConfigurationsSortOrderAsc  ListExadataConfigurationsSortOrderEnum = "ASC"
	ListExadataConfigurationsSortOrderDesc ListExadataConfigurationsSortOrderEnum = "DESC"
)

var mappingListExadataConfigurationsSortOrder = map[string]ListExadataConfigurationsSortOrderEnum{
	"ASC":  ListExadataConfigurationsSortOrderAsc,
	"DESC": ListExadataConfigurationsSortOrderDesc,
}

// GetListExadataConfigurationsSortOrderEnumValues Enumerates the set of values for ListExadataConfigurationsSortOrderEnum
func GetListExadataConfigurationsSortOrderEnumValues() []ListExadataConfigurationsSortOrderEnum {
	values := make([]ListExadataConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListExadataConfigurationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListExadataConfigurationsSortByEnum Enum with underlying type: string
type ListExadataConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListExadataConfigurationsSortByEnum
const (
	ListExadataConfigurationsSortByExadataname        ListExadataConfigurationsSortByEnum = "exadataName"
	ListExadataConfigurationsSortByExadatadisplayname ListExadataConfigurationsSortByEnum = "exadataDisplayName"
	ListExadataConfigurationsSortByExadatatype        ListExadataConfigurationsSortByEnum = "exadataType"
)

var mappingListExadataConfigurationsSortBy = map[string]ListExadataConfigurationsSortByEnum{
	"exadataName":        ListExadataConfigurationsSortByExadataname,
	"exadataDisplayName": ListExadataConfigurationsSortByExadatadisplayname,
	"exadataType":        ListExadataConfigurationsSortByExadatatype,
}

// GetListExadataConfigurationsSortByEnumValues Enumerates the set of values for ListExadataConfigurationsSortByEnum
func GetListExadataConfigurationsSortByEnumValues() []ListExadataConfigurationsSortByEnum {
	values := make([]ListExadataConfigurationsSortByEnum, 0)
	for _, v := range mappingListExadataConfigurationsSortBy {
		values = append(values, v)
	}
	return values
}
