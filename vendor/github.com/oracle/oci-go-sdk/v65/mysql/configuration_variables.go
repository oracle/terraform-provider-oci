// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// - EVENTUAL:
	//     Both RO and RW transactions do not wait for preceding transactions to be applied before executing.
	//     A RW transaction does not wait for other members to apply a transaction. This means that a transaction
	//     could be externalized on one member before the others. This also means that in the event of a primary failover,
	//     the new primary can accept new RO and RW transactions before the previous primary transactions are all applied.
	//     RO transactions could result in outdated values, RW transactions could result in a rollback due to conflicts.
	// - BEFORE_ON_PRIMARY_FAILOVER:
	//     New RO or RW transactions with a newly elected primary that is applying backlog from the old
	//     primary are held (not applied) until any backlog has been applied. This ensures that when a primary failover happens,
	//     intentionally or not, clients always see the latest value on the primary. This guarantees consistency, but means that
	//     clients must be able to handle the delay in the event that a backlog is being applied. Usually this delay should be minimal,
	//     but does depend on the size of the backlog.
	// - BEFORE:
	//     A RW transaction waits for all preceding transactions to complete before being applied. A RO transaction waits for all preceding
	//     transactions to complete before being executed. This ensures that this transaction reads the latest value by only affecting the
	//     latency of the transaction. This reduces the overhead of synchronization on every RW transaction, by ensuring synchronization is
	//     used only on RO transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
	// - AFTER:
	//     A RW transaction waits until its changes have been applied to all of the other members. This value has no effect on RO transactions.
	//     This mode ensures that when a transaction is committed on the local member, any subsequent transaction reads the written value or
	//     a more recent value on any group member. Use this mode with a group that is used for predominantly RO operations to ensure that
	//     applied RW transactions are applied everywhere once they commit. This could be used by your application to ensure that subsequent
	//     reads fetch the latest data which includes the latest writes. This reduces the overhead of synchronization on every RO transaction,
	//     by ensuring synchronization is used only on RW transactions. This consistency level also includes the consistency guarantees
	//     provided by BEFORE_ON_PRIMARY_FAILOVER.
	// - BEFORE_AND_AFTER:
	//     A RW transaction waits for 1) all preceding transactions to complete before being applied and 2) until its changes have been
	//     applied on other members. A RO transaction waits for all preceding transactions to complete before execution takes place.
	//     This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
	GroupReplicationConsistency ConfigurationVariablesGroupReplicationConsistencyEnum `mandatory:"false" json:"groupReplicationConsistency,omitempty"`

	// ("innodb_ft_enable_stopword")
	InnodbFtEnableStopword *bool `mandatory:"false" json:"innodbFtEnableStopword"`

	// ("local_infile")
	LocalInfile *bool `mandatory:"false" json:"localInfile"`

	// ("mysql_firewall_mode")
	MysqlFirewallMode *bool `mandatory:"false" json:"mysqlFirewallMode"`

	// ("mysqlx_enable_hello_notice") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxEnableHelloNotice *bool `mandatory:"false" json:"mysqlxEnableHelloNotice"`

	// ("sql_require_primary_key")
	SqlRequirePrimaryKey *bool `mandatory:"false" json:"sqlRequirePrimaryKey"`

	// ("sql_warnings")
	SqlWarnings *bool `mandatory:"false" json:"sqlWarnings"`

	// Sets the binary log expiration period in seconds.
	// binlogExpireLogsSeconds corresponds to the MySQL binary logging system variable binlog_expire_logs_seconds (https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_expire_logs_seconds).
	BinlogExpireLogsSeconds *int `mandatory:"false" json:"binlogExpireLogsSeconds"`

	// Configures the amount of table metadata added to the binary log when using row-based logging.
	// binlogRowMetadata corresponds to the MySQL binary logging system variable binlog_row_metadata (https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_metadata).
	BinlogRowMetadata ConfigurationVariablesBinlogRowMetadataEnum `mandatory:"false" json:"binlogRowMetadata,omitempty"`

	// When set to PARTIAL_JSON, this enables use of a space-efficient binary log format for updates that modify only a small portion of a JSON document.
	// binlogRowValueOptions corresponds to the MySQL binary logging system variable binlog_row_value_options (https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_value_options).
	BinlogRowValueOptions *string `mandatory:"false" json:"binlogRowValueOptions"`

	// Enables compression for transactions that are written to binary log files on this server.
	// binlogTransactionCompression corresponds to the MySQL binary logging system variable binlog_transaction_compression (https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_transaction_compression).
	BinlogTransactionCompression *bool `mandatory:"false" json:"binlogTransactionCompression"`

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

	// ("generated_random_password_length") DEPRECATED -- variable should not be settable and will be ignored
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

	// ("mysqlx_connect_timeout") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxConnectTimeout *int `mandatory:"false" json:"mysqlxConnectTimeout"`

	// ("mysqlx_document_id_unique_prefix") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxDocumentIdUniquePrefix *int `mandatory:"false" json:"mysqlxDocumentIdUniquePrefix"`

	// ("mysqlx_idle_worker_thread_timeout") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxIdleWorkerThreadTimeout *int `mandatory:"false" json:"mysqlxIdleWorkerThreadTimeout"`

	// ("mysqlx_interactive_timeout") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxInteractiveTimeout *int `mandatory:"false" json:"mysqlxInteractiveTimeout"`

	// ("mysqlx_max_allowed_packet") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxMaxAllowedPacket *int `mandatory:"false" json:"mysqlxMaxAllowedPacket"`

	// ("mysqlx_min_worker_threads") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxMinWorkerThreads *int `mandatory:"false" json:"mysqlxMinWorkerThreads"`

	// ("mysqlx_read_timeout") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxReadTimeout *int `mandatory:"false" json:"mysqlxReadTimeout"`

	// ("mysqlx_wait_timeout") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxWaitTimeout *int `mandatory:"false" json:"mysqlxWaitTimeout"`

	// ("mysqlx_write_timeout") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxWriteTimeout *int `mandatory:"false" json:"mysqlxWriteTimeout"`

	// ("parser_max_mem_size")
	ParserMaxMemSize *int `mandatory:"false" json:"parserMaxMemSize"`

	// ("query_alloc_block_size") DEPRECATED -- variable should not be settable and will be ignored
	QueryAllocBlockSize *int `mandatory:"false" json:"queryAllocBlockSize"`

	// ("query_prealloc_size") DEPRECATED -- variable should not be settable and will be ignored
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
	MysqlxZstdDefaultCompressionLevel *int `mandatory:"false" json:"mysqlxZstdDefaultCompressionLevel"`

	// DEPRECATED -- typo of mysqlx_zstd_default_compression_level. variable will be ignored.
	MysqlZstdDefaultCompressionLevel *int `mandatory:"false" json:"mysqlZstdDefaultCompressionLevel"`
}

func (m ConfigurationVariables) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigurationVariables) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigurationVariablesCompletionTypeEnum(string(m.CompletionType)); !ok && m.CompletionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompletionType: %s. Supported values are: %s.", m.CompletionType, strings.Join(GetConfigurationVariablesCompletionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigurationVariablesDefaultAuthenticationPluginEnum(string(m.DefaultAuthenticationPlugin)); !ok && m.DefaultAuthenticationPlugin != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultAuthenticationPlugin: %s. Supported values are: %s.", m.DefaultAuthenticationPlugin, strings.Join(GetConfigurationVariablesDefaultAuthenticationPluginEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigurationVariablesTransactionIsolationEnum(string(m.TransactionIsolation)); !ok && m.TransactionIsolation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransactionIsolation: %s. Supported values are: %s.", m.TransactionIsolation, strings.Join(GetConfigurationVariablesTransactionIsolationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigurationVariablesGroupReplicationConsistencyEnum(string(m.GroupReplicationConsistency)); !ok && m.GroupReplicationConsistency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupReplicationConsistency: %s. Supported values are: %s.", m.GroupReplicationConsistency, strings.Join(GetConfigurationVariablesGroupReplicationConsistencyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigurationVariablesBinlogRowMetadataEnum(string(m.BinlogRowMetadata)); !ok && m.BinlogRowMetadata != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BinlogRowMetadata: %s. Supported values are: %s.", m.BinlogRowMetadata, strings.Join(GetConfigurationVariablesBinlogRowMetadataEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigurationVariablesCompletionTypeEnum Enum with underlying type: string
type ConfigurationVariablesCompletionTypeEnum string

// Set of constants representing the allowable values for ConfigurationVariablesCompletionTypeEnum
const (
	ConfigurationVariablesCompletionTypeNoChain ConfigurationVariablesCompletionTypeEnum = "NO_CHAIN"
	ConfigurationVariablesCompletionTypeChain   ConfigurationVariablesCompletionTypeEnum = "CHAIN"
	ConfigurationVariablesCompletionTypeRelease ConfigurationVariablesCompletionTypeEnum = "RELEASE"
)

var mappingConfigurationVariablesCompletionTypeEnum = map[string]ConfigurationVariablesCompletionTypeEnum{
	"NO_CHAIN": ConfigurationVariablesCompletionTypeNoChain,
	"CHAIN":    ConfigurationVariablesCompletionTypeChain,
	"RELEASE":  ConfigurationVariablesCompletionTypeRelease,
}

var mappingConfigurationVariablesCompletionTypeEnumLowerCase = map[string]ConfigurationVariablesCompletionTypeEnum{
	"no_chain": ConfigurationVariablesCompletionTypeNoChain,
	"chain":    ConfigurationVariablesCompletionTypeChain,
	"release":  ConfigurationVariablesCompletionTypeRelease,
}

// GetConfigurationVariablesCompletionTypeEnumValues Enumerates the set of values for ConfigurationVariablesCompletionTypeEnum
func GetConfigurationVariablesCompletionTypeEnumValues() []ConfigurationVariablesCompletionTypeEnum {
	values := make([]ConfigurationVariablesCompletionTypeEnum, 0)
	for _, v := range mappingConfigurationVariablesCompletionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesCompletionTypeEnumStringValues Enumerates the set of values in String for ConfigurationVariablesCompletionTypeEnum
func GetConfigurationVariablesCompletionTypeEnumStringValues() []string {
	return []string{
		"NO_CHAIN",
		"CHAIN",
		"RELEASE",
	}
}

// GetMappingConfigurationVariablesCompletionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesCompletionTypeEnum(val string) (ConfigurationVariablesCompletionTypeEnum, bool) {
	enum, ok := mappingConfigurationVariablesCompletionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigurationVariablesDefaultAuthenticationPluginEnum Enum with underlying type: string
type ConfigurationVariablesDefaultAuthenticationPluginEnum string

// Set of constants representing the allowable values for ConfigurationVariablesDefaultAuthenticationPluginEnum
const (
	ConfigurationVariablesDefaultAuthenticationPluginMysqlNativePassword ConfigurationVariablesDefaultAuthenticationPluginEnum = "mysql_native_password"
	ConfigurationVariablesDefaultAuthenticationPluginSha256Password      ConfigurationVariablesDefaultAuthenticationPluginEnum = "sha256_password"
	ConfigurationVariablesDefaultAuthenticationPluginCachingSha2Password ConfigurationVariablesDefaultAuthenticationPluginEnum = "caching_sha2_password"
)

var mappingConfigurationVariablesDefaultAuthenticationPluginEnum = map[string]ConfigurationVariablesDefaultAuthenticationPluginEnum{
	"mysql_native_password": ConfigurationVariablesDefaultAuthenticationPluginMysqlNativePassword,
	"sha256_password":       ConfigurationVariablesDefaultAuthenticationPluginSha256Password,
	"caching_sha2_password": ConfigurationVariablesDefaultAuthenticationPluginCachingSha2Password,
}

var mappingConfigurationVariablesDefaultAuthenticationPluginEnumLowerCase = map[string]ConfigurationVariablesDefaultAuthenticationPluginEnum{
	"mysql_native_password": ConfigurationVariablesDefaultAuthenticationPluginMysqlNativePassword,
	"sha256_password":       ConfigurationVariablesDefaultAuthenticationPluginSha256Password,
	"caching_sha2_password": ConfigurationVariablesDefaultAuthenticationPluginCachingSha2Password,
}

// GetConfigurationVariablesDefaultAuthenticationPluginEnumValues Enumerates the set of values for ConfigurationVariablesDefaultAuthenticationPluginEnum
func GetConfigurationVariablesDefaultAuthenticationPluginEnumValues() []ConfigurationVariablesDefaultAuthenticationPluginEnum {
	values := make([]ConfigurationVariablesDefaultAuthenticationPluginEnum, 0)
	for _, v := range mappingConfigurationVariablesDefaultAuthenticationPluginEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesDefaultAuthenticationPluginEnumStringValues Enumerates the set of values in String for ConfigurationVariablesDefaultAuthenticationPluginEnum
func GetConfigurationVariablesDefaultAuthenticationPluginEnumStringValues() []string {
	return []string{
		"mysql_native_password",
		"sha256_password",
		"caching_sha2_password",
	}
}

// GetMappingConfigurationVariablesDefaultAuthenticationPluginEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesDefaultAuthenticationPluginEnum(val string) (ConfigurationVariablesDefaultAuthenticationPluginEnum, bool) {
	enum, ok := mappingConfigurationVariablesDefaultAuthenticationPluginEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigurationVariablesTransactionIsolationEnum Enum with underlying type: string
type ConfigurationVariablesTransactionIsolationEnum string

// Set of constants representing the allowable values for ConfigurationVariablesTransactionIsolationEnum
const (
	ConfigurationVariablesTransactionIsolationReadUncommitted ConfigurationVariablesTransactionIsolationEnum = "READ-UNCOMMITTED"
	ConfigurationVariablesTransactionIsolationReadCommited    ConfigurationVariablesTransactionIsolationEnum = "READ-COMMITED"
	ConfigurationVariablesTransactionIsolationReadCommitted   ConfigurationVariablesTransactionIsolationEnum = "READ-COMMITTED"
	ConfigurationVariablesTransactionIsolationRepeatableRead  ConfigurationVariablesTransactionIsolationEnum = "REPEATABLE-READ"
	ConfigurationVariablesTransactionIsolationSerializable    ConfigurationVariablesTransactionIsolationEnum = "SERIALIZABLE"
)

var mappingConfigurationVariablesTransactionIsolationEnum = map[string]ConfigurationVariablesTransactionIsolationEnum{
	"READ-UNCOMMITTED": ConfigurationVariablesTransactionIsolationReadUncommitted,
	"READ-COMMITED":    ConfigurationVariablesTransactionIsolationReadCommited,
	"READ-COMMITTED":   ConfigurationVariablesTransactionIsolationReadCommitted,
	"REPEATABLE-READ":  ConfigurationVariablesTransactionIsolationRepeatableRead,
	"SERIALIZABLE":     ConfigurationVariablesTransactionIsolationSerializable,
}

var mappingConfigurationVariablesTransactionIsolationEnumLowerCase = map[string]ConfigurationVariablesTransactionIsolationEnum{
	"read-uncommitted": ConfigurationVariablesTransactionIsolationReadUncommitted,
	"read-commited":    ConfigurationVariablesTransactionIsolationReadCommited,
	"read-committed":   ConfigurationVariablesTransactionIsolationReadCommitted,
	"repeatable-read":  ConfigurationVariablesTransactionIsolationRepeatableRead,
	"serializable":     ConfigurationVariablesTransactionIsolationSerializable,
}

// GetConfigurationVariablesTransactionIsolationEnumValues Enumerates the set of values for ConfigurationVariablesTransactionIsolationEnum
func GetConfigurationVariablesTransactionIsolationEnumValues() []ConfigurationVariablesTransactionIsolationEnum {
	values := make([]ConfigurationVariablesTransactionIsolationEnum, 0)
	for _, v := range mappingConfigurationVariablesTransactionIsolationEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesTransactionIsolationEnumStringValues Enumerates the set of values in String for ConfigurationVariablesTransactionIsolationEnum
func GetConfigurationVariablesTransactionIsolationEnumStringValues() []string {
	return []string{
		"READ-UNCOMMITTED",
		"READ-COMMITED",
		"READ-COMMITTED",
		"REPEATABLE-READ",
		"SERIALIZABLE",
	}
}

// GetMappingConfigurationVariablesTransactionIsolationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesTransactionIsolationEnum(val string) (ConfigurationVariablesTransactionIsolationEnum, bool) {
	enum, ok := mappingConfigurationVariablesTransactionIsolationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigurationVariablesGroupReplicationConsistencyEnum Enum with underlying type: string
type ConfigurationVariablesGroupReplicationConsistencyEnum string

// Set of constants representing the allowable values for ConfigurationVariablesGroupReplicationConsistencyEnum
const (
	ConfigurationVariablesGroupReplicationConsistencyEventual                ConfigurationVariablesGroupReplicationConsistencyEnum = "EVENTUAL"
	ConfigurationVariablesGroupReplicationConsistencyBeforeOnPrimaryFailover ConfigurationVariablesGroupReplicationConsistencyEnum = "BEFORE_ON_PRIMARY_FAILOVER"
	ConfigurationVariablesGroupReplicationConsistencyBefore                  ConfigurationVariablesGroupReplicationConsistencyEnum = "BEFORE"
	ConfigurationVariablesGroupReplicationConsistencyAfter                   ConfigurationVariablesGroupReplicationConsistencyEnum = "AFTER"
	ConfigurationVariablesGroupReplicationConsistencyBeforeAndAfter          ConfigurationVariablesGroupReplicationConsistencyEnum = "BEFORE_AND_AFTER"
)

var mappingConfigurationVariablesGroupReplicationConsistencyEnum = map[string]ConfigurationVariablesGroupReplicationConsistencyEnum{
	"EVENTUAL":                   ConfigurationVariablesGroupReplicationConsistencyEventual,
	"BEFORE_ON_PRIMARY_FAILOVER": ConfigurationVariablesGroupReplicationConsistencyBeforeOnPrimaryFailover,
	"BEFORE":                     ConfigurationVariablesGroupReplicationConsistencyBefore,
	"AFTER":                      ConfigurationVariablesGroupReplicationConsistencyAfter,
	"BEFORE_AND_AFTER":           ConfigurationVariablesGroupReplicationConsistencyBeforeAndAfter,
}

var mappingConfigurationVariablesGroupReplicationConsistencyEnumLowerCase = map[string]ConfigurationVariablesGroupReplicationConsistencyEnum{
	"eventual":                   ConfigurationVariablesGroupReplicationConsistencyEventual,
	"before_on_primary_failover": ConfigurationVariablesGroupReplicationConsistencyBeforeOnPrimaryFailover,
	"before":                     ConfigurationVariablesGroupReplicationConsistencyBefore,
	"after":                      ConfigurationVariablesGroupReplicationConsistencyAfter,
	"before_and_after":           ConfigurationVariablesGroupReplicationConsistencyBeforeAndAfter,
}

// GetConfigurationVariablesGroupReplicationConsistencyEnumValues Enumerates the set of values for ConfigurationVariablesGroupReplicationConsistencyEnum
func GetConfigurationVariablesGroupReplicationConsistencyEnumValues() []ConfigurationVariablesGroupReplicationConsistencyEnum {
	values := make([]ConfigurationVariablesGroupReplicationConsistencyEnum, 0)
	for _, v := range mappingConfigurationVariablesGroupReplicationConsistencyEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesGroupReplicationConsistencyEnumStringValues Enumerates the set of values in String for ConfigurationVariablesGroupReplicationConsistencyEnum
func GetConfigurationVariablesGroupReplicationConsistencyEnumStringValues() []string {
	return []string{
		"EVENTUAL",
		"BEFORE_ON_PRIMARY_FAILOVER",
		"BEFORE",
		"AFTER",
		"BEFORE_AND_AFTER",
	}
}

// GetMappingConfigurationVariablesGroupReplicationConsistencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesGroupReplicationConsistencyEnum(val string) (ConfigurationVariablesGroupReplicationConsistencyEnum, bool) {
	enum, ok := mappingConfigurationVariablesGroupReplicationConsistencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigurationVariablesBinlogRowMetadataEnum Enum with underlying type: string
type ConfigurationVariablesBinlogRowMetadataEnum string

// Set of constants representing the allowable values for ConfigurationVariablesBinlogRowMetadataEnum
const (
	ConfigurationVariablesBinlogRowMetadataFull    ConfigurationVariablesBinlogRowMetadataEnum = "FULL"
	ConfigurationVariablesBinlogRowMetadataMinimal ConfigurationVariablesBinlogRowMetadataEnum = "MINIMAL"
)

var mappingConfigurationVariablesBinlogRowMetadataEnum = map[string]ConfigurationVariablesBinlogRowMetadataEnum{
	"FULL":    ConfigurationVariablesBinlogRowMetadataFull,
	"MINIMAL": ConfigurationVariablesBinlogRowMetadataMinimal,
}

var mappingConfigurationVariablesBinlogRowMetadataEnumLowerCase = map[string]ConfigurationVariablesBinlogRowMetadataEnum{
	"full":    ConfigurationVariablesBinlogRowMetadataFull,
	"minimal": ConfigurationVariablesBinlogRowMetadataMinimal,
}

// GetConfigurationVariablesBinlogRowMetadataEnumValues Enumerates the set of values for ConfigurationVariablesBinlogRowMetadataEnum
func GetConfigurationVariablesBinlogRowMetadataEnumValues() []ConfigurationVariablesBinlogRowMetadataEnum {
	values := make([]ConfigurationVariablesBinlogRowMetadataEnum, 0)
	for _, v := range mappingConfigurationVariablesBinlogRowMetadataEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesBinlogRowMetadataEnumStringValues Enumerates the set of values in String for ConfigurationVariablesBinlogRowMetadataEnum
func GetConfigurationVariablesBinlogRowMetadataEnumStringValues() []string {
	return []string{
		"FULL",
		"MINIMAL",
	}
}

// GetMappingConfigurationVariablesBinlogRowMetadataEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesBinlogRowMetadataEnum(val string) (ConfigurationVariablesBinlogRowMetadataEnum, bool) {
	enum, ok := mappingConfigurationVariablesBinlogRowMetadataEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
