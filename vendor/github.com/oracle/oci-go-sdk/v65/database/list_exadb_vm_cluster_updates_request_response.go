// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExadbVmClusterUpdatesRequest wrapper for the ListExadbVmClusterUpdates operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExadbVmClusterUpdates.go.html to see an example of how to use ListExadbVmClusterUpdatesRequest.
type ListExadbVmClusterUpdatesRequest struct {

	// The Exadata VM cluster OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) on Exascale Infrastructure.
	ExadbVmClusterId *string `mandatory:"true" contributesTo:"path" name:"exadbVmClusterId"`

	// A filter to return only resources that match the given update type exactly.
	UpdateType ListExadbVmClusterUpdatesUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

	// A filter to return only resources that match the given update version exactly.
	Version *string `mandatory:"false" contributesTo:"query" name:"version"`

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

func (request ListExadbVmClusterUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExadbVmClusterUpdatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExadbVmClusterUpdatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExadbVmClusterUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExadbVmClusterUpdatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExadbVmClusterUpdatesUpdateTypeEnum(string(request.UpdateType)); !ok && request.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", request.UpdateType, strings.Join(GetListExadbVmClusterUpdatesUpdateTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExadbVmClusterUpdatesResponse wrapper for the ListExadbVmClusterUpdates operation
type ListExadbVmClusterUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExadbVmClusterUpdateSummary instances
	Items []ExadbVmClusterUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExadbVmClusterUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExadbVmClusterUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExadbVmClusterUpdatesUpdateTypeEnum Enum with underlying type: string
type ListExadbVmClusterUpdatesUpdateTypeEnum string

// Set of constants representing the allowable values for ListExadbVmClusterUpdatesUpdateTypeEnum
const (
	ListExadbVmClusterUpdatesUpdateTypeGiUpgrade ListExadbVmClusterUpdatesUpdateTypeEnum = "GI_UPGRADE"
	ListExadbVmClusterUpdatesUpdateTypeGiPatch   ListExadbVmClusterUpdatesUpdateTypeEnum = "GI_PATCH"
	ListExadbVmClusterUpdatesUpdateTypeOsUpdate  ListExadbVmClusterUpdatesUpdateTypeEnum = "OS_UPDATE"
)

var mappingListExadbVmClusterUpdatesUpdateTypeEnum = map[string]ListExadbVmClusterUpdatesUpdateTypeEnum{
	"GI_UPGRADE": ListExadbVmClusterUpdatesUpdateTypeGiUpgrade,
	"GI_PATCH":   ListExadbVmClusterUpdatesUpdateTypeGiPatch,
	"OS_UPDATE":  ListExadbVmClusterUpdatesUpdateTypeOsUpdate,
}

var mappingListExadbVmClusterUpdatesUpdateTypeEnumLowerCase = map[string]ListExadbVmClusterUpdatesUpdateTypeEnum{
	"gi_upgrade": ListExadbVmClusterUpdatesUpdateTypeGiUpgrade,
	"gi_patch":   ListExadbVmClusterUpdatesUpdateTypeGiPatch,
	"os_update":  ListExadbVmClusterUpdatesUpdateTypeOsUpdate,
}

// GetListExadbVmClusterUpdatesUpdateTypeEnumValues Enumerates the set of values for ListExadbVmClusterUpdatesUpdateTypeEnum
func GetListExadbVmClusterUpdatesUpdateTypeEnumValues() []ListExadbVmClusterUpdatesUpdateTypeEnum {
	values := make([]ListExadbVmClusterUpdatesUpdateTypeEnum, 0)
	for _, v := range mappingListExadbVmClusterUpdatesUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListExadbVmClusterUpdatesUpdateTypeEnumStringValues Enumerates the set of values in String for ListExadbVmClusterUpdatesUpdateTypeEnum
func GetListExadbVmClusterUpdatesUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingListExadbVmClusterUpdatesUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExadbVmClusterUpdatesUpdateTypeEnum(val string) (ListExadbVmClusterUpdatesUpdateTypeEnum, bool) {
	enum, ok := mappingListExadbVmClusterUpdatesUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
