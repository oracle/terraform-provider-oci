// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableSoftwareSourcesToAddInCompartmentRequest wrapper for the ListAvailableSoftwareSourcesToAddInCompartment operation
type ListAvailableSoftwareSourcesToAddInCompartmentRequest struct {

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
	SortOrder ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableSoftwareSourcesToAddInCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableSoftwareSourcesToAddInCompartmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableSoftwareSourcesToAddInCompartmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableSoftwareSourcesToAddInCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableSoftwareSourcesToAddInCompartmentRequest) ValidateEnumValue() (bool, error) {
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

	if _, ok := GetMappingListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableSoftwareSourcesToAddInCompartmentSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableSoftwareSourcesToAddInCompartmentSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableSoftwareSourcesToAddInCompartmentResponse wrapper for the ListAvailableSoftwareSourcesToAddInCompartment operation
type ListAvailableSoftwareSourcesToAddInCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwareSourceRepoCollection instances
	SoftwareSourceRepoCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableSoftwareSourcesToAddInCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableSoftwareSourcesToAddInCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum Enum with underlying type: string
type ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum
const (
	ListAvailableSoftwareSourcesToAddInCompartmentSortOrderAsc  ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum = "ASC"
	ListAvailableSoftwareSourcesToAddInCompartmentSortOrderDesc ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum = "DESC"
)

var mappingListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum = map[string]ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum{
	"ASC":  ListAvailableSoftwareSourcesToAddInCompartmentSortOrderAsc,
	"DESC": ListAvailableSoftwareSourcesToAddInCompartmentSortOrderDesc,
}

var mappingListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumLowerCase = map[string]ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum{
	"asc":  ListAvailableSoftwareSourcesToAddInCompartmentSortOrderAsc,
	"desc": ListAvailableSoftwareSourcesToAddInCompartmentSortOrderDesc,
}

// GetListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumValues Enumerates the set of values for ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum
func GetListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumValues() []ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum {
	values := make([]ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum, 0)
	for _, v := range mappingListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum
func GetListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum(val string) (ListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum, bool) {
	enum, ok := mappingListAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum Enum with underlying type: string
type ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum string

// Set of constants representing the allowable values for ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum
const (
	ListAvailableSoftwareSourcesToAddInCompartmentSortByTimecreated ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum = "timeCreated"
	ListAvailableSoftwareSourcesToAddInCompartmentSortByDisplayname ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum = "displayName"
)

var mappingListAvailableSoftwareSourcesToAddInCompartmentSortByEnum = map[string]ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum{
	"timeCreated": ListAvailableSoftwareSourcesToAddInCompartmentSortByTimecreated,
	"displayName": ListAvailableSoftwareSourcesToAddInCompartmentSortByDisplayname,
}

var mappingListAvailableSoftwareSourcesToAddInCompartmentSortByEnumLowerCase = map[string]ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum{
	"timecreated": ListAvailableSoftwareSourcesToAddInCompartmentSortByTimecreated,
	"displayname": ListAvailableSoftwareSourcesToAddInCompartmentSortByDisplayname,
}

// GetListAvailableSoftwareSourcesToAddInCompartmentSortByEnumValues Enumerates the set of values for ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum
func GetListAvailableSoftwareSourcesToAddInCompartmentSortByEnumValues() []ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum {
	values := make([]ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum, 0)
	for _, v := range mappingListAvailableSoftwareSourcesToAddInCompartmentSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableSoftwareSourcesToAddInCompartmentSortByEnumStringValues Enumerates the set of values in String for ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum
func GetListAvailableSoftwareSourcesToAddInCompartmentSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAvailableSoftwareSourcesToAddInCompartmentSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableSoftwareSourcesToAddInCompartmentSortByEnum(val string) (ListAvailableSoftwareSourcesToAddInCompartmentSortByEnum, bool) {
	enum, ok := mappingListAvailableSoftwareSourcesToAddInCompartmentSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
