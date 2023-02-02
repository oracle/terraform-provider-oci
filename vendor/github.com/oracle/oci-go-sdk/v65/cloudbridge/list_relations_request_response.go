// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRelationsRequest wrapper for the ListRelations operation
type ListRelationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only relation whose lifecycleState matches the given lifecycleState.
	LifecycleState RelationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique relation identifier.
	RelationId *string `mandatory:"false" contributesTo:"query" name:"relationId"`

	// The source ocid where the relation is starting from.
	SourceId *string `mandatory:"false" contributesTo:"query" name:"sourceId"`

	// The target ocid where the relation is ended to.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// External relation key.
	ExternalRelationKey *string `mandatory:"false" contributesTo:"query" name:"externalRelationKey"`

	// Relation type.
	RelationType ListRelationsRelationTypeEnum `mandatory:"false" contributesTo:"query" name:"relationType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRelationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListRelationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRelationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRelationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRelationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRelationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRelationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRelationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRelationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRelationsRelationTypeEnum(string(request.RelationType)); !ok && request.RelationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelationType: %s. Supported values are: %s.", request.RelationType, strings.Join(GetListRelationsRelationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRelationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRelationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRelationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRelationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRelationsResponse wrapper for the ListRelations operation
type ListRelationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RelationCollection instances
	RelationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRelationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRelationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRelationsRelationTypeEnum Enum with underlying type: string
type ListRelationsRelationTypeEnum string

// Set of constants representing the allowable values for ListRelationsRelationTypeEnum
const (
	ListRelationsRelationTypeAssociation ListRelationsRelationTypeEnum = "ASSOCIATION"
	ListRelationsRelationTypeDependency  ListRelationsRelationTypeEnum = "DEPENDENCY"
	ListRelationsRelationTypeComposition ListRelationsRelationTypeEnum = "COMPOSITION"
)

var mappingListRelationsRelationTypeEnum = map[string]ListRelationsRelationTypeEnum{
	"ASSOCIATION": ListRelationsRelationTypeAssociation,
	"DEPENDENCY":  ListRelationsRelationTypeDependency,
	"COMPOSITION": ListRelationsRelationTypeComposition,
}

var mappingListRelationsRelationTypeEnumLowerCase = map[string]ListRelationsRelationTypeEnum{
	"association": ListRelationsRelationTypeAssociation,
	"dependency":  ListRelationsRelationTypeDependency,
	"composition": ListRelationsRelationTypeComposition,
}

// GetListRelationsRelationTypeEnumValues Enumerates the set of values for ListRelationsRelationTypeEnum
func GetListRelationsRelationTypeEnumValues() []ListRelationsRelationTypeEnum {
	values := make([]ListRelationsRelationTypeEnum, 0)
	for _, v := range mappingListRelationsRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListRelationsRelationTypeEnumStringValues Enumerates the set of values in String for ListRelationsRelationTypeEnum
func GetListRelationsRelationTypeEnumStringValues() []string {
	return []string{
		"ASSOCIATION",
		"DEPENDENCY",
		"COMPOSITION",
	}
}

// GetMappingListRelationsRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRelationsRelationTypeEnum(val string) (ListRelationsRelationTypeEnum, bool) {
	enum, ok := mappingListRelationsRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRelationsSortOrderEnum Enum with underlying type: string
type ListRelationsSortOrderEnum string

// Set of constants representing the allowable values for ListRelationsSortOrderEnum
const (
	ListRelationsSortOrderAsc  ListRelationsSortOrderEnum = "ASC"
	ListRelationsSortOrderDesc ListRelationsSortOrderEnum = "DESC"
)

var mappingListRelationsSortOrderEnum = map[string]ListRelationsSortOrderEnum{
	"ASC":  ListRelationsSortOrderAsc,
	"DESC": ListRelationsSortOrderDesc,
}

var mappingListRelationsSortOrderEnumLowerCase = map[string]ListRelationsSortOrderEnum{
	"asc":  ListRelationsSortOrderAsc,
	"desc": ListRelationsSortOrderDesc,
}

// GetListRelationsSortOrderEnumValues Enumerates the set of values for ListRelationsSortOrderEnum
func GetListRelationsSortOrderEnumValues() []ListRelationsSortOrderEnum {
	values := make([]ListRelationsSortOrderEnum, 0)
	for _, v := range mappingListRelationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRelationsSortOrderEnumStringValues Enumerates the set of values in String for ListRelationsSortOrderEnum
func GetListRelationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRelationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRelationsSortOrderEnum(val string) (ListRelationsSortOrderEnum, bool) {
	enum, ok := mappingListRelationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRelationsSortByEnum Enum with underlying type: string
type ListRelationsSortByEnum string

// Set of constants representing the allowable values for ListRelationsSortByEnum
const (
	ListRelationsSortByTimecreated ListRelationsSortByEnum = "timeCreated"
	ListRelationsSortByTimeupdated ListRelationsSortByEnum = "timeUpdated"
	ListRelationsSortByDisplayname ListRelationsSortByEnum = "displayName"
)

var mappingListRelationsSortByEnum = map[string]ListRelationsSortByEnum{
	"timeCreated": ListRelationsSortByTimecreated,
	"timeUpdated": ListRelationsSortByTimeupdated,
	"displayName": ListRelationsSortByDisplayname,
}

var mappingListRelationsSortByEnumLowerCase = map[string]ListRelationsSortByEnum{
	"timecreated": ListRelationsSortByTimecreated,
	"timeupdated": ListRelationsSortByTimeupdated,
	"displayname": ListRelationsSortByDisplayname,
}

// GetListRelationsSortByEnumValues Enumerates the set of values for ListRelationsSortByEnum
func GetListRelationsSortByEnumValues() []ListRelationsSortByEnum {
	values := make([]ListRelationsSortByEnum, 0)
	for _, v := range mappingListRelationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRelationsSortByEnumStringValues Enumerates the set of values in String for ListRelationsSortByEnum
func GetListRelationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListRelationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRelationsSortByEnum(val string) (ListRelationsSortByEnum, bool) {
	enum, ok := mappingListRelationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
