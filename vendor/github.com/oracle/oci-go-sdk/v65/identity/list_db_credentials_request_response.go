// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbCredentialsRequest wrapper for the ListDbCredentials operation
//
// # See also
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbCredentialsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbCredentialsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbCredentialsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbCredentialsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbCredentialsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbCredentialLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbCredentialLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListDbCredentialsSortByEnum = map[string]ListDbCredentialsSortByEnum{
	"TIMECREATED": ListDbCredentialsSortByTimecreated,
	"NAME":        ListDbCredentialsSortByName,
}

var mappingListDbCredentialsSortByEnumLowerCase = map[string]ListDbCredentialsSortByEnum{
	"timecreated": ListDbCredentialsSortByTimecreated,
	"name":        ListDbCredentialsSortByName,
}

// GetListDbCredentialsSortByEnumValues Enumerates the set of values for ListDbCredentialsSortByEnum
func GetListDbCredentialsSortByEnumValues() []ListDbCredentialsSortByEnum {
	values := make([]ListDbCredentialsSortByEnum, 0)
	for _, v := range mappingListDbCredentialsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbCredentialsSortByEnumStringValues Enumerates the set of values in String for ListDbCredentialsSortByEnum
func GetListDbCredentialsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListDbCredentialsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbCredentialsSortByEnum(val string) (ListDbCredentialsSortByEnum, bool) {
	enum, ok := mappingListDbCredentialsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbCredentialsSortOrderEnum Enum with underlying type: string
type ListDbCredentialsSortOrderEnum string

// Set of constants representing the allowable values for ListDbCredentialsSortOrderEnum
const (
	ListDbCredentialsSortOrderAsc  ListDbCredentialsSortOrderEnum = "ASC"
	ListDbCredentialsSortOrderDesc ListDbCredentialsSortOrderEnum = "DESC"
)

var mappingListDbCredentialsSortOrderEnum = map[string]ListDbCredentialsSortOrderEnum{
	"ASC":  ListDbCredentialsSortOrderAsc,
	"DESC": ListDbCredentialsSortOrderDesc,
}

var mappingListDbCredentialsSortOrderEnumLowerCase = map[string]ListDbCredentialsSortOrderEnum{
	"asc":  ListDbCredentialsSortOrderAsc,
	"desc": ListDbCredentialsSortOrderDesc,
}

// GetListDbCredentialsSortOrderEnumValues Enumerates the set of values for ListDbCredentialsSortOrderEnum
func GetListDbCredentialsSortOrderEnumValues() []ListDbCredentialsSortOrderEnum {
	values := make([]ListDbCredentialsSortOrderEnum, 0)
	for _, v := range mappingListDbCredentialsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbCredentialsSortOrderEnumStringValues Enumerates the set of values in String for ListDbCredentialsSortOrderEnum
func GetListDbCredentialsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbCredentialsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbCredentialsSortOrderEnum(val string) (ListDbCredentialsSortOrderEnum, bool) {
	enum, ok := mappingListDbCredentialsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
