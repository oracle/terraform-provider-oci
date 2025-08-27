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

// GetOracleDbAzureVaultAssociationRequest wrapper for the GetOracleDbAzureVaultAssociation operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/GetOracleDbAzureVaultAssociation.go.html to see an example of how to use GetOracleDbAzureVaultAssociationRequest.
type GetOracleDbAzureVaultAssociationRequest struct {

	// The ID of the Oracle DB Azure Vault Association Resource.
	OracleDbAzureVaultAssociationId *string `mandatory:"true" contributesTo:"path" name:"oracleDbAzureVaultAssociationId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder GetOracleDbAzureVaultAssociationSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOracleDbAzureVaultAssociationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOracleDbAzureVaultAssociationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetOracleDbAzureVaultAssociationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOracleDbAzureVaultAssociationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetOracleDbAzureVaultAssociationRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetOracleDbAzureVaultAssociationSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetOracleDbAzureVaultAssociationSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetOracleDbAzureVaultAssociationResponse wrapper for the GetOracleDbAzureVaultAssociation operation
type GetOracleDbAzureVaultAssociationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OracleDbAzureVaultAssociation instances
	OracleDbAzureVaultAssociation `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetOracleDbAzureVaultAssociationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOracleDbAzureVaultAssociationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetOracleDbAzureVaultAssociationSortOrderEnum Enum with underlying type: string
type GetOracleDbAzureVaultAssociationSortOrderEnum string

// Set of constants representing the allowable values for GetOracleDbAzureVaultAssociationSortOrderEnum
const (
	GetOracleDbAzureVaultAssociationSortOrderAsc  GetOracleDbAzureVaultAssociationSortOrderEnum = "ASC"
	GetOracleDbAzureVaultAssociationSortOrderDesc GetOracleDbAzureVaultAssociationSortOrderEnum = "DESC"
)

var mappingGetOracleDbAzureVaultAssociationSortOrderEnum = map[string]GetOracleDbAzureVaultAssociationSortOrderEnum{
	"ASC":  GetOracleDbAzureVaultAssociationSortOrderAsc,
	"DESC": GetOracleDbAzureVaultAssociationSortOrderDesc,
}

var mappingGetOracleDbAzureVaultAssociationSortOrderEnumLowerCase = map[string]GetOracleDbAzureVaultAssociationSortOrderEnum{
	"asc":  GetOracleDbAzureVaultAssociationSortOrderAsc,
	"desc": GetOracleDbAzureVaultAssociationSortOrderDesc,
}

// GetGetOracleDbAzureVaultAssociationSortOrderEnumValues Enumerates the set of values for GetOracleDbAzureVaultAssociationSortOrderEnum
func GetGetOracleDbAzureVaultAssociationSortOrderEnumValues() []GetOracleDbAzureVaultAssociationSortOrderEnum {
	values := make([]GetOracleDbAzureVaultAssociationSortOrderEnum, 0)
	for _, v := range mappingGetOracleDbAzureVaultAssociationSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetOracleDbAzureVaultAssociationSortOrderEnumStringValues Enumerates the set of values in String for GetOracleDbAzureVaultAssociationSortOrderEnum
func GetGetOracleDbAzureVaultAssociationSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetOracleDbAzureVaultAssociationSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetOracleDbAzureVaultAssociationSortOrderEnum(val string) (GetOracleDbAzureVaultAssociationSortOrderEnum, bool) {
	enum, ok := mappingGetOracleDbAzureVaultAssociationSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
