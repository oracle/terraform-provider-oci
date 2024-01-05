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

// GetAwrDatabaseReportRequest wrapper for the GetAwrDatabaseReport operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrDatabaseReport.go.html to see an example of how to use GetAwrDatabaseReportRequest.
type GetAwrDatabaseReportRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabases
	AwrSourceDatabaseIdentifier *string `mandatory:"true" contributesTo:"query" name:"awrSourceDatabaseIdentifier"`

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

	// The query parameter to filter the AWR report types.
	ReportType GetAwrDatabaseReportReportTypeEnum `mandatory:"false" contributesTo:"query" name:"reportType" omitEmpty:"true"`

	// The format of the AWR report.
	ReportFormat GetAwrDatabaseReportReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAwrDatabaseReportRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAwrDatabaseReportRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAwrDatabaseReportRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAwrDatabaseReportRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAwrDatabaseReportRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetAwrDatabaseReportReportTypeEnum(string(request.ReportType)); !ok && request.ReportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportType: %s. Supported values are: %s.", request.ReportType, strings.Join(GetGetAwrDatabaseReportReportTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetAwrDatabaseReportReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetAwrDatabaseReportReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAwrDatabaseReportResponse wrapper for the GetAwrDatabaseReport operation
type GetAwrDatabaseReportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AwrDatabaseReport instance
	AwrDatabaseReport `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAwrDatabaseReportResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAwrDatabaseReportResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAwrDatabaseReportReportTypeEnum Enum with underlying type: string
type GetAwrDatabaseReportReportTypeEnum string

// Set of constants representing the allowable values for GetAwrDatabaseReportReportTypeEnum
const (
	GetAwrDatabaseReportReportTypeAwr GetAwrDatabaseReportReportTypeEnum = "AWR"
	GetAwrDatabaseReportReportTypeAsh GetAwrDatabaseReportReportTypeEnum = "ASH"
)

var mappingGetAwrDatabaseReportReportTypeEnum = map[string]GetAwrDatabaseReportReportTypeEnum{
	"AWR": GetAwrDatabaseReportReportTypeAwr,
	"ASH": GetAwrDatabaseReportReportTypeAsh,
}

var mappingGetAwrDatabaseReportReportTypeEnumLowerCase = map[string]GetAwrDatabaseReportReportTypeEnum{
	"awr": GetAwrDatabaseReportReportTypeAwr,
	"ash": GetAwrDatabaseReportReportTypeAsh,
}

// GetGetAwrDatabaseReportReportTypeEnumValues Enumerates the set of values for GetAwrDatabaseReportReportTypeEnum
func GetGetAwrDatabaseReportReportTypeEnumValues() []GetAwrDatabaseReportReportTypeEnum {
	values := make([]GetAwrDatabaseReportReportTypeEnum, 0)
	for _, v := range mappingGetAwrDatabaseReportReportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAwrDatabaseReportReportTypeEnumStringValues Enumerates the set of values in String for GetAwrDatabaseReportReportTypeEnum
func GetGetAwrDatabaseReportReportTypeEnumStringValues() []string {
	return []string{
		"AWR",
		"ASH",
	}
}

// GetMappingGetAwrDatabaseReportReportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAwrDatabaseReportReportTypeEnum(val string) (GetAwrDatabaseReportReportTypeEnum, bool) {
	enum, ok := mappingGetAwrDatabaseReportReportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetAwrDatabaseReportReportFormatEnum Enum with underlying type: string
type GetAwrDatabaseReportReportFormatEnum string

// Set of constants representing the allowable values for GetAwrDatabaseReportReportFormatEnum
const (
	GetAwrDatabaseReportReportFormatHtml GetAwrDatabaseReportReportFormatEnum = "HTML"
	GetAwrDatabaseReportReportFormatText GetAwrDatabaseReportReportFormatEnum = "TEXT"
)

var mappingGetAwrDatabaseReportReportFormatEnum = map[string]GetAwrDatabaseReportReportFormatEnum{
	"HTML": GetAwrDatabaseReportReportFormatHtml,
	"TEXT": GetAwrDatabaseReportReportFormatText,
}

var mappingGetAwrDatabaseReportReportFormatEnumLowerCase = map[string]GetAwrDatabaseReportReportFormatEnum{
	"html": GetAwrDatabaseReportReportFormatHtml,
	"text": GetAwrDatabaseReportReportFormatText,
}

// GetGetAwrDatabaseReportReportFormatEnumValues Enumerates the set of values for GetAwrDatabaseReportReportFormatEnum
func GetGetAwrDatabaseReportReportFormatEnumValues() []GetAwrDatabaseReportReportFormatEnum {
	values := make([]GetAwrDatabaseReportReportFormatEnum, 0)
	for _, v := range mappingGetAwrDatabaseReportReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAwrDatabaseReportReportFormatEnumStringValues Enumerates the set of values in String for GetAwrDatabaseReportReportFormatEnum
func GetGetAwrDatabaseReportReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
	}
}

// GetMappingGetAwrDatabaseReportReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAwrDatabaseReportReportFormatEnum(val string) (GetAwrDatabaseReportReportFormatEnum, bool) {
	enum, ok := mappingGetAwrDatabaseReportReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
