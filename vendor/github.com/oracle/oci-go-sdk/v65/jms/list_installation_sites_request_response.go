// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInstallationSitesRequest wrapper for the ListInstallationSites operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListInstallationSites.go.html to see an example of how to use ListInstallationSitesRequest.
type ListInstallationSitesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The vendor of the related Java Runtime.
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the related Java Runtime.
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the related Java Runtime.
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The file system path of the installation.
	InstallationPath *string `mandatory:"false" contributesTo:"query" name:"installationPath"`

	// The Fleet-unique identifier of the related application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListInstallationSitesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort installation sites. Only one sort order may be provided.
	// Default order for _timeLastSeen_, and _jreVersion_, _approximateApplicationCount_ is **descending**.
	// Default order for _managedInstanceId_, _jreDistribution_, _jreVendor_ and _osName_ is **ascending**.
	// If no value is specified _managedInstanceId_ is default.
	SortBy ListInstallationSitesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The operating system type.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// The security status of the Java Runtime.
	JreSecurityStatus ListInstallationSitesJreSecurityStatusEnum `mandatory:"false" contributesTo:"query" name:"jreSecurityStatus" omitEmpty:"true"`

	// Filter the list with path contains the given value.
	PathContains *string `mandatory:"false" contributesTo:"query" name:"pathContains"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInstallationSitesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInstallationSitesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInstallationSitesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInstallationSitesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInstallationSitesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInstallationSitesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInstallationSitesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInstallationSitesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInstallationSitesSortByEnumStringValues(), ",")))
	}
	for _, val := range request.OsFamily {
		if _, ok := GetMappingOsFamilyEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", val, strings.Join(GetOsFamilyEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListInstallationSitesJreSecurityStatusEnum(string(request.JreSecurityStatus)); !ok && request.JreSecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JreSecurityStatus: %s. Supported values are: %s.", request.JreSecurityStatus, strings.Join(GetListInstallationSitesJreSecurityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInstallationSitesResponse wrapper for the ListInstallationSites operation
type ListInstallationSitesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InstallationSiteCollection instances
	InstallationSiteCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInstallationSitesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInstallationSitesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInstallationSitesSortOrderEnum Enum with underlying type: string
type ListInstallationSitesSortOrderEnum string

// Set of constants representing the allowable values for ListInstallationSitesSortOrderEnum
const (
	ListInstallationSitesSortOrderAsc  ListInstallationSitesSortOrderEnum = "ASC"
	ListInstallationSitesSortOrderDesc ListInstallationSitesSortOrderEnum = "DESC"
)

var mappingListInstallationSitesSortOrderEnum = map[string]ListInstallationSitesSortOrderEnum{
	"ASC":  ListInstallationSitesSortOrderAsc,
	"DESC": ListInstallationSitesSortOrderDesc,
}

var mappingListInstallationSitesSortOrderEnumLowerCase = map[string]ListInstallationSitesSortOrderEnum{
	"asc":  ListInstallationSitesSortOrderAsc,
	"desc": ListInstallationSitesSortOrderDesc,
}

// GetListInstallationSitesSortOrderEnumValues Enumerates the set of values for ListInstallationSitesSortOrderEnum
func GetListInstallationSitesSortOrderEnumValues() []ListInstallationSitesSortOrderEnum {
	values := make([]ListInstallationSitesSortOrderEnum, 0)
	for _, v := range mappingListInstallationSitesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstallationSitesSortOrderEnumStringValues Enumerates the set of values in String for ListInstallationSitesSortOrderEnum
func GetListInstallationSitesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInstallationSitesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstallationSitesSortOrderEnum(val string) (ListInstallationSitesSortOrderEnum, bool) {
	enum, ok := mappingListInstallationSitesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstallationSitesSortByEnum Enum with underlying type: string
type ListInstallationSitesSortByEnum string

// Set of constants representing the allowable values for ListInstallationSitesSortByEnum
const (
	ListInstallationSitesSortByManagedinstanceid           ListInstallationSitesSortByEnum = "managedInstanceId"
	ListInstallationSitesSortByJredistribution             ListInstallationSitesSortByEnum = "jreDistribution"
	ListInstallationSitesSortByJrevendor                   ListInstallationSitesSortByEnum = "jreVendor"
	ListInstallationSitesSortByJreversion                  ListInstallationSitesSortByEnum = "jreVersion"
	ListInstallationSitesSortByPath                        ListInstallationSitesSortByEnum = "path"
	ListInstallationSitesSortByApproximateapplicationcount ListInstallationSitesSortByEnum = "approximateApplicationCount"
	ListInstallationSitesSortByOsname                      ListInstallationSitesSortByEnum = "osName"
	ListInstallationSitesSortBySecuritystatus              ListInstallationSitesSortByEnum = "securityStatus"
)

var mappingListInstallationSitesSortByEnum = map[string]ListInstallationSitesSortByEnum{
	"managedInstanceId":           ListInstallationSitesSortByManagedinstanceid,
	"jreDistribution":             ListInstallationSitesSortByJredistribution,
	"jreVendor":                   ListInstallationSitesSortByJrevendor,
	"jreVersion":                  ListInstallationSitesSortByJreversion,
	"path":                        ListInstallationSitesSortByPath,
	"approximateApplicationCount": ListInstallationSitesSortByApproximateapplicationcount,
	"osName":                      ListInstallationSitesSortByOsname,
	"securityStatus":              ListInstallationSitesSortBySecuritystatus,
}

var mappingListInstallationSitesSortByEnumLowerCase = map[string]ListInstallationSitesSortByEnum{
	"managedinstanceid":           ListInstallationSitesSortByManagedinstanceid,
	"jredistribution":             ListInstallationSitesSortByJredistribution,
	"jrevendor":                   ListInstallationSitesSortByJrevendor,
	"jreversion":                  ListInstallationSitesSortByJreversion,
	"path":                        ListInstallationSitesSortByPath,
	"approximateapplicationcount": ListInstallationSitesSortByApproximateapplicationcount,
	"osname":                      ListInstallationSitesSortByOsname,
	"securitystatus":              ListInstallationSitesSortBySecuritystatus,
}

// GetListInstallationSitesSortByEnumValues Enumerates the set of values for ListInstallationSitesSortByEnum
func GetListInstallationSitesSortByEnumValues() []ListInstallationSitesSortByEnum {
	values := make([]ListInstallationSitesSortByEnum, 0)
	for _, v := range mappingListInstallationSitesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstallationSitesSortByEnumStringValues Enumerates the set of values in String for ListInstallationSitesSortByEnum
func GetListInstallationSitesSortByEnumStringValues() []string {
	return []string{
		"managedInstanceId",
		"jreDistribution",
		"jreVendor",
		"jreVersion",
		"path",
		"approximateApplicationCount",
		"osName",
		"securityStatus",
	}
}

// GetMappingListInstallationSitesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstallationSitesSortByEnum(val string) (ListInstallationSitesSortByEnum, bool) {
	enum, ok := mappingListInstallationSitesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInstallationSitesJreSecurityStatusEnum Enum with underlying type: string
type ListInstallationSitesJreSecurityStatusEnum string

// Set of constants representing the allowable values for ListInstallationSitesJreSecurityStatusEnum
const (
	ListInstallationSitesJreSecurityStatusEarlyAccess     ListInstallationSitesJreSecurityStatusEnum = "EARLY_ACCESS"
	ListInstallationSitesJreSecurityStatusUnknown         ListInstallationSitesJreSecurityStatusEnum = "UNKNOWN"
	ListInstallationSitesJreSecurityStatusUpToDate        ListInstallationSitesJreSecurityStatusEnum = "UP_TO_DATE"
	ListInstallationSitesJreSecurityStatusUpdateRequired  ListInstallationSitesJreSecurityStatusEnum = "UPDATE_REQUIRED"
	ListInstallationSitesJreSecurityStatusUpgradeRequired ListInstallationSitesJreSecurityStatusEnum = "UPGRADE_REQUIRED"
)

var mappingListInstallationSitesJreSecurityStatusEnum = map[string]ListInstallationSitesJreSecurityStatusEnum{
	"EARLY_ACCESS":     ListInstallationSitesJreSecurityStatusEarlyAccess,
	"UNKNOWN":          ListInstallationSitesJreSecurityStatusUnknown,
	"UP_TO_DATE":       ListInstallationSitesJreSecurityStatusUpToDate,
	"UPDATE_REQUIRED":  ListInstallationSitesJreSecurityStatusUpdateRequired,
	"UPGRADE_REQUIRED": ListInstallationSitesJreSecurityStatusUpgradeRequired,
}

var mappingListInstallationSitesJreSecurityStatusEnumLowerCase = map[string]ListInstallationSitesJreSecurityStatusEnum{
	"early_access":     ListInstallationSitesJreSecurityStatusEarlyAccess,
	"unknown":          ListInstallationSitesJreSecurityStatusUnknown,
	"up_to_date":       ListInstallationSitesJreSecurityStatusUpToDate,
	"update_required":  ListInstallationSitesJreSecurityStatusUpdateRequired,
	"upgrade_required": ListInstallationSitesJreSecurityStatusUpgradeRequired,
}

// GetListInstallationSitesJreSecurityStatusEnumValues Enumerates the set of values for ListInstallationSitesJreSecurityStatusEnum
func GetListInstallationSitesJreSecurityStatusEnumValues() []ListInstallationSitesJreSecurityStatusEnum {
	values := make([]ListInstallationSitesJreSecurityStatusEnum, 0)
	for _, v := range mappingListInstallationSitesJreSecurityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListInstallationSitesJreSecurityStatusEnumStringValues Enumerates the set of values in String for ListInstallationSitesJreSecurityStatusEnum
func GetListInstallationSitesJreSecurityStatusEnumStringValues() []string {
	return []string{
		"EARLY_ACCESS",
		"UNKNOWN",
		"UP_TO_DATE",
		"UPDATE_REQUIRED",
		"UPGRADE_REQUIRED",
	}
}

// GetMappingListInstallationSitesJreSecurityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInstallationSitesJreSecurityStatusEnum(val string) (ListInstallationSitesJreSecurityStatusEnum, bool) {
	enum, ok := mappingListInstallationSitesJreSecurityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
