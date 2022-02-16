// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmcontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDataKeysRequest wrapper for the ListDataKeys operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmcontrolplane/ListDataKeys.go.html to see an example of how to use ListDataKeysRequest.
type ListDataKeysRequest struct {

	// The OCID of the APM domain
	ApmDomainId *string `mandatory:"true" contributesTo:"path" name:"apmDomainId"`

	// Data key type.
	DataKeyType ListDataKeysDataKeyTypeEnum `mandatory:"false" contributesTo:"query" name:"dataKeyType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataKeysDataKeyTypeEnum(string(request.DataKeyType)); !ok && request.DataKeyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataKeyType: %s. Supported values are: %s.", request.DataKeyType, strings.Join(GetListDataKeysDataKeyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataKeysResponse wrapper for the ListDataKeys operation
type ListDataKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []DataKeySummary instance
	Items []DataKeySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDataKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataKeysDataKeyTypeEnum Enum with underlying type: string
type ListDataKeysDataKeyTypeEnum string

// Set of constants representing the allowable values for ListDataKeysDataKeyTypeEnum
const (
	ListDataKeysDataKeyTypePrivate ListDataKeysDataKeyTypeEnum = "PRIVATE"
	ListDataKeysDataKeyTypePublic  ListDataKeysDataKeyTypeEnum = "PUBLIC"
)

var mappingListDataKeysDataKeyTypeEnum = map[string]ListDataKeysDataKeyTypeEnum{
	"PRIVATE": ListDataKeysDataKeyTypePrivate,
	"PUBLIC":  ListDataKeysDataKeyTypePublic,
}

// GetListDataKeysDataKeyTypeEnumValues Enumerates the set of values for ListDataKeysDataKeyTypeEnum
func GetListDataKeysDataKeyTypeEnumValues() []ListDataKeysDataKeyTypeEnum {
	values := make([]ListDataKeysDataKeyTypeEnum, 0)
	for _, v := range mappingListDataKeysDataKeyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataKeysDataKeyTypeEnumStringValues Enumerates the set of values in String for ListDataKeysDataKeyTypeEnum
func GetListDataKeysDataKeyTypeEnumStringValues() []string {
	return []string{
		"PRIVATE",
		"PUBLIC",
	}
}

// GetMappingListDataKeysDataKeyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataKeysDataKeyTypeEnum(val string) (ListDataKeysDataKeyTypeEnum, bool) {
	mappingListDataKeysDataKeyTypeEnumIgnoreCase := make(map[string]ListDataKeysDataKeyTypeEnum)
	for k, v := range mappingListDataKeysDataKeyTypeEnum {
		mappingListDataKeysDataKeyTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataKeysDataKeyTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
