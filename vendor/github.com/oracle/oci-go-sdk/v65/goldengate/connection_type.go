// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// ConnectionTypeEnum Enum with underlying type: string
type ConnectionTypeEnum string

// Set of constants representing the allowable values for ConnectionTypeEnum
const (
	ConnectionTypeGoldengate            ConnectionTypeEnum = "GOLDENGATE"
	ConnectionTypeKafka                 ConnectionTypeEnum = "KAFKA"
	ConnectionTypeKafkaSchemaRegistry   ConnectionTypeEnum = "KAFKA_SCHEMA_REGISTRY"
	ConnectionTypeMysql                 ConnectionTypeEnum = "MYSQL"
	ConnectionTypeJavaMessageService    ConnectionTypeEnum = "JAVA_MESSAGE_SERVICE"
	ConnectionTypeMicrosoftSqlserver    ConnectionTypeEnum = "MICROSOFT_SQLSERVER"
	ConnectionTypeOciObjectStorage      ConnectionTypeEnum = "OCI_OBJECT_STORAGE"
	ConnectionTypeOracle                ConnectionTypeEnum = "ORACLE"
	ConnectionTypeAzureDataLakeStorage  ConnectionTypeEnum = "AZURE_DATA_LAKE_STORAGE"
	ConnectionTypePostgresql            ConnectionTypeEnum = "POSTGRESQL"
	ConnectionTypeAzureSynapseAnalytics ConnectionTypeEnum = "AZURE_SYNAPSE_ANALYTICS"
	ConnectionTypeSnowflake             ConnectionTypeEnum = "SNOWFLAKE"
	ConnectionTypeAmazonS3              ConnectionTypeEnum = "AMAZON_S3"
	ConnectionTypeHdfs                  ConnectionTypeEnum = "HDFS"
	ConnectionTypeOracleNosql           ConnectionTypeEnum = "ORACLE_NOSQL"
	ConnectionTypeMongodb               ConnectionTypeEnum = "MONGODB"
	ConnectionTypeAmazonKinesis         ConnectionTypeEnum = "AMAZON_KINESIS"
	ConnectionTypeAmazonRedshift        ConnectionTypeEnum = "AMAZON_REDSHIFT"
	ConnectionTypeRedis                 ConnectionTypeEnum = "REDIS"
	ConnectionTypeElasticsearch         ConnectionTypeEnum = "ELASTICSEARCH"
	ConnectionTypeGeneric               ConnectionTypeEnum = "GENERIC"
	ConnectionTypeGoogleCloudStorage    ConnectionTypeEnum = "GOOGLE_CLOUD_STORAGE"
	ConnectionTypeGoogleBigquery        ConnectionTypeEnum = "GOOGLE_BIGQUERY"
)

var mappingConnectionTypeEnum = map[string]ConnectionTypeEnum{
	"GOLDENGATE":              ConnectionTypeGoldengate,
	"KAFKA":                   ConnectionTypeKafka,
	"KAFKA_SCHEMA_REGISTRY":   ConnectionTypeKafkaSchemaRegistry,
	"MYSQL":                   ConnectionTypeMysql,
	"JAVA_MESSAGE_SERVICE":    ConnectionTypeJavaMessageService,
	"MICROSOFT_SQLSERVER":     ConnectionTypeMicrosoftSqlserver,
	"OCI_OBJECT_STORAGE":      ConnectionTypeOciObjectStorage,
	"ORACLE":                  ConnectionTypeOracle,
	"AZURE_DATA_LAKE_STORAGE": ConnectionTypeAzureDataLakeStorage,
	"POSTGRESQL":              ConnectionTypePostgresql,
	"AZURE_SYNAPSE_ANALYTICS": ConnectionTypeAzureSynapseAnalytics,
	"SNOWFLAKE":               ConnectionTypeSnowflake,
	"AMAZON_S3":               ConnectionTypeAmazonS3,
	"HDFS":                    ConnectionTypeHdfs,
	"ORACLE_NOSQL":            ConnectionTypeOracleNosql,
	"MONGODB":                 ConnectionTypeMongodb,
	"AMAZON_KINESIS":          ConnectionTypeAmazonKinesis,
	"AMAZON_REDSHIFT":         ConnectionTypeAmazonRedshift,
	"REDIS":                   ConnectionTypeRedis,
	"ELASTICSEARCH":           ConnectionTypeElasticsearch,
	"GENERIC":                 ConnectionTypeGeneric,
	"GOOGLE_CLOUD_STORAGE":    ConnectionTypeGoogleCloudStorage,
	"GOOGLE_BIGQUERY":         ConnectionTypeGoogleBigquery,
}

var mappingConnectionTypeEnumLowerCase = map[string]ConnectionTypeEnum{
	"goldengate":              ConnectionTypeGoldengate,
	"kafka":                   ConnectionTypeKafka,
	"kafka_schema_registry":   ConnectionTypeKafkaSchemaRegistry,
	"mysql":                   ConnectionTypeMysql,
	"java_message_service":    ConnectionTypeJavaMessageService,
	"microsoft_sqlserver":     ConnectionTypeMicrosoftSqlserver,
	"oci_object_storage":      ConnectionTypeOciObjectStorage,
	"oracle":                  ConnectionTypeOracle,
	"azure_data_lake_storage": ConnectionTypeAzureDataLakeStorage,
	"postgresql":              ConnectionTypePostgresql,
	"azure_synapse_analytics": ConnectionTypeAzureSynapseAnalytics,
	"snowflake":               ConnectionTypeSnowflake,
	"amazon_s3":               ConnectionTypeAmazonS3,
	"hdfs":                    ConnectionTypeHdfs,
	"oracle_nosql":            ConnectionTypeOracleNosql,
	"mongodb":                 ConnectionTypeMongodb,
	"amazon_kinesis":          ConnectionTypeAmazonKinesis,
	"amazon_redshift":         ConnectionTypeAmazonRedshift,
	"redis":                   ConnectionTypeRedis,
	"elasticsearch":           ConnectionTypeElasticsearch,
	"generic":                 ConnectionTypeGeneric,
	"google_cloud_storage":    ConnectionTypeGoogleCloudStorage,
	"google_bigquery":         ConnectionTypeGoogleBigquery,
}

// GetConnectionTypeEnumValues Enumerates the set of values for ConnectionTypeEnum
func GetConnectionTypeEnumValues() []ConnectionTypeEnum {
	values := make([]ConnectionTypeEnum, 0)
	for _, v := range mappingConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionTypeEnumStringValues Enumerates the set of values in String for ConnectionTypeEnum
func GetConnectionTypeEnumStringValues() []string {
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

// GetMappingConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionTypeEnum(val string) (ConnectionTypeEnum, bool) {
	enum, ok := mappingConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
