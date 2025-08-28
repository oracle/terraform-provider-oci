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

// ListManagedInstancesRequest wrapper for the ListManagedInstances operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListManagedInstances.go.html to see an example of how to use ListManagedInstancesRequest.
type ListManagedInstancesRequest struct {

	// The OCID of the compartment that contains the resources to list. This filter returns
	// only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The plugin status of the managed instance.
	PluginStatus ListManagedInstancesPluginStatusEnum `mandatory:"false" contributesTo:"query" name:"pluginStatus" omitEmpty:"true"`

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
	SortOrder ListManagedInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _timeCreated_ is **descending**.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified, _timeCreated_ is default.
	SortBy ListManagedInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstancesPluginStatusEnum(string(request.PluginStatus)); !ok && request.PluginStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PluginStatus: %s. Supported values are: %s.", request.PluginStatus, strings.Join(GetListManagedInstancesPluginStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstancesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstancesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstancesResponse wrapper for the ListManagedInstances operation
type ListManagedInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceCollection instances
	ManagedInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstancesPluginStatusEnum Enum with underlying type: string
type ListManagedInstancesPluginStatusEnum string

// Set of constants representing the allowable values for ListManagedInstancesPluginStatusEnum
const (
	ListManagedInstancesPluginStatusActive   ListManagedInstancesPluginStatusEnum = "ACTIVE"
	ListManagedInstancesPluginStatusInactive ListManagedInstancesPluginStatusEnum = "INACTIVE"
)

var mappingListManagedInstancesPluginStatusEnum = map[string]ListManagedInstancesPluginStatusEnum{
	"ACTIVE":   ListManagedInstancesPluginStatusActive,
	"INACTIVE": ListManagedInstancesPluginStatusInactive,
}

var mappingListManagedInstancesPluginStatusEnumLowerCase = map[string]ListManagedInstancesPluginStatusEnum{
	"active":   ListManagedInstancesPluginStatusActive,
	"inactive": ListManagedInstancesPluginStatusInactive,
}

// GetListManagedInstancesPluginStatusEnumValues Enumerates the set of values for ListManagedInstancesPluginStatusEnum
func GetListManagedInstancesPluginStatusEnumValues() []ListManagedInstancesPluginStatusEnum {
	values := make([]ListManagedInstancesPluginStatusEnum, 0)
	for _, v := range mappingListManagedInstancesPluginStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstancesPluginStatusEnumStringValues Enumerates the set of values in String for ListManagedInstancesPluginStatusEnum
func GetListManagedInstancesPluginStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListManagedInstancesPluginStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstancesPluginStatusEnum(val string) (ListManagedInstancesPluginStatusEnum, bool) {
	enum, ok := mappingListManagedInstancesPluginStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstancesSortOrderEnum Enum with underlying type: string
type ListManagedInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstancesSortOrderEnum
const (
	ListManagedInstancesSortOrderAsc  ListManagedInstancesSortOrderEnum = "ASC"
	ListManagedInstancesSortOrderDesc ListManagedInstancesSortOrderEnum = "DESC"
)

var mappingListManagedInstancesSortOrderEnum = map[string]ListManagedInstancesSortOrderEnum{
	"ASC":  ListManagedInstancesSortOrderAsc,
	"DESC": ListManagedInstancesSortOrderDesc,
}

var mappingListManagedInstancesSortOrderEnumLowerCase = map[string]ListManagedInstancesSortOrderEnum{
	"asc":  ListManagedInstancesSortOrderAsc,
	"desc": ListManagedInstancesSortOrderDesc,
}

// GetListManagedInstancesSortOrderEnumValues Enumerates the set of values for ListManagedInstancesSortOrderEnum
func GetListManagedInstancesSortOrderEnumValues() []ListManagedInstancesSortOrderEnum {
	values := make([]ListManagedInstancesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstancesSortOrderEnum
func GetListManagedInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstancesSortOrderEnum(val string) (ListManagedInstancesSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstancesSortByEnum Enum with underlying type: string
type ListManagedInstancesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstancesSortByEnum
const (
	ListManagedInstancesSortByTimecreated ListManagedInstancesSortByEnum = "timeCreated"
	ListManagedInstancesSortByDisplayname ListManagedInstancesSortByEnum = "displayName"
)

var mappingListManagedInstancesSortByEnum = map[string]ListManagedInstancesSortByEnum{
	"timeCreated": ListManagedInstancesSortByTimecreated,
	"displayName": ListManagedInstancesSortByDisplayname,
}

var mappingListManagedInstancesSortByEnumLowerCase = map[string]ListManagedInstancesSortByEnum{
	"timecreated": ListManagedInstancesSortByTimecreated,
	"displayname": ListManagedInstancesSortByDisplayname,
}

// GetListManagedInstancesSortByEnumValues Enumerates the set of values for ListManagedInstancesSortByEnum
func GetListManagedInstancesSortByEnumValues() []ListManagedInstancesSortByEnum {
	values := make([]ListManagedInstancesSortByEnum, 0)
	for _, v := range mappingListManagedInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstancesSortByEnumStringValues Enumerates the set of values in String for ListManagedInstancesSortByEnum
func GetListManagedInstancesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListManagedInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstancesSortByEnum(val string) (ListManagedInstancesSortByEnum, bool) {
	enum, ok := mappingListManagedInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
