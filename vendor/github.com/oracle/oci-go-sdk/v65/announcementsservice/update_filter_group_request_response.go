// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package announcementsservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// UpdateFilterGroupRequest wrapper for the UpdateFilterGroup operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/announcementsservice/UpdateFilterGroup.go.html to see an example of how to use UpdateFilterGroupRequest.
type UpdateFilterGroupRequest struct {

	// The OCID of the announcement subscription.
	AnnouncementSubscriptionId *string `mandatory:"true" contributesTo:"path" name:"announcementSubscriptionId"`

	// The name of the filter group.
	FilterGroupName *string `mandatory:"true" contributesTo:"path" name:"filterGroupName"`

	// The filter group information to update.
	UpdateFilterGroupDetails `contributesTo:"body"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The locking version, used for optimistic concurrency control.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateFilterGroupRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateFilterGroupRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateFilterGroupRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateFilterGroupRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateFilterGroupRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateFilterGroupResponse wrapper for the UpdateFilterGroup operation
type UpdateFilterGroupResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The FilterGroup instance
	FilterGroup `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response UpdateFilterGroupResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateFilterGroupResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
