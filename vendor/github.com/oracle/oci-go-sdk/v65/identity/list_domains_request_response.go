// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDomainsRequest wrapper for the ListDomains operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListDomains.go.html to see an example of how to use ListDomainsRequest.
type ListDomainsRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The mutable display name of the identity domain.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The region-agnostic identity domain URL.
	Url *string `mandatory:"false" contributesTo:"query" name:"url"`

	// The region-specific identity domain URL.
	HomeRegionUrl *string `mandatory:"false" contributesTo:"query" name:"homeRegionUrl"`

	// The identity domain type.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// The license type of the identity domain.
	LicenseType *string `mandatory:"false" contributesTo:"query" name:"licenseType"`

	// Indicates whether or not the identity domain is visible at the sign-in screen.
	IsHiddenOnLogin *bool `mandatory:"false" contributesTo:"query" name:"isHiddenOnLogin"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListDomainsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListDomainsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState DomainLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDomainsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDomainsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDomainsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDomainsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDomainsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDomainsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDomainsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDomainsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDomainsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDomainLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDomainLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDomainsResponse wrapper for the ListDomains operation
type ListDomainsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DomainSummary instances
	Items []DomainSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDomainsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDomainsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDomainsSortByEnum Enum with underlying type: string
type ListDomainsSortByEnum string

// Set of constants representing the allowable values for ListDomainsSortByEnum
const (
	ListDomainsSortByTimecreated ListDomainsSortByEnum = "TIMECREATED"
	ListDomainsSortByName        ListDomainsSortByEnum = "NAME"
)

var mappingListDomainsSortByEnum = map[string]ListDomainsSortByEnum{
	"TIMECREATED": ListDomainsSortByTimecreated,
	"NAME":        ListDomainsSortByName,
}

var mappingListDomainsSortByEnumLowerCase = map[string]ListDomainsSortByEnum{
	"timecreated": ListDomainsSortByTimecreated,
	"name":        ListDomainsSortByName,
}

// GetListDomainsSortByEnumValues Enumerates the set of values for ListDomainsSortByEnum
func GetListDomainsSortByEnumValues() []ListDomainsSortByEnum {
	values := make([]ListDomainsSortByEnum, 0)
	for _, v := range mappingListDomainsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDomainsSortByEnumStringValues Enumerates the set of values in String for ListDomainsSortByEnum
func GetListDomainsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListDomainsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDomainsSortByEnum(val string) (ListDomainsSortByEnum, bool) {
	enum, ok := mappingListDomainsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDomainsSortOrderEnum Enum with underlying type: string
type ListDomainsSortOrderEnum string

// Set of constants representing the allowable values for ListDomainsSortOrderEnum
const (
	ListDomainsSortOrderAsc  ListDomainsSortOrderEnum = "ASC"
	ListDomainsSortOrderDesc ListDomainsSortOrderEnum = "DESC"
)

var mappingListDomainsSortOrderEnum = map[string]ListDomainsSortOrderEnum{
	"ASC":  ListDomainsSortOrderAsc,
	"DESC": ListDomainsSortOrderDesc,
}

var mappingListDomainsSortOrderEnumLowerCase = map[string]ListDomainsSortOrderEnum{
	"asc":  ListDomainsSortOrderAsc,
	"desc": ListDomainsSortOrderDesc,
}

// GetListDomainsSortOrderEnumValues Enumerates the set of values for ListDomainsSortOrderEnum
func GetListDomainsSortOrderEnumValues() []ListDomainsSortOrderEnum {
	values := make([]ListDomainsSortOrderEnum, 0)
	for _, v := range mappingListDomainsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDomainsSortOrderEnumStringValues Enumerates the set of values in String for ListDomainsSortOrderEnum
func GetListDomainsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDomainsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDomainsSortOrderEnum(val string) (ListDomainsSortOrderEnum, bool) {
	enum, ok := mappingListDomainsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
