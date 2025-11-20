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

// ListOracleDbGcpIdentityConnectorsRequest wrapper for the ListOracleDbGcpIdentityConnectors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbGcpIdentityConnectors.go.html to see an example of how to use ListOracleDbGcpIdentityConnectorsRequest.
type ListOracleDbGcpIdentityConnectorsRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB GCP Identity Connector resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB Identity Connector resource that match the given resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only resources that match the specified lifecycle state. The state value is case-insensitive.
	LifecycleState OracleDbGcpIdentityConnectorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbGcpIdentityConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbGcpIdentityConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbGcpIdentityConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbGcpIdentityConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbGcpIdentityConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbGcpIdentityConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbGcpIdentityConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbGcpIdentityConnectorLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbGcpIdentityConnectorLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbGcpIdentityConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbGcpIdentityConnectorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbGcpIdentityConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbGcpIdentityConnectorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbGcpIdentityConnectorsResponse wrapper for the ListOracleDbGcpIdentityConnectors operation
type ListOracleDbGcpIdentityConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbGcpIdentityConnectorSummaryCollection instances
	OracleDbGcpIdentityConnectorSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbGcpIdentityConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbGcpIdentityConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbGcpIdentityConnectorsSortOrderEnum Enum with underlying type: string
type ListOracleDbGcpIdentityConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbGcpIdentityConnectorsSortOrderEnum
const (
	ListOracleDbGcpIdentityConnectorsSortOrderAsc  ListOracleDbGcpIdentityConnectorsSortOrderEnum = "ASC"
	ListOracleDbGcpIdentityConnectorsSortOrderDesc ListOracleDbGcpIdentityConnectorsSortOrderEnum = "DESC"
)

var mappingListOracleDbGcpIdentityConnectorsSortOrderEnum = map[string]ListOracleDbGcpIdentityConnectorsSortOrderEnum{
	"ASC":  ListOracleDbGcpIdentityConnectorsSortOrderAsc,
	"DESC": ListOracleDbGcpIdentityConnectorsSortOrderDesc,
}

var mappingListOracleDbGcpIdentityConnectorsSortOrderEnumLowerCase = map[string]ListOracleDbGcpIdentityConnectorsSortOrderEnum{
	"asc":  ListOracleDbGcpIdentityConnectorsSortOrderAsc,
	"desc": ListOracleDbGcpIdentityConnectorsSortOrderDesc,
}

// GetListOracleDbGcpIdentityConnectorsSortOrderEnumValues Enumerates the set of values for ListOracleDbGcpIdentityConnectorsSortOrderEnum
func GetListOracleDbGcpIdentityConnectorsSortOrderEnumValues() []ListOracleDbGcpIdentityConnectorsSortOrderEnum {
	values := make([]ListOracleDbGcpIdentityConnectorsSortOrderEnum, 0)
	for _, v := range mappingListOracleDbGcpIdentityConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbGcpIdentityConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbGcpIdentityConnectorsSortOrderEnum
func GetListOracleDbGcpIdentityConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbGcpIdentityConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbGcpIdentityConnectorsSortOrderEnum(val string) (ListOracleDbGcpIdentityConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbGcpIdentityConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbGcpIdentityConnectorsSortByEnum Enum with underlying type: string
type ListOracleDbGcpIdentityConnectorsSortByEnum string

// Set of constants representing the allowable values for ListOracleDbGcpIdentityConnectorsSortByEnum
const (
	ListOracleDbGcpIdentityConnectorsSortByTimecreated ListOracleDbGcpIdentityConnectorsSortByEnum = "timeCreated"
	ListOracleDbGcpIdentityConnectorsSortByDisplayname ListOracleDbGcpIdentityConnectorsSortByEnum = "displayName"
)

var mappingListOracleDbGcpIdentityConnectorsSortByEnum = map[string]ListOracleDbGcpIdentityConnectorsSortByEnum{
	"timeCreated": ListOracleDbGcpIdentityConnectorsSortByTimecreated,
	"displayName": ListOracleDbGcpIdentityConnectorsSortByDisplayname,
}

var mappingListOracleDbGcpIdentityConnectorsSortByEnumLowerCase = map[string]ListOracleDbGcpIdentityConnectorsSortByEnum{
	"timecreated": ListOracleDbGcpIdentityConnectorsSortByTimecreated,
	"displayname": ListOracleDbGcpIdentityConnectorsSortByDisplayname,
}

// GetListOracleDbGcpIdentityConnectorsSortByEnumValues Enumerates the set of values for ListOracleDbGcpIdentityConnectorsSortByEnum
func GetListOracleDbGcpIdentityConnectorsSortByEnumValues() []ListOracleDbGcpIdentityConnectorsSortByEnum {
	values := make([]ListOracleDbGcpIdentityConnectorsSortByEnum, 0)
	for _, v := range mappingListOracleDbGcpIdentityConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbGcpIdentityConnectorsSortByEnumStringValues Enumerates the set of values in String for ListOracleDbGcpIdentityConnectorsSortByEnum
func GetListOracleDbGcpIdentityConnectorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbGcpIdentityConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbGcpIdentityConnectorsSortByEnum(val string) (ListOracleDbGcpIdentityConnectorsSortByEnum, bool) {
	enum, ok := mappingListOracleDbGcpIdentityConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
