// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// InstallAllPackageUpdatesOnManagedInstanceRequest wrapper for the InstallAllPackageUpdatesOnManagedInstance operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/InstallAllPackageUpdatesOnManagedInstance.go.html to see an example of how to use InstallAllPackageUpdatesOnManagedInstanceRequest.
type InstallAllPackageUpdatesOnManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The type of updates to be applied
	UpdateType InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

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

func (request InstallAllPackageUpdatesOnManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request InstallAllPackageUpdatesOnManagedInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request InstallAllPackageUpdatesOnManagedInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request InstallAllPackageUpdatesOnManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request InstallAllPackageUpdatesOnManagedInstanceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum(string(request.UpdateType)); !ok && request.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", request.UpdateType, strings.Join(GetInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstallAllPackageUpdatesOnManagedInstanceResponse wrapper for the InstallAllPackageUpdatesOnManagedInstance operation
type InstallAllPackageUpdatesOnManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query the status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response InstallAllPackageUpdatesOnManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response InstallAllPackageUpdatesOnManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum Enum with underlying type: string
type InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum string

// Set of constants representing the allowable values for InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum
const (
	InstallAllPackageUpdatesOnManagedInstanceUpdateTypeSecurity    InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum = "SECURITY"
	InstallAllPackageUpdatesOnManagedInstanceUpdateTypeBugfix      InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum = "BUGFIX"
	InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnhancement InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum = "ENHANCEMENT"
	InstallAllPackageUpdatesOnManagedInstanceUpdateTypeOther       InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum = "OTHER"
	InstallAllPackageUpdatesOnManagedInstanceUpdateTypeKsplice     InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum = "KSPLICE"
	InstallAllPackageUpdatesOnManagedInstanceUpdateTypeAll         InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum = "ALL"
)

var mappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum = map[string]InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum{
	"SECURITY":    InstallAllPackageUpdatesOnManagedInstanceUpdateTypeSecurity,
	"BUGFIX":      InstallAllPackageUpdatesOnManagedInstanceUpdateTypeBugfix,
	"ENHANCEMENT": InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnhancement,
	"OTHER":       InstallAllPackageUpdatesOnManagedInstanceUpdateTypeOther,
	"KSPLICE":     InstallAllPackageUpdatesOnManagedInstanceUpdateTypeKsplice,
	"ALL":         InstallAllPackageUpdatesOnManagedInstanceUpdateTypeAll,
}

// GetInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnumValues Enumerates the set of values for InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum
func GetInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnumValues() []InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum {
	values := make([]InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum, 0)
	for _, v := range mappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnumStringValues Enumerates the set of values in String for InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum
func GetInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnumStringValues() []string {
	return []string{
		"SECURITY",
		"BUGFIX",
		"ENHANCEMENT",
		"OTHER",
		"KSPLICE",
		"ALL",
	}
}

// GetMappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum(val string) (InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum, bool) {
	mappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnumIgnoreCase := make(map[string]InstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum)
	for k, v := range mappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnum {
		mappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingInstallAllPackageUpdatesOnManagedInstanceUpdateTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
