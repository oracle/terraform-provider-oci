// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaskingObjectsRequest wrapper for the ListMaskingObjects operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingObjects.go.html to see an example of how to use ListMaskingObjectsRequest.
type ListMaskingObjectsRequest struct {

	// The OCID of the masking policy.
	MaskingPolicyId *string `mandatory:"true" contributesTo:"path" name:"maskingPolicyId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder).
	// The default order is ascending.
	SortBy ListMaskingObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object type.
	ObjectType []ListMaskingObjectsObjectTypeEnum `contributesTo:"query" name:"objectType" omitEmpty:"true" collectionFormat:"multi"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingObjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingObjectsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.ObjectType {
		if _, ok := GetMappingListMaskingObjectsObjectTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", val, strings.Join(GetListMaskingObjectsObjectTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingObjectsResponse wrapper for the ListMaskingObjects operation
type ListMaskingObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingObjectCollection instances
	MaskingObjectCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingObjectsSortOrderEnum Enum with underlying type: string
type ListMaskingObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingObjectsSortOrderEnum
const (
	ListMaskingObjectsSortOrderAsc  ListMaskingObjectsSortOrderEnum = "ASC"
	ListMaskingObjectsSortOrderDesc ListMaskingObjectsSortOrderEnum = "DESC"
)

var mappingListMaskingObjectsSortOrderEnum = map[string]ListMaskingObjectsSortOrderEnum{
	"ASC":  ListMaskingObjectsSortOrderAsc,
	"DESC": ListMaskingObjectsSortOrderDesc,
}

var mappingListMaskingObjectsSortOrderEnumLowerCase = map[string]ListMaskingObjectsSortOrderEnum{
	"asc":  ListMaskingObjectsSortOrderAsc,
	"desc": ListMaskingObjectsSortOrderDesc,
}

// GetListMaskingObjectsSortOrderEnumValues Enumerates the set of values for ListMaskingObjectsSortOrderEnum
func GetListMaskingObjectsSortOrderEnumValues() []ListMaskingObjectsSortOrderEnum {
	values := make([]ListMaskingObjectsSortOrderEnum, 0)
	for _, v := range mappingListMaskingObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingObjectsSortOrderEnum
func GetListMaskingObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingObjectsSortOrderEnum(val string) (ListMaskingObjectsSortOrderEnum, bool) {
	enum, ok := mappingListMaskingObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingObjectsSortByEnum Enum with underlying type: string
type ListMaskingObjectsSortByEnum string

// Set of constants representing the allowable values for ListMaskingObjectsSortByEnum
const (
	ListMaskingObjectsSortBySchemaname ListMaskingObjectsSortByEnum = "schemaName"
	ListMaskingObjectsSortByObjectname ListMaskingObjectsSortByEnum = "objectName"
	ListMaskingObjectsSortByObjecttype ListMaskingObjectsSortByEnum = "objectType"
)

var mappingListMaskingObjectsSortByEnum = map[string]ListMaskingObjectsSortByEnum{
	"schemaName": ListMaskingObjectsSortBySchemaname,
	"objectName": ListMaskingObjectsSortByObjectname,
	"objectType": ListMaskingObjectsSortByObjecttype,
}

var mappingListMaskingObjectsSortByEnumLowerCase = map[string]ListMaskingObjectsSortByEnum{
	"schemaname": ListMaskingObjectsSortBySchemaname,
	"objectname": ListMaskingObjectsSortByObjectname,
	"objecttype": ListMaskingObjectsSortByObjecttype,
}

// GetListMaskingObjectsSortByEnumValues Enumerates the set of values for ListMaskingObjectsSortByEnum
func GetListMaskingObjectsSortByEnumValues() []ListMaskingObjectsSortByEnum {
	values := make([]ListMaskingObjectsSortByEnum, 0)
	for _, v := range mappingListMaskingObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingObjectsSortByEnumStringValues Enumerates the set of values in String for ListMaskingObjectsSortByEnum
func GetListMaskingObjectsSortByEnumStringValues() []string {
	return []string{
		"schemaName",
		"objectName",
		"objectType",
	}
}

// GetMappingListMaskingObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingObjectsSortByEnum(val string) (ListMaskingObjectsSortByEnum, bool) {
	enum, ok := mappingListMaskingObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingObjectsObjectTypeEnum Enum with underlying type: string
type ListMaskingObjectsObjectTypeEnum string

// Set of constants representing the allowable values for ListMaskingObjectsObjectTypeEnum
const (
	ListMaskingObjectsObjectTypeAll            ListMaskingObjectsObjectTypeEnum = "ALL"
	ListMaskingObjectsObjectTypeTable          ListMaskingObjectsObjectTypeEnum = "TABLE"
	ListMaskingObjectsObjectTypeEditioningView ListMaskingObjectsObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingListMaskingObjectsObjectTypeEnum = map[string]ListMaskingObjectsObjectTypeEnum{
	"ALL":             ListMaskingObjectsObjectTypeAll,
	"TABLE":           ListMaskingObjectsObjectTypeTable,
	"EDITIONING_VIEW": ListMaskingObjectsObjectTypeEditioningView,
}

var mappingListMaskingObjectsObjectTypeEnumLowerCase = map[string]ListMaskingObjectsObjectTypeEnum{
	"all":             ListMaskingObjectsObjectTypeAll,
	"table":           ListMaskingObjectsObjectTypeTable,
	"editioning_view": ListMaskingObjectsObjectTypeEditioningView,
}

// GetListMaskingObjectsObjectTypeEnumValues Enumerates the set of values for ListMaskingObjectsObjectTypeEnum
func GetListMaskingObjectsObjectTypeEnumValues() []ListMaskingObjectsObjectTypeEnum {
	values := make([]ListMaskingObjectsObjectTypeEnum, 0)
	for _, v := range mappingListMaskingObjectsObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingObjectsObjectTypeEnumStringValues Enumerates the set of values in String for ListMaskingObjectsObjectTypeEnum
func GetListMaskingObjectsObjectTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingListMaskingObjectsObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingObjectsObjectTypeEnum(val string) (ListMaskingObjectsObjectTypeEnum, bool) {
	enum, ok := mappingListMaskingObjectsObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
