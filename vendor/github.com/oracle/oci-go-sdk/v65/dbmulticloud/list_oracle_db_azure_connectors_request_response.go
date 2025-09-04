// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOracleDbAzureConnectorsRequest wrapper for the ListOracleDbAzureConnectors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureConnectors.go.html to see an example of how to use ListOracleDbAzureConnectorsRequest.
type ListOracleDbAzureConnectorsRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB Azure Connector resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB Azure Azure Identity Connector resources.
	OracleDbAzureConnectorId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureConnectorId"`

	// A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState OracleDbAzureConnectorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database resource.
	DbClusterResourceId *string `mandatory:"false" contributesTo:"query" name:"dbClusterResourceId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAzureConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAzureConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAzureConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAzureConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAzureConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAzureConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAzureConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAzureConnectorLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAzureConnectorLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAzureConnectorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAzureConnectorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAzureConnectorsResponse wrapper for the ListOracleDbAzureConnectors operation
type ListOracleDbAzureConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureConnectorSummaryCollection instances
	OracleDbAzureConnectorSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAzureConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAzureConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAzureConnectorsSortOrderEnum Enum with underlying type: string
type ListOracleDbAzureConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAzureConnectorsSortOrderEnum
const (
	ListOracleDbAzureConnectorsSortOrderAsc  ListOracleDbAzureConnectorsSortOrderEnum = "ASC"
	ListOracleDbAzureConnectorsSortOrderDesc ListOracleDbAzureConnectorsSortOrderEnum = "DESC"
)

var mappingListOracleDbAzureConnectorsSortOrderEnum = map[string]ListOracleDbAzureConnectorsSortOrderEnum{
	"ASC":  ListOracleDbAzureConnectorsSortOrderAsc,
	"DESC": ListOracleDbAzureConnectorsSortOrderDesc,
}

var mappingListOracleDbAzureConnectorsSortOrderEnumLowerCase = map[string]ListOracleDbAzureConnectorsSortOrderEnum{
	"asc":  ListOracleDbAzureConnectorsSortOrderAsc,
	"desc": ListOracleDbAzureConnectorsSortOrderDesc,
}

// GetListOracleDbAzureConnectorsSortOrderEnumValues Enumerates the set of values for ListOracleDbAzureConnectorsSortOrderEnum
func GetListOracleDbAzureConnectorsSortOrderEnumValues() []ListOracleDbAzureConnectorsSortOrderEnum {
	values := make([]ListOracleDbAzureConnectorsSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAzureConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAzureConnectorsSortOrderEnum
func GetListOracleDbAzureConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAzureConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureConnectorsSortOrderEnum(val string) (ListOracleDbAzureConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAzureConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAzureConnectorsSortByEnum Enum with underlying type: string
type ListOracleDbAzureConnectorsSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAzureConnectorsSortByEnum
const (
	ListOracleDbAzureConnectorsSortByTimecreated ListOracleDbAzureConnectorsSortByEnum = "timeCreated"
	ListOracleDbAzureConnectorsSortByDisplayname ListOracleDbAzureConnectorsSortByEnum = "displayName"
)

var mappingListOracleDbAzureConnectorsSortByEnum = map[string]ListOracleDbAzureConnectorsSortByEnum{
	"timeCreated": ListOracleDbAzureConnectorsSortByTimecreated,
	"displayName": ListOracleDbAzureConnectorsSortByDisplayname,
}

var mappingListOracleDbAzureConnectorsSortByEnumLowerCase = map[string]ListOracleDbAzureConnectorsSortByEnum{
	"timecreated": ListOracleDbAzureConnectorsSortByTimecreated,
	"displayname": ListOracleDbAzureConnectorsSortByDisplayname,
}

// GetListOracleDbAzureConnectorsSortByEnumValues Enumerates the set of values for ListOracleDbAzureConnectorsSortByEnum
func GetListOracleDbAzureConnectorsSortByEnumValues() []ListOracleDbAzureConnectorsSortByEnum {
	values := make([]ListOracleDbAzureConnectorsSortByEnum, 0)
	for _, v := range mappingListOracleDbAzureConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureConnectorsSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAzureConnectorsSortByEnum
func GetListOracleDbAzureConnectorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAzureConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureConnectorsSortByEnum(val string) (ListOracleDbAzureConnectorsSortByEnum, bool) {
	enum, ok := mappingListOracleDbAzureConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
