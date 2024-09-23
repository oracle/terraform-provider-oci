// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFleetCredentialsRequest wrapper for the ListFleetCredentials operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListFleetCredentials.go.html to see an example of how to use ListFleetCredentialsRequest.
type ListFleetCredentialsRequest struct {

	// unique Fleet identifier
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState FleetCredentialLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Credential Level.
	CredentialLevel CredentialEntitySpecificDetailsCredentialLevelEnum `mandatory:"false" contributesTo:"query" name:"credentialLevel" omitEmpty:"true"`

	// unique FleetCredential identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFleetCredentialsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFleetCredentialsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetCredentialsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetCredentialsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetCredentialsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetCredentialsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFleetCredentialsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFleetCredentialLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFleetCredentialLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCredentialEntitySpecificDetailsCredentialLevelEnum(string(request.CredentialLevel)); !ok && request.CredentialLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CredentialLevel: %s. Supported values are: %s.", request.CredentialLevel, strings.Join(GetCredentialEntitySpecificDetailsCredentialLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetCredentialsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFleetCredentialsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFleetCredentialsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFleetCredentialsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFleetCredentialsResponse wrapper for the ListFleetCredentials operation
type ListFleetCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetCredentialCollection instances
	FleetCredentialCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetCredentialsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetCredentialsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetCredentialsSortOrderEnum Enum with underlying type: string
type ListFleetCredentialsSortOrderEnum string

// Set of constants representing the allowable values for ListFleetCredentialsSortOrderEnum
const (
	ListFleetCredentialsSortOrderAsc  ListFleetCredentialsSortOrderEnum = "ASC"
	ListFleetCredentialsSortOrderDesc ListFleetCredentialsSortOrderEnum = "DESC"
)

var mappingListFleetCredentialsSortOrderEnum = map[string]ListFleetCredentialsSortOrderEnum{
	"ASC":  ListFleetCredentialsSortOrderAsc,
	"DESC": ListFleetCredentialsSortOrderDesc,
}

var mappingListFleetCredentialsSortOrderEnumLowerCase = map[string]ListFleetCredentialsSortOrderEnum{
	"asc":  ListFleetCredentialsSortOrderAsc,
	"desc": ListFleetCredentialsSortOrderDesc,
}

// GetListFleetCredentialsSortOrderEnumValues Enumerates the set of values for ListFleetCredentialsSortOrderEnum
func GetListFleetCredentialsSortOrderEnumValues() []ListFleetCredentialsSortOrderEnum {
	values := make([]ListFleetCredentialsSortOrderEnum, 0)
	for _, v := range mappingListFleetCredentialsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetCredentialsSortOrderEnumStringValues Enumerates the set of values in String for ListFleetCredentialsSortOrderEnum
func GetListFleetCredentialsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFleetCredentialsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetCredentialsSortOrderEnum(val string) (ListFleetCredentialsSortOrderEnum, bool) {
	enum, ok := mappingListFleetCredentialsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFleetCredentialsSortByEnum Enum with underlying type: string
type ListFleetCredentialsSortByEnum string

// Set of constants representing the allowable values for ListFleetCredentialsSortByEnum
const (
	ListFleetCredentialsSortByTimecreated ListFleetCredentialsSortByEnum = "timeCreated"
	ListFleetCredentialsSortByDisplayname ListFleetCredentialsSortByEnum = "displayName"
)

var mappingListFleetCredentialsSortByEnum = map[string]ListFleetCredentialsSortByEnum{
	"timeCreated": ListFleetCredentialsSortByTimecreated,
	"displayName": ListFleetCredentialsSortByDisplayname,
}

var mappingListFleetCredentialsSortByEnumLowerCase = map[string]ListFleetCredentialsSortByEnum{
	"timecreated": ListFleetCredentialsSortByTimecreated,
	"displayname": ListFleetCredentialsSortByDisplayname,
}

// GetListFleetCredentialsSortByEnumValues Enumerates the set of values for ListFleetCredentialsSortByEnum
func GetListFleetCredentialsSortByEnumValues() []ListFleetCredentialsSortByEnum {
	values := make([]ListFleetCredentialsSortByEnum, 0)
	for _, v := range mappingListFleetCredentialsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFleetCredentialsSortByEnumStringValues Enumerates the set of values in String for ListFleetCredentialsSortByEnum
func GetListFleetCredentialsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFleetCredentialsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFleetCredentialsSortByEnum(val string) (ListFleetCredentialsSortByEnum, bool) {
	enum, ok := mappingListFleetCredentialsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
