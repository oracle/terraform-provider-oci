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

// ListConnectionsRequest wrapper for the ListConnections operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListConnections.go.html to see an example of how to use ListConnectionsRequest.
type ListConnectionsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListConnectionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Unique external identifier of this resource in the external source system.
	ExternalKey *string `mandatory:"false" contributesTo:"query" name:"externalKey"`

	// Time that the resource's status was last updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStatusUpdated"`

	// Indicates whether this connection is the default connection.
	IsDefault *bool `mandatory:"false" contributesTo:"query" name:"isDefault"`

	// Specifies the fields to return in a connection summary response.
	Fields []ListConnectionsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListConnectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConnectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListConnectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConnectionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConnectionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListConnectionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListConnectionsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListConnectionsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListConnectionsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListConnectionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConnectionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConnectionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConnectionsResponse wrapper for the ListConnections operation
type ListConnectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectionCollection instances
	ConnectionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConnectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectionsLifecycleStateEnum Enum with underlying type: string
type ListConnectionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListConnectionsLifecycleStateEnum
const (
	ListConnectionsLifecycleStateCreating ListConnectionsLifecycleStateEnum = "CREATING"
	ListConnectionsLifecycleStateActive   ListConnectionsLifecycleStateEnum = "ACTIVE"
	ListConnectionsLifecycleStateInactive ListConnectionsLifecycleStateEnum = "INACTIVE"
	ListConnectionsLifecycleStateUpdating ListConnectionsLifecycleStateEnum = "UPDATING"
	ListConnectionsLifecycleStateDeleting ListConnectionsLifecycleStateEnum = "DELETING"
	ListConnectionsLifecycleStateDeleted  ListConnectionsLifecycleStateEnum = "DELETED"
	ListConnectionsLifecycleStateFailed   ListConnectionsLifecycleStateEnum = "FAILED"
	ListConnectionsLifecycleStateMoving   ListConnectionsLifecycleStateEnum = "MOVING"
)

var mappingListConnectionsLifecycleStateEnum = map[string]ListConnectionsLifecycleStateEnum{
	"CREATING": ListConnectionsLifecycleStateCreating,
	"ACTIVE":   ListConnectionsLifecycleStateActive,
	"INACTIVE": ListConnectionsLifecycleStateInactive,
	"UPDATING": ListConnectionsLifecycleStateUpdating,
	"DELETING": ListConnectionsLifecycleStateDeleting,
	"DELETED":  ListConnectionsLifecycleStateDeleted,
	"FAILED":   ListConnectionsLifecycleStateFailed,
	"MOVING":   ListConnectionsLifecycleStateMoving,
}

var mappingListConnectionsLifecycleStateEnumLowerCase = map[string]ListConnectionsLifecycleStateEnum{
	"creating": ListConnectionsLifecycleStateCreating,
	"active":   ListConnectionsLifecycleStateActive,
	"inactive": ListConnectionsLifecycleStateInactive,
	"updating": ListConnectionsLifecycleStateUpdating,
	"deleting": ListConnectionsLifecycleStateDeleting,
	"deleted":  ListConnectionsLifecycleStateDeleted,
	"failed":   ListConnectionsLifecycleStateFailed,
	"moving":   ListConnectionsLifecycleStateMoving,
}

// GetListConnectionsLifecycleStateEnumValues Enumerates the set of values for ListConnectionsLifecycleStateEnum
func GetListConnectionsLifecycleStateEnumValues() []ListConnectionsLifecycleStateEnum {
	values := make([]ListConnectionsLifecycleStateEnum, 0)
	for _, v := range mappingListConnectionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListConnectionsLifecycleStateEnum
func GetListConnectionsLifecycleStateEnumStringValues() []string {
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

// GetMappingListConnectionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsLifecycleStateEnum(val string) (ListConnectionsLifecycleStateEnum, bool) {
	enum, ok := mappingListConnectionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsFieldsEnum Enum with underlying type: string
type ListConnectionsFieldsEnum string

// Set of constants representing the allowable values for ListConnectionsFieldsEnum
const (
	ListConnectionsFieldsKey            ListConnectionsFieldsEnum = "key"
	ListConnectionsFieldsDisplayname    ListConnectionsFieldsEnum = "displayName"
	ListConnectionsFieldsDescription    ListConnectionsFieldsEnum = "description"
	ListConnectionsFieldsDataassetkey   ListConnectionsFieldsEnum = "dataAssetKey"
	ListConnectionsFieldsTypekey        ListConnectionsFieldsEnum = "typeKey"
	ListConnectionsFieldsTimecreated    ListConnectionsFieldsEnum = "timeCreated"
	ListConnectionsFieldsExternalkey    ListConnectionsFieldsEnum = "externalKey"
	ListConnectionsFieldsLifecyclestate ListConnectionsFieldsEnum = "lifecycleState"
	ListConnectionsFieldsIsdefault      ListConnectionsFieldsEnum = "isDefault"
	ListConnectionsFieldsUri            ListConnectionsFieldsEnum = "uri"
)

var mappingListConnectionsFieldsEnum = map[string]ListConnectionsFieldsEnum{
	"key":            ListConnectionsFieldsKey,
	"displayName":    ListConnectionsFieldsDisplayname,
	"description":    ListConnectionsFieldsDescription,
	"dataAssetKey":   ListConnectionsFieldsDataassetkey,
	"typeKey":        ListConnectionsFieldsTypekey,
	"timeCreated":    ListConnectionsFieldsTimecreated,
	"externalKey":    ListConnectionsFieldsExternalkey,
	"lifecycleState": ListConnectionsFieldsLifecyclestate,
	"isDefault":      ListConnectionsFieldsIsdefault,
	"uri":            ListConnectionsFieldsUri,
}

var mappingListConnectionsFieldsEnumLowerCase = map[string]ListConnectionsFieldsEnum{
	"key":            ListConnectionsFieldsKey,
	"displayname":    ListConnectionsFieldsDisplayname,
	"description":    ListConnectionsFieldsDescription,
	"dataassetkey":   ListConnectionsFieldsDataassetkey,
	"typekey":        ListConnectionsFieldsTypekey,
	"timecreated":    ListConnectionsFieldsTimecreated,
	"externalkey":    ListConnectionsFieldsExternalkey,
	"lifecyclestate": ListConnectionsFieldsLifecyclestate,
	"isdefault":      ListConnectionsFieldsIsdefault,
	"uri":            ListConnectionsFieldsUri,
}

// GetListConnectionsFieldsEnumValues Enumerates the set of values for ListConnectionsFieldsEnum
func GetListConnectionsFieldsEnumValues() []ListConnectionsFieldsEnum {
	values := make([]ListConnectionsFieldsEnum, 0)
	for _, v := range mappingListConnectionsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsFieldsEnumStringValues Enumerates the set of values in String for ListConnectionsFieldsEnum
func GetListConnectionsFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"dataAssetKey",
		"typeKey",
		"timeCreated",
		"externalKey",
		"lifecycleState",
		"isDefault",
		"uri",
	}
}

// GetMappingListConnectionsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsFieldsEnum(val string) (ListConnectionsFieldsEnum, bool) {
	enum, ok := mappingListConnectionsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsSortByEnum Enum with underlying type: string
type ListConnectionsSortByEnum string

// Set of constants representing the allowable values for ListConnectionsSortByEnum
const (
	ListConnectionsSortByTimecreated ListConnectionsSortByEnum = "TIMECREATED"
	ListConnectionsSortByDisplayname ListConnectionsSortByEnum = "DISPLAYNAME"
)

var mappingListConnectionsSortByEnum = map[string]ListConnectionsSortByEnum{
	"TIMECREATED": ListConnectionsSortByTimecreated,
	"DISPLAYNAME": ListConnectionsSortByDisplayname,
}

var mappingListConnectionsSortByEnumLowerCase = map[string]ListConnectionsSortByEnum{
	"timecreated": ListConnectionsSortByTimecreated,
	"displayname": ListConnectionsSortByDisplayname,
}

// GetListConnectionsSortByEnumValues Enumerates the set of values for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumValues() []ListConnectionsSortByEnum {
	values := make([]ListConnectionsSortByEnum, 0)
	for _, v := range mappingListConnectionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsSortByEnumStringValues Enumerates the set of values in String for ListConnectionsSortByEnum
func GetListConnectionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListConnectionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsSortByEnum(val string) (ListConnectionsSortByEnum, bool) {
	enum, ok := mappingListConnectionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionsSortOrderEnum Enum with underlying type: string
type ListConnectionsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectionsSortOrderEnum
const (
	ListConnectionsSortOrderAsc  ListConnectionsSortOrderEnum = "ASC"
	ListConnectionsSortOrderDesc ListConnectionsSortOrderEnum = "DESC"
)

var mappingListConnectionsSortOrderEnum = map[string]ListConnectionsSortOrderEnum{
	"ASC":  ListConnectionsSortOrderAsc,
	"DESC": ListConnectionsSortOrderDesc,
}

var mappingListConnectionsSortOrderEnumLowerCase = map[string]ListConnectionsSortOrderEnum{
	"asc":  ListConnectionsSortOrderAsc,
	"desc": ListConnectionsSortOrderDesc,
}

// GetListConnectionsSortOrderEnumValues Enumerates the set of values for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumValues() []ListConnectionsSortOrderEnum {
	values := make([]ListConnectionsSortOrderEnum, 0)
	for _, v := range mappingListConnectionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionsSortOrderEnumStringValues Enumerates the set of values in String for ListConnectionsSortOrderEnum
func GetListConnectionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConnectionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionsSortOrderEnum(val string) (ListConnectionsSortOrderEnum, bool) {
	enum, ok := mappingListConnectionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
