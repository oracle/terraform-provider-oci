// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListSensitiveColumnsRequest wrapper for the ListSensitiveColumns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveColumns.go.html to see an example of how to use ListSensitiveColumnsRequest.
type ListSensitiveColumnsRequest struct {

	// The OCID of the sensitive data model.
	SensitiveDataModelId *string `mandatory:"true" contributesTo:"path" name:"sensitiveDataModelId"`

	// A filter to return only the resources that were created after the specified date and time, as defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Search for resources that were created before a specific date.
	// Specifying this parameter corresponding `timeCreatedLessThan`
	// parameter will retrieve all resources created before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Search for resources that were updated after a specific date.
	// Specifying this parameter corresponding `timeUpdatedGreaterThanOrEqualTo`
	// parameter will retrieve all resources updated after the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	TimeUpdatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedGreaterThanOrEqualTo"`

	// Search for resources that were updated before a specific date.
	// Specifying this parameter corresponding `timeUpdatedLessThan`
	// parameter will retrieve all resources updated before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	TimeUpdatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedLessThan"`

	// Filters the sensitive column resources with the given lifecycle state values.
	SensitiveColumnLifecycleState ListSensitiveColumnsSensitiveColumnLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"sensitiveColumnLifecycleState" omitEmpty:"true"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object type.
	ObjectType []ListSensitiveColumnsObjectTypeEnum `contributesTo:"query" name:"objectType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only the resources that match the specified data types.
	DataType []string `contributesTo:"query" name:"dataType" collectionFormat:"multi"`

	// A filter to return only the sensitive columns that match the specified status.
	Status []ListSensitiveColumnsStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only the sensitive columns that are associated with one of the sensitive types identified by the specified OCIDs.
	SensitiveTypeId []string `contributesTo:"query" name:"sensitiveTypeId" collectionFormat:"multi"`

	// A filter to return only the sensitive columns that are children of one of the columns identified by the specified keys.
	ParentColumnKey []string `contributesTo:"query" name:"parentColumnKey" collectionFormat:"multi"`

	// A filter to return sensitive columns based on their relationship with their parent columns. If set to NONE,
	// it returns the sensitive columns that do not have any parent. The response includes the parent columns as
	// well as the independent columns that are not in any relationship. If set to APP_DEFINED, it returns all the
	// child columns that have application-level (non-dictionary) relationship with their parents. If set to DB_DEFINED,
	// it returns all the child columns that have database-level (dictionary-defined) relationship with their parents.
	RelationType []ListSensitiveColumnsRelationTypeEnum `contributesTo:"query" name:"relationType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only the sensitive columns that belong to the specified column group.
	ColumnGroup *string `mandatory:"false" contributesTo:"query" name:"columnGroup"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSensitiveColumnsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for timeCreated is descending.
	// The default order for schemaName, objectName, and columnName is ascending.
	SortBy ListSensitiveColumnsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveColumnsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveColumnsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveColumnsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveColumnsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveColumnsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSensitiveColumnsSensitiveColumnLifecycleStateEnum(string(request.SensitiveColumnLifecycleState)); !ok && request.SensitiveColumnLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SensitiveColumnLifecycleState: %s. Supported values are: %s.", request.SensitiveColumnLifecycleState, strings.Join(GetListSensitiveColumnsSensitiveColumnLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.ObjectType {
		if _, ok := GetMappingListSensitiveColumnsObjectTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", val, strings.Join(GetListSensitiveColumnsObjectTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.Status {
		if _, ok := GetMappingListSensitiveColumnsStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetListSensitiveColumnsStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.RelationType {
		if _, ok := GetMappingListSensitiveColumnsRelationTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", val, strings.Join(GetListSensitiveColumnsRelationTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListSensitiveColumnsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSensitiveColumnsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveColumnsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSensitiveColumnsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveColumnsResponse wrapper for the ListSensitiveColumns operation
type ListSensitiveColumnsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveColumnCollection instances
	SensitiveColumnCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveColumnsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveColumnsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveColumnsSensitiveColumnLifecycleStateEnum Enum with underlying type: string
type ListSensitiveColumnsSensitiveColumnLifecycleStateEnum string

// Set of constants representing the allowable values for ListSensitiveColumnsSensitiveColumnLifecycleStateEnum
const (
	ListSensitiveColumnsSensitiveColumnLifecycleStateCreating ListSensitiveColumnsSensitiveColumnLifecycleStateEnum = "CREATING"
	ListSensitiveColumnsSensitiveColumnLifecycleStateActive   ListSensitiveColumnsSensitiveColumnLifecycleStateEnum = "ACTIVE"
	ListSensitiveColumnsSensitiveColumnLifecycleStateUpdating ListSensitiveColumnsSensitiveColumnLifecycleStateEnum = "UPDATING"
	ListSensitiveColumnsSensitiveColumnLifecycleStateDeleting ListSensitiveColumnsSensitiveColumnLifecycleStateEnum = "DELETING"
	ListSensitiveColumnsSensitiveColumnLifecycleStateFailed   ListSensitiveColumnsSensitiveColumnLifecycleStateEnum = "FAILED"
)

var mappingListSensitiveColumnsSensitiveColumnLifecycleStateEnum = map[string]ListSensitiveColumnsSensitiveColumnLifecycleStateEnum{
	"CREATING": ListSensitiveColumnsSensitiveColumnLifecycleStateCreating,
	"ACTIVE":   ListSensitiveColumnsSensitiveColumnLifecycleStateActive,
	"UPDATING": ListSensitiveColumnsSensitiveColumnLifecycleStateUpdating,
	"DELETING": ListSensitiveColumnsSensitiveColumnLifecycleStateDeleting,
	"FAILED":   ListSensitiveColumnsSensitiveColumnLifecycleStateFailed,
}

// GetListSensitiveColumnsSensitiveColumnLifecycleStateEnumValues Enumerates the set of values for ListSensitiveColumnsSensitiveColumnLifecycleStateEnum
func GetListSensitiveColumnsSensitiveColumnLifecycleStateEnumValues() []ListSensitiveColumnsSensitiveColumnLifecycleStateEnum {
	values := make([]ListSensitiveColumnsSensitiveColumnLifecycleStateEnum, 0)
	for _, v := range mappingListSensitiveColumnsSensitiveColumnLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveColumnsSensitiveColumnLifecycleStateEnumStringValues Enumerates the set of values in String for ListSensitiveColumnsSensitiveColumnLifecycleStateEnum
func GetListSensitiveColumnsSensitiveColumnLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

// GetMappingListSensitiveColumnsSensitiveColumnLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveColumnsSensitiveColumnLifecycleStateEnum(val string) (ListSensitiveColumnsSensitiveColumnLifecycleStateEnum, bool) {
	mappingListSensitiveColumnsSensitiveColumnLifecycleStateEnumIgnoreCase := make(map[string]ListSensitiveColumnsSensitiveColumnLifecycleStateEnum)
	for k, v := range mappingListSensitiveColumnsSensitiveColumnLifecycleStateEnum {
		mappingListSensitiveColumnsSensitiveColumnLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSensitiveColumnsSensitiveColumnLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveColumnsObjectTypeEnum Enum with underlying type: string
type ListSensitiveColumnsObjectTypeEnum string

// Set of constants representing the allowable values for ListSensitiveColumnsObjectTypeEnum
const (
	ListSensitiveColumnsObjectTypeAll            ListSensitiveColumnsObjectTypeEnum = "ALL"
	ListSensitiveColumnsObjectTypeTable          ListSensitiveColumnsObjectTypeEnum = "TABLE"
	ListSensitiveColumnsObjectTypeEditioningView ListSensitiveColumnsObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingListSensitiveColumnsObjectTypeEnum = map[string]ListSensitiveColumnsObjectTypeEnum{
	"ALL":             ListSensitiveColumnsObjectTypeAll,
	"TABLE":           ListSensitiveColumnsObjectTypeTable,
	"EDITIONING_VIEW": ListSensitiveColumnsObjectTypeEditioningView,
}

// GetListSensitiveColumnsObjectTypeEnumValues Enumerates the set of values for ListSensitiveColumnsObjectTypeEnum
func GetListSensitiveColumnsObjectTypeEnumValues() []ListSensitiveColumnsObjectTypeEnum {
	values := make([]ListSensitiveColumnsObjectTypeEnum, 0)
	for _, v := range mappingListSensitiveColumnsObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveColumnsObjectTypeEnumStringValues Enumerates the set of values in String for ListSensitiveColumnsObjectTypeEnum
func GetListSensitiveColumnsObjectTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingListSensitiveColumnsObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveColumnsObjectTypeEnum(val string) (ListSensitiveColumnsObjectTypeEnum, bool) {
	mappingListSensitiveColumnsObjectTypeEnumIgnoreCase := make(map[string]ListSensitiveColumnsObjectTypeEnum)
	for k, v := range mappingListSensitiveColumnsObjectTypeEnum {
		mappingListSensitiveColumnsObjectTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSensitiveColumnsObjectTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveColumnsStatusEnum Enum with underlying type: string
type ListSensitiveColumnsStatusEnum string

// Set of constants representing the allowable values for ListSensitiveColumnsStatusEnum
const (
	ListSensitiveColumnsStatusValid   ListSensitiveColumnsStatusEnum = "VALID"
	ListSensitiveColumnsStatusInvalid ListSensitiveColumnsStatusEnum = "INVALID"
)

var mappingListSensitiveColumnsStatusEnum = map[string]ListSensitiveColumnsStatusEnum{
	"VALID":   ListSensitiveColumnsStatusValid,
	"INVALID": ListSensitiveColumnsStatusInvalid,
}

// GetListSensitiveColumnsStatusEnumValues Enumerates the set of values for ListSensitiveColumnsStatusEnum
func GetListSensitiveColumnsStatusEnumValues() []ListSensitiveColumnsStatusEnum {
	values := make([]ListSensitiveColumnsStatusEnum, 0)
	for _, v := range mappingListSensitiveColumnsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveColumnsStatusEnumStringValues Enumerates the set of values in String for ListSensitiveColumnsStatusEnum
func GetListSensitiveColumnsStatusEnumStringValues() []string {
	return []string{
		"VALID",
		"INVALID",
	}
}

// GetMappingListSensitiveColumnsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveColumnsStatusEnum(val string) (ListSensitiveColumnsStatusEnum, bool) {
	mappingListSensitiveColumnsStatusEnumIgnoreCase := make(map[string]ListSensitiveColumnsStatusEnum)
	for k, v := range mappingListSensitiveColumnsStatusEnum {
		mappingListSensitiveColumnsStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSensitiveColumnsStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveColumnsRelationTypeEnum Enum with underlying type: string
type ListSensitiveColumnsRelationTypeEnum string

// Set of constants representing the allowable values for ListSensitiveColumnsRelationTypeEnum
const (
	ListSensitiveColumnsRelationTypeNone       ListSensitiveColumnsRelationTypeEnum = "NONE"
	ListSensitiveColumnsRelationTypeAppDefined ListSensitiveColumnsRelationTypeEnum = "APP_DEFINED"
	ListSensitiveColumnsRelationTypeDbDefined  ListSensitiveColumnsRelationTypeEnum = "DB_DEFINED"
)

var mappingListSensitiveColumnsRelationTypeEnum = map[string]ListSensitiveColumnsRelationTypeEnum{
	"NONE":        ListSensitiveColumnsRelationTypeNone,
	"APP_DEFINED": ListSensitiveColumnsRelationTypeAppDefined,
	"DB_DEFINED":  ListSensitiveColumnsRelationTypeDbDefined,
}

// GetListSensitiveColumnsRelationTypeEnumValues Enumerates the set of values for ListSensitiveColumnsRelationTypeEnum
func GetListSensitiveColumnsRelationTypeEnumValues() []ListSensitiveColumnsRelationTypeEnum {
	values := make([]ListSensitiveColumnsRelationTypeEnum, 0)
	for _, v := range mappingListSensitiveColumnsRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveColumnsRelationTypeEnumStringValues Enumerates the set of values in String for ListSensitiveColumnsRelationTypeEnum
func GetListSensitiveColumnsRelationTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"APP_DEFINED",
		"DB_DEFINED",
	}
}

// GetMappingListSensitiveColumnsRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveColumnsRelationTypeEnum(val string) (ListSensitiveColumnsRelationTypeEnum, bool) {
	mappingListSensitiveColumnsRelationTypeEnumIgnoreCase := make(map[string]ListSensitiveColumnsRelationTypeEnum)
	for k, v := range mappingListSensitiveColumnsRelationTypeEnum {
		mappingListSensitiveColumnsRelationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSensitiveColumnsRelationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveColumnsSortOrderEnum Enum with underlying type: string
type ListSensitiveColumnsSortOrderEnum string

// Set of constants representing the allowable values for ListSensitiveColumnsSortOrderEnum
const (
	ListSensitiveColumnsSortOrderAsc  ListSensitiveColumnsSortOrderEnum = "ASC"
	ListSensitiveColumnsSortOrderDesc ListSensitiveColumnsSortOrderEnum = "DESC"
)

var mappingListSensitiveColumnsSortOrderEnum = map[string]ListSensitiveColumnsSortOrderEnum{
	"ASC":  ListSensitiveColumnsSortOrderAsc,
	"DESC": ListSensitiveColumnsSortOrderDesc,
}

// GetListSensitiveColumnsSortOrderEnumValues Enumerates the set of values for ListSensitiveColumnsSortOrderEnum
func GetListSensitiveColumnsSortOrderEnumValues() []ListSensitiveColumnsSortOrderEnum {
	values := make([]ListSensitiveColumnsSortOrderEnum, 0)
	for _, v := range mappingListSensitiveColumnsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveColumnsSortOrderEnumStringValues Enumerates the set of values in String for ListSensitiveColumnsSortOrderEnum
func GetListSensitiveColumnsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSensitiveColumnsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveColumnsSortOrderEnum(val string) (ListSensitiveColumnsSortOrderEnum, bool) {
	mappingListSensitiveColumnsSortOrderEnumIgnoreCase := make(map[string]ListSensitiveColumnsSortOrderEnum)
	for k, v := range mappingListSensitiveColumnsSortOrderEnum {
		mappingListSensitiveColumnsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSensitiveColumnsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveColumnsSortByEnum Enum with underlying type: string
type ListSensitiveColumnsSortByEnum string

// Set of constants representing the allowable values for ListSensitiveColumnsSortByEnum
const (
	ListSensitiveColumnsSortByTimecreated ListSensitiveColumnsSortByEnum = "timeCreated"
	ListSensitiveColumnsSortBySchemaname  ListSensitiveColumnsSortByEnum = "schemaName"
	ListSensitiveColumnsSortByObjectname  ListSensitiveColumnsSortByEnum = "objectName"
	ListSensitiveColumnsSortByColumnname  ListSensitiveColumnsSortByEnum = "columnName"
)

var mappingListSensitiveColumnsSortByEnum = map[string]ListSensitiveColumnsSortByEnum{
	"timeCreated": ListSensitiveColumnsSortByTimecreated,
	"schemaName":  ListSensitiveColumnsSortBySchemaname,
	"objectName":  ListSensitiveColumnsSortByObjectname,
	"columnName":  ListSensitiveColumnsSortByColumnname,
}

// GetListSensitiveColumnsSortByEnumValues Enumerates the set of values for ListSensitiveColumnsSortByEnum
func GetListSensitiveColumnsSortByEnumValues() []ListSensitiveColumnsSortByEnum {
	values := make([]ListSensitiveColumnsSortByEnum, 0)
	for _, v := range mappingListSensitiveColumnsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveColumnsSortByEnumStringValues Enumerates the set of values in String for ListSensitiveColumnsSortByEnum
func GetListSensitiveColumnsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"schemaName",
		"objectName",
		"columnName",
	}
}

// GetMappingListSensitiveColumnsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveColumnsSortByEnum(val string) (ListSensitiveColumnsSortByEnum, bool) {
	mappingListSensitiveColumnsSortByEnumIgnoreCase := make(map[string]ListSensitiveColumnsSortByEnum)
	for k, v := range mappingListSensitiveColumnsSortByEnum {
		mappingListSensitiveColumnsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSensitiveColumnsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
