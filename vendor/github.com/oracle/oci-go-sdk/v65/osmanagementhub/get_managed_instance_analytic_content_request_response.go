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

// GetManagedInstanceAnalyticContentRequest wrapper for the GetManagedInstanceAnalyticContent operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetManagedInstanceAnalyticContent.go.html to see an example of how to use GetManagedInstanceAnalyticContentRequest.
type GetManagedInstanceAnalyticContentRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	// This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group. This filter returns resources associated with this group.
	ManagedInstanceGroupId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceGroupId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle environment. This filter returns only resource contained with the specified lifecycle environment.
	LifecycleEnvironmentId *string `mandatory:"false" contributesTo:"query" name:"lifecycleEnvironmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage. This resource returns resources associated with this lifecycle stage.
	LifecycleStageId *string `mandatory:"false" contributesTo:"query" name:"lifecycleStageId"`

	// A filter to return only managed instances whose status matches the status provided.
	Status []ManagedInstanceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return instances that have the specified number of available security updates.
	SecurityUpdatesAvailableEqualsTo *int `mandatory:"false" contributesTo:"query" name:"securityUpdatesAvailableEqualsTo"`

	// A filter to return instances that have the specified number of available bug updates.
	BugUpdatesAvailableEqualsTo *int `mandatory:"false" contributesTo:"query" name:"bugUpdatesAvailableEqualsTo"`

	// A filter to return instances that have more available security updates than the number specified.
	SecurityUpdatesAvailableGreaterThan *int `mandatory:"false" contributesTo:"query" name:"securityUpdatesAvailableGreaterThan"`

	// A filter to return instances that have more available bug updates than the number specified.
	BugUpdatesAvailableGreaterThan *int `mandatory:"false" contributesTo:"query" name:"bugUpdatesAvailableGreaterThan"`

	// A filter to return only resources whose location matches the given value.
	Location []ManagedInstanceLocationEnum `contributesTo:"query" name:"location" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources whose location does not match the given value.
	LocationNotEqualTo []ManagedInstanceLocationEnum `contributesTo:"query" name:"locationNotEqualTo" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the given operating system family.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// Indicates whether to list only resources managed by the Autonomous Linux service.
	IsManagedByAutonomousLinux *bool `mandatory:"false" contributesTo:"query" name:"isManagedByAutonomousLinux"`

	// The format of the report to download. Default is CSV.
	ReportFormat GetManagedInstanceAnalyticContentReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

	// The type of the report the user wants to download. Default is ALL.
	ReportType GetManagedInstanceAnalyticContentReportTypeEnum `mandatory:"false" contributesTo:"query" name:"reportType" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetManagedInstanceAnalyticContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetManagedInstanceAnalyticContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetManagedInstanceAnalyticContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetManagedInstanceAnalyticContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetManagedInstanceAnalyticContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Status {
		if _, ok := GetMappingManagedInstanceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetManagedInstanceStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.Location {
		if _, ok := GetMappingManagedInstanceLocationEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Location: %s. Supported values are: %s.", val, strings.Join(GetManagedInstanceLocationEnumStringValues(), ",")))
		}
	}

	for _, val := range request.LocationNotEqualTo {
		if _, ok := GetMappingManagedInstanceLocationEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LocationNotEqualTo: %s. Supported values are: %s.", val, strings.Join(GetManagedInstanceLocationEnumStringValues(), ",")))
		}
	}

	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingGetManagedInstanceAnalyticContentReportFormatEnum(string(request.ReportFormat)); !ok && request.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", request.ReportFormat, strings.Join(GetGetManagedInstanceAnalyticContentReportFormatEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetManagedInstanceAnalyticContentReportTypeEnum(string(request.ReportType)); !ok && request.ReportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportType: %s. Supported values are: %s.", request.ReportType, strings.Join(GetGetManagedInstanceAnalyticContentReportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetManagedInstanceAnalyticContentResponse wrapper for the GetManagedInstanceAnalyticContent operation
type GetManagedInstanceAnalyticContentResponse struct {

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

func (response GetManagedInstanceAnalyticContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetManagedInstanceAnalyticContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetManagedInstanceAnalyticContentReportFormatEnum Enum with underlying type: string
type GetManagedInstanceAnalyticContentReportFormatEnum string

// Set of constants representing the allowable values for GetManagedInstanceAnalyticContentReportFormatEnum
const (
	GetManagedInstanceAnalyticContentReportFormatCsv  GetManagedInstanceAnalyticContentReportFormatEnum = "csv"
	GetManagedInstanceAnalyticContentReportFormatJson GetManagedInstanceAnalyticContentReportFormatEnum = "json"
	GetManagedInstanceAnalyticContentReportFormatXml  GetManagedInstanceAnalyticContentReportFormatEnum = "xml"
)

var mappingGetManagedInstanceAnalyticContentReportFormatEnum = map[string]GetManagedInstanceAnalyticContentReportFormatEnum{
	"csv":  GetManagedInstanceAnalyticContentReportFormatCsv,
	"json": GetManagedInstanceAnalyticContentReportFormatJson,
	"xml":  GetManagedInstanceAnalyticContentReportFormatXml,
}

var mappingGetManagedInstanceAnalyticContentReportFormatEnumLowerCase = map[string]GetManagedInstanceAnalyticContentReportFormatEnum{
	"csv":  GetManagedInstanceAnalyticContentReportFormatCsv,
	"json": GetManagedInstanceAnalyticContentReportFormatJson,
	"xml":  GetManagedInstanceAnalyticContentReportFormatXml,
}

// GetGetManagedInstanceAnalyticContentReportFormatEnumValues Enumerates the set of values for GetManagedInstanceAnalyticContentReportFormatEnum
func GetGetManagedInstanceAnalyticContentReportFormatEnumValues() []GetManagedInstanceAnalyticContentReportFormatEnum {
	values := make([]GetManagedInstanceAnalyticContentReportFormatEnum, 0)
	for _, v := range mappingGetManagedInstanceAnalyticContentReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetManagedInstanceAnalyticContentReportFormatEnumStringValues Enumerates the set of values in String for GetManagedInstanceAnalyticContentReportFormatEnum
func GetGetManagedInstanceAnalyticContentReportFormatEnumStringValues() []string {
	return []string{
		"csv",
		"json",
		"xml",
	}
}

// GetMappingGetManagedInstanceAnalyticContentReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetManagedInstanceAnalyticContentReportFormatEnum(val string) (GetManagedInstanceAnalyticContentReportFormatEnum, bool) {
	enum, ok := mappingGetManagedInstanceAnalyticContentReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetManagedInstanceAnalyticContentReportTypeEnum Enum with underlying type: string
type GetManagedInstanceAnalyticContentReportTypeEnum string

// Set of constants representing the allowable values for GetManagedInstanceAnalyticContentReportTypeEnum
const (
	GetManagedInstanceAnalyticContentReportTypeSecurity GetManagedInstanceAnalyticContentReportTypeEnum = "SECURITY"
	GetManagedInstanceAnalyticContentReportTypeBugfix   GetManagedInstanceAnalyticContentReportTypeEnum = "BUGFIX"
	GetManagedInstanceAnalyticContentReportTypeActivity GetManagedInstanceAnalyticContentReportTypeEnum = "ACTIVITY"
	GetManagedInstanceAnalyticContentReportTypeAll      GetManagedInstanceAnalyticContentReportTypeEnum = "ALL"
)

var mappingGetManagedInstanceAnalyticContentReportTypeEnum = map[string]GetManagedInstanceAnalyticContentReportTypeEnum{
	"SECURITY": GetManagedInstanceAnalyticContentReportTypeSecurity,
	"BUGFIX":   GetManagedInstanceAnalyticContentReportTypeBugfix,
	"ACTIVITY": GetManagedInstanceAnalyticContentReportTypeActivity,
	"ALL":      GetManagedInstanceAnalyticContentReportTypeAll,
}

var mappingGetManagedInstanceAnalyticContentReportTypeEnumLowerCase = map[string]GetManagedInstanceAnalyticContentReportTypeEnum{
	"security": GetManagedInstanceAnalyticContentReportTypeSecurity,
	"bugfix":   GetManagedInstanceAnalyticContentReportTypeBugfix,
	"activity": GetManagedInstanceAnalyticContentReportTypeActivity,
	"all":      GetManagedInstanceAnalyticContentReportTypeAll,
}

// GetGetManagedInstanceAnalyticContentReportTypeEnumValues Enumerates the set of values for GetManagedInstanceAnalyticContentReportTypeEnum
func GetGetManagedInstanceAnalyticContentReportTypeEnumValues() []GetManagedInstanceAnalyticContentReportTypeEnum {
	values := make([]GetManagedInstanceAnalyticContentReportTypeEnum, 0)
	for _, v := range mappingGetManagedInstanceAnalyticContentReportTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetManagedInstanceAnalyticContentReportTypeEnumStringValues Enumerates the set of values in String for GetManagedInstanceAnalyticContentReportTypeEnum
func GetGetManagedInstanceAnalyticContentReportTypeEnumStringValues() []string {
	return []string{
		"SECURITY",
		"BUGFIX",
		"ACTIVITY",
		"ALL",
	}
}

// GetMappingGetManagedInstanceAnalyticContentReportTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetManagedInstanceAnalyticContentReportTypeEnum(val string) (GetManagedInstanceAnalyticContentReportTypeEnum, bool) {
	enum, ok := mappingGetManagedInstanceAnalyticContentReportTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
