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

// ListAssociableEntitiesRequest wrapper for the ListAssociableEntities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListAssociableEntities.go.html to see an example of how to use ListAssociableEntitiesRequest.
type ListAssociableEntitiesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The source name.
	SourceName *string `mandatory:"true" contributesTo:"path" name:"sourceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The entity type - either eligible or ineligible for association.
	Type ListAssociableEntitiesTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The text used for filtering returned entities.  Search text is applicable to the
	// entity name and the entity description.
	SearchText *string `mandatory:"false" contributesTo:"query" name:"searchText"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The attribute used to sort the returned entities
	SortBy ListAssociableEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAssociableEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociableEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociableEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociableEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociableEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssociableEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAssociableEntitiesTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListAssociableEntitiesTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssociableEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssociableEntitiesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssociableEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssociableEntitiesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssociableEntitiesResponse wrapper for the ListAssociableEntities operation
type ListAssociableEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssociableEntityCollection instances
	AssociableEntityCollection `presentIn:"body"`

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

func (response ListAssociableEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociableEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociableEntitiesTypeEnum Enum with underlying type: string
type ListAssociableEntitiesTypeEnum string

// Set of constants representing the allowable values for ListAssociableEntitiesTypeEnum
const (
	ListAssociableEntitiesTypeEligible   ListAssociableEntitiesTypeEnum = "ELIGIBLE"
	ListAssociableEntitiesTypeIneligible ListAssociableEntitiesTypeEnum = "INELIGIBLE"
)

var mappingListAssociableEntitiesTypeEnum = map[string]ListAssociableEntitiesTypeEnum{
	"ELIGIBLE":   ListAssociableEntitiesTypeEligible,
	"INELIGIBLE": ListAssociableEntitiesTypeIneligible,
}

// GetListAssociableEntitiesTypeEnumValues Enumerates the set of values for ListAssociableEntitiesTypeEnum
func GetListAssociableEntitiesTypeEnumValues() []ListAssociableEntitiesTypeEnum {
	values := make([]ListAssociableEntitiesTypeEnum, 0)
	for _, v := range mappingListAssociableEntitiesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociableEntitiesTypeEnumStringValues Enumerates the set of values in String for ListAssociableEntitiesTypeEnum
func GetListAssociableEntitiesTypeEnumStringValues() []string {
	return []string{
		"ELIGIBLE",
		"INELIGIBLE",
	}
}

// GetMappingListAssociableEntitiesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociableEntitiesTypeEnum(val string) (ListAssociableEntitiesTypeEnum, bool) {
	mappingListAssociableEntitiesTypeEnumIgnoreCase := make(map[string]ListAssociableEntitiesTypeEnum)
	for k, v := range mappingListAssociableEntitiesTypeEnum {
		mappingListAssociableEntitiesTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAssociableEntitiesTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssociableEntitiesSortByEnum Enum with underlying type: string
type ListAssociableEntitiesSortByEnum string

// Set of constants representing the allowable values for ListAssociableEntitiesSortByEnum
const (
	ListAssociableEntitiesSortByEntityname     ListAssociableEntitiesSortByEnum = "entityName"
	ListAssociableEntitiesSortByEntitytypename ListAssociableEntitiesSortByEnum = "entityTypeName"
	ListAssociableEntitiesSortByHost           ListAssociableEntitiesSortByEnum = "host"
	ListAssociableEntitiesSortByAgentid        ListAssociableEntitiesSortByEnum = "agentId"
)

var mappingListAssociableEntitiesSortByEnum = map[string]ListAssociableEntitiesSortByEnum{
	"entityName":     ListAssociableEntitiesSortByEntityname,
	"entityTypeName": ListAssociableEntitiesSortByEntitytypename,
	"host":           ListAssociableEntitiesSortByHost,
	"agentId":        ListAssociableEntitiesSortByAgentid,
}

// GetListAssociableEntitiesSortByEnumValues Enumerates the set of values for ListAssociableEntitiesSortByEnum
func GetListAssociableEntitiesSortByEnumValues() []ListAssociableEntitiesSortByEnum {
	values := make([]ListAssociableEntitiesSortByEnum, 0)
	for _, v := range mappingListAssociableEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociableEntitiesSortByEnumStringValues Enumerates the set of values in String for ListAssociableEntitiesSortByEnum
func GetListAssociableEntitiesSortByEnumStringValues() []string {
	return []string{
		"entityName",
		"entityTypeName",
		"host",
		"agentId",
	}
}

// GetMappingListAssociableEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociableEntitiesSortByEnum(val string) (ListAssociableEntitiesSortByEnum, bool) {
	mappingListAssociableEntitiesSortByEnumIgnoreCase := make(map[string]ListAssociableEntitiesSortByEnum)
	for k, v := range mappingListAssociableEntitiesSortByEnum {
		mappingListAssociableEntitiesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAssociableEntitiesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssociableEntitiesSortOrderEnum Enum with underlying type: string
type ListAssociableEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListAssociableEntitiesSortOrderEnum
const (
	ListAssociableEntitiesSortOrderAsc  ListAssociableEntitiesSortOrderEnum = "ASC"
	ListAssociableEntitiesSortOrderDesc ListAssociableEntitiesSortOrderEnum = "DESC"
)

var mappingListAssociableEntitiesSortOrderEnum = map[string]ListAssociableEntitiesSortOrderEnum{
	"ASC":  ListAssociableEntitiesSortOrderAsc,
	"DESC": ListAssociableEntitiesSortOrderDesc,
}

// GetListAssociableEntitiesSortOrderEnumValues Enumerates the set of values for ListAssociableEntitiesSortOrderEnum
func GetListAssociableEntitiesSortOrderEnumValues() []ListAssociableEntitiesSortOrderEnum {
	values := make([]ListAssociableEntitiesSortOrderEnum, 0)
	for _, v := range mappingListAssociableEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssociableEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListAssociableEntitiesSortOrderEnum
func GetListAssociableEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssociableEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssociableEntitiesSortOrderEnum(val string) (ListAssociableEntitiesSortOrderEnum, bool) {
	mappingListAssociableEntitiesSortOrderEnumIgnoreCase := make(map[string]ListAssociableEntitiesSortOrderEnum)
	for k, v := range mappingListAssociableEntitiesSortOrderEnum {
		mappingListAssociableEntitiesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAssociableEntitiesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
