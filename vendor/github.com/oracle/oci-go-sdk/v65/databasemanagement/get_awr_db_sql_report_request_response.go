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

// GetAwrDbSqlReportRequest wrapper for the GetAwrDbSqlReport operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetAwrDbSqlReport.go.html to see an example of how to use GetAwrDbSqlReportRequest.
type GetAwrDbSqlReportRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The parameter to filter the database by internal ID.
	// Note that the internal ID of the database can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" contributesTo:"path" name:"awrDbId"`

	// The parameter to filter SQL by ID. Note that the SQL ID is generated internally by Oracle for each SQL statement and can be retrieved from AWR Report API (/managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbReport) or Performance Hub API (/internal/managedDatabases/{managedDatabaseId}/actions/retrievePerformanceData)
	SqlId *string `mandatory:"true" contributesTo:"query" name:"sqlId"`

	// The optional single value query parameter to filter the database instance number.
	InstNum *string `mandatory:"false" contributesTo:"query" name:"instNum"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnIdGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnIdGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot ID.
	EndSnIdLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnIdLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The format of the AWR report.
	ReportFormat GetAwrDbSqlReportReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAwrDbSqlReportRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAwrDbSqlReportRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAwrDbSqlReportRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAwrDbSqlReportRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAwrDbSqlReportRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetAwrDbSqlReportReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetAwrDbSqlReportReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAwrDbSqlReportResponse wrapper for the GetAwrDbSqlReport operation
type GetAwrDbSqlReportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AwrDbSqlReport instance
	AwrDbSqlReport `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAwrDbSqlReportResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAwrDbSqlReportResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAwrDbSqlReportReportFormatEnum Enum with underlying type: string
type GetAwrDbSqlReportReportFormatEnum string

// Set of constants representing the allowable values for GetAwrDbSqlReportReportFormatEnum
const (
	GetAwrDbSqlReportReportFormatHtml GetAwrDbSqlReportReportFormatEnum = "HTML"
	GetAwrDbSqlReportReportFormatText GetAwrDbSqlReportReportFormatEnum = "TEXT"
)

var mappingGetAwrDbSqlReportReportFormatEnum = map[string]GetAwrDbSqlReportReportFormatEnum{
	"HTML": GetAwrDbSqlReportReportFormatHtml,
	"TEXT": GetAwrDbSqlReportReportFormatText,
}

var mappingGetAwrDbSqlReportReportFormatEnumLowerCase = map[string]GetAwrDbSqlReportReportFormatEnum{
	"html": GetAwrDbSqlReportReportFormatHtml,
	"text": GetAwrDbSqlReportReportFormatText,
}

// GetGetAwrDbSqlReportReportFormatEnumValues Enumerates the set of values for GetAwrDbSqlReportReportFormatEnum
func GetGetAwrDbSqlReportReportFormatEnumValues() []GetAwrDbSqlReportReportFormatEnum {
	values := make([]GetAwrDbSqlReportReportFormatEnum, 0)
	for _, v := range mappingGetAwrDbSqlReportReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAwrDbSqlReportReportFormatEnumStringValues Enumerates the set of values in String for GetAwrDbSqlReportReportFormatEnum
func GetGetAwrDbSqlReportReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
	}
}

// GetMappingGetAwrDbSqlReportReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAwrDbSqlReportReportFormatEnum(val string) (GetAwrDbSqlReportReportFormatEnum, bool) {
	enum, ok := mappingGetAwrDbSqlReportReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
