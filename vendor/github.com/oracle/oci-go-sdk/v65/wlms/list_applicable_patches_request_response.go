// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListApplicablePatchesRequest wrapper for the ListApplicablePatches operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListApplicablePatches.go.html to see an example of how to use ListApplicablePatchesRequest.
type ListApplicablePatchesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" contributesTo:"path" name:"wlsDomainId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListApplicablePatchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified, _displayName_ is default.
	SortBy ListApplicablePatchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplicablePatchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplicablePatchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplicablePatchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplicablePatchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApplicablePatchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListApplicablePatchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApplicablePatchesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplicablePatchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApplicablePatchesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApplicablePatchesResponse wrapper for the ListApplicablePatches operation
type ListApplicablePatchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApplicablePatchCollection instances
	ApplicablePatchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApplicablePatchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicablePatchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicablePatchesSortOrderEnum Enum with underlying type: string
type ListApplicablePatchesSortOrderEnum string

// Set of constants representing the allowable values for ListApplicablePatchesSortOrderEnum
const (
	ListApplicablePatchesSortOrderAsc  ListApplicablePatchesSortOrderEnum = "ASC"
	ListApplicablePatchesSortOrderDesc ListApplicablePatchesSortOrderEnum = "DESC"
)

var mappingListApplicablePatchesSortOrderEnum = map[string]ListApplicablePatchesSortOrderEnum{
	"ASC":  ListApplicablePatchesSortOrderAsc,
	"DESC": ListApplicablePatchesSortOrderDesc,
}

var mappingListApplicablePatchesSortOrderEnumLowerCase = map[string]ListApplicablePatchesSortOrderEnum{
	"asc":  ListApplicablePatchesSortOrderAsc,
	"desc": ListApplicablePatchesSortOrderDesc,
}

// GetListApplicablePatchesSortOrderEnumValues Enumerates the set of values for ListApplicablePatchesSortOrderEnum
func GetListApplicablePatchesSortOrderEnumValues() []ListApplicablePatchesSortOrderEnum {
	values := make([]ListApplicablePatchesSortOrderEnum, 0)
	for _, v := range mappingListApplicablePatchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicablePatchesSortOrderEnumStringValues Enumerates the set of values in String for ListApplicablePatchesSortOrderEnum
func GetListApplicablePatchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApplicablePatchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicablePatchesSortOrderEnum(val string) (ListApplicablePatchesSortOrderEnum, bool) {
	enum, ok := mappingListApplicablePatchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApplicablePatchesSortByEnum Enum with underlying type: string
type ListApplicablePatchesSortByEnum string

// Set of constants representing the allowable values for ListApplicablePatchesSortByEnum
const (
	ListApplicablePatchesSortByDisplayname ListApplicablePatchesSortByEnum = "displayName"
)

var mappingListApplicablePatchesSortByEnum = map[string]ListApplicablePatchesSortByEnum{
	"displayName": ListApplicablePatchesSortByDisplayname,
}

var mappingListApplicablePatchesSortByEnumLowerCase = map[string]ListApplicablePatchesSortByEnum{
	"displayname": ListApplicablePatchesSortByDisplayname,
}

// GetListApplicablePatchesSortByEnumValues Enumerates the set of values for ListApplicablePatchesSortByEnum
func GetListApplicablePatchesSortByEnumValues() []ListApplicablePatchesSortByEnum {
	values := make([]ListApplicablePatchesSortByEnum, 0)
	for _, v := range mappingListApplicablePatchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicablePatchesSortByEnumStringValues Enumerates the set of values in String for ListApplicablePatchesSortByEnum
func GetListApplicablePatchesSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListApplicablePatchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicablePatchesSortByEnum(val string) (ListApplicablePatchesSortByEnum, bool) {
	enum, ok := mappingListApplicablePatchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
