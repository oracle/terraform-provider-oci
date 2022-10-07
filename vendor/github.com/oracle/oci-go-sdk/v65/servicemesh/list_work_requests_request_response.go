// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
type ListWorkRequestsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// A filter to return work requests that match the given resourceId.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only resources that match the operation status given.
	OperationStatus ListWorkRequestsOperationStatusEnum `mandatory:"false" contributesTo:"query" name:"operationStatus" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
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
	if _, ok := GetMappingListWorkRequestsOperationStatusEnum(string(request.OperationStatus)); !ok && request.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", request.OperationStatus, strings.Join(GetListWorkRequestsOperationStatusEnumStringValues(), ",")))
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

	// A list of WorkRequestCollection instances
	WorkRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsOperationStatusEnum Enum with underlying type: string
type ListWorkRequestsOperationStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsOperationStatusEnum
const (
	ListWorkRequestsOperationStatusAccepted   ListWorkRequestsOperationStatusEnum = "ACCEPTED"
	ListWorkRequestsOperationStatusInProgress ListWorkRequestsOperationStatusEnum = "IN_PROGRESS"
	ListWorkRequestsOperationStatusFailed     ListWorkRequestsOperationStatusEnum = "FAILED"
	ListWorkRequestsOperationStatusSucceeded  ListWorkRequestsOperationStatusEnum = "SUCCEEDED"
	ListWorkRequestsOperationStatusWaiting    ListWorkRequestsOperationStatusEnum = "WAITING"
	ListWorkRequestsOperationStatusCanceling  ListWorkRequestsOperationStatusEnum = "CANCELING"
	ListWorkRequestsOperationStatusCanceled   ListWorkRequestsOperationStatusEnum = "CANCELED"
)

var mappingListWorkRequestsOperationStatusEnum = map[string]ListWorkRequestsOperationStatusEnum{
	"ACCEPTED":    ListWorkRequestsOperationStatusAccepted,
	"IN_PROGRESS": ListWorkRequestsOperationStatusInProgress,
	"FAILED":      ListWorkRequestsOperationStatusFailed,
	"SUCCEEDED":   ListWorkRequestsOperationStatusSucceeded,
	"WAITING":     ListWorkRequestsOperationStatusWaiting,
	"CANCELING":   ListWorkRequestsOperationStatusCanceling,
	"CANCELED":    ListWorkRequestsOperationStatusCanceled,
}

var mappingListWorkRequestsOperationStatusEnumLowerCase = map[string]ListWorkRequestsOperationStatusEnum{
	"accepted":    ListWorkRequestsOperationStatusAccepted,
	"in_progress": ListWorkRequestsOperationStatusInProgress,
	"failed":      ListWorkRequestsOperationStatusFailed,
	"succeeded":   ListWorkRequestsOperationStatusSucceeded,
	"waiting":     ListWorkRequestsOperationStatusWaiting,
	"canceling":   ListWorkRequestsOperationStatusCanceling,
	"canceled":    ListWorkRequestsOperationStatusCanceled,
}

// GetListWorkRequestsOperationStatusEnumValues Enumerates the set of values for ListWorkRequestsOperationStatusEnum
func GetListWorkRequestsOperationStatusEnumValues() []ListWorkRequestsOperationStatusEnum {
	values := make([]ListWorkRequestsOperationStatusEnum, 0)
	for _, v := range mappingListWorkRequestsOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsOperationStatusEnumStringValues Enumerates the set of values in String for ListWorkRequestsOperationStatusEnum
func GetListWorkRequestsOperationStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"WAITING",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingListWorkRequestsOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsOperationStatusEnum(val string) (ListWorkRequestsOperationStatusEnum, bool) {
	enum, ok := mappingListWorkRequestsOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
