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

// ListModuleStreamsOnManagedInstanceRequest wrapper for the ListModuleStreamsOnManagedInstance operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListModuleStreamsOnManagedInstance.go.html to see an example of how to use ListModuleStreamsOnManagedInstanceRequest.
type ListModuleStreamsOnManagedInstanceRequest struct {

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

	// The status of the stream
	// A stream with the "ENABLED" status can be used as a source for installing
	// profiles.  Streams with this status are also "ACTIVE".
	// A stream with the "DISABLED" status cannot be the source for installing
	// profiles.  To install profiles and packages from this stream, it must be
	// enabled.
	// A stream with the "ACTIVE" status can be used as a source for installing
	// profiles.  The packages that comprise the stream are also used when a
	// matching package is installed directly.  In general, a stream can have
	// this status if it is the default stream for the module and no stream has
	// been explicitly enabled.
	StreamStatus ListModuleStreamsOnManagedInstanceStreamStatusEnum `mandatory:"false" contributesTo:"query" name:"streamStatus" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListModuleStreamsOnManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListModuleStreamsOnManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModuleStreamsOnManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModuleStreamsOnManagedInstanceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModuleStreamsOnManagedInstanceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModuleStreamsOnManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModuleStreamsOnManagedInstanceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModuleStreamsOnManagedInstanceStreamStatusEnum(string(request.StreamStatus)); !ok && request.StreamStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StreamStatus: %s. Supported values are: %s.", request.StreamStatus, strings.Join(GetListModuleStreamsOnManagedInstanceStreamStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModuleStreamsOnManagedInstanceSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModuleStreamsOnManagedInstanceSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModuleStreamsOnManagedInstanceSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModuleStreamsOnManagedInstanceSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModuleStreamsOnManagedInstanceResponse wrapper for the ListModuleStreamsOnManagedInstance operation
type ListModuleStreamsOnManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModuleStreamOnManagedInstanceSummary instances
	Items []ModuleStreamOnManagedInstanceSummary `presentIn:"body"`

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

func (response ListModuleStreamsOnManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModuleStreamsOnManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModuleStreamsOnManagedInstanceStreamStatusEnum Enum with underlying type: string
type ListModuleStreamsOnManagedInstanceStreamStatusEnum string

// Set of constants representing the allowable values for ListModuleStreamsOnManagedInstanceStreamStatusEnum
const (
	ListModuleStreamsOnManagedInstanceStreamStatusEnabled  ListModuleStreamsOnManagedInstanceStreamStatusEnum = "ENABLED"
	ListModuleStreamsOnManagedInstanceStreamStatusDisabled ListModuleStreamsOnManagedInstanceStreamStatusEnum = "DISABLED"
	ListModuleStreamsOnManagedInstanceStreamStatusActive   ListModuleStreamsOnManagedInstanceStreamStatusEnum = "ACTIVE"
)

var mappingListModuleStreamsOnManagedInstanceStreamStatusEnum = map[string]ListModuleStreamsOnManagedInstanceStreamStatusEnum{
	"ENABLED":  ListModuleStreamsOnManagedInstanceStreamStatusEnabled,
	"DISABLED": ListModuleStreamsOnManagedInstanceStreamStatusDisabled,
	"ACTIVE":   ListModuleStreamsOnManagedInstanceStreamStatusActive,
}

var mappingListModuleStreamsOnManagedInstanceStreamStatusEnumLowerCase = map[string]ListModuleStreamsOnManagedInstanceStreamStatusEnum{
	"enabled":  ListModuleStreamsOnManagedInstanceStreamStatusEnabled,
	"disabled": ListModuleStreamsOnManagedInstanceStreamStatusDisabled,
	"active":   ListModuleStreamsOnManagedInstanceStreamStatusActive,
}

// GetListModuleStreamsOnManagedInstanceStreamStatusEnumValues Enumerates the set of values for ListModuleStreamsOnManagedInstanceStreamStatusEnum
func GetListModuleStreamsOnManagedInstanceStreamStatusEnumValues() []ListModuleStreamsOnManagedInstanceStreamStatusEnum {
	values := make([]ListModuleStreamsOnManagedInstanceStreamStatusEnum, 0)
	for _, v := range mappingListModuleStreamsOnManagedInstanceStreamStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamsOnManagedInstanceStreamStatusEnumStringValues Enumerates the set of values in String for ListModuleStreamsOnManagedInstanceStreamStatusEnum
func GetListModuleStreamsOnManagedInstanceStreamStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"ACTIVE",
	}
}

// GetMappingListModuleStreamsOnManagedInstanceStreamStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamsOnManagedInstanceStreamStatusEnum(val string) (ListModuleStreamsOnManagedInstanceStreamStatusEnum, bool) {
	enum, ok := mappingListModuleStreamsOnManagedInstanceStreamStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModuleStreamsOnManagedInstanceSortOrderEnum Enum with underlying type: string
type ListModuleStreamsOnManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListModuleStreamsOnManagedInstanceSortOrderEnum
const (
	ListModuleStreamsOnManagedInstanceSortOrderAsc  ListModuleStreamsOnManagedInstanceSortOrderEnum = "ASC"
	ListModuleStreamsOnManagedInstanceSortOrderDesc ListModuleStreamsOnManagedInstanceSortOrderEnum = "DESC"
)

var mappingListModuleStreamsOnManagedInstanceSortOrderEnum = map[string]ListModuleStreamsOnManagedInstanceSortOrderEnum{
	"ASC":  ListModuleStreamsOnManagedInstanceSortOrderAsc,
	"DESC": ListModuleStreamsOnManagedInstanceSortOrderDesc,
}

var mappingListModuleStreamsOnManagedInstanceSortOrderEnumLowerCase = map[string]ListModuleStreamsOnManagedInstanceSortOrderEnum{
	"asc":  ListModuleStreamsOnManagedInstanceSortOrderAsc,
	"desc": ListModuleStreamsOnManagedInstanceSortOrderDesc,
}

// GetListModuleStreamsOnManagedInstanceSortOrderEnumValues Enumerates the set of values for ListModuleStreamsOnManagedInstanceSortOrderEnum
func GetListModuleStreamsOnManagedInstanceSortOrderEnumValues() []ListModuleStreamsOnManagedInstanceSortOrderEnum {
	values := make([]ListModuleStreamsOnManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListModuleStreamsOnManagedInstanceSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamsOnManagedInstanceSortOrderEnumStringValues Enumerates the set of values in String for ListModuleStreamsOnManagedInstanceSortOrderEnum
func GetListModuleStreamsOnManagedInstanceSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModuleStreamsOnManagedInstanceSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamsOnManagedInstanceSortOrderEnum(val string) (ListModuleStreamsOnManagedInstanceSortOrderEnum, bool) {
	enum, ok := mappingListModuleStreamsOnManagedInstanceSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModuleStreamsOnManagedInstanceSortByEnum Enum with underlying type: string
type ListModuleStreamsOnManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListModuleStreamsOnManagedInstanceSortByEnum
const (
	ListModuleStreamsOnManagedInstanceSortByTimecreated ListModuleStreamsOnManagedInstanceSortByEnum = "TIMECREATED"
	ListModuleStreamsOnManagedInstanceSortByDisplayname ListModuleStreamsOnManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListModuleStreamsOnManagedInstanceSortByEnum = map[string]ListModuleStreamsOnManagedInstanceSortByEnum{
	"TIMECREATED": ListModuleStreamsOnManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListModuleStreamsOnManagedInstanceSortByDisplayname,
}

var mappingListModuleStreamsOnManagedInstanceSortByEnumLowerCase = map[string]ListModuleStreamsOnManagedInstanceSortByEnum{
	"timecreated": ListModuleStreamsOnManagedInstanceSortByTimecreated,
	"displayname": ListModuleStreamsOnManagedInstanceSortByDisplayname,
}

// GetListModuleStreamsOnManagedInstanceSortByEnumValues Enumerates the set of values for ListModuleStreamsOnManagedInstanceSortByEnum
func GetListModuleStreamsOnManagedInstanceSortByEnumValues() []ListModuleStreamsOnManagedInstanceSortByEnum {
	values := make([]ListModuleStreamsOnManagedInstanceSortByEnum, 0)
	for _, v := range mappingListModuleStreamsOnManagedInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamsOnManagedInstanceSortByEnumStringValues Enumerates the set of values in String for ListModuleStreamsOnManagedInstanceSortByEnum
func GetListModuleStreamsOnManagedInstanceSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListModuleStreamsOnManagedInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamsOnManagedInstanceSortByEnum(val string) (ListModuleStreamsOnManagedInstanceSortByEnum, bool) {
	enum, ok := mappingListModuleStreamsOnManagedInstanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
