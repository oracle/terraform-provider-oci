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

// ListAttributesRequest wrapper for the ListAttributes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListAttributes.go.html to see an example of how to use ListAttributesRequest.
type ListAttributesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the entire business name given. The match is not case sensitive.
	BusinessName *string `mandatory:"false" contributesTo:"query" name:"businessName"`

	// A filter to return only resources that match display name or business name pattern given. The match is not case sensitive.
	// For Example : /folders?displayOrBusinessNameContains=Cu.*
	// The above would match all folders with display name or business name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayOrBusinessNameContains *string `mandatory:"false" contributesTo:"query" name:"displayOrBusinessNameContains"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListAttributesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExternal"`

	// Data type as defined in an external system.
	ExternalTypeName *string `mandatory:"false" contributesTo:"query" name:"externalTypeName"`

	// Identifies whether this attribute can be used as a watermark to extract incremental data.
	IsIncrementalData *bool `mandatory:"false" contributesTo:"query" name:"isIncrementalData"`

	// Identifies whether this attribute can be assigned null value.
	IsNullable *bool `mandatory:"false" contributesTo:"query" name:"isNullable"`

	// Max allowed length of the attribute value.
	Length *int64 `mandatory:"false" contributesTo:"query" name:"length"`

	// Position of the attribute in the record definition.
	Position *int `mandatory:"false" contributesTo:"query" name:"position"`

	// Precision of the attribute value usually applies to float data type.
	Precision *int `mandatory:"false" contributesTo:"query" name:"precision"`

	// Scale of the attribute value usually applies to float data type.
	Scale *int `mandatory:"false" contributesTo:"query" name:"scale"`

	// Specifies the fields to return in an entity attribute summary response.
	Fields []ListAttributesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. Default order for POSITION is ascending. If no value is specified POSITION is default.
	SortBy ListAttributesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAttributesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListAttributesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAttributesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAttributesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAttributesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAttributesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAttributesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAttributesLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListAttributesFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListAttributesFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListAttributesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAttributesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttributesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAttributesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAttributesResponse wrapper for the ListAttributes operation
type ListAttributesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AttributeCollection instances
	AttributeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAttributesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAttributesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAttributesLifecycleStateEnum Enum with underlying type: string
type ListAttributesLifecycleStateEnum string

// Set of constants representing the allowable values for ListAttributesLifecycleStateEnum
const (
	ListAttributesLifecycleStateCreating ListAttributesLifecycleStateEnum = "CREATING"
	ListAttributesLifecycleStateActive   ListAttributesLifecycleStateEnum = "ACTIVE"
	ListAttributesLifecycleStateInactive ListAttributesLifecycleStateEnum = "INACTIVE"
	ListAttributesLifecycleStateUpdating ListAttributesLifecycleStateEnum = "UPDATING"
	ListAttributesLifecycleStateDeleting ListAttributesLifecycleStateEnum = "DELETING"
	ListAttributesLifecycleStateDeleted  ListAttributesLifecycleStateEnum = "DELETED"
	ListAttributesLifecycleStateFailed   ListAttributesLifecycleStateEnum = "FAILED"
	ListAttributesLifecycleStateMoving   ListAttributesLifecycleStateEnum = "MOVING"
)

var mappingListAttributesLifecycleStateEnum = map[string]ListAttributesLifecycleStateEnum{
	"CREATING": ListAttributesLifecycleStateCreating,
	"ACTIVE":   ListAttributesLifecycleStateActive,
	"INACTIVE": ListAttributesLifecycleStateInactive,
	"UPDATING": ListAttributesLifecycleStateUpdating,
	"DELETING": ListAttributesLifecycleStateDeleting,
	"DELETED":  ListAttributesLifecycleStateDeleted,
	"FAILED":   ListAttributesLifecycleStateFailed,
	"MOVING":   ListAttributesLifecycleStateMoving,
}

// GetListAttributesLifecycleStateEnumValues Enumerates the set of values for ListAttributesLifecycleStateEnum
func GetListAttributesLifecycleStateEnumValues() []ListAttributesLifecycleStateEnum {
	values := make([]ListAttributesLifecycleStateEnum, 0)
	for _, v := range mappingListAttributesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributesLifecycleStateEnumStringValues Enumerates the set of values in String for ListAttributesLifecycleStateEnum
func GetListAttributesLifecycleStateEnumStringValues() []string {
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

// GetMappingListAttributesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributesLifecycleStateEnum(val string) (ListAttributesLifecycleStateEnum, bool) {
	mappingListAttributesLifecycleStateEnumIgnoreCase := make(map[string]ListAttributesLifecycleStateEnum)
	for k, v := range mappingListAttributesLifecycleStateEnum {
		mappingListAttributesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAttributesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributesFieldsEnum Enum with underlying type: string
type ListAttributesFieldsEnum string

// Set of constants representing the allowable values for ListAttributesFieldsEnum
const (
	ListAttributesFieldsKey                        ListAttributesFieldsEnum = "key"
	ListAttributesFieldsDisplayname                ListAttributesFieldsEnum = "displayName"
	ListAttributesFieldsDescription                ListAttributesFieldsEnum = "description"
	ListAttributesFieldsEntitykey                  ListAttributesFieldsEnum = "entityKey"
	ListAttributesFieldsLifecyclestate             ListAttributesFieldsEnum = "lifecycleState"
	ListAttributesFieldsTimecreated                ListAttributesFieldsEnum = "timeCreated"
	ListAttributesFieldsExternaldatatype           ListAttributesFieldsEnum = "externalDataType"
	ListAttributesFieldsExternalkey                ListAttributesFieldsEnum = "externalKey"
	ListAttributesFieldsLength                     ListAttributesFieldsEnum = "length"
	ListAttributesFieldsPrecision                  ListAttributesFieldsEnum = "precision"
	ListAttributesFieldsScale                      ListAttributesFieldsEnum = "scale"
	ListAttributesFieldsIsnullable                 ListAttributesFieldsEnum = "isNullable"
	ListAttributesFieldsUri                        ListAttributesFieldsEnum = "uri"
	ListAttributesFieldsPath                       ListAttributesFieldsEnum = "path"
	ListAttributesFieldsMincollectioncount         ListAttributesFieldsEnum = "minCollectionCount"
	ListAttributesFieldsMaxcollectioncount         ListAttributesFieldsEnum = "maxCollectionCount"
	ListAttributesFieldsDatatypeentitykey          ListAttributesFieldsEnum = "datatypeEntityKey"
	ListAttributesFieldsExternaldatatypeentitykey  ListAttributesFieldsEnum = "externalDatatypeEntityKey"
	ListAttributesFieldsParentattributekey         ListAttributesFieldsEnum = "parentAttributeKey"
	ListAttributesFieldsExternalparentattributekey ListAttributesFieldsEnum = "externalParentAttributeKey"
	ListAttributesFieldsPosition                   ListAttributesFieldsEnum = "position"
	ListAttributesFieldsTypekey                    ListAttributesFieldsEnum = "typeKey"
)

var mappingListAttributesFieldsEnum = map[string]ListAttributesFieldsEnum{
	"key":                        ListAttributesFieldsKey,
	"displayName":                ListAttributesFieldsDisplayname,
	"description":                ListAttributesFieldsDescription,
	"entityKey":                  ListAttributesFieldsEntitykey,
	"lifecycleState":             ListAttributesFieldsLifecyclestate,
	"timeCreated":                ListAttributesFieldsTimecreated,
	"externalDataType":           ListAttributesFieldsExternaldatatype,
	"externalKey":                ListAttributesFieldsExternalkey,
	"length":                     ListAttributesFieldsLength,
	"precision":                  ListAttributesFieldsPrecision,
	"scale":                      ListAttributesFieldsScale,
	"isNullable":                 ListAttributesFieldsIsnullable,
	"uri":                        ListAttributesFieldsUri,
	"path":                       ListAttributesFieldsPath,
	"minCollectionCount":         ListAttributesFieldsMincollectioncount,
	"maxCollectionCount":         ListAttributesFieldsMaxcollectioncount,
	"datatypeEntityKey":          ListAttributesFieldsDatatypeentitykey,
	"externalDatatypeEntityKey":  ListAttributesFieldsExternaldatatypeentitykey,
	"parentAttributeKey":         ListAttributesFieldsParentattributekey,
	"externalParentAttributeKey": ListAttributesFieldsExternalparentattributekey,
	"position":                   ListAttributesFieldsPosition,
	"typeKey":                    ListAttributesFieldsTypekey,
}

// GetListAttributesFieldsEnumValues Enumerates the set of values for ListAttributesFieldsEnum
func GetListAttributesFieldsEnumValues() []ListAttributesFieldsEnum {
	values := make([]ListAttributesFieldsEnum, 0)
	for _, v := range mappingListAttributesFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributesFieldsEnumStringValues Enumerates the set of values in String for ListAttributesFieldsEnum
func GetListAttributesFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"entityKey",
		"lifecycleState",
		"timeCreated",
		"externalDataType",
		"externalKey",
		"length",
		"precision",
		"scale",
		"isNullable",
		"uri",
		"path",
		"minCollectionCount",
		"maxCollectionCount",
		"datatypeEntityKey",
		"externalDatatypeEntityKey",
		"parentAttributeKey",
		"externalParentAttributeKey",
		"position",
		"typeKey",
	}
}

// GetMappingListAttributesFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributesFieldsEnum(val string) (ListAttributesFieldsEnum, bool) {
	mappingListAttributesFieldsEnumIgnoreCase := make(map[string]ListAttributesFieldsEnum)
	for k, v := range mappingListAttributesFieldsEnum {
		mappingListAttributesFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAttributesFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributesSortByEnum Enum with underlying type: string
type ListAttributesSortByEnum string

// Set of constants representing the allowable values for ListAttributesSortByEnum
const (
	ListAttributesSortByTimecreated ListAttributesSortByEnum = "TIMECREATED"
	ListAttributesSortByDisplayname ListAttributesSortByEnum = "DISPLAYNAME"
	ListAttributesSortByPosition    ListAttributesSortByEnum = "POSITION"
)

var mappingListAttributesSortByEnum = map[string]ListAttributesSortByEnum{
	"TIMECREATED": ListAttributesSortByTimecreated,
	"DISPLAYNAME": ListAttributesSortByDisplayname,
	"POSITION":    ListAttributesSortByPosition,
}

// GetListAttributesSortByEnumValues Enumerates the set of values for ListAttributesSortByEnum
func GetListAttributesSortByEnumValues() []ListAttributesSortByEnum {
	values := make([]ListAttributesSortByEnum, 0)
	for _, v := range mappingListAttributesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributesSortByEnumStringValues Enumerates the set of values in String for ListAttributesSortByEnum
func GetListAttributesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
		"POSITION",
	}
}

// GetMappingListAttributesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributesSortByEnum(val string) (ListAttributesSortByEnum, bool) {
	mappingListAttributesSortByEnumIgnoreCase := make(map[string]ListAttributesSortByEnum)
	for k, v := range mappingListAttributesSortByEnum {
		mappingListAttributesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAttributesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttributesSortOrderEnum Enum with underlying type: string
type ListAttributesSortOrderEnum string

// Set of constants representing the allowable values for ListAttributesSortOrderEnum
const (
	ListAttributesSortOrderAsc  ListAttributesSortOrderEnum = "ASC"
	ListAttributesSortOrderDesc ListAttributesSortOrderEnum = "DESC"
)

var mappingListAttributesSortOrderEnum = map[string]ListAttributesSortOrderEnum{
	"ASC":  ListAttributesSortOrderAsc,
	"DESC": ListAttributesSortOrderDesc,
}

// GetListAttributesSortOrderEnumValues Enumerates the set of values for ListAttributesSortOrderEnum
func GetListAttributesSortOrderEnumValues() []ListAttributesSortOrderEnum {
	values := make([]ListAttributesSortOrderEnum, 0)
	for _, v := range mappingListAttributesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttributesSortOrderEnumStringValues Enumerates the set of values in String for ListAttributesSortOrderEnum
func GetListAttributesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAttributesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttributesSortOrderEnum(val string) (ListAttributesSortOrderEnum, bool) {
	mappingListAttributesSortOrderEnumIgnoreCase := make(map[string]ListAttributesSortOrderEnum)
	for k, v := range mappingListAttributesSortOrderEnum {
		mappingListAttributesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAttributesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
