// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDedicatedAiClustersRequest wrapper for the ListDedicatedAiClusters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListDedicatedAiClusters.go.html to see an example of how to use ListDedicatedAiClustersRequest.
type ListDedicatedAiClustersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the dedicated AI clusters that their lifecycle state matches the given lifecycle state.
	LifecycleState DedicatedAiClusterLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated AI cluster.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDedicatedAiClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListDedicatedAiClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDedicatedAiClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDedicatedAiClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDedicatedAiClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDedicatedAiClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDedicatedAiClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedAiClusterLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDedicatedAiClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDedicatedAiClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDedicatedAiClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDedicatedAiClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDedicatedAiClustersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDedicatedAiClustersResponse wrapper for the ListDedicatedAiClusters operation
type ListDedicatedAiClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DedicatedAiClusterCollection instances
	DedicatedAiClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDedicatedAiClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDedicatedAiClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDedicatedAiClustersSortOrderEnum Enum with underlying type: string
type ListDedicatedAiClustersSortOrderEnum string

// Set of constants representing the allowable values for ListDedicatedAiClustersSortOrderEnum
const (
	ListDedicatedAiClustersSortOrderAsc  ListDedicatedAiClustersSortOrderEnum = "ASC"
	ListDedicatedAiClustersSortOrderDesc ListDedicatedAiClustersSortOrderEnum = "DESC"
)

var mappingListDedicatedAiClustersSortOrderEnum = map[string]ListDedicatedAiClustersSortOrderEnum{
	"ASC":  ListDedicatedAiClustersSortOrderAsc,
	"DESC": ListDedicatedAiClustersSortOrderDesc,
}

var mappingListDedicatedAiClustersSortOrderEnumLowerCase = map[string]ListDedicatedAiClustersSortOrderEnum{
	"asc":  ListDedicatedAiClustersSortOrderAsc,
	"desc": ListDedicatedAiClustersSortOrderDesc,
}

// GetListDedicatedAiClustersSortOrderEnumValues Enumerates the set of values for ListDedicatedAiClustersSortOrderEnum
func GetListDedicatedAiClustersSortOrderEnumValues() []ListDedicatedAiClustersSortOrderEnum {
	values := make([]ListDedicatedAiClustersSortOrderEnum, 0)
	for _, v := range mappingListDedicatedAiClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDedicatedAiClustersSortOrderEnumStringValues Enumerates the set of values in String for ListDedicatedAiClustersSortOrderEnum
func GetListDedicatedAiClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDedicatedAiClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDedicatedAiClustersSortOrderEnum(val string) (ListDedicatedAiClustersSortOrderEnum, bool) {
	enum, ok := mappingListDedicatedAiClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDedicatedAiClustersSortByEnum Enum with underlying type: string
type ListDedicatedAiClustersSortByEnum string

// Set of constants representing the allowable values for ListDedicatedAiClustersSortByEnum
const (
	ListDedicatedAiClustersSortByTimecreated    ListDedicatedAiClustersSortByEnum = "timeCreated"
	ListDedicatedAiClustersSortByDisplayname    ListDedicatedAiClustersSortByEnum = "displayName"
	ListDedicatedAiClustersSortByLifecyclestate ListDedicatedAiClustersSortByEnum = "lifecycleState"
)

var mappingListDedicatedAiClustersSortByEnum = map[string]ListDedicatedAiClustersSortByEnum{
	"timeCreated":    ListDedicatedAiClustersSortByTimecreated,
	"displayName":    ListDedicatedAiClustersSortByDisplayname,
	"lifecycleState": ListDedicatedAiClustersSortByLifecyclestate,
}

var mappingListDedicatedAiClustersSortByEnumLowerCase = map[string]ListDedicatedAiClustersSortByEnum{
	"timecreated":    ListDedicatedAiClustersSortByTimecreated,
	"displayname":    ListDedicatedAiClustersSortByDisplayname,
	"lifecyclestate": ListDedicatedAiClustersSortByLifecyclestate,
}

// GetListDedicatedAiClustersSortByEnumValues Enumerates the set of values for ListDedicatedAiClustersSortByEnum
func GetListDedicatedAiClustersSortByEnumValues() []ListDedicatedAiClustersSortByEnum {
	values := make([]ListDedicatedAiClustersSortByEnum, 0)
	for _, v := range mappingListDedicatedAiClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDedicatedAiClustersSortByEnumStringValues Enumerates the set of values in String for ListDedicatedAiClustersSortByEnum
func GetListDedicatedAiClustersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
	}
}

// GetMappingListDedicatedAiClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDedicatedAiClustersSortByEnum(val string) (ListDedicatedAiClustersSortByEnum, bool) {
	enum, ok := mappingListDedicatedAiClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
