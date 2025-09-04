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

// GetOracleDbAzureBlobMountRequest wrapper for the GetOracleDbAzureBlobMount operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureBlobMount.go.html to see an example of how to use GetOracleDbAzureBlobMountRequest.
type GetOracleDbAzureBlobMountRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle DB Azure Blob Mount resource.
	OracleDbAzureBlobMountId *string `mandatory:"true" contributesTo:"path" name:"oracleDbAzureBlobMountId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder GetOracleDbAzureBlobMountSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOracleDbAzureBlobMountRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOracleDbAzureBlobMountRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetOracleDbAzureBlobMountRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOracleDbAzureBlobMountRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetOracleDbAzureBlobMountRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetOracleDbAzureBlobMountSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetOracleDbAzureBlobMountSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetOracleDbAzureBlobMountResponse wrapper for the GetOracleDbAzureBlobMount operation
type GetOracleDbAzureBlobMountResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureBlobMount instances
	OracleDbAzureBlobMount `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetOracleDbAzureBlobMountResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOracleDbAzureBlobMountResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetOracleDbAzureBlobMountSortOrderEnum Enum with underlying type: string
type GetOracleDbAzureBlobMountSortOrderEnum string

// Set of constants representing the allowable values for GetOracleDbAzureBlobMountSortOrderEnum
const (
	GetOracleDbAzureBlobMountSortOrderAsc  GetOracleDbAzureBlobMountSortOrderEnum = "ASC"
	GetOracleDbAzureBlobMountSortOrderDesc GetOracleDbAzureBlobMountSortOrderEnum = "DESC"
)

var mappingGetOracleDbAzureBlobMountSortOrderEnum = map[string]GetOracleDbAzureBlobMountSortOrderEnum{
	"ASC":  GetOracleDbAzureBlobMountSortOrderAsc,
	"DESC": GetOracleDbAzureBlobMountSortOrderDesc,
}

var mappingGetOracleDbAzureBlobMountSortOrderEnumLowerCase = map[string]GetOracleDbAzureBlobMountSortOrderEnum{
	"asc":  GetOracleDbAzureBlobMountSortOrderAsc,
	"desc": GetOracleDbAzureBlobMountSortOrderDesc,
}

// GetGetOracleDbAzureBlobMountSortOrderEnumValues Enumerates the set of values for GetOracleDbAzureBlobMountSortOrderEnum
func GetGetOracleDbAzureBlobMountSortOrderEnumValues() []GetOracleDbAzureBlobMountSortOrderEnum {
	values := make([]GetOracleDbAzureBlobMountSortOrderEnum, 0)
	for _, v := range mappingGetOracleDbAzureBlobMountSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetOracleDbAzureBlobMountSortOrderEnumStringValues Enumerates the set of values in String for GetOracleDbAzureBlobMountSortOrderEnum
func GetGetOracleDbAzureBlobMountSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetOracleDbAzureBlobMountSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetOracleDbAzureBlobMountSortOrderEnum(val string) (GetOracleDbAzureBlobMountSortOrderEnum, bool) {
	enum, ok := mappingGetOracleDbAzureBlobMountSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
