// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCryptoAssessmentTdeObjectsRequest wrapper for the ListCryptoAssessmentTdeObjects operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessmentTdeObjects.go.html to see an example of how to use ListCryptoAssessmentTdeObjectsRequest.
type ListCryptoAssessmentTdeObjectsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A required filter to return only TDE objects of the specified type.
	ObjectType ListCryptoAssessmentTdeObjectsObjectTypeEnum `mandatory:"true" contributesTo:"query" name:"objectType" omitEmpty:"true"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListCryptoAssessmentTdeObjectsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources associated with the specified crypto assessment OCID.
	AssessmentId *string `mandatory:"false" contributesTo:"query" name:"assessmentId"`

	// A filter to return targets from assessments of the specified type.
	AssessmentType CryptoAssessmentTypeEnum `mandatory:"false" contributesTo:"query" name:"assessmentType" omitEmpty:"true"`

	// A filter to return only inventory rows associated with the specified target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only resources associated with any of the specified target OCIDs.
	TargetIds []string `contributesTo:"query" name:"targetIds" collectionFormat:"multi"`

	// Filters TDE object summary rows by quantum-readiness category.
	QuantumReadiness ListCryptoAssessmentTdeObjectsQuantumReadinessEnum `mandatory:"false" contributesTo:"query" name:"quantumReadiness" omitEmpty:"true"`

	// Filters TDE object summary rows by any of the specified observed encryption algorithms.
	EncryptionObserved []string `contributesTo:"query" name:"encryptionObserved" collectionFormat:"multi"`

	// Filters TDE object summary rows by derived encryption status. NOT_SUPPORTED maps to rows where encryptionObserved is NOT_SUPPORTED, UNENCRYPTED maps to rows where encryptionObserved is null or NONE, and ENCRYPTED maps to rows where the observed encryption algorithm is any other value.
	EncryptionStatus ListCryptoAssessmentTdeObjectsEncryptionStatusEnum `mandatory:"false" contributesTo:"query" name:"encryptionStatus" omitEmpty:"true"`

	// The field used to sort TDE object summary results.
	SortBy ListCryptoAssessmentTdeObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCryptoAssessmentTdeObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCryptoAssessmentTdeObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentTdeObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentTdeObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentTdeObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentTdeObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAssessmentTdeObjectsObjectTypeEnum(string(request.ObjectType)); !ok && request.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", request.ObjectType, strings.Join(GetListCryptoAssessmentTdeObjectsObjectTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentTdeObjectsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCryptoAssessmentTdeObjectsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(request.AssessmentType)); !ok && request.AssessmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentType: %s. Supported values are: %s.", request.AssessmentType, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentTdeObjectsQuantumReadinessEnum(string(request.QuantumReadiness)); !ok && request.QuantumReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuantumReadiness: %s. Supported values are: %s.", request.QuantumReadiness, strings.Join(GetListCryptoAssessmentTdeObjectsQuantumReadinessEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentTdeObjectsEncryptionStatusEnum(string(request.EncryptionStatus)); !ok && request.EncryptionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EncryptionStatus: %s. Supported values are: %s.", request.EncryptionStatus, strings.Join(GetListCryptoAssessmentTdeObjectsEncryptionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentTdeObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAssessmentTdeObjectsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentTdeObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAssessmentTdeObjectsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentTdeObjectsResponse wrapper for the ListCryptoAssessmentTdeObjects operation
type ListCryptoAssessmentTdeObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentTdeObjectCollection instances
	CryptoAssessmentTdeObjectCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCryptoAssessmentTdeObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentTdeObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentTdeObjectsObjectTypeEnum Enum with underlying type: string
type ListCryptoAssessmentTdeObjectsObjectTypeEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentTdeObjectsObjectTypeEnum
const (
	ListCryptoAssessmentTdeObjectsObjectTypeTablespace ListCryptoAssessmentTdeObjectsObjectTypeEnum = "TABLESPACE"
	ListCryptoAssessmentTdeObjectsObjectTypeColumn     ListCryptoAssessmentTdeObjectsObjectTypeEnum = "COLUMN"
)

var mappingListCryptoAssessmentTdeObjectsObjectTypeEnum = map[string]ListCryptoAssessmentTdeObjectsObjectTypeEnum{
	"TABLESPACE": ListCryptoAssessmentTdeObjectsObjectTypeTablespace,
	"COLUMN":     ListCryptoAssessmentTdeObjectsObjectTypeColumn,
}

var mappingListCryptoAssessmentTdeObjectsObjectTypeEnumLowerCase = map[string]ListCryptoAssessmentTdeObjectsObjectTypeEnum{
	"tablespace": ListCryptoAssessmentTdeObjectsObjectTypeTablespace,
	"column":     ListCryptoAssessmentTdeObjectsObjectTypeColumn,
}

// GetListCryptoAssessmentTdeObjectsObjectTypeEnumValues Enumerates the set of values for ListCryptoAssessmentTdeObjectsObjectTypeEnum
func GetListCryptoAssessmentTdeObjectsObjectTypeEnumValues() []ListCryptoAssessmentTdeObjectsObjectTypeEnum {
	values := make([]ListCryptoAssessmentTdeObjectsObjectTypeEnum, 0)
	for _, v := range mappingListCryptoAssessmentTdeObjectsObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentTdeObjectsObjectTypeEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentTdeObjectsObjectTypeEnum
func GetListCryptoAssessmentTdeObjectsObjectTypeEnumStringValues() []string {
	return []string{
		"TABLESPACE",
		"COLUMN",
	}
}

// GetMappingListCryptoAssessmentTdeObjectsObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentTdeObjectsObjectTypeEnum(val string) (ListCryptoAssessmentTdeObjectsObjectTypeEnum, bool) {
	enum, ok := mappingListCryptoAssessmentTdeObjectsObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentTdeObjectsAccessLevelEnum Enum with underlying type: string
type ListCryptoAssessmentTdeObjectsAccessLevelEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentTdeObjectsAccessLevelEnum
const (
	ListCryptoAssessmentTdeObjectsAccessLevelRestricted ListCryptoAssessmentTdeObjectsAccessLevelEnum = "RESTRICTED"
	ListCryptoAssessmentTdeObjectsAccessLevelAccessible ListCryptoAssessmentTdeObjectsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCryptoAssessmentTdeObjectsAccessLevelEnum = map[string]ListCryptoAssessmentTdeObjectsAccessLevelEnum{
	"RESTRICTED": ListCryptoAssessmentTdeObjectsAccessLevelRestricted,
	"ACCESSIBLE": ListCryptoAssessmentTdeObjectsAccessLevelAccessible,
}

var mappingListCryptoAssessmentTdeObjectsAccessLevelEnumLowerCase = map[string]ListCryptoAssessmentTdeObjectsAccessLevelEnum{
	"restricted": ListCryptoAssessmentTdeObjectsAccessLevelRestricted,
	"accessible": ListCryptoAssessmentTdeObjectsAccessLevelAccessible,
}

// GetListCryptoAssessmentTdeObjectsAccessLevelEnumValues Enumerates the set of values for ListCryptoAssessmentTdeObjectsAccessLevelEnum
func GetListCryptoAssessmentTdeObjectsAccessLevelEnumValues() []ListCryptoAssessmentTdeObjectsAccessLevelEnum {
	values := make([]ListCryptoAssessmentTdeObjectsAccessLevelEnum, 0)
	for _, v := range mappingListCryptoAssessmentTdeObjectsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentTdeObjectsAccessLevelEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentTdeObjectsAccessLevelEnum
func GetListCryptoAssessmentTdeObjectsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCryptoAssessmentTdeObjectsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentTdeObjectsAccessLevelEnum(val string) (ListCryptoAssessmentTdeObjectsAccessLevelEnum, bool) {
	enum, ok := mappingListCryptoAssessmentTdeObjectsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentTdeObjectsQuantumReadinessEnum Enum with underlying type: string
type ListCryptoAssessmentTdeObjectsQuantumReadinessEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentTdeObjectsQuantumReadinessEnum
const (
	ListCryptoAssessmentTdeObjectsQuantumReadinessResistant     ListCryptoAssessmentTdeObjectsQuantumReadinessEnum = "RESISTANT"
	ListCryptoAssessmentTdeObjectsQuantumReadinessNotResistant  ListCryptoAssessmentTdeObjectsQuantumReadinessEnum = "NOT_RESISTANT"
	ListCryptoAssessmentTdeObjectsQuantumReadinessNotAvailable  ListCryptoAssessmentTdeObjectsQuantumReadinessEnum = "NOT_AVAILABLE"
	ListCryptoAssessmentTdeObjectsQuantumReadinessNotApplicable ListCryptoAssessmentTdeObjectsQuantumReadinessEnum = "NOT_APPLICABLE"
	ListCryptoAssessmentTdeObjectsQuantumReadinessNotSupported  ListCryptoAssessmentTdeObjectsQuantumReadinessEnum = "NOT_SUPPORTED"
)

var mappingListCryptoAssessmentTdeObjectsQuantumReadinessEnum = map[string]ListCryptoAssessmentTdeObjectsQuantumReadinessEnum{
	"RESISTANT":      ListCryptoAssessmentTdeObjectsQuantumReadinessResistant,
	"NOT_RESISTANT":  ListCryptoAssessmentTdeObjectsQuantumReadinessNotResistant,
	"NOT_AVAILABLE":  ListCryptoAssessmentTdeObjectsQuantumReadinessNotAvailable,
	"NOT_APPLICABLE": ListCryptoAssessmentTdeObjectsQuantumReadinessNotApplicable,
	"NOT_SUPPORTED":  ListCryptoAssessmentTdeObjectsQuantumReadinessNotSupported,
}

var mappingListCryptoAssessmentTdeObjectsQuantumReadinessEnumLowerCase = map[string]ListCryptoAssessmentTdeObjectsQuantumReadinessEnum{
	"resistant":      ListCryptoAssessmentTdeObjectsQuantumReadinessResistant,
	"not_resistant":  ListCryptoAssessmentTdeObjectsQuantumReadinessNotResistant,
	"not_available":  ListCryptoAssessmentTdeObjectsQuantumReadinessNotAvailable,
	"not_applicable": ListCryptoAssessmentTdeObjectsQuantumReadinessNotApplicable,
	"not_supported":  ListCryptoAssessmentTdeObjectsQuantumReadinessNotSupported,
}

// GetListCryptoAssessmentTdeObjectsQuantumReadinessEnumValues Enumerates the set of values for ListCryptoAssessmentTdeObjectsQuantumReadinessEnum
func GetListCryptoAssessmentTdeObjectsQuantumReadinessEnumValues() []ListCryptoAssessmentTdeObjectsQuantumReadinessEnum {
	values := make([]ListCryptoAssessmentTdeObjectsQuantumReadinessEnum, 0)
	for _, v := range mappingListCryptoAssessmentTdeObjectsQuantumReadinessEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentTdeObjectsQuantumReadinessEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentTdeObjectsQuantumReadinessEnum
func GetListCryptoAssessmentTdeObjectsQuantumReadinessEnumStringValues() []string {
	return []string{
		"RESISTANT",
		"NOT_RESISTANT",
		"NOT_AVAILABLE",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingListCryptoAssessmentTdeObjectsQuantumReadinessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentTdeObjectsQuantumReadinessEnum(val string) (ListCryptoAssessmentTdeObjectsQuantumReadinessEnum, bool) {
	enum, ok := mappingListCryptoAssessmentTdeObjectsQuantumReadinessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentTdeObjectsEncryptionStatusEnum Enum with underlying type: string
type ListCryptoAssessmentTdeObjectsEncryptionStatusEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentTdeObjectsEncryptionStatusEnum
const (
	ListCryptoAssessmentTdeObjectsEncryptionStatusEncrypted    ListCryptoAssessmentTdeObjectsEncryptionStatusEnum = "ENCRYPTED"
	ListCryptoAssessmentTdeObjectsEncryptionStatusUnencrypted  ListCryptoAssessmentTdeObjectsEncryptionStatusEnum = "UNENCRYPTED"
	ListCryptoAssessmentTdeObjectsEncryptionStatusNotSupported ListCryptoAssessmentTdeObjectsEncryptionStatusEnum = "NOT_SUPPORTED"
)

var mappingListCryptoAssessmentTdeObjectsEncryptionStatusEnum = map[string]ListCryptoAssessmentTdeObjectsEncryptionStatusEnum{
	"ENCRYPTED":     ListCryptoAssessmentTdeObjectsEncryptionStatusEncrypted,
	"UNENCRYPTED":   ListCryptoAssessmentTdeObjectsEncryptionStatusUnencrypted,
	"NOT_SUPPORTED": ListCryptoAssessmentTdeObjectsEncryptionStatusNotSupported,
}

var mappingListCryptoAssessmentTdeObjectsEncryptionStatusEnumLowerCase = map[string]ListCryptoAssessmentTdeObjectsEncryptionStatusEnum{
	"encrypted":     ListCryptoAssessmentTdeObjectsEncryptionStatusEncrypted,
	"unencrypted":   ListCryptoAssessmentTdeObjectsEncryptionStatusUnencrypted,
	"not_supported": ListCryptoAssessmentTdeObjectsEncryptionStatusNotSupported,
}

// GetListCryptoAssessmentTdeObjectsEncryptionStatusEnumValues Enumerates the set of values for ListCryptoAssessmentTdeObjectsEncryptionStatusEnum
func GetListCryptoAssessmentTdeObjectsEncryptionStatusEnumValues() []ListCryptoAssessmentTdeObjectsEncryptionStatusEnum {
	values := make([]ListCryptoAssessmentTdeObjectsEncryptionStatusEnum, 0)
	for _, v := range mappingListCryptoAssessmentTdeObjectsEncryptionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentTdeObjectsEncryptionStatusEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentTdeObjectsEncryptionStatusEnum
func GetListCryptoAssessmentTdeObjectsEncryptionStatusEnumStringValues() []string {
	return []string{
		"ENCRYPTED",
		"UNENCRYPTED",
		"NOT_SUPPORTED",
	}
}

// GetMappingListCryptoAssessmentTdeObjectsEncryptionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentTdeObjectsEncryptionStatusEnum(val string) (ListCryptoAssessmentTdeObjectsEncryptionStatusEnum, bool) {
	enum, ok := mappingListCryptoAssessmentTdeObjectsEncryptionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentTdeObjectsSortByEnum Enum with underlying type: string
type ListCryptoAssessmentTdeObjectsSortByEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentTdeObjectsSortByEnum
const (
	ListCryptoAssessmentTdeObjectsSortByTablespacename   ListCryptoAssessmentTdeObjectsSortByEnum = "tablespaceName"
	ListCryptoAssessmentTdeObjectsSortByColumnname       ListCryptoAssessmentTdeObjectsSortByEnum = "columnName"
	ListCryptoAssessmentTdeObjectsSortByQuantumreadiness ListCryptoAssessmentTdeObjectsSortByEnum = "quantumReadiness"
)

var mappingListCryptoAssessmentTdeObjectsSortByEnum = map[string]ListCryptoAssessmentTdeObjectsSortByEnum{
	"tablespaceName":   ListCryptoAssessmentTdeObjectsSortByTablespacename,
	"columnName":       ListCryptoAssessmentTdeObjectsSortByColumnname,
	"quantumReadiness": ListCryptoAssessmentTdeObjectsSortByQuantumreadiness,
}

var mappingListCryptoAssessmentTdeObjectsSortByEnumLowerCase = map[string]ListCryptoAssessmentTdeObjectsSortByEnum{
	"tablespacename":   ListCryptoAssessmentTdeObjectsSortByTablespacename,
	"columnname":       ListCryptoAssessmentTdeObjectsSortByColumnname,
	"quantumreadiness": ListCryptoAssessmentTdeObjectsSortByQuantumreadiness,
}

// GetListCryptoAssessmentTdeObjectsSortByEnumValues Enumerates the set of values for ListCryptoAssessmentTdeObjectsSortByEnum
func GetListCryptoAssessmentTdeObjectsSortByEnumValues() []ListCryptoAssessmentTdeObjectsSortByEnum {
	values := make([]ListCryptoAssessmentTdeObjectsSortByEnum, 0)
	for _, v := range mappingListCryptoAssessmentTdeObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentTdeObjectsSortByEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentTdeObjectsSortByEnum
func GetListCryptoAssessmentTdeObjectsSortByEnumStringValues() []string {
	return []string{
		"tablespaceName",
		"columnName",
		"quantumReadiness",
	}
}

// GetMappingListCryptoAssessmentTdeObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentTdeObjectsSortByEnum(val string) (ListCryptoAssessmentTdeObjectsSortByEnum, bool) {
	enum, ok := mappingListCryptoAssessmentTdeObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentTdeObjectsSortOrderEnum Enum with underlying type: string
type ListCryptoAssessmentTdeObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentTdeObjectsSortOrderEnum
const (
	ListCryptoAssessmentTdeObjectsSortOrderAsc  ListCryptoAssessmentTdeObjectsSortOrderEnum = "ASC"
	ListCryptoAssessmentTdeObjectsSortOrderDesc ListCryptoAssessmentTdeObjectsSortOrderEnum = "DESC"
)

var mappingListCryptoAssessmentTdeObjectsSortOrderEnum = map[string]ListCryptoAssessmentTdeObjectsSortOrderEnum{
	"ASC":  ListCryptoAssessmentTdeObjectsSortOrderAsc,
	"DESC": ListCryptoAssessmentTdeObjectsSortOrderDesc,
}

var mappingListCryptoAssessmentTdeObjectsSortOrderEnumLowerCase = map[string]ListCryptoAssessmentTdeObjectsSortOrderEnum{
	"asc":  ListCryptoAssessmentTdeObjectsSortOrderAsc,
	"desc": ListCryptoAssessmentTdeObjectsSortOrderDesc,
}

// GetListCryptoAssessmentTdeObjectsSortOrderEnumValues Enumerates the set of values for ListCryptoAssessmentTdeObjectsSortOrderEnum
func GetListCryptoAssessmentTdeObjectsSortOrderEnumValues() []ListCryptoAssessmentTdeObjectsSortOrderEnum {
	values := make([]ListCryptoAssessmentTdeObjectsSortOrderEnum, 0)
	for _, v := range mappingListCryptoAssessmentTdeObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentTdeObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentTdeObjectsSortOrderEnum
func GetListCryptoAssessmentTdeObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAssessmentTdeObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentTdeObjectsSortOrderEnum(val string) (ListCryptoAssessmentTdeObjectsSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAssessmentTdeObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
