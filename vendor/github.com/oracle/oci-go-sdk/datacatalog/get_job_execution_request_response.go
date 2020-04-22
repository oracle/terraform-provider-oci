// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetJobExecutionRequest wrapper for the GetJobExecution operation
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
func (request GetJobExecutionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobExecutionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
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

var mappingGetJobExecutionFields = map[string]GetJobExecutionFieldsEnum{
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

// GetGetJobExecutionFieldsEnumValues Enumerates the set of values for GetJobExecutionFieldsEnum
func GetGetJobExecutionFieldsEnumValues() []GetJobExecutionFieldsEnum {
	values := make([]GetJobExecutionFieldsEnum, 0)
	for _, v := range mappingGetJobExecutionFields {
		values = append(values, v)
	}
	return values
}
