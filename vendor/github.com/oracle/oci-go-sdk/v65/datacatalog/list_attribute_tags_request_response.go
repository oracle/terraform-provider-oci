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

// ListAttributeTagsRequest wrapper for the ListAttributeTags operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListAttributeTags.go.html to see an example of how to use ListAttributeTagsRequest.
type ListAttributeTagsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Unique attribute key.
	AttributeKey *string `mandatory:"true" contributesTo:"path" name:"attributeKey"`

	// Immutable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListAttributeTagsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique key of the related term.
	TermKey *string `mandatory:"false" contributesTo:"query" name:"termKey"`

	// Path of the related term.
	TermPath *string `mandatory:"false" contributesTo:"query" name:"termPath"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// Specifies the fields to return in an entity attribute tag summary response.
	Fields []ListAttributeTagsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListAttributeTagsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAttributeTagsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListAttributeTagsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAttributeTagsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAttributeTagsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAttributeTagsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAttributeTagsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAttributeTagsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAttributeTagsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListAttributeTagsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListAttributeTagsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListAttributeTagsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAttributeTagsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttributeTagsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAttributeTagsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAttributeTagsResponse wrapper for the ListAttributeTags operation
type ListAttributeTagsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AttributeTagCollection instances
	AttributeTagCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAttributeTagsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAttributeTagsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAttributeTagsLifecycleStateEnum Enum with underlying type: string
type ListAttributeTagsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAttributeTagsLifecycleStateEnum
const (
	ListAttributeTagsLifecycleStateCreating ListAttributeTagsLifecycleStateEnum = "CREATING"
	ListAttributeTagsLifecycleStateActive   ListAttributeTagsLifecycleStateEnum = "ACTIVE"
	ListAttributeTagsLifecycleStateInactive ListAttributeTagsLifecycleStateEnum = "INACTIVE"
	ListAttributeTagsLifecycleStateUpdating ListAttributeTagsLifecycleStateEnum = "UPDATING"
	ListAttributeTagsLifecycleStateDeleting ListAttributeTagsLifecycleStateEnum = "DELETING"
	ListAttributeTagsLifecycleStateDeleted  ListAttributeTagsLifecycleStateEnum = "DELETED"
	ListAttributeTagsLifecycleStateFailed   ListAttributeTagsLifecycleStateEnum = "FAILED"
	ListAttributeTagsLifecycleStateMoving   ListAttributeTagsLifecycleStateEnum = "MOVING"
)

var mappingListAttributeTagsLifecycleStateEnum = map[string]ListAttributeTagsLifecycleStateEnum{
	"CREATING": ListAttributeTagsLifecycleStateCreating,
	"ACTIVE":   ListAttributeTagsLifecycleStateActive,
	"INACTIVE": ListAttributeTagsLifecycleStateInactive,
	"UPDATING": ListAttributeTagsLifecycleStateUpdating,
	"DELETING": ListAttributeTagsLifecycleStateDeleting,
	"DELETED":  ListAttributeTagsLifecycleStateDeleted,
	"FAILED":   ListAttributeTagsLifecycleStateFailed,
	"MOVING":   ListAttributeTagsLifecycleStateMoving,
}

var mappingListAttributeTagsLifecycleStateEnumLowerCase = map[string]ListAttributeTagsLifecycleStateEnum{
	"creating": ListAttributeTagsLifecycleStateCreating,
	"active":   ListAttributeTagsLifecycleStateActive,
	"inactive": ListAttributeTagsLifecycleStateInactive,
	"updating": ListAttributeTagsLifecycleStateUpdating,
	"deleting": ListAttributeTagsLifecycleStateDeleting,
	"deleted":  ListAttributeTagsLifecycleStateDeleted,
	"failed":   ListAttributeTagsLifecycleStateFailed,
	"moving":   ListAttributeTagsLifecycleStateMoving,
}

// GetListAttributeTagsLifecycleStateEnumValues Enumerates the set of values for ListAttributeTagsLifecycleStateEnum
func GetListAttributeTagsLifecycleStateEnumValues() []ListAttributeTagsLifecycleStateEnum {
	values := make([]ListAttributeTagsLifecycleStateEnum, 0)
	for _, v := range mappingListAttributeTagsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributeTagsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAttributeTagsLifecycleStateEnum
func GetListAttributeTagsLifecycleStateEnumStringValues() []string {
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

// GetMappingListAttributeTagsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributeTagsLifecycleStateEnum(val string) (ListAttributeTagsLifecycleStateEnum, bool) {
	enum, ok := mappingListAttributeTagsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributeTagsFieldsEnum Enum with underlying type: string
type ListAttributeTagsFieldsEnum string

// Set of constants representing the allowable values for ListAttributeTagsFieldsEnum
const (
	ListAttributeTagsFieldsKey             ListAttributeTagsFieldsEnum = "key"
	ListAttributeTagsFieldsName            ListAttributeTagsFieldsEnum = "name"
	ListAttributeTagsFieldsTermkey         ListAttributeTagsFieldsEnum = "termKey"
	ListAttributeTagsFieldsTermpath        ListAttributeTagsFieldsEnum = "termPath"
	ListAttributeTagsFieldsTermdescription ListAttributeTagsFieldsEnum = "termDescription"
	ListAttributeTagsFieldsLifecyclestate  ListAttributeTagsFieldsEnum = "lifecycleState"
	ListAttributeTagsFieldsTimecreated     ListAttributeTagsFieldsEnum = "timeCreated"
	ListAttributeTagsFieldsUri             ListAttributeTagsFieldsEnum = "uri"
	ListAttributeTagsFieldsGlossarykey     ListAttributeTagsFieldsEnum = "glossaryKey"
	ListAttributeTagsFieldsAttributekey    ListAttributeTagsFieldsEnum = "attributeKey"
)

var mappingListAttributeTagsFieldsEnum = map[string]ListAttributeTagsFieldsEnum{
	"key":             ListAttributeTagsFieldsKey,
	"name":            ListAttributeTagsFieldsName,
	"termKey":         ListAttributeTagsFieldsTermkey,
	"termPath":        ListAttributeTagsFieldsTermpath,
	"termDescription": ListAttributeTagsFieldsTermdescription,
	"lifecycleState":  ListAttributeTagsFieldsLifecyclestate,
	"timeCreated":     ListAttributeTagsFieldsTimecreated,
	"uri":             ListAttributeTagsFieldsUri,
	"glossaryKey":     ListAttributeTagsFieldsGlossarykey,
	"attributeKey":    ListAttributeTagsFieldsAttributekey,
}

var mappingListAttributeTagsFieldsEnumLowerCase = map[string]ListAttributeTagsFieldsEnum{
	"key":             ListAttributeTagsFieldsKey,
	"name":            ListAttributeTagsFieldsName,
	"termkey":         ListAttributeTagsFieldsTermkey,
	"termpath":        ListAttributeTagsFieldsTermpath,
	"termdescription": ListAttributeTagsFieldsTermdescription,
	"lifecyclestate":  ListAttributeTagsFieldsLifecyclestate,
	"timecreated":     ListAttributeTagsFieldsTimecreated,
	"uri":             ListAttributeTagsFieldsUri,
	"glossarykey":     ListAttributeTagsFieldsGlossarykey,
	"attributekey":    ListAttributeTagsFieldsAttributekey,
}

// GetListAttributeTagsFieldsEnumValues Enumerates the set of values for ListAttributeTagsFieldsEnum
func GetListAttributeTagsFieldsEnumValues() []ListAttributeTagsFieldsEnum {
	values := make([]ListAttributeTagsFieldsEnum, 0)
	for _, v := range mappingListAttributeTagsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributeTagsFieldsEnumStringValues Enumerates the set of values in String for ListAttributeTagsFieldsEnum
func GetListAttributeTagsFieldsEnumStringValues() []string {
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
		"attributeKey",
	}
}

// GetMappingListAttributeTagsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributeTagsFieldsEnum(val string) (ListAttributeTagsFieldsEnum, bool) {
	enum, ok := mappingListAttributeTagsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributeTagsSortByEnum Enum with underlying type: string
type ListAttributeTagsSortByEnum string

// Set of constants representing the allowable values for ListAttributeTagsSortByEnum
const (
	ListAttributeTagsSortByTimecreated ListAttributeTagsSortByEnum = "TIMECREATED"
	ListAttributeTagsSortByDisplayname ListAttributeTagsSortByEnum = "DISPLAYNAME"
)

var mappingListAttributeTagsSortByEnum = map[string]ListAttributeTagsSortByEnum{
	"TIMECREATED": ListAttributeTagsSortByTimecreated,
	"DISPLAYNAME": ListAttributeTagsSortByDisplayname,
}

var mappingListAttributeTagsSortByEnumLowerCase = map[string]ListAttributeTagsSortByEnum{
	"timecreated": ListAttributeTagsSortByTimecreated,
	"displayname": ListAttributeTagsSortByDisplayname,
}

// GetListAttributeTagsSortByEnumValues Enumerates the set of values for ListAttributeTagsSortByEnum
func GetListAttributeTagsSortByEnumValues() []ListAttributeTagsSortByEnum {
	values := make([]ListAttributeTagsSortByEnum, 0)
	for _, v := range mappingListAttributeTagsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributeTagsSortByEnumStringValues Enumerates the set of values in String for ListAttributeTagsSortByEnum
func GetListAttributeTagsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAttributeTagsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributeTagsSortByEnum(val string) (ListAttributeTagsSortByEnum, bool) {
	enum, ok := mappingListAttributeTagsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributeTagsSortOrderEnum Enum with underlying type: string
type ListAttributeTagsSortOrderEnum string

// Set of constants representing the allowable values for ListAttributeTagsSortOrderEnum
const (
	ListAttributeTagsSortOrderAsc  ListAttributeTagsSortOrderEnum = "ASC"
	ListAttributeTagsSortOrderDesc ListAttributeTagsSortOrderEnum = "DESC"
)

var mappingListAttributeTagsSortOrderEnum = map[string]ListAttributeTagsSortOrderEnum{
	"ASC":  ListAttributeTagsSortOrderAsc,
	"DESC": ListAttributeTagsSortOrderDesc,
}

var mappingListAttributeTagsSortOrderEnumLowerCase = map[string]ListAttributeTagsSortOrderEnum{
	"asc":  ListAttributeTagsSortOrderAsc,
	"desc": ListAttributeTagsSortOrderDesc,
}

// GetListAttributeTagsSortOrderEnumValues Enumerates the set of values for ListAttributeTagsSortOrderEnum
func GetListAttributeTagsSortOrderEnumValues() []ListAttributeTagsSortOrderEnum {
	values := make([]ListAttributeTagsSortOrderEnum, 0)
	for _, v := range mappingListAttributeTagsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributeTagsSortOrderEnumStringValues Enumerates the set of values in String for ListAttributeTagsSortOrderEnum
func GetListAttributeTagsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAttributeTagsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributeTagsSortOrderEnum(val string) (ListAttributeTagsSortOrderEnum, bool) {
	enum, ok := mappingListAttributeTagsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
