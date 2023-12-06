// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseAutonomousDatabasesClonesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabasesClones,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"clone_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"actual_used_data_storage_size_in_tbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"allocated_storage_size_in_tbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"apex_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"apex_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ords_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"are_primary_whitelisted_ips_used": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"autonomous_container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"autonomous_maintenance_schedule_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"available_upgrade_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"backup_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"manual_backup_bucket_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"manual_backup_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"backup_retention_period_in_days": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"character_set": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compute_count": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"compute_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connection_strings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"all_connection_strings": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"dedicated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"high": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"low": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"medium": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"profiles": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"consumer_group": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"host_format": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_regional": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"protocol": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"session_mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"syntax_format": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"tls_authentication": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"connection_urls": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"apex_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_transforms_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"graph_studio_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"machine_learning_notebook_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"machine_learning_user_management_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mongo_db_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ords_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_dev_web_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"cpu_core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"customer_contacts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"email": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"data_safe_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_storage_size_in_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"data_storage_size_in_tbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"database_edition": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_management_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dataguard_region_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_tools_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compute_count": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"max_idle_time_in_minutes": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"db_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_workload": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"disaster_recovery_region_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"failed_data_recovery_in_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"in_memory_area_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"in_memory_percentage": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"infrastructure_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_access_control_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_auto_scaling_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_auto_scaling_for_storage_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_data_guard_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_dedicated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_free_tier": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_local_data_guard_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_mtls_connection_required": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_preview": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_reconnect_clone_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_refreshable_clone": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_remote_data_guard_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key_history_entry": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"kms_key_version_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_activated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vault_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"key_store_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_store_wallet_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kms_key_lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kms_key_version_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"local_adg_auto_failover_max_data_loss_limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_disaster_recovery_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"local_standby_db": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"lag_time_in_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_data_guard_role_changed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_disaster_recovery_role_changed": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"memory_per_oracle_compute_unit_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"long_term_backup_schedule": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"is_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"repeat_cadence": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"retention_period_in_days": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"time_of_backup": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"max_cpu_core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ncharacter_set": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_long_term_backup_time_stamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nsg_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ocpu_count": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"open_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_db_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"private_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_endpoint_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_endpoint_label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"provisionable_cpus": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeFloat,
							},
						},
						"refreshable_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"refreshable_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_disaster_recovery_configuration": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"disaster_recovery_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"resource_pool_leader_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_pool_summary": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"is_disabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"pool_size": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scheduled_operations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"day_of_week": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"scheduled_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"scheduled_stop_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"service_console_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"standby_db": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"lag_time_in_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_data_guard_role_changed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_disaster_recovery_role_changed": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"standby_whitelisted_ips": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supported_regions_to_clone_to": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_data_guard_role_changed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_deletion_of_free_autonomous_database": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_disaster_recovery_role_changed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_local_data_guard_enabled": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_maintenance_begin": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_maintenance_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_joining_resource_pool": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_last_failover": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_last_refresh": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_last_refresh_point": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_last_switchover": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_next_refresh": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_reclamation_of_free_autonomous_database": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_until_reconnect_clone_enabled": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_backup_storage_size_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"used_data_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_data_storage_size_in_tbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"whitelisted_ips": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDatabasesClones(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabasesClonesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabasesClonesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabaseClonesResponse
}

func (s *DatabaseAutonomousDatabasesClonesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabasesClonesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabaseClonesRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if cloneType, ok := s.D.GetOkExists("clone_type"); ok {
		request.CloneType = oci_database.ListAutonomousDatabaseClonesCloneTypeEnum(cloneType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousDatabaseSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabaseClones(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabaseClones(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabasesClonesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabasesClonesDataSource-", DatabaseAutonomousDatabasesClonesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDatabasesClone := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ActualUsedDataStorageSizeInTBs != nil {
			autonomousDatabasesClone["actual_used_data_storage_size_in_tbs"] = *r.ActualUsedDataStorageSizeInTBs
		}

		if r.AllocatedStorageSizeInTBs != nil {
			autonomousDatabasesClone["allocated_storage_size_in_tbs"] = *r.AllocatedStorageSizeInTBs
		}

		if r.ApexDetails != nil {
			autonomousDatabasesClone["apex_details"] = []interface{}{AutonomousDatabaseApexToMap(r.ApexDetails)}
		} else {
			autonomousDatabasesClone["apex_details"] = nil
		}

		if r.ArePrimaryWhitelistedIpsUsed != nil {
			autonomousDatabasesClone["are_primary_whitelisted_ips_used"] = *r.ArePrimaryWhitelistedIpsUsed
		}

		if r.AutonomousContainerDatabaseId != nil {
			autonomousDatabasesClone["autonomous_container_database_id"] = *r.AutonomousContainerDatabaseId
		}

		autonomousDatabasesClone["autonomous_maintenance_schedule_type"] = r.AutonomousMaintenanceScheduleType

		autonomousDatabasesClone["available_upgrade_versions"] = r.AvailableUpgradeVersions

		if r.BackupConfig != nil {
			autonomousDatabasesClone["backup_config"] = []interface{}{AutonomousDatabaseBackupConfigToMap(r.BackupConfig)}
		} else {
			autonomousDatabasesClone["backup_config"] = nil
		}

		if r.BackupRetentionPeriodInDays != nil {
			autonomousDatabasesClone["backup_retention_period_in_days"] = *r.BackupRetentionPeriodInDays
		}

		if r.CharacterSet != nil {
			autonomousDatabasesClone["character_set"] = *r.CharacterSet
		}

		if r.ComputeCount != nil {
			autonomousDatabasesClone["compute_count"] = *r.ComputeCount
		}

		autonomousDatabasesClone["compute_model"] = r.ComputeModel

		if r.ConnectionStrings != nil {
			autonomousDatabasesClone["connection_strings"] = []interface{}{AutonomousDatabaseConnectionStringsToMap(r.ConnectionStrings)}
		} else {
			autonomousDatabasesClone["connection_strings"] = nil
		}

		if r.ConnectionUrls != nil {
			autonomousDatabasesClone["connection_urls"] = []interface{}{AutonomousDatabaseConnectionUrlsToMap(r.ConnectionUrls)}
		} else {
			autonomousDatabasesClone["connection_urls"] = nil
		}

		if r.CpuCoreCount != nil {
			autonomousDatabasesClone["cpu_core_count"] = *r.CpuCoreCount
		}

		customerContacts := []interface{}{}
		for _, item := range r.CustomerContacts {
			customerContacts = append(customerContacts, CustomerContactToMap(item))
		}
		autonomousDatabasesClone["customer_contacts"] = customerContacts

		autonomousDatabasesClone["data_safe_status"] = r.DataSafeStatus

		if r.DataStorageSizeInGBs != nil {
			autonomousDatabasesClone["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		if r.DataStorageSizeInTBs != nil {
			autonomousDatabasesClone["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		autonomousDatabasesClone["database_edition"] = r.DatabaseEdition

		autonomousDatabasesClone["database_management_status"] = r.DatabaseManagementStatus

		autonomousDatabasesClone["dataguard_region_type"] = r.DataguardRegionType

		if r.DbName != nil {
			autonomousDatabasesClone["db_name"] = *r.DbName
		}

		dbToolsDetails := []interface{}{}
		for _, item := range r.DbToolsDetails {
			dbToolsDetails = append(dbToolsDetails, DatabaseToolToMap(item))
		}

		autonomousDatabasesClone["db_tools_details"] = dbToolsDetails
		if r.DbVersion != nil {
			autonomousDatabasesClone["db_version"] = *r.DbVersion
		}

		autonomousDatabasesClone["db_workload"] = r.DbWorkload

		if r.DefinedTags != nil {
			autonomousDatabasesClone["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		autonomousDatabasesClone["disaster_recovery_region_type"] = r.DisasterRecoveryRegionType

		if r.DisplayName != nil {
			autonomousDatabasesClone["display_name"] = *r.DisplayName
		}

		if r.FailedDataRecoveryInSeconds != nil {
			autonomousDatabasesClone["failed_data_recovery_in_seconds"] = *r.FailedDataRecoveryInSeconds
		}

		autonomousDatabasesClone["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousDatabasesClone["id"] = *r.Id
		}

		if r.InMemoryAreaInGBs != nil {
			autonomousDatabasesClone["in_memory_area_in_gbs"] = *r.InMemoryAreaInGBs
		}

		if r.InMemoryPercentage != nil {
			autonomousDatabasesClone["in_memory_percentage"] = *r.InMemoryPercentage
		}

		autonomousDatabasesClone["infrastructure_type"] = r.InfrastructureType

		if r.IsAccessControlEnabled != nil {
			autonomousDatabasesClone["is_access_control_enabled"] = *r.IsAccessControlEnabled
		}

		if r.IsAutoScalingEnabled != nil {
			autonomousDatabasesClone["is_auto_scaling_enabled"] = *r.IsAutoScalingEnabled
		}

		if r.IsAutoScalingForStorageEnabled != nil {
			autonomousDatabasesClone["is_auto_scaling_for_storage_enabled"] = *r.IsAutoScalingForStorageEnabled
		}

		if r.IsDataGuardEnabled != nil {
			autonomousDatabasesClone["is_data_guard_enabled"] = *r.IsDataGuardEnabled
		}

		if r.IsDedicated != nil {
			autonomousDatabasesClone["is_dedicated"] = *r.IsDedicated
		}

		if r.IsFreeTier != nil {
			autonomousDatabasesClone["is_free_tier"] = *r.IsFreeTier
		}

		if r.IsLocalDataGuardEnabled != nil {
			autonomousDatabasesClone["is_local_data_guard_enabled"] = *r.IsLocalDataGuardEnabled
		}

		if r.IsMtlsConnectionRequired != nil {
			autonomousDatabasesClone["is_mtls_connection_required"] = *r.IsMtlsConnectionRequired
		}

		if r.IsPreview != nil {
			autonomousDatabasesClone["is_preview"] = *r.IsPreview
		}

		if r.IsReconnectCloneEnabled != nil {
			autonomousDatabasesClone["is_reconnect_clone_enabled"] = *r.IsReconnectCloneEnabled
		}

		if r.IsRefreshableClone != nil {
			autonomousDatabasesClone["is_refreshable_clone"] = *r.IsRefreshableClone
		}

		if r.IsRemoteDataGuardEnabled != nil {
			autonomousDatabasesClone["is_remote_data_guard_enabled"] = *r.IsRemoteDataGuardEnabled
		}

		keyHistoryEntry := []interface{}{}
		for _, item := range r.KeyHistoryEntry {
			keyHistoryEntry = append(keyHistoryEntry, AutonomousDatabaseKeyHistoryEntryToMap(item))
		}
		autonomousDatabasesClone["key_history_entry"] = keyHistoryEntry

		if r.KeyStoreId != nil {
			autonomousDatabasesClone["key_store_id"] = *r.KeyStoreId
		}

		if r.KeyStoreWalletName != nil {
			autonomousDatabasesClone["key_store_wallet_name"] = *r.KeyStoreWalletName
		}

		if r.KmsKeyId != nil {
			autonomousDatabasesClone["kms_key_id"] = *r.KmsKeyId
		}

		if r.KmsKeyLifecycleDetails != nil {
			autonomousDatabasesClone["kms_key_lifecycle_details"] = *r.KmsKeyLifecycleDetails
		}

		if r.KmsKeyVersionId != nil {
			autonomousDatabasesClone["kms_key_version_id"] = *r.KmsKeyVersionId
		}

		autonomousDatabasesClone["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			autonomousDatabasesClone["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.LocalAdgAutoFailoverMaxDataLossLimit != nil {
			autonomousDatabasesClone["local_adg_auto_failover_max_data_loss_limit"] = *r.LocalAdgAutoFailoverMaxDataLossLimit
		}

		autonomousDatabasesClone["local_disaster_recovery_type"] = r.LocalDisasterRecoveryType

		if r.LocalStandbyDb != nil {
			autonomousDatabasesClone["local_standby_db"] = []interface{}{AutonomousDatabaseStandbySummaryToMap(r.LocalStandbyDb)}
		} else {
			autonomousDatabasesClone["local_standby_db"] = nil
		}

		if r.LongTermBackupSchedule != nil {
			autonomousDatabasesClone["long_term_backup_schedule"] = []interface{}{LongTermBackUpScheduleDetailsToMap(r.LongTermBackupSchedule)}
		} else {
			autonomousDatabasesClone["long_term_backup_schedule"] = nil
		}

		if r.MaxCpuCoreCount != nil {
			autonomousDatabasesClone["max_cpu_core_count"] = *r.MaxCpuCoreCount
		}

		if r.MemoryPerOracleComputeUnitInGBs != nil {
			autonomousDatabasesClone["memory_per_oracle_compute_unit_in_gbs"] = *r.MemoryPerOracleComputeUnitInGBs
		}

		if r.MaxCpuCoreCount != nil {
			autonomousDatabasesClone["max_cpu_core_count"] = *r.MaxCpuCoreCount
		}

		if r.NcharacterSet != nil {
			autonomousDatabasesClone["ncharacter_set"] = *r.NcharacterSet
		}

		if r.NextLongTermBackupTimeStamp != nil {
			autonomousDatabasesClone["next_long_term_backup_time_stamp"] = r.NextLongTermBackupTimeStamp.String()
		}

		autonomousDatabasesClone["nsg_ids"] = r.NsgIds

		if r.OcpuCount != nil {
			autonomousDatabasesClone["ocpu_count"] = *r.OcpuCount
		}

		autonomousDatabasesClone["open_mode"] = r.OpenMode

		autonomousDatabasesClone["peer_db_ids"] = r.PeerDbIds

		if r.PrivateEndpoint != nil {
			autonomousDatabasesClone["private_endpoint"] = *r.PrivateEndpoint
		}

		if r.PrivateEndpointIp != nil {
			autonomousDatabasesClone["private_endpoint_ip"] = *r.PrivateEndpointIp
		}

		if r.PrivateEndpointLabel != nil {
			autonomousDatabasesClone["private_endpoint_label"] = *r.PrivateEndpointLabel
		}

		autonomousDatabasesClone["provisionable_cpus"] = r.ProvisionableCpus

		autonomousDatabasesClone["refreshable_mode"] = r.RefreshableMode

		autonomousDatabasesClone["refreshable_status"] = r.RefreshableStatus

		if r.RemoteDisasterRecoveryConfiguration != nil {
			autonomousDatabasesClone["remote_disaster_recovery_configuration"] = []interface{}{DisasterRecoveryConfigurationToMap(r.RemoteDisasterRecoveryConfiguration)}
		} else {
			autonomousDatabasesClone["remote_disaster_recovery_configuration"] = nil
		}

		if r.ResourcePoolLeaderId != nil {
			autonomousDatabasesClone["resource_pool_leader_id"] = *r.ResourcePoolLeaderId
		}

		if r.ResourcePoolSummary != nil {
			autonomousDatabasesClone["resource_pool_summary"] = []interface{}{ResourcePoolSummaryToMap(r.ResourcePoolSummary)}
		} else {
			autonomousDatabasesClone["resource_pool_summary"] = nil
		}

		autonomousDatabasesClone["role"] = r.Role

		scheduledOperations := []interface{}{}
		for _, item := range r.ScheduledOperations {
			scheduledOperations = append(scheduledOperations, ScheduledOperationDetailsToMap(item))
		}
		autonomousDatabasesClone["scheduled_operations"] = scheduledOperations

		if r.ServiceConsoleUrl != nil {
			autonomousDatabasesClone["service_console_url"] = *r.ServiceConsoleUrl
		}

		if r.SourceId != nil {
			autonomousDatabasesClone["source_id"] = *r.SourceId
		}

		if r.StandbyDb != nil {
			autonomousDatabasesClone["standby_db"] = []interface{}{AutonomousDatabaseStandbySummaryToMap(r.StandbyDb)}
		} else {
			autonomousDatabasesClone["standby_db"] = nil
		}

		autonomousDatabasesClone["standby_whitelisted_ips"] = r.StandbyWhitelistedIps

		autonomousDatabasesClone["state"] = r.LifecycleState

		if r.SubnetId != nil {
			autonomousDatabasesClone["subnet_id"] = *r.SubnetId
		}

		autonomousDatabasesClone["supported_regions_to_clone_to"] = r.SupportedRegionsToCloneTo

		if r.SystemTags != nil {
			autonomousDatabasesClone["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			autonomousDatabasesClone["time_created"] = r.TimeCreated.String()
		}

		if r.TimeDataGuardRoleChanged != nil {
			autonomousDatabasesClone["time_data_guard_role_changed"] = r.TimeDataGuardRoleChanged.String()
		}

		if r.TimeDeletionOfFreeAutonomousDatabase != nil {
			autonomousDatabasesClone["time_deletion_of_free_autonomous_database"] = r.TimeDeletionOfFreeAutonomousDatabase.String()
		}

		if r.TimeDisasterRecoveryRoleChanged != nil {
			autonomousDatabasesClone["time_disaster_recovery_role_changed"] = r.TimeDisasterRecoveryRoleChanged.String()
		}

		if r.TimeLocalDataGuardEnabled != nil {
			autonomousDatabasesClone["time_local_data_guard_enabled"] = r.TimeLocalDataGuardEnabled.String()
		}

		if r.TimeMaintenanceBegin != nil {
			autonomousDatabasesClone["time_maintenance_begin"] = r.TimeMaintenanceBegin.String()
		}

		if r.TimeMaintenanceEnd != nil {
			autonomousDatabasesClone["time_maintenance_end"] = r.TimeMaintenanceEnd.String()
		}

		if r.TimeOfJoiningResourcePool != nil {
			autonomousDatabasesClone["time_of_joining_resource_pool"] = r.TimeOfJoiningResourcePool.String()
		}

		if r.TimeOfLastFailover != nil {
			autonomousDatabasesClone["time_of_last_failover"] = r.TimeOfLastFailover.String()
		}

		if r.TimeOfLastRefresh != nil {
			autonomousDatabasesClone["time_of_last_refresh"] = r.TimeOfLastRefresh.String()
		}

		if r.TimeOfLastRefreshPoint != nil {
			autonomousDatabasesClone["time_of_last_refresh_point"] = r.TimeOfLastRefreshPoint.String()
		}

		if r.TimeOfLastSwitchover != nil {
			autonomousDatabasesClone["time_of_last_switchover"] = r.TimeOfLastSwitchover.String()
		}

		if r.TimeOfNextRefresh != nil {
			autonomousDatabasesClone["time_of_next_refresh"] = r.TimeOfNextRefresh.String()
		}

		if r.TimeReclamationOfFreeAutonomousDatabase != nil {
			autonomousDatabasesClone["time_reclamation_of_free_autonomous_database"] = r.TimeReclamationOfFreeAutonomousDatabase.String()
		}

		if r.TimeUntilReconnectCloneEnabled != nil {
			autonomousDatabasesClone["time_until_reconnect_clone_enabled"] = r.TimeUntilReconnectCloneEnabled.String()
		}

		if r.TotalBackupStorageSizeInGBs != nil {
			autonomousDatabasesClone["total_backup_storage_size_in_gbs"] = *r.TotalBackupStorageSizeInGBs
		}

		if r.UsedDataStorageSizeInGBs != nil {
			autonomousDatabasesClone["used_data_storage_size_in_gbs"] = *r.UsedDataStorageSizeInGBs
		}

		if r.UsedDataStorageSizeInTBs != nil {
			autonomousDatabasesClone["used_data_storage_size_in_tbs"] = *r.UsedDataStorageSizeInTBs
		}

		if r.VaultId != nil {
			autonomousDatabasesClone["vault_id"] = *r.VaultId
		}

		autonomousDatabasesClone["whitelisted_ips"] = r.WhitelistedIps

		resources = append(resources, autonomousDatabasesClone)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDatabasesClonesDataSource().Schema["autonomous_databases"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_databases", resources); err != nil {
		return err
	}

	return nil
}

func LongTermBackUpScheduleDetailsToMap(obj *oci_database.LongTermBackUpScheduleDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDisabled != nil {
		result["is_disabled"] = bool(*obj.IsDisabled)
	}

	result["repeat_cadence"] = string(obj.RepeatCadence)

	if obj.RetentionPeriodInDays != nil {
		result["retention_period_in_days"] = int(*obj.RetentionPeriodInDays)
	}

	if obj.TimeOfBackup != nil {
		result["time_of_backup"] = obj.TimeOfBackup.String()
	}

	return result
}
