// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOsPatchesRequest wrapper for the ListOsPatches operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListOsPatches.go.html to see an example of how to use ListOsPatchesRequest.
type ListOsPatchesRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOsPatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOsPatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListOsPatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOsPatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOsPatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOsPatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOsPatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOsPatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOsPatchesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOsPatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOsPatchesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOsPatchesResponse wrapper for the ListOsPatches operation
type ListOsPatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OsPatchSummary instances
	Items []OsPatchSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response ListOsPatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOsPatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOsPatchesSortByEnum Enum with underlying type: string
type ListOsPatchesSortByEnum string

// Set of constants representing the allowable values for ListOsPatchesSortByEnum
const (
	ListOsPatchesSortByTimecreated ListOsPatchesSortByEnum = "timeCreated"
	ListOsPatchesSortByDisplayname ListOsPatchesSortByEnum = "displayName"
)

var mappingListOsPatchesSortByEnum = map[string]ListOsPatchesSortByEnum{
	"timeCreated": ListOsPatchesSortByTimecreated,
	"displayName": ListOsPatchesSortByDisplayname,
}

var mappingListOsPatchesSortByEnumLowerCase = map[string]ListOsPatchesSortByEnum{
	"timecreated": ListOsPatchesSortByTimecreated,
	"displayname": ListOsPatchesSortByDisplayname,
}

// GetListOsPatchesSortByEnumValues Enumerates the set of values for ListOsPatchesSortByEnum
func GetListOsPatchesSortByEnumValues() []ListOsPatchesSortByEnum {
	values := make([]ListOsPatchesSortByEnum, 0)
	for _, v := range mappingListOsPatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOsPatchesSortByEnumStringValues Enumerates the set of values in String for ListOsPatchesSortByEnum
func GetListOsPatchesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOsPatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOsPatchesSortByEnum(val string) (ListOsPatchesSortByEnum, bool) {
	enum, ok := mappingListOsPatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOsPatchesSortOrderEnum Enum with underlying type: string
type ListOsPatchesSortOrderEnum string

// Set of constants representing the allowable values for ListOsPatchesSortOrderEnum
const (
	ListOsPatchesSortOrderAsc  ListOsPatchesSortOrderEnum = "ASC"
	ListOsPatchesSortOrderDesc ListOsPatchesSortOrderEnum = "DESC"
)

var mappingListOsPatchesSortOrderEnum = map[string]ListOsPatchesSortOrderEnum{
	"ASC":  ListOsPatchesSortOrderAsc,
	"DESC": ListOsPatchesSortOrderDesc,
}

var mappingListOsPatchesSortOrderEnumLowerCase = map[string]ListOsPatchesSortOrderEnum{
	"asc":  ListOsPatchesSortOrderAsc,
	"desc": ListOsPatchesSortOrderDesc,
}

// GetListOsPatchesSortOrderEnumValues Enumerates the set of values for ListOsPatchesSortOrderEnum
func GetListOsPatchesSortOrderEnumValues() []ListOsPatchesSortOrderEnum {
	values := make([]ListOsPatchesSortOrderEnum, 0)
	for _, v := range mappingListOsPatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOsPatchesSortOrderEnumStringValues Enumerates the set of values in String for ListOsPatchesSortOrderEnum
func GetListOsPatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOsPatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOsPatchesSortOrderEnum(val string) (ListOsPatchesSortOrderEnum, bool) {
	enum, ok := mappingListOsPatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
