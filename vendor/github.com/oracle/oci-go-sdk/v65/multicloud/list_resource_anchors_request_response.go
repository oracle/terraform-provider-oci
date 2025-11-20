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

// ListResourceAnchorsRequest wrapper for the ListResourceAnchors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListResourceAnchors.go.html to see an example of how to use ListResourceAnchorsRequest.
type ListResourceAnchorsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud base compartment or sub-compartment in which to list resources.
	// A Multicloud base compartment is an OCI compartment that maps to a subscription in a Cloud Service Provider (such as Azure, AWS, or Google Cloud).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment linked to the resource.
	LinkedCompartmentId *string `mandatory:"false" contributesTo:"query" name:"linkedCompartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ResourceAnchorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnchor.
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
	SortOrder ListResourceAnchorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListResourceAnchorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Check the sub-compartments of a given compartmentId
	IsCompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isCompartmentIdInSubtree"`

	// Whether to fetch and include the compartment name, setting this field to yes may introduce additional latency.
	ShouldFetchCompartmentName *bool `mandatory:"false" contributesTo:"query" name:"shouldFetchCompartmentName"`

	// The subscription service name of the Cloud Service Provider.
	SubscriptionServiceName ListResourceAnchorsSubscriptionServiceNameEnum `mandatory:"false" contributesTo:"query" name:"subscriptionServiceName" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
	SubscriptionId *string `mandatory:"false" contributesTo:"query" name:"subscriptionId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceAnchorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceAnchorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceAnchorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceAnchorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceAnchorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceAnchorLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetResourceAnchorLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceAnchorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceAnchorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceAnchorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceAnchorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceAnchorsSubscriptionServiceNameEnum(string(request.SubscriptionServiceName)); !ok && request.SubscriptionServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionServiceName: %s. Supported values are: %s.", request.SubscriptionServiceName, strings.Join(GetListResourceAnchorsSubscriptionServiceNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceAnchorsResponse wrapper for the ListResourceAnchors operation
type ListResourceAnchorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceAnchorCollection instances
	ResourceAnchorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceAnchorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceAnchorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceAnchorsSortOrderEnum Enum with underlying type: string
type ListResourceAnchorsSortOrderEnum string

// Set of constants representing the allowable values for ListResourceAnchorsSortOrderEnum
const (
	ListResourceAnchorsSortOrderAsc  ListResourceAnchorsSortOrderEnum = "ASC"
	ListResourceAnchorsSortOrderDesc ListResourceAnchorsSortOrderEnum = "DESC"
)

var mappingListResourceAnchorsSortOrderEnum = map[string]ListResourceAnchorsSortOrderEnum{
	"ASC":  ListResourceAnchorsSortOrderAsc,
	"DESC": ListResourceAnchorsSortOrderDesc,
}

var mappingListResourceAnchorsSortOrderEnumLowerCase = map[string]ListResourceAnchorsSortOrderEnum{
	"asc":  ListResourceAnchorsSortOrderAsc,
	"desc": ListResourceAnchorsSortOrderDesc,
}

// GetListResourceAnchorsSortOrderEnumValues Enumerates the set of values for ListResourceAnchorsSortOrderEnum
func GetListResourceAnchorsSortOrderEnumValues() []ListResourceAnchorsSortOrderEnum {
	values := make([]ListResourceAnchorsSortOrderEnum, 0)
	for _, v := range mappingListResourceAnchorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceAnchorsSortOrderEnumStringValues Enumerates the set of values in String for ListResourceAnchorsSortOrderEnum
func GetListResourceAnchorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceAnchorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceAnchorsSortOrderEnum(val string) (ListResourceAnchorsSortOrderEnum, bool) {
	enum, ok := mappingListResourceAnchorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceAnchorsSortByEnum Enum with underlying type: string
type ListResourceAnchorsSortByEnum string

// Set of constants representing the allowable values for ListResourceAnchorsSortByEnum
const (
	ListResourceAnchorsSortByTimecreated ListResourceAnchorsSortByEnum = "timeCreated"
	ListResourceAnchorsSortByDisplayname ListResourceAnchorsSortByEnum = "displayName"
)

var mappingListResourceAnchorsSortByEnum = map[string]ListResourceAnchorsSortByEnum{
	"timeCreated": ListResourceAnchorsSortByTimecreated,
	"displayName": ListResourceAnchorsSortByDisplayname,
}

var mappingListResourceAnchorsSortByEnumLowerCase = map[string]ListResourceAnchorsSortByEnum{
	"timecreated": ListResourceAnchorsSortByTimecreated,
	"displayname": ListResourceAnchorsSortByDisplayname,
}

// GetListResourceAnchorsSortByEnumValues Enumerates the set of values for ListResourceAnchorsSortByEnum
func GetListResourceAnchorsSortByEnumValues() []ListResourceAnchorsSortByEnum {
	values := make([]ListResourceAnchorsSortByEnum, 0)
	for _, v := range mappingListResourceAnchorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceAnchorsSortByEnumStringValues Enumerates the set of values in String for ListResourceAnchorsSortByEnum
func GetListResourceAnchorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListResourceAnchorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceAnchorsSortByEnum(val string) (ListResourceAnchorsSortByEnum, bool) {
	enum, ok := mappingListResourceAnchorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceAnchorsSubscriptionServiceNameEnum Enum with underlying type: string
type ListResourceAnchorsSubscriptionServiceNameEnum string

// Set of constants representing the allowable values for ListResourceAnchorsSubscriptionServiceNameEnum
const (
	ListResourceAnchorsSubscriptionServiceNameOracledbatazure  ListResourceAnchorsSubscriptionServiceNameEnum = "ORACLEDBATAZURE"
	ListResourceAnchorsSubscriptionServiceNameOracledbatgoogle ListResourceAnchorsSubscriptionServiceNameEnum = "ORACLEDBATGOOGLE"
	ListResourceAnchorsSubscriptionServiceNameOracledbataws    ListResourceAnchorsSubscriptionServiceNameEnum = "ORACLEDBATAWS"
)

var mappingListResourceAnchorsSubscriptionServiceNameEnum = map[string]ListResourceAnchorsSubscriptionServiceNameEnum{
	"ORACLEDBATAZURE":  ListResourceAnchorsSubscriptionServiceNameOracledbatazure,
	"ORACLEDBATGOOGLE": ListResourceAnchorsSubscriptionServiceNameOracledbatgoogle,
	"ORACLEDBATAWS":    ListResourceAnchorsSubscriptionServiceNameOracledbataws,
}

var mappingListResourceAnchorsSubscriptionServiceNameEnumLowerCase = map[string]ListResourceAnchorsSubscriptionServiceNameEnum{
	"oracledbatazure":  ListResourceAnchorsSubscriptionServiceNameOracledbatazure,
	"oracledbatgoogle": ListResourceAnchorsSubscriptionServiceNameOracledbatgoogle,
	"oracledbataws":    ListResourceAnchorsSubscriptionServiceNameOracledbataws,
}

// GetListResourceAnchorsSubscriptionServiceNameEnumValues Enumerates the set of values for ListResourceAnchorsSubscriptionServiceNameEnum
func GetListResourceAnchorsSubscriptionServiceNameEnumValues() []ListResourceAnchorsSubscriptionServiceNameEnum {
	values := make([]ListResourceAnchorsSubscriptionServiceNameEnum, 0)
	for _, v := range mappingListResourceAnchorsSubscriptionServiceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceAnchorsSubscriptionServiceNameEnumStringValues Enumerates the set of values in String for ListResourceAnchorsSubscriptionServiceNameEnum
func GetListResourceAnchorsSubscriptionServiceNameEnumStringValues() []string {
	return []string{
		"ORACLEDBATAZURE",
		"ORACLEDBATGOOGLE",
		"ORACLEDBATAWS",
	}
}

// GetMappingListResourceAnchorsSubscriptionServiceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceAnchorsSubscriptionServiceNameEnum(val string) (ListResourceAnchorsSubscriptionServiceNameEnum, bool) {
	enum, ok := mappingListResourceAnchorsSubscriptionServiceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
