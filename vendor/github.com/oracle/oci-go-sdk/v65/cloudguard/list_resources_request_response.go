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

// ListResourcesRequest wrapper for the ListResources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResources.go.html to see an example of how to use ListResourcesRequest.
type ListResourcesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the target in which to list resources.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// OCI monitoring region.
	Region *string `mandatory:"false" contributesTo:"query" name:"region"`

	// Cvss score associated with the resource.
	CvssScore *int `mandatory:"false" contributesTo:"query" name:"cvssScore"`

	// Cvss score greater than associated with the resource.
	CvssScoreGreaterThan *int `mandatory:"false" contributesTo:"query" name:"cvssScoreGreaterThan"`

	// Cvss score less than associated with the resource.
	CvssScoreLessThan *int `mandatory:"false" contributesTo:"query" name:"cvssScoreLessThan"`

	// CVE ID associated with the resource.
	CveId *string `mandatory:"false" contributesTo:"query" name:"cveId"`

	// Risk level of the problem.
	RiskLevel *string `mandatory:"false" contributesTo:"query" name:"riskLevel"`

	// To filter risk level greater than the one mentioned in query param
	RiskLevelGreaterThan *string `mandatory:"false" contributesTo:"query" name:"riskLevelGreaterThan"`

	// To filter risk level less than the one mentioned in query param
	RiskLevelLessThan *string `mandatory:"false" contributesTo:"query" name:"riskLevelLessThan"`

	// Comma seperated list of detector rule IDs to be passed in to match against Problems.
	DetectorRuleIdList []string `contributesTo:"query" name:"detectorRuleIdList" collectionFormat:"multi"`

	// The field to list the problems by detector type.
	DetectorType ListResourcesDetectorTypeEnum `mandatory:"false" contributesTo:"query" name:"detectorType" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListResourcesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use
	SortOrder ListResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourcesDetectorTypeEnum(string(request.DetectorType)); !ok && request.DetectorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DetectorType: %s. Supported values are: %s.", request.DetectorType, strings.Join(GetListResourcesDetectorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourcesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListResourcesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourcesResponse wrapper for the ListResources operation
type ListResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceCollection instances
	ResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourcesDetectorTypeEnum Enum with underlying type: string
type ListResourcesDetectorTypeEnum string

// Set of constants representing the allowable values for ListResourcesDetectorTypeEnum
const (
	ListResourcesDetectorTypeActivityDetector         ListResourcesDetectorTypeEnum = "IAAS_ACTIVITY_DETECTOR"
	ListResourcesDetectorTypeConfigurationDetector    ListResourcesDetectorTypeEnum = "IAAS_CONFIGURATION_DETECTOR"
	ListResourcesDetectorTypeThreatDetector           ListResourcesDetectorTypeEnum = "IAAS_THREAT_DETECTOR"
	ListResourcesDetectorTypeLogInsightDetector       ListResourcesDetectorTypeEnum = "IAAS_LOG_INSIGHT_DETECTOR"
	ListResourcesDetectorTypeInstanceSecurityDetector ListResourcesDetectorTypeEnum = "IAAS_INSTANCE_SECURITY_DETECTOR"
)

var mappingListResourcesDetectorTypeEnum = map[string]ListResourcesDetectorTypeEnum{
	"IAAS_ACTIVITY_DETECTOR":          ListResourcesDetectorTypeActivityDetector,
	"IAAS_CONFIGURATION_DETECTOR":     ListResourcesDetectorTypeConfigurationDetector,
	"IAAS_THREAT_DETECTOR":            ListResourcesDetectorTypeThreatDetector,
	"IAAS_LOG_INSIGHT_DETECTOR":       ListResourcesDetectorTypeLogInsightDetector,
	"IAAS_INSTANCE_SECURITY_DETECTOR": ListResourcesDetectorTypeInstanceSecurityDetector,
}

var mappingListResourcesDetectorTypeEnumLowerCase = map[string]ListResourcesDetectorTypeEnum{
	"iaas_activity_detector":          ListResourcesDetectorTypeActivityDetector,
	"iaas_configuration_detector":     ListResourcesDetectorTypeConfigurationDetector,
	"iaas_threat_detector":            ListResourcesDetectorTypeThreatDetector,
	"iaas_log_insight_detector":       ListResourcesDetectorTypeLogInsightDetector,
	"iaas_instance_security_detector": ListResourcesDetectorTypeInstanceSecurityDetector,
}

// GetListResourcesDetectorTypeEnumValues Enumerates the set of values for ListResourcesDetectorTypeEnum
func GetListResourcesDetectorTypeEnumValues() []ListResourcesDetectorTypeEnum {
	values := make([]ListResourcesDetectorTypeEnum, 0)
	for _, v := range mappingListResourcesDetectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourcesDetectorTypeEnumStringValues Enumerates the set of values in String for ListResourcesDetectorTypeEnum
func GetListResourcesDetectorTypeEnumStringValues() []string {
	return []string{
		"IAAS_ACTIVITY_DETECTOR",
		"IAAS_CONFIGURATION_DETECTOR",
		"IAAS_THREAT_DETECTOR",
		"IAAS_LOG_INSIGHT_DETECTOR",
		"IAAS_INSTANCE_SECURITY_DETECTOR",
	}
}

// GetMappingListResourcesDetectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourcesDetectorTypeEnum(val string) (ListResourcesDetectorTypeEnum, bool) {
	enum, ok := mappingListResourcesDetectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourcesAccessLevelEnum Enum with underlying type: string
type ListResourcesAccessLevelEnum string

// Set of constants representing the allowable values for ListResourcesAccessLevelEnum
const (
	ListResourcesAccessLevelRestricted ListResourcesAccessLevelEnum = "RESTRICTED"
	ListResourcesAccessLevelAccessible ListResourcesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListResourcesAccessLevelEnum = map[string]ListResourcesAccessLevelEnum{
	"RESTRICTED": ListResourcesAccessLevelRestricted,
	"ACCESSIBLE": ListResourcesAccessLevelAccessible,
}

var mappingListResourcesAccessLevelEnumLowerCase = map[string]ListResourcesAccessLevelEnum{
	"restricted": ListResourcesAccessLevelRestricted,
	"accessible": ListResourcesAccessLevelAccessible,
}

// GetListResourcesAccessLevelEnumValues Enumerates the set of values for ListResourcesAccessLevelEnum
func GetListResourcesAccessLevelEnumValues() []ListResourcesAccessLevelEnum {
	values := make([]ListResourcesAccessLevelEnum, 0)
	for _, v := range mappingListResourcesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourcesAccessLevelEnumStringValues Enumerates the set of values in String for ListResourcesAccessLevelEnum
func GetListResourcesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListResourcesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourcesAccessLevelEnum(val string) (ListResourcesAccessLevelEnum, bool) {
	enum, ok := mappingListResourcesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourcesSortOrderEnum Enum with underlying type: string
type ListResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListResourcesSortOrderEnum
const (
	ListResourcesSortOrderAsc  ListResourcesSortOrderEnum = "ASC"
	ListResourcesSortOrderDesc ListResourcesSortOrderEnum = "DESC"
)

var mappingListResourcesSortOrderEnum = map[string]ListResourcesSortOrderEnum{
	"ASC":  ListResourcesSortOrderAsc,
	"DESC": ListResourcesSortOrderDesc,
}

var mappingListResourcesSortOrderEnumLowerCase = map[string]ListResourcesSortOrderEnum{
	"asc":  ListResourcesSortOrderAsc,
	"desc": ListResourcesSortOrderDesc,
}

// GetListResourcesSortOrderEnumValues Enumerates the set of values for ListResourcesSortOrderEnum
func GetListResourcesSortOrderEnumValues() []ListResourcesSortOrderEnum {
	values := make([]ListResourcesSortOrderEnum, 0)
	for _, v := range mappingListResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListResourcesSortOrderEnum
func GetListResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourcesSortOrderEnum(val string) (ListResourcesSortOrderEnum, bool) {
	enum, ok := mappingListResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourcesSortByEnum Enum with underlying type: string
type ListResourcesSortByEnum string

// Set of constants representing the allowable values for ListResourcesSortByEnum
const (
	ListResourcesSortByTimecreated ListResourcesSortByEnum = "timeCreated"
	ListResourcesSortByDisplayname ListResourcesSortByEnum = "displayName"
)

var mappingListResourcesSortByEnum = map[string]ListResourcesSortByEnum{
	"timeCreated": ListResourcesSortByTimecreated,
	"displayName": ListResourcesSortByDisplayname,
}

var mappingListResourcesSortByEnumLowerCase = map[string]ListResourcesSortByEnum{
	"timecreated": ListResourcesSortByTimecreated,
	"displayname": ListResourcesSortByDisplayname,
}

// GetListResourcesSortByEnumValues Enumerates the set of values for ListResourcesSortByEnum
func GetListResourcesSortByEnumValues() []ListResourcesSortByEnum {
	values := make([]ListResourcesSortByEnum, 0)
	for _, v := range mappingListResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourcesSortByEnumStringValues Enumerates the set of values in String for ListResourcesSortByEnum
func GetListResourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourcesSortByEnum(val string) (ListResourcesSortByEnum, bool) {
	enum, ok := mappingListResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
