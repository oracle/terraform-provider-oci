// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMySmtpCredentialsRequest wrapper for the ListMySmtpCredentials operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/ListMySmtpCredentials.go.html to see an example of how to use ListMySmtpCredentialsRequest.
type ListMySmtpCredentialsRequest struct {

	// OPTIONAL. The filter string that is used to request a subset of resources. The filter string MUST be a valid filter expression. See the Filtering section of the SCIM specification for more information (Section 3.4.2.2). The string should contain at least one condition that each item must match in order to be returned in the search results. Each condition specifies an attribute, an operator, and a value. Conditions within a filter can be connected by logical operators (such as AND and OR). Sets of conditions can be grouped together using parentheses.
	Filter *string `mandatory:"false" contributesTo:"query" name:"filter"`

	// OPTIONAL. A string that indicates the attribute whose value SHALL be used to order the returned responses. The sortBy attribute MUST be in standard attribute notation form. See the Attribute Notation section of the SCIM specification for more information (Section 3.10). Also, see the Sorting section of the SCIM specification for more information (Section 3.4.2.3).
	SortBy *string `mandatory:"false" contributesTo:"query" name:"sortBy"`

	// A string that indicates the order in which the sortBy parameter is applied. Allowed values are 'ascending' and 'descending'. See (Sorting Section (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.4.2.3)). OPTIONAL.
	SortOrder ListMySmtpCredentialsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// OPTIONAL. An integer that indicates the 1-based index of the first query result. See the Pagination section of the SCIM specification for more information. (Section 3.4.2.4). The number of results pages to return. The first page is 1. Specify 2 to access the second page of results, and so on.
	StartIndex *int `mandatory:"false" contributesTo:"query" name:"startIndex"`

	// OPTIONAL. An integer that indicates the desired maximum number of query results per page. 1000 is the largest value that you can use. See the Pagination section of the System for Cross-Domain Identity Management Protocol specification for more information. (Section 3.4.2.4).
	Count *int `mandatory:"false" contributesTo:"query" name:"count"`

	// The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
	Authorization *string `mandatory:"false" contributesTo:"header" name:"authorization"`

	// An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
	ResourceTypeSchemaVersion *string `mandatory:"false" contributesTo:"header" name:"resource_type_schema_version"`

	// A token you supply to uniquely identify the request and provide idempotency if the request is retried. Idempotency tokens expire after 24 hours.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The value of the `opc-next-page` response header from the previous 'List' call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated 'List' call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMySmtpCredentialsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMySmtpCredentialsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMySmtpCredentialsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMySmtpCredentialsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMySmtpCredentialsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMySmtpCredentialsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMySmtpCredentialsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMySmtpCredentialsResponse wrapper for the ListMySmtpCredentials operation
type ListMySmtpCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MySmtpCredentials instances
	MySmtpCredentials `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMySmtpCredentialsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMySmtpCredentialsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMySmtpCredentialsSortOrderEnum Enum with underlying type: string
type ListMySmtpCredentialsSortOrderEnum string

// Set of constants representing the allowable values for ListMySmtpCredentialsSortOrderEnum
const (
	ListMySmtpCredentialsSortOrderAscending  ListMySmtpCredentialsSortOrderEnum = "ASCENDING"
	ListMySmtpCredentialsSortOrderDescending ListMySmtpCredentialsSortOrderEnum = "DESCENDING"
)

var mappingListMySmtpCredentialsSortOrderEnum = map[string]ListMySmtpCredentialsSortOrderEnum{
	"ASCENDING":  ListMySmtpCredentialsSortOrderAscending,
	"DESCENDING": ListMySmtpCredentialsSortOrderDescending,
}

var mappingListMySmtpCredentialsSortOrderEnumLowerCase = map[string]ListMySmtpCredentialsSortOrderEnum{
	"ascending":  ListMySmtpCredentialsSortOrderAscending,
	"descending": ListMySmtpCredentialsSortOrderDescending,
}

// GetListMySmtpCredentialsSortOrderEnumValues Enumerates the set of values for ListMySmtpCredentialsSortOrderEnum
func GetListMySmtpCredentialsSortOrderEnumValues() []ListMySmtpCredentialsSortOrderEnum {
	values := make([]ListMySmtpCredentialsSortOrderEnum, 0)
	for _, v := range mappingListMySmtpCredentialsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMySmtpCredentialsSortOrderEnumStringValues Enumerates the set of values in String for ListMySmtpCredentialsSortOrderEnum
func GetListMySmtpCredentialsSortOrderEnumStringValues() []string {
	return []string{
		"ASCENDING",
		"DESCENDING",
	}
}

// GetMappingListMySmtpCredentialsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMySmtpCredentialsSortOrderEnum(val string) (ListMySmtpCredentialsSortOrderEnum, bool) {
	enum, ok := mappingListMySmtpCredentialsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
