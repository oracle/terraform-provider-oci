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

// ListCredentialPublicSynonymsRequest wrapper for the ListCredentialPublicSynonyms operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/ListCredentialPublicSynonyms.go.html to see an example of how to use ListCredentialPublicSynonymsRequest.
type ListCredentialPublicSynonymsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
	DatabaseToolsConnectionId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsConnectionId"`

	// The name of the credential
	CredentialKey *string `mandatory:"true" contributesTo:"path" name:"credentialKey"`

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
	SortOrder ListCredentialPublicSynonymsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCredentialPublicSynonymsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCredentialPublicSynonymsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCredentialPublicSynonymsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCredentialPublicSynonymsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCredentialPublicSynonymsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCredentialPublicSynonymsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCredentialPublicSynonymsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCredentialPublicSynonymsResponse wrapper for the ListCredentialPublicSynonyms operation
type ListCredentialPublicSynonymsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CredentialPublicSynonymCollection instances
	CredentialPublicSynonymCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCredentialPublicSynonymsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCredentialPublicSynonymsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCredentialPublicSynonymsSortOrderEnum Enum with underlying type: string
type ListCredentialPublicSynonymsSortOrderEnum string

// Set of constants representing the allowable values for ListCredentialPublicSynonymsSortOrderEnum
const (
	ListCredentialPublicSynonymsSortOrderAsc  ListCredentialPublicSynonymsSortOrderEnum = "ASC"
	ListCredentialPublicSynonymsSortOrderDesc ListCredentialPublicSynonymsSortOrderEnum = "DESC"
)

var mappingListCredentialPublicSynonymsSortOrderEnum = map[string]ListCredentialPublicSynonymsSortOrderEnum{
	"ASC":  ListCredentialPublicSynonymsSortOrderAsc,
	"DESC": ListCredentialPublicSynonymsSortOrderDesc,
}

var mappingListCredentialPublicSynonymsSortOrderEnumLowerCase = map[string]ListCredentialPublicSynonymsSortOrderEnum{
	"asc":  ListCredentialPublicSynonymsSortOrderAsc,
	"desc": ListCredentialPublicSynonymsSortOrderDesc,
}

// GetListCredentialPublicSynonymsSortOrderEnumValues Enumerates the set of values for ListCredentialPublicSynonymsSortOrderEnum
func GetListCredentialPublicSynonymsSortOrderEnumValues() []ListCredentialPublicSynonymsSortOrderEnum {
	values := make([]ListCredentialPublicSynonymsSortOrderEnum, 0)
	for _, v := range mappingListCredentialPublicSynonymsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCredentialPublicSynonymsSortOrderEnumStringValues Enumerates the set of values in String for ListCredentialPublicSynonymsSortOrderEnum
func GetListCredentialPublicSynonymsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCredentialPublicSynonymsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCredentialPublicSynonymsSortOrderEnum(val string) (ListCredentialPublicSynonymsSortOrderEnum, bool) {
	enum, ok := mappingListCredentialPublicSynonymsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
