// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCredentialsRequest wrapper for the ListCredentials operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListCredentials.go.html to see an example of how to use ListCredentialsRequest.
type ListCredentialsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
	DatabaseToolsConnectionId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsConnectionId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If-Match is most often used with state-changing methods (e.g., POST, PUT, DELETE) to prevent
	// accidental overwrites when multiple user agentss might be acting in parallel on the same
	// resource (i.e., to prevent the "lost update" problem). In general, it can be used with any
	// method that involves the selection or modification of a representation to abort the request
	// if the selected representation's current entity tag is not a member within the If-Match field value.
	// When specified on an action-specific subresource, the ETag value of the resource on which the
	// action is requested should be provided.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCredentialsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCredentialsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCredentialsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCredentialsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCredentialsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCredentialsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCredentialsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCredentialsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCredentialsResponse wrapper for the ListCredentials operation
type ListCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CredentialCollection instances
	CredentialCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCredentialsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCredentialsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCredentialsSortOrderEnum Enum with underlying type: string
type ListCredentialsSortOrderEnum string

// Set of constants representing the allowable values for ListCredentialsSortOrderEnum
const (
	ListCredentialsSortOrderAsc  ListCredentialsSortOrderEnum = "ASC"
	ListCredentialsSortOrderDesc ListCredentialsSortOrderEnum = "DESC"
)

var mappingListCredentialsSortOrderEnum = map[string]ListCredentialsSortOrderEnum{
	"ASC":  ListCredentialsSortOrderAsc,
	"DESC": ListCredentialsSortOrderDesc,
}

var mappingListCredentialsSortOrderEnumLowerCase = map[string]ListCredentialsSortOrderEnum{
	"asc":  ListCredentialsSortOrderAsc,
	"desc": ListCredentialsSortOrderDesc,
}

// GetListCredentialsSortOrderEnumValues Enumerates the set of values for ListCredentialsSortOrderEnum
func GetListCredentialsSortOrderEnumValues() []ListCredentialsSortOrderEnum {
	values := make([]ListCredentialsSortOrderEnum, 0)
	for _, v := range mappingListCredentialsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCredentialsSortOrderEnumStringValues Enumerates the set of values in String for ListCredentialsSortOrderEnum
func GetListCredentialsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCredentialsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCredentialsSortOrderEnum(val string) (ListCredentialsSortOrderEnum, bool) {
	enum, ok := mappingListCredentialsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
