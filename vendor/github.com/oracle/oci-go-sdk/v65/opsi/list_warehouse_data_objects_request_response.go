// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWarehouseDataObjectsRequest wrapper for the ListWarehouseDataObjects operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListWarehouseDataObjects.go.html to see an example of how to use ListWarehouseDataObjectsRequest.
type ListWarehouseDataObjectsRequest struct {

	// Type of the Warehouse.
	WarehouseType ListWarehouseDataObjectsWarehouseTypeEnum `mandatory:"true" contributesTo:"path" name:"warehouseType"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Warehouse.
	WarehouseId *string `mandatory:"true" contributesTo:"path" name:"warehouseId"`

	// A filter to return only data objects that match the data object type. By default, no filtering will be applied on data object type.
	DataObjectType []DataObjectTypeEnum `contributesTo:"query" name:"dataObjectType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only data objects that match the entire data object name. By default, no filtering will be applied on data object name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only data objects that match the entire data object owner name.  By default, no filtering will be applied on data object owner name.
	Owner *string `mandatory:"false" contributesTo:"query" name:"owner"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListWarehouseDataObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort options for Warehouse data objects list.
	SortBy ListWarehouseDataObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies the optional fields to return in a WarehouseDataObjectSummary. Unless requested, these fields are not returned by default.
	SummaryField []ListWarehouseDataObjectsSummaryFieldEnum `contributesTo:"query" name:"summaryField" omitEmpty:"true" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWarehouseDataObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWarehouseDataObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWarehouseDataObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWarehouseDataObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWarehouseDataObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWarehouseDataObjectsWarehouseTypeEnum(string(request.WarehouseType)); !ok && request.WarehouseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WarehouseType: %s. Supported values are: %s.", request.WarehouseType, strings.Join(GetListWarehouseDataObjectsWarehouseTypeEnumStringValues(), ",")))
	}
	for _, val := range request.DataObjectType {
		if _, ok := GetMappingDataObjectTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataObjectType: %s. Supported values are: %s.", val, strings.Join(GetDataObjectTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListWarehouseDataObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWarehouseDataObjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWarehouseDataObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWarehouseDataObjectsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.SummaryField {
		if _, ok := GetMappingListWarehouseDataObjectsSummaryFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SummaryField: %s. Supported values are: %s.", val, strings.Join(GetListWarehouseDataObjectsSummaryFieldEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWarehouseDataObjectsResponse wrapper for the ListWarehouseDataObjects operation
type ListWarehouseDataObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WarehouseDataObjectCollection instances
	WarehouseDataObjectCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWarehouseDataObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWarehouseDataObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWarehouseDataObjectsWarehouseTypeEnum Enum with underlying type: string
type ListWarehouseDataObjectsWarehouseTypeEnum string

// Set of constants representing the allowable values for ListWarehouseDataObjectsWarehouseTypeEnum
const (
	ListWarehouseDataObjectsWarehouseTypeAwrhubs ListWarehouseDataObjectsWarehouseTypeEnum = "awrHubs"
)

var mappingListWarehouseDataObjectsWarehouseTypeEnum = map[string]ListWarehouseDataObjectsWarehouseTypeEnum{
	"awrHubs": ListWarehouseDataObjectsWarehouseTypeAwrhubs,
}

var mappingListWarehouseDataObjectsWarehouseTypeEnumLowerCase = map[string]ListWarehouseDataObjectsWarehouseTypeEnum{
	"awrhubs": ListWarehouseDataObjectsWarehouseTypeAwrhubs,
}

// GetListWarehouseDataObjectsWarehouseTypeEnumValues Enumerates the set of values for ListWarehouseDataObjectsWarehouseTypeEnum
func GetListWarehouseDataObjectsWarehouseTypeEnumValues() []ListWarehouseDataObjectsWarehouseTypeEnum {
	values := make([]ListWarehouseDataObjectsWarehouseTypeEnum, 0)
	for _, v := range mappingListWarehouseDataObjectsWarehouseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListWarehouseDataObjectsWarehouseTypeEnumStringValues Enumerates the set of values in String for ListWarehouseDataObjectsWarehouseTypeEnum
func GetListWarehouseDataObjectsWarehouseTypeEnumStringValues() []string {
	return []string{
		"awrHubs",
	}
}

// GetMappingListWarehouseDataObjectsWarehouseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWarehouseDataObjectsWarehouseTypeEnum(val string) (ListWarehouseDataObjectsWarehouseTypeEnum, bool) {
	enum, ok := mappingListWarehouseDataObjectsWarehouseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWarehouseDataObjectsSortOrderEnum Enum with underlying type: string
type ListWarehouseDataObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListWarehouseDataObjectsSortOrderEnum
const (
	ListWarehouseDataObjectsSortOrderAsc  ListWarehouseDataObjectsSortOrderEnum = "ASC"
	ListWarehouseDataObjectsSortOrderDesc ListWarehouseDataObjectsSortOrderEnum = "DESC"
)

var mappingListWarehouseDataObjectsSortOrderEnum = map[string]ListWarehouseDataObjectsSortOrderEnum{
	"ASC":  ListWarehouseDataObjectsSortOrderAsc,
	"DESC": ListWarehouseDataObjectsSortOrderDesc,
}

var mappingListWarehouseDataObjectsSortOrderEnumLowerCase = map[string]ListWarehouseDataObjectsSortOrderEnum{
	"asc":  ListWarehouseDataObjectsSortOrderAsc,
	"desc": ListWarehouseDataObjectsSortOrderDesc,
}

// GetListWarehouseDataObjectsSortOrderEnumValues Enumerates the set of values for ListWarehouseDataObjectsSortOrderEnum
func GetListWarehouseDataObjectsSortOrderEnumValues() []ListWarehouseDataObjectsSortOrderEnum {
	values := make([]ListWarehouseDataObjectsSortOrderEnum, 0)
	for _, v := range mappingListWarehouseDataObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWarehouseDataObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListWarehouseDataObjectsSortOrderEnum
func GetListWarehouseDataObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWarehouseDataObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWarehouseDataObjectsSortOrderEnum(val string) (ListWarehouseDataObjectsSortOrderEnum, bool) {
	enum, ok := mappingListWarehouseDataObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWarehouseDataObjectsSortByEnum Enum with underlying type: string
type ListWarehouseDataObjectsSortByEnum string

// Set of constants representing the allowable values for ListWarehouseDataObjectsSortByEnum
const (
	ListWarehouseDataObjectsSortByDataobjecttype ListWarehouseDataObjectsSortByEnum = "dataObjectType"
	ListWarehouseDataObjectsSortByName           ListWarehouseDataObjectsSortByEnum = "name"
	ListWarehouseDataObjectsSortByOwner          ListWarehouseDataObjectsSortByEnum = "owner"
)

var mappingListWarehouseDataObjectsSortByEnum = map[string]ListWarehouseDataObjectsSortByEnum{
	"dataObjectType": ListWarehouseDataObjectsSortByDataobjecttype,
	"name":           ListWarehouseDataObjectsSortByName,
	"owner":          ListWarehouseDataObjectsSortByOwner,
}

var mappingListWarehouseDataObjectsSortByEnumLowerCase = map[string]ListWarehouseDataObjectsSortByEnum{
	"dataobjecttype": ListWarehouseDataObjectsSortByDataobjecttype,
	"name":           ListWarehouseDataObjectsSortByName,
	"owner":          ListWarehouseDataObjectsSortByOwner,
}

// GetListWarehouseDataObjectsSortByEnumValues Enumerates the set of values for ListWarehouseDataObjectsSortByEnum
func GetListWarehouseDataObjectsSortByEnumValues() []ListWarehouseDataObjectsSortByEnum {
	values := make([]ListWarehouseDataObjectsSortByEnum, 0)
	for _, v := range mappingListWarehouseDataObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWarehouseDataObjectsSortByEnumStringValues Enumerates the set of values in String for ListWarehouseDataObjectsSortByEnum
func GetListWarehouseDataObjectsSortByEnumStringValues() []string {
	return []string{
		"dataObjectType",
		"name",
		"owner",
	}
}

// GetMappingListWarehouseDataObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWarehouseDataObjectsSortByEnum(val string) (ListWarehouseDataObjectsSortByEnum, bool) {
	enum, ok := mappingListWarehouseDataObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWarehouseDataObjectsSummaryFieldEnum Enum with underlying type: string
type ListWarehouseDataObjectsSummaryFieldEnum string

// Set of constants representing the allowable values for ListWarehouseDataObjectsSummaryFieldEnum
const (
	ListWarehouseDataObjectsSummaryFieldDetails ListWarehouseDataObjectsSummaryFieldEnum = "details"
)

var mappingListWarehouseDataObjectsSummaryFieldEnum = map[string]ListWarehouseDataObjectsSummaryFieldEnum{
	"details": ListWarehouseDataObjectsSummaryFieldDetails,
}

var mappingListWarehouseDataObjectsSummaryFieldEnumLowerCase = map[string]ListWarehouseDataObjectsSummaryFieldEnum{
	"details": ListWarehouseDataObjectsSummaryFieldDetails,
}

// GetListWarehouseDataObjectsSummaryFieldEnumValues Enumerates the set of values for ListWarehouseDataObjectsSummaryFieldEnum
func GetListWarehouseDataObjectsSummaryFieldEnumValues() []ListWarehouseDataObjectsSummaryFieldEnum {
	values := make([]ListWarehouseDataObjectsSummaryFieldEnum, 0)
	for _, v := range mappingListWarehouseDataObjectsSummaryFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetListWarehouseDataObjectsSummaryFieldEnumStringValues Enumerates the set of values in String for ListWarehouseDataObjectsSummaryFieldEnum
func GetListWarehouseDataObjectsSummaryFieldEnumStringValues() []string {
	return []string{
		"details",
	}
}

// GetMappingListWarehouseDataObjectsSummaryFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWarehouseDataObjectsSummaryFieldEnum(val string) (ListWarehouseDataObjectsSummaryFieldEnum, bool) {
	enum, ok := mappingListWarehouseDataObjectsSummaryFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
