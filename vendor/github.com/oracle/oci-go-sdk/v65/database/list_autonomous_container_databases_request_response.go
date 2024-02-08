// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousContainerDatabasesRequest wrapper for the ListAutonomousContainerDatabases operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousContainerDatabases.go.html to see an example of how to use ListAutonomousContainerDatabasesRequest.
type ListAutonomousContainerDatabasesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The Autonomous Exadata Infrastructure OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousExadataInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"autonomousExadataInfrastructureId"`

	// The Autonomous VM Cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousVmClusterId *string `mandatory:"false" contributesTo:"query" name:"autonomousVmClusterId"`

	// A filter to return only resources that match the given Infrastructure Type.
	InfrastructureType AutonomousContainerDatabaseSummaryInfrastructureTypeEnum `mandatory:"false" contributesTo:"query" name:"infrastructureType" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListAutonomousContainerDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousContainerDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState AutonomousContainerDatabaseSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given availability domain exactly.
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given service-level agreement type exactly.
	ServiceLevelAgreementType *string `mandatory:"false" contributesTo:"query" name:"serviceLevelAgreementType"`

	// The cloud Autonomous VM Cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CloudAutonomousVmClusterId *string `mandatory:"false" contributesTo:"query" name:"cloudAutonomousVmClusterId"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousContainerDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousContainerDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousContainerDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousContainerDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousContainerDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousContainerDatabaseSummaryInfrastructureTypeEnum(string(request.InfrastructureType)); !ok && request.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", request.InfrastructureType, strings.Join(GetAutonomousContainerDatabaseSummaryInfrastructureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousContainerDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutonomousContainerDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousContainerDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutonomousContainerDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAutonomousContainerDatabaseSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousContainerDatabasesResponse wrapper for the ListAutonomousContainerDatabases operation
type ListAutonomousContainerDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousContainerDatabaseSummary instances
	Items []AutonomousContainerDatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you must contact Oracle about
	// a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousContainerDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousContainerDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousContainerDatabasesSortByEnum Enum with underlying type: string
type ListAutonomousContainerDatabasesSortByEnum string

// Set of constants representing the allowable values for ListAutonomousContainerDatabasesSortByEnum
const (
	ListAutonomousContainerDatabasesSortByTimecreated ListAutonomousContainerDatabasesSortByEnum = "TIMECREATED"
	ListAutonomousContainerDatabasesSortByDisplayname ListAutonomousContainerDatabasesSortByEnum = "DISPLAYNAME"
)

var mappingListAutonomousContainerDatabasesSortByEnum = map[string]ListAutonomousContainerDatabasesSortByEnum{
	"TIMECREATED": ListAutonomousContainerDatabasesSortByTimecreated,
	"DISPLAYNAME": ListAutonomousContainerDatabasesSortByDisplayname,
}

var mappingListAutonomousContainerDatabasesSortByEnumLowerCase = map[string]ListAutonomousContainerDatabasesSortByEnum{
	"timecreated": ListAutonomousContainerDatabasesSortByTimecreated,
	"displayname": ListAutonomousContainerDatabasesSortByDisplayname,
}

// GetListAutonomousContainerDatabasesSortByEnumValues Enumerates the set of values for ListAutonomousContainerDatabasesSortByEnum
func GetListAutonomousContainerDatabasesSortByEnumValues() []ListAutonomousContainerDatabasesSortByEnum {
	values := make([]ListAutonomousContainerDatabasesSortByEnum, 0)
	for _, v := range mappingListAutonomousContainerDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousContainerDatabasesSortByEnumStringValues Enumerates the set of values in String for ListAutonomousContainerDatabasesSortByEnum
func GetListAutonomousContainerDatabasesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAutonomousContainerDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousContainerDatabasesSortByEnum(val string) (ListAutonomousContainerDatabasesSortByEnum, bool) {
	enum, ok := mappingListAutonomousContainerDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutonomousContainerDatabasesSortOrderEnum Enum with underlying type: string
type ListAutonomousContainerDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousContainerDatabasesSortOrderEnum
const (
	ListAutonomousContainerDatabasesSortOrderAsc  ListAutonomousContainerDatabasesSortOrderEnum = "ASC"
	ListAutonomousContainerDatabasesSortOrderDesc ListAutonomousContainerDatabasesSortOrderEnum = "DESC"
)

var mappingListAutonomousContainerDatabasesSortOrderEnum = map[string]ListAutonomousContainerDatabasesSortOrderEnum{
	"ASC":  ListAutonomousContainerDatabasesSortOrderAsc,
	"DESC": ListAutonomousContainerDatabasesSortOrderDesc,
}

var mappingListAutonomousContainerDatabasesSortOrderEnumLowerCase = map[string]ListAutonomousContainerDatabasesSortOrderEnum{
	"asc":  ListAutonomousContainerDatabasesSortOrderAsc,
	"desc": ListAutonomousContainerDatabasesSortOrderDesc,
}

// GetListAutonomousContainerDatabasesSortOrderEnumValues Enumerates the set of values for ListAutonomousContainerDatabasesSortOrderEnum
func GetListAutonomousContainerDatabasesSortOrderEnumValues() []ListAutonomousContainerDatabasesSortOrderEnum {
	values := make([]ListAutonomousContainerDatabasesSortOrderEnum, 0)
	for _, v := range mappingListAutonomousContainerDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousContainerDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListAutonomousContainerDatabasesSortOrderEnum
func GetListAutonomousContainerDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutonomousContainerDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousContainerDatabasesSortOrderEnum(val string) (ListAutonomousContainerDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListAutonomousContainerDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
