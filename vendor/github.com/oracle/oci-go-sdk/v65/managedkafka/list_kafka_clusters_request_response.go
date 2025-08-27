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

// ListKafkaClustersRequest wrapper for the ListKafkaClusters operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managedkafka/ListKafkaClusters.go.html to see an example of how to use ListKafkaClustersRequest.
type ListKafkaClustersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState KafkaClusterLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.
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
	SortOrder ListKafkaClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListKafkaClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKafkaClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKafkaClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListKafkaClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKafkaClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListKafkaClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaClusterLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetKafkaClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListKafkaClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListKafkaClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListKafkaClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListKafkaClustersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListKafkaClustersResponse wrapper for the ListKafkaClusters operation
type ListKafkaClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of KafkaClusterCollection instances
	KafkaClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListKafkaClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKafkaClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKafkaClustersSortOrderEnum Enum with underlying type: string
type ListKafkaClustersSortOrderEnum string

// Set of constants representing the allowable values for ListKafkaClustersSortOrderEnum
const (
	ListKafkaClustersSortOrderAsc  ListKafkaClustersSortOrderEnum = "ASC"
	ListKafkaClustersSortOrderDesc ListKafkaClustersSortOrderEnum = "DESC"
)

var mappingListKafkaClustersSortOrderEnum = map[string]ListKafkaClustersSortOrderEnum{
	"ASC":  ListKafkaClustersSortOrderAsc,
	"DESC": ListKafkaClustersSortOrderDesc,
}

var mappingListKafkaClustersSortOrderEnumLowerCase = map[string]ListKafkaClustersSortOrderEnum{
	"asc":  ListKafkaClustersSortOrderAsc,
	"desc": ListKafkaClustersSortOrderDesc,
}

// GetListKafkaClustersSortOrderEnumValues Enumerates the set of values for ListKafkaClustersSortOrderEnum
func GetListKafkaClustersSortOrderEnumValues() []ListKafkaClustersSortOrderEnum {
	values := make([]ListKafkaClustersSortOrderEnum, 0)
	for _, v := range mappingListKafkaClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListKafkaClustersSortOrderEnumStringValues Enumerates the set of values in String for ListKafkaClustersSortOrderEnum
func GetListKafkaClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListKafkaClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKafkaClustersSortOrderEnum(val string) (ListKafkaClustersSortOrderEnum, bool) {
	enum, ok := mappingListKafkaClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListKafkaClustersSortByEnum Enum with underlying type: string
type ListKafkaClustersSortByEnum string

// Set of constants representing the allowable values for ListKafkaClustersSortByEnum
const (
	ListKafkaClustersSortByTimecreated ListKafkaClustersSortByEnum = "timeCreated"
	ListKafkaClustersSortByDisplayname ListKafkaClustersSortByEnum = "displayName"
)

var mappingListKafkaClustersSortByEnum = map[string]ListKafkaClustersSortByEnum{
	"timeCreated": ListKafkaClustersSortByTimecreated,
	"displayName": ListKafkaClustersSortByDisplayname,
}

var mappingListKafkaClustersSortByEnumLowerCase = map[string]ListKafkaClustersSortByEnum{
	"timecreated": ListKafkaClustersSortByTimecreated,
	"displayname": ListKafkaClustersSortByDisplayname,
}

// GetListKafkaClustersSortByEnumValues Enumerates the set of values for ListKafkaClustersSortByEnum
func GetListKafkaClustersSortByEnumValues() []ListKafkaClustersSortByEnum {
	values := make([]ListKafkaClustersSortByEnum, 0)
	for _, v := range mappingListKafkaClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListKafkaClustersSortByEnumStringValues Enumerates the set of values in String for ListKafkaClustersSortByEnum
func GetListKafkaClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListKafkaClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKafkaClustersSortByEnum(val string) (ListKafkaClustersSortByEnum, bool) {
	enum, ok := mappingListKafkaClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
