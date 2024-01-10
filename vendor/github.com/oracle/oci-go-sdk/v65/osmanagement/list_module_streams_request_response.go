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

// ListModuleStreamsRequest wrapper for the ListModuleStreams operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListModuleStreams.go.html to see an example of how to use ListModuleStreamsRequest.
type ListModuleStreamsRequest struct {

	// The OCID of the software source.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The name of a module.  This parameter is required if a
	// streamName is specified.
	ModuleName *string `mandatory:"false" contributesTo:"query" name:"moduleName"`

	// The name of the stream of the containing module.  This parameter
	// is required if a profileName is specified.
	StreamName *string `mandatory:"false" contributesTo:"query" name:"streamName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListModuleStreamsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListModuleStreamsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModuleStreamsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModuleStreamsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModuleStreamsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModuleStreamsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModuleStreamsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModuleStreamsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModuleStreamsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModuleStreamsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModuleStreamsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModuleStreamsResponse wrapper for the ListModuleStreams operation
type ListModuleStreamsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModuleStreamSummary instances
	Items []ModuleStreamSummary `presentIn:"body"`

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

func (response ListModuleStreamsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModuleStreamsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModuleStreamsSortOrderEnum Enum with underlying type: string
type ListModuleStreamsSortOrderEnum string

// Set of constants representing the allowable values for ListModuleStreamsSortOrderEnum
const (
	ListModuleStreamsSortOrderAsc  ListModuleStreamsSortOrderEnum = "ASC"
	ListModuleStreamsSortOrderDesc ListModuleStreamsSortOrderEnum = "DESC"
)

var mappingListModuleStreamsSortOrderEnum = map[string]ListModuleStreamsSortOrderEnum{
	"ASC":  ListModuleStreamsSortOrderAsc,
	"DESC": ListModuleStreamsSortOrderDesc,
}

var mappingListModuleStreamsSortOrderEnumLowerCase = map[string]ListModuleStreamsSortOrderEnum{
	"asc":  ListModuleStreamsSortOrderAsc,
	"desc": ListModuleStreamsSortOrderDesc,
}

// GetListModuleStreamsSortOrderEnumValues Enumerates the set of values for ListModuleStreamsSortOrderEnum
func GetListModuleStreamsSortOrderEnumValues() []ListModuleStreamsSortOrderEnum {
	values := make([]ListModuleStreamsSortOrderEnum, 0)
	for _, v := range mappingListModuleStreamsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamsSortOrderEnumStringValues Enumerates the set of values in String for ListModuleStreamsSortOrderEnum
func GetListModuleStreamsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModuleStreamsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamsSortOrderEnum(val string) (ListModuleStreamsSortOrderEnum, bool) {
	enum, ok := mappingListModuleStreamsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModuleStreamsSortByEnum Enum with underlying type: string
type ListModuleStreamsSortByEnum string

// Set of constants representing the allowable values for ListModuleStreamsSortByEnum
const (
	ListModuleStreamsSortByTimecreated ListModuleStreamsSortByEnum = "TIMECREATED"
	ListModuleStreamsSortByDisplayname ListModuleStreamsSortByEnum = "DISPLAYNAME"
)

var mappingListModuleStreamsSortByEnum = map[string]ListModuleStreamsSortByEnum{
	"TIMECREATED": ListModuleStreamsSortByTimecreated,
	"DISPLAYNAME": ListModuleStreamsSortByDisplayname,
}

var mappingListModuleStreamsSortByEnumLowerCase = map[string]ListModuleStreamsSortByEnum{
	"timecreated": ListModuleStreamsSortByTimecreated,
	"displayname": ListModuleStreamsSortByDisplayname,
}

// GetListModuleStreamsSortByEnumValues Enumerates the set of values for ListModuleStreamsSortByEnum
func GetListModuleStreamsSortByEnumValues() []ListModuleStreamsSortByEnum {
	values := make([]ListModuleStreamsSortByEnum, 0)
	for _, v := range mappingListModuleStreamsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModuleStreamsSortByEnumStringValues Enumerates the set of values in String for ListModuleStreamsSortByEnum
func GetListModuleStreamsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListModuleStreamsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModuleStreamsSortByEnum(val string) (ListModuleStreamsSortByEnum, bool) {
	enum, ok := mappingListModuleStreamsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
