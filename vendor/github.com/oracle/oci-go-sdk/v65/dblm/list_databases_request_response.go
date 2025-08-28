// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabasesRequest wrapper for the ListDatabases operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dblm/ListDatabases.go.html to see an example of how to use ListDatabasesRequest.
type ListDatabasesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState DblmVulnerabilityLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only database that match the given release version.
	DatabaseRelease *string `mandatory:"false" contributesTo:"query" name:"databaseRelease"`

	// Filter by database type.
	// Possible values Single Instance or RAC.
	DatabaseType ListDatabasesDatabaseTypeEnum `mandatory:"false" contributesTo:"query" name:"databaseType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by.
	SortBy ListDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Subscribed image
	ImageId *string `mandatory:"false" contributesTo:"query" name:"imageId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only database that have given patchId as additional patch (drifter from image version).
	DrifterPatchId *int `mandatory:"false" contributesTo:"query" name:"drifterPatchId"`

	// Filter databases by image compliance status.
	ImageCompliance ListDatabasesImageComplianceEnum `mandatory:"false" contributesTo:"query" name:"imageCompliance" omitEmpty:"true"`

	// Filter by one or more severity types.
	// Possible values are critical, high, medium, low, info and none.
	SeverityType []ResourcesSeveritiesEnum `contributesTo:"query" name:"severityType" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDblmVulnerabilityLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDblmVulnerabilityLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabasesDatabaseTypeEnum(string(request.DatabaseType)); !ok && request.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", request.DatabaseType, strings.Join(GetListDatabasesDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabasesImageComplianceEnum(string(request.ImageCompliance)); !ok && request.ImageCompliance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ImageCompliance: %s. Supported values are: %s.", request.ImageCompliance, strings.Join(GetListDatabasesImageComplianceEnumStringValues(), ",")))
	}
	for _, val := range request.SeverityType {
		if _, ok := GetMappingResourcesSeveritiesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SeverityType: %s. Supported values are: %s.", val, strings.Join(GetResourcesSeveritiesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabasesResponse wrapper for the ListDatabases operation
type ListDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatchDatabasesCollection instances
	PatchDatabasesCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabasesDatabaseTypeEnum Enum with underlying type: string
type ListDatabasesDatabaseTypeEnum string

// Set of constants representing the allowable values for ListDatabasesDatabaseTypeEnum
const (
	ListDatabasesDatabaseTypeSi  ListDatabasesDatabaseTypeEnum = "SI"
	ListDatabasesDatabaseTypeRac ListDatabasesDatabaseTypeEnum = "RAC"
)

var mappingListDatabasesDatabaseTypeEnum = map[string]ListDatabasesDatabaseTypeEnum{
	"SI":  ListDatabasesDatabaseTypeSi,
	"RAC": ListDatabasesDatabaseTypeRac,
}

var mappingListDatabasesDatabaseTypeEnumLowerCase = map[string]ListDatabasesDatabaseTypeEnum{
	"si":  ListDatabasesDatabaseTypeSi,
	"rac": ListDatabasesDatabaseTypeRac,
}

// GetListDatabasesDatabaseTypeEnumValues Enumerates the set of values for ListDatabasesDatabaseTypeEnum
func GetListDatabasesDatabaseTypeEnumValues() []ListDatabasesDatabaseTypeEnum {
	values := make([]ListDatabasesDatabaseTypeEnum, 0)
	for _, v := range mappingListDatabasesDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabasesDatabaseTypeEnumStringValues Enumerates the set of values in String for ListDatabasesDatabaseTypeEnum
func GetListDatabasesDatabaseTypeEnumStringValues() []string {
	return []string{
		"SI",
		"RAC",
	}
}

// GetMappingListDatabasesDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabasesDatabaseTypeEnum(val string) (ListDatabasesDatabaseTypeEnum, bool) {
	enum, ok := mappingListDatabasesDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabasesSortOrderEnum Enum with underlying type: string
type ListDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListDatabasesSortOrderEnum
const (
	ListDatabasesSortOrderAsc  ListDatabasesSortOrderEnum = "ASC"
	ListDatabasesSortOrderDesc ListDatabasesSortOrderEnum = "DESC"
)

var mappingListDatabasesSortOrderEnum = map[string]ListDatabasesSortOrderEnum{
	"ASC":  ListDatabasesSortOrderAsc,
	"DESC": ListDatabasesSortOrderDesc,
}

var mappingListDatabasesSortOrderEnumLowerCase = map[string]ListDatabasesSortOrderEnum{
	"asc":  ListDatabasesSortOrderAsc,
	"desc": ListDatabasesSortOrderDesc,
}

// GetListDatabasesSortOrderEnumValues Enumerates the set of values for ListDatabasesSortOrderEnum
func GetListDatabasesSortOrderEnumValues() []ListDatabasesSortOrderEnum {
	values := make([]ListDatabasesSortOrderEnum, 0)
	for _, v := range mappingListDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListDatabasesSortOrderEnum
func GetListDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabasesSortOrderEnum(val string) (ListDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabasesSortByEnum Enum with underlying type: string
type ListDatabasesSortByEnum string

// Set of constants representing the allowable values for ListDatabasesSortByEnum
const (
	ListDatabasesSortByTimecreated     ListDatabasesSortByEnum = "timeCreated"
	ListDatabasesSortByName            ListDatabasesSortByEnum = "name"
	ListDatabasesSortByResourcetype    ListDatabasesSortByEnum = "resourceType"
	ListDatabasesSortByRelease         ListDatabasesSortByEnum = "release"
	ListDatabasesSortBySubscribedimage ListDatabasesSortByEnum = "subscribedImage"
	ListDatabasesSortByPatchcompliance ListDatabasesSortByEnum = "patchCompliance"
)

var mappingListDatabasesSortByEnum = map[string]ListDatabasesSortByEnum{
	"timeCreated":     ListDatabasesSortByTimecreated,
	"name":            ListDatabasesSortByName,
	"resourceType":    ListDatabasesSortByResourcetype,
	"release":         ListDatabasesSortByRelease,
	"subscribedImage": ListDatabasesSortBySubscribedimage,
	"patchCompliance": ListDatabasesSortByPatchcompliance,
}

var mappingListDatabasesSortByEnumLowerCase = map[string]ListDatabasesSortByEnum{
	"timecreated":     ListDatabasesSortByTimecreated,
	"name":            ListDatabasesSortByName,
	"resourcetype":    ListDatabasesSortByResourcetype,
	"release":         ListDatabasesSortByRelease,
	"subscribedimage": ListDatabasesSortBySubscribedimage,
	"patchcompliance": ListDatabasesSortByPatchcompliance,
}

// GetListDatabasesSortByEnumValues Enumerates the set of values for ListDatabasesSortByEnum
func GetListDatabasesSortByEnumValues() []ListDatabasesSortByEnum {
	values := make([]ListDatabasesSortByEnum, 0)
	for _, v := range mappingListDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabasesSortByEnumStringValues Enumerates the set of values in String for ListDatabasesSortByEnum
func GetListDatabasesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"resourceType",
		"release",
		"subscribedImage",
		"patchCompliance",
	}
}

// GetMappingListDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabasesSortByEnum(val string) (ListDatabasesSortByEnum, bool) {
	enum, ok := mappingListDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabasesImageComplianceEnum Enum with underlying type: string
type ListDatabasesImageComplianceEnum string

// Set of constants representing the allowable values for ListDatabasesImageComplianceEnum
const (
	ListDatabasesImageComplianceNotSubscribed          ListDatabasesImageComplianceEnum = "NOT_SUBSCRIBED"
	ListDatabasesImageComplianceNotCompliantWithImages ListDatabasesImageComplianceEnum = "NOT_COMPLIANT_WITH_IMAGES"
	ListDatabasesImageComplianceAllDatabases           ListDatabasesImageComplianceEnum = "ALL_DATABASES"
)

var mappingListDatabasesImageComplianceEnum = map[string]ListDatabasesImageComplianceEnum{
	"NOT_SUBSCRIBED":            ListDatabasesImageComplianceNotSubscribed,
	"NOT_COMPLIANT_WITH_IMAGES": ListDatabasesImageComplianceNotCompliantWithImages,
	"ALL_DATABASES":             ListDatabasesImageComplianceAllDatabases,
}

var mappingListDatabasesImageComplianceEnumLowerCase = map[string]ListDatabasesImageComplianceEnum{
	"not_subscribed":            ListDatabasesImageComplianceNotSubscribed,
	"not_compliant_with_images": ListDatabasesImageComplianceNotCompliantWithImages,
	"all_databases":             ListDatabasesImageComplianceAllDatabases,
}

// GetListDatabasesImageComplianceEnumValues Enumerates the set of values for ListDatabasesImageComplianceEnum
func GetListDatabasesImageComplianceEnumValues() []ListDatabasesImageComplianceEnum {
	values := make([]ListDatabasesImageComplianceEnum, 0)
	for _, v := range mappingListDatabasesImageComplianceEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabasesImageComplianceEnumStringValues Enumerates the set of values in String for ListDatabasesImageComplianceEnum
func GetListDatabasesImageComplianceEnumStringValues() []string {
	return []string{
		"NOT_SUBSCRIBED",
		"NOT_COMPLIANT_WITH_IMAGES",
		"ALL_DATABASES",
	}
}

// GetMappingListDatabasesImageComplianceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabasesImageComplianceEnum(val string) (ListDatabasesImageComplianceEnum, bool) {
	enum, ok := mappingListDatabasesImageComplianceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
