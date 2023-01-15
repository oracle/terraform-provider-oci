// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListReplicasRequest wrapper for the ListReplicas operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/ListReplicas.go.html to see an example of how to use ListReplicasRequest.
type ListReplicasRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated list call. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from
	// the previous list call. For information about pagination, see List
	// Pagination (https://docs.cloud.oracle.comAPI/Concepts/usingapi.htm#List_Pagination).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resource matching the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The DB System OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"false" contributesTo:"query" name:"dbSystemId"`

	// The LifecycleState of the read replica.
	LifecycleState ReplicaSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The read replica OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ReplicaId *string `mandatory:"false" contributesTo:"query" name:"replicaId"`

	// The field to sort by. You can sort by one field only. By default, the Time field is sorted in descending order and the Display Name field in ascending order.
	SortBy ListReplicasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use (ASC or DESC).
	SortOrder ListReplicasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReplicasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReplicasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListReplicasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReplicasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListReplicasRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicaSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetReplicaSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReplicasSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListReplicasSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListReplicasSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListReplicasSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListReplicasResponse wrapper for the ListReplicas operation
type ListReplicasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ReplicaSummary instances
	Items []ReplicaSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListReplicasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReplicasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReplicasSortByEnum Enum with underlying type: string
type ListReplicasSortByEnum string

// Set of constants representing the allowable values for ListReplicasSortByEnum
const (
	ListReplicasSortByTimecreated ListReplicasSortByEnum = "timeCreated"
	ListReplicasSortByDisplayname ListReplicasSortByEnum = "displayName"
)

var mappingListReplicasSortByEnum = map[string]ListReplicasSortByEnum{
	"timeCreated": ListReplicasSortByTimecreated,
	"displayName": ListReplicasSortByDisplayname,
}

var mappingListReplicasSortByEnumLowerCase = map[string]ListReplicasSortByEnum{
	"timecreated": ListReplicasSortByTimecreated,
	"displayname": ListReplicasSortByDisplayname,
}

// GetListReplicasSortByEnumValues Enumerates the set of values for ListReplicasSortByEnum
func GetListReplicasSortByEnumValues() []ListReplicasSortByEnum {
	values := make([]ListReplicasSortByEnum, 0)
	for _, v := range mappingListReplicasSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicasSortByEnumStringValues Enumerates the set of values in String for ListReplicasSortByEnum
func GetListReplicasSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListReplicasSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicasSortByEnum(val string) (ListReplicasSortByEnum, bool) {
	enum, ok := mappingListReplicasSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListReplicasSortOrderEnum Enum with underlying type: string
type ListReplicasSortOrderEnum string

// Set of constants representing the allowable values for ListReplicasSortOrderEnum
const (
	ListReplicasSortOrderAsc  ListReplicasSortOrderEnum = "ASC"
	ListReplicasSortOrderDesc ListReplicasSortOrderEnum = "DESC"
)

var mappingListReplicasSortOrderEnum = map[string]ListReplicasSortOrderEnum{
	"ASC":  ListReplicasSortOrderAsc,
	"DESC": ListReplicasSortOrderDesc,
}

var mappingListReplicasSortOrderEnumLowerCase = map[string]ListReplicasSortOrderEnum{
	"asc":  ListReplicasSortOrderAsc,
	"desc": ListReplicasSortOrderDesc,
}

// GetListReplicasSortOrderEnumValues Enumerates the set of values for ListReplicasSortOrderEnum
func GetListReplicasSortOrderEnumValues() []ListReplicasSortOrderEnum {
	values := make([]ListReplicasSortOrderEnum, 0)
	for _, v := range mappingListReplicasSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListReplicasSortOrderEnumStringValues Enumerates the set of values in String for ListReplicasSortOrderEnum
func GetListReplicasSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListReplicasSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListReplicasSortOrderEnum(val string) (ListReplicasSortOrderEnum, bool) {
	enum, ok := mappingListReplicasSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
