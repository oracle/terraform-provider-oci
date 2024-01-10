// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOperationsInsightsPrivateEndpointsRequest wrapper for the ListOperationsInsightsPrivateEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOperationsInsightsPrivateEndpoints.go.html to see an example of how to use ListOperationsInsightsPrivateEndpointsRequest.
type ListOperationsInsightsPrivateEndpointsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Operations Insights PrivateEndpoint identifier
	OpsiPrivateEndpointId *string `mandatory:"false" contributesTo:"query" name:"opsiPrivateEndpointId"`

	// The option to filter OPSI private endpoints that can used for RAC. Should be used along with vcnId query parameter.
	IsUsedForRacDbs *bool `mandatory:"false" contributesTo:"query" name:"isUsedForRacDbs"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// Lifecycle states
	LifecycleState []OperationsInsightsPrivateEndpointLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOperationsInsightsPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort private endpoints.
	SortBy ListOperationsInsightsPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOperationsInsightsPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOperationsInsightsPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOperationsInsightsPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOperationsInsightsPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOperationsInsightsPrivateEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingOperationsInsightsPrivateEndpointLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetOperationsInsightsPrivateEndpointLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListOperationsInsightsPrivateEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOperationsInsightsPrivateEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOperationsInsightsPrivateEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOperationsInsightsPrivateEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOperationsInsightsPrivateEndpointsResponse wrapper for the ListOperationsInsightsPrivateEndpoints operation
type ListOperationsInsightsPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OperationsInsightsPrivateEndpointCollection instances
	OperationsInsightsPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOperationsInsightsPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOperationsInsightsPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOperationsInsightsPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListOperationsInsightsPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListOperationsInsightsPrivateEndpointsSortOrderEnum
const (
	ListOperationsInsightsPrivateEndpointsSortOrderAsc  ListOperationsInsightsPrivateEndpointsSortOrderEnum = "ASC"
	ListOperationsInsightsPrivateEndpointsSortOrderDesc ListOperationsInsightsPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListOperationsInsightsPrivateEndpointsSortOrderEnum = map[string]ListOperationsInsightsPrivateEndpointsSortOrderEnum{
	"ASC":  ListOperationsInsightsPrivateEndpointsSortOrderAsc,
	"DESC": ListOperationsInsightsPrivateEndpointsSortOrderDesc,
}

var mappingListOperationsInsightsPrivateEndpointsSortOrderEnumLowerCase = map[string]ListOperationsInsightsPrivateEndpointsSortOrderEnum{
	"asc":  ListOperationsInsightsPrivateEndpointsSortOrderAsc,
	"desc": ListOperationsInsightsPrivateEndpointsSortOrderDesc,
}

// GetListOperationsInsightsPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListOperationsInsightsPrivateEndpointsSortOrderEnum
func GetListOperationsInsightsPrivateEndpointsSortOrderEnumValues() []ListOperationsInsightsPrivateEndpointsSortOrderEnum {
	values := make([]ListOperationsInsightsPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListOperationsInsightsPrivateEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperationsInsightsPrivateEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListOperationsInsightsPrivateEndpointsSortOrderEnum
func GetListOperationsInsightsPrivateEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOperationsInsightsPrivateEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperationsInsightsPrivateEndpointsSortOrderEnum(val string) (ListOperationsInsightsPrivateEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListOperationsInsightsPrivateEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOperationsInsightsPrivateEndpointsSortByEnum Enum with underlying type: string
type ListOperationsInsightsPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListOperationsInsightsPrivateEndpointsSortByEnum
const (
	ListOperationsInsightsPrivateEndpointsSortByTimecreated ListOperationsInsightsPrivateEndpointsSortByEnum = "timeCreated"
	ListOperationsInsightsPrivateEndpointsSortById          ListOperationsInsightsPrivateEndpointsSortByEnum = "id"
	ListOperationsInsightsPrivateEndpointsSortByDisplayname ListOperationsInsightsPrivateEndpointsSortByEnum = "displayName"
)

var mappingListOperationsInsightsPrivateEndpointsSortByEnum = map[string]ListOperationsInsightsPrivateEndpointsSortByEnum{
	"timeCreated": ListOperationsInsightsPrivateEndpointsSortByTimecreated,
	"id":          ListOperationsInsightsPrivateEndpointsSortById,
	"displayName": ListOperationsInsightsPrivateEndpointsSortByDisplayname,
}

var mappingListOperationsInsightsPrivateEndpointsSortByEnumLowerCase = map[string]ListOperationsInsightsPrivateEndpointsSortByEnum{
	"timecreated": ListOperationsInsightsPrivateEndpointsSortByTimecreated,
	"id":          ListOperationsInsightsPrivateEndpointsSortById,
	"displayname": ListOperationsInsightsPrivateEndpointsSortByDisplayname,
}

// GetListOperationsInsightsPrivateEndpointsSortByEnumValues Enumerates the set of values for ListOperationsInsightsPrivateEndpointsSortByEnum
func GetListOperationsInsightsPrivateEndpointsSortByEnumValues() []ListOperationsInsightsPrivateEndpointsSortByEnum {
	values := make([]ListOperationsInsightsPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListOperationsInsightsPrivateEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperationsInsightsPrivateEndpointsSortByEnumStringValues Enumerates the set of values in String for ListOperationsInsightsPrivateEndpointsSortByEnum
func GetListOperationsInsightsPrivateEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"id",
		"displayName",
	}
}

// GetMappingListOperationsInsightsPrivateEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperationsInsightsPrivateEndpointsSortByEnum(val string) (ListOperationsInsightsPrivateEndpointsSortByEnum, bool) {
	enum, ok := mappingListOperationsInsightsPrivateEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
