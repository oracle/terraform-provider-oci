// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProfileAvailableSoftwareSourcesRequest wrapper for the ListProfileAvailableSoftwareSources operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListProfileAvailableSoftwareSources.go.html to see an example of how to use ListProfileAvailableSoftwareSourcesRequest.
type ListProfileAvailableSoftwareSourcesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.
	ProfileId *string `mandatory:"true" contributesTo:"path" name:"profileId"`

	// A filter to return resources that match the given display names.
	DisplayName []string `contributesTo:"query" name:"displayName" collectionFormat:"multi"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListProfileAvailableSoftwareSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending.
	// Default order for displayName is ascending.
	SortBy ListProfileAvailableSoftwareSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProfileAvailableSoftwareSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProfileAvailableSoftwareSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProfileAvailableSoftwareSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProfileAvailableSoftwareSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProfileAvailableSoftwareSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProfileAvailableSoftwareSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProfileAvailableSoftwareSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProfileAvailableSoftwareSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProfileAvailableSoftwareSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProfileAvailableSoftwareSourcesResponse wrapper for the ListProfileAvailableSoftwareSources operation
type ListProfileAvailableSoftwareSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AvailableSoftwareSourceCollection instances
	AvailableSoftwareSourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProfileAvailableSoftwareSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProfileAvailableSoftwareSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProfileAvailableSoftwareSourcesSortOrderEnum Enum with underlying type: string
type ListProfileAvailableSoftwareSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListProfileAvailableSoftwareSourcesSortOrderEnum
const (
	ListProfileAvailableSoftwareSourcesSortOrderAsc  ListProfileAvailableSoftwareSourcesSortOrderEnum = "ASC"
	ListProfileAvailableSoftwareSourcesSortOrderDesc ListProfileAvailableSoftwareSourcesSortOrderEnum = "DESC"
)

var mappingListProfileAvailableSoftwareSourcesSortOrderEnum = map[string]ListProfileAvailableSoftwareSourcesSortOrderEnum{
	"ASC":  ListProfileAvailableSoftwareSourcesSortOrderAsc,
	"DESC": ListProfileAvailableSoftwareSourcesSortOrderDesc,
}

var mappingListProfileAvailableSoftwareSourcesSortOrderEnumLowerCase = map[string]ListProfileAvailableSoftwareSourcesSortOrderEnum{
	"asc":  ListProfileAvailableSoftwareSourcesSortOrderAsc,
	"desc": ListProfileAvailableSoftwareSourcesSortOrderDesc,
}

// GetListProfileAvailableSoftwareSourcesSortOrderEnumValues Enumerates the set of values for ListProfileAvailableSoftwareSourcesSortOrderEnum
func GetListProfileAvailableSoftwareSourcesSortOrderEnumValues() []ListProfileAvailableSoftwareSourcesSortOrderEnum {
	values := make([]ListProfileAvailableSoftwareSourcesSortOrderEnum, 0)
	for _, v := range mappingListProfileAvailableSoftwareSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfileAvailableSoftwareSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListProfileAvailableSoftwareSourcesSortOrderEnum
func GetListProfileAvailableSoftwareSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProfileAvailableSoftwareSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfileAvailableSoftwareSourcesSortOrderEnum(val string) (ListProfileAvailableSoftwareSourcesSortOrderEnum, bool) {
	enum, ok := mappingListProfileAvailableSoftwareSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfileAvailableSoftwareSourcesSortByEnum Enum with underlying type: string
type ListProfileAvailableSoftwareSourcesSortByEnum string

// Set of constants representing the allowable values for ListProfileAvailableSoftwareSourcesSortByEnum
const (
	ListProfileAvailableSoftwareSourcesSortByTimecreated ListProfileAvailableSoftwareSourcesSortByEnum = "timeCreated"
	ListProfileAvailableSoftwareSourcesSortByDisplayname ListProfileAvailableSoftwareSourcesSortByEnum = "displayName"
)

var mappingListProfileAvailableSoftwareSourcesSortByEnum = map[string]ListProfileAvailableSoftwareSourcesSortByEnum{
	"timeCreated": ListProfileAvailableSoftwareSourcesSortByTimecreated,
	"displayName": ListProfileAvailableSoftwareSourcesSortByDisplayname,
}

var mappingListProfileAvailableSoftwareSourcesSortByEnumLowerCase = map[string]ListProfileAvailableSoftwareSourcesSortByEnum{
	"timecreated": ListProfileAvailableSoftwareSourcesSortByTimecreated,
	"displayname": ListProfileAvailableSoftwareSourcesSortByDisplayname,
}

// GetListProfileAvailableSoftwareSourcesSortByEnumValues Enumerates the set of values for ListProfileAvailableSoftwareSourcesSortByEnum
func GetListProfileAvailableSoftwareSourcesSortByEnumValues() []ListProfileAvailableSoftwareSourcesSortByEnum {
	values := make([]ListProfileAvailableSoftwareSourcesSortByEnum, 0)
	for _, v := range mappingListProfileAvailableSoftwareSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfileAvailableSoftwareSourcesSortByEnumStringValues Enumerates the set of values in String for ListProfileAvailableSoftwareSourcesSortByEnum
func GetListProfileAvailableSoftwareSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListProfileAvailableSoftwareSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfileAvailableSoftwareSourcesSortByEnum(val string) (ListProfileAvailableSoftwareSourcesSortByEnum, bool) {
	enum, ok := mappingListProfileAvailableSoftwareSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
