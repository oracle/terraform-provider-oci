// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTranslatorsRequest wrapper for the ListTranslators operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListTranslators.go.html to see an example of how to use ListTranslatorsRequest.
type ListTranslatorsRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Translator identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// List only Translators of this type.
	Type ListTranslatorsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// List only Translators with this name. Translator names are unique and may not change.
	// Example: `MyTranslator`
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// List only the resources that are in this lifecycle state.
	LifecycleState ListTranslatorsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTranslatorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `timeCreated`.
	// The default sort order for `timeCreated` and `timeUpdated` is descending.
	// For all other sort fields the default sort order is ascending.
	SortBy ListTranslatorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTranslatorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTranslatorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTranslatorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTranslatorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTranslatorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTranslatorsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListTranslatorsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTranslatorsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTranslatorsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTranslatorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTranslatorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTranslatorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTranslatorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTranslatorsResponse wrapper for the ListTranslators operation
type ListTranslatorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TranslatorCollection instances
	TranslatorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of results that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListTranslatorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTranslatorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTranslatorsTypeEnum Enum with underlying type: string
type ListTranslatorsTypeEnum string

// Set of constants representing the allowable values for ListTranslatorsTypeEnum
const (
	ListTranslatorsTypeGoogle    ListTranslatorsTypeEnum = "GOOGLE"
	ListTranslatorsTypeMicrosoft ListTranslatorsTypeEnum = "MICROSOFT"
)

var mappingListTranslatorsTypeEnum = map[string]ListTranslatorsTypeEnum{
	"GOOGLE":    ListTranslatorsTypeGoogle,
	"MICROSOFT": ListTranslatorsTypeMicrosoft,
}

var mappingListTranslatorsTypeEnumLowerCase = map[string]ListTranslatorsTypeEnum{
	"google":    ListTranslatorsTypeGoogle,
	"microsoft": ListTranslatorsTypeMicrosoft,
}

// GetListTranslatorsTypeEnumValues Enumerates the set of values for ListTranslatorsTypeEnum
func GetListTranslatorsTypeEnumValues() []ListTranslatorsTypeEnum {
	values := make([]ListTranslatorsTypeEnum, 0)
	for _, v := range mappingListTranslatorsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListTranslatorsTypeEnumStringValues Enumerates the set of values in String for ListTranslatorsTypeEnum
func GetListTranslatorsTypeEnumStringValues() []string {
	return []string{
		"GOOGLE",
		"MICROSOFT",
	}
}

// GetMappingListTranslatorsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTranslatorsTypeEnum(val string) (ListTranslatorsTypeEnum, bool) {
	enum, ok := mappingListTranslatorsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTranslatorsLifecycleStateEnum Enum with underlying type: string
type ListTranslatorsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTranslatorsLifecycleStateEnum
const (
	ListTranslatorsLifecycleStateCreating ListTranslatorsLifecycleStateEnum = "CREATING"
	ListTranslatorsLifecycleStateUpdating ListTranslatorsLifecycleStateEnum = "UPDATING"
	ListTranslatorsLifecycleStateActive   ListTranslatorsLifecycleStateEnum = "ACTIVE"
	ListTranslatorsLifecycleStateInactive ListTranslatorsLifecycleStateEnum = "INACTIVE"
	ListTranslatorsLifecycleStateDeleting ListTranslatorsLifecycleStateEnum = "DELETING"
	ListTranslatorsLifecycleStateDeleted  ListTranslatorsLifecycleStateEnum = "DELETED"
	ListTranslatorsLifecycleStateFailed   ListTranslatorsLifecycleStateEnum = "FAILED"
)

var mappingListTranslatorsLifecycleStateEnum = map[string]ListTranslatorsLifecycleStateEnum{
	"CREATING": ListTranslatorsLifecycleStateCreating,
	"UPDATING": ListTranslatorsLifecycleStateUpdating,
	"ACTIVE":   ListTranslatorsLifecycleStateActive,
	"INACTIVE": ListTranslatorsLifecycleStateInactive,
	"DELETING": ListTranslatorsLifecycleStateDeleting,
	"DELETED":  ListTranslatorsLifecycleStateDeleted,
	"FAILED":   ListTranslatorsLifecycleStateFailed,
}

var mappingListTranslatorsLifecycleStateEnumLowerCase = map[string]ListTranslatorsLifecycleStateEnum{
	"creating": ListTranslatorsLifecycleStateCreating,
	"updating": ListTranslatorsLifecycleStateUpdating,
	"active":   ListTranslatorsLifecycleStateActive,
	"inactive": ListTranslatorsLifecycleStateInactive,
	"deleting": ListTranslatorsLifecycleStateDeleting,
	"deleted":  ListTranslatorsLifecycleStateDeleted,
	"failed":   ListTranslatorsLifecycleStateFailed,
}

// GetListTranslatorsLifecycleStateEnumValues Enumerates the set of values for ListTranslatorsLifecycleStateEnum
func GetListTranslatorsLifecycleStateEnumValues() []ListTranslatorsLifecycleStateEnum {
	values := make([]ListTranslatorsLifecycleStateEnum, 0)
	for _, v := range mappingListTranslatorsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTranslatorsLifecycleStateEnumStringValues Enumerates the set of values in String for ListTranslatorsLifecycleStateEnum
func GetListTranslatorsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListTranslatorsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTranslatorsLifecycleStateEnum(val string) (ListTranslatorsLifecycleStateEnum, bool) {
	enum, ok := mappingListTranslatorsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTranslatorsSortOrderEnum Enum with underlying type: string
type ListTranslatorsSortOrderEnum string

// Set of constants representing the allowable values for ListTranslatorsSortOrderEnum
const (
	ListTranslatorsSortOrderAsc  ListTranslatorsSortOrderEnum = "ASC"
	ListTranslatorsSortOrderDesc ListTranslatorsSortOrderEnum = "DESC"
)

var mappingListTranslatorsSortOrderEnum = map[string]ListTranslatorsSortOrderEnum{
	"ASC":  ListTranslatorsSortOrderAsc,
	"DESC": ListTranslatorsSortOrderDesc,
}

var mappingListTranslatorsSortOrderEnumLowerCase = map[string]ListTranslatorsSortOrderEnum{
	"asc":  ListTranslatorsSortOrderAsc,
	"desc": ListTranslatorsSortOrderDesc,
}

// GetListTranslatorsSortOrderEnumValues Enumerates the set of values for ListTranslatorsSortOrderEnum
func GetListTranslatorsSortOrderEnumValues() []ListTranslatorsSortOrderEnum {
	values := make([]ListTranslatorsSortOrderEnum, 0)
	for _, v := range mappingListTranslatorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTranslatorsSortOrderEnumStringValues Enumerates the set of values in String for ListTranslatorsSortOrderEnum
func GetListTranslatorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTranslatorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTranslatorsSortOrderEnum(val string) (ListTranslatorsSortOrderEnum, bool) {
	enum, ok := mappingListTranslatorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTranslatorsSortByEnum Enum with underlying type: string
type ListTranslatorsSortByEnum string

// Set of constants representing the allowable values for ListTranslatorsSortByEnum
const (
	ListTranslatorsSortByTimecreated ListTranslatorsSortByEnum = "timeCreated"
	ListTranslatorsSortByTimeupdated ListTranslatorsSortByEnum = "timeUpdated"
	ListTranslatorsSortByName        ListTranslatorsSortByEnum = "name"
	ListTranslatorsSortByType        ListTranslatorsSortByEnum = "type"
)

var mappingListTranslatorsSortByEnum = map[string]ListTranslatorsSortByEnum{
	"timeCreated": ListTranslatorsSortByTimecreated,
	"timeUpdated": ListTranslatorsSortByTimeupdated,
	"name":        ListTranslatorsSortByName,
	"type":        ListTranslatorsSortByType,
}

var mappingListTranslatorsSortByEnumLowerCase = map[string]ListTranslatorsSortByEnum{
	"timecreated": ListTranslatorsSortByTimecreated,
	"timeupdated": ListTranslatorsSortByTimeupdated,
	"name":        ListTranslatorsSortByName,
	"type":        ListTranslatorsSortByType,
}

// GetListTranslatorsSortByEnumValues Enumerates the set of values for ListTranslatorsSortByEnum
func GetListTranslatorsSortByEnumValues() []ListTranslatorsSortByEnum {
	values := make([]ListTranslatorsSortByEnum, 0)
	for _, v := range mappingListTranslatorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTranslatorsSortByEnumStringValues Enumerates the set of values in String for ListTranslatorsSortByEnum
func GetListTranslatorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"name",
		"type",
	}
}

// GetMappingListTranslatorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTranslatorsSortByEnum(val string) (ListTranslatorsSortByEnum, bool) {
	enum, ok := mappingListTranslatorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
