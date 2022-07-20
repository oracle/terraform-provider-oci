// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
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
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
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
