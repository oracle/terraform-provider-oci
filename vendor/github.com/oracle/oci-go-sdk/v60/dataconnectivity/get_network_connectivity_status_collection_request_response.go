// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"net/http"
	"strings"
)

// GetNetworkConnectivityStatusCollectionRequest wrapper for the GetNetworkConnectivityStatusCollection operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataconnectivity/GetNetworkConnectivityStatusCollection.go.html to see an example of how to use GetNetworkConnectivityStatusCollectionRequest.
type GetNetworkConnectivityStatusCollectionRequest struct {

	// The registry Ocid.
	RegistryId *string `mandatory:"true" contributesTo:"path" name:"registryId"`

	// The endpoint key.
	EndpointKey *string `mandatory:"true" contributesTo:"path" name:"endpointKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or server error without risk of executing that same action again.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy GetNetworkConnectivityStatusCollectionSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder GetNetworkConnectivityStatusCollectionSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetNetworkConnectivityStatusCollectionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetNetworkConnectivityStatusCollectionRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetNetworkConnectivityStatusCollectionRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetNetworkConnectivityStatusCollectionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetNetworkConnectivityStatusCollectionRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetNetworkConnectivityStatusCollectionSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetNetworkConnectivityStatusCollectionSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetNetworkConnectivityStatusCollectionSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetNetworkConnectivityStatusCollectionSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetNetworkConnectivityStatusCollectionResponse wrapper for the GetNetworkConnectivityStatusCollection operation
type GetNetworkConnectivityStatusCollectionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NetworkConnectivityStatusCollection instances
	NetworkConnectivityStatusCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetNetworkConnectivityStatusCollectionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetNetworkConnectivityStatusCollectionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetNetworkConnectivityStatusCollectionSortByEnum Enum with underlying type: string
type GetNetworkConnectivityStatusCollectionSortByEnum string

// Set of constants representing the allowable values for GetNetworkConnectivityStatusCollectionSortByEnum
const (
	GetNetworkConnectivityStatusCollectionSortById          GetNetworkConnectivityStatusCollectionSortByEnum = "id"
	GetNetworkConnectivityStatusCollectionSortByTimecreated GetNetworkConnectivityStatusCollectionSortByEnum = "timeCreated"
	GetNetworkConnectivityStatusCollectionSortByDisplayname GetNetworkConnectivityStatusCollectionSortByEnum = "displayName"
)

var mappingGetNetworkConnectivityStatusCollectionSortByEnum = map[string]GetNetworkConnectivityStatusCollectionSortByEnum{
	"id":          GetNetworkConnectivityStatusCollectionSortById,
	"timeCreated": GetNetworkConnectivityStatusCollectionSortByTimecreated,
	"displayName": GetNetworkConnectivityStatusCollectionSortByDisplayname,
}

var mappingGetNetworkConnectivityStatusCollectionSortByEnumLowerCase = map[string]GetNetworkConnectivityStatusCollectionSortByEnum{
	"id":          GetNetworkConnectivityStatusCollectionSortById,
	"timecreated": GetNetworkConnectivityStatusCollectionSortByTimecreated,
	"displayname": GetNetworkConnectivityStatusCollectionSortByDisplayname,
}

// GetGetNetworkConnectivityStatusCollectionSortByEnumValues Enumerates the set of values for GetNetworkConnectivityStatusCollectionSortByEnum
func GetGetNetworkConnectivityStatusCollectionSortByEnumValues() []GetNetworkConnectivityStatusCollectionSortByEnum {
	values := make([]GetNetworkConnectivityStatusCollectionSortByEnum, 0)
	for _, v := range mappingGetNetworkConnectivityStatusCollectionSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetNetworkConnectivityStatusCollectionSortByEnumStringValues Enumerates the set of values in String for GetNetworkConnectivityStatusCollectionSortByEnum
func GetGetNetworkConnectivityStatusCollectionSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"displayName",
	}
}

// GetMappingGetNetworkConnectivityStatusCollectionSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetNetworkConnectivityStatusCollectionSortByEnum(val string) (GetNetworkConnectivityStatusCollectionSortByEnum, bool) {
	enum, ok := mappingGetNetworkConnectivityStatusCollectionSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetNetworkConnectivityStatusCollectionSortOrderEnum Enum with underlying type: string
type GetNetworkConnectivityStatusCollectionSortOrderEnum string

// Set of constants representing the allowable values for GetNetworkConnectivityStatusCollectionSortOrderEnum
const (
	GetNetworkConnectivityStatusCollectionSortOrderAsc  GetNetworkConnectivityStatusCollectionSortOrderEnum = "ASC"
	GetNetworkConnectivityStatusCollectionSortOrderDesc GetNetworkConnectivityStatusCollectionSortOrderEnum = "DESC"
)

var mappingGetNetworkConnectivityStatusCollectionSortOrderEnum = map[string]GetNetworkConnectivityStatusCollectionSortOrderEnum{
	"ASC":  GetNetworkConnectivityStatusCollectionSortOrderAsc,
	"DESC": GetNetworkConnectivityStatusCollectionSortOrderDesc,
}

var mappingGetNetworkConnectivityStatusCollectionSortOrderEnumLowerCase = map[string]GetNetworkConnectivityStatusCollectionSortOrderEnum{
	"asc":  GetNetworkConnectivityStatusCollectionSortOrderAsc,
	"desc": GetNetworkConnectivityStatusCollectionSortOrderDesc,
}

// GetGetNetworkConnectivityStatusCollectionSortOrderEnumValues Enumerates the set of values for GetNetworkConnectivityStatusCollectionSortOrderEnum
func GetGetNetworkConnectivityStatusCollectionSortOrderEnumValues() []GetNetworkConnectivityStatusCollectionSortOrderEnum {
	values := make([]GetNetworkConnectivityStatusCollectionSortOrderEnum, 0)
	for _, v := range mappingGetNetworkConnectivityStatusCollectionSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetNetworkConnectivityStatusCollectionSortOrderEnumStringValues Enumerates the set of values in String for GetNetworkConnectivityStatusCollectionSortOrderEnum
func GetGetNetworkConnectivityStatusCollectionSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetNetworkConnectivityStatusCollectionSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetNetworkConnectivityStatusCollectionSortOrderEnum(val string) (GetNetworkConnectivityStatusCollectionSortOrderEnum, bool) {
	enum, ok := mappingGetNetworkConnectivityStatusCollectionSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
