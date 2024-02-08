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

// ListModuleStreamProfilesOnManagedInstanceRequest wrapper for the ListModuleStreamProfilesOnManagedInstance operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListModuleStreamProfilesOnManagedInstance.go.html to see an example of how to use ListModuleStreamProfilesOnManagedInstanceRequest.
type ListModuleStreamProfilesOnManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The name of a module.  This parameter is required if a
	// streamName is specified.
	ModuleName *string `mandatory:"false" contributesTo:"query" name:"moduleName"`

	// The name of the stream of the containing module.  This parameter
	// is required if a profileName is specified.
	StreamName *string `mandatory:"false" contributesTo:"query" name:"streamName"`

	// The name of the profile of the containing module stream
	ProfileName *string `mandatory:"false" contributesTo:"query" name:"profileName"`

	// The status of the profile.
	// A profile with the "INSTALLED" status indicates that the
	// profile has been installed.
	// A profile with the "AVAILABLE" status indicates that the
	// profile is not installed, but can be.
	ProfileStatus ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum `mandatory:"false" contributesTo:"query" name:"profileStatus" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListModuleStreamProfilesOnManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListModuleStreamProfilesOnManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModuleStreamProfilesOnManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModuleStreamProfilesOnManagedInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModuleStreamProfilesOnManagedInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModuleStreamProfilesOnManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModuleStreamProfilesOnManagedInstanceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModuleStreamProfilesOnManagedInstanceProfileStatusEnum(string(request.ProfileStatus)); !ok && request.ProfileStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProfileStatus: %s. Supported values are: %s.", request.ProfileStatus, strings.Join(GetListModuleStreamProfilesOnManagedInstanceProfileStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModuleStreamProfilesOnManagedInstanceSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModuleStreamProfilesOnManagedInstanceSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModuleStreamProfilesOnManagedInstanceSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModuleStreamProfilesOnManagedInstanceSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModuleStreamProfilesOnManagedInstanceResponse wrapper for the ListModuleStreamProfilesOnManagedInstance operation
type ListModuleStreamProfilesOnManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModuleStreamProfileOnManagedInstanceSummary instances
	Items []ModuleStreamProfileOnManagedInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query the status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListModuleStreamProfilesOnManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModuleStreamProfilesOnManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum Enum with underlying type: string
type ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum string

// Set of constants representing the allowable values for ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum
const (
	ListModuleStreamProfilesOnManagedInstanceProfileStatusInstalled ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum = "INSTALLED"
	ListModuleStreamProfilesOnManagedInstanceProfileStatusAvailable ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum = "AVAILABLE"
)

var mappingListModuleStreamProfilesOnManagedInstanceProfileStatusEnum = map[string]ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum{
	"INSTALLED": ListModuleStreamProfilesOnManagedInstanceProfileStatusInstalled,
	"AVAILABLE": ListModuleStreamProfilesOnManagedInstanceProfileStatusAvailable,
}

var mappingListModuleStreamProfilesOnManagedInstanceProfileStatusEnumLowerCase = map[string]ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum{
	"installed": ListModuleStreamProfilesOnManagedInstanceProfileStatusInstalled,
	"available": ListModuleStreamProfilesOnManagedInstanceProfileStatusAvailable,
}

// GetListModuleStreamProfilesOnManagedInstanceProfileStatusEnumValues Enumerates the set of values for ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum
func GetListModuleStreamProfilesOnManagedInstanceProfileStatusEnumValues() []ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum {
	values := make([]ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum, 0)
	for _, v := range mappingListModuleStreamProfilesOnManagedInstanceProfileStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamProfilesOnManagedInstanceProfileStatusEnumStringValues Enumerates the set of values in String for ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum
func GetListModuleStreamProfilesOnManagedInstanceProfileStatusEnumStringValues() []string {
	return []string{
		"INSTALLED",
		"AVAILABLE",
	}
}

// GetMappingListModuleStreamProfilesOnManagedInstanceProfileStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamProfilesOnManagedInstanceProfileStatusEnum(val string) (ListModuleStreamProfilesOnManagedInstanceProfileStatusEnum, bool) {
	enum, ok := mappingListModuleStreamProfilesOnManagedInstanceProfileStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModuleStreamProfilesOnManagedInstanceSortOrderEnum Enum with underlying type: string
type ListModuleStreamProfilesOnManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListModuleStreamProfilesOnManagedInstanceSortOrderEnum
const (
	ListModuleStreamProfilesOnManagedInstanceSortOrderAsc  ListModuleStreamProfilesOnManagedInstanceSortOrderEnum = "ASC"
	ListModuleStreamProfilesOnManagedInstanceSortOrderDesc ListModuleStreamProfilesOnManagedInstanceSortOrderEnum = "DESC"
)

var mappingListModuleStreamProfilesOnManagedInstanceSortOrderEnum = map[string]ListModuleStreamProfilesOnManagedInstanceSortOrderEnum{
	"ASC":  ListModuleStreamProfilesOnManagedInstanceSortOrderAsc,
	"DESC": ListModuleStreamProfilesOnManagedInstanceSortOrderDesc,
}

var mappingListModuleStreamProfilesOnManagedInstanceSortOrderEnumLowerCase = map[string]ListModuleStreamProfilesOnManagedInstanceSortOrderEnum{
	"asc":  ListModuleStreamProfilesOnManagedInstanceSortOrderAsc,
	"desc": ListModuleStreamProfilesOnManagedInstanceSortOrderDesc,
}

// GetListModuleStreamProfilesOnManagedInstanceSortOrderEnumValues Enumerates the set of values for ListModuleStreamProfilesOnManagedInstanceSortOrderEnum
func GetListModuleStreamProfilesOnManagedInstanceSortOrderEnumValues() []ListModuleStreamProfilesOnManagedInstanceSortOrderEnum {
	values := make([]ListModuleStreamProfilesOnManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListModuleStreamProfilesOnManagedInstanceSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamProfilesOnManagedInstanceSortOrderEnumStringValues Enumerates the set of values in String for ListModuleStreamProfilesOnManagedInstanceSortOrderEnum
func GetListModuleStreamProfilesOnManagedInstanceSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModuleStreamProfilesOnManagedInstanceSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamProfilesOnManagedInstanceSortOrderEnum(val string) (ListModuleStreamProfilesOnManagedInstanceSortOrderEnum, bool) {
	enum, ok := mappingListModuleStreamProfilesOnManagedInstanceSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModuleStreamProfilesOnManagedInstanceSortByEnum Enum with underlying type: string
type ListModuleStreamProfilesOnManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListModuleStreamProfilesOnManagedInstanceSortByEnum
const (
	ListModuleStreamProfilesOnManagedInstanceSortByTimecreated ListModuleStreamProfilesOnManagedInstanceSortByEnum = "TIMECREATED"
	ListModuleStreamProfilesOnManagedInstanceSortByDisplayname ListModuleStreamProfilesOnManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListModuleStreamProfilesOnManagedInstanceSortByEnum = map[string]ListModuleStreamProfilesOnManagedInstanceSortByEnum{
	"TIMECREATED": ListModuleStreamProfilesOnManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListModuleStreamProfilesOnManagedInstanceSortByDisplayname,
}

var mappingListModuleStreamProfilesOnManagedInstanceSortByEnumLowerCase = map[string]ListModuleStreamProfilesOnManagedInstanceSortByEnum{
	"timecreated": ListModuleStreamProfilesOnManagedInstanceSortByTimecreated,
	"displayname": ListModuleStreamProfilesOnManagedInstanceSortByDisplayname,
}

// GetListModuleStreamProfilesOnManagedInstanceSortByEnumValues Enumerates the set of values for ListModuleStreamProfilesOnManagedInstanceSortByEnum
func GetListModuleStreamProfilesOnManagedInstanceSortByEnumValues() []ListModuleStreamProfilesOnManagedInstanceSortByEnum {
	values := make([]ListModuleStreamProfilesOnManagedInstanceSortByEnum, 0)
	for _, v := range mappingListModuleStreamProfilesOnManagedInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamProfilesOnManagedInstanceSortByEnumStringValues Enumerates the set of values in String for ListModuleStreamProfilesOnManagedInstanceSortByEnum
func GetListModuleStreamProfilesOnManagedInstanceSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListModuleStreamProfilesOnManagedInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamProfilesOnManagedInstanceSortByEnum(val string) (ListModuleStreamProfilesOnManagedInstanceSortByEnum, bool) {
	enum, ok := mappingListModuleStreamProfilesOnManagedInstanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
