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

// ListDifferenceColumnsRequest wrapper for the ListDifferenceColumns operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDifferenceColumns.go.html to see an example of how to use ListDifferenceColumnsRequest.
type ListDifferenceColumnsRequest struct {

	// The OCID of the SDM masking policy difference.
	SdmMaskingPolicyDifferenceId *string `mandatory:"true" contributesTo:"path" name:"sdmMaskingPolicyDifferenceId"`

	// A filter to return only the SDM masking policy difference columns that match the specified difference type
	DifferenceType SdmMaskingPolicyDifferenceDifferenceTypeEnum `mandatory:"false" contributesTo:"query" name:"differenceType" omitEmpty:"true"`

	// A filter to return only the SDM masking policy difference columns that match the specified planned action.
	PlannedAction DifferenceColumnPlannedActionEnum `mandatory:"false" contributesTo:"query" name:"plannedAction" omitEmpty:"true"`

	// A filter to return the SDM masking policy difference columns based on the value of their syncStatus attribute.
	SyncStatus DifferenceColumnSyncStatusEnum `mandatory:"false" contributesTo:"query" name:"syncStatus" omitEmpty:"true"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDifferenceColumnsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for schemaName is descending.
	// The default order for differenceType, schemaName, objectName, columnName and plannedAction is ascending.
	SortBy ListDifferenceColumnsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDifferenceColumnsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDifferenceColumnsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDifferenceColumnsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDifferenceColumnsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDifferenceColumnsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSdmMaskingPolicyDifferenceDifferenceTypeEnum(string(request.DifferenceType)); !ok && request.DifferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DifferenceType: %s. Supported values are: %s.", request.DifferenceType, strings.Join(GetSdmMaskingPolicyDifferenceDifferenceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDifferenceColumnPlannedActionEnum(string(request.PlannedAction)); !ok && request.PlannedAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlannedAction: %s. Supported values are: %s.", request.PlannedAction, strings.Join(GetDifferenceColumnPlannedActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDifferenceColumnSyncStatusEnum(string(request.SyncStatus)); !ok && request.SyncStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SyncStatus: %s. Supported values are: %s.", request.SyncStatus, strings.Join(GetDifferenceColumnSyncStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDifferenceColumnsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDifferenceColumnsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDifferenceColumnsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDifferenceColumnsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDifferenceColumnsResponse wrapper for the ListDifferenceColumns operation
type ListDifferenceColumnsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SdmMaskingPolicyDifferenceColumnCollection instances
	SdmMaskingPolicyDifferenceColumnCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDifferenceColumnsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDifferenceColumnsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDifferenceColumnsSortOrderEnum Enum with underlying type: string
type ListDifferenceColumnsSortOrderEnum string

// Set of constants representing the allowable values for ListDifferenceColumnsSortOrderEnum
const (
	ListDifferenceColumnsSortOrderAsc  ListDifferenceColumnsSortOrderEnum = "ASC"
	ListDifferenceColumnsSortOrderDesc ListDifferenceColumnsSortOrderEnum = "DESC"
)

var mappingListDifferenceColumnsSortOrderEnum = map[string]ListDifferenceColumnsSortOrderEnum{
	"ASC":  ListDifferenceColumnsSortOrderAsc,
	"DESC": ListDifferenceColumnsSortOrderDesc,
}

var mappingListDifferenceColumnsSortOrderEnumLowerCase = map[string]ListDifferenceColumnsSortOrderEnum{
	"asc":  ListDifferenceColumnsSortOrderAsc,
	"desc": ListDifferenceColumnsSortOrderDesc,
}

// GetListDifferenceColumnsSortOrderEnumValues Enumerates the set of values for ListDifferenceColumnsSortOrderEnum
func GetListDifferenceColumnsSortOrderEnumValues() []ListDifferenceColumnsSortOrderEnum {
	values := make([]ListDifferenceColumnsSortOrderEnum, 0)
	for _, v := range mappingListDifferenceColumnsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDifferenceColumnsSortOrderEnumStringValues Enumerates the set of values in String for ListDifferenceColumnsSortOrderEnum
func GetListDifferenceColumnsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDifferenceColumnsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDifferenceColumnsSortOrderEnum(val string) (ListDifferenceColumnsSortOrderEnum, bool) {
	enum, ok := mappingListDifferenceColumnsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDifferenceColumnsSortByEnum Enum with underlying type: string
type ListDifferenceColumnsSortByEnum string

// Set of constants representing the allowable values for ListDifferenceColumnsSortByEnum
const (
	ListDifferenceColumnsSortByDifferencetype ListDifferenceColumnsSortByEnum = "differenceType"
	ListDifferenceColumnsSortBySchemaname     ListDifferenceColumnsSortByEnum = "schemaName"
	ListDifferenceColumnsSortByObjectname     ListDifferenceColumnsSortByEnum = "objectName"
	ListDifferenceColumnsSortByColumnname     ListDifferenceColumnsSortByEnum = "columnName"
	ListDifferenceColumnsSortByPlannedaction  ListDifferenceColumnsSortByEnum = "plannedAction"
)

var mappingListDifferenceColumnsSortByEnum = map[string]ListDifferenceColumnsSortByEnum{
	"differenceType": ListDifferenceColumnsSortByDifferencetype,
	"schemaName":     ListDifferenceColumnsSortBySchemaname,
	"objectName":     ListDifferenceColumnsSortByObjectname,
	"columnName":     ListDifferenceColumnsSortByColumnname,
	"plannedAction":  ListDifferenceColumnsSortByPlannedaction,
}

var mappingListDifferenceColumnsSortByEnumLowerCase = map[string]ListDifferenceColumnsSortByEnum{
	"differencetype": ListDifferenceColumnsSortByDifferencetype,
	"schemaname":     ListDifferenceColumnsSortBySchemaname,
	"objectname":     ListDifferenceColumnsSortByObjectname,
	"columnname":     ListDifferenceColumnsSortByColumnname,
	"plannedaction":  ListDifferenceColumnsSortByPlannedaction,
}

// GetListDifferenceColumnsSortByEnumValues Enumerates the set of values for ListDifferenceColumnsSortByEnum
func GetListDifferenceColumnsSortByEnumValues() []ListDifferenceColumnsSortByEnum {
	values := make([]ListDifferenceColumnsSortByEnum, 0)
	for _, v := range mappingListDifferenceColumnsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDifferenceColumnsSortByEnumStringValues Enumerates the set of values in String for ListDifferenceColumnsSortByEnum
func GetListDifferenceColumnsSortByEnumStringValues() []string {
	return []string{
		"differenceType",
		"schemaName",
		"objectName",
		"columnName",
		"plannedAction",
	}
}

// GetMappingListDifferenceColumnsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDifferenceColumnsSortByEnum(val string) (ListDifferenceColumnsSortByEnum, bool) {
	enum, ok := mappingListDifferenceColumnsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
