// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetJobLogRequest wrapper for the GetJobLog operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetJobLog.go.html to see an example of how to use GetJobLogRequest.
type GetJobLogRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique job key.
	JobKey *string `mandatory:"true" contributesTo:"path" name:"jobKey"`

	// The key of the job execution.
	JobExecutionKey *string `mandatory:"true" contributesTo:"path" name:"jobExecutionKey"`

	// Unique job log key.
	JobLogKey *string `mandatory:"true" contributesTo:"path" name:"jobLogKey"`

	// Specifies the fields to return in a job log response.
	Fields []GetJobLogFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJobLogRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobLogRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetJobLogRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobLogRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetJobLogRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetJobLogFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetJobLogFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetJobLogResponse wrapper for the GetJobLog operation
type GetJobLogResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The JobLog instance
	JobLog `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetJobLogResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobLogResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJobLogFieldsEnum Enum with underlying type: string
type GetJobLogFieldsEnum string

// Set of constants representing the allowable values for GetJobLogFieldsEnum
const (
	GetJobLogFieldsKey             GetJobLogFieldsEnum = "key"
	GetJobLogFieldsJobexecutionkey GetJobLogFieldsEnum = "jobExecutionKey"
	GetJobLogFieldsCreatedbyid     GetJobLogFieldsEnum = "createdById"
	GetJobLogFieldsUpdatedbyid     GetJobLogFieldsEnum = "updatedById"
	GetJobLogFieldsTimeupdated     GetJobLogFieldsEnum = "timeUpdated"
	GetJobLogFieldsTimecreated     GetJobLogFieldsEnum = "timeCreated"
	GetJobLogFieldsSeverity        GetJobLogFieldsEnum = "severity"
	GetJobLogFieldsLogmessage      GetJobLogFieldsEnum = "logMessage"
	GetJobLogFieldsUri             GetJobLogFieldsEnum = "uri"
)

var mappingGetJobLogFieldsEnum = map[string]GetJobLogFieldsEnum{
	"key":             GetJobLogFieldsKey,
	"jobExecutionKey": GetJobLogFieldsJobexecutionkey,
	"createdById":     GetJobLogFieldsCreatedbyid,
	"updatedById":     GetJobLogFieldsUpdatedbyid,
	"timeUpdated":     GetJobLogFieldsTimeupdated,
	"timeCreated":     GetJobLogFieldsTimecreated,
	"severity":        GetJobLogFieldsSeverity,
	"logMessage":      GetJobLogFieldsLogmessage,
	"uri":             GetJobLogFieldsUri,
}

var mappingGetJobLogFieldsEnumLowerCase = map[string]GetJobLogFieldsEnum{
	"key":             GetJobLogFieldsKey,
	"jobexecutionkey": GetJobLogFieldsJobexecutionkey,
	"createdbyid":     GetJobLogFieldsCreatedbyid,
	"updatedbyid":     GetJobLogFieldsUpdatedbyid,
	"timeupdated":     GetJobLogFieldsTimeupdated,
	"timecreated":     GetJobLogFieldsTimecreated,
	"severity":        GetJobLogFieldsSeverity,
	"logmessage":      GetJobLogFieldsLogmessage,
	"uri":             GetJobLogFieldsUri,
}

// GetGetJobLogFieldsEnumValues Enumerates the set of values for GetJobLogFieldsEnum
func GetGetJobLogFieldsEnumValues() []GetJobLogFieldsEnum {
	values := make([]GetJobLogFieldsEnum, 0)
	for _, v := range mappingGetJobLogFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetJobLogFieldsEnumStringValues Enumerates the set of values in String for GetJobLogFieldsEnum
func GetGetJobLogFieldsEnumStringValues() []string {
	return []string{
		"key",
		"jobExecutionKey",
		"createdById",
		"updatedById",
		"timeUpdated",
		"timeCreated",
		"severity",
		"logMessage",
		"uri",
	}
}

// GetMappingGetJobLogFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetJobLogFieldsEnum(val string) (GetJobLogFieldsEnum, bool) {
	enum, ok := mappingGetJobLogFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
