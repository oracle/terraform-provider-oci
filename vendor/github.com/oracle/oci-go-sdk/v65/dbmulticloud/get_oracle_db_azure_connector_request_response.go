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

// GetOracleDbAzureConnectorRequest wrapper for the GetOracleDbAzureConnector operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureConnector.go.html to see an example of how to use GetOracleDbAzureConnectorRequest.
type GetOracleDbAzureConnectorRequest struct {

	// The ID of the Oracle DB Azure Connector Resource.
	OracleDbAzureConnectorId *string `mandatory:"true" contributesTo:"path" name:"oracleDbAzureConnectorId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder GetOracleDbAzureConnectorSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOracleDbAzureConnectorRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOracleDbAzureConnectorRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetOracleDbAzureConnectorRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOracleDbAzureConnectorRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetOracleDbAzureConnectorRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetOracleDbAzureConnectorSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetOracleDbAzureConnectorSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetOracleDbAzureConnectorResponse wrapper for the GetOracleDbAzureConnector operation
type GetOracleDbAzureConnectorResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureConnector instances
	OracleDbAzureConnector `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetOracleDbAzureConnectorResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOracleDbAzureConnectorResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetOracleDbAzureConnectorSortOrderEnum Enum with underlying type: string
type GetOracleDbAzureConnectorSortOrderEnum string

// Set of constants representing the allowable values for GetOracleDbAzureConnectorSortOrderEnum
const (
	GetOracleDbAzureConnectorSortOrderAsc  GetOracleDbAzureConnectorSortOrderEnum = "ASC"
	GetOracleDbAzureConnectorSortOrderDesc GetOracleDbAzureConnectorSortOrderEnum = "DESC"
)

var mappingGetOracleDbAzureConnectorSortOrderEnum = map[string]GetOracleDbAzureConnectorSortOrderEnum{
	"ASC":  GetOracleDbAzureConnectorSortOrderAsc,
	"DESC": GetOracleDbAzureConnectorSortOrderDesc,
}

var mappingGetOracleDbAzureConnectorSortOrderEnumLowerCase = map[string]GetOracleDbAzureConnectorSortOrderEnum{
	"asc":  GetOracleDbAzureConnectorSortOrderAsc,
	"desc": GetOracleDbAzureConnectorSortOrderDesc,
}

// GetGetOracleDbAzureConnectorSortOrderEnumValues Enumerates the set of values for GetOracleDbAzureConnectorSortOrderEnum
func GetGetOracleDbAzureConnectorSortOrderEnumValues() []GetOracleDbAzureConnectorSortOrderEnum {
	values := make([]GetOracleDbAzureConnectorSortOrderEnum, 0)
	for _, v := range mappingGetOracleDbAzureConnectorSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetOracleDbAzureConnectorSortOrderEnumStringValues Enumerates the set of values in String for GetOracleDbAzureConnectorSortOrderEnum
func GetGetOracleDbAzureConnectorSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetOracleDbAzureConnectorSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetOracleDbAzureConnectorSortOrderEnum(val string) (GetOracleDbAzureConnectorSortOrderEnum, bool) {
	enum, ok := mappingGetOracleDbAzureConnectorSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
