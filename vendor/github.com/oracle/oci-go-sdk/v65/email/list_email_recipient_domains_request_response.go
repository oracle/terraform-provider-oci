// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailRecipientDomainsRequest wrapper for the ListEmailRecipientDomains operation
type ListEmailRecipientDomainsRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to only return resources that match the email delivery config exactly.
	EmailDeliveryConfigId *string `mandatory:"false" contributesTo:"query" name:"emailDeliveryConfigId"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailRecipientDomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState EmailRecipientDomainLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies the attribute with which to sort the recipient domains.
	// Default: `timeCreated`
	// * **timeCreated:** Sorts by timeCreated.
	// * **name:** Sorts by name.
	SortBy ListEmailRecipientDomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailRecipientDomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailRecipientDomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailRecipientDomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailRecipientDomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailRecipientDomainsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailRecipientDomainsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailRecipientDomainsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailRecipientDomainLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailRecipientDomainLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailRecipientDomainsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailRecipientDomainsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailRecipientDomainsResponse wrapper for the ListEmailRecipientDomains operation
type ListEmailRecipientDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailRecipientDomainCollection instances
	EmailRecipientDomainCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListEmailRecipientDomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailRecipientDomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailRecipientDomainsSortOrderEnum Enum with underlying type: string
type ListEmailRecipientDomainsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailRecipientDomainsSortOrderEnum
const (
	ListEmailRecipientDomainsSortOrderAsc  ListEmailRecipientDomainsSortOrderEnum = "ASC"
	ListEmailRecipientDomainsSortOrderDesc ListEmailRecipientDomainsSortOrderEnum = "DESC"
)

var mappingListEmailRecipientDomainsSortOrderEnum = map[string]ListEmailRecipientDomainsSortOrderEnum{
	"ASC":  ListEmailRecipientDomainsSortOrderAsc,
	"DESC": ListEmailRecipientDomainsSortOrderDesc,
}

var mappingListEmailRecipientDomainsSortOrderEnumLowerCase = map[string]ListEmailRecipientDomainsSortOrderEnum{
	"asc":  ListEmailRecipientDomainsSortOrderAsc,
	"desc": ListEmailRecipientDomainsSortOrderDesc,
}

// GetListEmailRecipientDomainsSortOrderEnumValues Enumerates the set of values for ListEmailRecipientDomainsSortOrderEnum
func GetListEmailRecipientDomainsSortOrderEnumValues() []ListEmailRecipientDomainsSortOrderEnum {
	values := make([]ListEmailRecipientDomainsSortOrderEnum, 0)
	for _, v := range mappingListEmailRecipientDomainsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailRecipientDomainsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailRecipientDomainsSortOrderEnum
func GetListEmailRecipientDomainsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailRecipientDomainsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailRecipientDomainsSortOrderEnum(val string) (ListEmailRecipientDomainsSortOrderEnum, bool) {
	enum, ok := mappingListEmailRecipientDomainsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailRecipientDomainsSortByEnum Enum with underlying type: string
type ListEmailRecipientDomainsSortByEnum string

// Set of constants representing the allowable values for ListEmailRecipientDomainsSortByEnum
const (
	ListEmailRecipientDomainsSortByTimecreated ListEmailRecipientDomainsSortByEnum = "timeCreated"
	ListEmailRecipientDomainsSortByName        ListEmailRecipientDomainsSortByEnum = "name"
)

var mappingListEmailRecipientDomainsSortByEnum = map[string]ListEmailRecipientDomainsSortByEnum{
	"timeCreated": ListEmailRecipientDomainsSortByTimecreated,
	"name":        ListEmailRecipientDomainsSortByName,
}

var mappingListEmailRecipientDomainsSortByEnumLowerCase = map[string]ListEmailRecipientDomainsSortByEnum{
	"timecreated": ListEmailRecipientDomainsSortByTimecreated,
	"name":        ListEmailRecipientDomainsSortByName,
}

// GetListEmailRecipientDomainsSortByEnumValues Enumerates the set of values for ListEmailRecipientDomainsSortByEnum
func GetListEmailRecipientDomainsSortByEnumValues() []ListEmailRecipientDomainsSortByEnum {
	values := make([]ListEmailRecipientDomainsSortByEnum, 0)
	for _, v := range mappingListEmailRecipientDomainsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailRecipientDomainsSortByEnumStringValues Enumerates the set of values in String for ListEmailRecipientDomainsSortByEnum
func GetListEmailRecipientDomainsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListEmailRecipientDomainsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailRecipientDomainsSortByEnum(val string) (ListEmailRecipientDomainsSortByEnum, bool) {
	enum, ok := mappingListEmailRecipientDomainsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
