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

// GetOracleDbAzureBlobContainerRequest wrapper for the GetOracleDbAzureBlobContainer operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureBlobContainer.go.html to see an example of how to use GetOracleDbAzureBlobContainerRequest.
type GetOracleDbAzureBlobContainerRequest struct {

	// The ID of the Oracle DB Azure Blob Container Resource.
	OracleDbAzureBlobContainerId *string `mandatory:"true" contributesTo:"path" name:"oracleDbAzureBlobContainerId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder GetOracleDbAzureBlobContainerSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOracleDbAzureBlobContainerRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOracleDbAzureBlobContainerRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetOracleDbAzureBlobContainerRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOracleDbAzureBlobContainerRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetOracleDbAzureBlobContainerRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetOracleDbAzureBlobContainerSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetOracleDbAzureBlobContainerSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetOracleDbAzureBlobContainerResponse wrapper for the GetOracleDbAzureBlobContainer operation
type GetOracleDbAzureBlobContainerResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureBlobContainer instances
	OracleDbAzureBlobContainer `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetOracleDbAzureBlobContainerResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOracleDbAzureBlobContainerResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetOracleDbAzureBlobContainerSortOrderEnum Enum with underlying type: string
type GetOracleDbAzureBlobContainerSortOrderEnum string

// Set of constants representing the allowable values for GetOracleDbAzureBlobContainerSortOrderEnum
const (
	GetOracleDbAzureBlobContainerSortOrderAsc  GetOracleDbAzureBlobContainerSortOrderEnum = "ASC"
	GetOracleDbAzureBlobContainerSortOrderDesc GetOracleDbAzureBlobContainerSortOrderEnum = "DESC"
)

var mappingGetOracleDbAzureBlobContainerSortOrderEnum = map[string]GetOracleDbAzureBlobContainerSortOrderEnum{
	"ASC":  GetOracleDbAzureBlobContainerSortOrderAsc,
	"DESC": GetOracleDbAzureBlobContainerSortOrderDesc,
}

var mappingGetOracleDbAzureBlobContainerSortOrderEnumLowerCase = map[string]GetOracleDbAzureBlobContainerSortOrderEnum{
	"asc":  GetOracleDbAzureBlobContainerSortOrderAsc,
	"desc": GetOracleDbAzureBlobContainerSortOrderDesc,
}

// GetGetOracleDbAzureBlobContainerSortOrderEnumValues Enumerates the set of values for GetOracleDbAzureBlobContainerSortOrderEnum
func GetGetOracleDbAzureBlobContainerSortOrderEnumValues() []GetOracleDbAzureBlobContainerSortOrderEnum {
	values := make([]GetOracleDbAzureBlobContainerSortOrderEnum, 0)
	for _, v := range mappingGetOracleDbAzureBlobContainerSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetOracleDbAzureBlobContainerSortOrderEnumStringValues Enumerates the set of values in String for GetOracleDbAzureBlobContainerSortOrderEnum
func GetGetOracleDbAzureBlobContainerSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetOracleDbAzureBlobContainerSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetOracleDbAzureBlobContainerSortOrderEnum(val string) (GetOracleDbAzureBlobContainerSortOrderEnum, bool) {
	enum, ok := mappingGetOracleDbAzureBlobContainerSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
