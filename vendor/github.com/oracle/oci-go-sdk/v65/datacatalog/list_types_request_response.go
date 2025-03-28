// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTypesRequest wrapper for the ListTypes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListTypes.go.html to see an example of how to use ListTypesRequest.
type ListTypesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListTypesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Indicates whether the type is internal, making it unavailable for use by metadata elements.
	IsInternal *string `mandatory:"false" contributesTo:"query" name:"isInternal"`

	// Indicates whether the type can be used for tagging metadata elements.
	IsTag *string `mandatory:"false" contributesTo:"query" name:"isTag"`

	// Indicates whether the type is approved for use as a classifying object.
	IsApproved *string `mandatory:"false" contributesTo:"query" name:"isApproved"`

	// Data type as defined in an external system.
	ExternalTypeName *string `mandatory:"false" contributesTo:"query" name:"externalTypeName"`

	// Indicates the category of this type . For example, data assets or connections.
	TypeCategory *string `mandatory:"false" contributesTo:"query" name:"typeCategory"`

	// Specifies the fields to return in a type summary response.
	Fields []ListTypesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTypesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTypesLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListTypesFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListTypesFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTypesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTypesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTypesResponse wrapper for the ListTypes operation
type ListTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TypeCollection instances
	TypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTypesLifecycleStateEnum Enum with underlying type: string
type ListTypesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTypesLifecycleStateEnum
const (
	ListTypesLifecycleStateCreating ListTypesLifecycleStateEnum = "CREATING"
	ListTypesLifecycleStateActive   ListTypesLifecycleStateEnum = "ACTIVE"
	ListTypesLifecycleStateInactive ListTypesLifecycleStateEnum = "INACTIVE"
	ListTypesLifecycleStateUpdating ListTypesLifecycleStateEnum = "UPDATING"
	ListTypesLifecycleStateDeleting ListTypesLifecycleStateEnum = "DELETING"
	ListTypesLifecycleStateDeleted  ListTypesLifecycleStateEnum = "DELETED"
	ListTypesLifecycleStateFailed   ListTypesLifecycleStateEnum = "FAILED"
	ListTypesLifecycleStateMoving   ListTypesLifecycleStateEnum = "MOVING"
)

var mappingListTypesLifecycleStateEnum = map[string]ListTypesLifecycleStateEnum{
	"CREATING": ListTypesLifecycleStateCreating,
	"ACTIVE":   ListTypesLifecycleStateActive,
	"INACTIVE": ListTypesLifecycleStateInactive,
	"UPDATING": ListTypesLifecycleStateUpdating,
	"DELETING": ListTypesLifecycleStateDeleting,
	"DELETED":  ListTypesLifecycleStateDeleted,
	"FAILED":   ListTypesLifecycleStateFailed,
	"MOVING":   ListTypesLifecycleStateMoving,
}

var mappingListTypesLifecycleStateEnumLowerCase = map[string]ListTypesLifecycleStateEnum{
	"creating": ListTypesLifecycleStateCreating,
	"active":   ListTypesLifecycleStateActive,
	"inactive": ListTypesLifecycleStateInactive,
	"updating": ListTypesLifecycleStateUpdating,
	"deleting": ListTypesLifecycleStateDeleting,
	"deleted":  ListTypesLifecycleStateDeleted,
	"failed":   ListTypesLifecycleStateFailed,
	"moving":   ListTypesLifecycleStateMoving,
}

// GetListTypesLifecycleStateEnumValues Enumerates the set of values for ListTypesLifecycleStateEnum
func GetListTypesLifecycleStateEnumValues() []ListTypesLifecycleStateEnum {
	values := make([]ListTypesLifecycleStateEnum, 0)
	for _, v := range mappingListTypesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTypesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTypesLifecycleStateEnum
func GetListTypesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"MOVING",
	}
}

// GetMappingListTypesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTypesLifecycleStateEnum(val string) (ListTypesLifecycleStateEnum, bool) {
	enum, ok := mappingListTypesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTypesFieldsEnum Enum with underlying type: string
type ListTypesFieldsEnum string

// Set of constants representing the allowable values for ListTypesFieldsEnum
const (
	ListTypesFieldsKey            ListTypesFieldsEnum = "key"
	ListTypesFieldsDescription    ListTypesFieldsEnum = "description"
	ListTypesFieldsName           ListTypesFieldsEnum = "name"
	ListTypesFieldsCatalogid      ListTypesFieldsEnum = "catalogId"
	ListTypesFieldsLifecyclestate ListTypesFieldsEnum = "lifecycleState"
	ListTypesFieldsTypecategory   ListTypesFieldsEnum = "typeCategory"
	ListTypesFieldsUri            ListTypesFieldsEnum = "uri"
)

var mappingListTypesFieldsEnum = map[string]ListTypesFieldsEnum{
	"key":            ListTypesFieldsKey,
	"description":    ListTypesFieldsDescription,
	"name":           ListTypesFieldsName,
	"catalogId":      ListTypesFieldsCatalogid,
	"lifecycleState": ListTypesFieldsLifecyclestate,
	"typeCategory":   ListTypesFieldsTypecategory,
	"uri":            ListTypesFieldsUri,
}

var mappingListTypesFieldsEnumLowerCase = map[string]ListTypesFieldsEnum{
	"key":            ListTypesFieldsKey,
	"description":    ListTypesFieldsDescription,
	"name":           ListTypesFieldsName,
	"catalogid":      ListTypesFieldsCatalogid,
	"lifecyclestate": ListTypesFieldsLifecyclestate,
	"typecategory":   ListTypesFieldsTypecategory,
	"uri":            ListTypesFieldsUri,
}

// GetListTypesFieldsEnumValues Enumerates the set of values for ListTypesFieldsEnum
func GetListTypesFieldsEnumValues() []ListTypesFieldsEnum {
	values := make([]ListTypesFieldsEnum, 0)
	for _, v := range mappingListTypesFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListTypesFieldsEnumStringValues Enumerates the set of values in String for ListTypesFieldsEnum
func GetListTypesFieldsEnumStringValues() []string {
	return []string{
		"key",
		"description",
		"name",
		"catalogId",
		"lifecycleState",
		"typeCategory",
		"uri",
	}
}

// GetMappingListTypesFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTypesFieldsEnum(val string) (ListTypesFieldsEnum, bool) {
	enum, ok := mappingListTypesFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTypesSortByEnum Enum with underlying type: string
type ListTypesSortByEnum string

// Set of constants representing the allowable values for ListTypesSortByEnum
const (
	ListTypesSortByTimecreated ListTypesSortByEnum = "TIMECREATED"
	ListTypesSortByDisplayname ListTypesSortByEnum = "DISPLAYNAME"
)

var mappingListTypesSortByEnum = map[string]ListTypesSortByEnum{
	"TIMECREATED": ListTypesSortByTimecreated,
	"DISPLAYNAME": ListTypesSortByDisplayname,
}

var mappingListTypesSortByEnumLowerCase = map[string]ListTypesSortByEnum{
	"timecreated": ListTypesSortByTimecreated,
	"displayname": ListTypesSortByDisplayname,
}

// GetListTypesSortByEnumValues Enumerates the set of values for ListTypesSortByEnum
func GetListTypesSortByEnumValues() []ListTypesSortByEnum {
	values := make([]ListTypesSortByEnum, 0)
	for _, v := range mappingListTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTypesSortByEnumStringValues Enumerates the set of values in String for ListTypesSortByEnum
func GetListTypesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTypesSortByEnum(val string) (ListTypesSortByEnum, bool) {
	enum, ok := mappingListTypesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTypesSortOrderEnum Enum with underlying type: string
type ListTypesSortOrderEnum string

// Set of constants representing the allowable values for ListTypesSortOrderEnum
const (
	ListTypesSortOrderAsc  ListTypesSortOrderEnum = "ASC"
	ListTypesSortOrderDesc ListTypesSortOrderEnum = "DESC"
)

var mappingListTypesSortOrderEnum = map[string]ListTypesSortOrderEnum{
	"ASC":  ListTypesSortOrderAsc,
	"DESC": ListTypesSortOrderDesc,
}

var mappingListTypesSortOrderEnumLowerCase = map[string]ListTypesSortOrderEnum{
	"asc":  ListTypesSortOrderAsc,
	"desc": ListTypesSortOrderDesc,
}

// GetListTypesSortOrderEnumValues Enumerates the set of values for ListTypesSortOrderEnum
func GetListTypesSortOrderEnumValues() []ListTypesSortOrderEnum {
	values := make([]ListTypesSortOrderEnum, 0)
	for _, v := range mappingListTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTypesSortOrderEnumStringValues Enumerates the set of values in String for ListTypesSortOrderEnum
func GetListTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTypesSortOrderEnum(val string) (ListTypesSortOrderEnum, bool) {
	enum, ok := mappingListTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
