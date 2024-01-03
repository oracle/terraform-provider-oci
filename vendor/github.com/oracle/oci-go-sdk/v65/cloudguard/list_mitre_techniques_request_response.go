// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMitreTechniquesRequest wrapper for the ListMitreTechniques operation
type ListMitreTechniquesRequest struct {

	// unique id of the mitre tactic
	MitreTacticId *string `mandatory:"true" contributesTo:"path" name:"mitreTacticId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMitreTechniquesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMitreTechniquesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMitreTechniquesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMitreTechniquesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMitreTechniquesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMitreTechniquesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMitreTechniquesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMitreTechniquesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMitreTechniquesResponse wrapper for the ListMitreTechniques operation
type ListMitreTechniquesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MitreTechniqueCollection instances
	MitreTechniqueCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMitreTechniquesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMitreTechniquesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMitreTechniquesSortOrderEnum Enum with underlying type: string
type ListMitreTechniquesSortOrderEnum string

// Set of constants representing the allowable values for ListMitreTechniquesSortOrderEnum
const (
	ListMitreTechniquesSortOrderAsc  ListMitreTechniquesSortOrderEnum = "ASC"
	ListMitreTechniquesSortOrderDesc ListMitreTechniquesSortOrderEnum = "DESC"
)

var mappingListMitreTechniquesSortOrderEnum = map[string]ListMitreTechniquesSortOrderEnum{
	"ASC":  ListMitreTechniquesSortOrderAsc,
	"DESC": ListMitreTechniquesSortOrderDesc,
}

var mappingListMitreTechniquesSortOrderEnumLowerCase = map[string]ListMitreTechniquesSortOrderEnum{
	"asc":  ListMitreTechniquesSortOrderAsc,
	"desc": ListMitreTechniquesSortOrderDesc,
}

// GetListMitreTechniquesSortOrderEnumValues Enumerates the set of values for ListMitreTechniquesSortOrderEnum
func GetListMitreTechniquesSortOrderEnumValues() []ListMitreTechniquesSortOrderEnum {
	values := make([]ListMitreTechniquesSortOrderEnum, 0)
	for _, v := range mappingListMitreTechniquesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMitreTechniquesSortOrderEnumStringValues Enumerates the set of values in String for ListMitreTechniquesSortOrderEnum
func GetListMitreTechniquesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMitreTechniquesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMitreTechniquesSortOrderEnum(val string) (ListMitreTechniquesSortOrderEnum, bool) {
	enum, ok := mappingListMitreTechniquesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
