// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUncorrelatedPackageUsageRequest wrapper for the ListUncorrelatedPackageUsage operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListUncorrelatedPackageUsage.go.html to see an example of how to use ListUncorrelatedPackageUsageRequest.
type ListUncorrelatedPackageUsageRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The unique identifier of a Java package.
	PackageName *string `mandatory:"false" contributesTo:"query" name:"packageName"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The Fleet-unique identifier of the application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListUncorrelatedPackageUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort Java packages.  Only one sort order may be provided.
	// If no value is specified _timeLastSeen_ is default.
	SortBy ListUncorrelatedPackageUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUncorrelatedPackageUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUncorrelatedPackageUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUncorrelatedPackageUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUncorrelatedPackageUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUncorrelatedPackageUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUncorrelatedPackageUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUncorrelatedPackageUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUncorrelatedPackageUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUncorrelatedPackageUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUncorrelatedPackageUsageResponse wrapper for the ListUncorrelatedPackageUsage operation
type ListUncorrelatedPackageUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UncorrelatedPackageUsageCollection instances
	UncorrelatedPackageUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUncorrelatedPackageUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUncorrelatedPackageUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUncorrelatedPackageUsageSortOrderEnum Enum with underlying type: string
type ListUncorrelatedPackageUsageSortOrderEnum string

// Set of constants representing the allowable values for ListUncorrelatedPackageUsageSortOrderEnum
const (
	ListUncorrelatedPackageUsageSortOrderAsc  ListUncorrelatedPackageUsageSortOrderEnum = "ASC"
	ListUncorrelatedPackageUsageSortOrderDesc ListUncorrelatedPackageUsageSortOrderEnum = "DESC"
)

var mappingListUncorrelatedPackageUsageSortOrderEnum = map[string]ListUncorrelatedPackageUsageSortOrderEnum{
	"ASC":  ListUncorrelatedPackageUsageSortOrderAsc,
	"DESC": ListUncorrelatedPackageUsageSortOrderDesc,
}

var mappingListUncorrelatedPackageUsageSortOrderEnumLowerCase = map[string]ListUncorrelatedPackageUsageSortOrderEnum{
	"asc":  ListUncorrelatedPackageUsageSortOrderAsc,
	"desc": ListUncorrelatedPackageUsageSortOrderDesc,
}

// GetListUncorrelatedPackageUsageSortOrderEnumValues Enumerates the set of values for ListUncorrelatedPackageUsageSortOrderEnum
func GetListUncorrelatedPackageUsageSortOrderEnumValues() []ListUncorrelatedPackageUsageSortOrderEnum {
	values := make([]ListUncorrelatedPackageUsageSortOrderEnum, 0)
	for _, v := range mappingListUncorrelatedPackageUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUncorrelatedPackageUsageSortOrderEnumStringValues Enumerates the set of values in String for ListUncorrelatedPackageUsageSortOrderEnum
func GetListUncorrelatedPackageUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUncorrelatedPackageUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUncorrelatedPackageUsageSortOrderEnum(val string) (ListUncorrelatedPackageUsageSortOrderEnum, bool) {
	enum, ok := mappingListUncorrelatedPackageUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUncorrelatedPackageUsageSortByEnum Enum with underlying type: string
type ListUncorrelatedPackageUsageSortByEnum string

// Set of constants representing the allowable values for ListUncorrelatedPackageUsageSortByEnum
const (
	ListUncorrelatedPackageUsageSortByPackagename             ListUncorrelatedPackageUsageSortByEnum = "packageName"
	ListUncorrelatedPackageUsageSortByManagedinstancecount    ListUncorrelatedPackageUsageSortByEnum = "managedInstanceCount"
	ListUncorrelatedPackageUsageSortByApplicationcount        ListUncorrelatedPackageUsageSortByEnum = "applicationCount"
	ListUncorrelatedPackageUsageSortByLastdetecteddynamically ListUncorrelatedPackageUsageSortByEnum = "lastDetectedDynamically"
)

var mappingListUncorrelatedPackageUsageSortByEnum = map[string]ListUncorrelatedPackageUsageSortByEnum{
	"packageName":             ListUncorrelatedPackageUsageSortByPackagename,
	"managedInstanceCount":    ListUncorrelatedPackageUsageSortByManagedinstancecount,
	"applicationCount":        ListUncorrelatedPackageUsageSortByApplicationcount,
	"lastDetectedDynamically": ListUncorrelatedPackageUsageSortByLastdetecteddynamically,
}

var mappingListUncorrelatedPackageUsageSortByEnumLowerCase = map[string]ListUncorrelatedPackageUsageSortByEnum{
	"packagename":             ListUncorrelatedPackageUsageSortByPackagename,
	"managedinstancecount":    ListUncorrelatedPackageUsageSortByManagedinstancecount,
	"applicationcount":        ListUncorrelatedPackageUsageSortByApplicationcount,
	"lastdetecteddynamically": ListUncorrelatedPackageUsageSortByLastdetecteddynamically,
}

// GetListUncorrelatedPackageUsageSortByEnumValues Enumerates the set of values for ListUncorrelatedPackageUsageSortByEnum
func GetListUncorrelatedPackageUsageSortByEnumValues() []ListUncorrelatedPackageUsageSortByEnum {
	values := make([]ListUncorrelatedPackageUsageSortByEnum, 0)
	for _, v := range mappingListUncorrelatedPackageUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUncorrelatedPackageUsageSortByEnumStringValues Enumerates the set of values in String for ListUncorrelatedPackageUsageSortByEnum
func GetListUncorrelatedPackageUsageSortByEnumStringValues() []string {
	return []string{
		"packageName",
		"managedInstanceCount",
		"applicationCount",
		"lastDetectedDynamically",
	}
}

// GetMappingListUncorrelatedPackageUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUncorrelatedPackageUsageSortByEnum(val string) (ListUncorrelatedPackageUsageSortByEnum, bool) {
	enum, ok := mappingListUncorrelatedPackageUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
