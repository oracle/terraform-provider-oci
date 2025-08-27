// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRecipientInvitationsRequest wrapper for the ListRecipientInvitations operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListRecipientInvitations.go.html to see an example of how to use ListRecipientInvitationsRequest.
type ListRecipientInvitationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The tenancy that sent the invitation.
	SenderTenancyId *string `mandatory:"false" contributesTo:"query" name:"senderTenancyId"`

	// The lifecycle state of the resource.
	LifecycleState ListRecipientInvitationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The status of the recipient invitation.
	Status ListRecipientInvitationsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecipientInvitationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecipientInvitationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecipientInvitationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecipientInvitationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRecipientInvitationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRecipientInvitationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRecipientInvitationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRecipientInvitationsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListRecipientInvitationsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRecipientInvitationsResponse wrapper for the ListRecipientInvitations operation
type ListRecipientInvitationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecipientInvitationCollection instances
	RecipientInvitationCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListRecipientInvitationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecipientInvitationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecipientInvitationsLifecycleStateEnum Enum with underlying type: string
type ListRecipientInvitationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListRecipientInvitationsLifecycleStateEnum
const (
	ListRecipientInvitationsLifecycleStateCreating   ListRecipientInvitationsLifecycleStateEnum = "CREATING"
	ListRecipientInvitationsLifecycleStateActive     ListRecipientInvitationsLifecycleStateEnum = "ACTIVE"
	ListRecipientInvitationsLifecycleStateInactive   ListRecipientInvitationsLifecycleStateEnum = "INACTIVE"
	ListRecipientInvitationsLifecycleStateUpdating   ListRecipientInvitationsLifecycleStateEnum = "UPDATING"
	ListRecipientInvitationsLifecycleStateFailed     ListRecipientInvitationsLifecycleStateEnum = "FAILED"
	ListRecipientInvitationsLifecycleStateTerminated ListRecipientInvitationsLifecycleStateEnum = "TERMINATED"
)

var mappingListRecipientInvitationsLifecycleStateEnum = map[string]ListRecipientInvitationsLifecycleStateEnum{
	"CREATING":   ListRecipientInvitationsLifecycleStateCreating,
	"ACTIVE":     ListRecipientInvitationsLifecycleStateActive,
	"INACTIVE":   ListRecipientInvitationsLifecycleStateInactive,
	"UPDATING":   ListRecipientInvitationsLifecycleStateUpdating,
	"FAILED":     ListRecipientInvitationsLifecycleStateFailed,
	"TERMINATED": ListRecipientInvitationsLifecycleStateTerminated,
}

var mappingListRecipientInvitationsLifecycleStateEnumLowerCase = map[string]ListRecipientInvitationsLifecycleStateEnum{
	"creating":   ListRecipientInvitationsLifecycleStateCreating,
	"active":     ListRecipientInvitationsLifecycleStateActive,
	"inactive":   ListRecipientInvitationsLifecycleStateInactive,
	"updating":   ListRecipientInvitationsLifecycleStateUpdating,
	"failed":     ListRecipientInvitationsLifecycleStateFailed,
	"terminated": ListRecipientInvitationsLifecycleStateTerminated,
}

// GetListRecipientInvitationsLifecycleStateEnumValues Enumerates the set of values for ListRecipientInvitationsLifecycleStateEnum
func GetListRecipientInvitationsLifecycleStateEnumValues() []ListRecipientInvitationsLifecycleStateEnum {
	values := make([]ListRecipientInvitationsLifecycleStateEnum, 0)
	for _, v := range mappingListRecipientInvitationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecipientInvitationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListRecipientInvitationsLifecycleStateEnum
func GetListRecipientInvitationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"FAILED",
		"TERMINATED",
	}
}

// GetMappingListRecipientInvitationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecipientInvitationsLifecycleStateEnum(val string) (ListRecipientInvitationsLifecycleStateEnum, bool) {
	enum, ok := mappingListRecipientInvitationsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRecipientInvitationsStatusEnum Enum with underlying type: string
type ListRecipientInvitationsStatusEnum string

// Set of constants representing the allowable values for ListRecipientInvitationsStatusEnum
const (
	ListRecipientInvitationsStatusPending  ListRecipientInvitationsStatusEnum = "PENDING"
	ListRecipientInvitationsStatusCanceled ListRecipientInvitationsStatusEnum = "CANCELED"
	ListRecipientInvitationsStatusAccepted ListRecipientInvitationsStatusEnum = "ACCEPTED"
	ListRecipientInvitationsStatusIgnored  ListRecipientInvitationsStatusEnum = "IGNORED"
	ListRecipientInvitationsStatusExpired  ListRecipientInvitationsStatusEnum = "EXPIRED"
	ListRecipientInvitationsStatusFailed   ListRecipientInvitationsStatusEnum = "FAILED"
)

var mappingListRecipientInvitationsStatusEnum = map[string]ListRecipientInvitationsStatusEnum{
	"PENDING":  ListRecipientInvitationsStatusPending,
	"CANCELED": ListRecipientInvitationsStatusCanceled,
	"ACCEPTED": ListRecipientInvitationsStatusAccepted,
	"IGNORED":  ListRecipientInvitationsStatusIgnored,
	"EXPIRED":  ListRecipientInvitationsStatusExpired,
	"FAILED":   ListRecipientInvitationsStatusFailed,
}

var mappingListRecipientInvitationsStatusEnumLowerCase = map[string]ListRecipientInvitationsStatusEnum{
	"pending":  ListRecipientInvitationsStatusPending,
	"canceled": ListRecipientInvitationsStatusCanceled,
	"accepted": ListRecipientInvitationsStatusAccepted,
	"ignored":  ListRecipientInvitationsStatusIgnored,
	"expired":  ListRecipientInvitationsStatusExpired,
	"failed":   ListRecipientInvitationsStatusFailed,
}

// GetListRecipientInvitationsStatusEnumValues Enumerates the set of values for ListRecipientInvitationsStatusEnum
func GetListRecipientInvitationsStatusEnumValues() []ListRecipientInvitationsStatusEnum {
	values := make([]ListRecipientInvitationsStatusEnum, 0)
	for _, v := range mappingListRecipientInvitationsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListRecipientInvitationsStatusEnumStringValues Enumerates the set of values in String for ListRecipientInvitationsStatusEnum
func GetListRecipientInvitationsStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"CANCELED",
		"ACCEPTED",
		"IGNORED",
		"EXPIRED",
		"FAILED",
	}
}

// GetMappingListRecipientInvitationsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRecipientInvitationsStatusEnum(val string) (ListRecipientInvitationsStatusEnum, bool) {
	enum, ok := mappingListRecipientInvitationsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
