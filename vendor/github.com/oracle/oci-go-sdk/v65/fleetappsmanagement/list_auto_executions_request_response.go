// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutoExecutionsRequest wrapper for the ListAutoExecutions operation
type ListAutoExecutionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Resource Identifier
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Returns results where the product name matches specified value.
	ProductName *string `mandatory:"false" contributesTo:"query" name:"productName"`

	// Returns results where the execution time is greater than or equal to the specified value
	TimeExecutionGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExecutionGreaterThanOrEqualTo"`

	// Returns results where the execution time is less than the specified value
	TimeExecutionLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExecutionLessThan"`

	// If set to true, resources will be returned for not only the provided compartment, but all compartments which
	// descend from it. Which resources are returned and their field contents depends on the value of accessLevel.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The field to sort by. Only one sort order may be provided.  Default order for timeCreated and timeScheduled is descending.
	SortBy ListAutoExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAutoExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutoExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutoExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutoExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutoExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutoExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutoExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutoExecutionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutoExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutoExecutionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutoExecutionsResponse wrapper for the ListAutoExecutions operation
type ListAutoExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AutoExecutionCollection instances
	AutoExecutionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutoExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutoExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutoExecutionsSortByEnum Enum with underlying type: string
type ListAutoExecutionsSortByEnum string

// Set of constants representing the allowable values for ListAutoExecutionsSortByEnum
const (
	ListAutoExecutionsSortByTimecreated   ListAutoExecutionsSortByEnum = "timeCreated"
	ListAutoExecutionsSortByTimescheduled ListAutoExecutionsSortByEnum = "timeScheduled"
)

var mappingListAutoExecutionsSortByEnum = map[string]ListAutoExecutionsSortByEnum{
	"timeCreated":   ListAutoExecutionsSortByTimecreated,
	"timeScheduled": ListAutoExecutionsSortByTimescheduled,
}

var mappingListAutoExecutionsSortByEnumLowerCase = map[string]ListAutoExecutionsSortByEnum{
	"timecreated":   ListAutoExecutionsSortByTimecreated,
	"timescheduled": ListAutoExecutionsSortByTimescheduled,
}

// GetListAutoExecutionsSortByEnumValues Enumerates the set of values for ListAutoExecutionsSortByEnum
func GetListAutoExecutionsSortByEnumValues() []ListAutoExecutionsSortByEnum {
	values := make([]ListAutoExecutionsSortByEnum, 0)
	for _, v := range mappingListAutoExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutoExecutionsSortByEnumStringValues Enumerates the set of values in String for ListAutoExecutionsSortByEnum
func GetListAutoExecutionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeScheduled",
	}
}

// GetMappingListAutoExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutoExecutionsSortByEnum(val string) (ListAutoExecutionsSortByEnum, bool) {
	enum, ok := mappingListAutoExecutionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutoExecutionsSortOrderEnum Enum with underlying type: string
type ListAutoExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListAutoExecutionsSortOrderEnum
const (
	ListAutoExecutionsSortOrderAsc  ListAutoExecutionsSortOrderEnum = "ASC"
	ListAutoExecutionsSortOrderDesc ListAutoExecutionsSortOrderEnum = "DESC"
)

var mappingListAutoExecutionsSortOrderEnum = map[string]ListAutoExecutionsSortOrderEnum{
	"ASC":  ListAutoExecutionsSortOrderAsc,
	"DESC": ListAutoExecutionsSortOrderDesc,
}

var mappingListAutoExecutionsSortOrderEnumLowerCase = map[string]ListAutoExecutionsSortOrderEnum{
	"asc":  ListAutoExecutionsSortOrderAsc,
	"desc": ListAutoExecutionsSortOrderDesc,
}

// GetListAutoExecutionsSortOrderEnumValues Enumerates the set of values for ListAutoExecutionsSortOrderEnum
func GetListAutoExecutionsSortOrderEnumValues() []ListAutoExecutionsSortOrderEnum {
	values := make([]ListAutoExecutionsSortOrderEnum, 0)
	for _, v := range mappingListAutoExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutoExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListAutoExecutionsSortOrderEnum
func GetListAutoExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutoExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutoExecutionsSortOrderEnum(val string) (ListAutoExecutionsSortOrderEnum, bool) {
	enum, ok := mappingListAutoExecutionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
