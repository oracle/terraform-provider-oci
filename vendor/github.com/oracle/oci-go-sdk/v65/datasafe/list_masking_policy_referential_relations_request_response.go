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

// ListMaskingPolicyReferentialRelationsRequest wrapper for the ListMaskingPolicyReferentialRelations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingPolicyReferentialRelations.go.html to see an example of how to use ListMaskingPolicyReferentialRelationsRequest.
type ListMaskingPolicyReferentialRelationsRequest struct {

	// The OCID of the masking policy.
	MaskingPolicyId *string `mandatory:"true" contributesTo:"path" name:"maskingPolicyId"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// A filter to return columns based on their relationship with their parent columns. If set to NONE,
	// it returns the columns that do not have any parent. The response includes the parent columns as
	// well as the independent columns that are not in any relationship. If set to APP_DEFINED, it returns all the
	// child columns that have application-level (non-dictionary) relationship with their parents. If set to DB_DEFINED,
	// it returns all the child columns that have database-level (dictionary-defined) relationship with their parents.
	RelationType []ListMaskingPolicyReferentialRelationsRelationTypeEnum `contributesTo:"query" name:"relationType" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingPolicyReferentialRelationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder).
	SortBy ListMaskingPolicyReferentialRelationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingPolicyReferentialRelationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingPolicyReferentialRelationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingPolicyReferentialRelationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingPolicyReferentialRelationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingPolicyReferentialRelationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.RelationType {
		if _, ok := GetMappingListMaskingPolicyReferentialRelationsRelationTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", val, strings.Join(GetListMaskingPolicyReferentialRelationsRelationTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListMaskingPolicyReferentialRelationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingPolicyReferentialRelationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPolicyReferentialRelationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingPolicyReferentialRelationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingPolicyReferentialRelationsResponse wrapper for the ListMaskingPolicyReferentialRelations operation
type ListMaskingPolicyReferentialRelationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingPolicyReferentialRelationCollection instances
	MaskingPolicyReferentialRelationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingPolicyReferentialRelationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingPolicyReferentialRelationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingPolicyReferentialRelationsRelationTypeEnum Enum with underlying type: string
type ListMaskingPolicyReferentialRelationsRelationTypeEnum string

// Set of constants representing the allowable values for ListMaskingPolicyReferentialRelationsRelationTypeEnum
const (
	ListMaskingPolicyReferentialRelationsRelationTypeNone       ListMaskingPolicyReferentialRelationsRelationTypeEnum = "NONE"
	ListMaskingPolicyReferentialRelationsRelationTypeAppDefined ListMaskingPolicyReferentialRelationsRelationTypeEnum = "APP_DEFINED"
	ListMaskingPolicyReferentialRelationsRelationTypeDbDefined  ListMaskingPolicyReferentialRelationsRelationTypeEnum = "DB_DEFINED"
)

var mappingListMaskingPolicyReferentialRelationsRelationTypeEnum = map[string]ListMaskingPolicyReferentialRelationsRelationTypeEnum{
	"NONE":        ListMaskingPolicyReferentialRelationsRelationTypeNone,
	"APP_DEFINED": ListMaskingPolicyReferentialRelationsRelationTypeAppDefined,
	"DB_DEFINED":  ListMaskingPolicyReferentialRelationsRelationTypeDbDefined,
}

var mappingListMaskingPolicyReferentialRelationsRelationTypeEnumLowerCase = map[string]ListMaskingPolicyReferentialRelationsRelationTypeEnum{
	"none":        ListMaskingPolicyReferentialRelationsRelationTypeNone,
	"app_defined": ListMaskingPolicyReferentialRelationsRelationTypeAppDefined,
	"db_defined":  ListMaskingPolicyReferentialRelationsRelationTypeDbDefined,
}

// GetListMaskingPolicyReferentialRelationsRelationTypeEnumValues Enumerates the set of values for ListMaskingPolicyReferentialRelationsRelationTypeEnum
func GetListMaskingPolicyReferentialRelationsRelationTypeEnumValues() []ListMaskingPolicyReferentialRelationsRelationTypeEnum {
	values := make([]ListMaskingPolicyReferentialRelationsRelationTypeEnum, 0)
	for _, v := range mappingListMaskingPolicyReferentialRelationsRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyReferentialRelationsRelationTypeEnumStringValues Enumerates the set of values in String for ListMaskingPolicyReferentialRelationsRelationTypeEnum
func GetListMaskingPolicyReferentialRelationsRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingListMaskingPolicyReferentialRelationsRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyReferentialRelationsRelationTypeEnum(val string) (ListMaskingPolicyReferentialRelationsRelationTypeEnum, bool) {
	enum, ok := mappingListMaskingPolicyReferentialRelationsRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPolicyReferentialRelationsSortOrderEnum Enum with underlying type: string
type ListMaskingPolicyReferentialRelationsSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingPolicyReferentialRelationsSortOrderEnum
const (
	ListMaskingPolicyReferentialRelationsSortOrderAsc  ListMaskingPolicyReferentialRelationsSortOrderEnum = "ASC"
	ListMaskingPolicyReferentialRelationsSortOrderDesc ListMaskingPolicyReferentialRelationsSortOrderEnum = "DESC"
)

var mappingListMaskingPolicyReferentialRelationsSortOrderEnum = map[string]ListMaskingPolicyReferentialRelationsSortOrderEnum{
	"ASC":  ListMaskingPolicyReferentialRelationsSortOrderAsc,
	"DESC": ListMaskingPolicyReferentialRelationsSortOrderDesc,
}

var mappingListMaskingPolicyReferentialRelationsSortOrderEnumLowerCase = map[string]ListMaskingPolicyReferentialRelationsSortOrderEnum{
	"asc":  ListMaskingPolicyReferentialRelationsSortOrderAsc,
	"desc": ListMaskingPolicyReferentialRelationsSortOrderDesc,
}

// GetListMaskingPolicyReferentialRelationsSortOrderEnumValues Enumerates the set of values for ListMaskingPolicyReferentialRelationsSortOrderEnum
func GetListMaskingPolicyReferentialRelationsSortOrderEnumValues() []ListMaskingPolicyReferentialRelationsSortOrderEnum {
	values := make([]ListMaskingPolicyReferentialRelationsSortOrderEnum, 0)
	for _, v := range mappingListMaskingPolicyReferentialRelationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyReferentialRelationsSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingPolicyReferentialRelationsSortOrderEnum
func GetListMaskingPolicyReferentialRelationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingPolicyReferentialRelationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyReferentialRelationsSortOrderEnum(val string) (ListMaskingPolicyReferentialRelationsSortOrderEnum, bool) {
	enum, ok := mappingListMaskingPolicyReferentialRelationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPolicyReferentialRelationsSortByEnum Enum with underlying type: string
type ListMaskingPolicyReferentialRelationsSortByEnum string

// Set of constants representing the allowable values for ListMaskingPolicyReferentialRelationsSortByEnum
const (
	ListMaskingPolicyReferentialRelationsSortByRelationtype ListMaskingPolicyReferentialRelationsSortByEnum = "relationType"
	ListMaskingPolicyReferentialRelationsSortBySchemaname   ListMaskingPolicyReferentialRelationsSortByEnum = "schemaName"
	ListMaskingPolicyReferentialRelationsSortByObjectname   ListMaskingPolicyReferentialRelationsSortByEnum = "objectName"
)

var mappingListMaskingPolicyReferentialRelationsSortByEnum = map[string]ListMaskingPolicyReferentialRelationsSortByEnum{
	"relationType": ListMaskingPolicyReferentialRelationsSortByRelationtype,
	"schemaName":   ListMaskingPolicyReferentialRelationsSortBySchemaname,
	"objectName":   ListMaskingPolicyReferentialRelationsSortByObjectname,
}

var mappingListMaskingPolicyReferentialRelationsSortByEnumLowerCase = map[string]ListMaskingPolicyReferentialRelationsSortByEnum{
	"relationtype": ListMaskingPolicyReferentialRelationsSortByRelationtype,
	"schemaname":   ListMaskingPolicyReferentialRelationsSortBySchemaname,
	"objectname":   ListMaskingPolicyReferentialRelationsSortByObjectname,
}

// GetListMaskingPolicyReferentialRelationsSortByEnumValues Enumerates the set of values for ListMaskingPolicyReferentialRelationsSortByEnum
func GetListMaskingPolicyReferentialRelationsSortByEnumValues() []ListMaskingPolicyReferentialRelationsSortByEnum {
	values := make([]ListMaskingPolicyReferentialRelationsSortByEnum, 0)
	for _, v := range mappingListMaskingPolicyReferentialRelationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyReferentialRelationsSortByEnumStringValues Enumerates the set of values in String for ListMaskingPolicyReferentialRelationsSortByEnum
func GetListMaskingPolicyReferentialRelationsSortByEnumStringValues() []string {
	return []string{
		"relationType",
		"schemaName",
		"objectName",
	}
}

// GetMappingListMaskingPolicyReferentialRelationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyReferentialRelationsSortByEnum(val string) (ListMaskingPolicyReferentialRelationsSortByEnum, bool) {
	enum, ok := mappingListMaskingPolicyReferentialRelationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
