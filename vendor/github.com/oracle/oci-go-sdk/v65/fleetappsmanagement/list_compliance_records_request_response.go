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

// ListComplianceRecordsRequest wrapper for the ListComplianceRecords operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListComplianceRecords.go.html to see an example of how to use ListComplianceRecordsRequest.
type ListComplianceRecordsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Resource identifier.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Entity identifier.Ex:FleetId
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// Product Name.
	ProductName *string `mandatory:"false" contributesTo:"query" name:"productName"`

	// ProductStack name.
	ProductStack *string `mandatory:"false" contributesTo:"query" name:"productStack"`

	// Unique target name
	TargetName *string `mandatory:"false" contributesTo:"query" name:"targetName"`

	// Target Compliance State.
	ComplianceState *string `mandatory:"false" contributesTo:"query" name:"complianceState"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListComplianceRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListComplianceRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListComplianceRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListComplianceRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListComplianceRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListComplianceRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListComplianceRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListComplianceRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListComplianceRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListComplianceRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListComplianceRecordsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListComplianceRecordsResponse wrapper for the ListComplianceRecords operation
type ListComplianceRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ComplianceRecordCollection instances
	ComplianceRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListComplianceRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListComplianceRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListComplianceRecordsSortOrderEnum Enum with underlying type: string
type ListComplianceRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListComplianceRecordsSortOrderEnum
const (
	ListComplianceRecordsSortOrderAsc  ListComplianceRecordsSortOrderEnum = "ASC"
	ListComplianceRecordsSortOrderDesc ListComplianceRecordsSortOrderEnum = "DESC"
)

var mappingListComplianceRecordsSortOrderEnum = map[string]ListComplianceRecordsSortOrderEnum{
	"ASC":  ListComplianceRecordsSortOrderAsc,
	"DESC": ListComplianceRecordsSortOrderDesc,
}

var mappingListComplianceRecordsSortOrderEnumLowerCase = map[string]ListComplianceRecordsSortOrderEnum{
	"asc":  ListComplianceRecordsSortOrderAsc,
	"desc": ListComplianceRecordsSortOrderDesc,
}

// GetListComplianceRecordsSortOrderEnumValues Enumerates the set of values for ListComplianceRecordsSortOrderEnum
func GetListComplianceRecordsSortOrderEnumValues() []ListComplianceRecordsSortOrderEnum {
	values := make([]ListComplianceRecordsSortOrderEnum, 0)
	for _, v := range mappingListComplianceRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListComplianceRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListComplianceRecordsSortOrderEnum
func GetListComplianceRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListComplianceRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComplianceRecordsSortOrderEnum(val string) (ListComplianceRecordsSortOrderEnum, bool) {
	enum, ok := mappingListComplianceRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListComplianceRecordsSortByEnum Enum with underlying type: string
type ListComplianceRecordsSortByEnum string

// Set of constants representing the allowable values for ListComplianceRecordsSortByEnum
const (
	ListComplianceRecordsSortByTimecreated ListComplianceRecordsSortByEnum = "timeCreated"
	ListComplianceRecordsSortByDisplayname ListComplianceRecordsSortByEnum = "displayName"
)

var mappingListComplianceRecordsSortByEnum = map[string]ListComplianceRecordsSortByEnum{
	"timeCreated": ListComplianceRecordsSortByTimecreated,
	"displayName": ListComplianceRecordsSortByDisplayname,
}

var mappingListComplianceRecordsSortByEnumLowerCase = map[string]ListComplianceRecordsSortByEnum{
	"timecreated": ListComplianceRecordsSortByTimecreated,
	"displayname": ListComplianceRecordsSortByDisplayname,
}

// GetListComplianceRecordsSortByEnumValues Enumerates the set of values for ListComplianceRecordsSortByEnum
func GetListComplianceRecordsSortByEnumValues() []ListComplianceRecordsSortByEnum {
	values := make([]ListComplianceRecordsSortByEnum, 0)
	for _, v := range mappingListComplianceRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListComplianceRecordsSortByEnumStringValues Enumerates the set of values in String for ListComplianceRecordsSortByEnum
func GetListComplianceRecordsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListComplianceRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComplianceRecordsSortByEnum(val string) (ListComplianceRecordsSortByEnum, bool) {
	enum, ok := mappingListComplianceRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
