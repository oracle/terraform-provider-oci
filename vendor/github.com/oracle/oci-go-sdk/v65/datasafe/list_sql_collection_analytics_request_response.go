// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSqlCollectionAnalyticsRequest wrapper for the ListSqlCollectionAnalytics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlCollectionAnalytics.go.html to see an example of how to use ListSqlCollectionAnalyticsRequest.
type ListSqlCollectionAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSqlCollectionAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the SQL collection.
	LifecycleState ListSqlCollectionAnalyticsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The group by parameter to summarize SQL collection aggregation.
	GroupBy []ListSqlCollectionAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// An optional filter to return the stats of the SQL collection logs collected after the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStarted"`

	// An optional filter to return the stats of the SQL collection logs collected before the date-time specified,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnded *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnded"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlCollectionAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlCollectionAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlCollectionAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlCollectionAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlCollectionAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlCollectionAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSqlCollectionAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlCollectionAnalyticsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSqlCollectionAnalyticsLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingListSqlCollectionAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListSqlCollectionAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlCollectionAnalyticsResponse wrapper for the ListSqlCollectionAnalytics operation
type ListSqlCollectionAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlCollectionAnalyticsCollection instances
	SqlCollectionAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlCollectionAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlCollectionAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlCollectionAnalyticsAccessLevelEnum Enum with underlying type: string
type ListSqlCollectionAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListSqlCollectionAnalyticsAccessLevelEnum
const (
	ListSqlCollectionAnalyticsAccessLevelRestricted ListSqlCollectionAnalyticsAccessLevelEnum = "RESTRICTED"
	ListSqlCollectionAnalyticsAccessLevelAccessible ListSqlCollectionAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSqlCollectionAnalyticsAccessLevelEnum = map[string]ListSqlCollectionAnalyticsAccessLevelEnum{
	"RESTRICTED": ListSqlCollectionAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListSqlCollectionAnalyticsAccessLevelAccessible,
}

var mappingListSqlCollectionAnalyticsAccessLevelEnumLowerCase = map[string]ListSqlCollectionAnalyticsAccessLevelEnum{
	"restricted": ListSqlCollectionAnalyticsAccessLevelRestricted,
	"accessible": ListSqlCollectionAnalyticsAccessLevelAccessible,
}

// GetListSqlCollectionAnalyticsAccessLevelEnumValues Enumerates the set of values for ListSqlCollectionAnalyticsAccessLevelEnum
func GetListSqlCollectionAnalyticsAccessLevelEnumValues() []ListSqlCollectionAnalyticsAccessLevelEnum {
	values := make([]ListSqlCollectionAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListSqlCollectionAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlCollectionAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListSqlCollectionAnalyticsAccessLevelEnum
func GetListSqlCollectionAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSqlCollectionAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlCollectionAnalyticsAccessLevelEnum(val string) (ListSqlCollectionAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListSqlCollectionAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlCollectionAnalyticsLifecycleStateEnum Enum with underlying type: string
type ListSqlCollectionAnalyticsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSqlCollectionAnalyticsLifecycleStateEnum
const (
	ListSqlCollectionAnalyticsLifecycleStateCreating       ListSqlCollectionAnalyticsLifecycleStateEnum = "CREATING"
	ListSqlCollectionAnalyticsLifecycleStateUpdating       ListSqlCollectionAnalyticsLifecycleStateEnum = "UPDATING"
	ListSqlCollectionAnalyticsLifecycleStateCollecting     ListSqlCollectionAnalyticsLifecycleStateEnum = "COLLECTING"
	ListSqlCollectionAnalyticsLifecycleStateCompleted      ListSqlCollectionAnalyticsLifecycleStateEnum = "COMPLETED"
	ListSqlCollectionAnalyticsLifecycleStateInactive       ListSqlCollectionAnalyticsLifecycleStateEnum = "INACTIVE"
	ListSqlCollectionAnalyticsLifecycleStateFailed         ListSqlCollectionAnalyticsLifecycleStateEnum = "FAILED"
	ListSqlCollectionAnalyticsLifecycleStateDeleting       ListSqlCollectionAnalyticsLifecycleStateEnum = "DELETING"
	ListSqlCollectionAnalyticsLifecycleStateDeleted        ListSqlCollectionAnalyticsLifecycleStateEnum = "DELETED"
	ListSqlCollectionAnalyticsLifecycleStateNeedsAttention ListSqlCollectionAnalyticsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListSqlCollectionAnalyticsLifecycleStateEnum = map[string]ListSqlCollectionAnalyticsLifecycleStateEnum{
	"CREATING":        ListSqlCollectionAnalyticsLifecycleStateCreating,
	"UPDATING":        ListSqlCollectionAnalyticsLifecycleStateUpdating,
	"COLLECTING":      ListSqlCollectionAnalyticsLifecycleStateCollecting,
	"COMPLETED":       ListSqlCollectionAnalyticsLifecycleStateCompleted,
	"INACTIVE":        ListSqlCollectionAnalyticsLifecycleStateInactive,
	"FAILED":          ListSqlCollectionAnalyticsLifecycleStateFailed,
	"DELETING":        ListSqlCollectionAnalyticsLifecycleStateDeleting,
	"DELETED":         ListSqlCollectionAnalyticsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListSqlCollectionAnalyticsLifecycleStateNeedsAttention,
}

var mappingListSqlCollectionAnalyticsLifecycleStateEnumLowerCase = map[string]ListSqlCollectionAnalyticsLifecycleStateEnum{
	"creating":        ListSqlCollectionAnalyticsLifecycleStateCreating,
	"updating":        ListSqlCollectionAnalyticsLifecycleStateUpdating,
	"collecting":      ListSqlCollectionAnalyticsLifecycleStateCollecting,
	"completed":       ListSqlCollectionAnalyticsLifecycleStateCompleted,
	"inactive":        ListSqlCollectionAnalyticsLifecycleStateInactive,
	"failed":          ListSqlCollectionAnalyticsLifecycleStateFailed,
	"deleting":        ListSqlCollectionAnalyticsLifecycleStateDeleting,
	"deleted":         ListSqlCollectionAnalyticsLifecycleStateDeleted,
	"needs_attention": ListSqlCollectionAnalyticsLifecycleStateNeedsAttention,
}

// GetListSqlCollectionAnalyticsLifecycleStateEnumValues Enumerates the set of values for ListSqlCollectionAnalyticsLifecycleStateEnum
func GetListSqlCollectionAnalyticsLifecycleStateEnumValues() []ListSqlCollectionAnalyticsLifecycleStateEnum {
	values := make([]ListSqlCollectionAnalyticsLifecycleStateEnum, 0)
	for _, v := range mappingListSqlCollectionAnalyticsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlCollectionAnalyticsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSqlCollectionAnalyticsLifecycleStateEnum
func GetListSqlCollectionAnalyticsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"COLLECTING",
		"COMPLETED",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListSqlCollectionAnalyticsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlCollectionAnalyticsLifecycleStateEnum(val string) (ListSqlCollectionAnalyticsLifecycleStateEnum, bool) {
	enum, ok := mappingListSqlCollectionAnalyticsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlCollectionAnalyticsGroupByEnum Enum with underlying type: string
type ListSqlCollectionAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListSqlCollectionAnalyticsGroupByEnum
const (
	ListSqlCollectionAnalyticsGroupByTargetid       ListSqlCollectionAnalyticsGroupByEnum = "targetId"
	ListSqlCollectionAnalyticsGroupByLifecyclestate ListSqlCollectionAnalyticsGroupByEnum = "lifecycleState"
)

var mappingListSqlCollectionAnalyticsGroupByEnum = map[string]ListSqlCollectionAnalyticsGroupByEnum{
	"targetId":       ListSqlCollectionAnalyticsGroupByTargetid,
	"lifecycleState": ListSqlCollectionAnalyticsGroupByLifecyclestate,
}

var mappingListSqlCollectionAnalyticsGroupByEnumLowerCase = map[string]ListSqlCollectionAnalyticsGroupByEnum{
	"targetid":       ListSqlCollectionAnalyticsGroupByTargetid,
	"lifecyclestate": ListSqlCollectionAnalyticsGroupByLifecyclestate,
}

// GetListSqlCollectionAnalyticsGroupByEnumValues Enumerates the set of values for ListSqlCollectionAnalyticsGroupByEnum
func GetListSqlCollectionAnalyticsGroupByEnumValues() []ListSqlCollectionAnalyticsGroupByEnum {
	values := make([]ListSqlCollectionAnalyticsGroupByEnum, 0)
	for _, v := range mappingListSqlCollectionAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlCollectionAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListSqlCollectionAnalyticsGroupByEnum
func GetListSqlCollectionAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"targetId",
		"lifecycleState",
	}
}

// GetMappingListSqlCollectionAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlCollectionAnalyticsGroupByEnum(val string) (ListSqlCollectionAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListSqlCollectionAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
