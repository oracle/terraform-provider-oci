// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListStackResourceDriftDetailsRequest wrapper for the ListStackResourceDriftDetails operation
type ListStackResourceDriftDetailsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack.
	StackId *string `mandatory:"true" contributesTo:"path" name:"stackId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only resources that match the given drift status. The value is case-insensitive.
	// Allowable values -
	//   - NOT_CHECKED
	//   - MODIFIED
	//   - IN_SYNC
	//   - DELETED
	ResourceDriftStatus []StackResourceDriftSummaryResourceDriftStatusEnum `contributesTo:"query" name:"resourceDriftStatus" omitEmpty:"true" collectionFormat:"multi"`

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

func (request ListStackResourceDriftDetailsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListStackResourceDriftDetailsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListStackResourceDriftDetailsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListStackResourceDriftDetailsResponse wrapper for the ListStackResourceDriftDetails operation
type ListStackResourceDriftDetailsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of StackResourceDriftCollection instances
	StackResourceDriftCollection `presentIn:"body"`

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

func (response ListStackResourceDriftDetailsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListStackResourceDriftDetailsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
