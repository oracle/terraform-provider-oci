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

// ListResourceProfilesRequest wrapper for the ListResourceProfiles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceProfiles.go.html to see an example of how to use ListResourceProfilesRequest.
type ListResourceProfilesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Start time for a filter. If start time is not specified, start time will be set to current time - 30 days.
	TimeLastDetectedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastDetectedGreaterThanOrEqualTo"`

	// End time for a filter. If end time is not specified, end time will be set to current time.
	TimeLastDetectedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastDetectedLessThanOrEqualTo"`

	// Start time for a filter. If start time is not specified, start time will be set to current time - 30 days.
	TimeFirstDetectedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFirstDetectedGreaterThanOrEqualTo"`

	// End time for a filter. If end time is not specified, end time will be set to current time.
	TimeFirstDetectedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeFirstDetectedLessThanOrEqualTo"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListResourceProfilesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the list of resource types given
	ResourceTypes []string `contributesTo:"query" name:"resourceTypes" collectionFormat:"multi"`

	// risk score filter
	RiskScoreGreaterThanOrEqualTo *float64 `mandatory:"false" contributesTo:"query" name:"riskScoreGreaterThanOrEqualTo"`

	// risk score filter
	RiskScoreLessThanOrEqualTo *float64 `mandatory:"false" contributesTo:"query" name:"riskScoreLessThanOrEqualTo"`

	// A filter to return only resources that match the list of techniques given
	Techniques []string `contributesTo:"query" name:"techniques" collectionFormat:"multi"`

	// A filter to return only resources that match the list of tactics given.
	Tactics []string `contributesTo:"query" name:"tactics" collectionFormat:"multi"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResourceProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort resource profiles. Only one sort order may be provided. Default order for timeLastDetected is descending. If no value is specified timeLastDetected is default.
	SortBy ListResourceProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceProfilesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListResourceProfilesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceProfilesResponse wrapper for the ListResourceProfiles operation
type ListResourceProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceProfileCollection instances
	ResourceProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceProfilesAccessLevelEnum Enum with underlying type: string
type ListResourceProfilesAccessLevelEnum string

// Set of constants representing the allowable values for ListResourceProfilesAccessLevelEnum
const (
	ListResourceProfilesAccessLevelRestricted ListResourceProfilesAccessLevelEnum = "RESTRICTED"
	ListResourceProfilesAccessLevelAccessible ListResourceProfilesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListResourceProfilesAccessLevelEnum = map[string]ListResourceProfilesAccessLevelEnum{
	"RESTRICTED": ListResourceProfilesAccessLevelRestricted,
	"ACCESSIBLE": ListResourceProfilesAccessLevelAccessible,
}

var mappingListResourceProfilesAccessLevelEnumLowerCase = map[string]ListResourceProfilesAccessLevelEnum{
	"restricted": ListResourceProfilesAccessLevelRestricted,
	"accessible": ListResourceProfilesAccessLevelAccessible,
}

// GetListResourceProfilesAccessLevelEnumValues Enumerates the set of values for ListResourceProfilesAccessLevelEnum
func GetListResourceProfilesAccessLevelEnumValues() []ListResourceProfilesAccessLevelEnum {
	values := make([]ListResourceProfilesAccessLevelEnum, 0)
	for _, v := range mappingListResourceProfilesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceProfilesAccessLevelEnumStringValues Enumerates the set of values in String for ListResourceProfilesAccessLevelEnum
func GetListResourceProfilesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListResourceProfilesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceProfilesAccessLevelEnum(val string) (ListResourceProfilesAccessLevelEnum, bool) {
	enum, ok := mappingListResourceProfilesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceProfilesSortOrderEnum Enum with underlying type: string
type ListResourceProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListResourceProfilesSortOrderEnum
const (
	ListResourceProfilesSortOrderAsc  ListResourceProfilesSortOrderEnum = "ASC"
	ListResourceProfilesSortOrderDesc ListResourceProfilesSortOrderEnum = "DESC"
)

var mappingListResourceProfilesSortOrderEnum = map[string]ListResourceProfilesSortOrderEnum{
	"ASC":  ListResourceProfilesSortOrderAsc,
	"DESC": ListResourceProfilesSortOrderDesc,
}

var mappingListResourceProfilesSortOrderEnumLowerCase = map[string]ListResourceProfilesSortOrderEnum{
	"asc":  ListResourceProfilesSortOrderAsc,
	"desc": ListResourceProfilesSortOrderDesc,
}

// GetListResourceProfilesSortOrderEnumValues Enumerates the set of values for ListResourceProfilesSortOrderEnum
func GetListResourceProfilesSortOrderEnumValues() []ListResourceProfilesSortOrderEnum {
	values := make([]ListResourceProfilesSortOrderEnum, 0)
	for _, v := range mappingListResourceProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListResourceProfilesSortOrderEnum
func GetListResourceProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceProfilesSortOrderEnum(val string) (ListResourceProfilesSortOrderEnum, bool) {
	enum, ok := mappingListResourceProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceProfilesSortByEnum Enum with underlying type: string
type ListResourceProfilesSortByEnum string

// Set of constants representing the allowable values for ListResourceProfilesSortByEnum
const (
	ListResourceProfilesSortByRiskscore         ListResourceProfilesSortByEnum = "riskScore"
	ListResourceProfilesSortByRiskscoregrowth   ListResourceProfilesSortByEnum = "riskScoreGrowth"
	ListResourceProfilesSortByTimefirstdetected ListResourceProfilesSortByEnum = "timeFirstDetected"
	ListResourceProfilesSortByTimelastdetected  ListResourceProfilesSortByEnum = "timeLastDetected"
	ListResourceProfilesSortBySightingscount    ListResourceProfilesSortByEnum = "sightingsCount"
	ListResourceProfilesSortByDisplayname       ListResourceProfilesSortByEnum = "displayName"
	ListResourceProfilesSortByType              ListResourceProfilesSortByEnum = "type"
)

var mappingListResourceProfilesSortByEnum = map[string]ListResourceProfilesSortByEnum{
	"riskScore":         ListResourceProfilesSortByRiskscore,
	"riskScoreGrowth":   ListResourceProfilesSortByRiskscoregrowth,
	"timeFirstDetected": ListResourceProfilesSortByTimefirstdetected,
	"timeLastDetected":  ListResourceProfilesSortByTimelastdetected,
	"sightingsCount":    ListResourceProfilesSortBySightingscount,
	"displayName":       ListResourceProfilesSortByDisplayname,
	"type":              ListResourceProfilesSortByType,
}

var mappingListResourceProfilesSortByEnumLowerCase = map[string]ListResourceProfilesSortByEnum{
	"riskscore":         ListResourceProfilesSortByRiskscore,
	"riskscoregrowth":   ListResourceProfilesSortByRiskscoregrowth,
	"timefirstdetected": ListResourceProfilesSortByTimefirstdetected,
	"timelastdetected":  ListResourceProfilesSortByTimelastdetected,
	"sightingscount":    ListResourceProfilesSortBySightingscount,
	"displayname":       ListResourceProfilesSortByDisplayname,
	"type":              ListResourceProfilesSortByType,
}

// GetListResourceProfilesSortByEnumValues Enumerates the set of values for ListResourceProfilesSortByEnum
func GetListResourceProfilesSortByEnumValues() []ListResourceProfilesSortByEnum {
	values := make([]ListResourceProfilesSortByEnum, 0)
	for _, v := range mappingListResourceProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceProfilesSortByEnumStringValues Enumerates the set of values in String for ListResourceProfilesSortByEnum
func GetListResourceProfilesSortByEnumStringValues() []string {
	return []string{
		"riskScore",
		"riskScoreGrowth",
		"timeFirstDetected",
		"timeLastDetected",
		"sightingsCount",
		"displayName",
		"type",
	}
}

// GetMappingListResourceProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceProfilesSortByEnum(val string) (ListResourceProfilesSortByEnum, bool) {
	enum, ok := mappingListResourceProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
