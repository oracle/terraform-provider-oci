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

// ConfigurationVariables User-defined service variables.
type ConfigurationVariables struct {

	// ("completion_type")
	CompletionType ConfigurationVariablesCompletionTypeEnum `mandatory:"false" json:"completionType,omitempty"`

	// If enabled, the server stores all temporary tables on disk rather than in memory.
	// bigTables corresponds to the MySQL server variable big_tables (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_big_tables).
	BigTables *bool `mandatory:"false" json:"bigTables"`

	// Set the chunking size for updates to the global memory usage counter Global_connection_memory.
	// connectionMemoryChunkSize corresponds to the MySQL system variable connection_memory_chunk_size (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_chunk_size).
	ConnectionMemoryChunkSize *int `mandatory:"false" json:"connectionMemoryChunkSize"`

	// Set the maximum amount of memory that can be used by a single user connection.
	// connectionMemoryLimit corresponds to the MySQL system variable connection_memory_limit (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_limit).
	ConnectionMemoryLimit *int64 `mandatory:"false" json:"connectionMemoryLimit"`

	// ("default_authentication_plugin")
	DefaultAuthenticationPlugin ConfigurationVariablesDefaultAuthenticationPluginEnum `mandatory:"false" json:"defaultAuthenticationPlugin,omitempty"`

	// Set the total amount of memory that can be used by all user connections.
	// globalConnectionMemoryLimit corresponds to the MySQL system variable global_connection_memory_limit (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_global_connection_memory_limit).
	GlobalConnectionMemoryLimit *int64 `mandatory:"false" json:"globalConnectionMemoryLimit"`

	// Determines whether the MySQL server calculates Global_connection_memory.
	// globalConnectionMemoryTracking corresponds to the MySQL system variable global_connection_memory_tracking (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_global_connection_memory_tracking).
	GlobalConnectionMemoryTracking *bool `mandatory:"false" json:"globalConnectionMemoryTracking"`

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

	// Enables dedicated log writer threads for writing redo log records from the log buffer to the system buffers and flushing the system buffers to the redo log files.
	// This is the MySQL variable "innodb_log_writer_threads". For more information, please see the MySQL documentation (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_log_writer_threads)
	InnodbLogWriterThreads *bool `mandatory:"false" json:"innodbLogWriterThreads"`

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

	// The size (in bytes) of the buffer pool, that is, the memory area where InnoDB caches table and index data.
	// innodbBufferPoolSize corresponds to the MySQL server system variable
	// innodb_buffer_pool_size (https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_buffer_pool_size).
	// The default and maximum values depend on the amount of RAM provisioned by the shape.
	// See Default User Variables (https://docs.cloud.oracle.com/mysql-database/doc/configuring-db-system.html#GUID-B5504C19-F6F4-4DAB-8506-189A4E8F4A6A).
	InnodbBufferPoolSize *int64 `mandatory:"false" json:"innodbBufferPoolSize"`

	// ("innodb_ft_result_cache_limit")
	InnodbFtResultCacheLimit *int64 `mandatory:"false" json:"innodbFtResultCacheLimit"`

	// Sets the size of the transaction cache.
	// maxBinlogCacheSize corresponds to the MySQL server system variable max_binlog_cache_size (https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_max_binlog_cache_size).
	MaxBinlogCacheSize *int64 `mandatory:"false" json:"maxBinlogCacheSize"`

	// ("max_connect_errors")
	MaxConnectErrors *int64 `mandatory:"false" json:"maxConnectErrors"`

	// This variable sets the maximum size to which user-created MEMORY tables are permitted to grow.
	// maxHeapTableSize corresponds to the MySQL system variable
	// max_heap_table_size (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_heap_table_size)
	MaxHeapTableSize *int64 `mandatory:"false" json:"maxHeapTableSize"`

	// ("max_connections")
	MaxConnections *int `mandatory:"false" json:"maxConnections"`

	// ("max_prepared_stmt_count")
	MaxPreparedStmtCount *int `mandatory:"false" json:"maxPreparedStmtCount"`

	// The number of seconds that the mysqld server waits for a connect packet before responding with Bad handshake.
	// connectTimeout corresponds to the MySQL system variable
	// connect_timeout (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_connect_timeout)
	// Increasing the connect_timeout value might help if clients frequently encounter errors of the form
	// "Lost connection to MySQL server at 'XXX', system error: errno".
	ConnectTimeout *int `mandatory:"false" json:"connectTimeout"`

	// ("cte_max_recursion_depth")
	CteMaxRecursionDepth *int64 `mandatory:"false" json:"cteMaxRecursionDepth"`

	// ("generated_random_password_length") DEPRECATED -- variable should not be settable and will be ignored
	GeneratedRandomPasswordLength *int `mandatory:"false" json:"generatedRandomPasswordLength"`

	// ("information_schema_stats_expiry")
	InformationSchemaStatsExpiry *int `mandatory:"false" json:"informationSchemaStatsExpiry"`

	// Specifies the percentage of the most recently used pages for each buffer pool to read out and dump.
	// innodbBufferPoolDumpPct corresponds to the MySQL InnoDB system variable
	// innodb_buffer_pool_dump_pct (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_buffer_pool_dump_pct).
	// The range is 1 to 100. The default value is 25.
	// For example, if there are 4 buffer pools with 100 pages each, and innodb_buffer_pool_dump_pct is set to 25,
	// the 25 most recently used pages from each buffer pool are dumped.
	InnodbBufferPoolDumpPct *int `mandatory:"false" json:"innodbBufferPoolDumpPct"`

	// ("innodb_buffer_pool_instances")
	InnodbBufferPoolInstances *int `mandatory:"false" json:"innodbBufferPoolInstances"`

	// innodbDdlBufferSize corresponds to the MySQL system variable innodb_ddl_buffer_size  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_ddl_buffer_size)
	InnodbDdlBufferSize *int64 `mandatory:"false" json:"innodbDdlBufferSize"`

	// innodbDdlThreads corresponds to the MySQL system variable innodb_ddl_threads  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_ddl_threads)
	InnodbDdlThreads *int `mandatory:"false" json:"innodbDdlThreads"`

	// ("innodb_ft_max_token_size")
	InnodbFtMaxTokenSize *int `mandatory:"false" json:"innodbFtMaxTokenSize"`

	// ("innodb_ft_min_token_size")
	InnodbFtMinTokenSize *int `mandatory:"false" json:"innodbFtMinTokenSize"`

	// ("innodb_ft_num_word_optimize")
	InnodbFtNumWordOptimize *int `mandatory:"false" json:"innodbFtNumWordOptimize"`

	// ("innodb_lock_wait_timeout")
	InnodbLockWaitTimeout *int `mandatory:"false" json:"innodbLockWaitTimeout"`

	// The desired maximum purge lag in terms of transactions.
	// InnoDB maintains a list of transactions that have index records delete-marked by UPDATE or DELETE operations. The length of the list is the purge lag.
	// If this value is exceeded, a delay is imposed on INSERT, UPDATE, and DELETE operations to allow time for purge to catch up.
	// The default value is 0, which means there is no maximum purge lag and no delay.
	// innodbMaxPurgeLag corresponds to the MySQL server system variable
	// innodb_max_purge_lag (https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_max_purge_lag).
	InnodbMaxPurgeLag *int64 `mandatory:"false" json:"innodbMaxPurgeLag"`

	// The maximum delay in microseconds for the delay imposed when the innodb_max_purge_lag threshold is exceeded.
	// The specified innodb_max_purge_lag_delay value is an upper limit on the delay period.
	// innodbMaxPurgeLagDelay corresponds to the MySQL server system variable
	// innodb_max_purge_lag_delay (https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_max_purge_lag_delay).
	InnodbMaxPurgeLagDelay *int `mandatory:"false" json:"innodbMaxPurgeLagDelay"`

	// The number of seconds the server waits for activity on an interactive connection before closing it.
	// interactiveTimeout corresponds to the MySQL system variable.
	// interactive_timeout (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_interactive_timeout)
	InteractiveTimeout *int `mandatory:"false" json:"interactiveTimeout"`

	// The number of index pages to sample when estimating cardinality and other statistics for an indexed column,
	// such as those calculated by ANALYZE TABLE.
	// innodbStatsPersistentSamplePages corresponds to the MySQL InnoDB system variable
	// innodb_stats_persistent_sample_pages (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_stats_persistent_sample_pages)
	// innodb_stats_persistent_sample_pages only applies when innodb_stats_persistent is enabled for a table;
	// when innodb_stats_persistent is disabled, innodb_stats_transient_sample_pages applies instead.
	InnodbStatsPersistentSamplePages *int64 `mandatory:"false" json:"innodbStatsPersistentSamplePages"`

	// The number of index pages to sample when estimating cardinality and other statistics for an indexed column,
	// such as those calculated by ANALYZE TABLE (https://dev.mysql.com/doc/refman/8.0/en/analyze-table.html).
	// innodbStatsTransientSamplePages corresponds to the MySQL InnoDB system variable
	// innodb_stats_transient_sample_pages (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_stats_transient_sample_pages)
	// innodb_stats_transient_sample_pages only applies when innodb_stats_persistent is disabled for a table;
	// when innodb_stats_persistent is enabled, innodb_stats_persistent_sample_pages applies instead.
	// innodb_stats_persistent is ON by default and cannot be changed. It is possible to override it using the
	// STATS_PERSISTENT clause of the CREATE TABLE (https://dev.mysql.com/doc/refman/8.0/en/create-table.html) and
	// ALTER TABLE (https://dev.mysql.com/doc/refman/8.0/en/alter-table.html) statements.
	InnodbStatsTransientSamplePages *int64 `mandatory:"false" json:"innodbStatsTransientSamplePages"`

	// The maximum size of one packet or any generated/intermediate string.
	// This is the mysql variable "max_allowed_packet".
	MaxAllowedPacket *int `mandatory:"false" json:"maxAllowedPacket"`

	// ("max_execution_time")
	MaxExecutionTime *int64 `mandatory:"false" json:"maxExecutionTime"`

	// The number of seconds X Plugin waits for the first packet to be received from newly connected clients.
	// mysqlxConnectTimeout corresponds to the MySQL X Plugin system variable
	// mysqlx_connect_timeout (https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_connect_timeout)
	MysqlxConnectTimeout *int `mandatory:"false" json:"mysqlxConnectTimeout"`

	// ("mysqlx_document_id_unique_prefix") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxDocumentIdUniquePrefix *int `mandatory:"false" json:"mysqlxDocumentIdUniquePrefix"`

	// ("mysqlx_idle_worker_thread_timeout") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxIdleWorkerThreadTimeout *int `mandatory:"false" json:"mysqlxIdleWorkerThreadTimeout"`

	// The number of seconds to wait for interactive clients to timeout.
	// mysqlxInteractiveTimeout corresponds to the MySQL X Plugin system variable.
	// mysqlx_interactive_timeout (https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_interactive_timeout)
	MysqlxInteractiveTimeout *int `mandatory:"false" json:"mysqlxInteractiveTimeout"`

	// The maximum size of network packets that can be received by X Plugin.
	// This is the mysql variable "mysqlx_max_allowed_packet".
	MysqlxMaxAllowedPacket *int `mandatory:"false" json:"mysqlxMaxAllowedPacket"`

	// ("mysqlx_min_worker_threads") DEPRECATED -- variable should not be settable and will be ignored
	MysqlxMinWorkerThreads *int `mandatory:"false" json:"mysqlxMinWorkerThreads"`

	// The number of seconds that X Plugin waits for blocking read operations to complete. After this time, if the
	// read operation is not successful, X Plugin closes the connection and returns a warning notice with the error
	// code ER_IO_READ_ERROR to the client application.
	// mysqlxReadTimeout corresponds to the MySQL X Plugin system variable
	// mysqlx_read_timeout (https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_read_timeout)
	MysqlxReadTimeout *int `mandatory:"false" json:"mysqlxReadTimeout"`

	// The number of seconds that X Plugin waits for activity on a connection.
	// mysqlxWaitTimeout corresponds to the MySQL X Plugin system variable.
	// mysqlx_wait_timeout (https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_wait_timeout)
	MysqlxWaitTimeout *int `mandatory:"false" json:"mysqlxWaitTimeout"`

	// The number of seconds that X Plugin waits for blocking write operations to complete. After this time, if the
	// write operation is not successful, X Plugin closes the connection.
	// mysqlxReadmysqlxWriteTimeoutTimeout corresponds to the MySQL X Plugin system variable
	// mysqlx_write_timeout (https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_write_timeout)
	MysqlxWriteTimeout *int `mandatory:"false" json:"mysqlxWriteTimeout"`

	// The number of seconds to wait for more data from a connection before aborting the read.
	// netReadTimeout corresponds to the MySQL system variable
	// net_read_timeout (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_net_read_timeout)
	NetReadTimeout *int `mandatory:"false" json:"netReadTimeout"`

	// The number of seconds to wait for a block to be written to a connection before aborting the write.
	// netWriteTimeout corresponds to the MySQL system variable
	// net_write_timeout (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_net_write_timeout)
	NetWriteTimeout *int `mandatory:"false" json:"netWriteTimeout"`

	// ("parser_max_mem_size")
	ParserMaxMemSize *int64 `mandatory:"false" json:"parserMaxMemSize"`

	// ("query_alloc_block_size") DEPRECATED -- variable should not be settable and will be ignored
	QueryAllocBlockSize *int64 `mandatory:"false" json:"queryAllocBlockSize"`

	// ("query_prealloc_size") DEPRECATED -- variable should not be settable and will be ignored
	QueryPreallocSize *int64 `mandatory:"false" json:"queryPreallocSize"`

	// regexpTimeLimit corresponds to the MySQL system variable regexp_time_limit  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_regexp_time_limit)
	RegexpTimeLimit *int `mandatory:"false" json:"regexpTimeLimit"`

	// ("sql_mode")
	SqlMode *string `mandatory:"false" json:"sqlMode"`

	// The maximum size of internal in-memory temporary tables. This variable does not apply to user-created MEMORY tables.
	// tmp_table_size corresponds to the MySQL system variable
	// tmp_table_size (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_tmp_table_size)
	TmpTableSize *int64 `mandatory:"false" json:"tmpTableSize"`

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

	// Each session that must perform a sort allocates a buffer of this size.
	// sortBufferSize corresponds to the MySQL system variable sort_buffer_size (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_sort_buffer_size)
	SortBufferSize *int64 `mandatory:"false" json:"sortBufferSize"`

	// The number of seconds the server waits for activity on a noninteractive connection before closing it.
	// waitTimeout corresponds to the MySQL system variable.
	// wait_timeout (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_wait_timeout)
	WaitTimeout *int `mandatory:"false" json:"waitTimeout"`

	// Controls whether the thread pool uses dedicated listener threads. If enabled, a listener thread in each thread group is dedicated to the task of listening
	// for network events from clients, ensuring that the maximum number of query worker threads is no more than the value specified by threadPoolMaxTransactionsLimit.
	// threadPoolDedicatedListeners corresponds to the MySQL Database Service-specific system variable thread_pool_dedicated_listeners.
	ThreadPoolDedicatedListeners *bool `mandatory:"false" json:"threadPoolDedicatedListeners"`

	// Limits the maximum number of open transactions to the defined value. The default value is 0, which enforces no limit.
	// threadPoolMaxTransactionsLimit corresponds to the MySQL Database Service-specific system variable thread_pool_max_transactions_limit.
	ThreadPoolMaxTransactionsLimit *int `mandatory:"false" json:"threadPoolMaxTransactionsLimit"`

	// Initializes the time zone for each client that connects.
	// This corresponds to the MySQL System Variable "time_zone".
	// The values can be given in one of the following formats, none of which are case-sensitive:
	// - As a string indicating an offset from UTC of the form [H]H:MM, prefixed with a + or -, such as '+10:00', '-6:00', or '+05:30'. The permitted range is '-13:59' to '+14:00', inclusive.
	// - As a named time zone, as defined by the "IANA Time Zone database", such as 'Europe/Helsinki', 'US/Eastern', 'MET', or 'UTC'.
	TimeZone *string `mandatory:"false" json:"timeZone"`
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
