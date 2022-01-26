// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeInstallationUsageRequest wrapper for the SummarizeInstallationUsage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeInstallationUsage.go.html to see an example of how to use SummarizeInstallationUsageRequest.
type SummarizeInstallationUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The vendor of the related Java Runtime.
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the related Java Runtime.
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the related Java Runtime.
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The file system path of the installation.
	InstallationPath *string `mandatory:"false" contributesTo:"query" name:"installationPath"`

	// The Fleet-unique identifier of the related application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// Additional fields to include into the returned model on top of the required ones.
	// This parameter can also include 'approximateApplicationCount' and 'approximateManagedInstanceCount'.
	// For example 'approximateApplicationCount,approximateManagedInstanceCount'.
	Fields []SummarizeInstallationUsageFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeInstallationUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort installation views. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, and _jreVersion_, _approximateApplicationCount_
	// and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _jreDistribution_ and _jreVendor_ is **ascending**. If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeInstallationUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The operating system type.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeInstallationUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeInstallationUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeInstallationUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeInstallationUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeInstallationUsageResponse wrapper for the SummarizeInstallationUsage operation
type SummarizeInstallationUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstallationUsageCollection instances
	InstallationUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeInstallationUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeInstallationUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeInstallationUsageSortOrderEnum Enum with underlying type: string
type SummarizeInstallationUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeInstallationUsageSortOrderEnum
const (
	SummarizeInstallationUsageSortOrderAsc  SummarizeInstallationUsageSortOrderEnum = "ASC"
	SummarizeInstallationUsageSortOrderDesc SummarizeInstallationUsageSortOrderEnum = "DESC"
)

var mappingSummarizeInstallationUsageSortOrder = map[string]SummarizeInstallationUsageSortOrderEnum{
	"ASC":  SummarizeInstallationUsageSortOrderAsc,
	"DESC": SummarizeInstallationUsageSortOrderDesc,
}

// GetSummarizeInstallationUsageSortOrderEnumValues Enumerates the set of values for SummarizeInstallationUsageSortOrderEnum
func GetSummarizeInstallationUsageSortOrderEnumValues() []SummarizeInstallationUsageSortOrderEnum {
	values := make([]SummarizeInstallationUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeInstallationUsageSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeInstallationUsageSortByEnum Enum with underlying type: string
type SummarizeInstallationUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeInstallationUsageSortByEnum
const (
	SummarizeInstallationUsageSortByJredistribution                 SummarizeInstallationUsageSortByEnum = "jreDistribution"
	SummarizeInstallationUsageSortByJrevendor                       SummarizeInstallationUsageSortByEnum = "jreVendor"
	SummarizeInstallationUsageSortByJreversion                      SummarizeInstallationUsageSortByEnum = "jreVersion"
	SummarizeInstallationUsageSortByPath                            SummarizeInstallationUsageSortByEnum = "path"
	SummarizeInstallationUsageSortByTimefirstseen                   SummarizeInstallationUsageSortByEnum = "timeFirstSeen"
	SummarizeInstallationUsageSortByTimelastseen                    SummarizeInstallationUsageSortByEnum = "timeLastSeen"
	SummarizeInstallationUsageSortByApproximateapplicationcount     SummarizeInstallationUsageSortByEnum = "approximateApplicationCount"
	SummarizeInstallationUsageSortByApproximatemanagedinstancecount SummarizeInstallationUsageSortByEnum = "approximateManagedInstanceCount"
	SummarizeInstallationUsageSortByOsname                          SummarizeInstallationUsageSortByEnum = "osName"
)

var mappingSummarizeInstallationUsageSortBy = map[string]SummarizeInstallationUsageSortByEnum{
	"jreDistribution":                 SummarizeInstallationUsageSortByJredistribution,
	"jreVendor":                       SummarizeInstallationUsageSortByJrevendor,
	"jreVersion":                      SummarizeInstallationUsageSortByJreversion,
	"path":                            SummarizeInstallationUsageSortByPath,
	"timeFirstSeen":                   SummarizeInstallationUsageSortByTimefirstseen,
	"timeLastSeen":                    SummarizeInstallationUsageSortByTimelastseen,
	"approximateApplicationCount":     SummarizeInstallationUsageSortByApproximateapplicationcount,
	"approximateManagedInstanceCount": SummarizeInstallationUsageSortByApproximatemanagedinstancecount,
	"osName":                          SummarizeInstallationUsageSortByOsname,
}

// GetSummarizeInstallationUsageSortByEnumValues Enumerates the set of values for SummarizeInstallationUsageSortByEnum
func GetSummarizeInstallationUsageSortByEnumValues() []SummarizeInstallationUsageSortByEnum {
	values := make([]SummarizeInstallationUsageSortByEnum, 0)
	for _, v := range mappingSummarizeInstallationUsageSortBy {
		values = append(values, v)
	}
	return values
}
