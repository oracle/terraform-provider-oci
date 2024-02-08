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

// ListMediaWorkflowTaskDeclarationsRequest wrapper for the ListMediaWorkflowTaskDeclarations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflowTaskDeclarations.go.html to see an example of how to use ListMediaWorkflowTaskDeclarationsRequest.
type ListMediaWorkflowTaskDeclarationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources with their system defined, unique name matching the given name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to select MediaWorkflowTaskDeclaration by version.
	Version *int `mandatory:"false" contributesTo:"query" name:"version"`

	// A filter to only select the newest version for each MediaWorkflowTaskDeclaration name.
	IsCurrent *bool `mandatory:"false" contributesTo:"query" name:"isCurrent"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListMediaWorkflowTaskDeclarationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMediaWorkflowTaskDeclarationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMediaWorkflowTaskDeclarationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMediaWorkflowTaskDeclarationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMediaWorkflowTaskDeclarationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMediaWorkflowTaskDeclarationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMediaWorkflowTaskDeclarationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMediaWorkflowTaskDeclarationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMediaWorkflowTaskDeclarationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowTaskDeclarationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMediaWorkflowTaskDeclarationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMediaWorkflowTaskDeclarationsResponse wrapper for the ListMediaWorkflowTaskDeclarations operation
type ListMediaWorkflowTaskDeclarationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MediaWorkflowTaskDeclarationCollection instances
	MediaWorkflowTaskDeclarationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMediaWorkflowTaskDeclarationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMediaWorkflowTaskDeclarationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMediaWorkflowTaskDeclarationsSortByEnum Enum with underlying type: string
type ListMediaWorkflowTaskDeclarationsSortByEnum string

// Set of constants representing the allowable values for ListMediaWorkflowTaskDeclarationsSortByEnum
const (
	ListMediaWorkflowTaskDeclarationsSortByName    ListMediaWorkflowTaskDeclarationsSortByEnum = "name"
	ListMediaWorkflowTaskDeclarationsSortByVersion ListMediaWorkflowTaskDeclarationsSortByEnum = "version"
)

var mappingListMediaWorkflowTaskDeclarationsSortByEnum = map[string]ListMediaWorkflowTaskDeclarationsSortByEnum{
	"name":    ListMediaWorkflowTaskDeclarationsSortByName,
	"version": ListMediaWorkflowTaskDeclarationsSortByVersion,
}

var mappingListMediaWorkflowTaskDeclarationsSortByEnumLowerCase = map[string]ListMediaWorkflowTaskDeclarationsSortByEnum{
	"name":    ListMediaWorkflowTaskDeclarationsSortByName,
	"version": ListMediaWorkflowTaskDeclarationsSortByVersion,
}

// GetListMediaWorkflowTaskDeclarationsSortByEnumValues Enumerates the set of values for ListMediaWorkflowTaskDeclarationsSortByEnum
func GetListMediaWorkflowTaskDeclarationsSortByEnumValues() []ListMediaWorkflowTaskDeclarationsSortByEnum {
	values := make([]ListMediaWorkflowTaskDeclarationsSortByEnum, 0)
	for _, v := range mappingListMediaWorkflowTaskDeclarationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowTaskDeclarationsSortByEnumStringValues Enumerates the set of values in String for ListMediaWorkflowTaskDeclarationsSortByEnum
func GetListMediaWorkflowTaskDeclarationsSortByEnumStringValues() []string {
	return []string{
		"name",
		"version",
	}
}

// GetMappingListMediaWorkflowTaskDeclarationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowTaskDeclarationsSortByEnum(val string) (ListMediaWorkflowTaskDeclarationsSortByEnum, bool) {
	enum, ok := mappingListMediaWorkflowTaskDeclarationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaWorkflowTaskDeclarationsSortOrderEnum Enum with underlying type: string
type ListMediaWorkflowTaskDeclarationsSortOrderEnum string

// Set of constants representing the allowable values for ListMediaWorkflowTaskDeclarationsSortOrderEnum
const (
	ListMediaWorkflowTaskDeclarationsSortOrderAsc  ListMediaWorkflowTaskDeclarationsSortOrderEnum = "ASC"
	ListMediaWorkflowTaskDeclarationsSortOrderDesc ListMediaWorkflowTaskDeclarationsSortOrderEnum = "DESC"
)

var mappingListMediaWorkflowTaskDeclarationsSortOrderEnum = map[string]ListMediaWorkflowTaskDeclarationsSortOrderEnum{
	"ASC":  ListMediaWorkflowTaskDeclarationsSortOrderAsc,
	"DESC": ListMediaWorkflowTaskDeclarationsSortOrderDesc,
}

var mappingListMediaWorkflowTaskDeclarationsSortOrderEnumLowerCase = map[string]ListMediaWorkflowTaskDeclarationsSortOrderEnum{
	"asc":  ListMediaWorkflowTaskDeclarationsSortOrderAsc,
	"desc": ListMediaWorkflowTaskDeclarationsSortOrderDesc,
}

// GetListMediaWorkflowTaskDeclarationsSortOrderEnumValues Enumerates the set of values for ListMediaWorkflowTaskDeclarationsSortOrderEnum
func GetListMediaWorkflowTaskDeclarationsSortOrderEnumValues() []ListMediaWorkflowTaskDeclarationsSortOrderEnum {
	values := make([]ListMediaWorkflowTaskDeclarationsSortOrderEnum, 0)
	for _, v := range mappingListMediaWorkflowTaskDeclarationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowTaskDeclarationsSortOrderEnumStringValues Enumerates the set of values in String for ListMediaWorkflowTaskDeclarationsSortOrderEnum
func GetListMediaWorkflowTaskDeclarationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMediaWorkflowTaskDeclarationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowTaskDeclarationsSortOrderEnum(val string) (ListMediaWorkflowTaskDeclarationsSortOrderEnum, bool) {
	enum, ok := mappingListMediaWorkflowTaskDeclarationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
