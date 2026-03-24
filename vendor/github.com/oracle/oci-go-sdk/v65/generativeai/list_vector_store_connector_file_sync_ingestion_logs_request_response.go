// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVectorStoreConnectorFileSyncIngestionLogsRequest wrapper for the ListVectorStoreConnectorFileSyncIngestionLogs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListVectorStoreConnectorFileSyncIngestionLogs.go.html to see an example of how to use ListVectorStoreConnectorFileSyncIngestionLogsRequest.
type ListVectorStoreConnectorFileSyncIngestionLogsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vectorStoreConnectorFileSync.
	VectorStoreConnectorFileSyncId *string `mandatory:"true" contributesTo:"path" name:"vectorStoreConnectorFileSyncId"`

	// A filter to return only resources whose lifecycle state matches the given value
	LifecycleState VectorStoreConnectorFileSyncLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide only one sort order. Default order for `status` is ascending.
	SortBy ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVectorStoreConnectorFileSyncIngestionLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVectorStoreConnectorFileSyncIngestionLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVectorStoreConnectorFileSyncIngestionLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVectorStoreConnectorFileSyncIngestionLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVectorStoreConnectorFileSyncIngestionLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVectorStoreConnectorFileSyncLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVectorStoreConnectorFileSyncLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVectorStoreConnectorFileSyncIngestionLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVectorStoreConnectorFileSyncIngestionLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVectorStoreConnectorFileSyncIngestionLogsResponse wrapper for the ListVectorStoreConnectorFileSyncIngestionLogs operation
type ListVectorStoreConnectorFileSyncIngestionLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FileSyncIngestionLogsCollection instances
	FileSyncIngestionLogsCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVectorStoreConnectorFileSyncIngestionLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVectorStoreConnectorFileSyncIngestionLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum Enum with underlying type: string
type ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum string

// Set of constants representing the allowable values for ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum
const (
	ListVectorStoreConnectorFileSyncIngestionLogsSortByStatus   ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum = "status"
	ListVectorStoreConnectorFileSyncIngestionLogsSortByFilesize ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum = "fileSize"
	ListVectorStoreConnectorFileSyncIngestionLogsSortByFilepath ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum = "filePath"
)

var mappingListVectorStoreConnectorFileSyncIngestionLogsSortByEnum = map[string]ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum{
	"status":   ListVectorStoreConnectorFileSyncIngestionLogsSortByStatus,
	"fileSize": ListVectorStoreConnectorFileSyncIngestionLogsSortByFilesize,
	"filePath": ListVectorStoreConnectorFileSyncIngestionLogsSortByFilepath,
}

var mappingListVectorStoreConnectorFileSyncIngestionLogsSortByEnumLowerCase = map[string]ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum{
	"status":   ListVectorStoreConnectorFileSyncIngestionLogsSortByStatus,
	"filesize": ListVectorStoreConnectorFileSyncIngestionLogsSortByFilesize,
	"filepath": ListVectorStoreConnectorFileSyncIngestionLogsSortByFilepath,
}

// GetListVectorStoreConnectorFileSyncIngestionLogsSortByEnumValues Enumerates the set of values for ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum
func GetListVectorStoreConnectorFileSyncIngestionLogsSortByEnumValues() []ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum {
	values := make([]ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum, 0)
	for _, v := range mappingListVectorStoreConnectorFileSyncIngestionLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVectorStoreConnectorFileSyncIngestionLogsSortByEnumStringValues Enumerates the set of values in String for ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum
func GetListVectorStoreConnectorFileSyncIngestionLogsSortByEnumStringValues() []string {
	return []string{
		"status",
		"fileSize",
		"filePath",
	}
}

// GetMappingListVectorStoreConnectorFileSyncIngestionLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVectorStoreConnectorFileSyncIngestionLogsSortByEnum(val string) (ListVectorStoreConnectorFileSyncIngestionLogsSortByEnum, bool) {
	enum, ok := mappingListVectorStoreConnectorFileSyncIngestionLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum Enum with underlying type: string
type ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum string

// Set of constants representing the allowable values for ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum
const (
	ListVectorStoreConnectorFileSyncIngestionLogsSortOrderAsc  ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum = "ASC"
	ListVectorStoreConnectorFileSyncIngestionLogsSortOrderDesc ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum = "DESC"
)

var mappingListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum = map[string]ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum{
	"ASC":  ListVectorStoreConnectorFileSyncIngestionLogsSortOrderAsc,
	"DESC": ListVectorStoreConnectorFileSyncIngestionLogsSortOrderDesc,
}

var mappingListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnumLowerCase = map[string]ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum{
	"asc":  ListVectorStoreConnectorFileSyncIngestionLogsSortOrderAsc,
	"desc": ListVectorStoreConnectorFileSyncIngestionLogsSortOrderDesc,
}

// GetListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnumValues Enumerates the set of values for ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum
func GetListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnumValues() []ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum {
	values := make([]ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum, 0)
	for _, v := range mappingListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnumStringValues Enumerates the set of values in String for ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum
func GetListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum(val string) (ListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnum, bool) {
	enum, ok := mappingListVectorStoreConnectorFileSyncIngestionLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
