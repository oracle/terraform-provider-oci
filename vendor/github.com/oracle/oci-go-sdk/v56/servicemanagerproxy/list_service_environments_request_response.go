// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemanagerproxy

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
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

var mappingListServiceEnvironmentsSortBy = map[string]ListServiceEnvironmentsSortByEnum{
	"ID": ListServiceEnvironmentsSortById,
}

// GetListServiceEnvironmentsSortByEnumValues Enumerates the set of values for ListServiceEnvironmentsSortByEnum
func GetListServiceEnvironmentsSortByEnumValues() []ListServiceEnvironmentsSortByEnum {
	values := make([]ListServiceEnvironmentsSortByEnum, 0)
	for _, v := range mappingListServiceEnvironmentsSortBy {
		values = append(values, v)
	}
	return values
}

// ListServiceEnvironmentsSortOrderEnum Enum with underlying type: string
type ListServiceEnvironmentsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceEnvironmentsSortOrderEnum
const (
	ListServiceEnvironmentsSortOrderAsc  ListServiceEnvironmentsSortOrderEnum = "ASC"
	ListServiceEnvironmentsSortOrderDesc ListServiceEnvironmentsSortOrderEnum = "DESC"
)

var mappingListServiceEnvironmentsSortOrder = map[string]ListServiceEnvironmentsSortOrderEnum{
	"ASC":  ListServiceEnvironmentsSortOrderAsc,
	"DESC": ListServiceEnvironmentsSortOrderDesc,
}

// GetListServiceEnvironmentsSortOrderEnumValues Enumerates the set of values for ListServiceEnvironmentsSortOrderEnum
func GetListServiceEnvironmentsSortOrderEnumValues() []ListServiceEnvironmentsSortOrderEnum {
	values := make([]ListServiceEnvironmentsSortOrderEnum, 0)
	for _, v := range mappingListServiceEnvironmentsSortOrder {
		values = append(values, v)
	}
	return values
}
