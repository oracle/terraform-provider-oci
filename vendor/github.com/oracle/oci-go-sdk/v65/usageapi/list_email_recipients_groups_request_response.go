// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailRecipientsGroupsRequest wrapper for the ListEmailRecipientsGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListEmailRecipientsGroups.go.html to see an example of how to use ListEmailRecipientsGroupsRequest.
type ListEmailRecipientsGroupsRequest struct {

	// The UsageStatement Subscription unique OCID.
	SubscriptionId *string `mandatory:"true" contributesTo:"path" name:"subscriptionId"`

	// The compartment ID in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximumimum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results.
	// This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error, without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The field to sort by. If not specified, the default is displayName.
	SortBy ListEmailRecipientsGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListEmailRecipientsGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailRecipientsGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailRecipientsGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailRecipientsGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailRecipientsGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailRecipientsGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailRecipientsGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailRecipientsGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailRecipientsGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailRecipientsGroupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailRecipientsGroupsResponse wrapper for the ListEmailRecipientsGroups operation
type ListEmailRecipientsGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailRecipientsGroupCollection instances
	EmailRecipientsGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEmailRecipientsGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailRecipientsGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailRecipientsGroupsSortByEnum Enum with underlying type: string
type ListEmailRecipientsGroupsSortByEnum string

// Set of constants representing the allowable values for ListEmailRecipientsGroupsSortByEnum
const (
	ListEmailRecipientsGroupsSortByDisplayname ListEmailRecipientsGroupsSortByEnum = "displayName"
)

var mappingListEmailRecipientsGroupsSortByEnum = map[string]ListEmailRecipientsGroupsSortByEnum{
	"displayName": ListEmailRecipientsGroupsSortByDisplayname,
}

var mappingListEmailRecipientsGroupsSortByEnumLowerCase = map[string]ListEmailRecipientsGroupsSortByEnum{
	"displayname": ListEmailRecipientsGroupsSortByDisplayname,
}

// GetListEmailRecipientsGroupsSortByEnumValues Enumerates the set of values for ListEmailRecipientsGroupsSortByEnum
func GetListEmailRecipientsGroupsSortByEnumValues() []ListEmailRecipientsGroupsSortByEnum {
	values := make([]ListEmailRecipientsGroupsSortByEnum, 0)
	for _, v := range mappingListEmailRecipientsGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailRecipientsGroupsSortByEnumStringValues Enumerates the set of values in String for ListEmailRecipientsGroupsSortByEnum
func GetListEmailRecipientsGroupsSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListEmailRecipientsGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailRecipientsGroupsSortByEnum(val string) (ListEmailRecipientsGroupsSortByEnum, bool) {
	enum, ok := mappingListEmailRecipientsGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailRecipientsGroupsSortOrderEnum Enum with underlying type: string
type ListEmailRecipientsGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailRecipientsGroupsSortOrderEnum
const (
	ListEmailRecipientsGroupsSortOrderAsc  ListEmailRecipientsGroupsSortOrderEnum = "ASC"
	ListEmailRecipientsGroupsSortOrderDesc ListEmailRecipientsGroupsSortOrderEnum = "DESC"
)

var mappingListEmailRecipientsGroupsSortOrderEnum = map[string]ListEmailRecipientsGroupsSortOrderEnum{
	"ASC":  ListEmailRecipientsGroupsSortOrderAsc,
	"DESC": ListEmailRecipientsGroupsSortOrderDesc,
}

var mappingListEmailRecipientsGroupsSortOrderEnumLowerCase = map[string]ListEmailRecipientsGroupsSortOrderEnum{
	"asc":  ListEmailRecipientsGroupsSortOrderAsc,
	"desc": ListEmailRecipientsGroupsSortOrderDesc,
}

// GetListEmailRecipientsGroupsSortOrderEnumValues Enumerates the set of values for ListEmailRecipientsGroupsSortOrderEnum
func GetListEmailRecipientsGroupsSortOrderEnumValues() []ListEmailRecipientsGroupsSortOrderEnum {
	values := make([]ListEmailRecipientsGroupsSortOrderEnum, 0)
	for _, v := range mappingListEmailRecipientsGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailRecipientsGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailRecipientsGroupsSortOrderEnum
func GetListEmailRecipientsGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailRecipientsGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailRecipientsGroupsSortOrderEnum(val string) (ListEmailRecipientsGroupsSortOrderEnum, bool) {
	enum, ok := mappingListEmailRecipientsGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
