// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListJavaLicenseAcceptanceRecordsRequest wrapper for the ListJavaLicenseAcceptanceRecords operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaLicenseAcceptanceRecords.go.html to see an example of how to use ListJavaLicenseAcceptanceRecordsRequest.
type ListJavaLicenseAcceptanceRecordsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the user principal detail.
	// The search string can be any of the property values from the Principal object.
	// This object is used as a response datatype for the `createdBy` and `lastUpdatedBy` fields in applicable resource.
	SearchByUser *string `mandatory:"false" contributesTo:"query" name:"searchByUser"`

	// Unique Java license acceptance record identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Unique Java license type.
	LicenseType ListJavaLicenseAcceptanceRecordsLicenseTypeEnum `mandatory:"false" contributesTo:"query" name:"licenseType" omitEmpty:"true"`

	// The status of license acceptance.
	Status ListJavaLicenseAcceptanceRecordsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListJavaLicenseAcceptanceRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If no value is specified, _timeAccepted_ is the default.
	SortBy ListJavaLicenseAcceptanceRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJavaLicenseAcceptanceRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJavaLicenseAcceptanceRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJavaLicenseAcceptanceRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJavaLicenseAcceptanceRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListJavaLicenseAcceptanceRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListJavaLicenseAcceptanceRecordsLicenseTypeEnum(string(request.LicenseType)); !ok && request.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", request.LicenseType, strings.Join(GetListJavaLicenseAcceptanceRecordsLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaLicenseAcceptanceRecordsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListJavaLicenseAcceptanceRecordsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaLicenseAcceptanceRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListJavaLicenseAcceptanceRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListJavaLicenseAcceptanceRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListJavaLicenseAcceptanceRecordsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListJavaLicenseAcceptanceRecordsResponse wrapper for the ListJavaLicenseAcceptanceRecords operation
type ListJavaLicenseAcceptanceRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JavaLicenseAcceptanceRecordCollection instances
	JavaLicenseAcceptanceRecordCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJavaLicenseAcceptanceRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJavaLicenseAcceptanceRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJavaLicenseAcceptanceRecordsLicenseTypeEnum Enum with underlying type: string
type ListJavaLicenseAcceptanceRecordsLicenseTypeEnum string

// Set of constants representing the allowable values for ListJavaLicenseAcceptanceRecordsLicenseTypeEnum
const (
	ListJavaLicenseAcceptanceRecordsLicenseTypeOtn        ListJavaLicenseAcceptanceRecordsLicenseTypeEnum = "OTN"
	ListJavaLicenseAcceptanceRecordsLicenseTypeNftc       ListJavaLicenseAcceptanceRecordsLicenseTypeEnum = "NFTC"
	ListJavaLicenseAcceptanceRecordsLicenseTypeRestricted ListJavaLicenseAcceptanceRecordsLicenseTypeEnum = "RESTRICTED"
)

var mappingListJavaLicenseAcceptanceRecordsLicenseTypeEnum = map[string]ListJavaLicenseAcceptanceRecordsLicenseTypeEnum{
	"OTN":        ListJavaLicenseAcceptanceRecordsLicenseTypeOtn,
	"NFTC":       ListJavaLicenseAcceptanceRecordsLicenseTypeNftc,
	"RESTRICTED": ListJavaLicenseAcceptanceRecordsLicenseTypeRestricted,
}

var mappingListJavaLicenseAcceptanceRecordsLicenseTypeEnumLowerCase = map[string]ListJavaLicenseAcceptanceRecordsLicenseTypeEnum{
	"otn":        ListJavaLicenseAcceptanceRecordsLicenseTypeOtn,
	"nftc":       ListJavaLicenseAcceptanceRecordsLicenseTypeNftc,
	"restricted": ListJavaLicenseAcceptanceRecordsLicenseTypeRestricted,
}

// GetListJavaLicenseAcceptanceRecordsLicenseTypeEnumValues Enumerates the set of values for ListJavaLicenseAcceptanceRecordsLicenseTypeEnum
func GetListJavaLicenseAcceptanceRecordsLicenseTypeEnumValues() []ListJavaLicenseAcceptanceRecordsLicenseTypeEnum {
	values := make([]ListJavaLicenseAcceptanceRecordsLicenseTypeEnum, 0)
	for _, v := range mappingListJavaLicenseAcceptanceRecordsLicenseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaLicenseAcceptanceRecordsLicenseTypeEnumStringValues Enumerates the set of values in String for ListJavaLicenseAcceptanceRecordsLicenseTypeEnum
func GetListJavaLicenseAcceptanceRecordsLicenseTypeEnumStringValues() []string {
	return []string{
		"OTN",
		"NFTC",
		"RESTRICTED",
	}
}

// GetMappingListJavaLicenseAcceptanceRecordsLicenseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaLicenseAcceptanceRecordsLicenseTypeEnum(val string) (ListJavaLicenseAcceptanceRecordsLicenseTypeEnum, bool) {
	enum, ok := mappingListJavaLicenseAcceptanceRecordsLicenseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaLicenseAcceptanceRecordsStatusEnum Enum with underlying type: string
type ListJavaLicenseAcceptanceRecordsStatusEnum string

// Set of constants representing the allowable values for ListJavaLicenseAcceptanceRecordsStatusEnum
const (
	ListJavaLicenseAcceptanceRecordsStatusAccepted ListJavaLicenseAcceptanceRecordsStatusEnum = "ACCEPTED"
	ListJavaLicenseAcceptanceRecordsStatusRevoked  ListJavaLicenseAcceptanceRecordsStatusEnum = "REVOKED"
)

var mappingListJavaLicenseAcceptanceRecordsStatusEnum = map[string]ListJavaLicenseAcceptanceRecordsStatusEnum{
	"ACCEPTED": ListJavaLicenseAcceptanceRecordsStatusAccepted,
	"REVOKED":  ListJavaLicenseAcceptanceRecordsStatusRevoked,
}

var mappingListJavaLicenseAcceptanceRecordsStatusEnumLowerCase = map[string]ListJavaLicenseAcceptanceRecordsStatusEnum{
	"accepted": ListJavaLicenseAcceptanceRecordsStatusAccepted,
	"revoked":  ListJavaLicenseAcceptanceRecordsStatusRevoked,
}

// GetListJavaLicenseAcceptanceRecordsStatusEnumValues Enumerates the set of values for ListJavaLicenseAcceptanceRecordsStatusEnum
func GetListJavaLicenseAcceptanceRecordsStatusEnumValues() []ListJavaLicenseAcceptanceRecordsStatusEnum {
	values := make([]ListJavaLicenseAcceptanceRecordsStatusEnum, 0)
	for _, v := range mappingListJavaLicenseAcceptanceRecordsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaLicenseAcceptanceRecordsStatusEnumStringValues Enumerates the set of values in String for ListJavaLicenseAcceptanceRecordsStatusEnum
func GetListJavaLicenseAcceptanceRecordsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"REVOKED",
	}
}

// GetMappingListJavaLicenseAcceptanceRecordsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaLicenseAcceptanceRecordsStatusEnum(val string) (ListJavaLicenseAcceptanceRecordsStatusEnum, bool) {
	enum, ok := mappingListJavaLicenseAcceptanceRecordsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaLicenseAcceptanceRecordsSortOrderEnum Enum with underlying type: string
type ListJavaLicenseAcceptanceRecordsSortOrderEnum string

// Set of constants representing the allowable values for ListJavaLicenseAcceptanceRecordsSortOrderEnum
const (
	ListJavaLicenseAcceptanceRecordsSortOrderAsc  ListJavaLicenseAcceptanceRecordsSortOrderEnum = "ASC"
	ListJavaLicenseAcceptanceRecordsSortOrderDesc ListJavaLicenseAcceptanceRecordsSortOrderEnum = "DESC"
)

var mappingListJavaLicenseAcceptanceRecordsSortOrderEnum = map[string]ListJavaLicenseAcceptanceRecordsSortOrderEnum{
	"ASC":  ListJavaLicenseAcceptanceRecordsSortOrderAsc,
	"DESC": ListJavaLicenseAcceptanceRecordsSortOrderDesc,
}

var mappingListJavaLicenseAcceptanceRecordsSortOrderEnumLowerCase = map[string]ListJavaLicenseAcceptanceRecordsSortOrderEnum{
	"asc":  ListJavaLicenseAcceptanceRecordsSortOrderAsc,
	"desc": ListJavaLicenseAcceptanceRecordsSortOrderDesc,
}

// GetListJavaLicenseAcceptanceRecordsSortOrderEnumValues Enumerates the set of values for ListJavaLicenseAcceptanceRecordsSortOrderEnum
func GetListJavaLicenseAcceptanceRecordsSortOrderEnumValues() []ListJavaLicenseAcceptanceRecordsSortOrderEnum {
	values := make([]ListJavaLicenseAcceptanceRecordsSortOrderEnum, 0)
	for _, v := range mappingListJavaLicenseAcceptanceRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaLicenseAcceptanceRecordsSortOrderEnumStringValues Enumerates the set of values in String for ListJavaLicenseAcceptanceRecordsSortOrderEnum
func GetListJavaLicenseAcceptanceRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListJavaLicenseAcceptanceRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaLicenseAcceptanceRecordsSortOrderEnum(val string) (ListJavaLicenseAcceptanceRecordsSortOrderEnum, bool) {
	enum, ok := mappingListJavaLicenseAcceptanceRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListJavaLicenseAcceptanceRecordsSortByEnum Enum with underlying type: string
type ListJavaLicenseAcceptanceRecordsSortByEnum string

// Set of constants representing the allowable values for ListJavaLicenseAcceptanceRecordsSortByEnum
const (
	ListJavaLicenseAcceptanceRecordsSortByTimeaccepted            ListJavaLicenseAcceptanceRecordsSortByEnum = "timeAccepted"
	ListJavaLicenseAcceptanceRecordsSortByTimelastupdated         ListJavaLicenseAcceptanceRecordsSortByEnum = "timeLastUpdated"
	ListJavaLicenseAcceptanceRecordsSortByLicenseacceptancestatus ListJavaLicenseAcceptanceRecordsSortByEnum = "licenseAcceptanceStatus"
)

var mappingListJavaLicenseAcceptanceRecordsSortByEnum = map[string]ListJavaLicenseAcceptanceRecordsSortByEnum{
	"timeAccepted":            ListJavaLicenseAcceptanceRecordsSortByTimeaccepted,
	"timeLastUpdated":         ListJavaLicenseAcceptanceRecordsSortByTimelastupdated,
	"licenseAcceptanceStatus": ListJavaLicenseAcceptanceRecordsSortByLicenseacceptancestatus,
}

var mappingListJavaLicenseAcceptanceRecordsSortByEnumLowerCase = map[string]ListJavaLicenseAcceptanceRecordsSortByEnum{
	"timeaccepted":            ListJavaLicenseAcceptanceRecordsSortByTimeaccepted,
	"timelastupdated":         ListJavaLicenseAcceptanceRecordsSortByTimelastupdated,
	"licenseacceptancestatus": ListJavaLicenseAcceptanceRecordsSortByLicenseacceptancestatus,
}

// GetListJavaLicenseAcceptanceRecordsSortByEnumValues Enumerates the set of values for ListJavaLicenseAcceptanceRecordsSortByEnum
func GetListJavaLicenseAcceptanceRecordsSortByEnumValues() []ListJavaLicenseAcceptanceRecordsSortByEnum {
	values := make([]ListJavaLicenseAcceptanceRecordsSortByEnum, 0)
	for _, v := range mappingListJavaLicenseAcceptanceRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListJavaLicenseAcceptanceRecordsSortByEnumStringValues Enumerates the set of values in String for ListJavaLicenseAcceptanceRecordsSortByEnum
func GetListJavaLicenseAcceptanceRecordsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
		"timeLastUpdated",
		"licenseAcceptanceStatus",
	}
}

// GetMappingListJavaLicenseAcceptanceRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListJavaLicenseAcceptanceRecordsSortByEnum(val string) (ListJavaLicenseAcceptanceRecordsSortByEnum, bool) {
	enum, ok := mappingListJavaLicenseAcceptanceRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
