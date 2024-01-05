// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetJobExecutionRequest wrapper for the GetJobExecution operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetJobExecution.go.html to see an example of how to use GetJobExecutionRequest.
type GetJobExecutionRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique job key.
	JobKey *string `mandatory:"true" contributesTo:"path" name:"jobKey"`

	// The key of the job execution.
	JobExecutionKey *string `mandatory:"true" contributesTo:"path" name:"jobExecutionKey"`

	// Specifies the fields to return in a job execution response.
	Fields []GetJobExecutionFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJobExecutionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobExecutionRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetJobExecutionRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobExecutionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetJobExecutionRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetJobExecutionFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetJobExecutionFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetJobExecutionResponse wrapper for the GetJobExecution operation
type GetJobExecutionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The JobExecution instance
	JobExecution `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetJobExecutionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobExecutionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJobExecutionFieldsEnum Enum with underlying type: string
type GetJobExecutionFieldsEnum string

// Set of constants representing the allowable values for GetJobExecutionFieldsEnum
const (
	GetJobExecutionFieldsKey                 GetJobExecutionFieldsEnum = "key"
	GetJobExecutionFieldsJobkey              GetJobExecutionFieldsEnum = "jobKey"
	GetJobExecutionFieldsJobtype             GetJobExecutionFieldsEnum = "jobType"
	GetJobExecutionFieldsSubtype             GetJobExecutionFieldsEnum = "subType"
	GetJobExecutionFieldsParentkey           GetJobExecutionFieldsEnum = "parentKey"
	GetJobExecutionFieldsScheduleinstancekey GetJobExecutionFieldsEnum = "scheduleInstanceKey"
	GetJobExecutionFieldsLifecyclestate      GetJobExecutionFieldsEnum = "lifecycleState"
	GetJobExecutionFieldsTimecreated         GetJobExecutionFieldsEnum = "timeCreated"
	GetJobExecutionFieldsTimestarted         GetJobExecutionFieldsEnum = "timeStarted"
	GetJobExecutionFieldsTimeended           GetJobExecutionFieldsEnum = "timeEnded"
	GetJobExecutionFieldsErrorcode           GetJobExecutionFieldsEnum = "errorCode"
	GetJobExecutionFieldsErrormessage        GetJobExecutionFieldsEnum = "errorMessage"
	GetJobExecutionFieldsProcesskey          GetJobExecutionFieldsEnum = "processKey"
	GetJobExecutionFieldsExternalurl         GetJobExecutionFieldsEnum = "externalUrl"
	GetJobExecutionFieldsEventkey            GetJobExecutionFieldsEnum = "eventKey"
	GetJobExecutionFieldsDataentitykey       GetJobExecutionFieldsEnum = "dataEntityKey"
	GetJobExecutionFieldsCreatedbyid         GetJobExecutionFieldsEnum = "createdById"
	GetJobExecutionFieldsUpdatedbyid         GetJobExecutionFieldsEnum = "updatedById"
	GetJobExecutionFieldsProperties          GetJobExecutionFieldsEnum = "properties"
	GetJobExecutionFieldsUri                 GetJobExecutionFieldsEnum = "uri"
)

var mappingGetJobExecutionFieldsEnum = map[string]GetJobExecutionFieldsEnum{
	"key":                 GetJobExecutionFieldsKey,
	"jobKey":              GetJobExecutionFieldsJobkey,
	"jobType":             GetJobExecutionFieldsJobtype,
	"subType":             GetJobExecutionFieldsSubtype,
	"parentKey":           GetJobExecutionFieldsParentkey,
	"scheduleInstanceKey": GetJobExecutionFieldsScheduleinstancekey,
	"lifecycleState":      GetJobExecutionFieldsLifecyclestate,
	"timeCreated":         GetJobExecutionFieldsTimecreated,
	"timeStarted":         GetJobExecutionFieldsTimestarted,
	"timeEnded":           GetJobExecutionFieldsTimeended,
	"errorCode":           GetJobExecutionFieldsErrorcode,
	"errorMessage":        GetJobExecutionFieldsErrormessage,
	"processKey":          GetJobExecutionFieldsProcesskey,
	"externalUrl":         GetJobExecutionFieldsExternalurl,
	"eventKey":            GetJobExecutionFieldsEventkey,
	"dataEntityKey":       GetJobExecutionFieldsDataentitykey,
	"createdById":         GetJobExecutionFieldsCreatedbyid,
	"updatedById":         GetJobExecutionFieldsUpdatedbyid,
	"properties":          GetJobExecutionFieldsProperties,
	"uri":                 GetJobExecutionFieldsUri,
}

var mappingGetJobExecutionFieldsEnumLowerCase = map[string]GetJobExecutionFieldsEnum{
	"key":                 GetJobExecutionFieldsKey,
	"jobkey":              GetJobExecutionFieldsJobkey,
	"jobtype":             GetJobExecutionFieldsJobtype,
	"subtype":             GetJobExecutionFieldsSubtype,
	"parentkey":           GetJobExecutionFieldsParentkey,
	"scheduleinstancekey": GetJobExecutionFieldsScheduleinstancekey,
	"lifecyclestate":      GetJobExecutionFieldsLifecyclestate,
	"timecreated":         GetJobExecutionFieldsTimecreated,
	"timestarted":         GetJobExecutionFieldsTimestarted,
	"timeended":           GetJobExecutionFieldsTimeended,
	"errorcode":           GetJobExecutionFieldsErrorcode,
	"errormessage":        GetJobExecutionFieldsErrormessage,
	"processkey":          GetJobExecutionFieldsProcesskey,
	"externalurl":         GetJobExecutionFieldsExternalurl,
	"eventkey":            GetJobExecutionFieldsEventkey,
	"dataentitykey":       GetJobExecutionFieldsDataentitykey,
	"createdbyid":         GetJobExecutionFieldsCreatedbyid,
	"updatedbyid":         GetJobExecutionFieldsUpdatedbyid,
	"properties":          GetJobExecutionFieldsProperties,
	"uri":                 GetJobExecutionFieldsUri,
}

// GetGetJobExecutionFieldsEnumValues Enumerates the set of values for GetJobExecutionFieldsEnum
func GetGetJobExecutionFieldsEnumValues() []GetJobExecutionFieldsEnum {
	values := make([]GetJobExecutionFieldsEnum, 0)
	for _, v := range mappingGetJobExecutionFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetJobExecutionFieldsEnumStringValues Enumerates the set of values in String for GetJobExecutionFieldsEnum
func GetGetJobExecutionFieldsEnumStringValues() []string {
	return []string{
		"key",
		"jobKey",
		"jobType",
		"subType",
		"parentKey",
		"scheduleInstanceKey",
		"lifecycleState",
		"timeCreated",
		"timeStarted",
		"timeEnded",
		"errorCode",
		"errorMessage",
		"processKey",
		"externalUrl",
		"eventKey",
		"dataEntityKey",
		"createdById",
		"updatedById",
		"properties",
		"uri",
	}
}

// GetMappingGetJobExecutionFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetJobExecutionFieldsEnum(val string) (GetJobExecutionFieldsEnum, bool) {
	enum, ok := mappingGetJobExecutionFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
