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

// ListOciCacheConfigSetsRequest wrapper for the ListOciCacheConfigSets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheConfigSets.go.html to see an example of how to use ListOciCacheConfigSetsRequest.
type ListOciCacheConfigSetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return the OCI Cache Config Set resources, whose lifecycle state matches with the given lifecycle state.
	LifecycleState OciCacheConfigSetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return the OCI Cache Config Set resources, whose software version matches with the given software version.
	SoftwareVersion OciCacheConfigSetSoftwareVersionEnum `mandatory:"false" contributesTo:"query" name:"softwareVersion" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique OCI Cache Config Set identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOciCacheConfigSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListOciCacheConfigSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOciCacheConfigSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOciCacheConfigSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOciCacheConfigSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOciCacheConfigSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOciCacheConfigSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheConfigSetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOciCacheConfigSetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciCacheConfigSetSoftwareVersionEnum(string(request.SoftwareVersion)); !ok && request.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", request.SoftwareVersion, strings.Join(GetOciCacheConfigSetSoftwareVersionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheConfigSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOciCacheConfigSetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOciCacheConfigSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOciCacheConfigSetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOciCacheConfigSetsResponse wrapper for the ListOciCacheConfigSets operation
type ListOciCacheConfigSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OciCacheConfigSetCollection instances
	OciCacheConfigSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOciCacheConfigSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOciCacheConfigSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOciCacheConfigSetsSortOrderEnum Enum with underlying type: string
type ListOciCacheConfigSetsSortOrderEnum string

// Set of constants representing the allowable values for ListOciCacheConfigSetsSortOrderEnum
const (
	ListOciCacheConfigSetsSortOrderAsc  ListOciCacheConfigSetsSortOrderEnum = "ASC"
	ListOciCacheConfigSetsSortOrderDesc ListOciCacheConfigSetsSortOrderEnum = "DESC"
)

var mappingListOciCacheConfigSetsSortOrderEnum = map[string]ListOciCacheConfigSetsSortOrderEnum{
	"ASC":  ListOciCacheConfigSetsSortOrderAsc,
	"DESC": ListOciCacheConfigSetsSortOrderDesc,
}

var mappingListOciCacheConfigSetsSortOrderEnumLowerCase = map[string]ListOciCacheConfigSetsSortOrderEnum{
	"asc":  ListOciCacheConfigSetsSortOrderAsc,
	"desc": ListOciCacheConfigSetsSortOrderDesc,
}

// GetListOciCacheConfigSetsSortOrderEnumValues Enumerates the set of values for ListOciCacheConfigSetsSortOrderEnum
func GetListOciCacheConfigSetsSortOrderEnumValues() []ListOciCacheConfigSetsSortOrderEnum {
	values := make([]ListOciCacheConfigSetsSortOrderEnum, 0)
	for _, v := range mappingListOciCacheConfigSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheConfigSetsSortOrderEnumStringValues Enumerates the set of values in String for ListOciCacheConfigSetsSortOrderEnum
func GetListOciCacheConfigSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOciCacheConfigSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheConfigSetsSortOrderEnum(val string) (ListOciCacheConfigSetsSortOrderEnum, bool) {
	enum, ok := mappingListOciCacheConfigSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOciCacheConfigSetsSortByEnum Enum with underlying type: string
type ListOciCacheConfigSetsSortByEnum string

// Set of constants representing the allowable values for ListOciCacheConfigSetsSortByEnum
const (
	ListOciCacheConfigSetsSortByTimecreated ListOciCacheConfigSetsSortByEnum = "timeCreated"
	ListOciCacheConfigSetsSortByDisplayname ListOciCacheConfigSetsSortByEnum = "displayName"
)

var mappingListOciCacheConfigSetsSortByEnum = map[string]ListOciCacheConfigSetsSortByEnum{
	"timeCreated": ListOciCacheConfigSetsSortByTimecreated,
	"displayName": ListOciCacheConfigSetsSortByDisplayname,
}

var mappingListOciCacheConfigSetsSortByEnumLowerCase = map[string]ListOciCacheConfigSetsSortByEnum{
	"timecreated": ListOciCacheConfigSetsSortByTimecreated,
	"displayname": ListOciCacheConfigSetsSortByDisplayname,
}

// GetListOciCacheConfigSetsSortByEnumValues Enumerates the set of values for ListOciCacheConfigSetsSortByEnum
func GetListOciCacheConfigSetsSortByEnumValues() []ListOciCacheConfigSetsSortByEnum {
	values := make([]ListOciCacheConfigSetsSortByEnum, 0)
	for _, v := range mappingListOciCacheConfigSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOciCacheConfigSetsSortByEnumStringValues Enumerates the set of values in String for ListOciCacheConfigSetsSortByEnum
func GetListOciCacheConfigSetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOciCacheConfigSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOciCacheConfigSetsSortByEnum(val string) (ListOciCacheConfigSetsSortByEnum, bool) {
	enum, ok := mappingListOciCacheConfigSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
