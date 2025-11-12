// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetNodePoolOptionsRequest wrapper for the GetNodePoolOptions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/containerengine/GetNodePoolOptions.go.html to see an example of how to use GetNodePoolOptionsRequest.
type GetNodePoolOptionsRequest struct {

	// The id of the option set to retrieve. Use "all" get all options, or use a cluster ID to get options specific to the provided cluster.
	NodePoolOptionId *string `mandatory:"true" contributesTo:"path" name:"nodePoolOptionId"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Option to show all kubernetes patch versions
	ShouldListAllPatchVersions *bool `mandatory:"false" contributesTo:"query" name:"shouldListAllPatchVersions"`

	// Filter node pool options by OS type.
	NodePoolOsType GetNodePoolOptionsNodePoolOsTypeEnum `mandatory:"false" contributesTo:"query" name:"nodePoolOsType" omitEmpty:"true"`

	// Filter node pool options by OS architecture.
	NodePoolOsArch GetNodePoolOptionsNodePoolOsArchEnum `mandatory:"false" contributesTo:"query" name:"nodePoolOsArch" omitEmpty:"true"`

	// Filter node pool options by Kubernetes version.
	NodePoolK8sVersion *string `mandatory:"false" contributesTo:"query" name:"nodePoolK8sVersion"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetNodePoolOptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetNodePoolOptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetNodePoolOptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetNodePoolOptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetNodePoolOptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetNodePoolOptionsNodePoolOsTypeEnum(string(request.NodePoolOsType)); !ok && request.NodePoolOsType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodePoolOsType: %s. Supported values are: %s.", request.NodePoolOsType, strings.Join(GetGetNodePoolOptionsNodePoolOsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetNodePoolOptionsNodePoolOsArchEnum(string(request.NodePoolOsArch)); !ok && request.NodePoolOsArch != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodePoolOsArch: %s. Supported values are: %s.", request.NodePoolOsArch, strings.Join(GetGetNodePoolOptionsNodePoolOsArchEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetNodePoolOptionsResponse wrapper for the GetNodePoolOptions operation
type GetNodePoolOptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The NodePoolOptions instance
	NodePoolOptions `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetNodePoolOptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetNodePoolOptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetNodePoolOptionsNodePoolOsTypeEnum Enum with underlying type: string
type GetNodePoolOptionsNodePoolOsTypeEnum string

// Set of constants representing the allowable values for GetNodePoolOptionsNodePoolOsTypeEnum
const (
	GetNodePoolOptionsNodePoolOsTypeOl7    GetNodePoolOptionsNodePoolOsTypeEnum = "OL7"
	GetNodePoolOptionsNodePoolOsTypeOl8    GetNodePoolOptionsNodePoolOsTypeEnum = "OL8"
	GetNodePoolOptionsNodePoolOsTypeUbuntu GetNodePoolOptionsNodePoolOsTypeEnum = "UBUNTU"
)

var mappingGetNodePoolOptionsNodePoolOsTypeEnum = map[string]GetNodePoolOptionsNodePoolOsTypeEnum{
	"OL7":    GetNodePoolOptionsNodePoolOsTypeOl7,
	"OL8":    GetNodePoolOptionsNodePoolOsTypeOl8,
	"UBUNTU": GetNodePoolOptionsNodePoolOsTypeUbuntu,
}

var mappingGetNodePoolOptionsNodePoolOsTypeEnumLowerCase = map[string]GetNodePoolOptionsNodePoolOsTypeEnum{
	"ol7":    GetNodePoolOptionsNodePoolOsTypeOl7,
	"ol8":    GetNodePoolOptionsNodePoolOsTypeOl8,
	"ubuntu": GetNodePoolOptionsNodePoolOsTypeUbuntu,
}

// GetGetNodePoolOptionsNodePoolOsTypeEnumValues Enumerates the set of values for GetNodePoolOptionsNodePoolOsTypeEnum
func GetGetNodePoolOptionsNodePoolOsTypeEnumValues() []GetNodePoolOptionsNodePoolOsTypeEnum {
	values := make([]GetNodePoolOptionsNodePoolOsTypeEnum, 0)
	for _, v := range mappingGetNodePoolOptionsNodePoolOsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetNodePoolOptionsNodePoolOsTypeEnumStringValues Enumerates the set of values in String for GetNodePoolOptionsNodePoolOsTypeEnum
func GetGetNodePoolOptionsNodePoolOsTypeEnumStringValues() []string {
	return []string{
		"OL7",
		"OL8",
		"UBUNTU",
	}
}

// GetMappingGetNodePoolOptionsNodePoolOsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetNodePoolOptionsNodePoolOsTypeEnum(val string) (GetNodePoolOptionsNodePoolOsTypeEnum, bool) {
	enum, ok := mappingGetNodePoolOptionsNodePoolOsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetNodePoolOptionsNodePoolOsArchEnum Enum with underlying type: string
type GetNodePoolOptionsNodePoolOsArchEnum string

// Set of constants representing the allowable values for GetNodePoolOptionsNodePoolOsArchEnum
const (
	GetNodePoolOptionsNodePoolOsArchX8664   GetNodePoolOptionsNodePoolOsArchEnum = "X86_64"
	GetNodePoolOptionsNodePoolOsArchAarch64 GetNodePoolOptionsNodePoolOsArchEnum = "AARCH64"
)

var mappingGetNodePoolOptionsNodePoolOsArchEnum = map[string]GetNodePoolOptionsNodePoolOsArchEnum{
	"X86_64":  GetNodePoolOptionsNodePoolOsArchX8664,
	"AARCH64": GetNodePoolOptionsNodePoolOsArchAarch64,
}

var mappingGetNodePoolOptionsNodePoolOsArchEnumLowerCase = map[string]GetNodePoolOptionsNodePoolOsArchEnum{
	"x86_64":  GetNodePoolOptionsNodePoolOsArchX8664,
	"aarch64": GetNodePoolOptionsNodePoolOsArchAarch64,
}

// GetGetNodePoolOptionsNodePoolOsArchEnumValues Enumerates the set of values for GetNodePoolOptionsNodePoolOsArchEnum
func GetGetNodePoolOptionsNodePoolOsArchEnumValues() []GetNodePoolOptionsNodePoolOsArchEnum {
	values := make([]GetNodePoolOptionsNodePoolOsArchEnum, 0)
	for _, v := range mappingGetNodePoolOptionsNodePoolOsArchEnum {
		values = append(values, v)
	}
	return values
}

// GetGetNodePoolOptionsNodePoolOsArchEnumStringValues Enumerates the set of values in String for GetNodePoolOptionsNodePoolOsArchEnum
func GetGetNodePoolOptionsNodePoolOsArchEnumStringValues() []string {
	return []string{
		"X86_64",
		"AARCH64",
	}
}

// GetMappingGetNodePoolOptionsNodePoolOsArchEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetNodePoolOptionsNodePoolOsArchEnum(val string) (GetNodePoolOptionsNodePoolOsArchEnum, bool) {
	enum, ok := mappingGetNodePoolOptionsNodePoolOsArchEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
