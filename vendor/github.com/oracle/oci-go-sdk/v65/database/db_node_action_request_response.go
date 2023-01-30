// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// DbNodeActionRequest wrapper for the DbNodeAction operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DbNodeAction.go.html to see an example of how to use DbNodeActionRequest.
type DbNodeActionRequest struct {

	// The database node OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbNodeId *string `mandatory:"true" contributesTo:"path" name:"dbNodeId"`

	// The action to perform on the DB Node.
	Action DbNodeActionActionEnum `mandatory:"true" contributesTo:"query" name:"action" omitEmpty:"true"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DbNodeActionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DbNodeActionRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DbNodeActionRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DbNodeActionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DbNodeActionRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbNodeActionActionEnum(string(request.Action)); !ok && request.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", request.Action, strings.Join(GetDbNodeActionActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbNodeActionResponse wrapper for the DbNodeAction operation
type DbNodeActionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DbNode instance
	DbNode `presentIn:"body"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request. Multiple OCID values are returned in a comma-separated list. Use GetWorkRequest with a work request OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DbNodeActionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DbNodeActionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DbNodeActionActionEnum Enum with underlying type: string
type DbNodeActionActionEnum string

// Set of constants representing the allowable values for DbNodeActionActionEnum
const (
	DbNodeActionActionStop      DbNodeActionActionEnum = "STOP"
	DbNodeActionActionStart     DbNodeActionActionEnum = "START"
	DbNodeActionActionSoftreset DbNodeActionActionEnum = "SOFTRESET"
	DbNodeActionActionReset     DbNodeActionActionEnum = "RESET"
)

var mappingDbNodeActionActionEnum = map[string]DbNodeActionActionEnum{
	"STOP":      DbNodeActionActionStop,
	"START":     DbNodeActionActionStart,
	"SOFTRESET": DbNodeActionActionSoftreset,
	"RESET":     DbNodeActionActionReset,
}

var mappingDbNodeActionActionEnumLowerCase = map[string]DbNodeActionActionEnum{
	"stop":      DbNodeActionActionStop,
	"start":     DbNodeActionActionStart,
	"softreset": DbNodeActionActionSoftreset,
	"reset":     DbNodeActionActionReset,
}

// GetDbNodeActionActionEnumValues Enumerates the set of values for DbNodeActionActionEnum
func GetDbNodeActionActionEnumValues() []DbNodeActionActionEnum {
	values := make([]DbNodeActionActionEnum, 0)
	for _, v := range mappingDbNodeActionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDbNodeActionActionEnumStringValues Enumerates the set of values in String for DbNodeActionActionEnum
func GetDbNodeActionActionEnumStringValues() []string {
	return []string{
		"STOP",
		"START",
		"SOFTRESET",
		"RESET",
	}
}

// GetMappingDbNodeActionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbNodeActionActionEnum(val string) (DbNodeActionActionEnum, bool) {
	enum, ok := mappingDbNodeActionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
