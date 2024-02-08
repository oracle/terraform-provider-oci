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

// ListFolderTagsRequest wrapper for the ListFolderTags operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListFolderTags.go.html to see an example of how to use ListFolderTagsRequest.
type ListFolderTagsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique folder key.
	FolderKey *string `mandatory:"true" contributesTo:"path" name:"folderKey"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListFolderTagsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique key of the related term.
	TermKey *string `mandatory:"false" contributesTo:"query" name:"termKey"`

	// Path of the related term.
	TermPath *string `mandatory:"false" contributesTo:"query" name:"termPath"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Specifies the fields to return in a folder tag summary response.
	Fields []ListFolderTagsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListFolderTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListFolderTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListFolderTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFolderTagsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFolderTagsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFolderTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFolderTagsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFolderTagsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFolderTagsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListFolderTagsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListFolderTagsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListFolderTagsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFolderTagsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFolderTagsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFolderTagsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFolderTagsResponse wrapper for the ListFolderTags operation
type ListFolderTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FolderTagCollection instances
	FolderTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFolderTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFolderTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFolderTagsLifecycleStateEnum Enum with underlying type: string
type ListFolderTagsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFolderTagsLifecycleStateEnum
const (
	ListFolderTagsLifecycleStateCreating ListFolderTagsLifecycleStateEnum = "CREATING"
	ListFolderTagsLifecycleStateActive   ListFolderTagsLifecycleStateEnum = "ACTIVE"
	ListFolderTagsLifecycleStateInactive ListFolderTagsLifecycleStateEnum = "INACTIVE"
	ListFolderTagsLifecycleStateUpdating ListFolderTagsLifecycleStateEnum = "UPDATING"
	ListFolderTagsLifecycleStateDeleting ListFolderTagsLifecycleStateEnum = "DELETING"
	ListFolderTagsLifecycleStateDeleted  ListFolderTagsLifecycleStateEnum = "DELETED"
	ListFolderTagsLifecycleStateFailed   ListFolderTagsLifecycleStateEnum = "FAILED"
	ListFolderTagsLifecycleStateMoving   ListFolderTagsLifecycleStateEnum = "MOVING"
)

var mappingListFolderTagsLifecycleStateEnum = map[string]ListFolderTagsLifecycleStateEnum{
	"CREATING": ListFolderTagsLifecycleStateCreating,
	"ACTIVE":   ListFolderTagsLifecycleStateActive,
	"INACTIVE": ListFolderTagsLifecycleStateInactive,
	"UPDATING": ListFolderTagsLifecycleStateUpdating,
	"DELETING": ListFolderTagsLifecycleStateDeleting,
	"DELETED":  ListFolderTagsLifecycleStateDeleted,
	"FAILED":   ListFolderTagsLifecycleStateFailed,
	"MOVING":   ListFolderTagsLifecycleStateMoving,
}

var mappingListFolderTagsLifecycleStateEnumLowerCase = map[string]ListFolderTagsLifecycleStateEnum{
	"creating": ListFolderTagsLifecycleStateCreating,
	"active":   ListFolderTagsLifecycleStateActive,
	"inactive": ListFolderTagsLifecycleStateInactive,
	"updating": ListFolderTagsLifecycleStateUpdating,
	"deleting": ListFolderTagsLifecycleStateDeleting,
	"deleted":  ListFolderTagsLifecycleStateDeleted,
	"failed":   ListFolderTagsLifecycleStateFailed,
	"moving":   ListFolderTagsLifecycleStateMoving,
}

// GetListFolderTagsLifecycleStateEnumValues Enumerates the set of values for ListFolderTagsLifecycleStateEnum
func GetListFolderTagsLifecycleStateEnumValues() []ListFolderTagsLifecycleStateEnum {
	values := make([]ListFolderTagsLifecycleStateEnum, 0)
	for _, v := range mappingListFolderTagsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFolderTagsLifecycleStateEnumStringValues Enumerates the set of values in String for ListFolderTagsLifecycleStateEnum
func GetListFolderTagsLifecycleStateEnumStringValues() []string {
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

// GetMappingListFolderTagsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFolderTagsLifecycleStateEnum(val string) (ListFolderTagsLifecycleStateEnum, bool) {
	enum, ok := mappingListFolderTagsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFolderTagsFieldsEnum Enum with underlying type: string
type ListFolderTagsFieldsEnum string

// Set of constants representing the allowable values for ListFolderTagsFieldsEnum
const (
	ListFolderTagsFieldsKey             ListFolderTagsFieldsEnum = "key"
	ListFolderTagsFieldsName            ListFolderTagsFieldsEnum = "name"
	ListFolderTagsFieldsTermkey         ListFolderTagsFieldsEnum = "termKey"
	ListFolderTagsFieldsTermpath        ListFolderTagsFieldsEnum = "termPath"
	ListFolderTagsFieldsTermdescription ListFolderTagsFieldsEnum = "termDescription"
	ListFolderTagsFieldsLifecyclestate  ListFolderTagsFieldsEnum = "lifecycleState"
	ListFolderTagsFieldsTimecreated     ListFolderTagsFieldsEnum = "timeCreated"
	ListFolderTagsFieldsUri             ListFolderTagsFieldsEnum = "uri"
	ListFolderTagsFieldsGlossarykey     ListFolderTagsFieldsEnum = "glossaryKey"
	ListFolderTagsFieldsFolderkey       ListFolderTagsFieldsEnum = "folderKey"
)

var mappingListFolderTagsFieldsEnum = map[string]ListFolderTagsFieldsEnum{
	"key":             ListFolderTagsFieldsKey,
	"name":            ListFolderTagsFieldsName,
	"termKey":         ListFolderTagsFieldsTermkey,
	"termPath":        ListFolderTagsFieldsTermpath,
	"termDescription": ListFolderTagsFieldsTermdescription,
	"lifecycleState":  ListFolderTagsFieldsLifecyclestate,
	"timeCreated":     ListFolderTagsFieldsTimecreated,
	"uri":             ListFolderTagsFieldsUri,
	"glossaryKey":     ListFolderTagsFieldsGlossarykey,
	"folderKey":       ListFolderTagsFieldsFolderkey,
}

var mappingListFolderTagsFieldsEnumLowerCase = map[string]ListFolderTagsFieldsEnum{
	"key":             ListFolderTagsFieldsKey,
	"name":            ListFolderTagsFieldsName,
	"termkey":         ListFolderTagsFieldsTermkey,
	"termpath":        ListFolderTagsFieldsTermpath,
	"termdescription": ListFolderTagsFieldsTermdescription,
	"lifecyclestate":  ListFolderTagsFieldsLifecyclestate,
	"timecreated":     ListFolderTagsFieldsTimecreated,
	"uri":             ListFolderTagsFieldsUri,
	"glossarykey":     ListFolderTagsFieldsGlossarykey,
	"folderkey":       ListFolderTagsFieldsFolderkey,
}

// GetListFolderTagsFieldsEnumValues Enumerates the set of values for ListFolderTagsFieldsEnum
func GetListFolderTagsFieldsEnumValues() []ListFolderTagsFieldsEnum {
	values := make([]ListFolderTagsFieldsEnum, 0)
	for _, v := range mappingListFolderTagsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListFolderTagsFieldsEnumStringValues Enumerates the set of values in String for ListFolderTagsFieldsEnum
func GetListFolderTagsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"name",
		"termKey",
		"termPath",
		"termDescription",
		"lifecycleState",
		"timeCreated",
		"uri",
		"glossaryKey",
		"folderKey",
	}
}

// GetMappingListFolderTagsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFolderTagsFieldsEnum(val string) (ListFolderTagsFieldsEnum, bool) {
	enum, ok := mappingListFolderTagsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFolderTagsSortByEnum Enum with underlying type: string
type ListFolderTagsSortByEnum string

// Set of constants representing the allowable values for ListFolderTagsSortByEnum
const (
	ListFolderTagsSortByTimecreated ListFolderTagsSortByEnum = "TIMECREATED"
	ListFolderTagsSortByDisplayname ListFolderTagsSortByEnum = "DISPLAYNAME"
)

var mappingListFolderTagsSortByEnum = map[string]ListFolderTagsSortByEnum{
	"TIMECREATED": ListFolderTagsSortByTimecreated,
	"DISPLAYNAME": ListFolderTagsSortByDisplayname,
}

var mappingListFolderTagsSortByEnumLowerCase = map[string]ListFolderTagsSortByEnum{
	"timecreated": ListFolderTagsSortByTimecreated,
	"displayname": ListFolderTagsSortByDisplayname,
}

// GetListFolderTagsSortByEnumValues Enumerates the set of values for ListFolderTagsSortByEnum
func GetListFolderTagsSortByEnumValues() []ListFolderTagsSortByEnum {
	values := make([]ListFolderTagsSortByEnum, 0)
	for _, v := range mappingListFolderTagsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFolderTagsSortByEnumStringValues Enumerates the set of values in String for ListFolderTagsSortByEnum
func GetListFolderTagsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListFolderTagsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFolderTagsSortByEnum(val string) (ListFolderTagsSortByEnum, bool) {
	enum, ok := mappingListFolderTagsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFolderTagsSortOrderEnum Enum with underlying type: string
type ListFolderTagsSortOrderEnum string

// Set of constants representing the allowable values for ListFolderTagsSortOrderEnum
const (
	ListFolderTagsSortOrderAsc  ListFolderTagsSortOrderEnum = "ASC"
	ListFolderTagsSortOrderDesc ListFolderTagsSortOrderEnum = "DESC"
)

var mappingListFolderTagsSortOrderEnum = map[string]ListFolderTagsSortOrderEnum{
	"ASC":  ListFolderTagsSortOrderAsc,
	"DESC": ListFolderTagsSortOrderDesc,
}

var mappingListFolderTagsSortOrderEnumLowerCase = map[string]ListFolderTagsSortOrderEnum{
	"asc":  ListFolderTagsSortOrderAsc,
	"desc": ListFolderTagsSortOrderDesc,
}

// GetListFolderTagsSortOrderEnumValues Enumerates the set of values for ListFolderTagsSortOrderEnum
func GetListFolderTagsSortOrderEnumValues() []ListFolderTagsSortOrderEnum {
	values := make([]ListFolderTagsSortOrderEnum, 0)
	for _, v := range mappingListFolderTagsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFolderTagsSortOrderEnumStringValues Enumerates the set of values in String for ListFolderTagsSortOrderEnum
func GetListFolderTagsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFolderTagsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFolderTagsSortOrderEnum(val string) (ListFolderTagsSortOrderEnum, bool) {
	enum, ok := mappingListFolderTagsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
