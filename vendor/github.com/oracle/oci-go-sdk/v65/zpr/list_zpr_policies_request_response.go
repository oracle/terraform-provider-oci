// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package zpr

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListZprPoliciesRequest wrapper for the ListZprPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/zpr/ListZprPolicies.go.html to see an example of how to use ListZprPoliciesRequest.
type ListZprPoliciesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState ZprPolicyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ZprPolicy.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListZprPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `name` is ascending.
	SortBy ListZprPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListZprPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListZprPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListZprPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListZprPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListZprPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingZprPolicyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetZprPolicyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListZprPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListZprPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListZprPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListZprPoliciesResponse wrapper for the ListZprPolicies operation
type ListZprPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ZprPolicyCollection instances
	ZprPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListZprPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListZprPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListZprPoliciesSortOrderEnum Enum with underlying type: string
type ListZprPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListZprPoliciesSortOrderEnum
const (
	ListZprPoliciesSortOrderAsc  ListZprPoliciesSortOrderEnum = "ASC"
	ListZprPoliciesSortOrderDesc ListZprPoliciesSortOrderEnum = "DESC"
)

var mappingListZprPoliciesSortOrderEnum = map[string]ListZprPoliciesSortOrderEnum{
	"ASC":  ListZprPoliciesSortOrderAsc,
	"DESC": ListZprPoliciesSortOrderDesc,
}

var mappingListZprPoliciesSortOrderEnumLowerCase = map[string]ListZprPoliciesSortOrderEnum{
	"asc":  ListZprPoliciesSortOrderAsc,
	"desc": ListZprPoliciesSortOrderDesc,
}

// GetListZprPoliciesSortOrderEnumValues Enumerates the set of values for ListZprPoliciesSortOrderEnum
func GetListZprPoliciesSortOrderEnumValues() []ListZprPoliciesSortOrderEnum {
	values := make([]ListZprPoliciesSortOrderEnum, 0)
	for _, v := range mappingListZprPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListZprPoliciesSortOrderEnum
func GetListZprPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListZprPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPoliciesSortOrderEnum(val string) (ListZprPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListZprPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListZprPoliciesSortByEnum Enum with underlying type: string
type ListZprPoliciesSortByEnum string

// Set of constants representing the allowable values for ListZprPoliciesSortByEnum
const (
	ListZprPoliciesSortByTimecreated ListZprPoliciesSortByEnum = "timeCreated"
	ListZprPoliciesSortByName        ListZprPoliciesSortByEnum = "name"
)

var mappingListZprPoliciesSortByEnum = map[string]ListZprPoliciesSortByEnum{
	"timeCreated": ListZprPoliciesSortByTimecreated,
	"name":        ListZprPoliciesSortByName,
}

var mappingListZprPoliciesSortByEnumLowerCase = map[string]ListZprPoliciesSortByEnum{
	"timecreated": ListZprPoliciesSortByTimecreated,
	"name":        ListZprPoliciesSortByName,
}

// GetListZprPoliciesSortByEnumValues Enumerates the set of values for ListZprPoliciesSortByEnum
func GetListZprPoliciesSortByEnumValues() []ListZprPoliciesSortByEnum {
	values := make([]ListZprPoliciesSortByEnum, 0)
	for _, v := range mappingListZprPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListZprPoliciesSortByEnumStringValues Enumerates the set of values in String for ListZprPoliciesSortByEnum
func GetListZprPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListZprPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListZprPoliciesSortByEnum(val string) (ListZprPoliciesSortByEnum, bool) {
	enum, ok := mappingListZprPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
