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

// ListTransferStatusPipelinesRequest wrapper for the ListTransferStatusPipelines operation
type ListTransferStatusPipelinesRequest struct {

	// CDS pipeline OCID.
	PipelineId *string `mandatory:"true" contributesTo:"path" name:"pipelineId"`

	// Optional filename filter for transfer status rows.
	Filename *string `mandatory:"false" contributesTo:"query" name:"filename"`

	// Optional checksum filter for transfer status rows.
	Checksum *string `mandatory:"false" contributesTo:"query" name:"checksum"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTransferStatusPipelinesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeLastUpdated is descending.
	SortBy ListTransferStatusPipelinesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTransferStatusPipelinesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTransferStatusPipelinesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTransferStatusPipelinesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTransferStatusPipelinesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTransferStatusPipelinesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTransferStatusPipelinesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTransferStatusPipelinesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTransferStatusPipelinesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTransferStatusPipelinesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTransferStatusPipelinesResponse wrapper for the ListTransferStatusPipelines operation
type ListTransferStatusPipelinesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TransferStatusPipelineSummaryCollection instances
	TransferStatusPipelineSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTransferStatusPipelinesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTransferStatusPipelinesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTransferStatusPipelinesSortOrderEnum Enum with underlying type: string
type ListTransferStatusPipelinesSortOrderEnum string

// Set of constants representing the allowable values for ListTransferStatusPipelinesSortOrderEnum
const (
	ListTransferStatusPipelinesSortOrderAsc  ListTransferStatusPipelinesSortOrderEnum = "ASC"
	ListTransferStatusPipelinesSortOrderDesc ListTransferStatusPipelinesSortOrderEnum = "DESC"
)

var mappingListTransferStatusPipelinesSortOrderEnum = map[string]ListTransferStatusPipelinesSortOrderEnum{
	"ASC":  ListTransferStatusPipelinesSortOrderAsc,
	"DESC": ListTransferStatusPipelinesSortOrderDesc,
}

var mappingListTransferStatusPipelinesSortOrderEnumLowerCase = map[string]ListTransferStatusPipelinesSortOrderEnum{
	"asc":  ListTransferStatusPipelinesSortOrderAsc,
	"desc": ListTransferStatusPipelinesSortOrderDesc,
}

// GetListTransferStatusPipelinesSortOrderEnumValues Enumerates the set of values for ListTransferStatusPipelinesSortOrderEnum
func GetListTransferStatusPipelinesSortOrderEnumValues() []ListTransferStatusPipelinesSortOrderEnum {
	values := make([]ListTransferStatusPipelinesSortOrderEnum, 0)
	for _, v := range mappingListTransferStatusPipelinesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTransferStatusPipelinesSortOrderEnumStringValues Enumerates the set of values in String for ListTransferStatusPipelinesSortOrderEnum
func GetListTransferStatusPipelinesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTransferStatusPipelinesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTransferStatusPipelinesSortOrderEnum(val string) (ListTransferStatusPipelinesSortOrderEnum, bool) {
	enum, ok := mappingListTransferStatusPipelinesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTransferStatusPipelinesSortByEnum Enum with underlying type: string
type ListTransferStatusPipelinesSortByEnum string

// Set of constants representing the allowable values for ListTransferStatusPipelinesSortByEnum
const (
	ListTransferStatusPipelinesSortByTimelastupdated ListTransferStatusPipelinesSortByEnum = "timeLastUpdated"
	ListTransferStatusPipelinesSortByTimeuploaded    ListTransferStatusPipelinesSortByEnum = "timeUploaded"
	ListTransferStatusPipelinesSortByTransferid      ListTransferStatusPipelinesSortByEnum = "transferId"
)

var mappingListTransferStatusPipelinesSortByEnum = map[string]ListTransferStatusPipelinesSortByEnum{
	"timeLastUpdated": ListTransferStatusPipelinesSortByTimelastupdated,
	"timeUploaded":    ListTransferStatusPipelinesSortByTimeuploaded,
	"transferId":      ListTransferStatusPipelinesSortByTransferid,
}

var mappingListTransferStatusPipelinesSortByEnumLowerCase = map[string]ListTransferStatusPipelinesSortByEnum{
	"timelastupdated": ListTransferStatusPipelinesSortByTimelastupdated,
	"timeuploaded":    ListTransferStatusPipelinesSortByTimeuploaded,
	"transferid":      ListTransferStatusPipelinesSortByTransferid,
}

// GetListTransferStatusPipelinesSortByEnumValues Enumerates the set of values for ListTransferStatusPipelinesSortByEnum
func GetListTransferStatusPipelinesSortByEnumValues() []ListTransferStatusPipelinesSortByEnum {
	values := make([]ListTransferStatusPipelinesSortByEnum, 0)
	for _, v := range mappingListTransferStatusPipelinesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTransferStatusPipelinesSortByEnumStringValues Enumerates the set of values in String for ListTransferStatusPipelinesSortByEnum
func GetListTransferStatusPipelinesSortByEnumStringValues() []string {
	return []string{
		"timeLastUpdated",
		"timeUploaded",
		"transferId",
	}
}

// GetMappingListTransferStatusPipelinesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTransferStatusPipelinesSortByEnum(val string) (ListTransferStatusPipelinesSortByEnum, bool) {
	enum, ok := mappingListTransferStatusPipelinesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
