// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetJobDefinitionRequest wrapper for the GetJobDefinition operation
type GetJobDefinitionRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique job definition key.
	JobDefinitionKey *string `mandatory:"true" contributesTo:"path" name:"jobDefinitionKey"`

	// Specifies the fields to return in a job definition response.
	Fields []GetJobDefinitionFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJobDefinitionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobDefinitionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobDefinitionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetJobDefinitionResponse wrapper for the GetJobDefinition operation
type GetJobDefinitionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The JobDefinition instance
	JobDefinition `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetJobDefinitionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobDefinitionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJobDefinitionFieldsEnum Enum with underlying type: string
type GetJobDefinitionFieldsEnum string

// Set of constants representing the allowable values for GetJobDefinitionFieldsEnum
const (
	GetJobDefinitionFieldsKey                        GetJobDefinitionFieldsEnum = "key"
	GetJobDefinitionFieldsDisplayname                GetJobDefinitionFieldsEnum = "displayName"
	GetJobDefinitionFieldsDescription                GetJobDefinitionFieldsEnum = "description"
	GetJobDefinitionFieldsCatalogid                  GetJobDefinitionFieldsEnum = "catalogId"
	GetJobDefinitionFieldsJobtype                    GetJobDefinitionFieldsEnum = "jobType"
	GetJobDefinitionFieldsIsincremental              GetJobDefinitionFieldsEnum = "isIncremental"
	GetJobDefinitionFieldsDataassetkey               GetJobDefinitionFieldsEnum = "dataAssetKey"
	GetJobDefinitionFieldsConnectionkey              GetJobDefinitionFieldsEnum = "connectionKey"
	GetJobDefinitionFieldsInternalversion            GetJobDefinitionFieldsEnum = "internalVersion"
	GetJobDefinitionFieldsLifecyclestate             GetJobDefinitionFieldsEnum = "lifecycleState"
	GetJobDefinitionFieldsTimecreated                GetJobDefinitionFieldsEnum = "timeCreated"
	GetJobDefinitionFieldsTimeupdated                GetJobDefinitionFieldsEnum = "timeUpdated"
	GetJobDefinitionFieldsCreatedbyid                GetJobDefinitionFieldsEnum = "createdById"
	GetJobDefinitionFieldsUpdatedbyid                GetJobDefinitionFieldsEnum = "updatedById"
	GetJobDefinitionFieldsUri                        GetJobDefinitionFieldsEnum = "uri"
	GetJobDefinitionFieldsIssampledataextracted      GetJobDefinitionFieldsEnum = "isSampleDataExtracted"
	GetJobDefinitionFieldsSampledatasizeinmbs        GetJobDefinitionFieldsEnum = "sampleDataSizeInMBs"
	GetJobDefinitionFieldsTimelatestexecutionstarted GetJobDefinitionFieldsEnum = "timeLatestExecutionStarted"
	GetJobDefinitionFieldsTimelatestexecutionended   GetJobDefinitionFieldsEnum = "timeLatestExecutionEnded"
	GetJobDefinitionFieldsJobexecutionstate          GetJobDefinitionFieldsEnum = "jobExecutionState"
	GetJobDefinitionFieldsScheduletype               GetJobDefinitionFieldsEnum = "scheduleType"
	GetJobDefinitionFieldsProperties                 GetJobDefinitionFieldsEnum = "properties"
)

var mappingGetJobDefinitionFields = map[string]GetJobDefinitionFieldsEnum{
	"key":             GetJobDefinitionFieldsKey,
	"displayName":     GetJobDefinitionFieldsDisplayname,
	"description":     GetJobDefinitionFieldsDescription,
	"catalogId":       GetJobDefinitionFieldsCatalogid,
	"jobType":         GetJobDefinitionFieldsJobtype,
	"isIncremental":   GetJobDefinitionFieldsIsincremental,
	"dataAssetKey":    GetJobDefinitionFieldsDataassetkey,
	"connectionKey":   GetJobDefinitionFieldsConnectionkey,
	"internalVersion": GetJobDefinitionFieldsInternalversion,
	"lifecycleState":  GetJobDefinitionFieldsLifecyclestate,
	"timeCreated":     GetJobDefinitionFieldsTimecreated,
	"timeUpdated":     GetJobDefinitionFieldsTimeupdated,
	"createdById":     GetJobDefinitionFieldsCreatedbyid,
	"updatedById":     GetJobDefinitionFieldsUpdatedbyid,
	"uri":             GetJobDefinitionFieldsUri,
	"isSampleDataExtracted":      GetJobDefinitionFieldsIssampledataextracted,
	"sampleDataSizeInMBs":        GetJobDefinitionFieldsSampledatasizeinmbs,
	"timeLatestExecutionStarted": GetJobDefinitionFieldsTimelatestexecutionstarted,
	"timeLatestExecutionEnded":   GetJobDefinitionFieldsTimelatestexecutionended,
	"jobExecutionState":          GetJobDefinitionFieldsJobexecutionstate,
	"scheduleType":               GetJobDefinitionFieldsScheduletype,
	"properties":                 GetJobDefinitionFieldsProperties,
}

// GetGetJobDefinitionFieldsEnumValues Enumerates the set of values for GetJobDefinitionFieldsEnum
func GetGetJobDefinitionFieldsEnumValues() []GetJobDefinitionFieldsEnum {
	values := make([]GetJobDefinitionFieldsEnum, 0)
	for _, v := range mappingGetJobDefinitionFields {
		values = append(values, v)
	}
	return values
}
