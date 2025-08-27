// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetMultiCloudResourceDiscoveryRequest wrapper for the GetMultiCloudResourceDiscovery operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetMultiCloudResourceDiscovery.go.html to see an example of how to use GetMultiCloudResourceDiscoveryRequest.
type GetMultiCloudResourceDiscoveryRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multi Cloud Discovery Resource.
	MultiCloudResourceDiscoveryId *string `mandatory:"true" contributesTo:"path" name:"multiCloudResourceDiscoveryId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder GetMultiCloudResourceDiscoverySortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetMultiCloudResourceDiscoveryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetMultiCloudResourceDiscoveryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetMultiCloudResourceDiscoveryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetMultiCloudResourceDiscoveryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetMultiCloudResourceDiscoveryRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetMultiCloudResourceDiscoverySortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetMultiCloudResourceDiscoverySortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetMultiCloudResourceDiscoveryResponse wrapper for the GetMultiCloudResourceDiscovery operation
type GetMultiCloudResourceDiscoveryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MultiCloudResourceDiscovery instances
	MultiCloudResourceDiscovery `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetMultiCloudResourceDiscoveryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetMultiCloudResourceDiscoveryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetMultiCloudResourceDiscoverySortOrderEnum Enum with underlying type: string
type GetMultiCloudResourceDiscoverySortOrderEnum string

// Set of constants representing the allowable values for GetMultiCloudResourceDiscoverySortOrderEnum
const (
	GetMultiCloudResourceDiscoverySortOrderAsc  GetMultiCloudResourceDiscoverySortOrderEnum = "ASC"
	GetMultiCloudResourceDiscoverySortOrderDesc GetMultiCloudResourceDiscoverySortOrderEnum = "DESC"
)

var mappingGetMultiCloudResourceDiscoverySortOrderEnum = map[string]GetMultiCloudResourceDiscoverySortOrderEnum{
	"ASC":  GetMultiCloudResourceDiscoverySortOrderAsc,
	"DESC": GetMultiCloudResourceDiscoverySortOrderDesc,
}

var mappingGetMultiCloudResourceDiscoverySortOrderEnumLowerCase = map[string]GetMultiCloudResourceDiscoverySortOrderEnum{
	"asc":  GetMultiCloudResourceDiscoverySortOrderAsc,
	"desc": GetMultiCloudResourceDiscoverySortOrderDesc,
}

// GetGetMultiCloudResourceDiscoverySortOrderEnumValues Enumerates the set of values for GetMultiCloudResourceDiscoverySortOrderEnum
func GetGetMultiCloudResourceDiscoverySortOrderEnumValues() []GetMultiCloudResourceDiscoverySortOrderEnum {
	values := make([]GetMultiCloudResourceDiscoverySortOrderEnum, 0)
	for _, v := range mappingGetMultiCloudResourceDiscoverySortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetMultiCloudResourceDiscoverySortOrderEnumStringValues Enumerates the set of values in String for GetMultiCloudResourceDiscoverySortOrderEnum
func GetGetMultiCloudResourceDiscoverySortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetMultiCloudResourceDiscoverySortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetMultiCloudResourceDiscoverySortOrderEnum(val string) (GetMultiCloudResourceDiscoverySortOrderEnum, bool) {
	enum, ok := mappingGetMultiCloudResourceDiscoverySortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
