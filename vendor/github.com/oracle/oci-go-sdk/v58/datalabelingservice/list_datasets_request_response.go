// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datalabelingservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDatasetsRequest wrapper for the ListDatasets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datalabelingservice/ListDatasets.go.html to see an example of how to use ListDatasetsRequest.
type ListDatasetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Dataset OCID
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire annotation format given.
	AnnotationFormat *string `mandatory:"false" contributesTo:"query" name:"annotationFormat"`

	// A filter to return only resources whose lifecycleState matches this query param.
	LifecycleState DatasetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatasetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatasetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatasetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatasetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatasetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatasetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatasetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatasetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDatasetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatasetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatasetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatasetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatasetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatasetsResponse wrapper for the ListDatasets operation
type ListDatasetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatasetCollection instances
	DatasetCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For the pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatasetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatasetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatasetsSortOrderEnum Enum with underlying type: string
type ListDatasetsSortOrderEnum string

// Set of constants representing the allowable values for ListDatasetsSortOrderEnum
const (
	ListDatasetsSortOrderAsc  ListDatasetsSortOrderEnum = "ASC"
	ListDatasetsSortOrderDesc ListDatasetsSortOrderEnum = "DESC"
)

var mappingListDatasetsSortOrderEnum = map[string]ListDatasetsSortOrderEnum{
	"ASC":  ListDatasetsSortOrderAsc,
	"DESC": ListDatasetsSortOrderDesc,
}

// GetListDatasetsSortOrderEnumValues Enumerates the set of values for ListDatasetsSortOrderEnum
func GetListDatasetsSortOrderEnumValues() []ListDatasetsSortOrderEnum {
	values := make([]ListDatasetsSortOrderEnum, 0)
	for _, v := range mappingListDatasetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatasetsSortOrderEnumStringValues Enumerates the set of values in String for ListDatasetsSortOrderEnum
func GetListDatasetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatasetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatasetsSortOrderEnum(val string) (ListDatasetsSortOrderEnum, bool) {
	mappingListDatasetsSortOrderEnumIgnoreCase := make(map[string]ListDatasetsSortOrderEnum)
	for k, v := range mappingListDatasetsSortOrderEnum {
		mappingListDatasetsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDatasetsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatasetsSortByEnum Enum with underlying type: string
type ListDatasetsSortByEnum string

// Set of constants representing the allowable values for ListDatasetsSortByEnum
const (
	ListDatasetsSortByTimecreated ListDatasetsSortByEnum = "timeCreated"
	ListDatasetsSortByDisplayname ListDatasetsSortByEnum = "displayName"
)

var mappingListDatasetsSortByEnum = map[string]ListDatasetsSortByEnum{
	"timeCreated": ListDatasetsSortByTimecreated,
	"displayName": ListDatasetsSortByDisplayname,
}

// GetListDatasetsSortByEnumValues Enumerates the set of values for ListDatasetsSortByEnum
func GetListDatasetsSortByEnumValues() []ListDatasetsSortByEnum {
	values := make([]ListDatasetsSortByEnum, 0)
	for _, v := range mappingListDatasetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatasetsSortByEnumStringValues Enumerates the set of values in String for ListDatasetsSortByEnum
func GetListDatasetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDatasetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatasetsSortByEnum(val string) (ListDatasetsSortByEnum, bool) {
	mappingListDatasetsSortByEnumIgnoreCase := make(map[string]ListDatasetsSortByEnum)
	for k, v := range mappingListDatasetsSortByEnum {
		mappingListDatasetsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDatasetsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
