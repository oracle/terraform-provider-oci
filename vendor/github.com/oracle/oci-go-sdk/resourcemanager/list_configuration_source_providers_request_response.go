// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConfigurationSourceProvidersRequest wrapper for the ListConfigurationSourceProviders operation
type ListConfigurationSourceProvidersRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that exist in the compartment, identified by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only configuration source providers that match the provided OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ConfigurationSourceProviderId *string `mandatory:"false" contributesTo:"query" name:"configurationSourceProviderId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to use when sorting returned resources.
	// By default, `TIMECREATED` is ordered descending.
	// By default, `DISPLAYNAME` is ordered ascending. Note that you can sort only on one field.
	SortBy ListConfigurationSourceProvidersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use when sorting returned resources. Ascending (`ASC`) or descending (`DESC`).
	SortOrder ListConfigurationSourceProvidersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The number of items returned in a paginated `List` call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the preceding `List` call.
	// For information about pagination, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConfigurationSourceProvidersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConfigurationSourceProvidersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConfigurationSourceProvidersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConfigurationSourceProvidersResponse wrapper for the ListConfigurationSourceProviders operation
type ListConfigurationSourceProvidersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConfigurationSourceProviderCollection instances
	ConfigurationSourceProviderCollection `presentIn:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of paginated list items. If the `opc-next-page`
	// header appears in the response, additional pages of results remain.
	// To receive the next page, include the header value in the `page` param.
	// If the `opc-next-page` header does not appear in the response, there
	// are no more list items to get. For more information about list pagination,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConfigurationSourceProvidersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConfigurationSourceProvidersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConfigurationSourceProvidersSortByEnum Enum with underlying type: string
type ListConfigurationSourceProvidersSortByEnum string

// Set of constants representing the allowable values for ListConfigurationSourceProvidersSortByEnum
const (
	ListConfigurationSourceProvidersSortByTimecreated ListConfigurationSourceProvidersSortByEnum = "TIMECREATED"
	ListConfigurationSourceProvidersSortByDisplayname ListConfigurationSourceProvidersSortByEnum = "DISPLAYNAME"
)

var mappingListConfigurationSourceProvidersSortBy = map[string]ListConfigurationSourceProvidersSortByEnum{
	"TIMECREATED": ListConfigurationSourceProvidersSortByTimecreated,
	"DISPLAYNAME": ListConfigurationSourceProvidersSortByDisplayname,
}

// GetListConfigurationSourceProvidersSortByEnumValues Enumerates the set of values for ListConfigurationSourceProvidersSortByEnum
func GetListConfigurationSourceProvidersSortByEnumValues() []ListConfigurationSourceProvidersSortByEnum {
	values := make([]ListConfigurationSourceProvidersSortByEnum, 0)
	for _, v := range mappingListConfigurationSourceProvidersSortBy {
		values = append(values, v)
	}
	return values
}

// ListConfigurationSourceProvidersSortOrderEnum Enum with underlying type: string
type ListConfigurationSourceProvidersSortOrderEnum string

// Set of constants representing the allowable values for ListConfigurationSourceProvidersSortOrderEnum
const (
	ListConfigurationSourceProvidersSortOrderAsc  ListConfigurationSourceProvidersSortOrderEnum = "ASC"
	ListConfigurationSourceProvidersSortOrderDesc ListConfigurationSourceProvidersSortOrderEnum = "DESC"
)

var mappingListConfigurationSourceProvidersSortOrder = map[string]ListConfigurationSourceProvidersSortOrderEnum{
	"ASC":  ListConfigurationSourceProvidersSortOrderAsc,
	"DESC": ListConfigurationSourceProvidersSortOrderDesc,
}

// GetListConfigurationSourceProvidersSortOrderEnumValues Enumerates the set of values for ListConfigurationSourceProvidersSortOrderEnum
func GetListConfigurationSourceProvidersSortOrderEnumValues() []ListConfigurationSourceProvidersSortOrderEnum {
	values := make([]ListConfigurationSourceProvidersSortOrderEnum, 0)
	for _, v := range mappingListConfigurationSourceProvidersSortOrder {
		values = append(values, v)
	}
	return values
}
