// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDatabaseInsightsRequest wrapper for the ListDatabaseInsights operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseInsights.go.html to see an example of how to use ListDatabaseInsightsRequest.
type ListDatabaseInsightsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique Enterprise Manager bridge identifier
	EnterpriseManagerBridgeId *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerBridgeId"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Lifecycle states
	LifecycleState []LifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []ListDatabaseInsightsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Specifies the fields to return in a database summary response. By default all fields are returned if omitted.
	Fields []ListDatabaseInsightsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDatabaseInsightsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Database insight list sort options. If `fields` parameter is selected, the `sortBy` parameter must be one of the fields specified.
	SortBy ListDatabaseInsightsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
	ExadataInsightId *string `mandatory:"false" contributesTo:"query" name:"exadataInsightId"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Unique Operations Insights PrivateEndpoint identifier
	OpsiPrivateEndpointId *string `mandatory:"false" contributesTo:"query" name:"opsiPrivateEndpointId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseInsightsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseInsightsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseInsightsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseInsightsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseInsightsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Status {
		if _, ok := GetMappingResourceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetResourceStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.LifecycleState {
		if _, ok := GetMappingLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
		}
	}

	for _, val := range request.DatabaseType {
		if _, ok := GetMappingListDatabaseInsightsDatabaseTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", val, strings.Join(GetListDatabaseInsightsDatabaseTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.Fields {
		if _, ok := GetMappingListDatabaseInsightsFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListDatabaseInsightsFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListDatabaseInsightsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseInsightsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseInsightsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseInsightsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseInsightsResponse wrapper for the ListDatabaseInsights operation
type ListDatabaseInsightsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseInsightsCollection instances
	DatabaseInsightsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseInsightsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseInsightsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseInsightsDatabaseTypeEnum Enum with underlying type: string
type ListDatabaseInsightsDatabaseTypeEnum string

// Set of constants representing the allowable values for ListDatabaseInsightsDatabaseTypeEnum
const (
	ListDatabaseInsightsDatabaseTypeAdwS                 ListDatabaseInsightsDatabaseTypeEnum = "ADW-S"
	ListDatabaseInsightsDatabaseTypeAtpS                 ListDatabaseInsightsDatabaseTypeEnum = "ATP-S"
	ListDatabaseInsightsDatabaseTypeAdwD                 ListDatabaseInsightsDatabaseTypeEnum = "ADW-D"
	ListDatabaseInsightsDatabaseTypeAtpD                 ListDatabaseInsightsDatabaseTypeEnum = "ATP-D"
	ListDatabaseInsightsDatabaseTypeExternalPdb          ListDatabaseInsightsDatabaseTypeEnum = "EXTERNAL-PDB"
	ListDatabaseInsightsDatabaseTypeExternalNoncdb       ListDatabaseInsightsDatabaseTypeEnum = "EXTERNAL-NONCDB"
	ListDatabaseInsightsDatabaseTypeComanagedVmCdb       ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-VM-CDB"
	ListDatabaseInsightsDatabaseTypeComanagedVmPdb       ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-VM-PDB"
	ListDatabaseInsightsDatabaseTypeComanagedVmNoncdb    ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-VM-NONCDB"
	ListDatabaseInsightsDatabaseTypeComanagedBmCdb       ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-BM-CDB"
	ListDatabaseInsightsDatabaseTypeComanagedBmPdb       ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-BM-PDB"
	ListDatabaseInsightsDatabaseTypeComanagedBmNoncdb    ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-BM-NONCDB"
	ListDatabaseInsightsDatabaseTypeComanagedExacsCdb    ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-EXACS-CDB"
	ListDatabaseInsightsDatabaseTypeComanagedExacsPdb    ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-EXACS-PDB"
	ListDatabaseInsightsDatabaseTypeComanagedExacsNoncdb ListDatabaseInsightsDatabaseTypeEnum = "COMANAGED-EXACS-NONCDB"
)

var mappingListDatabaseInsightsDatabaseTypeEnum = map[string]ListDatabaseInsightsDatabaseTypeEnum{
	"ADW-S":                  ListDatabaseInsightsDatabaseTypeAdwS,
	"ATP-S":                  ListDatabaseInsightsDatabaseTypeAtpS,
	"ADW-D":                  ListDatabaseInsightsDatabaseTypeAdwD,
	"ATP-D":                  ListDatabaseInsightsDatabaseTypeAtpD,
	"EXTERNAL-PDB":           ListDatabaseInsightsDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB":        ListDatabaseInsightsDatabaseTypeExternalNoncdb,
	"COMANAGED-VM-CDB":       ListDatabaseInsightsDatabaseTypeComanagedVmCdb,
	"COMANAGED-VM-PDB":       ListDatabaseInsightsDatabaseTypeComanagedVmPdb,
	"COMANAGED-VM-NONCDB":    ListDatabaseInsightsDatabaseTypeComanagedVmNoncdb,
	"COMANAGED-BM-CDB":       ListDatabaseInsightsDatabaseTypeComanagedBmCdb,
	"COMANAGED-BM-PDB":       ListDatabaseInsightsDatabaseTypeComanagedBmPdb,
	"COMANAGED-BM-NONCDB":    ListDatabaseInsightsDatabaseTypeComanagedBmNoncdb,
	"COMANAGED-EXACS-CDB":    ListDatabaseInsightsDatabaseTypeComanagedExacsCdb,
	"COMANAGED-EXACS-PDB":    ListDatabaseInsightsDatabaseTypeComanagedExacsPdb,
	"COMANAGED-EXACS-NONCDB": ListDatabaseInsightsDatabaseTypeComanagedExacsNoncdb,
}

var mappingListDatabaseInsightsDatabaseTypeEnumLowerCase = map[string]ListDatabaseInsightsDatabaseTypeEnum{
	"adw-s":                  ListDatabaseInsightsDatabaseTypeAdwS,
	"atp-s":                  ListDatabaseInsightsDatabaseTypeAtpS,
	"adw-d":                  ListDatabaseInsightsDatabaseTypeAdwD,
	"atp-d":                  ListDatabaseInsightsDatabaseTypeAtpD,
	"external-pdb":           ListDatabaseInsightsDatabaseTypeExternalPdb,
	"external-noncdb":        ListDatabaseInsightsDatabaseTypeExternalNoncdb,
	"comanaged-vm-cdb":       ListDatabaseInsightsDatabaseTypeComanagedVmCdb,
	"comanaged-vm-pdb":       ListDatabaseInsightsDatabaseTypeComanagedVmPdb,
	"comanaged-vm-noncdb":    ListDatabaseInsightsDatabaseTypeComanagedVmNoncdb,
	"comanaged-bm-cdb":       ListDatabaseInsightsDatabaseTypeComanagedBmCdb,
	"comanaged-bm-pdb":       ListDatabaseInsightsDatabaseTypeComanagedBmPdb,
	"comanaged-bm-noncdb":    ListDatabaseInsightsDatabaseTypeComanagedBmNoncdb,
	"comanaged-exacs-cdb":    ListDatabaseInsightsDatabaseTypeComanagedExacsCdb,
	"comanaged-exacs-pdb":    ListDatabaseInsightsDatabaseTypeComanagedExacsPdb,
	"comanaged-exacs-noncdb": ListDatabaseInsightsDatabaseTypeComanagedExacsNoncdb,
}

// GetListDatabaseInsightsDatabaseTypeEnumValues Enumerates the set of values for ListDatabaseInsightsDatabaseTypeEnum
func GetListDatabaseInsightsDatabaseTypeEnumValues() []ListDatabaseInsightsDatabaseTypeEnum {
	values := make([]ListDatabaseInsightsDatabaseTypeEnum, 0)
	for _, v := range mappingListDatabaseInsightsDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseInsightsDatabaseTypeEnumStringValues Enumerates the set of values in String for ListDatabaseInsightsDatabaseTypeEnum
func GetListDatabaseInsightsDatabaseTypeEnumStringValues() []string {
	return []string{
		"ADW-S",
		"ATP-S",
		"ADW-D",
		"ATP-D",
		"EXTERNAL-PDB",
		"EXTERNAL-NONCDB",
		"COMANAGED-VM-CDB",
		"COMANAGED-VM-PDB",
		"COMANAGED-VM-NONCDB",
		"COMANAGED-BM-CDB",
		"COMANAGED-BM-PDB",
		"COMANAGED-BM-NONCDB",
		"COMANAGED-EXACS-CDB",
		"COMANAGED-EXACS-PDB",
		"COMANAGED-EXACS-NONCDB",
	}
}

// GetMappingListDatabaseInsightsDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseInsightsDatabaseTypeEnum(val string) (ListDatabaseInsightsDatabaseTypeEnum, bool) {
	enum, ok := mappingListDatabaseInsightsDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseInsightsFieldsEnum Enum with underlying type: string
type ListDatabaseInsightsFieldsEnum string

// Set of constants representing the allowable values for ListDatabaseInsightsFieldsEnum
const (
	ListDatabaseInsightsFieldsCompartmentid       ListDatabaseInsightsFieldsEnum = "compartmentId"
	ListDatabaseInsightsFieldsDatabasename        ListDatabaseInsightsFieldsEnum = "databaseName"
	ListDatabaseInsightsFieldsDatabasedisplayname ListDatabaseInsightsFieldsEnum = "databaseDisplayName"
	ListDatabaseInsightsFieldsDatabasetype        ListDatabaseInsightsFieldsEnum = "databaseType"
	ListDatabaseInsightsFieldsDatabaseversion     ListDatabaseInsightsFieldsEnum = "databaseVersion"
	ListDatabaseInsightsFieldsDatabasehostnames   ListDatabaseInsightsFieldsEnum = "databaseHostNames"
	ListDatabaseInsightsFieldsFreeformtags        ListDatabaseInsightsFieldsEnum = "freeformTags"
	ListDatabaseInsightsFieldsDefinedtags         ListDatabaseInsightsFieldsEnum = "definedTags"
)

var mappingListDatabaseInsightsFieldsEnum = map[string]ListDatabaseInsightsFieldsEnum{
	"compartmentId":       ListDatabaseInsightsFieldsCompartmentid,
	"databaseName":        ListDatabaseInsightsFieldsDatabasename,
	"databaseDisplayName": ListDatabaseInsightsFieldsDatabasedisplayname,
	"databaseType":        ListDatabaseInsightsFieldsDatabasetype,
	"databaseVersion":     ListDatabaseInsightsFieldsDatabaseversion,
	"databaseHostNames":   ListDatabaseInsightsFieldsDatabasehostnames,
	"freeformTags":        ListDatabaseInsightsFieldsFreeformtags,
	"definedTags":         ListDatabaseInsightsFieldsDefinedtags,
}

var mappingListDatabaseInsightsFieldsEnumLowerCase = map[string]ListDatabaseInsightsFieldsEnum{
	"compartmentid":       ListDatabaseInsightsFieldsCompartmentid,
	"databasename":        ListDatabaseInsightsFieldsDatabasename,
	"databasedisplayname": ListDatabaseInsightsFieldsDatabasedisplayname,
	"databasetype":        ListDatabaseInsightsFieldsDatabasetype,
	"databaseversion":     ListDatabaseInsightsFieldsDatabaseversion,
	"databasehostnames":   ListDatabaseInsightsFieldsDatabasehostnames,
	"freeformtags":        ListDatabaseInsightsFieldsFreeformtags,
	"definedtags":         ListDatabaseInsightsFieldsDefinedtags,
}

// GetListDatabaseInsightsFieldsEnumValues Enumerates the set of values for ListDatabaseInsightsFieldsEnum
func GetListDatabaseInsightsFieldsEnumValues() []ListDatabaseInsightsFieldsEnum {
	values := make([]ListDatabaseInsightsFieldsEnum, 0)
	for _, v := range mappingListDatabaseInsightsFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseInsightsFieldsEnumStringValues Enumerates the set of values in String for ListDatabaseInsightsFieldsEnum
func GetListDatabaseInsightsFieldsEnumStringValues() []string {
	return []string{
		"compartmentId",
		"databaseName",
		"databaseDisplayName",
		"databaseType",
		"databaseVersion",
		"databaseHostNames",
		"freeformTags",
		"definedTags",
	}
}

// GetMappingListDatabaseInsightsFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseInsightsFieldsEnum(val string) (ListDatabaseInsightsFieldsEnum, bool) {
	enum, ok := mappingListDatabaseInsightsFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseInsightsSortOrderEnum Enum with underlying type: string
type ListDatabaseInsightsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseInsightsSortOrderEnum
const (
	ListDatabaseInsightsSortOrderAsc  ListDatabaseInsightsSortOrderEnum = "ASC"
	ListDatabaseInsightsSortOrderDesc ListDatabaseInsightsSortOrderEnum = "DESC"
)

var mappingListDatabaseInsightsSortOrderEnum = map[string]ListDatabaseInsightsSortOrderEnum{
	"ASC":  ListDatabaseInsightsSortOrderAsc,
	"DESC": ListDatabaseInsightsSortOrderDesc,
}

var mappingListDatabaseInsightsSortOrderEnumLowerCase = map[string]ListDatabaseInsightsSortOrderEnum{
	"asc":  ListDatabaseInsightsSortOrderAsc,
	"desc": ListDatabaseInsightsSortOrderDesc,
}

// GetListDatabaseInsightsSortOrderEnumValues Enumerates the set of values for ListDatabaseInsightsSortOrderEnum
func GetListDatabaseInsightsSortOrderEnumValues() []ListDatabaseInsightsSortOrderEnum {
	values := make([]ListDatabaseInsightsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseInsightsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseInsightsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseInsightsSortOrderEnum
func GetListDatabaseInsightsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseInsightsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseInsightsSortOrderEnum(val string) (ListDatabaseInsightsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseInsightsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseInsightsSortByEnum Enum with underlying type: string
type ListDatabaseInsightsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseInsightsSortByEnum
const (
	ListDatabaseInsightsSortByDatabasename        ListDatabaseInsightsSortByEnum = "databaseName"
	ListDatabaseInsightsSortByDatabasedisplayname ListDatabaseInsightsSortByEnum = "databaseDisplayName"
	ListDatabaseInsightsSortByDatabasetype        ListDatabaseInsightsSortByEnum = "databaseType"
)

var mappingListDatabaseInsightsSortByEnum = map[string]ListDatabaseInsightsSortByEnum{
	"databaseName":        ListDatabaseInsightsSortByDatabasename,
	"databaseDisplayName": ListDatabaseInsightsSortByDatabasedisplayname,
	"databaseType":        ListDatabaseInsightsSortByDatabasetype,
}

var mappingListDatabaseInsightsSortByEnumLowerCase = map[string]ListDatabaseInsightsSortByEnum{
	"databasename":        ListDatabaseInsightsSortByDatabasename,
	"databasedisplayname": ListDatabaseInsightsSortByDatabasedisplayname,
	"databasetype":        ListDatabaseInsightsSortByDatabasetype,
}

// GetListDatabaseInsightsSortByEnumValues Enumerates the set of values for ListDatabaseInsightsSortByEnum
func GetListDatabaseInsightsSortByEnumValues() []ListDatabaseInsightsSortByEnum {
	values := make([]ListDatabaseInsightsSortByEnum, 0)
	for _, v := range mappingListDatabaseInsightsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseInsightsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseInsightsSortByEnum
func GetListDatabaseInsightsSortByEnumStringValues() []string {
	return []string{
		"databaseName",
		"databaseDisplayName",
		"databaseType",
	}
}

// GetMappingListDatabaseInsightsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseInsightsSortByEnum(val string) (ListDatabaseInsightsSortByEnum, bool) {
	enum, ok := mappingListDatabaseInsightsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
