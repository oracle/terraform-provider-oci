// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBulkActionResourceTypesRequest wrapper for the ListBulkActionResourceTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListBulkActionResourceTypes.go.html to see an example of how to use ListBulkActionResourceTypesRequest.
type ListBulkActionResourceTypesRequest struct {

	// The type of bulk action.
	BulkActionType ListBulkActionResourceTypesBulkActionTypeEnum `mandatory:"true" contributesTo:"query" name:"bulkActionType" omitEmpty:"true"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBulkActionResourceTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBulkActionResourceTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBulkActionResourceTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBulkActionResourceTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBulkActionResourceTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBulkActionResourceTypesBulkActionTypeEnum(string(request.BulkActionType)); !ok && request.BulkActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BulkActionType: %s. Supported values are: %s.", request.BulkActionType, strings.Join(GetListBulkActionResourceTypesBulkActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBulkActionResourceTypesResponse wrapper for the ListBulkActionResourceTypes operation
type ListBulkActionResourceTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BulkActionResourceTypeCollection instances
	BulkActionResourceTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBulkActionResourceTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBulkActionResourceTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBulkActionResourceTypesBulkActionTypeEnum Enum with underlying type: string
type ListBulkActionResourceTypesBulkActionTypeEnum string

// Set of constants representing the allowable values for ListBulkActionResourceTypesBulkActionTypeEnum
const (
	ListBulkActionResourceTypesBulkActionTypeMoveResources   ListBulkActionResourceTypesBulkActionTypeEnum = "BULK_MOVE_RESOURCES"
	ListBulkActionResourceTypesBulkActionTypeDeleteResources ListBulkActionResourceTypesBulkActionTypeEnum = "BULK_DELETE_RESOURCES"
)

var mappingListBulkActionResourceTypesBulkActionTypeEnum = map[string]ListBulkActionResourceTypesBulkActionTypeEnum{
	"BULK_MOVE_RESOURCES":   ListBulkActionResourceTypesBulkActionTypeMoveResources,
	"BULK_DELETE_RESOURCES": ListBulkActionResourceTypesBulkActionTypeDeleteResources,
}

var mappingListBulkActionResourceTypesBulkActionTypeEnumLowerCase = map[string]ListBulkActionResourceTypesBulkActionTypeEnum{
	"bulk_move_resources":   ListBulkActionResourceTypesBulkActionTypeMoveResources,
	"bulk_delete_resources": ListBulkActionResourceTypesBulkActionTypeDeleteResources,
}

// GetListBulkActionResourceTypesBulkActionTypeEnumValues Enumerates the set of values for ListBulkActionResourceTypesBulkActionTypeEnum
func GetListBulkActionResourceTypesBulkActionTypeEnumValues() []ListBulkActionResourceTypesBulkActionTypeEnum {
	values := make([]ListBulkActionResourceTypesBulkActionTypeEnum, 0)
	for _, v := range mappingListBulkActionResourceTypesBulkActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListBulkActionResourceTypesBulkActionTypeEnumStringValues Enumerates the set of values in String for ListBulkActionResourceTypesBulkActionTypeEnum
func GetListBulkActionResourceTypesBulkActionTypeEnumStringValues() []string {
	return []string{
		"BULK_MOVE_RESOURCES",
		"BULK_DELETE_RESOURCES",
	}
}

// GetMappingListBulkActionResourceTypesBulkActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBulkActionResourceTypesBulkActionTypeEnum(val string) (ListBulkActionResourceTypesBulkActionTypeEnum, bool) {
	enum, ok := mappingListBulkActionResourceTypesBulkActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
