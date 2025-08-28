// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListReferentialRelationsRequest wrapper for the ListReferentialRelations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListReferentialRelations.go.html to see an example of how to use ListReferentialRelationsRequest.
type ListReferentialRelationsRequest struct {

	// The OCID of the sensitive data model.
	SensitiveDataModelId *string `mandatory:"true" contributesTo:"path" name:"sensitiveDataModelId"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// Returns referential relations containing sensitive columns when true.
	// Returns referential relations containing non sensitive columns when false.
	IsSensitive *bool `mandatory:"false" contributesTo:"query" name:"isSensitive"`

	// A filter to return sensitive columns based on their relationship with their parent columns. If set to NONE,
	// it returns the sensitive columns that do not have any parent. The response includes the parent columns as
	// well as the independent columns that are not in any relationship. If set to APP_DEFINED, it returns all the
	// child columns that have application-level (non-dictionary) relationship with their parents. If set to DB_DEFINED,
	// it returns all the child columns that have database-level (dictionary-defined) relationship with their parents.
	RelationType []ListReferentialRelationsRelationTypeEnum `contributesTo:"query" name:"relationType" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListReferentialRelationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for key is descending.
	SortBy ListReferentialRelationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReferentialRelationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReferentialRelationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReferentialRelationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReferentialRelationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReferentialRelationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.RelationType {
		if _, ok := GetMappingListReferentialRelationsRelationTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", val, strings.Join(GetListReferentialRelationsRelationTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListReferentialRelationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReferentialRelationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReferentialRelationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReferentialRelationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReferentialRelationsResponse wrapper for the ListReferentialRelations operation
type ListReferentialRelationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ReferentialRelationCollection instances
	ReferentialRelationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListReferentialRelationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReferentialRelationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReferentialRelationsRelationTypeEnum Enum with underlying type: string
type ListReferentialRelationsRelationTypeEnum string

// Set of constants representing the allowable values for ListReferentialRelationsRelationTypeEnum
const (
	ListReferentialRelationsRelationTypeNone       ListReferentialRelationsRelationTypeEnum = "NONE"
	ListReferentialRelationsRelationTypeAppDefined ListReferentialRelationsRelationTypeEnum = "APP_DEFINED"
	ListReferentialRelationsRelationTypeDbDefined  ListReferentialRelationsRelationTypeEnum = "DB_DEFINED"
)

var mappingListReferentialRelationsRelationTypeEnum = map[string]ListReferentialRelationsRelationTypeEnum{
	"NONE":        ListReferentialRelationsRelationTypeNone,
	"APP_DEFINED": ListReferentialRelationsRelationTypeAppDefined,
	"DB_DEFINED":  ListReferentialRelationsRelationTypeDbDefined,
}

var mappingListReferentialRelationsRelationTypeEnumLowerCase = map[string]ListReferentialRelationsRelationTypeEnum{
	"none":        ListReferentialRelationsRelationTypeNone,
	"app_defined": ListReferentialRelationsRelationTypeAppDefined,
	"db_defined":  ListReferentialRelationsRelationTypeDbDefined,
}

// GetListReferentialRelationsRelationTypeEnumValues Enumerates the set of values for ListReferentialRelationsRelationTypeEnum
func GetListReferentialRelationsRelationTypeEnumValues() []ListReferentialRelationsRelationTypeEnum {
	values := make([]ListReferentialRelationsRelationTypeEnum, 0)
	for _, v := range mappingListReferentialRelationsRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListReferentialRelationsRelationTypeEnumStringValues Enumerates the set of values in String for ListReferentialRelationsRelationTypeEnum
func GetListReferentialRelationsRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingListReferentialRelationsRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReferentialRelationsRelationTypeEnum(val string) (ListReferentialRelationsRelationTypeEnum, bool) {
	enum, ok := mappingListReferentialRelationsRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReferentialRelationsSortOrderEnum Enum with underlying type: string
type ListReferentialRelationsSortOrderEnum string

// Set of constants representing the allowable values for ListReferentialRelationsSortOrderEnum
const (
	ListReferentialRelationsSortOrderAsc  ListReferentialRelationsSortOrderEnum = "ASC"
	ListReferentialRelationsSortOrderDesc ListReferentialRelationsSortOrderEnum = "DESC"
)

var mappingListReferentialRelationsSortOrderEnum = map[string]ListReferentialRelationsSortOrderEnum{
	"ASC":  ListReferentialRelationsSortOrderAsc,
	"DESC": ListReferentialRelationsSortOrderDesc,
}

var mappingListReferentialRelationsSortOrderEnumLowerCase = map[string]ListReferentialRelationsSortOrderEnum{
	"asc":  ListReferentialRelationsSortOrderAsc,
	"desc": ListReferentialRelationsSortOrderDesc,
}

// GetListReferentialRelationsSortOrderEnumValues Enumerates the set of values for ListReferentialRelationsSortOrderEnum
func GetListReferentialRelationsSortOrderEnumValues() []ListReferentialRelationsSortOrderEnum {
	values := make([]ListReferentialRelationsSortOrderEnum, 0)
	for _, v := range mappingListReferentialRelationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReferentialRelationsSortOrderEnumStringValues Enumerates the set of values in String for ListReferentialRelationsSortOrderEnum
func GetListReferentialRelationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReferentialRelationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReferentialRelationsSortOrderEnum(val string) (ListReferentialRelationsSortOrderEnum, bool) {
	enum, ok := mappingListReferentialRelationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReferentialRelationsSortByEnum Enum with underlying type: string
type ListReferentialRelationsSortByEnum string

// Set of constants representing the allowable values for ListReferentialRelationsSortByEnum
const (
	ListReferentialRelationsSortByKey          ListReferentialRelationsSortByEnum = "key"
	ListReferentialRelationsSortByRelationtype ListReferentialRelationsSortByEnum = "relationType"
	ListReferentialRelationsSortBySchemaname   ListReferentialRelationsSortByEnum = "schemaName"
	ListReferentialRelationsSortByTablename    ListReferentialRelationsSortByEnum = "tableName"
)

var mappingListReferentialRelationsSortByEnum = map[string]ListReferentialRelationsSortByEnum{
	"key":          ListReferentialRelationsSortByKey,
	"relationType": ListReferentialRelationsSortByRelationtype,
	"schemaName":   ListReferentialRelationsSortBySchemaname,
	"tableName":    ListReferentialRelationsSortByTablename,
}

var mappingListReferentialRelationsSortByEnumLowerCase = map[string]ListReferentialRelationsSortByEnum{
	"key":          ListReferentialRelationsSortByKey,
	"relationtype": ListReferentialRelationsSortByRelationtype,
	"schemaname":   ListReferentialRelationsSortBySchemaname,
	"tablename":    ListReferentialRelationsSortByTablename,
}

// GetListReferentialRelationsSortByEnumValues Enumerates the set of values for ListReferentialRelationsSortByEnum
func GetListReferentialRelationsSortByEnumValues() []ListReferentialRelationsSortByEnum {
	values := make([]ListReferentialRelationsSortByEnum, 0)
	for _, v := range mappingListReferentialRelationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReferentialRelationsSortByEnumStringValues Enumerates the set of values in String for ListReferentialRelationsSortByEnum
func GetListReferentialRelationsSortByEnumStringValues() []string {
	return []string{
		"key",
		"relationType",
		"schemaName",
		"tableName",
	}
}

// GetMappingListReferentialRelationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReferentialRelationsSortByEnum(val string) (ListReferentialRelationsSortByEnum, bool) {
	enum, ok := mappingListReferentialRelationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
