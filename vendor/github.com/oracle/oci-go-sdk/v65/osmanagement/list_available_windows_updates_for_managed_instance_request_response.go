// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableWindowsUpdatesForManagedInstanceRequest wrapper for the ListAvailableWindowsUpdatesForManagedInstance operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListAvailableWindowsUpdatesForManagedInstance.go.html to see an example of how to use ListAvailableWindowsUpdatesForManagedInstanceRequest.
type ListAvailableWindowsUpdatesForManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListAvailableWindowsUpdatesForManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicator of whether the update can be installed using OSMS.
	IsEligibleForInstallation ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum `mandatory:"false" contributesTo:"query" name:"isEligibleForInstallation" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableWindowsUpdatesForManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableWindowsUpdatesForManagedInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableWindowsUpdatesForManagedInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableWindowsUpdatesForManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableWindowsUpdatesForManagedInstanceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableWindowsUpdatesForManagedInstanceSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableWindowsUpdatesForManagedInstanceSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum(string(request.IsEligibleForInstallation)); !ok && request.IsEligibleForInstallation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsEligibleForInstallation: %s. Supported values are: %s.", request.IsEligibleForInstallation, strings.Join(GetListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableWindowsUpdatesForManagedInstanceResponse wrapper for the ListAvailableWindowsUpdatesForManagedInstance operation
type ListAvailableWindowsUpdatesForManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailableWindowsUpdateSummary instances
	Items []AvailableWindowsUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableWindowsUpdatesForManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableWindowsUpdatesForManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum Enum with underlying type: string
type ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum
const (
	ListAvailableWindowsUpdatesForManagedInstanceSortOrderAsc  ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum = "ASC"
	ListAvailableWindowsUpdatesForManagedInstanceSortOrderDesc ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum = "DESC"
)

var mappingListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum = map[string]ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum{
	"ASC":  ListAvailableWindowsUpdatesForManagedInstanceSortOrderAsc,
	"DESC": ListAvailableWindowsUpdatesForManagedInstanceSortOrderDesc,
}

var mappingListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumLowerCase = map[string]ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum{
	"asc":  ListAvailableWindowsUpdatesForManagedInstanceSortOrderAsc,
	"desc": ListAvailableWindowsUpdatesForManagedInstanceSortOrderDesc,
}

// GetListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumValues Enumerates the set of values for ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum
func GetListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumValues() []ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum {
	values := make([]ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum
func GetListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum(val string) (ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum, bool) {
	enum, ok := mappingListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableWindowsUpdatesForManagedInstanceSortByEnum Enum with underlying type: string
type ListAvailableWindowsUpdatesForManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListAvailableWindowsUpdatesForManagedInstanceSortByEnum
const (
	ListAvailableWindowsUpdatesForManagedInstanceSortByTimecreated ListAvailableWindowsUpdatesForManagedInstanceSortByEnum = "TIMECREATED"
	ListAvailableWindowsUpdatesForManagedInstanceSortByDisplayname ListAvailableWindowsUpdatesForManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListAvailableWindowsUpdatesForManagedInstanceSortByEnum = map[string]ListAvailableWindowsUpdatesForManagedInstanceSortByEnum{
	"TIMECREATED": ListAvailableWindowsUpdatesForManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListAvailableWindowsUpdatesForManagedInstanceSortByDisplayname,
}

var mappingListAvailableWindowsUpdatesForManagedInstanceSortByEnumLowerCase = map[string]ListAvailableWindowsUpdatesForManagedInstanceSortByEnum{
	"timecreated": ListAvailableWindowsUpdatesForManagedInstanceSortByTimecreated,
	"displayname": ListAvailableWindowsUpdatesForManagedInstanceSortByDisplayname,
}

// GetListAvailableWindowsUpdatesForManagedInstanceSortByEnumValues Enumerates the set of values for ListAvailableWindowsUpdatesForManagedInstanceSortByEnum
func GetListAvailableWindowsUpdatesForManagedInstanceSortByEnumValues() []ListAvailableWindowsUpdatesForManagedInstanceSortByEnum {
	values := make([]ListAvailableWindowsUpdatesForManagedInstanceSortByEnum, 0)
	for _, v := range mappingListAvailableWindowsUpdatesForManagedInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableWindowsUpdatesForManagedInstanceSortByEnumStringValues Enumerates the set of values in String for ListAvailableWindowsUpdatesForManagedInstanceSortByEnum
func GetListAvailableWindowsUpdatesForManagedInstanceSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAvailableWindowsUpdatesForManagedInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableWindowsUpdatesForManagedInstanceSortByEnum(val string) (ListAvailableWindowsUpdatesForManagedInstanceSortByEnum, bool) {
	enum, ok := mappingListAvailableWindowsUpdatesForManagedInstanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum Enum with underlying type: string
type ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum string

// Set of constants representing the allowable values for ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum
const (
	ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationInstallable    ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum = "INSTALLABLE"
	ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationNotInstallable ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum = "NOT_INSTALLABLE"
	ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationUnknown        ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum = "UNKNOWN"
)

var mappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum = map[string]ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum{
	"INSTALLABLE":     ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationInstallable,
	"NOT_INSTALLABLE": ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationNotInstallable,
	"UNKNOWN":         ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationUnknown,
}

var mappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumLowerCase = map[string]ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum{
	"installable":     ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationInstallable,
	"not_installable": ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationNotInstallable,
	"unknown":         ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationUnknown,
}

// GetListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumValues Enumerates the set of values for ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum
func GetListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumValues() []ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum {
	values := make([]ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum, 0)
	for _, v := range mappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumStringValues Enumerates the set of values in String for ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum
func GetListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumStringValues() []string {
	return []string{
		"INSTALLABLE",
		"NOT_INSTALLABLE",
		"UNKNOWN",
	}
}

// GetMappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum(val string) (ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum, bool) {
	enum, ok := mappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
