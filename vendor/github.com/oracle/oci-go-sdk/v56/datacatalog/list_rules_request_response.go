// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListRulesRequest wrapper for the ListRules operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListRules.go.html to see an example of how to use ListRulesRequest.
type ListRulesRequest struct {

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
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Rule type used to filter the response to a list rules call.
	RuleType ListRulesRuleTypeEnum `mandatory:"false" contributesTo:"query" name:"ruleType" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Rule origin type used to filter the response to a list rules call.
	OriginType ListRulesOriginTypeEnum `mandatory:"false" contributesTo:"query" name:"originType" omitEmpty:"true"`

	// Unique external identifier of this resource in the external source system.
	ExternalKey *string `mandatory:"false" contributesTo:"query" name:"externalKey"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Specifies the fields to return in a rule summary response.
	Fields []ListRulesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRulesResponse wrapper for the ListRules operation
type ListRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RuleCollection instances
	RuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRulesRuleTypeEnum Enum with underlying type: string
type ListRulesRuleTypeEnum string

// Set of constants representing the allowable values for ListRulesRuleTypeEnum
const (
	ListRulesRuleTypePrimarykey ListRulesRuleTypeEnum = "PRIMARYKEY"
	ListRulesRuleTypeForeignkey ListRulesRuleTypeEnum = "FOREIGNKEY"
	ListRulesRuleTypeUniquekey  ListRulesRuleTypeEnum = "UNIQUEKEY"
)

var mappingListRulesRuleType = map[string]ListRulesRuleTypeEnum{
	"PRIMARYKEY": ListRulesRuleTypePrimarykey,
	"FOREIGNKEY": ListRulesRuleTypeForeignkey,
	"UNIQUEKEY":  ListRulesRuleTypeUniquekey,
}

// GetListRulesRuleTypeEnumValues Enumerates the set of values for ListRulesRuleTypeEnum
func GetListRulesRuleTypeEnumValues() []ListRulesRuleTypeEnum {
	values := make([]ListRulesRuleTypeEnum, 0)
	for _, v := range mappingListRulesRuleType {
		values = append(values, v)
	}
	return values
}

// ListRulesLifecycleStateEnum Enum with underlying type: string
type ListRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListRulesLifecycleStateEnum
const (
	ListRulesLifecycleStateCreating ListRulesLifecycleStateEnum = "CREATING"
	ListRulesLifecycleStateActive   ListRulesLifecycleStateEnum = "ACTIVE"
	ListRulesLifecycleStateInactive ListRulesLifecycleStateEnum = "INACTIVE"
	ListRulesLifecycleStateUpdating ListRulesLifecycleStateEnum = "UPDATING"
	ListRulesLifecycleStateDeleting ListRulesLifecycleStateEnum = "DELETING"
	ListRulesLifecycleStateDeleted  ListRulesLifecycleStateEnum = "DELETED"
	ListRulesLifecycleStateFailed   ListRulesLifecycleStateEnum = "FAILED"
	ListRulesLifecycleStateMoving   ListRulesLifecycleStateEnum = "MOVING"
)

var mappingListRulesLifecycleState = map[string]ListRulesLifecycleStateEnum{
	"CREATING": ListRulesLifecycleStateCreating,
	"ACTIVE":   ListRulesLifecycleStateActive,
	"INACTIVE": ListRulesLifecycleStateInactive,
	"UPDATING": ListRulesLifecycleStateUpdating,
	"DELETING": ListRulesLifecycleStateDeleting,
	"DELETED":  ListRulesLifecycleStateDeleted,
	"FAILED":   ListRulesLifecycleStateFailed,
	"MOVING":   ListRulesLifecycleStateMoving,
}

// GetListRulesLifecycleStateEnumValues Enumerates the set of values for ListRulesLifecycleStateEnum
func GetListRulesLifecycleStateEnumValues() []ListRulesLifecycleStateEnum {
	values := make([]ListRulesLifecycleStateEnum, 0)
	for _, v := range mappingListRulesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListRulesOriginTypeEnum Enum with underlying type: string
type ListRulesOriginTypeEnum string

// Set of constants representing the allowable values for ListRulesOriginTypeEnum
const (
	ListRulesOriginTypeSource    ListRulesOriginTypeEnum = "SOURCE"
	ListRulesOriginTypeUser      ListRulesOriginTypeEnum = "USER"
	ListRulesOriginTypeProfiling ListRulesOriginTypeEnum = "PROFILING"
)

var mappingListRulesOriginType = map[string]ListRulesOriginTypeEnum{
	"SOURCE":    ListRulesOriginTypeSource,
	"USER":      ListRulesOriginTypeUser,
	"PROFILING": ListRulesOriginTypeProfiling,
}

// GetListRulesOriginTypeEnumValues Enumerates the set of values for ListRulesOriginTypeEnum
func GetListRulesOriginTypeEnumValues() []ListRulesOriginTypeEnum {
	values := make([]ListRulesOriginTypeEnum, 0)
	for _, v := range mappingListRulesOriginType {
		values = append(values, v)
	}
	return values
}

// ListRulesFieldsEnum Enum with underlying type: string
type ListRulesFieldsEnum string

// Set of constants representing the allowable values for ListRulesFieldsEnum
const (
	ListRulesFieldsKey                  ListRulesFieldsEnum = "key"
	ListRulesFieldsDisplayname          ListRulesFieldsEnum = "displayName"
	ListRulesFieldsRuletype             ListRulesFieldsEnum = "ruleType"
	ListRulesFieldsExternalkey          ListRulesFieldsEnum = "externalKey"
	ListRulesFieldsReferencedfolderkey  ListRulesFieldsEnum = "referencedFolderKey"
	ListRulesFieldsReferencedfoldername ListRulesFieldsEnum = "referencedFolderName"
	ListRulesFieldsReferencedentitykey  ListRulesFieldsEnum = "referencedEntityKey"
	ListRulesFieldsReferencedentityname ListRulesFieldsEnum = "referencedEntityName"
	ListRulesFieldsReferencedrulekey    ListRulesFieldsEnum = "referencedRuleKey"
	ListRulesFieldsReferencedrulename   ListRulesFieldsEnum = "referencedRuleName"
	ListRulesFieldsOrigintype           ListRulesFieldsEnum = "originType"
	ListRulesFieldsLifecyclestate       ListRulesFieldsEnum = "lifecycleState"
	ListRulesFieldsTimecreated          ListRulesFieldsEnum = "timeCreated"
	ListRulesFieldsUri                  ListRulesFieldsEnum = "uri"
)

var mappingListRulesFields = map[string]ListRulesFieldsEnum{
	"key":                  ListRulesFieldsKey,
	"displayName":          ListRulesFieldsDisplayname,
	"ruleType":             ListRulesFieldsRuletype,
	"externalKey":          ListRulesFieldsExternalkey,
	"referencedFolderKey":  ListRulesFieldsReferencedfolderkey,
	"referencedFolderName": ListRulesFieldsReferencedfoldername,
	"referencedEntityKey":  ListRulesFieldsReferencedentitykey,
	"referencedEntityName": ListRulesFieldsReferencedentityname,
	"referencedRuleKey":    ListRulesFieldsReferencedrulekey,
	"referencedRuleName":   ListRulesFieldsReferencedrulename,
	"originType":           ListRulesFieldsOrigintype,
	"lifecycleState":       ListRulesFieldsLifecyclestate,
	"timeCreated":          ListRulesFieldsTimecreated,
	"uri":                  ListRulesFieldsUri,
}

// GetListRulesFieldsEnumValues Enumerates the set of values for ListRulesFieldsEnum
func GetListRulesFieldsEnumValues() []ListRulesFieldsEnum {
	values := make([]ListRulesFieldsEnum, 0)
	for _, v := range mappingListRulesFields {
		values = append(values, v)
	}
	return values
}

// ListRulesSortByEnum Enum with underlying type: string
type ListRulesSortByEnum string

// Set of constants representing the allowable values for ListRulesSortByEnum
const (
	ListRulesSortByTimecreated ListRulesSortByEnum = "TIMECREATED"
	ListRulesSortByDisplayname ListRulesSortByEnum = "DISPLAYNAME"
)

var mappingListRulesSortBy = map[string]ListRulesSortByEnum{
	"TIMECREATED": ListRulesSortByTimecreated,
	"DISPLAYNAME": ListRulesSortByDisplayname,
}

// GetListRulesSortByEnumValues Enumerates the set of values for ListRulesSortByEnum
func GetListRulesSortByEnumValues() []ListRulesSortByEnum {
	values := make([]ListRulesSortByEnum, 0)
	for _, v := range mappingListRulesSortBy {
		values = append(values, v)
	}
	return values
}

// ListRulesSortOrderEnum Enum with underlying type: string
type ListRulesSortOrderEnum string

// Set of constants representing the allowable values for ListRulesSortOrderEnum
const (
	ListRulesSortOrderAsc  ListRulesSortOrderEnum = "ASC"
	ListRulesSortOrderDesc ListRulesSortOrderEnum = "DESC"
)

var mappingListRulesSortOrder = map[string]ListRulesSortOrderEnum{
	"ASC":  ListRulesSortOrderAsc,
	"DESC": ListRulesSortOrderDesc,
}

// GetListRulesSortOrderEnumValues Enumerates the set of values for ListRulesSortOrderEnum
func GetListRulesSortOrderEnumValues() []ListRulesSortOrderEnum {
	values := make([]ListRulesSortOrderEnum, 0)
	for _, v := range mappingListRulesSortOrder {
		values = append(values, v)
	}
	return values
}
