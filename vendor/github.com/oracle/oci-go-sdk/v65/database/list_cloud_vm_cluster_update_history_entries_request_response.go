// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCloudVmClusterUpdateHistoryEntriesRequest wrapper for the ListCloudVmClusterUpdateHistoryEntries operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListCloudVmClusterUpdateHistoryEntries.go.html to see an example of how to use ListCloudVmClusterUpdateHistoryEntriesRequest.
type ListCloudVmClusterUpdateHistoryEntriesRequest struct {

	// The cloud VM cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CloudVmClusterId *string `mandatory:"true" contributesTo:"path" name:"cloudVmClusterId"`

	// A filter to return only resources that match the given update type exactly.
	UpdateType ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

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

func (request ListCloudVmClusterUpdateHistoryEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCloudVmClusterUpdateHistoryEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCloudVmClusterUpdateHistoryEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCloudVmClusterUpdateHistoryEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCloudVmClusterUpdateHistoryEntriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum(string(request.UpdateType)); !ok && request.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", request.UpdateType, strings.Join(GetListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCloudVmClusterUpdateHistoryEntriesResponse wrapper for the ListCloudVmClusterUpdateHistoryEntries operation
type ListCloudVmClusterUpdateHistoryEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []UpdateHistoryEntrySummary instances
	Items []UpdateHistoryEntrySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCloudVmClusterUpdateHistoryEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCloudVmClusterUpdateHistoryEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum Enum with underlying type: string
type ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum string

// Set of constants representing the allowable values for ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum
const (
	ListCloudVmClusterUpdateHistoryEntriesUpdateTypeGiUpgrade ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum = "GI_UPGRADE"
	ListCloudVmClusterUpdateHistoryEntriesUpdateTypeGiPatch   ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum = "GI_PATCH"
	ListCloudVmClusterUpdateHistoryEntriesUpdateTypeOsUpdate  ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum = "OS_UPDATE"
)

var mappingListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum = map[string]ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum{
	"GI_UPGRADE": ListCloudVmClusterUpdateHistoryEntriesUpdateTypeGiUpgrade,
	"GI_PATCH":   ListCloudVmClusterUpdateHistoryEntriesUpdateTypeGiPatch,
	"OS_UPDATE":  ListCloudVmClusterUpdateHistoryEntriesUpdateTypeOsUpdate,
}

var mappingListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnumLowerCase = map[string]ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum{
	"gi_upgrade": ListCloudVmClusterUpdateHistoryEntriesUpdateTypeGiUpgrade,
	"gi_patch":   ListCloudVmClusterUpdateHistoryEntriesUpdateTypeGiPatch,
	"os_update":  ListCloudVmClusterUpdateHistoryEntriesUpdateTypeOsUpdate,
}

// GetListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnumValues Enumerates the set of values for ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum
func GetListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnumValues() []ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum {
	values := make([]ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum, 0)
	for _, v := range mappingListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnumStringValues Enumerates the set of values in String for ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum
func GetListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum(val string) (ListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnum, bool) {
	enum, ok := mappingListCloudVmClusterUpdateHistoryEntriesUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
