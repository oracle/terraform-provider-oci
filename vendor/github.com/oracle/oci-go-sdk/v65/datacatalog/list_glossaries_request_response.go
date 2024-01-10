// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGlossariesRequest wrapper for the ListGlossaries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListGlossaries.go.html to see an example of how to use ListGlossariesRequest.
type ListGlossariesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListGlossariesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Specifies the fields to return in a glossary summary response.
	Fields []ListGlossariesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListGlossariesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListGlossariesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListGlossariesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGlossariesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGlossariesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGlossariesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGlossariesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGlossariesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListGlossariesLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListGlossariesFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListGlossariesFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListGlossariesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGlossariesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGlossariesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGlossariesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGlossariesResponse wrapper for the ListGlossaries operation
type ListGlossariesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GlossaryCollection instances
	GlossaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGlossariesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGlossariesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGlossariesLifecycleStateEnum Enum with underlying type: string
type ListGlossariesLifecycleStateEnum string

// Set of constants representing the allowable values for ListGlossariesLifecycleStateEnum
const (
	ListGlossariesLifecycleStateCreating ListGlossariesLifecycleStateEnum = "CREATING"
	ListGlossariesLifecycleStateActive   ListGlossariesLifecycleStateEnum = "ACTIVE"
	ListGlossariesLifecycleStateInactive ListGlossariesLifecycleStateEnum = "INACTIVE"
	ListGlossariesLifecycleStateUpdating ListGlossariesLifecycleStateEnum = "UPDATING"
	ListGlossariesLifecycleStateDeleting ListGlossariesLifecycleStateEnum = "DELETING"
	ListGlossariesLifecycleStateDeleted  ListGlossariesLifecycleStateEnum = "DELETED"
	ListGlossariesLifecycleStateFailed   ListGlossariesLifecycleStateEnum = "FAILED"
	ListGlossariesLifecycleStateMoving   ListGlossariesLifecycleStateEnum = "MOVING"
)

var mappingListGlossariesLifecycleStateEnum = map[string]ListGlossariesLifecycleStateEnum{
	"CREATING": ListGlossariesLifecycleStateCreating,
	"ACTIVE":   ListGlossariesLifecycleStateActive,
	"INACTIVE": ListGlossariesLifecycleStateInactive,
	"UPDATING": ListGlossariesLifecycleStateUpdating,
	"DELETING": ListGlossariesLifecycleStateDeleting,
	"DELETED":  ListGlossariesLifecycleStateDeleted,
	"FAILED":   ListGlossariesLifecycleStateFailed,
	"MOVING":   ListGlossariesLifecycleStateMoving,
}

var mappingListGlossariesLifecycleStateEnumLowerCase = map[string]ListGlossariesLifecycleStateEnum{
	"creating": ListGlossariesLifecycleStateCreating,
	"active":   ListGlossariesLifecycleStateActive,
	"inactive": ListGlossariesLifecycleStateInactive,
	"updating": ListGlossariesLifecycleStateUpdating,
	"deleting": ListGlossariesLifecycleStateDeleting,
	"deleted":  ListGlossariesLifecycleStateDeleted,
	"failed":   ListGlossariesLifecycleStateFailed,
	"moving":   ListGlossariesLifecycleStateMoving,
}

// GetListGlossariesLifecycleStateEnumValues Enumerates the set of values for ListGlossariesLifecycleStateEnum
func GetListGlossariesLifecycleStateEnumValues() []ListGlossariesLifecycleStateEnum {
	values := make([]ListGlossariesLifecycleStateEnum, 0)
	for _, v := range mappingListGlossariesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListGlossariesLifecycleStateEnumStringValues Enumerates the set of values in String for ListGlossariesLifecycleStateEnum
func GetListGlossariesLifecycleStateEnumStringValues() []string {
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

// GetMappingListGlossariesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGlossariesLifecycleStateEnum(val string) (ListGlossariesLifecycleStateEnum, bool) {
	enum, ok := mappingListGlossariesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGlossariesFieldsEnum Enum with underlying type: string
type ListGlossariesFieldsEnum string

// Set of constants representing the allowable values for ListGlossariesFieldsEnum
const (
	ListGlossariesFieldsKey            ListGlossariesFieldsEnum = "key"
	ListGlossariesFieldsDisplayname    ListGlossariesFieldsEnum = "displayName"
	ListGlossariesFieldsDescription    ListGlossariesFieldsEnum = "description"
	ListGlossariesFieldsCatalogid      ListGlossariesFieldsEnum = "catalogId"
	ListGlossariesFieldsLifecyclestate ListGlossariesFieldsEnum = "lifecycleState"
	ListGlossariesFieldsTimecreated    ListGlossariesFieldsEnum = "timeCreated"
	ListGlossariesFieldsUri            ListGlossariesFieldsEnum = "uri"
	ListGlossariesFieldsWorkflowstatus ListGlossariesFieldsEnum = "workflowStatus"
)

var mappingListGlossariesFieldsEnum = map[string]ListGlossariesFieldsEnum{
	"key":            ListGlossariesFieldsKey,
	"displayName":    ListGlossariesFieldsDisplayname,
	"description":    ListGlossariesFieldsDescription,
	"catalogId":      ListGlossariesFieldsCatalogid,
	"lifecycleState": ListGlossariesFieldsLifecyclestate,
	"timeCreated":    ListGlossariesFieldsTimecreated,
	"uri":            ListGlossariesFieldsUri,
	"workflowStatus": ListGlossariesFieldsWorkflowstatus,
}

var mappingListGlossariesFieldsEnumLowerCase = map[string]ListGlossariesFieldsEnum{
	"key":            ListGlossariesFieldsKey,
	"displayname":    ListGlossariesFieldsDisplayname,
	"description":    ListGlossariesFieldsDescription,
	"catalogid":      ListGlossariesFieldsCatalogid,
	"lifecyclestate": ListGlossariesFieldsLifecyclestate,
	"timecreated":    ListGlossariesFieldsTimecreated,
	"uri":            ListGlossariesFieldsUri,
	"workflowstatus": ListGlossariesFieldsWorkflowstatus,
}

// GetListGlossariesFieldsEnumValues Enumerates the set of values for ListGlossariesFieldsEnum
func GetListGlossariesFieldsEnumValues() []ListGlossariesFieldsEnum {
	values := make([]ListGlossariesFieldsEnum, 0)
	for _, v := range mappingListGlossariesFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListGlossariesFieldsEnumStringValues Enumerates the set of values in String for ListGlossariesFieldsEnum
func GetListGlossariesFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"catalogId",
		"lifecycleState",
		"timeCreated",
		"uri",
		"workflowStatus",
	}
}

// GetMappingListGlossariesFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGlossariesFieldsEnum(val string) (ListGlossariesFieldsEnum, bool) {
	enum, ok := mappingListGlossariesFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGlossariesSortByEnum Enum with underlying type: string
type ListGlossariesSortByEnum string

// Set of constants representing the allowable values for ListGlossariesSortByEnum
const (
	ListGlossariesSortByTimecreated ListGlossariesSortByEnum = "TIMECREATED"
	ListGlossariesSortByDisplayname ListGlossariesSortByEnum = "DISPLAYNAME"
)

var mappingListGlossariesSortByEnum = map[string]ListGlossariesSortByEnum{
	"TIMECREATED": ListGlossariesSortByTimecreated,
	"DISPLAYNAME": ListGlossariesSortByDisplayname,
}

var mappingListGlossariesSortByEnumLowerCase = map[string]ListGlossariesSortByEnum{
	"timecreated": ListGlossariesSortByTimecreated,
	"displayname": ListGlossariesSortByDisplayname,
}

// GetListGlossariesSortByEnumValues Enumerates the set of values for ListGlossariesSortByEnum
func GetListGlossariesSortByEnumValues() []ListGlossariesSortByEnum {
	values := make([]ListGlossariesSortByEnum, 0)
	for _, v := range mappingListGlossariesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGlossariesSortByEnumStringValues Enumerates the set of values in String for ListGlossariesSortByEnum
func GetListGlossariesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListGlossariesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGlossariesSortByEnum(val string) (ListGlossariesSortByEnum, bool) {
	enum, ok := mappingListGlossariesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGlossariesSortOrderEnum Enum with underlying type: string
type ListGlossariesSortOrderEnum string

// Set of constants representing the allowable values for ListGlossariesSortOrderEnum
const (
	ListGlossariesSortOrderAsc  ListGlossariesSortOrderEnum = "ASC"
	ListGlossariesSortOrderDesc ListGlossariesSortOrderEnum = "DESC"
)

var mappingListGlossariesSortOrderEnum = map[string]ListGlossariesSortOrderEnum{
	"ASC":  ListGlossariesSortOrderAsc,
	"DESC": ListGlossariesSortOrderDesc,
}

var mappingListGlossariesSortOrderEnumLowerCase = map[string]ListGlossariesSortOrderEnum{
	"asc":  ListGlossariesSortOrderAsc,
	"desc": ListGlossariesSortOrderDesc,
}

// GetListGlossariesSortOrderEnumValues Enumerates the set of values for ListGlossariesSortOrderEnum
func GetListGlossariesSortOrderEnumValues() []ListGlossariesSortOrderEnum {
	values := make([]ListGlossariesSortOrderEnum, 0)
	for _, v := range mappingListGlossariesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGlossariesSortOrderEnumStringValues Enumerates the set of values in String for ListGlossariesSortOrderEnum
func GetListGlossariesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGlossariesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGlossariesSortOrderEnum(val string) (ListGlossariesSortOrderEnum, bool) {
	enum, ok := mappingListGlossariesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
