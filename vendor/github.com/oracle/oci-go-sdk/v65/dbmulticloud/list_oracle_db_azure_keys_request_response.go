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

// ListOracleDbAzureKeysRequest wrapper for the ListOracleDbAzureKeys operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureKeys.go.html to see an example of how to use ListOracleDbAzureKeysRequest.
type ListOracleDbAzureKeysRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Azure Vault Keys.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB Azure Vault Resources.
	OracleDbAzureVaultId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureVaultId"`

	// A filter to return Oracle DB Azure Vault Key Resources.
	OracleDbAzureKeyId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureKeyId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState OracleDbAzureKeyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAzureKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAzureKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAzureKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAzureKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAzureKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAzureKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAzureKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAzureKeyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAzureKeyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureKeysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAzureKeysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureKeysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAzureKeysSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAzureKeysResponse wrapper for the ListOracleDbAzureKeys operation
type ListOracleDbAzureKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureKeySummaryCollection instances
	OracleDbAzureKeySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAzureKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAzureKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAzureKeysSortOrderEnum Enum with underlying type: string
type ListOracleDbAzureKeysSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAzureKeysSortOrderEnum
const (
	ListOracleDbAzureKeysSortOrderAsc  ListOracleDbAzureKeysSortOrderEnum = "ASC"
	ListOracleDbAzureKeysSortOrderDesc ListOracleDbAzureKeysSortOrderEnum = "DESC"
)

var mappingListOracleDbAzureKeysSortOrderEnum = map[string]ListOracleDbAzureKeysSortOrderEnum{
	"ASC":  ListOracleDbAzureKeysSortOrderAsc,
	"DESC": ListOracleDbAzureKeysSortOrderDesc,
}

var mappingListOracleDbAzureKeysSortOrderEnumLowerCase = map[string]ListOracleDbAzureKeysSortOrderEnum{
	"asc":  ListOracleDbAzureKeysSortOrderAsc,
	"desc": ListOracleDbAzureKeysSortOrderDesc,
}

// GetListOracleDbAzureKeysSortOrderEnumValues Enumerates the set of values for ListOracleDbAzureKeysSortOrderEnum
func GetListOracleDbAzureKeysSortOrderEnumValues() []ListOracleDbAzureKeysSortOrderEnum {
	values := make([]ListOracleDbAzureKeysSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAzureKeysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureKeysSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAzureKeysSortOrderEnum
func GetListOracleDbAzureKeysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAzureKeysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureKeysSortOrderEnum(val string) (ListOracleDbAzureKeysSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAzureKeysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAzureKeysSortByEnum Enum with underlying type: string
type ListOracleDbAzureKeysSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAzureKeysSortByEnum
const (
	ListOracleDbAzureKeysSortByTimecreated ListOracleDbAzureKeysSortByEnum = "timeCreated"
	ListOracleDbAzureKeysSortByDisplayname ListOracleDbAzureKeysSortByEnum = "displayName"
)

var mappingListOracleDbAzureKeysSortByEnum = map[string]ListOracleDbAzureKeysSortByEnum{
	"timeCreated": ListOracleDbAzureKeysSortByTimecreated,
	"displayName": ListOracleDbAzureKeysSortByDisplayname,
}

var mappingListOracleDbAzureKeysSortByEnumLowerCase = map[string]ListOracleDbAzureKeysSortByEnum{
	"timecreated": ListOracleDbAzureKeysSortByTimecreated,
	"displayname": ListOracleDbAzureKeysSortByDisplayname,
}

// GetListOracleDbAzureKeysSortByEnumValues Enumerates the set of values for ListOracleDbAzureKeysSortByEnum
func GetListOracleDbAzureKeysSortByEnumValues() []ListOracleDbAzureKeysSortByEnum {
	values := make([]ListOracleDbAzureKeysSortByEnum, 0)
	for _, v := range mappingListOracleDbAzureKeysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureKeysSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAzureKeysSortByEnum
func GetListOracleDbAzureKeysSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAzureKeysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureKeysSortByEnum(val string) (ListOracleDbAzureKeysSortByEnum, bool) {
	enum, ok := mappingListOracleDbAzureKeysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
