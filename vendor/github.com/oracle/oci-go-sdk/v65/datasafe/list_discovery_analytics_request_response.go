// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDiscoveryAnalyticsRequest wrapper for the ListDiscoveryAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDiscoveryAnalytics.go.html to see an example of how to use ListDiscoveryAnalyticsRequest.
type ListDiscoveryAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Attribute by which the discovery analytics data should be grouped.
	GroupBy ListDiscoveryAnalyticsGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return the target database group that matches the specified OCID.
	TargetDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupId"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for all the fields is ascending.
	SortBy ListDiscoveryAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDiscoveryAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only the resources that match the specified sensitive data model OCID.
	SensitiveDataModelId *string `mandatory:"false" contributesTo:"query" name:"sensitiveDataModelId"`

	// A filter to return only items related to a specific sensitive type OCID.
	SensitiveTypeId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypeId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only the common sensitive type resources. Common sensitive types belong to
	// library sensitive types which are frequently used to perform sensitive data discovery.
	IsCommon *bool `mandatory:"false" contributesTo:"query" name:"isCommon"`

	// An optional filter to return only resources that match the specified OCID of the sensitive type group resource.
	SensitiveTypeGroupId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypeGroupId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDiscoveryAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDiscoveryAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDiscoveryAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDiscoveryAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDiscoveryAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDiscoveryAnalyticsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetListDiscoveryAnalyticsGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDiscoveryAnalyticsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDiscoveryAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDiscoveryAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDiscoveryAnalyticsResponse wrapper for the ListDiscoveryAnalytics operation
type ListDiscoveryAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DiscoveryAnalyticsCollection instances
	DiscoveryAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDiscoveryAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDiscoveryAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDiscoveryAnalyticsGroupByEnum Enum with underlying type: string
type ListDiscoveryAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListDiscoveryAnalyticsGroupByEnum
const (
	ListDiscoveryAnalyticsGroupByTargetid                               ListDiscoveryAnalyticsGroupByEnum = "targetId"
	ListDiscoveryAnalyticsGroupBySensitivedatamodelid                   ListDiscoveryAnalyticsGroupByEnum = "sensitiveDataModelId"
	ListDiscoveryAnalyticsGroupBySensitivetypeid                        ListDiscoveryAnalyticsGroupByEnum = "sensitiveTypeId"
	ListDiscoveryAnalyticsGroupByTargetidandsensitivedatamodelid        ListDiscoveryAnalyticsGroupByEnum = "targetIdAndSensitiveDataModelId"
	ListDiscoveryAnalyticsGroupBySensitivetypeidandtargetid             ListDiscoveryAnalyticsGroupByEnum = "sensitiveTypeIdAndTargetId"
	ListDiscoveryAnalyticsGroupBySensitivetypeidandsensitivedatamodelid ListDiscoveryAnalyticsGroupByEnum = "sensitiveTypeIdAndSensitiveDataModelId"
)

var mappingListDiscoveryAnalyticsGroupByEnum = map[string]ListDiscoveryAnalyticsGroupByEnum{
	"targetId":                               ListDiscoveryAnalyticsGroupByTargetid,
	"sensitiveDataModelId":                   ListDiscoveryAnalyticsGroupBySensitivedatamodelid,
	"sensitiveTypeId":                        ListDiscoveryAnalyticsGroupBySensitivetypeid,
	"targetIdAndSensitiveDataModelId":        ListDiscoveryAnalyticsGroupByTargetidandsensitivedatamodelid,
	"sensitiveTypeIdAndTargetId":             ListDiscoveryAnalyticsGroupBySensitivetypeidandtargetid,
	"sensitiveTypeIdAndSensitiveDataModelId": ListDiscoveryAnalyticsGroupBySensitivetypeidandsensitivedatamodelid,
}

var mappingListDiscoveryAnalyticsGroupByEnumLowerCase = map[string]ListDiscoveryAnalyticsGroupByEnum{
	"targetid":                               ListDiscoveryAnalyticsGroupByTargetid,
	"sensitivedatamodelid":                   ListDiscoveryAnalyticsGroupBySensitivedatamodelid,
	"sensitivetypeid":                        ListDiscoveryAnalyticsGroupBySensitivetypeid,
	"targetidandsensitivedatamodelid":        ListDiscoveryAnalyticsGroupByTargetidandsensitivedatamodelid,
	"sensitivetypeidandtargetid":             ListDiscoveryAnalyticsGroupBySensitivetypeidandtargetid,
	"sensitivetypeidandsensitivedatamodelid": ListDiscoveryAnalyticsGroupBySensitivetypeidandsensitivedatamodelid,
}

// GetListDiscoveryAnalyticsGroupByEnumValues Enumerates the set of values for ListDiscoveryAnalyticsGroupByEnum
func GetListDiscoveryAnalyticsGroupByEnumValues() []ListDiscoveryAnalyticsGroupByEnum {
	values := make([]ListDiscoveryAnalyticsGroupByEnum, 0)
	for _, v := range mappingListDiscoveryAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListDiscoveryAnalyticsGroupByEnum
func GetListDiscoveryAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"targetId",
		"sensitiveDataModelId",
		"sensitiveTypeId",
		"targetIdAndSensitiveDataModelId",
		"sensitiveTypeIdAndTargetId",
		"sensitiveTypeIdAndSensitiveDataModelId",
	}
}

// GetMappingListDiscoveryAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryAnalyticsGroupByEnum(val string) (ListDiscoveryAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListDiscoveryAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryAnalyticsSortByEnum Enum with underlying type: string
type ListDiscoveryAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListDiscoveryAnalyticsSortByEnum
const (
	ListDiscoveryAnalyticsSortByTimelastdiscovered ListDiscoveryAnalyticsSortByEnum = "timeLastDiscovered"
)

var mappingListDiscoveryAnalyticsSortByEnum = map[string]ListDiscoveryAnalyticsSortByEnum{
	"timeLastDiscovered": ListDiscoveryAnalyticsSortByTimelastdiscovered,
}

var mappingListDiscoveryAnalyticsSortByEnumLowerCase = map[string]ListDiscoveryAnalyticsSortByEnum{
	"timelastdiscovered": ListDiscoveryAnalyticsSortByTimelastdiscovered,
}

// GetListDiscoveryAnalyticsSortByEnumValues Enumerates the set of values for ListDiscoveryAnalyticsSortByEnum
func GetListDiscoveryAnalyticsSortByEnumValues() []ListDiscoveryAnalyticsSortByEnum {
	values := make([]ListDiscoveryAnalyticsSortByEnum, 0)
	for _, v := range mappingListDiscoveryAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListDiscoveryAnalyticsSortByEnum
func GetListDiscoveryAnalyticsSortByEnumStringValues() []string {
	return []string{
		"timeLastDiscovered",
	}
}

// GetMappingListDiscoveryAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryAnalyticsSortByEnum(val string) (ListDiscoveryAnalyticsSortByEnum, bool) {
	enum, ok := mappingListDiscoveryAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDiscoveryAnalyticsSortOrderEnum Enum with underlying type: string
type ListDiscoveryAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListDiscoveryAnalyticsSortOrderEnum
const (
	ListDiscoveryAnalyticsSortOrderAsc  ListDiscoveryAnalyticsSortOrderEnum = "ASC"
	ListDiscoveryAnalyticsSortOrderDesc ListDiscoveryAnalyticsSortOrderEnum = "DESC"
)

var mappingListDiscoveryAnalyticsSortOrderEnum = map[string]ListDiscoveryAnalyticsSortOrderEnum{
	"ASC":  ListDiscoveryAnalyticsSortOrderAsc,
	"DESC": ListDiscoveryAnalyticsSortOrderDesc,
}

var mappingListDiscoveryAnalyticsSortOrderEnumLowerCase = map[string]ListDiscoveryAnalyticsSortOrderEnum{
	"asc":  ListDiscoveryAnalyticsSortOrderAsc,
	"desc": ListDiscoveryAnalyticsSortOrderDesc,
}

// GetListDiscoveryAnalyticsSortOrderEnumValues Enumerates the set of values for ListDiscoveryAnalyticsSortOrderEnum
func GetListDiscoveryAnalyticsSortOrderEnumValues() []ListDiscoveryAnalyticsSortOrderEnum {
	values := make([]ListDiscoveryAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListDiscoveryAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDiscoveryAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListDiscoveryAnalyticsSortOrderEnum
func GetListDiscoveryAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDiscoveryAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDiscoveryAnalyticsSortOrderEnum(val string) (ListDiscoveryAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListDiscoveryAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
