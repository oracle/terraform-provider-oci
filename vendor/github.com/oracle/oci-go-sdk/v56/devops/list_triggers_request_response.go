// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListTriggersRequest wrapper for the ListTriggers operation
//
// See also
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

var mappingListTriggersSortOrder = map[string]ListTriggersSortOrderEnum{
	"ASC":  ListTriggersSortOrderAsc,
	"DESC": ListTriggersSortOrderDesc,
}

// GetListTriggersSortOrderEnumValues Enumerates the set of values for ListTriggersSortOrderEnum
func GetListTriggersSortOrderEnumValues() []ListTriggersSortOrderEnum {
	values := make([]ListTriggersSortOrderEnum, 0)
	for _, v := range mappingListTriggersSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTriggersSortByEnum Enum with underlying type: string
type ListTriggersSortByEnum string

// Set of constants representing the allowable values for ListTriggersSortByEnum
const (
	ListTriggersSortByTimecreated ListTriggersSortByEnum = "timeCreated"
	ListTriggersSortByDisplayname ListTriggersSortByEnum = "displayName"
)

var mappingListTriggersSortBy = map[string]ListTriggersSortByEnum{
	"timeCreated": ListTriggersSortByTimecreated,
	"displayName": ListTriggersSortByDisplayname,
}

// GetListTriggersSortByEnumValues Enumerates the set of values for ListTriggersSortByEnum
func GetListTriggersSortByEnumValues() []ListTriggersSortByEnum {
	values := make([]ListTriggersSortByEnum, 0)
	for _, v := range mappingListTriggersSortBy {
		values = append(values, v)
	}
	return values
}
