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

// ListMaskedColumnsRequest wrapper for the ListMaskedColumns operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskedColumns.go.html to see an example of how to use ListMaskedColumnsRequest.
type ListMaskedColumnsRequest struct {

	// The OCID of the masking report.
	MaskingReportId *string `mandatory:"true" contributesTo:"path" name:"maskingReportId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskedColumnsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for all the fields is ascending.
	SortBy ListMaskedColumnsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object type.
	ObjectType []ListMaskedColumnsObjectTypeEnum `contributesTo:"query" name:"objectType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only the resources that match the specified masking column group.
	MaskingColumnGroup []string `contributesTo:"query" name:"maskingColumnGroup" collectionFormat:"multi"`

	// A filter to return only items related to a specific sensitive type OCID.
	SensitiveTypeId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypeId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskedColumnsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskedColumnsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskedColumnsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskedColumnsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskedColumnsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskedColumnsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskedColumnsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskedColumnsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskedColumnsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.ObjectType {
		if _, ok := GetMappingListMaskedColumnsObjectTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", val, strings.Join(GetListMaskedColumnsObjectTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskedColumnsResponse wrapper for the ListMaskedColumns operation
type ListMaskedColumnsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskedColumnCollection instances
	MaskedColumnCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskedColumnsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskedColumnsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskedColumnsSortOrderEnum Enum with underlying type: string
type ListMaskedColumnsSortOrderEnum string

// Set of constants representing the allowable values for ListMaskedColumnsSortOrderEnum
const (
	ListMaskedColumnsSortOrderAsc  ListMaskedColumnsSortOrderEnum = "ASC"
	ListMaskedColumnsSortOrderDesc ListMaskedColumnsSortOrderEnum = "DESC"
)

var mappingListMaskedColumnsSortOrderEnum = map[string]ListMaskedColumnsSortOrderEnum{
	"ASC":  ListMaskedColumnsSortOrderAsc,
	"DESC": ListMaskedColumnsSortOrderDesc,
}

var mappingListMaskedColumnsSortOrderEnumLowerCase = map[string]ListMaskedColumnsSortOrderEnum{
	"asc":  ListMaskedColumnsSortOrderAsc,
	"desc": ListMaskedColumnsSortOrderDesc,
}

// GetListMaskedColumnsSortOrderEnumValues Enumerates the set of values for ListMaskedColumnsSortOrderEnum
func GetListMaskedColumnsSortOrderEnumValues() []ListMaskedColumnsSortOrderEnum {
	values := make([]ListMaskedColumnsSortOrderEnum, 0)
	for _, v := range mappingListMaskedColumnsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskedColumnsSortOrderEnumStringValues Enumerates the set of values in String for ListMaskedColumnsSortOrderEnum
func GetListMaskedColumnsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskedColumnsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskedColumnsSortOrderEnum(val string) (ListMaskedColumnsSortOrderEnum, bool) {
	enum, ok := mappingListMaskedColumnsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskedColumnsSortByEnum Enum with underlying type: string
type ListMaskedColumnsSortByEnum string

// Set of constants representing the allowable values for ListMaskedColumnsSortByEnum
const (
	ListMaskedColumnsSortBySchemaname ListMaskedColumnsSortByEnum = "schemaName"
	ListMaskedColumnsSortByObjectname ListMaskedColumnsSortByEnum = "objectName"
)

var mappingListMaskedColumnsSortByEnum = map[string]ListMaskedColumnsSortByEnum{
	"schemaName": ListMaskedColumnsSortBySchemaname,
	"objectName": ListMaskedColumnsSortByObjectname,
}

var mappingListMaskedColumnsSortByEnumLowerCase = map[string]ListMaskedColumnsSortByEnum{
	"schemaname": ListMaskedColumnsSortBySchemaname,
	"objectname": ListMaskedColumnsSortByObjectname,
}

// GetListMaskedColumnsSortByEnumValues Enumerates the set of values for ListMaskedColumnsSortByEnum
func GetListMaskedColumnsSortByEnumValues() []ListMaskedColumnsSortByEnum {
	values := make([]ListMaskedColumnsSortByEnum, 0)
	for _, v := range mappingListMaskedColumnsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskedColumnsSortByEnumStringValues Enumerates the set of values in String for ListMaskedColumnsSortByEnum
func GetListMaskedColumnsSortByEnumStringValues() []string {
	return []string{
		"schemaName",
		"objectName",
	}
}

// GetMappingListMaskedColumnsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskedColumnsSortByEnum(val string) (ListMaskedColumnsSortByEnum, bool) {
	enum, ok := mappingListMaskedColumnsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskedColumnsObjectTypeEnum Enum with underlying type: string
type ListMaskedColumnsObjectTypeEnum string

// Set of constants representing the allowable values for ListMaskedColumnsObjectTypeEnum
const (
	ListMaskedColumnsObjectTypeAll            ListMaskedColumnsObjectTypeEnum = "ALL"
	ListMaskedColumnsObjectTypeTable          ListMaskedColumnsObjectTypeEnum = "TABLE"
	ListMaskedColumnsObjectTypeEditioningView ListMaskedColumnsObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingListMaskedColumnsObjectTypeEnum = map[string]ListMaskedColumnsObjectTypeEnum{
	"ALL":             ListMaskedColumnsObjectTypeAll,
	"TABLE":           ListMaskedColumnsObjectTypeTable,
	"EDITIONING_VIEW": ListMaskedColumnsObjectTypeEditioningView,
}

var mappingListMaskedColumnsObjectTypeEnumLowerCase = map[string]ListMaskedColumnsObjectTypeEnum{
	"all":             ListMaskedColumnsObjectTypeAll,
	"table":           ListMaskedColumnsObjectTypeTable,
	"editioning_view": ListMaskedColumnsObjectTypeEditioningView,
}

// GetListMaskedColumnsObjectTypeEnumValues Enumerates the set of values for ListMaskedColumnsObjectTypeEnum
func GetListMaskedColumnsObjectTypeEnumValues() []ListMaskedColumnsObjectTypeEnum {
	values := make([]ListMaskedColumnsObjectTypeEnum, 0)
	for _, v := range mappingListMaskedColumnsObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskedColumnsObjectTypeEnumStringValues Enumerates the set of values in String for ListMaskedColumnsObjectTypeEnum
func GetListMaskedColumnsObjectTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingListMaskedColumnsObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskedColumnsObjectTypeEnum(val string) (ListMaskedColumnsObjectTypeEnum, bool) {
	enum, ok := mappingListMaskedColumnsObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
