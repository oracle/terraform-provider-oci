// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCloudAutonomousVmClustersRequest wrapper for the ListCloudAutonomousVmClusters operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudAutonomousVmClusters.go.html to see an example of how to use ListCloudAutonomousVmClustersRequest.
type ListCloudAutonomousVmClustersRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// If provided, filters the results for the specified cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"cloudExadataInfrastructureId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListCloudAutonomousVmClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCloudAutonomousVmClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState CloudAutonomousVmClusterSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given availability domain exactly.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudAutonomousVmClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudAutonomousVmClustersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudAutonomousVmClustersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudAutonomousVmClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudAutonomousVmClustersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudAutonomousVmClustersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudAutonomousVmClustersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudAutonomousVmClustersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudAutonomousVmClustersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCloudAutonomousVmClusterSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCloudAutonomousVmClusterSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudAutonomousVmClustersResponse wrapper for the ListCloudAutonomousVmClusters operation
type ListCloudAutonomousVmClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CloudAutonomousVmClusterSummary instances
	Items []CloudAutonomousVmClusterSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudAutonomousVmClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudAutonomousVmClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudAutonomousVmClustersSortByEnum Enum with underlying type: string
type ListCloudAutonomousVmClustersSortByEnum string

// Set of constants representing the allowable values for ListCloudAutonomousVmClustersSortByEnum
const (
	ListCloudAutonomousVmClustersSortByTimecreated ListCloudAutonomousVmClustersSortByEnum = "TIMECREATED"
	ListCloudAutonomousVmClustersSortByDisplayname ListCloudAutonomousVmClustersSortByEnum = "DISPLAYNAME"
)

var mappingListCloudAutonomousVmClustersSortByEnum = map[string]ListCloudAutonomousVmClustersSortByEnum{
	"TIMECREATED": ListCloudAutonomousVmClustersSortByTimecreated,
	"DISPLAYNAME": ListCloudAutonomousVmClustersSortByDisplayname,
}

var mappingListCloudAutonomousVmClustersSortByEnumLowerCase = map[string]ListCloudAutonomousVmClustersSortByEnum{
	"timecreated": ListCloudAutonomousVmClustersSortByTimecreated,
	"displayname": ListCloudAutonomousVmClustersSortByDisplayname,
}

// GetListCloudAutonomousVmClustersSortByEnumValues Enumerates the set of values for ListCloudAutonomousVmClustersSortByEnum
func GetListCloudAutonomousVmClustersSortByEnumValues() []ListCloudAutonomousVmClustersSortByEnum {
	values := make([]ListCloudAutonomousVmClustersSortByEnum, 0)
	for _, v := range mappingListCloudAutonomousVmClustersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAutonomousVmClustersSortByEnumStringValues Enumerates the set of values in String for ListCloudAutonomousVmClustersSortByEnum
func GetListCloudAutonomousVmClustersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListCloudAutonomousVmClustersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAutonomousVmClustersSortByEnum(val string) (ListCloudAutonomousVmClustersSortByEnum, bool) {
	enum, ok := mappingListCloudAutonomousVmClustersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudAutonomousVmClustersSortOrderEnum Enum with underlying type: string
type ListCloudAutonomousVmClustersSortOrderEnum string

// Set of constants representing the allowable values for ListCloudAutonomousVmClustersSortOrderEnum
const (
	ListCloudAutonomousVmClustersSortOrderAsc  ListCloudAutonomousVmClustersSortOrderEnum = "ASC"
	ListCloudAutonomousVmClustersSortOrderDesc ListCloudAutonomousVmClustersSortOrderEnum = "DESC"
)

var mappingListCloudAutonomousVmClustersSortOrderEnum = map[string]ListCloudAutonomousVmClustersSortOrderEnum{
	"ASC":  ListCloudAutonomousVmClustersSortOrderAsc,
	"DESC": ListCloudAutonomousVmClustersSortOrderDesc,
}

var mappingListCloudAutonomousVmClustersSortOrderEnumLowerCase = map[string]ListCloudAutonomousVmClustersSortOrderEnum{
	"asc":  ListCloudAutonomousVmClustersSortOrderAsc,
	"desc": ListCloudAutonomousVmClustersSortOrderDesc,
}

// GetListCloudAutonomousVmClustersSortOrderEnumValues Enumerates the set of values for ListCloudAutonomousVmClustersSortOrderEnum
func GetListCloudAutonomousVmClustersSortOrderEnumValues() []ListCloudAutonomousVmClustersSortOrderEnum {
	values := make([]ListCloudAutonomousVmClustersSortOrderEnum, 0)
	for _, v := range mappingListCloudAutonomousVmClustersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudAutonomousVmClustersSortOrderEnumStringValues Enumerates the set of values in String for ListCloudAutonomousVmClustersSortOrderEnum
func GetListCloudAutonomousVmClustersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudAutonomousVmClustersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudAutonomousVmClustersSortOrderEnum(val string) (ListCloudAutonomousVmClustersSortOrderEnum, bool) {
	enum, ok := mappingListCloudAutonomousVmClustersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
