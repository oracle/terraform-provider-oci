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

// ListOracleDbGcpKeyRingsRequest wrapper for the ListOracleDbGcpKeyRings operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbGcpKeyRings.go.html to see an example of how to use ListOracleDbGcpKeyRingsRequest.
type ListOracleDbGcpKeyRingsRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB GCP Key Ring resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB GCP Key Rings.
	OracleDbGcpKeyRingId *string `mandatory:"false" contributesTo:"query" name:"oracleDbGcpKeyRingId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState OracleDbGcpKeyRingLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Oracle DB GCP Identity Connector resources that match the specified resource OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	OracleDbGcpConnectorId *string `mandatory:"false" contributesTo:"query" name:"oracleDbGcpConnectorId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbGcpKeyRingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbGcpKeyRingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbGcpKeyRingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbGcpKeyRingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbGcpKeyRingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbGcpKeyRingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbGcpKeyRingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbGcpKeyRingLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbGcpKeyRingLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbGcpKeyRingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbGcpKeyRingsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbGcpKeyRingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbGcpKeyRingsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbGcpKeyRingsResponse wrapper for the ListOracleDbGcpKeyRings operation
type ListOracleDbGcpKeyRingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbGcpKeyRingSummaryCollection instances
	OracleDbGcpKeyRingSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbGcpKeyRingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbGcpKeyRingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbGcpKeyRingsSortOrderEnum Enum with underlying type: string
type ListOracleDbGcpKeyRingsSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbGcpKeyRingsSortOrderEnum
const (
	ListOracleDbGcpKeyRingsSortOrderAsc  ListOracleDbGcpKeyRingsSortOrderEnum = "ASC"
	ListOracleDbGcpKeyRingsSortOrderDesc ListOracleDbGcpKeyRingsSortOrderEnum = "DESC"
)

var mappingListOracleDbGcpKeyRingsSortOrderEnum = map[string]ListOracleDbGcpKeyRingsSortOrderEnum{
	"ASC":  ListOracleDbGcpKeyRingsSortOrderAsc,
	"DESC": ListOracleDbGcpKeyRingsSortOrderDesc,
}

var mappingListOracleDbGcpKeyRingsSortOrderEnumLowerCase = map[string]ListOracleDbGcpKeyRingsSortOrderEnum{
	"asc":  ListOracleDbGcpKeyRingsSortOrderAsc,
	"desc": ListOracleDbGcpKeyRingsSortOrderDesc,
}

// GetListOracleDbGcpKeyRingsSortOrderEnumValues Enumerates the set of values for ListOracleDbGcpKeyRingsSortOrderEnum
func GetListOracleDbGcpKeyRingsSortOrderEnumValues() []ListOracleDbGcpKeyRingsSortOrderEnum {
	values := make([]ListOracleDbGcpKeyRingsSortOrderEnum, 0)
	for _, v := range mappingListOracleDbGcpKeyRingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbGcpKeyRingsSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbGcpKeyRingsSortOrderEnum
func GetListOracleDbGcpKeyRingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbGcpKeyRingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbGcpKeyRingsSortOrderEnum(val string) (ListOracleDbGcpKeyRingsSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbGcpKeyRingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbGcpKeyRingsSortByEnum Enum with underlying type: string
type ListOracleDbGcpKeyRingsSortByEnum string

// Set of constants representing the allowable values for ListOracleDbGcpKeyRingsSortByEnum
const (
	ListOracleDbGcpKeyRingsSortByTimecreated ListOracleDbGcpKeyRingsSortByEnum = "timeCreated"
	ListOracleDbGcpKeyRingsSortByDisplayname ListOracleDbGcpKeyRingsSortByEnum = "displayName"
)

var mappingListOracleDbGcpKeyRingsSortByEnum = map[string]ListOracleDbGcpKeyRingsSortByEnum{
	"timeCreated": ListOracleDbGcpKeyRingsSortByTimecreated,
	"displayName": ListOracleDbGcpKeyRingsSortByDisplayname,
}

var mappingListOracleDbGcpKeyRingsSortByEnumLowerCase = map[string]ListOracleDbGcpKeyRingsSortByEnum{
	"timecreated": ListOracleDbGcpKeyRingsSortByTimecreated,
	"displayname": ListOracleDbGcpKeyRingsSortByDisplayname,
}

// GetListOracleDbGcpKeyRingsSortByEnumValues Enumerates the set of values for ListOracleDbGcpKeyRingsSortByEnum
func GetListOracleDbGcpKeyRingsSortByEnumValues() []ListOracleDbGcpKeyRingsSortByEnum {
	values := make([]ListOracleDbGcpKeyRingsSortByEnum, 0)
	for _, v := range mappingListOracleDbGcpKeyRingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbGcpKeyRingsSortByEnumStringValues Enumerates the set of values in String for ListOracleDbGcpKeyRingsSortByEnum
func GetListOracleDbGcpKeyRingsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbGcpKeyRingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbGcpKeyRingsSortByEnum(val string) (ListOracleDbGcpKeyRingsSortByEnum, bool) {
	enum, ok := mappingListOracleDbGcpKeyRingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
