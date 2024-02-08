// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mediaservices

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMediaWorkflowJobsRequest wrapper for the ListMediaWorkflowJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflowJobs.go.html to see an example of how to use ListMediaWorkflowJobsRequest.
type ListMediaWorkflowJobsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// unique MediaWorkflowJob identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Unique MediaWorkflow identifier.
	MediaWorkflowId *string `mandatory:"false" contributesTo:"query" name:"mediaWorkflowId"`

	// A filter to return only the resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources with lifecycleState matching the given lifecycleState.
	LifecycleState MediaWorkflowJobLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A token representing the position at which to start retrieving results. This must come from the
	// `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The parameter sort by.
	SortBy ListMediaWorkflowJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMediaWorkflowJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMediaWorkflowJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMediaWorkflowJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMediaWorkflowJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMediaWorkflowJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMediaWorkflowJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMediaWorkflowJobLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMediaWorkflowJobLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMediaWorkflowJobsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMediaWorkflowJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMediaWorkflowJobsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMediaWorkflowJobsResponse wrapper for the ListMediaWorkflowJobs operation
type ListMediaWorkflowJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MediaWorkflowJobCollection instances
	MediaWorkflowJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMediaWorkflowJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMediaWorkflowJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMediaWorkflowJobsSortByEnum Enum with underlying type: string
type ListMediaWorkflowJobsSortByEnum string

// Set of constants representing the allowable values for ListMediaWorkflowJobsSortByEnum
const (
	ListMediaWorkflowJobsSortByTimecreated    ListMediaWorkflowJobsSortByEnum = "timeCreated"
	ListMediaWorkflowJobsSortByWorkflowid     ListMediaWorkflowJobsSortByEnum = "workflowId"
	ListMediaWorkflowJobsSortByLifecyclestate ListMediaWorkflowJobsSortByEnum = "lifecycleState"
)

var mappingListMediaWorkflowJobsSortByEnum = map[string]ListMediaWorkflowJobsSortByEnum{
	"timeCreated":    ListMediaWorkflowJobsSortByTimecreated,
	"workflowId":     ListMediaWorkflowJobsSortByWorkflowid,
	"lifecycleState": ListMediaWorkflowJobsSortByLifecyclestate,
}

var mappingListMediaWorkflowJobsSortByEnumLowerCase = map[string]ListMediaWorkflowJobsSortByEnum{
	"timecreated":    ListMediaWorkflowJobsSortByTimecreated,
	"workflowid":     ListMediaWorkflowJobsSortByWorkflowid,
	"lifecyclestate": ListMediaWorkflowJobsSortByLifecyclestate,
}

// GetListMediaWorkflowJobsSortByEnumValues Enumerates the set of values for ListMediaWorkflowJobsSortByEnum
func GetListMediaWorkflowJobsSortByEnumValues() []ListMediaWorkflowJobsSortByEnum {
	values := make([]ListMediaWorkflowJobsSortByEnum, 0)
	for _, v := range mappingListMediaWorkflowJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowJobsSortByEnumStringValues Enumerates the set of values in String for ListMediaWorkflowJobsSortByEnum
func GetListMediaWorkflowJobsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"workflowId",
		"lifecycleState",
	}
}

// GetMappingListMediaWorkflowJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowJobsSortByEnum(val string) (ListMediaWorkflowJobsSortByEnum, bool) {
	enum, ok := mappingListMediaWorkflowJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMediaWorkflowJobsSortOrderEnum Enum with underlying type: string
type ListMediaWorkflowJobsSortOrderEnum string

// Set of constants representing the allowable values for ListMediaWorkflowJobsSortOrderEnum
const (
	ListMediaWorkflowJobsSortOrderAsc  ListMediaWorkflowJobsSortOrderEnum = "ASC"
	ListMediaWorkflowJobsSortOrderDesc ListMediaWorkflowJobsSortOrderEnum = "DESC"
)

var mappingListMediaWorkflowJobsSortOrderEnum = map[string]ListMediaWorkflowJobsSortOrderEnum{
	"ASC":  ListMediaWorkflowJobsSortOrderAsc,
	"DESC": ListMediaWorkflowJobsSortOrderDesc,
}

var mappingListMediaWorkflowJobsSortOrderEnumLowerCase = map[string]ListMediaWorkflowJobsSortOrderEnum{
	"asc":  ListMediaWorkflowJobsSortOrderAsc,
	"desc": ListMediaWorkflowJobsSortOrderDesc,
}

// GetListMediaWorkflowJobsSortOrderEnumValues Enumerates the set of values for ListMediaWorkflowJobsSortOrderEnum
func GetListMediaWorkflowJobsSortOrderEnumValues() []ListMediaWorkflowJobsSortOrderEnum {
	values := make([]ListMediaWorkflowJobsSortOrderEnum, 0)
	for _, v := range mappingListMediaWorkflowJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMediaWorkflowJobsSortOrderEnumStringValues Enumerates the set of values in String for ListMediaWorkflowJobsSortOrderEnum
func GetListMediaWorkflowJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMediaWorkflowJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMediaWorkflowJobsSortOrderEnum(val string) (ListMediaWorkflowJobsSortOrderEnum, bool) {
	enum, ok := mappingListMediaWorkflowJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
