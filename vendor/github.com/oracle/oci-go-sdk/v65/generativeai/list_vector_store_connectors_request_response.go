// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVectorStoreConnectorsRequest wrapper for the ListVectorStoreConnectors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListVectorStoreConnectors.go.html to see an example of how to use ListVectorStoreConnectorsRequest.
type ListVectorStoreConnectorsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The openai compatible id of the VectorStore.
	VectorStoreId *string `mandatory:"false" contributesTo:"query" name:"vectorStoreId"`

	// A filter to return only resources whose lifecycle state matches the given value.
	LifecycleState VectorStoreConnectorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the given name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VectorStoreConnector.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVectorStoreConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated` is descending.
	// Default order for `name` is ascending.
	SortBy ListVectorStoreConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVectorStoreConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVectorStoreConnectorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVectorStoreConnectorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVectorStoreConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVectorStoreConnectorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVectorStoreConnectorLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVectorStoreConnectorLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVectorStoreConnectorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVectorStoreConnectorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVectorStoreConnectorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVectorStoreConnectorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVectorStoreConnectorsResponse wrapper for the ListVectorStoreConnectors operation
type ListVectorStoreConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VectorStoreConnectorCollection instances
	VectorStoreConnectorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVectorStoreConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVectorStoreConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVectorStoreConnectorsSortOrderEnum Enum with underlying type: string
type ListVectorStoreConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListVectorStoreConnectorsSortOrderEnum
const (
	ListVectorStoreConnectorsSortOrderAsc  ListVectorStoreConnectorsSortOrderEnum = "ASC"
	ListVectorStoreConnectorsSortOrderDesc ListVectorStoreConnectorsSortOrderEnum = "DESC"
)

var mappingListVectorStoreConnectorsSortOrderEnum = map[string]ListVectorStoreConnectorsSortOrderEnum{
	"ASC":  ListVectorStoreConnectorsSortOrderAsc,
	"DESC": ListVectorStoreConnectorsSortOrderDesc,
}

var mappingListVectorStoreConnectorsSortOrderEnumLowerCase = map[string]ListVectorStoreConnectorsSortOrderEnum{
	"asc":  ListVectorStoreConnectorsSortOrderAsc,
	"desc": ListVectorStoreConnectorsSortOrderDesc,
}

// GetListVectorStoreConnectorsSortOrderEnumValues Enumerates the set of values for ListVectorStoreConnectorsSortOrderEnum
func GetListVectorStoreConnectorsSortOrderEnumValues() []ListVectorStoreConnectorsSortOrderEnum {
	values := make([]ListVectorStoreConnectorsSortOrderEnum, 0)
	for _, v := range mappingListVectorStoreConnectorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVectorStoreConnectorsSortOrderEnumStringValues Enumerates the set of values in String for ListVectorStoreConnectorsSortOrderEnum
func GetListVectorStoreConnectorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVectorStoreConnectorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVectorStoreConnectorsSortOrderEnum(val string) (ListVectorStoreConnectorsSortOrderEnum, bool) {
	enum, ok := mappingListVectorStoreConnectorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVectorStoreConnectorsSortByEnum Enum with underlying type: string
type ListVectorStoreConnectorsSortByEnum string

// Set of constants representing the allowable values for ListVectorStoreConnectorsSortByEnum
const (
	ListVectorStoreConnectorsSortByName        ListVectorStoreConnectorsSortByEnum = "name"
	ListVectorStoreConnectorsSortByTimecreated ListVectorStoreConnectorsSortByEnum = "timeCreated"
)

var mappingListVectorStoreConnectorsSortByEnum = map[string]ListVectorStoreConnectorsSortByEnum{
	"name":        ListVectorStoreConnectorsSortByName,
	"timeCreated": ListVectorStoreConnectorsSortByTimecreated,
}

var mappingListVectorStoreConnectorsSortByEnumLowerCase = map[string]ListVectorStoreConnectorsSortByEnum{
	"name":        ListVectorStoreConnectorsSortByName,
	"timecreated": ListVectorStoreConnectorsSortByTimecreated,
}

// GetListVectorStoreConnectorsSortByEnumValues Enumerates the set of values for ListVectorStoreConnectorsSortByEnum
func GetListVectorStoreConnectorsSortByEnumValues() []ListVectorStoreConnectorsSortByEnum {
	values := make([]ListVectorStoreConnectorsSortByEnum, 0)
	for _, v := range mappingListVectorStoreConnectorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVectorStoreConnectorsSortByEnumStringValues Enumerates the set of values in String for ListVectorStoreConnectorsSortByEnum
func GetListVectorStoreConnectorsSortByEnumStringValues() []string {
	return []string{
		"name",
		"timeCreated",
	}
}

// GetMappingListVectorStoreConnectorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVectorStoreConnectorsSortByEnum(val string) (ListVectorStoreConnectorsSortByEnum, bool) {
	enum, ok := mappingListVectorStoreConnectorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
