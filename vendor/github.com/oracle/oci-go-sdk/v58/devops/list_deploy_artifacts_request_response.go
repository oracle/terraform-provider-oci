// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDeployArtifactsRequest wrapper for the ListDeployArtifacts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployArtifacts.go.html to see an example of how to use ListDeployArtifactsRequest.
type ListDeployArtifactsRequest struct {

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// unique project identifier
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only DeployArtifacts that matches the given lifecycleState.
	LifecycleState DeployArtifactLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListDeployArtifactsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListDeployArtifactsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeployArtifactsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeployArtifactsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeployArtifactsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeployArtifactsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeployArtifactsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeployArtifactLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDeployArtifactLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeployArtifactsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeployArtifactsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeployArtifactsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeployArtifactsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeployArtifactsResponse wrapper for the ListDeployArtifacts operation
type ListDeployArtifactsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeployArtifactCollection instances
	DeployArtifactCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeployArtifactsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeployArtifactsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeployArtifactsSortOrderEnum Enum with underlying type: string
type ListDeployArtifactsSortOrderEnum string

// Set of constants representing the allowable values for ListDeployArtifactsSortOrderEnum
const (
	ListDeployArtifactsSortOrderAsc  ListDeployArtifactsSortOrderEnum = "ASC"
	ListDeployArtifactsSortOrderDesc ListDeployArtifactsSortOrderEnum = "DESC"
)

var mappingListDeployArtifactsSortOrderEnum = map[string]ListDeployArtifactsSortOrderEnum{
	"ASC":  ListDeployArtifactsSortOrderAsc,
	"DESC": ListDeployArtifactsSortOrderDesc,
}

// GetListDeployArtifactsSortOrderEnumValues Enumerates the set of values for ListDeployArtifactsSortOrderEnum
func GetListDeployArtifactsSortOrderEnumValues() []ListDeployArtifactsSortOrderEnum {
	values := make([]ListDeployArtifactsSortOrderEnum, 0)
	for _, v := range mappingListDeployArtifactsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeployArtifactsSortOrderEnumStringValues Enumerates the set of values in String for ListDeployArtifactsSortOrderEnum
func GetListDeployArtifactsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeployArtifactsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeployArtifactsSortOrderEnum(val string) (ListDeployArtifactsSortOrderEnum, bool) {
	mappingListDeployArtifactsSortOrderEnumIgnoreCase := make(map[string]ListDeployArtifactsSortOrderEnum)
	for k, v := range mappingListDeployArtifactsSortOrderEnum {
		mappingListDeployArtifactsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDeployArtifactsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeployArtifactsSortByEnum Enum with underlying type: string
type ListDeployArtifactsSortByEnum string

// Set of constants representing the allowable values for ListDeployArtifactsSortByEnum
const (
	ListDeployArtifactsSortByTimecreated ListDeployArtifactsSortByEnum = "timeCreated"
	ListDeployArtifactsSortByDisplayname ListDeployArtifactsSortByEnum = "displayName"
)

var mappingListDeployArtifactsSortByEnum = map[string]ListDeployArtifactsSortByEnum{
	"timeCreated": ListDeployArtifactsSortByTimecreated,
	"displayName": ListDeployArtifactsSortByDisplayname,
}

// GetListDeployArtifactsSortByEnumValues Enumerates the set of values for ListDeployArtifactsSortByEnum
func GetListDeployArtifactsSortByEnumValues() []ListDeployArtifactsSortByEnum {
	values := make([]ListDeployArtifactsSortByEnum, 0)
	for _, v := range mappingListDeployArtifactsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeployArtifactsSortByEnumStringValues Enumerates the set of values in String for ListDeployArtifactsSortByEnum
func GetListDeployArtifactsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeployArtifactsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeployArtifactsSortByEnum(val string) (ListDeployArtifactsSortByEnum, bool) {
	mappingListDeployArtifactsSortByEnumIgnoreCase := make(map[string]ListDeployArtifactsSortByEnum)
	for k, v := range mappingListDeployArtifactsSortByEnum {
		mappingListDeployArtifactsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDeployArtifactsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
