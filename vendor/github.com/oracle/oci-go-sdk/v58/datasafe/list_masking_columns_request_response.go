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

// ListMaskingColumnsRequest wrapper for the ListMaskingColumns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingColumns.go.html to see an example of how to use ListMaskingColumnsRequest.
type ListMaskingColumnsRequest struct {

	// The OCID of the masking policy.
	MaskingPolicyId *string `mandatory:"true" contributesTo:"path" name:"maskingPolicyId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingColumnsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for timeCreated is descending.
	// The default order for other fields is ascending.
	SortBy ListMaskingColumnsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only the resources that match the specified lifecycle states.
	MaskingColumnLifecycleState ListMaskingColumnsMaskingColumnLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"maskingColumnLifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the specified data types.
	DataType []ListMaskingColumnsDataTypeEnum `contributesTo:"query" name:"dataType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object type.
	ObjectType []ListMaskingColumnsObjectTypeEnum `contributesTo:"query" name:"objectType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only the resources that match the specified masking column group.
	MaskingColumnGroup []string `contributesTo:"query" name:"maskingColumnGroup" collectionFormat:"multi"`

	// A filter to return only items related to a specific sensitive type OCID.
	SensitiveTypeId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypeId"`

	// A filter to return the masking column resources based on the value of their isMaskingEnabled attribute.
	// A value of true returns only those columns for which masking is enabled. A value of false returns only those columns
	// for which masking is disabled. Omitting this parameter returns all the masking columns in a masking policy.
	IsMaskingEnabled *bool `mandatory:"false" contributesTo:"query" name:"isMaskingEnabled"`

	// A filter to return masking columns based on whether the assigned masking formats need a
	// seed value for masking. A value of true returns those masking columns that are using
	// Deterministic Encryption or Deterministic Substitution masking format.
	IsSeedRequired *bool `mandatory:"false" contributesTo:"query" name:"isSeedRequired"`

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

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingColumnsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingColumnsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingColumnsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingColumnsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingColumnsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingColumnsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingColumnsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingColumnsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingColumnsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingColumnsMaskingColumnLifecycleStateEnum(string(request.MaskingColumnLifecycleState)); !ok && request.MaskingColumnLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaskingColumnLifecycleState: %s. Supported values are: %s.", request.MaskingColumnLifecycleState, strings.Join(GetListMaskingColumnsMaskingColumnLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.DataType {
		if _, ok := GetMappingListMaskingColumnsDataTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", val, strings.Join(GetListMaskingColumnsDataTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ObjectType {
		if _, ok := GetMappingListMaskingColumnsObjectTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", val, strings.Join(GetListMaskingColumnsObjectTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingColumnsResponse wrapper for the ListMaskingColumns operation
type ListMaskingColumnsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingColumnCollection instances
	MaskingColumnCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingColumnsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingColumnsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingColumnsSortOrderEnum Enum with underlying type: string
type ListMaskingColumnsSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingColumnsSortOrderEnum
const (
	ListMaskingColumnsSortOrderAsc  ListMaskingColumnsSortOrderEnum = "ASC"
	ListMaskingColumnsSortOrderDesc ListMaskingColumnsSortOrderEnum = "DESC"
)

var mappingListMaskingColumnsSortOrderEnum = map[string]ListMaskingColumnsSortOrderEnum{
	"ASC":  ListMaskingColumnsSortOrderAsc,
	"DESC": ListMaskingColumnsSortOrderDesc,
}

// GetListMaskingColumnsSortOrderEnumValues Enumerates the set of values for ListMaskingColumnsSortOrderEnum
func GetListMaskingColumnsSortOrderEnumValues() []ListMaskingColumnsSortOrderEnum {
	values := make([]ListMaskingColumnsSortOrderEnum, 0)
	for _, v := range mappingListMaskingColumnsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingColumnsSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingColumnsSortOrderEnum
func GetListMaskingColumnsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingColumnsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingColumnsSortOrderEnum(val string) (ListMaskingColumnsSortOrderEnum, bool) {
	mappingListMaskingColumnsSortOrderEnumIgnoreCase := make(map[string]ListMaskingColumnsSortOrderEnum)
	for k, v := range mappingListMaskingColumnsSortOrderEnum {
		mappingListMaskingColumnsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingColumnsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingColumnsSortByEnum Enum with underlying type: string
type ListMaskingColumnsSortByEnum string

// Set of constants representing the allowable values for ListMaskingColumnsSortByEnum
const (
	ListMaskingColumnsSortByTimecreated ListMaskingColumnsSortByEnum = "timeCreated"
	ListMaskingColumnsSortBySchemaname  ListMaskingColumnsSortByEnum = "schemaName"
	ListMaskingColumnsSortByObjectname  ListMaskingColumnsSortByEnum = "objectName"
)

var mappingListMaskingColumnsSortByEnum = map[string]ListMaskingColumnsSortByEnum{
	"timeCreated": ListMaskingColumnsSortByTimecreated,
	"schemaName":  ListMaskingColumnsSortBySchemaname,
	"objectName":  ListMaskingColumnsSortByObjectname,
}

// GetListMaskingColumnsSortByEnumValues Enumerates the set of values for ListMaskingColumnsSortByEnum
func GetListMaskingColumnsSortByEnumValues() []ListMaskingColumnsSortByEnum {
	values := make([]ListMaskingColumnsSortByEnum, 0)
	for _, v := range mappingListMaskingColumnsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingColumnsSortByEnumStringValues Enumerates the set of values in String for ListMaskingColumnsSortByEnum
func GetListMaskingColumnsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"schemaName",
		"objectName",
	}
}

// GetMappingListMaskingColumnsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingColumnsSortByEnum(val string) (ListMaskingColumnsSortByEnum, bool) {
	mappingListMaskingColumnsSortByEnumIgnoreCase := make(map[string]ListMaskingColumnsSortByEnum)
	for k, v := range mappingListMaskingColumnsSortByEnum {
		mappingListMaskingColumnsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingColumnsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingColumnsMaskingColumnLifecycleStateEnum Enum with underlying type: string
type ListMaskingColumnsMaskingColumnLifecycleStateEnum string

// Set of constants representing the allowable values for ListMaskingColumnsMaskingColumnLifecycleStateEnum
const (
	ListMaskingColumnsMaskingColumnLifecycleStateCreating       ListMaskingColumnsMaskingColumnLifecycleStateEnum = "CREATING"
	ListMaskingColumnsMaskingColumnLifecycleStateActive         ListMaskingColumnsMaskingColumnLifecycleStateEnum = "ACTIVE"
	ListMaskingColumnsMaskingColumnLifecycleStateUpdating       ListMaskingColumnsMaskingColumnLifecycleStateEnum = "UPDATING"
	ListMaskingColumnsMaskingColumnLifecycleStateDeleting       ListMaskingColumnsMaskingColumnLifecycleStateEnum = "DELETING"
	ListMaskingColumnsMaskingColumnLifecycleStateNeedsAttention ListMaskingColumnsMaskingColumnLifecycleStateEnum = "NEEDS_ATTENTION"
	ListMaskingColumnsMaskingColumnLifecycleStateFailed         ListMaskingColumnsMaskingColumnLifecycleStateEnum = "FAILED"
)

var mappingListMaskingColumnsMaskingColumnLifecycleStateEnum = map[string]ListMaskingColumnsMaskingColumnLifecycleStateEnum{
	"CREATING":        ListMaskingColumnsMaskingColumnLifecycleStateCreating,
	"ACTIVE":          ListMaskingColumnsMaskingColumnLifecycleStateActive,
	"UPDATING":        ListMaskingColumnsMaskingColumnLifecycleStateUpdating,
	"DELETING":        ListMaskingColumnsMaskingColumnLifecycleStateDeleting,
	"NEEDS_ATTENTION": ListMaskingColumnsMaskingColumnLifecycleStateNeedsAttention,
	"FAILED":          ListMaskingColumnsMaskingColumnLifecycleStateFailed,
}

// GetListMaskingColumnsMaskingColumnLifecycleStateEnumValues Enumerates the set of values for ListMaskingColumnsMaskingColumnLifecycleStateEnum
func GetListMaskingColumnsMaskingColumnLifecycleStateEnumValues() []ListMaskingColumnsMaskingColumnLifecycleStateEnum {
	values := make([]ListMaskingColumnsMaskingColumnLifecycleStateEnum, 0)
	for _, v := range mappingListMaskingColumnsMaskingColumnLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingColumnsMaskingColumnLifecycleStateEnumStringValues Enumerates the set of values in String for ListMaskingColumnsMaskingColumnLifecycleStateEnum
func GetListMaskingColumnsMaskingColumnLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingListMaskingColumnsMaskingColumnLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingColumnsMaskingColumnLifecycleStateEnum(val string) (ListMaskingColumnsMaskingColumnLifecycleStateEnum, bool) {
	mappingListMaskingColumnsMaskingColumnLifecycleStateEnumIgnoreCase := make(map[string]ListMaskingColumnsMaskingColumnLifecycleStateEnum)
	for k, v := range mappingListMaskingColumnsMaskingColumnLifecycleStateEnum {
		mappingListMaskingColumnsMaskingColumnLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingColumnsMaskingColumnLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingColumnsDataTypeEnum Enum with underlying type: string
type ListMaskingColumnsDataTypeEnum string

// Set of constants representing the allowable values for ListMaskingColumnsDataTypeEnum
const (
	ListMaskingColumnsDataTypeCharacter ListMaskingColumnsDataTypeEnum = "CHARACTER"
	ListMaskingColumnsDataTypeDate      ListMaskingColumnsDataTypeEnum = "DATE"
	ListMaskingColumnsDataTypeLob       ListMaskingColumnsDataTypeEnum = "LOB"
	ListMaskingColumnsDataTypeNumeric   ListMaskingColumnsDataTypeEnum = "NUMERIC"
)

var mappingListMaskingColumnsDataTypeEnum = map[string]ListMaskingColumnsDataTypeEnum{
	"CHARACTER": ListMaskingColumnsDataTypeCharacter,
	"DATE":      ListMaskingColumnsDataTypeDate,
	"LOB":       ListMaskingColumnsDataTypeLob,
	"NUMERIC":   ListMaskingColumnsDataTypeNumeric,
}

// GetListMaskingColumnsDataTypeEnumValues Enumerates the set of values for ListMaskingColumnsDataTypeEnum
func GetListMaskingColumnsDataTypeEnumValues() []ListMaskingColumnsDataTypeEnum {
	values := make([]ListMaskingColumnsDataTypeEnum, 0)
	for _, v := range mappingListMaskingColumnsDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingColumnsDataTypeEnumStringValues Enumerates the set of values in String for ListMaskingColumnsDataTypeEnum
func GetListMaskingColumnsDataTypeEnumStringValues() []string {
	return []string{
		"CHARACTER",
		"DATE",
		"LOB",
		"NUMERIC",
	}
}

// GetMappingListMaskingColumnsDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingColumnsDataTypeEnum(val string) (ListMaskingColumnsDataTypeEnum, bool) {
	mappingListMaskingColumnsDataTypeEnumIgnoreCase := make(map[string]ListMaskingColumnsDataTypeEnum)
	for k, v := range mappingListMaskingColumnsDataTypeEnum {
		mappingListMaskingColumnsDataTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingColumnsDataTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingColumnsObjectTypeEnum Enum with underlying type: string
type ListMaskingColumnsObjectTypeEnum string

// Set of constants representing the allowable values for ListMaskingColumnsObjectTypeEnum
const (
	ListMaskingColumnsObjectTypeAll            ListMaskingColumnsObjectTypeEnum = "ALL"
	ListMaskingColumnsObjectTypeTable          ListMaskingColumnsObjectTypeEnum = "TABLE"
	ListMaskingColumnsObjectTypeEditioningView ListMaskingColumnsObjectTypeEnum = "EDITIONING_VIEW"
)

var mappingListMaskingColumnsObjectTypeEnum = map[string]ListMaskingColumnsObjectTypeEnum{
	"ALL":             ListMaskingColumnsObjectTypeAll,
	"TABLE":           ListMaskingColumnsObjectTypeTable,
	"EDITIONING_VIEW": ListMaskingColumnsObjectTypeEditioningView,
}

// GetListMaskingColumnsObjectTypeEnumValues Enumerates the set of values for ListMaskingColumnsObjectTypeEnum
func GetListMaskingColumnsObjectTypeEnumValues() []ListMaskingColumnsObjectTypeEnum {
	values := make([]ListMaskingColumnsObjectTypeEnum, 0)
	for _, v := range mappingListMaskingColumnsObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingColumnsObjectTypeEnumStringValues Enumerates the set of values in String for ListMaskingColumnsObjectTypeEnum
func GetListMaskingColumnsObjectTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"TABLE",
		"EDITIONING_VIEW",
	}
}

// GetMappingListMaskingColumnsObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingColumnsObjectTypeEnum(val string) (ListMaskingColumnsObjectTypeEnum, bool) {
	mappingListMaskingColumnsObjectTypeEnumIgnoreCase := make(map[string]ListMaskingColumnsObjectTypeEnum)
	for k, v := range mappingListMaskingColumnsObjectTypeEnum {
		mappingListMaskingColumnsObjectTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingColumnsObjectTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
