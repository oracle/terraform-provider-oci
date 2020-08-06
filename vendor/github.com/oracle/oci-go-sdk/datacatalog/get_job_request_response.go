// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetJobRequest wrapper for the GetJob operation
type GetJobRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique job key.
	JobKey *string `mandatory:"true" contributesTo:"path" name:"jobKey"`

	// Specifies the fields to return in a job response.
	Fields []GetJobFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJobRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetJobResponse wrapper for the GetJob operation
type GetJobResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Job instance
	Job `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetJobResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJobFieldsEnum Enum with underlying type: string
type GetJobFieldsEnum string

// Set of constants representing the allowable values for GetJobFieldsEnum
const (
	GetJobFieldsKey                    GetJobFieldsEnum = "key"
	GetJobFieldsDisplayname            GetJobFieldsEnum = "displayName"
	GetJobFieldsDescription            GetJobFieldsEnum = "description"
	GetJobFieldsCatalogid              GetJobFieldsEnum = "catalogId"
	GetJobFieldsLifecyclestate         GetJobFieldsEnum = "lifecycleState"
	GetJobFieldsTimecreated            GetJobFieldsEnum = "timeCreated"
	GetJobFieldsTimeupdated            GetJobFieldsEnum = "timeUpdated"
	GetJobFieldsJobtype                GetJobFieldsEnum = "jobType"
	GetJobFieldsSchedulecronexpression GetJobFieldsEnum = "scheduleCronExpression"
	GetJobFieldsTimeschedulebegin      GetJobFieldsEnum = "timeScheduleBegin"
	GetJobFieldsTimescheduleend        GetJobFieldsEnum = "timeScheduleEnd"
	GetJobFieldsScheduletype           GetJobFieldsEnum = "scheduleType"
	GetJobFieldsConnectionkey          GetJobFieldsEnum = "connectionKey"
	GetJobFieldsJobdefinitionkey       GetJobFieldsEnum = "jobDefinitionKey"
	GetJobFieldsInternalversion        GetJobFieldsEnum = "internalVersion"
	GetJobFieldsExecutioncount         GetJobFieldsEnum = "executionCount"
	GetJobFieldsTimeoflatestexecution  GetJobFieldsEnum = "timeOfLatestExecution"
	GetJobFieldsExecutions             GetJobFieldsEnum = "executions"
	GetJobFieldsCreatedbyid            GetJobFieldsEnum = "createdById"
	GetJobFieldsUpdatedbyid            GetJobFieldsEnum = "updatedById"
	GetJobFieldsUri                    GetJobFieldsEnum = "uri"
	GetJobFieldsJobdefinitionname      GetJobFieldsEnum = "jobDefinitionName"
	GetJobFieldsErrorcode              GetJobFieldsEnum = "errorCode"
	GetJobFieldsErrormessage           GetJobFieldsEnum = "errorMessage"
)

var mappingGetJobFields = map[string]GetJobFieldsEnum{
	"key":                    GetJobFieldsKey,
	"displayName":            GetJobFieldsDisplayname,
	"description":            GetJobFieldsDescription,
	"catalogId":              GetJobFieldsCatalogid,
	"lifecycleState":         GetJobFieldsLifecyclestate,
	"timeCreated":            GetJobFieldsTimecreated,
	"timeUpdated":            GetJobFieldsTimeupdated,
	"jobType":                GetJobFieldsJobtype,
	"scheduleCronExpression": GetJobFieldsSchedulecronexpression,
	"timeScheduleBegin":      GetJobFieldsTimeschedulebegin,
	"timeScheduleEnd":        GetJobFieldsTimescheduleend,
	"scheduleType":           GetJobFieldsScheduletype,
	"connectionKey":          GetJobFieldsConnectionkey,
	"jobDefinitionKey":       GetJobFieldsJobdefinitionkey,
	"internalVersion":        GetJobFieldsInternalversion,
	"executionCount":         GetJobFieldsExecutioncount,
	"timeOfLatestExecution":  GetJobFieldsTimeoflatestexecution,
	"executions":             GetJobFieldsExecutions,
	"createdById":            GetJobFieldsCreatedbyid,
	"updatedById":            GetJobFieldsUpdatedbyid,
	"uri":                    GetJobFieldsUri,
	"jobDefinitionName":      GetJobFieldsJobdefinitionname,
	"errorCode":              GetJobFieldsErrorcode,
	"errorMessage":           GetJobFieldsErrormessage,
}

// GetGetJobFieldsEnumValues Enumerates the set of values for GetJobFieldsEnum
func GetGetJobFieldsEnumValues() []GetJobFieldsEnum {
	values := make([]GetJobFieldsEnum, 0)
	for _, v := range mappingGetJobFields {
		values = append(values, v)
	}
	return values
}
