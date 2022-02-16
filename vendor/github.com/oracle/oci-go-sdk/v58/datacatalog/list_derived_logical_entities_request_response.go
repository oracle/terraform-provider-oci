// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListDerivedLogicalEntitiesRequest wrapper for the ListDerivedLogicalEntities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListDerivedLogicalEntities.go.html to see an example of how to use ListDerivedLogicalEntitiesRequest.
type ListDerivedLogicalEntitiesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique pattern key.
	PatternKey *string `mandatory:"true" contributesTo:"path" name:"patternKey"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListDerivedLogicalEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDerivedLogicalEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDerivedLogicalEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDerivedLogicalEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDerivedLogicalEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDerivedLogicalEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDerivedLogicalEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDerivedLogicalEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDerivedLogicalEntitiesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDerivedLogicalEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDerivedLogicalEntitiesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDerivedLogicalEntitiesResponse wrapper for the ListDerivedLogicalEntities operation
type ListDerivedLogicalEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EntityCollection instances
	EntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDerivedLogicalEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDerivedLogicalEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDerivedLogicalEntitiesSortByEnum Enum with underlying type: string
type ListDerivedLogicalEntitiesSortByEnum string

// Set of constants representing the allowable values for ListDerivedLogicalEntitiesSortByEnum
const (
	ListDerivedLogicalEntitiesSortByTimecreated ListDerivedLogicalEntitiesSortByEnum = "TIMECREATED"
	ListDerivedLogicalEntitiesSortByDisplayname ListDerivedLogicalEntitiesSortByEnum = "DISPLAYNAME"
)

var mappingListDerivedLogicalEntitiesSortByEnum = map[string]ListDerivedLogicalEntitiesSortByEnum{
	"TIMECREATED": ListDerivedLogicalEntitiesSortByTimecreated,
	"DISPLAYNAME": ListDerivedLogicalEntitiesSortByDisplayname,
}

// GetListDerivedLogicalEntitiesSortByEnumValues Enumerates the set of values for ListDerivedLogicalEntitiesSortByEnum
func GetListDerivedLogicalEntitiesSortByEnumValues() []ListDerivedLogicalEntitiesSortByEnum {
	values := make([]ListDerivedLogicalEntitiesSortByEnum, 0)
	for _, v := range mappingListDerivedLogicalEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDerivedLogicalEntitiesSortByEnumStringValues Enumerates the set of values in String for ListDerivedLogicalEntitiesSortByEnum
func GetListDerivedLogicalEntitiesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDerivedLogicalEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDerivedLogicalEntitiesSortByEnum(val string) (ListDerivedLogicalEntitiesSortByEnum, bool) {
	mappingListDerivedLogicalEntitiesSortByEnumIgnoreCase := make(map[string]ListDerivedLogicalEntitiesSortByEnum)
	for k, v := range mappingListDerivedLogicalEntitiesSortByEnum {
		mappingListDerivedLogicalEntitiesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDerivedLogicalEntitiesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListDerivedLogicalEntitiesSortOrderEnum Enum with underlying type: string
type ListDerivedLogicalEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListDerivedLogicalEntitiesSortOrderEnum
const (
	ListDerivedLogicalEntitiesSortOrderAsc  ListDerivedLogicalEntitiesSortOrderEnum = "ASC"
	ListDerivedLogicalEntitiesSortOrderDesc ListDerivedLogicalEntitiesSortOrderEnum = "DESC"
)

var mappingListDerivedLogicalEntitiesSortOrderEnum = map[string]ListDerivedLogicalEntitiesSortOrderEnum{
	"ASC":  ListDerivedLogicalEntitiesSortOrderAsc,
	"DESC": ListDerivedLogicalEntitiesSortOrderDesc,
}

// GetListDerivedLogicalEntitiesSortOrderEnumValues Enumerates the set of values for ListDerivedLogicalEntitiesSortOrderEnum
func GetListDerivedLogicalEntitiesSortOrderEnumValues() []ListDerivedLogicalEntitiesSortOrderEnum {
	values := make([]ListDerivedLogicalEntitiesSortOrderEnum, 0)
	for _, v := range mappingListDerivedLogicalEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDerivedLogicalEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListDerivedLogicalEntitiesSortOrderEnum
func GetListDerivedLogicalEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDerivedLogicalEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDerivedLogicalEntitiesSortOrderEnum(val string) (ListDerivedLogicalEntitiesSortOrderEnum, bool) {
	mappingListDerivedLogicalEntitiesSortOrderEnumIgnoreCase := make(map[string]ListDerivedLogicalEntitiesSortOrderEnum)
	for k, v := range mappingListDerivedLogicalEntitiesSortOrderEnum {
		mappingListDerivedLogicalEntitiesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListDerivedLogicalEntitiesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
