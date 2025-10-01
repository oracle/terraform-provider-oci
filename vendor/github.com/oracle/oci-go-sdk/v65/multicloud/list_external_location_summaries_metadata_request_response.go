// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExternalLocationSummariesMetadataRequest wrapper for the ListExternalLocationSummariesMetadata operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListExternalLocationSummariesMetadata.go.html to see an example of how to use ListExternalLocationSummariesMetadataRequest.
type ListExternalLocationSummariesMetadataRequest struct {

	// The subscription service name values from [ORACLEDBATAZURE, ORACLEDBATGOOGLE, ORACLEDBATAWS]
	SubscriptionServiceName ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum `mandatory:"true" contributesTo:"query" name:"subscriptionServiceName" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription in which to list resources.
	SubscriptionId *string `mandatory:"false" contributesTo:"query" name:"subscriptionId"`

	// The resource type query (i.e. dbsystem, instance etc.)
	EntityType ListExternalLocationSummariesMetadataEntityTypeEnum `mandatory:"false" contributesTo:"query" name:"entityType" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExternalLocationSummariesMetadataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListExternalLocationSummariesMetadataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalLocationSummariesMetadataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalLocationSummariesMetadataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalLocationSummariesMetadataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalLocationSummariesMetadataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalLocationSummariesMetadataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalLocationSummariesMetadataSubscriptionServiceNameEnum(string(request.SubscriptionServiceName)); !ok && request.SubscriptionServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionServiceName: %s. Supported values are: %s.", request.SubscriptionServiceName, strings.Join(GetListExternalLocationSummariesMetadataSubscriptionServiceNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalLocationSummariesMetadataEntityTypeEnum(string(request.EntityType)); !ok && request.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", request.EntityType, strings.Join(GetListExternalLocationSummariesMetadataEntityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalLocationSummariesMetadataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalLocationSummariesMetadataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalLocationSummariesMetadataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalLocationSummariesMetadataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalLocationSummariesMetadataResponse wrapper for the ListExternalLocationSummariesMetadata operation
type ListExternalLocationSummariesMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalLocationSummariesMetadatumSummaryCollection instances
	ExternalLocationSummariesMetadatumSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalLocationSummariesMetadataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalLocationSummariesMetadataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum Enum with underlying type: string
type ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum string

// Set of constants representing the allowable values for ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum
const (
	ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbatazure  ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum = "ORACLEDBATAZURE"
	ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbatgoogle ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum = "ORACLEDBATGOOGLE"
	ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbataws    ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum = "ORACLEDBATAWS"
)

var mappingListExternalLocationSummariesMetadataSubscriptionServiceNameEnum = map[string]ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum{
	"ORACLEDBATAZURE":  ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbatazure,
	"ORACLEDBATGOOGLE": ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbatgoogle,
	"ORACLEDBATAWS":    ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbataws,
}

var mappingListExternalLocationSummariesMetadataSubscriptionServiceNameEnumLowerCase = map[string]ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum{
	"oracledbatazure":  ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbatazure,
	"oracledbatgoogle": ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbatgoogle,
	"oracledbataws":    ListExternalLocationSummariesMetadataSubscriptionServiceNameOracledbataws,
}

// GetListExternalLocationSummariesMetadataSubscriptionServiceNameEnumValues Enumerates the set of values for ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum
func GetListExternalLocationSummariesMetadataSubscriptionServiceNameEnumValues() []ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum {
	values := make([]ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum, 0)
	for _, v := range mappingListExternalLocationSummariesMetadataSubscriptionServiceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationSummariesMetadataSubscriptionServiceNameEnumStringValues Enumerates the set of values in String for ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum
func GetListExternalLocationSummariesMetadataSubscriptionServiceNameEnumStringValues() []string {
	return []string{
		"ORACLEDBATAZURE",
		"ORACLEDBATGOOGLE",
		"ORACLEDBATAWS",
	}
}

// GetMappingListExternalLocationSummariesMetadataSubscriptionServiceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationSummariesMetadataSubscriptionServiceNameEnum(val string) (ListExternalLocationSummariesMetadataSubscriptionServiceNameEnum, bool) {
	enum, ok := mappingListExternalLocationSummariesMetadataSubscriptionServiceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalLocationSummariesMetadataEntityTypeEnum Enum with underlying type: string
type ListExternalLocationSummariesMetadataEntityTypeEnum string

// Set of constants representing the allowable values for ListExternalLocationSummariesMetadataEntityTypeEnum
const (
	ListExternalLocationSummariesMetadataEntityTypeDbsystem ListExternalLocationSummariesMetadataEntityTypeEnum = "dbsystem"
)

var mappingListExternalLocationSummariesMetadataEntityTypeEnum = map[string]ListExternalLocationSummariesMetadataEntityTypeEnum{
	"dbsystem": ListExternalLocationSummariesMetadataEntityTypeDbsystem,
}

var mappingListExternalLocationSummariesMetadataEntityTypeEnumLowerCase = map[string]ListExternalLocationSummariesMetadataEntityTypeEnum{
	"dbsystem": ListExternalLocationSummariesMetadataEntityTypeDbsystem,
}

// GetListExternalLocationSummariesMetadataEntityTypeEnumValues Enumerates the set of values for ListExternalLocationSummariesMetadataEntityTypeEnum
func GetListExternalLocationSummariesMetadataEntityTypeEnumValues() []ListExternalLocationSummariesMetadataEntityTypeEnum {
	values := make([]ListExternalLocationSummariesMetadataEntityTypeEnum, 0)
	for _, v := range mappingListExternalLocationSummariesMetadataEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationSummariesMetadataEntityTypeEnumStringValues Enumerates the set of values in String for ListExternalLocationSummariesMetadataEntityTypeEnum
func GetListExternalLocationSummariesMetadataEntityTypeEnumStringValues() []string {
	return []string{
		"dbsystem",
	}
}

// GetMappingListExternalLocationSummariesMetadataEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationSummariesMetadataEntityTypeEnum(val string) (ListExternalLocationSummariesMetadataEntityTypeEnum, bool) {
	enum, ok := mappingListExternalLocationSummariesMetadataEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalLocationSummariesMetadataSortOrderEnum Enum with underlying type: string
type ListExternalLocationSummariesMetadataSortOrderEnum string

// Set of constants representing the allowable values for ListExternalLocationSummariesMetadataSortOrderEnum
const (
	ListExternalLocationSummariesMetadataSortOrderAsc  ListExternalLocationSummariesMetadataSortOrderEnum = "ASC"
	ListExternalLocationSummariesMetadataSortOrderDesc ListExternalLocationSummariesMetadataSortOrderEnum = "DESC"
)

var mappingListExternalLocationSummariesMetadataSortOrderEnum = map[string]ListExternalLocationSummariesMetadataSortOrderEnum{
	"ASC":  ListExternalLocationSummariesMetadataSortOrderAsc,
	"DESC": ListExternalLocationSummariesMetadataSortOrderDesc,
}

var mappingListExternalLocationSummariesMetadataSortOrderEnumLowerCase = map[string]ListExternalLocationSummariesMetadataSortOrderEnum{
	"asc":  ListExternalLocationSummariesMetadataSortOrderAsc,
	"desc": ListExternalLocationSummariesMetadataSortOrderDesc,
}

// GetListExternalLocationSummariesMetadataSortOrderEnumValues Enumerates the set of values for ListExternalLocationSummariesMetadataSortOrderEnum
func GetListExternalLocationSummariesMetadataSortOrderEnumValues() []ListExternalLocationSummariesMetadataSortOrderEnum {
	values := make([]ListExternalLocationSummariesMetadataSortOrderEnum, 0)
	for _, v := range mappingListExternalLocationSummariesMetadataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationSummariesMetadataSortOrderEnumStringValues Enumerates the set of values in String for ListExternalLocationSummariesMetadataSortOrderEnum
func GetListExternalLocationSummariesMetadataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalLocationSummariesMetadataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationSummariesMetadataSortOrderEnum(val string) (ListExternalLocationSummariesMetadataSortOrderEnum, bool) {
	enum, ok := mappingListExternalLocationSummariesMetadataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalLocationSummariesMetadataSortByEnum Enum with underlying type: string
type ListExternalLocationSummariesMetadataSortByEnum string

// Set of constants representing the allowable values for ListExternalLocationSummariesMetadataSortByEnum
const (
	ListExternalLocationSummariesMetadataSortByTimecreated ListExternalLocationSummariesMetadataSortByEnum = "timeCreated"
	ListExternalLocationSummariesMetadataSortByDisplayname ListExternalLocationSummariesMetadataSortByEnum = "displayName"
)

var mappingListExternalLocationSummariesMetadataSortByEnum = map[string]ListExternalLocationSummariesMetadataSortByEnum{
	"timeCreated": ListExternalLocationSummariesMetadataSortByTimecreated,
	"displayName": ListExternalLocationSummariesMetadataSortByDisplayname,
}

var mappingListExternalLocationSummariesMetadataSortByEnumLowerCase = map[string]ListExternalLocationSummariesMetadataSortByEnum{
	"timecreated": ListExternalLocationSummariesMetadataSortByTimecreated,
	"displayname": ListExternalLocationSummariesMetadataSortByDisplayname,
}

// GetListExternalLocationSummariesMetadataSortByEnumValues Enumerates the set of values for ListExternalLocationSummariesMetadataSortByEnum
func GetListExternalLocationSummariesMetadataSortByEnumValues() []ListExternalLocationSummariesMetadataSortByEnum {
	values := make([]ListExternalLocationSummariesMetadataSortByEnum, 0)
	for _, v := range mappingListExternalLocationSummariesMetadataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationSummariesMetadataSortByEnumStringValues Enumerates the set of values in String for ListExternalLocationSummariesMetadataSortByEnum
func GetListExternalLocationSummariesMetadataSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListExternalLocationSummariesMetadataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationSummariesMetadataSortByEnum(val string) (ListExternalLocationSummariesMetadataSortByEnum, bool) {
	enum, ok := mappingListExternalLocationSummariesMetadataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
