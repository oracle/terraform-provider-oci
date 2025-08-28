// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSoftwareUpdatesRequest wrapper for the ListSoftwareUpdates operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListSoftwareUpdates.go.html to see an example of how to use ListSoftwareUpdatesRequest.
type ListSoftwareUpdatesRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListSoftwareUpdatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSoftwareUpdatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error, without risk of executing that same action again. Retry tokens expire after 24
	// hours but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareUpdatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwareUpdatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwareUpdatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSoftwareUpdatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwareUpdatesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareUpdatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwareUpdatesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwareUpdatesResponse wrapper for the ListSoftwareUpdates operation
type ListSoftwareUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwareUpdateCollection instances
	SoftwareUpdateCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSoftwareUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareUpdatesSortByEnum Enum with underlying type: string
type ListSoftwareUpdatesSortByEnum string

// Set of constants representing the allowable values for ListSoftwareUpdatesSortByEnum
const (
	ListSoftwareUpdatesSortByTimecreated ListSoftwareUpdatesSortByEnum = "timeCreated"
	ListSoftwareUpdatesSortByDisplayname ListSoftwareUpdatesSortByEnum = "displayName"
)

var mappingListSoftwareUpdatesSortByEnum = map[string]ListSoftwareUpdatesSortByEnum{
	"timeCreated": ListSoftwareUpdatesSortByTimecreated,
	"displayName": ListSoftwareUpdatesSortByDisplayname,
}

var mappingListSoftwareUpdatesSortByEnumLowerCase = map[string]ListSoftwareUpdatesSortByEnum{
	"timecreated": ListSoftwareUpdatesSortByTimecreated,
	"displayname": ListSoftwareUpdatesSortByDisplayname,
}

// GetListSoftwareUpdatesSortByEnumValues Enumerates the set of values for ListSoftwareUpdatesSortByEnum
func GetListSoftwareUpdatesSortByEnumValues() []ListSoftwareUpdatesSortByEnum {
	values := make([]ListSoftwareUpdatesSortByEnum, 0)
	for _, v := range mappingListSoftwareUpdatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareUpdatesSortByEnumStringValues Enumerates the set of values in String for ListSoftwareUpdatesSortByEnum
func GetListSoftwareUpdatesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSoftwareUpdatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareUpdatesSortByEnum(val string) (ListSoftwareUpdatesSortByEnum, bool) {
	enum, ok := mappingListSoftwareUpdatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareUpdatesSortOrderEnum Enum with underlying type: string
type ListSoftwareUpdatesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareUpdatesSortOrderEnum
const (
	ListSoftwareUpdatesSortOrderAsc  ListSoftwareUpdatesSortOrderEnum = "ASC"
	ListSoftwareUpdatesSortOrderDesc ListSoftwareUpdatesSortOrderEnum = "DESC"
)

var mappingListSoftwareUpdatesSortOrderEnum = map[string]ListSoftwareUpdatesSortOrderEnum{
	"ASC":  ListSoftwareUpdatesSortOrderAsc,
	"DESC": ListSoftwareUpdatesSortOrderDesc,
}

var mappingListSoftwareUpdatesSortOrderEnumLowerCase = map[string]ListSoftwareUpdatesSortOrderEnum{
	"asc":  ListSoftwareUpdatesSortOrderAsc,
	"desc": ListSoftwareUpdatesSortOrderDesc,
}

// GetListSoftwareUpdatesSortOrderEnumValues Enumerates the set of values for ListSoftwareUpdatesSortOrderEnum
func GetListSoftwareUpdatesSortOrderEnumValues() []ListSoftwareUpdatesSortOrderEnum {
	values := make([]ListSoftwareUpdatesSortOrderEnum, 0)
	for _, v := range mappingListSoftwareUpdatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareUpdatesSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwareUpdatesSortOrderEnum
func GetListSoftwareUpdatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwareUpdatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareUpdatesSortOrderEnum(val string) (ListSoftwareUpdatesSortOrderEnum, bool) {
	enum, ok := mappingListSoftwareUpdatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
