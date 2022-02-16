// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListCustomPropertiesRequest wrapper for the ListCustomProperties operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListCustomProperties.go.html to see an example of how to use ListCustomPropertiesRequest.
type ListCustomPropertiesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique namespace identifier.
	NamespaceId *string `mandatory:"true" contributesTo:"path" name:"namespaceId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Return the custom properties which has specified data types
	DataTypes []CustomPropertyDataTypeEnum `contributesTo:"query" name:"dataTypes" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the entire type name given. The match is not case sensitive
	TypeName []ListCustomPropertiesTypeNameEnum `contributesTo:"query" name:"typeName" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListCustomPropertiesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Specifies the fields to return in a custom property summary response.
	Fields []ListCustomPropertiesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCustomPropertiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for USAGECOUNT and DISPLAYNAME is Ascending
	SortBy ListCustomPropertiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCustomPropertiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCustomPropertiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCustomPropertiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCustomPropertiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCustomPropertiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.DataTypes {
		if _, ok := GetMappingCustomPropertyDataTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataTypes: %s. Supported values are: %s.", val, strings.Join(GetCustomPropertyDataTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.TypeName {
		if _, ok := GetMappingListCustomPropertiesTypeNameEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TypeName: %s. Supported values are: %s.", val, strings.Join(GetListCustomPropertiesTypeNameEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListCustomPropertiesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListCustomPropertiesLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListCustomPropertiesFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListCustomPropertiesFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListCustomPropertiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCustomPropertiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCustomPropertiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCustomPropertiesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCustomPropertiesResponse wrapper for the ListCustomProperties operation
type ListCustomPropertiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CustomPropertyCollection instances
	CustomPropertyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCustomPropertiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCustomPropertiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCustomPropertiesTypeNameEnum Enum with underlying type: string
type ListCustomPropertiesTypeNameEnum string

// Set of constants representing the allowable values for ListCustomPropertiesTypeNameEnum
const (
	ListCustomPropertiesTypeNameDataAsset                        ListCustomPropertiesTypeNameEnum = "DATA_ASSET"
	ListCustomPropertiesTypeNameAutonomousDataWarehouse          ListCustomPropertiesTypeNameEnum = "AUTONOMOUS_DATA_WAREHOUSE"
	ListCustomPropertiesTypeNameHive                             ListCustomPropertiesTypeNameEnum = "HIVE"
	ListCustomPropertiesTypeNameKafka                            ListCustomPropertiesTypeNameEnum = "KAFKA"
	ListCustomPropertiesTypeNameMysql                            ListCustomPropertiesTypeNameEnum = "MYSQL"
	ListCustomPropertiesTypeNameOracleObjectStorage              ListCustomPropertiesTypeNameEnum = "ORACLE_OBJECT_STORAGE"
	ListCustomPropertiesTypeNameAutonomousTransactionProcessing  ListCustomPropertiesTypeNameEnum = "AUTONOMOUS_TRANSACTION_PROCESSING"
	ListCustomPropertiesTypeNameOracle                           ListCustomPropertiesTypeNameEnum = "ORACLE"
	ListCustomPropertiesTypeNamePostgresql                       ListCustomPropertiesTypeNameEnum = "POSTGRESQL"
	ListCustomPropertiesTypeNameMicrosoftAzureSqlDatabase        ListCustomPropertiesTypeNameEnum = "MICROSOFT_AZURE_SQL_DATABASE"
	ListCustomPropertiesTypeNameMicrosoftSqlServer               ListCustomPropertiesTypeNameEnum = "MICROSOFT_SQL_SERVER"
	ListCustomPropertiesTypeNameIbmDb2                           ListCustomPropertiesTypeNameEnum = "IBM_DB2"
	ListCustomPropertiesTypeNameDataEntity                       ListCustomPropertiesTypeNameEnum = "DATA_ENTITY"
	ListCustomPropertiesTypeNameLogicalEntity                    ListCustomPropertiesTypeNameEnum = "LOGICAL_ENTITY"
	ListCustomPropertiesTypeNameTable                            ListCustomPropertiesTypeNameEnum = "TABLE"
	ListCustomPropertiesTypeNameView                             ListCustomPropertiesTypeNameEnum = "VIEW"
	ListCustomPropertiesTypeNameAttribute                        ListCustomPropertiesTypeNameEnum = "ATTRIBUTE"
	ListCustomPropertiesTypeNameFolder                           ListCustomPropertiesTypeNameEnum = "FOLDER"
	ListCustomPropertiesTypeNameOracleAnalyticsSubjectAreaColumn ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_SUBJECT_AREA_COLUMN"
	ListCustomPropertiesTypeNameOracleAnalyticsLogicalColumn     ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_LOGICAL_COLUMN"
	ListCustomPropertiesTypeNameOracleAnalyticsPhysicalColumn    ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_PHYSICAL_COLUMN"
	ListCustomPropertiesTypeNameOracleAnalyticsAnalysisColumn    ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_ANALYSIS_COLUMN"
	ListCustomPropertiesTypeNameOracleAnalyticsServer            ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_SERVER"
	ListCustomPropertiesTypeNameOracleAnalyticsCloud             ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_CLOUD"
	ListCustomPropertiesTypeNameOracleAnalyticsSubjectArea       ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_SUBJECT_AREA"
	ListCustomPropertiesTypeNameOracleAnalyticsDashboard         ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_DASHBOARD"
	ListCustomPropertiesTypeNameOracleAnalyticsBusinessModel     ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_BUSINESS_MODEL"
	ListCustomPropertiesTypeNameOracleAnalyticsPhysicalDatabase  ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_PHYSICAL_DATABASE"
	ListCustomPropertiesTypeNameOracleAnalyticsPhysicalSchema    ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_PHYSICAL_SCHEMA"
	ListCustomPropertiesTypeNameOracleAnalyticsPresentationTable ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_PRESENTATION_TABLE"
	ListCustomPropertiesTypeNameOracleAnalyticsLogicalTable      ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_LOGICAL_TABLE"
	ListCustomPropertiesTypeNameOracleAnalyticsPhysicalTable     ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_PHYSICAL_TABLE"
	ListCustomPropertiesTypeNameOracleAnalyticsAnalysis          ListCustomPropertiesTypeNameEnum = "ORACLE_ANALYTICS_ANALYSIS"
	ListCustomPropertiesTypeNameDatabaseSchema                   ListCustomPropertiesTypeNameEnum = "DATABASE_SCHEMA"
	ListCustomPropertiesTypeNameTopic                            ListCustomPropertiesTypeNameEnum = "TOPIC"
	ListCustomPropertiesTypeNameConnection                       ListCustomPropertiesTypeNameEnum = "CONNECTION"
	ListCustomPropertiesTypeNameGlossary                         ListCustomPropertiesTypeNameEnum = "GLOSSARY"
	ListCustomPropertiesTypeNameTerm                             ListCustomPropertiesTypeNameEnum = "TERM"
	ListCustomPropertiesTypeNameCategory                         ListCustomPropertiesTypeNameEnum = "CATEGORY"
	ListCustomPropertiesTypeNameFile                             ListCustomPropertiesTypeNameEnum = "FILE"
	ListCustomPropertiesTypeNameBucket                           ListCustomPropertiesTypeNameEnum = "BUCKET"
	ListCustomPropertiesTypeNameMessage                          ListCustomPropertiesTypeNameEnum = "MESSAGE"
	ListCustomPropertiesTypeNameUnrecognizedFile                 ListCustomPropertiesTypeNameEnum = "UNRECOGNIZED_FILE"
)

var mappingListCustomPropertiesTypeNameEnum = map[string]ListCustomPropertiesTypeNameEnum{
	"DATA_ASSET":                           ListCustomPropertiesTypeNameDataAsset,
	"AUTONOMOUS_DATA_WAREHOUSE":            ListCustomPropertiesTypeNameAutonomousDataWarehouse,
	"HIVE":                                 ListCustomPropertiesTypeNameHive,
	"KAFKA":                                ListCustomPropertiesTypeNameKafka,
	"MYSQL":                                ListCustomPropertiesTypeNameMysql,
	"ORACLE_OBJECT_STORAGE":                ListCustomPropertiesTypeNameOracleObjectStorage,
	"AUTONOMOUS_TRANSACTION_PROCESSING":    ListCustomPropertiesTypeNameAutonomousTransactionProcessing,
	"ORACLE":                               ListCustomPropertiesTypeNameOracle,
	"POSTGRESQL":                           ListCustomPropertiesTypeNamePostgresql,
	"MICROSOFT_AZURE_SQL_DATABASE":         ListCustomPropertiesTypeNameMicrosoftAzureSqlDatabase,
	"MICROSOFT_SQL_SERVER":                 ListCustomPropertiesTypeNameMicrosoftSqlServer,
	"IBM_DB2":                              ListCustomPropertiesTypeNameIbmDb2,
	"DATA_ENTITY":                          ListCustomPropertiesTypeNameDataEntity,
	"LOGICAL_ENTITY":                       ListCustomPropertiesTypeNameLogicalEntity,
	"TABLE":                                ListCustomPropertiesTypeNameTable,
	"VIEW":                                 ListCustomPropertiesTypeNameView,
	"ATTRIBUTE":                            ListCustomPropertiesTypeNameAttribute,
	"FOLDER":                               ListCustomPropertiesTypeNameFolder,
	"ORACLE_ANALYTICS_SUBJECT_AREA_COLUMN": ListCustomPropertiesTypeNameOracleAnalyticsSubjectAreaColumn,
	"ORACLE_ANALYTICS_LOGICAL_COLUMN":      ListCustomPropertiesTypeNameOracleAnalyticsLogicalColumn,
	"ORACLE_ANALYTICS_PHYSICAL_COLUMN":     ListCustomPropertiesTypeNameOracleAnalyticsPhysicalColumn,
	"ORACLE_ANALYTICS_ANALYSIS_COLUMN":     ListCustomPropertiesTypeNameOracleAnalyticsAnalysisColumn,
	"ORACLE_ANALYTICS_SERVER":              ListCustomPropertiesTypeNameOracleAnalyticsServer,
	"ORACLE_ANALYTICS_CLOUD":               ListCustomPropertiesTypeNameOracleAnalyticsCloud,
	"ORACLE_ANALYTICS_SUBJECT_AREA":        ListCustomPropertiesTypeNameOracleAnalyticsSubjectArea,
	"ORACLE_ANALYTICS_DASHBOARD":           ListCustomPropertiesTypeNameOracleAnalyticsDashboard,
	"ORACLE_ANALYTICS_BUSINESS_MODEL":      ListCustomPropertiesTypeNameOracleAnalyticsBusinessModel,
	"ORACLE_ANALYTICS_PHYSICAL_DATABASE":   ListCustomPropertiesTypeNameOracleAnalyticsPhysicalDatabase,
	"ORACLE_ANALYTICS_PHYSICAL_SCHEMA":     ListCustomPropertiesTypeNameOracleAnalyticsPhysicalSchema,
	"ORACLE_ANALYTICS_PRESENTATION_TABLE":  ListCustomPropertiesTypeNameOracleAnalyticsPresentationTable,
	"ORACLE_ANALYTICS_LOGICAL_TABLE":       ListCustomPropertiesTypeNameOracleAnalyticsLogicalTable,
	"ORACLE_ANALYTICS_PHYSICAL_TABLE":      ListCustomPropertiesTypeNameOracleAnalyticsPhysicalTable,
	"ORACLE_ANALYTICS_ANALYSIS":            ListCustomPropertiesTypeNameOracleAnalyticsAnalysis,
	"DATABASE_SCHEMA":                      ListCustomPropertiesTypeNameDatabaseSchema,
	"TOPIC":                                ListCustomPropertiesTypeNameTopic,
	"CONNECTION":                           ListCustomPropertiesTypeNameConnection,
	"GLOSSARY":                             ListCustomPropertiesTypeNameGlossary,
	"TERM":                                 ListCustomPropertiesTypeNameTerm,
	"CATEGORY":                             ListCustomPropertiesTypeNameCategory,
	"FILE":                                 ListCustomPropertiesTypeNameFile,
	"BUCKET":                               ListCustomPropertiesTypeNameBucket,
	"MESSAGE":                              ListCustomPropertiesTypeNameMessage,
	"UNRECOGNIZED_FILE":                    ListCustomPropertiesTypeNameUnrecognizedFile,
}

// GetListCustomPropertiesTypeNameEnumValues Enumerates the set of values for ListCustomPropertiesTypeNameEnum
func GetListCustomPropertiesTypeNameEnumValues() []ListCustomPropertiesTypeNameEnum {
	values := make([]ListCustomPropertiesTypeNameEnum, 0)
	for _, v := range mappingListCustomPropertiesTypeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomPropertiesTypeNameEnumStringValues Enumerates the set of values in String for ListCustomPropertiesTypeNameEnum
func GetListCustomPropertiesTypeNameEnumStringValues() []string {
	return []string{
		"DATA_ASSET",
		"AUTONOMOUS_DATA_WAREHOUSE",
		"HIVE",
		"KAFKA",
		"MYSQL",
		"ORACLE_OBJECT_STORAGE",
		"AUTONOMOUS_TRANSACTION_PROCESSING",
		"ORACLE",
		"POSTGRESQL",
		"MICROSOFT_AZURE_SQL_DATABASE",
		"MICROSOFT_SQL_SERVER",
		"IBM_DB2",
		"DATA_ENTITY",
		"LOGICAL_ENTITY",
		"TABLE",
		"VIEW",
		"ATTRIBUTE",
		"FOLDER",
		"ORACLE_ANALYTICS_SUBJECT_AREA_COLUMN",
		"ORACLE_ANALYTICS_LOGICAL_COLUMN",
		"ORACLE_ANALYTICS_PHYSICAL_COLUMN",
		"ORACLE_ANALYTICS_ANALYSIS_COLUMN",
		"ORACLE_ANALYTICS_SERVER",
		"ORACLE_ANALYTICS_CLOUD",
		"ORACLE_ANALYTICS_SUBJECT_AREA",
		"ORACLE_ANALYTICS_DASHBOARD",
		"ORACLE_ANALYTICS_BUSINESS_MODEL",
		"ORACLE_ANALYTICS_PHYSICAL_DATABASE",
		"ORACLE_ANALYTICS_PHYSICAL_SCHEMA",
		"ORACLE_ANALYTICS_PRESENTATION_TABLE",
		"ORACLE_ANALYTICS_LOGICAL_TABLE",
		"ORACLE_ANALYTICS_PHYSICAL_TABLE",
		"ORACLE_ANALYTICS_ANALYSIS",
		"DATABASE_SCHEMA",
		"TOPIC",
		"CONNECTION",
		"GLOSSARY",
		"TERM",
		"CATEGORY",
		"FILE",
		"BUCKET",
		"MESSAGE",
		"UNRECOGNIZED_FILE",
	}
}

// GetMappingListCustomPropertiesTypeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomPropertiesTypeNameEnum(val string) (ListCustomPropertiesTypeNameEnum, bool) {
	mappingListCustomPropertiesTypeNameEnumIgnoreCase := make(map[string]ListCustomPropertiesTypeNameEnum)
	for k, v := range mappingListCustomPropertiesTypeNameEnum {
		mappingListCustomPropertiesTypeNameEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCustomPropertiesTypeNameEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomPropertiesLifecycleStateEnum Enum with underlying type: string
type ListCustomPropertiesLifecycleStateEnum string

// Set of constants representing the allowable values for ListCustomPropertiesLifecycleStateEnum
const (
	ListCustomPropertiesLifecycleStateCreating ListCustomPropertiesLifecycleStateEnum = "CREATING"
	ListCustomPropertiesLifecycleStateActive   ListCustomPropertiesLifecycleStateEnum = "ACTIVE"
	ListCustomPropertiesLifecycleStateInactive ListCustomPropertiesLifecycleStateEnum = "INACTIVE"
	ListCustomPropertiesLifecycleStateUpdating ListCustomPropertiesLifecycleStateEnum = "UPDATING"
	ListCustomPropertiesLifecycleStateDeleting ListCustomPropertiesLifecycleStateEnum = "DELETING"
	ListCustomPropertiesLifecycleStateDeleted  ListCustomPropertiesLifecycleStateEnum = "DELETED"
	ListCustomPropertiesLifecycleStateFailed   ListCustomPropertiesLifecycleStateEnum = "FAILED"
	ListCustomPropertiesLifecycleStateMoving   ListCustomPropertiesLifecycleStateEnum = "MOVING"
)

var mappingListCustomPropertiesLifecycleStateEnum = map[string]ListCustomPropertiesLifecycleStateEnum{
	"CREATING": ListCustomPropertiesLifecycleStateCreating,
	"ACTIVE":   ListCustomPropertiesLifecycleStateActive,
	"INACTIVE": ListCustomPropertiesLifecycleStateInactive,
	"UPDATING": ListCustomPropertiesLifecycleStateUpdating,
	"DELETING": ListCustomPropertiesLifecycleStateDeleting,
	"DELETED":  ListCustomPropertiesLifecycleStateDeleted,
	"FAILED":   ListCustomPropertiesLifecycleStateFailed,
	"MOVING":   ListCustomPropertiesLifecycleStateMoving,
}

// GetListCustomPropertiesLifecycleStateEnumValues Enumerates the set of values for ListCustomPropertiesLifecycleStateEnum
func GetListCustomPropertiesLifecycleStateEnumValues() []ListCustomPropertiesLifecycleStateEnum {
	values := make([]ListCustomPropertiesLifecycleStateEnum, 0)
	for _, v := range mappingListCustomPropertiesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomPropertiesLifecycleStateEnumStringValues Enumerates the set of values in String for ListCustomPropertiesLifecycleStateEnum
func GetListCustomPropertiesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"MOVING",
	}
}

// GetMappingListCustomPropertiesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomPropertiesLifecycleStateEnum(val string) (ListCustomPropertiesLifecycleStateEnum, bool) {
	mappingListCustomPropertiesLifecycleStateEnumIgnoreCase := make(map[string]ListCustomPropertiesLifecycleStateEnum)
	for k, v := range mappingListCustomPropertiesLifecycleStateEnum {
		mappingListCustomPropertiesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCustomPropertiesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomPropertiesFieldsEnum Enum with underlying type: string
type ListCustomPropertiesFieldsEnum string

// Set of constants representing the allowable values for ListCustomPropertiesFieldsEnum
const (
	ListCustomPropertiesFieldsKey            ListCustomPropertiesFieldsEnum = "key"
	ListCustomPropertiesFieldsDisplayname    ListCustomPropertiesFieldsEnum = "displayName"
	ListCustomPropertiesFieldsDescription    ListCustomPropertiesFieldsEnum = "description"
	ListCustomPropertiesFieldsDatatype       ListCustomPropertiesFieldsEnum = "dataType"
	ListCustomPropertiesFieldsNamespacename  ListCustomPropertiesFieldsEnum = "namespaceName"
	ListCustomPropertiesFieldsLifecyclestate ListCustomPropertiesFieldsEnum = "lifecycleState"
	ListCustomPropertiesFieldsTimecreated    ListCustomPropertiesFieldsEnum = "timeCreated"
)

var mappingListCustomPropertiesFieldsEnum = map[string]ListCustomPropertiesFieldsEnum{
	"key":            ListCustomPropertiesFieldsKey,
	"displayName":    ListCustomPropertiesFieldsDisplayname,
	"description":    ListCustomPropertiesFieldsDescription,
	"dataType":       ListCustomPropertiesFieldsDatatype,
	"namespaceName":  ListCustomPropertiesFieldsNamespacename,
	"lifecycleState": ListCustomPropertiesFieldsLifecyclestate,
	"timeCreated":    ListCustomPropertiesFieldsTimecreated,
}

// GetListCustomPropertiesFieldsEnumValues Enumerates the set of values for ListCustomPropertiesFieldsEnum
func GetListCustomPropertiesFieldsEnumValues() []ListCustomPropertiesFieldsEnum {
	values := make([]ListCustomPropertiesFieldsEnum, 0)
	for _, v := range mappingListCustomPropertiesFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomPropertiesFieldsEnumStringValues Enumerates the set of values in String for ListCustomPropertiesFieldsEnum
func GetListCustomPropertiesFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"dataType",
		"namespaceName",
		"lifecycleState",
		"timeCreated",
	}
}

// GetMappingListCustomPropertiesFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomPropertiesFieldsEnum(val string) (ListCustomPropertiesFieldsEnum, bool) {
	mappingListCustomPropertiesFieldsEnumIgnoreCase := make(map[string]ListCustomPropertiesFieldsEnum)
	for k, v := range mappingListCustomPropertiesFieldsEnum {
		mappingListCustomPropertiesFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCustomPropertiesFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomPropertiesSortOrderEnum Enum with underlying type: string
type ListCustomPropertiesSortOrderEnum string

// Set of constants representing the allowable values for ListCustomPropertiesSortOrderEnum
const (
	ListCustomPropertiesSortOrderAsc  ListCustomPropertiesSortOrderEnum = "ASC"
	ListCustomPropertiesSortOrderDesc ListCustomPropertiesSortOrderEnum = "DESC"
)

var mappingListCustomPropertiesSortOrderEnum = map[string]ListCustomPropertiesSortOrderEnum{
	"ASC":  ListCustomPropertiesSortOrderAsc,
	"DESC": ListCustomPropertiesSortOrderDesc,
}

// GetListCustomPropertiesSortOrderEnumValues Enumerates the set of values for ListCustomPropertiesSortOrderEnum
func GetListCustomPropertiesSortOrderEnumValues() []ListCustomPropertiesSortOrderEnum {
	values := make([]ListCustomPropertiesSortOrderEnum, 0)
	for _, v := range mappingListCustomPropertiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomPropertiesSortOrderEnumStringValues Enumerates the set of values in String for ListCustomPropertiesSortOrderEnum
func GetListCustomPropertiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCustomPropertiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomPropertiesSortOrderEnum(val string) (ListCustomPropertiesSortOrderEnum, bool) {
	mappingListCustomPropertiesSortOrderEnumIgnoreCase := make(map[string]ListCustomPropertiesSortOrderEnum)
	for k, v := range mappingListCustomPropertiesSortOrderEnum {
		mappingListCustomPropertiesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCustomPropertiesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListCustomPropertiesSortByEnum Enum with underlying type: string
type ListCustomPropertiesSortByEnum string

// Set of constants representing the allowable values for ListCustomPropertiesSortByEnum
const (
	ListCustomPropertiesSortByDisplayname ListCustomPropertiesSortByEnum = "DISPLAYNAME"
	ListCustomPropertiesSortByUsagecount  ListCustomPropertiesSortByEnum = "USAGECOUNT"
)

var mappingListCustomPropertiesSortByEnum = map[string]ListCustomPropertiesSortByEnum{
	"DISPLAYNAME": ListCustomPropertiesSortByDisplayname,
	"USAGECOUNT":  ListCustomPropertiesSortByUsagecount,
}

// GetListCustomPropertiesSortByEnumValues Enumerates the set of values for ListCustomPropertiesSortByEnum
func GetListCustomPropertiesSortByEnumValues() []ListCustomPropertiesSortByEnum {
	values := make([]ListCustomPropertiesSortByEnum, 0)
	for _, v := range mappingListCustomPropertiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCustomPropertiesSortByEnumStringValues Enumerates the set of values in String for ListCustomPropertiesSortByEnum
func GetListCustomPropertiesSortByEnumStringValues() []string {
	return []string{
		"DISPLAYNAME",
		"USAGECOUNT",
	}
}

// GetMappingListCustomPropertiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCustomPropertiesSortByEnum(val string) (ListCustomPropertiesSortByEnum, bool) {
	mappingListCustomPropertiesSortByEnumIgnoreCase := make(map[string]ListCustomPropertiesSortByEnum)
	for k, v := range mappingListCustomPropertiesSortByEnum {
		mappingListCustomPropertiesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCustomPropertiesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
