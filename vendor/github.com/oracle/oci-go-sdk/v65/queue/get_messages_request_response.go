// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetMessagesRequest wrapper for the GetMessages operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/queue/GetMessages.go.html to see an example of how to use GetMessagesRequest.
type GetMessagesRequest struct {

	// The unique queue identifier.
	QueueId *string `mandatory:"true" contributesTo:"path" name:"queueId"`

	// If the `visibilityInSeconds` parameter is set, messages will be hidden for `visibilityInSeconds` seconds and won't be consumable by other consumers during that time.
	// If it isn't set it defaults to the value set at the queue level.
	// Using a `visibilityInSeconds` value of 0 effectively acts as a peek functionality.
	// Messages retrieved that way aren't meant to be deleted because they will most likely be delivered to another consumer as their visibility won't change, but will still increase the delivery count by one.
	VisibilityInSeconds *int `mandatory:"false" contributesTo:"query" name:"visibilityInSeconds"`

	// If the `timeoutInSeconds parameter` isn't set or it is set to a value greater than 0, the request is using the long-polling mode and will only return when a message is available for consumption (it does not wait for limit messages but still only returns at-most limit messages) or after `timeoutInSeconds` seconds (in which case it will return an empty response), whichever comes first.
	// If the parameter is set to 0, the request is using the short-polling mode and immediately returns whether messages have been retrieved or not.
	// In same rare-cases a long-polling request could be interrupted (returned with empty response) before the end of the timeout.
	TimeoutInSeconds *int `mandatory:"false" contributesTo:"query" name:"timeoutInSeconds"`

	// The limit parameter controls how many messages is returned at-most.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Optional parameter to filter the channels.
	ChannelFilter *string `mandatory:"false" contributesTo:"query" name:"channelFilter"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetMessagesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetMessagesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetMessagesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetMessagesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetMessagesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetMessagesResponse wrapper for the GetMessages operation
type GetMessagesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The GetMessages instance
	GetMessages `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetMessagesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetMessagesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
