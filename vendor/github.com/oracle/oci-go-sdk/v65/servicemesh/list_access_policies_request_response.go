// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAccessPoliciesRequest wrapper for the ListAccessPolicies operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListAccessPolicies.go.html to see an example of how to use ListAccessPoliciesRequest.
type ListAccessPoliciesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAccessPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for 'timeCreated' is descending. Default order for 'name' is ascending.
	SortBy ListAccessPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique Mesh identifier.
	MeshId *string `mandatory:"false" contributesTo:"query" name:"meshId"`

	// Unique AccessPolicy identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the life cycle state given.
	LifecycleState *string `mandatory:"false" contributesTo:"query" name:"lifecycleState"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAccessPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAccessPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAccessPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAccessPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAccessPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAccessPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAccessPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAccessPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAccessPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAccessPoliciesResponse wrapper for the ListAccessPolicies operation
type ListAccessPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AccessPolicyCollection instances
	AccessPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAccessPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAccessPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAccessPoliciesSortOrderEnum Enum with underlying type: string
type ListAccessPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListAccessPoliciesSortOrderEnum
const (
	ListAccessPoliciesSortOrderAsc  ListAccessPoliciesSortOrderEnum = "ASC"
	ListAccessPoliciesSortOrderDesc ListAccessPoliciesSortOrderEnum = "DESC"
)

var mappingListAccessPoliciesSortOrderEnum = map[string]ListAccessPoliciesSortOrderEnum{
	"ASC":  ListAccessPoliciesSortOrderAsc,
	"DESC": ListAccessPoliciesSortOrderDesc,
}

var mappingListAccessPoliciesSortOrderEnumLowerCase = map[string]ListAccessPoliciesSortOrderEnum{
	"asc":  ListAccessPoliciesSortOrderAsc,
	"desc": ListAccessPoliciesSortOrderDesc,
}

// GetListAccessPoliciesSortOrderEnumValues Enumerates the set of values for ListAccessPoliciesSortOrderEnum
func GetListAccessPoliciesSortOrderEnumValues() []ListAccessPoliciesSortOrderEnum {
	values := make([]ListAccessPoliciesSortOrderEnum, 0)
	for _, v := range mappingListAccessPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAccessPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListAccessPoliciesSortOrderEnum
func GetListAccessPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAccessPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAccessPoliciesSortOrderEnum(val string) (ListAccessPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListAccessPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAccessPoliciesSortByEnum Enum with underlying type: string
type ListAccessPoliciesSortByEnum string

// Set of constants representing the allowable values for ListAccessPoliciesSortByEnum
const (
	ListAccessPoliciesSortById          ListAccessPoliciesSortByEnum = "id"
	ListAccessPoliciesSortByTimecreated ListAccessPoliciesSortByEnum = "timeCreated"
	ListAccessPoliciesSortByName        ListAccessPoliciesSortByEnum = "name"
)

var mappingListAccessPoliciesSortByEnum = map[string]ListAccessPoliciesSortByEnum{
	"id":          ListAccessPoliciesSortById,
	"timeCreated": ListAccessPoliciesSortByTimecreated,
	"name":        ListAccessPoliciesSortByName,
}

var mappingListAccessPoliciesSortByEnumLowerCase = map[string]ListAccessPoliciesSortByEnum{
	"id":          ListAccessPoliciesSortById,
	"timecreated": ListAccessPoliciesSortByTimecreated,
	"name":        ListAccessPoliciesSortByName,
}

// GetListAccessPoliciesSortByEnumValues Enumerates the set of values for ListAccessPoliciesSortByEnum
func GetListAccessPoliciesSortByEnumValues() []ListAccessPoliciesSortByEnum {
	values := make([]ListAccessPoliciesSortByEnum, 0)
	for _, v := range mappingListAccessPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAccessPoliciesSortByEnumStringValues Enumerates the set of values in String for ListAccessPoliciesSortByEnum
func GetListAccessPoliciesSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"name",
	}
}

// GetMappingListAccessPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAccessPoliciesSortByEnum(val string) (ListAccessPoliciesSortByEnum, bool) {
	enum, ok := mappingListAccessPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
