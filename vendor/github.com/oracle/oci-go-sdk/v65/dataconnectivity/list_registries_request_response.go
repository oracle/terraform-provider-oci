// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRegistriesRequest wrapper for the ListRegistries operation
type ListRegistriesRequest struct {

	// The OCID of the compartment containing the resources you want to list.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter allows list registries to deep look at the whole tenancy.
	IsDeepLookup *bool `mandatory:"false" contributesTo:"query" name:"isDeepLookup"`

	// Lifecycle state of the resource.
	LifecycleState ListRegistriesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRegistriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRegistriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRegistriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRegistriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRegistriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRegistriesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListRegistriesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRegistriesResponse wrapper for the ListRegistries operation
type ListRegistriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RegistrySummaryCollection instances
	RegistrySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRegistriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRegistriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRegistriesLifecycleStateEnum Enum with underlying type: string
type ListRegistriesLifecycleStateEnum string

// Set of constants representing the allowable values for ListRegistriesLifecycleStateEnum
const (
	ListRegistriesLifecycleStateCreating ListRegistriesLifecycleStateEnum = "CREATING"
	ListRegistriesLifecycleStateActive   ListRegistriesLifecycleStateEnum = "ACTIVE"
	ListRegistriesLifecycleStateInactive ListRegistriesLifecycleStateEnum = "INACTIVE"
	ListRegistriesLifecycleStateUpdating ListRegistriesLifecycleStateEnum = "UPDATING"
	ListRegistriesLifecycleStateDeleting ListRegistriesLifecycleStateEnum = "DELETING"
	ListRegistriesLifecycleStateDeleted  ListRegistriesLifecycleStateEnum = "DELETED"
	ListRegistriesLifecycleStateFailed   ListRegistriesLifecycleStateEnum = "FAILED"
	ListRegistriesLifecycleStateStarting ListRegistriesLifecycleStateEnum = "STARTING"
	ListRegistriesLifecycleStateStopping ListRegistriesLifecycleStateEnum = "STOPPING"
	ListRegistriesLifecycleStateStopped  ListRegistriesLifecycleStateEnum = "STOPPED"
)

var mappingListRegistriesLifecycleStateEnum = map[string]ListRegistriesLifecycleStateEnum{
	"CREATING": ListRegistriesLifecycleStateCreating,
	"ACTIVE":   ListRegistriesLifecycleStateActive,
	"INACTIVE": ListRegistriesLifecycleStateInactive,
	"UPDATING": ListRegistriesLifecycleStateUpdating,
	"DELETING": ListRegistriesLifecycleStateDeleting,
	"DELETED":  ListRegistriesLifecycleStateDeleted,
	"FAILED":   ListRegistriesLifecycleStateFailed,
	"STARTING": ListRegistriesLifecycleStateStarting,
	"STOPPING": ListRegistriesLifecycleStateStopping,
	"STOPPED":  ListRegistriesLifecycleStateStopped,
}

var mappingListRegistriesLifecycleStateEnumLowerCase = map[string]ListRegistriesLifecycleStateEnum{
	"creating": ListRegistriesLifecycleStateCreating,
	"active":   ListRegistriesLifecycleStateActive,
	"inactive": ListRegistriesLifecycleStateInactive,
	"updating": ListRegistriesLifecycleStateUpdating,
	"deleting": ListRegistriesLifecycleStateDeleting,
	"deleted":  ListRegistriesLifecycleStateDeleted,
	"failed":   ListRegistriesLifecycleStateFailed,
	"starting": ListRegistriesLifecycleStateStarting,
	"stopping": ListRegistriesLifecycleStateStopping,
	"stopped":  ListRegistriesLifecycleStateStopped,
}

// GetListRegistriesLifecycleStateEnumValues Enumerates the set of values for ListRegistriesLifecycleStateEnum
func GetListRegistriesLifecycleStateEnumValues() []ListRegistriesLifecycleStateEnum {
	values := make([]ListRegistriesLifecycleStateEnum, 0)
	for _, v := range mappingListRegistriesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListRegistriesLifecycleStateEnumStringValues Enumerates the set of values in String for ListRegistriesLifecycleStateEnum
func GetListRegistriesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"STARTING",
		"STOPPING",
		"STOPPED",
	}
}

// GetMappingListRegistriesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRegistriesLifecycleStateEnum(val string) (ListRegistriesLifecycleStateEnum, bool) {
	enum, ok := mappingListRegistriesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
