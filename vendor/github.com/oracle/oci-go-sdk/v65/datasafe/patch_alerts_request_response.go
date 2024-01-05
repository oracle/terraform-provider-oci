// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// PatchAlertsRequest wrapper for the PatchAlerts operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/PatchAlerts.go.html to see an example of how to use PatchAlertsRequest.
type PatchAlertsRequest struct {

	// The alert details to update the status of one or more alert specified by the alert IDs.
	PatchAlertsDetails `contributesTo:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the if-match parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel PatchAlertsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PatchAlertsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PatchAlertsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request PatchAlertsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PatchAlertsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PatchAlertsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchAlertsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetPatchAlertsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchAlertsResponse wrapper for the PatchAlerts operation
type PatchAlertsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID of the work request. Use GetWorkRequest with this OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response PatchAlertsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PatchAlertsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// PatchAlertsAccessLevelEnum Enum with underlying type: string
type PatchAlertsAccessLevelEnum string

// Set of constants representing the allowable values for PatchAlertsAccessLevelEnum
const (
	PatchAlertsAccessLevelRestricted PatchAlertsAccessLevelEnum = "RESTRICTED"
	PatchAlertsAccessLevelAccessible PatchAlertsAccessLevelEnum = "ACCESSIBLE"
)

var mappingPatchAlertsAccessLevelEnum = map[string]PatchAlertsAccessLevelEnum{
	"RESTRICTED": PatchAlertsAccessLevelRestricted,
	"ACCESSIBLE": PatchAlertsAccessLevelAccessible,
}

var mappingPatchAlertsAccessLevelEnumLowerCase = map[string]PatchAlertsAccessLevelEnum{
	"restricted": PatchAlertsAccessLevelRestricted,
	"accessible": PatchAlertsAccessLevelAccessible,
}

// GetPatchAlertsAccessLevelEnumValues Enumerates the set of values for PatchAlertsAccessLevelEnum
func GetPatchAlertsAccessLevelEnumValues() []PatchAlertsAccessLevelEnum {
	values := make([]PatchAlertsAccessLevelEnum, 0)
	for _, v := range mappingPatchAlertsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchAlertsAccessLevelEnumStringValues Enumerates the set of values in String for PatchAlertsAccessLevelEnum
func GetPatchAlertsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingPatchAlertsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchAlertsAccessLevelEnum(val string) (PatchAlertsAccessLevelEnum, bool) {
	enum, ok := mappingPatchAlertsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
