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

// ListEntityTagsRequest wrapper for the ListEntityTags operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListEntityTags.go.html to see an example of how to use ListEntityTagsRequest.
type ListEntityTagsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListEntityTagsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique key of the related term.
	TermKey *string `mandatory:"false" contributesTo:"query" name:"termKey"`

	// Path of the related term.
	TermPath *string `mandatory:"false" contributesTo:"query" name:"termPath"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Specifies the fields to return in an entity tag summary response.
	Fields []ListEntityTagsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListEntityTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListEntityTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListEntityTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEntityTagsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEntityTagsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEntityTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEntityTagsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEntityTagsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListEntityTagsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListEntityTagsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListEntityTagsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListEntityTagsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEntityTagsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntityTagsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEntityTagsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEntityTagsResponse wrapper for the ListEntityTags operation
type ListEntityTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EntityTagCollection instances
	EntityTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEntityTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEntityTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEntityTagsLifecycleStateEnum Enum with underlying type: string
type ListEntityTagsLifecycleStateEnum string

// Set of constants representing the allowable values for ListEntityTagsLifecycleStateEnum
const (
	ListEntityTagsLifecycleStateCreating ListEntityTagsLifecycleStateEnum = "CREATING"
	ListEntityTagsLifecycleStateActive   ListEntityTagsLifecycleStateEnum = "ACTIVE"
	ListEntityTagsLifecycleStateInactive ListEntityTagsLifecycleStateEnum = "INACTIVE"
	ListEntityTagsLifecycleStateUpdating ListEntityTagsLifecycleStateEnum = "UPDATING"
	ListEntityTagsLifecycleStateDeleting ListEntityTagsLifecycleStateEnum = "DELETING"
	ListEntityTagsLifecycleStateDeleted  ListEntityTagsLifecycleStateEnum = "DELETED"
	ListEntityTagsLifecycleStateFailed   ListEntityTagsLifecycleStateEnum = "FAILED"
	ListEntityTagsLifecycleStateMoving   ListEntityTagsLifecycleStateEnum = "MOVING"
)

var mappingListEntityTagsLifecycleStateEnum = map[string]ListEntityTagsLifecycleStateEnum{
	"CREATING": ListEntityTagsLifecycleStateCreating,
	"ACTIVE":   ListEntityTagsLifecycleStateActive,
	"INACTIVE": ListEntityTagsLifecycleStateInactive,
	"UPDATING": ListEntityTagsLifecycleStateUpdating,
	"DELETING": ListEntityTagsLifecycleStateDeleting,
	"DELETED":  ListEntityTagsLifecycleStateDeleted,
	"FAILED":   ListEntityTagsLifecycleStateFailed,
	"MOVING":   ListEntityTagsLifecycleStateMoving,
}

// GetListEntityTagsLifecycleStateEnumValues Enumerates the set of values for ListEntityTagsLifecycleStateEnum
func GetListEntityTagsLifecycleStateEnumValues() []ListEntityTagsLifecycleStateEnum {
	values := make([]ListEntityTagsLifecycleStateEnum, 0)
	for _, v := range mappingListEntityTagsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntityTagsLifecycleStateEnumStringValues Enumerates the set of values in String for ListEntityTagsLifecycleStateEnum
func GetListEntityTagsLifecycleStateEnumStringValues() []string {
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

// GetMappingListEntityTagsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntityTagsLifecycleStateEnum(val string) (ListEntityTagsLifecycleStateEnum, bool) {
	mappingListEntityTagsLifecycleStateEnumIgnoreCase := make(map[string]ListEntityTagsLifecycleStateEnum)
	for k, v := range mappingListEntityTagsLifecycleStateEnum {
		mappingListEntityTagsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntityTagsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntityTagsFieldsEnum Enum with underlying type: string
type ListEntityTagsFieldsEnum string

// Set of constants representing the allowable values for ListEntityTagsFieldsEnum
const (
	ListEntityTagsFieldsKey             ListEntityTagsFieldsEnum = "key"
	ListEntityTagsFieldsName            ListEntityTagsFieldsEnum = "name"
	ListEntityTagsFieldsTermkey         ListEntityTagsFieldsEnum = "termKey"
	ListEntityTagsFieldsTermpath        ListEntityTagsFieldsEnum = "termPath"
	ListEntityTagsFieldsTermdescription ListEntityTagsFieldsEnum = "termDescription"
	ListEntityTagsFieldsLifecyclestate  ListEntityTagsFieldsEnum = "lifecycleState"
	ListEntityTagsFieldsTimecreated     ListEntityTagsFieldsEnum = "timeCreated"
	ListEntityTagsFieldsUri             ListEntityTagsFieldsEnum = "uri"
	ListEntityTagsFieldsGlossarykey     ListEntityTagsFieldsEnum = "glossaryKey"
	ListEntityTagsFieldsEntitykey       ListEntityTagsFieldsEnum = "entityKey"
)

var mappingListEntityTagsFieldsEnum = map[string]ListEntityTagsFieldsEnum{
	"key":             ListEntityTagsFieldsKey,
	"name":            ListEntityTagsFieldsName,
	"termKey":         ListEntityTagsFieldsTermkey,
	"termPath":        ListEntityTagsFieldsTermpath,
	"termDescription": ListEntityTagsFieldsTermdescription,
	"lifecycleState":  ListEntityTagsFieldsLifecyclestate,
	"timeCreated":     ListEntityTagsFieldsTimecreated,
	"uri":             ListEntityTagsFieldsUri,
	"glossaryKey":     ListEntityTagsFieldsGlossarykey,
	"entityKey":       ListEntityTagsFieldsEntitykey,
}

// GetListEntityTagsFieldsEnumValues Enumerates the set of values for ListEntityTagsFieldsEnum
func GetListEntityTagsFieldsEnumValues() []ListEntityTagsFieldsEnum {
	values := make([]ListEntityTagsFieldsEnum, 0)
	for _, v := range mappingListEntityTagsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntityTagsFieldsEnumStringValues Enumerates the set of values in String for ListEntityTagsFieldsEnum
func GetListEntityTagsFieldsEnumStringValues() []string {
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
		"entityKey",
	}
}

// GetMappingListEntityTagsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntityTagsFieldsEnum(val string) (ListEntityTagsFieldsEnum, bool) {
	mappingListEntityTagsFieldsEnumIgnoreCase := make(map[string]ListEntityTagsFieldsEnum)
	for k, v := range mappingListEntityTagsFieldsEnum {
		mappingListEntityTagsFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntityTagsFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntityTagsSortByEnum Enum with underlying type: string
type ListEntityTagsSortByEnum string

// Set of constants representing the allowable values for ListEntityTagsSortByEnum
const (
	ListEntityTagsSortByTimecreated ListEntityTagsSortByEnum = "TIMECREATED"
	ListEntityTagsSortByDisplayname ListEntityTagsSortByEnum = "DISPLAYNAME"
)

var mappingListEntityTagsSortByEnum = map[string]ListEntityTagsSortByEnum{
	"TIMECREATED": ListEntityTagsSortByTimecreated,
	"DISPLAYNAME": ListEntityTagsSortByDisplayname,
}

// GetListEntityTagsSortByEnumValues Enumerates the set of values for ListEntityTagsSortByEnum
func GetListEntityTagsSortByEnumValues() []ListEntityTagsSortByEnum {
	values := make([]ListEntityTagsSortByEnum, 0)
	for _, v := range mappingListEntityTagsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntityTagsSortByEnumStringValues Enumerates the set of values in String for ListEntityTagsSortByEnum
func GetListEntityTagsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListEntityTagsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntityTagsSortByEnum(val string) (ListEntityTagsSortByEnum, bool) {
	mappingListEntityTagsSortByEnumIgnoreCase := make(map[string]ListEntityTagsSortByEnum)
	for k, v := range mappingListEntityTagsSortByEnum {
		mappingListEntityTagsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntityTagsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntityTagsSortOrderEnum Enum with underlying type: string
type ListEntityTagsSortOrderEnum string

// Set of constants representing the allowable values for ListEntityTagsSortOrderEnum
const (
	ListEntityTagsSortOrderAsc  ListEntityTagsSortOrderEnum = "ASC"
	ListEntityTagsSortOrderDesc ListEntityTagsSortOrderEnum = "DESC"
)

var mappingListEntityTagsSortOrderEnum = map[string]ListEntityTagsSortOrderEnum{
	"ASC":  ListEntityTagsSortOrderAsc,
	"DESC": ListEntityTagsSortOrderDesc,
}

// GetListEntityTagsSortOrderEnumValues Enumerates the set of values for ListEntityTagsSortOrderEnum
func GetListEntityTagsSortOrderEnumValues() []ListEntityTagsSortOrderEnum {
	values := make([]ListEntityTagsSortOrderEnum, 0)
	for _, v := range mappingListEntityTagsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntityTagsSortOrderEnumStringValues Enumerates the set of values in String for ListEntityTagsSortOrderEnum
func GetListEntityTagsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEntityTagsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntityTagsSortOrderEnum(val string) (ListEntityTagsSortOrderEnum, bool) {
	mappingListEntityTagsSortOrderEnumIgnoreCase := make(map[string]ListEntityTagsSortOrderEnum)
	for k, v := range mappingListEntityTagsSortOrderEnum {
		mappingListEntityTagsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntityTagsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
