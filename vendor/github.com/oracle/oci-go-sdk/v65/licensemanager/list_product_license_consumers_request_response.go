// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package licensemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProductLicenseConsumersRequest wrapper for the ListProductLicenseConsumers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListProductLicenseConsumers.go.html to see an example of how to use ListProductLicenseConsumersRequest.
type ListProductLicenseConsumersRequest struct {

	// Unique product license identifier.
	ProductLicenseId *string `mandatory:"true" contributesTo:"query" name:"productLicenseId"`

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicates if the given compartment is the root compartment.
	IsCompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isCompartmentIdInSubtree"`

	// The sort order to use, whether `ASC` or `DESC`.
	SortOrder ListProductLicenseConsumersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the rules.
	// Default: `licenseUnitsRequired`
	// * **licenseUnitsRequired:** Sorts by licenseUnitsRequired of the Resource.
	SortBy ListProductLicenseConsumersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProductLicenseConsumersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProductLicenseConsumersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProductLicenseConsumersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProductLicenseConsumersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProductLicenseConsumersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProductLicenseConsumersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProductLicenseConsumersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProductLicenseConsumersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProductLicenseConsumersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProductLicenseConsumersResponse wrapper for the ListProductLicenseConsumers operation
type ListProductLicenseConsumersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProductLicenseConsumerCollection instances
	ProductLicenseConsumerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProductLicenseConsumersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProductLicenseConsumersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProductLicenseConsumersSortOrderEnum Enum with underlying type: string
type ListProductLicenseConsumersSortOrderEnum string

// Set of constants representing the allowable values for ListProductLicenseConsumersSortOrderEnum
const (
	ListProductLicenseConsumersSortOrderAsc  ListProductLicenseConsumersSortOrderEnum = "ASC"
	ListProductLicenseConsumersSortOrderDesc ListProductLicenseConsumersSortOrderEnum = "DESC"
)

var mappingListProductLicenseConsumersSortOrderEnum = map[string]ListProductLicenseConsumersSortOrderEnum{
	"ASC":  ListProductLicenseConsumersSortOrderAsc,
	"DESC": ListProductLicenseConsumersSortOrderDesc,
}

var mappingListProductLicenseConsumersSortOrderEnumLowerCase = map[string]ListProductLicenseConsumersSortOrderEnum{
	"asc":  ListProductLicenseConsumersSortOrderAsc,
	"desc": ListProductLicenseConsumersSortOrderDesc,
}

// GetListProductLicenseConsumersSortOrderEnumValues Enumerates the set of values for ListProductLicenseConsumersSortOrderEnum
func GetListProductLicenseConsumersSortOrderEnumValues() []ListProductLicenseConsumersSortOrderEnum {
	values := make([]ListProductLicenseConsumersSortOrderEnum, 0)
	for _, v := range mappingListProductLicenseConsumersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProductLicenseConsumersSortOrderEnumStringValues Enumerates the set of values in String for ListProductLicenseConsumersSortOrderEnum
func GetListProductLicenseConsumersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProductLicenseConsumersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProductLicenseConsumersSortOrderEnum(val string) (ListProductLicenseConsumersSortOrderEnum, bool) {
	enum, ok := mappingListProductLicenseConsumersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProductLicenseConsumersSortByEnum Enum with underlying type: string
type ListProductLicenseConsumersSortByEnum string

// Set of constants representing the allowable values for ListProductLicenseConsumersSortByEnum
const (
	ListProductLicenseConsumersSortByLicenseunitsrequired ListProductLicenseConsumersSortByEnum = "licenseUnitsRequired"
)

var mappingListProductLicenseConsumersSortByEnum = map[string]ListProductLicenseConsumersSortByEnum{
	"licenseUnitsRequired": ListProductLicenseConsumersSortByLicenseunitsrequired,
}

var mappingListProductLicenseConsumersSortByEnumLowerCase = map[string]ListProductLicenseConsumersSortByEnum{
	"licenseunitsrequired": ListProductLicenseConsumersSortByLicenseunitsrequired,
}

// GetListProductLicenseConsumersSortByEnumValues Enumerates the set of values for ListProductLicenseConsumersSortByEnum
func GetListProductLicenseConsumersSortByEnumValues() []ListProductLicenseConsumersSortByEnum {
	values := make([]ListProductLicenseConsumersSortByEnum, 0)
	for _, v := range mappingListProductLicenseConsumersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProductLicenseConsumersSortByEnumStringValues Enumerates the set of values in String for ListProductLicenseConsumersSortByEnum
func GetListProductLicenseConsumersSortByEnumStringValues() []string {
	return []string{
		"licenseUnitsRequired",
	}
}

// GetMappingListProductLicenseConsumersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProductLicenseConsumersSortByEnum(val string) (ListProductLicenseConsumersSortByEnum, bool) {
	enum, ok := mappingListProductLicenseConsumersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
