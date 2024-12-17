// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPipelineSchemaTablesRequest wrapper for the ListPipelineSchemaTables operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelineSchemaTables.go.html to see an example of how to use ListPipelineSchemaTablesRequest.
type ListPipelineSchemaTablesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the pipeline created.
	PipelineId *string `mandatory:"true" contributesTo:"path" name:"pipelineId"`

	// Name of the source schema obtained from get schema endpoint of the created pipeline.
	SourceSchemaName *string `mandatory:"true" contributesTo:"query" name:"sourceSchemaName"`

	// Name of the target schema obtained from get schema endpoint of the created pipeline.
	TargetSchemaName *string `mandatory:"true" contributesTo:"query" name:"targetSchemaName"`

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
	SortOrder ListPipelineSchemaTablesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListPipelineSchemaTablesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPipelineSchemaTablesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPipelineSchemaTablesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPipelineSchemaTablesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPipelineSchemaTablesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPipelineSchemaTablesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPipelineSchemaTablesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPipelineSchemaTablesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPipelineSchemaTablesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPipelineSchemaTablesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPipelineSchemaTablesResponse wrapper for the ListPipelineSchemaTables operation
type ListPipelineSchemaTablesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PipelineSchemaTableCollection instances
	PipelineSchemaTableCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPipelineSchemaTablesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPipelineSchemaTablesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPipelineSchemaTablesSortOrderEnum Enum with underlying type: string
type ListPipelineSchemaTablesSortOrderEnum string

// Set of constants representing the allowable values for ListPipelineSchemaTablesSortOrderEnum
const (
	ListPipelineSchemaTablesSortOrderAsc  ListPipelineSchemaTablesSortOrderEnum = "ASC"
	ListPipelineSchemaTablesSortOrderDesc ListPipelineSchemaTablesSortOrderEnum = "DESC"
)

var mappingListPipelineSchemaTablesSortOrderEnum = map[string]ListPipelineSchemaTablesSortOrderEnum{
	"ASC":  ListPipelineSchemaTablesSortOrderAsc,
	"DESC": ListPipelineSchemaTablesSortOrderDesc,
}

var mappingListPipelineSchemaTablesSortOrderEnumLowerCase = map[string]ListPipelineSchemaTablesSortOrderEnum{
	"asc":  ListPipelineSchemaTablesSortOrderAsc,
	"desc": ListPipelineSchemaTablesSortOrderDesc,
}

// GetListPipelineSchemaTablesSortOrderEnumValues Enumerates the set of values for ListPipelineSchemaTablesSortOrderEnum
func GetListPipelineSchemaTablesSortOrderEnumValues() []ListPipelineSchemaTablesSortOrderEnum {
	values := make([]ListPipelineSchemaTablesSortOrderEnum, 0)
	for _, v := range mappingListPipelineSchemaTablesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineSchemaTablesSortOrderEnumStringValues Enumerates the set of values in String for ListPipelineSchemaTablesSortOrderEnum
func GetListPipelineSchemaTablesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPipelineSchemaTablesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineSchemaTablesSortOrderEnum(val string) (ListPipelineSchemaTablesSortOrderEnum, bool) {
	enum, ok := mappingListPipelineSchemaTablesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPipelineSchemaTablesSortByEnum Enum with underlying type: string
type ListPipelineSchemaTablesSortByEnum string

// Set of constants representing the allowable values for ListPipelineSchemaTablesSortByEnum
const (
	ListPipelineSchemaTablesSortByTimecreated ListPipelineSchemaTablesSortByEnum = "timeCreated"
	ListPipelineSchemaTablesSortByDisplayname ListPipelineSchemaTablesSortByEnum = "displayName"
)

var mappingListPipelineSchemaTablesSortByEnum = map[string]ListPipelineSchemaTablesSortByEnum{
	"timeCreated": ListPipelineSchemaTablesSortByTimecreated,
	"displayName": ListPipelineSchemaTablesSortByDisplayname,
}

var mappingListPipelineSchemaTablesSortByEnumLowerCase = map[string]ListPipelineSchemaTablesSortByEnum{
	"timecreated": ListPipelineSchemaTablesSortByTimecreated,
	"displayname": ListPipelineSchemaTablesSortByDisplayname,
}

// GetListPipelineSchemaTablesSortByEnumValues Enumerates the set of values for ListPipelineSchemaTablesSortByEnum
func GetListPipelineSchemaTablesSortByEnumValues() []ListPipelineSchemaTablesSortByEnum {
	values := make([]ListPipelineSchemaTablesSortByEnum, 0)
	for _, v := range mappingListPipelineSchemaTablesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPipelineSchemaTablesSortByEnumStringValues Enumerates the set of values in String for ListPipelineSchemaTablesSortByEnum
func GetListPipelineSchemaTablesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPipelineSchemaTablesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPipelineSchemaTablesSortByEnum(val string) (ListPipelineSchemaTablesSortByEnum, bool) {
	enum, ok := mappingListPipelineSchemaTablesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
