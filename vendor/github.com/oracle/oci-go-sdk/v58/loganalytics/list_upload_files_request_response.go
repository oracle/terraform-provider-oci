// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListUploadFilesRequest wrapper for the ListUploadFiles operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListUploadFiles.go.html to see an example of how to use ListUploadFilesRequest.
type ListUploadFilesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Unique internal identifier to refer upload container.
	UploadReference *string `mandatory:"true" contributesTo:"path" name:"uploadReference"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListUploadFilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeStarted is descending.
	// timeCreated, fileName and logGroup are deprecated. Instead use timestarted, name, logGroup accordingly.
	SortBy ListUploadFilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// This can be used to filter upload files based on the file, log group and log source names.
	SearchStr *string `mandatory:"false" contributesTo:"query" name:"searchStr"`

	// Upload File processing state.
	Status []ListUploadFilesStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUploadFilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUploadFilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUploadFilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUploadFilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUploadFilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUploadFilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUploadFilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUploadFilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUploadFilesSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Status {
		if _, ok := GetMappingListUploadFilesStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetListUploadFilesStatusEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUploadFilesResponse wrapper for the ListUploadFiles operation
type ListUploadFilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UploadFileCollection instances
	UploadFileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUploadFilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUploadFilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUploadFilesSortOrderEnum Enum with underlying type: string
type ListUploadFilesSortOrderEnum string

// Set of constants representing the allowable values for ListUploadFilesSortOrderEnum
const (
	ListUploadFilesSortOrderAsc  ListUploadFilesSortOrderEnum = "ASC"
	ListUploadFilesSortOrderDesc ListUploadFilesSortOrderEnum = "DESC"
)

var mappingListUploadFilesSortOrderEnum = map[string]ListUploadFilesSortOrderEnum{
	"ASC":  ListUploadFilesSortOrderAsc,
	"DESC": ListUploadFilesSortOrderDesc,
}

// GetListUploadFilesSortOrderEnumValues Enumerates the set of values for ListUploadFilesSortOrderEnum
func GetListUploadFilesSortOrderEnumValues() []ListUploadFilesSortOrderEnum {
	values := make([]ListUploadFilesSortOrderEnum, 0)
	for _, v := range mappingListUploadFilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUploadFilesSortOrderEnumStringValues Enumerates the set of values in String for ListUploadFilesSortOrderEnum
func GetListUploadFilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUploadFilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUploadFilesSortOrderEnum(val string) (ListUploadFilesSortOrderEnum, bool) {
	mappingListUploadFilesSortOrderEnumIgnoreCase := make(map[string]ListUploadFilesSortOrderEnum)
	for k, v := range mappingListUploadFilesSortOrderEnum {
		mappingListUploadFilesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUploadFilesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUploadFilesSortByEnum Enum with underlying type: string
type ListUploadFilesSortByEnum string

// Set of constants representing the allowable values for ListUploadFilesSortByEnum
const (
	ListUploadFilesSortByTimestarted  ListUploadFilesSortByEnum = "timeStarted"
	ListUploadFilesSortByName         ListUploadFilesSortByEnum = "name"
	ListUploadFilesSortByLoggroupname ListUploadFilesSortByEnum = "logGroupName"
	ListUploadFilesSortBySourcename   ListUploadFilesSortByEnum = "sourceName"
	ListUploadFilesSortByStatus       ListUploadFilesSortByEnum = "status"
	ListUploadFilesSortByTimecreated  ListUploadFilesSortByEnum = "timeCreated"
	ListUploadFilesSortByFilename     ListUploadFilesSortByEnum = "fileName"
	ListUploadFilesSortByLoggroup     ListUploadFilesSortByEnum = "logGroup"
)

var mappingListUploadFilesSortByEnum = map[string]ListUploadFilesSortByEnum{
	"timeStarted":  ListUploadFilesSortByTimestarted,
	"name":         ListUploadFilesSortByName,
	"logGroupName": ListUploadFilesSortByLoggroupname,
	"sourceName":   ListUploadFilesSortBySourcename,
	"status":       ListUploadFilesSortByStatus,
	"timeCreated":  ListUploadFilesSortByTimecreated,
	"fileName":     ListUploadFilesSortByFilename,
	"logGroup":     ListUploadFilesSortByLoggroup,
}

// GetListUploadFilesSortByEnumValues Enumerates the set of values for ListUploadFilesSortByEnum
func GetListUploadFilesSortByEnumValues() []ListUploadFilesSortByEnum {
	values := make([]ListUploadFilesSortByEnum, 0)
	for _, v := range mappingListUploadFilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUploadFilesSortByEnumStringValues Enumerates the set of values in String for ListUploadFilesSortByEnum
func GetListUploadFilesSortByEnumStringValues() []string {
	return []string{
		"timeStarted",
		"name",
		"logGroupName",
		"sourceName",
		"status",
		"timeCreated",
		"fileName",
		"logGroup",
	}
}

// GetMappingListUploadFilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUploadFilesSortByEnum(val string) (ListUploadFilesSortByEnum, bool) {
	mappingListUploadFilesSortByEnumIgnoreCase := make(map[string]ListUploadFilesSortByEnum)
	for k, v := range mappingListUploadFilesSortByEnum {
		mappingListUploadFilesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUploadFilesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUploadFilesStatusEnum Enum with underlying type: string
type ListUploadFilesStatusEnum string

// Set of constants representing the allowable values for ListUploadFilesStatusEnum
const (
	ListUploadFilesStatusInProgress ListUploadFilesStatusEnum = "IN_PROGRESS"
	ListUploadFilesStatusSuccessful ListUploadFilesStatusEnum = "SUCCESSFUL"
	ListUploadFilesStatusFailed     ListUploadFilesStatusEnum = "FAILED"
)

var mappingListUploadFilesStatusEnum = map[string]ListUploadFilesStatusEnum{
	"IN_PROGRESS": ListUploadFilesStatusInProgress,
	"SUCCESSFUL":  ListUploadFilesStatusSuccessful,
	"FAILED":      ListUploadFilesStatusFailed,
}

// GetListUploadFilesStatusEnumValues Enumerates the set of values for ListUploadFilesStatusEnum
func GetListUploadFilesStatusEnumValues() []ListUploadFilesStatusEnum {
	values := make([]ListUploadFilesStatusEnum, 0)
	for _, v := range mappingListUploadFilesStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListUploadFilesStatusEnumStringValues Enumerates the set of values in String for ListUploadFilesStatusEnum
func GetListUploadFilesStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCESSFUL",
		"FAILED",
	}
}

// GetMappingListUploadFilesStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUploadFilesStatusEnum(val string) (ListUploadFilesStatusEnum, bool) {
	mappingListUploadFilesStatusEnumIgnoreCase := make(map[string]ListUploadFilesStatusEnum)
	for k, v := range mappingListUploadFilesStatusEnum {
		mappingListUploadFilesStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUploadFilesStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
