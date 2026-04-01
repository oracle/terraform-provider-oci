// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package self

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubscriptionsRequest wrapper for the ListSubscriptions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/self/ListSubscriptions.go.html to see an example of how to use ListSubscriptionsRequest.
type ListSubscriptionsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleDetails ListSubscriptionsLifecycleDetailsEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetails" omitEmpty:"true"`

	// A filter to return only resources that match the given name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Subscription.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubscriptionsLifecycleDetailsEnum(string(request.LifecycleDetails)); !ok && request.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", request.LifecycleDetails, strings.Join(GetListSubscriptionsLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubscriptionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubscriptionsResponse wrapper for the ListSubscriptions operation
type ListSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubscriptionCollection instances
	SubscriptionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubscriptionsLifecycleDetailsEnum Enum with underlying type: string
type ListSubscriptionsLifecycleDetailsEnum string

// Set of constants representing the allowable values for ListSubscriptionsLifecycleDetailsEnum
const (
	ListSubscriptionsLifecycleDetailsCreated               ListSubscriptionsLifecycleDetailsEnum = "CREATED"
	ListSubscriptionsLifecycleDetailsPendingActivation     ListSubscriptionsLifecycleDetailsEnum = "PENDING_ACTIVATION"
	ListSubscriptionsLifecycleDetailsProvisioningStarted   ListSubscriptionsLifecycleDetailsEnum = "PROVISIONING_STARTED"
	ListSubscriptionsLifecycleDetailsProvisioningCompleted ListSubscriptionsLifecycleDetailsEnum = "PROVISIONING_COMPLETED"
	ListSubscriptionsLifecycleDetailsProvisioningFailed    ListSubscriptionsLifecycleDetailsEnum = "PROVISIONING_FAILED"
	ListSubscriptionsLifecycleDetailsActive                ListSubscriptionsLifecycleDetailsEnum = "ACTIVE"
	ListSubscriptionsLifecycleDetailsExpired               ListSubscriptionsLifecycleDetailsEnum = "EXPIRED"
	ListSubscriptionsLifecycleDetailsTerminated            ListSubscriptionsLifecycleDetailsEnum = "TERMINATED"
	ListSubscriptionsLifecycleDetailsFailed                ListSubscriptionsLifecycleDetailsEnum = "FAILED"
	ListSubscriptionsLifecycleDetailsDeleting              ListSubscriptionsLifecycleDetailsEnum = "DELETING"
	ListSubscriptionsLifecycleDetailsUpdating              ListSubscriptionsLifecycleDetailsEnum = "UPDATING"
	ListSubscriptionsLifecycleDetailsDeleted               ListSubscriptionsLifecycleDetailsEnum = "DELETED"
)

var mappingListSubscriptionsLifecycleDetailsEnum = map[string]ListSubscriptionsLifecycleDetailsEnum{
	"CREATED":                ListSubscriptionsLifecycleDetailsCreated,
	"PENDING_ACTIVATION":     ListSubscriptionsLifecycleDetailsPendingActivation,
	"PROVISIONING_STARTED":   ListSubscriptionsLifecycleDetailsProvisioningStarted,
	"PROVISIONING_COMPLETED": ListSubscriptionsLifecycleDetailsProvisioningCompleted,
	"PROVISIONING_FAILED":    ListSubscriptionsLifecycleDetailsProvisioningFailed,
	"ACTIVE":                 ListSubscriptionsLifecycleDetailsActive,
	"EXPIRED":                ListSubscriptionsLifecycleDetailsExpired,
	"TERMINATED":             ListSubscriptionsLifecycleDetailsTerminated,
	"FAILED":                 ListSubscriptionsLifecycleDetailsFailed,
	"DELETING":               ListSubscriptionsLifecycleDetailsDeleting,
	"UPDATING":               ListSubscriptionsLifecycleDetailsUpdating,
	"DELETED":                ListSubscriptionsLifecycleDetailsDeleted,
}

var mappingListSubscriptionsLifecycleDetailsEnumLowerCase = map[string]ListSubscriptionsLifecycleDetailsEnum{
	"created":                ListSubscriptionsLifecycleDetailsCreated,
	"pending_activation":     ListSubscriptionsLifecycleDetailsPendingActivation,
	"provisioning_started":   ListSubscriptionsLifecycleDetailsProvisioningStarted,
	"provisioning_completed": ListSubscriptionsLifecycleDetailsProvisioningCompleted,
	"provisioning_failed":    ListSubscriptionsLifecycleDetailsProvisioningFailed,
	"active":                 ListSubscriptionsLifecycleDetailsActive,
	"expired":                ListSubscriptionsLifecycleDetailsExpired,
	"terminated":             ListSubscriptionsLifecycleDetailsTerminated,
	"failed":                 ListSubscriptionsLifecycleDetailsFailed,
	"deleting":               ListSubscriptionsLifecycleDetailsDeleting,
	"updating":               ListSubscriptionsLifecycleDetailsUpdating,
	"deleted":                ListSubscriptionsLifecycleDetailsDeleted,
}

// GetListSubscriptionsLifecycleDetailsEnumValues Enumerates the set of values for ListSubscriptionsLifecycleDetailsEnum
func GetListSubscriptionsLifecycleDetailsEnumValues() []ListSubscriptionsLifecycleDetailsEnum {
	values := make([]ListSubscriptionsLifecycleDetailsEnum, 0)
	for _, v := range mappingListSubscriptionsLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionsLifecycleDetailsEnumStringValues Enumerates the set of values in String for ListSubscriptionsLifecycleDetailsEnum
func GetListSubscriptionsLifecycleDetailsEnumStringValues() []string {
	return []string{
		"CREATED",
		"PENDING_ACTIVATION",
		"PROVISIONING_STARTED",
		"PROVISIONING_COMPLETED",
		"PROVISIONING_FAILED",
		"ACTIVE",
		"EXPIRED",
		"TERMINATED",
		"FAILED",
		"DELETING",
		"UPDATING",
		"DELETED",
	}
}

// GetMappingListSubscriptionsLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionsLifecycleDetailsEnum(val string) (ListSubscriptionsLifecycleDetailsEnum, bool) {
	enum, ok := mappingListSubscriptionsLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubscriptionsSortOrderEnum Enum with underlying type: string
type ListSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListSubscriptionsSortOrderEnum
const (
	ListSubscriptionsSortOrderAsc  ListSubscriptionsSortOrderEnum = "ASC"
	ListSubscriptionsSortOrderDesc ListSubscriptionsSortOrderEnum = "DESC"
)

var mappingListSubscriptionsSortOrderEnum = map[string]ListSubscriptionsSortOrderEnum{
	"ASC":  ListSubscriptionsSortOrderAsc,
	"DESC": ListSubscriptionsSortOrderDesc,
}

var mappingListSubscriptionsSortOrderEnumLowerCase = map[string]ListSubscriptionsSortOrderEnum{
	"asc":  ListSubscriptionsSortOrderAsc,
	"desc": ListSubscriptionsSortOrderDesc,
}

// GetListSubscriptionsSortOrderEnumValues Enumerates the set of values for ListSubscriptionsSortOrderEnum
func GetListSubscriptionsSortOrderEnumValues() []ListSubscriptionsSortOrderEnum {
	values := make([]ListSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListSubscriptionsSortOrderEnum
func GetListSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionsSortOrderEnum(val string) (ListSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubscriptionsSortByEnum Enum with underlying type: string
type ListSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListSubscriptionsSortByEnum
const (
	ListSubscriptionsSortByTimecreated ListSubscriptionsSortByEnum = "timeCreated"
	ListSubscriptionsSortByDisplayname ListSubscriptionsSortByEnum = "displayName"
	ListSubscriptionsSortBySelftokenid ListSubscriptionsSortByEnum = "selfTokenId"
	ListSubscriptionsSortByProductid   ListSubscriptionsSortByEnum = "productId"
)

var mappingListSubscriptionsSortByEnum = map[string]ListSubscriptionsSortByEnum{
	"timeCreated": ListSubscriptionsSortByTimecreated,
	"displayName": ListSubscriptionsSortByDisplayname,
	"selfTokenId": ListSubscriptionsSortBySelftokenid,
	"productId":   ListSubscriptionsSortByProductid,
}

var mappingListSubscriptionsSortByEnumLowerCase = map[string]ListSubscriptionsSortByEnum{
	"timecreated": ListSubscriptionsSortByTimecreated,
	"displayname": ListSubscriptionsSortByDisplayname,
	"selftokenid": ListSubscriptionsSortBySelftokenid,
	"productid":   ListSubscriptionsSortByProductid,
}

// GetListSubscriptionsSortByEnumValues Enumerates the set of values for ListSubscriptionsSortByEnum
func GetListSubscriptionsSortByEnumValues() []ListSubscriptionsSortByEnum {
	values := make([]ListSubscriptionsSortByEnum, 0)
	for _, v := range mappingListSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListSubscriptionsSortByEnum
func GetListSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"selfTokenId",
		"productId",
	}
}

// GetMappingListSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionsSortByEnum(val string) (ListSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
