// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

	// The server's default character set. If you set this variable, you should also set collation_server to specify the collation for the character set.
	// characterSetServer corresponds to the MySQL server variable character_set_server (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_character_set_server).
	CharacterSetServer ConfigurationVariablesCharacterSetServerEnum `mandatory:"false" json:"characterSetServer,omitempty"`

	// The server's default collation.
	// collationServer corresponds to the MySQL server variable collation_server (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_collation_server).
	CollationServer ConfigurationVariablesCollationServerEnum `mandatory:"false" json:"collationServer,omitempty"`

	// Set the chunking size for updates to the global memory usage counter Global_connection_memory.
	// connectionMemoryChunkSize corresponds to the MySQL system variable connection_memory_chunk_size (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_chunk_size).
	ConnectionMemoryChunkSize *int `mandatory:"false" json:"connectionMemoryChunkSize"`

	// Set the maximum amount of memory that can be used by a single user connection.
	// connectionMemoryLimit corresponds to the MySQL system variable connection_memory_limit (https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_limit).
	ConnectionMemoryLimit *int64 `mandatory:"false" json:"connectionMemoryLimit"`

	// The default authentication plugin. This must be a plugin that uses internal credentials storage, so these values are permitted:
	// mysql_native_password, sha256_password, caching_sha2_password.
	// As of MySQL 8.0.27, which introduces multifactor authentication, default_authentication_plugin is still used,
	// but in conjunction with and at a lower precedence than the authentication_policy system variable.
	// For details, see The Default Authentication Plugin. Because of this diminished role, default_authentication_plugin is deprecated as of MySQL 8.0.27
	// and subject to removal in a future MySQL version.
	// defaultAuthenticationPlugin corresponds to the MySQL system variable
	// default_authentication_plugin (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_default_authentication_plugin).
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

	// Specifies the maximum permitted result length in bytes for the GROUP_CONCAT() function.
	// This is the MySQL variable "group_concat_max_len". For more information, please see the MySQL documentation (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_group_concat_max_len)
	GroupConcatMaxLen *int64 `mandatory:"false" json:"groupConcatMaxLen"`

	// ("innodb_ft_enable_stopword")
	InnodbFtEnableStopword *bool `mandatory:"false" json:"innodbFtEnableStopword"`

	// Enables dedicated log writer threads for writing redo log records from the log buffer to the system buffers and flushing the system buffers to the redo log files.
	// This is the MySQL variable "innodb_log_writer_threads". For more information, please see the MySQL documentation (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_log_writer_threads)
	InnodbLogWriterThreads *bool `mandatory:"false" json:"innodbLogWriterThreads"`

	// This variable controls server-side LOCAL capability for LOAD DATA statements. Depending on the local_infile setting,
	// the server refuses or permits local data loading by clients that have LOCAL enabled on the client side.
	// local_infile corresponds to the MySQL Server system variable
	// local_infile (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_local_infile)
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
	// See Default User Variables (https://docs.oracle.com/iaas/mysql-database/doc/configuring-db-system.html#GUID-B5504C19-F6F4-4DAB-8506-189A4E8F4A6A).
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

	// When you enable innodbStrictMode, the InnoDB storage engine returns errors instead of warnings for invalid or incompatible table options.
	// innodbStrictMode corresponds to the MySQL InnoDB system variable
	// innodb_strict_mode (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_strict_mode)
	InnodbStrictMode *bool `mandatory:"false" json:"innodbStrictMode"`

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

	// This variable controls the block encryption mode for block-based algorithms such as AES. It affects encryption for AES_ENCRYPT() and AES_DECRYPT().
	//
	// block_encryption_mode takes a value in aes-keylen-mode format, where keylen is the key length in bits and mode is the encryption mode. The value is not case-sensitive.
	// Permitted keylen values are 128, 192, and 256. Permitted mode values are ECB, CBC, CFB1, CFB8, CFB128, and OFB.
	// block_encryption_mode corresponds to the MySQL Server Administration system variable
	// block_encryption_mode (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_block_encryption_mode)
	BlockEncryptionMode *string `mandatory:"false" json:"blockEncryptionMode"`

	// Controls how many microseconds the binary log commit waits before synchronizing the binary log file to disk.
	// There is no delay by default. Setting this variable to a microsecond delay enables more transactions to be synchronized
	// together to disk at once, reducing the overall time to commit a group of transactions because the larger groups required
	// fewer time units per group.
	// binlogGroupCommitSyncDelay corresponds to the MySQL Replication system variable
	// binlog_group_commit_sync_delay (https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_group_commit_sync_delay)
	BinlogGroupCommitSyncDelay *int `mandatory:"false" json:"binlogGroupCommitSyncDelay"`

	// The maximum number of transactions to wait for before aborting the current delay as specified by binlog_group_commit_sync_delay.
	// If binlog_group_commit_sync_delay is set to 0, then this option has no effect.
	// binlogGroupCommitSyncNoDelayCount corresponds to the MySQL Replication system variable
	// binlog_group_commit_sync_no_delay_count (https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_group_commit_sync_no_delay_count)
	BinlogGroupCommitSyncNoDelayCount *int `mandatory:"false" json:"binlogGroupCommitSyncNoDelayCount"`

	// Specifies the number of seconds to wait for more data or a heartbeat signal from the source before the replica considers the connection broken,
	// aborts the read, and tries to reconnect. Setting this variable has no immediate effect. The state of the variable applies on all subsequent START REPLICA commands.
	// replicaNetTimeout corresponds to the MySQL Replica server system variable
	// replica_net_timeout (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_net_timeout)
	ReplicaNetTimeout *int `mandatory:"false" json:"replicaNetTimeout"`

	// Whether client connections to the server are required to use some form of secure transport.
	// When this variable is enabled, the server permits only TCP/IP connections encrypted using TLS/SSL, or connections that use a socket file or shared memory.
	// The server rejects nonsecure connection attempts, which fail with an ER_SECURE_TRANSPORT_REQUIRED error.
	// require_secure_transport corresponds to the MySQL Server Administration system variable
	// require_secure_transport (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_require_secure_transport)
	RequireSecureTransport *bool `mandatory:"false" json:"requireSecureTransport"`

	// Defines the amount of disk space occupied by redo log files. innodb_redo_log_capacity supercedes the innodb_log_files_in_group and innodb_log_file_size variables,
	// which are both ignored if innodb_redo_log_capacity is defined. If innodb_redo_log_capacity is not defined, and if neither innodb_log_file_size or innodb_log_files_in_group are defined,
	// then the default innodb_redo_log_capacity value is used.
	// innodbRedoLogCapacity corresponds to the InnoDB Startup Options and System Variables
	// innodb_redo_log_capacity (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_redo_log_capacity)
	InnodbRedoLogCapacity *int64 `mandatory:"false" json:"innodbRedoLogCapacity"`

	// The delay period before executing a new transaction, in milliseconds. The maximum value is 300000 (5 minutes).
	// A transaction delay can be used in cases where parallel transactions affect the performance of other operations due to resource contention.
	// For example, if parallel transactions affect index creation or an online buffer pool resizing operation,
	// you can configure a transaction delay to reduce resource contention while those operations are running.
	// threadPoolTransactionDelay corresponds to the MySQL Server system variable
	// thread_pool_transaction_delay (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_transaction_delay)
	ThreadPoolTransactionDelay *int `mandatory:"false" json:"threadPoolTransactionDelay"`

	// The maximum number of query threads permitted in a thread group.
	// The maximum value is 4096, but if thread_pool_max_transactions_limit is set, thread_pool_query_threads_per_group must not exceed that value.
	// The default value of 1 means there is one active query thread in each thread group, which works well for many loads.
	// When you are using the high concurrency thread pool algorithm (thread_pool_algorithm = 1),
	// consider increasing the value if you experience slower response times due to long-running transactions.
	// threadPoolQueryThreadsPerGroup corresponds to the MySQL Server system variable
	// thread_pool_query_threads_per_group (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_query_threads_per_group)
	ThreadPoolQueryThreadsPerGroup *int `mandatory:"false" json:"threadPoolQueryThreadsPerGroup"`

	// This variable determines the default output format used by EXPLAIN in the absence of a FORMAT option when displaying a query execution plan.
	// explainFormat corresponds to the MySQL system variable
	// explain_format (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_explain_format).
	ExplainFormat ConfigurationVariablesExplainFormatEnum `mandatory:"false" json:"explainFormat,omitempty"`

	// This system variable determines whether the server enables certain nonstandard behaviors for default values and NULL-value handling in TIMESTAMP columns.
	// By default, explicit_defaults_for_timestamp is enabled, which disables the nonstandard behaviors. Disabling explicit_defaults_for_timestamp results in a warning.
	// explicit_defaults_for_timestamp corresponds to the MySQL Server Administration system variable
	// explicit_defaults_for_timestamp (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_explicit_defaults_for_timestamp)
	ExplicitDefaultsForTimestamp *bool `mandatory:"false" json:"explicitDefaultsForTimestamp"`

	// Whether GIPK mode is in effect, in which case a MySQL replication source server adds a generated invisible primary key to any InnoDB table that is created without one.
	// sqlGenerateInvisiblePrimaryKey corresponds to the MySQL system variable
	// sql_generate_invisible_primary_key  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_sql_generate_invisible_primary_key).
	SqlGenerateInvisiblePrimaryKey *bool `mandatory:"false" json:"sqlGenerateInvisiblePrimaryKey"`

	// Defines the maximum amount of memory that can be occupied by the TempTable storage engine before it starts storing data on disk.
	// The default value is 1073741824 bytes (1GiB). For more information, see Section 10.4.4, “Internal Temporary Table Use in MySQL”.
	// temptableMaxRam corresponds to the MySQL system variable
	// temptable_max_ram  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_temptable_max_ram).
	TemptableMaxRam *int64 `mandatory:"false" json:"temptableMaxRam"`

	// Whether InnoDB performs change buffering, an optimization that delays write operations to secondary indexes so that the I/O operations can be performed sequentially.
	// Permitted values are described in the following table. Values may also be specified numerically.
	// innodbChangeBuffering corresponds to the MySQL InnoDB Startup Options and System Variables
	// innodb_change_buffering  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_change_buffering).
	InnodbChangeBuffering ConfigurationVariablesInnodbChangeBufferingEnum `mandatory:"false" json:"innodbChangeBuffering,omitempty"`

	// Whether the InnoDB adaptive hash index is enabled or disabled.
	// It may be desirable, depending on your workload, to dynamically enable or disable adaptive hash indexing to improve query performance.
	// Because the adaptive hash index may not be useful for all workloads, conduct benchmarks with it both enabled and disabled, using realistic workloads.
	// innodbAdaptiveHashIndex corresponds to the MySQL InnoDB Startup Options and System Variables
	// innodb_adaptive_hash_index  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_adaptive_hash_index).
	InnodbAdaptiveHashIndex *bool `mandatory:"false" json:"innodbAdaptiveHashIndex"`

	// When enabled, undo tablespaces that exceed the threshold value defined by innodb_max_undo_log_size are marked for truncation.
	// Only undo tablespaces can be truncated. Truncating undo logs that reside in the system tablespace is not supported.
	// For truncation to occur, there must be at least two undo tablespaces.
	// innodbUndoLogTruncate corresponds to the MySQL InnoDB Startup Options and System Variables
	// innodb_undo_log_truncate  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_undo_log_truncate).
	InnodbUndoLogTruncate *bool `mandatory:"false" json:"innodbUndoLogTruncate"`

	// The number of table definitions that can be stored in the table definition cache.
	// If you use a large number of tables, you can create a large table definition cache to speed up opening of tables.
	// The table definition cache takes less space and does not use file descriptors, unlike the normal table cache.
	// table_definition_cache corresponds to the MySQL Server Administration system variable
	// table_definition_cache (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_table_definition_cache)
	TableDefinitionCache *int `mandatory:"false" json:"tableDefinitionCache"`

	// The number of open tables for all threads. Increasing this value increases the number of file descriptors that mysqld requires.
	// table_open_cache corresponds to the MySQL Server Administration system variable
	// table_open_cache (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_table_open_cache)
	TableOpenCache *int `mandatory:"false" json:"tableOpenCache"`

	// The maximum amount of space to use for all relay logs.
	// relayLogSpaceLimit corresponds to the MySQL Replica Server Options variable
	// relay_log_space_limit  (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_relay_log_space_limit).
	RelayLogSpaceLimit *int64 `mandatory:"false" json:"relayLogSpaceLimit"`

	// The optimizer_switch system variable enables control over optimizer behavior.
	// The value of this variable is a set of flags, each of which has a value of on or off to indicate whether the corresponding optimizer behavior is enabled or disabled.
	// This variable has global and session values and can be changed at runtime. The global default can be set at server startup.
	// Setting hypergraph_optimizer=on for cloud builds below 9.0.0 will fail.
	// optimizerSwitch corresponds to the MySQL Server System variable
	// optimizer_switch  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_optimizer_switch).
	OptimizerSwitch *string `mandatory:"false" json:"optimizerSwitch"`

	// From MySQL 8.0.26, use replica_type_conversions in place of slave_type_conversions, which is deprecated from that release.
	// In releases before MySQL 8.0.26, use slave_type_conversions.
	// replica_type_conversions controls the type conversion mode in effect on the replica when using row-based replication.
	// Its value is a comma-delimited set of zero or more elements from the list: ALL_LOSSY, ALL_NON_LOSSY, ALL_SIGNED, ALL_UNSIGNED.
	// Set this variable to an empty string to disallow type conversions between the source and the replica.
	// Setting this variable takes effect for all replication channels immediately, including running channels.
	// replica_type_conversions corresponds to the MySQL Replica Server Options variable
	// replica_type_conversions  (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_type_conversions).
	ReplicaTypeConversions *string `mandatory:"false" json:"replicaTypeConversions"`

	// Beginning with MySQL 8.0.26, slave_parallel_workers is deprecated, and you should use replica_parallel_workers instead.
	// (Prior to MySQL 8.0.26, you must use slave_parallel_workers to set the number of applier threads.)
	// replicaParallelWorkers corresponds to the MySQL Replica Server Options variable
	// replica_parallel_workers  (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_parallel_workers).
	ReplicaParallelWorkers *int `mandatory:"false" json:"replicaParallelWorkers"`

	// Whether to resolve host names when checking client connections. If this variable is OFF, mysqld resolves host names when checking client connections.
	// If it is ON, mysqld uses only IP numbers; in this case, all Host column values in the grant tables must be IP addresses.
	// See Section 7.1.12.3, “DNS Lookups and the Host Cache”.
	// skipNameResolve corresponds to the MySQL Server System variable
	// skip_name_resolve  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_skip_name_resolve).
	SkipNameResolve *bool `mandatory:"false" json:"skipNameResolve"`

	// The maximum number of simultaneous connections permitted to any given MySQL user account.
	// A value of 0 (the default) means “no limit.” This variable has a global value that can be set at server startup or runtime.
	// It also has a read-only session value that indicates the effective simultaneous-connection limit that applies to the account associated with the current session.
	// maxUserConnections corresponds to the MySQL Server System variable
	// max_user_connections  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_user_connections).
	MaxUserConnections *int64 `mandatory:"false" json:"maxUserConnections"`

	// The minimum size of the buffer that is used for plain index scans, range index scans, and joins that do not use indexes and thus perform full table scans.
	// In MySQL 8.0.18 and later, this variable also controls the amount of memory used for hash joins. Normally, the best way to get fast joins is to add indexes.
	// Increase the value of join_buffer_size to get a faster full join when adding indexes is not possible. One join buffer is allocated for each full join between two tables.
	// For a complex join between several tables for which indexes are not used, multiple join buffers might be necessary.
	// joinBufferSize corresponds to the MySQL Server System variable
	// join_buffer_size  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_join_buffer_size).
	JoinBufferSize *int64 `mandatory:"false" json:"joinBufferSize"`

	// Limit the assumed maximum number of seeks when looking up rows based on a key.
	// The MySQL optimizer assumes that no more than this number of key seeks are required when searching for matching rows in a table by scanning an index,
	// regardless of the actual cardinality of the index (see Section 15.7.7.22, “SHOW INDEX Statement”).
	// By setting this to a low value (say, 100), you can force MySQL to prefer indexes instead of table scans.
	// maxSeeksForKey corresponds to the MySQL Server System variable
	// max_seeks_for_key  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_seeks_for_key).
	MaxSeeksForKey *int64 `mandatory:"false" json:"maxSeeksForKey"`

	// The limit on memory consumption for the range optimizer. A value of 0 means “no limit.”
	// If an execution plan considered by the optimizer uses the range access method but the optimizer estimates that the amount of memory needed for this method would exceed the limit,
	// it abandons the plan and considers other plans.
	// rangeOptimizerMaxMemSize corresponds to the MySQL Server System variable
	// range_optimizer_max_mem_size  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_range_optimizer_max_mem_size).
	RangeOptimizerMaxMemSize *int64 `mandatory:"false" json:"rangeOptimizerMaxMemSize"`

	// auto_increment_increment and auto_increment_offset are intended for use with circular (source-to-source) replication,
	// and can be used to control the operation of AUTO_INCREMENT columns. Both variables have global and session values,
	// and each can assume an integer value between 1 and 65,535 inclusive.
	// autoIncrementIncrement corresponds to the MySQL Replication Source Options variable
	// auto_increment_increment  (https://dev.mysql.com/doc/refman/8.0/en/replication-options-source.html#sysvar_auto_increment_increment).
	AutoIncrementIncrement *int `mandatory:"false" json:"autoIncrementIncrement"`

	// This variable has a default value of 1. If it is left with its default value,
	// and Group Replication is started on the server in multi-primary mode, it is changed to the server ID.
	// autoIncrementOffset corresponds to the MySQL Replication Source Options variable
	// auto_increment_offset  (https://dev.mysql.com/doc/refman/8.0/en/replication-options-source.html#sysvar_auto_increment_offset).
	AutoIncrementOffset *int `mandatory:"false" json:"autoIncrementOffset"`

	// The lock mode to use for generating auto-increment values.
	// Permissible values are 0, 1, or 2, for traditional, consecutive, or interleaved, respectively.
	// innodbAutoincLockMode corresponds to the MySQL InnoDB Startup Options and System Variables
	// innodb_autoinc_lock_mode  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_autoinc_lock_mode).
	InnodbAutoincLockMode *int `mandatory:"false" json:"innodbAutoincLockMode"`

	// InnoDB rolls back only the last statement on a transaction timeout by default.
	// If --innodb-rollback-on-timeout is specified, a transaction timeout causes InnoDB to abort and roll back the entire transaction.
	// innodbRollbackOnTimeout corresponds to the MySQL InnoDB Startup Options and System Variables
	// innodb_rollback_on_timeout  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_rollback_on_timeout).
	InnodbRollbackOnTimeout *bool `mandatory:"false" json:"innodbRollbackOnTimeout"`

	// Specifies an upper limit in bytes on the size of the temporary log files used during online DDL operations for InnoDB tables.
	// There is one such log file for each index being created or table being altered.
	// This log file stores data inserted, updated, or deleted in the table during the DDL operation.
	// innodbOnlineAlterLogMaxSize corresponds to the MySQL InnoDB Startup Options and System Variables
	// innodb_online_alter_log_max_size  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_online_alter_log_max_size).
	InnodbOnlineAlterLogMaxSize *int64 `mandatory:"false" json:"innodbOnlineAlterLogMaxSize"`

	// This variable defines:
	// * The sort buffer size for online DDL operations that create or rebuild secondary indexes.
	//   However, as of MySQL 8.0.27, this responsibility is subsumed by the innodb_ddl_buffer_size variable.
	// * The amount by which the temporary log file is extended when recording concurrent DML during an online DDL operation,
	//   and the size of the temporary log file read buffer and write buffer.
	// innodbSortBufferSize corresponds to the MySQL InnoDB Startup Options and System Variables
	// innodb_sort_buffer_size  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_sort_buffer_size).
	InnodbSortBufferSize *int `mandatory:"false" json:"innodbSortBufferSize"`

	// Enables the NUMA interleave memory policy for allocation of the InnoDB buffer pool.
	// When innodb_numa_interleave is enabled, the NUMA memory policy is set to MPOL_INTERLEAVE for the mysqld process.
	// After the InnoDB buffer pool is allocated, the NUMA memory policy is set back to MPOL_DEFAULT.
	// For the innodb_numa_interleave option to be available, MySQL must be compiled on a NUMA-enabled Linux system.
	// innodbNumaInterleave corresponds to the MySQL InnoDB Startup Options and System Variables
	// innodb_numa_interleave  (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_numa_interleave).
	InnodbNumaInterleave *bool `mandatory:"false" json:"innodbNumaInterleave"`

	// The number of thread groups in the thread pool. This is the most important parameter controlling thread pool performance.
	// It affects how many statements can execute simultaneously. If a value outside the range of permissible values is specified,
	// the thread pool plugin does not load and the server writes a message to the error log.
	// threadPoolSize corresponds to the MySQL Server System variable
	// thread_pool_size  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_size).
	ThreadPoolSize *int `mandatory:"false" json:"threadPoolSize"`

	// If a query takes longer than this many seconds, the server increments the Slow_queries status variable.
	// If the slow query log is enabled, the query is logged to the slow query log file.
	// This value is measured in real time, not CPU time,
	// so a query that is under the threshold on a lightly loaded system might be above the threshold on a heavily loaded one.
	// longQueryTime corresponds to the MySQL Server System variable
	// long_query_time  (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_long_query_time).
	LongQueryTime *int `mandatory:"false" json:"longQueryTime"`
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
	if _, ok := GetMappingConfigurationVariablesCharacterSetServerEnum(string(m.CharacterSetServer)); !ok && m.CharacterSetServer != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CharacterSetServer: %s. Supported values are: %s.", m.CharacterSetServer, strings.Join(GetConfigurationVariablesCharacterSetServerEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigurationVariablesCollationServerEnum(string(m.CollationServer)); !ok && m.CollationServer != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CollationServer: %s. Supported values are: %s.", m.CollationServer, strings.Join(GetConfigurationVariablesCollationServerEnumStringValues(), ",")))
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
	if _, ok := GetMappingConfigurationVariablesExplainFormatEnum(string(m.ExplainFormat)); !ok && m.ExplainFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExplainFormat: %s. Supported values are: %s.", m.ExplainFormat, strings.Join(GetConfigurationVariablesExplainFormatEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfigurationVariablesInnodbChangeBufferingEnum(string(m.InnodbChangeBuffering)); !ok && m.InnodbChangeBuffering != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InnodbChangeBuffering: %s. Supported values are: %s.", m.InnodbChangeBuffering, strings.Join(GetConfigurationVariablesInnodbChangeBufferingEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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

// ConfigurationVariablesCharacterSetServerEnum Enum with underlying type: string
type ConfigurationVariablesCharacterSetServerEnum string

// Set of constants representing the allowable values for ConfigurationVariablesCharacterSetServerEnum
const (
	ConfigurationVariablesCharacterSetServerArmscii8 ConfigurationVariablesCharacterSetServerEnum = "ARMSCII8"
	ConfigurationVariablesCharacterSetServerAscii    ConfigurationVariablesCharacterSetServerEnum = "ASCII"
	ConfigurationVariablesCharacterSetServerBig5     ConfigurationVariablesCharacterSetServerEnum = "BIG5"
	ConfigurationVariablesCharacterSetServerBinary   ConfigurationVariablesCharacterSetServerEnum = "BINARY"
	ConfigurationVariablesCharacterSetServerCp1250   ConfigurationVariablesCharacterSetServerEnum = "CP1250"
	ConfigurationVariablesCharacterSetServerCp1251   ConfigurationVariablesCharacterSetServerEnum = "CP1251"
	ConfigurationVariablesCharacterSetServerCp1256   ConfigurationVariablesCharacterSetServerEnum = "CP1256"
	ConfigurationVariablesCharacterSetServerCp1257   ConfigurationVariablesCharacterSetServerEnum = "CP1257"
	ConfigurationVariablesCharacterSetServerCp850    ConfigurationVariablesCharacterSetServerEnum = "CP850"
	ConfigurationVariablesCharacterSetServerCp852    ConfigurationVariablesCharacterSetServerEnum = "CP852"
	ConfigurationVariablesCharacterSetServerCp866    ConfigurationVariablesCharacterSetServerEnum = "CP866"
	ConfigurationVariablesCharacterSetServerCp932    ConfigurationVariablesCharacterSetServerEnum = "CP932"
	ConfigurationVariablesCharacterSetServerDec8     ConfigurationVariablesCharacterSetServerEnum = "DEC8"
	ConfigurationVariablesCharacterSetServerEucjpms  ConfigurationVariablesCharacterSetServerEnum = "EUCJPMS"
	ConfigurationVariablesCharacterSetServerEuckr    ConfigurationVariablesCharacterSetServerEnum = "EUCKR"
	ConfigurationVariablesCharacterSetServerGb18030  ConfigurationVariablesCharacterSetServerEnum = "GB18030"
	ConfigurationVariablesCharacterSetServerGb2312   ConfigurationVariablesCharacterSetServerEnum = "GB2312"
	ConfigurationVariablesCharacterSetServerGbk      ConfigurationVariablesCharacterSetServerEnum = "GBK"
	ConfigurationVariablesCharacterSetServerGeostd8  ConfigurationVariablesCharacterSetServerEnum = "GEOSTD8"
	ConfigurationVariablesCharacterSetServerGreek    ConfigurationVariablesCharacterSetServerEnum = "GREEK"
	ConfigurationVariablesCharacterSetServerHebrew   ConfigurationVariablesCharacterSetServerEnum = "HEBREW"
	ConfigurationVariablesCharacterSetServerHp8      ConfigurationVariablesCharacterSetServerEnum = "HP8"
	ConfigurationVariablesCharacterSetServerKeybcs2  ConfigurationVariablesCharacterSetServerEnum = "KEYBCS2"
	ConfigurationVariablesCharacterSetServerKoi8r    ConfigurationVariablesCharacterSetServerEnum = "KOI8R"
	ConfigurationVariablesCharacterSetServerKoi8u    ConfigurationVariablesCharacterSetServerEnum = "KOI8U"
	ConfigurationVariablesCharacterSetServerLatin1   ConfigurationVariablesCharacterSetServerEnum = "LATIN1"
	ConfigurationVariablesCharacterSetServerLatin2   ConfigurationVariablesCharacterSetServerEnum = "LATIN2"
	ConfigurationVariablesCharacterSetServerLatin5   ConfigurationVariablesCharacterSetServerEnum = "LATIN5"
	ConfigurationVariablesCharacterSetServerLatin7   ConfigurationVariablesCharacterSetServerEnum = "LATIN7"
	ConfigurationVariablesCharacterSetServerMacce    ConfigurationVariablesCharacterSetServerEnum = "MACCE"
	ConfigurationVariablesCharacterSetServerMacroman ConfigurationVariablesCharacterSetServerEnum = "MACROMAN"
	ConfigurationVariablesCharacterSetServerSjis     ConfigurationVariablesCharacterSetServerEnum = "SJIS"
	ConfigurationVariablesCharacterSetServerSwe7     ConfigurationVariablesCharacterSetServerEnum = "SWE7"
	ConfigurationVariablesCharacterSetServerTis620   ConfigurationVariablesCharacterSetServerEnum = "TIS620"
	ConfigurationVariablesCharacterSetServerUcs2     ConfigurationVariablesCharacterSetServerEnum = "UCS2"
	ConfigurationVariablesCharacterSetServerUjis     ConfigurationVariablesCharacterSetServerEnum = "UJIS"
	ConfigurationVariablesCharacterSetServerUtf16    ConfigurationVariablesCharacterSetServerEnum = "UTF16"
	ConfigurationVariablesCharacterSetServerUtf16le  ConfigurationVariablesCharacterSetServerEnum = "UTF16LE"
	ConfigurationVariablesCharacterSetServerUtf32    ConfigurationVariablesCharacterSetServerEnum = "UTF32"
	ConfigurationVariablesCharacterSetServerUtf8mb3  ConfigurationVariablesCharacterSetServerEnum = "UTF8MB3"
	ConfigurationVariablesCharacterSetServerUtf8mb4  ConfigurationVariablesCharacterSetServerEnum = "UTF8MB4"
)

var mappingConfigurationVariablesCharacterSetServerEnum = map[string]ConfigurationVariablesCharacterSetServerEnum{
	"ARMSCII8": ConfigurationVariablesCharacterSetServerArmscii8,
	"ASCII":    ConfigurationVariablesCharacterSetServerAscii,
	"BIG5":     ConfigurationVariablesCharacterSetServerBig5,
	"BINARY":   ConfigurationVariablesCharacterSetServerBinary,
	"CP1250":   ConfigurationVariablesCharacterSetServerCp1250,
	"CP1251":   ConfigurationVariablesCharacterSetServerCp1251,
	"CP1256":   ConfigurationVariablesCharacterSetServerCp1256,
	"CP1257":   ConfigurationVariablesCharacterSetServerCp1257,
	"CP850":    ConfigurationVariablesCharacterSetServerCp850,
	"CP852":    ConfigurationVariablesCharacterSetServerCp852,
	"CP866":    ConfigurationVariablesCharacterSetServerCp866,
	"CP932":    ConfigurationVariablesCharacterSetServerCp932,
	"DEC8":     ConfigurationVariablesCharacterSetServerDec8,
	"EUCJPMS":  ConfigurationVariablesCharacterSetServerEucjpms,
	"EUCKR":    ConfigurationVariablesCharacterSetServerEuckr,
	"GB18030":  ConfigurationVariablesCharacterSetServerGb18030,
	"GB2312":   ConfigurationVariablesCharacterSetServerGb2312,
	"GBK":      ConfigurationVariablesCharacterSetServerGbk,
	"GEOSTD8":  ConfigurationVariablesCharacterSetServerGeostd8,
	"GREEK":    ConfigurationVariablesCharacterSetServerGreek,
	"HEBREW":   ConfigurationVariablesCharacterSetServerHebrew,
	"HP8":      ConfigurationVariablesCharacterSetServerHp8,
	"KEYBCS2":  ConfigurationVariablesCharacterSetServerKeybcs2,
	"KOI8R":    ConfigurationVariablesCharacterSetServerKoi8r,
	"KOI8U":    ConfigurationVariablesCharacterSetServerKoi8u,
	"LATIN1":   ConfigurationVariablesCharacterSetServerLatin1,
	"LATIN2":   ConfigurationVariablesCharacterSetServerLatin2,
	"LATIN5":   ConfigurationVariablesCharacterSetServerLatin5,
	"LATIN7":   ConfigurationVariablesCharacterSetServerLatin7,
	"MACCE":    ConfigurationVariablesCharacterSetServerMacce,
	"MACROMAN": ConfigurationVariablesCharacterSetServerMacroman,
	"SJIS":     ConfigurationVariablesCharacterSetServerSjis,
	"SWE7":     ConfigurationVariablesCharacterSetServerSwe7,
	"TIS620":   ConfigurationVariablesCharacterSetServerTis620,
	"UCS2":     ConfigurationVariablesCharacterSetServerUcs2,
	"UJIS":     ConfigurationVariablesCharacterSetServerUjis,
	"UTF16":    ConfigurationVariablesCharacterSetServerUtf16,
	"UTF16LE":  ConfigurationVariablesCharacterSetServerUtf16le,
	"UTF32":    ConfigurationVariablesCharacterSetServerUtf32,
	"UTF8MB3":  ConfigurationVariablesCharacterSetServerUtf8mb3,
	"UTF8MB4":  ConfigurationVariablesCharacterSetServerUtf8mb4,
}

var mappingConfigurationVariablesCharacterSetServerEnumLowerCase = map[string]ConfigurationVariablesCharacterSetServerEnum{
	"armscii8": ConfigurationVariablesCharacterSetServerArmscii8,
	"ascii":    ConfigurationVariablesCharacterSetServerAscii,
	"big5":     ConfigurationVariablesCharacterSetServerBig5,
	"binary":   ConfigurationVariablesCharacterSetServerBinary,
	"cp1250":   ConfigurationVariablesCharacterSetServerCp1250,
	"cp1251":   ConfigurationVariablesCharacterSetServerCp1251,
	"cp1256":   ConfigurationVariablesCharacterSetServerCp1256,
	"cp1257":   ConfigurationVariablesCharacterSetServerCp1257,
	"cp850":    ConfigurationVariablesCharacterSetServerCp850,
	"cp852":    ConfigurationVariablesCharacterSetServerCp852,
	"cp866":    ConfigurationVariablesCharacterSetServerCp866,
	"cp932":    ConfigurationVariablesCharacterSetServerCp932,
	"dec8":     ConfigurationVariablesCharacterSetServerDec8,
	"eucjpms":  ConfigurationVariablesCharacterSetServerEucjpms,
	"euckr":    ConfigurationVariablesCharacterSetServerEuckr,
	"gb18030":  ConfigurationVariablesCharacterSetServerGb18030,
	"gb2312":   ConfigurationVariablesCharacterSetServerGb2312,
	"gbk":      ConfigurationVariablesCharacterSetServerGbk,
	"geostd8":  ConfigurationVariablesCharacterSetServerGeostd8,
	"greek":    ConfigurationVariablesCharacterSetServerGreek,
	"hebrew":   ConfigurationVariablesCharacterSetServerHebrew,
	"hp8":      ConfigurationVariablesCharacterSetServerHp8,
	"keybcs2":  ConfigurationVariablesCharacterSetServerKeybcs2,
	"koi8r":    ConfigurationVariablesCharacterSetServerKoi8r,
	"koi8u":    ConfigurationVariablesCharacterSetServerKoi8u,
	"latin1":   ConfigurationVariablesCharacterSetServerLatin1,
	"latin2":   ConfigurationVariablesCharacterSetServerLatin2,
	"latin5":   ConfigurationVariablesCharacterSetServerLatin5,
	"latin7":   ConfigurationVariablesCharacterSetServerLatin7,
	"macce":    ConfigurationVariablesCharacterSetServerMacce,
	"macroman": ConfigurationVariablesCharacterSetServerMacroman,
	"sjis":     ConfigurationVariablesCharacterSetServerSjis,
	"swe7":     ConfigurationVariablesCharacterSetServerSwe7,
	"tis620":   ConfigurationVariablesCharacterSetServerTis620,
	"ucs2":     ConfigurationVariablesCharacterSetServerUcs2,
	"ujis":     ConfigurationVariablesCharacterSetServerUjis,
	"utf16":    ConfigurationVariablesCharacterSetServerUtf16,
	"utf16le":  ConfigurationVariablesCharacterSetServerUtf16le,
	"utf32":    ConfigurationVariablesCharacterSetServerUtf32,
	"utf8mb3":  ConfigurationVariablesCharacterSetServerUtf8mb3,
	"utf8mb4":  ConfigurationVariablesCharacterSetServerUtf8mb4,
}

// GetConfigurationVariablesCharacterSetServerEnumValues Enumerates the set of values for ConfigurationVariablesCharacterSetServerEnum
func GetConfigurationVariablesCharacterSetServerEnumValues() []ConfigurationVariablesCharacterSetServerEnum {
	values := make([]ConfigurationVariablesCharacterSetServerEnum, 0)
	for _, v := range mappingConfigurationVariablesCharacterSetServerEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesCharacterSetServerEnumStringValues Enumerates the set of values in String for ConfigurationVariablesCharacterSetServerEnum
func GetConfigurationVariablesCharacterSetServerEnumStringValues() []string {
	return []string{
		"ARMSCII8",
		"ASCII",
		"BIG5",
		"BINARY",
		"CP1250",
		"CP1251",
		"CP1256",
		"CP1257",
		"CP850",
		"CP852",
		"CP866",
		"CP932",
		"DEC8",
		"EUCJPMS",
		"EUCKR",
		"GB18030",
		"GB2312",
		"GBK",
		"GEOSTD8",
		"GREEK",
		"HEBREW",
		"HP8",
		"KEYBCS2",
		"KOI8R",
		"KOI8U",
		"LATIN1",
		"LATIN2",
		"LATIN5",
		"LATIN7",
		"MACCE",
		"MACROMAN",
		"SJIS",
		"SWE7",
		"TIS620",
		"UCS2",
		"UJIS",
		"UTF16",
		"UTF16LE",
		"UTF32",
		"UTF8MB3",
		"UTF8MB4",
	}
}

// GetMappingConfigurationVariablesCharacterSetServerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesCharacterSetServerEnum(val string) (ConfigurationVariablesCharacterSetServerEnum, bool) {
	enum, ok := mappingConfigurationVariablesCharacterSetServerEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigurationVariablesCollationServerEnum Enum with underlying type: string
type ConfigurationVariablesCollationServerEnum string

// Set of constants representing the allowable values for ConfigurationVariablesCollationServerEnum
const (
	ConfigurationVariablesCollationServerArmscii8GeneralCi     ConfigurationVariablesCollationServerEnum = "ARMSCII8_GENERAL_CI"
	ConfigurationVariablesCollationServerArmscii8Bin           ConfigurationVariablesCollationServerEnum = "ARMSCII8_BIN"
	ConfigurationVariablesCollationServerAsciiGeneralCi        ConfigurationVariablesCollationServerEnum = "ASCII_GENERAL_CI"
	ConfigurationVariablesCollationServerAsciiBin              ConfigurationVariablesCollationServerEnum = "ASCII_BIN"
	ConfigurationVariablesCollationServerBig5ChineseCi         ConfigurationVariablesCollationServerEnum = "BIG5_CHINESE_CI"
	ConfigurationVariablesCollationServerBig5Bin               ConfigurationVariablesCollationServerEnum = "BIG5_BIN"
	ConfigurationVariablesCollationServerBinary                ConfigurationVariablesCollationServerEnum = "BINARY"
	ConfigurationVariablesCollationServerCp1250GeneralCi       ConfigurationVariablesCollationServerEnum = "CP1250_GENERAL_CI"
	ConfigurationVariablesCollationServerCp1250Bin             ConfigurationVariablesCollationServerEnum = "CP1250_BIN"
	ConfigurationVariablesCollationServerCp1250CroatianCi      ConfigurationVariablesCollationServerEnum = "CP1250_CROATIAN_CI"
	ConfigurationVariablesCollationServerCp1250CzechCs         ConfigurationVariablesCollationServerEnum = "CP1250_CZECH_CS"
	ConfigurationVariablesCollationServerCp1250PolishCi        ConfigurationVariablesCollationServerEnum = "CP1250_POLISH_CI"
	ConfigurationVariablesCollationServerCp1251GeneralCi       ConfigurationVariablesCollationServerEnum = "CP1251_GENERAL_CI"
	ConfigurationVariablesCollationServerCp1251Bin             ConfigurationVariablesCollationServerEnum = "CP1251_BIN"
	ConfigurationVariablesCollationServerCp1251BulgarianCi     ConfigurationVariablesCollationServerEnum = "CP1251_BULGARIAN_CI"
	ConfigurationVariablesCollationServerCp1251GeneralCs       ConfigurationVariablesCollationServerEnum = "CP1251_GENERAL_CS"
	ConfigurationVariablesCollationServerCp1251UkrainianCi     ConfigurationVariablesCollationServerEnum = "CP1251_UKRAINIAN_CI"
	ConfigurationVariablesCollationServerCp1256GeneralCi       ConfigurationVariablesCollationServerEnum = "CP1256_GENERAL_CI"
	ConfigurationVariablesCollationServerCp1256Bin             ConfigurationVariablesCollationServerEnum = "CP1256_BIN"
	ConfigurationVariablesCollationServerCp1257GeneralCi       ConfigurationVariablesCollationServerEnum = "CP1257_GENERAL_CI"
	ConfigurationVariablesCollationServerCp1257Bin             ConfigurationVariablesCollationServerEnum = "CP1257_BIN"
	ConfigurationVariablesCollationServerCp1257LithuanianCi    ConfigurationVariablesCollationServerEnum = "CP1257_LITHUANIAN_CI"
	ConfigurationVariablesCollationServerCp850GeneralCi        ConfigurationVariablesCollationServerEnum = "CP850_GENERAL_CI"
	ConfigurationVariablesCollationServerCp850Bin              ConfigurationVariablesCollationServerEnum = "CP850_BIN"
	ConfigurationVariablesCollationServerCp852GeneralCi        ConfigurationVariablesCollationServerEnum = "CP852_GENERAL_CI"
	ConfigurationVariablesCollationServerCp852Bin              ConfigurationVariablesCollationServerEnum = "CP852_BIN"
	ConfigurationVariablesCollationServerCp866GeneralCi        ConfigurationVariablesCollationServerEnum = "CP866_GENERAL_CI"
	ConfigurationVariablesCollationServerCp866Bin              ConfigurationVariablesCollationServerEnum = "CP866_BIN"
	ConfigurationVariablesCollationServerCp932JapaneseCi       ConfigurationVariablesCollationServerEnum = "CP932_JAPANESE_CI"
	ConfigurationVariablesCollationServerCp932Bin              ConfigurationVariablesCollationServerEnum = "CP932_BIN"
	ConfigurationVariablesCollationServerDec8SwedishCi         ConfigurationVariablesCollationServerEnum = "DEC8_SWEDISH_CI"
	ConfigurationVariablesCollationServerDec8Bin               ConfigurationVariablesCollationServerEnum = "DEC8_BIN"
	ConfigurationVariablesCollationServerEucjpmsJapaneseCi     ConfigurationVariablesCollationServerEnum = "EUCJPMS_JAPANESE_CI"
	ConfigurationVariablesCollationServerEucjpmsBin            ConfigurationVariablesCollationServerEnum = "EUCJPMS_BIN"
	ConfigurationVariablesCollationServerEuckrKoreanCi         ConfigurationVariablesCollationServerEnum = "EUCKR_KOREAN_CI"
	ConfigurationVariablesCollationServerEuckrBin              ConfigurationVariablesCollationServerEnum = "EUCKR_BIN"
	ConfigurationVariablesCollationServerGb18030ChineseCi      ConfigurationVariablesCollationServerEnum = "GB18030_CHINESE_CI"
	ConfigurationVariablesCollationServerGb18030Bin            ConfigurationVariablesCollationServerEnum = "GB18030_BIN"
	ConfigurationVariablesCollationServerGb18030Unicode520Ci   ConfigurationVariablesCollationServerEnum = "GB18030_UNICODE_520_CI"
	ConfigurationVariablesCollationServerGb2312ChineseCi       ConfigurationVariablesCollationServerEnum = "GB2312_CHINESE_CI"
	ConfigurationVariablesCollationServerGb2312Bin             ConfigurationVariablesCollationServerEnum = "GB2312_BIN"
	ConfigurationVariablesCollationServerGbkChineseCi          ConfigurationVariablesCollationServerEnum = "GBK_CHINESE_CI"
	ConfigurationVariablesCollationServerGbkBin                ConfigurationVariablesCollationServerEnum = "GBK_BIN"
	ConfigurationVariablesCollationServerGeostd8GeneralCi      ConfigurationVariablesCollationServerEnum = "GEOSTD8_GENERAL_CI"
	ConfigurationVariablesCollationServerGeostd8Bin            ConfigurationVariablesCollationServerEnum = "GEOSTD8_BIN"
	ConfigurationVariablesCollationServerGreekGeneralCi        ConfigurationVariablesCollationServerEnum = "GREEK_GENERAL_CI"
	ConfigurationVariablesCollationServerGreekBin              ConfigurationVariablesCollationServerEnum = "GREEK_BIN"
	ConfigurationVariablesCollationServerHebrewGeneralCi       ConfigurationVariablesCollationServerEnum = "HEBREW_GENERAL_CI"
	ConfigurationVariablesCollationServerHebrewBin             ConfigurationVariablesCollationServerEnum = "HEBREW_BIN"
	ConfigurationVariablesCollationServerHp8EnglishCi          ConfigurationVariablesCollationServerEnum = "HP8_ENGLISH_CI"
	ConfigurationVariablesCollationServerHp8Bin                ConfigurationVariablesCollationServerEnum = "HP8_BIN"
	ConfigurationVariablesCollationServerKeybcs2GeneralCi      ConfigurationVariablesCollationServerEnum = "KEYBCS2_GENERAL_CI"
	ConfigurationVariablesCollationServerKeybcs2Bin            ConfigurationVariablesCollationServerEnum = "KEYBCS2_BIN"
	ConfigurationVariablesCollationServerKoi8rGeneralCi        ConfigurationVariablesCollationServerEnum = "KOI8R_GENERAL_CI"
	ConfigurationVariablesCollationServerKoi8rBin              ConfigurationVariablesCollationServerEnum = "KOI8R_BIN"
	ConfigurationVariablesCollationServerKoi8uGeneralCi        ConfigurationVariablesCollationServerEnum = "KOI8U_GENERAL_CI"
	ConfigurationVariablesCollationServerKoi8uBin              ConfigurationVariablesCollationServerEnum = "KOI8U_BIN"
	ConfigurationVariablesCollationServerLatin1SwedishCi       ConfigurationVariablesCollationServerEnum = "LATIN1_SWEDISH_CI"
	ConfigurationVariablesCollationServerLatin1Bin             ConfigurationVariablesCollationServerEnum = "LATIN1_BIN"
	ConfigurationVariablesCollationServerLatin1DanishCi        ConfigurationVariablesCollationServerEnum = "LATIN1_DANISH_CI"
	ConfigurationVariablesCollationServerLatin1GeneralCi       ConfigurationVariablesCollationServerEnum = "LATIN1_GENERAL_CI"
	ConfigurationVariablesCollationServerLatin1GeneralCs       ConfigurationVariablesCollationServerEnum = "LATIN1_GENERAL_CS"
	ConfigurationVariablesCollationServerLatin1German1Ci       ConfigurationVariablesCollationServerEnum = "LATIN1_GERMAN1_CI"
	ConfigurationVariablesCollationServerLatin1German2Ci       ConfigurationVariablesCollationServerEnum = "LATIN1_GERMAN2_CI"
	ConfigurationVariablesCollationServerLatin1SpanishCi       ConfigurationVariablesCollationServerEnum = "LATIN1_SPANISH_CI"
	ConfigurationVariablesCollationServerLatin2GeneralCi       ConfigurationVariablesCollationServerEnum = "LATIN2_GENERAL_CI"
	ConfigurationVariablesCollationServerLatin2Bin             ConfigurationVariablesCollationServerEnum = "LATIN2_BIN"
	ConfigurationVariablesCollationServerLatin2CroatianCi      ConfigurationVariablesCollationServerEnum = "LATIN2_CROATIAN_CI"
	ConfigurationVariablesCollationServerLatin2CzechCs         ConfigurationVariablesCollationServerEnum = "LATIN2_CZECH_CS"
	ConfigurationVariablesCollationServerLatin2HungarianCi     ConfigurationVariablesCollationServerEnum = "LATIN2_HUNGARIAN_CI"
	ConfigurationVariablesCollationServerLatin5TurkishCi       ConfigurationVariablesCollationServerEnum = "LATIN5_TURKISH_CI"
	ConfigurationVariablesCollationServerLatin5Bin             ConfigurationVariablesCollationServerEnum = "LATIN5_BIN"
	ConfigurationVariablesCollationServerLatin7GeneralCi       ConfigurationVariablesCollationServerEnum = "LATIN7_GENERAL_CI"
	ConfigurationVariablesCollationServerLatin7Bin             ConfigurationVariablesCollationServerEnum = "LATIN7_BIN"
	ConfigurationVariablesCollationServerLatin7EstonianCs      ConfigurationVariablesCollationServerEnum = "LATIN7_ESTONIAN_CS"
	ConfigurationVariablesCollationServerLatin7GeneralCs       ConfigurationVariablesCollationServerEnum = "LATIN7_GENERAL_CS"
	ConfigurationVariablesCollationServerMacceGeneralCi        ConfigurationVariablesCollationServerEnum = "MACCE_GENERAL_CI"
	ConfigurationVariablesCollationServerMacceBin              ConfigurationVariablesCollationServerEnum = "MACCE_BIN"
	ConfigurationVariablesCollationServerMacromanGeneralCi     ConfigurationVariablesCollationServerEnum = "MACROMAN_GENERAL_CI"
	ConfigurationVariablesCollationServerMacromanBin           ConfigurationVariablesCollationServerEnum = "MACROMAN_BIN"
	ConfigurationVariablesCollationServerSjisJapaneseCi        ConfigurationVariablesCollationServerEnum = "SJIS_JAPANESE_CI"
	ConfigurationVariablesCollationServerSjisBin               ConfigurationVariablesCollationServerEnum = "SJIS_BIN"
	ConfigurationVariablesCollationServerSwe7SwedishCi         ConfigurationVariablesCollationServerEnum = "SWE7_SWEDISH_CI"
	ConfigurationVariablesCollationServerSwe7Bin               ConfigurationVariablesCollationServerEnum = "SWE7_BIN"
	ConfigurationVariablesCollationServerTis620ThaiCi          ConfigurationVariablesCollationServerEnum = "TIS620_THAI_CI"
	ConfigurationVariablesCollationServerTis620Bin             ConfigurationVariablesCollationServerEnum = "TIS620_BIN"
	ConfigurationVariablesCollationServerUcs2GeneralCi         ConfigurationVariablesCollationServerEnum = "UCS2_GENERAL_CI"
	ConfigurationVariablesCollationServerUcs2Bin               ConfigurationVariablesCollationServerEnum = "UCS2_BIN"
	ConfigurationVariablesCollationServerUcs2CroatianCi        ConfigurationVariablesCollationServerEnum = "UCS2_CROATIAN_CI"
	ConfigurationVariablesCollationServerUcs2CzechCi           ConfigurationVariablesCollationServerEnum = "UCS2_CZECH_CI"
	ConfigurationVariablesCollationServerUcs2DanishCi          ConfigurationVariablesCollationServerEnum = "UCS2_DANISH_CI"
	ConfigurationVariablesCollationServerUcs2EsperantoCi       ConfigurationVariablesCollationServerEnum = "UCS2_ESPERANTO_CI"
	ConfigurationVariablesCollationServerUcs2EstonianCi        ConfigurationVariablesCollationServerEnum = "UCS2_ESTONIAN_CI"
	ConfigurationVariablesCollationServerUcs2GeneralMysql500Ci ConfigurationVariablesCollationServerEnum = "UCS2_GENERAL_MYSQL500_CI"
	ConfigurationVariablesCollationServerUcs2German2Ci         ConfigurationVariablesCollationServerEnum = "UCS2_GERMAN2_CI"
	ConfigurationVariablesCollationServerUcs2HungarianCi       ConfigurationVariablesCollationServerEnum = "UCS2_HUNGARIAN_CI"
	ConfigurationVariablesCollationServerUcs2IcelandicCi       ConfigurationVariablesCollationServerEnum = "UCS2_ICELANDIC_CI"
	ConfigurationVariablesCollationServerUcs2LatvianCi         ConfigurationVariablesCollationServerEnum = "UCS2_LATVIAN_CI"
	ConfigurationVariablesCollationServerUcs2LithuanianCi      ConfigurationVariablesCollationServerEnum = "UCS2_LITHUANIAN_CI"
	ConfigurationVariablesCollationServerUcs2PersianCi         ConfigurationVariablesCollationServerEnum = "UCS2_PERSIAN_CI"
	ConfigurationVariablesCollationServerUcs2PolishCi          ConfigurationVariablesCollationServerEnum = "UCS2_POLISH_CI"
	ConfigurationVariablesCollationServerUcs2RomanianCi        ConfigurationVariablesCollationServerEnum = "UCS2_ROMANIAN_CI"
	ConfigurationVariablesCollationServerUcs2RomanCi           ConfigurationVariablesCollationServerEnum = "UCS2_ROMAN_CI"
	ConfigurationVariablesCollationServerUcs2SinhalaCi         ConfigurationVariablesCollationServerEnum = "UCS2_SINHALA_CI"
	ConfigurationVariablesCollationServerUcs2SlovakCi          ConfigurationVariablesCollationServerEnum = "UCS2_SLOVAK_CI"
	ConfigurationVariablesCollationServerUcs2SlovenianCi       ConfigurationVariablesCollationServerEnum = "UCS2_SLOVENIAN_CI"
	ConfigurationVariablesCollationServerUcs2Spanish2Ci        ConfigurationVariablesCollationServerEnum = "UCS2_SPANISH2_CI"
	ConfigurationVariablesCollationServerUcs2SpanishCi         ConfigurationVariablesCollationServerEnum = "UCS2_SPANISH_CI"
	ConfigurationVariablesCollationServerUcs2SwedishCi         ConfigurationVariablesCollationServerEnum = "UCS2_SWEDISH_CI"
	ConfigurationVariablesCollationServerUcs2TurkishCi         ConfigurationVariablesCollationServerEnum = "UCS2_TURKISH_CI"
	ConfigurationVariablesCollationServerUcs2Unicode520Ci      ConfigurationVariablesCollationServerEnum = "UCS2_UNICODE_520_CI"
	ConfigurationVariablesCollationServerUcs2UnicodeCi         ConfigurationVariablesCollationServerEnum = "UCS2_UNICODE_CI"
	ConfigurationVariablesCollationServerUcs2VietnameseCi      ConfigurationVariablesCollationServerEnum = "UCS2_VIETNAMESE_CI"
	ConfigurationVariablesCollationServerUjisJapaneseCi        ConfigurationVariablesCollationServerEnum = "UJIS_JAPANESE_CI"
	ConfigurationVariablesCollationServerUjisBin               ConfigurationVariablesCollationServerEnum = "UJIS_BIN"
	ConfigurationVariablesCollationServerUtf16GeneralCi        ConfigurationVariablesCollationServerEnum = "UTF16_GENERAL_CI"
	ConfigurationVariablesCollationServerUtf16Bin              ConfigurationVariablesCollationServerEnum = "UTF16_BIN"
	ConfigurationVariablesCollationServerUtf16CroatianCi       ConfigurationVariablesCollationServerEnum = "UTF16_CROATIAN_CI"
	ConfigurationVariablesCollationServerUtf16CzechCi          ConfigurationVariablesCollationServerEnum = "UTF16_CZECH_CI"
	ConfigurationVariablesCollationServerUtf16DanishCi         ConfigurationVariablesCollationServerEnum = "UTF16_DANISH_CI"
	ConfigurationVariablesCollationServerUtf16EsperantoCi      ConfigurationVariablesCollationServerEnum = "UTF16_ESPERANTO_CI"
	ConfigurationVariablesCollationServerUtf16EstonianCi       ConfigurationVariablesCollationServerEnum = "UTF16_ESTONIAN_CI"
	ConfigurationVariablesCollationServerUtf16German2Ci        ConfigurationVariablesCollationServerEnum = "UTF16_GERMAN2_CI"
	ConfigurationVariablesCollationServerUtf16HungarianCi      ConfigurationVariablesCollationServerEnum = "UTF16_HUNGARIAN_CI"
	ConfigurationVariablesCollationServerUtf16IcelandicCi      ConfigurationVariablesCollationServerEnum = "UTF16_ICELANDIC_CI"
	ConfigurationVariablesCollationServerUtf16LatvianCi        ConfigurationVariablesCollationServerEnum = "UTF16_LATVIAN_CI"
	ConfigurationVariablesCollationServerUtf16LithuanianCi     ConfigurationVariablesCollationServerEnum = "UTF16_LITHUANIAN_CI"
	ConfigurationVariablesCollationServerUtf16PersianCi        ConfigurationVariablesCollationServerEnum = "UTF16_PERSIAN_CI"
	ConfigurationVariablesCollationServerUtf16PolishCi         ConfigurationVariablesCollationServerEnum = "UTF16_POLISH_CI"
	ConfigurationVariablesCollationServerUtf16RomanianCi       ConfigurationVariablesCollationServerEnum = "UTF16_ROMANIAN_CI"
	ConfigurationVariablesCollationServerUtf16RomanCi          ConfigurationVariablesCollationServerEnum = "UTF16_ROMAN_CI"
	ConfigurationVariablesCollationServerUtf16SinhalaCi        ConfigurationVariablesCollationServerEnum = "UTF16_SINHALA_CI"
	ConfigurationVariablesCollationServerUtf16SlovakCi         ConfigurationVariablesCollationServerEnum = "UTF16_SLOVAK_CI"
	ConfigurationVariablesCollationServerUtf16SlovenianCi      ConfigurationVariablesCollationServerEnum = "UTF16_SLOVENIAN_CI"
	ConfigurationVariablesCollationServerUtf16Spanish2Ci       ConfigurationVariablesCollationServerEnum = "UTF16_SPANISH2_CI"
	ConfigurationVariablesCollationServerUtf16SpanishCi        ConfigurationVariablesCollationServerEnum = "UTF16_SPANISH_CI"
	ConfigurationVariablesCollationServerUtf16SwedishCi        ConfigurationVariablesCollationServerEnum = "UTF16_SWEDISH_CI"
	ConfigurationVariablesCollationServerUtf16TurkishCi        ConfigurationVariablesCollationServerEnum = "UTF16_TURKISH_CI"
	ConfigurationVariablesCollationServerUtf16Unicode520Ci     ConfigurationVariablesCollationServerEnum = "UTF16_UNICODE_520_CI"
	ConfigurationVariablesCollationServerUtf16UnicodeCi        ConfigurationVariablesCollationServerEnum = "UTF16_UNICODE_CI"
	ConfigurationVariablesCollationServerUtf16VietnameseCi     ConfigurationVariablesCollationServerEnum = "UTF16_VIETNAMESE_CI"
	ConfigurationVariablesCollationServerUtf16leGeneralCi      ConfigurationVariablesCollationServerEnum = "UTF16LE_GENERAL_CI"
	ConfigurationVariablesCollationServerUtf16leBin            ConfigurationVariablesCollationServerEnum = "UTF16LE_BIN"
	ConfigurationVariablesCollationServerUtf32GeneralCi        ConfigurationVariablesCollationServerEnum = "UTF32_GENERAL_CI"
	ConfigurationVariablesCollationServerUtf32Bin              ConfigurationVariablesCollationServerEnum = "UTF32_BIN"
	ConfigurationVariablesCollationServerUtf32CroatianCi       ConfigurationVariablesCollationServerEnum = "UTF32_CROATIAN_CI"
	ConfigurationVariablesCollationServerUtf32CzechCi          ConfigurationVariablesCollationServerEnum = "UTF32_CZECH_CI"
	ConfigurationVariablesCollationServerUtf32DanishCi         ConfigurationVariablesCollationServerEnum = "UTF32_DANISH_CI"
	ConfigurationVariablesCollationServerUtf32EsperantoCi      ConfigurationVariablesCollationServerEnum = "UTF32_ESPERANTO_CI"
	ConfigurationVariablesCollationServerUtf32EstonianCi       ConfigurationVariablesCollationServerEnum = "UTF32_ESTONIAN_CI"
	ConfigurationVariablesCollationServerUtf32German2Ci        ConfigurationVariablesCollationServerEnum = "UTF32_GERMAN2_CI"
	ConfigurationVariablesCollationServerUtf32HungarianCi      ConfigurationVariablesCollationServerEnum = "UTF32_HUNGARIAN_CI"
	ConfigurationVariablesCollationServerUtf32IcelandicCi      ConfigurationVariablesCollationServerEnum = "UTF32_ICELANDIC_CI"
	ConfigurationVariablesCollationServerUtf32LatvianCi        ConfigurationVariablesCollationServerEnum = "UTF32_LATVIAN_CI"
	ConfigurationVariablesCollationServerUtf32LithuanianCi     ConfigurationVariablesCollationServerEnum = "UTF32_LITHUANIAN_CI"
	ConfigurationVariablesCollationServerUtf32PersianCi        ConfigurationVariablesCollationServerEnum = "UTF32_PERSIAN_CI"
	ConfigurationVariablesCollationServerUtf32PolishCi         ConfigurationVariablesCollationServerEnum = "UTF32_POLISH_CI"
	ConfigurationVariablesCollationServerUtf32RomanianCi       ConfigurationVariablesCollationServerEnum = "UTF32_ROMANIAN_CI"
	ConfigurationVariablesCollationServerUtf32RomanCi          ConfigurationVariablesCollationServerEnum = "UTF32_ROMAN_CI"
	ConfigurationVariablesCollationServerUtf32SinhalaCi        ConfigurationVariablesCollationServerEnum = "UTF32_SINHALA_CI"
	ConfigurationVariablesCollationServerUtf32SlovakCi         ConfigurationVariablesCollationServerEnum = "UTF32_SLOVAK_CI"
	ConfigurationVariablesCollationServerUtf32SlovenianCi      ConfigurationVariablesCollationServerEnum = "UTF32_SLOVENIAN_CI"
	ConfigurationVariablesCollationServerUtf32Spanish2Ci       ConfigurationVariablesCollationServerEnum = "UTF32_SPANISH2_CI"
	ConfigurationVariablesCollationServerUtf32SpanishCi        ConfigurationVariablesCollationServerEnum = "UTF32_SPANISH_CI"
	ConfigurationVariablesCollationServerUtf32SwedishCi        ConfigurationVariablesCollationServerEnum = "UTF32_SWEDISH_CI"
	ConfigurationVariablesCollationServerUtf32TurkishCi        ConfigurationVariablesCollationServerEnum = "UTF32_TURKISH_CI"
	ConfigurationVariablesCollationServerUtf32Unicode520Ci     ConfigurationVariablesCollationServerEnum = "UTF32_UNICODE_520_CI"
	ConfigurationVariablesCollationServerUtf32UnicodeCi        ConfigurationVariablesCollationServerEnum = "UTF32_UNICODE_CI"
	ConfigurationVariablesCollationServerUtf32VietnameseCi     ConfigurationVariablesCollationServerEnum = "UTF32_VIETNAMESE_CI"
	ConfigurationVariablesCollationServerUtf8GeneralCi         ConfigurationVariablesCollationServerEnum = "UTF8_GENERAL_CI"
	ConfigurationVariablesCollationServerUtf8Bin               ConfigurationVariablesCollationServerEnum = "UTF8_BIN"
	ConfigurationVariablesCollationServerUtf8CroatianCi        ConfigurationVariablesCollationServerEnum = "UTF8_CROATIAN_CI"
	ConfigurationVariablesCollationServerUtf8CzechCi           ConfigurationVariablesCollationServerEnum = "UTF8_CZECH_CI"
	ConfigurationVariablesCollationServerUtf8DanishCi          ConfigurationVariablesCollationServerEnum = "UTF8_DANISH_CI"
	ConfigurationVariablesCollationServerUtf8EsperantoCi       ConfigurationVariablesCollationServerEnum = "UTF8_ESPERANTO_CI"
	ConfigurationVariablesCollationServerUtf8EstonianCi        ConfigurationVariablesCollationServerEnum = "UTF8_ESTONIAN_CI"
	ConfigurationVariablesCollationServerUtf8GeneralMysql500Ci ConfigurationVariablesCollationServerEnum = "UTF8_GENERAL_MYSQL500_CI"
	ConfigurationVariablesCollationServerUtf8German2Ci         ConfigurationVariablesCollationServerEnum = "UTF8_GERMAN2_CI"
	ConfigurationVariablesCollationServerUtf8HungarianCi       ConfigurationVariablesCollationServerEnum = "UTF8_HUNGARIAN_CI"
	ConfigurationVariablesCollationServerUtf8IcelandicCi       ConfigurationVariablesCollationServerEnum = "UTF8_ICELANDIC_CI"
	ConfigurationVariablesCollationServerUtf8LatvianCi         ConfigurationVariablesCollationServerEnum = "UTF8_LATVIAN_CI"
	ConfigurationVariablesCollationServerUtf8LithuanianCi      ConfigurationVariablesCollationServerEnum = "UTF8_LITHUANIAN_CI"
	ConfigurationVariablesCollationServerUtf8PersianCi         ConfigurationVariablesCollationServerEnum = "UTF8_PERSIAN_CI"
	ConfigurationVariablesCollationServerUtf8PolishCi          ConfigurationVariablesCollationServerEnum = "UTF8_POLISH_CI"
	ConfigurationVariablesCollationServerUtf8RomanianCi        ConfigurationVariablesCollationServerEnum = "UTF8_ROMANIAN_CI"
	ConfigurationVariablesCollationServerUtf8RomanCi           ConfigurationVariablesCollationServerEnum = "UTF8_ROMAN_CI"
	ConfigurationVariablesCollationServerUtf8SinhalaCi         ConfigurationVariablesCollationServerEnum = "UTF8_SINHALA_CI"
	ConfigurationVariablesCollationServerUtf8SlovakCi          ConfigurationVariablesCollationServerEnum = "UTF8_SLOVAK_CI"
	ConfigurationVariablesCollationServerUtf8SlovenianCi       ConfigurationVariablesCollationServerEnum = "UTF8_SLOVENIAN_CI"
	ConfigurationVariablesCollationServerUtf8Spanish2Ci        ConfigurationVariablesCollationServerEnum = "UTF8_SPANISH2_CI"
	ConfigurationVariablesCollationServerUtf8SpanishCi         ConfigurationVariablesCollationServerEnum = "UTF8_SPANISH_CI"
	ConfigurationVariablesCollationServerUtf8SwedishCi         ConfigurationVariablesCollationServerEnum = "UTF8_SWEDISH_CI"
	ConfigurationVariablesCollationServerUtf8TolowerCi         ConfigurationVariablesCollationServerEnum = "UTF8_TOLOWER_CI"
	ConfigurationVariablesCollationServerUtf8TurkishCi         ConfigurationVariablesCollationServerEnum = "UTF8_TURKISH_CI"
	ConfigurationVariablesCollationServerUtf8Unicode520Ci      ConfigurationVariablesCollationServerEnum = "UTF8_UNICODE_520_CI"
	ConfigurationVariablesCollationServerUtf8UnicodeCi         ConfigurationVariablesCollationServerEnum = "UTF8_UNICODE_CI"
	ConfigurationVariablesCollationServerUtf8VietnameseCi      ConfigurationVariablesCollationServerEnum = "UTF8_VIETNAMESE_CI"
	ConfigurationVariablesCollationServerUtf8mb40900AiCi       ConfigurationVariablesCollationServerEnum = "UTF8MB4_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb40900AsCi       ConfigurationVariablesCollationServerEnum = "UTF8MB4_0900_AS_CI"
	ConfigurationVariablesCollationServerUtf8mb40900AsCs       ConfigurationVariablesCollationServerEnum = "UTF8MB4_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb40900Bin        ConfigurationVariablesCollationServerEnum = "UTF8MB4_0900_BIN"
	ConfigurationVariablesCollationServerUtf8mb4Bin            ConfigurationVariablesCollationServerEnum = "UTF8MB4_BIN"
	ConfigurationVariablesCollationServerUtf8mb4CroatianCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_CROATIAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4Cs0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_CS_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Cs0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_CS_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4CzechCi        ConfigurationVariablesCollationServerEnum = "UTF8MB4_CZECH_CI"
	ConfigurationVariablesCollationServerUtf8mb4DanishCi       ConfigurationVariablesCollationServerEnum = "UTF8MB4_DANISH_CI"
	ConfigurationVariablesCollationServerUtf8mb4Da0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_DA_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Da0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_DA_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4DePb0900AiCi   ConfigurationVariablesCollationServerEnum = "UTF8MB4_DE_PB_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4DePb0900AsCs   ConfigurationVariablesCollationServerEnum = "UTF8MB4_DE_PB_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4Eo0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_EO_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Eo0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_EO_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4EsperantoCi    ConfigurationVariablesCollationServerEnum = "UTF8MB4_ESPERANTO_CI"
	ConfigurationVariablesCollationServerUtf8mb4EstonianCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_ESTONIAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4Es0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_ES_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Es0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_ES_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4EsTrad0900AiCi ConfigurationVariablesCollationServerEnum = "UTF8MB4_ES_TRAD_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4EsTrad0900AsCs ConfigurationVariablesCollationServerEnum = "UTF8MB4_ES_TRAD_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4Et0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_ET_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Et0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_ET_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4GeneralCi      ConfigurationVariablesCollationServerEnum = "UTF8MB4_GENERAL_CI"
	ConfigurationVariablesCollationServerUtf8mb4German2Ci      ConfigurationVariablesCollationServerEnum = "UTF8MB4_GERMAN2_CI"
	ConfigurationVariablesCollationServerUtf8mb4Hr0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_HR_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Hr0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_HR_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4HungarianCi    ConfigurationVariablesCollationServerEnum = "UTF8MB4_HUNGARIAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4Hu0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_HU_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Hu0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_HU_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4IcelandicCi    ConfigurationVariablesCollationServerEnum = "UTF8MB4_ICELANDIC_CI"
	ConfigurationVariablesCollationServerUtf8mb4Is0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_IS_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Is0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_IS_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4Ja0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_JA_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4Ja0900AsCsKs   ConfigurationVariablesCollationServerEnum = "UTF8MB4_JA_0900_AS_CS_KS"
	ConfigurationVariablesCollationServerUtf8mb4LatvianCi      ConfigurationVariablesCollationServerEnum = "UTF8MB4_LATVIAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4La0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_LA_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4La0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_LA_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4LithuanianCi   ConfigurationVariablesCollationServerEnum = "UTF8MB4_LITHUANIAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4Lt0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_LT_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Lt0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_LT_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4Lv0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_LV_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Lv0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_LV_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4PersianCi      ConfigurationVariablesCollationServerEnum = "UTF8MB4_PERSIAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4Pl0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_PL_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Pl0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_PL_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4PolishCi       ConfigurationVariablesCollationServerEnum = "UTF8MB4_POLISH_CI"
	ConfigurationVariablesCollationServerUtf8mb4RomanianCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_ROMANIAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4RomanCi        ConfigurationVariablesCollationServerEnum = "UTF8MB4_ROMAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4Ro0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_RO_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Ro0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_RO_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4Ru0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_RU_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Ru0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_RU_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4SinhalaCi      ConfigurationVariablesCollationServerEnum = "UTF8MB4_SINHALA_CI"
	ConfigurationVariablesCollationServerUtf8mb4Sk0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_SK_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Sk0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_SK_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4SlovakCi       ConfigurationVariablesCollationServerEnum = "UTF8MB4_SLOVAK_CI"
	ConfigurationVariablesCollationServerUtf8mb4SlovenianCi    ConfigurationVariablesCollationServerEnum = "UTF8MB4_SLOVENIAN_CI"
	ConfigurationVariablesCollationServerUtf8mb4Sl0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_SL_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Sl0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_SL_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4Spanish2Ci     ConfigurationVariablesCollationServerEnum = "UTF8MB4_SPANISH2_CI"
	ConfigurationVariablesCollationServerUtf8mb4SpanishCi      ConfigurationVariablesCollationServerEnum = "UTF8MB4_SPANISH_CI"
	ConfigurationVariablesCollationServerUtf8mb4Sv0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_SV_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Sv0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_SV_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4SwedishCi      ConfigurationVariablesCollationServerEnum = "UTF8MB4_SWEDISH_CI"
	ConfigurationVariablesCollationServerUtf8mb4Tr0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_TR_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Tr0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_TR_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4TurkishCi      ConfigurationVariablesCollationServerEnum = "UTF8MB4_TURKISH_CI"
	ConfigurationVariablesCollationServerUtf8mb4Unicode520Ci   ConfigurationVariablesCollationServerEnum = "UTF8MB4_UNICODE_520_CI"
	ConfigurationVariablesCollationServerUtf8mb4UnicodeCi      ConfigurationVariablesCollationServerEnum = "UTF8MB4_UNICODE_CI"
	ConfigurationVariablesCollationServerUtf8mb4VietnameseCi   ConfigurationVariablesCollationServerEnum = "UTF8MB4_VIETNAMESE_CI"
	ConfigurationVariablesCollationServerUtf8mb4Vi0900AiCi     ConfigurationVariablesCollationServerEnum = "UTF8MB4_VI_0900_AI_CI"
	ConfigurationVariablesCollationServerUtf8mb4Vi0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_VI_0900_AS_CS"
	ConfigurationVariablesCollationServerUtf8mb4Zh0900AsCs     ConfigurationVariablesCollationServerEnum = "UTF8MB4_ZH_0900_AS_CS"
)

var mappingConfigurationVariablesCollationServerEnum = map[string]ConfigurationVariablesCollationServerEnum{
	"ARMSCII8_GENERAL_CI":        ConfigurationVariablesCollationServerArmscii8GeneralCi,
	"ARMSCII8_BIN":               ConfigurationVariablesCollationServerArmscii8Bin,
	"ASCII_GENERAL_CI":           ConfigurationVariablesCollationServerAsciiGeneralCi,
	"ASCII_BIN":                  ConfigurationVariablesCollationServerAsciiBin,
	"BIG5_CHINESE_CI":            ConfigurationVariablesCollationServerBig5ChineseCi,
	"BIG5_BIN":                   ConfigurationVariablesCollationServerBig5Bin,
	"BINARY":                     ConfigurationVariablesCollationServerBinary,
	"CP1250_GENERAL_CI":          ConfigurationVariablesCollationServerCp1250GeneralCi,
	"CP1250_BIN":                 ConfigurationVariablesCollationServerCp1250Bin,
	"CP1250_CROATIAN_CI":         ConfigurationVariablesCollationServerCp1250CroatianCi,
	"CP1250_CZECH_CS":            ConfigurationVariablesCollationServerCp1250CzechCs,
	"CP1250_POLISH_CI":           ConfigurationVariablesCollationServerCp1250PolishCi,
	"CP1251_GENERAL_CI":          ConfigurationVariablesCollationServerCp1251GeneralCi,
	"CP1251_BIN":                 ConfigurationVariablesCollationServerCp1251Bin,
	"CP1251_BULGARIAN_CI":        ConfigurationVariablesCollationServerCp1251BulgarianCi,
	"CP1251_GENERAL_CS":          ConfigurationVariablesCollationServerCp1251GeneralCs,
	"CP1251_UKRAINIAN_CI":        ConfigurationVariablesCollationServerCp1251UkrainianCi,
	"CP1256_GENERAL_CI":          ConfigurationVariablesCollationServerCp1256GeneralCi,
	"CP1256_BIN":                 ConfigurationVariablesCollationServerCp1256Bin,
	"CP1257_GENERAL_CI":          ConfigurationVariablesCollationServerCp1257GeneralCi,
	"CP1257_BIN":                 ConfigurationVariablesCollationServerCp1257Bin,
	"CP1257_LITHUANIAN_CI":       ConfigurationVariablesCollationServerCp1257LithuanianCi,
	"CP850_GENERAL_CI":           ConfigurationVariablesCollationServerCp850GeneralCi,
	"CP850_BIN":                  ConfigurationVariablesCollationServerCp850Bin,
	"CP852_GENERAL_CI":           ConfigurationVariablesCollationServerCp852GeneralCi,
	"CP852_BIN":                  ConfigurationVariablesCollationServerCp852Bin,
	"CP866_GENERAL_CI":           ConfigurationVariablesCollationServerCp866GeneralCi,
	"CP866_BIN":                  ConfigurationVariablesCollationServerCp866Bin,
	"CP932_JAPANESE_CI":          ConfigurationVariablesCollationServerCp932JapaneseCi,
	"CP932_BIN":                  ConfigurationVariablesCollationServerCp932Bin,
	"DEC8_SWEDISH_CI":            ConfigurationVariablesCollationServerDec8SwedishCi,
	"DEC8_BIN":                   ConfigurationVariablesCollationServerDec8Bin,
	"EUCJPMS_JAPANESE_CI":        ConfigurationVariablesCollationServerEucjpmsJapaneseCi,
	"EUCJPMS_BIN":                ConfigurationVariablesCollationServerEucjpmsBin,
	"EUCKR_KOREAN_CI":            ConfigurationVariablesCollationServerEuckrKoreanCi,
	"EUCKR_BIN":                  ConfigurationVariablesCollationServerEuckrBin,
	"GB18030_CHINESE_CI":         ConfigurationVariablesCollationServerGb18030ChineseCi,
	"GB18030_BIN":                ConfigurationVariablesCollationServerGb18030Bin,
	"GB18030_UNICODE_520_CI":     ConfigurationVariablesCollationServerGb18030Unicode520Ci,
	"GB2312_CHINESE_CI":          ConfigurationVariablesCollationServerGb2312ChineseCi,
	"GB2312_BIN":                 ConfigurationVariablesCollationServerGb2312Bin,
	"GBK_CHINESE_CI":             ConfigurationVariablesCollationServerGbkChineseCi,
	"GBK_BIN":                    ConfigurationVariablesCollationServerGbkBin,
	"GEOSTD8_GENERAL_CI":         ConfigurationVariablesCollationServerGeostd8GeneralCi,
	"GEOSTD8_BIN":                ConfigurationVariablesCollationServerGeostd8Bin,
	"GREEK_GENERAL_CI":           ConfigurationVariablesCollationServerGreekGeneralCi,
	"GREEK_BIN":                  ConfigurationVariablesCollationServerGreekBin,
	"HEBREW_GENERAL_CI":          ConfigurationVariablesCollationServerHebrewGeneralCi,
	"HEBREW_BIN":                 ConfigurationVariablesCollationServerHebrewBin,
	"HP8_ENGLISH_CI":             ConfigurationVariablesCollationServerHp8EnglishCi,
	"HP8_BIN":                    ConfigurationVariablesCollationServerHp8Bin,
	"KEYBCS2_GENERAL_CI":         ConfigurationVariablesCollationServerKeybcs2GeneralCi,
	"KEYBCS2_BIN":                ConfigurationVariablesCollationServerKeybcs2Bin,
	"KOI8R_GENERAL_CI":           ConfigurationVariablesCollationServerKoi8rGeneralCi,
	"KOI8R_BIN":                  ConfigurationVariablesCollationServerKoi8rBin,
	"KOI8U_GENERAL_CI":           ConfigurationVariablesCollationServerKoi8uGeneralCi,
	"KOI8U_BIN":                  ConfigurationVariablesCollationServerKoi8uBin,
	"LATIN1_SWEDISH_CI":          ConfigurationVariablesCollationServerLatin1SwedishCi,
	"LATIN1_BIN":                 ConfigurationVariablesCollationServerLatin1Bin,
	"LATIN1_DANISH_CI":           ConfigurationVariablesCollationServerLatin1DanishCi,
	"LATIN1_GENERAL_CI":          ConfigurationVariablesCollationServerLatin1GeneralCi,
	"LATIN1_GENERAL_CS":          ConfigurationVariablesCollationServerLatin1GeneralCs,
	"LATIN1_GERMAN1_CI":          ConfigurationVariablesCollationServerLatin1German1Ci,
	"LATIN1_GERMAN2_CI":          ConfigurationVariablesCollationServerLatin1German2Ci,
	"LATIN1_SPANISH_CI":          ConfigurationVariablesCollationServerLatin1SpanishCi,
	"LATIN2_GENERAL_CI":          ConfigurationVariablesCollationServerLatin2GeneralCi,
	"LATIN2_BIN":                 ConfigurationVariablesCollationServerLatin2Bin,
	"LATIN2_CROATIAN_CI":         ConfigurationVariablesCollationServerLatin2CroatianCi,
	"LATIN2_CZECH_CS":            ConfigurationVariablesCollationServerLatin2CzechCs,
	"LATIN2_HUNGARIAN_CI":        ConfigurationVariablesCollationServerLatin2HungarianCi,
	"LATIN5_TURKISH_CI":          ConfigurationVariablesCollationServerLatin5TurkishCi,
	"LATIN5_BIN":                 ConfigurationVariablesCollationServerLatin5Bin,
	"LATIN7_GENERAL_CI":          ConfigurationVariablesCollationServerLatin7GeneralCi,
	"LATIN7_BIN":                 ConfigurationVariablesCollationServerLatin7Bin,
	"LATIN7_ESTONIAN_CS":         ConfigurationVariablesCollationServerLatin7EstonianCs,
	"LATIN7_GENERAL_CS":          ConfigurationVariablesCollationServerLatin7GeneralCs,
	"MACCE_GENERAL_CI":           ConfigurationVariablesCollationServerMacceGeneralCi,
	"MACCE_BIN":                  ConfigurationVariablesCollationServerMacceBin,
	"MACROMAN_GENERAL_CI":        ConfigurationVariablesCollationServerMacromanGeneralCi,
	"MACROMAN_BIN":               ConfigurationVariablesCollationServerMacromanBin,
	"SJIS_JAPANESE_CI":           ConfigurationVariablesCollationServerSjisJapaneseCi,
	"SJIS_BIN":                   ConfigurationVariablesCollationServerSjisBin,
	"SWE7_SWEDISH_CI":            ConfigurationVariablesCollationServerSwe7SwedishCi,
	"SWE7_BIN":                   ConfigurationVariablesCollationServerSwe7Bin,
	"TIS620_THAI_CI":             ConfigurationVariablesCollationServerTis620ThaiCi,
	"TIS620_BIN":                 ConfigurationVariablesCollationServerTis620Bin,
	"UCS2_GENERAL_CI":            ConfigurationVariablesCollationServerUcs2GeneralCi,
	"UCS2_BIN":                   ConfigurationVariablesCollationServerUcs2Bin,
	"UCS2_CROATIAN_CI":           ConfigurationVariablesCollationServerUcs2CroatianCi,
	"UCS2_CZECH_CI":              ConfigurationVariablesCollationServerUcs2CzechCi,
	"UCS2_DANISH_CI":             ConfigurationVariablesCollationServerUcs2DanishCi,
	"UCS2_ESPERANTO_CI":          ConfigurationVariablesCollationServerUcs2EsperantoCi,
	"UCS2_ESTONIAN_CI":           ConfigurationVariablesCollationServerUcs2EstonianCi,
	"UCS2_GENERAL_MYSQL500_CI":   ConfigurationVariablesCollationServerUcs2GeneralMysql500Ci,
	"UCS2_GERMAN2_CI":            ConfigurationVariablesCollationServerUcs2German2Ci,
	"UCS2_HUNGARIAN_CI":          ConfigurationVariablesCollationServerUcs2HungarianCi,
	"UCS2_ICELANDIC_CI":          ConfigurationVariablesCollationServerUcs2IcelandicCi,
	"UCS2_LATVIAN_CI":            ConfigurationVariablesCollationServerUcs2LatvianCi,
	"UCS2_LITHUANIAN_CI":         ConfigurationVariablesCollationServerUcs2LithuanianCi,
	"UCS2_PERSIAN_CI":            ConfigurationVariablesCollationServerUcs2PersianCi,
	"UCS2_POLISH_CI":             ConfigurationVariablesCollationServerUcs2PolishCi,
	"UCS2_ROMANIAN_CI":           ConfigurationVariablesCollationServerUcs2RomanianCi,
	"UCS2_ROMAN_CI":              ConfigurationVariablesCollationServerUcs2RomanCi,
	"UCS2_SINHALA_CI":            ConfigurationVariablesCollationServerUcs2SinhalaCi,
	"UCS2_SLOVAK_CI":             ConfigurationVariablesCollationServerUcs2SlovakCi,
	"UCS2_SLOVENIAN_CI":          ConfigurationVariablesCollationServerUcs2SlovenianCi,
	"UCS2_SPANISH2_CI":           ConfigurationVariablesCollationServerUcs2Spanish2Ci,
	"UCS2_SPANISH_CI":            ConfigurationVariablesCollationServerUcs2SpanishCi,
	"UCS2_SWEDISH_CI":            ConfigurationVariablesCollationServerUcs2SwedishCi,
	"UCS2_TURKISH_CI":            ConfigurationVariablesCollationServerUcs2TurkishCi,
	"UCS2_UNICODE_520_CI":        ConfigurationVariablesCollationServerUcs2Unicode520Ci,
	"UCS2_UNICODE_CI":            ConfigurationVariablesCollationServerUcs2UnicodeCi,
	"UCS2_VIETNAMESE_CI":         ConfigurationVariablesCollationServerUcs2VietnameseCi,
	"UJIS_JAPANESE_CI":           ConfigurationVariablesCollationServerUjisJapaneseCi,
	"UJIS_BIN":                   ConfigurationVariablesCollationServerUjisBin,
	"UTF16_GENERAL_CI":           ConfigurationVariablesCollationServerUtf16GeneralCi,
	"UTF16_BIN":                  ConfigurationVariablesCollationServerUtf16Bin,
	"UTF16_CROATIAN_CI":          ConfigurationVariablesCollationServerUtf16CroatianCi,
	"UTF16_CZECH_CI":             ConfigurationVariablesCollationServerUtf16CzechCi,
	"UTF16_DANISH_CI":            ConfigurationVariablesCollationServerUtf16DanishCi,
	"UTF16_ESPERANTO_CI":         ConfigurationVariablesCollationServerUtf16EsperantoCi,
	"UTF16_ESTONIAN_CI":          ConfigurationVariablesCollationServerUtf16EstonianCi,
	"UTF16_GERMAN2_CI":           ConfigurationVariablesCollationServerUtf16German2Ci,
	"UTF16_HUNGARIAN_CI":         ConfigurationVariablesCollationServerUtf16HungarianCi,
	"UTF16_ICELANDIC_CI":         ConfigurationVariablesCollationServerUtf16IcelandicCi,
	"UTF16_LATVIAN_CI":           ConfigurationVariablesCollationServerUtf16LatvianCi,
	"UTF16_LITHUANIAN_CI":        ConfigurationVariablesCollationServerUtf16LithuanianCi,
	"UTF16_PERSIAN_CI":           ConfigurationVariablesCollationServerUtf16PersianCi,
	"UTF16_POLISH_CI":            ConfigurationVariablesCollationServerUtf16PolishCi,
	"UTF16_ROMANIAN_CI":          ConfigurationVariablesCollationServerUtf16RomanianCi,
	"UTF16_ROMAN_CI":             ConfigurationVariablesCollationServerUtf16RomanCi,
	"UTF16_SINHALA_CI":           ConfigurationVariablesCollationServerUtf16SinhalaCi,
	"UTF16_SLOVAK_CI":            ConfigurationVariablesCollationServerUtf16SlovakCi,
	"UTF16_SLOVENIAN_CI":         ConfigurationVariablesCollationServerUtf16SlovenianCi,
	"UTF16_SPANISH2_CI":          ConfigurationVariablesCollationServerUtf16Spanish2Ci,
	"UTF16_SPANISH_CI":           ConfigurationVariablesCollationServerUtf16SpanishCi,
	"UTF16_SWEDISH_CI":           ConfigurationVariablesCollationServerUtf16SwedishCi,
	"UTF16_TURKISH_CI":           ConfigurationVariablesCollationServerUtf16TurkishCi,
	"UTF16_UNICODE_520_CI":       ConfigurationVariablesCollationServerUtf16Unicode520Ci,
	"UTF16_UNICODE_CI":           ConfigurationVariablesCollationServerUtf16UnicodeCi,
	"UTF16_VIETNAMESE_CI":        ConfigurationVariablesCollationServerUtf16VietnameseCi,
	"UTF16LE_GENERAL_CI":         ConfigurationVariablesCollationServerUtf16leGeneralCi,
	"UTF16LE_BIN":                ConfigurationVariablesCollationServerUtf16leBin,
	"UTF32_GENERAL_CI":           ConfigurationVariablesCollationServerUtf32GeneralCi,
	"UTF32_BIN":                  ConfigurationVariablesCollationServerUtf32Bin,
	"UTF32_CROATIAN_CI":          ConfigurationVariablesCollationServerUtf32CroatianCi,
	"UTF32_CZECH_CI":             ConfigurationVariablesCollationServerUtf32CzechCi,
	"UTF32_DANISH_CI":            ConfigurationVariablesCollationServerUtf32DanishCi,
	"UTF32_ESPERANTO_CI":         ConfigurationVariablesCollationServerUtf32EsperantoCi,
	"UTF32_ESTONIAN_CI":          ConfigurationVariablesCollationServerUtf32EstonianCi,
	"UTF32_GERMAN2_CI":           ConfigurationVariablesCollationServerUtf32German2Ci,
	"UTF32_HUNGARIAN_CI":         ConfigurationVariablesCollationServerUtf32HungarianCi,
	"UTF32_ICELANDIC_CI":         ConfigurationVariablesCollationServerUtf32IcelandicCi,
	"UTF32_LATVIAN_CI":           ConfigurationVariablesCollationServerUtf32LatvianCi,
	"UTF32_LITHUANIAN_CI":        ConfigurationVariablesCollationServerUtf32LithuanianCi,
	"UTF32_PERSIAN_CI":           ConfigurationVariablesCollationServerUtf32PersianCi,
	"UTF32_POLISH_CI":            ConfigurationVariablesCollationServerUtf32PolishCi,
	"UTF32_ROMANIAN_CI":          ConfigurationVariablesCollationServerUtf32RomanianCi,
	"UTF32_ROMAN_CI":             ConfigurationVariablesCollationServerUtf32RomanCi,
	"UTF32_SINHALA_CI":           ConfigurationVariablesCollationServerUtf32SinhalaCi,
	"UTF32_SLOVAK_CI":            ConfigurationVariablesCollationServerUtf32SlovakCi,
	"UTF32_SLOVENIAN_CI":         ConfigurationVariablesCollationServerUtf32SlovenianCi,
	"UTF32_SPANISH2_CI":          ConfigurationVariablesCollationServerUtf32Spanish2Ci,
	"UTF32_SPANISH_CI":           ConfigurationVariablesCollationServerUtf32SpanishCi,
	"UTF32_SWEDISH_CI":           ConfigurationVariablesCollationServerUtf32SwedishCi,
	"UTF32_TURKISH_CI":           ConfigurationVariablesCollationServerUtf32TurkishCi,
	"UTF32_UNICODE_520_CI":       ConfigurationVariablesCollationServerUtf32Unicode520Ci,
	"UTF32_UNICODE_CI":           ConfigurationVariablesCollationServerUtf32UnicodeCi,
	"UTF32_VIETNAMESE_CI":        ConfigurationVariablesCollationServerUtf32VietnameseCi,
	"UTF8_GENERAL_CI":            ConfigurationVariablesCollationServerUtf8GeneralCi,
	"UTF8_BIN":                   ConfigurationVariablesCollationServerUtf8Bin,
	"UTF8_CROATIAN_CI":           ConfigurationVariablesCollationServerUtf8CroatianCi,
	"UTF8_CZECH_CI":              ConfigurationVariablesCollationServerUtf8CzechCi,
	"UTF8_DANISH_CI":             ConfigurationVariablesCollationServerUtf8DanishCi,
	"UTF8_ESPERANTO_CI":          ConfigurationVariablesCollationServerUtf8EsperantoCi,
	"UTF8_ESTONIAN_CI":           ConfigurationVariablesCollationServerUtf8EstonianCi,
	"UTF8_GENERAL_MYSQL500_CI":   ConfigurationVariablesCollationServerUtf8GeneralMysql500Ci,
	"UTF8_GERMAN2_CI":            ConfigurationVariablesCollationServerUtf8German2Ci,
	"UTF8_HUNGARIAN_CI":          ConfigurationVariablesCollationServerUtf8HungarianCi,
	"UTF8_ICELANDIC_CI":          ConfigurationVariablesCollationServerUtf8IcelandicCi,
	"UTF8_LATVIAN_CI":            ConfigurationVariablesCollationServerUtf8LatvianCi,
	"UTF8_LITHUANIAN_CI":         ConfigurationVariablesCollationServerUtf8LithuanianCi,
	"UTF8_PERSIAN_CI":            ConfigurationVariablesCollationServerUtf8PersianCi,
	"UTF8_POLISH_CI":             ConfigurationVariablesCollationServerUtf8PolishCi,
	"UTF8_ROMANIAN_CI":           ConfigurationVariablesCollationServerUtf8RomanianCi,
	"UTF8_ROMAN_CI":              ConfigurationVariablesCollationServerUtf8RomanCi,
	"UTF8_SINHALA_CI":            ConfigurationVariablesCollationServerUtf8SinhalaCi,
	"UTF8_SLOVAK_CI":             ConfigurationVariablesCollationServerUtf8SlovakCi,
	"UTF8_SLOVENIAN_CI":          ConfigurationVariablesCollationServerUtf8SlovenianCi,
	"UTF8_SPANISH2_CI":           ConfigurationVariablesCollationServerUtf8Spanish2Ci,
	"UTF8_SPANISH_CI":            ConfigurationVariablesCollationServerUtf8SpanishCi,
	"UTF8_SWEDISH_CI":            ConfigurationVariablesCollationServerUtf8SwedishCi,
	"UTF8_TOLOWER_CI":            ConfigurationVariablesCollationServerUtf8TolowerCi,
	"UTF8_TURKISH_CI":            ConfigurationVariablesCollationServerUtf8TurkishCi,
	"UTF8_UNICODE_520_CI":        ConfigurationVariablesCollationServerUtf8Unicode520Ci,
	"UTF8_UNICODE_CI":            ConfigurationVariablesCollationServerUtf8UnicodeCi,
	"UTF8_VIETNAMESE_CI":         ConfigurationVariablesCollationServerUtf8VietnameseCi,
	"UTF8MB4_0900_AI_CI":         ConfigurationVariablesCollationServerUtf8mb40900AiCi,
	"UTF8MB4_0900_AS_CI":         ConfigurationVariablesCollationServerUtf8mb40900AsCi,
	"UTF8MB4_0900_AS_CS":         ConfigurationVariablesCollationServerUtf8mb40900AsCs,
	"UTF8MB4_0900_BIN":           ConfigurationVariablesCollationServerUtf8mb40900Bin,
	"UTF8MB4_BIN":                ConfigurationVariablesCollationServerUtf8mb4Bin,
	"UTF8MB4_CROATIAN_CI":        ConfigurationVariablesCollationServerUtf8mb4CroatianCi,
	"UTF8MB4_CS_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Cs0900AiCi,
	"UTF8MB4_CS_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Cs0900AsCs,
	"UTF8MB4_CZECH_CI":           ConfigurationVariablesCollationServerUtf8mb4CzechCi,
	"UTF8MB4_DANISH_CI":          ConfigurationVariablesCollationServerUtf8mb4DanishCi,
	"UTF8MB4_DA_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Da0900AiCi,
	"UTF8MB4_DA_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Da0900AsCs,
	"UTF8MB4_DE_PB_0900_AI_CI":   ConfigurationVariablesCollationServerUtf8mb4DePb0900AiCi,
	"UTF8MB4_DE_PB_0900_AS_CS":   ConfigurationVariablesCollationServerUtf8mb4DePb0900AsCs,
	"UTF8MB4_EO_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Eo0900AiCi,
	"UTF8MB4_EO_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Eo0900AsCs,
	"UTF8MB4_ESPERANTO_CI":       ConfigurationVariablesCollationServerUtf8mb4EsperantoCi,
	"UTF8MB4_ESTONIAN_CI":        ConfigurationVariablesCollationServerUtf8mb4EstonianCi,
	"UTF8MB4_ES_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Es0900AiCi,
	"UTF8MB4_ES_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Es0900AsCs,
	"UTF8MB4_ES_TRAD_0900_AI_CI": ConfigurationVariablesCollationServerUtf8mb4EsTrad0900AiCi,
	"UTF8MB4_ES_TRAD_0900_AS_CS": ConfigurationVariablesCollationServerUtf8mb4EsTrad0900AsCs,
	"UTF8MB4_ET_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Et0900AiCi,
	"UTF8MB4_ET_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Et0900AsCs,
	"UTF8MB4_GENERAL_CI":         ConfigurationVariablesCollationServerUtf8mb4GeneralCi,
	"UTF8MB4_GERMAN2_CI":         ConfigurationVariablesCollationServerUtf8mb4German2Ci,
	"UTF8MB4_HR_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Hr0900AiCi,
	"UTF8MB4_HR_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Hr0900AsCs,
	"UTF8MB4_HUNGARIAN_CI":       ConfigurationVariablesCollationServerUtf8mb4HungarianCi,
	"UTF8MB4_HU_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Hu0900AiCi,
	"UTF8MB4_HU_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Hu0900AsCs,
	"UTF8MB4_ICELANDIC_CI":       ConfigurationVariablesCollationServerUtf8mb4IcelandicCi,
	"UTF8MB4_IS_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Is0900AiCi,
	"UTF8MB4_IS_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Is0900AsCs,
	"UTF8MB4_JA_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Ja0900AsCs,
	"UTF8MB4_JA_0900_AS_CS_KS":   ConfigurationVariablesCollationServerUtf8mb4Ja0900AsCsKs,
	"UTF8MB4_LATVIAN_CI":         ConfigurationVariablesCollationServerUtf8mb4LatvianCi,
	"UTF8MB4_LA_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4La0900AiCi,
	"UTF8MB4_LA_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4La0900AsCs,
	"UTF8MB4_LITHUANIAN_CI":      ConfigurationVariablesCollationServerUtf8mb4LithuanianCi,
	"UTF8MB4_LT_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Lt0900AiCi,
	"UTF8MB4_LT_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Lt0900AsCs,
	"UTF8MB4_LV_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Lv0900AiCi,
	"UTF8MB4_LV_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Lv0900AsCs,
	"UTF8MB4_PERSIAN_CI":         ConfigurationVariablesCollationServerUtf8mb4PersianCi,
	"UTF8MB4_PL_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Pl0900AiCi,
	"UTF8MB4_PL_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Pl0900AsCs,
	"UTF8MB4_POLISH_CI":          ConfigurationVariablesCollationServerUtf8mb4PolishCi,
	"UTF8MB4_ROMANIAN_CI":        ConfigurationVariablesCollationServerUtf8mb4RomanianCi,
	"UTF8MB4_ROMAN_CI":           ConfigurationVariablesCollationServerUtf8mb4RomanCi,
	"UTF8MB4_RO_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Ro0900AiCi,
	"UTF8MB4_RO_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Ro0900AsCs,
	"UTF8MB4_RU_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Ru0900AiCi,
	"UTF8MB4_RU_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Ru0900AsCs,
	"UTF8MB4_SINHALA_CI":         ConfigurationVariablesCollationServerUtf8mb4SinhalaCi,
	"UTF8MB4_SK_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Sk0900AiCi,
	"UTF8MB4_SK_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Sk0900AsCs,
	"UTF8MB4_SLOVAK_CI":          ConfigurationVariablesCollationServerUtf8mb4SlovakCi,
	"UTF8MB4_SLOVENIAN_CI":       ConfigurationVariablesCollationServerUtf8mb4SlovenianCi,
	"UTF8MB4_SL_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Sl0900AiCi,
	"UTF8MB4_SL_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Sl0900AsCs,
	"UTF8MB4_SPANISH2_CI":        ConfigurationVariablesCollationServerUtf8mb4Spanish2Ci,
	"UTF8MB4_SPANISH_CI":         ConfigurationVariablesCollationServerUtf8mb4SpanishCi,
	"UTF8MB4_SV_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Sv0900AiCi,
	"UTF8MB4_SV_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Sv0900AsCs,
	"UTF8MB4_SWEDISH_CI":         ConfigurationVariablesCollationServerUtf8mb4SwedishCi,
	"UTF8MB4_TR_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Tr0900AiCi,
	"UTF8MB4_TR_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Tr0900AsCs,
	"UTF8MB4_TURKISH_CI":         ConfigurationVariablesCollationServerUtf8mb4TurkishCi,
	"UTF8MB4_UNICODE_520_CI":     ConfigurationVariablesCollationServerUtf8mb4Unicode520Ci,
	"UTF8MB4_UNICODE_CI":         ConfigurationVariablesCollationServerUtf8mb4UnicodeCi,
	"UTF8MB4_VIETNAMESE_CI":      ConfigurationVariablesCollationServerUtf8mb4VietnameseCi,
	"UTF8MB4_VI_0900_AI_CI":      ConfigurationVariablesCollationServerUtf8mb4Vi0900AiCi,
	"UTF8MB4_VI_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Vi0900AsCs,
	"UTF8MB4_ZH_0900_AS_CS":      ConfigurationVariablesCollationServerUtf8mb4Zh0900AsCs,
}

var mappingConfigurationVariablesCollationServerEnumLowerCase = map[string]ConfigurationVariablesCollationServerEnum{
	"armscii8_general_ci":        ConfigurationVariablesCollationServerArmscii8GeneralCi,
	"armscii8_bin":               ConfigurationVariablesCollationServerArmscii8Bin,
	"ascii_general_ci":           ConfigurationVariablesCollationServerAsciiGeneralCi,
	"ascii_bin":                  ConfigurationVariablesCollationServerAsciiBin,
	"big5_chinese_ci":            ConfigurationVariablesCollationServerBig5ChineseCi,
	"big5_bin":                   ConfigurationVariablesCollationServerBig5Bin,
	"binary":                     ConfigurationVariablesCollationServerBinary,
	"cp1250_general_ci":          ConfigurationVariablesCollationServerCp1250GeneralCi,
	"cp1250_bin":                 ConfigurationVariablesCollationServerCp1250Bin,
	"cp1250_croatian_ci":         ConfigurationVariablesCollationServerCp1250CroatianCi,
	"cp1250_czech_cs":            ConfigurationVariablesCollationServerCp1250CzechCs,
	"cp1250_polish_ci":           ConfigurationVariablesCollationServerCp1250PolishCi,
	"cp1251_general_ci":          ConfigurationVariablesCollationServerCp1251GeneralCi,
	"cp1251_bin":                 ConfigurationVariablesCollationServerCp1251Bin,
	"cp1251_bulgarian_ci":        ConfigurationVariablesCollationServerCp1251BulgarianCi,
	"cp1251_general_cs":          ConfigurationVariablesCollationServerCp1251GeneralCs,
	"cp1251_ukrainian_ci":        ConfigurationVariablesCollationServerCp1251UkrainianCi,
	"cp1256_general_ci":          ConfigurationVariablesCollationServerCp1256GeneralCi,
	"cp1256_bin":                 ConfigurationVariablesCollationServerCp1256Bin,
	"cp1257_general_ci":          ConfigurationVariablesCollationServerCp1257GeneralCi,
	"cp1257_bin":                 ConfigurationVariablesCollationServerCp1257Bin,
	"cp1257_lithuanian_ci":       ConfigurationVariablesCollationServerCp1257LithuanianCi,
	"cp850_general_ci":           ConfigurationVariablesCollationServerCp850GeneralCi,
	"cp850_bin":                  ConfigurationVariablesCollationServerCp850Bin,
	"cp852_general_ci":           ConfigurationVariablesCollationServerCp852GeneralCi,
	"cp852_bin":                  ConfigurationVariablesCollationServerCp852Bin,
	"cp866_general_ci":           ConfigurationVariablesCollationServerCp866GeneralCi,
	"cp866_bin":                  ConfigurationVariablesCollationServerCp866Bin,
	"cp932_japanese_ci":          ConfigurationVariablesCollationServerCp932JapaneseCi,
	"cp932_bin":                  ConfigurationVariablesCollationServerCp932Bin,
	"dec8_swedish_ci":            ConfigurationVariablesCollationServerDec8SwedishCi,
	"dec8_bin":                   ConfigurationVariablesCollationServerDec8Bin,
	"eucjpms_japanese_ci":        ConfigurationVariablesCollationServerEucjpmsJapaneseCi,
	"eucjpms_bin":                ConfigurationVariablesCollationServerEucjpmsBin,
	"euckr_korean_ci":            ConfigurationVariablesCollationServerEuckrKoreanCi,
	"euckr_bin":                  ConfigurationVariablesCollationServerEuckrBin,
	"gb18030_chinese_ci":         ConfigurationVariablesCollationServerGb18030ChineseCi,
	"gb18030_bin":                ConfigurationVariablesCollationServerGb18030Bin,
	"gb18030_unicode_520_ci":     ConfigurationVariablesCollationServerGb18030Unicode520Ci,
	"gb2312_chinese_ci":          ConfigurationVariablesCollationServerGb2312ChineseCi,
	"gb2312_bin":                 ConfigurationVariablesCollationServerGb2312Bin,
	"gbk_chinese_ci":             ConfigurationVariablesCollationServerGbkChineseCi,
	"gbk_bin":                    ConfigurationVariablesCollationServerGbkBin,
	"geostd8_general_ci":         ConfigurationVariablesCollationServerGeostd8GeneralCi,
	"geostd8_bin":                ConfigurationVariablesCollationServerGeostd8Bin,
	"greek_general_ci":           ConfigurationVariablesCollationServerGreekGeneralCi,
	"greek_bin":                  ConfigurationVariablesCollationServerGreekBin,
	"hebrew_general_ci":          ConfigurationVariablesCollationServerHebrewGeneralCi,
	"hebrew_bin":                 ConfigurationVariablesCollationServerHebrewBin,
	"hp8_english_ci":             ConfigurationVariablesCollationServerHp8EnglishCi,
	"hp8_bin":                    ConfigurationVariablesCollationServerHp8Bin,
	"keybcs2_general_ci":         ConfigurationVariablesCollationServerKeybcs2GeneralCi,
	"keybcs2_bin":                ConfigurationVariablesCollationServerKeybcs2Bin,
	"koi8r_general_ci":           ConfigurationVariablesCollationServerKoi8rGeneralCi,
	"koi8r_bin":                  ConfigurationVariablesCollationServerKoi8rBin,
	"koi8u_general_ci":           ConfigurationVariablesCollationServerKoi8uGeneralCi,
	"koi8u_bin":                  ConfigurationVariablesCollationServerKoi8uBin,
	"latin1_swedish_ci":          ConfigurationVariablesCollationServerLatin1SwedishCi,
	"latin1_bin":                 ConfigurationVariablesCollationServerLatin1Bin,
	"latin1_danish_ci":           ConfigurationVariablesCollationServerLatin1DanishCi,
	"latin1_general_ci":          ConfigurationVariablesCollationServerLatin1GeneralCi,
	"latin1_general_cs":          ConfigurationVariablesCollationServerLatin1GeneralCs,
	"latin1_german1_ci":          ConfigurationVariablesCollationServerLatin1German1Ci,
	"latin1_german2_ci":          ConfigurationVariablesCollationServerLatin1German2Ci,
	"latin1_spanish_ci":          ConfigurationVariablesCollationServerLatin1SpanishCi,
	"latin2_general_ci":          ConfigurationVariablesCollationServerLatin2GeneralCi,
	"latin2_bin":                 ConfigurationVariablesCollationServerLatin2Bin,
	"latin2_croatian_ci":         ConfigurationVariablesCollationServerLatin2CroatianCi,
	"latin2_czech_cs":            ConfigurationVariablesCollationServerLatin2CzechCs,
	"latin2_hungarian_ci":        ConfigurationVariablesCollationServerLatin2HungarianCi,
	"latin5_turkish_ci":          ConfigurationVariablesCollationServerLatin5TurkishCi,
	"latin5_bin":                 ConfigurationVariablesCollationServerLatin5Bin,
	"latin7_general_ci":          ConfigurationVariablesCollationServerLatin7GeneralCi,
	"latin7_bin":                 ConfigurationVariablesCollationServerLatin7Bin,
	"latin7_estonian_cs":         ConfigurationVariablesCollationServerLatin7EstonianCs,
	"latin7_general_cs":          ConfigurationVariablesCollationServerLatin7GeneralCs,
	"macce_general_ci":           ConfigurationVariablesCollationServerMacceGeneralCi,
	"macce_bin":                  ConfigurationVariablesCollationServerMacceBin,
	"macroman_general_ci":        ConfigurationVariablesCollationServerMacromanGeneralCi,
	"macroman_bin":               ConfigurationVariablesCollationServerMacromanBin,
	"sjis_japanese_ci":           ConfigurationVariablesCollationServerSjisJapaneseCi,
	"sjis_bin":                   ConfigurationVariablesCollationServerSjisBin,
	"swe7_swedish_ci":            ConfigurationVariablesCollationServerSwe7SwedishCi,
	"swe7_bin":                   ConfigurationVariablesCollationServerSwe7Bin,
	"tis620_thai_ci":             ConfigurationVariablesCollationServerTis620ThaiCi,
	"tis620_bin":                 ConfigurationVariablesCollationServerTis620Bin,
	"ucs2_general_ci":            ConfigurationVariablesCollationServerUcs2GeneralCi,
	"ucs2_bin":                   ConfigurationVariablesCollationServerUcs2Bin,
	"ucs2_croatian_ci":           ConfigurationVariablesCollationServerUcs2CroatianCi,
	"ucs2_czech_ci":              ConfigurationVariablesCollationServerUcs2CzechCi,
	"ucs2_danish_ci":             ConfigurationVariablesCollationServerUcs2DanishCi,
	"ucs2_esperanto_ci":          ConfigurationVariablesCollationServerUcs2EsperantoCi,
	"ucs2_estonian_ci":           ConfigurationVariablesCollationServerUcs2EstonianCi,
	"ucs2_general_mysql500_ci":   ConfigurationVariablesCollationServerUcs2GeneralMysql500Ci,
	"ucs2_german2_ci":            ConfigurationVariablesCollationServerUcs2German2Ci,
	"ucs2_hungarian_ci":          ConfigurationVariablesCollationServerUcs2HungarianCi,
	"ucs2_icelandic_ci":          ConfigurationVariablesCollationServerUcs2IcelandicCi,
	"ucs2_latvian_ci":            ConfigurationVariablesCollationServerUcs2LatvianCi,
	"ucs2_lithuanian_ci":         ConfigurationVariablesCollationServerUcs2LithuanianCi,
	"ucs2_persian_ci":            ConfigurationVariablesCollationServerUcs2PersianCi,
	"ucs2_polish_ci":             ConfigurationVariablesCollationServerUcs2PolishCi,
	"ucs2_romanian_ci":           ConfigurationVariablesCollationServerUcs2RomanianCi,
	"ucs2_roman_ci":              ConfigurationVariablesCollationServerUcs2RomanCi,
	"ucs2_sinhala_ci":            ConfigurationVariablesCollationServerUcs2SinhalaCi,
	"ucs2_slovak_ci":             ConfigurationVariablesCollationServerUcs2SlovakCi,
	"ucs2_slovenian_ci":          ConfigurationVariablesCollationServerUcs2SlovenianCi,
	"ucs2_spanish2_ci":           ConfigurationVariablesCollationServerUcs2Spanish2Ci,
	"ucs2_spanish_ci":            ConfigurationVariablesCollationServerUcs2SpanishCi,
	"ucs2_swedish_ci":            ConfigurationVariablesCollationServerUcs2SwedishCi,
	"ucs2_turkish_ci":            ConfigurationVariablesCollationServerUcs2TurkishCi,
	"ucs2_unicode_520_ci":        ConfigurationVariablesCollationServerUcs2Unicode520Ci,
	"ucs2_unicode_ci":            ConfigurationVariablesCollationServerUcs2UnicodeCi,
	"ucs2_vietnamese_ci":         ConfigurationVariablesCollationServerUcs2VietnameseCi,
	"ujis_japanese_ci":           ConfigurationVariablesCollationServerUjisJapaneseCi,
	"ujis_bin":                   ConfigurationVariablesCollationServerUjisBin,
	"utf16_general_ci":           ConfigurationVariablesCollationServerUtf16GeneralCi,
	"utf16_bin":                  ConfigurationVariablesCollationServerUtf16Bin,
	"utf16_croatian_ci":          ConfigurationVariablesCollationServerUtf16CroatianCi,
	"utf16_czech_ci":             ConfigurationVariablesCollationServerUtf16CzechCi,
	"utf16_danish_ci":            ConfigurationVariablesCollationServerUtf16DanishCi,
	"utf16_esperanto_ci":         ConfigurationVariablesCollationServerUtf16EsperantoCi,
	"utf16_estonian_ci":          ConfigurationVariablesCollationServerUtf16EstonianCi,
	"utf16_german2_ci":           ConfigurationVariablesCollationServerUtf16German2Ci,
	"utf16_hungarian_ci":         ConfigurationVariablesCollationServerUtf16HungarianCi,
	"utf16_icelandic_ci":         ConfigurationVariablesCollationServerUtf16IcelandicCi,
	"utf16_latvian_ci":           ConfigurationVariablesCollationServerUtf16LatvianCi,
	"utf16_lithuanian_ci":        ConfigurationVariablesCollationServerUtf16LithuanianCi,
	"utf16_persian_ci":           ConfigurationVariablesCollationServerUtf16PersianCi,
	"utf16_polish_ci":            ConfigurationVariablesCollationServerUtf16PolishCi,
	"utf16_romanian_ci":          ConfigurationVariablesCollationServerUtf16RomanianCi,
	"utf16_roman_ci":             ConfigurationVariablesCollationServerUtf16RomanCi,
	"utf16_sinhala_ci":           ConfigurationVariablesCollationServerUtf16SinhalaCi,
	"utf16_slovak_ci":            ConfigurationVariablesCollationServerUtf16SlovakCi,
	"utf16_slovenian_ci":         ConfigurationVariablesCollationServerUtf16SlovenianCi,
	"utf16_spanish2_ci":          ConfigurationVariablesCollationServerUtf16Spanish2Ci,
	"utf16_spanish_ci":           ConfigurationVariablesCollationServerUtf16SpanishCi,
	"utf16_swedish_ci":           ConfigurationVariablesCollationServerUtf16SwedishCi,
	"utf16_turkish_ci":           ConfigurationVariablesCollationServerUtf16TurkishCi,
	"utf16_unicode_520_ci":       ConfigurationVariablesCollationServerUtf16Unicode520Ci,
	"utf16_unicode_ci":           ConfigurationVariablesCollationServerUtf16UnicodeCi,
	"utf16_vietnamese_ci":        ConfigurationVariablesCollationServerUtf16VietnameseCi,
	"utf16le_general_ci":         ConfigurationVariablesCollationServerUtf16leGeneralCi,
	"utf16le_bin":                ConfigurationVariablesCollationServerUtf16leBin,
	"utf32_general_ci":           ConfigurationVariablesCollationServerUtf32GeneralCi,
	"utf32_bin":                  ConfigurationVariablesCollationServerUtf32Bin,
	"utf32_croatian_ci":          ConfigurationVariablesCollationServerUtf32CroatianCi,
	"utf32_czech_ci":             ConfigurationVariablesCollationServerUtf32CzechCi,
	"utf32_danish_ci":            ConfigurationVariablesCollationServerUtf32DanishCi,
	"utf32_esperanto_ci":         ConfigurationVariablesCollationServerUtf32EsperantoCi,
	"utf32_estonian_ci":          ConfigurationVariablesCollationServerUtf32EstonianCi,
	"utf32_german2_ci":           ConfigurationVariablesCollationServerUtf32German2Ci,
	"utf32_hungarian_ci":         ConfigurationVariablesCollationServerUtf32HungarianCi,
	"utf32_icelandic_ci":         ConfigurationVariablesCollationServerUtf32IcelandicCi,
	"utf32_latvian_ci":           ConfigurationVariablesCollationServerUtf32LatvianCi,
	"utf32_lithuanian_ci":        ConfigurationVariablesCollationServerUtf32LithuanianCi,
	"utf32_persian_ci":           ConfigurationVariablesCollationServerUtf32PersianCi,
	"utf32_polish_ci":            ConfigurationVariablesCollationServerUtf32PolishCi,
	"utf32_romanian_ci":          ConfigurationVariablesCollationServerUtf32RomanianCi,
	"utf32_roman_ci":             ConfigurationVariablesCollationServerUtf32RomanCi,
	"utf32_sinhala_ci":           ConfigurationVariablesCollationServerUtf32SinhalaCi,
	"utf32_slovak_ci":            ConfigurationVariablesCollationServerUtf32SlovakCi,
	"utf32_slovenian_ci":         ConfigurationVariablesCollationServerUtf32SlovenianCi,
	"utf32_spanish2_ci":          ConfigurationVariablesCollationServerUtf32Spanish2Ci,
	"utf32_spanish_ci":           ConfigurationVariablesCollationServerUtf32SpanishCi,
	"utf32_swedish_ci":           ConfigurationVariablesCollationServerUtf32SwedishCi,
	"utf32_turkish_ci":           ConfigurationVariablesCollationServerUtf32TurkishCi,
	"utf32_unicode_520_ci":       ConfigurationVariablesCollationServerUtf32Unicode520Ci,
	"utf32_unicode_ci":           ConfigurationVariablesCollationServerUtf32UnicodeCi,
	"utf32_vietnamese_ci":        ConfigurationVariablesCollationServerUtf32VietnameseCi,
	"utf8_general_ci":            ConfigurationVariablesCollationServerUtf8GeneralCi,
	"utf8_bin":                   ConfigurationVariablesCollationServerUtf8Bin,
	"utf8_croatian_ci":           ConfigurationVariablesCollationServerUtf8CroatianCi,
	"utf8_czech_ci":              ConfigurationVariablesCollationServerUtf8CzechCi,
	"utf8_danish_ci":             ConfigurationVariablesCollationServerUtf8DanishCi,
	"utf8_esperanto_ci":          ConfigurationVariablesCollationServerUtf8EsperantoCi,
	"utf8_estonian_ci":           ConfigurationVariablesCollationServerUtf8EstonianCi,
	"utf8_general_mysql500_ci":   ConfigurationVariablesCollationServerUtf8GeneralMysql500Ci,
	"utf8_german2_ci":            ConfigurationVariablesCollationServerUtf8German2Ci,
	"utf8_hungarian_ci":          ConfigurationVariablesCollationServerUtf8HungarianCi,
	"utf8_icelandic_ci":          ConfigurationVariablesCollationServerUtf8IcelandicCi,
	"utf8_latvian_ci":            ConfigurationVariablesCollationServerUtf8LatvianCi,
	"utf8_lithuanian_ci":         ConfigurationVariablesCollationServerUtf8LithuanianCi,
	"utf8_persian_ci":            ConfigurationVariablesCollationServerUtf8PersianCi,
	"utf8_polish_ci":             ConfigurationVariablesCollationServerUtf8PolishCi,
	"utf8_romanian_ci":           ConfigurationVariablesCollationServerUtf8RomanianCi,
	"utf8_roman_ci":              ConfigurationVariablesCollationServerUtf8RomanCi,
	"utf8_sinhala_ci":            ConfigurationVariablesCollationServerUtf8SinhalaCi,
	"utf8_slovak_ci":             ConfigurationVariablesCollationServerUtf8SlovakCi,
	"utf8_slovenian_ci":          ConfigurationVariablesCollationServerUtf8SlovenianCi,
	"utf8_spanish2_ci":           ConfigurationVariablesCollationServerUtf8Spanish2Ci,
	"utf8_spanish_ci":            ConfigurationVariablesCollationServerUtf8SpanishCi,
	"utf8_swedish_ci":            ConfigurationVariablesCollationServerUtf8SwedishCi,
	"utf8_tolower_ci":            ConfigurationVariablesCollationServerUtf8TolowerCi,
	"utf8_turkish_ci":            ConfigurationVariablesCollationServerUtf8TurkishCi,
	"utf8_unicode_520_ci":        ConfigurationVariablesCollationServerUtf8Unicode520Ci,
	"utf8_unicode_ci":            ConfigurationVariablesCollationServerUtf8UnicodeCi,
	"utf8_vietnamese_ci":         ConfigurationVariablesCollationServerUtf8VietnameseCi,
	"utf8mb4_0900_ai_ci":         ConfigurationVariablesCollationServerUtf8mb40900AiCi,
	"utf8mb4_0900_as_ci":         ConfigurationVariablesCollationServerUtf8mb40900AsCi,
	"utf8mb4_0900_as_cs":         ConfigurationVariablesCollationServerUtf8mb40900AsCs,
	"utf8mb4_0900_bin":           ConfigurationVariablesCollationServerUtf8mb40900Bin,
	"utf8mb4_bin":                ConfigurationVariablesCollationServerUtf8mb4Bin,
	"utf8mb4_croatian_ci":        ConfigurationVariablesCollationServerUtf8mb4CroatianCi,
	"utf8mb4_cs_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Cs0900AiCi,
	"utf8mb4_cs_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Cs0900AsCs,
	"utf8mb4_czech_ci":           ConfigurationVariablesCollationServerUtf8mb4CzechCi,
	"utf8mb4_danish_ci":          ConfigurationVariablesCollationServerUtf8mb4DanishCi,
	"utf8mb4_da_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Da0900AiCi,
	"utf8mb4_da_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Da0900AsCs,
	"utf8mb4_de_pb_0900_ai_ci":   ConfigurationVariablesCollationServerUtf8mb4DePb0900AiCi,
	"utf8mb4_de_pb_0900_as_cs":   ConfigurationVariablesCollationServerUtf8mb4DePb0900AsCs,
	"utf8mb4_eo_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Eo0900AiCi,
	"utf8mb4_eo_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Eo0900AsCs,
	"utf8mb4_esperanto_ci":       ConfigurationVariablesCollationServerUtf8mb4EsperantoCi,
	"utf8mb4_estonian_ci":        ConfigurationVariablesCollationServerUtf8mb4EstonianCi,
	"utf8mb4_es_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Es0900AiCi,
	"utf8mb4_es_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Es0900AsCs,
	"utf8mb4_es_trad_0900_ai_ci": ConfigurationVariablesCollationServerUtf8mb4EsTrad0900AiCi,
	"utf8mb4_es_trad_0900_as_cs": ConfigurationVariablesCollationServerUtf8mb4EsTrad0900AsCs,
	"utf8mb4_et_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Et0900AiCi,
	"utf8mb4_et_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Et0900AsCs,
	"utf8mb4_general_ci":         ConfigurationVariablesCollationServerUtf8mb4GeneralCi,
	"utf8mb4_german2_ci":         ConfigurationVariablesCollationServerUtf8mb4German2Ci,
	"utf8mb4_hr_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Hr0900AiCi,
	"utf8mb4_hr_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Hr0900AsCs,
	"utf8mb4_hungarian_ci":       ConfigurationVariablesCollationServerUtf8mb4HungarianCi,
	"utf8mb4_hu_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Hu0900AiCi,
	"utf8mb4_hu_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Hu0900AsCs,
	"utf8mb4_icelandic_ci":       ConfigurationVariablesCollationServerUtf8mb4IcelandicCi,
	"utf8mb4_is_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Is0900AiCi,
	"utf8mb4_is_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Is0900AsCs,
	"utf8mb4_ja_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Ja0900AsCs,
	"utf8mb4_ja_0900_as_cs_ks":   ConfigurationVariablesCollationServerUtf8mb4Ja0900AsCsKs,
	"utf8mb4_latvian_ci":         ConfigurationVariablesCollationServerUtf8mb4LatvianCi,
	"utf8mb4_la_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4La0900AiCi,
	"utf8mb4_la_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4La0900AsCs,
	"utf8mb4_lithuanian_ci":      ConfigurationVariablesCollationServerUtf8mb4LithuanianCi,
	"utf8mb4_lt_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Lt0900AiCi,
	"utf8mb4_lt_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Lt0900AsCs,
	"utf8mb4_lv_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Lv0900AiCi,
	"utf8mb4_lv_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Lv0900AsCs,
	"utf8mb4_persian_ci":         ConfigurationVariablesCollationServerUtf8mb4PersianCi,
	"utf8mb4_pl_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Pl0900AiCi,
	"utf8mb4_pl_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Pl0900AsCs,
	"utf8mb4_polish_ci":          ConfigurationVariablesCollationServerUtf8mb4PolishCi,
	"utf8mb4_romanian_ci":        ConfigurationVariablesCollationServerUtf8mb4RomanianCi,
	"utf8mb4_roman_ci":           ConfigurationVariablesCollationServerUtf8mb4RomanCi,
	"utf8mb4_ro_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Ro0900AiCi,
	"utf8mb4_ro_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Ro0900AsCs,
	"utf8mb4_ru_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Ru0900AiCi,
	"utf8mb4_ru_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Ru0900AsCs,
	"utf8mb4_sinhala_ci":         ConfigurationVariablesCollationServerUtf8mb4SinhalaCi,
	"utf8mb4_sk_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Sk0900AiCi,
	"utf8mb4_sk_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Sk0900AsCs,
	"utf8mb4_slovak_ci":          ConfigurationVariablesCollationServerUtf8mb4SlovakCi,
	"utf8mb4_slovenian_ci":       ConfigurationVariablesCollationServerUtf8mb4SlovenianCi,
	"utf8mb4_sl_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Sl0900AiCi,
	"utf8mb4_sl_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Sl0900AsCs,
	"utf8mb4_spanish2_ci":        ConfigurationVariablesCollationServerUtf8mb4Spanish2Ci,
	"utf8mb4_spanish_ci":         ConfigurationVariablesCollationServerUtf8mb4SpanishCi,
	"utf8mb4_sv_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Sv0900AiCi,
	"utf8mb4_sv_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Sv0900AsCs,
	"utf8mb4_swedish_ci":         ConfigurationVariablesCollationServerUtf8mb4SwedishCi,
	"utf8mb4_tr_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Tr0900AiCi,
	"utf8mb4_tr_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Tr0900AsCs,
	"utf8mb4_turkish_ci":         ConfigurationVariablesCollationServerUtf8mb4TurkishCi,
	"utf8mb4_unicode_520_ci":     ConfigurationVariablesCollationServerUtf8mb4Unicode520Ci,
	"utf8mb4_unicode_ci":         ConfigurationVariablesCollationServerUtf8mb4UnicodeCi,
	"utf8mb4_vietnamese_ci":      ConfigurationVariablesCollationServerUtf8mb4VietnameseCi,
	"utf8mb4_vi_0900_ai_ci":      ConfigurationVariablesCollationServerUtf8mb4Vi0900AiCi,
	"utf8mb4_vi_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Vi0900AsCs,
	"utf8mb4_zh_0900_as_cs":      ConfigurationVariablesCollationServerUtf8mb4Zh0900AsCs,
}

// GetConfigurationVariablesCollationServerEnumValues Enumerates the set of values for ConfigurationVariablesCollationServerEnum
func GetConfigurationVariablesCollationServerEnumValues() []ConfigurationVariablesCollationServerEnum {
	values := make([]ConfigurationVariablesCollationServerEnum, 0)
	for _, v := range mappingConfigurationVariablesCollationServerEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesCollationServerEnumStringValues Enumerates the set of values in String for ConfigurationVariablesCollationServerEnum
func GetConfigurationVariablesCollationServerEnumStringValues() []string {
	return []string{
		"ARMSCII8_GENERAL_CI",
		"ARMSCII8_BIN",
		"ASCII_GENERAL_CI",
		"ASCII_BIN",
		"BIG5_CHINESE_CI",
		"BIG5_BIN",
		"BINARY",
		"CP1250_GENERAL_CI",
		"CP1250_BIN",
		"CP1250_CROATIAN_CI",
		"CP1250_CZECH_CS",
		"CP1250_POLISH_CI",
		"CP1251_GENERAL_CI",
		"CP1251_BIN",
		"CP1251_BULGARIAN_CI",
		"CP1251_GENERAL_CS",
		"CP1251_UKRAINIAN_CI",
		"CP1256_GENERAL_CI",
		"CP1256_BIN",
		"CP1257_GENERAL_CI",
		"CP1257_BIN",
		"CP1257_LITHUANIAN_CI",
		"CP850_GENERAL_CI",
		"CP850_BIN",
		"CP852_GENERAL_CI",
		"CP852_BIN",
		"CP866_GENERAL_CI",
		"CP866_BIN",
		"CP932_JAPANESE_CI",
		"CP932_BIN",
		"DEC8_SWEDISH_CI",
		"DEC8_BIN",
		"EUCJPMS_JAPANESE_CI",
		"EUCJPMS_BIN",
		"EUCKR_KOREAN_CI",
		"EUCKR_BIN",
		"GB18030_CHINESE_CI",
		"GB18030_BIN",
		"GB18030_UNICODE_520_CI",
		"GB2312_CHINESE_CI",
		"GB2312_BIN",
		"GBK_CHINESE_CI",
		"GBK_BIN",
		"GEOSTD8_GENERAL_CI",
		"GEOSTD8_BIN",
		"GREEK_GENERAL_CI",
		"GREEK_BIN",
		"HEBREW_GENERAL_CI",
		"HEBREW_BIN",
		"HP8_ENGLISH_CI",
		"HP8_BIN",
		"KEYBCS2_GENERAL_CI",
		"KEYBCS2_BIN",
		"KOI8R_GENERAL_CI",
		"KOI8R_BIN",
		"KOI8U_GENERAL_CI",
		"KOI8U_BIN",
		"LATIN1_SWEDISH_CI",
		"LATIN1_BIN",
		"LATIN1_DANISH_CI",
		"LATIN1_GENERAL_CI",
		"LATIN1_GENERAL_CS",
		"LATIN1_GERMAN1_CI",
		"LATIN1_GERMAN2_CI",
		"LATIN1_SPANISH_CI",
		"LATIN2_GENERAL_CI",
		"LATIN2_BIN",
		"LATIN2_CROATIAN_CI",
		"LATIN2_CZECH_CS",
		"LATIN2_HUNGARIAN_CI",
		"LATIN5_TURKISH_CI",
		"LATIN5_BIN",
		"LATIN7_GENERAL_CI",
		"LATIN7_BIN",
		"LATIN7_ESTONIAN_CS",
		"LATIN7_GENERAL_CS",
		"MACCE_GENERAL_CI",
		"MACCE_BIN",
		"MACROMAN_GENERAL_CI",
		"MACROMAN_BIN",
		"SJIS_JAPANESE_CI",
		"SJIS_BIN",
		"SWE7_SWEDISH_CI",
		"SWE7_BIN",
		"TIS620_THAI_CI",
		"TIS620_BIN",
		"UCS2_GENERAL_CI",
		"UCS2_BIN",
		"UCS2_CROATIAN_CI",
		"UCS2_CZECH_CI",
		"UCS2_DANISH_CI",
		"UCS2_ESPERANTO_CI",
		"UCS2_ESTONIAN_CI",
		"UCS2_GENERAL_MYSQL500_CI",
		"UCS2_GERMAN2_CI",
		"UCS2_HUNGARIAN_CI",
		"UCS2_ICELANDIC_CI",
		"UCS2_LATVIAN_CI",
		"UCS2_LITHUANIAN_CI",
		"UCS2_PERSIAN_CI",
		"UCS2_POLISH_CI",
		"UCS2_ROMANIAN_CI",
		"UCS2_ROMAN_CI",
		"UCS2_SINHALA_CI",
		"UCS2_SLOVAK_CI",
		"UCS2_SLOVENIAN_CI",
		"UCS2_SPANISH2_CI",
		"UCS2_SPANISH_CI",
		"UCS2_SWEDISH_CI",
		"UCS2_TURKISH_CI",
		"UCS2_UNICODE_520_CI",
		"UCS2_UNICODE_CI",
		"UCS2_VIETNAMESE_CI",
		"UJIS_JAPANESE_CI",
		"UJIS_BIN",
		"UTF16_GENERAL_CI",
		"UTF16_BIN",
		"UTF16_CROATIAN_CI",
		"UTF16_CZECH_CI",
		"UTF16_DANISH_CI",
		"UTF16_ESPERANTO_CI",
		"UTF16_ESTONIAN_CI",
		"UTF16_GERMAN2_CI",
		"UTF16_HUNGARIAN_CI",
		"UTF16_ICELANDIC_CI",
		"UTF16_LATVIAN_CI",
		"UTF16_LITHUANIAN_CI",
		"UTF16_PERSIAN_CI",
		"UTF16_POLISH_CI",
		"UTF16_ROMANIAN_CI",
		"UTF16_ROMAN_CI",
		"UTF16_SINHALA_CI",
		"UTF16_SLOVAK_CI",
		"UTF16_SLOVENIAN_CI",
		"UTF16_SPANISH2_CI",
		"UTF16_SPANISH_CI",
		"UTF16_SWEDISH_CI",
		"UTF16_TURKISH_CI",
		"UTF16_UNICODE_520_CI",
		"UTF16_UNICODE_CI",
		"UTF16_VIETNAMESE_CI",
		"UTF16LE_GENERAL_CI",
		"UTF16LE_BIN",
		"UTF32_GENERAL_CI",
		"UTF32_BIN",
		"UTF32_CROATIAN_CI",
		"UTF32_CZECH_CI",
		"UTF32_DANISH_CI",
		"UTF32_ESPERANTO_CI",
		"UTF32_ESTONIAN_CI",
		"UTF32_GERMAN2_CI",
		"UTF32_HUNGARIAN_CI",
		"UTF32_ICELANDIC_CI",
		"UTF32_LATVIAN_CI",
		"UTF32_LITHUANIAN_CI",
		"UTF32_PERSIAN_CI",
		"UTF32_POLISH_CI",
		"UTF32_ROMANIAN_CI",
		"UTF32_ROMAN_CI",
		"UTF32_SINHALA_CI",
		"UTF32_SLOVAK_CI",
		"UTF32_SLOVENIAN_CI",
		"UTF32_SPANISH2_CI",
		"UTF32_SPANISH_CI",
		"UTF32_SWEDISH_CI",
		"UTF32_TURKISH_CI",
		"UTF32_UNICODE_520_CI",
		"UTF32_UNICODE_CI",
		"UTF32_VIETNAMESE_CI",
		"UTF8_GENERAL_CI",
		"UTF8_BIN",
		"UTF8_CROATIAN_CI",
		"UTF8_CZECH_CI",
		"UTF8_DANISH_CI",
		"UTF8_ESPERANTO_CI",
		"UTF8_ESTONIAN_CI",
		"UTF8_GENERAL_MYSQL500_CI",
		"UTF8_GERMAN2_CI",
		"UTF8_HUNGARIAN_CI",
		"UTF8_ICELANDIC_CI",
		"UTF8_LATVIAN_CI",
		"UTF8_LITHUANIAN_CI",
		"UTF8_PERSIAN_CI",
		"UTF8_POLISH_CI",
		"UTF8_ROMANIAN_CI",
		"UTF8_ROMAN_CI",
		"UTF8_SINHALA_CI",
		"UTF8_SLOVAK_CI",
		"UTF8_SLOVENIAN_CI",
		"UTF8_SPANISH2_CI",
		"UTF8_SPANISH_CI",
		"UTF8_SWEDISH_CI",
		"UTF8_TOLOWER_CI",
		"UTF8_TURKISH_CI",
		"UTF8_UNICODE_520_CI",
		"UTF8_UNICODE_CI",
		"UTF8_VIETNAMESE_CI",
		"UTF8MB4_0900_AI_CI",
		"UTF8MB4_0900_AS_CI",
		"UTF8MB4_0900_AS_CS",
		"UTF8MB4_0900_BIN",
		"UTF8MB4_BIN",
		"UTF8MB4_CROATIAN_CI",
		"UTF8MB4_CS_0900_AI_CI",
		"UTF8MB4_CS_0900_AS_CS",
		"UTF8MB4_CZECH_CI",
		"UTF8MB4_DANISH_CI",
		"UTF8MB4_DA_0900_AI_CI",
		"UTF8MB4_DA_0900_AS_CS",
		"UTF8MB4_DE_PB_0900_AI_CI",
		"UTF8MB4_DE_PB_0900_AS_CS",
		"UTF8MB4_EO_0900_AI_CI",
		"UTF8MB4_EO_0900_AS_CS",
		"UTF8MB4_ESPERANTO_CI",
		"UTF8MB4_ESTONIAN_CI",
		"UTF8MB4_ES_0900_AI_CI",
		"UTF8MB4_ES_0900_AS_CS",
		"UTF8MB4_ES_TRAD_0900_AI_CI",
		"UTF8MB4_ES_TRAD_0900_AS_CS",
		"UTF8MB4_ET_0900_AI_CI",
		"UTF8MB4_ET_0900_AS_CS",
		"UTF8MB4_GENERAL_CI",
		"UTF8MB4_GERMAN2_CI",
		"UTF8MB4_HR_0900_AI_CI",
		"UTF8MB4_HR_0900_AS_CS",
		"UTF8MB4_HUNGARIAN_CI",
		"UTF8MB4_HU_0900_AI_CI",
		"UTF8MB4_HU_0900_AS_CS",
		"UTF8MB4_ICELANDIC_CI",
		"UTF8MB4_IS_0900_AI_CI",
		"UTF8MB4_IS_0900_AS_CS",
		"UTF8MB4_JA_0900_AS_CS",
		"UTF8MB4_JA_0900_AS_CS_KS",
		"UTF8MB4_LATVIAN_CI",
		"UTF8MB4_LA_0900_AI_CI",
		"UTF8MB4_LA_0900_AS_CS",
		"UTF8MB4_LITHUANIAN_CI",
		"UTF8MB4_LT_0900_AI_CI",
		"UTF8MB4_LT_0900_AS_CS",
		"UTF8MB4_LV_0900_AI_CI",
		"UTF8MB4_LV_0900_AS_CS",
		"UTF8MB4_PERSIAN_CI",
		"UTF8MB4_PL_0900_AI_CI",
		"UTF8MB4_PL_0900_AS_CS",
		"UTF8MB4_POLISH_CI",
		"UTF8MB4_ROMANIAN_CI",
		"UTF8MB4_ROMAN_CI",
		"UTF8MB4_RO_0900_AI_CI",
		"UTF8MB4_RO_0900_AS_CS",
		"UTF8MB4_RU_0900_AI_CI",
		"UTF8MB4_RU_0900_AS_CS",
		"UTF8MB4_SINHALA_CI",
		"UTF8MB4_SK_0900_AI_CI",
		"UTF8MB4_SK_0900_AS_CS",
		"UTF8MB4_SLOVAK_CI",
		"UTF8MB4_SLOVENIAN_CI",
		"UTF8MB4_SL_0900_AI_CI",
		"UTF8MB4_SL_0900_AS_CS",
		"UTF8MB4_SPANISH2_CI",
		"UTF8MB4_SPANISH_CI",
		"UTF8MB4_SV_0900_AI_CI",
		"UTF8MB4_SV_0900_AS_CS",
		"UTF8MB4_SWEDISH_CI",
		"UTF8MB4_TR_0900_AI_CI",
		"UTF8MB4_TR_0900_AS_CS",
		"UTF8MB4_TURKISH_CI",
		"UTF8MB4_UNICODE_520_CI",
		"UTF8MB4_UNICODE_CI",
		"UTF8MB4_VIETNAMESE_CI",
		"UTF8MB4_VI_0900_AI_CI",
		"UTF8MB4_VI_0900_AS_CS",
		"UTF8MB4_ZH_0900_AS_CS",
	}
}

// GetMappingConfigurationVariablesCollationServerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesCollationServerEnum(val string) (ConfigurationVariablesCollationServerEnum, bool) {
	enum, ok := mappingConfigurationVariablesCollationServerEnumLowerCase[strings.ToLower(val)]
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

// ConfigurationVariablesExplainFormatEnum Enum with underlying type: string
type ConfigurationVariablesExplainFormatEnum string

// Set of constants representing the allowable values for ConfigurationVariablesExplainFormatEnum
const (
	ConfigurationVariablesExplainFormatTraditional ConfigurationVariablesExplainFormatEnum = "TRADITIONAL"
	ConfigurationVariablesExplainFormatJson        ConfigurationVariablesExplainFormatEnum = "JSON"
	ConfigurationVariablesExplainFormatTree        ConfigurationVariablesExplainFormatEnum = "TREE"
)

var mappingConfigurationVariablesExplainFormatEnum = map[string]ConfigurationVariablesExplainFormatEnum{
	"TRADITIONAL": ConfigurationVariablesExplainFormatTraditional,
	"JSON":        ConfigurationVariablesExplainFormatJson,
	"TREE":        ConfigurationVariablesExplainFormatTree,
}

var mappingConfigurationVariablesExplainFormatEnumLowerCase = map[string]ConfigurationVariablesExplainFormatEnum{
	"traditional": ConfigurationVariablesExplainFormatTraditional,
	"json":        ConfigurationVariablesExplainFormatJson,
	"tree":        ConfigurationVariablesExplainFormatTree,
}

// GetConfigurationVariablesExplainFormatEnumValues Enumerates the set of values for ConfigurationVariablesExplainFormatEnum
func GetConfigurationVariablesExplainFormatEnumValues() []ConfigurationVariablesExplainFormatEnum {
	values := make([]ConfigurationVariablesExplainFormatEnum, 0)
	for _, v := range mappingConfigurationVariablesExplainFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesExplainFormatEnumStringValues Enumerates the set of values in String for ConfigurationVariablesExplainFormatEnum
func GetConfigurationVariablesExplainFormatEnumStringValues() []string {
	return []string{
		"TRADITIONAL",
		"JSON",
		"TREE",
	}
}

// GetMappingConfigurationVariablesExplainFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesExplainFormatEnum(val string) (ConfigurationVariablesExplainFormatEnum, bool) {
	enum, ok := mappingConfigurationVariablesExplainFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ConfigurationVariablesInnodbChangeBufferingEnum Enum with underlying type: string
type ConfigurationVariablesInnodbChangeBufferingEnum string

// Set of constants representing the allowable values for ConfigurationVariablesInnodbChangeBufferingEnum
const (
	ConfigurationVariablesInnodbChangeBufferingNone    ConfigurationVariablesInnodbChangeBufferingEnum = "NONE"
	ConfigurationVariablesInnodbChangeBufferingInserts ConfigurationVariablesInnodbChangeBufferingEnum = "INSERTS"
	ConfigurationVariablesInnodbChangeBufferingDeletes ConfigurationVariablesInnodbChangeBufferingEnum = "DELETES"
	ConfigurationVariablesInnodbChangeBufferingChanges ConfigurationVariablesInnodbChangeBufferingEnum = "CHANGES"
	ConfigurationVariablesInnodbChangeBufferingPurges  ConfigurationVariablesInnodbChangeBufferingEnum = "PURGES"
	ConfigurationVariablesInnodbChangeBufferingAll     ConfigurationVariablesInnodbChangeBufferingEnum = "ALL"
)

var mappingConfigurationVariablesInnodbChangeBufferingEnum = map[string]ConfigurationVariablesInnodbChangeBufferingEnum{
	"NONE":    ConfigurationVariablesInnodbChangeBufferingNone,
	"INSERTS": ConfigurationVariablesInnodbChangeBufferingInserts,
	"DELETES": ConfigurationVariablesInnodbChangeBufferingDeletes,
	"CHANGES": ConfigurationVariablesInnodbChangeBufferingChanges,
	"PURGES":  ConfigurationVariablesInnodbChangeBufferingPurges,
	"ALL":     ConfigurationVariablesInnodbChangeBufferingAll,
}

var mappingConfigurationVariablesInnodbChangeBufferingEnumLowerCase = map[string]ConfigurationVariablesInnodbChangeBufferingEnum{
	"none":    ConfigurationVariablesInnodbChangeBufferingNone,
	"inserts": ConfigurationVariablesInnodbChangeBufferingInserts,
	"deletes": ConfigurationVariablesInnodbChangeBufferingDeletes,
	"changes": ConfigurationVariablesInnodbChangeBufferingChanges,
	"purges":  ConfigurationVariablesInnodbChangeBufferingPurges,
	"all":     ConfigurationVariablesInnodbChangeBufferingAll,
}

// GetConfigurationVariablesInnodbChangeBufferingEnumValues Enumerates the set of values for ConfigurationVariablesInnodbChangeBufferingEnum
func GetConfigurationVariablesInnodbChangeBufferingEnumValues() []ConfigurationVariablesInnodbChangeBufferingEnum {
	values := make([]ConfigurationVariablesInnodbChangeBufferingEnum, 0)
	for _, v := range mappingConfigurationVariablesInnodbChangeBufferingEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationVariablesInnodbChangeBufferingEnumStringValues Enumerates the set of values in String for ConfigurationVariablesInnodbChangeBufferingEnum
func GetConfigurationVariablesInnodbChangeBufferingEnumStringValues() []string {
	return []string{
		"NONE",
		"INSERTS",
		"DELETES",
		"CHANGES",
		"PURGES",
		"ALL",
	}
}

// GetMappingConfigurationVariablesInnodbChangeBufferingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationVariablesInnodbChangeBufferingEnum(val string) (ConfigurationVariablesInnodbChangeBufferingEnum, bool) {
	enum, ok := mappingConfigurationVariablesInnodbChangeBufferingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
