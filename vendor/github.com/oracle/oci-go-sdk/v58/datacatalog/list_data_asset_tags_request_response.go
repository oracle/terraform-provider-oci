// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDataAssetTagsRequest wrapper for the ListDataAssetTags operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListDataAssetTags.go.html to see an example of how to use ListDataAssetTagsRequest.
type ListDataAssetTagsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListDataAssetTagsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique key of the related term.
	TermKey *string `mandatory:"false" contributesTo:"query" name:"termKey"`

	// Path of the related term.
	TermPath *string `mandatory:"false" contributesTo:"query" name:"termPath"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Specifies the fields to return in a data asset tag summary response.
	Fields []ListDataAssetTagsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListDataAssetTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataAssetTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListDataAssetTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataAssetTagsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataAssetTagsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataAssetTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataAssetTagsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataAssetTagsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDataAssetTagsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListDataAssetTagsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListDataAssetTagsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListDataAssetTagsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataAssetTagsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataAssetTagsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataAssetTagsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataAssetTagsResponse wrapper for the ListDataAssetTags operation
type ListDataAssetTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataAssetTagCollection instances
	DataAssetTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataAssetTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataAssetTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataAssetTagsLifecycleStateEnum Enum with underlying type: string
type ListDataAssetTagsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataAssetTagsLifecycleStateEnum
const (
	ListDataAssetTagsLifecycleStateCreating ListDataAssetTagsLifecycleStateEnum = "CREATING"
	ListDataAssetTagsLifecycleStateActive   ListDataAssetTagsLifecycleStateEnum = "ACTIVE"
	ListDataAssetTagsLifecycleStateInactive ListDataAssetTagsLifecycleStateEnum = "INACTIVE"
	ListDataAssetTagsLifecycleStateUpdating ListDataAssetTagsLifecycleStateEnum = "UPDATING"
	ListDataAssetTagsLifecycleStateDeleting ListDataAssetTagsLifecycleStateEnum = "DELETING"
	ListDataAssetTagsLifecycleStateDeleted  ListDataAssetTagsLifecycleStateEnum = "DELETED"
	ListDataAssetTagsLifecycleStateFailed   ListDataAssetTagsLifecycleStateEnum = "FAILED"
	ListDataAssetTagsLifecycleStateMoving   ListDataAssetTagsLifecycleStateEnum = "MOVING"
)

var mappingListDataAssetTagsLifecycleStateEnum = map[string]ListDataAssetTagsLifecycleStateEnum{
	"CREATING": ListDataAssetTagsLifecycleStateCreating,
	"ACTIVE":   ListDataAssetTagsLifecycleStateActive,
	"INACTIVE": ListDataAssetTagsLifecycleStateInactive,
	"UPDATING": ListDataAssetTagsLifecycleStateUpdating,
	"DELETING": ListDataAssetTagsLifecycleStateDeleting,
	"DELETED":  ListDataAssetTagsLifecycleStateDeleted,
	"FAILED":   ListDataAssetTagsLifecycleStateFailed,
	"MOVING":   ListDataAssetTagsLifecycleStateMoving,
}

// GetListDataAssetTagsLifecycleStateEnumValues Enumerates the set of values for ListDataAssetTagsLifecycleStateEnum
func GetListDataAssetTagsLifecycleStateEnumValues() []ListDataAssetTagsLifecycleStateEnum {
	values := make([]ListDataAssetTagsLifecycleStateEnum, 0)
	for _, v := range mappingListDataAssetTagsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataAssetTagsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDataAssetTagsLifecycleStateEnum
func GetListDataAssetTagsLifecycleStateEnumStringValues() []string {
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

// GetMappingListDataAssetTagsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataAssetTagsLifecycleStateEnum(val string) (ListDataAssetTagsLifecycleStateEnum, bool) {
	mappingListDataAssetTagsLifecycleStateEnumIgnoreCase := make(map[string]ListDataAssetTagsLifecycleStateEnum)
	for k, v := range mappingListDataAssetTagsLifecycleStateEnum {
		mappingListDataAssetTagsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataAssetTagsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataAssetTagsFieldsEnum Enum with underlying type: string
type ListDataAssetTagsFieldsEnum string

// Set of constants representing the allowable values for ListDataAssetTagsFieldsEnum
const (
	ListDataAssetTagsFieldsKey             ListDataAssetTagsFieldsEnum = "key"
	ListDataAssetTagsFieldsName            ListDataAssetTagsFieldsEnum = "name"
	ListDataAssetTagsFieldsTermkey         ListDataAssetTagsFieldsEnum = "termKey"
	ListDataAssetTagsFieldsTermpath        ListDataAssetTagsFieldsEnum = "termPath"
	ListDataAssetTagsFieldsTermdescription ListDataAssetTagsFieldsEnum = "termDescription"
	ListDataAssetTagsFieldsLifecyclestate  ListDataAssetTagsFieldsEnum = "lifecycleState"
	ListDataAssetTagsFieldsTimecreated     ListDataAssetTagsFieldsEnum = "timeCreated"
	ListDataAssetTagsFieldsUri             ListDataAssetTagsFieldsEnum = "uri"
	ListDataAssetTagsFieldsGlossarykey     ListDataAssetTagsFieldsEnum = "glossaryKey"
	ListDataAssetTagsFieldsDataassetkey    ListDataAssetTagsFieldsEnum = "dataAssetKey"
)

var mappingListDataAssetTagsFieldsEnum = map[string]ListDataAssetTagsFieldsEnum{
	"key":             ListDataAssetTagsFieldsKey,
	"name":            ListDataAssetTagsFieldsName,
	"termKey":         ListDataAssetTagsFieldsTermkey,
	"termPath":        ListDataAssetTagsFieldsTermpath,
	"termDescription": ListDataAssetTagsFieldsTermdescription,
	"lifecycleState":  ListDataAssetTagsFieldsLifecyclestate,
	"timeCreated":     ListDataAssetTagsFieldsTimecreated,
	"uri":             ListDataAssetTagsFieldsUri,
	"glossaryKey":     ListDataAssetTagsFieldsGlossarykey,
	"dataAssetKey":    ListDataAssetTagsFieldsDataassetkey,
}

// GetListDataAssetTagsFieldsEnumValues Enumerates the set of values for ListDataAssetTagsFieldsEnum
func GetListDataAssetTagsFieldsEnumValues() []ListDataAssetTagsFieldsEnum {
	values := make([]ListDataAssetTagsFieldsEnum, 0)
	for _, v := range mappingListDataAssetTagsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataAssetTagsFieldsEnumStringValues Enumerates the set of values in String for ListDataAssetTagsFieldsEnum
func GetListDataAssetTagsFieldsEnumStringValues() []string {
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
		"dataAssetKey",
	}
}

// GetMappingListDataAssetTagsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataAssetTagsFieldsEnum(val string) (ListDataAssetTagsFieldsEnum, bool) {
	mappingListDataAssetTagsFieldsEnumIgnoreCase := make(map[string]ListDataAssetTagsFieldsEnum)
	for k, v := range mappingListDataAssetTagsFieldsEnum {
		mappingListDataAssetTagsFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataAssetTagsFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataAssetTagsSortByEnum Enum with underlying type: string
type ListDataAssetTagsSortByEnum string

// Set of constants representing the allowable values for ListDataAssetTagsSortByEnum
const (
	ListDataAssetTagsSortByTimecreated ListDataAssetTagsSortByEnum = "TIMECREATED"
	ListDataAssetTagsSortByDisplayname ListDataAssetTagsSortByEnum = "DISPLAYNAME"
)

var mappingListDataAssetTagsSortByEnum = map[string]ListDataAssetTagsSortByEnum{
	"TIMECREATED": ListDataAssetTagsSortByTimecreated,
	"DISPLAYNAME": ListDataAssetTagsSortByDisplayname,
}

// GetListDataAssetTagsSortByEnumValues Enumerates the set of values for ListDataAssetTagsSortByEnum
func GetListDataAssetTagsSortByEnumValues() []ListDataAssetTagsSortByEnum {
	values := make([]ListDataAssetTagsSortByEnum, 0)
	for _, v := range mappingListDataAssetTagsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataAssetTagsSortByEnumStringValues Enumerates the set of values in String for ListDataAssetTagsSortByEnum
func GetListDataAssetTagsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDataAssetTagsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataAssetTagsSortByEnum(val string) (ListDataAssetTagsSortByEnum, bool) {
	mappingListDataAssetTagsSortByEnumIgnoreCase := make(map[string]ListDataAssetTagsSortByEnum)
	for k, v := range mappingListDataAssetTagsSortByEnum {
		mappingListDataAssetTagsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataAssetTagsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataAssetTagsSortOrderEnum Enum with underlying type: string
type ListDataAssetTagsSortOrderEnum string

// Set of constants representing the allowable values for ListDataAssetTagsSortOrderEnum
const (
	ListDataAssetTagsSortOrderAsc  ListDataAssetTagsSortOrderEnum = "ASC"
	ListDataAssetTagsSortOrderDesc ListDataAssetTagsSortOrderEnum = "DESC"
)

var mappingListDataAssetTagsSortOrderEnum = map[string]ListDataAssetTagsSortOrderEnum{
	"ASC":  ListDataAssetTagsSortOrderAsc,
	"DESC": ListDataAssetTagsSortOrderDesc,
}

// GetListDataAssetTagsSortOrderEnumValues Enumerates the set of values for ListDataAssetTagsSortOrderEnum
func GetListDataAssetTagsSortOrderEnumValues() []ListDataAssetTagsSortOrderEnum {
	values := make([]ListDataAssetTagsSortOrderEnum, 0)
	for _, v := range mappingListDataAssetTagsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataAssetTagsSortOrderEnumStringValues Enumerates the set of values in String for ListDataAssetTagsSortOrderEnum
func GetListDataAssetTagsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataAssetTagsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataAssetTagsSortOrderEnum(val string) (ListDataAssetTagsSortOrderEnum, bool) {
	mappingListDataAssetTagsSortOrderEnumIgnoreCase := make(map[string]ListDataAssetTagsSortOrderEnum)
	for k, v := range mappingListDataAssetTagsSortOrderEnum {
		mappingListDataAssetTagsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDataAssetTagsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
