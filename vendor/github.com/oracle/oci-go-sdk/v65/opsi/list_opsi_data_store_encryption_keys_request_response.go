// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOpsiDataStoreEncryptionKeysRequest wrapper for the ListOpsiDataStoreEncryptionKeys operation
type ListOpsiDataStoreEncryptionKeysRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOpsiDataStoreEncryptionKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeActivated is descending. If no value is specified timeActivated is default.
	SortBy ListOpsiDataStoreEncryptionKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Encryption Key Activation Status
	ActivationStates []ListOpsiDataStoreEncryptionKeysActivationStatesEnum `contributesTo:"query" name:"activationStates" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOpsiDataStoreEncryptionKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOpsiDataStoreEncryptionKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOpsiDataStoreEncryptionKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOpsiDataStoreEncryptionKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOpsiDataStoreEncryptionKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOpsiDataStoreEncryptionKeysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOpsiDataStoreEncryptionKeysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOpsiDataStoreEncryptionKeysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOpsiDataStoreEncryptionKeysSortByEnumStringValues(), ",")))
	}
	for _, val := range request.ActivationStates {
		if _, ok := GetMappingListOpsiDataStoreEncryptionKeysActivationStatesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActivationStates: %s. Supported values are: %s.", val, strings.Join(GetListOpsiDataStoreEncryptionKeysActivationStatesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOpsiDataStoreEncryptionKeysResponse wrapper for the ListOpsiDataStoreEncryptionKeys operation
type ListOpsiDataStoreEncryptionKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OpsiDataStoreEncryptionKeyCollection instances
	OpsiDataStoreEncryptionKeyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOpsiDataStoreEncryptionKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOpsiDataStoreEncryptionKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOpsiDataStoreEncryptionKeysSortOrderEnum Enum with underlying type: string
type ListOpsiDataStoreEncryptionKeysSortOrderEnum string

// Set of constants representing the allowable values for ListOpsiDataStoreEncryptionKeysSortOrderEnum
const (
	ListOpsiDataStoreEncryptionKeysSortOrderAsc  ListOpsiDataStoreEncryptionKeysSortOrderEnum = "ASC"
	ListOpsiDataStoreEncryptionKeysSortOrderDesc ListOpsiDataStoreEncryptionKeysSortOrderEnum = "DESC"
)

var mappingListOpsiDataStoreEncryptionKeysSortOrderEnum = map[string]ListOpsiDataStoreEncryptionKeysSortOrderEnum{
	"ASC":  ListOpsiDataStoreEncryptionKeysSortOrderAsc,
	"DESC": ListOpsiDataStoreEncryptionKeysSortOrderDesc,
}

var mappingListOpsiDataStoreEncryptionKeysSortOrderEnumLowerCase = map[string]ListOpsiDataStoreEncryptionKeysSortOrderEnum{
	"asc":  ListOpsiDataStoreEncryptionKeysSortOrderAsc,
	"desc": ListOpsiDataStoreEncryptionKeysSortOrderDesc,
}

// GetListOpsiDataStoreEncryptionKeysSortOrderEnumValues Enumerates the set of values for ListOpsiDataStoreEncryptionKeysSortOrderEnum
func GetListOpsiDataStoreEncryptionKeysSortOrderEnumValues() []ListOpsiDataStoreEncryptionKeysSortOrderEnum {
	values := make([]ListOpsiDataStoreEncryptionKeysSortOrderEnum, 0)
	for _, v := range mappingListOpsiDataStoreEncryptionKeysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpsiDataStoreEncryptionKeysSortOrderEnumStringValues Enumerates the set of values in String for ListOpsiDataStoreEncryptionKeysSortOrderEnum
func GetListOpsiDataStoreEncryptionKeysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOpsiDataStoreEncryptionKeysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpsiDataStoreEncryptionKeysSortOrderEnum(val string) (ListOpsiDataStoreEncryptionKeysSortOrderEnum, bool) {
	enum, ok := mappingListOpsiDataStoreEncryptionKeysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOpsiDataStoreEncryptionKeysSortByEnum Enum with underlying type: string
type ListOpsiDataStoreEncryptionKeysSortByEnum string

// Set of constants representing the allowable values for ListOpsiDataStoreEncryptionKeysSortByEnum
const (
	ListOpsiDataStoreEncryptionKeysSortByTimeactivated ListOpsiDataStoreEncryptionKeysSortByEnum = "timeActivated"
)

var mappingListOpsiDataStoreEncryptionKeysSortByEnum = map[string]ListOpsiDataStoreEncryptionKeysSortByEnum{
	"timeActivated": ListOpsiDataStoreEncryptionKeysSortByTimeactivated,
}

var mappingListOpsiDataStoreEncryptionKeysSortByEnumLowerCase = map[string]ListOpsiDataStoreEncryptionKeysSortByEnum{
	"timeactivated": ListOpsiDataStoreEncryptionKeysSortByTimeactivated,
}

// GetListOpsiDataStoreEncryptionKeysSortByEnumValues Enumerates the set of values for ListOpsiDataStoreEncryptionKeysSortByEnum
func GetListOpsiDataStoreEncryptionKeysSortByEnumValues() []ListOpsiDataStoreEncryptionKeysSortByEnum {
	values := make([]ListOpsiDataStoreEncryptionKeysSortByEnum, 0)
	for _, v := range mappingListOpsiDataStoreEncryptionKeysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpsiDataStoreEncryptionKeysSortByEnumStringValues Enumerates the set of values in String for ListOpsiDataStoreEncryptionKeysSortByEnum
func GetListOpsiDataStoreEncryptionKeysSortByEnumStringValues() []string {
	return []string{
		"timeActivated",
	}
}

// GetMappingListOpsiDataStoreEncryptionKeysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpsiDataStoreEncryptionKeysSortByEnum(val string) (ListOpsiDataStoreEncryptionKeysSortByEnum, bool) {
	enum, ok := mappingListOpsiDataStoreEncryptionKeysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOpsiDataStoreEncryptionKeysActivationStatesEnum Enum with underlying type: string
type ListOpsiDataStoreEncryptionKeysActivationStatesEnum string

// Set of constants representing the allowable values for ListOpsiDataStoreEncryptionKeysActivationStatesEnum
const (
	ListOpsiDataStoreEncryptionKeysActivationStatesActive   ListOpsiDataStoreEncryptionKeysActivationStatesEnum = "ACTIVE"
	ListOpsiDataStoreEncryptionKeysActivationStatesInactive ListOpsiDataStoreEncryptionKeysActivationStatesEnum = "INACTIVE"
)

var mappingListOpsiDataStoreEncryptionKeysActivationStatesEnum = map[string]ListOpsiDataStoreEncryptionKeysActivationStatesEnum{
	"ACTIVE":   ListOpsiDataStoreEncryptionKeysActivationStatesActive,
	"INACTIVE": ListOpsiDataStoreEncryptionKeysActivationStatesInactive,
}

var mappingListOpsiDataStoreEncryptionKeysActivationStatesEnumLowerCase = map[string]ListOpsiDataStoreEncryptionKeysActivationStatesEnum{
	"active":   ListOpsiDataStoreEncryptionKeysActivationStatesActive,
	"inactive": ListOpsiDataStoreEncryptionKeysActivationStatesInactive,
}

// GetListOpsiDataStoreEncryptionKeysActivationStatesEnumValues Enumerates the set of values for ListOpsiDataStoreEncryptionKeysActivationStatesEnum
func GetListOpsiDataStoreEncryptionKeysActivationStatesEnumValues() []ListOpsiDataStoreEncryptionKeysActivationStatesEnum {
	values := make([]ListOpsiDataStoreEncryptionKeysActivationStatesEnum, 0)
	for _, v := range mappingListOpsiDataStoreEncryptionKeysActivationStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetListOpsiDataStoreEncryptionKeysActivationStatesEnumStringValues Enumerates the set of values in String for ListOpsiDataStoreEncryptionKeysActivationStatesEnum
func GetListOpsiDataStoreEncryptionKeysActivationStatesEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListOpsiDataStoreEncryptionKeysActivationStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOpsiDataStoreEncryptionKeysActivationStatesEnum(val string) (ListOpsiDataStoreEncryptionKeysActivationStatesEnum, bool) {
	enum, ok := mappingListOpsiDataStoreEncryptionKeysActivationStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
