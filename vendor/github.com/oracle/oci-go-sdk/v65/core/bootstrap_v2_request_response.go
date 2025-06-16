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

// BootstrapV2Request wrapper for the BootstrapV2 operation
type BootstrapV2Request struct {

	// Identifies client asking for NAT gateway information.  Currently, must
	// be `"panamaagent-v1"`.
	Identity *string `mandatory:"true" contributesTo:"query" name:"identity"`

	// Filter the results by the query type.
	Type BootstrapV2TypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// Opaque token returned from previous `Bootstrap` or `GetUpdates` request.
	SequenceToken *string `mandatory:"false" contributesTo:"query" name:"sequenceToken"`

	// Maximum number of updates to be returned in response.
	MaxUpdates *int `mandatory:"false" contributesTo:"query" name:"maxUpdates"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request BootstrapV2Request) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request BootstrapV2Request) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request BootstrapV2Request) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request BootstrapV2Request) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request BootstrapV2Request) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBootstrapV2TypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetBootstrapV2TypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BootstrapV2Response wrapper for the BootstrapV2 operation
type BootstrapV2Response struct {

	// The underlying http response
	RawResponse *http.Response

	// The BootstrapOrGetUpdatesResponse instance
	BootstrapOrGetUpdatesResponse `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response BootstrapV2Response) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response BootstrapV2Response) HTTPResponse() *http.Response {
	return response.RawResponse
}

// BootstrapV2TypeEnum Enum with underlying type: string
type BootstrapV2TypeEnum string

// Set of constants representing the allowable values for BootstrapV2TypeEnum
const (
	BootstrapV2TypeNatgateway       BootstrapV2TypeEnum = "natgateway"
	BootstrapV2TypeDrgAttachmentVrf BootstrapV2TypeEnum = "drg-attachment-vrf"
	BootstrapV2TypeImportPolicy     BootstrapV2TypeEnum = "import-policy"
	BootstrapV2TypeExportPolicy     BootstrapV2TypeEnum = "export-policy"
	BootstrapV2TypeInetAttachCidrs  BootstrapV2TypeEnum = "inet-attach-cidrs"
	BootstrapV2TypeStaticRoutes     BootstrapV2TypeEnum = "static-routes"
	BootstrapV2TypeDrgRouteLimit    BootstrapV2TypeEnum = "drg-route-limit"
)

var mappingBootstrapV2TypeEnum = map[string]BootstrapV2TypeEnum{
	"natgateway":         BootstrapV2TypeNatgateway,
	"drg-attachment-vrf": BootstrapV2TypeDrgAttachmentVrf,
	"import-policy":      BootstrapV2TypeImportPolicy,
	"export-policy":      BootstrapV2TypeExportPolicy,
	"inet-attach-cidrs":  BootstrapV2TypeInetAttachCidrs,
	"static-routes":      BootstrapV2TypeStaticRoutes,
	"drg-route-limit":    BootstrapV2TypeDrgRouteLimit,
}

var mappingBootstrapV2TypeEnumLowerCase = map[string]BootstrapV2TypeEnum{
	"natgateway":         BootstrapV2TypeNatgateway,
	"drg-attachment-vrf": BootstrapV2TypeDrgAttachmentVrf,
	"import-policy":      BootstrapV2TypeImportPolicy,
	"export-policy":      BootstrapV2TypeExportPolicy,
	"inet-attach-cidrs":  BootstrapV2TypeInetAttachCidrs,
	"static-routes":      BootstrapV2TypeStaticRoutes,
	"drg-route-limit":    BootstrapV2TypeDrgRouteLimit,
}

// GetBootstrapV2TypeEnumValues Enumerates the set of values for BootstrapV2TypeEnum
func GetBootstrapV2TypeEnumValues() []BootstrapV2TypeEnum {
	values := make([]BootstrapV2TypeEnum, 0)
	for _, v := range mappingBootstrapV2TypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBootstrapV2TypeEnumStringValues Enumerates the set of values in String for BootstrapV2TypeEnum
func GetBootstrapV2TypeEnumStringValues() []string {
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

// GetMappingBootstrapV2TypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBootstrapV2TypeEnum(val string) (BootstrapV2TypeEnum, bool) {
	enum, ok := mappingBootstrapV2TypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
