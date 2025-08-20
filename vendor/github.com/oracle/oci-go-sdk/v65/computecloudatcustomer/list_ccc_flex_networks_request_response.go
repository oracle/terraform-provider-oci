// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCccFlexNetworksRequest wrapper for the ListCccFlexNetworks operation
type ListCccFlexNetworksRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to
	// list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter used to return only resources that match the given infrastructureId.
	CccInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"cccInfrastructureId"`

	// An OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a
	// Compute Cloud@Customer FlexNetwork.
	FlexNetworkId *string `mandatory:"false" contributesTo:"query" name:"flexNetworkId"`

	// A filter used to return only resources that match the given lifecycleState.
	LifecycleState CccFlexNetworkLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose display name contains the substring.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCccFlexNetworksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListCccFlexNetworksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCccFlexNetworksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCccFlexNetworksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCccFlexNetworksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCccFlexNetworksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCccFlexNetworksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccFlexNetworkLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCccFlexNetworkLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccFlexNetworksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCccFlexNetworksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccFlexNetworksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCccFlexNetworksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCccFlexNetworksResponse wrapper for the ListCccFlexNetworks operation
type ListCccFlexNetworksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CccFlexNetworkCollection instances
	CccFlexNetworkCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCccFlexNetworksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCccFlexNetworksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCccFlexNetworksSortOrderEnum Enum with underlying type: string
type ListCccFlexNetworksSortOrderEnum string

// Set of constants representing the allowable values for ListCccFlexNetworksSortOrderEnum
const (
	ListCccFlexNetworksSortOrderAsc  ListCccFlexNetworksSortOrderEnum = "ASC"
	ListCccFlexNetworksSortOrderDesc ListCccFlexNetworksSortOrderEnum = "DESC"
)

var mappingListCccFlexNetworksSortOrderEnum = map[string]ListCccFlexNetworksSortOrderEnum{
	"ASC":  ListCccFlexNetworksSortOrderAsc,
	"DESC": ListCccFlexNetworksSortOrderDesc,
}

var mappingListCccFlexNetworksSortOrderEnumLowerCase = map[string]ListCccFlexNetworksSortOrderEnum{
	"asc":  ListCccFlexNetworksSortOrderAsc,
	"desc": ListCccFlexNetworksSortOrderDesc,
}

// GetListCccFlexNetworksSortOrderEnumValues Enumerates the set of values for ListCccFlexNetworksSortOrderEnum
func GetListCccFlexNetworksSortOrderEnumValues() []ListCccFlexNetworksSortOrderEnum {
	values := make([]ListCccFlexNetworksSortOrderEnum, 0)
	for _, v := range mappingListCccFlexNetworksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccFlexNetworksSortOrderEnumStringValues Enumerates the set of values in String for ListCccFlexNetworksSortOrderEnum
func GetListCccFlexNetworksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCccFlexNetworksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccFlexNetworksSortOrderEnum(val string) (ListCccFlexNetworksSortOrderEnum, bool) {
	enum, ok := mappingListCccFlexNetworksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccFlexNetworksSortByEnum Enum with underlying type: string
type ListCccFlexNetworksSortByEnum string

// Set of constants representing the allowable values for ListCccFlexNetworksSortByEnum
const (
	ListCccFlexNetworksSortByTimecreated ListCccFlexNetworksSortByEnum = "timeCreated"
	ListCccFlexNetworksSortByDisplayname ListCccFlexNetworksSortByEnum = "displayName"
)

var mappingListCccFlexNetworksSortByEnum = map[string]ListCccFlexNetworksSortByEnum{
	"timeCreated": ListCccFlexNetworksSortByTimecreated,
	"displayName": ListCccFlexNetworksSortByDisplayname,
}

var mappingListCccFlexNetworksSortByEnumLowerCase = map[string]ListCccFlexNetworksSortByEnum{
	"timecreated": ListCccFlexNetworksSortByTimecreated,
	"displayname": ListCccFlexNetworksSortByDisplayname,
}

// GetListCccFlexNetworksSortByEnumValues Enumerates the set of values for ListCccFlexNetworksSortByEnum
func GetListCccFlexNetworksSortByEnumValues() []ListCccFlexNetworksSortByEnum {
	values := make([]ListCccFlexNetworksSortByEnum, 0)
	for _, v := range mappingListCccFlexNetworksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccFlexNetworksSortByEnumStringValues Enumerates the set of values in String for ListCccFlexNetworksSortByEnum
func GetListCccFlexNetworksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCccFlexNetworksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccFlexNetworksSortByEnum(val string) (ListCccFlexNetworksSortByEnum, bool) {
	enum, ok := mappingListCccFlexNetworksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
