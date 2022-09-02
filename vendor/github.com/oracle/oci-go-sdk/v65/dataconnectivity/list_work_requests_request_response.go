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

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataconnectivity/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// The OCID of the compartment containing the resources you want to list.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// DCMS registry ID
	RegistryId *string `mandatory:"false" contributesTo:"query" name:"registryId"`

	// Work request status.
	WorkRequestStatus ListWorkRequestsWorkRequestStatusEnum `mandatory:"false" contributesTo:"query" name:"workRequestStatus" omitEmpty:"true"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestsWorkRequestStatusEnum(string(request.WorkRequestStatus)); !ok && request.WorkRequestStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkRequestStatus: %s. Supported values are: %s.", request.WorkRequestStatus, strings.Join(GetListWorkRequestsWorkRequestStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestSummaryCollection instances
	WorkRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsWorkRequestStatusEnum Enum with underlying type: string
type ListWorkRequestsWorkRequestStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsWorkRequestStatusEnum
const (
	ListWorkRequestsWorkRequestStatusAccepted   ListWorkRequestsWorkRequestStatusEnum = "ACCEPTED"
	ListWorkRequestsWorkRequestStatusInProgress ListWorkRequestsWorkRequestStatusEnum = "IN_PROGRESS"
	ListWorkRequestsWorkRequestStatusFailed     ListWorkRequestsWorkRequestStatusEnum = "FAILED"
	ListWorkRequestsWorkRequestStatusSucceeded  ListWorkRequestsWorkRequestStatusEnum = "SUCCEEDED"
	ListWorkRequestsWorkRequestStatusCanceling  ListWorkRequestsWorkRequestStatusEnum = "CANCELING"
	ListWorkRequestsWorkRequestStatusCanceled   ListWorkRequestsWorkRequestStatusEnum = "CANCELED"
)

var mappingListWorkRequestsWorkRequestStatusEnum = map[string]ListWorkRequestsWorkRequestStatusEnum{
	"ACCEPTED":    ListWorkRequestsWorkRequestStatusAccepted,
	"IN_PROGRESS": ListWorkRequestsWorkRequestStatusInProgress,
	"FAILED":      ListWorkRequestsWorkRequestStatusFailed,
	"SUCCEEDED":   ListWorkRequestsWorkRequestStatusSucceeded,
	"CANCELING":   ListWorkRequestsWorkRequestStatusCanceling,
	"CANCELED":    ListWorkRequestsWorkRequestStatusCanceled,
}

var mappingListWorkRequestsWorkRequestStatusEnumLowerCase = map[string]ListWorkRequestsWorkRequestStatusEnum{
	"accepted":    ListWorkRequestsWorkRequestStatusAccepted,
	"in_progress": ListWorkRequestsWorkRequestStatusInProgress,
	"failed":      ListWorkRequestsWorkRequestStatusFailed,
	"succeeded":   ListWorkRequestsWorkRequestStatusSucceeded,
	"canceling":   ListWorkRequestsWorkRequestStatusCanceling,
	"canceled":    ListWorkRequestsWorkRequestStatusCanceled,
}

// GetListWorkRequestsWorkRequestStatusEnumValues Enumerates the set of values for ListWorkRequestsWorkRequestStatusEnum
func GetListWorkRequestsWorkRequestStatusEnumValues() []ListWorkRequestsWorkRequestStatusEnum {
	values := make([]ListWorkRequestsWorkRequestStatusEnum, 0)
	for _, v := range mappingListWorkRequestsWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsWorkRequestStatusEnumStringValues Enumerates the set of values in String for ListWorkRequestsWorkRequestStatusEnum
func GetListWorkRequestsWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListWorkRequestsWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsWorkRequestStatusEnum(val string) (ListWorkRequestsWorkRequestStatusEnum, bool) {
	enum, ok := mappingListWorkRequestsWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
