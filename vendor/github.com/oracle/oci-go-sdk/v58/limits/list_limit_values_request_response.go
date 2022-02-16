// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package limits

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListLimitValuesRequest wrapper for the ListLimitValues operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limits/ListLimitValues.go.html to see an example of how to use ListLimitValuesRequest.
type ListLimitValuesRequest struct {

	// The OCID of the parent compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The target service name.
	ServiceName *string `mandatory:"true" contributesTo:"query" name:"serviceName"`

	// Filter entries by scope type.
	ScopeType ListLimitValuesScopeTypeEnum `mandatory:"false" contributesTo:"query" name:"scopeType" omitEmpty:"true"`

	// Filter entries by availability domain. This implies that only AD-specific values are returned.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// Optional field, can be used to see a specific resource limit value.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. The sorting is by availabilityDomain, as a second level field, if available.
	SortBy ListLimitValuesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'. By default, it is ascending.
	SortOrder ListLimitValuesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLimitValuesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLimitValuesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLimitValuesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLimitValuesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLimitValuesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLimitValuesScopeTypeEnum(string(request.ScopeType)); !ok && request.ScopeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScopeType: %s. Supported values are: %s.", request.ScopeType, strings.Join(GetListLimitValuesScopeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLimitValuesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLimitValuesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLimitValuesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLimitValuesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLimitValuesResponse wrapper for the ListLimitValues operation
type ListLimitValuesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LimitValueSummary instances
	Items []LimitValueSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLimitValuesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLimitValuesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLimitValuesScopeTypeEnum Enum with underlying type: string
type ListLimitValuesScopeTypeEnum string

// Set of constants representing the allowable values for ListLimitValuesScopeTypeEnum
const (
	ListLimitValuesScopeTypeGlobal ListLimitValuesScopeTypeEnum = "GLOBAL"
	ListLimitValuesScopeTypeRegion ListLimitValuesScopeTypeEnum = "REGION"
	ListLimitValuesScopeTypeAd     ListLimitValuesScopeTypeEnum = "AD"
)

var mappingListLimitValuesScopeTypeEnum = map[string]ListLimitValuesScopeTypeEnum{
	"GLOBAL": ListLimitValuesScopeTypeGlobal,
	"REGION": ListLimitValuesScopeTypeRegion,
	"AD":     ListLimitValuesScopeTypeAd,
}

// GetListLimitValuesScopeTypeEnumValues Enumerates the set of values for ListLimitValuesScopeTypeEnum
func GetListLimitValuesScopeTypeEnumValues() []ListLimitValuesScopeTypeEnum {
	values := make([]ListLimitValuesScopeTypeEnum, 0)
	for _, v := range mappingListLimitValuesScopeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitValuesScopeTypeEnumStringValues Enumerates the set of values in String for ListLimitValuesScopeTypeEnum
func GetListLimitValuesScopeTypeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"REGION",
		"AD",
	}
}

// GetMappingListLimitValuesScopeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitValuesScopeTypeEnum(val string) (ListLimitValuesScopeTypeEnum, bool) {
	mappingListLimitValuesScopeTypeEnumIgnoreCase := make(map[string]ListLimitValuesScopeTypeEnum)
	for k, v := range mappingListLimitValuesScopeTypeEnum {
		mappingListLimitValuesScopeTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLimitValuesScopeTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListLimitValuesSortByEnum Enum with underlying type: string
type ListLimitValuesSortByEnum string

// Set of constants representing the allowable values for ListLimitValuesSortByEnum
const (
	ListLimitValuesSortByName ListLimitValuesSortByEnum = "name"
)

var mappingListLimitValuesSortByEnum = map[string]ListLimitValuesSortByEnum{
	"name": ListLimitValuesSortByName,
}

// GetListLimitValuesSortByEnumValues Enumerates the set of values for ListLimitValuesSortByEnum
func GetListLimitValuesSortByEnumValues() []ListLimitValuesSortByEnum {
	values := make([]ListLimitValuesSortByEnum, 0)
	for _, v := range mappingListLimitValuesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitValuesSortByEnumStringValues Enumerates the set of values in String for ListLimitValuesSortByEnum
func GetListLimitValuesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListLimitValuesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitValuesSortByEnum(val string) (ListLimitValuesSortByEnum, bool) {
	mappingListLimitValuesSortByEnumIgnoreCase := make(map[string]ListLimitValuesSortByEnum)
	for k, v := range mappingListLimitValuesSortByEnum {
		mappingListLimitValuesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLimitValuesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListLimitValuesSortOrderEnum Enum with underlying type: string
type ListLimitValuesSortOrderEnum string

// Set of constants representing the allowable values for ListLimitValuesSortOrderEnum
const (
	ListLimitValuesSortOrderAsc  ListLimitValuesSortOrderEnum = "ASC"
	ListLimitValuesSortOrderDesc ListLimitValuesSortOrderEnum = "DESC"
)

var mappingListLimitValuesSortOrderEnum = map[string]ListLimitValuesSortOrderEnum{
	"ASC":  ListLimitValuesSortOrderAsc,
	"DESC": ListLimitValuesSortOrderDesc,
}

// GetListLimitValuesSortOrderEnumValues Enumerates the set of values for ListLimitValuesSortOrderEnum
func GetListLimitValuesSortOrderEnumValues() []ListLimitValuesSortOrderEnum {
	values := make([]ListLimitValuesSortOrderEnum, 0)
	for _, v := range mappingListLimitValuesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLimitValuesSortOrderEnumStringValues Enumerates the set of values in String for ListLimitValuesSortOrderEnum
func GetListLimitValuesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLimitValuesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLimitValuesSortOrderEnum(val string) (ListLimitValuesSortOrderEnum, bool) {
	mappingListLimitValuesSortOrderEnumIgnoreCase := make(map[string]ListLimitValuesSortOrderEnum)
	for k, v := range mappingListLimitValuesSortOrderEnum {
		mappingListLimitValuesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLimitValuesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
