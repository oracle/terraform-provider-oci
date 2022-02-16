// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListServiceCatalogAssociationsRequest wrapper for the ListServiceCatalogAssociations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListServiceCatalogAssociations.go.html to see an example of how to use ListServiceCatalogAssociationsRequest.
type ListServiceCatalogAssociationsRequest struct {

	// The unique identifier for the service catalog association.
	ServiceCatalogAssociationId *string `mandatory:"false" contributesTo:"query" name:"serviceCatalogAssociationId"`

	// The unique identifier for the service catalog.
	ServiceCatalogId *string `mandatory:"false" contributesTo:"query" name:"serviceCatalogId"`

	// The unique identifier of the entity associated with service catalog.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// The type of the application in the service catalog.
	EntityType *string `mandatory:"false" contributesTo:"query" name:"entityType"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to apply, either `ASC` or `DESC`. Default is `ASC`.
	SortOrder ListServiceCatalogAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Default is `TIMECREATED`
	SortBy ListServiceCatalogAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServiceCatalogAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServiceCatalogAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListServiceCatalogAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServiceCatalogAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListServiceCatalogAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListServiceCatalogAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListServiceCatalogAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceCatalogAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListServiceCatalogAssociationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListServiceCatalogAssociationsResponse wrapper for the ListServiceCatalogAssociations operation
type ListServiceCatalogAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServiceCatalogAssociationCollection instances
	ServiceCatalogAssociationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListServiceCatalogAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServiceCatalogAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServiceCatalogAssociationsSortOrderEnum Enum with underlying type: string
type ListServiceCatalogAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceCatalogAssociationsSortOrderEnum
const (
	ListServiceCatalogAssociationsSortOrderAsc  ListServiceCatalogAssociationsSortOrderEnum = "ASC"
	ListServiceCatalogAssociationsSortOrderDesc ListServiceCatalogAssociationsSortOrderEnum = "DESC"
)

var mappingListServiceCatalogAssociationsSortOrderEnum = map[string]ListServiceCatalogAssociationsSortOrderEnum{
	"ASC":  ListServiceCatalogAssociationsSortOrderAsc,
	"DESC": ListServiceCatalogAssociationsSortOrderDesc,
}

// GetListServiceCatalogAssociationsSortOrderEnumValues Enumerates the set of values for ListServiceCatalogAssociationsSortOrderEnum
func GetListServiceCatalogAssociationsSortOrderEnumValues() []ListServiceCatalogAssociationsSortOrderEnum {
	values := make([]ListServiceCatalogAssociationsSortOrderEnum, 0)
	for _, v := range mappingListServiceCatalogAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceCatalogAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListServiceCatalogAssociationsSortOrderEnum
func GetListServiceCatalogAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListServiceCatalogAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceCatalogAssociationsSortOrderEnum(val string) (ListServiceCatalogAssociationsSortOrderEnum, bool) {
	mappingListServiceCatalogAssociationsSortOrderEnumIgnoreCase := make(map[string]ListServiceCatalogAssociationsSortOrderEnum)
	for k, v := range mappingListServiceCatalogAssociationsSortOrderEnum {
		mappingListServiceCatalogAssociationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListServiceCatalogAssociationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceCatalogAssociationsSortByEnum Enum with underlying type: string
type ListServiceCatalogAssociationsSortByEnum string

// Set of constants representing the allowable values for ListServiceCatalogAssociationsSortByEnum
const (
	ListServiceCatalogAssociationsSortByTimecreated ListServiceCatalogAssociationsSortByEnum = "TIMECREATED"
)

var mappingListServiceCatalogAssociationsSortByEnum = map[string]ListServiceCatalogAssociationsSortByEnum{
	"TIMECREATED": ListServiceCatalogAssociationsSortByTimecreated,
}

// GetListServiceCatalogAssociationsSortByEnumValues Enumerates the set of values for ListServiceCatalogAssociationsSortByEnum
func GetListServiceCatalogAssociationsSortByEnumValues() []ListServiceCatalogAssociationsSortByEnum {
	values := make([]ListServiceCatalogAssociationsSortByEnum, 0)
	for _, v := range mappingListServiceCatalogAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceCatalogAssociationsSortByEnumStringValues Enumerates the set of values in String for ListServiceCatalogAssociationsSortByEnum
func GetListServiceCatalogAssociationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
	}
}

// GetMappingListServiceCatalogAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceCatalogAssociationsSortByEnum(val string) (ListServiceCatalogAssociationsSortByEnum, bool) {
	mappingListServiceCatalogAssociationsSortByEnumIgnoreCase := make(map[string]ListServiceCatalogAssociationsSortByEnum)
	for k, v := range mappingListServiceCatalogAssociationsSortByEnum {
		mappingListServiceCatalogAssociationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListServiceCatalogAssociationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
