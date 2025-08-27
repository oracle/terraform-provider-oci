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

// ListWlsDomainServerBackupsRequest wrapper for the ListWlsDomainServerBackups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainServerBackups.go.html to see an example of how to use ListWlsDomainServerBackupsRequest.
type ListWlsDomainServerBackupsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" contributesTo:"path" name:"wlsDomainId"`

	// The unique identifier of a server.
	// **Note:** Not an OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ServerId *string `mandatory:"true" contributesTo:"path" name:"serverId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token that represents the page at which to start retrieving results. The token is usually retrieved from a previous List call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order is either 'ASC' or 'DESC'.
	SortOrder ListWlsDomainServerBackupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort the resource. Only one sort order may be provided.
	// Default order for timeCreated is **descending**.
	// If no value is specified, timeCreated is default.
	SortBy ListWlsDomainServerBackupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWlsDomainServerBackupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWlsDomainServerBackupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWlsDomainServerBackupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWlsDomainServerBackupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWlsDomainServerBackupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWlsDomainServerBackupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWlsDomainServerBackupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWlsDomainServerBackupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWlsDomainServerBackupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWlsDomainServerBackupsResponse wrapper for the ListWlsDomainServerBackups operation
type ListWlsDomainServerBackupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BackupCollection instances
	BackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWlsDomainServerBackupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWlsDomainServerBackupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWlsDomainServerBackupsSortOrderEnum Enum with underlying type: string
type ListWlsDomainServerBackupsSortOrderEnum string

// Set of constants representing the allowable values for ListWlsDomainServerBackupsSortOrderEnum
const (
	ListWlsDomainServerBackupsSortOrderAsc  ListWlsDomainServerBackupsSortOrderEnum = "ASC"
	ListWlsDomainServerBackupsSortOrderDesc ListWlsDomainServerBackupsSortOrderEnum = "DESC"
)

var mappingListWlsDomainServerBackupsSortOrderEnum = map[string]ListWlsDomainServerBackupsSortOrderEnum{
	"ASC":  ListWlsDomainServerBackupsSortOrderAsc,
	"DESC": ListWlsDomainServerBackupsSortOrderDesc,
}

var mappingListWlsDomainServerBackupsSortOrderEnumLowerCase = map[string]ListWlsDomainServerBackupsSortOrderEnum{
	"asc":  ListWlsDomainServerBackupsSortOrderAsc,
	"desc": ListWlsDomainServerBackupsSortOrderDesc,
}

// GetListWlsDomainServerBackupsSortOrderEnumValues Enumerates the set of values for ListWlsDomainServerBackupsSortOrderEnum
func GetListWlsDomainServerBackupsSortOrderEnumValues() []ListWlsDomainServerBackupsSortOrderEnum {
	values := make([]ListWlsDomainServerBackupsSortOrderEnum, 0)
	for _, v := range mappingListWlsDomainServerBackupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainServerBackupsSortOrderEnumStringValues Enumerates the set of values in String for ListWlsDomainServerBackupsSortOrderEnum
func GetListWlsDomainServerBackupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWlsDomainServerBackupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainServerBackupsSortOrderEnum(val string) (ListWlsDomainServerBackupsSortOrderEnum, bool) {
	enum, ok := mappingListWlsDomainServerBackupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWlsDomainServerBackupsSortByEnum Enum with underlying type: string
type ListWlsDomainServerBackupsSortByEnum string

// Set of constants representing the allowable values for ListWlsDomainServerBackupsSortByEnum
const (
	ListWlsDomainServerBackupsSortByTimecreated ListWlsDomainServerBackupsSortByEnum = "timeCreated"
)

var mappingListWlsDomainServerBackupsSortByEnum = map[string]ListWlsDomainServerBackupsSortByEnum{
	"timeCreated": ListWlsDomainServerBackupsSortByTimecreated,
}

var mappingListWlsDomainServerBackupsSortByEnumLowerCase = map[string]ListWlsDomainServerBackupsSortByEnum{
	"timecreated": ListWlsDomainServerBackupsSortByTimecreated,
}

// GetListWlsDomainServerBackupsSortByEnumValues Enumerates the set of values for ListWlsDomainServerBackupsSortByEnum
func GetListWlsDomainServerBackupsSortByEnumValues() []ListWlsDomainServerBackupsSortByEnum {
	values := make([]ListWlsDomainServerBackupsSortByEnum, 0)
	for _, v := range mappingListWlsDomainServerBackupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWlsDomainServerBackupsSortByEnumStringValues Enumerates the set of values in String for ListWlsDomainServerBackupsSortByEnum
func GetListWlsDomainServerBackupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListWlsDomainServerBackupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWlsDomainServerBackupsSortByEnum(val string) (ListWlsDomainServerBackupsSortByEnum, bool) {
	enum, ok := mappingListWlsDomainServerBackupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
