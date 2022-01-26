// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListZoneTransferServersRequest wrapper for the ListZoneTransferServers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListZoneTransferServers.go.html to see an example of how to use ListZoneTransferServersRequest.
type ListZoneTransferServersRequest struct {

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListZoneTransferServersScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZoneTransferServersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZoneTransferServersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZoneTransferServersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZoneTransferServersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListZoneTransferServersResponse wrapper for the ListZoneTransferServers operation
type ListZoneTransferServersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ZoneTransferServer instances
	Items []ZoneTransferServer `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListZoneTransferServersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZoneTransferServersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZoneTransferServersScopeEnum Enum with underlying type: string
type ListZoneTransferServersScopeEnum string

// Set of constants representing the allowable values for ListZoneTransferServersScopeEnum
const (
	ListZoneTransferServersScopeGlobal  ListZoneTransferServersScopeEnum = "GLOBAL"
	ListZoneTransferServersScopePrivate ListZoneTransferServersScopeEnum = "PRIVATE"
)

var mappingListZoneTransferServersScope = map[string]ListZoneTransferServersScopeEnum{
	"GLOBAL":  ListZoneTransferServersScopeGlobal,
	"PRIVATE": ListZoneTransferServersScopePrivate,
}

// GetListZoneTransferServersScopeEnumValues Enumerates the set of values for ListZoneTransferServersScopeEnum
func GetListZoneTransferServersScopeEnumValues() []ListZoneTransferServersScopeEnum {
	values := make([]ListZoneTransferServersScopeEnum, 0)
	for _, v := range mappingListZoneTransferServersScope {
		values = append(values, v)
	}
	return values
}
