// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// GetAwrReportRequest wrapper for the GetAwrReport operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrReport.go.html to see an example of how to use GetAwrReportRequest.
type GetAwrReportRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// AWR source database identifier.
	AwrSourceDatabaseIdentifier *string `mandatory:"true" contributesTo:"query" name:"awrSourceDatabaseIdentifier"`

	// The format of the AWR report. Default report format is HTML.
	ReportFormat GetAwrReportReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

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

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAwrReportRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAwrReportRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAwrReportRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAwrReportRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAwrReportRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetAwrReportReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetAwrReportReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAwrReportResponse wrapper for the GetAwrReport operation
type GetAwrReportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AwrReport instance
	AwrReport `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAwrReportResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAwrReportResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAwrReportReportFormatEnum Enum with underlying type: string
type GetAwrReportReportFormatEnum string

// Set of constants representing the allowable values for GetAwrReportReportFormatEnum
const (
	GetAwrReportReportFormatHtml GetAwrReportReportFormatEnum = "HTML"
	GetAwrReportReportFormatText GetAwrReportReportFormatEnum = "TEXT"
)

var mappingGetAwrReportReportFormatEnum = map[string]GetAwrReportReportFormatEnum{
	"HTML": GetAwrReportReportFormatHtml,
	"TEXT": GetAwrReportReportFormatText,
}

// GetGetAwrReportReportFormatEnumValues Enumerates the set of values for GetAwrReportReportFormatEnum
func GetGetAwrReportReportFormatEnumValues() []GetAwrReportReportFormatEnum {
	values := make([]GetAwrReportReportFormatEnum, 0)
	for _, v := range mappingGetAwrReportReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAwrReportReportFormatEnumStringValues Enumerates the set of values in String for GetAwrReportReportFormatEnum
func GetGetAwrReportReportFormatEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
	}
}

// GetMappingGetAwrReportReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAwrReportReportFormatEnum(val string) (GetAwrReportReportFormatEnum, bool) {
	mappingGetAwrReportReportFormatEnumIgnoreCase := make(map[string]GetAwrReportReportFormatEnum)
	for k, v := range mappingGetAwrReportReportFormatEnum {
		mappingGetAwrReportReportFormatEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetAwrReportReportFormatEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
