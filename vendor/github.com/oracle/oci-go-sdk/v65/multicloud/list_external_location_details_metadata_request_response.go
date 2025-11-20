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

// ListExternalLocationDetailsMetadataRequest wrapper for the ListExternalLocationDetailsMetadata operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListExternalLocationDetailsMetadata.go.html to see an example of how to use ListExternalLocationDetailsMetadataRequest.
type ListExternalLocationDetailsMetadataRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
	SubscriptionId *string `mandatory:"true" contributesTo:"query" name:"subscriptionId"`

	// The subscription service name of the Cloud Service Provider.
	SubscriptionServiceName ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum `mandatory:"true" contributesTo:"query" name:"subscriptionServiceName" omitEmpty:"true"`

	// The resource type query (i.e. dbsystem, instance etc.)
	EntityType ListExternalLocationDetailsMetadataEntityTypeEnum `mandatory:"false" contributesTo:"query" name:"entityType" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud base compartment in which to list resources.
	// A Multicloud base compartment is an OCI compartment that maps to a subscription in a Cloud Service Provider (such as Azure, AWS, or Google Cloud).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment linked to the resource.
	LinkedCompartmentId *string `mandatory:"false" contributesTo:"query" name:"linkedCompartmentId"`

	// The Cloud Service Provider region.
	ExternalLocation *string `mandatory:"false" contributesTo:"query" name:"externalLocation"`

	// OCI Logical AD to filter the response.
	LogicalZone *string `mandatory:"false" contributesTo:"query" name:"logicalZone"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cluster Placement Group.
	ClusterPlacementGroupId *string `mandatory:"false" contributesTo:"query" name:"clusterPlacementGroupId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExternalLocationDetailsMetadataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListExternalLocationDetailsMetadataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalLocationDetailsMetadataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalLocationDetailsMetadataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalLocationDetailsMetadataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalLocationDetailsMetadataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalLocationDetailsMetadataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalLocationDetailsMetadataSubscriptionServiceNameEnum(string(request.SubscriptionServiceName)); !ok && request.SubscriptionServiceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionServiceName: %s. Supported values are: %s.", request.SubscriptionServiceName, strings.Join(GetListExternalLocationDetailsMetadataSubscriptionServiceNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalLocationDetailsMetadataEntityTypeEnum(string(request.EntityType)); !ok && request.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", request.EntityType, strings.Join(GetListExternalLocationDetailsMetadataEntityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalLocationDetailsMetadataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalLocationDetailsMetadataSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalLocationDetailsMetadataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalLocationDetailsMetadataSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalLocationDetailsMetadataResponse wrapper for the ListExternalLocationDetailsMetadata operation
type ListExternalLocationDetailsMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalLocationsMetadatumCollection instances
	ExternalLocationsMetadatumCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalLocationDetailsMetadataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalLocationDetailsMetadataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum Enum with underlying type: string
type ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum string

// Set of constants representing the allowable values for ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum
const (
	ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbatazure  ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum = "ORACLEDBATAZURE"
	ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbatgoogle ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum = "ORACLEDBATGOOGLE"
	ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbataws    ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum = "ORACLEDBATAWS"
)

var mappingListExternalLocationDetailsMetadataSubscriptionServiceNameEnum = map[string]ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum{
	"ORACLEDBATAZURE":  ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbatazure,
	"ORACLEDBATGOOGLE": ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbatgoogle,
	"ORACLEDBATAWS":    ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbataws,
}

var mappingListExternalLocationDetailsMetadataSubscriptionServiceNameEnumLowerCase = map[string]ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum{
	"oracledbatazure":  ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbatazure,
	"oracledbatgoogle": ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbatgoogle,
	"oracledbataws":    ListExternalLocationDetailsMetadataSubscriptionServiceNameOracledbataws,
}

// GetListExternalLocationDetailsMetadataSubscriptionServiceNameEnumValues Enumerates the set of values for ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum
func GetListExternalLocationDetailsMetadataSubscriptionServiceNameEnumValues() []ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum {
	values := make([]ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum, 0)
	for _, v := range mappingListExternalLocationDetailsMetadataSubscriptionServiceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationDetailsMetadataSubscriptionServiceNameEnumStringValues Enumerates the set of values in String for ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum
func GetListExternalLocationDetailsMetadataSubscriptionServiceNameEnumStringValues() []string {
	return []string{
		"ORACLEDBATAZURE",
		"ORACLEDBATGOOGLE",
		"ORACLEDBATAWS",
	}
}

// GetMappingListExternalLocationDetailsMetadataSubscriptionServiceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationDetailsMetadataSubscriptionServiceNameEnum(val string) (ListExternalLocationDetailsMetadataSubscriptionServiceNameEnum, bool) {
	enum, ok := mappingListExternalLocationDetailsMetadataSubscriptionServiceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalLocationDetailsMetadataEntityTypeEnum Enum with underlying type: string
type ListExternalLocationDetailsMetadataEntityTypeEnum string

// Set of constants representing the allowable values for ListExternalLocationDetailsMetadataEntityTypeEnum
const (
	ListExternalLocationDetailsMetadataEntityTypeDbsystem ListExternalLocationDetailsMetadataEntityTypeEnum = "dbsystem"
)

var mappingListExternalLocationDetailsMetadataEntityTypeEnum = map[string]ListExternalLocationDetailsMetadataEntityTypeEnum{
	"dbsystem": ListExternalLocationDetailsMetadataEntityTypeDbsystem,
}

var mappingListExternalLocationDetailsMetadataEntityTypeEnumLowerCase = map[string]ListExternalLocationDetailsMetadataEntityTypeEnum{
	"dbsystem": ListExternalLocationDetailsMetadataEntityTypeDbsystem,
}

// GetListExternalLocationDetailsMetadataEntityTypeEnumValues Enumerates the set of values for ListExternalLocationDetailsMetadataEntityTypeEnum
func GetListExternalLocationDetailsMetadataEntityTypeEnumValues() []ListExternalLocationDetailsMetadataEntityTypeEnum {
	values := make([]ListExternalLocationDetailsMetadataEntityTypeEnum, 0)
	for _, v := range mappingListExternalLocationDetailsMetadataEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationDetailsMetadataEntityTypeEnumStringValues Enumerates the set of values in String for ListExternalLocationDetailsMetadataEntityTypeEnum
func GetListExternalLocationDetailsMetadataEntityTypeEnumStringValues() []string {
	return []string{
		"dbsystem",
	}
}

// GetMappingListExternalLocationDetailsMetadataEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationDetailsMetadataEntityTypeEnum(val string) (ListExternalLocationDetailsMetadataEntityTypeEnum, bool) {
	enum, ok := mappingListExternalLocationDetailsMetadataEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalLocationDetailsMetadataSortOrderEnum Enum with underlying type: string
type ListExternalLocationDetailsMetadataSortOrderEnum string

// Set of constants representing the allowable values for ListExternalLocationDetailsMetadataSortOrderEnum
const (
	ListExternalLocationDetailsMetadataSortOrderAsc  ListExternalLocationDetailsMetadataSortOrderEnum = "ASC"
	ListExternalLocationDetailsMetadataSortOrderDesc ListExternalLocationDetailsMetadataSortOrderEnum = "DESC"
)

var mappingListExternalLocationDetailsMetadataSortOrderEnum = map[string]ListExternalLocationDetailsMetadataSortOrderEnum{
	"ASC":  ListExternalLocationDetailsMetadataSortOrderAsc,
	"DESC": ListExternalLocationDetailsMetadataSortOrderDesc,
}

var mappingListExternalLocationDetailsMetadataSortOrderEnumLowerCase = map[string]ListExternalLocationDetailsMetadataSortOrderEnum{
	"asc":  ListExternalLocationDetailsMetadataSortOrderAsc,
	"desc": ListExternalLocationDetailsMetadataSortOrderDesc,
}

// GetListExternalLocationDetailsMetadataSortOrderEnumValues Enumerates the set of values for ListExternalLocationDetailsMetadataSortOrderEnum
func GetListExternalLocationDetailsMetadataSortOrderEnumValues() []ListExternalLocationDetailsMetadataSortOrderEnum {
	values := make([]ListExternalLocationDetailsMetadataSortOrderEnum, 0)
	for _, v := range mappingListExternalLocationDetailsMetadataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationDetailsMetadataSortOrderEnumStringValues Enumerates the set of values in String for ListExternalLocationDetailsMetadataSortOrderEnum
func GetListExternalLocationDetailsMetadataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalLocationDetailsMetadataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationDetailsMetadataSortOrderEnum(val string) (ListExternalLocationDetailsMetadataSortOrderEnum, bool) {
	enum, ok := mappingListExternalLocationDetailsMetadataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalLocationDetailsMetadataSortByEnum Enum with underlying type: string
type ListExternalLocationDetailsMetadataSortByEnum string

// Set of constants representing the allowable values for ListExternalLocationDetailsMetadataSortByEnum
const (
	ListExternalLocationDetailsMetadataSortByTimecreated ListExternalLocationDetailsMetadataSortByEnum = "timeCreated"
	ListExternalLocationDetailsMetadataSortByDisplayname ListExternalLocationDetailsMetadataSortByEnum = "displayName"
)

var mappingListExternalLocationDetailsMetadataSortByEnum = map[string]ListExternalLocationDetailsMetadataSortByEnum{
	"timeCreated": ListExternalLocationDetailsMetadataSortByTimecreated,
	"displayName": ListExternalLocationDetailsMetadataSortByDisplayname,
}

var mappingListExternalLocationDetailsMetadataSortByEnumLowerCase = map[string]ListExternalLocationDetailsMetadataSortByEnum{
	"timecreated": ListExternalLocationDetailsMetadataSortByTimecreated,
	"displayname": ListExternalLocationDetailsMetadataSortByDisplayname,
}

// GetListExternalLocationDetailsMetadataSortByEnumValues Enumerates the set of values for ListExternalLocationDetailsMetadataSortByEnum
func GetListExternalLocationDetailsMetadataSortByEnumValues() []ListExternalLocationDetailsMetadataSortByEnum {
	values := make([]ListExternalLocationDetailsMetadataSortByEnum, 0)
	for _, v := range mappingListExternalLocationDetailsMetadataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalLocationDetailsMetadataSortByEnumStringValues Enumerates the set of values in String for ListExternalLocationDetailsMetadataSortByEnum
func GetListExternalLocationDetailsMetadataSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListExternalLocationDetailsMetadataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalLocationDetailsMetadataSortByEnum(val string) (ListExternalLocationDetailsMetadataSortByEnum, bool) {
	enum, ok := mappingListExternalLocationDetailsMetadataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
