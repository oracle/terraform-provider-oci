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

// ListOracleDbAzureVaultAssociationsRequest wrapper for the ListOracleDbAzureVaultAssociations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAzureVaultAssociations.go.html to see an example of how to use ListOracleDbAzureVaultAssociationsRequest.
type ListOracleDbAzureVaultAssociationsRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB Azure Vault resources that match the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault resource.
	OracleDbAzureVaultId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureVaultId"`

	// A filter to return Oracle DB Azure Vault Association resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB Azure Vault resources that match the specified OCID](/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Vault Association resource.
	OracleDbAzureVaultAssociationId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureVaultAssociationId"`

	// A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState OracleDbAzureVaultAssociationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Oracle DB Azure Azure Identity Connector resources.
	OracleDbAzureConnectorId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAzureConnectorId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAzureVaultAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAzureVaultAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAzureVaultAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAzureVaultAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAzureVaultAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAzureVaultAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAzureVaultAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAzureVaultAssociationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAzureVaultAssociationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureVaultAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAzureVaultAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAzureVaultAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAzureVaultAssociationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAzureVaultAssociationsResponse wrapper for the ListOracleDbAzureVaultAssociations operation
type ListOracleDbAzureVaultAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureVaultAssociationSummaryCollection instances
	OracleDbAzureVaultAssociationSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAzureVaultAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAzureVaultAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAzureVaultAssociationsSortOrderEnum Enum with underlying type: string
type ListOracleDbAzureVaultAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAzureVaultAssociationsSortOrderEnum
const (
	ListOracleDbAzureVaultAssociationsSortOrderAsc  ListOracleDbAzureVaultAssociationsSortOrderEnum = "ASC"
	ListOracleDbAzureVaultAssociationsSortOrderDesc ListOracleDbAzureVaultAssociationsSortOrderEnum = "DESC"
)

var mappingListOracleDbAzureVaultAssociationsSortOrderEnum = map[string]ListOracleDbAzureVaultAssociationsSortOrderEnum{
	"ASC":  ListOracleDbAzureVaultAssociationsSortOrderAsc,
	"DESC": ListOracleDbAzureVaultAssociationsSortOrderDesc,
}

var mappingListOracleDbAzureVaultAssociationsSortOrderEnumLowerCase = map[string]ListOracleDbAzureVaultAssociationsSortOrderEnum{
	"asc":  ListOracleDbAzureVaultAssociationsSortOrderAsc,
	"desc": ListOracleDbAzureVaultAssociationsSortOrderDesc,
}

// GetListOracleDbAzureVaultAssociationsSortOrderEnumValues Enumerates the set of values for ListOracleDbAzureVaultAssociationsSortOrderEnum
func GetListOracleDbAzureVaultAssociationsSortOrderEnumValues() []ListOracleDbAzureVaultAssociationsSortOrderEnum {
	values := make([]ListOracleDbAzureVaultAssociationsSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAzureVaultAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureVaultAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAzureVaultAssociationsSortOrderEnum
func GetListOracleDbAzureVaultAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAzureVaultAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureVaultAssociationsSortOrderEnum(val string) (ListOracleDbAzureVaultAssociationsSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAzureVaultAssociationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAzureVaultAssociationsSortByEnum Enum with underlying type: string
type ListOracleDbAzureVaultAssociationsSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAzureVaultAssociationsSortByEnum
const (
	ListOracleDbAzureVaultAssociationsSortByTimecreated ListOracleDbAzureVaultAssociationsSortByEnum = "timeCreated"
	ListOracleDbAzureVaultAssociationsSortByDisplayname ListOracleDbAzureVaultAssociationsSortByEnum = "displayName"
)

var mappingListOracleDbAzureVaultAssociationsSortByEnum = map[string]ListOracleDbAzureVaultAssociationsSortByEnum{
	"timeCreated": ListOracleDbAzureVaultAssociationsSortByTimecreated,
	"displayName": ListOracleDbAzureVaultAssociationsSortByDisplayname,
}

var mappingListOracleDbAzureVaultAssociationsSortByEnumLowerCase = map[string]ListOracleDbAzureVaultAssociationsSortByEnum{
	"timecreated": ListOracleDbAzureVaultAssociationsSortByTimecreated,
	"displayname": ListOracleDbAzureVaultAssociationsSortByDisplayname,
}

// GetListOracleDbAzureVaultAssociationsSortByEnumValues Enumerates the set of values for ListOracleDbAzureVaultAssociationsSortByEnum
func GetListOracleDbAzureVaultAssociationsSortByEnumValues() []ListOracleDbAzureVaultAssociationsSortByEnum {
	values := make([]ListOracleDbAzureVaultAssociationsSortByEnum, 0)
	for _, v := range mappingListOracleDbAzureVaultAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAzureVaultAssociationsSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAzureVaultAssociationsSortByEnum
func GetListOracleDbAzureVaultAssociationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAzureVaultAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAzureVaultAssociationsSortByEnum(val string) (ListOracleDbAzureVaultAssociationsSortByEnum, bool) {
	enum, ok := mappingListOracleDbAzureVaultAssociationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
