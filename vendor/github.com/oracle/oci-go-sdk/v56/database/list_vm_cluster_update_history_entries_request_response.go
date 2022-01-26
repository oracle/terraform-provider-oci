// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListVmClusterUpdateHistoryEntriesRequest wrapper for the ListVmClusterUpdateHistoryEntries operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListVmClusterUpdateHistoryEntries.go.html to see an example of how to use ListVmClusterUpdateHistoryEntriesRequest.
type ListVmClusterUpdateHistoryEntriesRequest struct {

	// The VM cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	VmClusterId *string `mandatory:"true" contributesTo:"path" name:"vmClusterId"`

	// A filter to return only resources that match the given update type exactly.
	UpdateType ListVmClusterUpdateHistoryEntriesUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVmClusterUpdateHistoryEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVmClusterUpdateHistoryEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVmClusterUpdateHistoryEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVmClusterUpdateHistoryEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListVmClusterUpdateHistoryEntriesResponse wrapper for the ListVmClusterUpdateHistoryEntries operation
type ListVmClusterUpdateHistoryEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []VmClusterUpdateHistoryEntrySummary instances
	Items []VmClusterUpdateHistoryEntrySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVmClusterUpdateHistoryEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVmClusterUpdateHistoryEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVmClusterUpdateHistoryEntriesUpdateTypeEnum Enum with underlying type: string
type ListVmClusterUpdateHistoryEntriesUpdateTypeEnum string

// Set of constants representing the allowable values for ListVmClusterUpdateHistoryEntriesUpdateTypeEnum
const (
	ListVmClusterUpdateHistoryEntriesUpdateTypeGiUpgrade ListVmClusterUpdateHistoryEntriesUpdateTypeEnum = "GI_UPGRADE"
	ListVmClusterUpdateHistoryEntriesUpdateTypeGiPatch   ListVmClusterUpdateHistoryEntriesUpdateTypeEnum = "GI_PATCH"
	ListVmClusterUpdateHistoryEntriesUpdateTypeOsUpdate  ListVmClusterUpdateHistoryEntriesUpdateTypeEnum = "OS_UPDATE"
)

var mappingListVmClusterUpdateHistoryEntriesUpdateType = map[string]ListVmClusterUpdateHistoryEntriesUpdateTypeEnum{
	"GI_UPGRADE": ListVmClusterUpdateHistoryEntriesUpdateTypeGiUpgrade,
	"GI_PATCH":   ListVmClusterUpdateHistoryEntriesUpdateTypeGiPatch,
	"OS_UPDATE":  ListVmClusterUpdateHistoryEntriesUpdateTypeOsUpdate,
}

// GetListVmClusterUpdateHistoryEntriesUpdateTypeEnumValues Enumerates the set of values for ListVmClusterUpdateHistoryEntriesUpdateTypeEnum
func GetListVmClusterUpdateHistoryEntriesUpdateTypeEnumValues() []ListVmClusterUpdateHistoryEntriesUpdateTypeEnum {
	values := make([]ListVmClusterUpdateHistoryEntriesUpdateTypeEnum, 0)
	for _, v := range mappingListVmClusterUpdateHistoryEntriesUpdateType {
		values = append(values, v)
	}
	return values
}
