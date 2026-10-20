// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubsettingSchemaRelationsRequest wrapper for the ListSubsettingSchemaRelations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingSchemaRelations.go.html to see an example of how to use ListSubsettingSchemaRelationsRequest.
type ListSubsettingSchemaRelationsRequest struct {

	// The OCID of the subsetting policy.
	SubsettingPolicyId *string `mandatory:"true" contributesTo:"path" name:"subsettingPolicyId"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return columns based on their relationship with their parent columns. If set to APP_DEFINED, it returns all the
	// child columns that have application-level (non-dictionary) relationship with their parents. If set to DB_DEFINED,
	// it returns all the child columns that have database-level (dictionary-defined) relationship with their parents.
	RelationType ListSubsettingSchemaRelationsRelationTypeEnum `mandatory:"false" contributesTo:"query" name:"relationType" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingSchemaRelationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder).
	SortBy ListSubsettingSchemaRelationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingSchemaRelationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingSchemaRelationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingSchemaRelationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingSchemaRelationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingSchemaRelationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingSchemaRelationsRelationTypeEnum(string(request.RelationType)); !ok && request.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", request.RelationType, strings.Join(GetListSubsettingSchemaRelationsRelationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingSchemaRelationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingSchemaRelationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingSchemaRelationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingSchemaRelationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingSchemaRelationsResponse wrapper for the ListSubsettingSchemaRelations operation
type ListSubsettingSchemaRelationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingSchemaRelationCollection instances
	SubsettingSchemaRelationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingSchemaRelationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingSchemaRelationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingSchemaRelationsRelationTypeEnum Enum with underlying type: string
type ListSubsettingSchemaRelationsRelationTypeEnum string

// Set of constants representing the allowable values for ListSubsettingSchemaRelationsRelationTypeEnum
const (
	ListSubsettingSchemaRelationsRelationTypeAppDefined ListSubsettingSchemaRelationsRelationTypeEnum = "APP_DEFINED"
	ListSubsettingSchemaRelationsRelationTypeDbDefined  ListSubsettingSchemaRelationsRelationTypeEnum = "DB_DEFINED"
)

var mappingListSubsettingSchemaRelationsRelationTypeEnum = map[string]ListSubsettingSchemaRelationsRelationTypeEnum{
	"APP_DEFINED": ListSubsettingSchemaRelationsRelationTypeAppDefined,
	"DB_DEFINED":  ListSubsettingSchemaRelationsRelationTypeDbDefined,
}

var mappingListSubsettingSchemaRelationsRelationTypeEnumLowerCase = map[string]ListSubsettingSchemaRelationsRelationTypeEnum{
	"app_defined": ListSubsettingSchemaRelationsRelationTypeAppDefined,
	"db_defined":  ListSubsettingSchemaRelationsRelationTypeDbDefined,
}

// GetListSubsettingSchemaRelationsRelationTypeEnumValues Enumerates the set of values for ListSubsettingSchemaRelationsRelationTypeEnum
func GetListSubsettingSchemaRelationsRelationTypeEnumValues() []ListSubsettingSchemaRelationsRelationTypeEnum {
	values := make([]ListSubsettingSchemaRelationsRelationTypeEnum, 0)
	for _, v := range mappingListSubsettingSchemaRelationsRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingSchemaRelationsRelationTypeEnumStringValues Enumerates the set of values in String for ListSubsettingSchemaRelationsRelationTypeEnum
func GetListSubsettingSchemaRelationsRelationTypeEnumStringValues() []string {
	return []string{
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingListSubsettingSchemaRelationsRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingSchemaRelationsRelationTypeEnum(val string) (ListSubsettingSchemaRelationsRelationTypeEnum, bool) {
	enum, ok := mappingListSubsettingSchemaRelationsRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingSchemaRelationsSortOrderEnum Enum with underlying type: string
type ListSubsettingSchemaRelationsSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingSchemaRelationsSortOrderEnum
const (
	ListSubsettingSchemaRelationsSortOrderAsc  ListSubsettingSchemaRelationsSortOrderEnum = "ASC"
	ListSubsettingSchemaRelationsSortOrderDesc ListSubsettingSchemaRelationsSortOrderEnum = "DESC"
)

var mappingListSubsettingSchemaRelationsSortOrderEnum = map[string]ListSubsettingSchemaRelationsSortOrderEnum{
	"ASC":  ListSubsettingSchemaRelationsSortOrderAsc,
	"DESC": ListSubsettingSchemaRelationsSortOrderDesc,
}

var mappingListSubsettingSchemaRelationsSortOrderEnumLowerCase = map[string]ListSubsettingSchemaRelationsSortOrderEnum{
	"asc":  ListSubsettingSchemaRelationsSortOrderAsc,
	"desc": ListSubsettingSchemaRelationsSortOrderDesc,
}

// GetListSubsettingSchemaRelationsSortOrderEnumValues Enumerates the set of values for ListSubsettingSchemaRelationsSortOrderEnum
func GetListSubsettingSchemaRelationsSortOrderEnumValues() []ListSubsettingSchemaRelationsSortOrderEnum {
	values := make([]ListSubsettingSchemaRelationsSortOrderEnum, 0)
	for _, v := range mappingListSubsettingSchemaRelationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingSchemaRelationsSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingSchemaRelationsSortOrderEnum
func GetListSubsettingSchemaRelationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingSchemaRelationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingSchemaRelationsSortOrderEnum(val string) (ListSubsettingSchemaRelationsSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingSchemaRelationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingSchemaRelationsSortByEnum Enum with underlying type: string
type ListSubsettingSchemaRelationsSortByEnum string

// Set of constants representing the allowable values for ListSubsettingSchemaRelationsSortByEnum
const (
	ListSubsettingSchemaRelationsSortByRelationtype ListSubsettingSchemaRelationsSortByEnum = "relationType"
	ListSubsettingSchemaRelationsSortBySchemaname   ListSubsettingSchemaRelationsSortByEnum = "schemaName"
	ListSubsettingSchemaRelationsSortByObjectname   ListSubsettingSchemaRelationsSortByEnum = "objectName"
)

var mappingListSubsettingSchemaRelationsSortByEnum = map[string]ListSubsettingSchemaRelationsSortByEnum{
	"relationType": ListSubsettingSchemaRelationsSortByRelationtype,
	"schemaName":   ListSubsettingSchemaRelationsSortBySchemaname,
	"objectName":   ListSubsettingSchemaRelationsSortByObjectname,
}

var mappingListSubsettingSchemaRelationsSortByEnumLowerCase = map[string]ListSubsettingSchemaRelationsSortByEnum{
	"relationtype": ListSubsettingSchemaRelationsSortByRelationtype,
	"schemaname":   ListSubsettingSchemaRelationsSortBySchemaname,
	"objectname":   ListSubsettingSchemaRelationsSortByObjectname,
}

// GetListSubsettingSchemaRelationsSortByEnumValues Enumerates the set of values for ListSubsettingSchemaRelationsSortByEnum
func GetListSubsettingSchemaRelationsSortByEnumValues() []ListSubsettingSchemaRelationsSortByEnum {
	values := make([]ListSubsettingSchemaRelationsSortByEnum, 0)
	for _, v := range mappingListSubsettingSchemaRelationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingSchemaRelationsSortByEnumStringValues Enumerates the set of values in String for ListSubsettingSchemaRelationsSortByEnum
func GetListSubsettingSchemaRelationsSortByEnumStringValues() []string {
	return []string{
		"relationType",
		"schemaName",
		"objectName",
	}
}

// GetMappingListSubsettingSchemaRelationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingSchemaRelationsSortByEnum(val string) (ListSubsettingSchemaRelationsSortByEnum, bool) {
	enum, ok := mappingListSubsettingSchemaRelationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
