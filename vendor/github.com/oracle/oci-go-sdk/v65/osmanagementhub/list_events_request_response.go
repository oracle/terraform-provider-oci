// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEventsRequest wrapper for the ListEvents operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListEvents.go.html to see an example of how to use ListEventsRequest.
type ListEventsRequest struct {

	// A filter to return only events whose summary matches the given value.
	EventSummary *string `mandatory:"false" contributesTo:"query" name:"eventSummary"`

	// A filter to return only events with a summary that contains the value provided.
	EventSummaryContains *string `mandatory:"false" contributesTo:"query" name:"eventSummaryContains"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the event.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The eventFingerprint of the KernelEventData.
	EventFingerprint *string `mandatory:"false" contributesTo:"query" name:"eventFingerprint"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only events that match the state provided. The state value is case-insensitive.
	LifecycleState EventLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. This filter returns resources associated with the specified resource.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only resources whose type matches the given value.
	Type []EventTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter that returns events that occurred on or before the date provided.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// A filter that returns events that occurred on or after the date provided.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated, timeOccurredAt and timeUpdated is descending. Default order for eventSummary is ascending.
	SortBy ListEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicates whether to list only resources managed by the Autonomous Linux service.
	IsManagedByAutonomousLinux *bool `mandatory:"false" contributesTo:"query" name:"isManagedByAutonomousLinux"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEventLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEventLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Type {
		if _, ok := GetMappingEventTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", val, strings.Join(GetEventTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEventsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEventsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEventsResponse wrapper for the ListEvents operation
type ListEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EventCollection instances
	EventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEventsSortOrderEnum Enum with underlying type: string
type ListEventsSortOrderEnum string

// Set of constants representing the allowable values for ListEventsSortOrderEnum
const (
	ListEventsSortOrderAsc  ListEventsSortOrderEnum = "ASC"
	ListEventsSortOrderDesc ListEventsSortOrderEnum = "DESC"
)

var mappingListEventsSortOrderEnum = map[string]ListEventsSortOrderEnum{
	"ASC":  ListEventsSortOrderAsc,
	"DESC": ListEventsSortOrderDesc,
}

var mappingListEventsSortOrderEnumLowerCase = map[string]ListEventsSortOrderEnum{
	"asc":  ListEventsSortOrderAsc,
	"desc": ListEventsSortOrderDesc,
}

// GetListEventsSortOrderEnumValues Enumerates the set of values for ListEventsSortOrderEnum
func GetListEventsSortOrderEnumValues() []ListEventsSortOrderEnum {
	values := make([]ListEventsSortOrderEnum, 0)
	for _, v := range mappingListEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsSortOrderEnumStringValues Enumerates the set of values in String for ListEventsSortOrderEnum
func GetListEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsSortOrderEnum(val string) (ListEventsSortOrderEnum, bool) {
	enum, ok := mappingListEventsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEventsSortByEnum Enum with underlying type: string
type ListEventsSortByEnum string

// Set of constants representing the allowable values for ListEventsSortByEnum
const (
	ListEventsSortByTimecreated    ListEventsSortByEnum = "timeCreated"
	ListEventsSortByTimeoccurredat ListEventsSortByEnum = "timeOccurredAt"
	ListEventsSortByTimeupdated    ListEventsSortByEnum = "timeUpdated"
	ListEventsSortByEventsummary   ListEventsSortByEnum = "eventSummary"
)

var mappingListEventsSortByEnum = map[string]ListEventsSortByEnum{
	"timeCreated":    ListEventsSortByTimecreated,
	"timeOccurredAt": ListEventsSortByTimeoccurredat,
	"timeUpdated":    ListEventsSortByTimeupdated,
	"eventSummary":   ListEventsSortByEventsummary,
}

var mappingListEventsSortByEnumLowerCase = map[string]ListEventsSortByEnum{
	"timecreated":    ListEventsSortByTimecreated,
	"timeoccurredat": ListEventsSortByTimeoccurredat,
	"timeupdated":    ListEventsSortByTimeupdated,
	"eventsummary":   ListEventsSortByEventsummary,
}

// GetListEventsSortByEnumValues Enumerates the set of values for ListEventsSortByEnum
func GetListEventsSortByEnumValues() []ListEventsSortByEnum {
	values := make([]ListEventsSortByEnum, 0)
	for _, v := range mappingListEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsSortByEnumStringValues Enumerates the set of values in String for ListEventsSortByEnum
func GetListEventsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeOccurredAt",
		"timeUpdated",
		"eventSummary",
	}
}

// GetMappingListEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsSortByEnum(val string) (ListEventsSortByEnum, bool) {
	enum, ok := mappingListEventsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
