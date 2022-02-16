// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemanagerproxy

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListServiceEnvironmentsRequest wrapper for the ListServiceEnvironments operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemanagerproxy/ListServiceEnvironments.go.html to see an example of how to use ListServiceEnvironmentsRequest.
type ListServiceEnvironmentsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique identifier associated with the service environment.
	// **Note:** Not an OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ServiceEnvironmentId *string `mandatory:"false" contributesTo:"query" name:"serviceEnvironmentId"`

	// The environment's service definition type.
	// For example, "RGBUOROMS" is the service definition type for "Oracle Retail Order Management Cloud Service".
	ServiceEnvironmentType *string `mandatory:"false" contributesTo:"query" name:"serviceEnvironmentType"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. ID is default ordered as ascending.
	SortBy ListServiceEnvironmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either `ASC` or `DESC`.
	SortOrder ListServiceEnvironmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The display name of the resource.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServiceEnvironmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServiceEnvironmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListServiceEnvironmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServiceEnvironmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListServiceEnvironmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListServiceEnvironmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListServiceEnvironmentsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceEnvironmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListServiceEnvironmentsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListServiceEnvironmentsResponse wrapper for the ListServiceEnvironments operation
type ListServiceEnvironmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServiceEnvironmentCollection instances
	ServiceEnvironmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListServiceEnvironmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServiceEnvironmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServiceEnvironmentsSortByEnum Enum with underlying type: string
type ListServiceEnvironmentsSortByEnum string

// Set of constants representing the allowable values for ListServiceEnvironmentsSortByEnum
const (
	ListServiceEnvironmentsSortById ListServiceEnvironmentsSortByEnum = "ID"
)

var mappingListServiceEnvironmentsSortByEnum = map[string]ListServiceEnvironmentsSortByEnum{
	"ID": ListServiceEnvironmentsSortById,
}

// GetListServiceEnvironmentsSortByEnumValues Enumerates the set of values for ListServiceEnvironmentsSortByEnum
func GetListServiceEnvironmentsSortByEnumValues() []ListServiceEnvironmentsSortByEnum {
	values := make([]ListServiceEnvironmentsSortByEnum, 0)
	for _, v := range mappingListServiceEnvironmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceEnvironmentsSortByEnumStringValues Enumerates the set of values in String for ListServiceEnvironmentsSortByEnum
func GetListServiceEnvironmentsSortByEnumStringValues() []string {
	return []string{
		"ID",
	}
}

// GetMappingListServiceEnvironmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceEnvironmentsSortByEnum(val string) (ListServiceEnvironmentsSortByEnum, bool) {
	mappingListServiceEnvironmentsSortByEnumIgnoreCase := make(map[string]ListServiceEnvironmentsSortByEnum)
	for k, v := range mappingListServiceEnvironmentsSortByEnum {
		mappingListServiceEnvironmentsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListServiceEnvironmentsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceEnvironmentsSortOrderEnum Enum with underlying type: string
type ListServiceEnvironmentsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceEnvironmentsSortOrderEnum
const (
	ListServiceEnvironmentsSortOrderAsc  ListServiceEnvironmentsSortOrderEnum = "ASC"
	ListServiceEnvironmentsSortOrderDesc ListServiceEnvironmentsSortOrderEnum = "DESC"
)

var mappingListServiceEnvironmentsSortOrderEnum = map[string]ListServiceEnvironmentsSortOrderEnum{
	"ASC":  ListServiceEnvironmentsSortOrderAsc,
	"DESC": ListServiceEnvironmentsSortOrderDesc,
}

// GetListServiceEnvironmentsSortOrderEnumValues Enumerates the set of values for ListServiceEnvironmentsSortOrderEnum
func GetListServiceEnvironmentsSortOrderEnumValues() []ListServiceEnvironmentsSortOrderEnum {
	values := make([]ListServiceEnvironmentsSortOrderEnum, 0)
	for _, v := range mappingListServiceEnvironmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceEnvironmentsSortOrderEnumStringValues Enumerates the set of values in String for ListServiceEnvironmentsSortOrderEnum
func GetListServiceEnvironmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListServiceEnvironmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceEnvironmentsSortOrderEnum(val string) (ListServiceEnvironmentsSortOrderEnum, bool) {
	mappingListServiceEnvironmentsSortOrderEnumIgnoreCase := make(map[string]ListServiceEnvironmentsSortOrderEnum)
	for k, v := range mappingListServiceEnvironmentsSortOrderEnum {
		mappingListServiceEnvironmentsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListServiceEnvironmentsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
