// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// InstallAllUpdatesOnManagedInstanceGroupRequest wrapper for the InstallAllUpdatesOnManagedInstanceGroup operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallAllUpdatesOnManagedInstanceGroup.go.html to see an example of how to use InstallAllUpdatesOnManagedInstanceGroupRequest.
type InstallAllUpdatesOnManagedInstanceGroupRequest struct {

	// OCID for the managed instance group
	ManagedInstanceGroupId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceGroupId"`

	// The type of updates to be applied
	UpdateType InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request InstallAllUpdatesOnManagedInstanceGroupRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request InstallAllUpdatesOnManagedInstanceGroupRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request InstallAllUpdatesOnManagedInstanceGroupRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request InstallAllUpdatesOnManagedInstanceGroupRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// InstallAllUpdatesOnManagedInstanceGroupResponse wrapper for the InstallAllUpdatesOnManagedInstanceGroup operation
type InstallAllUpdatesOnManagedInstanceGroupResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query the status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response InstallAllUpdatesOnManagedInstanceGroupResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response InstallAllUpdatesOnManagedInstanceGroupResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum Enum with underlying type: string
type InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum string

// Set of constants representing the allowable values for InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum
const (
	InstallAllUpdatesOnManagedInstanceGroupUpdateTypeSecurity    InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum = "SECURITY"
	InstallAllUpdatesOnManagedInstanceGroupUpdateTypeBugfix      InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum = "BUGFIX"
	InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnhancement InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum = "ENHANCEMENT"
	InstallAllUpdatesOnManagedInstanceGroupUpdateTypeOther       InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum = "OTHER"
	InstallAllUpdatesOnManagedInstanceGroupUpdateTypeKsplice     InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum = "KSPLICE"
	InstallAllUpdatesOnManagedInstanceGroupUpdateTypeAll         InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum = "ALL"
)

var mappingInstallAllUpdatesOnManagedInstanceGroupUpdateType = map[string]InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum{
	"SECURITY":    InstallAllUpdatesOnManagedInstanceGroupUpdateTypeSecurity,
	"BUGFIX":      InstallAllUpdatesOnManagedInstanceGroupUpdateTypeBugfix,
	"ENHANCEMENT": InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnhancement,
	"OTHER":       InstallAllUpdatesOnManagedInstanceGroupUpdateTypeOther,
	"KSPLICE":     InstallAllUpdatesOnManagedInstanceGroupUpdateTypeKsplice,
	"ALL":         InstallAllUpdatesOnManagedInstanceGroupUpdateTypeAll,
}

// GetInstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnumValues Enumerates the set of values for InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum
func GetInstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnumValues() []InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum {
	values := make([]InstallAllUpdatesOnManagedInstanceGroupUpdateTypeEnum, 0)
	for _, v := range mappingInstallAllUpdatesOnManagedInstanceGroupUpdateType {
		values = append(values, v)
	}
	return values
}
