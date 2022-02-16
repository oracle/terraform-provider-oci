// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v58/mysql"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MysqlMysqlConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlMysqlConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"shape_name": {
							Type:     schema.TypeString,
							Required: true,
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
						"parent_configuration_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"variables": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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

						// Computed
						"id": {
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
					},
				},
			},
		},
	}
}

func readMysqlMysqlConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()

	return tfresource.ReadResource(sync)
}

type MysqlMysqlConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.MysqlaasClient
	Res    *oci_mysql.ListConfigurationsResponse
}

func (s *MysqlMysqlConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlConfigurationsDataSourceCrud) Get() error {
	request := oci_mysql.ListConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationId, ok := s.D.GetOkExists("configuration_id"); ok {
		tmp := configurationId.(string)
		request.ConfigurationId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.ConfigurationLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		interfaces := type_.([]interface{})
		tmp := make([]oci_mysql.ListConfigurationsTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_mysql.ListConfigurationsTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 {
			request.Type = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.ListConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MysqlMysqlConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlMysqlConfigurationsDataSource-", MysqlMysqlConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mysqlConfiguration := map[string]interface{}{}

		if r.CompartmentId != nil {
			mysqlConfiguration["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			mysqlConfiguration["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			mysqlConfiguration["description"] = *r.Description
		}

		if r.DisplayName != nil {
			mysqlConfiguration["display_name"] = *r.DisplayName
		}

		mysqlConfiguration["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			mysqlConfiguration["id"] = *r.Id
		}

		if r.ShapeName != nil {
			mysqlConfiguration["shape_name"] = *r.ShapeName
		}

		mysqlConfiguration["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			mysqlConfiguration["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			mysqlConfiguration["time_updated"] = r.TimeUpdated.String()
		}

		mysqlConfiguration["type"] = r.Type

		resources = append(resources, mysqlConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MysqlMysqlConfigurationsDataSource().Schema["configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("configurations", resources); err != nil {
		return err
	}

	return nil
}
