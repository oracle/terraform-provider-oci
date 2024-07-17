// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInternalInternetGatewaysRequest wrapper for the ListInternalInternetGateways operation
type ListInternalInternetGatewaysRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListInternalInternetGatewaysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListInternalInternetGatewaysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle
	// state. The state value is case-insensitive.
	LifecycleState InternalInternetGatewayLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalInternetGatewaysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalInternetGatewaysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalInternetGatewaysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListInternalInternetGatewaysRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["compartmentId"] != nil {
		templateParam := mandatoryParamMap["compartmentId"]
		for _, template := range templateParam {
			replacementParam := *request.CompartmentId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalInternetGatewaysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalInternetGatewaysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalInternetGatewaysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalInternetGatewaysSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalInternetGatewaysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalInternetGatewaysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInternalInternetGatewayLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetInternalInternetGatewayLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalInternetGatewaysResponse wrapper for the ListInternalInternetGateways operation
type ListInternalInternetGatewaysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []InternalInternetGateway instances
	Items []InternalInternetGateway `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListInternalInternetGatewaysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalInternetGatewaysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalInternetGatewaysSortByEnum Enum with underlying type: string
type ListInternalInternetGatewaysSortByEnum string

// Set of constants representing the allowable values for ListInternalInternetGatewaysSortByEnum
const (
	ListInternalInternetGatewaysSortByTimecreated ListInternalInternetGatewaysSortByEnum = "TIMECREATED"
	ListInternalInternetGatewaysSortByDisplayname ListInternalInternetGatewaysSortByEnum = "DISPLAYNAME"
)

var mappingListInternalInternetGatewaysSortByEnum = map[string]ListInternalInternetGatewaysSortByEnum{
	"TIMECREATED": ListInternalInternetGatewaysSortByTimecreated,
	"DISPLAYNAME": ListInternalInternetGatewaysSortByDisplayname,
}

var mappingListInternalInternetGatewaysSortByEnumLowerCase = map[string]ListInternalInternetGatewaysSortByEnum{
	"timecreated": ListInternalInternetGatewaysSortByTimecreated,
	"displayname": ListInternalInternetGatewaysSortByDisplayname,
}

// GetListInternalInternetGatewaysSortByEnumValues Enumerates the set of values for ListInternalInternetGatewaysSortByEnum
func GetListInternalInternetGatewaysSortByEnumValues() []ListInternalInternetGatewaysSortByEnum {
	values := make([]ListInternalInternetGatewaysSortByEnum, 0)
	for _, v := range mappingListInternalInternetGatewaysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalInternetGatewaysSortByEnumStringValues Enumerates the set of values in String for ListInternalInternetGatewaysSortByEnum
func GetListInternalInternetGatewaysSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListInternalInternetGatewaysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalInternetGatewaysSortByEnum(val string) (ListInternalInternetGatewaysSortByEnum, bool) {
	enum, ok := mappingListInternalInternetGatewaysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalInternetGatewaysSortOrderEnum Enum with underlying type: string
type ListInternalInternetGatewaysSortOrderEnum string

// Set of constants representing the allowable values for ListInternalInternetGatewaysSortOrderEnum
const (
	ListInternalInternetGatewaysSortOrderAsc  ListInternalInternetGatewaysSortOrderEnum = "ASC"
	ListInternalInternetGatewaysSortOrderDesc ListInternalInternetGatewaysSortOrderEnum = "DESC"
)

var mappingListInternalInternetGatewaysSortOrderEnum = map[string]ListInternalInternetGatewaysSortOrderEnum{
	"ASC":  ListInternalInternetGatewaysSortOrderAsc,
	"DESC": ListInternalInternetGatewaysSortOrderDesc,
}

var mappingListInternalInternetGatewaysSortOrderEnumLowerCase = map[string]ListInternalInternetGatewaysSortOrderEnum{
	"asc":  ListInternalInternetGatewaysSortOrderAsc,
	"desc": ListInternalInternetGatewaysSortOrderDesc,
}

// GetListInternalInternetGatewaysSortOrderEnumValues Enumerates the set of values for ListInternalInternetGatewaysSortOrderEnum
func GetListInternalInternetGatewaysSortOrderEnumValues() []ListInternalInternetGatewaysSortOrderEnum {
	values := make([]ListInternalInternetGatewaysSortOrderEnum, 0)
	for _, v := range mappingListInternalInternetGatewaysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalInternetGatewaysSortOrderEnumStringValues Enumerates the set of values in String for ListInternalInternetGatewaysSortOrderEnum
func GetListInternalInternetGatewaysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalInternetGatewaysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalInternetGatewaysSortOrderEnum(val string) (ListInternalInternetGatewaysSortOrderEnum, bool) {
	enum, ok := mappingListInternalInternetGatewaysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
