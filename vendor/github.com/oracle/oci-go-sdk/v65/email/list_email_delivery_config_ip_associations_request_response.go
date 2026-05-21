// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEmailDeliveryConfigIpAssociationsRequest wrapper for the ListEmailDeliveryConfigIpAssociations operation
type ListEmailDeliveryConfigIpAssociationsRequest struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value of the opc-next-page response header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. `1` is the minimum, `1000` is the maximum. For important details about
	// how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either ascending or descending order.
	SortOrder ListEmailDeliveryConfigIpAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the email delivery config exactly.
	EmailDeliveryConfigId *string `mandatory:"false" contributesTo:"query" name:"emailDeliveryConfigId"`

	// A filter to only return resources that match the public ip exactly.
	OutboundIp *string `mandatory:"false" contributesTo:"query" name:"outboundIp"`

	// A filter to only return resources that match the given id exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Filter returned list by specified lifecycle state. This parameter is case-insensitive.
	LifecycleState EmailDeliveryConfigIpAssociationLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies the attribute with which to sort the email delivery configuration ip associations.
	// Default: `timeCreated`
	// * **timeCreated:** Sorts by timeCreated.
	SortBy ListEmailDeliveryConfigIpAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEmailDeliveryConfigIpAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEmailDeliveryConfigIpAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEmailDeliveryConfigIpAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEmailDeliveryConfigIpAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEmailDeliveryConfigIpAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEmailDeliveryConfigIpAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEmailDeliveryConfigIpAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmailDeliveryConfigIpAssociationLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetEmailDeliveryConfigIpAssociationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEmailDeliveryConfigIpAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEmailDeliveryConfigIpAssociationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEmailDeliveryConfigIpAssociationsResponse wrapper for the ListEmailDeliveryConfigIpAssociations operation
type ListEmailDeliveryConfigIpAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EmailDeliveryConfigIpAssociationCollection instances
	EmailDeliveryConfigIpAssociationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results remain.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListEmailDeliveryConfigIpAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEmailDeliveryConfigIpAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEmailDeliveryConfigIpAssociationsSortOrderEnum Enum with underlying type: string
type ListEmailDeliveryConfigIpAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListEmailDeliveryConfigIpAssociationsSortOrderEnum
const (
	ListEmailDeliveryConfigIpAssociationsSortOrderAsc  ListEmailDeliveryConfigIpAssociationsSortOrderEnum = "ASC"
	ListEmailDeliveryConfigIpAssociationsSortOrderDesc ListEmailDeliveryConfigIpAssociationsSortOrderEnum = "DESC"
)

var mappingListEmailDeliveryConfigIpAssociationsSortOrderEnum = map[string]ListEmailDeliveryConfigIpAssociationsSortOrderEnum{
	"ASC":  ListEmailDeliveryConfigIpAssociationsSortOrderAsc,
	"DESC": ListEmailDeliveryConfigIpAssociationsSortOrderDesc,
}

var mappingListEmailDeliveryConfigIpAssociationsSortOrderEnumLowerCase = map[string]ListEmailDeliveryConfigIpAssociationsSortOrderEnum{
	"asc":  ListEmailDeliveryConfigIpAssociationsSortOrderAsc,
	"desc": ListEmailDeliveryConfigIpAssociationsSortOrderDesc,
}

// GetListEmailDeliveryConfigIpAssociationsSortOrderEnumValues Enumerates the set of values for ListEmailDeliveryConfigIpAssociationsSortOrderEnum
func GetListEmailDeliveryConfigIpAssociationsSortOrderEnumValues() []ListEmailDeliveryConfigIpAssociationsSortOrderEnum {
	values := make([]ListEmailDeliveryConfigIpAssociationsSortOrderEnum, 0)
	for _, v := range mappingListEmailDeliveryConfigIpAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailDeliveryConfigIpAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListEmailDeliveryConfigIpAssociationsSortOrderEnum
func GetListEmailDeliveryConfigIpAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEmailDeliveryConfigIpAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailDeliveryConfigIpAssociationsSortOrderEnum(val string) (ListEmailDeliveryConfigIpAssociationsSortOrderEnum, bool) {
	enum, ok := mappingListEmailDeliveryConfigIpAssociationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEmailDeliveryConfigIpAssociationsSortByEnum Enum with underlying type: string
type ListEmailDeliveryConfigIpAssociationsSortByEnum string

// Set of constants representing the allowable values for ListEmailDeliveryConfigIpAssociationsSortByEnum
const (
	ListEmailDeliveryConfigIpAssociationsSortByTimecreated ListEmailDeliveryConfigIpAssociationsSortByEnum = "timeCreated"
)

var mappingListEmailDeliveryConfigIpAssociationsSortByEnum = map[string]ListEmailDeliveryConfigIpAssociationsSortByEnum{
	"timeCreated": ListEmailDeliveryConfigIpAssociationsSortByTimecreated,
}

var mappingListEmailDeliveryConfigIpAssociationsSortByEnumLowerCase = map[string]ListEmailDeliveryConfigIpAssociationsSortByEnum{
	"timecreated": ListEmailDeliveryConfigIpAssociationsSortByTimecreated,
}

// GetListEmailDeliveryConfigIpAssociationsSortByEnumValues Enumerates the set of values for ListEmailDeliveryConfigIpAssociationsSortByEnum
func GetListEmailDeliveryConfigIpAssociationsSortByEnumValues() []ListEmailDeliveryConfigIpAssociationsSortByEnum {
	values := make([]ListEmailDeliveryConfigIpAssociationsSortByEnum, 0)
	for _, v := range mappingListEmailDeliveryConfigIpAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEmailDeliveryConfigIpAssociationsSortByEnumStringValues Enumerates the set of values in String for ListEmailDeliveryConfigIpAssociationsSortByEnum
func GetListEmailDeliveryConfigIpAssociationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListEmailDeliveryConfigIpAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEmailDeliveryConfigIpAssociationsSortByEnum(val string) (ListEmailDeliveryConfigIpAssociationsSortByEnum, bool) {
	enum, ok := mappingListEmailDeliveryConfigIpAssociationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
