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

// ListPrivateServiceAccessesRequest wrapper for the ListPrivateServiceAccesses operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPrivateServiceAccesses.go.html to see an example of how to use ListPrivateServiceAccessesRequest.
type ListPrivateServiceAccessesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState PrivateServiceAccessLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

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

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListPrivateServiceAccessesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPrivateServiceAccessesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPrivateServiceAccessesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPrivateServiceAccessesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPrivateServiceAccessesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPrivateServiceAccessesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPrivateServiceAccessesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateServiceAccessLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPrivateServiceAccessLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivateServiceAccessesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPrivateServiceAccessesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivateServiceAccessesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPrivateServiceAccessesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPrivateServiceAccessesResponse wrapper for the ListPrivateServiceAccesses operation
type ListPrivateServiceAccessesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PrivateServiceAccessCollection instances
	PrivateServiceAccessCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPrivateServiceAccessesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPrivateServiceAccessesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPrivateServiceAccessesSortByEnum Enum with underlying type: string
type ListPrivateServiceAccessesSortByEnum string

// Set of constants representing the allowable values for ListPrivateServiceAccessesSortByEnum
const (
	ListPrivateServiceAccessesSortByTimecreated ListPrivateServiceAccessesSortByEnum = "timeCreated"
	ListPrivateServiceAccessesSortByDisplayname ListPrivateServiceAccessesSortByEnum = "displayName"
)

var mappingListPrivateServiceAccessesSortByEnum = map[string]ListPrivateServiceAccessesSortByEnum{
	"timeCreated": ListPrivateServiceAccessesSortByTimecreated,
	"displayName": ListPrivateServiceAccessesSortByDisplayname,
}

var mappingListPrivateServiceAccessesSortByEnumLowerCase = map[string]ListPrivateServiceAccessesSortByEnum{
	"timecreated": ListPrivateServiceAccessesSortByTimecreated,
	"displayname": ListPrivateServiceAccessesSortByDisplayname,
}

// GetListPrivateServiceAccessesSortByEnumValues Enumerates the set of values for ListPrivateServiceAccessesSortByEnum
func GetListPrivateServiceAccessesSortByEnumValues() []ListPrivateServiceAccessesSortByEnum {
	values := make([]ListPrivateServiceAccessesSortByEnum, 0)
	for _, v := range mappingListPrivateServiceAccessesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateServiceAccessesSortByEnumStringValues Enumerates the set of values in String for ListPrivateServiceAccessesSortByEnum
func GetListPrivateServiceAccessesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListPrivateServiceAccessesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateServiceAccessesSortByEnum(val string) (ListPrivateServiceAccessesSortByEnum, bool) {
	enum, ok := mappingListPrivateServiceAccessesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPrivateServiceAccessesSortOrderEnum Enum with underlying type: string
type ListPrivateServiceAccessesSortOrderEnum string

// Set of constants representing the allowable values for ListPrivateServiceAccessesSortOrderEnum
const (
	ListPrivateServiceAccessesSortOrderAsc  ListPrivateServiceAccessesSortOrderEnum = "ASC"
	ListPrivateServiceAccessesSortOrderDesc ListPrivateServiceAccessesSortOrderEnum = "DESC"
)

var mappingListPrivateServiceAccessesSortOrderEnum = map[string]ListPrivateServiceAccessesSortOrderEnum{
	"ASC":  ListPrivateServiceAccessesSortOrderAsc,
	"DESC": ListPrivateServiceAccessesSortOrderDesc,
}

var mappingListPrivateServiceAccessesSortOrderEnumLowerCase = map[string]ListPrivateServiceAccessesSortOrderEnum{
	"asc":  ListPrivateServiceAccessesSortOrderAsc,
	"desc": ListPrivateServiceAccessesSortOrderDesc,
}

// GetListPrivateServiceAccessesSortOrderEnumValues Enumerates the set of values for ListPrivateServiceAccessesSortOrderEnum
func GetListPrivateServiceAccessesSortOrderEnumValues() []ListPrivateServiceAccessesSortOrderEnum {
	values := make([]ListPrivateServiceAccessesSortOrderEnum, 0)
	for _, v := range mappingListPrivateServiceAccessesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateServiceAccessesSortOrderEnumStringValues Enumerates the set of values in String for ListPrivateServiceAccessesSortOrderEnum
func GetListPrivateServiceAccessesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPrivateServiceAccessesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateServiceAccessesSortOrderEnum(val string) (ListPrivateServiceAccessesSortOrderEnum, bool) {
	enum, ok := mappingListPrivateServiceAccessesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
