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

// GetResourceAnchorRequest wrapper for the GetResourceAnchor operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/GetResourceAnchor.go.html to see an example of how to use GetResourceAnchorRequest.
type GetResourceAnchorRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
	ResourceAnchorId *string `mandatory:"true" contributesTo:"path" name:"resourceAnchorId"`

	// The subscription service name values from [ORACLEDBATAZURE, ORACLEDBATGOOGLE, ORACLEDBATAWS]
	SubscriptionServiceName GetResourceAnchorSubscriptionServiceNameEnum `mandatory:"true" contributesTo:"query" name:"subscriptionServiceName" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription in which to list resources.
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetResourceAnchorRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetResourceAnchorRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetResourceAnchorRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetResourceAnchorRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetResourceAnchorRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetResourceAnchorSubscriptionServiceNameEnum(string(request.SubscriptionServiceName)); !ok && request.SubscriptionServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionServiceName: %s. Supported values are: %s.", request.SubscriptionServiceName, strings.Join(GetGetResourceAnchorSubscriptionServiceNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetResourceAnchorResponse wrapper for the GetResourceAnchor operation
type GetResourceAnchorResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ResourceAnchor instance
	ResourceAnchor `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetResourceAnchorResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetResourceAnchorResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetResourceAnchorSubscriptionServiceNameEnum Enum with underlying type: string
type GetResourceAnchorSubscriptionServiceNameEnum string

// Set of constants representing the allowable values for GetResourceAnchorSubscriptionServiceNameEnum
const (
	GetResourceAnchorSubscriptionServiceNameOracledbatazure  GetResourceAnchorSubscriptionServiceNameEnum = "ORACLEDBATAZURE"
	GetResourceAnchorSubscriptionServiceNameOracledbatgoogle GetResourceAnchorSubscriptionServiceNameEnum = "ORACLEDBATGOOGLE"
	GetResourceAnchorSubscriptionServiceNameOracledbataws    GetResourceAnchorSubscriptionServiceNameEnum = "ORACLEDBATAWS"
)

var mappingGetResourceAnchorSubscriptionServiceNameEnum = map[string]GetResourceAnchorSubscriptionServiceNameEnum{
	"ORACLEDBATAZURE":  GetResourceAnchorSubscriptionServiceNameOracledbatazure,
	"ORACLEDBATGOOGLE": GetResourceAnchorSubscriptionServiceNameOracledbatgoogle,
	"ORACLEDBATAWS":    GetResourceAnchorSubscriptionServiceNameOracledbataws,
}

var mappingGetResourceAnchorSubscriptionServiceNameEnumLowerCase = map[string]GetResourceAnchorSubscriptionServiceNameEnum{
	"oracledbatazure":  GetResourceAnchorSubscriptionServiceNameOracledbatazure,
	"oracledbatgoogle": GetResourceAnchorSubscriptionServiceNameOracledbatgoogle,
	"oracledbataws":    GetResourceAnchorSubscriptionServiceNameOracledbataws,
}

// GetGetResourceAnchorSubscriptionServiceNameEnumValues Enumerates the set of values for GetResourceAnchorSubscriptionServiceNameEnum
func GetGetResourceAnchorSubscriptionServiceNameEnumValues() []GetResourceAnchorSubscriptionServiceNameEnum {
	values := make([]GetResourceAnchorSubscriptionServiceNameEnum, 0)
	for _, v := range mappingGetResourceAnchorSubscriptionServiceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetGetResourceAnchorSubscriptionServiceNameEnumStringValues Enumerates the set of values in String for GetResourceAnchorSubscriptionServiceNameEnum
func GetGetResourceAnchorSubscriptionServiceNameEnumStringValues() []string {
	return []string{
		"ORACLEDBATAZURE",
		"ORACLEDBATGOOGLE",
		"ORACLEDBATAWS",
	}
}

// GetMappingGetResourceAnchorSubscriptionServiceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetResourceAnchorSubscriptionServiceNameEnum(val string) (GetResourceAnchorSubscriptionServiceNameEnum, bool) {
	enum, ok := mappingGetResourceAnchorSubscriptionServiceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
