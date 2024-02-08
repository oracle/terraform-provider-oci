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

// GetJobMetricsRequest wrapper for the GetJobMetrics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetJobMetrics.go.html to see an example of how to use GetJobMetricsRequest.
type GetJobMetricsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique job key.
	JobKey *string `mandatory:"true" contributesTo:"path" name:"jobKey"`

	// The key of the job execution.
	JobExecutionKey *string `mandatory:"true" contributesTo:"path" name:"jobExecutionKey"`

	// Unique job metrics key.
	JobMetricsKey *string `mandatory:"true" contributesTo:"path" name:"jobMetricsKey"`

	// Specifies the fields to return in a job metric response.
	Fields []GetJobMetricsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJobMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetJobMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetJobMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetJobMetricsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetJobMetricsFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetJobMetricsResponse wrapper for the GetJobMetrics operation
type GetJobMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The JobMetric instance
	JobMetric `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetJobMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJobMetricsFieldsEnum Enum with underlying type: string
type GetJobMetricsFieldsEnum string

// Set of constants representing the allowable values for GetJobMetricsFieldsEnum
const (
	GetJobMetricsFieldsKey             GetJobMetricsFieldsEnum = "key"
	GetJobMetricsFieldsDescription     GetJobMetricsFieldsEnum = "description"
	GetJobMetricsFieldsDisplayname     GetJobMetricsFieldsEnum = "displayName"
	GetJobMetricsFieldsTimeinserted    GetJobMetricsFieldsEnum = "timeInserted"
	GetJobMetricsFieldsCategory        GetJobMetricsFieldsEnum = "category"
	GetJobMetricsFieldsSubcategory     GetJobMetricsFieldsEnum = "subCategory"
	GetJobMetricsFieldsUnit            GetJobMetricsFieldsEnum = "unit"
	GetJobMetricsFieldsValue           GetJobMetricsFieldsEnum = "value"
	GetJobMetricsFieldsBatchkey        GetJobMetricsFieldsEnum = "batchKey"
	GetJobMetricsFieldsJobexecutionkey GetJobMetricsFieldsEnum = "jobExecutionKey"
	GetJobMetricsFieldsCreatedbyid     GetJobMetricsFieldsEnum = "createdById"
	GetJobMetricsFieldsUpdatedbyid     GetJobMetricsFieldsEnum = "updatedById"
	GetJobMetricsFieldsTimeupdated     GetJobMetricsFieldsEnum = "timeUpdated"
	GetJobMetricsFieldsTimecreated     GetJobMetricsFieldsEnum = "timeCreated"
	GetJobMetricsFieldsUri             GetJobMetricsFieldsEnum = "uri"
)

var mappingGetJobMetricsFieldsEnum = map[string]GetJobMetricsFieldsEnum{
	"key":             GetJobMetricsFieldsKey,
	"description":     GetJobMetricsFieldsDescription,
	"displayName":     GetJobMetricsFieldsDisplayname,
	"timeInserted":    GetJobMetricsFieldsTimeinserted,
	"category":        GetJobMetricsFieldsCategory,
	"subCategory":     GetJobMetricsFieldsSubcategory,
	"unit":            GetJobMetricsFieldsUnit,
	"value":           GetJobMetricsFieldsValue,
	"batchKey":        GetJobMetricsFieldsBatchkey,
	"jobExecutionKey": GetJobMetricsFieldsJobexecutionkey,
	"createdById":     GetJobMetricsFieldsCreatedbyid,
	"updatedById":     GetJobMetricsFieldsUpdatedbyid,
	"timeUpdated":     GetJobMetricsFieldsTimeupdated,
	"timeCreated":     GetJobMetricsFieldsTimecreated,
	"uri":             GetJobMetricsFieldsUri,
}

var mappingGetJobMetricsFieldsEnumLowerCase = map[string]GetJobMetricsFieldsEnum{
	"key":             GetJobMetricsFieldsKey,
	"description":     GetJobMetricsFieldsDescription,
	"displayname":     GetJobMetricsFieldsDisplayname,
	"timeinserted":    GetJobMetricsFieldsTimeinserted,
	"category":        GetJobMetricsFieldsCategory,
	"subcategory":     GetJobMetricsFieldsSubcategory,
	"unit":            GetJobMetricsFieldsUnit,
	"value":           GetJobMetricsFieldsValue,
	"batchkey":        GetJobMetricsFieldsBatchkey,
	"jobexecutionkey": GetJobMetricsFieldsJobexecutionkey,
	"createdbyid":     GetJobMetricsFieldsCreatedbyid,
	"updatedbyid":     GetJobMetricsFieldsUpdatedbyid,
	"timeupdated":     GetJobMetricsFieldsTimeupdated,
	"timecreated":     GetJobMetricsFieldsTimecreated,
	"uri":             GetJobMetricsFieldsUri,
}

// GetGetJobMetricsFieldsEnumValues Enumerates the set of values for GetJobMetricsFieldsEnum
func GetGetJobMetricsFieldsEnumValues() []GetJobMetricsFieldsEnum {
	values := make([]GetJobMetricsFieldsEnum, 0)
	for _, v := range mappingGetJobMetricsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetJobMetricsFieldsEnumStringValues Enumerates the set of values in String for GetJobMetricsFieldsEnum
func GetGetJobMetricsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"description",
		"displayName",
		"timeInserted",
		"category",
		"subCategory",
		"unit",
		"value",
		"batchKey",
		"jobExecutionKey",
		"createdById",
		"updatedById",
		"timeUpdated",
		"timeCreated",
		"uri",
	}
}

// GetMappingGetJobMetricsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetJobMetricsFieldsEnum(val string) (GetJobMetricsFieldsEnum, bool) {
	enum, ok := mappingGetJobMetricsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
