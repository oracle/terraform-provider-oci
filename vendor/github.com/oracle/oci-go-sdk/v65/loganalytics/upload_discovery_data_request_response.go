// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// UploadDiscoveryDataRequest wrapper for the UploadDiscoveryData operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UploadDiscoveryData.go.html to see an example of how to use UploadDiscoveryDataRequest.
type UploadDiscoveryDataRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Discovery data
	UploadDiscoveryDataDetails io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata key and value pairs separated by a semicolon. Example k1:v1;k2:v2;k3:v3
	OpcMetaProperties *string `mandatory:"false" contributesTo:"header" name:"opc-meta-properties"`

	// Discovery data type
	DiscoveryDataType UploadDiscoveryDataDiscoveryDataTypeEnum `mandatory:"false" contributesTo:"query" name:"discoveryDataType" omitEmpty:"true"`

	// The log group OCID that gets mapped to the logs in the discovery data.
	LogGroupId *string `mandatory:"false" contributesTo:"query" name:"logGroupId"`

	// Identifies the type of request payload.
	PayloadType UploadDiscoveryDataPayloadTypeEnum `mandatory:"false" contributesTo:"query" name:"payloadType" omitEmpty:"true"`

	// The content type of the log data.
	ContentType *string `mandatory:"false" contributesTo:"header" name:"content-type"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// A value of `100-continue` requests preliminary verification of the request method, path, and headers before the request body is sent.
	// If no error results from such verification, the server will send a 100 (Continue) interim response to indicate readiness for the request body.
	// The only allowed value for this parameter is "100-Continue" (case-insensitive).
	Expect *string `mandatory:"false" contributesTo:"header" name:"expect"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UploadDiscoveryDataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UploadDiscoveryDataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request UploadDiscoveryDataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.UploadDiscoveryDataDetails)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UploadDiscoveryDataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UploadDiscoveryDataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUploadDiscoveryDataDiscoveryDataTypeEnum(string(request.DiscoveryDataType)); !ok && request.DiscoveryDataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiscoveryDataType: %s. Supported values are: %s.", request.DiscoveryDataType, strings.Join(GetUploadDiscoveryDataDiscoveryDataTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUploadDiscoveryDataPayloadTypeEnum(string(request.PayloadType)); !ok && request.PayloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PayloadType: %s. Supported values are: %s.", request.PayloadType, strings.Join(GetUploadDiscoveryDataPayloadTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UploadDiscoveryDataResponse wrapper for the UploadDiscoveryData operation
type UploadDiscoveryDataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for log data.
	OpcObjectId *string `presentIn:"header" name:"opc-object-id"`

	// The time the upload was created, in the format defined by RFC3339
	TimeCreated *common.SDKTime `presentIn:"header" name:"timecreated"`
}

func (response UploadDiscoveryDataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UploadDiscoveryDataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UploadDiscoveryDataDiscoveryDataTypeEnum Enum with underlying type: string
type UploadDiscoveryDataDiscoveryDataTypeEnum string

// Set of constants representing the allowable values for UploadDiscoveryDataDiscoveryDataTypeEnum
const (
	UploadDiscoveryDataDiscoveryDataTypeEntity     UploadDiscoveryDataDiscoveryDataTypeEnum = "ENTITY"
	UploadDiscoveryDataDiscoveryDataTypeK8sObjects UploadDiscoveryDataDiscoveryDataTypeEnum = "K8S_OBJECTS"
)

var mappingUploadDiscoveryDataDiscoveryDataTypeEnum = map[string]UploadDiscoveryDataDiscoveryDataTypeEnum{
	"ENTITY":      UploadDiscoveryDataDiscoveryDataTypeEntity,
	"K8S_OBJECTS": UploadDiscoveryDataDiscoveryDataTypeK8sObjects,
}

var mappingUploadDiscoveryDataDiscoveryDataTypeEnumLowerCase = map[string]UploadDiscoveryDataDiscoveryDataTypeEnum{
	"entity":      UploadDiscoveryDataDiscoveryDataTypeEntity,
	"k8s_objects": UploadDiscoveryDataDiscoveryDataTypeK8sObjects,
}

// GetUploadDiscoveryDataDiscoveryDataTypeEnumValues Enumerates the set of values for UploadDiscoveryDataDiscoveryDataTypeEnum
func GetUploadDiscoveryDataDiscoveryDataTypeEnumValues() []UploadDiscoveryDataDiscoveryDataTypeEnum {
	values := make([]UploadDiscoveryDataDiscoveryDataTypeEnum, 0)
	for _, v := range mappingUploadDiscoveryDataDiscoveryDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUploadDiscoveryDataDiscoveryDataTypeEnumStringValues Enumerates the set of values in String for UploadDiscoveryDataDiscoveryDataTypeEnum
func GetUploadDiscoveryDataDiscoveryDataTypeEnumStringValues() []string {
	return []string{
		"ENTITY",
		"K8S_OBJECTS",
	}
}

// GetMappingUploadDiscoveryDataDiscoveryDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUploadDiscoveryDataDiscoveryDataTypeEnum(val string) (UploadDiscoveryDataDiscoveryDataTypeEnum, bool) {
	enum, ok := mappingUploadDiscoveryDataDiscoveryDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UploadDiscoveryDataPayloadTypeEnum Enum with underlying type: string
type UploadDiscoveryDataPayloadTypeEnum string

// Set of constants representing the allowable values for UploadDiscoveryDataPayloadTypeEnum
const (
	UploadDiscoveryDataPayloadTypeJson UploadDiscoveryDataPayloadTypeEnum = "JSON"
	UploadDiscoveryDataPayloadTypeGzip UploadDiscoveryDataPayloadTypeEnum = "GZIP"
	UploadDiscoveryDataPayloadTypeZip  UploadDiscoveryDataPayloadTypeEnum = "ZIP"
)

var mappingUploadDiscoveryDataPayloadTypeEnum = map[string]UploadDiscoveryDataPayloadTypeEnum{
	"JSON": UploadDiscoveryDataPayloadTypeJson,
	"GZIP": UploadDiscoveryDataPayloadTypeGzip,
	"ZIP":  UploadDiscoveryDataPayloadTypeZip,
}

var mappingUploadDiscoveryDataPayloadTypeEnumLowerCase = map[string]UploadDiscoveryDataPayloadTypeEnum{
	"json": UploadDiscoveryDataPayloadTypeJson,
	"gzip": UploadDiscoveryDataPayloadTypeGzip,
	"zip":  UploadDiscoveryDataPayloadTypeZip,
}

// GetUploadDiscoveryDataPayloadTypeEnumValues Enumerates the set of values for UploadDiscoveryDataPayloadTypeEnum
func GetUploadDiscoveryDataPayloadTypeEnumValues() []UploadDiscoveryDataPayloadTypeEnum {
	values := make([]UploadDiscoveryDataPayloadTypeEnum, 0)
	for _, v := range mappingUploadDiscoveryDataPayloadTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUploadDiscoveryDataPayloadTypeEnumStringValues Enumerates the set of values in String for UploadDiscoveryDataPayloadTypeEnum
func GetUploadDiscoveryDataPayloadTypeEnumStringValues() []string {
	return []string{
		"JSON",
		"GZIP",
		"ZIP",
	}
}

// GetMappingUploadDiscoveryDataPayloadTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUploadDiscoveryDataPayloadTypeEnum(val string) (UploadDiscoveryDataPayloadTypeEnum, bool) {
	enum, ok := mappingUploadDiscoveryDataPayloadTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
