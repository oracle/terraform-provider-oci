// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBdsMetastoreConfigurationsRequest wrapper for the ListBdsMetastoreConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListBdsMetastoreConfigurations.go.html to see an example of how to use ListBdsMetastoreConfigurationsRequest.
type ListBdsMetastoreConfigurationsRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The type of the metastore in the metastore configuration
	MetastoreType BdsMetastoreConfigurationMetastoreTypeEnum `mandatory:"false" contributesTo:"query" name:"metastoreType" omitEmpty:"true"`

	// The OCID of the Data Catalog metastore in the metastore configuration
	MetastoreId *string `mandatory:"false" contributesTo:"query" name:"metastoreId"`

	// The lifecycle state of the metastore in the metastore configuration
	LifecycleState BdsMetastoreConfigurationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The ID of the API key that is associated with the external metastore in the metastore configuration
	BdsApiKeyId *string `mandatory:"false" contributesTo:"query" name:"bdsApiKeyId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListBdsMetastoreConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBdsMetastoreConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBdsMetastoreConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBdsMetastoreConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBdsMetastoreConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBdsMetastoreConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBdsMetastoreConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsMetastoreConfigurationMetastoreTypeEnum(string(request.MetastoreType)); !ok && request.MetastoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetastoreType: %s. Supported values are: %s.", request.MetastoreType, strings.Join(GetBdsMetastoreConfigurationMetastoreTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBdsMetastoreConfigurationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBdsMetastoreConfigurationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsMetastoreConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBdsMetastoreConfigurationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsMetastoreConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBdsMetastoreConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBdsMetastoreConfigurationsResponse wrapper for the ListBdsMetastoreConfigurations operation
type ListBdsMetastoreConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BdsMetastoreConfigurationSummary instances
	Items []BdsMetastoreConfigurationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBdsMetastoreConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBdsMetastoreConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBdsMetastoreConfigurationsSortByEnum Enum with underlying type: string
type ListBdsMetastoreConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListBdsMetastoreConfigurationsSortByEnum
const (
	ListBdsMetastoreConfigurationsSortByTimecreated ListBdsMetastoreConfigurationsSortByEnum = "timeCreated"
	ListBdsMetastoreConfigurationsSortByDisplayname ListBdsMetastoreConfigurationsSortByEnum = "displayName"
)

var mappingListBdsMetastoreConfigurationsSortByEnum = map[string]ListBdsMetastoreConfigurationsSortByEnum{
	"timeCreated": ListBdsMetastoreConfigurationsSortByTimecreated,
	"displayName": ListBdsMetastoreConfigurationsSortByDisplayname,
}

var mappingListBdsMetastoreConfigurationsSortByEnumLowerCase = map[string]ListBdsMetastoreConfigurationsSortByEnum{
	"timecreated": ListBdsMetastoreConfigurationsSortByTimecreated,
	"displayname": ListBdsMetastoreConfigurationsSortByDisplayname,
}

// GetListBdsMetastoreConfigurationsSortByEnumValues Enumerates the set of values for ListBdsMetastoreConfigurationsSortByEnum
func GetListBdsMetastoreConfigurationsSortByEnumValues() []ListBdsMetastoreConfigurationsSortByEnum {
	values := make([]ListBdsMetastoreConfigurationsSortByEnum, 0)
	for _, v := range mappingListBdsMetastoreConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsMetastoreConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListBdsMetastoreConfigurationsSortByEnum
func GetListBdsMetastoreConfigurationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBdsMetastoreConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsMetastoreConfigurationsSortByEnum(val string) (ListBdsMetastoreConfigurationsSortByEnum, bool) {
	enum, ok := mappingListBdsMetastoreConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBdsMetastoreConfigurationsSortOrderEnum Enum with underlying type: string
type ListBdsMetastoreConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListBdsMetastoreConfigurationsSortOrderEnum
const (
	ListBdsMetastoreConfigurationsSortOrderAsc  ListBdsMetastoreConfigurationsSortOrderEnum = "ASC"
	ListBdsMetastoreConfigurationsSortOrderDesc ListBdsMetastoreConfigurationsSortOrderEnum = "DESC"
)

var mappingListBdsMetastoreConfigurationsSortOrderEnum = map[string]ListBdsMetastoreConfigurationsSortOrderEnum{
	"ASC":  ListBdsMetastoreConfigurationsSortOrderAsc,
	"DESC": ListBdsMetastoreConfigurationsSortOrderDesc,
}

var mappingListBdsMetastoreConfigurationsSortOrderEnumLowerCase = map[string]ListBdsMetastoreConfigurationsSortOrderEnum{
	"asc":  ListBdsMetastoreConfigurationsSortOrderAsc,
	"desc": ListBdsMetastoreConfigurationsSortOrderDesc,
}

// GetListBdsMetastoreConfigurationsSortOrderEnumValues Enumerates the set of values for ListBdsMetastoreConfigurationsSortOrderEnum
func GetListBdsMetastoreConfigurationsSortOrderEnumValues() []ListBdsMetastoreConfigurationsSortOrderEnum {
	values := make([]ListBdsMetastoreConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListBdsMetastoreConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsMetastoreConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListBdsMetastoreConfigurationsSortOrderEnum
func GetListBdsMetastoreConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBdsMetastoreConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsMetastoreConfigurationsSortOrderEnum(val string) (ListBdsMetastoreConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListBdsMetastoreConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
