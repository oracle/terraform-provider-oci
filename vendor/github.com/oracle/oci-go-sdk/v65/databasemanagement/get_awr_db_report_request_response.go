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

// GetAwrDbReportRequest wrapper for the GetAwrDbReport operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetAwrDbReport.go.html to see an example of how to use GetAwrDbReportRequest.
type GetAwrDbReportRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The parameter to filter the database by internal ID.
	// Note that the internal ID of the database can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" contributesTo:"path" name:"awrDbId"`

	// The optional multiple value query parameter to filter the database instance numbers.
	InstNums []int `contributesTo:"query" name:"instNums" collectionFormat:"csv"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnIdGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnIdGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot ID.
	EndSnIdLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnIdLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The query parameter to filter the AWR report types.
	ReportType GetAwrDbReportReportTypeEnum `mandatory:"false" contributesTo:"query" name:"reportType" omitEmpty:"true"`

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The format of the AWR report.
	ReportFormat GetAwrDbReportReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAwrDbReportRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAwrDbReportRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAwrDbReportRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAwrDbReportRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAwrDbReportRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetAwrDbReportReportTypeEnum(string(request.ReportType)); !ok && request.ReportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportType: %s. Supported values are: %s.", request.ReportType, strings.Join(GetGetAwrDbReportReportTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetAwrDbReportReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetAwrDbReportReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAwrDbReportResponse wrapper for the GetAwrDbReport operation
type GetAwrDbReportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AwrDbReport instance
	AwrDbReport `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAwrDbReportResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAwrDbReportResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAwrDbReportReportTypeEnum Enum with underlying type: string
type GetAwrDbReportReportTypeEnum string

// Set of constants representing the allowable values for GetAwrDbReportReportTypeEnum
const (
	GetAwrDbReportReportTypeAwr GetAwrDbReportReportTypeEnum = "AWR"
	GetAwrDbReportReportTypeAsh GetAwrDbReportReportTypeEnum = "ASH"
)

var mappingGetAwrDbReportReportTypeEnum = map[string]GetAwrDbReportReportTypeEnum{
	"AWR": GetAwrDbReportReportTypeAwr,
	"ASH": GetAwrDbReportReportTypeAsh,
}

var mappingGetAwrDbReportReportTypeEnumLowerCase = map[string]GetAwrDbReportReportTypeEnum{
	"awr": GetAwrDbReportReportTypeAwr,
	"ash": GetAwrDbReportReportTypeAsh,
}

// GetGetAwrDbReportReportTypeEnumValues Enumerates the set of values for GetAwrDbReportReportTypeEnum
func GetGetAwrDbReportReportTypeEnumValues() []GetAwrDbReportReportTypeEnum {
	values := make([]GetAwrDbReportReportTypeEnum, 0)
	for _, v := range mappingGetAwrDbReportReportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAwrDbReportReportTypeEnumStringValues Enumerates the set of values in String for GetAwrDbReportReportTypeEnum
func GetGetAwrDbReportReportTypeEnumStringValues() []string {
	return []string{
		"AWR",
		"ASH",
	}
}

// GetMappingGetAwrDbReportReportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAwrDbReportReportTypeEnum(val string) (GetAwrDbReportReportTypeEnum, bool) {
	enum, ok := mappingGetAwrDbReportReportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetAwrDbReportReportFormatEnum Enum with underlying type: string
type GetAwrDbReportReportFormatEnum string

// Set of constants representing the allowable values for GetAwrDbReportReportFormatEnum
const (
	GetAwrDbReportReportFormatHtml GetAwrDbReportReportFormatEnum = "HTML"
	GetAwrDbReportReportFormatText GetAwrDbReportReportFormatEnum = "TEXT"
)

var mappingGetAwrDbReportReportFormatEnum = map[string]GetAwrDbReportReportFormatEnum{
	"HTML": GetAwrDbReportReportFormatHtml,
	"TEXT": GetAwrDbReportReportFormatText,
}

var mappingGetAwrDbReportReportFormatEnumLowerCase = map[string]GetAwrDbReportReportFormatEnum{
	"html": GetAwrDbReportReportFormatHtml,
	"text": GetAwrDbReportReportFormatText,
}

// GetGetAwrDbReportReportFormatEnumValues Enumerates the set of values for GetAwrDbReportReportFormatEnum
func GetGetAwrDbReportReportFormatEnumValues() []GetAwrDbReportReportFormatEnum {
	values := make([]GetAwrDbReportReportFormatEnum, 0)
	for _, v := range mappingGetAwrDbReportReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAwrDbReportReportFormatEnumStringValues Enumerates the set of values in String for GetAwrDbReportReportFormatEnum
func GetGetAwrDbReportReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
	}
}

// GetMappingGetAwrDbReportReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAwrDbReportReportFormatEnum(val string) (GetAwrDbReportReportFormatEnum, bool) {
	enum, ok := mappingGetAwrDbReportReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
