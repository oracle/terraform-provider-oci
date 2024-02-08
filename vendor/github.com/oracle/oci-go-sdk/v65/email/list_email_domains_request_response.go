// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailDomainsRequest wrapper for the ListEmailDomains operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/email/ListEmailDomains.go.html to see an example of how to use ListEmailDomainsRequest.
type ListEmailDomainsRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailDomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState EmailDomainLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies the attribute with which to sort the email domains.
	// Default: `TIMECREATED`
	// * **TIMECREATED:** Sorts by timeCreated.
	// * **NAME:** Sorts by name.
	// * **ID:** Sorts by id.
	SortBy ListEmailDomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailDomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailDomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailDomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailDomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailDomainsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailDomainsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailDomainsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailDomainLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailDomainLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailDomainsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailDomainsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailDomainsResponse wrapper for the ListEmailDomains operation
type ListEmailDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailDomainCollection instances
	EmailDomainCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListEmailDomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailDomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailDomainsSortOrderEnum Enum with underlying type: string
type ListEmailDomainsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailDomainsSortOrderEnum
const (
	ListEmailDomainsSortOrderAsc  ListEmailDomainsSortOrderEnum = "ASC"
	ListEmailDomainsSortOrderDesc ListEmailDomainsSortOrderEnum = "DESC"
)

var mappingListEmailDomainsSortOrderEnum = map[string]ListEmailDomainsSortOrderEnum{
	"ASC":  ListEmailDomainsSortOrderAsc,
	"DESC": ListEmailDomainsSortOrderDesc,
}

var mappingListEmailDomainsSortOrderEnumLowerCase = map[string]ListEmailDomainsSortOrderEnum{
	"asc":  ListEmailDomainsSortOrderAsc,
	"desc": ListEmailDomainsSortOrderDesc,
}

// GetListEmailDomainsSortOrderEnumValues Enumerates the set of values for ListEmailDomainsSortOrderEnum
func GetListEmailDomainsSortOrderEnumValues() []ListEmailDomainsSortOrderEnum {
	values := make([]ListEmailDomainsSortOrderEnum, 0)
	for _, v := range mappingListEmailDomainsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailDomainsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailDomainsSortOrderEnum
func GetListEmailDomainsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailDomainsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailDomainsSortOrderEnum(val string) (ListEmailDomainsSortOrderEnum, bool) {
	enum, ok := mappingListEmailDomainsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailDomainsSortByEnum Enum with underlying type: string
type ListEmailDomainsSortByEnum string

// Set of constants representing the allowable values for ListEmailDomainsSortByEnum
const (
	ListEmailDomainsSortByTimecreated ListEmailDomainsSortByEnum = "TIMECREATED"
	ListEmailDomainsSortById          ListEmailDomainsSortByEnum = "ID"
	ListEmailDomainsSortByName        ListEmailDomainsSortByEnum = "NAME"
)

var mappingListEmailDomainsSortByEnum = map[string]ListEmailDomainsSortByEnum{
	"TIMECREATED": ListEmailDomainsSortByTimecreated,
	"ID":          ListEmailDomainsSortById,
	"NAME":        ListEmailDomainsSortByName,
}

var mappingListEmailDomainsSortByEnumLowerCase = map[string]ListEmailDomainsSortByEnum{
	"timecreated": ListEmailDomainsSortByTimecreated,
	"id":          ListEmailDomainsSortById,
	"name":        ListEmailDomainsSortByName,
}

// GetListEmailDomainsSortByEnumValues Enumerates the set of values for ListEmailDomainsSortByEnum
func GetListEmailDomainsSortByEnumValues() []ListEmailDomainsSortByEnum {
	values := make([]ListEmailDomainsSortByEnum, 0)
	for _, v := range mappingListEmailDomainsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailDomainsSortByEnumStringValues Enumerates the set of values in String for ListEmailDomainsSortByEnum
func GetListEmailDomainsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"ID",
		"NAME",
	}
}

// GetMappingListEmailDomainsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailDomainsSortByEnum(val string) (ListEmailDomainsSortByEnum, bool) {
	enum, ok := mappingListEmailDomainsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
