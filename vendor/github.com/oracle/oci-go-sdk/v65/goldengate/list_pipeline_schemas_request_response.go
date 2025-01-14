// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPipelineSchemasRequest wrapper for the ListPipelineSchemas operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelineSchemas.go.html to see an example of how to use ListPipelineSchemasRequest.
type ListPipelineSchemasRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the pipeline created.
	PipelineId *string `mandatory:"true" contributesTo:"path" name:"pipelineId"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPipelineSchemasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListPipelineSchemasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPipelineSchemasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPipelineSchemasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPipelineSchemasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPipelineSchemasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPipelineSchemasRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPipelineSchemasSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPipelineSchemasSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelineSchemasSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPipelineSchemasSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPipelineSchemasResponse wrapper for the ListPipelineSchemas operation
type ListPipelineSchemasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PipelineSchemaCollection instances
	PipelineSchemaCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPipelineSchemasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPipelineSchemasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPipelineSchemasSortOrderEnum Enum with underlying type: string
type ListPipelineSchemasSortOrderEnum string

// Set of constants representing the allowable values for ListPipelineSchemasSortOrderEnum
const (
	ListPipelineSchemasSortOrderAsc  ListPipelineSchemasSortOrderEnum = "ASC"
	ListPipelineSchemasSortOrderDesc ListPipelineSchemasSortOrderEnum = "DESC"
)

var mappingListPipelineSchemasSortOrderEnum = map[string]ListPipelineSchemasSortOrderEnum{
	"ASC":  ListPipelineSchemasSortOrderAsc,
	"DESC": ListPipelineSchemasSortOrderDesc,
}

var mappingListPipelineSchemasSortOrderEnumLowerCase = map[string]ListPipelineSchemasSortOrderEnum{
	"asc":  ListPipelineSchemasSortOrderAsc,
	"desc": ListPipelineSchemasSortOrderDesc,
}

// GetListPipelineSchemasSortOrderEnumValues Enumerates the set of values for ListPipelineSchemasSortOrderEnum
func GetListPipelineSchemasSortOrderEnumValues() []ListPipelineSchemasSortOrderEnum {
	values := make([]ListPipelineSchemasSortOrderEnum, 0)
	for _, v := range mappingListPipelineSchemasSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineSchemasSortOrderEnumStringValues Enumerates the set of values in String for ListPipelineSchemasSortOrderEnum
func GetListPipelineSchemasSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPipelineSchemasSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineSchemasSortOrderEnum(val string) (ListPipelineSchemasSortOrderEnum, bool) {
	enum, ok := mappingListPipelineSchemasSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelineSchemasSortByEnum Enum with underlying type: string
type ListPipelineSchemasSortByEnum string

// Set of constants representing the allowable values for ListPipelineSchemasSortByEnum
const (
	ListPipelineSchemasSortByTimecreated ListPipelineSchemasSortByEnum = "timeCreated"
	ListPipelineSchemasSortByDisplayname ListPipelineSchemasSortByEnum = "displayName"
)

var mappingListPipelineSchemasSortByEnum = map[string]ListPipelineSchemasSortByEnum{
	"timeCreated": ListPipelineSchemasSortByTimecreated,
	"displayName": ListPipelineSchemasSortByDisplayname,
}

var mappingListPipelineSchemasSortByEnumLowerCase = map[string]ListPipelineSchemasSortByEnum{
	"timecreated": ListPipelineSchemasSortByTimecreated,
	"displayname": ListPipelineSchemasSortByDisplayname,
}

// GetListPipelineSchemasSortByEnumValues Enumerates the set of values for ListPipelineSchemasSortByEnum
func GetListPipelineSchemasSortByEnumValues() []ListPipelineSchemasSortByEnum {
	values := make([]ListPipelineSchemasSortByEnum, 0)
	for _, v := range mappingListPipelineSchemasSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineSchemasSortByEnumStringValues Enumerates the set of values in String for ListPipelineSchemasSortByEnum
func GetListPipelineSchemasSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPipelineSchemasSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineSchemasSortByEnum(val string) (ListPipelineSchemasSortByEnum, bool) {
	enum, ok := mappingListPipelineSchemasSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
