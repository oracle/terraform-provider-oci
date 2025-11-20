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

// ListOracleDbAwsKeysRequest wrapper for the ListOracleDbAwsKeys operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/ListOracleDbAwsKeys.go.html to see an example of how to use ListOracleDbAwsKeysRequest.
type ListOracleDbAwsKeysRequest struct {

	// The ID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return Oracle DB AWS Key Resource that match the given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return Oracle DB AWS Identity Connector Resource that match the given OCID](/Content/General/Concepts/identifiers.htm) of the Oracle DB AWS Key resource.
	OracleDbAwsKeyId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAwsKeyId"`

	// A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState OracleDbAwsKeyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Oracle DB AWS Identity Connector resources that match the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB AWS Identity Connector resource.
	OracleDbAwsConnectorId *string `mandatory:"false" contributesTo:"query" name:"oracleDbAwsConnectorId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOracleDbAwsKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, default is timeCreated.
	SortBy ListOracleDbAwsKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOracleDbAwsKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOracleDbAwsKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOracleDbAwsKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOracleDbAwsKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOracleDbAwsKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOracleDbAwsKeyLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOracleDbAwsKeyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAwsKeysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOracleDbAwsKeysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOracleDbAwsKeysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOracleDbAwsKeysSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOracleDbAwsKeysResponse wrapper for the ListOracleDbAwsKeys operation
type ListOracleDbAwsKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAwsKeySummaryCollection instances
	OracleDbAwsKeySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOracleDbAwsKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOracleDbAwsKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOracleDbAwsKeysSortOrderEnum Enum with underlying type: string
type ListOracleDbAwsKeysSortOrderEnum string

// Set of constants representing the allowable values for ListOracleDbAwsKeysSortOrderEnum
const (
	ListOracleDbAwsKeysSortOrderAsc  ListOracleDbAwsKeysSortOrderEnum = "ASC"
	ListOracleDbAwsKeysSortOrderDesc ListOracleDbAwsKeysSortOrderEnum = "DESC"
)

var mappingListOracleDbAwsKeysSortOrderEnum = map[string]ListOracleDbAwsKeysSortOrderEnum{
	"ASC":  ListOracleDbAwsKeysSortOrderAsc,
	"DESC": ListOracleDbAwsKeysSortOrderDesc,
}

var mappingListOracleDbAwsKeysSortOrderEnumLowerCase = map[string]ListOracleDbAwsKeysSortOrderEnum{
	"asc":  ListOracleDbAwsKeysSortOrderAsc,
	"desc": ListOracleDbAwsKeysSortOrderDesc,
}

// GetListOracleDbAwsKeysSortOrderEnumValues Enumerates the set of values for ListOracleDbAwsKeysSortOrderEnum
func GetListOracleDbAwsKeysSortOrderEnumValues() []ListOracleDbAwsKeysSortOrderEnum {
	values := make([]ListOracleDbAwsKeysSortOrderEnum, 0)
	for _, v := range mappingListOracleDbAwsKeysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAwsKeysSortOrderEnumStringValues Enumerates the set of values in String for ListOracleDbAwsKeysSortOrderEnum
func GetListOracleDbAwsKeysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOracleDbAwsKeysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAwsKeysSortOrderEnum(val string) (ListOracleDbAwsKeysSortOrderEnum, bool) {
	enum, ok := mappingListOracleDbAwsKeysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOracleDbAwsKeysSortByEnum Enum with underlying type: string
type ListOracleDbAwsKeysSortByEnum string

// Set of constants representing the allowable values for ListOracleDbAwsKeysSortByEnum
const (
	ListOracleDbAwsKeysSortByTimecreated ListOracleDbAwsKeysSortByEnum = "timeCreated"
	ListOracleDbAwsKeysSortByDisplayname ListOracleDbAwsKeysSortByEnum = "displayName"
)

var mappingListOracleDbAwsKeysSortByEnum = map[string]ListOracleDbAwsKeysSortByEnum{
	"timeCreated": ListOracleDbAwsKeysSortByTimecreated,
	"displayName": ListOracleDbAwsKeysSortByDisplayname,
}

var mappingListOracleDbAwsKeysSortByEnumLowerCase = map[string]ListOracleDbAwsKeysSortByEnum{
	"timecreated": ListOracleDbAwsKeysSortByTimecreated,
	"displayname": ListOracleDbAwsKeysSortByDisplayname,
}

// GetListOracleDbAwsKeysSortByEnumValues Enumerates the set of values for ListOracleDbAwsKeysSortByEnum
func GetListOracleDbAwsKeysSortByEnumValues() []ListOracleDbAwsKeysSortByEnum {
	values := make([]ListOracleDbAwsKeysSortByEnum, 0)
	for _, v := range mappingListOracleDbAwsKeysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOracleDbAwsKeysSortByEnumStringValues Enumerates the set of values in String for ListOracleDbAwsKeysSortByEnum
func GetListOracleDbAwsKeysSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOracleDbAwsKeysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOracleDbAwsKeysSortByEnum(val string) (ListOracleDbAwsKeysSortByEnum, bool) {
	enum, ok := mappingListOracleDbAwsKeysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
