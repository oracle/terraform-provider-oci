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

// ListUncorrelatedPackageApplicationUsageRequest wrapper for the ListUncorrelatedPackageApplicationUsage operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListUncorrelatedPackageApplicationUsage.go.html to see an example of how to use ListUncorrelatedPackageApplicationUsageRequest.
type ListUncorrelatedPackageApplicationUsageRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The unique identifier of a Java package.
	PackageName *string `mandatory:"true" contributesTo:"path" name:"packageName"`

	// The Fleet-unique identifier of the application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort detection events.  Only one sort order may be provided.
	// If no value is specified _dynamicallyLastDetected_ is default.
	SortBy ListUncorrelatedPackageApplicationUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListUncorrelatedPackageApplicationUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUncorrelatedPackageApplicationUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUncorrelatedPackageApplicationUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUncorrelatedPackageApplicationUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUncorrelatedPackageApplicationUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUncorrelatedPackageApplicationUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUncorrelatedPackageApplicationUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUncorrelatedPackageApplicationUsageSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUncorrelatedPackageApplicationUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUncorrelatedPackageApplicationUsageSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUncorrelatedPackageApplicationUsageResponse wrapper for the ListUncorrelatedPackageApplicationUsage operation
type ListUncorrelatedPackageApplicationUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UncorrelatedPackageApplicationUsageCollection instances
	UncorrelatedPackageApplicationUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUncorrelatedPackageApplicationUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUncorrelatedPackageApplicationUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUncorrelatedPackageApplicationUsageSortByEnum Enum with underlying type: string
type ListUncorrelatedPackageApplicationUsageSortByEnum string

// Set of constants representing the allowable values for ListUncorrelatedPackageApplicationUsageSortByEnum
const (
	ListUncorrelatedPackageApplicationUsageSortByApplicationname         ListUncorrelatedPackageApplicationUsageSortByEnum = "applicationName"
	ListUncorrelatedPackageApplicationUsageSortByManagedinstancecount    ListUncorrelatedPackageApplicationUsageSortByEnum = "managedInstanceCount"
	ListUncorrelatedPackageApplicationUsageSortByLastdetecteddynamically ListUncorrelatedPackageApplicationUsageSortByEnum = "lastDetectedDynamically"
)

var mappingListUncorrelatedPackageApplicationUsageSortByEnum = map[string]ListUncorrelatedPackageApplicationUsageSortByEnum{
	"applicationName":         ListUncorrelatedPackageApplicationUsageSortByApplicationname,
	"managedInstanceCount":    ListUncorrelatedPackageApplicationUsageSortByManagedinstancecount,
	"lastDetectedDynamically": ListUncorrelatedPackageApplicationUsageSortByLastdetecteddynamically,
}

var mappingListUncorrelatedPackageApplicationUsageSortByEnumLowerCase = map[string]ListUncorrelatedPackageApplicationUsageSortByEnum{
	"applicationname":         ListUncorrelatedPackageApplicationUsageSortByApplicationname,
	"managedinstancecount":    ListUncorrelatedPackageApplicationUsageSortByManagedinstancecount,
	"lastdetecteddynamically": ListUncorrelatedPackageApplicationUsageSortByLastdetecteddynamically,
}

// GetListUncorrelatedPackageApplicationUsageSortByEnumValues Enumerates the set of values for ListUncorrelatedPackageApplicationUsageSortByEnum
func GetListUncorrelatedPackageApplicationUsageSortByEnumValues() []ListUncorrelatedPackageApplicationUsageSortByEnum {
	values := make([]ListUncorrelatedPackageApplicationUsageSortByEnum, 0)
	for _, v := range mappingListUncorrelatedPackageApplicationUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUncorrelatedPackageApplicationUsageSortByEnumStringValues Enumerates the set of values in String for ListUncorrelatedPackageApplicationUsageSortByEnum
func GetListUncorrelatedPackageApplicationUsageSortByEnumStringValues() []string {
	return []string{
		"applicationName",
		"managedInstanceCount",
		"lastDetectedDynamically",
	}
}

// GetMappingListUncorrelatedPackageApplicationUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUncorrelatedPackageApplicationUsageSortByEnum(val string) (ListUncorrelatedPackageApplicationUsageSortByEnum, bool) {
	enum, ok := mappingListUncorrelatedPackageApplicationUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUncorrelatedPackageApplicationUsageSortOrderEnum Enum with underlying type: string
type ListUncorrelatedPackageApplicationUsageSortOrderEnum string

// Set of constants representing the allowable values for ListUncorrelatedPackageApplicationUsageSortOrderEnum
const (
	ListUncorrelatedPackageApplicationUsageSortOrderAsc  ListUncorrelatedPackageApplicationUsageSortOrderEnum = "ASC"
	ListUncorrelatedPackageApplicationUsageSortOrderDesc ListUncorrelatedPackageApplicationUsageSortOrderEnum = "DESC"
)

var mappingListUncorrelatedPackageApplicationUsageSortOrderEnum = map[string]ListUncorrelatedPackageApplicationUsageSortOrderEnum{
	"ASC":  ListUncorrelatedPackageApplicationUsageSortOrderAsc,
	"DESC": ListUncorrelatedPackageApplicationUsageSortOrderDesc,
}

var mappingListUncorrelatedPackageApplicationUsageSortOrderEnumLowerCase = map[string]ListUncorrelatedPackageApplicationUsageSortOrderEnum{
	"asc":  ListUncorrelatedPackageApplicationUsageSortOrderAsc,
	"desc": ListUncorrelatedPackageApplicationUsageSortOrderDesc,
}

// GetListUncorrelatedPackageApplicationUsageSortOrderEnumValues Enumerates the set of values for ListUncorrelatedPackageApplicationUsageSortOrderEnum
func GetListUncorrelatedPackageApplicationUsageSortOrderEnumValues() []ListUncorrelatedPackageApplicationUsageSortOrderEnum {
	values := make([]ListUncorrelatedPackageApplicationUsageSortOrderEnum, 0)
	for _, v := range mappingListUncorrelatedPackageApplicationUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUncorrelatedPackageApplicationUsageSortOrderEnumStringValues Enumerates the set of values in String for ListUncorrelatedPackageApplicationUsageSortOrderEnum
func GetListUncorrelatedPackageApplicationUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUncorrelatedPackageApplicationUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUncorrelatedPackageApplicationUsageSortOrderEnum(val string) (ListUncorrelatedPackageApplicationUsageSortOrderEnum, bool) {
	enum, ok := mappingListUncorrelatedPackageApplicationUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
