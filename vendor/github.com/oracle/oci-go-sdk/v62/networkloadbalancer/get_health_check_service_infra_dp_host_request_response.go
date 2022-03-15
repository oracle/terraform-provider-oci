// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"net/http"
	"strings"
)

// GetHealthCheckServiceInfraDpHostRequest wrapper for the GetHealthCheckServiceInfraDpHost operation
type GetHealthCheckServiceInfraDpHostRequest struct {

	// The AvailabilityDomain of the hcs dp host.
	AvailabilityDomain *string `mandatory:"true" contributesTo:"path" name:"availabilityDomain"`

	// The ID of the hcs dp host.
	DpHostId *string `mandatory:"true" contributesTo:"path" name:"dpHostId"`

	// The system returns the requested resource, with a 200 status, only if the resource has no etag
	// matching the one specified. If the condition fails for the GET and HEAD methods, then the system returns the
	// HTTP status code `304 (Not Modified)`.
	// Example: `example-etag`
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetHealthCheckServiceInfraDpHostRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetHealthCheckServiceInfraDpHostRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetHealthCheckServiceInfraDpHostRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetHealthCheckServiceInfraDpHostRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetHealthCheckServiceInfraDpHostRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetHealthCheckServiceInfraDpHostResponse wrapper for the GetHealthCheckServiceInfraDpHost operation
type GetHealthCheckServiceInfraDpHostResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DpHost instance
	DpHost `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetHealthCheckServiceInfraDpHostResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetHealthCheckServiceInfraDpHostResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
