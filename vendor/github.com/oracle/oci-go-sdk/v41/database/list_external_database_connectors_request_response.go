// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// ListExternalDatabaseConnectorsRequest wrapper for the ListExternalDatabaseConnectors operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalDatabaseConnectors.go.html to see an example of how to use ListExternalDatabaseConnectorsRequest.
type ListExternalDatabaseConnectorsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external database whose connectors will be listed.
	ExternalDatabaseId *string `mandatory:"true" contributesTo:"query" name:"externalDatabaseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListExternalDatabaseConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExternalDatabaseConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState ExternalDatabaseConnectorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalDatabaseConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalDatabaseConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalDatabaseConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalDatabaseConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListExternalDatabaseConnectorsResponse wrapper for the ListExternalDatabaseConnectors operation
type ListExternalDatabaseConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExternalDatabaseConnectorSummary instances
	Items []ExternalDatabaseConnectorSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalDatabaseConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalDatabaseConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalDatabaseConnectorsSortByEnum Enum with underlying type: string
type ListExternalDatabaseConnectorsSortByEnum string

// Set of constants representing the allowable values for ListExternalDatabaseConnectorsSortByEnum
const (
	ListExternalDatabaseConnectorsSortByDisplayname ListExternalDatabaseConnectorsSortByEnum = "DISPLAYNAME"
	ListExternalDatabaseConnectorsSortByTimecreated ListExternalDatabaseConnectorsSortByEnum = "TIMECREATED"
)

var mappingListExternalDatabaseConnectorsSortBy = map[string]ListExternalDatabaseConnectorsSortByEnum{
	"DISPLAYNAME": ListExternalDatabaseConnectorsSortByDisplayname,
	"TIMECREATED": ListExternalDatabaseConnectorsSortByTimecreated,
}

// GetListExternalDatabaseConnectorsSortByEnumValues Enumerates the set of values for ListExternalDatabaseConnectorsSortByEnum
func GetListExternalDatabaseConnectorsSortByEnumValues() []ListExternalDatabaseConnectorsSortByEnum {
	values := make([]ListExternalDatabaseConnectorsSortByEnum, 0)
	for _, v := range mappingListExternalDatabaseConnectorsSortBy {
		values = append(values, v)
	}
	return values
}

// ListExternalDatabaseConnectorsSortOrderEnum Enum with underlying type: string
type ListExternalDatabaseConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListExternalDatabaseConnectorsSortOrderEnum
const (
	ListExternalDatabaseConnectorsSortOrderAsc  ListExternalDatabaseConnectorsSortOrderEnum = "ASC"
	ListExternalDatabaseConnectorsSortOrderDesc ListExternalDatabaseConnectorsSortOrderEnum = "DESC"
)

var mappingListExternalDatabaseConnectorsSortOrder = map[string]ListExternalDatabaseConnectorsSortOrderEnum{
	"ASC":  ListExternalDatabaseConnectorsSortOrderAsc,
	"DESC": ListExternalDatabaseConnectorsSortOrderDesc,
}

// GetListExternalDatabaseConnectorsSortOrderEnumValues Enumerates the set of values for ListExternalDatabaseConnectorsSortOrderEnum
func GetListExternalDatabaseConnectorsSortOrderEnumValues() []ListExternalDatabaseConnectorsSortOrderEnum {
	values := make([]ListExternalDatabaseConnectorsSortOrderEnum, 0)
	for _, v := range mappingListExternalDatabaseConnectorsSortOrder {
		values = append(values, v)
	}
	return values
}
