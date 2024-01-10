// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSensitiveObjectsRequest wrapper for the ListSensitiveObjects operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveObjects.go.html to see an example of how to use ListSensitiveObjectsRequest.
type ListSensitiveObjectsRequest struct {

	// The OCID of the sensitive data model.
	SensitiveDataModelId *string `mandatory:"true" contributesTo:"path" name:"sensitiveDataModelId"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object type.
	ObjectType []ListSensitiveObjectsObjectTypeEnum `contributesTo:"query" name:"objectType" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSensitiveObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder).
	// The default order is ascending.
	SortBy ListSensitiveObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ObjectType {
		if _, ok := GetMappingListSensitiveObjectsObjectTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", val, strings.Join(GetListSensitiveObjectsObjectTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListSensitiveObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSensitiveObjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSensitiveObjectsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveObjectsResponse wrapper for the ListSensitiveObjects operation
type ListSensitiveObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveObjectCollection instances
	SensitiveObjectCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveObjectsObjectTypeEnum Enum with underlying type: string
type ListSensitiveObjectsObjectTypeEnum string

// Set of constants representing the allowable values for ListSensitiveObjectsObjectTypeEnum
const (
	ListSensitiveObjectsObjectTypeAll            ListSensitiveObjectsObjectTypeEnum = "ALL"
	ListSensitiveObjectsObjectTypeTable          ListSensitiveObjectsObjectTypeEnum = "TABLE"
	ListSensitiveObjectsObjectTypeEditioningView ListSensitiveObjectsObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingListSensitiveObjectsObjectTypeEnum = map[string]ListSensitiveObjectsObjectTypeEnum{
	"ALL":             ListSensitiveObjectsObjectTypeAll,
	"TABLE":           ListSensitiveObjectsObjectTypeTable,
	"EDITIONING_VIEW": ListSensitiveObjectsObjectTypeEditioningView,
}

var mappingListSensitiveObjectsObjectTypeEnumLowerCase = map[string]ListSensitiveObjectsObjectTypeEnum{
	"all":             ListSensitiveObjectsObjectTypeAll,
	"table":           ListSensitiveObjectsObjectTypeTable,
	"editioning_view": ListSensitiveObjectsObjectTypeEditioningView,
}

// GetListSensitiveObjectsObjectTypeEnumValues Enumerates the set of values for ListSensitiveObjectsObjectTypeEnum
func GetListSensitiveObjectsObjectTypeEnumValues() []ListSensitiveObjectsObjectTypeEnum {
	values := make([]ListSensitiveObjectsObjectTypeEnum, 0)
	for _, v := range mappingListSensitiveObjectsObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveObjectsObjectTypeEnumStringValues Enumerates the set of values in String for ListSensitiveObjectsObjectTypeEnum
func GetListSensitiveObjectsObjectTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingListSensitiveObjectsObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveObjectsObjectTypeEnum(val string) (ListSensitiveObjectsObjectTypeEnum, bool) {
	enum, ok := mappingListSensitiveObjectsObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveObjectsSortOrderEnum Enum with underlying type: string
type ListSensitiveObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListSensitiveObjectsSortOrderEnum
const (
	ListSensitiveObjectsSortOrderAsc  ListSensitiveObjectsSortOrderEnum = "ASC"
	ListSensitiveObjectsSortOrderDesc ListSensitiveObjectsSortOrderEnum = "DESC"
)

var mappingListSensitiveObjectsSortOrderEnum = map[string]ListSensitiveObjectsSortOrderEnum{
	"ASC":  ListSensitiveObjectsSortOrderAsc,
	"DESC": ListSensitiveObjectsSortOrderDesc,
}

var mappingListSensitiveObjectsSortOrderEnumLowerCase = map[string]ListSensitiveObjectsSortOrderEnum{
	"asc":  ListSensitiveObjectsSortOrderAsc,
	"desc": ListSensitiveObjectsSortOrderDesc,
}

// GetListSensitiveObjectsSortOrderEnumValues Enumerates the set of values for ListSensitiveObjectsSortOrderEnum
func GetListSensitiveObjectsSortOrderEnumValues() []ListSensitiveObjectsSortOrderEnum {
	values := make([]ListSensitiveObjectsSortOrderEnum, 0)
	for _, v := range mappingListSensitiveObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListSensitiveObjectsSortOrderEnum
func GetListSensitiveObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSensitiveObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveObjectsSortOrderEnum(val string) (ListSensitiveObjectsSortOrderEnum, bool) {
	enum, ok := mappingListSensitiveObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveObjectsSortByEnum Enum with underlying type: string
type ListSensitiveObjectsSortByEnum string

// Set of constants representing the allowable values for ListSensitiveObjectsSortByEnum
const (
	ListSensitiveObjectsSortBySchemaname ListSensitiveObjectsSortByEnum = "schemaName"
	ListSensitiveObjectsSortByObjectname ListSensitiveObjectsSortByEnum = "objectName"
	ListSensitiveObjectsSortByObjecttype ListSensitiveObjectsSortByEnum = "objectType"
)

var mappingListSensitiveObjectsSortByEnum = map[string]ListSensitiveObjectsSortByEnum{
	"schemaName": ListSensitiveObjectsSortBySchemaname,
	"objectName": ListSensitiveObjectsSortByObjectname,
	"objectType": ListSensitiveObjectsSortByObjecttype,
}

var mappingListSensitiveObjectsSortByEnumLowerCase = map[string]ListSensitiveObjectsSortByEnum{
	"schemaname": ListSensitiveObjectsSortBySchemaname,
	"objectname": ListSensitiveObjectsSortByObjectname,
	"objecttype": ListSensitiveObjectsSortByObjecttype,
}

// GetListSensitiveObjectsSortByEnumValues Enumerates the set of values for ListSensitiveObjectsSortByEnum
func GetListSensitiveObjectsSortByEnumValues() []ListSensitiveObjectsSortByEnum {
	values := make([]ListSensitiveObjectsSortByEnum, 0)
	for _, v := range mappingListSensitiveObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveObjectsSortByEnumStringValues Enumerates the set of values in String for ListSensitiveObjectsSortByEnum
func GetListSensitiveObjectsSortByEnumStringValues() []string {
	return []string{
		"schemaName",
		"objectName",
		"objectType",
	}
}

// GetMappingListSensitiveObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveObjectsSortByEnum(val string) (ListSensitiveObjectsSortByEnum, bool) {
	enum, ok := mappingListSensitiveObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
