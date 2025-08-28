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

// ListOracleDbAzureBlobMountsRequest wrapper for the ListOracleDbAzureBlobMounts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureBlobMounts.go.html to see an example of how to use ListOracleDbAzureBlobMountsRequest.
type ListOracleDbAzureBlobMountsRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB Azure Blob Mount Resources.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// ID of Oracle DB Azure Blob Mount Resource.
	OracleDbAzureBlobMountId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureBlobMountId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState OracleDbAzureBlobMountLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Oracle DB Azure Blob Mount Resources.
	OracleDbAzureBlobContainerId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureBlobContainerId"`

	// A filter to return Oracle DB Azure Blob Mount Resources.
	OracleDbAzureConnectorId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureConnectorId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAzureBlobMountsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAzureBlobMountsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAzureBlobMountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAzureBlobMountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAzureBlobMountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAzureBlobMountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAzureBlobMountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAzureBlobMountLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAzureBlobMountLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureBlobMountsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAzureBlobMountsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureBlobMountsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAzureBlobMountsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAzureBlobMountsResponse wrapper for the ListOracleDbAzureBlobMounts operation
type ListOracleDbAzureBlobMountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureBlobMountSummaryCollection instances
	OracleDbAzureBlobMountSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAzureBlobMountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAzureBlobMountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAzureBlobMountsSortOrderEnum Enum with underlying type: string
type ListOracleDbAzureBlobMountsSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAzureBlobMountsSortOrderEnum
const (
	ListOracleDbAzureBlobMountsSortOrderAsc  ListOracleDbAzureBlobMountsSortOrderEnum = "ASC"
	ListOracleDbAzureBlobMountsSortOrderDesc ListOracleDbAzureBlobMountsSortOrderEnum = "DESC"
)

var mappingListOracleDbAzureBlobMountsSortOrderEnum = map[string]ListOracleDbAzureBlobMountsSortOrderEnum{
	"ASC":  ListOracleDbAzureBlobMountsSortOrderAsc,
	"DESC": ListOracleDbAzureBlobMountsSortOrderDesc,
}

var mappingListOracleDbAzureBlobMountsSortOrderEnumLowerCase = map[string]ListOracleDbAzureBlobMountsSortOrderEnum{
	"asc":  ListOracleDbAzureBlobMountsSortOrderAsc,
	"desc": ListOracleDbAzureBlobMountsSortOrderDesc,
}

// GetListOracleDbAzureBlobMountsSortOrderEnumValues Enumerates the set of values for ListOracleDbAzureBlobMountsSortOrderEnum
func GetListOracleDbAzureBlobMountsSortOrderEnumValues() []ListOracleDbAzureBlobMountsSortOrderEnum {
	values := make([]ListOracleDbAzureBlobMountsSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAzureBlobMountsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureBlobMountsSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAzureBlobMountsSortOrderEnum
func GetListOracleDbAzureBlobMountsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAzureBlobMountsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureBlobMountsSortOrderEnum(val string) (ListOracleDbAzureBlobMountsSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAzureBlobMountsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAzureBlobMountsSortByEnum Enum with underlying type: string
type ListOracleDbAzureBlobMountsSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAzureBlobMountsSortByEnum
const (
	ListOracleDbAzureBlobMountsSortByTimecreated ListOracleDbAzureBlobMountsSortByEnum = "timeCreated"
	ListOracleDbAzureBlobMountsSortByDisplayname ListOracleDbAzureBlobMountsSortByEnum = "displayName"
)

var mappingListOracleDbAzureBlobMountsSortByEnum = map[string]ListOracleDbAzureBlobMountsSortByEnum{
	"timeCreated": ListOracleDbAzureBlobMountsSortByTimecreated,
	"displayName": ListOracleDbAzureBlobMountsSortByDisplayname,
}

var mappingListOracleDbAzureBlobMountsSortByEnumLowerCase = map[string]ListOracleDbAzureBlobMountsSortByEnum{
	"timecreated": ListOracleDbAzureBlobMountsSortByTimecreated,
	"displayname": ListOracleDbAzureBlobMountsSortByDisplayname,
}

// GetListOracleDbAzureBlobMountsSortByEnumValues Enumerates the set of values for ListOracleDbAzureBlobMountsSortByEnum
func GetListOracleDbAzureBlobMountsSortByEnumValues() []ListOracleDbAzureBlobMountsSortByEnum {
	values := make([]ListOracleDbAzureBlobMountsSortByEnum, 0)
	for _, v := range mappingListOracleDbAzureBlobMountsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureBlobMountsSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAzureBlobMountsSortByEnum
func GetListOracleDbAzureBlobMountsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAzureBlobMountsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureBlobMountsSortByEnum(val string) (ListOracleDbAzureBlobMountsSortByEnum, bool) {
	enum, ok := mappingListOracleDbAzureBlobMountsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
