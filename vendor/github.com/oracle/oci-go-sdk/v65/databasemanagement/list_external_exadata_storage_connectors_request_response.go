// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExternalExadataStorageConnectorsRequest wrapper for the ListExternalExadataStorageConnectors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalExadataStorageConnectors.go.html to see an example of how to use ListExternalExadataStorageConnectorsRequest.
type ListExternalExadataStorageConnectorsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExternalExadataInfrastructureId *string `mandatory:"true" contributesTo:"query" name:"externalExadataInfrastructureId"`

	// The optional single value query filter parameter on the entity display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListExternalExadataStorageConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalExadataStorageConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalExadataStorageConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalExadataStorageConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalExadataStorageConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalExadataStorageConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalExadataStorageConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalExadataStorageConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalExadataStorageConnectorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalExadataStorageConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalExadataStorageConnectorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalExadataStorageConnectorsResponse wrapper for the ListExternalExadataStorageConnectors operation
type ListExternalExadataStorageConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalExadataStorageConnectorCollection instances
	ExternalExadataStorageConnectorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalExadataStorageConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalExadataStorageConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalExadataStorageConnectorsSortByEnum Enum with underlying type: string
type ListExternalExadataStorageConnectorsSortByEnum string

// Set of constants representing the allowable values for ListExternalExadataStorageConnectorsSortByEnum
const (
	ListExternalExadataStorageConnectorsSortByTimecreated ListExternalExadataStorageConnectorsSortByEnum = "TIMECREATED"
	ListExternalExadataStorageConnectorsSortByName        ListExternalExadataStorageConnectorsSortByEnum = "NAME"
)

var mappingListExternalExadataStorageConnectorsSortByEnum = map[string]ListExternalExadataStorageConnectorsSortByEnum{
	"TIMECREATED": ListExternalExadataStorageConnectorsSortByTimecreated,
	"NAME":        ListExternalExadataStorageConnectorsSortByName,
}

var mappingListExternalExadataStorageConnectorsSortByEnumLowerCase = map[string]ListExternalExadataStorageConnectorsSortByEnum{
	"timecreated": ListExternalExadataStorageConnectorsSortByTimecreated,
	"name":        ListExternalExadataStorageConnectorsSortByName,
}

// GetListExternalExadataStorageConnectorsSortByEnumValues Enumerates the set of values for ListExternalExadataStorageConnectorsSortByEnum
func GetListExternalExadataStorageConnectorsSortByEnumValues() []ListExternalExadataStorageConnectorsSortByEnum {
	values := make([]ListExternalExadataStorageConnectorsSortByEnum, 0)
	for _, v := range mappingListExternalExadataStorageConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalExadataStorageConnectorsSortByEnumStringValues Enumerates the set of values in String for ListExternalExadataStorageConnectorsSortByEnum
func GetListExternalExadataStorageConnectorsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListExternalExadataStorageConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalExadataStorageConnectorsSortByEnum(val string) (ListExternalExadataStorageConnectorsSortByEnum, bool) {
	enum, ok := mappingListExternalExadataStorageConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalExadataStorageConnectorsSortOrderEnum Enum with underlying type: string
type ListExternalExadataStorageConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListExternalExadataStorageConnectorsSortOrderEnum
const (
	ListExternalExadataStorageConnectorsSortOrderAsc  ListExternalExadataStorageConnectorsSortOrderEnum = "ASC"
	ListExternalExadataStorageConnectorsSortOrderDesc ListExternalExadataStorageConnectorsSortOrderEnum = "DESC"
)

var mappingListExternalExadataStorageConnectorsSortOrderEnum = map[string]ListExternalExadataStorageConnectorsSortOrderEnum{
	"ASC":  ListExternalExadataStorageConnectorsSortOrderAsc,
	"DESC": ListExternalExadataStorageConnectorsSortOrderDesc,
}

var mappingListExternalExadataStorageConnectorsSortOrderEnumLowerCase = map[string]ListExternalExadataStorageConnectorsSortOrderEnum{
	"asc":  ListExternalExadataStorageConnectorsSortOrderAsc,
	"desc": ListExternalExadataStorageConnectorsSortOrderDesc,
}

// GetListExternalExadataStorageConnectorsSortOrderEnumValues Enumerates the set of values for ListExternalExadataStorageConnectorsSortOrderEnum
func GetListExternalExadataStorageConnectorsSortOrderEnumValues() []ListExternalExadataStorageConnectorsSortOrderEnum {
	values := make([]ListExternalExadataStorageConnectorsSortOrderEnum, 0)
	for _, v := range mappingListExternalExadataStorageConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalExadataStorageConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListExternalExadataStorageConnectorsSortOrderEnum
func GetListExternalExadataStorageConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalExadataStorageConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalExadataStorageConnectorsSortOrderEnum(val string) (ListExternalExadataStorageConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListExternalExadataStorageConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
