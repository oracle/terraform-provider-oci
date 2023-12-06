// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAlarmsStatusRequest wrapper for the ListAlarmsStatus operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/ListAlarmsStatus.go.html to see an example of how to use ListAlarmsStatusRequest.
type ListAlarmsStatusRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the
	// resources monitored by the metric that you are searching for. Use tenancyId to search in
	// the root compartment.
	// Example: `ocid1.compartment.oc1..exampleuniqueID`
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer part of the request identifier token. If you need to contact Oracle about a particular
	// request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// When true, returns resources from all compartments and subcompartments. The parameter can
	// only be set to true when compartmentId is the tenancy OCID (the tenancy is the root compartment).
	// A true value requires the user to have tenancy-level permissions. If this requirement is not met,
	// then the call is rejected. When false, returns resources from only the compartment specified in
	// compartmentId. Default is false.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Default: 1000
	// Example: 500
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to return only resources that match the given display name exactly.
	// Use this filter to list an alarm by name. Alternatively, when you know the alarm OCID, use the GetAlarm operation.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to use when sorting returned alarm definitions. Only one sorting level is provided.
	// Example: `severity`
	SortBy ListAlarmsStatusSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use when sorting returned alarm definitions. Ascending (ASC) or descending (DESC).
	// Example: `ASC`
	SortOrder ListAlarmsStatusSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a resource that is monitored by the
	// metric that you are searching for.
	// Example: `ocid1.instance.oc1.phx.exampleuniqueID`
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only resources that match the given service name exactly.
	// Use this filter to list all alarms containing metric streams that match the *exact* service-name dimension.
	// Example: `logging-analytics`
	ServiceName *string `mandatory:"false" contributesTo:"query" name:"serviceName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the entity monitored by the
	// metric that you are searching for.
	// Example: `ocid1.instance.oc1.phx.exampleuniqueID`
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// The status of the metric stream to use for alarm filtering. For example, set `StatusQueryParam` to
	// "FIRING" to filter results to metric streams of the alarm with that status. Default behaviour is to return
	// alarms irrespective of metric streams' status.
	// Example: `FIRING`
	Status ListAlarmsStatusStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlarmsStatusRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlarmsStatusRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAlarmsStatusRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlarmsStatusRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAlarmsStatusRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAlarmsStatusSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAlarmsStatusSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlarmsStatusSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAlarmsStatusSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlarmsStatusStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListAlarmsStatusStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAlarmsStatusResponse wrapper for the ListAlarmsStatus operation
type ListAlarmsStatusResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AlarmStatusSummary instances
	Items []AlarmStatusSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAlarmsStatusResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlarmsStatusResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAlarmsStatusSortByEnum Enum with underlying type: string
type ListAlarmsStatusSortByEnum string

// Set of constants representing the allowable values for ListAlarmsStatusSortByEnum
const (
	ListAlarmsStatusSortByDisplayname ListAlarmsStatusSortByEnum = "displayName"
	ListAlarmsStatusSortBySeverity    ListAlarmsStatusSortByEnum = "severity"
)

var mappingListAlarmsStatusSortByEnum = map[string]ListAlarmsStatusSortByEnum{
	"displayName": ListAlarmsStatusSortByDisplayname,
	"severity":    ListAlarmsStatusSortBySeverity,
}

var mappingListAlarmsStatusSortByEnumLowerCase = map[string]ListAlarmsStatusSortByEnum{
	"displayname": ListAlarmsStatusSortByDisplayname,
	"severity":    ListAlarmsStatusSortBySeverity,
}

// GetListAlarmsStatusSortByEnumValues Enumerates the set of values for ListAlarmsStatusSortByEnum
func GetListAlarmsStatusSortByEnumValues() []ListAlarmsStatusSortByEnum {
	values := make([]ListAlarmsStatusSortByEnum, 0)
	for _, v := range mappingListAlarmsStatusSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmsStatusSortByEnumStringValues Enumerates the set of values in String for ListAlarmsStatusSortByEnum
func GetListAlarmsStatusSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"severity",
	}
}

// GetMappingListAlarmsStatusSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmsStatusSortByEnum(val string) (ListAlarmsStatusSortByEnum, bool) {
	enum, ok := mappingListAlarmsStatusSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlarmsStatusSortOrderEnum Enum with underlying type: string
type ListAlarmsStatusSortOrderEnum string

// Set of constants representing the allowable values for ListAlarmsStatusSortOrderEnum
const (
	ListAlarmsStatusSortOrderAsc  ListAlarmsStatusSortOrderEnum = "ASC"
	ListAlarmsStatusSortOrderDesc ListAlarmsStatusSortOrderEnum = "DESC"
)

var mappingListAlarmsStatusSortOrderEnum = map[string]ListAlarmsStatusSortOrderEnum{
	"ASC":  ListAlarmsStatusSortOrderAsc,
	"DESC": ListAlarmsStatusSortOrderDesc,
}

var mappingListAlarmsStatusSortOrderEnumLowerCase = map[string]ListAlarmsStatusSortOrderEnum{
	"asc":  ListAlarmsStatusSortOrderAsc,
	"desc": ListAlarmsStatusSortOrderDesc,
}

// GetListAlarmsStatusSortOrderEnumValues Enumerates the set of values for ListAlarmsStatusSortOrderEnum
func GetListAlarmsStatusSortOrderEnumValues() []ListAlarmsStatusSortOrderEnum {
	values := make([]ListAlarmsStatusSortOrderEnum, 0)
	for _, v := range mappingListAlarmsStatusSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmsStatusSortOrderEnumStringValues Enumerates the set of values in String for ListAlarmsStatusSortOrderEnum
func GetListAlarmsStatusSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAlarmsStatusSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmsStatusSortOrderEnum(val string) (ListAlarmsStatusSortOrderEnum, bool) {
	enum, ok := mappingListAlarmsStatusSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlarmsStatusStatusEnum Enum with underlying type: string
type ListAlarmsStatusStatusEnum string

// Set of constants representing the allowable values for ListAlarmsStatusStatusEnum
const (
	ListAlarmsStatusStatusFiring ListAlarmsStatusStatusEnum = "FIRING"
	ListAlarmsStatusStatusOk     ListAlarmsStatusStatusEnum = "OK"
)

var mappingListAlarmsStatusStatusEnum = map[string]ListAlarmsStatusStatusEnum{
	"FIRING": ListAlarmsStatusStatusFiring,
	"OK":     ListAlarmsStatusStatusOk,
}

var mappingListAlarmsStatusStatusEnumLowerCase = map[string]ListAlarmsStatusStatusEnum{
	"firing": ListAlarmsStatusStatusFiring,
	"ok":     ListAlarmsStatusStatusOk,
}

// GetListAlarmsStatusStatusEnumValues Enumerates the set of values for ListAlarmsStatusStatusEnum
func GetListAlarmsStatusStatusEnumValues() []ListAlarmsStatusStatusEnum {
	values := make([]ListAlarmsStatusStatusEnum, 0)
	for _, v := range mappingListAlarmsStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmsStatusStatusEnumStringValues Enumerates the set of values in String for ListAlarmsStatusStatusEnum
func GetListAlarmsStatusStatusEnumStringValues() []string {
	return []string{
		"FIRING",
		"OK",
	}
}

// GetMappingListAlarmsStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmsStatusStatusEnum(val string) (ListAlarmsStatusStatusEnum, bool) {
	enum, ok := mappingListAlarmsStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
