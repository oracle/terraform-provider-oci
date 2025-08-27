// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListKafkaClusterConfigsRequest wrapper for the ListKafkaClusterConfigs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managedkafka/ListKafkaClusterConfigs.go.html to see an example of how to use ListKafkaClusterConfigsRequest.
type ListKafkaClusterConfigsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState KafkaClusterConfigLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaClusterConfig.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListKafkaClusterConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListKafkaClusterConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKafkaClusterConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKafkaClusterConfigsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListKafkaClusterConfigsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKafkaClusterConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListKafkaClusterConfigsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaClusterConfigLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetKafkaClusterConfigLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListKafkaClusterConfigsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListKafkaClusterConfigsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListKafkaClusterConfigsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListKafkaClusterConfigsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListKafkaClusterConfigsResponse wrapper for the ListKafkaClusterConfigs operation
type ListKafkaClusterConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of KafkaClusterConfigCollection instances
	KafkaClusterConfigCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListKafkaClusterConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKafkaClusterConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKafkaClusterConfigsSortOrderEnum Enum with underlying type: string
type ListKafkaClusterConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListKafkaClusterConfigsSortOrderEnum
const (
	ListKafkaClusterConfigsSortOrderAsc  ListKafkaClusterConfigsSortOrderEnum = "ASC"
	ListKafkaClusterConfigsSortOrderDesc ListKafkaClusterConfigsSortOrderEnum = "DESC"
)

var mappingListKafkaClusterConfigsSortOrderEnum = map[string]ListKafkaClusterConfigsSortOrderEnum{
	"ASC":  ListKafkaClusterConfigsSortOrderAsc,
	"DESC": ListKafkaClusterConfigsSortOrderDesc,
}

var mappingListKafkaClusterConfigsSortOrderEnumLowerCase = map[string]ListKafkaClusterConfigsSortOrderEnum{
	"asc":  ListKafkaClusterConfigsSortOrderAsc,
	"desc": ListKafkaClusterConfigsSortOrderDesc,
}

// GetListKafkaClusterConfigsSortOrderEnumValues Enumerates the set of values for ListKafkaClusterConfigsSortOrderEnum
func GetListKafkaClusterConfigsSortOrderEnumValues() []ListKafkaClusterConfigsSortOrderEnum {
	values := make([]ListKafkaClusterConfigsSortOrderEnum, 0)
	for _, v := range mappingListKafkaClusterConfigsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListKafkaClusterConfigsSortOrderEnumStringValues Enumerates the set of values in String for ListKafkaClusterConfigsSortOrderEnum
func GetListKafkaClusterConfigsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListKafkaClusterConfigsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKafkaClusterConfigsSortOrderEnum(val string) (ListKafkaClusterConfigsSortOrderEnum, bool) {
	enum, ok := mappingListKafkaClusterConfigsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListKafkaClusterConfigsSortByEnum Enum with underlying type: string
type ListKafkaClusterConfigsSortByEnum string

// Set of constants representing the allowable values for ListKafkaClusterConfigsSortByEnum
const (
	ListKafkaClusterConfigsSortByTimecreated ListKafkaClusterConfigsSortByEnum = "timeCreated"
	ListKafkaClusterConfigsSortByDisplayname ListKafkaClusterConfigsSortByEnum = "displayName"
)

var mappingListKafkaClusterConfigsSortByEnum = map[string]ListKafkaClusterConfigsSortByEnum{
	"timeCreated": ListKafkaClusterConfigsSortByTimecreated,
	"displayName": ListKafkaClusterConfigsSortByDisplayname,
}

var mappingListKafkaClusterConfigsSortByEnumLowerCase = map[string]ListKafkaClusterConfigsSortByEnum{
	"timecreated": ListKafkaClusterConfigsSortByTimecreated,
	"displayname": ListKafkaClusterConfigsSortByDisplayname,
}

// GetListKafkaClusterConfigsSortByEnumValues Enumerates the set of values for ListKafkaClusterConfigsSortByEnum
func GetListKafkaClusterConfigsSortByEnumValues() []ListKafkaClusterConfigsSortByEnum {
	values := make([]ListKafkaClusterConfigsSortByEnum, 0)
	for _, v := range mappingListKafkaClusterConfigsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListKafkaClusterConfigsSortByEnumStringValues Enumerates the set of values in String for ListKafkaClusterConfigsSortByEnum
func GetListKafkaClusterConfigsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListKafkaClusterConfigsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKafkaClusterConfigsSortByEnum(val string) (ListKafkaClusterConfigsSortByEnum, bool) {
	enum, ok := mappingListKafkaClusterConfigsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
