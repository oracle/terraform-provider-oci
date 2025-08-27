// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSensitiveColumnAnalyticsRequest wrapper for the ListSensitiveColumnAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveColumnAnalytics.go.html to see an example of how to use ListSensitiveColumnAnalyticsRequest.
type ListSensitiveColumnAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSensitiveColumnAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return the target database group that matches the specified OCID.
	TargetDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupId"`

	// A filter to return only the sensitive columns that are associated with one of the sensitive types identified by the specified OCIDs.
	SensitiveTypeId []string `contributesTo:"query" name:"sensitiveTypeId" collectionFormat:"multi"`

	// An optional filter to return only resources that match the specified OCID of the sensitive type group resource.
	SensitiveTypeGroupId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypeGroupId"`

	// A filter to return only the resources that match the specified sensitive data model OCID.
	SensitiveDataModelId *string `mandatory:"false" contributesTo:"query" name:"sensitiveDataModelId"`

	// The group by parameter to summarize the sensitive columns.
	GroupBy []ListSensitiveColumnAnalyticsGroupByEnum `contributesTo:"query" name:"groupBy" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to a specific object name.
	ObjectName []string `contributesTo:"query" name:"objectName" collectionFormat:"multi"`

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveColumnAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveColumnAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveColumnAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveColumnAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveColumnAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSensitiveColumnAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSensitiveColumnAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	for _, val := range request.GroupBy {
		if _, ok := GetMappingListSensitiveColumnAnalyticsGroupByEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", val, strings.Join(GetListSensitiveColumnAnalyticsGroupByEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveColumnAnalyticsResponse wrapper for the ListSensitiveColumnAnalytics operation
type ListSensitiveColumnAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveColumnAnalyticsCollection instances
	SensitiveColumnAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveColumnAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveColumnAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveColumnAnalyticsAccessLevelEnum Enum with underlying type: string
type ListSensitiveColumnAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListSensitiveColumnAnalyticsAccessLevelEnum
const (
	ListSensitiveColumnAnalyticsAccessLevelRestricted ListSensitiveColumnAnalyticsAccessLevelEnum = "RESTRICTED"
	ListSensitiveColumnAnalyticsAccessLevelAccessible ListSensitiveColumnAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSensitiveColumnAnalyticsAccessLevelEnum = map[string]ListSensitiveColumnAnalyticsAccessLevelEnum{
	"RESTRICTED": ListSensitiveColumnAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListSensitiveColumnAnalyticsAccessLevelAccessible,
}

var mappingListSensitiveColumnAnalyticsAccessLevelEnumLowerCase = map[string]ListSensitiveColumnAnalyticsAccessLevelEnum{
	"restricted": ListSensitiveColumnAnalyticsAccessLevelRestricted,
	"accessible": ListSensitiveColumnAnalyticsAccessLevelAccessible,
}

// GetListSensitiveColumnAnalyticsAccessLevelEnumValues Enumerates the set of values for ListSensitiveColumnAnalyticsAccessLevelEnum
func GetListSensitiveColumnAnalyticsAccessLevelEnumValues() []ListSensitiveColumnAnalyticsAccessLevelEnum {
	values := make([]ListSensitiveColumnAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListSensitiveColumnAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveColumnAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListSensitiveColumnAnalyticsAccessLevelEnum
func GetListSensitiveColumnAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSensitiveColumnAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveColumnAnalyticsAccessLevelEnum(val string) (ListSensitiveColumnAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListSensitiveColumnAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveColumnAnalyticsGroupByEnum Enum with underlying type: string
type ListSensitiveColumnAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ListSensitiveColumnAnalyticsGroupByEnum
const (
	ListSensitiveColumnAnalyticsGroupByTargetid             ListSensitiveColumnAnalyticsGroupByEnum = "targetId"
	ListSensitiveColumnAnalyticsGroupBySensitivetypeid      ListSensitiveColumnAnalyticsGroupByEnum = "sensitiveTypeId"
	ListSensitiveColumnAnalyticsGroupBySensitivedatamodelid ListSensitiveColumnAnalyticsGroupByEnum = "sensitiveDataModelId"
)

var mappingListSensitiveColumnAnalyticsGroupByEnum = map[string]ListSensitiveColumnAnalyticsGroupByEnum{
	"targetId":             ListSensitiveColumnAnalyticsGroupByTargetid,
	"sensitiveTypeId":      ListSensitiveColumnAnalyticsGroupBySensitivetypeid,
	"sensitiveDataModelId": ListSensitiveColumnAnalyticsGroupBySensitivedatamodelid,
}

var mappingListSensitiveColumnAnalyticsGroupByEnumLowerCase = map[string]ListSensitiveColumnAnalyticsGroupByEnum{
	"targetid":             ListSensitiveColumnAnalyticsGroupByTargetid,
	"sensitivetypeid":      ListSensitiveColumnAnalyticsGroupBySensitivetypeid,
	"sensitivedatamodelid": ListSensitiveColumnAnalyticsGroupBySensitivedatamodelid,
}

// GetListSensitiveColumnAnalyticsGroupByEnumValues Enumerates the set of values for ListSensitiveColumnAnalyticsGroupByEnum
func GetListSensitiveColumnAnalyticsGroupByEnumValues() []ListSensitiveColumnAnalyticsGroupByEnum {
	values := make([]ListSensitiveColumnAnalyticsGroupByEnum, 0)
	for _, v := range mappingListSensitiveColumnAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveColumnAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ListSensitiveColumnAnalyticsGroupByEnum
func GetListSensitiveColumnAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"targetId",
		"sensitiveTypeId",
		"sensitiveDataModelId",
	}
}

// GetMappingListSensitiveColumnAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveColumnAnalyticsGroupByEnum(val string) (ListSensitiveColumnAnalyticsGroupByEnum, bool) {
	enum, ok := mappingListSensitiveColumnAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
