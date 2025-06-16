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

// GetUpdatesV2Request wrapper for the GetUpdatesV2 operation
type GetUpdatesV2Request struct {

	// Identifies client asking for NAT gateway information.  Currently, must
	// be `"panamaagent-v1"`.
	Identity *string `mandatory:"true" contributesTo:"query" name:"identity"`

	// Filter the results by the query type.
	Type GetUpdatesV2TypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// Opaque token returned from previous `Bootstrap` or `GetUpdates` request.
	SequenceToken *string `mandatory:"false" contributesTo:"query" name:"sequenceToken"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetUpdatesV2Request) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetUpdatesV2Request) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetUpdatesV2Request) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetUpdatesV2Request) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetUpdatesV2Request) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetUpdatesV2TypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetGetUpdatesV2TypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetUpdatesV2Response wrapper for the GetUpdatesV2 operation
type GetUpdatesV2Response struct {

	// The underlying http response
	RawResponse *http.Response

	// The BootstrapOrGetUpdatesResponse instance
	BootstrapOrGetUpdatesResponse `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetUpdatesV2Response) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetUpdatesV2Response) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetUpdatesV2TypeEnum Enum with underlying type: string
type GetUpdatesV2TypeEnum string

// Set of constants representing the allowable values for GetUpdatesV2TypeEnum
const (
	GetUpdatesV2TypeNatgateway       GetUpdatesV2TypeEnum = "natgateway"
	GetUpdatesV2TypeDrgAttachmentVrf GetUpdatesV2TypeEnum = "drg-attachment-vrf"
	GetUpdatesV2TypeImportPolicy     GetUpdatesV2TypeEnum = "import-policy"
	GetUpdatesV2TypeExportPolicy     GetUpdatesV2TypeEnum = "export-policy"
	GetUpdatesV2TypeInetAttachCidrs  GetUpdatesV2TypeEnum = "inet-attach-cidrs"
	GetUpdatesV2TypeStaticRoutes     GetUpdatesV2TypeEnum = "static-routes"
	GetUpdatesV2TypeDrgRouteLimit    GetUpdatesV2TypeEnum = "drg-route-limit"
)

var mappingGetUpdatesV2TypeEnum = map[string]GetUpdatesV2TypeEnum{
	"natgateway":         GetUpdatesV2TypeNatgateway,
	"drg-attachment-vrf": GetUpdatesV2TypeDrgAttachmentVrf,
	"import-policy":      GetUpdatesV2TypeImportPolicy,
	"export-policy":      GetUpdatesV2TypeExportPolicy,
	"inet-attach-cidrs":  GetUpdatesV2TypeInetAttachCidrs,
	"static-routes":      GetUpdatesV2TypeStaticRoutes,
	"drg-route-limit":    GetUpdatesV2TypeDrgRouteLimit,
}

var mappingGetUpdatesV2TypeEnumLowerCase = map[string]GetUpdatesV2TypeEnum{
	"natgateway":         GetUpdatesV2TypeNatgateway,
	"drg-attachment-vrf": GetUpdatesV2TypeDrgAttachmentVrf,
	"import-policy":      GetUpdatesV2TypeImportPolicy,
	"export-policy":      GetUpdatesV2TypeExportPolicy,
	"inet-attach-cidrs":  GetUpdatesV2TypeInetAttachCidrs,
	"static-routes":      GetUpdatesV2TypeStaticRoutes,
	"drg-route-limit":    GetUpdatesV2TypeDrgRouteLimit,
}

// GetGetUpdatesV2TypeEnumValues Enumerates the set of values for GetUpdatesV2TypeEnum
func GetGetUpdatesV2TypeEnumValues() []GetUpdatesV2TypeEnum {
	values := make([]GetUpdatesV2TypeEnum, 0)
	for _, v := range mappingGetUpdatesV2TypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetUpdatesV2TypeEnumStringValues Enumerates the set of values in String for GetUpdatesV2TypeEnum
func GetGetUpdatesV2TypeEnumStringValues() []string {
	return []string{
		"natgateway",
		"drg-attachment-vrf",
		"import-policy",
		"export-policy",
		"inet-attach-cidrs",
		"static-routes",
		"drg-route-limit",
	}
}

// GetMappingGetUpdatesV2TypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetUpdatesV2TypeEnum(val string) (GetUpdatesV2TypeEnum, bool) {
	enum, ok := mappingGetUpdatesV2TypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
