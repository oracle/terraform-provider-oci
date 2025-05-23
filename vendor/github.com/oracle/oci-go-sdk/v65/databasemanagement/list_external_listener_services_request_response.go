// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExternalListenerServicesRequest wrapper for the ListExternalListenerServices operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalListenerServices.go.html to see an example of how to use ListExternalListenerServicesRequest.
type ListExternalListenerServicesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external listener.
	ExternalListenerId *string `mandatory:"true" contributesTo:"path" name:"externalListenerId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"query" name:"managedDatabaseId"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The
	// default sort order for `NAME` is ascending and it is case-sensitive.
	SortBy ListExternalListenerServicesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalListenerServicesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalListenerServicesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalListenerServicesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalListenerServicesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalListenerServicesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalListenerServicesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalListenerServicesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalListenerServicesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalListenerServicesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalListenerServicesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalListenerServicesResponse wrapper for the ListExternalListenerServices operation
type ListExternalListenerServicesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalListenerServiceCollection instances
	ExternalListenerServiceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalListenerServicesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalListenerServicesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalListenerServicesSortByEnum Enum with underlying type: string
type ListExternalListenerServicesSortByEnum string

// Set of constants representing the allowable values for ListExternalListenerServicesSortByEnum
const (
	ListExternalListenerServicesSortByName ListExternalListenerServicesSortByEnum = "NAME"
)

var mappingListExternalListenerServicesSortByEnum = map[string]ListExternalListenerServicesSortByEnum{
	"NAME": ListExternalListenerServicesSortByName,
}

var mappingListExternalListenerServicesSortByEnumLowerCase = map[string]ListExternalListenerServicesSortByEnum{
	"name": ListExternalListenerServicesSortByName,
}

// GetListExternalListenerServicesSortByEnumValues Enumerates the set of values for ListExternalListenerServicesSortByEnum
func GetListExternalListenerServicesSortByEnumValues() []ListExternalListenerServicesSortByEnum {
	values := make([]ListExternalListenerServicesSortByEnum, 0)
	for _, v := range mappingListExternalListenerServicesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalListenerServicesSortByEnumStringValues Enumerates the set of values in String for ListExternalListenerServicesSortByEnum
func GetListExternalListenerServicesSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListExternalListenerServicesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalListenerServicesSortByEnum(val string) (ListExternalListenerServicesSortByEnum, bool) {
	enum, ok := mappingListExternalListenerServicesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalListenerServicesSortOrderEnum Enum with underlying type: string
type ListExternalListenerServicesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalListenerServicesSortOrderEnum
const (
	ListExternalListenerServicesSortOrderAsc  ListExternalListenerServicesSortOrderEnum = "ASC"
	ListExternalListenerServicesSortOrderDesc ListExternalListenerServicesSortOrderEnum = "DESC"
)

var mappingListExternalListenerServicesSortOrderEnum = map[string]ListExternalListenerServicesSortOrderEnum{
	"ASC":  ListExternalListenerServicesSortOrderAsc,
	"DESC": ListExternalListenerServicesSortOrderDesc,
}

var mappingListExternalListenerServicesSortOrderEnumLowerCase = map[string]ListExternalListenerServicesSortOrderEnum{
	"asc":  ListExternalListenerServicesSortOrderAsc,
	"desc": ListExternalListenerServicesSortOrderDesc,
}

// GetListExternalListenerServicesSortOrderEnumValues Enumerates the set of values for ListExternalListenerServicesSortOrderEnum
func GetListExternalListenerServicesSortOrderEnumValues() []ListExternalListenerServicesSortOrderEnum {
	values := make([]ListExternalListenerServicesSortOrderEnum, 0)
	for _, v := range mappingListExternalListenerServicesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalListenerServicesSortOrderEnumStringValues Enumerates the set of values in String for ListExternalListenerServicesSortOrderEnum
func GetListExternalListenerServicesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalListenerServicesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalListenerServicesSortOrderEnum(val string) (ListExternalListenerServicesSortOrderEnum, bool) {
	enum, ok := mappingListExternalListenerServicesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
