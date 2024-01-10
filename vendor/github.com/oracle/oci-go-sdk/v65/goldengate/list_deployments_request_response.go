// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDeploymentsRequest wrapper for the ListDeployments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeployments.go.html to see an example of how to use ListDeploymentsRequest.
type ListDeploymentsRequest struct {

	// The OCID of the compartment that contains the work request. Work requests should be scoped
	// to the same compartment as the resource the work request affects. If the work request concerns
	// multiple resources, and those resources are not in the same compartment, it is up to the service team
	// to pick the primary resource whose compartment should be used.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The connection type which the deployment must support.
	SupportedConnectionType ListDeploymentsSupportedConnectionTypeEnum `mandatory:"false" contributesTo:"query" name:"supportedConnectionType" omitEmpty:"true"`

	// The OCID of the connection which for the deployment must be assigned.
	AssignedConnectionId *string `mandatory:"false" contributesTo:"query" name:"assignedConnectionId"`

	// Return the deployments to which the specified connectionId may be assigned.
	AssignableConnectionId *string `mandatory:"false" contributesTo:"query" name:"assignableConnectionId"`

	// A filter to return only the resources that match the 'lifecycleState' given.
	LifecycleState ListDeploymentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the 'lifecycleSubState' given.
	LifecycleSubState ListDeploymentsLifecycleSubStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleSubState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the 'fqdn' given.
	Fqdn *string `mandatory:"false" contributesTo:"query" name:"fqdn"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDeploymentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListDeploymentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDeploymentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDeploymentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDeploymentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDeploymentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDeploymentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDeploymentsSupportedConnectionTypeEnum(string(request.SupportedConnectionType)); !ok && request.SupportedConnectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SupportedConnectionType: %s. Supported values are: %s.", request.SupportedConnectionType, strings.Join(GetListDeploymentsSupportedConnectionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDeploymentsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentsLifecycleSubStateEnum(string(request.LifecycleSubState)); !ok && request.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", request.LifecycleSubState, strings.Join(GetListDeploymentsLifecycleSubStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDeploymentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDeploymentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDeploymentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDeploymentsResponse wrapper for the ListDeployments operation
type ListDeploymentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DeploymentCollection instances
	DeploymentCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDeploymentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDeploymentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDeploymentsSupportedConnectionTypeEnum Enum with underlying type: string
type ListDeploymentsSupportedConnectionTypeEnum string

// Set of constants representing the allowable values for ListDeploymentsSupportedConnectionTypeEnum
const (
	ListDeploymentsSupportedConnectionTypeGoldengate            ListDeploymentsSupportedConnectionTypeEnum = "GOLDENGATE"
	ListDeploymentsSupportedConnectionTypeKafka                 ListDeploymentsSupportedConnectionTypeEnum = "KAFKA"
	ListDeploymentsSupportedConnectionTypeKafkaSchemaRegistry   ListDeploymentsSupportedConnectionTypeEnum = "KAFKA_SCHEMA_REGISTRY"
	ListDeploymentsSupportedConnectionTypeMysql                 ListDeploymentsSupportedConnectionTypeEnum = "MYSQL"
	ListDeploymentsSupportedConnectionTypeJavaMessageService    ListDeploymentsSupportedConnectionTypeEnum = "JAVA_MESSAGE_SERVICE"
	ListDeploymentsSupportedConnectionTypeMicrosoftSqlserver    ListDeploymentsSupportedConnectionTypeEnum = "MICROSOFT_SQLSERVER"
	ListDeploymentsSupportedConnectionTypeOciObjectStorage      ListDeploymentsSupportedConnectionTypeEnum = "OCI_OBJECT_STORAGE"
	ListDeploymentsSupportedConnectionTypeOracle                ListDeploymentsSupportedConnectionTypeEnum = "ORACLE"
	ListDeploymentsSupportedConnectionTypeAzureDataLakeStorage  ListDeploymentsSupportedConnectionTypeEnum = "AZURE_DATA_LAKE_STORAGE"
	ListDeploymentsSupportedConnectionTypePostgresql            ListDeploymentsSupportedConnectionTypeEnum = "POSTGRESQL"
	ListDeploymentsSupportedConnectionTypeAzureSynapseAnalytics ListDeploymentsSupportedConnectionTypeEnum = "AZURE_SYNAPSE_ANALYTICS"
	ListDeploymentsSupportedConnectionTypeSnowflake             ListDeploymentsSupportedConnectionTypeEnum = "SNOWFLAKE"
	ListDeploymentsSupportedConnectionTypeAmazonS3              ListDeploymentsSupportedConnectionTypeEnum = "AMAZON_S3"
	ListDeploymentsSupportedConnectionTypeHdfs                  ListDeploymentsSupportedConnectionTypeEnum = "HDFS"
	ListDeploymentsSupportedConnectionTypeOracleNosql           ListDeploymentsSupportedConnectionTypeEnum = "ORACLE_NOSQL"
	ListDeploymentsSupportedConnectionTypeMongodb               ListDeploymentsSupportedConnectionTypeEnum = "MONGODB"
	ListDeploymentsSupportedConnectionTypeAmazonKinesis         ListDeploymentsSupportedConnectionTypeEnum = "AMAZON_KINESIS"
	ListDeploymentsSupportedConnectionTypeAmazonRedshift        ListDeploymentsSupportedConnectionTypeEnum = "AMAZON_REDSHIFT"
	ListDeploymentsSupportedConnectionTypeRedis                 ListDeploymentsSupportedConnectionTypeEnum = "REDIS"
	ListDeploymentsSupportedConnectionTypeElasticsearch         ListDeploymentsSupportedConnectionTypeEnum = "ELASTICSEARCH"
	ListDeploymentsSupportedConnectionTypeGeneric               ListDeploymentsSupportedConnectionTypeEnum = "GENERIC"
	ListDeploymentsSupportedConnectionTypeGoogleCloudStorage    ListDeploymentsSupportedConnectionTypeEnum = "GOOGLE_CLOUD_STORAGE"
	ListDeploymentsSupportedConnectionTypeGoogleBigquery        ListDeploymentsSupportedConnectionTypeEnum = "GOOGLE_BIGQUERY"
)

var mappingListDeploymentsSupportedConnectionTypeEnum = map[string]ListDeploymentsSupportedConnectionTypeEnum{
	"GOLDENGATE":              ListDeploymentsSupportedConnectionTypeGoldengate,
	"KAFKA":                   ListDeploymentsSupportedConnectionTypeKafka,
	"KAFKA_SCHEMA_REGISTRY":   ListDeploymentsSupportedConnectionTypeKafkaSchemaRegistry,
	"MYSQL":                   ListDeploymentsSupportedConnectionTypeMysql,
	"JAVA_MESSAGE_SERVICE":    ListDeploymentsSupportedConnectionTypeJavaMessageService,
	"MICROSOFT_SQLSERVER":     ListDeploymentsSupportedConnectionTypeMicrosoftSqlserver,
	"OCI_OBJECT_STORAGE":      ListDeploymentsSupportedConnectionTypeOciObjectStorage,
	"ORACLE":                  ListDeploymentsSupportedConnectionTypeOracle,
	"AZURE_DATA_LAKE_STORAGE": ListDeploymentsSupportedConnectionTypeAzureDataLakeStorage,
	"POSTGRESQL":              ListDeploymentsSupportedConnectionTypePostgresql,
	"AZURE_SYNAPSE_ANALYTICS": ListDeploymentsSupportedConnectionTypeAzureSynapseAnalytics,
	"SNOWFLAKE":               ListDeploymentsSupportedConnectionTypeSnowflake,
	"AMAZON_S3":               ListDeploymentsSupportedConnectionTypeAmazonS3,
	"HDFS":                    ListDeploymentsSupportedConnectionTypeHdfs,
	"ORACLE_NOSQL":            ListDeploymentsSupportedConnectionTypeOracleNosql,
	"MONGODB":                 ListDeploymentsSupportedConnectionTypeMongodb,
	"AMAZON_KINESIS":          ListDeploymentsSupportedConnectionTypeAmazonKinesis,
	"AMAZON_REDSHIFT":         ListDeploymentsSupportedConnectionTypeAmazonRedshift,
	"REDIS":                   ListDeploymentsSupportedConnectionTypeRedis,
	"ELASTICSEARCH":           ListDeploymentsSupportedConnectionTypeElasticsearch,
	"GENERIC":                 ListDeploymentsSupportedConnectionTypeGeneric,
	"GOOGLE_CLOUD_STORAGE":    ListDeploymentsSupportedConnectionTypeGoogleCloudStorage,
	"GOOGLE_BIGQUERY":         ListDeploymentsSupportedConnectionTypeGoogleBigquery,
}

var mappingListDeploymentsSupportedConnectionTypeEnumLowerCase = map[string]ListDeploymentsSupportedConnectionTypeEnum{
	"goldengate":              ListDeploymentsSupportedConnectionTypeGoldengate,
	"kafka":                   ListDeploymentsSupportedConnectionTypeKafka,
	"kafka_schema_registry":   ListDeploymentsSupportedConnectionTypeKafkaSchemaRegistry,
	"mysql":                   ListDeploymentsSupportedConnectionTypeMysql,
	"java_message_service":    ListDeploymentsSupportedConnectionTypeJavaMessageService,
	"microsoft_sqlserver":     ListDeploymentsSupportedConnectionTypeMicrosoftSqlserver,
	"oci_object_storage":      ListDeploymentsSupportedConnectionTypeOciObjectStorage,
	"oracle":                  ListDeploymentsSupportedConnectionTypeOracle,
	"azure_data_lake_storage": ListDeploymentsSupportedConnectionTypeAzureDataLakeStorage,
	"postgresql":              ListDeploymentsSupportedConnectionTypePostgresql,
	"azure_synapse_analytics": ListDeploymentsSupportedConnectionTypeAzureSynapseAnalytics,
	"snowflake":               ListDeploymentsSupportedConnectionTypeSnowflake,
	"amazon_s3":               ListDeploymentsSupportedConnectionTypeAmazonS3,
	"hdfs":                    ListDeploymentsSupportedConnectionTypeHdfs,
	"oracle_nosql":            ListDeploymentsSupportedConnectionTypeOracleNosql,
	"mongodb":                 ListDeploymentsSupportedConnectionTypeMongodb,
	"amazon_kinesis":          ListDeploymentsSupportedConnectionTypeAmazonKinesis,
	"amazon_redshift":         ListDeploymentsSupportedConnectionTypeAmazonRedshift,
	"redis":                   ListDeploymentsSupportedConnectionTypeRedis,
	"elasticsearch":           ListDeploymentsSupportedConnectionTypeElasticsearch,
	"generic":                 ListDeploymentsSupportedConnectionTypeGeneric,
	"google_cloud_storage":    ListDeploymentsSupportedConnectionTypeGoogleCloudStorage,
	"google_bigquery":         ListDeploymentsSupportedConnectionTypeGoogleBigquery,
}

// GetListDeploymentsSupportedConnectionTypeEnumValues Enumerates the set of values for ListDeploymentsSupportedConnectionTypeEnum
func GetListDeploymentsSupportedConnectionTypeEnumValues() []ListDeploymentsSupportedConnectionTypeEnum {
	values := make([]ListDeploymentsSupportedConnectionTypeEnum, 0)
	for _, v := range mappingListDeploymentsSupportedConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsSupportedConnectionTypeEnumStringValues Enumerates the set of values in String for ListDeploymentsSupportedConnectionTypeEnum
func GetListDeploymentsSupportedConnectionTypeEnumStringValues() []string {
	return []string{
		"GOLDENGATE",
		"KAFKA",
		"KAFKA_SCHEMA_REGISTRY",
		"MYSQL",
		"JAVA_MESSAGE_SERVICE",
		"MICROSOFT_SQLSERVER",
		"OCI_OBJECT_STORAGE",
		"ORACLE",
		"AZURE_DATA_LAKE_STORAGE",
		"POSTGRESQL",
		"AZURE_SYNAPSE_ANALYTICS",
		"SNOWFLAKE",
		"AMAZON_S3",
		"HDFS",
		"ORACLE_NOSQL",
		"MONGODB",
		"AMAZON_KINESIS",
		"AMAZON_REDSHIFT",
		"REDIS",
		"ELASTICSEARCH",
		"GENERIC",
		"GOOGLE_CLOUD_STORAGE",
		"GOOGLE_BIGQUERY",
	}
}

// GetMappingListDeploymentsSupportedConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsSupportedConnectionTypeEnum(val string) (ListDeploymentsSupportedConnectionTypeEnum, bool) {
	enum, ok := mappingListDeploymentsSupportedConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentsLifecycleStateEnum Enum with underlying type: string
type ListDeploymentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDeploymentsLifecycleStateEnum
const (
	ListDeploymentsLifecycleStateCreating       ListDeploymentsLifecycleStateEnum = "CREATING"
	ListDeploymentsLifecycleStateUpdating       ListDeploymentsLifecycleStateEnum = "UPDATING"
	ListDeploymentsLifecycleStateActive         ListDeploymentsLifecycleStateEnum = "ACTIVE"
	ListDeploymentsLifecycleStateInactive       ListDeploymentsLifecycleStateEnum = "INACTIVE"
	ListDeploymentsLifecycleStateDeleting       ListDeploymentsLifecycleStateEnum = "DELETING"
	ListDeploymentsLifecycleStateDeleted        ListDeploymentsLifecycleStateEnum = "DELETED"
	ListDeploymentsLifecycleStateFailed         ListDeploymentsLifecycleStateEnum = "FAILED"
	ListDeploymentsLifecycleStateNeedsAttention ListDeploymentsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDeploymentsLifecycleStateInProgress     ListDeploymentsLifecycleStateEnum = "IN_PROGRESS"
	ListDeploymentsLifecycleStateCanceling      ListDeploymentsLifecycleStateEnum = "CANCELING"
	ListDeploymentsLifecycleStateCanceled       ListDeploymentsLifecycleStateEnum = "CANCELED"
	ListDeploymentsLifecycleStateSucceeded      ListDeploymentsLifecycleStateEnum = "SUCCEEDED"
	ListDeploymentsLifecycleStateWaiting        ListDeploymentsLifecycleStateEnum = "WAITING"
)

var mappingListDeploymentsLifecycleStateEnum = map[string]ListDeploymentsLifecycleStateEnum{
	"CREATING":        ListDeploymentsLifecycleStateCreating,
	"UPDATING":        ListDeploymentsLifecycleStateUpdating,
	"ACTIVE":          ListDeploymentsLifecycleStateActive,
	"INACTIVE":        ListDeploymentsLifecycleStateInactive,
	"DELETING":        ListDeploymentsLifecycleStateDeleting,
	"DELETED":         ListDeploymentsLifecycleStateDeleted,
	"FAILED":          ListDeploymentsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDeploymentsLifecycleStateNeedsAttention,
	"IN_PROGRESS":     ListDeploymentsLifecycleStateInProgress,
	"CANCELING":       ListDeploymentsLifecycleStateCanceling,
	"CANCELED":        ListDeploymentsLifecycleStateCanceled,
	"SUCCEEDED":       ListDeploymentsLifecycleStateSucceeded,
	"WAITING":         ListDeploymentsLifecycleStateWaiting,
}

var mappingListDeploymentsLifecycleStateEnumLowerCase = map[string]ListDeploymentsLifecycleStateEnum{
	"creating":        ListDeploymentsLifecycleStateCreating,
	"updating":        ListDeploymentsLifecycleStateUpdating,
	"active":          ListDeploymentsLifecycleStateActive,
	"inactive":        ListDeploymentsLifecycleStateInactive,
	"deleting":        ListDeploymentsLifecycleStateDeleting,
	"deleted":         ListDeploymentsLifecycleStateDeleted,
	"failed":          ListDeploymentsLifecycleStateFailed,
	"needs_attention": ListDeploymentsLifecycleStateNeedsAttention,
	"in_progress":     ListDeploymentsLifecycleStateInProgress,
	"canceling":       ListDeploymentsLifecycleStateCanceling,
	"canceled":        ListDeploymentsLifecycleStateCanceled,
	"succeeded":       ListDeploymentsLifecycleStateSucceeded,
	"waiting":         ListDeploymentsLifecycleStateWaiting,
}

// GetListDeploymentsLifecycleStateEnumValues Enumerates the set of values for ListDeploymentsLifecycleStateEnum
func GetListDeploymentsLifecycleStateEnumValues() []ListDeploymentsLifecycleStateEnum {
	values := make([]ListDeploymentsLifecycleStateEnum, 0)
	for _, v := range mappingListDeploymentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDeploymentsLifecycleStateEnum
func GetListDeploymentsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
		"IN_PROGRESS",
		"CANCELING",
		"CANCELED",
		"SUCCEEDED",
		"WAITING",
	}
}

// GetMappingListDeploymentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsLifecycleStateEnum(val string) (ListDeploymentsLifecycleStateEnum, bool) {
	enum, ok := mappingListDeploymentsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentsLifecycleSubStateEnum Enum with underlying type: string
type ListDeploymentsLifecycleSubStateEnum string

// Set of constants representing the allowable values for ListDeploymentsLifecycleSubStateEnum
const (
	ListDeploymentsLifecycleSubStateRecovering         ListDeploymentsLifecycleSubStateEnum = "RECOVERING"
	ListDeploymentsLifecycleSubStateStarting           ListDeploymentsLifecycleSubStateEnum = "STARTING"
	ListDeploymentsLifecycleSubStateStopping           ListDeploymentsLifecycleSubStateEnum = "STOPPING"
	ListDeploymentsLifecycleSubStateMoving             ListDeploymentsLifecycleSubStateEnum = "MOVING"
	ListDeploymentsLifecycleSubStateUpgrading          ListDeploymentsLifecycleSubStateEnum = "UPGRADING"
	ListDeploymentsLifecycleSubStateRestoring          ListDeploymentsLifecycleSubStateEnum = "RESTORING"
	ListDeploymentsLifecycleSubStateBackupInProgress   ListDeploymentsLifecycleSubStateEnum = "BACKUP_IN_PROGRESS"
	ListDeploymentsLifecycleSubStateRollbackInProgress ListDeploymentsLifecycleSubStateEnum = "ROLLBACK_IN_PROGRESS"
)

var mappingListDeploymentsLifecycleSubStateEnum = map[string]ListDeploymentsLifecycleSubStateEnum{
	"RECOVERING":           ListDeploymentsLifecycleSubStateRecovering,
	"STARTING":             ListDeploymentsLifecycleSubStateStarting,
	"STOPPING":             ListDeploymentsLifecycleSubStateStopping,
	"MOVING":               ListDeploymentsLifecycleSubStateMoving,
	"UPGRADING":            ListDeploymentsLifecycleSubStateUpgrading,
	"RESTORING":            ListDeploymentsLifecycleSubStateRestoring,
	"BACKUP_IN_PROGRESS":   ListDeploymentsLifecycleSubStateBackupInProgress,
	"ROLLBACK_IN_PROGRESS": ListDeploymentsLifecycleSubStateRollbackInProgress,
}

var mappingListDeploymentsLifecycleSubStateEnumLowerCase = map[string]ListDeploymentsLifecycleSubStateEnum{
	"recovering":           ListDeploymentsLifecycleSubStateRecovering,
	"starting":             ListDeploymentsLifecycleSubStateStarting,
	"stopping":             ListDeploymentsLifecycleSubStateStopping,
	"moving":               ListDeploymentsLifecycleSubStateMoving,
	"upgrading":            ListDeploymentsLifecycleSubStateUpgrading,
	"restoring":            ListDeploymentsLifecycleSubStateRestoring,
	"backup_in_progress":   ListDeploymentsLifecycleSubStateBackupInProgress,
	"rollback_in_progress": ListDeploymentsLifecycleSubStateRollbackInProgress,
}

// GetListDeploymentsLifecycleSubStateEnumValues Enumerates the set of values for ListDeploymentsLifecycleSubStateEnum
func GetListDeploymentsLifecycleSubStateEnumValues() []ListDeploymentsLifecycleSubStateEnum {
	values := make([]ListDeploymentsLifecycleSubStateEnum, 0)
	for _, v := range mappingListDeploymentsLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsLifecycleSubStateEnumStringValues Enumerates the set of values in String for ListDeploymentsLifecycleSubStateEnum
func GetListDeploymentsLifecycleSubStateEnumStringValues() []string {
	return []string{
		"RECOVERING",
		"STARTING",
		"STOPPING",
		"MOVING",
		"UPGRADING",
		"RESTORING",
		"BACKUP_IN_PROGRESS",
		"ROLLBACK_IN_PROGRESS",
	}
}

// GetMappingListDeploymentsLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsLifecycleSubStateEnum(val string) (ListDeploymentsLifecycleSubStateEnum, bool) {
	enum, ok := mappingListDeploymentsLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentsSortOrderEnum Enum with underlying type: string
type ListDeploymentsSortOrderEnum string

// Set of constants representing the allowable values for ListDeploymentsSortOrderEnum
const (
	ListDeploymentsSortOrderAsc  ListDeploymentsSortOrderEnum = "ASC"
	ListDeploymentsSortOrderDesc ListDeploymentsSortOrderEnum = "DESC"
)

var mappingListDeploymentsSortOrderEnum = map[string]ListDeploymentsSortOrderEnum{
	"ASC":  ListDeploymentsSortOrderAsc,
	"DESC": ListDeploymentsSortOrderDesc,
}

var mappingListDeploymentsSortOrderEnumLowerCase = map[string]ListDeploymentsSortOrderEnum{
	"asc":  ListDeploymentsSortOrderAsc,
	"desc": ListDeploymentsSortOrderDesc,
}

// GetListDeploymentsSortOrderEnumValues Enumerates the set of values for ListDeploymentsSortOrderEnum
func GetListDeploymentsSortOrderEnumValues() []ListDeploymentsSortOrderEnum {
	values := make([]ListDeploymentsSortOrderEnum, 0)
	for _, v := range mappingListDeploymentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsSortOrderEnumStringValues Enumerates the set of values in String for ListDeploymentsSortOrderEnum
func GetListDeploymentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDeploymentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsSortOrderEnum(val string) (ListDeploymentsSortOrderEnum, bool) {
	enum, ok := mappingListDeploymentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDeploymentsSortByEnum Enum with underlying type: string
type ListDeploymentsSortByEnum string

// Set of constants representing the allowable values for ListDeploymentsSortByEnum
const (
	ListDeploymentsSortByTimecreated ListDeploymentsSortByEnum = "timeCreated"
	ListDeploymentsSortByDisplayname ListDeploymentsSortByEnum = "displayName"
)

var mappingListDeploymentsSortByEnum = map[string]ListDeploymentsSortByEnum{
	"timeCreated": ListDeploymentsSortByTimecreated,
	"displayName": ListDeploymentsSortByDisplayname,
}

var mappingListDeploymentsSortByEnumLowerCase = map[string]ListDeploymentsSortByEnum{
	"timecreated": ListDeploymentsSortByTimecreated,
	"displayname": ListDeploymentsSortByDisplayname,
}

// GetListDeploymentsSortByEnumValues Enumerates the set of values for ListDeploymentsSortByEnum
func GetListDeploymentsSortByEnumValues() []ListDeploymentsSortByEnum {
	values := make([]ListDeploymentsSortByEnum, 0)
	for _, v := range mappingListDeploymentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDeploymentsSortByEnumStringValues Enumerates the set of values in String for ListDeploymentsSortByEnum
func GetListDeploymentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDeploymentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDeploymentsSortByEnum(val string) (ListDeploymentsSortByEnum, bool) {
	enum, ok := mappingListDeploymentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
