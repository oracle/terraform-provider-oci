// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOpensearchClustersRequest wrapper for the ListOpensearchClusters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/ListOpensearchClusters.go.html to see an example of how to use ListOpensearchClustersRequest.
type ListOpensearchClustersRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only OpensearchClusters their lifecycleState matches the given lifecycleState.
	LifecycleState OpensearchClusterLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique OpensearchCluster identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOpensearchClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOpensearchClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOpensearchClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOpensearchClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOpensearchClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOpensearchClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOpensearchClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpensearchClusterLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOpensearchClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpensearchClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOpensearchClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpensearchClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOpensearchClustersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOpensearchClustersResponse wrapper for the ListOpensearchClusters operation
type ListOpensearchClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OpensearchClusterCollection instances
	OpensearchClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOpensearchClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOpensearchClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOpensearchClustersSortOrderEnum Enum with underlying type: string
type ListOpensearchClustersSortOrderEnum string

// Set of constants representing the allowable values for ListOpensearchClustersSortOrderEnum
const (
	ListOpensearchClustersSortOrderAsc  ListOpensearchClustersSortOrderEnum = "ASC"
	ListOpensearchClustersSortOrderDesc ListOpensearchClustersSortOrderEnum = "DESC"
)

var mappingListOpensearchClustersSortOrderEnum = map[string]ListOpensearchClustersSortOrderEnum{
	"ASC":  ListOpensearchClustersSortOrderAsc,
	"DESC": ListOpensearchClustersSortOrderDesc,
}

var mappingListOpensearchClustersSortOrderEnumLowerCase = map[string]ListOpensearchClustersSortOrderEnum{
	"asc":  ListOpensearchClustersSortOrderAsc,
	"desc": ListOpensearchClustersSortOrderDesc,
}

// GetListOpensearchClustersSortOrderEnumValues Enumerates the set of values for ListOpensearchClustersSortOrderEnum
func GetListOpensearchClustersSortOrderEnumValues() []ListOpensearchClustersSortOrderEnum {
	values := make([]ListOpensearchClustersSortOrderEnum, 0)
	for _, v := range mappingListOpensearchClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpensearchClustersSortOrderEnumStringValues Enumerates the set of values in String for ListOpensearchClustersSortOrderEnum
func GetListOpensearchClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOpensearchClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpensearchClustersSortOrderEnum(val string) (ListOpensearchClustersSortOrderEnum, bool) {
	enum, ok := mappingListOpensearchClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOpensearchClustersSortByEnum Enum with underlying type: string
type ListOpensearchClustersSortByEnum string

// Set of constants representing the allowable values for ListOpensearchClustersSortByEnum
const (
	ListOpensearchClustersSortByTimecreated ListOpensearchClustersSortByEnum = "timeCreated"
	ListOpensearchClustersSortByDisplayname ListOpensearchClustersSortByEnum = "displayName"
)

var mappingListOpensearchClustersSortByEnum = map[string]ListOpensearchClustersSortByEnum{
	"timeCreated": ListOpensearchClustersSortByTimecreated,
	"displayName": ListOpensearchClustersSortByDisplayname,
}

var mappingListOpensearchClustersSortByEnumLowerCase = map[string]ListOpensearchClustersSortByEnum{
	"timecreated": ListOpensearchClustersSortByTimecreated,
	"displayname": ListOpensearchClustersSortByDisplayname,
}

// GetListOpensearchClustersSortByEnumValues Enumerates the set of values for ListOpensearchClustersSortByEnum
func GetListOpensearchClustersSortByEnumValues() []ListOpensearchClustersSortByEnum {
	values := make([]ListOpensearchClustersSortByEnum, 0)
	for _, v := range mappingListOpensearchClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpensearchClustersSortByEnumStringValues Enumerates the set of values in String for ListOpensearchClustersSortByEnum
func GetListOpensearchClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOpensearchClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpensearchClustersSortByEnum(val string) (ListOpensearchClustersSortByEnum, bool) {
	enum, ok := mappingListOpensearchClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
