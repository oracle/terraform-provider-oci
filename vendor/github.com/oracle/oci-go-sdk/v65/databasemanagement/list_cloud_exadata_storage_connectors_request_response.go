// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCloudExadataStorageConnectorsRequest wrapper for the ListCloudExadataStorageConnectors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCloudExadataStorageConnectors.go.html to see an example of how to use ListCloudExadataStorageConnectorsRequest.
type ListCloudExadataStorageConnectorsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"true" contributesTo:"query" name:"cloudExadataInfrastructureId"`

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
	SortBy ListCloudExadataStorageConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCloudExadataStorageConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCloudExadataStorageConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudExadataStorageConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudExadataStorageConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudExadataStorageConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudExadataStorageConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudExadataStorageConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCloudExadataStorageConnectorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCloudExadataStorageConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCloudExadataStorageConnectorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudExadataStorageConnectorsResponse wrapper for the ListCloudExadataStorageConnectors operation
type ListCloudExadataStorageConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CloudExadataStorageConnectorCollection instances
	CloudExadataStorageConnectorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudExadataStorageConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudExadataStorageConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudExadataStorageConnectorsSortByEnum Enum with underlying type: string
type ListCloudExadataStorageConnectorsSortByEnum string

// Set of constants representing the allowable values for ListCloudExadataStorageConnectorsSortByEnum
const (
	ListCloudExadataStorageConnectorsSortByTimecreated ListCloudExadataStorageConnectorsSortByEnum = "TIMECREATED"
	ListCloudExadataStorageConnectorsSortByName        ListCloudExadataStorageConnectorsSortByEnum = "NAME"
)

var mappingListCloudExadataStorageConnectorsSortByEnum = map[string]ListCloudExadataStorageConnectorsSortByEnum{
	"TIMECREATED": ListCloudExadataStorageConnectorsSortByTimecreated,
	"NAME":        ListCloudExadataStorageConnectorsSortByName,
}

var mappingListCloudExadataStorageConnectorsSortByEnumLowerCase = map[string]ListCloudExadataStorageConnectorsSortByEnum{
	"timecreated": ListCloudExadataStorageConnectorsSortByTimecreated,
	"name":        ListCloudExadataStorageConnectorsSortByName,
}

// GetListCloudExadataStorageConnectorsSortByEnumValues Enumerates the set of values for ListCloudExadataStorageConnectorsSortByEnum
func GetListCloudExadataStorageConnectorsSortByEnumValues() []ListCloudExadataStorageConnectorsSortByEnum {
	values := make([]ListCloudExadataStorageConnectorsSortByEnum, 0)
	for _, v := range mappingListCloudExadataStorageConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudExadataStorageConnectorsSortByEnumStringValues Enumerates the set of values in String for ListCloudExadataStorageConnectorsSortByEnum
func GetListCloudExadataStorageConnectorsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListCloudExadataStorageConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudExadataStorageConnectorsSortByEnum(val string) (ListCloudExadataStorageConnectorsSortByEnum, bool) {
	enum, ok := mappingListCloudExadataStorageConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCloudExadataStorageConnectorsSortOrderEnum Enum with underlying type: string
type ListCloudExadataStorageConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListCloudExadataStorageConnectorsSortOrderEnum
const (
	ListCloudExadataStorageConnectorsSortOrderAsc  ListCloudExadataStorageConnectorsSortOrderEnum = "ASC"
	ListCloudExadataStorageConnectorsSortOrderDesc ListCloudExadataStorageConnectorsSortOrderEnum = "DESC"
)

var mappingListCloudExadataStorageConnectorsSortOrderEnum = map[string]ListCloudExadataStorageConnectorsSortOrderEnum{
	"ASC":  ListCloudExadataStorageConnectorsSortOrderAsc,
	"DESC": ListCloudExadataStorageConnectorsSortOrderDesc,
}

var mappingListCloudExadataStorageConnectorsSortOrderEnumLowerCase = map[string]ListCloudExadataStorageConnectorsSortOrderEnum{
	"asc":  ListCloudExadataStorageConnectorsSortOrderAsc,
	"desc": ListCloudExadataStorageConnectorsSortOrderDesc,
}

// GetListCloudExadataStorageConnectorsSortOrderEnumValues Enumerates the set of values for ListCloudExadataStorageConnectorsSortOrderEnum
func GetListCloudExadataStorageConnectorsSortOrderEnumValues() []ListCloudExadataStorageConnectorsSortOrderEnum {
	values := make([]ListCloudExadataStorageConnectorsSortOrderEnum, 0)
	for _, v := range mappingListCloudExadataStorageConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudExadataStorageConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListCloudExadataStorageConnectorsSortOrderEnum
func GetListCloudExadataStorageConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCloudExadataStorageConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudExadataStorageConnectorsSortOrderEnum(val string) (ListCloudExadataStorageConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListCloudExadataStorageConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
