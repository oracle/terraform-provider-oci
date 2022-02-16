// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waf

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListProtectionCapabilityGroupTagsRequest wrapper for the ListProtectionCapabilityGroupTags operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListProtectionCapabilityGroupTags.go.html to see an example of how to use ListProtectionCapabilityGroupTagsRequest.
type ListProtectionCapabilityGroupTagsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to return only resources that matches given type.
	Type ProtectionCapabilitySummaryTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListProtectionCapabilityGroupTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for name is ascending.
	// If no value is specified name is default.
	SortBy ListProtectionCapabilityGroupTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProtectionCapabilityGroupTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProtectionCapabilityGroupTagsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProtectionCapabilityGroupTagsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProtectionCapabilityGroupTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProtectionCapabilityGroupTagsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProtectionCapabilitySummaryTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetProtectionCapabilitySummaryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectionCapabilityGroupTagsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProtectionCapabilityGroupTagsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectionCapabilityGroupTagsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProtectionCapabilityGroupTagsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProtectionCapabilityGroupTagsResponse wrapper for the ListProtectionCapabilityGroupTags operation
type ListProtectionCapabilityGroupTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProtectionCapabilityGroupTagCollection instances
	ProtectionCapabilityGroupTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProtectionCapabilityGroupTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProtectionCapabilityGroupTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProtectionCapabilityGroupTagsSortOrderEnum Enum with underlying type: string
type ListProtectionCapabilityGroupTagsSortOrderEnum string

// Set of constants representing the allowable values for ListProtectionCapabilityGroupTagsSortOrderEnum
const (
	ListProtectionCapabilityGroupTagsSortOrderAsc  ListProtectionCapabilityGroupTagsSortOrderEnum = "ASC"
	ListProtectionCapabilityGroupTagsSortOrderDesc ListProtectionCapabilityGroupTagsSortOrderEnum = "DESC"
)

var mappingListProtectionCapabilityGroupTagsSortOrderEnum = map[string]ListProtectionCapabilityGroupTagsSortOrderEnum{
	"ASC":  ListProtectionCapabilityGroupTagsSortOrderAsc,
	"DESC": ListProtectionCapabilityGroupTagsSortOrderDesc,
}

// GetListProtectionCapabilityGroupTagsSortOrderEnumValues Enumerates the set of values for ListProtectionCapabilityGroupTagsSortOrderEnum
func GetListProtectionCapabilityGroupTagsSortOrderEnumValues() []ListProtectionCapabilityGroupTagsSortOrderEnum {
	values := make([]ListProtectionCapabilityGroupTagsSortOrderEnum, 0)
	for _, v := range mappingListProtectionCapabilityGroupTagsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionCapabilityGroupTagsSortOrderEnumStringValues Enumerates the set of values in String for ListProtectionCapabilityGroupTagsSortOrderEnum
func GetListProtectionCapabilityGroupTagsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProtectionCapabilityGroupTagsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionCapabilityGroupTagsSortOrderEnum(val string) (ListProtectionCapabilityGroupTagsSortOrderEnum, bool) {
	mappingListProtectionCapabilityGroupTagsSortOrderEnumIgnoreCase := make(map[string]ListProtectionCapabilityGroupTagsSortOrderEnum)
	for k, v := range mappingListProtectionCapabilityGroupTagsSortOrderEnum {
		mappingListProtectionCapabilityGroupTagsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListProtectionCapabilityGroupTagsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListProtectionCapabilityGroupTagsSortByEnum Enum with underlying type: string
type ListProtectionCapabilityGroupTagsSortByEnum string

// Set of constants representing the allowable values for ListProtectionCapabilityGroupTagsSortByEnum
const (
	ListProtectionCapabilityGroupTagsSortByName ListProtectionCapabilityGroupTagsSortByEnum = "name"
)

var mappingListProtectionCapabilityGroupTagsSortByEnum = map[string]ListProtectionCapabilityGroupTagsSortByEnum{
	"name": ListProtectionCapabilityGroupTagsSortByName,
}

// GetListProtectionCapabilityGroupTagsSortByEnumValues Enumerates the set of values for ListProtectionCapabilityGroupTagsSortByEnum
func GetListProtectionCapabilityGroupTagsSortByEnumValues() []ListProtectionCapabilityGroupTagsSortByEnum {
	values := make([]ListProtectionCapabilityGroupTagsSortByEnum, 0)
	for _, v := range mappingListProtectionCapabilityGroupTagsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionCapabilityGroupTagsSortByEnumStringValues Enumerates the set of values in String for ListProtectionCapabilityGroupTagsSortByEnum
func GetListProtectionCapabilityGroupTagsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListProtectionCapabilityGroupTagsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionCapabilityGroupTagsSortByEnum(val string) (ListProtectionCapabilityGroupTagsSortByEnum, bool) {
	mappingListProtectionCapabilityGroupTagsSortByEnumIgnoreCase := make(map[string]ListProtectionCapabilityGroupTagsSortByEnum)
	for k, v := range mappingListProtectionCapabilityGroupTagsSortByEnum {
		mappingListProtectionCapabilityGroupTagsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListProtectionCapabilityGroupTagsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
