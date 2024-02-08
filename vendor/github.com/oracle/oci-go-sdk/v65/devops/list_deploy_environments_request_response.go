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

// ListDeployEnvironmentsRequest wrapper for the ListDeployEnvironments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployEnvironments.go.html to see an example of how to use ListDeployEnvironmentsRequest.
type ListDeployEnvironmentsRequest struct {

	// unique project identifier
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only DeployEnvironments that matches the given lifecycleState.
	LifecycleState DeployEnvironmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListDeployEnvironmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListDeployEnvironmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeployEnvironmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeployEnvironmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeployEnvironmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeployEnvironmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeployEnvironmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeployEnvironmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDeployEnvironmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeployEnvironmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeployEnvironmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeployEnvironmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeployEnvironmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeployEnvironmentsResponse wrapper for the ListDeployEnvironments operation
type ListDeployEnvironmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeployEnvironmentCollection instances
	DeployEnvironmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeployEnvironmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeployEnvironmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeployEnvironmentsSortOrderEnum Enum with underlying type: string
type ListDeployEnvironmentsSortOrderEnum string

// Set of constants representing the allowable values for ListDeployEnvironmentsSortOrderEnum
const (
	ListDeployEnvironmentsSortOrderAsc  ListDeployEnvironmentsSortOrderEnum = "ASC"
	ListDeployEnvironmentsSortOrderDesc ListDeployEnvironmentsSortOrderEnum = "DESC"
)

var mappingListDeployEnvironmentsSortOrderEnum = map[string]ListDeployEnvironmentsSortOrderEnum{
	"ASC":  ListDeployEnvironmentsSortOrderAsc,
	"DESC": ListDeployEnvironmentsSortOrderDesc,
}

var mappingListDeployEnvironmentsSortOrderEnumLowerCase = map[string]ListDeployEnvironmentsSortOrderEnum{
	"asc":  ListDeployEnvironmentsSortOrderAsc,
	"desc": ListDeployEnvironmentsSortOrderDesc,
}

// GetListDeployEnvironmentsSortOrderEnumValues Enumerates the set of values for ListDeployEnvironmentsSortOrderEnum
func GetListDeployEnvironmentsSortOrderEnumValues() []ListDeployEnvironmentsSortOrderEnum {
	values := make([]ListDeployEnvironmentsSortOrderEnum, 0)
	for _, v := range mappingListDeployEnvironmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeployEnvironmentsSortOrderEnumStringValues Enumerates the set of values in String for ListDeployEnvironmentsSortOrderEnum
func GetListDeployEnvironmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeployEnvironmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeployEnvironmentsSortOrderEnum(val string) (ListDeployEnvironmentsSortOrderEnum, bool) {
	enum, ok := mappingListDeployEnvironmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeployEnvironmentsSortByEnum Enum with underlying type: string
type ListDeployEnvironmentsSortByEnum string

// Set of constants representing the allowable values for ListDeployEnvironmentsSortByEnum
const (
	ListDeployEnvironmentsSortByTimecreated ListDeployEnvironmentsSortByEnum = "timeCreated"
	ListDeployEnvironmentsSortByDisplayname ListDeployEnvironmentsSortByEnum = "displayName"
)

var mappingListDeployEnvironmentsSortByEnum = map[string]ListDeployEnvironmentsSortByEnum{
	"timeCreated": ListDeployEnvironmentsSortByTimecreated,
	"displayName": ListDeployEnvironmentsSortByDisplayname,
}

var mappingListDeployEnvironmentsSortByEnumLowerCase = map[string]ListDeployEnvironmentsSortByEnum{
	"timecreated": ListDeployEnvironmentsSortByTimecreated,
	"displayname": ListDeployEnvironmentsSortByDisplayname,
}

// GetListDeployEnvironmentsSortByEnumValues Enumerates the set of values for ListDeployEnvironmentsSortByEnum
func GetListDeployEnvironmentsSortByEnumValues() []ListDeployEnvironmentsSortByEnum {
	values := make([]ListDeployEnvironmentsSortByEnum, 0)
	for _, v := range mappingListDeployEnvironmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeployEnvironmentsSortByEnumStringValues Enumerates the set of values in String for ListDeployEnvironmentsSortByEnum
func GetListDeployEnvironmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeployEnvironmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeployEnvironmentsSortByEnum(val string) (ListDeployEnvironmentsSortByEnum, bool) {
	enum, ok := mappingListDeployEnvironmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
