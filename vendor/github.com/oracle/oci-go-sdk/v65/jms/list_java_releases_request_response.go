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

// ListJavaReleasesRequest wrapper for the ListJavaReleases operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListJavaReleases.go.html to see an example of how to use ListJavaReleasesRequest.
type ListJavaReleasesRequest struct {

	// Unique Java release version identifier
	ReleaseVersion *string `mandatory:"false" contributesTo:"query" name:"releaseVersion"`

	// The version identifier for the Java family.
	FamilyVersion *string `mandatory:"false" contributesTo:"query" name:"familyVersion"`

	// Java release type.
	ReleaseType ListJavaReleasesReleaseTypeEnum `mandatory:"false" contributesTo:"query" name:"releaseType" omitEmpty:"true"`

	// The security status of the Java Runtime.
	JreSecurityStatus ListJavaReleasesJreSecurityStatusEnum `mandatory:"false" contributesTo:"query" name:"jreSecurityStatus" omitEmpty:"true"`

	// Java license type.
	LicenseType ListJavaReleasesLicenseTypeEnum `mandatory:"false" contributesTo:"query" name:"licenseType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJavaReleasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If no value is specified _releaseDate_ is default.
	SortBy ListJavaReleasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaReleasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaReleasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaReleasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaReleasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaReleasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaReleasesReleaseTypeEnum(string(request.ReleaseType)); !ok && request.ReleaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReleaseType: %s. Supported values are: %s.", request.ReleaseType, strings.Join(GetListJavaReleasesReleaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaReleasesJreSecurityStatusEnum(string(request.JreSecurityStatus)); !ok && request.JreSecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JreSecurityStatus: %s. Supported values are: %s.", request.JreSecurityStatus, strings.Join(GetListJavaReleasesJreSecurityStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaReleasesLicenseTypeEnum(string(request.LicenseType)); !ok && request.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", request.LicenseType, strings.Join(GetListJavaReleasesLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaReleasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaReleasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaReleasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaReleasesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaReleasesResponse wrapper for the ListJavaReleases operation
type ListJavaReleasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaReleaseCollection instances
	JavaReleaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaReleasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaReleasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaReleasesReleaseTypeEnum Enum with underlying type: string
type ListJavaReleasesReleaseTypeEnum string

// Set of constants representing the allowable values for ListJavaReleasesReleaseTypeEnum
const (
	ListJavaReleasesReleaseTypeCpu          ListJavaReleasesReleaseTypeEnum = "CPU"
	ListJavaReleasesReleaseTypeFeature      ListJavaReleasesReleaseTypeEnum = "FEATURE"
	ListJavaReleasesReleaseTypeBpr          ListJavaReleasesReleaseTypeEnum = "BPR"
	ListJavaReleasesReleaseTypePatchRelease ListJavaReleasesReleaseTypeEnum = "PATCH_RELEASE"
)

var mappingListJavaReleasesReleaseTypeEnum = map[string]ListJavaReleasesReleaseTypeEnum{
	"CPU":           ListJavaReleasesReleaseTypeCpu,
	"FEATURE":       ListJavaReleasesReleaseTypeFeature,
	"BPR":           ListJavaReleasesReleaseTypeBpr,
	"PATCH_RELEASE": ListJavaReleasesReleaseTypePatchRelease,
}

var mappingListJavaReleasesReleaseTypeEnumLowerCase = map[string]ListJavaReleasesReleaseTypeEnum{
	"cpu":           ListJavaReleasesReleaseTypeCpu,
	"feature":       ListJavaReleasesReleaseTypeFeature,
	"bpr":           ListJavaReleasesReleaseTypeBpr,
	"patch_release": ListJavaReleasesReleaseTypePatchRelease,
}

// GetListJavaReleasesReleaseTypeEnumValues Enumerates the set of values for ListJavaReleasesReleaseTypeEnum
func GetListJavaReleasesReleaseTypeEnumValues() []ListJavaReleasesReleaseTypeEnum {
	values := make([]ListJavaReleasesReleaseTypeEnum, 0)
	for _, v := range mappingListJavaReleasesReleaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaReleasesReleaseTypeEnumStringValues Enumerates the set of values in String for ListJavaReleasesReleaseTypeEnum
func GetListJavaReleasesReleaseTypeEnumStringValues() []string {
	return []string{
		"CPU",
		"FEATURE",
		"BPR",
		"PATCH_RELEASE",
	}
}

// GetMappingListJavaReleasesReleaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaReleasesReleaseTypeEnum(val string) (ListJavaReleasesReleaseTypeEnum, bool) {
	enum, ok := mappingListJavaReleasesReleaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaReleasesJreSecurityStatusEnum Enum with underlying type: string
type ListJavaReleasesJreSecurityStatusEnum string

// Set of constants representing the allowable values for ListJavaReleasesJreSecurityStatusEnum
const (
	ListJavaReleasesJreSecurityStatusEarlyAccess     ListJavaReleasesJreSecurityStatusEnum = "EARLY_ACCESS"
	ListJavaReleasesJreSecurityStatusUnknown         ListJavaReleasesJreSecurityStatusEnum = "UNKNOWN"
	ListJavaReleasesJreSecurityStatusUpToDate        ListJavaReleasesJreSecurityStatusEnum = "UP_TO_DATE"
	ListJavaReleasesJreSecurityStatusUpdateRequired  ListJavaReleasesJreSecurityStatusEnum = "UPDATE_REQUIRED"
	ListJavaReleasesJreSecurityStatusUpgradeRequired ListJavaReleasesJreSecurityStatusEnum = "UPGRADE_REQUIRED"
)

var mappingListJavaReleasesJreSecurityStatusEnum = map[string]ListJavaReleasesJreSecurityStatusEnum{
	"EARLY_ACCESS":     ListJavaReleasesJreSecurityStatusEarlyAccess,
	"UNKNOWN":          ListJavaReleasesJreSecurityStatusUnknown,
	"UP_TO_DATE":       ListJavaReleasesJreSecurityStatusUpToDate,
	"UPDATE_REQUIRED":  ListJavaReleasesJreSecurityStatusUpdateRequired,
	"UPGRADE_REQUIRED": ListJavaReleasesJreSecurityStatusUpgradeRequired,
}

var mappingListJavaReleasesJreSecurityStatusEnumLowerCase = map[string]ListJavaReleasesJreSecurityStatusEnum{
	"early_access":     ListJavaReleasesJreSecurityStatusEarlyAccess,
	"unknown":          ListJavaReleasesJreSecurityStatusUnknown,
	"up_to_date":       ListJavaReleasesJreSecurityStatusUpToDate,
	"update_required":  ListJavaReleasesJreSecurityStatusUpdateRequired,
	"upgrade_required": ListJavaReleasesJreSecurityStatusUpgradeRequired,
}

// GetListJavaReleasesJreSecurityStatusEnumValues Enumerates the set of values for ListJavaReleasesJreSecurityStatusEnum
func GetListJavaReleasesJreSecurityStatusEnumValues() []ListJavaReleasesJreSecurityStatusEnum {
	values := make([]ListJavaReleasesJreSecurityStatusEnum, 0)
	for _, v := range mappingListJavaReleasesJreSecurityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaReleasesJreSecurityStatusEnumStringValues Enumerates the set of values in String for ListJavaReleasesJreSecurityStatusEnum
func GetListJavaReleasesJreSecurityStatusEnumStringValues() []string {
	return []string{
		"EARLY_ACCESS",
		"UNKNOWN",
		"UP_TO_DATE",
		"UPDATE_REQUIRED",
		"UPGRADE_REQUIRED",
	}
}

// GetMappingListJavaReleasesJreSecurityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaReleasesJreSecurityStatusEnum(val string) (ListJavaReleasesJreSecurityStatusEnum, bool) {
	enum, ok := mappingListJavaReleasesJreSecurityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaReleasesLicenseTypeEnum Enum with underlying type: string
type ListJavaReleasesLicenseTypeEnum string

// Set of constants representing the allowable values for ListJavaReleasesLicenseTypeEnum
const (
	ListJavaReleasesLicenseTypeOtn        ListJavaReleasesLicenseTypeEnum = "OTN"
	ListJavaReleasesLicenseTypeNftc       ListJavaReleasesLicenseTypeEnum = "NFTC"
	ListJavaReleasesLicenseTypeRestricted ListJavaReleasesLicenseTypeEnum = "RESTRICTED"
)

var mappingListJavaReleasesLicenseTypeEnum = map[string]ListJavaReleasesLicenseTypeEnum{
	"OTN":        ListJavaReleasesLicenseTypeOtn,
	"NFTC":       ListJavaReleasesLicenseTypeNftc,
	"RESTRICTED": ListJavaReleasesLicenseTypeRestricted,
}

var mappingListJavaReleasesLicenseTypeEnumLowerCase = map[string]ListJavaReleasesLicenseTypeEnum{
	"otn":        ListJavaReleasesLicenseTypeOtn,
	"nftc":       ListJavaReleasesLicenseTypeNftc,
	"restricted": ListJavaReleasesLicenseTypeRestricted,
}

// GetListJavaReleasesLicenseTypeEnumValues Enumerates the set of values for ListJavaReleasesLicenseTypeEnum
func GetListJavaReleasesLicenseTypeEnumValues() []ListJavaReleasesLicenseTypeEnum {
	values := make([]ListJavaReleasesLicenseTypeEnum, 0)
	for _, v := range mappingListJavaReleasesLicenseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaReleasesLicenseTypeEnumStringValues Enumerates the set of values in String for ListJavaReleasesLicenseTypeEnum
func GetListJavaReleasesLicenseTypeEnumStringValues() []string {
	return []string{
		"OTN",
		"NFTC",
		"RESTRICTED",
	}
}

// GetMappingListJavaReleasesLicenseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaReleasesLicenseTypeEnum(val string) (ListJavaReleasesLicenseTypeEnum, bool) {
	enum, ok := mappingListJavaReleasesLicenseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaReleasesSortOrderEnum Enum with underlying type: string
type ListJavaReleasesSortOrderEnum string

// Set of constants representing the allowable values for ListJavaReleasesSortOrderEnum
const (
	ListJavaReleasesSortOrderAsc  ListJavaReleasesSortOrderEnum = "ASC"
	ListJavaReleasesSortOrderDesc ListJavaReleasesSortOrderEnum = "DESC"
)

var mappingListJavaReleasesSortOrderEnum = map[string]ListJavaReleasesSortOrderEnum{
	"ASC":  ListJavaReleasesSortOrderAsc,
	"DESC": ListJavaReleasesSortOrderDesc,
}

var mappingListJavaReleasesSortOrderEnumLowerCase = map[string]ListJavaReleasesSortOrderEnum{
	"asc":  ListJavaReleasesSortOrderAsc,
	"desc": ListJavaReleasesSortOrderDesc,
}

// GetListJavaReleasesSortOrderEnumValues Enumerates the set of values for ListJavaReleasesSortOrderEnum
func GetListJavaReleasesSortOrderEnumValues() []ListJavaReleasesSortOrderEnum {
	values := make([]ListJavaReleasesSortOrderEnum, 0)
	for _, v := range mappingListJavaReleasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaReleasesSortOrderEnumStringValues Enumerates the set of values in String for ListJavaReleasesSortOrderEnum
func GetListJavaReleasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaReleasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaReleasesSortOrderEnum(val string) (ListJavaReleasesSortOrderEnum, bool) {
	enum, ok := mappingListJavaReleasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaReleasesSortByEnum Enum with underlying type: string
type ListJavaReleasesSortByEnum string

// Set of constants representing the allowable values for ListJavaReleasesSortByEnum
const (
	ListJavaReleasesSortByReleasedate    ListJavaReleasesSortByEnum = "releaseDate"
	ListJavaReleasesSortByReleaseversion ListJavaReleasesSortByEnum = "releaseVersion"
	ListJavaReleasesSortByFamilyversion  ListJavaReleasesSortByEnum = "familyVersion"
	ListJavaReleasesSortByLicensetype    ListJavaReleasesSortByEnum = "licenseType"
)

var mappingListJavaReleasesSortByEnum = map[string]ListJavaReleasesSortByEnum{
	"releaseDate":    ListJavaReleasesSortByReleasedate,
	"releaseVersion": ListJavaReleasesSortByReleaseversion,
	"familyVersion":  ListJavaReleasesSortByFamilyversion,
	"licenseType":    ListJavaReleasesSortByLicensetype,
}

var mappingListJavaReleasesSortByEnumLowerCase = map[string]ListJavaReleasesSortByEnum{
	"releasedate":    ListJavaReleasesSortByReleasedate,
	"releaseversion": ListJavaReleasesSortByReleaseversion,
	"familyversion":  ListJavaReleasesSortByFamilyversion,
	"licensetype":    ListJavaReleasesSortByLicensetype,
}

// GetListJavaReleasesSortByEnumValues Enumerates the set of values for ListJavaReleasesSortByEnum
func GetListJavaReleasesSortByEnumValues() []ListJavaReleasesSortByEnum {
	values := make([]ListJavaReleasesSortByEnum, 0)
	for _, v := range mappingListJavaReleasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaReleasesSortByEnumStringValues Enumerates the set of values in String for ListJavaReleasesSortByEnum
func GetListJavaReleasesSortByEnumStringValues() []string {
	return []string{
		"releaseDate",
		"releaseVersion",
		"familyVersion",
		"licenseType",
	}
}

// GetMappingListJavaReleasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaReleasesSortByEnum(val string) (ListJavaReleasesSortByEnum, bool) {
	enum, ok := mappingListJavaReleasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
