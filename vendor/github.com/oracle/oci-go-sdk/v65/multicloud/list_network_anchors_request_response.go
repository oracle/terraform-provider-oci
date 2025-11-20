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

// ListNetworkAnchorsRequest wrapper for the ListNetworkAnchors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListNetworkAnchors.go.html to see an example of how to use ListNetworkAnchorsRequest.
type ListNetworkAnchorsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud base compartment or sub-compartment in which to list resources.
	// A Multicloud base compartment is an OCI compartment that maps to a subscription in a Cloud Service Provider (such as Azure, AWS, or Google Cloud).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
	SubscriptionId *string `mandatory:"false" contributesTo:"query" name:"subscriptionId"`

	// The subscription service name of the Cloud Service Provider.
	SubscriptionServiceName ListNetworkAnchorsSubscriptionServiceNameEnum `mandatory:"false" contributesTo:"query" name:"subscriptionServiceName" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	NetworkAnchorLifecycleState NetworkAnchorNetworkAnchorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"networkAnchorLifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The Cloud Service Provider region.
	ExternalLocation *string `mandatory:"false" contributesTo:"query" name:"externalLocation"`

	// A filter to return only NetworkAnchor resources that match the given OCI subnet Id.
	NetworkAnchorOciSubnetId *string `mandatory:"false" contributesTo:"query" name:"networkAnchorOciSubnetId"`

	// If set to true, a list operation will return NetworkAnchors from all child compartments in the provided compartmentId parameter.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// A filter to return only NetworkAnchor resources that match the given OCI Vcn Id.
	NetworkAnchorOciVcnId *string `mandatory:"false" contributesTo:"query" name:"networkAnchorOciVcnId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAnchor.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Whether to fetch and include the vcn display name, which may introduce additional latency.
	ShouldFetchVcnName *bool `mandatory:"false" contributesTo:"query" name:"shouldFetchVcnName"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListNetworkAnchorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListNetworkAnchorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkAnchorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkAnchorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkAnchorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkAnchorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkAnchorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkAnchorsSubscriptionServiceNameEnum(string(request.SubscriptionServiceName)); !ok && request.SubscriptionServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionServiceName: %s. Supported values are: %s.", request.SubscriptionServiceName, strings.Join(GetListNetworkAnchorsSubscriptionServiceNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkAnchorNetworkAnchorLifecycleStateEnum(string(request.NetworkAnchorLifecycleState)); !ok && request.NetworkAnchorLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkAnchorLifecycleState: %s. Supported values are: %s.", request.NetworkAnchorLifecycleState, strings.Join(GetNetworkAnchorNetworkAnchorLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkAnchorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkAnchorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkAnchorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkAnchorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkAnchorsResponse wrapper for the ListNetworkAnchors operation
type ListNetworkAnchorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NetworkAnchorCollection instances
	NetworkAnchorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNetworkAnchorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkAnchorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkAnchorsSubscriptionServiceNameEnum Enum with underlying type: string
type ListNetworkAnchorsSubscriptionServiceNameEnum string

// Set of constants representing the allowable values for ListNetworkAnchorsSubscriptionServiceNameEnum
const (
	ListNetworkAnchorsSubscriptionServiceNameOracledbatazure  ListNetworkAnchorsSubscriptionServiceNameEnum = "ORACLEDBATAZURE"
	ListNetworkAnchorsSubscriptionServiceNameOracledbatgoogle ListNetworkAnchorsSubscriptionServiceNameEnum = "ORACLEDBATGOOGLE"
	ListNetworkAnchorsSubscriptionServiceNameOracledbataws    ListNetworkAnchorsSubscriptionServiceNameEnum = "ORACLEDBATAWS"
)

var mappingListNetworkAnchorsSubscriptionServiceNameEnum = map[string]ListNetworkAnchorsSubscriptionServiceNameEnum{
	"ORACLEDBATAZURE":  ListNetworkAnchorsSubscriptionServiceNameOracledbatazure,
	"ORACLEDBATGOOGLE": ListNetworkAnchorsSubscriptionServiceNameOracledbatgoogle,
	"ORACLEDBATAWS":    ListNetworkAnchorsSubscriptionServiceNameOracledbataws,
}

var mappingListNetworkAnchorsSubscriptionServiceNameEnumLowerCase = map[string]ListNetworkAnchorsSubscriptionServiceNameEnum{
	"oracledbatazure":  ListNetworkAnchorsSubscriptionServiceNameOracledbatazure,
	"oracledbatgoogle": ListNetworkAnchorsSubscriptionServiceNameOracledbatgoogle,
	"oracledbataws":    ListNetworkAnchorsSubscriptionServiceNameOracledbataws,
}

// GetListNetworkAnchorsSubscriptionServiceNameEnumValues Enumerates the set of values for ListNetworkAnchorsSubscriptionServiceNameEnum
func GetListNetworkAnchorsSubscriptionServiceNameEnumValues() []ListNetworkAnchorsSubscriptionServiceNameEnum {
	values := make([]ListNetworkAnchorsSubscriptionServiceNameEnum, 0)
	for _, v := range mappingListNetworkAnchorsSubscriptionServiceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkAnchorsSubscriptionServiceNameEnumStringValues Enumerates the set of values in String for ListNetworkAnchorsSubscriptionServiceNameEnum
func GetListNetworkAnchorsSubscriptionServiceNameEnumStringValues() []string {
	return []string{
		"ORACLEDBATAZURE",
		"ORACLEDBATGOOGLE",
		"ORACLEDBATAWS",
	}
}

// GetMappingListNetworkAnchorsSubscriptionServiceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkAnchorsSubscriptionServiceNameEnum(val string) (ListNetworkAnchorsSubscriptionServiceNameEnum, bool) {
	enum, ok := mappingListNetworkAnchorsSubscriptionServiceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkAnchorsSortOrderEnum Enum with underlying type: string
type ListNetworkAnchorsSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkAnchorsSortOrderEnum
const (
	ListNetworkAnchorsSortOrderAsc  ListNetworkAnchorsSortOrderEnum = "ASC"
	ListNetworkAnchorsSortOrderDesc ListNetworkAnchorsSortOrderEnum = "DESC"
)

var mappingListNetworkAnchorsSortOrderEnum = map[string]ListNetworkAnchorsSortOrderEnum{
	"ASC":  ListNetworkAnchorsSortOrderAsc,
	"DESC": ListNetworkAnchorsSortOrderDesc,
}

var mappingListNetworkAnchorsSortOrderEnumLowerCase = map[string]ListNetworkAnchorsSortOrderEnum{
	"asc":  ListNetworkAnchorsSortOrderAsc,
	"desc": ListNetworkAnchorsSortOrderDesc,
}

// GetListNetworkAnchorsSortOrderEnumValues Enumerates the set of values for ListNetworkAnchorsSortOrderEnum
func GetListNetworkAnchorsSortOrderEnumValues() []ListNetworkAnchorsSortOrderEnum {
	values := make([]ListNetworkAnchorsSortOrderEnum, 0)
	for _, v := range mappingListNetworkAnchorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkAnchorsSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkAnchorsSortOrderEnum
func GetListNetworkAnchorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkAnchorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkAnchorsSortOrderEnum(val string) (ListNetworkAnchorsSortOrderEnum, bool) {
	enum, ok := mappingListNetworkAnchorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkAnchorsSortByEnum Enum with underlying type: string
type ListNetworkAnchorsSortByEnum string

// Set of constants representing the allowable values for ListNetworkAnchorsSortByEnum
const (
	ListNetworkAnchorsSortByTimecreated ListNetworkAnchorsSortByEnum = "timeCreated"
	ListNetworkAnchorsSortByDisplayname ListNetworkAnchorsSortByEnum = "displayName"
)

var mappingListNetworkAnchorsSortByEnum = map[string]ListNetworkAnchorsSortByEnum{
	"timeCreated": ListNetworkAnchorsSortByTimecreated,
	"displayName": ListNetworkAnchorsSortByDisplayname,
}

var mappingListNetworkAnchorsSortByEnumLowerCase = map[string]ListNetworkAnchorsSortByEnum{
	"timecreated": ListNetworkAnchorsSortByTimecreated,
	"displayname": ListNetworkAnchorsSortByDisplayname,
}

// GetListNetworkAnchorsSortByEnumValues Enumerates the set of values for ListNetworkAnchorsSortByEnum
func GetListNetworkAnchorsSortByEnumValues() []ListNetworkAnchorsSortByEnum {
	values := make([]ListNetworkAnchorsSortByEnum, 0)
	for _, v := range mappingListNetworkAnchorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkAnchorsSortByEnumStringValues Enumerates the set of values in String for ListNetworkAnchorsSortByEnum
func GetListNetworkAnchorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNetworkAnchorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkAnchorsSortByEnum(val string) (ListNetworkAnchorsSortByEnum, bool) {
	enum, ok := mappingListNetworkAnchorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
