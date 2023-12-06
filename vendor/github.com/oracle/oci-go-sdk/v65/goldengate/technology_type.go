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

// TechnologyTypeEnum Enum with underlying type: string
type TechnologyTypeEnum string

// Set of constants representing the allowable values for TechnologyTypeEnum
const (
	TechnologyTypeGoldengate                       TechnologyTypeEnum = "GOLDENGATE"
	TechnologyTypeGeneric                          TechnologyTypeEnum = "GENERIC"
	TechnologyTypeOciAutonomousDatabase            TechnologyTypeEnum = "OCI_AUTONOMOUS_DATABASE"
	TechnologyTypeOciAutonomousJsonDatabase        TechnologyTypeEnum = "OCI_AUTONOMOUS_JSON_DATABASE"
	TechnologyTypeOciMysql                         TechnologyTypeEnum = "OCI_MYSQL"
	TechnologyTypeOciObjectStorage                 TechnologyTypeEnum = "OCI_OBJECT_STORAGE"
	TechnologyTypeOciStreaming                     TechnologyTypeEnum = "OCI_STREAMING"
	TechnologyTypeOracleDatabase                   TechnologyTypeEnum = "ORACLE_DATABASE"
	TechnologyTypeOracleExadata                    TechnologyTypeEnum = "ORACLE_EXADATA"
	TechnologyTypeOracleNosql                      TechnologyTypeEnum = "ORACLE_NOSQL"
	TechnologyTypeOracleWeblogicJms                TechnologyTypeEnum = "ORACLE_WEBLOGIC_JMS"
	TechnologyTypeAmazonRdsOracle                  TechnologyTypeEnum = "AMAZON_RDS_ORACLE"
	TechnologyTypeAmazonRdsSqlserver               TechnologyTypeEnum = "AMAZON_RDS_SQLSERVER"
	TechnologyTypeAmazonS3                         TechnologyTypeEnum = "AMAZON_S3"
	TechnologyTypeAmazonAuroraMysql                TechnologyTypeEnum = "AMAZON_AURORA_MYSQL"
	TechnologyTypeAmazonAuroraPostgresql           TechnologyTypeEnum = "AMAZON_AURORA_POSTGRESQL"
	TechnologyTypeAmazonKinesis                    TechnologyTypeEnum = "AMAZON_KINESIS"
	TechnologyTypeAmazonRedshift                   TechnologyTypeEnum = "AMAZON_REDSHIFT"
	TechnologyTypeAmazonRdsMariadb                 TechnologyTypeEnum = "AMAZON_RDS_MARIADB"
	TechnologyTypeAmazonRdsMysql                   TechnologyTypeEnum = "AMAZON_RDS_MYSQL"
	TechnologyTypeAmazonRdsPostgresql              TechnologyTypeEnum = "AMAZON_RDS_POSTGRESQL"
	TechnologyTypeApacheKafka                      TechnologyTypeEnum = "APACHE_KAFKA"
	TechnologyTypeAzureCosmosDbForMongodb          TechnologyTypeEnum = "AZURE_COSMOS_DB_FOR_MONGODB"
	TechnologyTypeAzureDataLakeStorage             TechnologyTypeEnum = "AZURE_DATA_LAKE_STORAGE"
	TechnologyTypeAzureEventHubs                   TechnologyTypeEnum = "AZURE_EVENT_HUBS"
	TechnologyTypeAzureMysql                       TechnologyTypeEnum = "AZURE_MYSQL"
	TechnologyTypeAzurePostgresql                  TechnologyTypeEnum = "AZURE_POSTGRESQL"
	TechnologyTypeAzureSqlserverManagedInstance    TechnologyTypeEnum = "AZURE_SQLSERVER_MANAGED_INSTANCE"
	TechnologyTypeAzureSqlserverNonManagedInstance TechnologyTypeEnum = "AZURE_SQLSERVER_NON_MANAGED_INSTANCE"
	TechnologyTypeAzureSynapseAnalytics            TechnologyTypeEnum = "AZURE_SYNAPSE_ANALYTICS"
	TechnologyTypeConfluentKafka                   TechnologyTypeEnum = "CONFLUENT_KAFKA"
	TechnologyTypeConfluentSchemaRegistry          TechnologyTypeEnum = "CONFLUENT_SCHEMA_REGISTRY"
	TechnologyTypeElasticsearch                    TechnologyTypeEnum = "ELASTICSEARCH"
	TechnologyTypeGoogleBigquery                   TechnologyTypeEnum = "GOOGLE_BIGQUERY"
	TechnologyTypeGoogleCloudStorage               TechnologyTypeEnum = "GOOGLE_CLOUD_STORAGE"
	TechnologyTypeGoogleCloudSqlMysql              TechnologyTypeEnum = "GOOGLE_CLOUD_SQL_MYSQL"
	TechnologyTypeGoogleCloudSqlPostgresql         TechnologyTypeEnum = "GOOGLE_CLOUD_SQL_POSTGRESQL"
	TechnologyTypeGoogleCloudSqlSqlserver          TechnologyTypeEnum = "GOOGLE_CLOUD_SQL_SQLSERVER"
	TechnologyTypeHdfs                             TechnologyTypeEnum = "HDFS"
	TechnologyTypeMariadb                          TechnologyTypeEnum = "MARIADB"
	TechnologyTypeMicrosoftSqlserver               TechnologyTypeEnum = "MICROSOFT_SQLSERVER"
	TechnologyTypeMongodb                          TechnologyTypeEnum = "MONGODB"
	TechnologyTypeMysqlServer                      TechnologyTypeEnum = "MYSQL_SERVER"
	TechnologyTypePostgresqlServer                 TechnologyTypeEnum = "POSTGRESQL_SERVER"
	TechnologyTypeRedis                            TechnologyTypeEnum = "REDIS"
	TechnologyTypeSinglestoredb                    TechnologyTypeEnum = "SINGLESTOREDB"
	TechnologyTypeSinglestoredbCloud               TechnologyTypeEnum = "SINGLESTOREDB_CLOUD"
	TechnologyTypeSnowflake                        TechnologyTypeEnum = "SNOWFLAKE"
)

var mappingTechnologyTypeEnum = map[string]TechnologyTypeEnum{
	"GOLDENGATE":                           TechnologyTypeGoldengate,
	"GENERIC":                              TechnologyTypeGeneric,
	"OCI_AUTONOMOUS_DATABASE":              TechnologyTypeOciAutonomousDatabase,
	"OCI_AUTONOMOUS_JSON_DATABASE":         TechnologyTypeOciAutonomousJsonDatabase,
	"OCI_MYSQL":                            TechnologyTypeOciMysql,
	"OCI_OBJECT_STORAGE":                   TechnologyTypeOciObjectStorage,
	"OCI_STREAMING":                        TechnologyTypeOciStreaming,
	"ORACLE_DATABASE":                      TechnologyTypeOracleDatabase,
	"ORACLE_EXADATA":                       TechnologyTypeOracleExadata,
	"ORACLE_NOSQL":                         TechnologyTypeOracleNosql,
	"ORACLE_WEBLOGIC_JMS":                  TechnologyTypeOracleWeblogicJms,
	"AMAZON_RDS_ORACLE":                    TechnologyTypeAmazonRdsOracle,
	"AMAZON_RDS_SQLSERVER":                 TechnologyTypeAmazonRdsSqlserver,
	"AMAZON_S3":                            TechnologyTypeAmazonS3,
	"AMAZON_AURORA_MYSQL":                  TechnologyTypeAmazonAuroraMysql,
	"AMAZON_AURORA_POSTGRESQL":             TechnologyTypeAmazonAuroraPostgresql,
	"AMAZON_KINESIS":                       TechnologyTypeAmazonKinesis,
	"AMAZON_REDSHIFT":                      TechnologyTypeAmazonRedshift,
	"AMAZON_RDS_MARIADB":                   TechnologyTypeAmazonRdsMariadb,
	"AMAZON_RDS_MYSQL":                     TechnologyTypeAmazonRdsMysql,
	"AMAZON_RDS_POSTGRESQL":                TechnologyTypeAmazonRdsPostgresql,
	"APACHE_KAFKA":                         TechnologyTypeApacheKafka,
	"AZURE_COSMOS_DB_FOR_MONGODB":          TechnologyTypeAzureCosmosDbForMongodb,
	"AZURE_DATA_LAKE_STORAGE":              TechnologyTypeAzureDataLakeStorage,
	"AZURE_EVENT_HUBS":                     TechnologyTypeAzureEventHubs,
	"AZURE_MYSQL":                          TechnologyTypeAzureMysql,
	"AZURE_POSTGRESQL":                     TechnologyTypeAzurePostgresql,
	"AZURE_SQLSERVER_MANAGED_INSTANCE":     TechnologyTypeAzureSqlserverManagedInstance,
	"AZURE_SQLSERVER_NON_MANAGED_INSTANCE": TechnologyTypeAzureSqlserverNonManagedInstance,
	"AZURE_SYNAPSE_ANALYTICS":              TechnologyTypeAzureSynapseAnalytics,
	"CONFLUENT_KAFKA":                      TechnologyTypeConfluentKafka,
	"CONFLUENT_SCHEMA_REGISTRY":            TechnologyTypeConfluentSchemaRegistry,
	"ELASTICSEARCH":                        TechnologyTypeElasticsearch,
	"GOOGLE_BIGQUERY":                      TechnologyTypeGoogleBigquery,
	"GOOGLE_CLOUD_STORAGE":                 TechnologyTypeGoogleCloudStorage,
	"GOOGLE_CLOUD_SQL_MYSQL":               TechnologyTypeGoogleCloudSqlMysql,
	"GOOGLE_CLOUD_SQL_POSTGRESQL":          TechnologyTypeGoogleCloudSqlPostgresql,
	"GOOGLE_CLOUD_SQL_SQLSERVER":           TechnologyTypeGoogleCloudSqlSqlserver,
	"HDFS":                                 TechnologyTypeHdfs,
	"MARIADB":                              TechnologyTypeMariadb,
	"MICROSOFT_SQLSERVER":                  TechnologyTypeMicrosoftSqlserver,
	"MONGODB":                              TechnologyTypeMongodb,
	"MYSQL_SERVER":                         TechnologyTypeMysqlServer,
	"POSTGRESQL_SERVER":                    TechnologyTypePostgresqlServer,
	"REDIS":                                TechnologyTypeRedis,
	"SINGLESTOREDB":                        TechnologyTypeSinglestoredb,
	"SINGLESTOREDB_CLOUD":                  TechnologyTypeSinglestoredbCloud,
	"SNOWFLAKE":                            TechnologyTypeSnowflake,
}

var mappingTechnologyTypeEnumLowerCase = map[string]TechnologyTypeEnum{
	"goldengate":                           TechnologyTypeGoldengate,
	"generic":                              TechnologyTypeGeneric,
	"oci_autonomous_database":              TechnologyTypeOciAutonomousDatabase,
	"oci_autonomous_json_database":         TechnologyTypeOciAutonomousJsonDatabase,
	"oci_mysql":                            TechnologyTypeOciMysql,
	"oci_object_storage":                   TechnologyTypeOciObjectStorage,
	"oci_streaming":                        TechnologyTypeOciStreaming,
	"oracle_database":                      TechnologyTypeOracleDatabase,
	"oracle_exadata":                       TechnologyTypeOracleExadata,
	"oracle_nosql":                         TechnologyTypeOracleNosql,
	"oracle_weblogic_jms":                  TechnologyTypeOracleWeblogicJms,
	"amazon_rds_oracle":                    TechnologyTypeAmazonRdsOracle,
	"amazon_rds_sqlserver":                 TechnologyTypeAmazonRdsSqlserver,
	"amazon_s3":                            TechnologyTypeAmazonS3,
	"amazon_aurora_mysql":                  TechnologyTypeAmazonAuroraMysql,
	"amazon_aurora_postgresql":             TechnologyTypeAmazonAuroraPostgresql,
	"amazon_kinesis":                       TechnologyTypeAmazonKinesis,
	"amazon_redshift":                      TechnologyTypeAmazonRedshift,
	"amazon_rds_mariadb":                   TechnologyTypeAmazonRdsMariadb,
	"amazon_rds_mysql":                     TechnologyTypeAmazonRdsMysql,
	"amazon_rds_postgresql":                TechnologyTypeAmazonRdsPostgresql,
	"apache_kafka":                         TechnologyTypeApacheKafka,
	"azure_cosmos_db_for_mongodb":          TechnologyTypeAzureCosmosDbForMongodb,
	"azure_data_lake_storage":              TechnologyTypeAzureDataLakeStorage,
	"azure_event_hubs":                     TechnologyTypeAzureEventHubs,
	"azure_mysql":                          TechnologyTypeAzureMysql,
	"azure_postgresql":                     TechnologyTypeAzurePostgresql,
	"azure_sqlserver_managed_instance":     TechnologyTypeAzureSqlserverManagedInstance,
	"azure_sqlserver_non_managed_instance": TechnologyTypeAzureSqlserverNonManagedInstance,
	"azure_synapse_analytics":              TechnologyTypeAzureSynapseAnalytics,
	"confluent_kafka":                      TechnologyTypeConfluentKafka,
	"confluent_schema_registry":            TechnologyTypeConfluentSchemaRegistry,
	"elasticsearch":                        TechnologyTypeElasticsearch,
	"google_bigquery":                      TechnologyTypeGoogleBigquery,
	"google_cloud_storage":                 TechnologyTypeGoogleCloudStorage,
	"google_cloud_sql_mysql":               TechnologyTypeGoogleCloudSqlMysql,
	"google_cloud_sql_postgresql":          TechnologyTypeGoogleCloudSqlPostgresql,
	"google_cloud_sql_sqlserver":           TechnologyTypeGoogleCloudSqlSqlserver,
	"hdfs":                                 TechnologyTypeHdfs,
	"mariadb":                              TechnologyTypeMariadb,
	"microsoft_sqlserver":                  TechnologyTypeMicrosoftSqlserver,
	"mongodb":                              TechnologyTypeMongodb,
	"mysql_server":                         TechnologyTypeMysqlServer,
	"postgresql_server":                    TechnologyTypePostgresqlServer,
	"redis":                                TechnologyTypeRedis,
	"singlestoredb":                        TechnologyTypeSinglestoredb,
	"singlestoredb_cloud":                  TechnologyTypeSinglestoredbCloud,
	"snowflake":                            TechnologyTypeSnowflake,
}

// GetTechnologyTypeEnumValues Enumerates the set of values for TechnologyTypeEnum
func GetTechnologyTypeEnumValues() []TechnologyTypeEnum {
	values := make([]TechnologyTypeEnum, 0)
	for _, v := range mappingTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTechnologyTypeEnumStringValues Enumerates the set of values in String for TechnologyTypeEnum
func GetTechnologyTypeEnumStringValues() []string {
	return []string{
		"GOLDENGATE",
		"GENERIC",
		"OCI_AUTONOMOUS_DATABASE",
		"OCI_AUTONOMOUS_JSON_DATABASE",
		"OCI_MYSQL",
		"OCI_OBJECT_STORAGE",
		"OCI_STREAMING",
		"ORACLE_DATABASE",
		"ORACLE_EXADATA",
		"ORACLE_NOSQL",
		"ORACLE_WEBLOGIC_JMS",
		"AMAZON_RDS_ORACLE",
		"AMAZON_RDS_SQLSERVER",
		"AMAZON_S3",
		"AMAZON_AURORA_MYSQL",
		"AMAZON_AURORA_POSTGRESQL",
		"AMAZON_KINESIS",
		"AMAZON_REDSHIFT",
		"AMAZON_RDS_MARIADB",
		"AMAZON_RDS_MYSQL",
		"AMAZON_RDS_POSTGRESQL",
		"APACHE_KAFKA",
		"AZURE_COSMOS_DB_FOR_MONGODB",
		"AZURE_DATA_LAKE_STORAGE",
		"AZURE_EVENT_HUBS",
		"AZURE_MYSQL",
		"AZURE_POSTGRESQL",
		"AZURE_SQLSERVER_MANAGED_INSTANCE",
		"AZURE_SQLSERVER_NON_MANAGED_INSTANCE",
		"AZURE_SYNAPSE_ANALYTICS",
		"CONFLUENT_KAFKA",
		"CONFLUENT_SCHEMA_REGISTRY",
		"ELASTICSEARCH",
		"GOOGLE_BIGQUERY",
		"GOOGLE_CLOUD_STORAGE",
		"GOOGLE_CLOUD_SQL_MYSQL",
		"GOOGLE_CLOUD_SQL_POSTGRESQL",
		"GOOGLE_CLOUD_SQL_SQLSERVER",
		"HDFS",
		"MARIADB",
		"MICROSOFT_SQLSERVER",
		"MONGODB",
		"MYSQL_SERVER",
		"POSTGRESQL_SERVER",
		"REDIS",
		"SINGLESTOREDB",
		"SINGLESTOREDB_CLOUD",
		"SNOWFLAKE",
	}
}

// GetMappingTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTechnologyTypeEnum(val string) (TechnologyTypeEnum, bool) {
	enum, ok := mappingTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
