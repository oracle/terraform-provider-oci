// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// QueryLogSetsRequest wrapper for the QueryLogSets operation
type QueryLogSetsRequest struct {

	// The Log Analytics namespace used for the request. The namespace can be obtained by running 'oci os ns get'
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// This is the input to query log sets.
	QueryLogSetsDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder QueryLogSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If this filter is present, each of the logsets returned must contain the value of this filter.
	LogSetNameContains []string `contributesTo:"query" name:"logSetNameContains" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request QueryLogSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request QueryLogSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request QueryLogSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request QueryLogSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request QueryLogSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueryLogSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetQueryLogSetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryLogSetsResponse wrapper for the QueryLogSets operation
type QueryLogSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogSetCollection instances
	LogSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response QueryLogSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response QueryLogSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// QueryLogSetsSortOrderEnum Enum with underlying type: string
type QueryLogSetsSortOrderEnum string

// Set of constants representing the allowable values for QueryLogSetsSortOrderEnum
const (
	QueryLogSetsSortOrderAsc  QueryLogSetsSortOrderEnum = "ASC"
	QueryLogSetsSortOrderDesc QueryLogSetsSortOrderEnum = "DESC"
)

var mappingQueryLogSetsSortOrderEnum = map[string]QueryLogSetsSortOrderEnum{
	"ASC":  QueryLogSetsSortOrderAsc,
	"DESC": QueryLogSetsSortOrderDesc,
}

var mappingQueryLogSetsSortOrderEnumLowerCase = map[string]QueryLogSetsSortOrderEnum{
	"asc":  QueryLogSetsSortOrderAsc,
	"desc": QueryLogSetsSortOrderDesc,
}

// GetQueryLogSetsSortOrderEnumValues Enumerates the set of values for QueryLogSetsSortOrderEnum
func GetQueryLogSetsSortOrderEnumValues() []QueryLogSetsSortOrderEnum {
	values := make([]QueryLogSetsSortOrderEnum, 0)
	for _, v := range mappingQueryLogSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryLogSetsSortOrderEnumStringValues Enumerates the set of values in String for QueryLogSetsSortOrderEnum
func GetQueryLogSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingQueryLogSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryLogSetsSortOrderEnum(val string) (QueryLogSetsSortOrderEnum, bool) {
	enum, ok := mappingQueryLogSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
