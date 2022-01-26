// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListTsigKeysRequest wrapper for the ListTsigKeys operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListTsigKeys.go.html to see an example of how to use ListTsigKeysRequest.
type ListTsigKeysRequest struct {

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of a resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The name of a resource.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The state of a resource.
	LifecycleState TsigKeySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field by which to sort TSIG keys. If unspecified, defaults to `timeCreated`.
	SortBy ListTsigKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the resources.
	SortOrder ListTsigKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListTsigKeysScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTsigKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTsigKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTsigKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTsigKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTsigKeysResponse wrapper for the ListTsigKeys operation
type ListTsigKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TsigKeySummary instances
	Items []TsigKeySummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTsigKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTsigKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTsigKeysSortByEnum Enum with underlying type: string
type ListTsigKeysSortByEnum string

// Set of constants representing the allowable values for ListTsigKeysSortByEnum
const (
	ListTsigKeysSortByName        ListTsigKeysSortByEnum = "name"
	ListTsigKeysSortByTimecreated ListTsigKeysSortByEnum = "timeCreated"
)

var mappingListTsigKeysSortBy = map[string]ListTsigKeysSortByEnum{
	"name":        ListTsigKeysSortByName,
	"timeCreated": ListTsigKeysSortByTimecreated,
}

// GetListTsigKeysSortByEnumValues Enumerates the set of values for ListTsigKeysSortByEnum
func GetListTsigKeysSortByEnumValues() []ListTsigKeysSortByEnum {
	values := make([]ListTsigKeysSortByEnum, 0)
	for _, v := range mappingListTsigKeysSortBy {
		values = append(values, v)
	}
	return values
}

// ListTsigKeysSortOrderEnum Enum with underlying type: string
type ListTsigKeysSortOrderEnum string

// Set of constants representing the allowable values for ListTsigKeysSortOrderEnum
const (
	ListTsigKeysSortOrderAsc  ListTsigKeysSortOrderEnum = "ASC"
	ListTsigKeysSortOrderDesc ListTsigKeysSortOrderEnum = "DESC"
)

var mappingListTsigKeysSortOrder = map[string]ListTsigKeysSortOrderEnum{
	"ASC":  ListTsigKeysSortOrderAsc,
	"DESC": ListTsigKeysSortOrderDesc,
}

// GetListTsigKeysSortOrderEnumValues Enumerates the set of values for ListTsigKeysSortOrderEnum
func GetListTsigKeysSortOrderEnumValues() []ListTsigKeysSortOrderEnum {
	values := make([]ListTsigKeysSortOrderEnum, 0)
	for _, v := range mappingListTsigKeysSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTsigKeysScopeEnum Enum with underlying type: string
type ListTsigKeysScopeEnum string

// Set of constants representing the allowable values for ListTsigKeysScopeEnum
const (
	ListTsigKeysScopeGlobal  ListTsigKeysScopeEnum = "GLOBAL"
	ListTsigKeysScopePrivate ListTsigKeysScopeEnum = "PRIVATE"
)

var mappingListTsigKeysScope = map[string]ListTsigKeysScopeEnum{
	"GLOBAL":  ListTsigKeysScopeGlobal,
	"PRIVATE": ListTsigKeysScopePrivate,
}

// GetListTsigKeysScopeEnumValues Enumerates the set of values for ListTsigKeysScopeEnum
func GetListTsigKeysScopeEnumValues() []ListTsigKeysScopeEnum {
	values := make([]ListTsigKeysScopeEnum, 0)
	for _, v := range mappingListTsigKeysScope {
		values = append(values, v)
	}
	return values
}
