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

// ListVectorStoreConnectorFileSyncsRequest wrapper for the ListVectorStoreConnectorFileSyncs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListVectorStoreConnectorFileSyncs.go.html to see an example of how to use ListVectorStoreConnectorFileSyncsRequest.
type ListVectorStoreConnectorFileSyncsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycle state matches the given value
	LifecycleState VectorStoreConnectorFileSyncLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vectorStoreConnectorFileSync.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VectorStoreConnector.
	VectorStoreConnectorId *string `mandatory:"false" contributesTo:"query" name:"vectorStoreConnectorId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVectorStoreConnectorFileSyncsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `timeCreated`
	// is descending. Default order for `displayName` is ascending.
	SortBy ListVectorStoreConnectorFileSyncsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVectorStoreConnectorFileSyncsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVectorStoreConnectorFileSyncsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVectorStoreConnectorFileSyncsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVectorStoreConnectorFileSyncsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVectorStoreConnectorFileSyncsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVectorStoreConnectorFileSyncLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetVectorStoreConnectorFileSyncLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVectorStoreConnectorFileSyncsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVectorStoreConnectorFileSyncsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVectorStoreConnectorFileSyncsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVectorStoreConnectorFileSyncsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVectorStoreConnectorFileSyncsResponse wrapper for the ListVectorStoreConnectorFileSyncs operation
type ListVectorStoreConnectorFileSyncsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VectorStoreConnectorFileSyncCollection instances
	VectorStoreConnectorFileSyncCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVectorStoreConnectorFileSyncsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVectorStoreConnectorFileSyncsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVectorStoreConnectorFileSyncsSortOrderEnum Enum with underlying type: string
type ListVectorStoreConnectorFileSyncsSortOrderEnum string

// Set of constants representing the allowable values for ListVectorStoreConnectorFileSyncsSortOrderEnum
const (
	ListVectorStoreConnectorFileSyncsSortOrderAsc  ListVectorStoreConnectorFileSyncsSortOrderEnum = "ASC"
	ListVectorStoreConnectorFileSyncsSortOrderDesc ListVectorStoreConnectorFileSyncsSortOrderEnum = "DESC"
)

var mappingListVectorStoreConnectorFileSyncsSortOrderEnum = map[string]ListVectorStoreConnectorFileSyncsSortOrderEnum{
	"ASC":  ListVectorStoreConnectorFileSyncsSortOrderAsc,
	"DESC": ListVectorStoreConnectorFileSyncsSortOrderDesc,
}

var mappingListVectorStoreConnectorFileSyncsSortOrderEnumLowerCase = map[string]ListVectorStoreConnectorFileSyncsSortOrderEnum{
	"asc":  ListVectorStoreConnectorFileSyncsSortOrderAsc,
	"desc": ListVectorStoreConnectorFileSyncsSortOrderDesc,
}

// GetListVectorStoreConnectorFileSyncsSortOrderEnumValues Enumerates the set of values for ListVectorStoreConnectorFileSyncsSortOrderEnum
func GetListVectorStoreConnectorFileSyncsSortOrderEnumValues() []ListVectorStoreConnectorFileSyncsSortOrderEnum {
	values := make([]ListVectorStoreConnectorFileSyncsSortOrderEnum, 0)
	for _, v := range mappingListVectorStoreConnectorFileSyncsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVectorStoreConnectorFileSyncsSortOrderEnumStringValues Enumerates the set of values in String for ListVectorStoreConnectorFileSyncsSortOrderEnum
func GetListVectorStoreConnectorFileSyncsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVectorStoreConnectorFileSyncsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVectorStoreConnectorFileSyncsSortOrderEnum(val string) (ListVectorStoreConnectorFileSyncsSortOrderEnum, bool) {
	enum, ok := mappingListVectorStoreConnectorFileSyncsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVectorStoreConnectorFileSyncsSortByEnum Enum with underlying type: string
type ListVectorStoreConnectorFileSyncsSortByEnum string

// Set of constants representing the allowable values for ListVectorStoreConnectorFileSyncsSortByEnum
const (
	ListVectorStoreConnectorFileSyncsSortByDisplayname ListVectorStoreConnectorFileSyncsSortByEnum = "displayName"
	ListVectorStoreConnectorFileSyncsSortByTimecreated ListVectorStoreConnectorFileSyncsSortByEnum = "timeCreated"
)

var mappingListVectorStoreConnectorFileSyncsSortByEnum = map[string]ListVectorStoreConnectorFileSyncsSortByEnum{
	"displayName": ListVectorStoreConnectorFileSyncsSortByDisplayname,
	"timeCreated": ListVectorStoreConnectorFileSyncsSortByTimecreated,
}

var mappingListVectorStoreConnectorFileSyncsSortByEnumLowerCase = map[string]ListVectorStoreConnectorFileSyncsSortByEnum{
	"displayname": ListVectorStoreConnectorFileSyncsSortByDisplayname,
	"timecreated": ListVectorStoreConnectorFileSyncsSortByTimecreated,
}

// GetListVectorStoreConnectorFileSyncsSortByEnumValues Enumerates the set of values for ListVectorStoreConnectorFileSyncsSortByEnum
func GetListVectorStoreConnectorFileSyncsSortByEnumValues() []ListVectorStoreConnectorFileSyncsSortByEnum {
	values := make([]ListVectorStoreConnectorFileSyncsSortByEnum, 0)
	for _, v := range mappingListVectorStoreConnectorFileSyncsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVectorStoreConnectorFileSyncsSortByEnumStringValues Enumerates the set of values in String for ListVectorStoreConnectorFileSyncsSortByEnum
func GetListVectorStoreConnectorFileSyncsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListVectorStoreConnectorFileSyncsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVectorStoreConnectorFileSyncsSortByEnum(val string) (ListVectorStoreConnectorFileSyncsSortByEnum, bool) {
	enum, ok := mappingListVectorStoreConnectorFileSyncsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
