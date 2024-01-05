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

// AlertsUpdateRequest wrapper for the AlertsUpdate operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/AlertsUpdate.go.html to see an example of how to use AlertsUpdateRequest.
type AlertsUpdateRequest struct {

	// The details to update the alerts in the specified compartment.
	AlertsUpdateDetails `contributesTo:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the if-match parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel AlertsUpdateAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request AlertsUpdateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request AlertsUpdateRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request AlertsUpdateRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request AlertsUpdateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request AlertsUpdateRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlertsUpdateAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetAlertsUpdateAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlertsUpdateResponse wrapper for the AlertsUpdate operation
type AlertsUpdateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID of the work request. Use GetWorkRequest with this OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response AlertsUpdateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response AlertsUpdateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// AlertsUpdateAccessLevelEnum Enum with underlying type: string
type AlertsUpdateAccessLevelEnum string

// Set of constants representing the allowable values for AlertsUpdateAccessLevelEnum
const (
	AlertsUpdateAccessLevelRestricted AlertsUpdateAccessLevelEnum = "RESTRICTED"
	AlertsUpdateAccessLevelAccessible AlertsUpdateAccessLevelEnum = "ACCESSIBLE"
)

var mappingAlertsUpdateAccessLevelEnum = map[string]AlertsUpdateAccessLevelEnum{
	"RESTRICTED": AlertsUpdateAccessLevelRestricted,
	"ACCESSIBLE": AlertsUpdateAccessLevelAccessible,
}

var mappingAlertsUpdateAccessLevelEnumLowerCase = map[string]AlertsUpdateAccessLevelEnum{
	"restricted": AlertsUpdateAccessLevelRestricted,
	"accessible": AlertsUpdateAccessLevelAccessible,
}

// GetAlertsUpdateAccessLevelEnumValues Enumerates the set of values for AlertsUpdateAccessLevelEnum
func GetAlertsUpdateAccessLevelEnumValues() []AlertsUpdateAccessLevelEnum {
	values := make([]AlertsUpdateAccessLevelEnum, 0)
	for _, v := range mappingAlertsUpdateAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertsUpdateAccessLevelEnumStringValues Enumerates the set of values in String for AlertsUpdateAccessLevelEnum
func GetAlertsUpdateAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingAlertsUpdateAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertsUpdateAccessLevelEnum(val string) (AlertsUpdateAccessLevelEnum, bool) {
	enum, ok := mappingAlertsUpdateAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
