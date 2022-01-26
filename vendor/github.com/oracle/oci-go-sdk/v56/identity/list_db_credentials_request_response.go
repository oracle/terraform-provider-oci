// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDbCredentialsRequest wrapper for the ListDbCredentials operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListDbCredentials.go.html to see an example of how to use ListDbCredentialsRequest.
type ListDbCredentialsRequest struct {

	// The OCID of the user.
	UserId *string `mandatory:"true" contributesTo:"path" name:"userId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListDbCredentialsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListDbCredentialsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState DbCredentialLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbCredentialsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbCredentialsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbCredentialsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbCredentialsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDbCredentialsResponse wrapper for the ListDbCredentials operation
type ListDbCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbCredentialSummary instances
	Items []DbCredentialSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbCredentialsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbCredentialsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbCredentialsSortByEnum Enum with underlying type: string
type ListDbCredentialsSortByEnum string

// Set of constants representing the allowable values for ListDbCredentialsSortByEnum
const (
	ListDbCredentialsSortByTimecreated ListDbCredentialsSortByEnum = "TIMECREATED"
	ListDbCredentialsSortByName        ListDbCredentialsSortByEnum = "NAME"
)

var mappingListDbCredentialsSortBy = map[string]ListDbCredentialsSortByEnum{
	"TIMECREATED": ListDbCredentialsSortByTimecreated,
	"NAME":        ListDbCredentialsSortByName,
}

// GetListDbCredentialsSortByEnumValues Enumerates the set of values for ListDbCredentialsSortByEnum
func GetListDbCredentialsSortByEnumValues() []ListDbCredentialsSortByEnum {
	values := make([]ListDbCredentialsSortByEnum, 0)
	for _, v := range mappingListDbCredentialsSortBy {
		values = append(values, v)
	}
	return values
}

// ListDbCredentialsSortOrderEnum Enum with underlying type: string
type ListDbCredentialsSortOrderEnum string

// Set of constants representing the allowable values for ListDbCredentialsSortOrderEnum
const (
	ListDbCredentialsSortOrderAsc  ListDbCredentialsSortOrderEnum = "ASC"
	ListDbCredentialsSortOrderDesc ListDbCredentialsSortOrderEnum = "DESC"
)

var mappingListDbCredentialsSortOrder = map[string]ListDbCredentialsSortOrderEnum{
	"ASC":  ListDbCredentialsSortOrderAsc,
	"DESC": ListDbCredentialsSortOrderDesc,
}

// GetListDbCredentialsSortOrderEnumValues Enumerates the set of values for ListDbCredentialsSortOrderEnum
func GetListDbCredentialsSortOrderEnumValues() []ListDbCredentialsSortOrderEnum {
	values := make([]ListDbCredentialsSortOrderEnum, 0)
	for _, v := range mappingListDbCredentialsSortOrder {
		values = append(values, v)
	}
	return values
}
