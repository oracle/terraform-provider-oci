// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ConfigurationVariables User controllable service variables.
type ConfigurationVariables struct {

	// ("completion_type")
	CompletionType ConfigurationVariablesCompletionTypeEnum `mandatory:"false" json:"completionType,omitempty"`

	// ("default_authentication_plugin")
	DefaultAuthenticationPlugin ConfigurationVariablesDefaultAuthenticationPluginEnum `mandatory:"false" json:"defaultAuthenticationPlugin,omitempty"`

	// ("transaction_isolation")
	TransactionIsolation ConfigurationVariablesTransactionIsolationEnum `mandatory:"false" json:"transactionIsolation,omitempty"`

	// ("innodb_ft_server_stopword_table")
	InnodbFtServerStopwordTable *string `mandatory:"false" json:"innodbFtServerStopwordTable"`

	// ("mandatory_roles")
	MandatoryRoles *string `mandatory:"false" json:"mandatoryRoles"`

	// ("autocommit")
	Autocommit *bool `mandatory:"false" json:"autocommit"`

	// ("foreign_key_checks")
	ForeignKeyChecks *bool `mandatory:"false" json:"foreignKeyChecks"`

	// ("innodb_ft_enable_stopword")
	InnodbFtEnableStopword *bool `mandatory:"false" json:"innodbFtEnableStopword"`

	// ("local_infile")
	LocalInfile *bool `mandatory:"false" json:"localInfile"`

	// ("mysql_firewall_mode")
	MysqlFirewallMode *bool `mandatory:"false" json:"mysqlFirewallMode"`

	// ("mysqlx_enable_hello_notice")
	MysqlxEnableHelloNotice *bool `mandatory:"false" json:"mysqlxEnableHelloNotice"`

	// ("sql_require_primary_key")
	SqlRequirePrimaryKey *bool `mandatory:"false" json:"sqlRequirePrimaryKey"`

	// ("sql_warnings")
	SqlWarnings *bool `mandatory:"false" json:"sqlWarnings"`

	// ("binlog_expire_logs_seconds")
	BinlogExpireLogsSeconds *int `mandatory:"false" json:"binlogExpireLogsSeconds"`

	// ("innodb_buffer_pool_size")
	InnodbBufferPoolSize *int64 `mandatory:"false" json:"innodbBufferPoolSize"`

	// ("innodb_ft_result_cache_limit")
	InnodbFtResultCacheLimit *int `mandatory:"false" json:"innodbFtResultCacheLimit"`

	// ("max_connections")
	MaxConnections *int `mandatory:"false" json:"maxConnections"`

	// ("max_prepared_stmt_count")
	MaxPreparedStmtCount *int `mandatory:"false" json:"maxPreparedStmtCount"`

	// ("connect_timeout")
	ConnectTimeout *int `mandatory:"false" json:"connectTimeout"`

	// ("cte_max_recursion_depth")
	CteMaxRecursionDepth *int `mandatory:"false" json:"cteMaxRecursionDepth"`

	// ("generated_random_password_length")
	GeneratedRandomPasswordLength *int `mandatory:"false" json:"generatedRandomPasswordLength"`

	// ("information_schema_stats_expiry")
	InformationSchemaStatsExpiry *int `mandatory:"false" json:"informationSchemaStatsExpiry"`

	// ("innodb_buffer_pool_instances")
	InnodbBufferPoolInstances *int `mandatory:"false" json:"innodbBufferPoolInstances"`

	// ("innodb_ft_max_token_size")
	InnodbFtMaxTokenSize *int `mandatory:"false" json:"innodbFtMaxTokenSize"`

	// ("innodb_ft_min_token_size")
	InnodbFtMinTokenSize *int `mandatory:"false" json:"innodbFtMinTokenSize"`

	// ("innodb_ft_num_word_optimize")
	InnodbFtNumWordOptimize *int `mandatory:"false" json:"innodbFtNumWordOptimize"`

	// ("innodb_lock_wait_timeout")
	InnodbLockWaitTimeout *int `mandatory:"false" json:"innodbLockWaitTimeout"`

	// ("innodb_max_purge_lag")
	InnodbMaxPurgeLag *int `mandatory:"false" json:"innodbMaxPurgeLag"`

	// ("innodb_max_purge_lag_delay")
	InnodbMaxPurgeLagDelay *int `mandatory:"false" json:"innodbMaxPurgeLagDelay"`

	// ("max_execution_time")
	MaxExecutionTime *int `mandatory:"false" json:"maxExecutionTime"`

	// ("mysqlx_connect_timeout")
	MysqlxConnectTimeout *int `mandatory:"false" json:"mysqlxConnectTimeout"`

	// ("mysqlx_document_id_unique_prefix")
	MysqlxDocumentIdUniquePrefix *int `mandatory:"false" json:"mysqlxDocumentIdUniquePrefix"`

	// ("mysqlx_idle_worker_thread_timeout")
	MysqlxIdleWorkerThreadTimeout *int `mandatory:"false" json:"mysqlxIdleWorkerThreadTimeout"`

	// ("mysqlx_interactive_timeout")
	MysqlxInteractiveTimeout *int `mandatory:"false" json:"mysqlxInteractiveTimeout"`

	// ("mysqlx_max_allowed_packet")
	MysqlxMaxAllowedPacket *int `mandatory:"false" json:"mysqlxMaxAllowedPacket"`

	// ("mysqlx_min_worker_threads")
	MysqlxMinWorkerThreads *int `mandatory:"false" json:"mysqlxMinWorkerThreads"`

	// ("mysqlx_read_timeout")
	MysqlxReadTimeout *int `mandatory:"false" json:"mysqlxReadTimeout"`

	// ("mysqlx_wait_timeout")
	MysqlxWaitTimeout *int `mandatory:"false" json:"mysqlxWaitTimeout"`

	// ("mysqlx_write_timeout")
	MysqlxWriteTimeout *int `mandatory:"false" json:"mysqlxWriteTimeout"`

	// ("parser_max_mem_size")
	ParserMaxMemSize *int `mandatory:"false" json:"parserMaxMemSize"`

	// ("query_alloc_block_size")
	QueryAllocBlockSize *int `mandatory:"false" json:"queryAllocBlockSize"`

	// ("query_prealloc_size")
	QueryPreallocSize *int `mandatory:"false" json:"queryPreallocSize"`

	// ("sql_mode")
	SqlMode *string `mandatory:"false" json:"sqlMode"`

	// Set the default compression level for the deflate algorithm. ("mysqlx_deflate_default_compression_level")
	MysqlxDeflateDefaultCompressionLevel *int `mandatory:"false" json:"mysqlxDeflateDefaultCompressionLevel"`

	// Limit the upper bound of accepted compression levels for the deflate algorithm. ("mysqlx_deflate_max_client_compression_level")
	MysqlxDeflateMaxClientCompressionLevel *int `mandatory:"false" json:"mysqlxDeflateMaxClientCompressionLevel"`

	// Limit the upper bound of accepted compression levels for the lz4 algorithm. ("mysqlx_lz4_max_client_compression_level")
	MysqlxLz4MaxClientCompressionLevel *int `mandatory:"false" json:"mysqlxLz4MaxClientCompressionLevel"`

	// Set the default compression level for the lz4 algorithm. ("mysqlx_lz4_default_compression_level")
	MysqlxLz4DefaultCompressionLevel *int `mandatory:"false" json:"mysqlxLz4DefaultCompressionLevel"`

	// Limit the upper bound of accepted compression levels for the zstd algorithm. ("mysqlx_zstd_max_client_compression_level")
	MysqlxZstdMaxClientCompressionLevel *int `mandatory:"false" json:"mysqlxZstdMaxClientCompressionLevel"`

	// Set the default compression level for the zstd algorithm. ("mysqlx_zstd_default_compression_level")
	MysqlZstdDefaultCompressionLevel *int `mandatory:"false" json:"mysqlZstdDefaultCompressionLevel"`
}

func (m ConfigurationVariables) String() string {
	return common.PointerString(m)
}

// ConfigurationVariablesCompletionTypeEnum Enum with underlying type: string
type ConfigurationVariablesCompletionTypeEnum string

// Set of constants representing the allowable values for ConfigurationVariablesCompletionTypeEnum
const (
	ConfigurationVariablesCompletionTypeNoChain ConfigurationVariablesCompletionTypeEnum = "NO_CHAIN"
	ConfigurationVariablesCompletionTypeChain   ConfigurationVariablesCompletionTypeEnum = "CHAIN"
	ConfigurationVariablesCompletionTypeRelease ConfigurationVariablesCompletionTypeEnum = "RELEASE"
)

var mappingConfigurationVariablesCompletionType = map[string]ConfigurationVariablesCompletionTypeEnum{
	"NO_CHAIN": ConfigurationVariablesCompletionTypeNoChain,
	"CHAIN":    ConfigurationVariablesCompletionTypeChain,
	"RELEASE":  ConfigurationVariablesCompletionTypeRelease,
}

// GetConfigurationVariablesCompletionTypeEnumValues Enumerates the set of values for ConfigurationVariablesCompletionTypeEnum
func GetConfigurationVariablesCompletionTypeEnumValues() []ConfigurationVariablesCompletionTypeEnum {
	values := make([]ConfigurationVariablesCompletionTypeEnum, 0)
	for _, v := range mappingConfigurationVariablesCompletionType {
		values = append(values, v)
	}
	return values
}

// ConfigurationVariablesDefaultAuthenticationPluginEnum Enum with underlying type: string
type ConfigurationVariablesDefaultAuthenticationPluginEnum string

// Set of constants representing the allowable values for ConfigurationVariablesDefaultAuthenticationPluginEnum
const (
	ConfigurationVariablesDefaultAuthenticationPluginMysqlNativePassword ConfigurationVariablesDefaultAuthenticationPluginEnum = "mysql_native_password"
	ConfigurationVariablesDefaultAuthenticationPluginSha256Password      ConfigurationVariablesDefaultAuthenticationPluginEnum = "sha256_password"
	ConfigurationVariablesDefaultAuthenticationPluginCachingSha2Password ConfigurationVariablesDefaultAuthenticationPluginEnum = "caching_sha2_password"
)

var mappingConfigurationVariablesDefaultAuthenticationPlugin = map[string]ConfigurationVariablesDefaultAuthenticationPluginEnum{
	"mysql_native_password": ConfigurationVariablesDefaultAuthenticationPluginMysqlNativePassword,
	"sha256_password":       ConfigurationVariablesDefaultAuthenticationPluginSha256Password,
	"caching_sha2_password": ConfigurationVariablesDefaultAuthenticationPluginCachingSha2Password,
}

// GetConfigurationVariablesDefaultAuthenticationPluginEnumValues Enumerates the set of values for ConfigurationVariablesDefaultAuthenticationPluginEnum
func GetConfigurationVariablesDefaultAuthenticationPluginEnumValues() []ConfigurationVariablesDefaultAuthenticationPluginEnum {
	values := make([]ConfigurationVariablesDefaultAuthenticationPluginEnum, 0)
	for _, v := range mappingConfigurationVariablesDefaultAuthenticationPlugin {
		values = append(values, v)
	}
	return values
}

// ConfigurationVariablesTransactionIsolationEnum Enum with underlying type: string
type ConfigurationVariablesTransactionIsolationEnum string

// Set of constants representing the allowable values for ConfigurationVariablesTransactionIsolationEnum
const (
	ConfigurationVariablesTransactionIsolationReadUncommitted ConfigurationVariablesTransactionIsolationEnum = "READ-UNCOMMITTED"
	ConfigurationVariablesTransactionIsolationReadCommited    ConfigurationVariablesTransactionIsolationEnum = "READ-COMMITED"
	ConfigurationVariablesTransactionIsolationRepeatableRead  ConfigurationVariablesTransactionIsolationEnum = "REPEATABLE-READ"
	ConfigurationVariablesTransactionIsolationSerializable    ConfigurationVariablesTransactionIsolationEnum = "SERIALIZABLE"
)

var mappingConfigurationVariablesTransactionIsolation = map[string]ConfigurationVariablesTransactionIsolationEnum{
	"READ-UNCOMMITTED": ConfigurationVariablesTransactionIsolationReadUncommitted,
	"READ-COMMITED":    ConfigurationVariablesTransactionIsolationReadCommited,
	"REPEATABLE-READ":  ConfigurationVariablesTransactionIsolationRepeatableRead,
	"SERIALIZABLE":     ConfigurationVariablesTransactionIsolationSerializable,
}

// GetConfigurationVariablesTransactionIsolationEnumValues Enumerates the set of values for ConfigurationVariablesTransactionIsolationEnum
func GetConfigurationVariablesTransactionIsolationEnumValues() []ConfigurationVariablesTransactionIsolationEnum {
	values := make([]ConfigurationVariablesTransactionIsolationEnum, 0)
	for _, v := range mappingConfigurationVariablesTransactionIsolation {
		values = append(values, v)
	}
	return values
}
