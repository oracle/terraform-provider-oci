// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package gdp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTransferStatusesRequest wrapper for the ListTransferStatuses operation
type ListTransferStatusesRequest struct {

	// CDS pipeline OCID whose transfer statuses should be listed.
	PipelineId *string `mandatory:"true" contributesTo:"query" name:"pipelineId"`

	// Filename of the transferred object.
	Filename *string `mandatory:"true" contributesTo:"query" name:"filename"`

	// Checksum captured for the transferred object.
	Checksum *string `mandatory:"true" contributesTo:"query" name:"checksum"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTransferStatusesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeLastUpdated is descending.
	SortBy ListTransferStatusesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTransferStatusesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTransferStatusesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTransferStatusesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTransferStatusesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTransferStatusesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTransferStatusesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTransferStatusesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTransferStatusesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTransferStatusesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTransferStatusesResponse wrapper for the ListTransferStatuses operation
type ListTransferStatusesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TransferStatusSummaryCollection instances
	TransferStatusSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTransferStatusesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTransferStatusesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTransferStatusesSortOrderEnum Enum with underlying type: string
type ListTransferStatusesSortOrderEnum string

// Set of constants representing the allowable values for ListTransferStatusesSortOrderEnum
const (
	ListTransferStatusesSortOrderAsc  ListTransferStatusesSortOrderEnum = "ASC"
	ListTransferStatusesSortOrderDesc ListTransferStatusesSortOrderEnum = "DESC"
)

var mappingListTransferStatusesSortOrderEnum = map[string]ListTransferStatusesSortOrderEnum{
	"ASC":  ListTransferStatusesSortOrderAsc,
	"DESC": ListTransferStatusesSortOrderDesc,
}

var mappingListTransferStatusesSortOrderEnumLowerCase = map[string]ListTransferStatusesSortOrderEnum{
	"asc":  ListTransferStatusesSortOrderAsc,
	"desc": ListTransferStatusesSortOrderDesc,
}

// GetListTransferStatusesSortOrderEnumValues Enumerates the set of values for ListTransferStatusesSortOrderEnum
func GetListTransferStatusesSortOrderEnumValues() []ListTransferStatusesSortOrderEnum {
	values := make([]ListTransferStatusesSortOrderEnum, 0)
	for _, v := range mappingListTransferStatusesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTransferStatusesSortOrderEnumStringValues Enumerates the set of values in String for ListTransferStatusesSortOrderEnum
func GetListTransferStatusesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTransferStatusesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTransferStatusesSortOrderEnum(val string) (ListTransferStatusesSortOrderEnum, bool) {
	enum, ok := mappingListTransferStatusesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTransferStatusesSortByEnum Enum with underlying type: string
type ListTransferStatusesSortByEnum string

// Set of constants representing the allowable values for ListTransferStatusesSortByEnum
const (
	ListTransferStatusesSortByTimelastupdated ListTransferStatusesSortByEnum = "timeLastUpdated"
	ListTransferStatusesSortByTimeuploaded    ListTransferStatusesSortByEnum = "timeUploaded"
	ListTransferStatusesSortByTransferid      ListTransferStatusesSortByEnum = "transferId"
)

var mappingListTransferStatusesSortByEnum = map[string]ListTransferStatusesSortByEnum{
	"timeLastUpdated": ListTransferStatusesSortByTimelastupdated,
	"timeUploaded":    ListTransferStatusesSortByTimeuploaded,
	"transferId":      ListTransferStatusesSortByTransferid,
}

var mappingListTransferStatusesSortByEnumLowerCase = map[string]ListTransferStatusesSortByEnum{
	"timelastupdated": ListTransferStatusesSortByTimelastupdated,
	"timeuploaded":    ListTransferStatusesSortByTimeuploaded,
	"transferid":      ListTransferStatusesSortByTransferid,
}

// GetListTransferStatusesSortByEnumValues Enumerates the set of values for ListTransferStatusesSortByEnum
func GetListTransferStatusesSortByEnumValues() []ListTransferStatusesSortByEnum {
	values := make([]ListTransferStatusesSortByEnum, 0)
	for _, v := range mappingListTransferStatusesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTransferStatusesSortByEnumStringValues Enumerates the set of values in String for ListTransferStatusesSortByEnum
func GetListTransferStatusesSortByEnumStringValues() []string {
	return []string{
		"timeLastUpdated",
		"timeUploaded",
		"transferId",
	}
}

// GetMappingListTransferStatusesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTransferStatusesSortByEnum(val string) (ListTransferStatusesSortByEnum, bool) {
	enum, ok := mappingListTransferStatusesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
