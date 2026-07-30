// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOracleDbAwsKeyPoolsRequest wrapper for the ListOracleDbAwsKeyPools operation
type ListOracleDbAwsKeyPoolsRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB AWS Key Pool Resource that match the given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given OCID](/Content/General/Concepts/identifiers.htm) of the Oracle DB AWS Key Pool resource.
	OracleDbAwsKeyPoolId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAwsKeyPoolId"`

	// A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState OracleDbAwsKeyPoolLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAwsKeyPoolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAwsKeyPoolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAwsKeyPoolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAwsKeyPoolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAwsKeyPoolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAwsKeyPoolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAwsKeyPoolsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAwsKeyPoolLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAwsKeyPoolLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAwsKeyPoolsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAwsKeyPoolsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAwsKeyPoolsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAwsKeyPoolsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAwsKeyPoolsResponse wrapper for the ListOracleDbAwsKeyPools operation
type ListOracleDbAwsKeyPoolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAwsKeyPoolCollection instances
	OracleDbAwsKeyPoolCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAwsKeyPoolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAwsKeyPoolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAwsKeyPoolsSortOrderEnum Enum with underlying type: string
type ListOracleDbAwsKeyPoolsSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAwsKeyPoolsSortOrderEnum
const (
	ListOracleDbAwsKeyPoolsSortOrderAsc  ListOracleDbAwsKeyPoolsSortOrderEnum = "ASC"
	ListOracleDbAwsKeyPoolsSortOrderDesc ListOracleDbAwsKeyPoolsSortOrderEnum = "DESC"
)

var mappingListOracleDbAwsKeyPoolsSortOrderEnum = map[string]ListOracleDbAwsKeyPoolsSortOrderEnum{
	"ASC":  ListOracleDbAwsKeyPoolsSortOrderAsc,
	"DESC": ListOracleDbAwsKeyPoolsSortOrderDesc,
}

var mappingListOracleDbAwsKeyPoolsSortOrderEnumLowerCase = map[string]ListOracleDbAwsKeyPoolsSortOrderEnum{
	"asc":  ListOracleDbAwsKeyPoolsSortOrderAsc,
	"desc": ListOracleDbAwsKeyPoolsSortOrderDesc,
}

// GetListOracleDbAwsKeyPoolsSortOrderEnumValues Enumerates the set of values for ListOracleDbAwsKeyPoolsSortOrderEnum
func GetListOracleDbAwsKeyPoolsSortOrderEnumValues() []ListOracleDbAwsKeyPoolsSortOrderEnum {
	values := make([]ListOracleDbAwsKeyPoolsSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAwsKeyPoolsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAwsKeyPoolsSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAwsKeyPoolsSortOrderEnum
func GetListOracleDbAwsKeyPoolsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAwsKeyPoolsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAwsKeyPoolsSortOrderEnum(val string) (ListOracleDbAwsKeyPoolsSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAwsKeyPoolsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAwsKeyPoolsSortByEnum Enum with underlying type: string
type ListOracleDbAwsKeyPoolsSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAwsKeyPoolsSortByEnum
const (
	ListOracleDbAwsKeyPoolsSortByTimecreated ListOracleDbAwsKeyPoolsSortByEnum = "timeCreated"
	ListOracleDbAwsKeyPoolsSortByDisplayname ListOracleDbAwsKeyPoolsSortByEnum = "displayName"
)

var mappingListOracleDbAwsKeyPoolsSortByEnum = map[string]ListOracleDbAwsKeyPoolsSortByEnum{
	"timeCreated": ListOracleDbAwsKeyPoolsSortByTimecreated,
	"displayName": ListOracleDbAwsKeyPoolsSortByDisplayname,
}

var mappingListOracleDbAwsKeyPoolsSortByEnumLowerCase = map[string]ListOracleDbAwsKeyPoolsSortByEnum{
	"timecreated": ListOracleDbAwsKeyPoolsSortByTimecreated,
	"displayname": ListOracleDbAwsKeyPoolsSortByDisplayname,
}

// GetListOracleDbAwsKeyPoolsSortByEnumValues Enumerates the set of values for ListOracleDbAwsKeyPoolsSortByEnum
func GetListOracleDbAwsKeyPoolsSortByEnumValues() []ListOracleDbAwsKeyPoolsSortByEnum {
	values := make([]ListOracleDbAwsKeyPoolsSortByEnum, 0)
	for _, v := range mappingListOracleDbAwsKeyPoolsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAwsKeyPoolsSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAwsKeyPoolsSortByEnum
func GetListOracleDbAwsKeyPoolsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAwsKeyPoolsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAwsKeyPoolsSortByEnum(val string) (ListOracleDbAwsKeyPoolsSortByEnum, bool) {
	enum, ok := mappingListOracleDbAwsKeyPoolsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
