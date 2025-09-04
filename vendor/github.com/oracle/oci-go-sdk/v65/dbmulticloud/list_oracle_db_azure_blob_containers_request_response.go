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

// ListOracleDbAzureBlobContainersRequest wrapper for the ListOracleDbAzureBlobContainers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureBlobContainers.go.html to see an example of how to use ListOracleDbAzureBlobContainersRequest.
type ListOracleDbAzureBlobContainersRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB Azure Blob Container resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB Azure Blob Container resource.
	OracleDbAzureBlobContainerId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureBlobContainerId"`

	// A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState OracleDbAzureBlobContainerLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Oracle DB Azure Blob Container resources that match the specified Azure Account name.
	AzureStorageAccountName *string `mandatory:"false" contributesTo:"query" name:"azureStorageAccountName"`

	// A filter to return Oracle DB Azure Blob Container resources that match the specified Azure Storage name.
	AzureStorageContainerName *string `mandatory:"false" contributesTo:"query" name:"azureStorageContainerName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAzureBlobContainersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAzureBlobContainersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAzureBlobContainersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAzureBlobContainersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAzureBlobContainersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAzureBlobContainersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAzureBlobContainersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAzureBlobContainerLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAzureBlobContainerLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureBlobContainersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAzureBlobContainersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureBlobContainersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAzureBlobContainersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAzureBlobContainersResponse wrapper for the ListOracleDbAzureBlobContainers operation
type ListOracleDbAzureBlobContainersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureBlobContainerSummaryCollection instances
	OracleDbAzureBlobContainerSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAzureBlobContainersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAzureBlobContainersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAzureBlobContainersSortOrderEnum Enum with underlying type: string
type ListOracleDbAzureBlobContainersSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAzureBlobContainersSortOrderEnum
const (
	ListOracleDbAzureBlobContainersSortOrderAsc  ListOracleDbAzureBlobContainersSortOrderEnum = "ASC"
	ListOracleDbAzureBlobContainersSortOrderDesc ListOracleDbAzureBlobContainersSortOrderEnum = "DESC"
)

var mappingListOracleDbAzureBlobContainersSortOrderEnum = map[string]ListOracleDbAzureBlobContainersSortOrderEnum{
	"ASC":  ListOracleDbAzureBlobContainersSortOrderAsc,
	"DESC": ListOracleDbAzureBlobContainersSortOrderDesc,
}

var mappingListOracleDbAzureBlobContainersSortOrderEnumLowerCase = map[string]ListOracleDbAzureBlobContainersSortOrderEnum{
	"asc":  ListOracleDbAzureBlobContainersSortOrderAsc,
	"desc": ListOracleDbAzureBlobContainersSortOrderDesc,
}

// GetListOracleDbAzureBlobContainersSortOrderEnumValues Enumerates the set of values for ListOracleDbAzureBlobContainersSortOrderEnum
func GetListOracleDbAzureBlobContainersSortOrderEnumValues() []ListOracleDbAzureBlobContainersSortOrderEnum {
	values := make([]ListOracleDbAzureBlobContainersSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAzureBlobContainersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureBlobContainersSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAzureBlobContainersSortOrderEnum
func GetListOracleDbAzureBlobContainersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAzureBlobContainersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureBlobContainersSortOrderEnum(val string) (ListOracleDbAzureBlobContainersSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAzureBlobContainersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAzureBlobContainersSortByEnum Enum with underlying type: string
type ListOracleDbAzureBlobContainersSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAzureBlobContainersSortByEnum
const (
	ListOracleDbAzureBlobContainersSortByTimecreated ListOracleDbAzureBlobContainersSortByEnum = "timeCreated"
	ListOracleDbAzureBlobContainersSortByDisplayname ListOracleDbAzureBlobContainersSortByEnum = "displayName"
)

var mappingListOracleDbAzureBlobContainersSortByEnum = map[string]ListOracleDbAzureBlobContainersSortByEnum{
	"timeCreated": ListOracleDbAzureBlobContainersSortByTimecreated,
	"displayName": ListOracleDbAzureBlobContainersSortByDisplayname,
}

var mappingListOracleDbAzureBlobContainersSortByEnumLowerCase = map[string]ListOracleDbAzureBlobContainersSortByEnum{
	"timecreated": ListOracleDbAzureBlobContainersSortByTimecreated,
	"displayname": ListOracleDbAzureBlobContainersSortByDisplayname,
}

// GetListOracleDbAzureBlobContainersSortByEnumValues Enumerates the set of values for ListOracleDbAzureBlobContainersSortByEnum
func GetListOracleDbAzureBlobContainersSortByEnumValues() []ListOracleDbAzureBlobContainersSortByEnum {
	values := make([]ListOracleDbAzureBlobContainersSortByEnum, 0)
	for _, v := range mappingListOracleDbAzureBlobContainersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureBlobContainersSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAzureBlobContainersSortByEnum
func GetListOracleDbAzureBlobContainersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAzureBlobContainersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureBlobContainersSortByEnum(val string) (ListOracleDbAzureBlobContainersSortByEnum, bool) {
	enum, ok := mappingListOracleDbAzureBlobContainersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
