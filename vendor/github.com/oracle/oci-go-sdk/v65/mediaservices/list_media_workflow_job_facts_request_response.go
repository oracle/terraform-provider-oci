// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMediaWorkflowJobFactsRequest wrapper for the ListMediaWorkflowJobFacts operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflowJobFacts.go.html to see an example of how to use ListMediaWorkflowJobFactsRequest.
type ListMediaWorkflowJobFactsRequest struct {

	// Unique MediaWorkflowJob identifier.
	MediaWorkflowJobId *string `mandatory:"true" contributesTo:"path" name:"mediaWorkflowJobId"`

	// Filter by MediaWorkflowJob ID and MediaWorkflowJobFact key.
	Key *int `mandatory:"false" contributesTo:"query" name:"key"`

	// Types of details to include.
	Type ListMediaWorkflowJobFactsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// Types of details to include.
	SortBy ListMediaWorkflowJobFactsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMediaWorkflowJobFactsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMediaWorkflowJobFactsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMediaWorkflowJobFactsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMediaWorkflowJobFactsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMediaWorkflowJobFactsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMediaWorkflowJobFactsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMediaWorkflowJobFactsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListMediaWorkflowJobFactsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowJobFactsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMediaWorkflowJobFactsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowJobFactsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMediaWorkflowJobFactsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMediaWorkflowJobFactsResponse wrapper for the ListMediaWorkflowJobFacts operation
type ListMediaWorkflowJobFactsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MediaWorkflowJobFactCollection instances
	MediaWorkflowJobFactCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMediaWorkflowJobFactsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMediaWorkflowJobFactsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMediaWorkflowJobFactsTypeEnum Enum with underlying type: string
type ListMediaWorkflowJobFactsTypeEnum string

// Set of constants representing the allowable values for ListMediaWorkflowJobFactsTypeEnum
const (
	ListMediaWorkflowJobFactsTypeRunnablejob              ListMediaWorkflowJobFactsTypeEnum = "runnableJob"
	ListMediaWorkflowJobFactsTypeTaskdeclaration          ListMediaWorkflowJobFactsTypeEnum = "taskDeclaration"
	ListMediaWorkflowJobFactsTypeWorkflow                 ListMediaWorkflowJobFactsTypeEnum = "workflow"
	ListMediaWorkflowJobFactsTypeConfiguration            ListMediaWorkflowJobFactsTypeEnum = "configuration"
	ListMediaWorkflowJobFactsTypeParameterresolutionevent ListMediaWorkflowJobFactsTypeEnum = "parameterResolutionEvent"
)

var mappingListMediaWorkflowJobFactsTypeEnum = map[string]ListMediaWorkflowJobFactsTypeEnum{
	"runnableJob":              ListMediaWorkflowJobFactsTypeRunnablejob,
	"taskDeclaration":          ListMediaWorkflowJobFactsTypeTaskdeclaration,
	"workflow":                 ListMediaWorkflowJobFactsTypeWorkflow,
	"configuration":            ListMediaWorkflowJobFactsTypeConfiguration,
	"parameterResolutionEvent": ListMediaWorkflowJobFactsTypeParameterresolutionevent,
}

var mappingListMediaWorkflowJobFactsTypeEnumLowerCase = map[string]ListMediaWorkflowJobFactsTypeEnum{
	"runnablejob":              ListMediaWorkflowJobFactsTypeRunnablejob,
	"taskdeclaration":          ListMediaWorkflowJobFactsTypeTaskdeclaration,
	"workflow":                 ListMediaWorkflowJobFactsTypeWorkflow,
	"configuration":            ListMediaWorkflowJobFactsTypeConfiguration,
	"parameterresolutionevent": ListMediaWorkflowJobFactsTypeParameterresolutionevent,
}

// GetListMediaWorkflowJobFactsTypeEnumValues Enumerates the set of values for ListMediaWorkflowJobFactsTypeEnum
func GetListMediaWorkflowJobFactsTypeEnumValues() []ListMediaWorkflowJobFactsTypeEnum {
	values := make([]ListMediaWorkflowJobFactsTypeEnum, 0)
	for _, v := range mappingListMediaWorkflowJobFactsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowJobFactsTypeEnumStringValues Enumerates the set of values in String for ListMediaWorkflowJobFactsTypeEnum
func GetListMediaWorkflowJobFactsTypeEnumStringValues() []string {
	return []string{
		"runnableJob",
		"taskDeclaration",
		"workflow",
		"configuration",
		"parameterResolutionEvent",
	}
}

// GetMappingListMediaWorkflowJobFactsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowJobFactsTypeEnum(val string) (ListMediaWorkflowJobFactsTypeEnum, bool) {
	enum, ok := mappingListMediaWorkflowJobFactsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaWorkflowJobFactsSortByEnum Enum with underlying type: string
type ListMediaWorkflowJobFactsSortByEnum string

// Set of constants representing the allowable values for ListMediaWorkflowJobFactsSortByEnum
const (
	ListMediaWorkflowJobFactsSortByKey ListMediaWorkflowJobFactsSortByEnum = "key"
)

var mappingListMediaWorkflowJobFactsSortByEnum = map[string]ListMediaWorkflowJobFactsSortByEnum{
	"key": ListMediaWorkflowJobFactsSortByKey,
}

var mappingListMediaWorkflowJobFactsSortByEnumLowerCase = map[string]ListMediaWorkflowJobFactsSortByEnum{
	"key": ListMediaWorkflowJobFactsSortByKey,
}

// GetListMediaWorkflowJobFactsSortByEnumValues Enumerates the set of values for ListMediaWorkflowJobFactsSortByEnum
func GetListMediaWorkflowJobFactsSortByEnumValues() []ListMediaWorkflowJobFactsSortByEnum {
	values := make([]ListMediaWorkflowJobFactsSortByEnum, 0)
	for _, v := range mappingListMediaWorkflowJobFactsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowJobFactsSortByEnumStringValues Enumerates the set of values in String for ListMediaWorkflowJobFactsSortByEnum
func GetListMediaWorkflowJobFactsSortByEnumStringValues() []string {
	return []string{
		"key",
	}
}

// GetMappingListMediaWorkflowJobFactsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowJobFactsSortByEnum(val string) (ListMediaWorkflowJobFactsSortByEnum, bool) {
	enum, ok := mappingListMediaWorkflowJobFactsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaWorkflowJobFactsSortOrderEnum Enum with underlying type: string
type ListMediaWorkflowJobFactsSortOrderEnum string

// Set of constants representing the allowable values for ListMediaWorkflowJobFactsSortOrderEnum
const (
	ListMediaWorkflowJobFactsSortOrderAsc  ListMediaWorkflowJobFactsSortOrderEnum = "ASC"
	ListMediaWorkflowJobFactsSortOrderDesc ListMediaWorkflowJobFactsSortOrderEnum = "DESC"
)

var mappingListMediaWorkflowJobFactsSortOrderEnum = map[string]ListMediaWorkflowJobFactsSortOrderEnum{
	"ASC":  ListMediaWorkflowJobFactsSortOrderAsc,
	"DESC": ListMediaWorkflowJobFactsSortOrderDesc,
}

var mappingListMediaWorkflowJobFactsSortOrderEnumLowerCase = map[string]ListMediaWorkflowJobFactsSortOrderEnum{
	"asc":  ListMediaWorkflowJobFactsSortOrderAsc,
	"desc": ListMediaWorkflowJobFactsSortOrderDesc,
}

// GetListMediaWorkflowJobFactsSortOrderEnumValues Enumerates the set of values for ListMediaWorkflowJobFactsSortOrderEnum
func GetListMediaWorkflowJobFactsSortOrderEnumValues() []ListMediaWorkflowJobFactsSortOrderEnum {
	values := make([]ListMediaWorkflowJobFactsSortOrderEnum, 0)
	for _, v := range mappingListMediaWorkflowJobFactsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowJobFactsSortOrderEnumStringValues Enumerates the set of values in String for ListMediaWorkflowJobFactsSortOrderEnum
func GetListMediaWorkflowJobFactsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMediaWorkflowJobFactsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowJobFactsSortOrderEnum(val string) (ListMediaWorkflowJobFactsSortOrderEnum, bool) {
	enum, ok := mappingListMediaWorkflowJobFactsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
