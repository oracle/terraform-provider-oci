// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// RetrievePasswordRequest wrapper for the RetrievePassword operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/RetrievePassword.go.html to see an example of how to use RetrievePasswordRequest.
type RetrievePasswordRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC.
	SddcId *string `mandatory:"true" contributesTo:"path" name:"sddcId"`

	// The SDDC password type.
	Type RetrievePasswordTypeEnum `mandatory:"true" contributesTo:"query" name:"type" omitEmpty:"true"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RetrievePasswordRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RetrievePasswordRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RetrievePasswordRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RetrievePasswordRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RetrievePasswordRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRetrievePasswordTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetRetrievePasswordTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RetrievePasswordResponse wrapper for the RetrievePassword operation
type RetrievePasswordResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SddcPassword instance
	SddcPassword `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response RetrievePasswordResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RetrievePasswordResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RetrievePasswordTypeEnum Enum with underlying type: string
type RetrievePasswordTypeEnum string

// Set of constants representing the allowable values for RetrievePasswordTypeEnum
const (
	RetrievePasswordTypeVcenter RetrievePasswordTypeEnum = "VCENTER"
	RetrievePasswordTypeNsx     RetrievePasswordTypeEnum = "NSX"
	RetrievePasswordTypeHcx     RetrievePasswordTypeEnum = "HCX"
)

var mappingRetrievePasswordTypeEnum = map[string]RetrievePasswordTypeEnum{
	"VCENTER": RetrievePasswordTypeVcenter,
	"NSX":     RetrievePasswordTypeNsx,
	"HCX":     RetrievePasswordTypeHcx,
}

var mappingRetrievePasswordTypeEnumLowerCase = map[string]RetrievePasswordTypeEnum{
	"vcenter": RetrievePasswordTypeVcenter,
	"nsx":     RetrievePasswordTypeNsx,
	"hcx":     RetrievePasswordTypeHcx,
}

// GetRetrievePasswordTypeEnumValues Enumerates the set of values for RetrievePasswordTypeEnum
func GetRetrievePasswordTypeEnumValues() []RetrievePasswordTypeEnum {
	values := make([]RetrievePasswordTypeEnum, 0)
	for _, v := range mappingRetrievePasswordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRetrievePasswordTypeEnumStringValues Enumerates the set of values in String for RetrievePasswordTypeEnum
func GetRetrievePasswordTypeEnumStringValues() []string {
	return []string{
		"VCENTER",
		"NSX",
		"HCX",
	}
}

// GetMappingRetrievePasswordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRetrievePasswordTypeEnum(val string) (RetrievePasswordTypeEnum, bool) {
	enum, ok := mappingRetrievePasswordTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
