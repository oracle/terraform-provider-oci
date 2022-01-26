// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// GetSqlExecutionPlanRequest wrapper for the GetSqlExecutionPlan operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetSqlExecutionPlan.go.html to see an example of how to use GetSqlExecutionPlanRequest.
type GetSqlExecutionPlanRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The SQL tuning task identifier. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" contributesTo:"path" name:"sqlTuningAdvisorTaskId"`

	// The SQL object id for the SQL tuning task. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlObjectId *int64 `mandatory:"true" contributesTo:"query" name:"sqlObjectId"`

	// The attribute of the SQL execution plan.
	Attribute GetSqlExecutionPlanAttributeEnum `mandatory:"true" contributesTo:"query" name:"attribute" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSqlExecutionPlanRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSqlExecutionPlanRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetSqlExecutionPlanRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetSqlExecutionPlanRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetSqlExecutionPlanResponse wrapper for the GetSqlExecutionPlan operation
type GetSqlExecutionPlanResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SqlTuningAdvisorTaskSqlExecutionPlan instance
	SqlTuningAdvisorTaskSqlExecutionPlan `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSqlExecutionPlanResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetSqlExecutionPlanResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetSqlExecutionPlanAttributeEnum Enum with underlying type: string
type GetSqlExecutionPlanAttributeEnum string

// Set of constants representing the allowable values for GetSqlExecutionPlanAttributeEnum
const (
	GetSqlExecutionPlanAttributeOriginal                 GetSqlExecutionPlanAttributeEnum = "ORIGINAL"
	GetSqlExecutionPlanAttributeOriginalWithAdjustedCost GetSqlExecutionPlanAttributeEnum = "ORIGINAL_WITH_ADJUSTED_COST"
	GetSqlExecutionPlanAttributeUsingSqlProfile          GetSqlExecutionPlanAttributeEnum = "USING_SQL_PROFILE"
	GetSqlExecutionPlanAttributeUsingNewIndices          GetSqlExecutionPlanAttributeEnum = "USING_NEW_INDICES"
)

var mappingGetSqlExecutionPlanAttribute = map[string]GetSqlExecutionPlanAttributeEnum{
	"ORIGINAL":                    GetSqlExecutionPlanAttributeOriginal,
	"ORIGINAL_WITH_ADJUSTED_COST": GetSqlExecutionPlanAttributeOriginalWithAdjustedCost,
	"USING_SQL_PROFILE":           GetSqlExecutionPlanAttributeUsingSqlProfile,
	"USING_NEW_INDICES":           GetSqlExecutionPlanAttributeUsingNewIndices,
}

// GetGetSqlExecutionPlanAttributeEnumValues Enumerates the set of values for GetSqlExecutionPlanAttributeEnum
func GetGetSqlExecutionPlanAttributeEnumValues() []GetSqlExecutionPlanAttributeEnum {
	values := make([]GetSqlExecutionPlanAttributeEnum, 0)
	for _, v := range mappingGetSqlExecutionPlanAttribute {
		values = append(values, v)
	}
	return values
}
