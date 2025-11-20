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

// ListOracleDbAwsIdentityConnectorsRequest wrapper for the ListOracleDbAwsIdentityConnectors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAwsIdentityConnectors.go.html to see an example of how to use ListOracleDbAwsIdentityConnectorsRequest.
type ListOracleDbAwsIdentityConnectorsRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB AWS Identity Connector Resource that match the given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState OracleDbAwsIdentityConnectorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Oracle DB Identity Connector resource that match the given resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAwsIdentityConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAwsIdentityConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAwsIdentityConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAwsIdentityConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAwsIdentityConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAwsIdentityConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAwsIdentityConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAwsIdentityConnectorLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAwsIdentityConnectorLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAwsIdentityConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAwsIdentityConnectorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAwsIdentityConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAwsIdentityConnectorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAwsIdentityConnectorsResponse wrapper for the ListOracleDbAwsIdentityConnectors operation
type ListOracleDbAwsIdentityConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAwsIdentityConnectorSummaryCollection instances
	OracleDbAwsIdentityConnectorSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAwsIdentityConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAwsIdentityConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAwsIdentityConnectorsSortOrderEnum Enum with underlying type: string
type ListOracleDbAwsIdentityConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAwsIdentityConnectorsSortOrderEnum
const (
	ListOracleDbAwsIdentityConnectorsSortOrderAsc  ListOracleDbAwsIdentityConnectorsSortOrderEnum = "ASC"
	ListOracleDbAwsIdentityConnectorsSortOrderDesc ListOracleDbAwsIdentityConnectorsSortOrderEnum = "DESC"
)

var mappingListOracleDbAwsIdentityConnectorsSortOrderEnum = map[string]ListOracleDbAwsIdentityConnectorsSortOrderEnum{
	"ASC":  ListOracleDbAwsIdentityConnectorsSortOrderAsc,
	"DESC": ListOracleDbAwsIdentityConnectorsSortOrderDesc,
}

var mappingListOracleDbAwsIdentityConnectorsSortOrderEnumLowerCase = map[string]ListOracleDbAwsIdentityConnectorsSortOrderEnum{
	"asc":  ListOracleDbAwsIdentityConnectorsSortOrderAsc,
	"desc": ListOracleDbAwsIdentityConnectorsSortOrderDesc,
}

// GetListOracleDbAwsIdentityConnectorsSortOrderEnumValues Enumerates the set of values for ListOracleDbAwsIdentityConnectorsSortOrderEnum
func GetListOracleDbAwsIdentityConnectorsSortOrderEnumValues() []ListOracleDbAwsIdentityConnectorsSortOrderEnum {
	values := make([]ListOracleDbAwsIdentityConnectorsSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAwsIdentityConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAwsIdentityConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAwsIdentityConnectorsSortOrderEnum
func GetListOracleDbAwsIdentityConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAwsIdentityConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAwsIdentityConnectorsSortOrderEnum(val string) (ListOracleDbAwsIdentityConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAwsIdentityConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAwsIdentityConnectorsSortByEnum Enum with underlying type: string
type ListOracleDbAwsIdentityConnectorsSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAwsIdentityConnectorsSortByEnum
const (
	ListOracleDbAwsIdentityConnectorsSortByTimecreated ListOracleDbAwsIdentityConnectorsSortByEnum = "timeCreated"
	ListOracleDbAwsIdentityConnectorsSortByDisplayname ListOracleDbAwsIdentityConnectorsSortByEnum = "displayName"
)

var mappingListOracleDbAwsIdentityConnectorsSortByEnum = map[string]ListOracleDbAwsIdentityConnectorsSortByEnum{
	"timeCreated": ListOracleDbAwsIdentityConnectorsSortByTimecreated,
	"displayName": ListOracleDbAwsIdentityConnectorsSortByDisplayname,
}

var mappingListOracleDbAwsIdentityConnectorsSortByEnumLowerCase = map[string]ListOracleDbAwsIdentityConnectorsSortByEnum{
	"timecreated": ListOracleDbAwsIdentityConnectorsSortByTimecreated,
	"displayname": ListOracleDbAwsIdentityConnectorsSortByDisplayname,
}

// GetListOracleDbAwsIdentityConnectorsSortByEnumValues Enumerates the set of values for ListOracleDbAwsIdentityConnectorsSortByEnum
func GetListOracleDbAwsIdentityConnectorsSortByEnumValues() []ListOracleDbAwsIdentityConnectorsSortByEnum {
	values := make([]ListOracleDbAwsIdentityConnectorsSortByEnum, 0)
	for _, v := range mappingListOracleDbAwsIdentityConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAwsIdentityConnectorsSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAwsIdentityConnectorsSortByEnum
func GetListOracleDbAwsIdentityConnectorsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAwsIdentityConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAwsIdentityConnectorsSortByEnum(val string) (ListOracleDbAwsIdentityConnectorsSortByEnum, bool) {
	enum, ok := mappingListOracleDbAwsIdentityConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
