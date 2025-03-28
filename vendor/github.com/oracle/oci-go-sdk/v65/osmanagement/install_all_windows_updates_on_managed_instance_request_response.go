// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// InstallAllWindowsUpdatesOnManagedInstanceRequest wrapper for the InstallAllWindowsUpdatesOnManagedInstance operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallAllWindowsUpdatesOnManagedInstance.go.html to see an example of how to use InstallAllWindowsUpdatesOnManagedInstanceRequest.
type InstallAllWindowsUpdatesOnManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The type of updates to be applied
	UpdateType InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

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

func (request InstallAllWindowsUpdatesOnManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request InstallAllWindowsUpdatesOnManagedInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request InstallAllWindowsUpdatesOnManagedInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request InstallAllWindowsUpdatesOnManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request InstallAllWindowsUpdatesOnManagedInstanceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum(string(request.UpdateType)); !ok && request.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", request.UpdateType, strings.Join(GetInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstallAllWindowsUpdatesOnManagedInstanceResponse wrapper for the InstallAllWindowsUpdatesOnManagedInstance operation
type InstallAllWindowsUpdatesOnManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query the status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response InstallAllWindowsUpdatesOnManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response InstallAllWindowsUpdatesOnManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum Enum with underlying type: string
type InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum string

// Set of constants representing the allowable values for InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum
const (
	InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeSecurity    InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum = "SECURITY"
	InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeBugfix      InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum = "BUGFIX"
	InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnhancement InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum = "ENHANCEMENT"
	InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeOther       InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum = "OTHER"
	InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeKsplice     InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum = "KSPLICE"
	InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeAll         InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum = "ALL"
)

var mappingInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum = map[string]InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum{
	"SECURITY":    InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeSecurity,
	"BUGFIX":      InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeBugfix,
	"ENHANCEMENT": InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnhancement,
	"OTHER":       InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeOther,
	"KSPLICE":     InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeKsplice,
	"ALL":         InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeAll,
}

var mappingInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnumLowerCase = map[string]InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum{
	"security":    InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeSecurity,
	"bugfix":      InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeBugfix,
	"enhancement": InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnhancement,
	"other":       InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeOther,
	"ksplice":     InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeKsplice,
	"all":         InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeAll,
}

// GetInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnumValues Enumerates the set of values for InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum
func GetInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnumValues() []InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum {
	values := make([]InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum, 0)
	for _, v := range mappingInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnumStringValues Enumerates the set of values in String for InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum
func GetInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnumStringValues() []string {
	return []string{
		"SECURITY",
		"BUGFIX",
		"ENHANCEMENT",
		"OTHER",
		"KSPLICE",
		"ALL",
	}
}

// GetMappingInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum(val string) (InstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnum, bool) {
	enum, ok := mappingInstallAllWindowsUpdatesOnManagedInstanceUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
