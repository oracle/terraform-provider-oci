// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMonitoringResourcesRequest wrapper for the ListMonitoringResources operation
type ListMonitoringResourcesRequest struct {

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Target resource type
	TargetResourceType ListMonitoringResourcesTargetResourceTypeEnum `mandatory:"false" contributesTo:"query" name:"targetResourceType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMonitoringResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListMonitoringResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMonitoringResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitoringResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitoringResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitoringResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMonitoringResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMonitoringResourcesTargetResourceTypeEnum(string(request.TargetResourceType)); !ok && request.TargetResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetResourceType: %s. Supported values are: %s.", request.TargetResourceType, strings.Join(GetListMonitoringResourcesTargetResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoringResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMonitoringResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitoringResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMonitoringResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMonitoringResourcesResponse wrapper for the ListMonitoringResources operation
type ListMonitoringResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoringResourceCollection instances
	MonitoringResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMonitoringResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitoringResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitoringResourcesTargetResourceTypeEnum Enum with underlying type: string
type ListMonitoringResourcesTargetResourceTypeEnum string

// Set of constants representing the allowable values for ListMonitoringResourcesTargetResourceTypeEnum
const (
	ListMonitoringResourcesTargetResourceTypeCompartment  ListMonitoringResourcesTargetResourceTypeEnum = "COMPARTMENT"
	ListMonitoringResourcesTargetResourceTypeErpcloud     ListMonitoringResourcesTargetResourceTypeEnum = "ERPCLOUD"
	ListMonitoringResourcesTargetResourceTypeHcmcloud     ListMonitoringResourcesTargetResourceTypeEnum = "HCMCLOUD"
	ListMonitoringResourcesTargetResourceTypeFacloud      ListMonitoringResourcesTargetResourceTypeEnum = "FACLOUD"
	ListMonitoringResourcesTargetResourceTypeSecurityZone ListMonitoringResourcesTargetResourceTypeEnum = "SECURITY_ZONE"
)

var mappingListMonitoringResourcesTargetResourceTypeEnum = map[string]ListMonitoringResourcesTargetResourceTypeEnum{
	"COMPARTMENT":   ListMonitoringResourcesTargetResourceTypeCompartment,
	"ERPCLOUD":      ListMonitoringResourcesTargetResourceTypeErpcloud,
	"HCMCLOUD":      ListMonitoringResourcesTargetResourceTypeHcmcloud,
	"FACLOUD":       ListMonitoringResourcesTargetResourceTypeFacloud,
	"SECURITY_ZONE": ListMonitoringResourcesTargetResourceTypeSecurityZone,
}

var mappingListMonitoringResourcesTargetResourceTypeEnumLowerCase = map[string]ListMonitoringResourcesTargetResourceTypeEnum{
	"compartment":   ListMonitoringResourcesTargetResourceTypeCompartment,
	"erpcloud":      ListMonitoringResourcesTargetResourceTypeErpcloud,
	"hcmcloud":      ListMonitoringResourcesTargetResourceTypeHcmcloud,
	"facloud":       ListMonitoringResourcesTargetResourceTypeFacloud,
	"security_zone": ListMonitoringResourcesTargetResourceTypeSecurityZone,
}

// GetListMonitoringResourcesTargetResourceTypeEnumValues Enumerates the set of values for ListMonitoringResourcesTargetResourceTypeEnum
func GetListMonitoringResourcesTargetResourceTypeEnumValues() []ListMonitoringResourcesTargetResourceTypeEnum {
	values := make([]ListMonitoringResourcesTargetResourceTypeEnum, 0)
	for _, v := range mappingListMonitoringResourcesTargetResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoringResourcesTargetResourceTypeEnumStringValues Enumerates the set of values in String for ListMonitoringResourcesTargetResourceTypeEnum
func GetListMonitoringResourcesTargetResourceTypeEnumStringValues() []string {
	return []string{
		"COMPARTMENT",
		"ERPCLOUD",
		"HCMCLOUD",
		"FACLOUD",
		"SECURITY_ZONE",
	}
}

// GetMappingListMonitoringResourcesTargetResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoringResourcesTargetResourceTypeEnum(val string) (ListMonitoringResourcesTargetResourceTypeEnum, bool) {
	enum, ok := mappingListMonitoringResourcesTargetResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoringResourcesSortOrderEnum Enum with underlying type: string
type ListMonitoringResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListMonitoringResourcesSortOrderEnum
const (
	ListMonitoringResourcesSortOrderAsc  ListMonitoringResourcesSortOrderEnum = "ASC"
	ListMonitoringResourcesSortOrderDesc ListMonitoringResourcesSortOrderEnum = "DESC"
)

var mappingListMonitoringResourcesSortOrderEnum = map[string]ListMonitoringResourcesSortOrderEnum{
	"ASC":  ListMonitoringResourcesSortOrderAsc,
	"DESC": ListMonitoringResourcesSortOrderDesc,
}

var mappingListMonitoringResourcesSortOrderEnumLowerCase = map[string]ListMonitoringResourcesSortOrderEnum{
	"asc":  ListMonitoringResourcesSortOrderAsc,
	"desc": ListMonitoringResourcesSortOrderDesc,
}

// GetListMonitoringResourcesSortOrderEnumValues Enumerates the set of values for ListMonitoringResourcesSortOrderEnum
func GetListMonitoringResourcesSortOrderEnumValues() []ListMonitoringResourcesSortOrderEnum {
	values := make([]ListMonitoringResourcesSortOrderEnum, 0)
	for _, v := range mappingListMonitoringResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoringResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListMonitoringResourcesSortOrderEnum
func GetListMonitoringResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMonitoringResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoringResourcesSortOrderEnum(val string) (ListMonitoringResourcesSortOrderEnum, bool) {
	enum, ok := mappingListMonitoringResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitoringResourcesSortByEnum Enum with underlying type: string
type ListMonitoringResourcesSortByEnum string

// Set of constants representing the allowable values for ListMonitoringResourcesSortByEnum
const (
	ListMonitoringResourcesSortByTimecreated ListMonitoringResourcesSortByEnum = "timeCreated"
	ListMonitoringResourcesSortByDisplayname ListMonitoringResourcesSortByEnum = "displayName"
)

var mappingListMonitoringResourcesSortByEnum = map[string]ListMonitoringResourcesSortByEnum{
	"timeCreated": ListMonitoringResourcesSortByTimecreated,
	"displayName": ListMonitoringResourcesSortByDisplayname,
}

var mappingListMonitoringResourcesSortByEnumLowerCase = map[string]ListMonitoringResourcesSortByEnum{
	"timecreated": ListMonitoringResourcesSortByTimecreated,
	"displayname": ListMonitoringResourcesSortByDisplayname,
}

// GetListMonitoringResourcesSortByEnumValues Enumerates the set of values for ListMonitoringResourcesSortByEnum
func GetListMonitoringResourcesSortByEnumValues() []ListMonitoringResourcesSortByEnum {
	values := make([]ListMonitoringResourcesSortByEnum, 0)
	for _, v := range mappingListMonitoringResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitoringResourcesSortByEnumStringValues Enumerates the set of values in String for ListMonitoringResourcesSortByEnum
func GetListMonitoringResourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMonitoringResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitoringResourcesSortByEnum(val string) (ListMonitoringResourcesSortByEnum, bool) {
	enum, ok := mappingListMonitoringResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
