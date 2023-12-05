// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAwrHubObjectsRequest wrapper for the ListAwrHubObjects operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrHubObjects.go.html to see an example of how to use ListAwrHubObjectsRequest.
type ListAwrHubObjectsRequest struct {

	// Unique Awr Hub Source identifier
	AwrHubSourceId *string `mandatory:"true" contributesTo:"path" name:"awrHubSourceId"`

	// The string to use for matching against the start of object names in a Awr Hub list objects query.
	Prefix *string `mandatory:"false" contributesTo:"query" name:"prefix"`

	// Object names returned by Awr Hub list objects query must be greater or equal to this parameter.
	Start *string `mandatory:"false" contributesTo:"query" name:"start"`

	// Object names returned by Awr Hub list objects query must be strictly less than this parameter.
	End *string `mandatory:"false" contributesTo:"query" name:"end"`

	// When this parameter is set, only objects whose names do not contain the delimiter character
	// (after an optionally specified prefix) are returned in the Awr Hub list objects key of the response body.
	// Scanned objects whose names contain the delimiter have the part of their name up to the first
	// occurrence of the delimiter (including the optional prefix) returned as a set of prefixes.
	// Note that only '/' is a supported delimiter character at this time.
	Delimiter *string `mandatory:"false" contributesTo:"query" name:"delimiter"`

	// Awr Hub Object name after which remaining objects are listed
	StartAfter *string `mandatory:"false" contributesTo:"query" name:"startAfter"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// By default all the fields are returned. Use this parameter to fetch specific fields 'size', 'etag', 'md5',
	// 'timeCreated', 'timeModified', 'storageTier' and 'archivalState' fields. List the names of those fields
	// in a comma-separated, case-insensitive list as the value of this parameter.
	// For example: 'name,etag,timeCreated,md5,timeModified,storageTier,archivalState'.
	Fields ListAwrHubObjectsFieldsEnum `mandatory:"false" contributesTo:"query" name:"fields" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAwrHubObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAwrHubObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAwrHubObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAwrHubObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAwrHubObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAwrHubObjectsFieldsEnum(string(request.Fields)); !ok && request.Fields != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", request.Fields, strings.Join(GetListAwrHubObjectsFieldsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAwrHubObjectsResponse wrapper for the ListAwrHubObjects operation
type ListAwrHubObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ListObjects instances
	ListObjects `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAwrHubObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAwrHubObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAwrHubObjectsFieldsEnum Enum with underlying type: string
type ListAwrHubObjectsFieldsEnum string

// Set of constants representing the allowable values for ListAwrHubObjectsFieldsEnum
const (
	ListAwrHubObjectsFieldsName          ListAwrHubObjectsFieldsEnum = "name"
	ListAwrHubObjectsFieldsSize          ListAwrHubObjectsFieldsEnum = "size"
	ListAwrHubObjectsFieldsEtag          ListAwrHubObjectsFieldsEnum = "etag"
	ListAwrHubObjectsFieldsTimecreated   ListAwrHubObjectsFieldsEnum = "timeCreated"
	ListAwrHubObjectsFieldsMd5           ListAwrHubObjectsFieldsEnum = "md5"
	ListAwrHubObjectsFieldsArchivalstate ListAwrHubObjectsFieldsEnum = "archivalState"
	ListAwrHubObjectsFieldsTimemodified  ListAwrHubObjectsFieldsEnum = "timeModified"
	ListAwrHubObjectsFieldsStoragetier   ListAwrHubObjectsFieldsEnum = "storageTier"
)

var mappingListAwrHubObjectsFieldsEnum = map[string]ListAwrHubObjectsFieldsEnum{
	"name":          ListAwrHubObjectsFieldsName,
	"size":          ListAwrHubObjectsFieldsSize,
	"etag":          ListAwrHubObjectsFieldsEtag,
	"timeCreated":   ListAwrHubObjectsFieldsTimecreated,
	"md5":           ListAwrHubObjectsFieldsMd5,
	"archivalState": ListAwrHubObjectsFieldsArchivalstate,
	"timeModified":  ListAwrHubObjectsFieldsTimemodified,
	"storageTier":   ListAwrHubObjectsFieldsStoragetier,
}

var mappingListAwrHubObjectsFieldsEnumLowerCase = map[string]ListAwrHubObjectsFieldsEnum{
	"name":          ListAwrHubObjectsFieldsName,
	"size":          ListAwrHubObjectsFieldsSize,
	"etag":          ListAwrHubObjectsFieldsEtag,
	"timecreated":   ListAwrHubObjectsFieldsTimecreated,
	"md5":           ListAwrHubObjectsFieldsMd5,
	"archivalstate": ListAwrHubObjectsFieldsArchivalstate,
	"timemodified":  ListAwrHubObjectsFieldsTimemodified,
	"storagetier":   ListAwrHubObjectsFieldsStoragetier,
}

// GetListAwrHubObjectsFieldsEnumValues Enumerates the set of values for ListAwrHubObjectsFieldsEnum
func GetListAwrHubObjectsFieldsEnumValues() []ListAwrHubObjectsFieldsEnum {
	values := make([]ListAwrHubObjectsFieldsEnum, 0)
	for _, v := range mappingListAwrHubObjectsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrHubObjectsFieldsEnumStringValues Enumerates the set of values in String for ListAwrHubObjectsFieldsEnum
func GetListAwrHubObjectsFieldsEnumStringValues() []string {
	return []string{
		"name",
		"size",
		"etag",
		"timeCreated",
		"md5",
		"archivalState",
		"timeModified",
		"storageTier",
	}
}

// GetMappingListAwrHubObjectsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrHubObjectsFieldsEnum(val string) (ListAwrHubObjectsFieldsEnum, bool) {
	enum, ok := mappingListAwrHubObjectsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
