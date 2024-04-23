// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSoftwarePackageSoftwareSourcesRequest wrapper for the ListSoftwarePackageSoftwareSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListSoftwarePackageSoftwareSources.go.html to see an example of how to use ListSoftwarePackageSoftwareSourcesRequest.
type ListSoftwarePackageSoftwareSourcesRequest struct {

	// The name of the software package.
	SoftwarePackageName *string `mandatory:"true" contributesTo:"path" name:"softwarePackageName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This parameter is required and returns only resources contained within the specified compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The type of the software source.
	SoftwareSourceType []SoftwareSourceTypeEnum `contributesTo:"query" name:"softwareSourceType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the given operating system family.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only instances whose architecture type matches the given architecture.
	ArchType []ArchTypeEnum `contributesTo:"query" name:"archType" omitEmpty:"true" collectionFormat:"multi"`

	// The availabilities of the software source in a non-OCI environment for a tenancy.
	Availability []AvailabilityEnum `contributesTo:"query" name:"availability" omitEmpty:"true" collectionFormat:"multi"`

	// The availabilities of the software source in an OCI environment for a tenancy.
	AvailabilityAtOci []AvailabilityEnum `contributesTo:"query" name:"availabilityAtOci" omitEmpty:"true" collectionFormat:"multi"`

	// The availabilities of the software source. Use this query parameter to filter across availabilities in different environments.
	AvailabilityAnywhere []AvailabilityEnum `contributesTo:"query" name:"availabilityAnywhere" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return resources that match the given user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return resources that may partially match the given display name.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListSoftwarePackageSoftwareSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListSoftwarePackageSoftwareSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only software sources whose state matches the given state.
	LifecycleState []SoftwareSourceLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwarePackageSoftwareSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwarePackageSoftwareSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSoftwarePackageSoftwareSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwarePackageSoftwareSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSoftwarePackageSoftwareSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.SoftwareSourceType {
		if _, ok := GetMappingSoftwareSourceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareSourceType: %s. Supported values are: %s.", val, strings.Join(GetSoftwareSourceTypeEnumStringValues(), ",")))
		}
	}

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

	for _, val := range request.Availability {
		if _, ok := GetMappingAvailabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Availability: %s. Supported values are: %s.", val, strings.Join(GetAvailabilityEnumStringValues(), ",")))
		}
	}

	for _, val := range request.AvailabilityAtOci {
		if _, ok := GetMappingAvailabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityAtOci: %s. Supported values are: %s.", val, strings.Join(GetAvailabilityEnumStringValues(), ",")))
		}
	}

	for _, val := range request.AvailabilityAnywhere {
		if _, ok := GetMappingAvailabilityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityAnywhere: %s. Supported values are: %s.", val, strings.Join(GetAvailabilityEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListSoftwarePackageSoftwareSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSoftwarePackageSoftwareSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSoftwarePackageSoftwareSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSoftwarePackageSoftwareSourcesSortByEnumStringValues(), ",")))
	}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingSoftwareSourceLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetSoftwareSourceLifecycleStateEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSoftwarePackageSoftwareSourcesResponse wrapper for the ListSoftwarePackageSoftwareSources operation
type ListSoftwarePackageSoftwareSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SoftwareSourceCollection instances
	SoftwareSourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwarePackageSoftwareSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwarePackageSoftwareSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwarePackageSoftwareSourcesSortOrderEnum Enum with underlying type: string
type ListSoftwarePackageSoftwareSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwarePackageSoftwareSourcesSortOrderEnum
const (
	ListSoftwarePackageSoftwareSourcesSortOrderAsc  ListSoftwarePackageSoftwareSourcesSortOrderEnum = "ASC"
	ListSoftwarePackageSoftwareSourcesSortOrderDesc ListSoftwarePackageSoftwareSourcesSortOrderEnum = "DESC"
)

var mappingListSoftwarePackageSoftwareSourcesSortOrderEnum = map[string]ListSoftwarePackageSoftwareSourcesSortOrderEnum{
	"ASC":  ListSoftwarePackageSoftwareSourcesSortOrderAsc,
	"DESC": ListSoftwarePackageSoftwareSourcesSortOrderDesc,
}

var mappingListSoftwarePackageSoftwareSourcesSortOrderEnumLowerCase = map[string]ListSoftwarePackageSoftwareSourcesSortOrderEnum{
	"asc":  ListSoftwarePackageSoftwareSourcesSortOrderAsc,
	"desc": ListSoftwarePackageSoftwareSourcesSortOrderDesc,
}

// GetListSoftwarePackageSoftwareSourcesSortOrderEnumValues Enumerates the set of values for ListSoftwarePackageSoftwareSourcesSortOrderEnum
func GetListSoftwarePackageSoftwareSourcesSortOrderEnumValues() []ListSoftwarePackageSoftwareSourcesSortOrderEnum {
	values := make([]ListSoftwarePackageSoftwareSourcesSortOrderEnum, 0)
	for _, v := range mappingListSoftwarePackageSoftwareSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwarePackageSoftwareSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListSoftwarePackageSoftwareSourcesSortOrderEnum
func GetListSoftwarePackageSoftwareSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSoftwarePackageSoftwareSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwarePackageSoftwareSourcesSortOrderEnum(val string) (ListSoftwarePackageSoftwareSourcesSortOrderEnum, bool) {
	enum, ok := mappingListSoftwarePackageSoftwareSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSoftwarePackageSoftwareSourcesSortByEnum Enum with underlying type: string
type ListSoftwarePackageSoftwareSourcesSortByEnum string

// Set of constants representing the allowable values for ListSoftwarePackageSoftwareSourcesSortByEnum
const (
	ListSoftwarePackageSoftwareSourcesSortByTimecreated ListSoftwarePackageSoftwareSourcesSortByEnum = "timeCreated"
	ListSoftwarePackageSoftwareSourcesSortByDisplayname ListSoftwarePackageSoftwareSourcesSortByEnum = "displayName"
)

var mappingListSoftwarePackageSoftwareSourcesSortByEnum = map[string]ListSoftwarePackageSoftwareSourcesSortByEnum{
	"timeCreated": ListSoftwarePackageSoftwareSourcesSortByTimecreated,
	"displayName": ListSoftwarePackageSoftwareSourcesSortByDisplayname,
}

var mappingListSoftwarePackageSoftwareSourcesSortByEnumLowerCase = map[string]ListSoftwarePackageSoftwareSourcesSortByEnum{
	"timecreated": ListSoftwarePackageSoftwareSourcesSortByTimecreated,
	"displayname": ListSoftwarePackageSoftwareSourcesSortByDisplayname,
}

// GetListSoftwarePackageSoftwareSourcesSortByEnumValues Enumerates the set of values for ListSoftwarePackageSoftwareSourcesSortByEnum
func GetListSoftwarePackageSoftwareSourcesSortByEnumValues() []ListSoftwarePackageSoftwareSourcesSortByEnum {
	values := make([]ListSoftwarePackageSoftwareSourcesSortByEnum, 0)
	for _, v := range mappingListSoftwarePackageSoftwareSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSoftwarePackageSoftwareSourcesSortByEnumStringValues Enumerates the set of values in String for ListSoftwarePackageSoftwareSourcesSortByEnum
func GetListSoftwarePackageSoftwareSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSoftwarePackageSoftwareSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSoftwarePackageSoftwareSourcesSortByEnum(val string) (ListSoftwarePackageSoftwareSourcesSortByEnum, bool) {
	enum, ok := mappingListSoftwarePackageSoftwareSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
