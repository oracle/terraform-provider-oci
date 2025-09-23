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

// ListLibraryManagedInstanceUsageRequest wrapper for the ListLibraryManagedInstanceUsage operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListLibraryManagedInstanceUsage.go.html to see an example of how to use ListLibraryManagedInstanceUsageRequest.
type ListLibraryManagedInstanceUsageRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The unique identifier of a Java library.
	LibraryKey *string `mandatory:"true" contributesTo:"path" name:"libraryKey"`

	// The Fleet-unique identifier of the application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The host OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
	HostName *string `mandatory:"false" contributesTo:"query" name:"hostName"`

	// Filter the list with hostname contains the given value.
	HostnameContains *string `mandatory:"false" contributesTo:"query" name:"hostnameContains"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort library application usage summaries.  Only one sort order may be provided.
	// If no value is specified _lastSeenInClasspath_ is default.
	SortBy ListLibraryManagedInstanceUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListLibraryManagedInstanceUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLibraryManagedInstanceUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLibraryManagedInstanceUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLibraryManagedInstanceUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLibraryManagedInstanceUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLibraryManagedInstanceUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLibraryManagedInstanceUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLibraryManagedInstanceUsageSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLibraryManagedInstanceUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLibraryManagedInstanceUsageSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLibraryManagedInstanceUsageResponse wrapper for the ListLibraryManagedInstanceUsage operation
type ListLibraryManagedInstanceUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LibraryManagedInstanceUsageCollection instances
	LibraryManagedInstanceUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLibraryManagedInstanceUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLibraryManagedInstanceUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLibraryManagedInstanceUsageSortByEnum Enum with underlying type: string
type ListLibraryManagedInstanceUsageSortByEnum string

// Set of constants representing the allowable values for ListLibraryManagedInstanceUsageSortByEnum
const (
	ListLibraryManagedInstanceUsageSortByHostname                ListLibraryManagedInstanceUsageSortByEnum = "hostname"
	ListLibraryManagedInstanceUsageSortByApplicationcount        ListLibraryManagedInstanceUsageSortByEnum = "applicationCount"
	ListLibraryManagedInstanceUsageSortByLastdetecteddynamically ListLibraryManagedInstanceUsageSortByEnum = "lastDetectedDynamically"
	ListLibraryManagedInstanceUsageSortByFirstseeninclasspath    ListLibraryManagedInstanceUsageSortByEnum = "firstSeenInClasspath"
	ListLibraryManagedInstanceUsageSortByLastseeninclasspath     ListLibraryManagedInstanceUsageSortByEnum = "lastSeenInClasspath"
)

var mappingListLibraryManagedInstanceUsageSortByEnum = map[string]ListLibraryManagedInstanceUsageSortByEnum{
	"hostname":                ListLibraryManagedInstanceUsageSortByHostname,
	"applicationCount":        ListLibraryManagedInstanceUsageSortByApplicationcount,
	"lastDetectedDynamically": ListLibraryManagedInstanceUsageSortByLastdetecteddynamically,
	"firstSeenInClasspath":    ListLibraryManagedInstanceUsageSortByFirstseeninclasspath,
	"lastSeenInClasspath":     ListLibraryManagedInstanceUsageSortByLastseeninclasspath,
}

var mappingListLibraryManagedInstanceUsageSortByEnumLowerCase = map[string]ListLibraryManagedInstanceUsageSortByEnum{
	"hostname":                ListLibraryManagedInstanceUsageSortByHostname,
	"applicationcount":        ListLibraryManagedInstanceUsageSortByApplicationcount,
	"lastdetecteddynamically": ListLibraryManagedInstanceUsageSortByLastdetecteddynamically,
	"firstseeninclasspath":    ListLibraryManagedInstanceUsageSortByFirstseeninclasspath,
	"lastseeninclasspath":     ListLibraryManagedInstanceUsageSortByLastseeninclasspath,
}

// GetListLibraryManagedInstanceUsageSortByEnumValues Enumerates the set of values for ListLibraryManagedInstanceUsageSortByEnum
func GetListLibraryManagedInstanceUsageSortByEnumValues() []ListLibraryManagedInstanceUsageSortByEnum {
	values := make([]ListLibraryManagedInstanceUsageSortByEnum, 0)
	for _, v := range mappingListLibraryManagedInstanceUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryManagedInstanceUsageSortByEnumStringValues Enumerates the set of values in String for ListLibraryManagedInstanceUsageSortByEnum
func GetListLibraryManagedInstanceUsageSortByEnumStringValues() []string {
	return []string{
		"hostname",
		"applicationCount",
		"lastDetectedDynamically",
		"firstSeenInClasspath",
		"lastSeenInClasspath",
	}
}

// GetMappingListLibraryManagedInstanceUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryManagedInstanceUsageSortByEnum(val string) (ListLibraryManagedInstanceUsageSortByEnum, bool) {
	enum, ok := mappingListLibraryManagedInstanceUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLibraryManagedInstanceUsageSortOrderEnum Enum with underlying type: string
type ListLibraryManagedInstanceUsageSortOrderEnum string

// Set of constants representing the allowable values for ListLibraryManagedInstanceUsageSortOrderEnum
const (
	ListLibraryManagedInstanceUsageSortOrderAsc  ListLibraryManagedInstanceUsageSortOrderEnum = "ASC"
	ListLibraryManagedInstanceUsageSortOrderDesc ListLibraryManagedInstanceUsageSortOrderEnum = "DESC"
)

var mappingListLibraryManagedInstanceUsageSortOrderEnum = map[string]ListLibraryManagedInstanceUsageSortOrderEnum{
	"ASC":  ListLibraryManagedInstanceUsageSortOrderAsc,
	"DESC": ListLibraryManagedInstanceUsageSortOrderDesc,
}

var mappingListLibraryManagedInstanceUsageSortOrderEnumLowerCase = map[string]ListLibraryManagedInstanceUsageSortOrderEnum{
	"asc":  ListLibraryManagedInstanceUsageSortOrderAsc,
	"desc": ListLibraryManagedInstanceUsageSortOrderDesc,
}

// GetListLibraryManagedInstanceUsageSortOrderEnumValues Enumerates the set of values for ListLibraryManagedInstanceUsageSortOrderEnum
func GetListLibraryManagedInstanceUsageSortOrderEnumValues() []ListLibraryManagedInstanceUsageSortOrderEnum {
	values := make([]ListLibraryManagedInstanceUsageSortOrderEnum, 0)
	for _, v := range mappingListLibraryManagedInstanceUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryManagedInstanceUsageSortOrderEnumStringValues Enumerates the set of values in String for ListLibraryManagedInstanceUsageSortOrderEnum
func GetListLibraryManagedInstanceUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLibraryManagedInstanceUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryManagedInstanceUsageSortOrderEnum(val string) (ListLibraryManagedInstanceUsageSortOrderEnum, bool) {
	enum, ok := mappingListLibraryManagedInstanceUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
