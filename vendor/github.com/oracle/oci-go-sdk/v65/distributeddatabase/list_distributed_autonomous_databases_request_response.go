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

// ListDistributedAutonomousDatabasesRequest wrapper for the ListDistributedAutonomousDatabases operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ListDistributedAutonomousDatabases.go.html to see an example of how to use ListDistributedAutonomousDatabasesRequest.
type ListDistributedAutonomousDatabasesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that are associated with the given privateEndpointId.
	PrivateEndpointId *string `mandatory:"false" contributesTo:"query" name:"privateEndpointId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState DistributedAutonomousDatabaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDistributedAutonomousDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListDistributedAutonomousDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only Globally distributed autonomous databases that match the entire name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their dbDeploymentType matches the given dbDeploymentType.
	DbDeploymentType DistributedAutonomousDatabaseDbDeploymentTypeEnum `mandatory:"false" contributesTo:"query" name:"dbDeploymentType" omitEmpty:"true"`

	// Comma separated names of argument corresponding to which metadata need to be retrived.
	Metadata *string `mandatory:"false" contributesTo:"query" name:"metadata"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDistributedAutonomousDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDistributedAutonomousDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDistributedAutonomousDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDistributedAutonomousDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDistributedAutonomousDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedAutonomousDatabaseLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDistributedAutonomousDatabaseLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDistributedAutonomousDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDistributedAutonomousDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDistributedAutonomousDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDistributedAutonomousDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedAutonomousDatabaseDbDeploymentTypeEnum(string(request.DbDeploymentType)); !ok && request.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", request.DbDeploymentType, strings.Join(GetDistributedAutonomousDatabaseDbDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDistributedAutonomousDatabasesResponse wrapper for the ListDistributedAutonomousDatabases operation
type ListDistributedAutonomousDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DistributedAutonomousDatabaseCollection instances
	DistributedAutonomousDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDistributedAutonomousDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDistributedAutonomousDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDistributedAutonomousDatabasesSortOrderEnum Enum with underlying type: string
type ListDistributedAutonomousDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListDistributedAutonomousDatabasesSortOrderEnum
const (
	ListDistributedAutonomousDatabasesSortOrderAsc  ListDistributedAutonomousDatabasesSortOrderEnum = "ASC"
	ListDistributedAutonomousDatabasesSortOrderDesc ListDistributedAutonomousDatabasesSortOrderEnum = "DESC"
)

var mappingListDistributedAutonomousDatabasesSortOrderEnum = map[string]ListDistributedAutonomousDatabasesSortOrderEnum{
	"ASC":  ListDistributedAutonomousDatabasesSortOrderAsc,
	"DESC": ListDistributedAutonomousDatabasesSortOrderDesc,
}

var mappingListDistributedAutonomousDatabasesSortOrderEnumLowerCase = map[string]ListDistributedAutonomousDatabasesSortOrderEnum{
	"asc":  ListDistributedAutonomousDatabasesSortOrderAsc,
	"desc": ListDistributedAutonomousDatabasesSortOrderDesc,
}

// GetListDistributedAutonomousDatabasesSortOrderEnumValues Enumerates the set of values for ListDistributedAutonomousDatabasesSortOrderEnum
func GetListDistributedAutonomousDatabasesSortOrderEnumValues() []ListDistributedAutonomousDatabasesSortOrderEnum {
	values := make([]ListDistributedAutonomousDatabasesSortOrderEnum, 0)
	for _, v := range mappingListDistributedAutonomousDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDistributedAutonomousDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListDistributedAutonomousDatabasesSortOrderEnum
func GetListDistributedAutonomousDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDistributedAutonomousDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDistributedAutonomousDatabasesSortOrderEnum(val string) (ListDistributedAutonomousDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListDistributedAutonomousDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDistributedAutonomousDatabasesSortByEnum Enum with underlying type: string
type ListDistributedAutonomousDatabasesSortByEnum string

// Set of constants representing the allowable values for ListDistributedAutonomousDatabasesSortByEnum
const (
	ListDistributedAutonomousDatabasesSortByTimecreated ListDistributedAutonomousDatabasesSortByEnum = "timeCreated"
	ListDistributedAutonomousDatabasesSortByTimeupdated ListDistributedAutonomousDatabasesSortByEnum = "timeUpdated"
)

var mappingListDistributedAutonomousDatabasesSortByEnum = map[string]ListDistributedAutonomousDatabasesSortByEnum{
	"timeCreated": ListDistributedAutonomousDatabasesSortByTimecreated,
	"timeUpdated": ListDistributedAutonomousDatabasesSortByTimeupdated,
}

var mappingListDistributedAutonomousDatabasesSortByEnumLowerCase = map[string]ListDistributedAutonomousDatabasesSortByEnum{
	"timecreated": ListDistributedAutonomousDatabasesSortByTimecreated,
	"timeupdated": ListDistributedAutonomousDatabasesSortByTimeupdated,
}

// GetListDistributedAutonomousDatabasesSortByEnumValues Enumerates the set of values for ListDistributedAutonomousDatabasesSortByEnum
func GetListDistributedAutonomousDatabasesSortByEnumValues() []ListDistributedAutonomousDatabasesSortByEnum {
	values := make([]ListDistributedAutonomousDatabasesSortByEnum, 0)
	for _, v := range mappingListDistributedAutonomousDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDistributedAutonomousDatabasesSortByEnumStringValues Enumerates the set of values in String for ListDistributedAutonomousDatabasesSortByEnum
func GetListDistributedAutonomousDatabasesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListDistributedAutonomousDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDistributedAutonomousDatabasesSortByEnum(val string) (ListDistributedAutonomousDatabasesSortByEnum, bool) {
	enum, ok := mappingListDistributedAutonomousDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
