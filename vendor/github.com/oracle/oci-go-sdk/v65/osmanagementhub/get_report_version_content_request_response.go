// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GetReportVersionContentRequest wrapper for the GetReportVersionContent operation
type GetReportVersionContentRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Report.
	ReportId *string `mandatory:"true" contributesTo:"path" name:"reportId"`

	// The version of the Report.
	ReportVersion *string `mandatory:"true" contributesTo:"path" name:"reportVersion"`

	// The format of the report to download. Default is CSV.
	ReportDownloadFormat GetReportVersionContentReportDownloadFormatEnum `mandatory:"false" contributesTo:"query" name:"reportDownloadFormat" omitEmpty:"true"`

	// Whether to include detailed report content or not. Default is false.
	ShouldIncludeDetails *bool `mandatory:"false" contributesTo:"query" name:"shouldIncludeDetails"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetReportVersionContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetReportVersionContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetReportVersionContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetReportVersionContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetReportVersionContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetReportVersionContentReportDownloadFormatEnum(string(request.ReportDownloadFormat)); !ok && request.ReportDownloadFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportDownloadFormat: %s. Supported values are: %s.", request.ReportDownloadFormat, strings.Join(GetGetReportVersionContentReportDownloadFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetReportVersionContentResponse wrapper for the GetReportVersionContent operation
type GetReportVersionContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetReportVersionContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetReportVersionContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetReportVersionContentReportDownloadFormatEnum Enum with underlying type: string
type GetReportVersionContentReportDownloadFormatEnum string

// Set of constants representing the allowable values for GetReportVersionContentReportDownloadFormatEnum
const (
	GetReportVersionContentReportDownloadFormatCsv  GetReportVersionContentReportDownloadFormatEnum = "csv"
	GetReportVersionContentReportDownloadFormatJson GetReportVersionContentReportDownloadFormatEnum = "json"
	GetReportVersionContentReportDownloadFormatXml  GetReportVersionContentReportDownloadFormatEnum = "xml"
)

var mappingGetReportVersionContentReportDownloadFormatEnum = map[string]GetReportVersionContentReportDownloadFormatEnum{
	"csv":  GetReportVersionContentReportDownloadFormatCsv,
	"json": GetReportVersionContentReportDownloadFormatJson,
	"xml":  GetReportVersionContentReportDownloadFormatXml,
}

var mappingGetReportVersionContentReportDownloadFormatEnumLowerCase = map[string]GetReportVersionContentReportDownloadFormatEnum{
	"csv":  GetReportVersionContentReportDownloadFormatCsv,
	"json": GetReportVersionContentReportDownloadFormatJson,
	"xml":  GetReportVersionContentReportDownloadFormatXml,
}

// GetGetReportVersionContentReportDownloadFormatEnumValues Enumerates the set of values for GetReportVersionContentReportDownloadFormatEnum
func GetGetReportVersionContentReportDownloadFormatEnumValues() []GetReportVersionContentReportDownloadFormatEnum {
	values := make([]GetReportVersionContentReportDownloadFormatEnum, 0)
	for _, v := range mappingGetReportVersionContentReportDownloadFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetReportVersionContentReportDownloadFormatEnumStringValues Enumerates the set of values in String for GetReportVersionContentReportDownloadFormatEnum
func GetGetReportVersionContentReportDownloadFormatEnumStringValues() []string {
	return []string{
		"csv",
		"json",
		"xml",
	}
}

// GetMappingGetReportVersionContentReportDownloadFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetReportVersionContentReportDownloadFormatEnum(val string) (GetReportVersionContentReportDownloadFormatEnum, bool) {
	enum, ok := mappingGetReportVersionContentReportDownloadFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
