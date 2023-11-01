// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSoftwareSourceVendorsRequest wrapper for the ListSoftwareSourceVendors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListSoftwareSourceVendors.go.html to see an example of how to use ListSoftwareSourceVendorsRequest.
type ListSoftwareSourceVendorsRequest struct {

	// The OCID of the compartment that contains the resources to list. This parameter is required.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSoftwareSourceVendorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort software source vendors by. Only one sort order may be provided. Default order for name is ascending.
	SortBy ListSoftwareSourceVendorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The name of the entity to be queried.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareSourceVendorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareSourceVendorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwareSourceVendorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareSourceVendorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwareSourceVendorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSoftwareSourceVendorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwareSourceVendorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwareSourceVendorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwareSourceVendorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwareSourceVendorsResponse wrapper for the ListSoftwareSourceVendors operation
type ListSoftwareSourceVendorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SoftwareSourceVendorCollection instance
	SoftwareSourceVendorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSoftwareSourceVendorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareSourceVendorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareSourceVendorsSortOrderEnum Enum with underlying type: string
type ListSoftwareSourceVendorsSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareSourceVendorsSortOrderEnum
const (
	ListSoftwareSourceVendorsSortOrderAsc  ListSoftwareSourceVendorsSortOrderEnum = "ASC"
	ListSoftwareSourceVendorsSortOrderDesc ListSoftwareSourceVendorsSortOrderEnum = "DESC"
)

var mappingListSoftwareSourceVendorsSortOrderEnum = map[string]ListSoftwareSourceVendorsSortOrderEnum{
	"ASC":  ListSoftwareSourceVendorsSortOrderAsc,
	"DESC": ListSoftwareSourceVendorsSortOrderDesc,
}

var mappingListSoftwareSourceVendorsSortOrderEnumLowerCase = map[string]ListSoftwareSourceVendorsSortOrderEnum{
	"asc":  ListSoftwareSourceVendorsSortOrderAsc,
	"desc": ListSoftwareSourceVendorsSortOrderDesc,
}

// GetListSoftwareSourceVendorsSortOrderEnumValues Enumerates the set of values for ListSoftwareSourceVendorsSortOrderEnum
func GetListSoftwareSourceVendorsSortOrderEnumValues() []ListSoftwareSourceVendorsSortOrderEnum {
	values := make([]ListSoftwareSourceVendorsSortOrderEnum, 0)
	for _, v := range mappingListSoftwareSourceVendorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourceVendorsSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwareSourceVendorsSortOrderEnum
func GetListSoftwareSourceVendorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwareSourceVendorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourceVendorsSortOrderEnum(val string) (ListSoftwareSourceVendorsSortOrderEnum, bool) {
	enum, ok := mappingListSoftwareSourceVendorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwareSourceVendorsSortByEnum Enum with underlying type: string
type ListSoftwareSourceVendorsSortByEnum string

// Set of constants representing the allowable values for ListSoftwareSourceVendorsSortByEnum
const (
	ListSoftwareSourceVendorsSortByName ListSoftwareSourceVendorsSortByEnum = "name"
)

var mappingListSoftwareSourceVendorsSortByEnum = map[string]ListSoftwareSourceVendorsSortByEnum{
	"name": ListSoftwareSourceVendorsSortByName,
}

var mappingListSoftwareSourceVendorsSortByEnumLowerCase = map[string]ListSoftwareSourceVendorsSortByEnum{
	"name": ListSoftwareSourceVendorsSortByName,
}

// GetListSoftwareSourceVendorsSortByEnumValues Enumerates the set of values for ListSoftwareSourceVendorsSortByEnum
func GetListSoftwareSourceVendorsSortByEnumValues() []ListSoftwareSourceVendorsSortByEnum {
	values := make([]ListSoftwareSourceVendorsSortByEnum, 0)
	for _, v := range mappingListSoftwareSourceVendorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwareSourceVendorsSortByEnumStringValues Enumerates the set of values in String for ListSoftwareSourceVendorsSortByEnum
func GetListSoftwareSourceVendorsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListSoftwareSourceVendorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwareSourceVendorsSortByEnum(val string) (ListSoftwareSourceVendorsSortByEnum, bool) {
	enum, ok := mappingListSoftwareSourceVendorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
