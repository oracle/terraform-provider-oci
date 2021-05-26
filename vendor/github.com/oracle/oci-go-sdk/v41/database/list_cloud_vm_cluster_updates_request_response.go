// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// ListCloudVmClusterUpdatesRequest wrapper for the ListCloudVmClusterUpdates operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudVmClusterUpdates.go.html to see an example of how to use ListCloudVmClusterUpdatesRequest.
type ListCloudVmClusterUpdatesRequest struct {

	// The cloud VM cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CloudVmClusterId *string `mandatory:"true" contributesTo:"path" name:"cloudVmClusterId"`

	// A filter to return only resources that match the given update type exactly.
	UpdateType ListCloudVmClusterUpdatesUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

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

func (request ListCloudVmClusterUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudVmClusterUpdatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudVmClusterUpdatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudVmClusterUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCloudVmClusterUpdatesResponse wrapper for the ListCloudVmClusterUpdates operation
type ListCloudVmClusterUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []UpdateSummary instances
	Items []UpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudVmClusterUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudVmClusterUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudVmClusterUpdatesUpdateTypeEnum Enum with underlying type: string
type ListCloudVmClusterUpdatesUpdateTypeEnum string

// Set of constants representing the allowable values for ListCloudVmClusterUpdatesUpdateTypeEnum
const (
	ListCloudVmClusterUpdatesUpdateTypeGiUpgrade ListCloudVmClusterUpdatesUpdateTypeEnum = "GI_UPGRADE"
	ListCloudVmClusterUpdatesUpdateTypeGiPatch   ListCloudVmClusterUpdatesUpdateTypeEnum = "GI_PATCH"
	ListCloudVmClusterUpdatesUpdateTypeOsUpdate  ListCloudVmClusterUpdatesUpdateTypeEnum = "OS_UPDATE"
)

var mappingListCloudVmClusterUpdatesUpdateType = map[string]ListCloudVmClusterUpdatesUpdateTypeEnum{
	"GI_UPGRADE": ListCloudVmClusterUpdatesUpdateTypeGiUpgrade,
	"GI_PATCH":   ListCloudVmClusterUpdatesUpdateTypeGiPatch,
	"OS_UPDATE":  ListCloudVmClusterUpdatesUpdateTypeOsUpdate,
}

// GetListCloudVmClusterUpdatesUpdateTypeEnumValues Enumerates the set of values for ListCloudVmClusterUpdatesUpdateTypeEnum
func GetListCloudVmClusterUpdatesUpdateTypeEnumValues() []ListCloudVmClusterUpdatesUpdateTypeEnum {
	values := make([]ListCloudVmClusterUpdatesUpdateTypeEnum, 0)
	for _, v := range mappingListCloudVmClusterUpdatesUpdateType {
		values = append(values, v)
	}
	return values
}
