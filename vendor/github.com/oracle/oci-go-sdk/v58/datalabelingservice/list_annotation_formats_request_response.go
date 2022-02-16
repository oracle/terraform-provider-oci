// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListAnnotationFormatsRequest wrapper for the ListAnnotationFormats operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservice/ListAnnotationFormats.go.html to see an example of how to use ListAnnotationFormatsRequest.
type ListAnnotationFormatsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAnnotationFormatsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAnnotationFormatsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAnnotationFormatsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAnnotationFormatsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAnnotationFormatsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAnnotationFormatsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAnnotationFormatsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAnnotationFormatsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAnnotationFormatsResponse wrapper for the ListAnnotationFormats operation
type ListAnnotationFormatsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AnnotationFormatCollection instances
	AnnotationFormatCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For the pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAnnotationFormatsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAnnotationFormatsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAnnotationFormatsSortOrderEnum Enum with underlying type: string
type ListAnnotationFormatsSortOrderEnum string

// Set of constants representing the allowable values for ListAnnotationFormatsSortOrderEnum
const (
	ListAnnotationFormatsSortOrderAsc  ListAnnotationFormatsSortOrderEnum = "ASC"
	ListAnnotationFormatsSortOrderDesc ListAnnotationFormatsSortOrderEnum = "DESC"
)

var mappingListAnnotationFormatsSortOrderEnum = map[string]ListAnnotationFormatsSortOrderEnum{
	"ASC":  ListAnnotationFormatsSortOrderAsc,
	"DESC": ListAnnotationFormatsSortOrderDesc,
}

// GetListAnnotationFormatsSortOrderEnumValues Enumerates the set of values for ListAnnotationFormatsSortOrderEnum
func GetListAnnotationFormatsSortOrderEnumValues() []ListAnnotationFormatsSortOrderEnum {
	values := make([]ListAnnotationFormatsSortOrderEnum, 0)
	for _, v := range mappingListAnnotationFormatsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAnnotationFormatsSortOrderEnumStringValues Enumerates the set of values in String for ListAnnotationFormatsSortOrderEnum
func GetListAnnotationFormatsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAnnotationFormatsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAnnotationFormatsSortOrderEnum(val string) (ListAnnotationFormatsSortOrderEnum, bool) {
	mappingListAnnotationFormatsSortOrderEnumIgnoreCase := make(map[string]ListAnnotationFormatsSortOrderEnum)
	for k, v := range mappingListAnnotationFormatsSortOrderEnum {
		mappingListAnnotationFormatsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAnnotationFormatsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
