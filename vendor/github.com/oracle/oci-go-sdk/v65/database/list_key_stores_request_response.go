// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListKeyStoresRequest wrapper for the ListKeyStores operation
type ListKeyStoresRequest struct {

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the key store type given. The match is not case sensitive.
	KeyStoreType ListKeyStoresKeyStoreTypeEnum `mandatory:"false" contributesTo:"query" name:"keyStoreType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListKeyStoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListKeyStoresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListKeyStoresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListKeyStoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListKeyStoresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListKeyStoresKeyStoreTypeEnum(string(request.KeyStoreType)); !ok && request.KeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyStoreType: %s. Supported values are: %s.", request.KeyStoreType, strings.Join(GetListKeyStoresKeyStoreTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListKeyStoresResponse wrapper for the ListKeyStores operation
type ListKeyStoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []KeyStoreSummary instances
	Items []KeyStoreSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListKeyStoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListKeyStoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListKeyStoresKeyStoreTypeEnum Enum with underlying type: string
type ListKeyStoresKeyStoreTypeEnum string

// Set of constants representing the allowable values for ListKeyStoresKeyStoreTypeEnum
const (
	ListKeyStoresKeyStoreTypeOracleKeyVault ListKeyStoresKeyStoreTypeEnum = "ORACLE_KEY_VAULT"
	ListKeyStoresKeyStoreTypeThales         ListKeyStoresKeyStoreTypeEnum = "THALES"
)

var mappingListKeyStoresKeyStoreTypeEnum = map[string]ListKeyStoresKeyStoreTypeEnum{
	"ORACLE_KEY_VAULT": ListKeyStoresKeyStoreTypeOracleKeyVault,
	"THALES":           ListKeyStoresKeyStoreTypeThales,
}

var mappingListKeyStoresKeyStoreTypeEnumLowerCase = map[string]ListKeyStoresKeyStoreTypeEnum{
	"oracle_key_vault": ListKeyStoresKeyStoreTypeOracleKeyVault,
	"thales":           ListKeyStoresKeyStoreTypeThales,
}

// GetListKeyStoresKeyStoreTypeEnumValues Enumerates the set of values for ListKeyStoresKeyStoreTypeEnum
func GetListKeyStoresKeyStoreTypeEnumValues() []ListKeyStoresKeyStoreTypeEnum {
	values := make([]ListKeyStoresKeyStoreTypeEnum, 0)
	for _, v := range mappingListKeyStoresKeyStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListKeyStoresKeyStoreTypeEnumStringValues Enumerates the set of values in String for ListKeyStoresKeyStoreTypeEnum
func GetListKeyStoresKeyStoreTypeEnumStringValues() []string {
	return []string{
		"ORACLE_KEY_VAULT",
		"THALES",
	}
}

// GetMappingListKeyStoresKeyStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListKeyStoresKeyStoreTypeEnum(val string) (ListKeyStoresKeyStoreTypeEnum, bool) {
	enum, ok := mappingListKeyStoresKeyStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
