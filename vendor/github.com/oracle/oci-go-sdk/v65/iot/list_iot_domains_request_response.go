// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIotDomainsRequest wrapper for the ListIotDomains operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListIotDomains.go.html to see an example of how to use ListIotDomainsRequest.
type ListIotDomainsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter resources by OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Filter resources that match the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain group.
	IotDomainGroupId *string `mandatory:"false" contributesTo:"query" name:"iotDomainGroupId"`

	// Filter resources whose display name matches the specified value.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter resources whose lifecycleState matches the specified value.
	LifecycleState IotDomainLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination: The value of the opc-next-page response header from the previous "List" call.
	// For important details on how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either ASC (ascending) or DESC (descending).
	SortOrder ListIotDomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListIotDomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIotDomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIotDomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIotDomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIotDomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIotDomainsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIotDomainLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetIotDomainLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIotDomainsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIotDomainsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIotDomainsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIotDomainsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIotDomainsResponse wrapper for the ListIotDomains operation
type ListIotDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IotDomainCollection instances
	IotDomainCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListIotDomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIotDomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIotDomainsSortOrderEnum Enum with underlying type: string
type ListIotDomainsSortOrderEnum string

// Set of constants representing the allowable values for ListIotDomainsSortOrderEnum
const (
	ListIotDomainsSortOrderAsc  ListIotDomainsSortOrderEnum = "ASC"
	ListIotDomainsSortOrderDesc ListIotDomainsSortOrderEnum = "DESC"
)

var mappingListIotDomainsSortOrderEnum = map[string]ListIotDomainsSortOrderEnum{
	"ASC":  ListIotDomainsSortOrderAsc,
	"DESC": ListIotDomainsSortOrderDesc,
}

var mappingListIotDomainsSortOrderEnumLowerCase = map[string]ListIotDomainsSortOrderEnum{
	"asc":  ListIotDomainsSortOrderAsc,
	"desc": ListIotDomainsSortOrderDesc,
}

// GetListIotDomainsSortOrderEnumValues Enumerates the set of values for ListIotDomainsSortOrderEnum
func GetListIotDomainsSortOrderEnumValues() []ListIotDomainsSortOrderEnum {
	values := make([]ListIotDomainsSortOrderEnum, 0)
	for _, v := range mappingListIotDomainsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIotDomainsSortOrderEnumStringValues Enumerates the set of values in String for ListIotDomainsSortOrderEnum
func GetListIotDomainsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIotDomainsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIotDomainsSortOrderEnum(val string) (ListIotDomainsSortOrderEnum, bool) {
	enum, ok := mappingListIotDomainsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIotDomainsSortByEnum Enum with underlying type: string
type ListIotDomainsSortByEnum string

// Set of constants representing the allowable values for ListIotDomainsSortByEnum
const (
	ListIotDomainsSortByTimecreated ListIotDomainsSortByEnum = "timeCreated"
	ListIotDomainsSortByDisplayname ListIotDomainsSortByEnum = "displayName"
)

var mappingListIotDomainsSortByEnum = map[string]ListIotDomainsSortByEnum{
	"timeCreated": ListIotDomainsSortByTimecreated,
	"displayName": ListIotDomainsSortByDisplayname,
}

var mappingListIotDomainsSortByEnumLowerCase = map[string]ListIotDomainsSortByEnum{
	"timecreated": ListIotDomainsSortByTimecreated,
	"displayname": ListIotDomainsSortByDisplayname,
}

// GetListIotDomainsSortByEnumValues Enumerates the set of values for ListIotDomainsSortByEnum
func GetListIotDomainsSortByEnumValues() []ListIotDomainsSortByEnum {
	values := make([]ListIotDomainsSortByEnum, 0)
	for _, v := range mappingListIotDomainsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIotDomainsSortByEnumStringValues Enumerates the set of values in String for ListIotDomainsSortByEnum
func GetListIotDomainsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListIotDomainsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIotDomainsSortByEnum(val string) (ListIotDomainsSortByEnum, bool) {
	enum, ok := mappingListIotDomainsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
