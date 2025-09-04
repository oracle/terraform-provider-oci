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

// ListOracleDbGcpKeysRequest wrapper for the ListOracleDbGcpKeys operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbGcpKeys.go.html to see an example of how to use ListOracleDbGcpKeysRequest.
type ListOracleDbGcpKeysRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB Google Cloud Key resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB GCP Key Rings.
	OracleDbGcpKeyRingId *string `mandatory:"false" contributesTo:"query" name:"oracleDbGcpKeyRingId"`

	// A filter to return Oracle DB Google Cloud Key resources.
	OracleDbGcpKeyId *string `mandatory:"false" contributesTo:"query" name:"oracleDbGcpKeyId"`

	// A filter to return only resources that match the specified lifecycle state. The state value is case-insensitive.
	LifecycleState OracleDbGcpKeyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbGcpKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbGcpKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbGcpKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbGcpKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbGcpKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbGcpKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbGcpKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbGcpKeyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbGcpKeyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbGcpKeysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbGcpKeysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbGcpKeysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbGcpKeysSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbGcpKeysResponse wrapper for the ListOracleDbGcpKeys operation
type ListOracleDbGcpKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbGcpKeySummaryCollection instances
	OracleDbGcpKeySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbGcpKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbGcpKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbGcpKeysSortOrderEnum Enum with underlying type: string
type ListOracleDbGcpKeysSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbGcpKeysSortOrderEnum
const (
	ListOracleDbGcpKeysSortOrderAsc  ListOracleDbGcpKeysSortOrderEnum = "ASC"
	ListOracleDbGcpKeysSortOrderDesc ListOracleDbGcpKeysSortOrderEnum = "DESC"
)

var mappingListOracleDbGcpKeysSortOrderEnum = map[string]ListOracleDbGcpKeysSortOrderEnum{
	"ASC":  ListOracleDbGcpKeysSortOrderAsc,
	"DESC": ListOracleDbGcpKeysSortOrderDesc,
}

var mappingListOracleDbGcpKeysSortOrderEnumLowerCase = map[string]ListOracleDbGcpKeysSortOrderEnum{
	"asc":  ListOracleDbGcpKeysSortOrderAsc,
	"desc": ListOracleDbGcpKeysSortOrderDesc,
}

// GetListOracleDbGcpKeysSortOrderEnumValues Enumerates the set of values for ListOracleDbGcpKeysSortOrderEnum
func GetListOracleDbGcpKeysSortOrderEnumValues() []ListOracleDbGcpKeysSortOrderEnum {
	values := make([]ListOracleDbGcpKeysSortOrderEnum, 0)
	for _, v := range mappingListOracleDbGcpKeysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbGcpKeysSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbGcpKeysSortOrderEnum
func GetListOracleDbGcpKeysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbGcpKeysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbGcpKeysSortOrderEnum(val string) (ListOracleDbGcpKeysSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbGcpKeysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbGcpKeysSortByEnum Enum with underlying type: string
type ListOracleDbGcpKeysSortByEnum string

// Set of constants representing the allowable values for ListOracleDbGcpKeysSortByEnum
const (
	ListOracleDbGcpKeysSortByTimecreated ListOracleDbGcpKeysSortByEnum = "timeCreated"
	ListOracleDbGcpKeysSortByDisplayname ListOracleDbGcpKeysSortByEnum = "displayName"
)

var mappingListOracleDbGcpKeysSortByEnum = map[string]ListOracleDbGcpKeysSortByEnum{
	"timeCreated": ListOracleDbGcpKeysSortByTimecreated,
	"displayName": ListOracleDbGcpKeysSortByDisplayname,
}

var mappingListOracleDbGcpKeysSortByEnumLowerCase = map[string]ListOracleDbGcpKeysSortByEnum{
	"timecreated": ListOracleDbGcpKeysSortByTimecreated,
	"displayname": ListOracleDbGcpKeysSortByDisplayname,
}

// GetListOracleDbGcpKeysSortByEnumValues Enumerates the set of values for ListOracleDbGcpKeysSortByEnum
func GetListOracleDbGcpKeysSortByEnumValues() []ListOracleDbGcpKeysSortByEnum {
	values := make([]ListOracleDbGcpKeysSortByEnum, 0)
	for _, v := range mappingListOracleDbGcpKeysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbGcpKeysSortByEnumStringValues Enumerates the set of values in String for ListOracleDbGcpKeysSortByEnum
func GetListOracleDbGcpKeysSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbGcpKeysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbGcpKeysSortByEnum(val string) (ListOracleDbGcpKeysSortByEnum, bool) {
	enum, ok := mappingListOracleDbGcpKeysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
