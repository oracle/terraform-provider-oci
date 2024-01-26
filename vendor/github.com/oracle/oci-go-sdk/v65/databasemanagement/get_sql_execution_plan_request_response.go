// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetSqlExecutionPlanRequest wrapper for the GetSqlExecutionPlan operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetSqlExecutionPlan.go.html to see an example of how to use GetSqlExecutionPlanRequest.
type GetSqlExecutionPlanRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The SQL tuning task identifier. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlTuningAdvisorTaskId *int64 `mandatory:"true" contributesTo:"path" name:"sqlTuningAdvisorTaskId"`

	// The SQL object ID for the SQL tuning task. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SqlObjectId *int64 `mandatory:"true" contributesTo:"query" name:"sqlObjectId"`

	// The attribute of the SQL execution plan.
	Attribute GetSqlExecutionPlanAttributeEnum `mandatory:"true" contributesTo:"query" name:"attribute" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSqlExecutionPlanRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSqlExecutionPlanRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetSqlExecutionPlanRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetSqlExecutionPlanAttributeEnum(string(request.Attribute)); !ok && request.Attribute != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Attribute: %s. Supported values are: %s.", request.Attribute, strings.Join(GetGetSqlExecutionPlanAttributeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	GetSqlExecutionPlanAttributeUsingParallelExecution   GetSqlExecutionPlanAttributeEnum = "USING_PARALLEL_EXECUTION"
)

var mappingGetSqlExecutionPlanAttributeEnum = map[string]GetSqlExecutionPlanAttributeEnum{
	"ORIGINAL":                    GetSqlExecutionPlanAttributeOriginal,
	"ORIGINAL_WITH_ADJUSTED_COST": GetSqlExecutionPlanAttributeOriginalWithAdjustedCost,
	"USING_SQL_PROFILE":           GetSqlExecutionPlanAttributeUsingSqlProfile,
	"USING_NEW_INDICES":           GetSqlExecutionPlanAttributeUsingNewIndices,
	"USING_PARALLEL_EXECUTION":    GetSqlExecutionPlanAttributeUsingParallelExecution,
}

var mappingGetSqlExecutionPlanAttributeEnumLowerCase = map[string]GetSqlExecutionPlanAttributeEnum{
	"original":                    GetSqlExecutionPlanAttributeOriginal,
	"original_with_adjusted_cost": GetSqlExecutionPlanAttributeOriginalWithAdjustedCost,
	"using_sql_profile":           GetSqlExecutionPlanAttributeUsingSqlProfile,
	"using_new_indices":           GetSqlExecutionPlanAttributeUsingNewIndices,
	"using_parallel_execution":    GetSqlExecutionPlanAttributeUsingParallelExecution,
}

// GetGetSqlExecutionPlanAttributeEnumValues Enumerates the set of values for GetSqlExecutionPlanAttributeEnum
func GetGetSqlExecutionPlanAttributeEnumValues() []GetSqlExecutionPlanAttributeEnum {
	values := make([]GetSqlExecutionPlanAttributeEnum, 0)
	for _, v := range mappingGetSqlExecutionPlanAttributeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetSqlExecutionPlanAttributeEnumStringValues Enumerates the set of values in String for GetSqlExecutionPlanAttributeEnum
func GetGetSqlExecutionPlanAttributeEnumStringValues() []string {
	return []string{
		"ORIGINAL",
		"ORIGINAL_WITH_ADJUSTED_COST",
		"USING_SQL_PROFILE",
		"USING_NEW_INDICES",
		"USING_PARALLEL_EXECUTION",
	}
}

// GetMappingGetSqlExecutionPlanAttributeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetSqlExecutionPlanAttributeEnum(val string) (GetSqlExecutionPlanAttributeEnum, bool) {
	enum, ok := mappingGetSqlExecutionPlanAttributeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
