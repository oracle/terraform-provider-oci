// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWlsDomainsRequest wrapper for the ListWlsDomains operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomains.go.html to see an example of how to use ListWlsDomainsRequest.
type ListWlsDomainsRequest struct {

	// The OCID of the compartment that contains the resources to list. This filter returns
	// only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState WlsDomainLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return WebLogic domains based on the WebLogic version.
	WeblogicVersion ListWlsDomainsWeblogicVersionEnum `mandatory:"false" contributesTo:"query" name:"weblogicVersion" omitEmpty:"true"`

	// A filter to return WebLogic domains based on the type of middleware of the WebLogic domain.
	MiddlewareType ListWlsDomainsMiddlewareTypeEnum `mandatory:"false" contributesTo:"query" name:"middlewareType" omitEmpty:"true"`

	// A filter to return domains based on the patch readiness status.
	PatchReadinessStatus ListWlsDomainsPatchReadinessStatusEnum `mandatory:"false" contributesTo:"query" name:"patchReadinessStatus" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListWlsDomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _timeCreated_ is **descending**.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified, _timeCreated_ is default.
	SortBy ListWlsDomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlsDomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlsDomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlsDomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlsDomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlsDomainsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWlsDomainLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetWlsDomainLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainsWeblogicVersionEnum(string(request.WeblogicVersion)); !ok && request.WeblogicVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WeblogicVersion: %s. Supported values are: %s.", request.WeblogicVersion, strings.Join(GetListWlsDomainsWeblogicVersionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainsMiddlewareTypeEnum(string(request.MiddlewareType)); !ok && request.MiddlewareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MiddlewareType: %s. Supported values are: %s.", request.MiddlewareType, strings.Join(GetListWlsDomainsMiddlewareTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainsPatchReadinessStatusEnum(string(request.PatchReadinessStatus)); !ok && request.PatchReadinessStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchReadinessStatus: %s. Supported values are: %s.", request.PatchReadinessStatus, strings.Join(GetListWlsDomainsPatchReadinessStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlsDomainsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlsDomainsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlsDomainsResponse wrapper for the ListWlsDomains operation
type ListWlsDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WlsDomainCollection instances
	WlsDomainCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWlsDomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlsDomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlsDomainsWeblogicVersionEnum Enum with underlying type: string
type ListWlsDomainsWeblogicVersionEnum string

// Set of constants representing the allowable values for ListWlsDomainsWeblogicVersionEnum
const (
	ListWlsDomainsWeblogicVersionV12214 ListWlsDomainsWeblogicVersionEnum = "v12.2.1.4"
	ListWlsDomainsWeblogicVersionV14110 ListWlsDomainsWeblogicVersionEnum = "v14.1.1.0"
	ListWlsDomainsWeblogicVersionV14120 ListWlsDomainsWeblogicVersionEnum = "v14.1.2.0"
)

var mappingListWlsDomainsWeblogicVersionEnum = map[string]ListWlsDomainsWeblogicVersionEnum{
	"v12.2.1.4": ListWlsDomainsWeblogicVersionV12214,
	"v14.1.1.0": ListWlsDomainsWeblogicVersionV14110,
	"v14.1.2.0": ListWlsDomainsWeblogicVersionV14120,
}

var mappingListWlsDomainsWeblogicVersionEnumLowerCase = map[string]ListWlsDomainsWeblogicVersionEnum{
	"v12.2.1.4": ListWlsDomainsWeblogicVersionV12214,
	"v14.1.1.0": ListWlsDomainsWeblogicVersionV14110,
	"v14.1.2.0": ListWlsDomainsWeblogicVersionV14120,
}

// GetListWlsDomainsWeblogicVersionEnumValues Enumerates the set of values for ListWlsDomainsWeblogicVersionEnum
func GetListWlsDomainsWeblogicVersionEnumValues() []ListWlsDomainsWeblogicVersionEnum {
	values := make([]ListWlsDomainsWeblogicVersionEnum, 0)
	for _, v := range mappingListWlsDomainsWeblogicVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainsWeblogicVersionEnumStringValues Enumerates the set of values in String for ListWlsDomainsWeblogicVersionEnum
func GetListWlsDomainsWeblogicVersionEnumStringValues() []string {
	return []string{
		"v12.2.1.4",
		"v14.1.1.0",
		"v14.1.2.0",
	}
}

// GetMappingListWlsDomainsWeblogicVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainsWeblogicVersionEnum(val string) (ListWlsDomainsWeblogicVersionEnum, bool) {
	enum, ok := mappingListWlsDomainsWeblogicVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlsDomainsMiddlewareTypeEnum Enum with underlying type: string
type ListWlsDomainsMiddlewareTypeEnum string

// Set of constants representing the allowable values for ListWlsDomainsMiddlewareTypeEnum
const (
	ListWlsDomainsMiddlewareTypeFmw ListWlsDomainsMiddlewareTypeEnum = "FMW"
	ListWlsDomainsMiddlewareTypeWls ListWlsDomainsMiddlewareTypeEnum = "WLS"
)

var mappingListWlsDomainsMiddlewareTypeEnum = map[string]ListWlsDomainsMiddlewareTypeEnum{
	"FMW": ListWlsDomainsMiddlewareTypeFmw,
	"WLS": ListWlsDomainsMiddlewareTypeWls,
}

var mappingListWlsDomainsMiddlewareTypeEnumLowerCase = map[string]ListWlsDomainsMiddlewareTypeEnum{
	"fmw": ListWlsDomainsMiddlewareTypeFmw,
	"wls": ListWlsDomainsMiddlewareTypeWls,
}

// GetListWlsDomainsMiddlewareTypeEnumValues Enumerates the set of values for ListWlsDomainsMiddlewareTypeEnum
func GetListWlsDomainsMiddlewareTypeEnumValues() []ListWlsDomainsMiddlewareTypeEnum {
	values := make([]ListWlsDomainsMiddlewareTypeEnum, 0)
	for _, v := range mappingListWlsDomainsMiddlewareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainsMiddlewareTypeEnumStringValues Enumerates the set of values in String for ListWlsDomainsMiddlewareTypeEnum
func GetListWlsDomainsMiddlewareTypeEnumStringValues() []string {
	return []string{
		"FMW",
		"WLS",
	}
}

// GetMappingListWlsDomainsMiddlewareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainsMiddlewareTypeEnum(val string) (ListWlsDomainsMiddlewareTypeEnum, bool) {
	enum, ok := mappingListWlsDomainsMiddlewareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlsDomainsPatchReadinessStatusEnum Enum with underlying type: string
type ListWlsDomainsPatchReadinessStatusEnum string

// Set of constants representing the allowable values for ListWlsDomainsPatchReadinessStatusEnum
const (
	ListWlsDomainsPatchReadinessStatusOk      ListWlsDomainsPatchReadinessStatusEnum = "OK"
	ListWlsDomainsPatchReadinessStatusWarning ListWlsDomainsPatchReadinessStatusEnum = "WARNING"
	ListWlsDomainsPatchReadinessStatusError   ListWlsDomainsPatchReadinessStatusEnum = "ERROR"
	ListWlsDomainsPatchReadinessStatusUnknown ListWlsDomainsPatchReadinessStatusEnum = "UNKNOWN"
)

var mappingListWlsDomainsPatchReadinessStatusEnum = map[string]ListWlsDomainsPatchReadinessStatusEnum{
	"OK":      ListWlsDomainsPatchReadinessStatusOk,
	"WARNING": ListWlsDomainsPatchReadinessStatusWarning,
	"ERROR":   ListWlsDomainsPatchReadinessStatusError,
	"UNKNOWN": ListWlsDomainsPatchReadinessStatusUnknown,
}

var mappingListWlsDomainsPatchReadinessStatusEnumLowerCase = map[string]ListWlsDomainsPatchReadinessStatusEnum{
	"ok":      ListWlsDomainsPatchReadinessStatusOk,
	"warning": ListWlsDomainsPatchReadinessStatusWarning,
	"error":   ListWlsDomainsPatchReadinessStatusError,
	"unknown": ListWlsDomainsPatchReadinessStatusUnknown,
}

// GetListWlsDomainsPatchReadinessStatusEnumValues Enumerates the set of values for ListWlsDomainsPatchReadinessStatusEnum
func GetListWlsDomainsPatchReadinessStatusEnumValues() []ListWlsDomainsPatchReadinessStatusEnum {
	values := make([]ListWlsDomainsPatchReadinessStatusEnum, 0)
	for _, v := range mappingListWlsDomainsPatchReadinessStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainsPatchReadinessStatusEnumStringValues Enumerates the set of values in String for ListWlsDomainsPatchReadinessStatusEnum
func GetListWlsDomainsPatchReadinessStatusEnumStringValues() []string {
	return []string{
		"OK",
		"WARNING",
		"ERROR",
		"UNKNOWN",
	}
}

// GetMappingListWlsDomainsPatchReadinessStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainsPatchReadinessStatusEnum(val string) (ListWlsDomainsPatchReadinessStatusEnum, bool) {
	enum, ok := mappingListWlsDomainsPatchReadinessStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlsDomainsSortOrderEnum Enum with underlying type: string
type ListWlsDomainsSortOrderEnum string

// Set of constants representing the allowable values for ListWlsDomainsSortOrderEnum
const (
	ListWlsDomainsSortOrderAsc  ListWlsDomainsSortOrderEnum = "ASC"
	ListWlsDomainsSortOrderDesc ListWlsDomainsSortOrderEnum = "DESC"
)

var mappingListWlsDomainsSortOrderEnum = map[string]ListWlsDomainsSortOrderEnum{
	"ASC":  ListWlsDomainsSortOrderAsc,
	"DESC": ListWlsDomainsSortOrderDesc,
}

var mappingListWlsDomainsSortOrderEnumLowerCase = map[string]ListWlsDomainsSortOrderEnum{
	"asc":  ListWlsDomainsSortOrderAsc,
	"desc": ListWlsDomainsSortOrderDesc,
}

// GetListWlsDomainsSortOrderEnumValues Enumerates the set of values for ListWlsDomainsSortOrderEnum
func GetListWlsDomainsSortOrderEnumValues() []ListWlsDomainsSortOrderEnum {
	values := make([]ListWlsDomainsSortOrderEnum, 0)
	for _, v := range mappingListWlsDomainsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainsSortOrderEnumStringValues Enumerates the set of values in String for ListWlsDomainsSortOrderEnum
func GetListWlsDomainsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlsDomainsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainsSortOrderEnum(val string) (ListWlsDomainsSortOrderEnum, bool) {
	enum, ok := mappingListWlsDomainsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlsDomainsSortByEnum Enum with underlying type: string
type ListWlsDomainsSortByEnum string

// Set of constants representing the allowable values for ListWlsDomainsSortByEnum
const (
	ListWlsDomainsSortByTimecreated ListWlsDomainsSortByEnum = "timeCreated"
	ListWlsDomainsSortByDisplayname ListWlsDomainsSortByEnum = "displayName"
)

var mappingListWlsDomainsSortByEnum = map[string]ListWlsDomainsSortByEnum{
	"timeCreated": ListWlsDomainsSortByTimecreated,
	"displayName": ListWlsDomainsSortByDisplayname,
}

var mappingListWlsDomainsSortByEnumLowerCase = map[string]ListWlsDomainsSortByEnum{
	"timecreated": ListWlsDomainsSortByTimecreated,
	"displayname": ListWlsDomainsSortByDisplayname,
}

// GetListWlsDomainsSortByEnumValues Enumerates the set of values for ListWlsDomainsSortByEnum
func GetListWlsDomainsSortByEnumValues() []ListWlsDomainsSortByEnum {
	values := make([]ListWlsDomainsSortByEnum, 0)
	for _, v := range mappingListWlsDomainsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainsSortByEnumStringValues Enumerates the set of values in String for ListWlsDomainsSortByEnum
func GetListWlsDomainsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListWlsDomainsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainsSortByEnum(val string) (ListWlsDomainsSortByEnum, bool) {
	enum, ok := mappingListWlsDomainsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
