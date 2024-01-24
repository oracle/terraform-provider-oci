// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
)

func MysqlMysqlConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMysqlMysqlConfiguration,
		Read:     readMysqlMysqlConfiguration,
		Update:   updateMysqlMysqlConfiguration,
		Delete:   deleteMysqlMysqlConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"init_variables": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"lower_case_table_names": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"parent_configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"variables": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"autocommit": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"big_tables": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"binlog_expire_logs_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"binlog_row_metadata": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"binlog_row_value_options": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"binlog_transaction_compression": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"completion_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"connect_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"connection_memory_chunk_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"connection_memory_limit": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"cte_max_recursion_depth": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"default_authentication_plugin": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"foreign_key_checks": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"generated_random_password_length": {
							Type:       schema.TypeInt,
							Optional:   true,
							Computed:   true,
							ForceNew:   true,
							Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("generated_random_password_length"),
						},
						"global_connection_memory_limit": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"global_connection_memory_tracking": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"group_replication_consistency": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"information_schema_stats_expiry": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_buffer_pool_dump_pct": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_buffer_pool_instances": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_buffer_pool_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"innodb_ddl_buffer_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"innodb_ddl_threads": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_ft_enable_stopword": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_ft_max_token_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_ft_min_token_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_ft_num_word_optimize": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_ft_result_cache_limit": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"innodb_ft_server_stopword_table": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_lock_wait_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_log_writer_threads": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_max_purge_lag": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"innodb_max_purge_lag_delay": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"innodb_stats_persistent_sample_pages": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"innodb_stats_transient_sample_pages": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"interactive_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"local_infile": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mandatory_roles": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"max_allowed_packet": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"max_binlog_cache_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"max_connect_errors": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"max_connections": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"max_execution_time": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"max_heap_table_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"max_prepared_stmt_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysql_firewall_mode": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysql_zstd_default_compression_level": {
							Type:       schema.TypeInt,
							Optional:   true,
							Computed:   true,
							ForceNew:   true,
							Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("mysql_zstd_default_compression_level"),
						},
						"mysqlx_connect_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_deflate_default_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_deflate_max_client_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_document_id_unique_prefix": {
							Type:       schema.TypeInt,
							Optional:   true,
							Computed:   true,
							ForceNew:   true,
							Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("mysqlx_document_id_unique_prefix"),
						},
						"mysqlx_enable_hello_notice": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_idle_worker_thread_timeout": {
							Type:       schema.TypeInt,
							Optional:   true,
							Computed:   true,
							ForceNew:   true,
							Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("mysqlx_idle_worker_thread_timeout"),
						},
						"mysqlx_interactive_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_lz4default_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_lz4max_client_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_max_allowed_packet": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_min_worker_threads": {
							Type:       schema.TypeInt,
							Optional:   true,
							Computed:   true,
							ForceNew:   true,
							Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("mysqlx_min_worker_threads"),
						},
						"mysqlx_read_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_wait_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_write_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_zstd_default_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mysqlx_zstd_max_client_compression_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"net_read_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"net_write_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"parser_max_mem_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"query_alloc_block_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
							Deprecated:       tfresource.FieldDeprecatedAndAvoidReferences("query_alloc_block_size"),
						},
						"query_prealloc_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
							Deprecated:       tfresource.FieldDeprecatedAndAvoidReferences("query_prealloc_size"),
						},
						"regexp_time_limit": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"sort_buffer_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"sql_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"sql_require_primary_key": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"sql_warnings": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"thread_pool_dedicated_listeners": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"thread_pool_max_transactions_limit": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_zone": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"tmp_table_size": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"transaction_isolation": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"wait_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
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
		},
	}
}

func createMysqlMysqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()

	return tfresource.CreateResource(d, sync)
}

func readMysqlMysqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()

	return tfresource.ReadResource(sync)
}

func updateMysqlMysqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMysqlMysqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MysqlMysqlConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.MysqlaasClient
	Res                    *oci_mysql.Configuration
	DisableNotFoundRetries bool
}

func (s *MysqlMysqlConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlMysqlConfigurationResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MysqlMysqlConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.ConfigurationLifecycleStateActive),
	}
}

func (s *MysqlMysqlConfigurationResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *MysqlMysqlConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.ConfigurationLifecycleStateDeleted),
	}
}

func (s *MysqlMysqlConfigurationResourceCrud) UpdatedPending() []string {
	return []string{}
}

func (s *MysqlMysqlConfigurationResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.ConfigurationLifecycleStateActive),
	}
}

func (s *MysqlMysqlConfigurationResourceCrud) Create() error {
	request := oci_mysql.CreateConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if initVariables, ok := s.D.GetOkExists("init_variables"); ok {
		if tmpList := initVariables.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "init_variables", 0)
			tmp, err := s.mapToInitializationVariables(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InitVariables = &tmp
		}
	}

	if parentConfigurationId, ok := s.D.GetOkExists("parent_configuration_id"); ok {
		tmp := parentConfigurationId.(string)
		request.ParentConfigurationId = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	if variables, ok := s.D.GetOkExists("variables"); ok {
		if tmpList := variables.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "variables", 0)
			tmp, err := s.mapToConfigurationVariables(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Variables = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *MysqlMysqlConfigurationResourceCrud) Get() error {
	request := oci_mysql.GetConfigurationRequest{}

	tmp := s.D.Id()
	request.ConfigurationId = &tmp

	configurationId, err := parseMysqlConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.ConfigurationId = &configurationId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *MysqlMysqlConfigurationResourceCrud) Update() error {
	request := oci_mysql.UpdateConfigurationRequest{}

	tmp := s.D.Id()
	request.ConfigurationId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.UpdateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *MysqlMysqlConfigurationResourceCrud) Delete() error {
	request := oci_mysql.DeleteConfigurationRequest{}

	tmp := s.D.Id()
	request.ConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteConfiguration(context.Background(), request)
	return err
}

func (s *MysqlMysqlConfigurationResourceCrud) SetData() error {

	configurationId, err := parseMysqlConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("configuration_id", &configurationId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InitVariables != nil {
		s.D.Set("init_variables", []interface{}{InitializationVariablesToMap(s.Res.InitVariables)})
	} else {
		s.D.Set("init_variables", nil)
	}

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

func GetMysqlConfigurationCompositeId(configurationId string) string {
	configurationId = url.PathEscape(configurationId)
	compositeId := "configurations/" + configurationId
	return compositeId
}

func parseMysqlConfigurationCompositeId(compositeId string) (configurationId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("configurations/.*", compositeId)
	if !match || len(parts) != 2 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	configurationId, _ = url.PathUnescape(parts[1])

	return
}

func (s *MysqlMysqlConfigurationResourceCrud) mapToConfigurationVariables(fieldKeyFormat string) (oci_mysql.ConfigurationVariables, error) {
	result := oci_mysql.ConfigurationVariables{}

	if autocommit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "autocommit")); ok {
		tmp := autocommit.(bool)
		result.Autocommit = &tmp
	}

	if bigTables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "big_tables")); ok {
		tmp := bigTables.(bool)
		result.BigTables = &tmp
	}

	if binlogExpireLogsSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "binlog_expire_logs_seconds")); ok {
		tmp := binlogExpireLogsSeconds.(int)
		result.BinlogExpireLogsSeconds = &tmp
	}

	if binlogRowMetadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "binlog_row_metadata")); ok {
		result.BinlogRowMetadata = oci_mysql.ConfigurationVariablesBinlogRowMetadataEnum(binlogRowMetadata.(string))
	}

	if binlogRowValueOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "binlog_row_value_options")); ok {
		tmp := binlogRowValueOptions.(string)
		result.BinlogRowValueOptions = &tmp
	}

	if binlogTransactionCompression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "binlog_transaction_compression")); ok {
		tmp := binlogTransactionCompression.(bool)
		result.BinlogTransactionCompression = &tmp
	}

	if completionType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "completion_type")); ok {
		result.CompletionType = oci_mysql.ConfigurationVariablesCompletionTypeEnum(completionType.(string))
	}

	if connectTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connect_timeout")); ok {
		tmp := connectTimeout.(int)
		result.ConnectTimeout = &tmp
	}

	if connectionMemoryChunkSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_memory_chunk_size")); ok {
		tmp := connectionMemoryChunkSize.(int)
		result.ConnectionMemoryChunkSize = &tmp
	}

	if connectionMemoryLimit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_memory_limit")); ok {
		tmp := connectionMemoryLimit.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert connectionMemoryLimit string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ConnectionMemoryLimit = &tmpInt64
	}

	if cteMaxRecursionDepth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cte_max_recursion_depth")); ok {
		tmp := cteMaxRecursionDepth.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert cteMaxRecursionDepth string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.CteMaxRecursionDepth = &tmpInt64
	}

	if defaultAuthenticationPlugin, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_authentication_plugin")); ok {
		result.DefaultAuthenticationPlugin = oci_mysql.ConfigurationVariablesDefaultAuthenticationPluginEnum(defaultAuthenticationPlugin.(string))
	}

	if foreignKeyChecks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "foreign_key_checks")); ok {
		tmp := foreignKeyChecks.(bool)
		result.ForeignKeyChecks = &tmp
	}

	if generatedRandomPasswordLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "generated_random_password_length")); ok {
		tmp := generatedRandomPasswordLength.(int)
		result.GeneratedRandomPasswordLength = &tmp
	}

	if globalConnectionMemoryLimit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "global_connection_memory_limit")); ok {
		tmp := globalConnectionMemoryLimit.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert globalConnectionMemoryLimit string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.GlobalConnectionMemoryLimit = &tmpInt64
	}

	if globalConnectionMemoryTracking, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "global_connection_memory_tracking")); ok {
		tmp := globalConnectionMemoryTracking.(bool)
		result.GlobalConnectionMemoryTracking = &tmp
	}

	if groupReplicationConsistency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_replication_consistency")); ok {
		result.GroupReplicationConsistency = oci_mysql.ConfigurationVariablesGroupReplicationConsistencyEnum(groupReplicationConsistency.(string))
	}

	if informationSchemaStatsExpiry, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "information_schema_stats_expiry")); ok {
		tmp := informationSchemaStatsExpiry.(int)
		result.InformationSchemaStatsExpiry = &tmp
	}

	if innodbBufferPoolDumpPct, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_buffer_pool_dump_pct")); ok {
		tmp := innodbBufferPoolDumpPct.(int)
		result.InnodbBufferPoolDumpPct = &tmp
	}

	if innodbBufferPoolInstances, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_buffer_pool_instances")); ok {
		tmp := innodbBufferPoolInstances.(int)
		result.InnodbBufferPoolInstances = &tmp
	}

	if innodbBufferPoolSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_buffer_pool_size")); ok {
		tmp := innodbBufferPoolSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert innodbBufferPoolSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.InnodbBufferPoolSize = &tmpInt64
	}

	if innodbDdlBufferSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_ddl_buffer_size")); ok {
		tmp := innodbDdlBufferSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert innodbDdlBufferSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.InnodbDdlBufferSize = &tmpInt64
	}

	if innodbDdlThreads, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_ddl_threads")); ok {
		tmp := innodbDdlThreads.(int)
		result.InnodbDdlThreads = &tmp
	}

	if innodbFtEnableStopword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_ft_enable_stopword")); ok {
		tmp := innodbFtEnableStopword.(bool)
		result.InnodbFtEnableStopword = &tmp
	}

	if innodbFtMaxTokenSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_ft_max_token_size")); ok {
		tmp := innodbFtMaxTokenSize.(int)
		result.InnodbFtMaxTokenSize = &tmp
	}

	if innodbFtMinTokenSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_ft_min_token_size")); ok {
		tmp := innodbFtMinTokenSize.(int)
		result.InnodbFtMinTokenSize = &tmp
	}

	if innodbFtNumWordOptimize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_ft_num_word_optimize")); ok {
		tmp := innodbFtNumWordOptimize.(int)
		result.InnodbFtNumWordOptimize = &tmp
	}

	if innodbFtResultCacheLimit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_ft_result_cache_limit")); ok {
		tmp := innodbFtResultCacheLimit.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert innodbFtResultCacheLimit string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.InnodbFtResultCacheLimit = &tmpInt64
	}

	if innodbFtServerStopwordTable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_ft_server_stopword_table")); ok {
		tmp := innodbFtServerStopwordTable.(string)
		result.InnodbFtServerStopwordTable = &tmp
	}

	if innodbLockWaitTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_lock_wait_timeout")); ok {
		tmp := innodbLockWaitTimeout.(int)
		result.InnodbLockWaitTimeout = &tmp
	}

	if innodbLogWriterThreads, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_log_writer_threads")); ok {
		tmp := innodbLogWriterThreads.(bool)
		result.InnodbLogWriterThreads = &tmp
	}

	if innodbMaxPurgeLag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_max_purge_lag")); ok {
		tmp := innodbMaxPurgeLag.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert innodbMaxPurgeLag string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.InnodbMaxPurgeLag = &tmpInt64
	}

	if innodbMaxPurgeLagDelay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_max_purge_lag_delay")); ok {
		tmp := innodbMaxPurgeLagDelay.(int)
		result.InnodbMaxPurgeLagDelay = &tmp
	}

	if innodbStatsPersistentSamplePages, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_stats_persistent_sample_pages")); ok {
		tmp := innodbStatsPersistentSamplePages.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert innodbStatsPersistentSamplePages string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.InnodbStatsPersistentSamplePages = &tmpInt64
	}

	if innodbStatsTransientSamplePages, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "innodb_stats_transient_sample_pages")); ok {
		tmp := innodbStatsTransientSamplePages.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert innodbStatsTransientSamplePages string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.InnodbStatsTransientSamplePages = &tmpInt64
	}

	if interactiveTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interactive_timeout")); ok {
		tmp := interactiveTimeout.(int)
		result.InteractiveTimeout = &tmp
	}

	if localInfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "local_infile")); ok {
		tmp := localInfile.(bool)
		result.LocalInfile = &tmp
	}

	if mandatoryRoles, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mandatory_roles")); ok {
		tmp := mandatoryRoles.(string)
		result.MandatoryRoles = &tmp
	}

	if maxAllowedPacket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_allowed_packet")); ok {
		tmp := maxAllowedPacket.(int)
		result.MaxAllowedPacket = &tmp
	}

	if maxBinlogCacheSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_binlog_cache_size")); ok {
		tmp := maxBinlogCacheSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert maxBinlogCacheSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MaxBinlogCacheSize = &tmpInt64
	}

	if maxConnectErrors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_connect_errors")); ok {
		tmp := maxConnectErrors.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert maxConnectErrors string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MaxConnectErrors = &tmpInt64
	}

	if maxConnections, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_connections")); ok {
		tmp := maxConnections.(int)
		result.MaxConnections = &tmp
	}

	if maxExecutionTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_execution_time")); ok {
		tmp := maxExecutionTime.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert maxExecutionTime string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MaxExecutionTime = &tmpInt64
	}

	if maxHeapTableSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_heap_table_size")); ok {
		tmp := maxHeapTableSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert maxHeapTableSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MaxHeapTableSize = &tmpInt64
	}

	if maxPreparedStmtCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_prepared_stmt_count")); ok {
		tmp := maxPreparedStmtCount.(int)
		result.MaxPreparedStmtCount = &tmp
	}

	if mysqlFirewallMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysql_firewall_mode")); ok {
		tmp := mysqlFirewallMode.(bool)
		result.MysqlFirewallMode = &tmp
	}

	if mysqlZstdDefaultCompressionLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysql_zstd_default_compression_level")); ok {
		tmp := mysqlZstdDefaultCompressionLevel.(int)
		result.MysqlZstdDefaultCompressionLevel = &tmp
	}

	if mysqlxConnectTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_connect_timeout")); ok {
		tmp := mysqlxConnectTimeout.(int)
		result.MysqlxConnectTimeout = &tmp
	}

	if mysqlxDeflateDefaultCompressionLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_deflate_default_compression_level")); ok {
		tmp := mysqlxDeflateDefaultCompressionLevel.(int)
		result.MysqlxDeflateDefaultCompressionLevel = &tmp
	}

	if mysqlxDeflateMaxClientCompressionLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_deflate_max_client_compression_level")); ok {
		tmp := mysqlxDeflateMaxClientCompressionLevel.(int)
		result.MysqlxDeflateMaxClientCompressionLevel = &tmp
	}

	if mysqlxDocumentIdUniquePrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_document_id_unique_prefix")); ok {
		tmp := mysqlxDocumentIdUniquePrefix.(int)
		result.MysqlxDocumentIdUniquePrefix = &tmp
	}

	if mysqlxEnableHelloNotice, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_enable_hello_notice")); ok {
		tmp := mysqlxEnableHelloNotice.(bool)
		result.MysqlxEnableHelloNotice = &tmp
	}

	if mysqlxIdleWorkerThreadTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_idle_worker_thread_timeout")); ok {
		tmp := mysqlxIdleWorkerThreadTimeout.(int)
		result.MysqlxIdleWorkerThreadTimeout = &tmp
	}

	if mysqlxInteractiveTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_interactive_timeout")); ok {
		tmp := mysqlxInteractiveTimeout.(int)
		result.MysqlxInteractiveTimeout = &tmp
	}

	if mysqlxLz4DefaultCompressionLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_lz4default_compression_level")); ok {
		tmp := mysqlxLz4DefaultCompressionLevel.(int)
		result.MysqlxLz4DefaultCompressionLevel = &tmp
	}

	if mysqlxLz4MaxClientCompressionLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_lz4max_client_compression_level")); ok {
		tmp := mysqlxLz4MaxClientCompressionLevel.(int)
		result.MysqlxLz4MaxClientCompressionLevel = &tmp
	}

	if mysqlxMaxAllowedPacket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_max_allowed_packet")); ok {
		tmp := mysqlxMaxAllowedPacket.(int)
		result.MysqlxMaxAllowedPacket = &tmp
	}

	if mysqlxMinWorkerThreads, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_min_worker_threads")); ok {
		tmp := mysqlxMinWorkerThreads.(int)
		result.MysqlxMinWorkerThreads = &tmp
	}

	if mysqlxReadTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_read_timeout")); ok {
		tmp := mysqlxReadTimeout.(int)
		result.MysqlxReadTimeout = &tmp
	}

	if mysqlxWaitTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_wait_timeout")); ok {
		tmp := mysqlxWaitTimeout.(int)
		result.MysqlxWaitTimeout = &tmp
	}

	if mysqlxWriteTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_write_timeout")); ok {
		tmp := mysqlxWriteTimeout.(int)
		result.MysqlxWriteTimeout = &tmp
	}

	if mysqlxZstdDefaultCompressionLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_zstd_default_compression_level")); ok {
		tmp := mysqlxZstdDefaultCompressionLevel.(int)
		result.MysqlxZstdDefaultCompressionLevel = &tmp
	}

	if mysqlxZstdMaxClientCompressionLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysqlx_zstd_max_client_compression_level")); ok {
		tmp := mysqlxZstdMaxClientCompressionLevel.(int)
		result.MysqlxZstdMaxClientCompressionLevel = &tmp
	}

	if netReadTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "net_read_timeout")); ok {
		tmp := netReadTimeout.(int)
		result.NetReadTimeout = &tmp
	}

	if netWriteTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "net_write_timeout")); ok {
		tmp := netWriteTimeout.(int)
		result.NetWriteTimeout = &tmp
	}

	if parserMaxMemSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parser_max_mem_size")); ok {
		tmp := parserMaxMemSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert parserMaxMemSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ParserMaxMemSize = &tmpInt64
	}

	if queryAllocBlockSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_alloc_block_size")); ok {
		tmp := queryAllocBlockSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert queryAllocBlockSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.QueryAllocBlockSize = &tmpInt64
	}

	if queryPreallocSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_prealloc_size")); ok {
		tmp := queryPreallocSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert queryPreallocSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.QueryPreallocSize = &tmpInt64
	}

	if regexpTimeLimit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "regexp_time_limit")); ok {
		tmp := regexpTimeLimit.(int)
		result.RegexpTimeLimit = &tmp
	}

	if sortBufferSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sort_buffer_size")); ok {
		tmp := sortBufferSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert sortBufferSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.SortBufferSize = &tmpInt64
	}

	if sqlMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_mode")); ok {
		tmp := sqlMode.(string)
		result.SqlMode = &tmp
	}

	if sqlRequirePrimaryKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_require_primary_key")); ok {
		tmp := sqlRequirePrimaryKey.(bool)
		result.SqlRequirePrimaryKey = &tmp
	}

	if sqlWarnings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sql_warnings")); ok {
		tmp := sqlWarnings.(bool)
		result.SqlWarnings = &tmp
	}

	if threadPoolDedicatedListeners, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "thread_pool_dedicated_listeners")); ok {
		tmp := threadPoolDedicatedListeners.(bool)
		result.ThreadPoolDedicatedListeners = &tmp
	}

	if threadPoolMaxTransactionsLimit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "thread_pool_max_transactions_limit")); ok {
		tmp := threadPoolMaxTransactionsLimit.(int)
		result.ThreadPoolMaxTransactionsLimit = &tmp
	}

	if timeZone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_zone")); ok {
		tmp := timeZone.(string)
		result.TimeZone = &tmp
	}

	if tmpTableSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tmp_table_size")); ok {
		tmp := tmpTableSize.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert tmpTableSize string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.TmpTableSize = &tmpInt64
	}

	if transactionIsolation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transaction_isolation")); ok {
		result.TransactionIsolation = oci_mysql.ConfigurationVariablesTransactionIsolationEnum(transactionIsolation.(string))
	}

	if waitTimeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wait_timeout")); ok {
		tmp := waitTimeout.(int)
		result.WaitTimeout = &tmp
	}

	return result, nil
}

func ConfigurationVariablesToMap(obj *oci_mysql.ConfigurationVariables) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Autocommit != nil {
		result["autocommit"] = bool(*obj.Autocommit)
	}

	if obj.BigTables != nil {
		result["big_tables"] = bool(*obj.BigTables)
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

	if obj.ConnectionMemoryChunkSize != nil {
		result["connection_memory_chunk_size"] = int(*obj.ConnectionMemoryChunkSize)
	}

	if obj.ConnectionMemoryLimit != nil {
		result["connection_memory_limit"] = strconv.FormatInt(*obj.ConnectionMemoryLimit, 10)
	}

	if obj.CteMaxRecursionDepth != nil {
		result["cte_max_recursion_depth"] = strconv.FormatInt(*obj.CteMaxRecursionDepth, 10)
	}

	result["default_authentication_plugin"] = string(obj.DefaultAuthenticationPlugin)

	if obj.ForeignKeyChecks != nil {
		result["foreign_key_checks"] = bool(*obj.ForeignKeyChecks)
	}

	if obj.GeneratedRandomPasswordLength != nil {
		result["generated_random_password_length"] = int(*obj.GeneratedRandomPasswordLength)
	}

	if obj.GlobalConnectionMemoryLimit != nil {
		result["global_connection_memory_limit"] = strconv.FormatInt(*obj.GlobalConnectionMemoryLimit, 10)
	}

	if obj.GlobalConnectionMemoryTracking != nil {
		result["global_connection_memory_tracking"] = bool(*obj.GlobalConnectionMemoryTracking)
	}

	result["group_replication_consistency"] = string(obj.GroupReplicationConsistency)

	if obj.InformationSchemaStatsExpiry != nil {
		result["information_schema_stats_expiry"] = int(*obj.InformationSchemaStatsExpiry)
	}

	if obj.InnodbBufferPoolDumpPct != nil {
		result["innodb_buffer_pool_dump_pct"] = int(*obj.InnodbBufferPoolDumpPct)
	}

	if obj.InnodbBufferPoolInstances != nil {
		result["innodb_buffer_pool_instances"] = int(*obj.InnodbBufferPoolInstances)
	}

	if obj.InnodbBufferPoolSize != nil {
		result["innodb_buffer_pool_size"] = strconv.FormatInt(*obj.InnodbBufferPoolSize, 10)
	}

	if obj.InnodbDdlBufferSize != nil {
		result["innodb_ddl_buffer_size"] = strconv.FormatInt(*obj.InnodbDdlBufferSize, 10)
	}

	if obj.InnodbDdlThreads != nil {
		result["innodb_ddl_threads"] = int(*obj.InnodbDdlThreads)
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
		result["innodb_ft_result_cache_limit"] = strconv.FormatInt(*obj.InnodbFtResultCacheLimit, 10)
	}

	if obj.InnodbFtServerStopwordTable != nil {
		result["innodb_ft_server_stopword_table"] = string(*obj.InnodbFtServerStopwordTable)
	}

	if obj.InnodbLockWaitTimeout != nil {
		result["innodb_lock_wait_timeout"] = int(*obj.InnodbLockWaitTimeout)
	}

	if obj.InnodbLogWriterThreads != nil {
		result["innodb_log_writer_threads"] = bool(*obj.InnodbLogWriterThreads)
	}

	if obj.InnodbMaxPurgeLag != nil {
		result["innodb_max_purge_lag"] = strconv.FormatInt(*obj.InnodbMaxPurgeLag, 10)
	}

	if obj.InnodbMaxPurgeLagDelay != nil {
		result["innodb_max_purge_lag_delay"] = int(*obj.InnodbMaxPurgeLagDelay)
	}

	if obj.InnodbStatsPersistentSamplePages != nil {
		result["innodb_stats_persistent_sample_pages"] = strconv.FormatInt(*obj.InnodbStatsPersistentSamplePages, 10)
	}

	if obj.InnodbStatsTransientSamplePages != nil {
		result["innodb_stats_transient_sample_pages"] = strconv.FormatInt(*obj.InnodbStatsTransientSamplePages, 10)
	}

	if obj.InteractiveTimeout != nil {
		result["interactive_timeout"] = int(*obj.InteractiveTimeout)
	}

	if obj.LocalInfile != nil {
		result["local_infile"] = bool(*obj.LocalInfile)
	}

	if obj.MandatoryRoles != nil {
		result["mandatory_roles"] = string(*obj.MandatoryRoles)
	}

	if obj.MaxAllowedPacket != nil {
		result["max_allowed_packet"] = int(*obj.MaxAllowedPacket)
	}

	if obj.MaxBinlogCacheSize != nil {
		result["max_binlog_cache_size"] = strconv.FormatInt(*obj.MaxBinlogCacheSize, 10)
	}

	if obj.MaxConnectErrors != nil {
		result["max_connect_errors"] = strconv.FormatInt(*obj.MaxConnectErrors, 10)
	}

	if obj.MaxConnections != nil {
		result["max_connections"] = int(*obj.MaxConnections)
	}

	if obj.MaxExecutionTime != nil {
		result["max_execution_time"] = strconv.FormatInt(*obj.MaxExecutionTime, 10)
	}

	if obj.MaxHeapTableSize != nil {
		result["max_heap_table_size"] = strconv.FormatInt(*obj.MaxHeapTableSize, 10)
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

	if obj.NetReadTimeout != nil {
		result["net_read_timeout"] = int(*obj.NetReadTimeout)
	}

	if obj.NetWriteTimeout != nil {
		result["net_write_timeout"] = int(*obj.NetWriteTimeout)
	}

	if obj.ParserMaxMemSize != nil {
		result["parser_max_mem_size"] = strconv.FormatInt(*obj.ParserMaxMemSize, 10)
	}

	if obj.QueryAllocBlockSize != nil {
		result["query_alloc_block_size"] = strconv.FormatInt(*obj.QueryAllocBlockSize, 10)
	}

	if obj.QueryPreallocSize != nil {
		result["query_prealloc_size"] = strconv.FormatInt(*obj.QueryPreallocSize, 10)
	}

	if obj.RegexpTimeLimit != nil {
		result["regexp_time_limit"] = int(*obj.RegexpTimeLimit)
	}

	if obj.SortBufferSize != nil {
		result["sort_buffer_size"] = strconv.FormatInt(*obj.SortBufferSize, 10)
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

	if obj.ThreadPoolDedicatedListeners != nil {
		result["thread_pool_dedicated_listeners"] = bool(*obj.ThreadPoolDedicatedListeners)
	}

	if obj.ThreadPoolMaxTransactionsLimit != nil {
		result["thread_pool_max_transactions_limit"] = int(*obj.ThreadPoolMaxTransactionsLimit)
	}

	if obj.TimeZone != nil {
		result["time_zone"] = string(*obj.TimeZone)
	}

	if obj.TmpTableSize != nil {
		result["tmp_table_size"] = strconv.FormatInt(*obj.TmpTableSize, 10)
	}

	result["transaction_isolation"] = string(obj.TransactionIsolation)

	if obj.WaitTimeout != nil {
		result["wait_timeout"] = int(*obj.WaitTimeout)
	}

	return result
}

func (s *MysqlMysqlConfigurationResourceCrud) mapToInitializationVariables(fieldKeyFormat string) (oci_mysql.InitializationVariables, error) {
	result := oci_mysql.InitializationVariables{}

	if lowerCaseTableNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lower_case_table_names")); ok {
		result.LowerCaseTableNames = oci_mysql.InitializationVariablesLowerCaseTableNamesEnum(lowerCaseTableNames.(string))
	}

	return result, nil
}

func InitializationVariablesToMap(obj *oci_mysql.InitializationVariables) map[string]interface{} {
	result := map[string]interface{}{}

	result["lower_case_table_names"] = string(obj.LowerCaseTableNames)

	return result
}
