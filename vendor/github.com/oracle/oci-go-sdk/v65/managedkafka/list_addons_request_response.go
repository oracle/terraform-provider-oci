// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managedkafka

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAddonsRequest wrapper for the ListAddons operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managedkafka/ListAddons.go.html to see an example of how to use ListAddonsRequest.
type ListAddonsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KafkaCluster.
	KafkaClusterId *string `mandatory:"true" contributesTo:"path" name:"kafkaClusterId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState KafkaClusterAddonLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The name to filter on.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAddonsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending.
	SortBy ListAddonsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAddonsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAddonsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAddonsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAddonsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAddonsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKafkaClusterAddonLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetKafkaClusterAddonLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAddonsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAddonsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAddonsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAddonsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAddonsResponse wrapper for the ListAddons operation
type ListAddonsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AddonCollection instances
	AddonCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAddonsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAddonsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAddonsSortOrderEnum Enum with underlying type: string
type ListAddonsSortOrderEnum string

// Set of constants representing the allowable values for ListAddonsSortOrderEnum
const (
	ListAddonsSortOrderAsc  ListAddonsSortOrderEnum = "ASC"
	ListAddonsSortOrderDesc ListAddonsSortOrderEnum = "DESC"
)

var mappingListAddonsSortOrderEnum = map[string]ListAddonsSortOrderEnum{
	"ASC":  ListAddonsSortOrderAsc,
	"DESC": ListAddonsSortOrderDesc,
}

var mappingListAddonsSortOrderEnumLowerCase = map[string]ListAddonsSortOrderEnum{
	"asc":  ListAddonsSortOrderAsc,
	"desc": ListAddonsSortOrderDesc,
}

// GetListAddonsSortOrderEnumValues Enumerates the set of values for ListAddonsSortOrderEnum
func GetListAddonsSortOrderEnumValues() []ListAddonsSortOrderEnum {
	values := make([]ListAddonsSortOrderEnum, 0)
	for _, v := range mappingListAddonsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddonsSortOrderEnumStringValues Enumerates the set of values in String for ListAddonsSortOrderEnum
func GetListAddonsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAddonsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddonsSortOrderEnum(val string) (ListAddonsSortOrderEnum, bool) {
	enum, ok := mappingListAddonsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAddonsSortByEnum Enum with underlying type: string
type ListAddonsSortByEnum string

// Set of constants representing the allowable values for ListAddonsSortByEnum
const (
	ListAddonsSortByTimecreated ListAddonsSortByEnum = "timeCreated"
)

var mappingListAddonsSortByEnum = map[string]ListAddonsSortByEnum{
	"timeCreated": ListAddonsSortByTimecreated,
}

var mappingListAddonsSortByEnumLowerCase = map[string]ListAddonsSortByEnum{
	"timecreated": ListAddonsSortByTimecreated,
}

// GetListAddonsSortByEnumValues Enumerates the set of values for ListAddonsSortByEnum
func GetListAddonsSortByEnumValues() []ListAddonsSortByEnum {
	values := make([]ListAddonsSortByEnum, 0)
	for _, v := range mappingListAddonsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddonsSortByEnumStringValues Enumerates the set of values in String for ListAddonsSortByEnum
func GetListAddonsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListAddonsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddonsSortByEnum(val string) (ListAddonsSortByEnum, bool) {
	enum, ok := mappingListAddonsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
