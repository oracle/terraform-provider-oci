// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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
	ReportFormat GetReportVersionContentReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

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
	if _, ok := GetMappingGetReportVersionContentReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetReportVersionContentReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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

// GetReportVersionContentReportFormatEnum Enum with underlying type: string
type GetReportVersionContentReportFormatEnum string

// Set of constants representing the allowable values for GetReportVersionContentReportFormatEnum
const (
	GetReportVersionContentReportFormatCsv  GetReportVersionContentReportFormatEnum = "csv"
	GetReportVersionContentReportFormatJson GetReportVersionContentReportFormatEnum = "json"
	GetReportVersionContentReportFormatXml  GetReportVersionContentReportFormatEnum = "xml"
)

var mappingGetReportVersionContentReportFormatEnum = map[string]GetReportVersionContentReportFormatEnum{
	"csv":  GetReportVersionContentReportFormatCsv,
	"json": GetReportVersionContentReportFormatJson,
	"xml":  GetReportVersionContentReportFormatXml,
}

var mappingGetReportVersionContentReportFormatEnumLowerCase = map[string]GetReportVersionContentReportFormatEnum{
	"csv":  GetReportVersionContentReportFormatCsv,
	"json": GetReportVersionContentReportFormatJson,
	"xml":  GetReportVersionContentReportFormatXml,
}

// GetGetReportVersionContentReportFormatEnumValues Enumerates the set of values for GetReportVersionContentReportFormatEnum
func GetGetReportVersionContentReportFormatEnumValues() []GetReportVersionContentReportFormatEnum {
	values := make([]GetReportVersionContentReportFormatEnum, 0)
	for _, v := range mappingGetReportVersionContentReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetReportVersionContentReportFormatEnumStringValues Enumerates the set of values in String for GetReportVersionContentReportFormatEnum
func GetGetReportVersionContentReportFormatEnumStringValues() []string {
	return []string{
		"csv",
		"json",
		"xml",
	}
}

// GetMappingGetReportVersionContentReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetReportVersionContentReportFormatEnum(val string) (GetReportVersionContentReportFormatEnum, bool) {
	enum, ok := mappingGetReportVersionContentReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
