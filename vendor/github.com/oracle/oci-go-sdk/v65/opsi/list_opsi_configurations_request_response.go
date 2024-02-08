// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOpsiConfigurationsRequest wrapper for the ListOpsiConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOpsiConfigurations.go.html to see an example of how to use ListOpsiConfigurationsRequest.
type ListOpsiConfigurationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter to return based on resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter to return based on Lifecycle state of OPSI configuration.
	LifecycleState []OpsiConfigurationLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// Filter to return based on configuration type of OPSI configuration.
	OpsiConfigType []OpsiConfigurationTypeEnum `contributesTo:"query" name:"opsiConfigType" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOpsiConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// OPSI configurations list sort options.
	SortBy ListOpsiConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOpsiConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOpsiConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOpsiConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOpsiConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOpsiConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingOpsiConfigurationLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetOpsiConfigurationLifecycleStateEnumStringValues(), ",")))
		}
	}

	for _, val := range request.OpsiConfigType {
		if _, ok := GetMappingOpsiConfigurationTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpsiConfigType: %s. Supported values are: %s.", val, strings.Join(GetOpsiConfigurationTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListOpsiConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOpsiConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpsiConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOpsiConfigurationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOpsiConfigurationsResponse wrapper for the ListOpsiConfigurations operation
type ListOpsiConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OpsiConfigurationsCollection instances
	OpsiConfigurationsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOpsiConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOpsiConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOpsiConfigurationsSortOrderEnum Enum with underlying type: string
type ListOpsiConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListOpsiConfigurationsSortOrderEnum
const (
	ListOpsiConfigurationsSortOrderAsc  ListOpsiConfigurationsSortOrderEnum = "ASC"
	ListOpsiConfigurationsSortOrderDesc ListOpsiConfigurationsSortOrderEnum = "DESC"
)

var mappingListOpsiConfigurationsSortOrderEnum = map[string]ListOpsiConfigurationsSortOrderEnum{
	"ASC":  ListOpsiConfigurationsSortOrderAsc,
	"DESC": ListOpsiConfigurationsSortOrderDesc,
}

var mappingListOpsiConfigurationsSortOrderEnumLowerCase = map[string]ListOpsiConfigurationsSortOrderEnum{
	"asc":  ListOpsiConfigurationsSortOrderAsc,
	"desc": ListOpsiConfigurationsSortOrderDesc,
}

// GetListOpsiConfigurationsSortOrderEnumValues Enumerates the set of values for ListOpsiConfigurationsSortOrderEnum
func GetListOpsiConfigurationsSortOrderEnumValues() []ListOpsiConfigurationsSortOrderEnum {
	values := make([]ListOpsiConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListOpsiConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpsiConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListOpsiConfigurationsSortOrderEnum
func GetListOpsiConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOpsiConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpsiConfigurationsSortOrderEnum(val string) (ListOpsiConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListOpsiConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOpsiConfigurationsSortByEnum Enum with underlying type: string
type ListOpsiConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListOpsiConfigurationsSortByEnum
const (
	ListOpsiConfigurationsSortByDisplayname ListOpsiConfigurationsSortByEnum = "displayName"
)

var mappingListOpsiConfigurationsSortByEnum = map[string]ListOpsiConfigurationsSortByEnum{
	"displayName": ListOpsiConfigurationsSortByDisplayname,
}

var mappingListOpsiConfigurationsSortByEnumLowerCase = map[string]ListOpsiConfigurationsSortByEnum{
	"displayname": ListOpsiConfigurationsSortByDisplayname,
}

// GetListOpsiConfigurationsSortByEnumValues Enumerates the set of values for ListOpsiConfigurationsSortByEnum
func GetListOpsiConfigurationsSortByEnumValues() []ListOpsiConfigurationsSortByEnum {
	values := make([]ListOpsiConfigurationsSortByEnum, 0)
	for _, v := range mappingListOpsiConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpsiConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListOpsiConfigurationsSortByEnum
func GetListOpsiConfigurationsSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListOpsiConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpsiConfigurationsSortByEnum(val string) (ListOpsiConfigurationsSortByEnum, bool) {
	enum, ok := mappingListOpsiConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
