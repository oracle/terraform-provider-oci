// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNamedCredentialsRequest wrapper for the ListNamedCredentials operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementagent/ListNamedCredentials.go.html to see an example of how to use ListNamedCredentialsRequest.
type ListNamedCredentialsRequest struct {

	// The ManagementAgentID of the agent from which the named credentials are associated.
	ManagementAgentId *string `mandatory:"true" contributesTo:"query" name:"managementAgentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListNamedCredentialsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. If no value is specified timeCreated is default.
	SortBy ListNamedCredentialsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState []NamedCredentialLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// Filter list for these name items.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// Filter list for these type values.
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// Filter list for these Named credentials identifiers (ocid) values.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNamedCredentialsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNamedCredentialsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNamedCredentialsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNamedCredentialsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNamedCredentialsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNamedCredentialsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNamedCredentialsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNamedCredentialsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNamedCredentialsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingNamedCredentialLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetNamedCredentialLifecycleStateEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNamedCredentialsResponse wrapper for the ListNamedCredentials operation
type ListNamedCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NamedCredentialCollection instances
	NamedCredentialCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListNamedCredentialsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNamedCredentialsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNamedCredentialsSortOrderEnum Enum with underlying type: string
type ListNamedCredentialsSortOrderEnum string

// Set of constants representing the allowable values for ListNamedCredentialsSortOrderEnum
const (
	ListNamedCredentialsSortOrderAsc  ListNamedCredentialsSortOrderEnum = "ASC"
	ListNamedCredentialsSortOrderDesc ListNamedCredentialsSortOrderEnum = "DESC"
)

var mappingListNamedCredentialsSortOrderEnum = map[string]ListNamedCredentialsSortOrderEnum{
	"ASC":  ListNamedCredentialsSortOrderAsc,
	"DESC": ListNamedCredentialsSortOrderDesc,
}

var mappingListNamedCredentialsSortOrderEnumLowerCase = map[string]ListNamedCredentialsSortOrderEnum{
	"asc":  ListNamedCredentialsSortOrderAsc,
	"desc": ListNamedCredentialsSortOrderDesc,
}

// GetListNamedCredentialsSortOrderEnumValues Enumerates the set of values for ListNamedCredentialsSortOrderEnum
func GetListNamedCredentialsSortOrderEnumValues() []ListNamedCredentialsSortOrderEnum {
	values := make([]ListNamedCredentialsSortOrderEnum, 0)
	for _, v := range mappingListNamedCredentialsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNamedCredentialsSortOrderEnumStringValues Enumerates the set of values in String for ListNamedCredentialsSortOrderEnum
func GetListNamedCredentialsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNamedCredentialsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNamedCredentialsSortOrderEnum(val string) (ListNamedCredentialsSortOrderEnum, bool) {
	enum, ok := mappingListNamedCredentialsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNamedCredentialsSortByEnum Enum with underlying type: string
type ListNamedCredentialsSortByEnum string

// Set of constants representing the allowable values for ListNamedCredentialsSortByEnum
const (
	ListNamedCredentialsSortByName           ListNamedCredentialsSortByEnum = "name"
	ListNamedCredentialsSortByType           ListNamedCredentialsSortByEnum = "type"
	ListNamedCredentialsSortByTimecreated    ListNamedCredentialsSortByEnum = "timeCreated"
	ListNamedCredentialsSortByLifecyclestate ListNamedCredentialsSortByEnum = "lifecycleState"
)

var mappingListNamedCredentialsSortByEnum = map[string]ListNamedCredentialsSortByEnum{
	"name":           ListNamedCredentialsSortByName,
	"type":           ListNamedCredentialsSortByType,
	"timeCreated":    ListNamedCredentialsSortByTimecreated,
	"lifecycleState": ListNamedCredentialsSortByLifecyclestate,
}

var mappingListNamedCredentialsSortByEnumLowerCase = map[string]ListNamedCredentialsSortByEnum{
	"name":           ListNamedCredentialsSortByName,
	"type":           ListNamedCredentialsSortByType,
	"timecreated":    ListNamedCredentialsSortByTimecreated,
	"lifecyclestate": ListNamedCredentialsSortByLifecyclestate,
}

// GetListNamedCredentialsSortByEnumValues Enumerates the set of values for ListNamedCredentialsSortByEnum
func GetListNamedCredentialsSortByEnumValues() []ListNamedCredentialsSortByEnum {
	values := make([]ListNamedCredentialsSortByEnum, 0)
	for _, v := range mappingListNamedCredentialsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNamedCredentialsSortByEnumStringValues Enumerates the set of values in String for ListNamedCredentialsSortByEnum
func GetListNamedCredentialsSortByEnumStringValues() []string {
	return []string{
		"name",
		"type",
		"timeCreated",
		"lifecycleState",
	}
}

// GetMappingListNamedCredentialsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNamedCredentialsSortByEnum(val string) (ListNamedCredentialsSortByEnum, bool) {
	enum, ok := mappingListNamedCredentialsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
