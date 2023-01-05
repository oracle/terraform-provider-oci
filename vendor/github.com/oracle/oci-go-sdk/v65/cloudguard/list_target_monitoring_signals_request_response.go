// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTargetMonitoringSignalsRequest wrapper for the ListTargetMonitoringSignals operation
type ListTargetMonitoringSignalsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// OCID of target
	TargetId *string `mandatory:"true" contributesTo:"path" name:"targetId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field scan status. Only one status can be provided. Default value for scan status is failed. If no value is specified scan status is failed.
	ScanStatus ListTargetMonitoringSignalsScanStatusEnum `mandatory:"false" contributesTo:"query" name:"scanStatus" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTargetMonitoringSignalsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListTargetMonitoringSignalsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetMonitoringSignalsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetMonitoringSignalsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetMonitoringSignalsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetMonitoringSignalsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetMonitoringSignalsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetMonitoringSignalsScanStatusEnum(string(request.ScanStatus)); !ok && request.ScanStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScanStatus: %s. Supported values are: %s.", request.ScanStatus, strings.Join(GetListTargetMonitoringSignalsScanStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetMonitoringSignalsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetMonitoringSignalsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetMonitoringSignalsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetMonitoringSignalsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetMonitoringSignalsResponse wrapper for the ListTargetMonitoringSignals operation
type ListTargetMonitoringSignalsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetMonitoringSignalCollection instances
	TargetMonitoringSignalCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetMonitoringSignalsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetMonitoringSignalsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetMonitoringSignalsScanStatusEnum Enum with underlying type: string
type ListTargetMonitoringSignalsScanStatusEnum string

// Set of constants representing the allowable values for ListTargetMonitoringSignalsScanStatusEnum
const (
	ListTargetMonitoringSignalsScanStatusSucceeded ListTargetMonitoringSignalsScanStatusEnum = "SUCCEEDED"
	ListTargetMonitoringSignalsScanStatusFailed    ListTargetMonitoringSignalsScanStatusEnum = "FAILED"
)

var mappingListTargetMonitoringSignalsScanStatusEnum = map[string]ListTargetMonitoringSignalsScanStatusEnum{
	"SUCCEEDED": ListTargetMonitoringSignalsScanStatusSucceeded,
	"FAILED":    ListTargetMonitoringSignalsScanStatusFailed,
}

var mappingListTargetMonitoringSignalsScanStatusEnumLowerCase = map[string]ListTargetMonitoringSignalsScanStatusEnum{
	"succeeded": ListTargetMonitoringSignalsScanStatusSucceeded,
	"failed":    ListTargetMonitoringSignalsScanStatusFailed,
}

// GetListTargetMonitoringSignalsScanStatusEnumValues Enumerates the set of values for ListTargetMonitoringSignalsScanStatusEnum
func GetListTargetMonitoringSignalsScanStatusEnumValues() []ListTargetMonitoringSignalsScanStatusEnum {
	values := make([]ListTargetMonitoringSignalsScanStatusEnum, 0)
	for _, v := range mappingListTargetMonitoringSignalsScanStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetMonitoringSignalsScanStatusEnumStringValues Enumerates the set of values in String for ListTargetMonitoringSignalsScanStatusEnum
func GetListTargetMonitoringSignalsScanStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingListTargetMonitoringSignalsScanStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetMonitoringSignalsScanStatusEnum(val string) (ListTargetMonitoringSignalsScanStatusEnum, bool) {
	enum, ok := mappingListTargetMonitoringSignalsScanStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetMonitoringSignalsSortOrderEnum Enum with underlying type: string
type ListTargetMonitoringSignalsSortOrderEnum string

// Set of constants representing the allowable values for ListTargetMonitoringSignalsSortOrderEnum
const (
	ListTargetMonitoringSignalsSortOrderAsc  ListTargetMonitoringSignalsSortOrderEnum = "ASC"
	ListTargetMonitoringSignalsSortOrderDesc ListTargetMonitoringSignalsSortOrderEnum = "DESC"
)

var mappingListTargetMonitoringSignalsSortOrderEnum = map[string]ListTargetMonitoringSignalsSortOrderEnum{
	"ASC":  ListTargetMonitoringSignalsSortOrderAsc,
	"DESC": ListTargetMonitoringSignalsSortOrderDesc,
}

var mappingListTargetMonitoringSignalsSortOrderEnumLowerCase = map[string]ListTargetMonitoringSignalsSortOrderEnum{
	"asc":  ListTargetMonitoringSignalsSortOrderAsc,
	"desc": ListTargetMonitoringSignalsSortOrderDesc,
}

// GetListTargetMonitoringSignalsSortOrderEnumValues Enumerates the set of values for ListTargetMonitoringSignalsSortOrderEnum
func GetListTargetMonitoringSignalsSortOrderEnumValues() []ListTargetMonitoringSignalsSortOrderEnum {
	values := make([]ListTargetMonitoringSignalsSortOrderEnum, 0)
	for _, v := range mappingListTargetMonitoringSignalsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetMonitoringSignalsSortOrderEnumStringValues Enumerates the set of values in String for ListTargetMonitoringSignalsSortOrderEnum
func GetListTargetMonitoringSignalsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetMonitoringSignalsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetMonitoringSignalsSortOrderEnum(val string) (ListTargetMonitoringSignalsSortOrderEnum, bool) {
	enum, ok := mappingListTargetMonitoringSignalsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetMonitoringSignalsSortByEnum Enum with underlying type: string
type ListTargetMonitoringSignalsSortByEnum string

// Set of constants representing the allowable values for ListTargetMonitoringSignalsSortByEnum
const (
	ListTargetMonitoringSignalsSortByTimecreated ListTargetMonitoringSignalsSortByEnum = "timeCreated"
	ListTargetMonitoringSignalsSortByDisplayname ListTargetMonitoringSignalsSortByEnum = "displayName"
)

var mappingListTargetMonitoringSignalsSortByEnum = map[string]ListTargetMonitoringSignalsSortByEnum{
	"timeCreated": ListTargetMonitoringSignalsSortByTimecreated,
	"displayName": ListTargetMonitoringSignalsSortByDisplayname,
}

var mappingListTargetMonitoringSignalsSortByEnumLowerCase = map[string]ListTargetMonitoringSignalsSortByEnum{
	"timecreated": ListTargetMonitoringSignalsSortByTimecreated,
	"displayname": ListTargetMonitoringSignalsSortByDisplayname,
}

// GetListTargetMonitoringSignalsSortByEnumValues Enumerates the set of values for ListTargetMonitoringSignalsSortByEnum
func GetListTargetMonitoringSignalsSortByEnumValues() []ListTargetMonitoringSignalsSortByEnum {
	values := make([]ListTargetMonitoringSignalsSortByEnum, 0)
	for _, v := range mappingListTargetMonitoringSignalsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetMonitoringSignalsSortByEnumStringValues Enumerates the set of values in String for ListTargetMonitoringSignalsSortByEnum
func GetListTargetMonitoringSignalsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTargetMonitoringSignalsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetMonitoringSignalsSortByEnum(val string) (ListTargetMonitoringSignalsSortByEnum, bool) {
	enum, ok := mappingListTargetMonitoringSignalsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
