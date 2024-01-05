// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetAwrDatabaseSqlReportRequest wrapper for the GetAwrDatabaseSqlReport operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrDatabaseSqlReport.go.html to see an example of how to use GetAwrDatabaseSqlReportRequest.
type GetAwrDatabaseSqlReportRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabases
	AwrSourceDatabaseIdentifier *string `mandatory:"true" contributesTo:"query" name:"awrSourceDatabaseIdentifier"`

	// The parameter to filter SQL by ID. Note that the SQL ID is generated internally by Oracle for each SQL statement and can be retrieved from AWR Report API (/awrHubs/{awrHubId}/awrDbReport).
	SqlId *string `mandatory:"true" contributesTo:"query" name:"sqlId"`

	// The optional single value query parameter to filter by database instance number.
	InstanceNumber *string `mandatory:"false" contributesTo:"query" name:"instanceNumber"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnapshotIdentifierGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnapshotIdentifierGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot Identifier.
	EndSnapshotIdentifierLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnapshotIdentifierLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The format of the AWR report.
	ReportFormat GetAwrDatabaseSqlReportReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAwrDatabaseSqlReportRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAwrDatabaseSqlReportRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAwrDatabaseSqlReportRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAwrDatabaseSqlReportRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAwrDatabaseSqlReportRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetAwrDatabaseSqlReportReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetAwrDatabaseSqlReportReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAwrDatabaseSqlReportResponse wrapper for the GetAwrDatabaseSqlReport operation
type GetAwrDatabaseSqlReportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AwrDatabaseSqlReport instance
	AwrDatabaseSqlReport `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAwrDatabaseSqlReportResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAwrDatabaseSqlReportResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAwrDatabaseSqlReportReportFormatEnum Enum with underlying type: string
type GetAwrDatabaseSqlReportReportFormatEnum string

// Set of constants representing the allowable values for GetAwrDatabaseSqlReportReportFormatEnum
const (
	GetAwrDatabaseSqlReportReportFormatHtml GetAwrDatabaseSqlReportReportFormatEnum = "HTML"
	GetAwrDatabaseSqlReportReportFormatText GetAwrDatabaseSqlReportReportFormatEnum = "TEXT"
)

var mappingGetAwrDatabaseSqlReportReportFormatEnum = map[string]GetAwrDatabaseSqlReportReportFormatEnum{
	"HTML": GetAwrDatabaseSqlReportReportFormatHtml,
	"TEXT": GetAwrDatabaseSqlReportReportFormatText,
}

var mappingGetAwrDatabaseSqlReportReportFormatEnumLowerCase = map[string]GetAwrDatabaseSqlReportReportFormatEnum{
	"html": GetAwrDatabaseSqlReportReportFormatHtml,
	"text": GetAwrDatabaseSqlReportReportFormatText,
}

// GetGetAwrDatabaseSqlReportReportFormatEnumValues Enumerates the set of values for GetAwrDatabaseSqlReportReportFormatEnum
func GetGetAwrDatabaseSqlReportReportFormatEnumValues() []GetAwrDatabaseSqlReportReportFormatEnum {
	values := make([]GetAwrDatabaseSqlReportReportFormatEnum, 0)
	for _, v := range mappingGetAwrDatabaseSqlReportReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAwrDatabaseSqlReportReportFormatEnumStringValues Enumerates the set of values in String for GetAwrDatabaseSqlReportReportFormatEnum
func GetGetAwrDatabaseSqlReportReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
	}
}

// GetMappingGetAwrDatabaseSqlReportReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAwrDatabaseSqlReportReportFormatEnum(val string) (GetAwrDatabaseSqlReportReportFormatEnum, bool) {
	enum, ok := mappingGetAwrDatabaseSqlReportReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
