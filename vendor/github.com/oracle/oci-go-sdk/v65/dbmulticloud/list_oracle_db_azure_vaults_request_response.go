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

// ListOracleDbAzureVaultsRequest wrapper for the ListOracleDbAzureVaults operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureVaults.go.html to see an example of how to use ListOracleDbAzureVaultsRequest.
type ListOracleDbAzureVaultsRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Azure Vaults.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB Azure Vault Resources.
	OracleDbAzureVaultId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureVaultId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState OracleDbAzureVaultLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Azure Vaults.
	OracleDbAzureResourceGroup *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureResourceGroup"`

	// A filter to return Oracle DB Azure Blob Mount Resources.
	OracleDbAzureConnectorId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureConnectorId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAzureVaultsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAzureVaultsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAzureVaultsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAzureVaultsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAzureVaultsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAzureVaultsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAzureVaultsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAzureVaultLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAzureVaultLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureVaultsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAzureVaultsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureVaultsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAzureVaultsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAzureVaultsResponse wrapper for the ListOracleDbAzureVaults operation
type ListOracleDbAzureVaultsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureVaultSummaryCollection instances
	OracleDbAzureVaultSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAzureVaultsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAzureVaultsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAzureVaultsSortOrderEnum Enum with underlying type: string
type ListOracleDbAzureVaultsSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAzureVaultsSortOrderEnum
const (
	ListOracleDbAzureVaultsSortOrderAsc  ListOracleDbAzureVaultsSortOrderEnum = "ASC"
	ListOracleDbAzureVaultsSortOrderDesc ListOracleDbAzureVaultsSortOrderEnum = "DESC"
)

var mappingListOracleDbAzureVaultsSortOrderEnum = map[string]ListOracleDbAzureVaultsSortOrderEnum{
	"ASC":  ListOracleDbAzureVaultsSortOrderAsc,
	"DESC": ListOracleDbAzureVaultsSortOrderDesc,
}

var mappingListOracleDbAzureVaultsSortOrderEnumLowerCase = map[string]ListOracleDbAzureVaultsSortOrderEnum{
	"asc":  ListOracleDbAzureVaultsSortOrderAsc,
	"desc": ListOracleDbAzureVaultsSortOrderDesc,
}

// GetListOracleDbAzureVaultsSortOrderEnumValues Enumerates the set of values for ListOracleDbAzureVaultsSortOrderEnum
func GetListOracleDbAzureVaultsSortOrderEnumValues() []ListOracleDbAzureVaultsSortOrderEnum {
	values := make([]ListOracleDbAzureVaultsSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAzureVaultsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureVaultsSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAzureVaultsSortOrderEnum
func GetListOracleDbAzureVaultsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAzureVaultsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureVaultsSortOrderEnum(val string) (ListOracleDbAzureVaultsSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAzureVaultsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAzureVaultsSortByEnum Enum with underlying type: string
type ListOracleDbAzureVaultsSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAzureVaultsSortByEnum
const (
	ListOracleDbAzureVaultsSortByTimecreated ListOracleDbAzureVaultsSortByEnum = "timeCreated"
	ListOracleDbAzureVaultsSortByDisplayname ListOracleDbAzureVaultsSortByEnum = "displayName"
)

var mappingListOracleDbAzureVaultsSortByEnum = map[string]ListOracleDbAzureVaultsSortByEnum{
	"timeCreated": ListOracleDbAzureVaultsSortByTimecreated,
	"displayName": ListOracleDbAzureVaultsSortByDisplayname,
}

var mappingListOracleDbAzureVaultsSortByEnumLowerCase = map[string]ListOracleDbAzureVaultsSortByEnum{
	"timecreated": ListOracleDbAzureVaultsSortByTimecreated,
	"displayname": ListOracleDbAzureVaultsSortByDisplayname,
}

// GetListOracleDbAzureVaultsSortByEnumValues Enumerates the set of values for ListOracleDbAzureVaultsSortByEnum
func GetListOracleDbAzureVaultsSortByEnumValues() []ListOracleDbAzureVaultsSortByEnum {
	values := make([]ListOracleDbAzureVaultsSortByEnum, 0)
	for _, v := range mappingListOracleDbAzureVaultsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureVaultsSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAzureVaultsSortByEnum
func GetListOracleDbAzureVaultsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAzureVaultsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureVaultsSortByEnum(val string) (ListOracleDbAzureVaultsSortByEnum, bool) {
	enum, ok := mappingListOracleDbAzureVaultsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
