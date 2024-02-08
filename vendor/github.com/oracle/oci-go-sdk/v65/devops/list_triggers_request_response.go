// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTriggersRequest wrapper for the ListTriggers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListTriggers.go.html to see an example of how to use ListTriggersRequest.
type ListTriggersRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// unique project identifier
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// A filter to return only triggers that matches the given lifecycle state.
	LifecycleState TriggerLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique trigger identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListTriggersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListTriggersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTriggersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTriggersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTriggersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTriggersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTriggersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTriggerLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetTriggerLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTriggersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTriggersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTriggersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTriggersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTriggersResponse wrapper for the ListTriggers operation
type ListTriggersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TriggerCollection instances
	TriggerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTriggersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTriggersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTriggersSortOrderEnum Enum with underlying type: string
type ListTriggersSortOrderEnum string

// Set of constants representing the allowable values for ListTriggersSortOrderEnum
const (
	ListTriggersSortOrderAsc  ListTriggersSortOrderEnum = "ASC"
	ListTriggersSortOrderDesc ListTriggersSortOrderEnum = "DESC"
)

var mappingListTriggersSortOrderEnum = map[string]ListTriggersSortOrderEnum{
	"ASC":  ListTriggersSortOrderAsc,
	"DESC": ListTriggersSortOrderDesc,
}

var mappingListTriggersSortOrderEnumLowerCase = map[string]ListTriggersSortOrderEnum{
	"asc":  ListTriggersSortOrderAsc,
	"desc": ListTriggersSortOrderDesc,
}

// GetListTriggersSortOrderEnumValues Enumerates the set of values for ListTriggersSortOrderEnum
func GetListTriggersSortOrderEnumValues() []ListTriggersSortOrderEnum {
	values := make([]ListTriggersSortOrderEnum, 0)
	for _, v := range mappingListTriggersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTriggersSortOrderEnumStringValues Enumerates the set of values in String for ListTriggersSortOrderEnum
func GetListTriggersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTriggersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTriggersSortOrderEnum(val string) (ListTriggersSortOrderEnum, bool) {
	enum, ok := mappingListTriggersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTriggersSortByEnum Enum with underlying type: string
type ListTriggersSortByEnum string

// Set of constants representing the allowable values for ListTriggersSortByEnum
const (
	ListTriggersSortByTimecreated ListTriggersSortByEnum = "timeCreated"
	ListTriggersSortByDisplayname ListTriggersSortByEnum = "displayName"
)

var mappingListTriggersSortByEnum = map[string]ListTriggersSortByEnum{
	"timeCreated": ListTriggersSortByTimecreated,
	"displayName": ListTriggersSortByDisplayname,
}

var mappingListTriggersSortByEnumLowerCase = map[string]ListTriggersSortByEnum{
	"timecreated": ListTriggersSortByTimecreated,
	"displayname": ListTriggersSortByDisplayname,
}

// GetListTriggersSortByEnumValues Enumerates the set of values for ListTriggersSortByEnum
func GetListTriggersSortByEnumValues() []ListTriggersSortByEnum {
	values := make([]ListTriggersSortByEnum, 0)
	for _, v := range mappingListTriggersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTriggersSortByEnumStringValues Enumerates the set of values in String for ListTriggersSortByEnum
func GetListTriggersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTriggersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTriggersSortByEnum(val string) (ListTriggersSortByEnum, bool) {
	enum, ok := mappingListTriggersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
