// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLinksRequest wrapper for the ListLinks operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListLinks.go.html to see an example of how to use ListLinksRequest.
type ListLinksRequest struct {

	// The ID of the parent tenancy this link is associated with.
	ParentTenancyId *string `mandatory:"false" contributesTo:"query" name:"parentTenancyId"`

	// The ID of the child tenancy this link is associated with.
	ChildTenancyId *string `mandatory:"false" contributesTo:"query" name:"childTenancyId"`

	// The lifecycle state of the resource.
	LifecycleState ListLinksLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListLinksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLinksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLinksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLinksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLinksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLinksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLinksLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListLinksLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLinksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLinksSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLinksResponse wrapper for the ListLinks operation
type ListLinksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LinkCollection instances
	LinkCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLinksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLinksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLinksLifecycleStateEnum Enum with underlying type: string
type ListLinksLifecycleStateEnum string

// Set of constants representing the allowable values for ListLinksLifecycleStateEnum
const (
	ListLinksLifecycleStateCreating   ListLinksLifecycleStateEnum = "CREATING"
	ListLinksLifecycleStateActive     ListLinksLifecycleStateEnum = "ACTIVE"
	ListLinksLifecycleStateInactive   ListLinksLifecycleStateEnum = "INACTIVE"
	ListLinksLifecycleStateUpdating   ListLinksLifecycleStateEnum = "UPDATING"
	ListLinksLifecycleStateFailed     ListLinksLifecycleStateEnum = "FAILED"
	ListLinksLifecycleStateTerminated ListLinksLifecycleStateEnum = "TERMINATED"
)

var mappingListLinksLifecycleStateEnum = map[string]ListLinksLifecycleStateEnum{
	"CREATING":   ListLinksLifecycleStateCreating,
	"ACTIVE":     ListLinksLifecycleStateActive,
	"INACTIVE":   ListLinksLifecycleStateInactive,
	"UPDATING":   ListLinksLifecycleStateUpdating,
	"FAILED":     ListLinksLifecycleStateFailed,
	"TERMINATED": ListLinksLifecycleStateTerminated,
}

var mappingListLinksLifecycleStateEnumLowerCase = map[string]ListLinksLifecycleStateEnum{
	"creating":   ListLinksLifecycleStateCreating,
	"active":     ListLinksLifecycleStateActive,
	"inactive":   ListLinksLifecycleStateInactive,
	"updating":   ListLinksLifecycleStateUpdating,
	"failed":     ListLinksLifecycleStateFailed,
	"terminated": ListLinksLifecycleStateTerminated,
}

// GetListLinksLifecycleStateEnumValues Enumerates the set of values for ListLinksLifecycleStateEnum
func GetListLinksLifecycleStateEnumValues() []ListLinksLifecycleStateEnum {
	values := make([]ListLinksLifecycleStateEnum, 0)
	for _, v := range mappingListLinksLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListLinksLifecycleStateEnumStringValues Enumerates the set of values in String for ListLinksLifecycleStateEnum
func GetListLinksLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"FAILED",
		"TERMINATED",
	}
}

// GetMappingListLinksLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLinksLifecycleStateEnum(val string) (ListLinksLifecycleStateEnum, bool) {
	enum, ok := mappingListLinksLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLinksSortOrderEnum Enum with underlying type: string
type ListLinksSortOrderEnum string

// Set of constants representing the allowable values for ListLinksSortOrderEnum
const (
	ListLinksSortOrderAsc  ListLinksSortOrderEnum = "ASC"
	ListLinksSortOrderDesc ListLinksSortOrderEnum = "DESC"
)

var mappingListLinksSortOrderEnum = map[string]ListLinksSortOrderEnum{
	"ASC":  ListLinksSortOrderAsc,
	"DESC": ListLinksSortOrderDesc,
}

var mappingListLinksSortOrderEnumLowerCase = map[string]ListLinksSortOrderEnum{
	"asc":  ListLinksSortOrderAsc,
	"desc": ListLinksSortOrderDesc,
}

// GetListLinksSortOrderEnumValues Enumerates the set of values for ListLinksSortOrderEnum
func GetListLinksSortOrderEnumValues() []ListLinksSortOrderEnum {
	values := make([]ListLinksSortOrderEnum, 0)
	for _, v := range mappingListLinksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLinksSortOrderEnumStringValues Enumerates the set of values in String for ListLinksSortOrderEnum
func GetListLinksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLinksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLinksSortOrderEnum(val string) (ListLinksSortOrderEnum, bool) {
	enum, ok := mappingListLinksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
