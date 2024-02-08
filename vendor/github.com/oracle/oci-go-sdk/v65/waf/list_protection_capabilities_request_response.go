// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waf

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProtectionCapabilitiesRequest wrapper for the ListProtectionCapabilities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waf/ListProtectionCapabilities.go.html to see an example of how to use ListProtectionCapabilitiesRequest.
type ListProtectionCapabilitiesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The unique key of protection capability to filter by.
	Key *string `mandatory:"false" contributesTo:"query" name:"key"`

	// A filter to return only resources that matches given isLatestVersion.
	IsLatestVersion []bool `contributesTo:"query" name:"isLatestVersion" collectionFormat:"multi"`

	// A filter to return only resources that matches given type.
	Type ProtectionCapabilitySummaryTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter to return only resources that are accociated given group tag.
	GroupTag []string `contributesTo:"query" name:"groupTag" collectionFormat:"multi"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListProtectionCapabilitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for key is descending.
	// Default order for type is descending.
	// Default order for displayName is ascending.
	// If no value is specified key is default.
	SortBy ListProtectionCapabilitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProtectionCapabilitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProtectionCapabilitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProtectionCapabilitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProtectionCapabilitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProtectionCapabilitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProtectionCapabilitySummaryTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetProtectionCapabilitySummaryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectionCapabilitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProtectionCapabilitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProtectionCapabilitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProtectionCapabilitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProtectionCapabilitiesResponse wrapper for the ListProtectionCapabilities operation
type ListProtectionCapabilitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProtectionCapabilityCollection instances
	ProtectionCapabilityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProtectionCapabilitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProtectionCapabilitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProtectionCapabilitiesSortOrderEnum Enum with underlying type: string
type ListProtectionCapabilitiesSortOrderEnum string

// Set of constants representing the allowable values for ListProtectionCapabilitiesSortOrderEnum
const (
	ListProtectionCapabilitiesSortOrderAsc  ListProtectionCapabilitiesSortOrderEnum = "ASC"
	ListProtectionCapabilitiesSortOrderDesc ListProtectionCapabilitiesSortOrderEnum = "DESC"
)

var mappingListProtectionCapabilitiesSortOrderEnum = map[string]ListProtectionCapabilitiesSortOrderEnum{
	"ASC":  ListProtectionCapabilitiesSortOrderAsc,
	"DESC": ListProtectionCapabilitiesSortOrderDesc,
}

var mappingListProtectionCapabilitiesSortOrderEnumLowerCase = map[string]ListProtectionCapabilitiesSortOrderEnum{
	"asc":  ListProtectionCapabilitiesSortOrderAsc,
	"desc": ListProtectionCapabilitiesSortOrderDesc,
}

// GetListProtectionCapabilitiesSortOrderEnumValues Enumerates the set of values for ListProtectionCapabilitiesSortOrderEnum
func GetListProtectionCapabilitiesSortOrderEnumValues() []ListProtectionCapabilitiesSortOrderEnum {
	values := make([]ListProtectionCapabilitiesSortOrderEnum, 0)
	for _, v := range mappingListProtectionCapabilitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionCapabilitiesSortOrderEnumStringValues Enumerates the set of values in String for ListProtectionCapabilitiesSortOrderEnum
func GetListProtectionCapabilitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProtectionCapabilitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionCapabilitiesSortOrderEnum(val string) (ListProtectionCapabilitiesSortOrderEnum, bool) {
	enum, ok := mappingListProtectionCapabilitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProtectionCapabilitiesSortByEnum Enum with underlying type: string
type ListProtectionCapabilitiesSortByEnum string

// Set of constants representing the allowable values for ListProtectionCapabilitiesSortByEnum
const (
	ListProtectionCapabilitiesSortByKey         ListProtectionCapabilitiesSortByEnum = "key"
	ListProtectionCapabilitiesSortByType        ListProtectionCapabilitiesSortByEnum = "type"
	ListProtectionCapabilitiesSortByDisplayname ListProtectionCapabilitiesSortByEnum = "displayName"
)

var mappingListProtectionCapabilitiesSortByEnum = map[string]ListProtectionCapabilitiesSortByEnum{
	"key":         ListProtectionCapabilitiesSortByKey,
	"type":        ListProtectionCapabilitiesSortByType,
	"displayName": ListProtectionCapabilitiesSortByDisplayname,
}

var mappingListProtectionCapabilitiesSortByEnumLowerCase = map[string]ListProtectionCapabilitiesSortByEnum{
	"key":         ListProtectionCapabilitiesSortByKey,
	"type":        ListProtectionCapabilitiesSortByType,
	"displayname": ListProtectionCapabilitiesSortByDisplayname,
}

// GetListProtectionCapabilitiesSortByEnumValues Enumerates the set of values for ListProtectionCapabilitiesSortByEnum
func GetListProtectionCapabilitiesSortByEnumValues() []ListProtectionCapabilitiesSortByEnum {
	values := make([]ListProtectionCapabilitiesSortByEnum, 0)
	for _, v := range mappingListProtectionCapabilitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProtectionCapabilitiesSortByEnumStringValues Enumerates the set of values in String for ListProtectionCapabilitiesSortByEnum
func GetListProtectionCapabilitiesSortByEnumStringValues() []string {
	return []string{
		"key",
		"type",
		"displayName",
	}
}

// GetMappingListProtectionCapabilitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProtectionCapabilitiesSortByEnum(val string) (ListProtectionCapabilitiesSortByEnum, bool) {
	enum, ok := mappingListProtectionCapabilitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
