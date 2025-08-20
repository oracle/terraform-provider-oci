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

// ListBaseccVmClusterUpdatesRequest wrapper for the ListBaseccVmClusterUpdates operation
type ListBaseccVmClusterUpdatesRequest struct {

	// The VM cluster OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	BaseccVmClusterId *string `mandatory:"true" contributesTo:"path" name:"baseccVmClusterId"`

	// A filter to return only resources that match the given update type exactly.
	UpdateType ListBaseccVmClusterUpdatesUpdateTypeEnum `mandatory:"false" contributesTo:"query" name:"updateType" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState BaseccVmClusterUpdateSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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

func (request ListBaseccVmClusterUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBaseccVmClusterUpdatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBaseccVmClusterUpdatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBaseccVmClusterUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBaseccVmClusterUpdatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBaseccVmClusterUpdatesUpdateTypeEnum(string(request.UpdateType)); !ok && request.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", request.UpdateType, strings.Join(GetListBaseccVmClusterUpdatesUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterUpdateSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetBaseccVmClusterUpdateSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBaseccVmClusterUpdatesResponse wrapper for the ListBaseccVmClusterUpdates operation
type ListBaseccVmClusterUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BaseccVmClusterUpdateSummary instances
	Items []BaseccVmClusterUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBaseccVmClusterUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBaseccVmClusterUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBaseccVmClusterUpdatesUpdateTypeEnum Enum with underlying type: string
type ListBaseccVmClusterUpdatesUpdateTypeEnum string

// Set of constants representing the allowable values for ListBaseccVmClusterUpdatesUpdateTypeEnum
const (
	ListBaseccVmClusterUpdatesUpdateTypeGiUpgrade ListBaseccVmClusterUpdatesUpdateTypeEnum = "GI_UPGRADE"
	ListBaseccVmClusterUpdatesUpdateTypeGiPatch   ListBaseccVmClusterUpdatesUpdateTypeEnum = "GI_PATCH"
	ListBaseccVmClusterUpdatesUpdateTypeOsUpdate  ListBaseccVmClusterUpdatesUpdateTypeEnum = "OS_UPDATE"
)

var mappingListBaseccVmClusterUpdatesUpdateTypeEnum = map[string]ListBaseccVmClusterUpdatesUpdateTypeEnum{
	"GI_UPGRADE": ListBaseccVmClusterUpdatesUpdateTypeGiUpgrade,
	"GI_PATCH":   ListBaseccVmClusterUpdatesUpdateTypeGiPatch,
	"OS_UPDATE":  ListBaseccVmClusterUpdatesUpdateTypeOsUpdate,
}

var mappingListBaseccVmClusterUpdatesUpdateTypeEnumLowerCase = map[string]ListBaseccVmClusterUpdatesUpdateTypeEnum{
	"gi_upgrade": ListBaseccVmClusterUpdatesUpdateTypeGiUpgrade,
	"gi_patch":   ListBaseccVmClusterUpdatesUpdateTypeGiPatch,
	"os_update":  ListBaseccVmClusterUpdatesUpdateTypeOsUpdate,
}

// GetListBaseccVmClusterUpdatesUpdateTypeEnumValues Enumerates the set of values for ListBaseccVmClusterUpdatesUpdateTypeEnum
func GetListBaseccVmClusterUpdatesUpdateTypeEnumValues() []ListBaseccVmClusterUpdatesUpdateTypeEnum {
	values := make([]ListBaseccVmClusterUpdatesUpdateTypeEnum, 0)
	for _, v := range mappingListBaseccVmClusterUpdatesUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListBaseccVmClusterUpdatesUpdateTypeEnumStringValues Enumerates the set of values in String for ListBaseccVmClusterUpdatesUpdateTypeEnum
func GetListBaseccVmClusterUpdatesUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingListBaseccVmClusterUpdatesUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBaseccVmClusterUpdatesUpdateTypeEnum(val string) (ListBaseccVmClusterUpdatesUpdateTypeEnum, bool) {
	enum, ok := mappingListBaseccVmClusterUpdatesUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
