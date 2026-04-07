// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableSoftwareSourcesToAddRequest wrapper for the ListAvailableSoftwareSourcesToAdd operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListAvailableSoftwareSourcesToAdd.go.html to see an example of how to use ListAvailableSoftwareSourcesToAddRequest.
type ListAvailableSoftwareSourcesToAddRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given operating system family.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only instances whose architecture type matches the given architecture.
	ArchType []ArchTypeEnum `contributesTo:"query" name:"archType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAvailableSoftwareSourcesToAddSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAvailableSoftwareSourcesToAddSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableSoftwareSourcesToAddRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableSoftwareSourcesToAddRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableSoftwareSourcesToAddRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableSoftwareSourcesToAddRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableSoftwareSourcesToAddRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ArchType {
		if _, ok := GetMappingArchTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArchType: %s. Supported values are: %s.", val, strings.Join(GetArchTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListAvailableSoftwareSourcesToAddSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableSoftwareSourcesToAddSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableSoftwareSourcesToAddSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableSoftwareSourcesToAddSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableSoftwareSourcesToAddResponse wrapper for the ListAvailableSoftwareSourcesToAdd operation
type ListAvailableSoftwareSourcesToAddResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwareSourceRepoCollection instances
	SoftwareSourceRepoCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items in the result. Used for pagination of a list of items.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListAvailableSoftwareSourcesToAddResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableSoftwareSourcesToAddResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableSoftwareSourcesToAddSortOrderEnum Enum with underlying type: string
type ListAvailableSoftwareSourcesToAddSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableSoftwareSourcesToAddSortOrderEnum
const (
	ListAvailableSoftwareSourcesToAddSortOrderAsc  ListAvailableSoftwareSourcesToAddSortOrderEnum = "ASC"
	ListAvailableSoftwareSourcesToAddSortOrderDesc ListAvailableSoftwareSourcesToAddSortOrderEnum = "DESC"
)

var mappingListAvailableSoftwareSourcesToAddSortOrderEnum = map[string]ListAvailableSoftwareSourcesToAddSortOrderEnum{
	"ASC":  ListAvailableSoftwareSourcesToAddSortOrderAsc,
	"DESC": ListAvailableSoftwareSourcesToAddSortOrderDesc,
}

var mappingListAvailableSoftwareSourcesToAddSortOrderEnumLowerCase = map[string]ListAvailableSoftwareSourcesToAddSortOrderEnum{
	"asc":  ListAvailableSoftwareSourcesToAddSortOrderAsc,
	"desc": ListAvailableSoftwareSourcesToAddSortOrderDesc,
}

// GetListAvailableSoftwareSourcesToAddSortOrderEnumValues Enumerates the set of values for ListAvailableSoftwareSourcesToAddSortOrderEnum
func GetListAvailableSoftwareSourcesToAddSortOrderEnumValues() []ListAvailableSoftwareSourcesToAddSortOrderEnum {
	values := make([]ListAvailableSoftwareSourcesToAddSortOrderEnum, 0)
	for _, v := range mappingListAvailableSoftwareSourcesToAddSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableSoftwareSourcesToAddSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableSoftwareSourcesToAddSortOrderEnum
func GetListAvailableSoftwareSourcesToAddSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableSoftwareSourcesToAddSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableSoftwareSourcesToAddSortOrderEnum(val string) (ListAvailableSoftwareSourcesToAddSortOrderEnum, bool) {
	enum, ok := mappingListAvailableSoftwareSourcesToAddSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableSoftwareSourcesToAddSortByEnum Enum with underlying type: string
type ListAvailableSoftwareSourcesToAddSortByEnum string

// Set of constants representing the allowable values for ListAvailableSoftwareSourcesToAddSortByEnum
const (
	ListAvailableSoftwareSourcesToAddSortByTimecreated ListAvailableSoftwareSourcesToAddSortByEnum = "timeCreated"
	ListAvailableSoftwareSourcesToAddSortByDisplayname ListAvailableSoftwareSourcesToAddSortByEnum = "displayName"
)

var mappingListAvailableSoftwareSourcesToAddSortByEnum = map[string]ListAvailableSoftwareSourcesToAddSortByEnum{
	"timeCreated": ListAvailableSoftwareSourcesToAddSortByTimecreated,
	"displayName": ListAvailableSoftwareSourcesToAddSortByDisplayname,
}

var mappingListAvailableSoftwareSourcesToAddSortByEnumLowerCase = map[string]ListAvailableSoftwareSourcesToAddSortByEnum{
	"timecreated": ListAvailableSoftwareSourcesToAddSortByTimecreated,
	"displayname": ListAvailableSoftwareSourcesToAddSortByDisplayname,
}

// GetListAvailableSoftwareSourcesToAddSortByEnumValues Enumerates the set of values for ListAvailableSoftwareSourcesToAddSortByEnum
func GetListAvailableSoftwareSourcesToAddSortByEnumValues() []ListAvailableSoftwareSourcesToAddSortByEnum {
	values := make([]ListAvailableSoftwareSourcesToAddSortByEnum, 0)
	for _, v := range mappingListAvailableSoftwareSourcesToAddSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableSoftwareSourcesToAddSortByEnumStringValues Enumerates the set of values in String for ListAvailableSoftwareSourcesToAddSortByEnum
func GetListAvailableSoftwareSourcesToAddSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAvailableSoftwareSourcesToAddSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableSoftwareSourcesToAddSortByEnum(val string) (ListAvailableSoftwareSourcesToAddSortByEnum, bool) {
	enum, ok := mappingListAvailableSoftwareSourcesToAddSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
