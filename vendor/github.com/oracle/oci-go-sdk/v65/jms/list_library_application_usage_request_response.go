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

// ListLibraryApplicationUsageRequest wrapper for the ListLibraryApplicationUsage operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListLibraryApplicationUsage.go.html to see an example of how to use ListLibraryApplicationUsageRequest.
type ListLibraryApplicationUsageRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The unique identifier of a Java library.
	LibraryKey *string `mandatory:"true" contributesTo:"path" name:"libraryKey"`

	// The Fleet-unique identifier of the application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The Fleet-unique identifier of the managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The name of the application.
	ApplicationName *string `mandatory:"false" contributesTo:"query" name:"applicationName"`

	// Filter the list with application name contains the given value.
	ApplicationNameContains *string `mandatory:"false" contributesTo:"query" name:"applicationNameContains"`

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
	SortBy ListLibraryApplicationUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListLibraryApplicationUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLibraryApplicationUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLibraryApplicationUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLibraryApplicationUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLibraryApplicationUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLibraryApplicationUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLibraryApplicationUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLibraryApplicationUsageSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLibraryApplicationUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLibraryApplicationUsageSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLibraryApplicationUsageResponse wrapper for the ListLibraryApplicationUsage operation
type ListLibraryApplicationUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LibraryApplicationUsageCollection instances
	LibraryApplicationUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLibraryApplicationUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLibraryApplicationUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLibraryApplicationUsageSortByEnum Enum with underlying type: string
type ListLibraryApplicationUsageSortByEnum string

// Set of constants representing the allowable values for ListLibraryApplicationUsageSortByEnum
const (
	ListLibraryApplicationUsageSortByApplicationname         ListLibraryApplicationUsageSortByEnum = "applicationName"
	ListLibraryApplicationUsageSortByManagedinstancecount    ListLibraryApplicationUsageSortByEnum = "managedInstanceCount"
	ListLibraryApplicationUsageSortByLastdetecteddynamically ListLibraryApplicationUsageSortByEnum = "lastDetectedDynamically"
	ListLibraryApplicationUsageSortByFirstseeninclasspath    ListLibraryApplicationUsageSortByEnum = "firstSeenInClasspath"
	ListLibraryApplicationUsageSortByLastseeninclasspath     ListLibraryApplicationUsageSortByEnum = "lastSeenInClasspath"
)

var mappingListLibraryApplicationUsageSortByEnum = map[string]ListLibraryApplicationUsageSortByEnum{
	"applicationName":         ListLibraryApplicationUsageSortByApplicationname,
	"managedInstanceCount":    ListLibraryApplicationUsageSortByManagedinstancecount,
	"lastDetectedDynamically": ListLibraryApplicationUsageSortByLastdetecteddynamically,
	"firstSeenInClasspath":    ListLibraryApplicationUsageSortByFirstseeninclasspath,
	"lastSeenInClasspath":     ListLibraryApplicationUsageSortByLastseeninclasspath,
}

var mappingListLibraryApplicationUsageSortByEnumLowerCase = map[string]ListLibraryApplicationUsageSortByEnum{
	"applicationname":         ListLibraryApplicationUsageSortByApplicationname,
	"managedinstancecount":    ListLibraryApplicationUsageSortByManagedinstancecount,
	"lastdetecteddynamically": ListLibraryApplicationUsageSortByLastdetecteddynamically,
	"firstseeninclasspath":    ListLibraryApplicationUsageSortByFirstseeninclasspath,
	"lastseeninclasspath":     ListLibraryApplicationUsageSortByLastseeninclasspath,
}

// GetListLibraryApplicationUsageSortByEnumValues Enumerates the set of values for ListLibraryApplicationUsageSortByEnum
func GetListLibraryApplicationUsageSortByEnumValues() []ListLibraryApplicationUsageSortByEnum {
	values := make([]ListLibraryApplicationUsageSortByEnum, 0)
	for _, v := range mappingListLibraryApplicationUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryApplicationUsageSortByEnumStringValues Enumerates the set of values in String for ListLibraryApplicationUsageSortByEnum
func GetListLibraryApplicationUsageSortByEnumStringValues() []string {
	return []string{
		"applicationName",
		"managedInstanceCount",
		"lastDetectedDynamically",
		"firstSeenInClasspath",
		"lastSeenInClasspath",
	}
}

// GetMappingListLibraryApplicationUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryApplicationUsageSortByEnum(val string) (ListLibraryApplicationUsageSortByEnum, bool) {
	enum, ok := mappingListLibraryApplicationUsageSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLibraryApplicationUsageSortOrderEnum Enum with underlying type: string
type ListLibraryApplicationUsageSortOrderEnum string

// Set of constants representing the allowable values for ListLibraryApplicationUsageSortOrderEnum
const (
	ListLibraryApplicationUsageSortOrderAsc  ListLibraryApplicationUsageSortOrderEnum = "ASC"
	ListLibraryApplicationUsageSortOrderDesc ListLibraryApplicationUsageSortOrderEnum = "DESC"
)

var mappingListLibraryApplicationUsageSortOrderEnum = map[string]ListLibraryApplicationUsageSortOrderEnum{
	"ASC":  ListLibraryApplicationUsageSortOrderAsc,
	"DESC": ListLibraryApplicationUsageSortOrderDesc,
}

var mappingListLibraryApplicationUsageSortOrderEnumLowerCase = map[string]ListLibraryApplicationUsageSortOrderEnum{
	"asc":  ListLibraryApplicationUsageSortOrderAsc,
	"desc": ListLibraryApplicationUsageSortOrderDesc,
}

// GetListLibraryApplicationUsageSortOrderEnumValues Enumerates the set of values for ListLibraryApplicationUsageSortOrderEnum
func GetListLibraryApplicationUsageSortOrderEnumValues() []ListLibraryApplicationUsageSortOrderEnum {
	values := make([]ListLibraryApplicationUsageSortOrderEnum, 0)
	for _, v := range mappingListLibraryApplicationUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryApplicationUsageSortOrderEnumStringValues Enumerates the set of values in String for ListLibraryApplicationUsageSortOrderEnum
func GetListLibraryApplicationUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLibraryApplicationUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryApplicationUsageSortOrderEnum(val string) (ListLibraryApplicationUsageSortOrderEnum, bool) {
	enum, ok := mappingListLibraryApplicationUsageSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
