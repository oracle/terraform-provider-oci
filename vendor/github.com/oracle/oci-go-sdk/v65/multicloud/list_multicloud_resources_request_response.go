// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMulticloudResourcesRequest wrapper for the ListMulticloudResources operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListMulticloudResources.go.html to see an example of how to use ListMulticloudResourcesRequest.
type ListMulticloudResourcesRequest struct {

	// The subscription service name of the Cloud Service Provider.
	SubscriptionServiceName ListMulticloudResourcesSubscriptionServiceNameEnum `mandatory:"true" contributesTo:"query" name:"subscriptionServiceName" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
	ResourceAnchorId *string `mandatory:"false" contributesTo:"query" name:"resourceAnchorId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMulticloudResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListMulticloudResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The Cloud Service Provider region.
	ExternalLocation *string `mandatory:"false" contributesTo:"query" name:"externalLocation"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMulticloudResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMulticloudResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMulticloudResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMulticloudResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMulticloudResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMulticloudResourcesSubscriptionServiceNameEnum(string(request.SubscriptionServiceName)); !ok && request.SubscriptionServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionServiceName: %s. Supported values are: %s.", request.SubscriptionServiceName, strings.Join(GetListMulticloudResourcesSubscriptionServiceNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMulticloudResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMulticloudResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMulticloudResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMulticloudResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMulticloudResourcesResponse wrapper for the ListMulticloudResources operation
type ListMulticloudResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MulticloudResourceCollection instances
	MulticloudResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMulticloudResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMulticloudResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMulticloudResourcesSubscriptionServiceNameEnum Enum with underlying type: string
type ListMulticloudResourcesSubscriptionServiceNameEnum string

// Set of constants representing the allowable values for ListMulticloudResourcesSubscriptionServiceNameEnum
const (
	ListMulticloudResourcesSubscriptionServiceNameOracledbatazure  ListMulticloudResourcesSubscriptionServiceNameEnum = "ORACLEDBATAZURE"
	ListMulticloudResourcesSubscriptionServiceNameOracledbatgoogle ListMulticloudResourcesSubscriptionServiceNameEnum = "ORACLEDBATGOOGLE"
	ListMulticloudResourcesSubscriptionServiceNameOracledbataws    ListMulticloudResourcesSubscriptionServiceNameEnum = "ORACLEDBATAWS"
)

var mappingListMulticloudResourcesSubscriptionServiceNameEnum = map[string]ListMulticloudResourcesSubscriptionServiceNameEnum{
	"ORACLEDBATAZURE":  ListMulticloudResourcesSubscriptionServiceNameOracledbatazure,
	"ORACLEDBATGOOGLE": ListMulticloudResourcesSubscriptionServiceNameOracledbatgoogle,
	"ORACLEDBATAWS":    ListMulticloudResourcesSubscriptionServiceNameOracledbataws,
}

var mappingListMulticloudResourcesSubscriptionServiceNameEnumLowerCase = map[string]ListMulticloudResourcesSubscriptionServiceNameEnum{
	"oracledbatazure":  ListMulticloudResourcesSubscriptionServiceNameOracledbatazure,
	"oracledbatgoogle": ListMulticloudResourcesSubscriptionServiceNameOracledbatgoogle,
	"oracledbataws":    ListMulticloudResourcesSubscriptionServiceNameOracledbataws,
}

// GetListMulticloudResourcesSubscriptionServiceNameEnumValues Enumerates the set of values for ListMulticloudResourcesSubscriptionServiceNameEnum
func GetListMulticloudResourcesSubscriptionServiceNameEnumValues() []ListMulticloudResourcesSubscriptionServiceNameEnum {
	values := make([]ListMulticloudResourcesSubscriptionServiceNameEnum, 0)
	for _, v := range mappingListMulticloudResourcesSubscriptionServiceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListMulticloudResourcesSubscriptionServiceNameEnumStringValues Enumerates the set of values in String for ListMulticloudResourcesSubscriptionServiceNameEnum
func GetListMulticloudResourcesSubscriptionServiceNameEnumStringValues() []string {
	return []string{
		"ORACLEDBATAZURE",
		"ORACLEDBATGOOGLE",
		"ORACLEDBATAWS",
	}
}

// GetMappingListMulticloudResourcesSubscriptionServiceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMulticloudResourcesSubscriptionServiceNameEnum(val string) (ListMulticloudResourcesSubscriptionServiceNameEnum, bool) {
	enum, ok := mappingListMulticloudResourcesSubscriptionServiceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMulticloudResourcesSortOrderEnum Enum with underlying type: string
type ListMulticloudResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListMulticloudResourcesSortOrderEnum
const (
	ListMulticloudResourcesSortOrderAsc  ListMulticloudResourcesSortOrderEnum = "ASC"
	ListMulticloudResourcesSortOrderDesc ListMulticloudResourcesSortOrderEnum = "DESC"
)

var mappingListMulticloudResourcesSortOrderEnum = map[string]ListMulticloudResourcesSortOrderEnum{
	"ASC":  ListMulticloudResourcesSortOrderAsc,
	"DESC": ListMulticloudResourcesSortOrderDesc,
}

var mappingListMulticloudResourcesSortOrderEnumLowerCase = map[string]ListMulticloudResourcesSortOrderEnum{
	"asc":  ListMulticloudResourcesSortOrderAsc,
	"desc": ListMulticloudResourcesSortOrderDesc,
}

// GetListMulticloudResourcesSortOrderEnumValues Enumerates the set of values for ListMulticloudResourcesSortOrderEnum
func GetListMulticloudResourcesSortOrderEnumValues() []ListMulticloudResourcesSortOrderEnum {
	values := make([]ListMulticloudResourcesSortOrderEnum, 0)
	for _, v := range mappingListMulticloudResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMulticloudResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListMulticloudResourcesSortOrderEnum
func GetListMulticloudResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMulticloudResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMulticloudResourcesSortOrderEnum(val string) (ListMulticloudResourcesSortOrderEnum, bool) {
	enum, ok := mappingListMulticloudResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMulticloudResourcesSortByEnum Enum with underlying type: string
type ListMulticloudResourcesSortByEnum string

// Set of constants representing the allowable values for ListMulticloudResourcesSortByEnum
const (
	ListMulticloudResourcesSortByTimecreated ListMulticloudResourcesSortByEnum = "timeCreated"
	ListMulticloudResourcesSortByDisplayname ListMulticloudResourcesSortByEnum = "displayName"
)

var mappingListMulticloudResourcesSortByEnum = map[string]ListMulticloudResourcesSortByEnum{
	"timeCreated": ListMulticloudResourcesSortByTimecreated,
	"displayName": ListMulticloudResourcesSortByDisplayname,
}

var mappingListMulticloudResourcesSortByEnumLowerCase = map[string]ListMulticloudResourcesSortByEnum{
	"timecreated": ListMulticloudResourcesSortByTimecreated,
	"displayname": ListMulticloudResourcesSortByDisplayname,
}

// GetListMulticloudResourcesSortByEnumValues Enumerates the set of values for ListMulticloudResourcesSortByEnum
func GetListMulticloudResourcesSortByEnumValues() []ListMulticloudResourcesSortByEnum {
	values := make([]ListMulticloudResourcesSortByEnum, 0)
	for _, v := range mappingListMulticloudResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMulticloudResourcesSortByEnumStringValues Enumerates the set of values in String for ListMulticloudResourcesSortByEnum
func GetListMulticloudResourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMulticloudResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMulticloudResourcesSortByEnum(val string) (ListMulticloudResourcesSortByEnum, bool) {
	enum, ok := mappingListMulticloudResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
