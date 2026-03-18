// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDistributedDatabasesRequest wrapper for the ListDistributedDatabases operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ListDistributedDatabases.go.html to see an example of how to use ListDistributedDatabasesRequest.
type ListDistributedDatabasesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that are associated with the given privateEndpointId.
	PrivateEndpointId *string `mandatory:"false" contributesTo:"query" name:"privateEndpointId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState DistributedDatabaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDistributedDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDistributedDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only Globally distributed databases that match the entire name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their dbDeploymentType matches the given dbDeploymentType.
	DbDeploymentType DistributedDatabaseDbDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"dbDeploymentType" omitEmpty:"true"`

	// Comma separated names of argument corresponding to which metadata need to be retrived.
	Metadata *string `mandatory:"false" contributesTo:"query" name:"metadata"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDistributedDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDistributedDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDistributedDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDistributedDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDistributedDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDistributedDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDistributedDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDistributedDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDistributedDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDistributedDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDatabaseDbDeploymentTypeEnum(string(request.DbDeploymentType)); !ok && request.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", request.DbDeploymentType, strings.Join(GetDistributedDatabaseDbDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDistributedDatabasesResponse wrapper for the ListDistributedDatabases operation
type ListDistributedDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DistributedDatabaseCollection instances
	DistributedDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDistributedDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDistributedDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDistributedDatabasesSortOrderEnum Enum with underlying type: string
type ListDistributedDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListDistributedDatabasesSortOrderEnum
const (
	ListDistributedDatabasesSortOrderAsc  ListDistributedDatabasesSortOrderEnum = "ASC"
	ListDistributedDatabasesSortOrderDesc ListDistributedDatabasesSortOrderEnum = "DESC"
)

var mappingListDistributedDatabasesSortOrderEnum = map[string]ListDistributedDatabasesSortOrderEnum{
	"ASC":  ListDistributedDatabasesSortOrderAsc,
	"DESC": ListDistributedDatabasesSortOrderDesc,
}

var mappingListDistributedDatabasesSortOrderEnumLowerCase = map[string]ListDistributedDatabasesSortOrderEnum{
	"asc":  ListDistributedDatabasesSortOrderAsc,
	"desc": ListDistributedDatabasesSortOrderDesc,
}

// GetListDistributedDatabasesSortOrderEnumValues Enumerates the set of values for ListDistributedDatabasesSortOrderEnum
func GetListDistributedDatabasesSortOrderEnumValues() []ListDistributedDatabasesSortOrderEnum {
	values := make([]ListDistributedDatabasesSortOrderEnum, 0)
	for _, v := range mappingListDistributedDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDistributedDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListDistributedDatabasesSortOrderEnum
func GetListDistributedDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDistributedDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDistributedDatabasesSortOrderEnum(val string) (ListDistributedDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListDistributedDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDistributedDatabasesSortByEnum Enum with underlying type: string
type ListDistributedDatabasesSortByEnum string

// Set of constants representing the allowable values for ListDistributedDatabasesSortByEnum
const (
	ListDistributedDatabasesSortByTimecreated ListDistributedDatabasesSortByEnum = "timeCreated"
	ListDistributedDatabasesSortByTimeupdated ListDistributedDatabasesSortByEnum = "timeUpdated"
)

var mappingListDistributedDatabasesSortByEnum = map[string]ListDistributedDatabasesSortByEnum{
	"timeCreated": ListDistributedDatabasesSortByTimecreated,
	"timeUpdated": ListDistributedDatabasesSortByTimeupdated,
}

var mappingListDistributedDatabasesSortByEnumLowerCase = map[string]ListDistributedDatabasesSortByEnum{
	"timecreated": ListDistributedDatabasesSortByTimecreated,
	"timeupdated": ListDistributedDatabasesSortByTimeupdated,
}

// GetListDistributedDatabasesSortByEnumValues Enumerates the set of values for ListDistributedDatabasesSortByEnum
func GetListDistributedDatabasesSortByEnumValues() []ListDistributedDatabasesSortByEnum {
	values := make([]ListDistributedDatabasesSortByEnum, 0)
	for _, v := range mappingListDistributedDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDistributedDatabasesSortByEnumStringValues Enumerates the set of values in String for ListDistributedDatabasesSortByEnum
func GetListDistributedDatabasesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListDistributedDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDistributedDatabasesSortByEnum(val string) (ListDistributedDatabasesSortByEnum, bool) {
	enum, ok := mappingListDistributedDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
