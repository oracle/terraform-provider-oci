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

// ListDatabaseConfigurationsRequest wrapper for the ListDatabaseConfigurations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseConfigurations.go.html to see an example of how to use ListDatabaseConfigurationsRequest.
type ListDatabaseConfigurationsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Unique Enterprise Manager bridge identifier
	EnterpriseManagerBridgeId *string `mandatory:"false" contributesTo:"query" name:"enterpriseManagerBridgeId"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []ListDatabaseConfigurationsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

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
	SortOrder ListDatabaseConfigurationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Database configuration list sort options. If `fields` parameter is selected, the `sortBy` parameter must be one of the fields specified.
	SortBy ListDatabaseConfigurationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Optional list of Exadata Insight VM cluster name.
	VmclusterName []string `contributesTo:"query" name:"vmclusterName" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.DatabaseType {
		if _, ok := GetMappingListDatabaseConfigurationsDatabaseTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", val, strings.Join(GetListDatabaseConfigurationsDatabaseTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListDatabaseConfigurationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseConfigurationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseConfigurationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseConfigurationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseConfigurationsResponse wrapper for the ListDatabaseConfigurations operation
type ListDatabaseConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseConfigurationCollection instances
	DatabaseConfigurationCollection `presentIn:"body"`

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

func (response ListDatabaseConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseConfigurationsDatabaseTypeEnum Enum with underlying type: string
type ListDatabaseConfigurationsDatabaseTypeEnum string

// Set of constants representing the allowable values for ListDatabaseConfigurationsDatabaseTypeEnum
const (
	ListDatabaseConfigurationsDatabaseTypeAdwS                 ListDatabaseConfigurationsDatabaseTypeEnum = "ADW-S"
	ListDatabaseConfigurationsDatabaseTypeAtpS                 ListDatabaseConfigurationsDatabaseTypeEnum = "ATP-S"
	ListDatabaseConfigurationsDatabaseTypeAdwD                 ListDatabaseConfigurationsDatabaseTypeEnum = "ADW-D"
	ListDatabaseConfigurationsDatabaseTypeAtpD                 ListDatabaseConfigurationsDatabaseTypeEnum = "ATP-D"
	ListDatabaseConfigurationsDatabaseTypeExternalPdb          ListDatabaseConfigurationsDatabaseTypeEnum = "EXTERNAL-PDB"
	ListDatabaseConfigurationsDatabaseTypeExternalNoncdb       ListDatabaseConfigurationsDatabaseTypeEnum = "EXTERNAL-NONCDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedVmCdb       ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-VM-CDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedVmPdb       ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-VM-PDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedVmNoncdb    ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-VM-NONCDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedBmCdb       ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-BM-CDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedBmPdb       ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-BM-PDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedBmNoncdb    ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-BM-NONCDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedExacsCdb    ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-EXACS-CDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedExacsPdb    ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-EXACS-PDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedExacsNoncdb ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-EXACS-NONCDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedExaccCdb    ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-EXACC-CDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedExaccPdb    ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-EXACC-PDB"
	ListDatabaseConfigurationsDatabaseTypeComanagedExaccNoncdb ListDatabaseConfigurationsDatabaseTypeEnum = "COMANAGED-EXACC-NONCDB"
	ListDatabaseConfigurationsDatabaseTypeMdsMysql             ListDatabaseConfigurationsDatabaseTypeEnum = "MDS-MYSQL"
)

var mappingListDatabaseConfigurationsDatabaseTypeEnum = map[string]ListDatabaseConfigurationsDatabaseTypeEnum{
	"ADW-S":                  ListDatabaseConfigurationsDatabaseTypeAdwS,
	"ATP-S":                  ListDatabaseConfigurationsDatabaseTypeAtpS,
	"ADW-D":                  ListDatabaseConfigurationsDatabaseTypeAdwD,
	"ATP-D":                  ListDatabaseConfigurationsDatabaseTypeAtpD,
	"EXTERNAL-PDB":           ListDatabaseConfigurationsDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB":        ListDatabaseConfigurationsDatabaseTypeExternalNoncdb,
	"COMANAGED-VM-CDB":       ListDatabaseConfigurationsDatabaseTypeComanagedVmCdb,
	"COMANAGED-VM-PDB":       ListDatabaseConfigurationsDatabaseTypeComanagedVmPdb,
	"COMANAGED-VM-NONCDB":    ListDatabaseConfigurationsDatabaseTypeComanagedVmNoncdb,
	"COMANAGED-BM-CDB":       ListDatabaseConfigurationsDatabaseTypeComanagedBmCdb,
	"COMANAGED-BM-PDB":       ListDatabaseConfigurationsDatabaseTypeComanagedBmPdb,
	"COMANAGED-BM-NONCDB":    ListDatabaseConfigurationsDatabaseTypeComanagedBmNoncdb,
	"COMANAGED-EXACS-CDB":    ListDatabaseConfigurationsDatabaseTypeComanagedExacsCdb,
	"COMANAGED-EXACS-PDB":    ListDatabaseConfigurationsDatabaseTypeComanagedExacsPdb,
	"COMANAGED-EXACS-NONCDB": ListDatabaseConfigurationsDatabaseTypeComanagedExacsNoncdb,
	"COMANAGED-EXACC-CDB":    ListDatabaseConfigurationsDatabaseTypeComanagedExaccCdb,
	"COMANAGED-EXACC-PDB":    ListDatabaseConfigurationsDatabaseTypeComanagedExaccPdb,
	"COMANAGED-EXACC-NONCDB": ListDatabaseConfigurationsDatabaseTypeComanagedExaccNoncdb,
	"MDS-MYSQL":              ListDatabaseConfigurationsDatabaseTypeMdsMysql,
}

var mappingListDatabaseConfigurationsDatabaseTypeEnumLowerCase = map[string]ListDatabaseConfigurationsDatabaseTypeEnum{
	"adw-s":                  ListDatabaseConfigurationsDatabaseTypeAdwS,
	"atp-s":                  ListDatabaseConfigurationsDatabaseTypeAtpS,
	"adw-d":                  ListDatabaseConfigurationsDatabaseTypeAdwD,
	"atp-d":                  ListDatabaseConfigurationsDatabaseTypeAtpD,
	"external-pdb":           ListDatabaseConfigurationsDatabaseTypeExternalPdb,
	"external-noncdb":        ListDatabaseConfigurationsDatabaseTypeExternalNoncdb,
	"comanaged-vm-cdb":       ListDatabaseConfigurationsDatabaseTypeComanagedVmCdb,
	"comanaged-vm-pdb":       ListDatabaseConfigurationsDatabaseTypeComanagedVmPdb,
	"comanaged-vm-noncdb":    ListDatabaseConfigurationsDatabaseTypeComanagedVmNoncdb,
	"comanaged-bm-cdb":       ListDatabaseConfigurationsDatabaseTypeComanagedBmCdb,
	"comanaged-bm-pdb":       ListDatabaseConfigurationsDatabaseTypeComanagedBmPdb,
	"comanaged-bm-noncdb":    ListDatabaseConfigurationsDatabaseTypeComanagedBmNoncdb,
	"comanaged-exacs-cdb":    ListDatabaseConfigurationsDatabaseTypeComanagedExacsCdb,
	"comanaged-exacs-pdb":    ListDatabaseConfigurationsDatabaseTypeComanagedExacsPdb,
	"comanaged-exacs-noncdb": ListDatabaseConfigurationsDatabaseTypeComanagedExacsNoncdb,
	"comanaged-exacc-cdb":    ListDatabaseConfigurationsDatabaseTypeComanagedExaccCdb,
	"comanaged-exacc-pdb":    ListDatabaseConfigurationsDatabaseTypeComanagedExaccPdb,
	"comanaged-exacc-noncdb": ListDatabaseConfigurationsDatabaseTypeComanagedExaccNoncdb,
	"mds-mysql":              ListDatabaseConfigurationsDatabaseTypeMdsMysql,
}

// GetListDatabaseConfigurationsDatabaseTypeEnumValues Enumerates the set of values for ListDatabaseConfigurationsDatabaseTypeEnum
func GetListDatabaseConfigurationsDatabaseTypeEnumValues() []ListDatabaseConfigurationsDatabaseTypeEnum {
	values := make([]ListDatabaseConfigurationsDatabaseTypeEnum, 0)
	for _, v := range mappingListDatabaseConfigurationsDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseConfigurationsDatabaseTypeEnumStringValues Enumerates the set of values in String for ListDatabaseConfigurationsDatabaseTypeEnum
func GetListDatabaseConfigurationsDatabaseTypeEnumStringValues() []string {
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
		"COMANAGED-EXACC-CDB",
		"COMANAGED-EXACC-PDB",
		"COMANAGED-EXACC-NONCDB",
		"MDS-MYSQL",
	}
}

// GetMappingListDatabaseConfigurationsDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseConfigurationsDatabaseTypeEnum(val string) (ListDatabaseConfigurationsDatabaseTypeEnum, bool) {
	enum, ok := mappingListDatabaseConfigurationsDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseConfigurationsSortOrderEnum Enum with underlying type: string
type ListDatabaseConfigurationsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseConfigurationsSortOrderEnum
const (
	ListDatabaseConfigurationsSortOrderAsc  ListDatabaseConfigurationsSortOrderEnum = "ASC"
	ListDatabaseConfigurationsSortOrderDesc ListDatabaseConfigurationsSortOrderEnum = "DESC"
)

var mappingListDatabaseConfigurationsSortOrderEnum = map[string]ListDatabaseConfigurationsSortOrderEnum{
	"ASC":  ListDatabaseConfigurationsSortOrderAsc,
	"DESC": ListDatabaseConfigurationsSortOrderDesc,
}

var mappingListDatabaseConfigurationsSortOrderEnumLowerCase = map[string]ListDatabaseConfigurationsSortOrderEnum{
	"asc":  ListDatabaseConfigurationsSortOrderAsc,
	"desc": ListDatabaseConfigurationsSortOrderDesc,
}

// GetListDatabaseConfigurationsSortOrderEnumValues Enumerates the set of values for ListDatabaseConfigurationsSortOrderEnum
func GetListDatabaseConfigurationsSortOrderEnumValues() []ListDatabaseConfigurationsSortOrderEnum {
	values := make([]ListDatabaseConfigurationsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseConfigurationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseConfigurationsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseConfigurationsSortOrderEnum
func GetListDatabaseConfigurationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseConfigurationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseConfigurationsSortOrderEnum(val string) (ListDatabaseConfigurationsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseConfigurationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseConfigurationsSortByEnum Enum with underlying type: string
type ListDatabaseConfigurationsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseConfigurationsSortByEnum
const (
	ListDatabaseConfigurationsSortByDatabasename        ListDatabaseConfigurationsSortByEnum = "databaseName"
	ListDatabaseConfigurationsSortByDatabasedisplayname ListDatabaseConfigurationsSortByEnum = "databaseDisplayName"
	ListDatabaseConfigurationsSortByDatabasetype        ListDatabaseConfigurationsSortByEnum = "databaseType"
)

var mappingListDatabaseConfigurationsSortByEnum = map[string]ListDatabaseConfigurationsSortByEnum{
	"databaseName":        ListDatabaseConfigurationsSortByDatabasename,
	"databaseDisplayName": ListDatabaseConfigurationsSortByDatabasedisplayname,
	"databaseType":        ListDatabaseConfigurationsSortByDatabasetype,
}

var mappingListDatabaseConfigurationsSortByEnumLowerCase = map[string]ListDatabaseConfigurationsSortByEnum{
	"databasename":        ListDatabaseConfigurationsSortByDatabasename,
	"databasedisplayname": ListDatabaseConfigurationsSortByDatabasedisplayname,
	"databasetype":        ListDatabaseConfigurationsSortByDatabasetype,
}

// GetListDatabaseConfigurationsSortByEnumValues Enumerates the set of values for ListDatabaseConfigurationsSortByEnum
func GetListDatabaseConfigurationsSortByEnumValues() []ListDatabaseConfigurationsSortByEnum {
	values := make([]ListDatabaseConfigurationsSortByEnum, 0)
	for _, v := range mappingListDatabaseConfigurationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseConfigurationsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseConfigurationsSortByEnum
func GetListDatabaseConfigurationsSortByEnumStringValues() []string {
	return []string{
		"databaseName",
		"databaseDisplayName",
		"databaseType",
	}
}

// GetMappingListDatabaseConfigurationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseConfigurationsSortByEnum(val string) (ListDatabaseConfigurationsSortByEnum, bool) {
	enum, ok := mappingListDatabaseConfigurationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
