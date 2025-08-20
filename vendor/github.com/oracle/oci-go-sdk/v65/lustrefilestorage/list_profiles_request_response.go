// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProfilesRequest wrapper for the ListProfiles operation
type ListProfilesRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only profile associated with given name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProfilesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProfilesResponse wrapper for the ListProfiles operation
type ListProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProfileCollection instances
	ProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProfilesSortOrderEnum Enum with underlying type: string
type ListProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListProfilesSortOrderEnum
const (
	ListProfilesSortOrderAsc  ListProfilesSortOrderEnum = "ASC"
	ListProfilesSortOrderDesc ListProfilesSortOrderEnum = "DESC"
)

var mappingListProfilesSortOrderEnum = map[string]ListProfilesSortOrderEnum{
	"ASC":  ListProfilesSortOrderAsc,
	"DESC": ListProfilesSortOrderDesc,
}

var mappingListProfilesSortOrderEnumLowerCase = map[string]ListProfilesSortOrderEnum{
	"asc":  ListProfilesSortOrderAsc,
	"desc": ListProfilesSortOrderDesc,
}

// GetListProfilesSortOrderEnumValues Enumerates the set of values for ListProfilesSortOrderEnum
func GetListProfilesSortOrderEnumValues() []ListProfilesSortOrderEnum {
	values := make([]ListProfilesSortOrderEnum, 0)
	for _, v := range mappingListProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListProfilesSortOrderEnum
func GetListProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfilesSortOrderEnum(val string) (ListProfilesSortOrderEnum, bool) {
	enum, ok := mappingListProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
