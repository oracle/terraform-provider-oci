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

// ListVectorStoreConnectorIngestionLogsRequest wrapper for the ListVectorStoreConnectorIngestionLogs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListVectorStoreConnectorIngestionLogs.go.html to see an example of how to use ListVectorStoreConnectorIngestionLogsRequest.
type ListVectorStoreConnectorIngestionLogsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VectorStoreConnector.
	VectorStoreConnectorId *string `mandatory:"true" contributesTo:"path" name:"vectorStoreConnectorId"`

	// A filter to return only the Sync Logs whose status matches the given value.
	Status VectorStoreConnectorIngestionLogsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

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
	SortBy ListVectorStoreConnectorIngestionLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVectorStoreConnectorIngestionLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVectorStoreConnectorIngestionLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVectorStoreConnectorIngestionLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVectorStoreConnectorIngestionLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVectorStoreConnectorIngestionLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVectorStoreConnectorIngestionLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVectorStoreConnectorIngestionLogsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetVectorStoreConnectorIngestionLogsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVectorStoreConnectorIngestionLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVectorStoreConnectorIngestionLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVectorStoreConnectorIngestionLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVectorStoreConnectorIngestionLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVectorStoreConnectorIngestionLogsResponse wrapper for the ListVectorStoreConnectorIngestionLogs operation
type ListVectorStoreConnectorIngestionLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VectorStoreConnectorIngestionLogsCollection instances
	VectorStoreConnectorIngestionLogsCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVectorStoreConnectorIngestionLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVectorStoreConnectorIngestionLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVectorStoreConnectorIngestionLogsSortByEnum Enum with underlying type: string
type ListVectorStoreConnectorIngestionLogsSortByEnum string

// Set of constants representing the allowable values for ListVectorStoreConnectorIngestionLogsSortByEnum
const (
	ListVectorStoreConnectorIngestionLogsSortByStatus   ListVectorStoreConnectorIngestionLogsSortByEnum = "status"
	ListVectorStoreConnectorIngestionLogsSortByFilesize ListVectorStoreConnectorIngestionLogsSortByEnum = "fileSize"
	ListVectorStoreConnectorIngestionLogsSortByFilepath ListVectorStoreConnectorIngestionLogsSortByEnum = "filePath"
)

var mappingListVectorStoreConnectorIngestionLogsSortByEnum = map[string]ListVectorStoreConnectorIngestionLogsSortByEnum{
	"status":   ListVectorStoreConnectorIngestionLogsSortByStatus,
	"fileSize": ListVectorStoreConnectorIngestionLogsSortByFilesize,
	"filePath": ListVectorStoreConnectorIngestionLogsSortByFilepath,
}

var mappingListVectorStoreConnectorIngestionLogsSortByEnumLowerCase = map[string]ListVectorStoreConnectorIngestionLogsSortByEnum{
	"status":   ListVectorStoreConnectorIngestionLogsSortByStatus,
	"filesize": ListVectorStoreConnectorIngestionLogsSortByFilesize,
	"filepath": ListVectorStoreConnectorIngestionLogsSortByFilepath,
}

// GetListVectorStoreConnectorIngestionLogsSortByEnumValues Enumerates the set of values for ListVectorStoreConnectorIngestionLogsSortByEnum
func GetListVectorStoreConnectorIngestionLogsSortByEnumValues() []ListVectorStoreConnectorIngestionLogsSortByEnum {
	values := make([]ListVectorStoreConnectorIngestionLogsSortByEnum, 0)
	for _, v := range mappingListVectorStoreConnectorIngestionLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVectorStoreConnectorIngestionLogsSortByEnumStringValues Enumerates the set of values in String for ListVectorStoreConnectorIngestionLogsSortByEnum
func GetListVectorStoreConnectorIngestionLogsSortByEnumStringValues() []string {
	return []string{
		"status",
		"fileSize",
		"filePath",
	}
}

// GetMappingListVectorStoreConnectorIngestionLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVectorStoreConnectorIngestionLogsSortByEnum(val string) (ListVectorStoreConnectorIngestionLogsSortByEnum, bool) {
	enum, ok := mappingListVectorStoreConnectorIngestionLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVectorStoreConnectorIngestionLogsSortOrderEnum Enum with underlying type: string
type ListVectorStoreConnectorIngestionLogsSortOrderEnum string

// Set of constants representing the allowable values for ListVectorStoreConnectorIngestionLogsSortOrderEnum
const (
	ListVectorStoreConnectorIngestionLogsSortOrderAsc  ListVectorStoreConnectorIngestionLogsSortOrderEnum = "ASC"
	ListVectorStoreConnectorIngestionLogsSortOrderDesc ListVectorStoreConnectorIngestionLogsSortOrderEnum = "DESC"
)

var mappingListVectorStoreConnectorIngestionLogsSortOrderEnum = map[string]ListVectorStoreConnectorIngestionLogsSortOrderEnum{
	"ASC":  ListVectorStoreConnectorIngestionLogsSortOrderAsc,
	"DESC": ListVectorStoreConnectorIngestionLogsSortOrderDesc,
}

var mappingListVectorStoreConnectorIngestionLogsSortOrderEnumLowerCase = map[string]ListVectorStoreConnectorIngestionLogsSortOrderEnum{
	"asc":  ListVectorStoreConnectorIngestionLogsSortOrderAsc,
	"desc": ListVectorStoreConnectorIngestionLogsSortOrderDesc,
}

// GetListVectorStoreConnectorIngestionLogsSortOrderEnumValues Enumerates the set of values for ListVectorStoreConnectorIngestionLogsSortOrderEnum
func GetListVectorStoreConnectorIngestionLogsSortOrderEnumValues() []ListVectorStoreConnectorIngestionLogsSortOrderEnum {
	values := make([]ListVectorStoreConnectorIngestionLogsSortOrderEnum, 0)
	for _, v := range mappingListVectorStoreConnectorIngestionLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVectorStoreConnectorIngestionLogsSortOrderEnumStringValues Enumerates the set of values in String for ListVectorStoreConnectorIngestionLogsSortOrderEnum
func GetListVectorStoreConnectorIngestionLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVectorStoreConnectorIngestionLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVectorStoreConnectorIngestionLogsSortOrderEnum(val string) (ListVectorStoreConnectorIngestionLogsSortOrderEnum, bool) {
	enum, ok := mappingListVectorStoreConnectorIngestionLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
