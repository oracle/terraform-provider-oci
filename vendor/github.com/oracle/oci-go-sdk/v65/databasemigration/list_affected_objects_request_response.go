// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAffectedObjectsRequest wrapper for the ListAffectedObjects operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAffectedObjects.go.html to see an example of how to use ListAffectedObjectsRequest.
type ListAffectedObjectsRequest struct {

	// The OCID of the Assessment
	AssessmentId *string `mandatory:"true" contributesTo:"path" name:"assessmentId"`

	// The name of the Assessor
	AssessorName *string `mandatory:"true" contributesTo:"path" name:"assessorName"`

	// The Name of the assessor check
	CheckName *string `mandatory:"true" contributesTo:"path" name:"checkName"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for name is custom based on it's usage frequency. If no value is specified name is default.
	SortBy ListAffectedObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAffectedObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAffectedObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAffectedObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAffectedObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAffectedObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAffectedObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAffectedObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAffectedObjectsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAffectedObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAffectedObjectsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAffectedObjectsResponse wrapper for the ListAffectedObjects operation
type ListAffectedObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AffectedObjectsCollection instances
	AffectedObjectsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAffectedObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAffectedObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAffectedObjectsSortByEnum Enum with underlying type: string
type ListAffectedObjectsSortByEnum string

// Set of constants representing the allowable values for ListAffectedObjectsSortByEnum
const (
	ListAffectedObjectsSortByName ListAffectedObjectsSortByEnum = "name"
)

var mappingListAffectedObjectsSortByEnum = map[string]ListAffectedObjectsSortByEnum{
	"name": ListAffectedObjectsSortByName,
}

var mappingListAffectedObjectsSortByEnumLowerCase = map[string]ListAffectedObjectsSortByEnum{
	"name": ListAffectedObjectsSortByName,
}

// GetListAffectedObjectsSortByEnumValues Enumerates the set of values for ListAffectedObjectsSortByEnum
func GetListAffectedObjectsSortByEnumValues() []ListAffectedObjectsSortByEnum {
	values := make([]ListAffectedObjectsSortByEnum, 0)
	for _, v := range mappingListAffectedObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAffectedObjectsSortByEnumStringValues Enumerates the set of values in String for ListAffectedObjectsSortByEnum
func GetListAffectedObjectsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListAffectedObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAffectedObjectsSortByEnum(val string) (ListAffectedObjectsSortByEnum, bool) {
	enum, ok := mappingListAffectedObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAffectedObjectsSortOrderEnum Enum with underlying type: string
type ListAffectedObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListAffectedObjectsSortOrderEnum
const (
	ListAffectedObjectsSortOrderAsc  ListAffectedObjectsSortOrderEnum = "ASC"
	ListAffectedObjectsSortOrderDesc ListAffectedObjectsSortOrderEnum = "DESC"
)

var mappingListAffectedObjectsSortOrderEnum = map[string]ListAffectedObjectsSortOrderEnum{
	"ASC":  ListAffectedObjectsSortOrderAsc,
	"DESC": ListAffectedObjectsSortOrderDesc,
}

var mappingListAffectedObjectsSortOrderEnumLowerCase = map[string]ListAffectedObjectsSortOrderEnum{
	"asc":  ListAffectedObjectsSortOrderAsc,
	"desc": ListAffectedObjectsSortOrderDesc,
}

// GetListAffectedObjectsSortOrderEnumValues Enumerates the set of values for ListAffectedObjectsSortOrderEnum
func GetListAffectedObjectsSortOrderEnumValues() []ListAffectedObjectsSortOrderEnum {
	values := make([]ListAffectedObjectsSortOrderEnum, 0)
	for _, v := range mappingListAffectedObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAffectedObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListAffectedObjectsSortOrderEnum
func GetListAffectedObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAffectedObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAffectedObjectsSortOrderEnum(val string) (ListAffectedObjectsSortOrderEnum, bool) {
	enum, ok := mappingListAffectedObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
