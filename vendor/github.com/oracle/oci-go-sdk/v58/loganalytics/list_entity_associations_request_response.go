// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListEntityAssociationsRequest wrapper for the ListEntityAssociations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListEntityAssociations.go.html to see an example of how to use ListEntityAssociationsRequest.
type ListEntityAssociationsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The log analytics entity OCID.
	LogAnalyticsEntityId *string `mandatory:"true" contributesTo:"path" name:"logAnalyticsEntityId"`

	// Indicates whether to return direct associated entities or direct and inferred associated entities.
	DirectOrAllAssociations ListEntityAssociationsDirectOrAllAssociationsEnum `mandatory:"false" contributesTo:"query" name:"directOrAllAssociations" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListEntityAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort entities by. Only one sort order may be provided. Default order for timeCreated and timeUpdated
	// is descending. Default order for entity name is ascending. If no value is specified timeCreated is default.
	SortBy ListEntityAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEntityAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEntityAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEntityAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEntityAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEntityAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEntityAssociationsDirectOrAllAssociationsEnum(string(request.DirectOrAllAssociations)); !ok && request.DirectOrAllAssociations != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DirectOrAllAssociations: %s. Supported values are: %s.", request.DirectOrAllAssociations, strings.Join(GetListEntityAssociationsDirectOrAllAssociationsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntityAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEntityAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntityAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEntityAssociationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEntityAssociationsResponse wrapper for the ListEntityAssociations operation
type ListEntityAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsEntityCollection instances
	LogAnalyticsEntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEntityAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEntityAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEntityAssociationsDirectOrAllAssociationsEnum Enum with underlying type: string
type ListEntityAssociationsDirectOrAllAssociationsEnum string

// Set of constants representing the allowable values for ListEntityAssociationsDirectOrAllAssociationsEnum
const (
	ListEntityAssociationsDirectOrAllAssociationsDirect ListEntityAssociationsDirectOrAllAssociationsEnum = "DIRECT"
	ListEntityAssociationsDirectOrAllAssociationsAll    ListEntityAssociationsDirectOrAllAssociationsEnum = "ALL"
)

var mappingListEntityAssociationsDirectOrAllAssociationsEnum = map[string]ListEntityAssociationsDirectOrAllAssociationsEnum{
	"DIRECT": ListEntityAssociationsDirectOrAllAssociationsDirect,
	"ALL":    ListEntityAssociationsDirectOrAllAssociationsAll,
}

// GetListEntityAssociationsDirectOrAllAssociationsEnumValues Enumerates the set of values for ListEntityAssociationsDirectOrAllAssociationsEnum
func GetListEntityAssociationsDirectOrAllAssociationsEnumValues() []ListEntityAssociationsDirectOrAllAssociationsEnum {
	values := make([]ListEntityAssociationsDirectOrAllAssociationsEnum, 0)
	for _, v := range mappingListEntityAssociationsDirectOrAllAssociationsEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntityAssociationsDirectOrAllAssociationsEnumStringValues Enumerates the set of values in String for ListEntityAssociationsDirectOrAllAssociationsEnum
func GetListEntityAssociationsDirectOrAllAssociationsEnumStringValues() []string {
	return []string{
		"DIRECT",
		"ALL",
	}
}

// GetMappingListEntityAssociationsDirectOrAllAssociationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntityAssociationsDirectOrAllAssociationsEnum(val string) (ListEntityAssociationsDirectOrAllAssociationsEnum, bool) {
	mappingListEntityAssociationsDirectOrAllAssociationsEnumIgnoreCase := make(map[string]ListEntityAssociationsDirectOrAllAssociationsEnum)
	for k, v := range mappingListEntityAssociationsDirectOrAllAssociationsEnum {
		mappingListEntityAssociationsDirectOrAllAssociationsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntityAssociationsDirectOrAllAssociationsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntityAssociationsSortOrderEnum Enum with underlying type: string
type ListEntityAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListEntityAssociationsSortOrderEnum
const (
	ListEntityAssociationsSortOrderAsc  ListEntityAssociationsSortOrderEnum = "ASC"
	ListEntityAssociationsSortOrderDesc ListEntityAssociationsSortOrderEnum = "DESC"
)

var mappingListEntityAssociationsSortOrderEnum = map[string]ListEntityAssociationsSortOrderEnum{
	"ASC":  ListEntityAssociationsSortOrderAsc,
	"DESC": ListEntityAssociationsSortOrderDesc,
}

// GetListEntityAssociationsSortOrderEnumValues Enumerates the set of values for ListEntityAssociationsSortOrderEnum
func GetListEntityAssociationsSortOrderEnumValues() []ListEntityAssociationsSortOrderEnum {
	values := make([]ListEntityAssociationsSortOrderEnum, 0)
	for _, v := range mappingListEntityAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntityAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListEntityAssociationsSortOrderEnum
func GetListEntityAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEntityAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntityAssociationsSortOrderEnum(val string) (ListEntityAssociationsSortOrderEnum, bool) {
	mappingListEntityAssociationsSortOrderEnumIgnoreCase := make(map[string]ListEntityAssociationsSortOrderEnum)
	for k, v := range mappingListEntityAssociationsSortOrderEnum {
		mappingListEntityAssociationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntityAssociationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntityAssociationsSortByEnum Enum with underlying type: string
type ListEntityAssociationsSortByEnum string

// Set of constants representing the allowable values for ListEntityAssociationsSortByEnum
const (
	ListEntityAssociationsSortByTimecreated ListEntityAssociationsSortByEnum = "timeCreated"
	ListEntityAssociationsSortByTimeupdated ListEntityAssociationsSortByEnum = "timeUpdated"
	ListEntityAssociationsSortByName        ListEntityAssociationsSortByEnum = "name"
)

var mappingListEntityAssociationsSortByEnum = map[string]ListEntityAssociationsSortByEnum{
	"timeCreated": ListEntityAssociationsSortByTimecreated,
	"timeUpdated": ListEntityAssociationsSortByTimeupdated,
	"name":        ListEntityAssociationsSortByName,
}

// GetListEntityAssociationsSortByEnumValues Enumerates the set of values for ListEntityAssociationsSortByEnum
func GetListEntityAssociationsSortByEnumValues() []ListEntityAssociationsSortByEnum {
	values := make([]ListEntityAssociationsSortByEnum, 0)
	for _, v := range mappingListEntityAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntityAssociationsSortByEnumStringValues Enumerates the set of values in String for ListEntityAssociationsSortByEnum
func GetListEntityAssociationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
	}
}

// GetMappingListEntityAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntityAssociationsSortByEnum(val string) (ListEntityAssociationsSortByEnum, bool) {
	mappingListEntityAssociationsSortByEnumIgnoreCase := make(map[string]ListEntityAssociationsSortByEnum)
	for k, v := range mappingListEntityAssociationsSortByEnum {
		mappingListEntityAssociationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntityAssociationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
