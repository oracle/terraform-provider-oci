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

// ListAssociatedEntitiesRequest wrapper for the ListAssociatedEntities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListAssociatedEntities.go.html to see an example of how to use ListAssociatedEntitiesRequest.
type ListAssociatedEntitiesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The entity OCID.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// The entity type used for filtering.  Only associations on an entity with the
	// specified type will be returned.
	EntityType *string `mandatory:"false" contributesTo:"query" name:"entityType"`

	// The entity type display name used for filtering.  Only items associated with the entity
	// with the specified type display name will be returned.
	EntityTypeDisplayName *string `mandatory:"false" contributesTo:"query" name:"entityTypeDisplayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAssociatedEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned entities
	SortBy ListAssociatedEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociatedEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociatedEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociatedEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociatedEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssociatedEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssociatedEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssociatedEntitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssociatedEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssociatedEntitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssociatedEntitiesResponse wrapper for the ListAssociatedEntities operation
type ListAssociatedEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsAssociatedEntityCollection instances
	LogAnalyticsAssociatedEntityCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAssociatedEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociatedEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociatedEntitiesSortOrderEnum Enum with underlying type: string
type ListAssociatedEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListAssociatedEntitiesSortOrderEnum
const (
	ListAssociatedEntitiesSortOrderAsc  ListAssociatedEntitiesSortOrderEnum = "ASC"
	ListAssociatedEntitiesSortOrderDesc ListAssociatedEntitiesSortOrderEnum = "DESC"
)

var mappingListAssociatedEntitiesSortOrderEnum = map[string]ListAssociatedEntitiesSortOrderEnum{
	"ASC":  ListAssociatedEntitiesSortOrderAsc,
	"DESC": ListAssociatedEntitiesSortOrderDesc,
}

// GetListAssociatedEntitiesSortOrderEnumValues Enumerates the set of values for ListAssociatedEntitiesSortOrderEnum
func GetListAssociatedEntitiesSortOrderEnumValues() []ListAssociatedEntitiesSortOrderEnum {
	values := make([]ListAssociatedEntitiesSortOrderEnum, 0)
	for _, v := range mappingListAssociatedEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociatedEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListAssociatedEntitiesSortOrderEnum
func GetListAssociatedEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssociatedEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociatedEntitiesSortOrderEnum(val string) (ListAssociatedEntitiesSortOrderEnum, bool) {
	mappingListAssociatedEntitiesSortOrderEnumIgnoreCase := make(map[string]ListAssociatedEntitiesSortOrderEnum)
	for k, v := range mappingListAssociatedEntitiesSortOrderEnum {
		mappingListAssociatedEntitiesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAssociatedEntitiesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssociatedEntitiesSortByEnum Enum with underlying type: string
type ListAssociatedEntitiesSortByEnum string

// Set of constants representing the allowable values for ListAssociatedEntitiesSortByEnum
const (
	ListAssociatedEntitiesSortByEntityname            ListAssociatedEntitiesSortByEnum = "entityName"
	ListAssociatedEntitiesSortByEntitytypedisplayname ListAssociatedEntitiesSortByEnum = "entityTypeDisplayName"
	ListAssociatedEntitiesSortByAssociationcount      ListAssociatedEntitiesSortByEnum = "associationCount"
)

var mappingListAssociatedEntitiesSortByEnum = map[string]ListAssociatedEntitiesSortByEnum{
	"entityName":            ListAssociatedEntitiesSortByEntityname,
	"entityTypeDisplayName": ListAssociatedEntitiesSortByEntitytypedisplayname,
	"associationCount":      ListAssociatedEntitiesSortByAssociationcount,
}

// GetListAssociatedEntitiesSortByEnumValues Enumerates the set of values for ListAssociatedEntitiesSortByEnum
func GetListAssociatedEntitiesSortByEnumValues() []ListAssociatedEntitiesSortByEnum {
	values := make([]ListAssociatedEntitiesSortByEnum, 0)
	for _, v := range mappingListAssociatedEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociatedEntitiesSortByEnumStringValues Enumerates the set of values in String for ListAssociatedEntitiesSortByEnum
func GetListAssociatedEntitiesSortByEnumStringValues() []string {
	return []string{
		"entityName",
		"entityTypeDisplayName",
		"associationCount",
	}
}

// GetMappingListAssociatedEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociatedEntitiesSortByEnum(val string) (ListAssociatedEntitiesSortByEnum, bool) {
	mappingListAssociatedEntitiesSortByEnumIgnoreCase := make(map[string]ListAssociatedEntitiesSortByEnum)
	for k, v := range mappingListAssociatedEntitiesSortByEnum {
		mappingListAssociatedEntitiesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAssociatedEntitiesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
