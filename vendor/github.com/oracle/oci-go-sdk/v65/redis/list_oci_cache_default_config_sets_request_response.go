// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOciCacheDefaultConfigSetsRequest wrapper for the ListOciCacheDefaultConfigSets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheDefaultConfigSets.go.html to see an example of how to use ListOciCacheDefaultConfigSetsRequest.
type ListOciCacheDefaultConfigSetsRequest struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique OCI Cache Default Config Set identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return the OCI Cache Default Config Set resources, whose lifecycle state matches with the given lifecycle state.
	LifecycleState OciCacheDefaultConfigSetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return the OCI Cache Config Set resources, whose software version matches with the given software version.
	SoftwareVersion OciCacheConfigSetSoftwareVersionEnum `mandatory:"false" contributesTo:"query" name:"softwareVersion" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOciCacheDefaultConfigSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOciCacheDefaultConfigSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOciCacheDefaultConfigSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOciCacheDefaultConfigSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOciCacheDefaultConfigSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOciCacheDefaultConfigSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOciCacheDefaultConfigSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheDefaultConfigSetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOciCacheDefaultConfigSetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciCacheConfigSetSoftwareVersionEnum(string(request.SoftwareVersion)); !ok && request.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", request.SoftwareVersion, strings.Join(GetOciCacheConfigSetSoftwareVersionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheDefaultConfigSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOciCacheDefaultConfigSetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheDefaultConfigSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOciCacheDefaultConfigSetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOciCacheDefaultConfigSetsResponse wrapper for the ListOciCacheDefaultConfigSets operation
type ListOciCacheDefaultConfigSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OciCacheDefaultConfigSetCollection instances
	OciCacheDefaultConfigSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOciCacheDefaultConfigSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOciCacheDefaultConfigSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOciCacheDefaultConfigSetsSortByEnum Enum with underlying type: string
type ListOciCacheDefaultConfigSetsSortByEnum string

// Set of constants representing the allowable values for ListOciCacheDefaultConfigSetsSortByEnum
const (
	ListOciCacheDefaultConfigSetsSortByTimecreated ListOciCacheDefaultConfigSetsSortByEnum = "timeCreated"
	ListOciCacheDefaultConfigSetsSortByDisplayname ListOciCacheDefaultConfigSetsSortByEnum = "displayName"
)

var mappingListOciCacheDefaultConfigSetsSortByEnum = map[string]ListOciCacheDefaultConfigSetsSortByEnum{
	"timeCreated": ListOciCacheDefaultConfigSetsSortByTimecreated,
	"displayName": ListOciCacheDefaultConfigSetsSortByDisplayname,
}

var mappingListOciCacheDefaultConfigSetsSortByEnumLowerCase = map[string]ListOciCacheDefaultConfigSetsSortByEnum{
	"timecreated": ListOciCacheDefaultConfigSetsSortByTimecreated,
	"displayname": ListOciCacheDefaultConfigSetsSortByDisplayname,
}

// GetListOciCacheDefaultConfigSetsSortByEnumValues Enumerates the set of values for ListOciCacheDefaultConfigSetsSortByEnum
func GetListOciCacheDefaultConfigSetsSortByEnumValues() []ListOciCacheDefaultConfigSetsSortByEnum {
	values := make([]ListOciCacheDefaultConfigSetsSortByEnum, 0)
	for _, v := range mappingListOciCacheDefaultConfigSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheDefaultConfigSetsSortByEnumStringValues Enumerates the set of values in String for ListOciCacheDefaultConfigSetsSortByEnum
func GetListOciCacheDefaultConfigSetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOciCacheDefaultConfigSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheDefaultConfigSetsSortByEnum(val string) (ListOciCacheDefaultConfigSetsSortByEnum, bool) {
	enum, ok := mappingListOciCacheDefaultConfigSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOciCacheDefaultConfigSetsSortOrderEnum Enum with underlying type: string
type ListOciCacheDefaultConfigSetsSortOrderEnum string

// Set of constants representing the allowable values for ListOciCacheDefaultConfigSetsSortOrderEnum
const (
	ListOciCacheDefaultConfigSetsSortOrderAsc  ListOciCacheDefaultConfigSetsSortOrderEnum = "ASC"
	ListOciCacheDefaultConfigSetsSortOrderDesc ListOciCacheDefaultConfigSetsSortOrderEnum = "DESC"
)

var mappingListOciCacheDefaultConfigSetsSortOrderEnum = map[string]ListOciCacheDefaultConfigSetsSortOrderEnum{
	"ASC":  ListOciCacheDefaultConfigSetsSortOrderAsc,
	"DESC": ListOciCacheDefaultConfigSetsSortOrderDesc,
}

var mappingListOciCacheDefaultConfigSetsSortOrderEnumLowerCase = map[string]ListOciCacheDefaultConfigSetsSortOrderEnum{
	"asc":  ListOciCacheDefaultConfigSetsSortOrderAsc,
	"desc": ListOciCacheDefaultConfigSetsSortOrderDesc,
}

// GetListOciCacheDefaultConfigSetsSortOrderEnumValues Enumerates the set of values for ListOciCacheDefaultConfigSetsSortOrderEnum
func GetListOciCacheDefaultConfigSetsSortOrderEnumValues() []ListOciCacheDefaultConfigSetsSortOrderEnum {
	values := make([]ListOciCacheDefaultConfigSetsSortOrderEnum, 0)
	for _, v := range mappingListOciCacheDefaultConfigSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheDefaultConfigSetsSortOrderEnumStringValues Enumerates the set of values in String for ListOciCacheDefaultConfigSetsSortOrderEnum
func GetListOciCacheDefaultConfigSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOciCacheDefaultConfigSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheDefaultConfigSetsSortOrderEnum(val string) (ListOciCacheDefaultConfigSetsSortOrderEnum, bool) {
	enum, ok := mappingListOciCacheDefaultConfigSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
