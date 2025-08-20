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

// GetReportContentRequest wrapper for the GetReportContent operation
type GetReportContentRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Report.
	ReportId *string `mandatory:"true" contributesTo:"path" name:"reportId"`

	// The format of the report to download. Default is CSV.
	ReportFormat GetReportContentReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

	// Whether to include detailed report content or not. Default is false.
	ShouldIncludeDetails *bool `mandatory:"false" contributesTo:"query" name:"shouldIncludeDetails"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetReportContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetReportContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetReportContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetReportContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetReportContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetReportContentReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetReportContentReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetReportContentResponse wrapper for the GetReportContent operation
type GetReportContentResponse struct {

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

func (response GetReportContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetReportContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetReportContentReportFormatEnum Enum with underlying type: string
type GetReportContentReportFormatEnum string

// Set of constants representing the allowable values for GetReportContentReportFormatEnum
const (
	GetReportContentReportFormatCsv  GetReportContentReportFormatEnum = "csv"
	GetReportContentReportFormatJson GetReportContentReportFormatEnum = "json"
	GetReportContentReportFormatXml  GetReportContentReportFormatEnum = "xml"
)

var mappingGetReportContentReportFormatEnum = map[string]GetReportContentReportFormatEnum{
	"csv":  GetReportContentReportFormatCsv,
	"json": GetReportContentReportFormatJson,
	"xml":  GetReportContentReportFormatXml,
}

var mappingGetReportContentReportFormatEnumLowerCase = map[string]GetReportContentReportFormatEnum{
	"csv":  GetReportContentReportFormatCsv,
	"json": GetReportContentReportFormatJson,
	"xml":  GetReportContentReportFormatXml,
}

// GetGetReportContentReportFormatEnumValues Enumerates the set of values for GetReportContentReportFormatEnum
func GetGetReportContentReportFormatEnumValues() []GetReportContentReportFormatEnum {
	values := make([]GetReportContentReportFormatEnum, 0)
	for _, v := range mappingGetReportContentReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetReportContentReportFormatEnumStringValues Enumerates the set of values in String for GetReportContentReportFormatEnum
func GetGetReportContentReportFormatEnumStringValues() []string {
	return []string{
		"csv",
		"json",
		"xml",
	}
}

// GetMappingGetReportContentReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetReportContentReportFormatEnum(val string) (GetReportContentReportFormatEnum, bool) {
	enum, ok := mappingGetReportContentReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
