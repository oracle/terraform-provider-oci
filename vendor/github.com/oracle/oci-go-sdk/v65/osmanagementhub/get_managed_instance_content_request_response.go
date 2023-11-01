// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GetManagedInstanceContentRequest wrapper for the GetManagedInstanceContent operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetManagedInstanceContent.go.html to see an example of how to use GetManagedInstanceContentRequest.
type GetManagedInstanceContentRequest struct {

	// The OCID of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The assigned erratum name. It's unique and not changeable.
	// Example: `ELSA-2020-5804`
	AdvisoryName []string `contributesTo:"query" name:"advisoryName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the erratum advisory name given.
	AdvisoryNameContains *string `mandatory:"false" contributesTo:"query" name:"advisoryNameContains"`

	// A filter to return only errata that match the given advisory types.
	AdvisoryType []GetManagedInstanceContentAdvisoryTypeEnum `contributesTo:"query" name:"advisoryType" omitEmpty:"true" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetManagedInstanceContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetManagedInstanceContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetManagedInstanceContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetManagedInstanceContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetManagedInstanceContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.AdvisoryType {
		if _, ok := GetMappingGetManagedInstanceContentAdvisoryTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisoryType: %s. Supported values are: %s.", val, strings.Join(GetGetManagedInstanceContentAdvisoryTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetManagedInstanceContentResponse wrapper for the GetManagedInstanceContent operation
type GetManagedInstanceContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetManagedInstanceContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetManagedInstanceContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetManagedInstanceContentAdvisoryTypeEnum Enum with underlying type: string
type GetManagedInstanceContentAdvisoryTypeEnum string

// Set of constants representing the allowable values for GetManagedInstanceContentAdvisoryTypeEnum
const (
	GetManagedInstanceContentAdvisoryTypeSecurity    GetManagedInstanceContentAdvisoryTypeEnum = "SECURITY"
	GetManagedInstanceContentAdvisoryTypeBugfix      GetManagedInstanceContentAdvisoryTypeEnum = "BUGFIX"
	GetManagedInstanceContentAdvisoryTypeEnhancement GetManagedInstanceContentAdvisoryTypeEnum = "ENHANCEMENT"
)

var mappingGetManagedInstanceContentAdvisoryTypeEnum = map[string]GetManagedInstanceContentAdvisoryTypeEnum{
	"SECURITY":    GetManagedInstanceContentAdvisoryTypeSecurity,
	"BUGFIX":      GetManagedInstanceContentAdvisoryTypeBugfix,
	"ENHANCEMENT": GetManagedInstanceContentAdvisoryTypeEnhancement,
}

var mappingGetManagedInstanceContentAdvisoryTypeEnumLowerCase = map[string]GetManagedInstanceContentAdvisoryTypeEnum{
	"security":    GetManagedInstanceContentAdvisoryTypeSecurity,
	"bugfix":      GetManagedInstanceContentAdvisoryTypeBugfix,
	"enhancement": GetManagedInstanceContentAdvisoryTypeEnhancement,
}

// GetGetManagedInstanceContentAdvisoryTypeEnumValues Enumerates the set of values for GetManagedInstanceContentAdvisoryTypeEnum
func GetGetManagedInstanceContentAdvisoryTypeEnumValues() []GetManagedInstanceContentAdvisoryTypeEnum {
	values := make([]GetManagedInstanceContentAdvisoryTypeEnum, 0)
	for _, v := range mappingGetManagedInstanceContentAdvisoryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetManagedInstanceContentAdvisoryTypeEnumStringValues Enumerates the set of values in String for GetManagedInstanceContentAdvisoryTypeEnum
func GetGetManagedInstanceContentAdvisoryTypeEnumStringValues() []string {
	return []string{
		"SECURITY",
		"BUGFIX",
		"ENHANCEMENT",
	}
}

// GetMappingGetManagedInstanceContentAdvisoryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetManagedInstanceContentAdvisoryTypeEnum(val string) (GetManagedInstanceContentAdvisoryTypeEnum, bool) {
	enum, ok := mappingGetManagedInstanceContentAdvisoryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
