// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListReportsRequest wrapper for the ListReports operation
type ListReportsRequest struct {

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Indicates whether to include subcompartments in the returned results. Default is false.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return only reports currently in the given state.
	LifecycleState ReportLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the report. This resource returns resources associated with this report.
	ReportId *string `mandatory:"false" contributesTo:"query" name:"reportId"`

	// The advisory severity.
	AdvisorySeverity []AdvisorySeverityEnum `contributesTo:"query" name:"advisorySeverity" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only packages that match the given update classification type.
	ClassificationType []ClassificationTypesEnum `contributesTo:"query" name:"classificationType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the given operating system family.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the given vendor name.
	VendorName []VendorNameEnum `contributesTo:"query" name:"vendorName" omitEmpty:"true" collectionFormat:"multi"`

	// The type of the Report.
	ReportType []ReportTypeEnum `contributesTo:"query" name:"reportType" omitEmpty:"true" collectionFormat:"multi"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance. This filter returns resources associated with this managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group. This filter returns resources associated with this group.
	ManagedInstanceGroupId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceGroupId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dynamic set. This filter returns resources associated with this dynamic set.
	DynamicSetId *string `mandatory:"false" contributesTo:"query" name:"dynamicSetId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source. This filter returns resources associated with this software source.
	SoftwareSourceId *string `mandatory:"false" contributesTo:"query" name:"softwareSourceId"`

	// The OCID of the compartment included in the report.
	CompartmentIdInReport *string `mandatory:"false" contributesTo:"query" name:"compartmentIdInReport"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeUpdated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources with a date on or before the given value, in ISO 8601 format.
	// Example: 2017-07-14T02:40:00.000Z
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// A filter to return only resources with a date on or after the given value, in ISO 8601 format.
	// Example: 2017-07-14T02:40:00.000Z
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReportLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetReportLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.AdvisorySeverity {
		if _, ok := GetMappingAdvisorySeverityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdvisorySeverity: %s. Supported values are: %s.", val, strings.Join(GetAdvisorySeverityEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ClassificationType {
		if _, ok := GetMappingClassificationTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationType: %s. Supported values are: %s.", val, strings.Join(GetClassificationTypesEnumStringValues(), ",")))
		}
	}

	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	for _, val := range request.VendorName {
		if _, ok := GetMappingVendorNameEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VendorName: %s. Supported values are: %s.", val, strings.Join(GetVendorNameEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ReportType {
		if _, ok := GetMappingReportTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportType: %s. Supported values are: %s.", val, strings.Join(GetReportTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReportsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReportsResponse wrapper for the ListReports operation
type ListReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReportCollection instances
	ReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReportsSortOrderEnum Enum with underlying type: string
type ListReportsSortOrderEnum string

// Set of constants representing the allowable values for ListReportsSortOrderEnum
const (
	ListReportsSortOrderAsc  ListReportsSortOrderEnum = "ASC"
	ListReportsSortOrderDesc ListReportsSortOrderEnum = "DESC"
)

var mappingListReportsSortOrderEnum = map[string]ListReportsSortOrderEnum{
	"ASC":  ListReportsSortOrderAsc,
	"DESC": ListReportsSortOrderDesc,
}

var mappingListReportsSortOrderEnumLowerCase = map[string]ListReportsSortOrderEnum{
	"asc":  ListReportsSortOrderAsc,
	"desc": ListReportsSortOrderDesc,
}

// GetListReportsSortOrderEnumValues Enumerates the set of values for ListReportsSortOrderEnum
func GetListReportsSortOrderEnumValues() []ListReportsSortOrderEnum {
	values := make([]ListReportsSortOrderEnum, 0)
	for _, v := range mappingListReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportsSortOrderEnumStringValues Enumerates the set of values in String for ListReportsSortOrderEnum
func GetListReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportsSortOrderEnum(val string) (ListReportsSortOrderEnum, bool) {
	enum, ok := mappingListReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReportsSortByEnum Enum with underlying type: string
type ListReportsSortByEnum string

// Set of constants representing the allowable values for ListReportsSortByEnum
const (
	ListReportsSortByTimecreated ListReportsSortByEnum = "timeCreated"
	ListReportsSortByTimeupdated ListReportsSortByEnum = "timeUpdated"
	ListReportsSortByDisplayname ListReportsSortByEnum = "displayName"
)

var mappingListReportsSortByEnum = map[string]ListReportsSortByEnum{
	"timeCreated": ListReportsSortByTimecreated,
	"timeUpdated": ListReportsSortByTimeupdated,
	"displayName": ListReportsSortByDisplayname,
}

var mappingListReportsSortByEnumLowerCase = map[string]ListReportsSortByEnum{
	"timecreated": ListReportsSortByTimecreated,
	"timeupdated": ListReportsSortByTimeupdated,
	"displayname": ListReportsSortByDisplayname,
}

// GetListReportsSortByEnumValues Enumerates the set of values for ListReportsSortByEnum
func GetListReportsSortByEnumValues() []ListReportsSortByEnum {
	values := make([]ListReportsSortByEnum, 0)
	for _, v := range mappingListReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReportsSortByEnumStringValues Enumerates the set of values in String for ListReportsSortByEnum
func GetListReportsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReportsSortByEnum(val string) (ListReportsSortByEnum, bool) {
	enum, ok := mappingListReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
