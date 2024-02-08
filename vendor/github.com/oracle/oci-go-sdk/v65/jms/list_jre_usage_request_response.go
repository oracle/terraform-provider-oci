// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJreUsageRequest wrapper for the ListJreUsage operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJreUsage.go.html to see an example of how to use ListJreUsageRequest.
type ListJreUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The host OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the managed instance.
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// The Fleet-unique identifier of the application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The name of the application.
	ApplicationName *string `mandatory:"false" contributesTo:"query" name:"applicationName"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJreUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort JRE usages. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, and _version_ is **descending**.
	// Default order for _timeFirstSeen_, _timeLastSeen_, _version_, _approximateInstallationCount_,
	// _approximateApplicationCount_ and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _distribution_, _vendor_, and _osName_ is **ascending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy ListJreUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJreUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJreUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJreUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJreUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJreUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJreUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJreUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJreUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJreUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJreUsageResponse wrapper for the ListJreUsage operation
type ListJreUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JreUsageCollection instances
	JreUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJreUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJreUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJreUsageSortOrderEnum Enum with underlying type: string
type ListJreUsageSortOrderEnum string

// Set of constants representing the allowable values for ListJreUsageSortOrderEnum
const (
	ListJreUsageSortOrderAsc  ListJreUsageSortOrderEnum = "ASC"
	ListJreUsageSortOrderDesc ListJreUsageSortOrderEnum = "DESC"
)

var mappingListJreUsageSortOrderEnum = map[string]ListJreUsageSortOrderEnum{
	"ASC":  ListJreUsageSortOrderAsc,
	"DESC": ListJreUsageSortOrderDesc,
}

var mappingListJreUsageSortOrderEnumLowerCase = map[string]ListJreUsageSortOrderEnum{
	"asc":  ListJreUsageSortOrderAsc,
	"desc": ListJreUsageSortOrderDesc,
}

// GetListJreUsageSortOrderEnumValues Enumerates the set of values for ListJreUsageSortOrderEnum
func GetListJreUsageSortOrderEnumValues() []ListJreUsageSortOrderEnum {
	values := make([]ListJreUsageSortOrderEnum, 0)
	for _, v := range mappingListJreUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJreUsageSortOrderEnumStringValues Enumerates the set of values in String for ListJreUsageSortOrderEnum
func GetListJreUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJreUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJreUsageSortOrderEnum(val string) (ListJreUsageSortOrderEnum, bool) {
	enum, ok := mappingListJreUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJreUsageSortByEnum Enum with underlying type: string
type ListJreUsageSortByEnum string

// Set of constants representing the allowable values for ListJreUsageSortByEnum
const (
	ListJreUsageSortByDistribution                    ListJreUsageSortByEnum = "distribution"
	ListJreUsageSortByTimefirstseen                   ListJreUsageSortByEnum = "timeFirstSeen"
	ListJreUsageSortByTimelastseen                    ListJreUsageSortByEnum = "timeLastSeen"
	ListJreUsageSortByVendor                          ListJreUsageSortByEnum = "vendor"
	ListJreUsageSortByVersion                         ListJreUsageSortByEnum = "version"
	ListJreUsageSortByApproximateinstallationcount    ListJreUsageSortByEnum = "approximateInstallationCount"
	ListJreUsageSortByApproximateapplicationcount     ListJreUsageSortByEnum = "approximateApplicationCount"
	ListJreUsageSortByApproximatemanagedinstancecount ListJreUsageSortByEnum = "approximateManagedInstanceCount"
	ListJreUsageSortByOsname                          ListJreUsageSortByEnum = "osName"
	ListJreUsageSortBySecuritystatus                  ListJreUsageSortByEnum = "securityStatus"
)

var mappingListJreUsageSortByEnum = map[string]ListJreUsageSortByEnum{
	"distribution":                    ListJreUsageSortByDistribution,
	"timeFirstSeen":                   ListJreUsageSortByTimefirstseen,
	"timeLastSeen":                    ListJreUsageSortByTimelastseen,
	"vendor":                          ListJreUsageSortByVendor,
	"version":                         ListJreUsageSortByVersion,
	"approximateInstallationCount":    ListJreUsageSortByApproximateinstallationcount,
	"approximateApplicationCount":     ListJreUsageSortByApproximateapplicationcount,
	"approximateManagedInstanceCount": ListJreUsageSortByApproximatemanagedinstancecount,
	"osName":                          ListJreUsageSortByOsname,
	"securityStatus":                  ListJreUsageSortBySecuritystatus,
}

var mappingListJreUsageSortByEnumLowerCase = map[string]ListJreUsageSortByEnum{
	"distribution":                    ListJreUsageSortByDistribution,
	"timefirstseen":                   ListJreUsageSortByTimefirstseen,
	"timelastseen":                    ListJreUsageSortByTimelastseen,
	"vendor":                          ListJreUsageSortByVendor,
	"version":                         ListJreUsageSortByVersion,
	"approximateinstallationcount":    ListJreUsageSortByApproximateinstallationcount,
	"approximateapplicationcount":     ListJreUsageSortByApproximateapplicationcount,
	"approximatemanagedinstancecount": ListJreUsageSortByApproximatemanagedinstancecount,
	"osname":                          ListJreUsageSortByOsname,
	"securitystatus":                  ListJreUsageSortBySecuritystatus,
}

// GetListJreUsageSortByEnumValues Enumerates the set of values for ListJreUsageSortByEnum
func GetListJreUsageSortByEnumValues() []ListJreUsageSortByEnum {
	values := make([]ListJreUsageSortByEnum, 0)
	for _, v := range mappingListJreUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJreUsageSortByEnumStringValues Enumerates the set of values in String for ListJreUsageSortByEnum
func GetListJreUsageSortByEnumStringValues() []string {
	return []string{
		"distribution",
		"timeFirstSeen",
		"timeLastSeen",
		"vendor",
		"version",
		"approximateInstallationCount",
		"approximateApplicationCount",
		"approximateManagedInstanceCount",
		"osName",
		"securityStatus",
	}
}

// GetMappingListJreUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJreUsageSortByEnum(val string) (ListJreUsageSortByEnum, bool) {
	enum, ok := mappingListJreUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
