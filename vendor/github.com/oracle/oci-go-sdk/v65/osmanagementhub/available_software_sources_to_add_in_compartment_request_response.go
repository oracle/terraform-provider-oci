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

// AvailableSoftwareSourcesToAddInCompartmentRequest wrapper for the AvailableSoftwareSourcesToAddInCompartment operation
type AvailableSoftwareSourcesToAddInCompartmentRequest struct {

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
	SortOrder AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy AvailableSoftwareSourcesToAddInCompartmentSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

func (request AvailableSoftwareSourcesToAddInCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request AvailableSoftwareSourcesToAddInCompartmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request AvailableSoftwareSourcesToAddInCompartmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request AvailableSoftwareSourcesToAddInCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request AvailableSoftwareSourcesToAddInCompartmentRequest) ValidateEnumValue() (bool, error) {
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

	if _, ok := GetMappingAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAvailableSoftwareSourcesToAddInCompartmentSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetAvailableSoftwareSourcesToAddInCompartmentSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AvailableSoftwareSourcesToAddInCompartmentResponse wrapper for the AvailableSoftwareSourcesToAddInCompartment operation
type AvailableSoftwareSourcesToAddInCompartmentResponse struct {

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

func (response AvailableSoftwareSourcesToAddInCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response AvailableSoftwareSourcesToAddInCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum Enum with underlying type: string
type AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum string

// Set of constants representing the allowable values for AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum
const (
	AvailableSoftwareSourcesToAddInCompartmentSortOrderAsc  AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum = "ASC"
	AvailableSoftwareSourcesToAddInCompartmentSortOrderDesc AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum = "DESC"
)

var mappingAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum = map[string]AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum{
	"ASC":  AvailableSoftwareSourcesToAddInCompartmentSortOrderAsc,
	"DESC": AvailableSoftwareSourcesToAddInCompartmentSortOrderDesc,
}

var mappingAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumLowerCase = map[string]AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum{
	"asc":  AvailableSoftwareSourcesToAddInCompartmentSortOrderAsc,
	"desc": AvailableSoftwareSourcesToAddInCompartmentSortOrderDesc,
}

// GetAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumValues Enumerates the set of values for AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum
func GetAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumValues() []AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum {
	values := make([]AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum, 0)
	for _, v := range mappingAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumStringValues Enumerates the set of values in String for AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum
func GetAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAvailableSoftwareSourcesToAddInCompartmentSortOrderEnum(val string) (AvailableSoftwareSourcesToAddInCompartmentSortOrderEnum, bool) {
	enum, ok := mappingAvailableSoftwareSourcesToAddInCompartmentSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AvailableSoftwareSourcesToAddInCompartmentSortByEnum Enum with underlying type: string
type AvailableSoftwareSourcesToAddInCompartmentSortByEnum string

// Set of constants representing the allowable values for AvailableSoftwareSourcesToAddInCompartmentSortByEnum
const (
	AvailableSoftwareSourcesToAddInCompartmentSortByTimecreated AvailableSoftwareSourcesToAddInCompartmentSortByEnum = "timeCreated"
	AvailableSoftwareSourcesToAddInCompartmentSortByDisplayname AvailableSoftwareSourcesToAddInCompartmentSortByEnum = "displayName"
)

var mappingAvailableSoftwareSourcesToAddInCompartmentSortByEnum = map[string]AvailableSoftwareSourcesToAddInCompartmentSortByEnum{
	"timeCreated": AvailableSoftwareSourcesToAddInCompartmentSortByTimecreated,
	"displayName": AvailableSoftwareSourcesToAddInCompartmentSortByDisplayname,
}

var mappingAvailableSoftwareSourcesToAddInCompartmentSortByEnumLowerCase = map[string]AvailableSoftwareSourcesToAddInCompartmentSortByEnum{
	"timecreated": AvailableSoftwareSourcesToAddInCompartmentSortByTimecreated,
	"displayname": AvailableSoftwareSourcesToAddInCompartmentSortByDisplayname,
}

// GetAvailableSoftwareSourcesToAddInCompartmentSortByEnumValues Enumerates the set of values for AvailableSoftwareSourcesToAddInCompartmentSortByEnum
func GetAvailableSoftwareSourcesToAddInCompartmentSortByEnumValues() []AvailableSoftwareSourcesToAddInCompartmentSortByEnum {
	values := make([]AvailableSoftwareSourcesToAddInCompartmentSortByEnum, 0)
	for _, v := range mappingAvailableSoftwareSourcesToAddInCompartmentSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetAvailableSoftwareSourcesToAddInCompartmentSortByEnumStringValues Enumerates the set of values in String for AvailableSoftwareSourcesToAddInCompartmentSortByEnum
func GetAvailableSoftwareSourcesToAddInCompartmentSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingAvailableSoftwareSourcesToAddInCompartmentSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAvailableSoftwareSourcesToAddInCompartmentSortByEnum(val string) (AvailableSoftwareSourcesToAddInCompartmentSortByEnum, bool) {
	enum, ok := mappingAvailableSoftwareSourcesToAddInCompartmentSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
