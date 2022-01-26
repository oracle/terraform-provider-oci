// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waf

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListProtectionCapabilityGroupTagsRequest wrapper for the ListProtectionCapabilityGroupTags operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListProtectionCapabilityGroupTags.go.html to see an example of how to use ListProtectionCapabilityGroupTagsRequest.
type ListProtectionCapabilityGroupTagsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to return only resources that matches given type.
	Type ProtectionCapabilitySummaryTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListProtectionCapabilityGroupTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for name is ascending.
	// If no value is specified name is default.
	SortBy ListProtectionCapabilityGroupTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProtectionCapabilityGroupTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProtectionCapabilityGroupTagsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProtectionCapabilityGroupTagsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProtectionCapabilityGroupTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListProtectionCapabilityGroupTagsResponse wrapper for the ListProtectionCapabilityGroupTags operation
type ListProtectionCapabilityGroupTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProtectionCapabilityGroupTagCollection instances
	ProtectionCapabilityGroupTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProtectionCapabilityGroupTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProtectionCapabilityGroupTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProtectionCapabilityGroupTagsSortOrderEnum Enum with underlying type: string
type ListProtectionCapabilityGroupTagsSortOrderEnum string

// Set of constants representing the allowable values for ListProtectionCapabilityGroupTagsSortOrderEnum
const (
	ListProtectionCapabilityGroupTagsSortOrderAsc  ListProtectionCapabilityGroupTagsSortOrderEnum = "ASC"
	ListProtectionCapabilityGroupTagsSortOrderDesc ListProtectionCapabilityGroupTagsSortOrderEnum = "DESC"
)

var mappingListProtectionCapabilityGroupTagsSortOrder = map[string]ListProtectionCapabilityGroupTagsSortOrderEnum{
	"ASC":  ListProtectionCapabilityGroupTagsSortOrderAsc,
	"DESC": ListProtectionCapabilityGroupTagsSortOrderDesc,
}

// GetListProtectionCapabilityGroupTagsSortOrderEnumValues Enumerates the set of values for ListProtectionCapabilityGroupTagsSortOrderEnum
func GetListProtectionCapabilityGroupTagsSortOrderEnumValues() []ListProtectionCapabilityGroupTagsSortOrderEnum {
	values := make([]ListProtectionCapabilityGroupTagsSortOrderEnum, 0)
	for _, v := range mappingListProtectionCapabilityGroupTagsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListProtectionCapabilityGroupTagsSortByEnum Enum with underlying type: string
type ListProtectionCapabilityGroupTagsSortByEnum string

// Set of constants representing the allowable values for ListProtectionCapabilityGroupTagsSortByEnum
const (
	ListProtectionCapabilityGroupTagsSortByName ListProtectionCapabilityGroupTagsSortByEnum = "name"
)

var mappingListProtectionCapabilityGroupTagsSortBy = map[string]ListProtectionCapabilityGroupTagsSortByEnum{
	"name": ListProtectionCapabilityGroupTagsSortByName,
}

// GetListProtectionCapabilityGroupTagsSortByEnumValues Enumerates the set of values for ListProtectionCapabilityGroupTagsSortByEnum
func GetListProtectionCapabilityGroupTagsSortByEnumValues() []ListProtectionCapabilityGroupTagsSortByEnum {
	values := make([]ListProtectionCapabilityGroupTagsSortByEnum, 0)
	for _, v := range mappingListProtectionCapabilityGroupTagsSortBy {
		values = append(values, v)
	}
	return values
}
