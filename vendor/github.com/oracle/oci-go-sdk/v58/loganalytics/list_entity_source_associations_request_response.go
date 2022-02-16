// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListEntitySourceAssociationsRequest wrapper for the ListEntitySourceAssociations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListEntitySourceAssociations.go.html to see an example of how to use ListEntitySourceAssociationsRequest.
type ListEntitySourceAssociationsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The entity OCID.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// The entity type used for filtering.  Only associations on an entity with the
	// specified type will be returned.
	EntityType *string `mandatory:"false" contributesTo:"query" name:"entityType"`

	// The entity type display name used for filtering.  Only items associated with the entity
	// with the specified type display name will be returned.
	EntityTypeDisplayName *string `mandatory:"false" contributesTo:"query" name:"entityTypeDisplayName"`

	// The life cycle state used for filtering.  Only associations with the specified
	// life cycle state will be returned.
	LifeCycleState ListEntitySourceAssociationsLifeCycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifeCycleState" omitEmpty:"true"`

	// A flag indicating whether or not to return the total number of items returned.
	IsShowTotal *bool `mandatory:"false" contributesTo:"query" name:"isShowTotal"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListEntitySourceAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned associations
	SortBy ListEntitySourceAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEntitySourceAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEntitySourceAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEntitySourceAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEntitySourceAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEntitySourceAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEntitySourceAssociationsLifeCycleStateEnum(string(request.LifeCycleState)); !ok && request.LifeCycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifeCycleState: %s. Supported values are: %s.", request.LifeCycleState, strings.Join(GetListEntitySourceAssociationsLifeCycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntitySourceAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEntitySourceAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntitySourceAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEntitySourceAssociationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEntitySourceAssociationsResponse wrapper for the ListEntitySourceAssociations operation
type ListEntitySourceAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsAssociationCollection instances
	LogAnalyticsAssociationCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListEntitySourceAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEntitySourceAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEntitySourceAssociationsLifeCycleStateEnum Enum with underlying type: string
type ListEntitySourceAssociationsLifeCycleStateEnum string

// Set of constants representing the allowable values for ListEntitySourceAssociationsLifeCycleStateEnum
const (
	ListEntitySourceAssociationsLifeCycleStateAll        ListEntitySourceAssociationsLifeCycleStateEnum = "ALL"
	ListEntitySourceAssociationsLifeCycleStateAccepted   ListEntitySourceAssociationsLifeCycleStateEnum = "ACCEPTED"
	ListEntitySourceAssociationsLifeCycleStateInProgress ListEntitySourceAssociationsLifeCycleStateEnum = "IN_PROGRESS"
	ListEntitySourceAssociationsLifeCycleStateSucceeded  ListEntitySourceAssociationsLifeCycleStateEnum = "SUCCEEDED"
	ListEntitySourceAssociationsLifeCycleStateFailed     ListEntitySourceAssociationsLifeCycleStateEnum = "FAILED"
)

var mappingListEntitySourceAssociationsLifeCycleStateEnum = map[string]ListEntitySourceAssociationsLifeCycleStateEnum{
	"ALL":         ListEntitySourceAssociationsLifeCycleStateAll,
	"ACCEPTED":    ListEntitySourceAssociationsLifeCycleStateAccepted,
	"IN_PROGRESS": ListEntitySourceAssociationsLifeCycleStateInProgress,
	"SUCCEEDED":   ListEntitySourceAssociationsLifeCycleStateSucceeded,
	"FAILED":      ListEntitySourceAssociationsLifeCycleStateFailed,
}

// GetListEntitySourceAssociationsLifeCycleStateEnumValues Enumerates the set of values for ListEntitySourceAssociationsLifeCycleStateEnum
func GetListEntitySourceAssociationsLifeCycleStateEnumValues() []ListEntitySourceAssociationsLifeCycleStateEnum {
	values := make([]ListEntitySourceAssociationsLifeCycleStateEnum, 0)
	for _, v := range mappingListEntitySourceAssociationsLifeCycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitySourceAssociationsLifeCycleStateEnumStringValues Enumerates the set of values in String for ListEntitySourceAssociationsLifeCycleStateEnum
func GetListEntitySourceAssociationsLifeCycleStateEnumStringValues() []string {
	return []string{
		"ALL",
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingListEntitySourceAssociationsLifeCycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitySourceAssociationsLifeCycleStateEnum(val string) (ListEntitySourceAssociationsLifeCycleStateEnum, bool) {
	mappingListEntitySourceAssociationsLifeCycleStateEnumIgnoreCase := make(map[string]ListEntitySourceAssociationsLifeCycleStateEnum)
	for k, v := range mappingListEntitySourceAssociationsLifeCycleStateEnum {
		mappingListEntitySourceAssociationsLifeCycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntitySourceAssociationsLifeCycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntitySourceAssociationsSortOrderEnum Enum with underlying type: string
type ListEntitySourceAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListEntitySourceAssociationsSortOrderEnum
const (
	ListEntitySourceAssociationsSortOrderAsc  ListEntitySourceAssociationsSortOrderEnum = "ASC"
	ListEntitySourceAssociationsSortOrderDesc ListEntitySourceAssociationsSortOrderEnum = "DESC"
)

var mappingListEntitySourceAssociationsSortOrderEnum = map[string]ListEntitySourceAssociationsSortOrderEnum{
	"ASC":  ListEntitySourceAssociationsSortOrderAsc,
	"DESC": ListEntitySourceAssociationsSortOrderDesc,
}

// GetListEntitySourceAssociationsSortOrderEnumValues Enumerates the set of values for ListEntitySourceAssociationsSortOrderEnum
func GetListEntitySourceAssociationsSortOrderEnumValues() []ListEntitySourceAssociationsSortOrderEnum {
	values := make([]ListEntitySourceAssociationsSortOrderEnum, 0)
	for _, v := range mappingListEntitySourceAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitySourceAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListEntitySourceAssociationsSortOrderEnum
func GetListEntitySourceAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEntitySourceAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitySourceAssociationsSortOrderEnum(val string) (ListEntitySourceAssociationsSortOrderEnum, bool) {
	mappingListEntitySourceAssociationsSortOrderEnumIgnoreCase := make(map[string]ListEntitySourceAssociationsSortOrderEnum)
	for k, v := range mappingListEntitySourceAssociationsSortOrderEnum {
		mappingListEntitySourceAssociationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntitySourceAssociationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntitySourceAssociationsSortByEnum Enum with underlying type: string
type ListEntitySourceAssociationsSortByEnum string

// Set of constants representing the allowable values for ListEntitySourceAssociationsSortByEnum
const (
	ListEntitySourceAssociationsSortBySourcedisplayname ListEntitySourceAssociationsSortByEnum = "sourceDisplayName"
	ListEntitySourceAssociationsSortByTimelastattempted ListEntitySourceAssociationsSortByEnum = "timeLastAttempted"
	ListEntitySourceAssociationsSortByStatus            ListEntitySourceAssociationsSortByEnum = "status"
)

var mappingListEntitySourceAssociationsSortByEnum = map[string]ListEntitySourceAssociationsSortByEnum{
	"sourceDisplayName": ListEntitySourceAssociationsSortBySourcedisplayname,
	"timeLastAttempted": ListEntitySourceAssociationsSortByTimelastattempted,
	"status":            ListEntitySourceAssociationsSortByStatus,
}

// GetListEntitySourceAssociationsSortByEnumValues Enumerates the set of values for ListEntitySourceAssociationsSortByEnum
func GetListEntitySourceAssociationsSortByEnumValues() []ListEntitySourceAssociationsSortByEnum {
	values := make([]ListEntitySourceAssociationsSortByEnum, 0)
	for _, v := range mappingListEntitySourceAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitySourceAssociationsSortByEnumStringValues Enumerates the set of values in String for ListEntitySourceAssociationsSortByEnum
func GetListEntitySourceAssociationsSortByEnumStringValues() []string {
	return []string{
		"sourceDisplayName",
		"timeLastAttempted",
		"status",
	}
}

// GetMappingListEntitySourceAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitySourceAssociationsSortByEnum(val string) (ListEntitySourceAssociationsSortByEnum, bool) {
	mappingListEntitySourceAssociationsSortByEnumIgnoreCase := make(map[string]ListEntitySourceAssociationsSortByEnum)
	for k, v := range mappingListEntitySourceAssociationsSortByEnum {
		mappingListEntitySourceAssociationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListEntitySourceAssociationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
