// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeApplicationUsageRequest wrapper for the SummarizeApplicationUsage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeApplicationUsage.go.html to see an example of how to use SummarizeApplicationUsageRequest.
type SummarizeApplicationUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The Fleet-unique identifier of the application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The type of the application.
	ApplicationType *string `mandatory:"false" contributesTo:"query" name:"applicationType"`

	// The vendor of the related Java Runtime.
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the related Java Runtime.
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the related Java Runtime.
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The file system path of the installation.
	InstallationPath *string `mandatory:"false" contributesTo:"query" name:"installationPath"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// Additional fields to include into the returned model on top of the required ones.
	// This parameter can also include 'approximateJreCount', 'approximateInstallationCount' and 'approximateManagedInstanceCount'.
	// For example 'approximateJreCount,approximateInstallationCount'.
	Fields []SummarizeApplicationUsageFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeApplicationUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort application views. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, _approximateJreCount_, _approximateInstallationCount_
	// and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _displayName_ and _osName_ is **ascending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeApplicationUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The operating system type.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeApplicationUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeApplicationUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeApplicationUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeApplicationUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeApplicationUsageResponse wrapper for the SummarizeApplicationUsage operation
type SummarizeApplicationUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApplicationUsageCollection instances
	ApplicationUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeApplicationUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeApplicationUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeApplicationUsageSortOrderEnum Enum with underlying type: string
type SummarizeApplicationUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeApplicationUsageSortOrderEnum
const (
	SummarizeApplicationUsageSortOrderAsc  SummarizeApplicationUsageSortOrderEnum = "ASC"
	SummarizeApplicationUsageSortOrderDesc SummarizeApplicationUsageSortOrderEnum = "DESC"
)

var mappingSummarizeApplicationUsageSortOrder = map[string]SummarizeApplicationUsageSortOrderEnum{
	"ASC":  SummarizeApplicationUsageSortOrderAsc,
	"DESC": SummarizeApplicationUsageSortOrderDesc,
}

// GetSummarizeApplicationUsageSortOrderEnumValues Enumerates the set of values for SummarizeApplicationUsageSortOrderEnum
func GetSummarizeApplicationUsageSortOrderEnumValues() []SummarizeApplicationUsageSortOrderEnum {
	values := make([]SummarizeApplicationUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeApplicationUsageSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeApplicationUsageSortByEnum Enum with underlying type: string
type SummarizeApplicationUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeApplicationUsageSortByEnum
const (
	SummarizeApplicationUsageSortByTimefirstseen                   SummarizeApplicationUsageSortByEnum = "timeFirstSeen"
	SummarizeApplicationUsageSortByTimelastseen                    SummarizeApplicationUsageSortByEnum = "timeLastSeen"
	SummarizeApplicationUsageSortByDisplayname                     SummarizeApplicationUsageSortByEnum = "displayName"
	SummarizeApplicationUsageSortByApproximatejrecount             SummarizeApplicationUsageSortByEnum = "approximateJreCount"
	SummarizeApplicationUsageSortByApproximateinstallationcount    SummarizeApplicationUsageSortByEnum = "approximateInstallationCount"
	SummarizeApplicationUsageSortByApproximatemanagedinstancecount SummarizeApplicationUsageSortByEnum = "approximateManagedInstanceCount"
	SummarizeApplicationUsageSortByOsname                          SummarizeApplicationUsageSortByEnum = "osName"
)

var mappingSummarizeApplicationUsageSortBy = map[string]SummarizeApplicationUsageSortByEnum{
	"timeFirstSeen":                   SummarizeApplicationUsageSortByTimefirstseen,
	"timeLastSeen":                    SummarizeApplicationUsageSortByTimelastseen,
	"displayName":                     SummarizeApplicationUsageSortByDisplayname,
	"approximateJreCount":             SummarizeApplicationUsageSortByApproximatejrecount,
	"approximateInstallationCount":    SummarizeApplicationUsageSortByApproximateinstallationcount,
	"approximateManagedInstanceCount": SummarizeApplicationUsageSortByApproximatemanagedinstancecount,
	"osName":                          SummarizeApplicationUsageSortByOsname,
}

// GetSummarizeApplicationUsageSortByEnumValues Enumerates the set of values for SummarizeApplicationUsageSortByEnum
func GetSummarizeApplicationUsageSortByEnumValues() []SummarizeApplicationUsageSortByEnum {
	values := make([]SummarizeApplicationUsageSortByEnum, 0)
	for _, v := range mappingSummarizeApplicationUsageSortBy {
		values = append(values, v)
	}
	return values
}
