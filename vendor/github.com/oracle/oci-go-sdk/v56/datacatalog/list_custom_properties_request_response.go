// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
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

var mappingListCustomPropertiesTypeName = map[string]ListCustomPropertiesTypeNameEnum{
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
	for _, v := range mappingListCustomPropertiesTypeName {
		values = append(values, v)
	}
	return values
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

var mappingListCustomPropertiesLifecycleState = map[string]ListCustomPropertiesLifecycleStateEnum{
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
	for _, v := range mappingListCustomPropertiesLifecycleState {
		values = append(values, v)
	}
	return values
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

var mappingListCustomPropertiesFields = map[string]ListCustomPropertiesFieldsEnum{
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
	for _, v := range mappingListCustomPropertiesFields {
		values = append(values, v)
	}
	return values
}

// ListCustomPropertiesSortOrderEnum Enum with underlying type: string
type ListCustomPropertiesSortOrderEnum string

// Set of constants representing the allowable values for ListCustomPropertiesSortOrderEnum
const (
	ListCustomPropertiesSortOrderAsc  ListCustomPropertiesSortOrderEnum = "ASC"
	ListCustomPropertiesSortOrderDesc ListCustomPropertiesSortOrderEnum = "DESC"
)

var mappingListCustomPropertiesSortOrder = map[string]ListCustomPropertiesSortOrderEnum{
	"ASC":  ListCustomPropertiesSortOrderAsc,
	"DESC": ListCustomPropertiesSortOrderDesc,
}

// GetListCustomPropertiesSortOrderEnumValues Enumerates the set of values for ListCustomPropertiesSortOrderEnum
func GetListCustomPropertiesSortOrderEnumValues() []ListCustomPropertiesSortOrderEnum {
	values := make([]ListCustomPropertiesSortOrderEnum, 0)
	for _, v := range mappingListCustomPropertiesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListCustomPropertiesSortByEnum Enum with underlying type: string
type ListCustomPropertiesSortByEnum string

// Set of constants representing the allowable values for ListCustomPropertiesSortByEnum
const (
	ListCustomPropertiesSortByDisplayname ListCustomPropertiesSortByEnum = "DISPLAYNAME"
	ListCustomPropertiesSortByUsagecount  ListCustomPropertiesSortByEnum = "USAGECOUNT"
)

var mappingListCustomPropertiesSortBy = map[string]ListCustomPropertiesSortByEnum{
	"DISPLAYNAME": ListCustomPropertiesSortByDisplayname,
	"USAGECOUNT":  ListCustomPropertiesSortByUsagecount,
}

// GetListCustomPropertiesSortByEnumValues Enumerates the set of values for ListCustomPropertiesSortByEnum
func GetListCustomPropertiesSortByEnumValues() []ListCustomPropertiesSortByEnum {
	values := make([]ListCustomPropertiesSortByEnum, 0)
	for _, v := range mappingListCustomPropertiesSortBy {
		values = append(values, v)
	}
	return values
}
