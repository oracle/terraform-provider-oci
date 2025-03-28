// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package audit

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEventsRequest wrapper for the ListEvents operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/audit/ListEvents.go.html to see an example of how to use ListEventsRequest.
type ListEventsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Returns events that were processed at or after this start date and time, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// For example, a start value of `2017-01-15T11:30:00Z` will retrieve a list of all events processed
	// since 30 minutes after the 11th hour of January 15, 2017, in Coordinated Universal Time (UTC).
	// You can specify a value with granularity to the minute. Seconds (and milliseconds, if included) must
	// be set to `0`.
	StartTime *common.SDKTime `mandatory:"true" contributesTo:"query" name:"startTime"`

	// Returns events that were processed before this end date and time, expressed in
	// RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// For example, a start value of `2017-01-01T00:00:00Z` and an end value of `2017-01-02T00:00:00Z`
	// will retrieve a list of all events processed on January 1, 2017. Similarly, a start value of
	// `2017-01-01T00:00:00Z` and an end value of `2017-02-01T00:00:00Z` will result in a list of all
	// events processed between January 1, 2017 and January 31, 2017. You can specify a value with
	// granularity to the minute. Seconds (and milliseconds, if included) must be set to `0`.
	EndTime *common.SDKTime `mandatory:"true" contributesTo:"query" name:"endTime"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEventsResponse wrapper for the ListEvents operation
type ListEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AuditEvent instances
	Items []AuditEvent `presentIn:"body"`

	// For pagination of a list of audit events. When this header appears in the response,
	// it means you received a partial list and there are more results. Include this value as the `page`
	// parameter for the subsequent ListEvents request to get the next batch of events. For important
	// details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
