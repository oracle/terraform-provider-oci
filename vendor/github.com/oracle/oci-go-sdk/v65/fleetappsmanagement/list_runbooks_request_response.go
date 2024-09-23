// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRunbooksRequest wrapper for the ListRunbooks operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbooks.go.html to see an example of how to use ListRunbooksRequest.
type ListRunbooksRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState RunbookLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique Runbook identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The ID of the runbook type.
	Type RunbookTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The runbook relevance of product or full-stack.
	RunbookRelevance RunbookRunbookRelevanceEnum `mandatory:"false" contributesTo:"query" name:"runbookRelevance" omitEmpty:"true"`

	// The ID of the runbook platform.
	Platform *string `mandatory:"false" contributesTo:"query" name:"platform"`

	// The runbook lifecycle.
	Operation *string `mandatory:"false" contributesTo:"query" name:"operation"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListRunbooksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListRunbooksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRunbooksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRunbooksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRunbooksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRunbooksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRunbooksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRunbookLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetRunbookLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRunbookTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetRunbookTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRunbookRunbookRelevanceEnum(string(request.RunbookRelevance)); !ok && request.RunbookRelevance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RunbookRelevance: %s. Supported values are: %s.", request.RunbookRelevance, strings.Join(GetRunbookRunbookRelevanceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRunbooksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRunbooksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRunbooksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRunbooksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRunbooksResponse wrapper for the ListRunbooks operation
type ListRunbooksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RunbookCollection instances
	RunbookCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRunbooksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRunbooksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRunbooksSortOrderEnum Enum with underlying type: string
type ListRunbooksSortOrderEnum string

// Set of constants representing the allowable values for ListRunbooksSortOrderEnum
const (
	ListRunbooksSortOrderAsc  ListRunbooksSortOrderEnum = "ASC"
	ListRunbooksSortOrderDesc ListRunbooksSortOrderEnum = "DESC"
)

var mappingListRunbooksSortOrderEnum = map[string]ListRunbooksSortOrderEnum{
	"ASC":  ListRunbooksSortOrderAsc,
	"DESC": ListRunbooksSortOrderDesc,
}

var mappingListRunbooksSortOrderEnumLowerCase = map[string]ListRunbooksSortOrderEnum{
	"asc":  ListRunbooksSortOrderAsc,
	"desc": ListRunbooksSortOrderDesc,
}

// GetListRunbooksSortOrderEnumValues Enumerates the set of values for ListRunbooksSortOrderEnum
func GetListRunbooksSortOrderEnumValues() []ListRunbooksSortOrderEnum {
	values := make([]ListRunbooksSortOrderEnum, 0)
	for _, v := range mappingListRunbooksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunbooksSortOrderEnumStringValues Enumerates the set of values in String for ListRunbooksSortOrderEnum
func GetListRunbooksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRunbooksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunbooksSortOrderEnum(val string) (ListRunbooksSortOrderEnum, bool) {
	enum, ok := mappingListRunbooksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRunbooksSortByEnum Enum with underlying type: string
type ListRunbooksSortByEnum string

// Set of constants representing the allowable values for ListRunbooksSortByEnum
const (
	ListRunbooksSortByTimecreated ListRunbooksSortByEnum = "timeCreated"
	ListRunbooksSortByDisplayname ListRunbooksSortByEnum = "displayName"
)

var mappingListRunbooksSortByEnum = map[string]ListRunbooksSortByEnum{
	"timeCreated": ListRunbooksSortByTimecreated,
	"displayName": ListRunbooksSortByDisplayname,
}

var mappingListRunbooksSortByEnumLowerCase = map[string]ListRunbooksSortByEnum{
	"timecreated": ListRunbooksSortByTimecreated,
	"displayname": ListRunbooksSortByDisplayname,
}

// GetListRunbooksSortByEnumValues Enumerates the set of values for ListRunbooksSortByEnum
func GetListRunbooksSortByEnumValues() []ListRunbooksSortByEnum {
	values := make([]ListRunbooksSortByEnum, 0)
	for _, v := range mappingListRunbooksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRunbooksSortByEnumStringValues Enumerates the set of values in String for ListRunbooksSortByEnum
func GetListRunbooksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListRunbooksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRunbooksSortByEnum(val string) (ListRunbooksSortByEnum, bool) {
	enum, ok := mappingListRunbooksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
