// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// RetrieveByolRealmAllocationsRequest wrapper for the RetrieveByolRealmAllocations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/RetrieveByolRealmAllocations.go.html to see an example of how to use RetrieveByolRealmAllocationsRequest.
type RetrieveByolRealmAllocationsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL.
	ByolId *string `mandatory:"true" contributesTo:"path" name:"byolId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder RetrieveByolRealmAllocationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending.
	SortBy RetrieveByolRealmAllocationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RetrieveByolRealmAllocationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RetrieveByolRealmAllocationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RetrieveByolRealmAllocationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RetrieveByolRealmAllocationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RetrieveByolRealmAllocationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRetrieveByolRealmAllocationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetRetrieveByolRealmAllocationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRetrieveByolRealmAllocationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetRetrieveByolRealmAllocationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RetrieveByolRealmAllocationsResponse wrapper for the RetrieveByolRealmAllocations operation
type RetrieveByolRealmAllocationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ByolRealmAllocationCollection instances
	ByolRealmAllocationCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response RetrieveByolRealmAllocationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RetrieveByolRealmAllocationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RetrieveByolRealmAllocationsSortOrderEnum Enum with underlying type: string
type RetrieveByolRealmAllocationsSortOrderEnum string

// Set of constants representing the allowable values for RetrieveByolRealmAllocationsSortOrderEnum
const (
	RetrieveByolRealmAllocationsSortOrderAsc  RetrieveByolRealmAllocationsSortOrderEnum = "ASC"
	RetrieveByolRealmAllocationsSortOrderDesc RetrieveByolRealmAllocationsSortOrderEnum = "DESC"
)

var mappingRetrieveByolRealmAllocationsSortOrderEnum = map[string]RetrieveByolRealmAllocationsSortOrderEnum{
	"ASC":  RetrieveByolRealmAllocationsSortOrderAsc,
	"DESC": RetrieveByolRealmAllocationsSortOrderDesc,
}

var mappingRetrieveByolRealmAllocationsSortOrderEnumLowerCase = map[string]RetrieveByolRealmAllocationsSortOrderEnum{
	"asc":  RetrieveByolRealmAllocationsSortOrderAsc,
	"desc": RetrieveByolRealmAllocationsSortOrderDesc,
}

// GetRetrieveByolRealmAllocationsSortOrderEnumValues Enumerates the set of values for RetrieveByolRealmAllocationsSortOrderEnum
func GetRetrieveByolRealmAllocationsSortOrderEnumValues() []RetrieveByolRealmAllocationsSortOrderEnum {
	values := make([]RetrieveByolRealmAllocationsSortOrderEnum, 0)
	for _, v := range mappingRetrieveByolRealmAllocationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetRetrieveByolRealmAllocationsSortOrderEnumStringValues Enumerates the set of values in String for RetrieveByolRealmAllocationsSortOrderEnum
func GetRetrieveByolRealmAllocationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingRetrieveByolRealmAllocationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRetrieveByolRealmAllocationsSortOrderEnum(val string) (RetrieveByolRealmAllocationsSortOrderEnum, bool) {
	enum, ok := mappingRetrieveByolRealmAllocationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RetrieveByolRealmAllocationsSortByEnum Enum with underlying type: string
type RetrieveByolRealmAllocationsSortByEnum string

// Set of constants representing the allowable values for RetrieveByolRealmAllocationsSortByEnum
const (
	RetrieveByolRealmAllocationsSortByTimecreated RetrieveByolRealmAllocationsSortByEnum = "timeCreated"
)

var mappingRetrieveByolRealmAllocationsSortByEnum = map[string]RetrieveByolRealmAllocationsSortByEnum{
	"timeCreated": RetrieveByolRealmAllocationsSortByTimecreated,
}

var mappingRetrieveByolRealmAllocationsSortByEnumLowerCase = map[string]RetrieveByolRealmAllocationsSortByEnum{
	"timecreated": RetrieveByolRealmAllocationsSortByTimecreated,
}

// GetRetrieveByolRealmAllocationsSortByEnumValues Enumerates the set of values for RetrieveByolRealmAllocationsSortByEnum
func GetRetrieveByolRealmAllocationsSortByEnumValues() []RetrieveByolRealmAllocationsSortByEnum {
	values := make([]RetrieveByolRealmAllocationsSortByEnum, 0)
	for _, v := range mappingRetrieveByolRealmAllocationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetRetrieveByolRealmAllocationsSortByEnumStringValues Enumerates the set of values in String for RetrieveByolRealmAllocationsSortByEnum
func GetRetrieveByolRealmAllocationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingRetrieveByolRealmAllocationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRetrieveByolRealmAllocationsSortByEnum(val string) (RetrieveByolRealmAllocationsSortByEnum, bool) {
	enum, ok := mappingRetrieveByolRealmAllocationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
