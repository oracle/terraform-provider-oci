// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExadbVmClusterUpdateHistoryEntriesRequest wrapper for the ListExadbVmClusterUpdateHistoryEntries operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExadbVmClusterUpdateHistoryEntries.go.html to see an example of how to use ListExadbVmClusterUpdateHistoryEntriesRequest.
type ListExadbVmClusterUpdateHistoryEntriesRequest struct {

	// The Exadata VM cluster OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on Exascale Infrastructure.
	ExadbVmClusterId *string `mandatory:"true" contributesTo:"path" name:"exadbVmClusterId"`

	// A filter to return only resources that match the given update type exactly.
	UpdateType ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExadbVmClusterUpdateHistoryEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExadbVmClusterUpdateHistoryEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExadbVmClusterUpdateHistoryEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExadbVmClusterUpdateHistoryEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExadbVmClusterUpdateHistoryEntriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum(string(request.UpdateType)); !ok && request.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", request.UpdateType, strings.Join(GetListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExadbVmClusterUpdateHistoryEntriesResponse wrapper for the ListExadbVmClusterUpdateHistoryEntries operation
type ListExadbVmClusterUpdateHistoryEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExadbVmClusterUpdateHistoryEntrySummary instances
	Items []ExadbVmClusterUpdateHistoryEntrySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExadbVmClusterUpdateHistoryEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExadbVmClusterUpdateHistoryEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum Enum with underlying type: string
type ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum string

// Set of constants representing the allowable values for ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum
const (
	ListExadbVmClusterUpdateHistoryEntriesUpdateTypeGiUpgrade ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum = "GI_UPGRADE"
	ListExadbVmClusterUpdateHistoryEntriesUpdateTypeGiPatch   ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum = "GI_PATCH"
	ListExadbVmClusterUpdateHistoryEntriesUpdateTypeOsUpdate  ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum = "OS_UPDATE"
)

var mappingListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum = map[string]ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum{
	"GI_UPGRADE": ListExadbVmClusterUpdateHistoryEntriesUpdateTypeGiUpgrade,
	"GI_PATCH":   ListExadbVmClusterUpdateHistoryEntriesUpdateTypeGiPatch,
	"OS_UPDATE":  ListExadbVmClusterUpdateHistoryEntriesUpdateTypeOsUpdate,
}

var mappingListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnumLowerCase = map[string]ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum{
	"gi_upgrade": ListExadbVmClusterUpdateHistoryEntriesUpdateTypeGiUpgrade,
	"gi_patch":   ListExadbVmClusterUpdateHistoryEntriesUpdateTypeGiPatch,
	"os_update":  ListExadbVmClusterUpdateHistoryEntriesUpdateTypeOsUpdate,
}

// GetListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnumValues Enumerates the set of values for ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum
func GetListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnumValues() []ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum {
	values := make([]ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum, 0)
	for _, v := range mappingListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnumStringValues Enumerates the set of values in String for ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum
func GetListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum(val string) (ListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnum, bool) {
	enum, ok := mappingListExadbVmClusterUpdateHistoryEntriesUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
