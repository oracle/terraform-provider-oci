// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAvailableWindowsUpdatesForManagedInstanceRequest wrapper for the ListAvailableWindowsUpdatesForManagedInstance operation
type ListAvailableWindowsUpdatesForManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListAvailableWindowsUpdatesForManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicator of whether the update can be installed using OSMS.
	IsEligibleForInstallation ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum `mandatory:"false" contributesTo:"query" name:"isEligibleForInstallation" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableWindowsUpdatesForManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableWindowsUpdatesForManagedInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableWindowsUpdatesForManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAvailableWindowsUpdatesForManagedInstanceResponse wrapper for the ListAvailableWindowsUpdatesForManagedInstance operation
type ListAvailableWindowsUpdatesForManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailableWindowsUpdateSummary instances
	Items []AvailableWindowsUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `AvailableWindowsUpdateSummary`s. If this header
	// appears in the response, then this is a partial list of
	// `AvailableWindowsUpdateSummary`s available to be installed on the managed instance. Include this value
	// as the `page` parameter in a subsequent
	// GET request to get the next batch of managed instances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableWindowsUpdatesForManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableWindowsUpdatesForManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum Enum with underlying type: string
type ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum
const (
	ListAvailableWindowsUpdatesForManagedInstanceSortOrderAsc  ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum = "ASC"
	ListAvailableWindowsUpdatesForManagedInstanceSortOrderDesc ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum = "DESC"
)

var mappingListAvailableWindowsUpdatesForManagedInstanceSortOrder = map[string]ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum{
	"ASC":  ListAvailableWindowsUpdatesForManagedInstanceSortOrderAsc,
	"DESC": ListAvailableWindowsUpdatesForManagedInstanceSortOrderDesc,
}

// GetListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumValues Enumerates the set of values for ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum
func GetListAvailableWindowsUpdatesForManagedInstanceSortOrderEnumValues() []ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum {
	values := make([]ListAvailableWindowsUpdatesForManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListAvailableWindowsUpdatesForManagedInstanceSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAvailableWindowsUpdatesForManagedInstanceSortByEnum Enum with underlying type: string
type ListAvailableWindowsUpdatesForManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListAvailableWindowsUpdatesForManagedInstanceSortByEnum
const (
	ListAvailableWindowsUpdatesForManagedInstanceSortByTimecreated ListAvailableWindowsUpdatesForManagedInstanceSortByEnum = "TIMECREATED"
	ListAvailableWindowsUpdatesForManagedInstanceSortByDisplayname ListAvailableWindowsUpdatesForManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListAvailableWindowsUpdatesForManagedInstanceSortBy = map[string]ListAvailableWindowsUpdatesForManagedInstanceSortByEnum{
	"TIMECREATED": ListAvailableWindowsUpdatesForManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListAvailableWindowsUpdatesForManagedInstanceSortByDisplayname,
}

// GetListAvailableWindowsUpdatesForManagedInstanceSortByEnumValues Enumerates the set of values for ListAvailableWindowsUpdatesForManagedInstanceSortByEnum
func GetListAvailableWindowsUpdatesForManagedInstanceSortByEnumValues() []ListAvailableWindowsUpdatesForManagedInstanceSortByEnum {
	values := make([]ListAvailableWindowsUpdatesForManagedInstanceSortByEnum, 0)
	for _, v := range mappingListAvailableWindowsUpdatesForManagedInstanceSortBy {
		values = append(values, v)
	}
	return values
}

// ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum Enum with underlying type: string
type ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum string

// Set of constants representing the allowable values for ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum
const (
	ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationInstallable    ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum = "INSTALLABLE"
	ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationNotInstallable ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum = "NOT_INSTALLABLE"
	ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationUnknown        ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum = "UNKNOWN"
)

var mappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallation = map[string]ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum{
	"INSTALLABLE":     ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationInstallable,
	"NOT_INSTALLABLE": ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationNotInstallable,
	"UNKNOWN":         ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationUnknown,
}

// GetListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumValues Enumerates the set of values for ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum
func GetListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnumValues() []ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum {
	values := make([]ListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallationEnum, 0)
	for _, v := range mappingListAvailableWindowsUpdatesForManagedInstanceIsEligibleForInstallation {
		values = append(values, v)
	}
	return values
}
