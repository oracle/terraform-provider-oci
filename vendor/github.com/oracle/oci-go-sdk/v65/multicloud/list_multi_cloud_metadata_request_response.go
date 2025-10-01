// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMultiCloudMetadataRequest wrapper for the ListMultiCloudMetadata operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListMultiCloudMetadata.go.html to see an example of how to use ListMultiCloudMetadataRequest.
type ListMultiCloudMetadataRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMultiCloudMetadataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListMultiCloudMetadataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMultiCloudMetadataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMultiCloudMetadataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMultiCloudMetadataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMultiCloudMetadataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMultiCloudMetadataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMultiCloudMetadataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMultiCloudMetadataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMultiCloudMetadataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMultiCloudMetadataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMultiCloudMetadataResponse wrapper for the ListMultiCloudMetadata operation
type ListMultiCloudMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MultiCloudMetadataCollection instances
	MultiCloudMetadataCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMultiCloudMetadataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMultiCloudMetadataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMultiCloudMetadataSortOrderEnum Enum with underlying type: string
type ListMultiCloudMetadataSortOrderEnum string

// Set of constants representing the allowable values for ListMultiCloudMetadataSortOrderEnum
const (
	ListMultiCloudMetadataSortOrderAsc  ListMultiCloudMetadataSortOrderEnum = "ASC"
	ListMultiCloudMetadataSortOrderDesc ListMultiCloudMetadataSortOrderEnum = "DESC"
)

var mappingListMultiCloudMetadataSortOrderEnum = map[string]ListMultiCloudMetadataSortOrderEnum{
	"ASC":  ListMultiCloudMetadataSortOrderAsc,
	"DESC": ListMultiCloudMetadataSortOrderDesc,
}

var mappingListMultiCloudMetadataSortOrderEnumLowerCase = map[string]ListMultiCloudMetadataSortOrderEnum{
	"asc":  ListMultiCloudMetadataSortOrderAsc,
	"desc": ListMultiCloudMetadataSortOrderDesc,
}

// GetListMultiCloudMetadataSortOrderEnumValues Enumerates the set of values for ListMultiCloudMetadataSortOrderEnum
func GetListMultiCloudMetadataSortOrderEnumValues() []ListMultiCloudMetadataSortOrderEnum {
	values := make([]ListMultiCloudMetadataSortOrderEnum, 0)
	for _, v := range mappingListMultiCloudMetadataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMultiCloudMetadataSortOrderEnumStringValues Enumerates the set of values in String for ListMultiCloudMetadataSortOrderEnum
func GetListMultiCloudMetadataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMultiCloudMetadataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMultiCloudMetadataSortOrderEnum(val string) (ListMultiCloudMetadataSortOrderEnum, bool) {
	enum, ok := mappingListMultiCloudMetadataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMultiCloudMetadataSortByEnum Enum with underlying type: string
type ListMultiCloudMetadataSortByEnum string

// Set of constants representing the allowable values for ListMultiCloudMetadataSortByEnum
const (
	ListMultiCloudMetadataSortByTimecreated ListMultiCloudMetadataSortByEnum = "timeCreated"
	ListMultiCloudMetadataSortByDisplayname ListMultiCloudMetadataSortByEnum = "displayName"
)

var mappingListMultiCloudMetadataSortByEnum = map[string]ListMultiCloudMetadataSortByEnum{
	"timeCreated": ListMultiCloudMetadataSortByTimecreated,
	"displayName": ListMultiCloudMetadataSortByDisplayname,
}

var mappingListMultiCloudMetadataSortByEnumLowerCase = map[string]ListMultiCloudMetadataSortByEnum{
	"timecreated": ListMultiCloudMetadataSortByTimecreated,
	"displayname": ListMultiCloudMetadataSortByDisplayname,
}

// GetListMultiCloudMetadataSortByEnumValues Enumerates the set of values for ListMultiCloudMetadataSortByEnum
func GetListMultiCloudMetadataSortByEnumValues() []ListMultiCloudMetadataSortByEnum {
	values := make([]ListMultiCloudMetadataSortByEnum, 0)
	for _, v := range mappingListMultiCloudMetadataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMultiCloudMetadataSortByEnumStringValues Enumerates the set of values in String for ListMultiCloudMetadataSortByEnum
func GetListMultiCloudMetadataSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMultiCloudMetadataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMultiCloudMetadataSortByEnum(val string) (ListMultiCloudMetadataSortByEnum, bool) {
	enum, ok := mappingListMultiCloudMetadataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
