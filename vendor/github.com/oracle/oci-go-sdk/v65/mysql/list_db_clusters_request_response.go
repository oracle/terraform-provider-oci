// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbClustersRequest wrapper for the ListDbClusters operation
type ListDbClustersRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The shared-storage DB cluster OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbClusterId *string `mandatory:"false" contributesTo:"query" name:"dbClusterId"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The shared-storage DB cluster lifecycle state.
	LifecycleState DbClusterLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The requested Configuration instance.
	ConfigurationId *string `mandatory:"false" contributesTo:"query" name:"configurationId"`

	// The field to use for sorting.
	SortBy ListDbClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListDbClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.oracle.com/iaasAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbClusterLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbClusterLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbClustersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbClustersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbClustersResponse wrapper for the ListDbClusters operation
type ListDbClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DbClusterCollection instances
	DbClusterCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbClustersSortByEnum Enum with underlying type: string
type ListDbClustersSortByEnum string

// Set of constants representing the allowable values for ListDbClustersSortByEnum
const (
	ListDbClustersSortByDisplayname ListDbClustersSortByEnum = "displayName"
	ListDbClustersSortByTimecreated ListDbClustersSortByEnum = "timeCreated"
)

var mappingListDbClustersSortByEnum = map[string]ListDbClustersSortByEnum{
	"displayName": ListDbClustersSortByDisplayname,
	"timeCreated": ListDbClustersSortByTimecreated,
}

var mappingListDbClustersSortByEnumLowerCase = map[string]ListDbClustersSortByEnum{
	"displayname": ListDbClustersSortByDisplayname,
	"timecreated": ListDbClustersSortByTimecreated,
}

// GetListDbClustersSortByEnumValues Enumerates the set of values for ListDbClustersSortByEnum
func GetListDbClustersSortByEnumValues() []ListDbClustersSortByEnum {
	values := make([]ListDbClustersSortByEnum, 0)
	for _, v := range mappingListDbClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbClustersSortByEnumStringValues Enumerates the set of values in String for ListDbClustersSortByEnum
func GetListDbClustersSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListDbClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbClustersSortByEnum(val string) (ListDbClustersSortByEnum, bool) {
	enum, ok := mappingListDbClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbClustersSortOrderEnum Enum with underlying type: string
type ListDbClustersSortOrderEnum string

// Set of constants representing the allowable values for ListDbClustersSortOrderEnum
const (
	ListDbClustersSortOrderAsc  ListDbClustersSortOrderEnum = "ASC"
	ListDbClustersSortOrderDesc ListDbClustersSortOrderEnum = "DESC"
)

var mappingListDbClustersSortOrderEnum = map[string]ListDbClustersSortOrderEnum{
	"ASC":  ListDbClustersSortOrderAsc,
	"DESC": ListDbClustersSortOrderDesc,
}

var mappingListDbClustersSortOrderEnumLowerCase = map[string]ListDbClustersSortOrderEnum{
	"asc":  ListDbClustersSortOrderAsc,
	"desc": ListDbClustersSortOrderDesc,
}

// GetListDbClustersSortOrderEnumValues Enumerates the set of values for ListDbClustersSortOrderEnum
func GetListDbClustersSortOrderEnumValues() []ListDbClustersSortOrderEnum {
	values := make([]ListDbClustersSortOrderEnum, 0)
	for _, v := range mappingListDbClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbClustersSortOrderEnumStringValues Enumerates the set of values in String for ListDbClustersSortOrderEnum
func GetListDbClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbClustersSortOrderEnum(val string) (ListDbClustersSortOrderEnum, bool) {
	enum, ok := mappingListDbClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
