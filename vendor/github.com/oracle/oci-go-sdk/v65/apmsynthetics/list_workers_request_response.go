// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWorkersRequest wrapper for the ListWorkers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmsynthetics/ListWorkers.go.html to see an example of how to use ListWorkersRequest.
type ListWorkersRequest struct {

	// The APM domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// The OCID of the On-premise vantage point.
	OnPremiseVantagePointId *string `mandatory:"true" contributesTo:"path" name:"onPremiseVantagePointId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The maximum number of results per page, or items to return in a paginated
	// "List" call. For information on how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only On-premise VP workers that match the status given.
	Status ListWorkersStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return only On-premise VP workers that match the capability given.
	Capability *string `mandatory:"false" contributesTo:"query" name:"capability"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). Default sort order is ascending.
	SortOrder ListWorkersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order of displayName is ascending.
	// Default order of timeCreated, timeUpdated and timeLastSyncup is descending.
	// The displayName sort by is case-sensitive.
	SortBy ListWorkersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkersStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListWorkersStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkersResponse wrapper for the ListWorkers operation
type ListWorkersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkerCollection instances
	WorkerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkersStatusEnum Enum with underlying type: string
type ListWorkersStatusEnum string

// Set of constants representing the allowable values for ListWorkersStatusEnum
const (
	ListWorkersStatusEnabled  ListWorkersStatusEnum = "ENABLED"
	ListWorkersStatusDisabled ListWorkersStatusEnum = "DISABLED"
)

var mappingListWorkersStatusEnum = map[string]ListWorkersStatusEnum{
	"ENABLED":  ListWorkersStatusEnabled,
	"DISABLED": ListWorkersStatusDisabled,
}

var mappingListWorkersStatusEnumLowerCase = map[string]ListWorkersStatusEnum{
	"enabled":  ListWorkersStatusEnabled,
	"disabled": ListWorkersStatusDisabled,
}

// GetListWorkersStatusEnumValues Enumerates the set of values for ListWorkersStatusEnum
func GetListWorkersStatusEnumValues() []ListWorkersStatusEnum {
	values := make([]ListWorkersStatusEnum, 0)
	for _, v := range mappingListWorkersStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkersStatusEnumStringValues Enumerates the set of values in String for ListWorkersStatusEnum
func GetListWorkersStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingListWorkersStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkersStatusEnum(val string) (ListWorkersStatusEnum, bool) {
	enum, ok := mappingListWorkersStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkersSortOrderEnum Enum with underlying type: string
type ListWorkersSortOrderEnum string

// Set of constants representing the allowable values for ListWorkersSortOrderEnum
const (
	ListWorkersSortOrderAsc  ListWorkersSortOrderEnum = "ASC"
	ListWorkersSortOrderDesc ListWorkersSortOrderEnum = "DESC"
)

var mappingListWorkersSortOrderEnum = map[string]ListWorkersSortOrderEnum{
	"ASC":  ListWorkersSortOrderAsc,
	"DESC": ListWorkersSortOrderDesc,
}

var mappingListWorkersSortOrderEnumLowerCase = map[string]ListWorkersSortOrderEnum{
	"asc":  ListWorkersSortOrderAsc,
	"desc": ListWorkersSortOrderDesc,
}

// GetListWorkersSortOrderEnumValues Enumerates the set of values for ListWorkersSortOrderEnum
func GetListWorkersSortOrderEnumValues() []ListWorkersSortOrderEnum {
	values := make([]ListWorkersSortOrderEnum, 0)
	for _, v := range mappingListWorkersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkersSortOrderEnumStringValues Enumerates the set of values in String for ListWorkersSortOrderEnum
func GetListWorkersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkersSortOrderEnum(val string) (ListWorkersSortOrderEnum, bool) {
	enum, ok := mappingListWorkersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkersSortByEnum Enum with underlying type: string
type ListWorkersSortByEnum string

// Set of constants representing the allowable values for ListWorkersSortByEnum
const (
	ListWorkersSortByDisplayname    ListWorkersSortByEnum = "displayName"
	ListWorkersSortByName           ListWorkersSortByEnum = "name"
	ListWorkersSortByStatus         ListWorkersSortByEnum = "status"
	ListWorkersSortByPriority       ListWorkersSortByEnum = "priority"
	ListWorkersSortByTimecreated    ListWorkersSortByEnum = "timeCreated"
	ListWorkersSortByTimeupdated    ListWorkersSortByEnum = "timeUpdated"
	ListWorkersSortByTimelastsyncup ListWorkersSortByEnum = "timeLastSyncup"
)

var mappingListWorkersSortByEnum = map[string]ListWorkersSortByEnum{
	"displayName":    ListWorkersSortByDisplayname,
	"name":           ListWorkersSortByName,
	"status":         ListWorkersSortByStatus,
	"priority":       ListWorkersSortByPriority,
	"timeCreated":    ListWorkersSortByTimecreated,
	"timeUpdated":    ListWorkersSortByTimeupdated,
	"timeLastSyncup": ListWorkersSortByTimelastsyncup,
}

var mappingListWorkersSortByEnumLowerCase = map[string]ListWorkersSortByEnum{
	"displayname":    ListWorkersSortByDisplayname,
	"name":           ListWorkersSortByName,
	"status":         ListWorkersSortByStatus,
	"priority":       ListWorkersSortByPriority,
	"timecreated":    ListWorkersSortByTimecreated,
	"timeupdated":    ListWorkersSortByTimeupdated,
	"timelastsyncup": ListWorkersSortByTimelastsyncup,
}

// GetListWorkersSortByEnumValues Enumerates the set of values for ListWorkersSortByEnum
func GetListWorkersSortByEnumValues() []ListWorkersSortByEnum {
	values := make([]ListWorkersSortByEnum, 0)
	for _, v := range mappingListWorkersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkersSortByEnumStringValues Enumerates the set of values in String for ListWorkersSortByEnum
func GetListWorkersSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"name",
		"status",
		"priority",
		"timeCreated",
		"timeUpdated",
		"timeLastSyncup",
	}
}

// GetMappingListWorkersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkersSortByEnum(val string) (ListWorkersSortByEnum, bool) {
	enum, ok := mappingListWorkersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
