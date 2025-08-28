// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceServersRequest wrapper for the ListManagedInstanceServers operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListManagedInstanceServers.go.html to see an example of how to use ListManagedInstanceServersRequest.
type ListManagedInstanceServersRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The name of the resource.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListManagedInstanceServersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for _timeCreated_ is **descending**.
	// Default order for _name_ is **ascending**.
	// If no value is specified, _timeCreated_ is default.
	SortBy ListManagedInstanceServersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceServersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceServersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceServersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceServersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceServersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceServersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceServersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceServersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceServersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceServersResponse wrapper for the ListManagedInstanceServers operation
type ListManagedInstanceServersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServerCollection instances
	ServerCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListManagedInstanceServersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceServersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceServersSortOrderEnum Enum with underlying type: string
type ListManagedInstanceServersSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceServersSortOrderEnum
const (
	ListManagedInstanceServersSortOrderAsc  ListManagedInstanceServersSortOrderEnum = "ASC"
	ListManagedInstanceServersSortOrderDesc ListManagedInstanceServersSortOrderEnum = "DESC"
)

var mappingListManagedInstanceServersSortOrderEnum = map[string]ListManagedInstanceServersSortOrderEnum{
	"ASC":  ListManagedInstanceServersSortOrderAsc,
	"DESC": ListManagedInstanceServersSortOrderDesc,
}

var mappingListManagedInstanceServersSortOrderEnumLowerCase = map[string]ListManagedInstanceServersSortOrderEnum{
	"asc":  ListManagedInstanceServersSortOrderAsc,
	"desc": ListManagedInstanceServersSortOrderDesc,
}

// GetListManagedInstanceServersSortOrderEnumValues Enumerates the set of values for ListManagedInstanceServersSortOrderEnum
func GetListManagedInstanceServersSortOrderEnumValues() []ListManagedInstanceServersSortOrderEnum {
	values := make([]ListManagedInstanceServersSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceServersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceServersSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceServersSortOrderEnum
func GetListManagedInstanceServersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceServersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceServersSortOrderEnum(val string) (ListManagedInstanceServersSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceServersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceServersSortByEnum Enum with underlying type: string
type ListManagedInstanceServersSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceServersSortByEnum
const (
	ListManagedInstanceServersSortByTimecreated ListManagedInstanceServersSortByEnum = "timeCreated"
	ListManagedInstanceServersSortByName        ListManagedInstanceServersSortByEnum = "name"
)

var mappingListManagedInstanceServersSortByEnum = map[string]ListManagedInstanceServersSortByEnum{
	"timeCreated": ListManagedInstanceServersSortByTimecreated,
	"name":        ListManagedInstanceServersSortByName,
}

var mappingListManagedInstanceServersSortByEnumLowerCase = map[string]ListManagedInstanceServersSortByEnum{
	"timecreated": ListManagedInstanceServersSortByTimecreated,
	"name":        ListManagedInstanceServersSortByName,
}

// GetListManagedInstanceServersSortByEnumValues Enumerates the set of values for ListManagedInstanceServersSortByEnum
func GetListManagedInstanceServersSortByEnumValues() []ListManagedInstanceServersSortByEnum {
	values := make([]ListManagedInstanceServersSortByEnum, 0)
	for _, v := range mappingListManagedInstanceServersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceServersSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceServersSortByEnum
func GetListManagedInstanceServersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListManagedInstanceServersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceServersSortByEnum(val string) (ListManagedInstanceServersSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceServersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
