// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListBulkActionResourceTypesRequest wrapper for the ListBulkActionResourceTypes operation
type ListBulkActionResourceTypesRequest struct {

	// The type of bulk action.
	BulkActionType ListBulkActionResourceTypesBulkActionTypeEnum `mandatory:"true" contributesTo:"query" name:"bulkActionType" omitEmpty:"true"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBulkActionResourceTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBulkActionResourceTypesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBulkActionResourceTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBulkActionResourceTypesResponse wrapper for the ListBulkActionResourceTypes operation
type ListBulkActionResourceTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BulkActionResourceTypeCollection instances
	BulkActionResourceTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBulkActionResourceTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBulkActionResourceTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBulkActionResourceTypesBulkActionTypeEnum Enum with underlying type: string
type ListBulkActionResourceTypesBulkActionTypeEnum string

// Set of constants representing the allowable values for ListBulkActionResourceTypesBulkActionTypeEnum
const (
	ListBulkActionResourceTypesBulkActionTypeMoveResources   ListBulkActionResourceTypesBulkActionTypeEnum = "BULK_MOVE_RESOURCES"
	ListBulkActionResourceTypesBulkActionTypeDeleteResources ListBulkActionResourceTypesBulkActionTypeEnum = "BULK_DELETE_RESOURCES"
)

var mappingListBulkActionResourceTypesBulkActionType = map[string]ListBulkActionResourceTypesBulkActionTypeEnum{
	"BULK_MOVE_RESOURCES":   ListBulkActionResourceTypesBulkActionTypeMoveResources,
	"BULK_DELETE_RESOURCES": ListBulkActionResourceTypesBulkActionTypeDeleteResources,
}

// GetListBulkActionResourceTypesBulkActionTypeEnumValues Enumerates the set of values for ListBulkActionResourceTypesBulkActionTypeEnum
func GetListBulkActionResourceTypesBulkActionTypeEnumValues() []ListBulkActionResourceTypesBulkActionTypeEnum {
	values := make([]ListBulkActionResourceTypesBulkActionTypeEnum, 0)
	for _, v := range mappingListBulkActionResourceTypesBulkActionType {
		values = append(values, v)
	}
	return values
}
