// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetNetworkAnchorRequest wrapper for the GetNetworkAnchor operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/GetNetworkAnchor.go.html to see an example of how to use GetNetworkAnchorRequest.
type GetNetworkAnchorRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAnchor.
	NetworkAnchorId *string `mandatory:"true" contributesTo:"path" name:"networkAnchorId"`

	// The subscription service name values from [ORACLEDBATAZURE, ORACLEDBATGOOGLE, ORACLEDBATAWS]
	SubscriptionServiceName GetNetworkAnchorSubscriptionServiceNameEnum `mandatory:"true" contributesTo:"query" name:"subscriptionServiceName" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription in which to list resources.
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// OMHub Control Plane must know underlying CSP CP Region External Location Name.
	ExternalLocation *string `mandatory:"false" contributesTo:"query" name:"externalLocation"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetNetworkAnchorRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetNetworkAnchorRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetNetworkAnchorRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetNetworkAnchorRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetNetworkAnchorRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetNetworkAnchorSubscriptionServiceNameEnum(string(request.SubscriptionServiceName)); !ok && request.SubscriptionServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionServiceName: %s. Supported values are: %s.", request.SubscriptionServiceName, strings.Join(GetGetNetworkAnchorSubscriptionServiceNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetNetworkAnchorResponse wrapper for the GetNetworkAnchor operation
type GetNetworkAnchorResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The NetworkAnchor instance
	NetworkAnchor `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetNetworkAnchorResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetNetworkAnchorResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetNetworkAnchorSubscriptionServiceNameEnum Enum with underlying type: string
type GetNetworkAnchorSubscriptionServiceNameEnum string

// Set of constants representing the allowable values for GetNetworkAnchorSubscriptionServiceNameEnum
const (
	GetNetworkAnchorSubscriptionServiceNameOracledbatazure  GetNetworkAnchorSubscriptionServiceNameEnum = "ORACLEDBATAZURE"
	GetNetworkAnchorSubscriptionServiceNameOracledbatgoogle GetNetworkAnchorSubscriptionServiceNameEnum = "ORACLEDBATGOOGLE"
	GetNetworkAnchorSubscriptionServiceNameOracledbataws    GetNetworkAnchorSubscriptionServiceNameEnum = "ORACLEDBATAWS"
)

var mappingGetNetworkAnchorSubscriptionServiceNameEnum = map[string]GetNetworkAnchorSubscriptionServiceNameEnum{
	"ORACLEDBATAZURE":  GetNetworkAnchorSubscriptionServiceNameOracledbatazure,
	"ORACLEDBATGOOGLE": GetNetworkAnchorSubscriptionServiceNameOracledbatgoogle,
	"ORACLEDBATAWS":    GetNetworkAnchorSubscriptionServiceNameOracledbataws,
}

var mappingGetNetworkAnchorSubscriptionServiceNameEnumLowerCase = map[string]GetNetworkAnchorSubscriptionServiceNameEnum{
	"oracledbatazure":  GetNetworkAnchorSubscriptionServiceNameOracledbatazure,
	"oracledbatgoogle": GetNetworkAnchorSubscriptionServiceNameOracledbatgoogle,
	"oracledbataws":    GetNetworkAnchorSubscriptionServiceNameOracledbataws,
}

// GetGetNetworkAnchorSubscriptionServiceNameEnumValues Enumerates the set of values for GetNetworkAnchorSubscriptionServiceNameEnum
func GetGetNetworkAnchorSubscriptionServiceNameEnumValues() []GetNetworkAnchorSubscriptionServiceNameEnum {
	values := make([]GetNetworkAnchorSubscriptionServiceNameEnum, 0)
	for _, v := range mappingGetNetworkAnchorSubscriptionServiceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetGetNetworkAnchorSubscriptionServiceNameEnumStringValues Enumerates the set of values in String for GetNetworkAnchorSubscriptionServiceNameEnum
func GetGetNetworkAnchorSubscriptionServiceNameEnumStringValues() []string {
	return []string{
		"ORACLEDBATAZURE",
		"ORACLEDBATGOOGLE",
		"ORACLEDBATAWS",
	}
}

// GetMappingGetNetworkAnchorSubscriptionServiceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetNetworkAnchorSubscriptionServiceNameEnum(val string) (GetNetworkAnchorSubscriptionServiceNameEnum, bool) {
	enum, ok := mappingGetNetworkAnchorSubscriptionServiceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
