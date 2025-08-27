// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// TerminateInstanceRequest wrapper for the TerminateInstance operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/TerminateInstance.go.html to see an example of how to use TerminateInstanceRequest.
type TerminateInstanceRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
	InstanceId *string `mandatory:"true" contributesTo:"path" name:"instanceId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource. The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Specifies whether to delete or preserve the boot volume when terminating an instance.
	// When set to `true`, the boot volume is preserved. The default value is `false`.
	PreserveBootVolume *bool `mandatory:"false" contributesTo:"query" name:"preserveBootVolume"`

	// Specifies whether to delete or preserve the data volumes created during launch when
	// terminating an instance. When set to `true`, the data volumes are preserved. The
	// default value is `true`.
	PreserveDataVolumesCreatedAtLaunch *bool `mandatory:"false" contributesTo:"query" name:"preserveDataVolumesCreatedAtLaunch"`

	// This optional parameter overrides recycle level for hosts. The parameter can be used when hosts are associated
	// with a Capacity Reservation.
	// * `FULL_RECYCLE` - Does not skip host wipe. This is the default behavior.
	RecycleLevel TerminateInstanceRecycleLevelEnum `mandatory:"false" contributesTo:"query" name:"recycleLevel" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request TerminateInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request TerminateInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request TerminateInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request TerminateInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request TerminateInstanceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTerminateInstanceRecycleLevelEnum(string(request.RecycleLevel)); !ok && request.RecycleLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecycleLevel: %s. Supported values are: %s.", request.RecycleLevel, strings.Join(GetTerminateInstanceRecycleLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TerminateInstanceResponse wrapper for the TerminateInstance operation
type TerminateInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response TerminateInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response TerminateInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// TerminateInstanceRecycleLevelEnum Enum with underlying type: string
type TerminateInstanceRecycleLevelEnum string

// Set of constants representing the allowable values for TerminateInstanceRecycleLevelEnum
const (
	TerminateInstanceRecycleLevelFullRecycle TerminateInstanceRecycleLevelEnum = "FULL_RECYCLE"
)

var mappingTerminateInstanceRecycleLevelEnum = map[string]TerminateInstanceRecycleLevelEnum{
	"FULL_RECYCLE": TerminateInstanceRecycleLevelFullRecycle,
}

var mappingTerminateInstanceRecycleLevelEnumLowerCase = map[string]TerminateInstanceRecycleLevelEnum{
	"full_recycle": TerminateInstanceRecycleLevelFullRecycle,
}

// GetTerminateInstanceRecycleLevelEnumValues Enumerates the set of values for TerminateInstanceRecycleLevelEnum
func GetTerminateInstanceRecycleLevelEnumValues() []TerminateInstanceRecycleLevelEnum {
	values := make([]TerminateInstanceRecycleLevelEnum, 0)
	for _, v := range mappingTerminateInstanceRecycleLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetTerminateInstanceRecycleLevelEnumStringValues Enumerates the set of values in String for TerminateInstanceRecycleLevelEnum
func GetTerminateInstanceRecycleLevelEnumStringValues() []string {
	return []string{
		"FULL_RECYCLE",
	}
}

// GetMappingTerminateInstanceRecycleLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTerminateInstanceRecycleLevelEnum(val string) (TerminateInstanceRecycleLevelEnum, bool) {
	enum, ok := mappingTerminateInstanceRecycleLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
