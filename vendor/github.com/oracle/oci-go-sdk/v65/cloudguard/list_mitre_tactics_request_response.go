// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMitreTacticsRequest wrapper for the ListMitreTactics operation
type ListMitreTacticsRequest struct {

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMitreTacticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMitreTacticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMitreTacticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMitreTacticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMitreTacticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMitreTacticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMitreTacticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMitreTacticsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMitreTacticsResponse wrapper for the ListMitreTactics operation
type ListMitreTacticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MitreTacticCollection instances
	MitreTacticCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMitreTacticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMitreTacticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMitreTacticsSortOrderEnum Enum with underlying type: string
type ListMitreTacticsSortOrderEnum string

// Set of constants representing the allowable values for ListMitreTacticsSortOrderEnum
const (
	ListMitreTacticsSortOrderAsc  ListMitreTacticsSortOrderEnum = "ASC"
	ListMitreTacticsSortOrderDesc ListMitreTacticsSortOrderEnum = "DESC"
)

var mappingListMitreTacticsSortOrderEnum = map[string]ListMitreTacticsSortOrderEnum{
	"ASC":  ListMitreTacticsSortOrderAsc,
	"DESC": ListMitreTacticsSortOrderDesc,
}

var mappingListMitreTacticsSortOrderEnumLowerCase = map[string]ListMitreTacticsSortOrderEnum{
	"asc":  ListMitreTacticsSortOrderAsc,
	"desc": ListMitreTacticsSortOrderDesc,
}

// GetListMitreTacticsSortOrderEnumValues Enumerates the set of values for ListMitreTacticsSortOrderEnum
func GetListMitreTacticsSortOrderEnumValues() []ListMitreTacticsSortOrderEnum {
	values := make([]ListMitreTacticsSortOrderEnum, 0)
	for _, v := range mappingListMitreTacticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMitreTacticsSortOrderEnumStringValues Enumerates the set of values in String for ListMitreTacticsSortOrderEnum
func GetListMitreTacticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMitreTacticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMitreTacticsSortOrderEnum(val string) (ListMitreTacticsSortOrderEnum, bool) {
	enum, ok := mappingListMitreTacticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
