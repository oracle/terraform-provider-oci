// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProblemEntitiesRequest wrapper for the ListProblemEntities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListProblemEntities.go.html to see an example of how to use ListProblemEntitiesRequest.
type ListProblemEntitiesRequest struct {

	// OCID of the problem.
	ProblemId *string `mandatory:"true" contributesTo:"path" name:"problemId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use
	SortOrder ListProblemEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListProblemEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProblemEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProblemEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProblemEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProblemEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProblemEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProblemEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProblemEntitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProblemEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProblemEntitiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProblemEntitiesResponse wrapper for the ListProblemEntities operation
type ListProblemEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProblemEntityCollection instances
	ProblemEntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProblemEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProblemEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProblemEntitiesSortOrderEnum Enum with underlying type: string
type ListProblemEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListProblemEntitiesSortOrderEnum
const (
	ListProblemEntitiesSortOrderAsc  ListProblemEntitiesSortOrderEnum = "ASC"
	ListProblemEntitiesSortOrderDesc ListProblemEntitiesSortOrderEnum = "DESC"
)

var mappingListProblemEntitiesSortOrderEnum = map[string]ListProblemEntitiesSortOrderEnum{
	"ASC":  ListProblemEntitiesSortOrderAsc,
	"DESC": ListProblemEntitiesSortOrderDesc,
}

var mappingListProblemEntitiesSortOrderEnumLowerCase = map[string]ListProblemEntitiesSortOrderEnum{
	"asc":  ListProblemEntitiesSortOrderAsc,
	"desc": ListProblemEntitiesSortOrderDesc,
}

// GetListProblemEntitiesSortOrderEnumValues Enumerates the set of values for ListProblemEntitiesSortOrderEnum
func GetListProblemEntitiesSortOrderEnumValues() []ListProblemEntitiesSortOrderEnum {
	values := make([]ListProblemEntitiesSortOrderEnum, 0)
	for _, v := range mappingListProblemEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListProblemEntitiesSortOrderEnum
func GetListProblemEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProblemEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemEntitiesSortOrderEnum(val string) (ListProblemEntitiesSortOrderEnum, bool) {
	enum, ok := mappingListProblemEntitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProblemEntitiesSortByEnum Enum with underlying type: string
type ListProblemEntitiesSortByEnum string

// Set of constants representing the allowable values for ListProblemEntitiesSortByEnum
const (
	ListProblemEntitiesSortByTimecreated ListProblemEntitiesSortByEnum = "timeCreated"
)

var mappingListProblemEntitiesSortByEnum = map[string]ListProblemEntitiesSortByEnum{
	"timeCreated": ListProblemEntitiesSortByTimecreated,
}

var mappingListProblemEntitiesSortByEnumLowerCase = map[string]ListProblemEntitiesSortByEnum{
	"timecreated": ListProblemEntitiesSortByTimecreated,
}

// GetListProblemEntitiesSortByEnumValues Enumerates the set of values for ListProblemEntitiesSortByEnum
func GetListProblemEntitiesSortByEnumValues() []ListProblemEntitiesSortByEnum {
	values := make([]ListProblemEntitiesSortByEnum, 0)
	for _, v := range mappingListProblemEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemEntitiesSortByEnumStringValues Enumerates the set of values in String for ListProblemEntitiesSortByEnum
func GetListProblemEntitiesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListProblemEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProblemEntitiesSortByEnum(val string) (ListProblemEntitiesSortByEnum, bool) {
	enum, ok := mappingListProblemEntitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
