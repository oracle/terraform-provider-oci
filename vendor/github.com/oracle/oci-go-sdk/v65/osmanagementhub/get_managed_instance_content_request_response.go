// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// GetManagedInstanceContentRequest wrapper for the GetManagedInstanceContent operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetManagedInstanceContent.go.html to see an example of how to use GetManagedInstanceContentRequest.
type GetManagedInstanceContentRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A filter to return only vulnerabilities matching the given types.
	VulnerabilityType []VulnerabilityTypesEnum `contributesTo:"query" name:"vulnerabilityType" omitEmpty:"true" collectionFormat:"multi"`

	// The assigned erratum name. It's unique and not changeable.
	// Example: `ELSA-2020-5804`
	AdvisoryName []string `contributesTo:"query" name:"advisoryName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the erratum advisory name given.
	AdvisoryNameContains *string `mandatory:"false" contributesTo:"query" name:"advisoryNameContains"`

	// A filter to return only errata that match the given advisory types.
	AdvisoryType []AdvisoryTypesEnum `contributesTo:"query" name:"advisoryType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return vulnerabilities that match the given name. For Linux instances, this refers to the advisory name. For Windows instances, this refers to the Windows update display name.
	VulnerabilityName []string `contributesTo:"query" name:"vulnerabilityName" collectionFormat:"multi"`

	// A filter to return vulnerabilities that partially match the given name. For Linux instances, this refers to the advisory name. For Windows instances, this refers to the Windows update display name.
	VulnerabilityNameContains *string `mandatory:"false" contributesTo:"query" name:"vulnerabilityNameContains"`

	// The format of the report to download. Default is CSV.
	ReportFormat GetManagedInstanceContentReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetManagedInstanceContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetManagedInstanceContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetManagedInstanceContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetManagedInstanceContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetManagedInstanceContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.VulnerabilityType {
		if _, ok := GetMappingVulnerabilityTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VulnerabilityType: %s. Supported values are: %s.", val, strings.Join(GetVulnerabilityTypesEnumStringValues(), ",")))
		}
	}

	for _, val := range request.AdvisoryType {
		if _, ok := GetMappingAdvisoryTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisoryType: %s. Supported values are: %s.", val, strings.Join(GetAdvisoryTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingGetManagedInstanceContentReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetManagedInstanceContentReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetManagedInstanceContentResponse wrapper for the GetManagedInstanceContent operation
type GetManagedInstanceContentResponse struct {

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

func (response GetManagedInstanceContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetManagedInstanceContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetManagedInstanceContentReportFormatEnum Enum with underlying type: string
type GetManagedInstanceContentReportFormatEnum string

// Set of constants representing the allowable values for GetManagedInstanceContentReportFormatEnum
const (
	GetManagedInstanceContentReportFormatCsv  GetManagedInstanceContentReportFormatEnum = "csv"
	GetManagedInstanceContentReportFormatJson GetManagedInstanceContentReportFormatEnum = "json"
	GetManagedInstanceContentReportFormatXml  GetManagedInstanceContentReportFormatEnum = "xml"
)

var mappingGetManagedInstanceContentReportFormatEnum = map[string]GetManagedInstanceContentReportFormatEnum{
	"csv":  GetManagedInstanceContentReportFormatCsv,
	"json": GetManagedInstanceContentReportFormatJson,
	"xml":  GetManagedInstanceContentReportFormatXml,
}

var mappingGetManagedInstanceContentReportFormatEnumLowerCase = map[string]GetManagedInstanceContentReportFormatEnum{
	"csv":  GetManagedInstanceContentReportFormatCsv,
	"json": GetManagedInstanceContentReportFormatJson,
	"xml":  GetManagedInstanceContentReportFormatXml,
}

// GetGetManagedInstanceContentReportFormatEnumValues Enumerates the set of values for GetManagedInstanceContentReportFormatEnum
func GetGetManagedInstanceContentReportFormatEnumValues() []GetManagedInstanceContentReportFormatEnum {
	values := make([]GetManagedInstanceContentReportFormatEnum, 0)
	for _, v := range mappingGetManagedInstanceContentReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetManagedInstanceContentReportFormatEnumStringValues Enumerates the set of values in String for GetManagedInstanceContentReportFormatEnum
func GetGetManagedInstanceContentReportFormatEnumStringValues() []string {
	return []string{
		"csv",
		"json",
		"xml",
	}
}

// GetMappingGetManagedInstanceContentReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetManagedInstanceContentReportFormatEnum(val string) (GetManagedInstanceContentReportFormatEnum, bool) {
	enum, ok := mappingGetManagedInstanceContentReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
