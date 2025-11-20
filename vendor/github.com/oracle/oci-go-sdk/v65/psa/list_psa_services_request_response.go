// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPsaServicesRequest wrapper for the ListPsaServices operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPsaServices.go.html to see an example of how to use ListPsaServicesRequest.
type ListPsaServicesRequest struct {

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The unique identifier of the OCI service.
	ServiceId *string `mandatory:"false" contributesTo:"query" name:"serviceId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for `displayName` is descending.
	SortBy ListPsaServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPsaServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPsaServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPsaServicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPsaServicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPsaServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPsaServicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPsaServicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPsaServicesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPsaServicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPsaServicesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPsaServicesResponse wrapper for the ListPsaServices operation
type ListPsaServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PsaServiceCollection instances
	PsaServiceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPsaServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPsaServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPsaServicesSortByEnum Enum with underlying type: string
type ListPsaServicesSortByEnum string

// Set of constants representing the allowable values for ListPsaServicesSortByEnum
const (
	ListPsaServicesSortByDisplayname ListPsaServicesSortByEnum = "displayName"
)

var mappingListPsaServicesSortByEnum = map[string]ListPsaServicesSortByEnum{
	"displayName": ListPsaServicesSortByDisplayname,
}

var mappingListPsaServicesSortByEnumLowerCase = map[string]ListPsaServicesSortByEnum{
	"displayname": ListPsaServicesSortByDisplayname,
}

// GetListPsaServicesSortByEnumValues Enumerates the set of values for ListPsaServicesSortByEnum
func GetListPsaServicesSortByEnumValues() []ListPsaServicesSortByEnum {
	values := make([]ListPsaServicesSortByEnum, 0)
	for _, v := range mappingListPsaServicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPsaServicesSortByEnumStringValues Enumerates the set of values in String for ListPsaServicesSortByEnum
func GetListPsaServicesSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListPsaServicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPsaServicesSortByEnum(val string) (ListPsaServicesSortByEnum, bool) {
	enum, ok := mappingListPsaServicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPsaServicesSortOrderEnum Enum with underlying type: string
type ListPsaServicesSortOrderEnum string

// Set of constants representing the allowable values for ListPsaServicesSortOrderEnum
const (
	ListPsaServicesSortOrderAsc  ListPsaServicesSortOrderEnum = "ASC"
	ListPsaServicesSortOrderDesc ListPsaServicesSortOrderEnum = "DESC"
)

var mappingListPsaServicesSortOrderEnum = map[string]ListPsaServicesSortOrderEnum{
	"ASC":  ListPsaServicesSortOrderAsc,
	"DESC": ListPsaServicesSortOrderDesc,
}

var mappingListPsaServicesSortOrderEnumLowerCase = map[string]ListPsaServicesSortOrderEnum{
	"asc":  ListPsaServicesSortOrderAsc,
	"desc": ListPsaServicesSortOrderDesc,
}

// GetListPsaServicesSortOrderEnumValues Enumerates the set of values for ListPsaServicesSortOrderEnum
func GetListPsaServicesSortOrderEnumValues() []ListPsaServicesSortOrderEnum {
	values := make([]ListPsaServicesSortOrderEnum, 0)
	for _, v := range mappingListPsaServicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPsaServicesSortOrderEnumStringValues Enumerates the set of values in String for ListPsaServicesSortOrderEnum
func GetListPsaServicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPsaServicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPsaServicesSortOrderEnum(val string) (ListPsaServicesSortOrderEnum, bool) {
	enum, ok := mappingListPsaServicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
