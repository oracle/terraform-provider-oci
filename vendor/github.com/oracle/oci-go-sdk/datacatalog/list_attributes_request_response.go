// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAttributesRequest wrapper for the ListAttributes operation
type ListAttributesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
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

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
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
func (request ListAttributesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAttributesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
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

var mappingListAttributesLifecycleState = map[string]ListAttributesLifecycleStateEnum{
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
	for _, v := range mappingListAttributesLifecycleState {
		values = append(values, v)
	}
	return values
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
	ListAttributesFieldsIsnullable                 ListAttributesFieldsEnum = "isNullable"
	ListAttributesFieldsUri                        ListAttributesFieldsEnum = "uri"
	ListAttributesFieldsPath                       ListAttributesFieldsEnum = "path"
	ListAttributesFieldsMincollectioncount         ListAttributesFieldsEnum = "minCollectionCount"
	ListAttributesFieldsMaxcollectioncount         ListAttributesFieldsEnum = "maxCollectionCount"
	ListAttributesFieldsDatatypeentitykey          ListAttributesFieldsEnum = "datatypeEntityKey"
	ListAttributesFieldsExternaldatatypeentitykey  ListAttributesFieldsEnum = "externalDatatypeEntityKey"
	ListAttributesFieldsParentattributekey         ListAttributesFieldsEnum = "parentAttributeKey"
	ListAttributesFieldsExternalparentattributekey ListAttributesFieldsEnum = "externalParentAttributeKey"
)

var mappingListAttributesFields = map[string]ListAttributesFieldsEnum{
	"key":                        ListAttributesFieldsKey,
	"displayName":                ListAttributesFieldsDisplayname,
	"description":                ListAttributesFieldsDescription,
	"entityKey":                  ListAttributesFieldsEntitykey,
	"lifecycleState":             ListAttributesFieldsLifecyclestate,
	"timeCreated":                ListAttributesFieldsTimecreated,
	"externalDataType":           ListAttributesFieldsExternaldatatype,
	"externalKey":                ListAttributesFieldsExternalkey,
	"length":                     ListAttributesFieldsLength,
	"isNullable":                 ListAttributesFieldsIsnullable,
	"uri":                        ListAttributesFieldsUri,
	"path":                       ListAttributesFieldsPath,
	"minCollectionCount":         ListAttributesFieldsMincollectioncount,
	"maxCollectionCount":         ListAttributesFieldsMaxcollectioncount,
	"datatypeEntityKey":          ListAttributesFieldsDatatypeentitykey,
	"externalDatatypeEntityKey":  ListAttributesFieldsExternaldatatypeentitykey,
	"parentAttributeKey":         ListAttributesFieldsParentattributekey,
	"externalParentAttributeKey": ListAttributesFieldsExternalparentattributekey,
}

// GetListAttributesFieldsEnumValues Enumerates the set of values for ListAttributesFieldsEnum
func GetListAttributesFieldsEnumValues() []ListAttributesFieldsEnum {
	values := make([]ListAttributesFieldsEnum, 0)
	for _, v := range mappingListAttributesFields {
		values = append(values, v)
	}
	return values
}

// ListAttributesSortByEnum Enum with underlying type: string
type ListAttributesSortByEnum string

// Set of constants representing the allowable values for ListAttributesSortByEnum
const (
	ListAttributesSortByTimecreated ListAttributesSortByEnum = "TIMECREATED"
	ListAttributesSortByDisplayname ListAttributesSortByEnum = "DISPLAYNAME"
)

var mappingListAttributesSortBy = map[string]ListAttributesSortByEnum{
	"TIMECREATED": ListAttributesSortByTimecreated,
	"DISPLAYNAME": ListAttributesSortByDisplayname,
}

// GetListAttributesSortByEnumValues Enumerates the set of values for ListAttributesSortByEnum
func GetListAttributesSortByEnumValues() []ListAttributesSortByEnum {
	values := make([]ListAttributesSortByEnum, 0)
	for _, v := range mappingListAttributesSortBy {
		values = append(values, v)
	}
	return values
}

// ListAttributesSortOrderEnum Enum with underlying type: string
type ListAttributesSortOrderEnum string

// Set of constants representing the allowable values for ListAttributesSortOrderEnum
const (
	ListAttributesSortOrderAsc  ListAttributesSortOrderEnum = "ASC"
	ListAttributesSortOrderDesc ListAttributesSortOrderEnum = "DESC"
)

var mappingListAttributesSortOrder = map[string]ListAttributesSortOrderEnum{
	"ASC":  ListAttributesSortOrderAsc,
	"DESC": ListAttributesSortOrderDesc,
}

// GetListAttributesSortOrderEnumValues Enumerates the set of values for ListAttributesSortOrderEnum
func GetListAttributesSortOrderEnumValues() []ListAttributesSortOrderEnum {
	values := make([]ListAttributesSortOrderEnum, 0)
	for _, v := range mappingListAttributesSortOrder {
		values = append(values, v)
	}
	return values
}
