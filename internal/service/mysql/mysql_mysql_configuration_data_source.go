// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v58/mysql"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MysqlMysqlConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMysqlMysqlConfiguration,
		Schema: map[string]*schema.Schema{
			"configuration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"parent_configuration_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"variables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"autocommit": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"binlog_expire_logs_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"binlog_row_metadata": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"binlog_row_value_options": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"binlog_transaction_compression": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"completion_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"connect_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cte_max_recursion_depth": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"default_authentication_plugin": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"foreign_key_checks": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"generated_random_password_length": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"group_replication_consistency": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"information_schema_stats_expiry": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"innodb_buffer_pool_instances": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"innodb_buffer_pool_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     utils.ValidateInt64TypeString,
							DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
						},
						"innodb_ft_enable_stopword": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"innodb_ft_max_token_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"innodb_ft_min_token_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"innodb_ft_num_word_optimize": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"innodb_ft_result_cache_limit": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"innodb_ft_server_stopword_table": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"innodb_lock_wait_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"innodb_max_purge_lag": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"innodb_max_purge_lag_delay": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"local_infile": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"mandatory_roles": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_connections": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_execution_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_prepared_stmt_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysql_firewall_mode": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"mysql_zstd_default_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_connect_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_deflate_default_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_deflate_max_client_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_document_id_unique_prefix": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_enable_hello_notice": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"mysqlx_idle_worker_thread_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_interactive_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_lz4default_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_lz4max_client_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_max_allowed_packet": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_min_worker_threads": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_read_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_wait_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_write_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_zstd_default_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mysqlx_zstd_max_client_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"parser_max_mem_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"query_alloc_block_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"query_prealloc_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sql_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sql_require_primary_key": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"sql_warnings": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"transaction_isolation": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
		},
	}
}

func readSingularMysqlMysqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()

	return tfresource.ReadResource(sync)
}

type MysqlMysqlConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.MysqlaasClient
	Res    *oci_mysql.GetConfigurationResponse
}

func (s *MysqlMysqlConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlConfigurationDataSourceCrud) Get() error {
	request := oci_mysql.GetConfigurationRequest{}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlMysqlConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ParentConfigurationId != nil {
		s.D.Set("parent_configuration_id", *s.Res.ParentConfigurationId)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.Variables != nil {
		s.D.Set("variables", []interface{}{ConfigurationVariablesToMap(s.Res.Variables)})
	} else {
		s.D.Set("variables", nil)
	}

	return nil
}

func ConfigurationVariablesToMap(obj *oci_mysql.ConfigurationVariables) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Autocommit != nil {
		result["autocommit"] = bool(*obj.Autocommit)
	}

	if obj.BinlogExpireLogsSeconds != nil {
		result["binlog_expire_logs_seconds"] = int(*obj.BinlogExpireLogsSeconds)
	}

	result["binlog_row_metadata"] = string(obj.BinlogRowMetadata)

	if obj.BinlogRowValueOptions != nil {
		result["binlog_row_value_options"] = string(*obj.BinlogRowValueOptions)
	}

	if obj.BinlogTransactionCompression != nil {
		result["binlog_transaction_compression"] = bool(*obj.BinlogTransactionCompression)
	}

	result["completion_type"] = string(obj.CompletionType)

	if obj.ConnectTimeout != nil {
		result["connect_timeout"] = int(*obj.ConnectTimeout)
	}

	if obj.CteMaxRecursionDepth != nil {
		result["cte_max_recursion_depth"] = int(*obj.CteMaxRecursionDepth)
	}

	result["default_authentication_plugin"] = string(obj.DefaultAuthenticationPlugin)

	if obj.ForeignKeyChecks != nil {
		result["foreign_key_checks"] = bool(*obj.ForeignKeyChecks)
	}

	if obj.GeneratedRandomPasswordLength != nil {
		result["generated_random_password_length"] = int(*obj.GeneratedRandomPasswordLength)
	}

	result["group_replication_consistency"] = string(obj.GroupReplicationConsistency)

	if obj.InformationSchemaStatsExpiry != nil {
		result["information_schema_stats_expiry"] = int(*obj.InformationSchemaStatsExpiry)
	}

	if obj.InnodbBufferPoolInstances != nil {
		result["innodb_buffer_pool_instances"] = int(*obj.InnodbBufferPoolInstances)
	}

	if obj.InnodbBufferPoolSize != nil {
		result["innodb_buffer_pool_size"] = strconv.FormatInt(*obj.InnodbBufferPoolSize, 10)
	}

	if obj.InnodbFtEnableStopword != nil {
		result["innodb_ft_enable_stopword"] = bool(*obj.InnodbFtEnableStopword)
	}

	if obj.InnodbFtMaxTokenSize != nil {
		result["innodb_ft_max_token_size"] = int(*obj.InnodbFtMaxTokenSize)
	}

	if obj.InnodbFtMinTokenSize != nil {
		result["innodb_ft_min_token_size"] = int(*obj.InnodbFtMinTokenSize)
	}

	if obj.InnodbFtNumWordOptimize != nil {
		result["innodb_ft_num_word_optimize"] = int(*obj.InnodbFtNumWordOptimize)
	}

	if obj.InnodbFtResultCacheLimit != nil {
		result["innodb_ft_result_cache_limit"] = int(*obj.InnodbFtResultCacheLimit)
	}

	if obj.InnodbFtServerStopwordTable != nil {
		result["innodb_ft_server_stopword_table"] = string(*obj.InnodbFtServerStopwordTable)
	}

	if obj.InnodbLockWaitTimeout != nil {
		result["innodb_lock_wait_timeout"] = int(*obj.InnodbLockWaitTimeout)
	}

	if obj.InnodbMaxPurgeLag != nil {
		result["innodb_max_purge_lag"] = int(*obj.InnodbMaxPurgeLag)
	}

	if obj.InnodbMaxPurgeLagDelay != nil {
		result["innodb_max_purge_lag_delay"] = int(*obj.InnodbMaxPurgeLagDelay)
	}

	if obj.LocalInfile != nil {
		result["local_infile"] = bool(*obj.LocalInfile)
	}

	if obj.MandatoryRoles != nil {
		result["mandatory_roles"] = string(*obj.MandatoryRoles)
	}

	if obj.MaxConnections != nil {
		result["max_connections"] = int(*obj.MaxConnections)
	}

	if obj.MaxExecutionTime != nil {
		result["max_execution_time"] = int(*obj.MaxExecutionTime)
	}

	if obj.MaxPreparedStmtCount != nil {
		result["max_prepared_stmt_count"] = int(*obj.MaxPreparedStmtCount)
	}

	if obj.MysqlFirewallMode != nil {
		result["mysql_firewall_mode"] = bool(*obj.MysqlFirewallMode)
	}

	if obj.MysqlZstdDefaultCompressionLevel != nil {
		result["mysql_zstd_default_compression_level"] = int(*obj.MysqlZstdDefaultCompressionLevel)
	}

	if obj.MysqlxConnectTimeout != nil {
		result["mysqlx_connect_timeout"] = int(*obj.MysqlxConnectTimeout)
	}

	if obj.MysqlxDeflateDefaultCompressionLevel != nil {
		result["mysqlx_deflate_default_compression_level"] = int(*obj.MysqlxDeflateDefaultCompressionLevel)
	}

	if obj.MysqlxDeflateMaxClientCompressionLevel != nil {
		result["mysqlx_deflate_max_client_compression_level"] = int(*obj.MysqlxDeflateMaxClientCompressionLevel)
	}

	if obj.MysqlxDocumentIdUniquePrefix != nil {
		result["mysqlx_document_id_unique_prefix"] = int(*obj.MysqlxDocumentIdUniquePrefix)
	}

	if obj.MysqlxEnableHelloNotice != nil {
		result["mysqlx_enable_hello_notice"] = bool(*obj.MysqlxEnableHelloNotice)
	}

	if obj.MysqlxIdleWorkerThreadTimeout != nil {
		result["mysqlx_idle_worker_thread_timeout"] = int(*obj.MysqlxIdleWorkerThreadTimeout)
	}

	if obj.MysqlxInteractiveTimeout != nil {
		result["mysqlx_interactive_timeout"] = int(*obj.MysqlxInteractiveTimeout)
	}

	if obj.MysqlxLz4DefaultCompressionLevel != nil {
		result["mysqlx_lz4default_compression_level"] = int(*obj.MysqlxLz4DefaultCompressionLevel)
	}

	if obj.MysqlxLz4MaxClientCompressionLevel != nil {
		result["mysqlx_lz4max_client_compression_level"] = int(*obj.MysqlxLz4MaxClientCompressionLevel)
	}

	if obj.MysqlxMaxAllowedPacket != nil {
		result["mysqlx_max_allowed_packet"] = int(*obj.MysqlxMaxAllowedPacket)
	}

	if obj.MysqlxMinWorkerThreads != nil {
		result["mysqlx_min_worker_threads"] = int(*obj.MysqlxMinWorkerThreads)
	}

	if obj.MysqlxReadTimeout != nil {
		result["mysqlx_read_timeout"] = int(*obj.MysqlxReadTimeout)
	}

	if obj.MysqlxWaitTimeout != nil {
		result["mysqlx_wait_timeout"] = int(*obj.MysqlxWaitTimeout)
	}

	if obj.MysqlxWriteTimeout != nil {
		result["mysqlx_write_timeout"] = int(*obj.MysqlxWriteTimeout)
	}

	if obj.MysqlxZstdDefaultCompressionLevel != nil {
		result["mysqlx_zstd_default_compression_level"] = int(*obj.MysqlxZstdDefaultCompressionLevel)
	}

	if obj.MysqlxZstdMaxClientCompressionLevel != nil {
		result["mysqlx_zstd_max_client_compression_level"] = int(*obj.MysqlxZstdMaxClientCompressionLevel)
	}

	if obj.ParserMaxMemSize != nil {
		result["parser_max_mem_size"] = int(*obj.ParserMaxMemSize)
	}

	if obj.QueryAllocBlockSize != nil {
		result["query_alloc_block_size"] = int(*obj.QueryAllocBlockSize)
	}

	if obj.QueryPreallocSize != nil {
		result["query_prealloc_size"] = int(*obj.QueryPreallocSize)
	}

	if obj.SqlMode != nil {
		result["sql_mode"] = string(*obj.SqlMode)
	}

	if obj.SqlRequirePrimaryKey != nil {
		result["sql_require_primary_key"] = bool(*obj.SqlRequirePrimaryKey)
	}

	if obj.SqlWarnings != nil {
		result["sql_warnings"] = bool(*obj.SqlWarnings)
	}

	result["transaction_isolation"] = string(obj.TransactionIsolation)

	return result
}
